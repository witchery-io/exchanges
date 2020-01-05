package domain

// OrderBook represents universal order book structure for all exchanges
type OrderBook struct {
	Pair   CurrencyPair
	Price  int64
	Count  int64
	Amount int64
	Side   OrderDirection
}

// OrderBookEvent used when any exchange order book related update fires
type OrderBookEvent struct {
	OrderBook OrderBook
}
