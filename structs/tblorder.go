package structs

import "time"

type Order struct {
	IdOrder string `json:"idOrder" gorm:"primary_key;type:varchar(16);unique_index"`
	Audit
	TanggalOrder     time.Time `json:"tanggalOrder" time_format:"02-01-2006"`
	IdAlamat         int       `json:"idAlamat" gorm:"index"`
	IdNasabah        int       `json:"idNasabah" gorm:"index"`
	TotalHarga       float64   `json:"totalHarga"`
	TotalItem        int       `json:"totalItem"`
	MetodePengiriman int       `json:metodePengiriman`
}
