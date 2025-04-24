package accounts

import (
	"context"

	"github.com/hafidhirsyad/account-svc/logger"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func (c *AccountRepository) UpdateBalance(ctx context.Context, trx *gorm.DB, accId int64, updateNominal map[string]any) (err error) {
	if trx == nil {
		trx = c.conn.WithContext(ctx)
	}

	sql := trx.Table(TableWallets).Where("account_id = ?", accId).Updates(&updateNominal)
	if sql.Error != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "Error when withdraw", map[string]any{"error": sql.Error, "func": "Withdraw", "path": "repository.accounts.withdraw", "payload": updateNominal})
		return sql.Error
	}

	return nil
}
