package domain

import "time"

/*
ENUM(
buy
sell
)
*/
type OrderDirection int

/*
ENUM(
exchange
margin
funding
)
*/
type OrderContext int

/*
ENUM(
stop
market
limit
)
*/
type OrderType int

/*
ENUM(
active
executed
canceled
partially
)
*/
type OrderStatus int

type Order struct {
	OrderNumber           string
	Direction             OrderDirection
	Context               OrderContext
	Type                  OrderType
	Pair                  CurrencyPair
	OriginalAmount        int64
	RemainingAmount       int64
	ExecutedAmount        int64
	Price                 int64
	AverageExecutionPrice int64
	OpenedAt              time.Time
	UpdatedAt             time.Time
	CanceledAt            time.Time
	AccountId             string
	Status                OrderStatus
	Exchange              string
}

/*
ENUM(
snapshot
new
update
cancel
)
*/
type OrderEventType int

type OrderEvent struct {
	Type  OrderEventType
	Order Order
}
