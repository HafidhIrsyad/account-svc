package handler

import (
	"net/http"

	"github.com/hafidhirsyad/account-svc/entity"
	"github.com/hafidhirsyad/account-svc/logger"
	"github.com/hafidhirsyad/account-svc/usecase/accounts"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type AccountHandler struct {
	accService accounts.AccountServiceI
}

func NewAccountHandler(accService accounts.AccountServiceI) *AccountHandler {
	return &AccountHandler{accService: accService}
}

func (h *AccountHandler) Register(c echo.Context) error {
	ctx := c.Request().Context()

	logger.Log(ctx, zerolog.InfoLevel, "register api called", map[string]any{"func": "Register", "path": "api.handler.accounts"})

	var req entity.RegisterReq
	if err := c.Bind(&req); err != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "invalid request", map[string]any{"func": "Register", "path": "api.handler.accounts"})
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "invalid request"})
	}

	err := req.ValidateRegister()
	if err != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "error validate request", map[string]any{"error": err.Error(), "func": "Register", "path": "api.handler.accounts"})
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	noRek, err := h.accService.Register(ctx, req)
	if err != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "error register", map[string]any{"error": err.Error(), "func": "Register", "path": "api.handler.accounts", "request": req})
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	logger.Log(ctx, zerolog.InfoLevel, "register api called success", map[string]any{"func": "Register", "path": "api.handler.accounts"})

	return c.JSON(http.StatusOK, map[string]int64{"no_rekening": noRek})
}
