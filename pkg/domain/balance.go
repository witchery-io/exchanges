package domain

// Balance represents universal balance structure for all exchanges
type Balance struct {
	AccountID string
	Currency  Currency
	Name      string
	Total     int64
	Exchange  string
}

// BalanceEventType used by subscriber universal for all exchanges
/*
ENUM(
snapshot
update
)
*/
type BalanceEventType int

// BalanceEvent used when any exchange balance related update fires
type BalanceEvent struct {
	Type    BalanceEventType
	Balance Balance
}
