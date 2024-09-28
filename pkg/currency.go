package pkg

import (
	"simple_bank/constant"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case constant.CNY, constant.USD, constant.CAD:
		return true
	}
	return false
}
