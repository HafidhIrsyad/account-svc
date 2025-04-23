package accounts

import (
	"context"

	"gorm.io/gorm"
)

// Deposit implements AccountRepositoryI.
func (c *AccountRepository) Deposit(ctx context.Context, trx *gorm.DB, payload Deposit) (id int64, err error) {
	if trx == nil {
		trx = c.conn.WithContext(ctx)
	}

	sql := trx.Table(TableWallets).Create(&payload)
	if sql.Error != nil {
		return 0, sql.Error
	}

	return sql.RowsAffected, nil
}
