package domain

import "time"

/*
ENUM(
active
closed
)
*/
type PositionStatus int

/*
ENUM(
snapshot
new
update
cancel
)
*/
type PositionEventType int

type Position struct {
	Id                   string
	Pair                 CurrencyPair
	Status               PositionStatus
	Amount               int64
	BasePrice            int64
	MarginFunding        float64
	MarginFundingType    int64
	ProfitLoss           int64
	ProfitLossPercentage float64
	LiquidationPrice     int64
	AccountId            string
	Exchange             string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	//Leverage             int64
}

type PositionEvent struct {
	Type     PositionEventType
	Position Position
}
