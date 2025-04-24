package router

import (
	"github.com/hafidhirsyad/account-svc/api/handler"
	"github.com/labstack/echo/v4"
)

func AccountsPath(e *echo.Echo, ah handler.AccountHandler) {
	e.POST("/account/register", ah.Register)
}
