package accounts

import (
	"context"
	"time"

	"github.com/hafidhirsyad/account-svc/logger"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

// Withdraw implements AccountRepositoryI.
func (c *AccountRepository) Withdraw(ctx context.Context, trx *gorm.DB, req Withdraw) (err error) {
	if trx == nil {
		trx = c.conn.WithContext(ctx)
	}

	now := time.Now()
	updateNominal := map[string]any{
		"nominal":    req.Nominal,
		"updated_at": &now,
	}

	sql := trx.Table(TableWallets).Where("acoount_id = ?", req.AccountId).Updates(&updateNominal)
	if sql.Error != nil {
		logger.Log(ctx, zerolog.ErrorLevel, "Error when withdraw", map[string]any{"error": sql.Error, "func": "Withdraw", "path": "repository.accounts.withdraw", "payload": req})
		return sql.Error
	}

	return nil
}
