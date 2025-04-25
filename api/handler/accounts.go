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

func (h *AccountHandler) Deposit(c echo.Context) error {
	ctx := c.Request().Context()

	logger.Log(ctx, zerolog.InfoLevel, "deposit api called", map[string]any{"func": "Deposit", "path": "api.handler.accounts"})

	var req entity.DepositReq
	if err := c.Bind(&req); err != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "invalid request", map[string]any{"func": "Deposit", "path": "api.handler.accounts"})
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "invalid request"})
	}

	err := req.ValidateDeposit()
	if err != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "error validate request", map[string]any{"error": err.Error(), "func": "Deposit", "path": "api.handler.accounts"})
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	saldo, err := h.accService.Deposit(ctx, req)
	if err != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "error Deposit", map[string]any{"error": err.Error(), "func": "Deposit", "path": "api.handler.accounts", "request": req})
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	logger.Log(ctx, zerolog.InfoLevel, "deposit api called success", map[string]any{"func": "Deposit", "path": "api.handler.accounts"})

	return c.JSON(http.StatusOK, map[string]int64{"saldo": saldo})
}

func (h *AccountHandler) Withdraw(c echo.Context) error {
	ctx := c.Request().Context()

	logger.Log(ctx, zerolog.InfoLevel, "Withdraw api called", map[string]any{"func": "Withdraw", "path": "api.handler.accounts"})

	var req entity.WithdrawReq
	if err := c.Bind(&req); err != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "invalid request", map[string]any{"func": "Withdraw", "path": "api.handler.accounts"})
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": "invalid request"})
	}

	err := req.ValidateWithdraw()
	if err != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "error validate request", map[string]any{"error": err.Error(), "func": "Withdraw", "path": "api.handler.accounts"})
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	saldo, err := h.accService.Withdraw(ctx, req)
	if err != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "error Withdraw", map[string]any{"error": err.Error(), "func": "Withdraw", "path": "api.handler.accounts", "request": req})
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	logger.Log(ctx, zerolog.InfoLevel, "Withdraw api called success", map[string]any{"func": "Withdraw", "path": "api.handler.accounts"})

	return c.JSON(http.StatusOK, map[string]int64{"saldo": saldo})
}

func (h *AccountHandler) GetBalanceByNoRekening(c echo.Context) error {
	ctx := c.Request().Context()

	logger.Log(ctx, zerolog.InfoLevel, "get balance api called", map[string]any{"func": "GetBalanceByNoRekening", "path": "api.handler.accounts"})

	rekeningStr := c.Param("no_rekening")
	norekInt := accounts.ParseStrToInt64(rekeningStr)
	if norekInt == 0 {
		msg := "Nomor Rekening tidak boleh kosong"
		logger.Log(ctx, zerolog.ErrorLevel, msg, map[string]any{"func": "GetBalanceByNoRekening", "path": "api.handler.accounts", "request": norekInt})
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": msg})
	}

	saldo, err := h.accService.GetBalanceByNoRekening(ctx, norekInt)
	if err != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "error get balance", map[string]any{"error": err.Error(), "func": "GetBalanceByNoRekening", "path": "api.handler.accounts", "request": norekInt})
		return c.JSON(http.StatusBadRequest, map[string]string{"remark": err.Error()})
	}

	logger.Log(ctx, zerolog.InfoLevel, "get balance api called success", map[string]any{"func": "GetBalanceByNoRekening", "path": "api.handler.accounts"})

	return c.JSON(http.StatusOK, map[string]int64{"saldo": saldo})
}
