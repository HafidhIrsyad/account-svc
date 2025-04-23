package accounts

import (
	"context"

	"github.com/hafidhirsyad/account-svc/logger"
	"github.com/rs/zerolog"
)

func (c *AccountRepository) GetBalanceByNoRekening(ctx context.Context, norek int64) (resp Balance, err error) {
	query := `
		select
			acc.id as account_id, acc."name", acc.no_rekening,
			wl.id as wallet_id, wl.nominal
		from accounts acc
		left join wallets wl on acc.id = wl.account_id
		where acc.no_rekening = ?
	`

	sql := c.conn.WithContext(ctx).Raw(query, norek).Scan(&resp)
	if sql.Error != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "Error when get balance by no rekening", map[string]any{"error": sql.Error, "func": "GetBalanceByNoRekening", "path": "repository.accounts.get_balance", "payload": norek})
		return resp, sql.Error
	}

	return resp, nil
}
