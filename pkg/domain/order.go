package domain

import "time"

// OrderDirection direction for Order
/*
ENUM(
buy
sell
)
*/
type OrderDirection int

// OrderContext context for Order
/*
ENUM(
exchange
margin
funding
)
*/
type OrderContext int

// OrderType type for Order
/*
ENUM(
stop
market
limit
)
*/
type OrderType int

// OrderStatus status for Order
/*
ENUM(
active
executed
canceled
partially
)
*/
type OrderStatus int

// Order represents universal order structure for all exchanges
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
	AccountID             string
	Status                OrderStatus
	Exchange              string
}

// OrderEventType used by subscriber universal for all exchanges
/*
ENUM(
snapshot
new
update
cancel
)
*/
type OrderEventType int

// OrderEvent used when any exchange order related update event fires
type OrderEvent struct {
	Type  OrderEventType
	Order Order
}
