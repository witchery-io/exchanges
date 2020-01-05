package domain

import "strings"

// CurrencyPair is universal pair type for all exchange clients
type CurrencyPair string

func (c CurrencyPair) String() string {
	return string(c)
}

// GetCur1 get first Currency from pair
func (c CurrencyPair) GetCur1() Currency {
	return Currency(strings.Split(string(c), "|")[0])
}

// GetCur2 get second Currency from pair
func (c CurrencyPair) GetCur2() Currency {
	return Currency(strings.Split(string(c), "|")[1])
}

// NewCurrencyPairFrom2Currencies create currency pair from 2 currencies
func NewCurrencyPairFrom2Currencies(cur1, cur2 Currency) CurrencyPair {
	return CurrencyPair(cur1 + "|" + cur2)
}
