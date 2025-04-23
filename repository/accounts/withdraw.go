package accounts

import (
	"context"
	"time"

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
		return sql.Error
	}

	return nil
}
