package accounts

const TableAccounts = "accounts"
const TableWallets = "wallets"

type (
	Register struct {
		Name       string `gorm:"column:name"`
		NIK        int64  `gorm:"column:nik"`
		NoHP       int64  `gorm:"column:no_hp"`
		NoRekening int64  `gorm:"column:no_rekening"`
	}

	Balance struct {
		WalletId   int64  `gorm:"column:wallet_id"`
		AccountId  int64  `gorm:"column:account_id"`
		NIK        int64  `gorm:"column:nik"`
		NoHP       int64  `gorm:"column:no_hp"`
		NoRekening int64  `gorm:"column:no_rekening"`
		Name       string `gorm:"column:name"`
		Nominal    int64  `gorm:"column:nominal"`
	}

	Deposit struct {
		Nominal   int64 `gorm:"column:nominal"`
		AccountId int64 `gorm:"column:account_id"`
	}

	Filter struct {
		NoHP       int64 `gorm:"column:no_hp"`
		NoRekening int64 `gorm:"column:no_rekening"`
		NIK        int64 `gorm:"column:nik"`
	}
)
