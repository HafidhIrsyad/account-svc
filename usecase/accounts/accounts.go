package accounts

import (
	"context"
	"math/rand"
	"time"

	"github.com/hafidhirsyad/account-svc/entity"
	repo "github.com/hafidhirsyad/account-svc/repository/accounts"
	trxRepo "github.com/hafidhirsyad/account-svc/repository/transaction"
)

type AccountServiceI interface {
	Register(ctx context.Context, req entity.RegisterReq) (norek int64, err error)
	Deposit(ctx context.Context, req entity.DepositReq) (saldo int64, err error)
	Withdraw(ctx context.Context, req entity.WithdrawReq) (saldo int64, err error)
	GetBalanceByNoRekening(ctx context.Context, norek int64) (saldo int64, err error)
}

type AccountService struct {
	accRepo repo.AccountRepositoryI
	trxRepo trxRepo.TransactionRepositoryI
	rand    *rand.Rand
}

func NewAccountService(accRepo repo.AccountRepositoryI, trxRepo trxRepo.TransactionRepositoryI) AccountServiceI {
	return &AccountService{
		accRepo: accRepo,
		trxRepo: trxRepo,
		rand:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
