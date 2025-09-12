package main

import (
	"github.com/anlt-sudo/bingo/internal/handler"
	"github.com/rs/zerolog/log"
)

func main() {

	bingoService, err := handler.NewBingoHandler()
	if err != nil {
		log.Error().Err(err).Msg("Error while init service")
		return
	}
	defer bingoService.Close()

	err = bingoService.RunGame()
	if err != nil {
		log.Error().Err(err).Msg("Loi khi chay Bingo")
		return
	}

	log.Info().Msg("Da Ket Thuc")
}
