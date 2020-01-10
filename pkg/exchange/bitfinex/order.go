package bitfinex

import (
	"strconv"
	"time"

	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"github.com/witchery-io/go-exchanges/pkg/domain"
	"github.com/witchery-io/go-exchanges/pkg/util"
)

// GetOrderType converts order type to domain Order type
func GetOrderType(o *bitfinex.Order) domain.OrderType {
	var dt domain.OrderType

	switch o.Type {
	case bitfinex.OrderTypeMarket:
		dt = domain.OrderTypeMarket
	case bitfinex.OrderTypeExchangeMarket:
		dt = domain.OrderTypeMarket
	case bitfinex.OrderTypeLimit:
		dt = domain.OrderTypeLimit
	case bitfinex.OrderTypeExchangeLimit:
		dt = domain.OrderTypeLimit
	case bitfinex.OrderTypeStop:
		dt = domain.OrderTypeStop
	case bitfinex.OrderTypeExchangeStop:
		dt = domain.OrderTypeStop
	case bitfinex.OrderTypeTrailingStop:
	case bitfinex.OrderTypeExchangeTrailingStop:
	case bitfinex.OrderTypeFOK:
	case bitfinex.OrderTypeExchangeFOK:
	case bitfinex.OrderTypeStopLimit:
	case bitfinex.OrderTypeExchangeStopLimit:
	}

	return dt
}

// GetOrderContext returns domain type Order Context
func GetOrderContext(o *bitfinex.Order) domain.OrderContext {
	var c domain.OrderContext

	switch o.Type {
	case bitfinex.OrderTypeExchangeMarket:
		c = domain.OrderContextExchange
	case bitfinex.OrderTypeExchangeLimit:
		c = domain.OrderContextExchange
	case bitfinex.OrderTypeExchangeStop:
		c = domain.OrderContextExchange
	default:
		c = domain.OrderContextMargin
	}

	return c
}

// GetOrderDirection returns domain type direction for given order
func GetOrderDirection(o *bitfinex.Order) domain.OrderDirection {
	if o.Amount < 0 {
		return domain.OrderDirectionSell
	}

	return domain.OrderDirectionBuy
}

// NewOrderEvent makes domain type OrderEvent from o and eventType
func (c *client) NewOrderEvent(o *bitfinex.Order, eventType domain.OrderEventType) *domain.OrderEvent {
	pair := domain.NewCurrencyPairFrom2Currencies(domain.Currency(o.Symbol[1:4]), domain.Currency(o.Symbol[4:]))

	return &domain.OrderEvent{
		Type: eventType,
		Order: domain.Order{
			OrderNumber:           strconv.Itoa(int(o.ID)),
			Direction:             GetOrderDirection(o),
			Context:               GetOrderContext(o),
			Type:                  GetOrderType(o),
			Pair:                  pair,
			OriginalAmount:        util.CurrencyFloat64ToInt64(o.AmountOrig, pair.GetCur1().String()),
			RemainingAmount:       util.CurrencyFloat64ToInt64(o.Amount, pair.GetCur1().String()),
			Price:                 util.CurrencyFloat64ToInt64(o.Price, pair.GetCur2().String()),
			AverageExecutionPrice: util.CurrencyFloat64ToInt64(o.PriceAvg, pair.GetCur2().String()),
			OpenedAt:              time.Unix(o.MTSCreated, 0),
			UpdatedAt:             time.Unix(o.MTSUpdated, 0),
			AccountID:             c.AccountID,
			Status:                statuses[o.Status],
			Exchange:              c.GetName(),
		},
	}
}
