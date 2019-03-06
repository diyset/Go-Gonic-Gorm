package structs

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Nasabah struct {
	gorm.Model
	First_Name   string    `json:"firstName"`
	Last_Name    string    `json:"lastName"`
	Email        string    `gorm:"type:varchar(100);unique_index" json:"email"`
	JenisKelamin string    `json:"jenisKelamin"`
	TanggalLahir time.Time `json:"tanggalLahir" time_format:"02-01-2006"`
	IsAdult      bool      `json:"isAdult"`
}

func GetFullNameNasabah(nasabah Nasabah) (string) {
	return nasabah.First_Name + " " + nasabah.Last_Name
}

type Alamat struct {
	IdAlamat uint `gorm:"primary_key" json:"idAlamat"`
	Audit //extend Audit.go
	NamaJalan string `json:"namaJalan"`
	Rt        string `json:"rt"`
	Rw        string `json:"rw"`
	No        string `json:"no"`
	Provinsi  string `json:"provinsi"`
	IdPerson  int    `gorm:"index" json:"idPerson"`
}
