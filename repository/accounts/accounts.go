package accounts

import (
	"context"

	"gorm.io/gorm"
)

type AccountRepositoryI interface {
	Register(ctx context.Context, trx *gorm.DB, payload Register) (id int64, err error)
	Deposit(ctx context.Context, trx *gorm.DB, payload Deposit) (id int64, err error)
	Withdraw(ctx context.Context, trx *gorm.DB, req Withdraw) (err error)
	GetBalanceByNoRekening(ctx context.Context, norek int64) (resp Balance, err error)
}

type AccountRepository struct {
	conn *gorm.DB
}

func NewAccountRepository(conn *gorm.DB) AccountRepositoryI {
	return &AccountRepository{conn}
}
