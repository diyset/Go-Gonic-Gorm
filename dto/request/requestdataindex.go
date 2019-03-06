package request


type DataList struct {
	List []DetailData `json:"list"`
}

type DetailData struct {
	Test string `json:"test"`
}