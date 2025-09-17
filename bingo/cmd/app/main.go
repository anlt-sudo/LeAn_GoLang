package main

import (
	"github.com/anlt-sudo/bingo/internal/handler"
	"github.com/rs/zerolog/log"
)

func main() {

	bingoHandler, err := handler.NewBingoHandler()
	if err != nil {
		log.Error().Err(err).Msg("Error while init service")
		return
	}
	defer bingoHandler.Close()

	err = bingoHandler.RunGame()
	if err != nil {
		log.Error().Err(err).Msg("Loi khi chay Bingo")
		return
	}

	log.Info().Msg("Da Ket Thuc")
}
