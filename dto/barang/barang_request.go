package barangdto

type BarangRequest struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama" validate:"required"`
	Foto      string `json:"foto" validate:"required"`
	Hargabeli string `json:"hargabeli" `
	Hargajual string `json:"hargajual"`
	Stok      string `json:"stok"`
}

type UpdateBarangRequest struct {
	Nama      string `json:"nama" validate:"required"`
	Foto      string `json:"foto" validate:"required"`
	Hargabeli string `json:"hargabeli" `
	Hargajual string `json:"hargajual"`
	Stok      string `json:"stok"`
}
