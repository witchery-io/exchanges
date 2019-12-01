package bitfinex

import (
	"context"
	"fmt"
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"github.com/bitfinexcom/bitfinex-api-go/v2/websocket"
	"github.com/sacOO7/gowebsocket"
	"github.com/witchery-io/go-exchanges/pkg/domain"
	"github.com/witchery-io/go-exchanges/pkg/exchange"
	"github.com/witchery-io/go-exchanges/pkg/util"
	"reflect"
	"time"
)

type client struct {
	*exchange.BaseExchangeClient

	publicClient *websocket.Client
	privateWS    gowebsocket.Socket
}

func (c *client) Start() {

	for obj := range c.publicClient.Listen() {
		switch v := obj.(type) {
		case error:
			c.BaseExchangeClient.ErrorsChannel <- v
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
		_, err := c.publicClient.SubscribeTrades(ctx, pair.String())
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
		_, err := c.publicClient.SubscribeTicker(ctx, pair.String())
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *client) Authenticate(map[string]string) error {
	panic("implement me")
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

	if c.publicClient.IsConnected() {
		return nil
	}

	if err := c.publicClient.Connect(); err != nil {
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
		publicClient: nil,
		privateWS:    gowebsocket.Socket{},
	}

	p := websocket.NewDefaultParameters()
	c.publicClient = websocket.NewWithParams(p)

	return c
}
