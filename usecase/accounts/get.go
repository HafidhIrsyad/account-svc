package accounts

import (
	"context"
	"errors"
	"time"

	"github.com/hafidhirsyad/account-svc/logger"
	repo "github.com/hafidhirsyad/account-svc/repository/accounts"
	"github.com/rs/zerolog"
)

func (a *AccountService) GetBalanceByNoRekening(ctx context.Context, norek int64) (saldo int64, err error) {
	defer TimeTrack(time.Now(), "AccountService.GetBalanceByNoRekening")

	resp, err := a.accRepo.GetBalanceByFilter(ctx, repo.Filter{
		NoRekening: norek,
	})
	if err != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "error when get balance by filter", map[string]any{"error": err.Error(), "request": norek, "func": "GetBalanceByNoRekening", "path": "usecase.accounts.get"})
		return 0, err
	}

	if resp.AccountId == 0 {
		msg := "Nomor rekening tidak ditemukan, pastikan memasukkan nomor rekening yang benar"
		logger.Log(ctx, zerolog.ErrorLevel, msg, map[string]any{"request": norek, "func": "GetBalanceByNoRekening", "path": "usecase.accounts.get"})
		return 0, errors.New(msg)
	}

	return resp.Nominal, nil
}
