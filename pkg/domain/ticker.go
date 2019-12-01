package domain

import "time"

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

type TickerEvent struct {
	Ticker Ticker
}
