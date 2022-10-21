package models

type Barang struct {
	ID        int    `json:"id" gorm:"primary_key:auto_increment" `
	Nama      string `json:"nama" gorm:"type: varchar(255)"`
	Foto      string `json:"foto" gorm:"type: varchar(255)"`
	Hargabeli string `json:"hargabeli" gorm:"type:varchar(255)"`
	Hargajual string `json:"hargajual" gorm:"type:varchar(255)"`
	Stok      string `json:"stok" gorm:"type:varchar(255)"`
	// UserID    int    `json:"userid"`
	// User      string `json:"user"`
}

type BarangResponse struct {
	ID        int    `json:"id" `
	Nama      string `json:"nama" `
	Foto      string `json:"foto" `
	Hargabeli string `json:"hargabeli" `
	Hargajual string `json:"hargajual" `
	Stok      string `json:"stok" `
}

func (BarangResponse) TableName() string {
	return "barangs"
}
