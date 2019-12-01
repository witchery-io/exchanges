package domain

type Balance struct {
	AccountId string
	Currency  string
	Name      string
	Value1    int64
	Value2    int64
	Exchange  Exchange
}
