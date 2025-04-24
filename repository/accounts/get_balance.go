package accounts

import (
	"context"

	"github.com/hafidhirsyad/account-svc/logger"
	"github.com/rs/zerolog"
)

func (c *AccountRepository) GetBalanceByFilter(ctx context.Context, filter Filter) (resp Balance, err error) {
	query := `
		select
			acc.id as account_id, acc."name", acc.no_rekening, acc.nik, acc.no_hp,
			wl.id as wallet_id, wl.nominal
		from accounts acc
		left join wallets wl on acc.id = wl.account_id
	`
	args := []any{}
	if filter.NIK != 0 || filter.NoHP != 0 {
		query += ` where acc.nik = ? or acc.no_hp = ? `
		args = append(args, filter.NIK)
		args = append(args, filter.NoHP)
	}

	if filter.NoRekening != 0 {
		query += ` where acc.no_rekening = ? `
		args = append(args, filter.NoRekening)
	}

	sql := c.conn.WithContext(ctx).Raw(query, args...).Scan(&resp)
	if sql.Error != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "Error when get balance by no rekening", map[string]any{"error": sql.Error, "func": "GetBalanceByNoRekening", "path": "repository.accounts.get_balance", "filter": filter})
		return resp, sql.Error
	}

	return resp, nil
}
