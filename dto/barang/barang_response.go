package barangdto

import "nutech/models"

type BarangResponse struct {
	ID        int         `json:"id"`
	Nama      string      `json:"nama" validate:"required"`
	Foto      string      `json:"foto" validate:"required"`
	Hargabeli string      `json:"hargabeli" validate:"required"`
	Hargajual string      `json:"hargajual"`
	Stok      string      `json:"stock"`
	UserID    int         `json:"user_id"`
	User      models.User `json:"user"`
}
