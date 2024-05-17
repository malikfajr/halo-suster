package customvalidator

import (
	"log"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

func validNip(fl validator.FieldLevel) bool {
	role := fl.Param()

	v := fl.Field().Int()

	nip := strconv.Itoa(int(v))
	if len(nip) < 10 {
		log.Println("length below 10 char")
		return false
	}

	prefix := nip[:3]
	gender := nip[3:4]
	year := nip[4:8]
	month := nip[8:10]
	unix := nip[10:]

	if role == "it" && prefix != "615" {
		log.Println("not valid it prefix")
		return false
	}

	if role == "nurse" && prefix != "303" {
		log.Println("not valid nurse prefix")
		return false
	}

	if validGender(gender) == false {
		log.Println("wrong gender code")
		return false
	}

	if validYear(year) == false {
		log.Println("year not valid")
		return false
	}

	if validMonth(month) == false {
		log.Println("month not valid")
		return false
	}

	if len(unix) < 3 || len(unix) > 5 {
		log.Println("unique not valid.", len(unix))
		return false
	}

	return true
}

func validGender(key string) bool {
	genders := map[string]bool{
		"1": true,
		"2": true,
	}

	_, ok := genders[key]
	return ok
}

func validYear(key string) bool {
	year, err := strconv.Atoi(key)
	if err != nil {
		return false
	}

	if year < 2000 || year > time.Now().Year() {
		return false
	}

	return true
}

func validMonth(key string) bool {
	months := map[string]bool{
		"01": true,
		"02": true,
		"03": true,
		"04": true,
		"05": true,
		"06": true,
		"07": true,
		"08": true,
		"09": true,
		"10": true,
		"11": true,
		"12": true,
	}

	_, ok := months[key]
	return ok
}
