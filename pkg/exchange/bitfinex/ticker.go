package bitfinex

import (
	"time"

	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"github.com/witchery-io/go-exchanges/pkg/domain"
	"github.com/witchery-io/go-exchanges/pkg/util"
)

// NewTickerEvent makes domain type TickerEvent from t
func (c *client) NewTickerEvent(t *bitfinex.Ticker) *domain.TickerEvent {
	pair := domain.NewCurrencyPairFrom2Currencies(domain.Currency(t.Symbol[0:3]), domain.Currency(t.Symbol[3:]))

	return &domain.TickerEvent{
		Ticker: domain.Ticker{
			Bid:                util.CurrencyFloat64ToInt64(t.Bid, pair.GetCur1().String()),
			Ask:                util.CurrencyFloat64ToInt64(t.Ask, pair.GetCur1().String()),
			Last:               util.CurrencyFloat64ToInt64(t.LastPrice, pair.GetCur1().String()),
			Volume:             util.CurrencyFloat64ToInt64(t.Volume, pair.GetCur1().String()),
			DailyChange:        util.CurrencyFloat64ToInt64(t.DailyChange, pair.GetCur1().String()),
			DailyChangePercent: t.DailyChangePerc,
			High:               util.CurrencyFloat64ToInt64(t.High, pair.GetCur1().String()),
			Low:                util.CurrencyFloat64ToInt64(t.Low, pair.GetCur1().String()),
			Pair:               pair,
			UpdatedAt:          time.Now(),
			Exchange:           c.GetName(),
		},
	}
}
