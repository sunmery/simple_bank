package pkg

import (
	"github.com/go-playground/validator/v10"
)

const (
	CNY = "CNY"
	USD = "USD"
	EDR = "EDR"
)

func ValidateCurrency(fl validator.FieldLevel) bool {
	currency, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	switch currency {
	case CNY, USD, EDR:
		return true
	}
	return false
}
