package barangdto

type BarangResponse struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama" validate:"required"`
	Foto      string `json:"foto" validate:"required"`
	Hargabeli string `json:"hargabeli" validate:"required"`
	Hargajual string `json:"hargajual"`
	Stok      string `json:"stock"`
}
