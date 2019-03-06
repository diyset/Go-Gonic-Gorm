package structs

type OrderDetail struct {
	IdOrderDetail string `json:"idOrderDetail" gorm:"primary_key;type:varchar(12);unique_index"`
	Audit
	IdOrder             string  `json:"idOrder" gorm:"type:varchar(16);index"`
	IdProduct           int     `json:"idProduct" gorm:"index"`
	Quantity            int     `json:"quantity"`
	TotaHargaPerProduct float64 `json:"totalHargaPerProduct"`
}
