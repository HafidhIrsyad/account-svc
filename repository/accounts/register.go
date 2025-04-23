package accounts

import (
	"context"

	"gorm.io/gorm"
)

func (c *AccountRepository) Register(ctx context.Context, trx *gorm.DB, payload Register) (id int64, err error) {
	if trx == nil {
		trx = c.conn.WithContext(ctx)
	}

	sql := trx.Table(TableAccounts).Create(&payload)
	if sql.Error != nil {
		return 0, sql.Error
	}

	return sql.RowsAffected, nil
}
