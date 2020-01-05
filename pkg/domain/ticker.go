package domain

import "time"

// Ticker is universal ticker type for all exchange clients
type Ticker struct {
	Bid                int64        `json:"bid"`
	Ask                int64        `json:"ask"`
	Last               int64        `json:"last"`
	Volume             int64        `json:"volume"`
	DailyChange        int64        `json:"dailyChange"`
	DailyChangePercent float64      `json:"dailyChangePercent"`
	High               int64        `json:"high"`
	Low                int64        `json:"low"`
	Pair               CurrencyPair `json:"pair"`
	UpdatedAt          time.Time    `json:"updatedAt"`
	Exchange           string       `json:"exchange"`
}

// TickerEvent used when any exchange ticker related update event fires
type TickerEvent struct {
	Ticker Ticker
}
