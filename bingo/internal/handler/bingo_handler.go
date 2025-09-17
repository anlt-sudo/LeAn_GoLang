package handler

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/anlt-sudo/bingo/internal/model"
	"github.com/anlt-sudo/bingo/internal/service"
	"github.com/rs/zerolog/log"
)

type BingoHandler struct {
    called       map[string]bool
    calledList   []string
    fileService  *service.FileService
    bingoService *service.BingoService
    eventChan    chan model.GameEvent   
    consoleChan  chan model.GameEvent   
    fileChan     chan model.GameEvent   
}

func NewBingoHandler() (*BingoHandler, error) {
	fileService, err := service.NewFileService("bingo.txt")
	bingoService := service.NewBingoBoard()
	eventChan := make(chan model.GameEvent, 10)
	consoleChan := make(chan model.GameEvent, 100)
	fileChan := make(chan model.GameEvent, 100)


	if err != nil {
		log.Error().Err(err).Msg("Lỗi tạo file handler")
		return nil, fmt.Errorf("lỗi tạo file handler: %w", err)
	}


	return &BingoHandler{
		called:       make(map[string]bool),
		calledList:   []string{},
		fileService:  fileService,
		bingoService: bingoService,
		eventChan:    eventChan,
		consoleChan:  consoleChan,
		fileChan:     fileChan,
	}, nil
}

func (bs *BingoHandler) Close() error {
	return bs.fileService.Close()
}

func (bs *BingoHandler) RunGame() error {
    var wg sync.WaitGroup
    wg.Add(3)

    go func() { defer wg.Done(); bs.fanOut() }()
    go func() { defer wg.Done(); bs.consoleWriter() }()
    go func() { defer wg.Done(); bs.fileWriter() }()

    bs.eventChan <- model.GameEvent{EventType: "start", Board: bs.bingoService.Board}

    var bingoMsg string
    var bingoPos [][2]int

    for {
        calledNumber, ok := bs.callNextNumber()
        if !ok {
            break
        }

        bs.bingoService.Used[calledNumber] = true

        bs.eventChan <- model.GameEvent{
            EventType: "called",
            Message:   calledNumber,
            Board:     bs.bingoService.Board,
        }

        hasBingo, msg, pos := bs.bingoService.CheckBingo()
        if hasBingo {
            bingoMsg = msg
            bingoPos = pos
            break
        }

        time.Sleep(2 * time.Second)
    }

    bs.eventChan <- model.GameEvent{
        EventType: "bingo",
        Message:   bingoMsg,
        Board:     bs.bingoService.Board,
        BingoPos:  bingoPos,
    }

    bs.eventChan <- model.GameEvent{EventType: "final"}

    close(bs.eventChan)
    wg.Wait()
    return nil
}



func (bs *BingoHandler) callNextNumber() (string, bool) {
	for {
		num := rand.Intn(50) + 1
		str := fmt.Sprintf("%d", num)
		if bs.called[str] {
			continue
		}
		bs.called[str] = true
		bs.calledList = append(bs.calledList, str)
		return str, true
	}
}

func (bs *BingoHandler) consoleWriter() {
    for event := range bs.consoleChan {
        switch event.EventType {
        case "start":
            fmt.Println("Bắt đầu game Bingo!")
            printBoard(event.Board, nil)
        case "called":
            fmt.Println("Số được gọi:", event.Message)
            printBoard(event.Board, bs.bingoService.Used)
        case "bingo":
            fmt.Println("BINGO!!!", event.Message)
            printBoard(event.Board, bs.bingoService.Used)
            if len(event.BingoPos) > 0 {
                fmt.Println("Vị trí Bingo:", event.BingoPos)
            }
        case "final":
            fmt.Println("Kết thúc game sau", len(bs.calledList), "lần gọi")
            fmt.Println("Các số đã gọi:", bs.calledList)
        }
    }
}

func (bs *BingoHandler) fileWriter() {
    for event := range bs.fileChan {
        switch event.EventType {
        case "start":
            bs.fileService.WriteBoard(event.Board)
        case "called":
            bs.fileService.WriteCalledNumber(event.Message)
        case "bingo":
            bs.fileService.WriteBingoResult(event.Message)
            bs.fileService.WriteFinalBoard(event.Board, bs.called, event.BingoPos)
        case "final":
            bs.fileService.WriteNewline()
        }
    }
}


func printBoard(board [][]string, used map[string]bool) {
    fmt.Println("----- BOARD -----")
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            val := board[i][j]
            if used != nil && used[val] {
                fmt.Printf("[%2s] ", val)
            } else {
                fmt.Printf(" %2s  ", val)
            }
        }
        fmt.Println()
    }
    fmt.Println("-----------------")
}

func (bs *BingoHandler) fanOut() {
    for event := range bs.eventChan {
        bs.consoleChan <- event
        bs.fileChan <- event
    }
    close(bs.consoleChan)
    close(bs.fileChan)
}


