package router

import (
	"github.com/hafidhirsyad/account-svc/api/handler"
	"github.com/labstack/echo/v4"
)

func AccountsPath(e *echo.Echo, ah handler.AccountHandler) {
	e.POST("/account/daftar", ah.Register)
	e.POST("/account/tabung", ah.Deposit)
	e.POST("/account/tarik", ah.Withdraw)
	e.GET("/account/saldo/:no_rekening", ah.GetBalanceByNoRekening)
}
