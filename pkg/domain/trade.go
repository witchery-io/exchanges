package domain

import "time"

// Trade is universal trade type for all exchange clients
type Trade struct {
	ID             string
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

// TradeEventType event type for Trade
/*
ENUM(
execute
update
)
*/
type TradeEventType int

// TradeEvent used when any exchange trade related update event fires
type TradeEvent struct {
	Type  TradeEventType
	Trade Trade
}
