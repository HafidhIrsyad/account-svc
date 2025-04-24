package accounts

import (
	"context"

	"gorm.io/gorm"
)

type AccountRepositoryI interface {
	Register(ctx context.Context, trx *gorm.DB, payload Register) (id int64, err error)
	InsertBalance(ctx context.Context, trx *gorm.DB, payload Deposit) (id int64, err error)
	UpdateBalance(ctx context.Context, trx *gorm.DB, accId int64, updateNominal map[string]any) (err error)
	GetBalanceByFilter(ctx context.Context, filter Filter) (resp Balance, err error)
}

type AccountRepository struct {
	conn *gorm.DB
}

func NewAccountRepository(conn *gorm.DB) AccountRepositoryI {
	return &AccountRepository{conn}
}
