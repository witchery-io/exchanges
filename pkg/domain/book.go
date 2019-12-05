package domain

//type OrderBook struct {
//	Pair CurrencyPair
//	Asks []BookUpdate
//	Bids []BookUpdate
//}

type OrderBookEvent struct {
	OrderBook OrderBook
}

type OrderBook struct {
	Pair   CurrencyPair
	Price  int64
	Count  int64
	Amount int64
	Side   OrderDirection
}
