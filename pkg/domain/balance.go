package domain

type Balance struct {
	AccountId string
	Currency  Currency
	Name      string
	Total     int64
	Exchange  string
}

/*
ENUM(
snapshot
update
)
*/
type BalanceEventType int

type BalanceEvent struct {
	Type    BalanceEventType
	Balance Balance
}
