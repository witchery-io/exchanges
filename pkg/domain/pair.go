package domain

type CurrencyPair struct {
	Currency1 Currency
	Currency2 Currency
}

func (c *CurrencyPair) String() string {
	return c.Currency1.String() + c.Currency2.String()
}

func NewCurrencyPairFromString(pair string) CurrencyPair {
	return CurrencyPair{
		Currency1: NewCurrencyFromString(pair[0:3]),
		Currency2: NewCurrencyFromString(pair[3:]),
	}
}
