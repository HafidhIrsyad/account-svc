package accounts

import "context"

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
		return resp, sql.Error
	}

	return resp, nil
}
