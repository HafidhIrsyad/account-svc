package server

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/hafidhirsyad/account-svc/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func Set(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(logger.LogMiddleware)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}

func SetConfig(dirpath string, filename string) {
	filePath := filepath.Join(dirpath, filename)
	fileExist := isFileExist(filePath)

	if fileExist {
		viper.AddConfigPath(dirpath)
		viper.SetConfigFile(filePath)

		if err := viper.ReadInConfig(); err != nil {
			log.Fatal().Err(err).Msgf("error reading config file: %+v", err)
		}
	} else {
		viper.AutomaticEnv()
	}

	SecretConfig()
}

// isFileExist check if the file exist on the given file path
func isFileExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return false
}
