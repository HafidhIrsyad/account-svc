package server

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func Start() {
	SetConfig(".", ".env")

	_, err := DBConnection()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	e := echo.New()
	Set(e)

	startServerWithGracefulShutdown(e)
}
