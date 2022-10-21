package models

type Barang struct {
	ID        int    `json:"id" gorm:"primary_key:auto_increment" `
	Nama      string `json:"nama" form:"nama" gorm:"type: varchar(255)"`
	Foto      string `json:"foto" form:"foto" gorm:"type: varchar(255)"`
	Hargabeli string `json:"hargabeli" form:"hargabeli" gorm:"type:varchar(255)"`
	Hargajual string `json:"hargajual" form:"hargajual" gorm:"type:varchar(255)"`
	Stok      string `json:"stok" form:"stok" gorm:"type:varchar(255)"`
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
