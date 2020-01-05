package domain

import "time"

// PositionStatus status for position
/*
ENUM(
active
closed
)
*/
type PositionStatus int

// PositionEventType event type for position
/*
ENUM(
snapshot
new
update
cancel
)
*/
type PositionEventType int

// Position is universal position type for all exchange clients
type Position struct {
	ID                   string
	Pair                 CurrencyPair
	Status               PositionStatus
	Amount               int64
	BasePrice            int64
	MarginFunding        float64
	MarginFundingType    int64
	ProfitLoss           int64
	ProfitLossPercentage float64
	LiquidationPrice     int64
	AccountID            string
	Exchange             string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	//Leverage             int64
}

// PositionEvent used when any exchange position related update event fires
type PositionEvent struct {
	Type     PositionEventType
	Position Position
}
