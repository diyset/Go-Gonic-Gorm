package structs

import "time"

type Warehouse struct {
	IdWarehouse uint `gorm:"primary_key" json:"idWarehouse"`
	Audit
	IdProduct int `gorm:"index" json:"idProduct"`
	Quantity int `json:"quantity"`
	TanggalKirim time.Time `json:"tanggalKirim" time_format:"02-01-2006"`
}