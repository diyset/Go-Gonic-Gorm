package utility

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func ValidateEmailFormat(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	valid := false
	if re.MatchString(email) {
		valid = true
	}
	return valid
}

func ValidateDateFormat(date string) bool {
	re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])-(0?[1-9]|1[012])-((19|20)\\d\\d)")
	valid := false
	if re.MatchString(date) {
		valid = true
	}
	return valid
}

func DateFormatMyApp(date string) (time.Time, error) {
	parsingTime, err := time.Parse("02-01-2006", date)
	if err != nil {
		panic("Error Parsing Date!")
	}
	return parsingTime, err
}

func DateFormatMyLayout(date time.Time) (string) {
	return date.Format("01-02-2006")
}

func GeneratorIdOrder(metodePengiriman int, idNasabah int) (string) {
	timeSuffix := int(time.Time.UnixNano(time.Now()))/100000
	prefixId := strconv.Itoa(timeSuffix) + strconv.Itoa(metodePengiriman) + strconv.Itoa(idNasabah)
	fmt.Println(len(prefixId))
	return prefixId
}
