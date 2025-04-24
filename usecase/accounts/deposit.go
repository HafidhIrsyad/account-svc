package accounts

import (
	"context"
	"errors"
	"time"

	"github.com/hafidhirsyad/account-svc/entity"
	"github.com/hafidhirsyad/account-svc/logger"
	repo "github.com/hafidhirsyad/account-svc/repository/accounts"
	"github.com/rs/zerolog"
)

func (a *AccountService) Deposit(ctx context.Context, req entity.DepositReq) (saldo int64, err error) {
	defer TimeTrack(time.Now(), "AccountService.Deposit")

	norekInt := ParseStrToInt64(req.NoRekening)
	now := time.Now()
	saldo = req.Nominal

	resp, err := a.accRepo.GetBalanceByFilter(ctx, repo.Filter{
		NoRekening: int64(norekInt),
	})
	if err != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "error when get balance by filter", map[string]any{"error": err.Error(), "request": req, "func": "Deposit", "path": "usecase.accounts.deposit"})
		return 0, err
	}

	if resp.AccountId == 0 {
		msg := "Nomor rekening tidak ditemukan, pastikan memasukkan nomor rekening yang benar"
		logger.Log(ctx, zerolog.ErrorLevel, msg, map[string]any{"request": req, "func": "Deposit", "path": "usecase.accounts.deposit"})
		return 0, errors.New(msg)
	}

	trx := a.trxRepo.BeginTransaction(ctx)

	if resp.WalletId != 0 {
		saldo = resp.Nominal + req.Nominal
		err = a.accRepo.UpdateBalance(ctx, trx, resp.AccountId, map[string]any{
			"nominal":    saldo,
			"updated_at": &now,
		})
		if err != nil {
			logger.Log(ctx, zerolog.ErrorLevel, "error when update balance", map[string]any{"error": err.Error(), "request": req, "func": "Deposit", "path": "usecase.accounts.deposit"})
			a.trxRepo.RollBackTransaction(ctx, trx)
			return 0, err
		}
	} else {
		_, err = a.accRepo.InsertBalance(ctx, trx, repo.Deposit{
			AccountId: resp.AccountId,
			Nominal:   req.Nominal,
		})
		if err != nil {
			logger.Log(ctx, zerolog.ErrorLevel, "error when insert balance", map[string]any{"error": err.Error(), "request": req, "func": "Deposit", "path": "usecase.accounts.deposit"})
			a.trxRepo.RollBackTransaction(ctx, trx)
			return 0, err
		}
	}

	a.trxRepo.CommitTransaction(ctx, trx)

	return saldo, nil
}
