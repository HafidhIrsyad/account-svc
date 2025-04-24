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

	if d.Nominal == 0 {
		return errors.New("Nominal Saldo Wajib diisi")
	}

	return nil
}
