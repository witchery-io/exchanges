package util

import (
	"math"

	"github.com/Rhymond/go-money"
)

// CurrencyFloat64ToInt64 convert money from float to integer
func CurrencyFloat64ToInt64(in float64, currency string) int64 {
	return int64(in * math.Pow(10, float64(money.GetCurrency(currency).Fraction)))
}
