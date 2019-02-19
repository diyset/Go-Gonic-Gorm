package structs

type Product struct {
	IdProduct uint `gorm:"primary_key" json:"idProduct"`
	Audit
	NameProduct string  `json:"nameProduct" gorm:"type:varchar(100)"`
	IdCategory  int     `json:"idCategory"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}
