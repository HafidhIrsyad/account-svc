package server

import (
	"github.com/hafidhirsyad/account-svc/api/handler"
	"github.com/hafidhirsyad/account-svc/api/router"
	accRepo "github.com/hafidhirsyad/account-svc/repository/accounts"
	"github.com/hafidhirsyad/account-svc/repository/transaction"
	accService "github.com/hafidhirsyad/account-svc/usecase/accounts"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func Start() {
	SetConfig(".", ".env")

	db, err := DBConnection()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	accountsRepo := accRepo.NewAccountRepository(db)
	trxRepository := transaction.NewTransactionRepository(db)
	accountSvc := accService.NewAccountService(accountsRepo, trxRepository)
	accountHandler := handler.NewAccountHandler(accountSvc)

	e := echo.New()
	Set(e)

	router.AccountsPath(e, *accountHandler)

	startServerWithGracefulShutdown(e)
}
