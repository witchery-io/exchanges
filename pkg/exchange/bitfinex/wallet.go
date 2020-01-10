package bitfinex

import (
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"github.com/witchery-io/go-exchanges/pkg/domain"
	"github.com/witchery-io/go-exchanges/pkg/util"
)

// NewBalanceEvent makes domain type BalanceEvent from w and eventType
func (c *client) NewBalanceEvent(w *bitfinex.Wallet, eventType domain.BalanceEventType) *domain.BalanceEvent {
	currency := domain.Currency(w.Currency)

	return &domain.BalanceEvent{
		Type: eventType,
		Balance: domain.Balance{
			AccountID: c.AccountID,
			Currency:  currency,
			Name:      w.Type,
			Total:     util.CurrencyFloat64ToInt64(w.Balance, currency.String()),
			Exchange:  c.Name,
		},
	}
}
