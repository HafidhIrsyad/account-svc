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

func (a *AccountService) generateNoRekening() int64 {
	return int64(1000000000 + a.rand.Intn(900000000))
}

func (a *AccountService) Register(ctx context.Context, req entity.RegisterReq) (norek int64, err error) {
	defer TimeTrack(time.Now(), "AccountService.Register")

	noRekening := a.generateNoRekening()
	noHpInt := ParseStrToInt64(req.NoHP)
	nikInt := ParseStrToInt64(req.NIK)

	resp, err := a.accRepo.GetBalanceByFilter(ctx, repo.Filter{
		NIK:  int64(nikInt),
		NoHP: int64(noHpInt),
	})
	if err != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "error when get balance by filter", map[string]any{"error": err.Error(), "request": req, "func": "Register", "path": "usecase.accounts.register"})
		return 0, err
	}

	if resp.NoHP != 0 || resp.NIK != 0 {
		msg := "No HP Atau NIK sudah terdaftar"
		logger.Log(ctx, zerolog.ErrorLevel, msg, map[string]any{"request": req, "func": "Register", "path": "usecase.accounts.register"})
		return 0, errors.New(msg)
	}

	trx := a.trxRepo.BeginTransaction(ctx)

	_, err = a.accRepo.Register(ctx, trx, repo.Register{
		Name:       req.Name,
		NIK:        int64(nikInt),
		NoHP:       int64(noHpInt),
		NoRekening: noRekening,
	})
	if err != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "error when register", map[string]any{"error": err.Error(), "request": req, "func": "Register", "path": "usecase.accounts.register"})
		a.trxRepo.RollBackTransaction(ctx, trx)
		return 0, err
	}

	a.trxRepo.CommitTransaction(ctx, trx)

	return noRekening, nil
}
