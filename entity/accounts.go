package entity

import "errors"

type (
	RegisterReq struct {
		Name string `json:"nama"`
		NIK  string `json:"nik"`
		NoHP string `json:"no_hp"`
	}
	DepositReq struct {
		NoRekening string `json:"no_rekening"`
		Nominal    int64  `json:"nominal"`
	}
	WithdrawReq struct {
		NoRekening string `json:"no_rekening"`
		Nominal    int64  `json:"nominal"`
	}
	Balance struct {
		WalletId   int64  `json:"wallet_id"`
		AccountId  int64  `json:"account_id"`
		NIK        int64  `json:"nik"`
		NoHP       int64  `json:"no_hp"`
		NoRekening int64  `json:"no_rekening"`
		Name       string `json:"name"`
		Nominal    int64  `json:"nominal"`
	}
)

func (r *RegisterReq) ValidateRegister() (err error) {
	if r.Name == "" {
		return errors.New("Nama Wajib diisi")
	}

	if r.NIK == "" {
		return errors.New("NIK Wajib diisi")
	}

	if len(r.NIK) < 16 {
		return errors.New("NIK minimal 16 charactes")
	}

	if r.NoHP == "" {
		return errors.New("No HP Wajib diisi")
	}

	return nil
}

func (d *DepositReq) ValidateDeposit() (err error) {
	if d.NoRekening == "" {
		return errors.New("Nomor Rekening Wajib diisi")
	}

	if d.Nominal <= 0 {
		return errors.New("Nomila tidak boleh 0 atau minus")
	}

	return nil
}

func (w *WithdrawReq) ValidateWithdraw() (err error) {
	if w.NoRekening == "" {
		return errors.New("Nomor Rekening Wajib diisi")
	}

	if w.Nominal <= 0 {
		return errors.New("Nomila tidak boleh 0 atau minus")
	}

	return nil
}
