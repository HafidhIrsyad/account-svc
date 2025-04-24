package entity

import "errors"

type (
	RegisterReq struct {
		Name string `json:"nama"`
		NIK  string `json:"nik"`
		NoHP string `json:"no_hp"`
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
