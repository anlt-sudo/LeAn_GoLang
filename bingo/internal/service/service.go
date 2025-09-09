package service

import (
	"fmt"
	"github.com/anlt-sudo/bingo/internal/bingo"
	"github.com/anlt-sudo/bingo/internal/file"
	"github.com/rs/zerolog/log"
	"math/rand"
	"time"
)

type BingoService struct {
	board       [][]string
	called      map[string]bool
	calledList  []string
	fileHandler *file.Handler
}

func NewBingoService() (*BingoService, error) {
	fileHandler, err := file.NewFileHandler("bingo_output.txt")
	if err != nil {
		log.Error().Err(err).Msg("Lỗi tạo file handler")
		return nil, fmt.Errorf("lỗi tạo file handler: %w", err)
	}

	board := bingo.CreateBingoBoard()

	return &BingoService{
		board:       board,
		called:      make(map[string]bool),
		calledList:  []string{},
		fileHandler: fileHandler,
	}, nil
}

func (bs *BingoService) Close() error {
	return bs.fileHandler.Close()
}

func (bs *BingoService) RunGame() error {
	err := bs.fileHandler.WriteBoard(bs.board)
	if err != nil {
		log.Error().Err(err).Msg("Lỗi ghi bảng")
		return fmt.Errorf("lỗi ghi bảng: %w", err)
	}

	var bingoMsg string
	var bingoPos [][2]int

	for {
		calledNumber, ok := bs.callNextNumber()
		if !ok {
			break
		}

		if len(bs.calledList) == 1 {
			err = bs.fileHandler.WriteNewline()
			if err != nil {
				log.Error().Err(err).Msg("Lỗi ghi xuống dòng")
				return fmt.Errorf("lỗi ghi xuống dòng: %w", err)
			}
		}

		err = bs.fileHandler.WriteCalledNumber(calledNumber)
		if err != nil {
			log.Error().Err(err).Msg("Lỗi ghi số")
			return fmt.Errorf("lỗi ghi số: %w", err)
		}

		hasBingo, msg, pos := bingo.CheckBingo(bs.board, bs.called)
		if hasBingo {
			bingoMsg = msg
			bingoPos = pos
			break
		}

		time.Sleep(2 * time.Second)
	}

	err = bs.fileHandler.WriteBingoResult(bingoMsg)
	if err != nil {
		log.Error().Err(err).Msg("Lỗi ghi kết quả")
		return fmt.Errorf("lỗi ghi kết quả: %w", err)
	}

	err = bs.fileHandler.WriteFinalBoard(bs.board, bs.called, bingoPos)
	if err != nil {
		log.Error().Err(err).Msg("Lỗi ghi bảng cuối")
		return fmt.Errorf("lỗi ghi bảng cuối: %w", err)
	}

	return nil
}

func (bs *BingoService) callNextNumber() (string, bool) {
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
