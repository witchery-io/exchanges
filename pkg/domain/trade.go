package domain

import "time"

type Trade struct {
	Id             string
	Pair           CurrencyPair
	ExecutedAt     time.Time
	OrderNumber    string
	ExecutedAmount int64
	ExecutionPrice int64
	OrderType      OrderType
	OrderPrice     int64
	IsMaker        bool
	Fee            int64
	FeeCurrency    string
	Exchange       Exchange
}

/*
ENUM(
execute
update
)
*/
type TradeEventType int

type TradeEvent struct {
	Type  TradeEventType
	Trade Trade
}
