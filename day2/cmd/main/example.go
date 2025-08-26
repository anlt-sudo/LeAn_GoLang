package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger.Info().Msg("starting app")
	fmt.Println("Hello, tidy + zerolog!")
}
