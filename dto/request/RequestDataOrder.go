package request

type DataListOrder struct {
	Products         []DataProduct `json:"products"`
	IdNasabah        int           `json:"idNasabah"`
	IdAlamat         int           `json:"idAlamat"`
	MetodePengiriman int           `json:"metodePengiriman"`
}

type DataProduct struct {
	IdProduct string `json:"idProduct"`
	Quantity  int    `json:"quantity"`
}
