package accounts

import (
	"context"

	"github.com/hafidhirsyad/account-svc/logger"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func (c *AccountRepository) Register(ctx context.Context, trx *gorm.DB, payload Register) (id int64, err error) {
	if trx == nil {
		trx = c.conn.WithContext(ctx)
	}

	sql := trx.Table(TableAccounts).Create(&payload)
	if sql.Error != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "Error Register", map[string]any{"error": sql.Error, "func": "Register", "path": "repository.accounts.register", "payload": payload})
		return 0, sql.Error
	}

	return sql.RowsAffected, nil
}
