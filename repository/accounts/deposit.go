package accounts

import (
	"context"

	"github.com/hafidhirsyad/account-svc/logger"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

// Deposit implements AccountRepositoryI.
func (c *AccountRepository) Deposit(ctx context.Context, trx *gorm.DB, payload Deposit) (id int64, err error) {
	if trx == nil {
		trx = c.conn.WithContext(ctx)
	}

	sql := trx.Table(TableWallets).Create(&payload)
	if sql.Error != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "Error when deposit", map[string]any{"error": sql.Error, "func": "Deposit", "path": "repository.accounts.deposit", "payload": payload})
		return 0, sql.Error
	}

	return sql.RowsAffected, nil
}
