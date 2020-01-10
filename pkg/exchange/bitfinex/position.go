package bitfinex

import (
	"strconv"

	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"github.com/witchery-io/go-exchanges/pkg/domain"
	"github.com/witchery-io/go-exchanges/pkg/util"
)

// NewPositionEvent makes domain type PositionEvent p and sets given eventType
func (c *client) NewPositionEvent(p *bitfinex.Position, eventType domain.PositionEventType) *domain.PositionEvent {
	pair := domain.NewCurrencyPairFrom2Currencies(domain.Currency(p.Symbol[1:4]), domain.Currency(p.Symbol[4:]))

	status := domain.PositionStatusActive
	if p.Status == bitfinex.PositionStatusClosed {
		status = domain.PositionStatusClosed
	}

	return &domain.PositionEvent{
		Type: eventType,
		Position: domain.Position{
			ID:                   strconv.Itoa(int(p.Id)),
			Pair:                 pair,
			Status:               status,
			Amount:               util.CurrencyFloat64ToInt64(p.Amount, pair.GetCur1().String()),
			BasePrice:            util.CurrencyFloat64ToInt64(p.BasePrice, pair.GetCur2().String()),
			MarginFunding:        p.MarginFunding,
			MarginFundingType:    p.MarginFundingType,
			ProfitLoss:           util.CurrencyFloat64ToInt64(p.ProfitLoss, pair.GetCur2().String()),
			ProfitLossPercentage: p.ProfitLossPercentage,
			LiquidationPrice:     util.CurrencyFloat64ToInt64(p.LiquidationPrice, pair.GetCur2().String()),
			Exchange:             c.GetName(),
			AccountID:            c.AccountID,
		},
	}
}