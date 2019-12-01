package domain

type Currency string

func (c *Currency) String() string {
	return string(*c)
}

func NewCurrencyFromString(currency string) Currency {
	return Currency(currency)
}
