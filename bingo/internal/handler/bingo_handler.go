package handler

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/anlt-sudo/bingo/internal/service"
	"github.com/rs/zerolog/log"
)

type BingoHandler struct {
	called      map[string]bool
	calledList  []string
	fileService *service.FileService
	bingoService *service.BingoService
}

func NewBingoHandler() (*BingoHandler, error) {
	fileService, err := service.NewFileService("bingo.txt")
	bingoService := service.NewBingoBoard()

	if err != nil {
		log.Error().Err(err).Msg("Lỗi tạo file handler")
		return nil, fmt.Errorf("lỗi tạo file handler: %w", err)
	}


	return &BingoHandler{
		called:      make(map[string]bool),
		calledList:  []string{},
		fileService: fileService,
		bingoService: bingoService,
	}, nil
}

func (bs *BingoHandler) Close() error {
	return bs.fileService.Close()
}

func (bs *BingoHandler) RunGame() error {
	err := bs.fileService.WriteBoard(bs.bingoService.Board)
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

		bs.bingoService.Used[calledNumber] = true

		if len(bs.calledList) == 1 {
			err = bs.fileService.WriteNewline()
			if err != nil {
				log.Error().Err(err).Msg("Lỗi ghi xuống dòng")
				return fmt.Errorf("lỗi ghi xuống dòng: %w", err)
			}
		}

		err = bs.fileService.WriteCalledNumber(calledNumber)
		if err != nil {
			log.Error().Err(err).Msg("Lỗi ghi số")
			return fmt.Errorf("lỗi ghi số: %w", err)
		}

		hasBingo, msg, pos := bs.bingoService.CheckBingo()
		if hasBingo {
			bingoMsg = msg
			bingoPos = pos
			break
		}

		time.Sleep(2 * time.Second)
	}

	err = bs.fileService.WriteBingoResult(bingoMsg)
	if err != nil {
		log.Error().Err(err).Msg("Lỗi ghi kết quả")
		return fmt.Errorf("lỗi ghi kết quả: %w", err)
	}

	err = bs.fileService.WriteFinalBoard(bs.bingoService.Board, bs.called, bingoPos)
	if err != nil {
		log.Error().Err(err).Msg("Lỗi ghi bảng cuối")
		return fmt.Errorf("lỗi ghi bảng cuối: %w", err)
	}

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
