package accounts

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hafidhirsyad/account-svc/entity"
	"github.com/hafidhirsyad/account-svc/logger"
	repo "github.com/hafidhirsyad/account-svc/repository/accounts"
	"github.com/rs/zerolog"
)

// Withdraw implements AccountServiceI.
func (a *AccountService) Withdraw(ctx context.Context, req entity.WithdrawReq) (saldo int64, err error) {
	defer TimeTrack(time.Now(), "AccountService.Withdraw")

	norekInt := ParseStrToInt64(req.NoRekening)
	now := time.Now()
	saldo = req.Nominal

	resp, err := a.accRepo.GetBalanceByFilter(ctx, repo.Filter{
		NoRekening: int64(norekInt),
	})
	if err != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "error when get balance by filter", map[string]any{"error": err.Error(), "request": req, "func": "Withdraw", "path": "usecase.accounts.withdraw"})
		return 0, err
	}

	if resp.WalletId == 0 {
		msg := "Anda tidak memiliki saldo"
		logger.Log(ctx, zerolog.ErrorLevel, msg, map[string]any{"request": req, "func": "Withdraw", "path": "usecase.accounts.withdraw"})
		return 0, errors.New(msg)
	}

	if resp.Nominal < req.Nominal {
		msg := fmt.Sprintf("Saldo anda tidak mencukup untuk penarikan, saldo anda saat ini %d", resp.Nominal)
		logger.Log(ctx, zerolog.ErrorLevel, msg, map[string]any{"request": req, "func": "Withdraw", "path": "usecase.accounts.withdraw"})
		return 0, errors.New(msg)
	}

	trx := a.trxRepo.BeginTransaction(ctx)

	saldo = resp.Nominal - req.Nominal
	err = a.accRepo.UpdateBalance(ctx, trx, resp.AccountId, map[string]any{
		"nominal":    saldo,
		"updated_at": &now,
	})
	if err != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "error when update balance", map[string]any{"error": err.Error(), "request": req, "func": "Withdraw", "path": "usecase.accounts.withdraw"})
		a.trxRepo.RollBackTransaction(ctx, trx)
		return 0, err
	}

	a.trxRepo.CommitTransaction(ctx, trx)

	return saldo, nil
}
