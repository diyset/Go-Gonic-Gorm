package structs

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"time"
)

type Warehouse struct {
	IdWarehouse uint `gorm:"primary_key" json:"idWarehouse"`
	Audit
	IdProduct    int       `gorm:"index" json:"idProduct"`
	Quantity     int       `json:"quantity" validate:"required"`
	TanggalKirim time.Time `json:"tanggalKirim" time_format:"02-01-2006" binding:"required"`
}

func WarehouseDateValidate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}
	return true
}
