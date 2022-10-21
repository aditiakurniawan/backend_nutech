package barangdto

type BarangRequest struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama" form:"nama" validate:"required"`
	Foto      string `json:"foto" form:"foto" `
	Hargabeli string `json:"hargabeli" form:"hargabeli" `
	Hargajual string `json:"hargajual" form:"hargajual"`
	Stok      string `json:"stok" form:"stok"`
}

type UpdateBarangRequest struct {
	Nama      string `json:"nama" form:"nama" validate:"required"`
	Foto      string `json:"foto" form:"foto"`
	Hargabeli string `json:"hargabeli" form:"hargabeli" `
	Hargajual string `json:"hargajual" form:"hargajual"`
	Stok      string `json:"stok" form:"stok"`
}
