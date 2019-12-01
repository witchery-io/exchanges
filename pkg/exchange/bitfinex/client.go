package bitfinex

// @todo split file

import (
	"context"
	"fmt"
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"github.com/bitfinexcom/bitfinex-api-go/v2/websocket"
	"github.com/witchery-io/go-exchanges/pkg/domain"
	"github.com/witchery-io/go-exchanges/pkg/exchange"
	"github.com/witchery-io/go-exchanges/pkg/util"
	"reflect"
	"strconv"
	"time"
)

var statuses = map[bitfinex.OrderStatus]domain.OrderStatus{
	bitfinex.OrderStatusActive:          domain.OrderStatusActive,
	bitfinex.OrderStatusCanceled:        domain.OrderStatusCanceled,
	bitfinex.OrderStatusExecuted:        domain.OrderStatusExecuted,
	bitfinex.OrderStatusPartiallyFilled: domain.OrderStatusPartially,
}

type client struct {
	*exchange.BaseExchangeClient

	wsClient *websocket.Client
}

func (c *client) Start() {

	for obj := range c.wsClient.Listen() {
		switch v := obj.(type) {
		case error:
			c.BaseExchangeClient.ErrorsChannel <- v
		case *bitfinex.OrderNew, *bitfinex.OrderUpdate:

			o := obj.(bitfinex.Order)

			pair := domain.NewCurrencyPairFromString(o.Symbol)
			direction := domain.OrderDirectionBuy
			if o.Amount < 0 {
				direction = domain.OrderDirectionSell
			}

			var oType domain.OrderType
			oContext := domain.OrderContextMargin

			switch o.Type {
			case "MARKET":
				oType = domain.OrderTypeMarket
			case "EXCHANGE MARKET":
				oContext = domain.OrderContextExchange
				oType = domain.OrderTypeMarket
			case "LIMIT":
				oType = domain.OrderTypeLimit
			case "EXCHANGE LIMIT":
				oContext = domain.OrderContextExchange
				oType = domain.OrderTypeLimit
			case "STOP":
				oType = domain.OrderTypeStop
			case "EXCHANGE STOP":
				oContext = domain.OrderContextExchange
				oType = domain.OrderTypeStop
			case "TRAILING STOP":
			case "EXCHANGE TRAILING STOP":
			case "FOK":
			case "EXCHANGE FOK":
			case "STOP LIMIT":
			case "EXCHANGE STOP LIMIT":
			}

			c.BaseExchangeClient.OrdersChannel <- &domain.OrderEvent{
				Type: 0,
				Order: domain.Order{
					OrderNumber:           strconv.Itoa(int(o.ID)),
					Direction:             direction,
					Context:               oContext,
					Type:                  oType,
					Pair:                  pair,
					OriginalAmount:        util.CurrencyFloat64ToInt64(o.AmountOrig, pair.Currency1.String()),
					RemainingAmount:       util.CurrencyFloat64ToInt64(o.Amount, pair.Currency1.String()),
					Price:                 util.CurrencyFloat64ToInt64(o.Price, pair.Currency2.String()),
					AverageExecutionPrice: util.CurrencyFloat64ToInt64(o.PriceAvg, pair.Currency2.String()),
					OpenedAt:              time.Unix(o.MTSCreated, 0),
					UpdatedAt:             time.Unix(o.MTSUpdated, 0),
					AccountId:             "",
					Status:                statuses[o.Status],
					Exchange:              c.GetName(),
				},
			}
		case *bitfinex.OrderSnapshot:

			for _, o := range v.Snapshot {
				pair := domain.NewCurrencyPairFromString(o.Symbol)
				direction := domain.OrderDirectionBuy
				if o.Amount < 0 {
					direction = domain.OrderDirectionSell
				}

				var oType domain.OrderType
				oContext := domain.OrderContextMargin

				switch o.Type {
				case "MARKET":
					oType = domain.OrderTypeMarket
				case "EXCHANGE MARKET":
					oContext = domain.OrderContextExchange
					oType = domain.OrderTypeMarket
				case "LIMIT":
					oType = domain.OrderTypeLimit
				case "EXCHANGE LIMIT":
					oContext = domain.OrderContextExchange
					oType = domain.OrderTypeLimit
				case "STOP":
					oType = domain.OrderTypeStop
				case "EXCHANGE STOP":
					oContext = domain.OrderContextExchange
					oType = domain.OrderTypeStop
				case "TRAILING STOP":
				case "EXCHANGE TRAILING STOP":
				case "FOK":
				case "EXCHANGE FOK":
				case "STOP LIMIT":
				case "EXCHANGE STOP LIMIT":
				}

				c.BaseExchangeClient.OrdersChannel <- &domain.OrderEvent{
					Type: domain.OrderEventTypeSnapshot,
					Order: domain.Order{
						OrderNumber:           strconv.Itoa(int(o.ID)),
						Direction:             direction,
						Context:               oContext,
						Type:                  oType,
						Pair:                  pair,
						OriginalAmount:        util.CurrencyFloat64ToInt64(o.AmountOrig, pair.Currency1.String()),
						RemainingAmount:       util.CurrencyFloat64ToInt64(o.Amount, pair.Currency1.String()),
						Price:                 util.CurrencyFloat64ToInt64(o.Price, pair.Currency2.String()),
						AverageExecutionPrice: util.CurrencyFloat64ToInt64(o.PriceAvg, pair.Currency2.String()),
						OpenedAt:              time.Unix(o.MTSCreated, 0),
						UpdatedAt:             time.Unix(o.MTSUpdated, 0),
						AccountId:             "",
						Status:                statuses[o.Status],
						Exchange:              c.GetName(),
					},
				}
			}
		case *bitfinex.Trade:
		case *bitfinex.Ticker:
			pair := domain.NewCurrencyPairFromString(v.Symbol)

			c.BaseExchangeClient.TickersChannel <- &domain.TickerEvent{
				Ticker: domain.Ticker{
					Bid:                util.CurrencyFloat64ToInt64(v.Bid, pair.Currency1.String()),
					Ask:                util.CurrencyFloat64ToInt64(v.Ask, pair.Currency1.String()),
					Last:               util.CurrencyFloat64ToInt64(v.LastPrice, pair.Currency1.String()),
					Volume:             util.CurrencyFloat64ToInt64(v.Volume, pair.Currency1.String()),
					DailyChange:        util.CurrencyFloat64ToInt64(v.DailyChange, pair.Currency1.String()),
					DailyChangePercent: v.DailyChangePerc,
					High:               util.CurrencyFloat64ToInt64(v.High, pair.Currency1.String()),
					Low:                util.CurrencyFloat64ToInt64(v.Low, pair.Currency1.String()),
					Pair:               pair,
					UpdatedAt:          time.Now(),
					Exchange:           c.GetName(),
				},
			}
		default:
			fmt.Println(reflect.TypeOf(obj))
		}
	}

	return
}

func (c *client) InitOrdersWatcher(ctx context.Context) error {
	panic("implement me")
}

func (c *client) InitTradesWatcher(ctx context.Context, pairs []domain.CurrencyPair) error {
	err := c.connectPublicWS()
	if err != nil {
		return err
	}

	for _, pair := range pairs {
		_, err := c.wsClient.SubscribeTrades(ctx, pair.String())
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *client) InitTickersWatcher(ctx context.Context, pairs []domain.CurrencyPair) error {

	err := c.connectPublicWS()
	if err != nil {
		return err
	}

	for _, pair := range pairs {
		_, err := c.wsClient.SubscribeTicker(ctx, pair.String())
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *client) Authenticate(credentials map[string]string) error {

	c.wsClient = c.wsClient.Credentials(credentials["key"], credentials["secret"])

	return nil

}

func (c *client) SubmitOrder(ctx context.Context, order *domain.Order) error {
	panic("implement me")
}

func (c *client) UpdateOrder(ctx context.Context, orderId string, order *domain.Order) error {
	panic("implement me")
}

func (c *client) CancelOrder(ctx context.Context, orderId string) error {
	panic("implement me")
}

func (c *client) GetOrder(ctx context.Context, orderId string) error {
	panic("implement me")
}

func (c *client) GetOrders(ctx context.Context) ([]*domain.Order, error) {
	panic("implement me")
}

func (c *client) connectPublicWS() error {

	if c.wsClient.IsConnected() {
		return nil
	}

	if err := c.wsClient.Connect(); err != nil {
		return err
	}

	return nil
}

func (c *client) connectPrivateWS() error {

	if c.wsClient.IsConnected() {
		return nil
	}

	if err := c.wsClient.Connect(); err != nil {
		return err
	}

	return nil
}

func New() exchange.Client {

	c := &client{
		BaseExchangeClient: &exchange.BaseExchangeClient{
			Name:           "bitfinex",
			OrdersChannel:  make(chan *domain.OrderEvent, 1000),
			TradesChannel:  make(chan *domain.TradeEvent, 1000),
			TickersChannel: make(chan *domain.TickerEvent, 1000),
			ErrorsChannel:  make(chan error, 1000),
		},
		wsClient: nil,
	}

	p := websocket.NewDefaultParameters()
	c.wsClient = websocket.NewWithParams(p)

	return c
}
