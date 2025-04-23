package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func startServerWithGracefulShutdown(e *echo.Echo) {
	addr := fmt.Sprintf(":%d", AppPort)
	server := &http.Server{
		Addr:    addr,
		Handler: e,
	}

	// Create server context
	serverCtx, cancelServerCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig

		graceful_timeout := 10

		shutDownCtx, cancel := context.WithTimeout(serverCtx, time.Duration(graceful_timeout)*time.Second)
		defer cancel()

		go func() {
			<-shutDownCtx.Done()
			if shutDownCtx.Err() == context.DeadlineExceeded {
				log.Fatal().Err(shutDownCtx.Err()).Msg("graceful shutdown timed out, forcing exit..")
			}
		}()

		err := server.Shutdown(shutDownCtx)
		if err != nil {
			log.Fatal().Err(err).Msgf("error on shutting down gracefully: %v", err)
		}

		cancelServerCtx()
	}()

	log.Info().Msgf("starting account-svc api in port: %d", AppPort)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msgf("error on starting up server: %v", err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()

	log.Info().Msg("server is shut down!")
}
