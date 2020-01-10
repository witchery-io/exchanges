package bitfinex

// @todo split file

import (
	"context"
	"fmt"
	"reflect"

	"github.com/Rhymond/go-money"
	"github.com/bitfinexcom/bitfinex-api-go/v2"
	"github.com/bitfinexcom/bitfinex-api-go/v2/websocket"
	"github.com/witchery-io/go-exchanges/pkg/domain"
	"github.com/witchery-io/go-exchanges/pkg/exchange"
	"github.com/witchery-io/go-exchanges/pkg/util"
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

func (c *client) InitBalancesWatcher(ctx context.Context) error {
	err := c.connectPublicWS()
	if err != nil {
		return err
	}

	return nil
}

// @todo simplify
func (c *client) Start() {
	for obj := range c.wsClient.Listen() {
		switch v := obj.(type) {
		case *websocket.ErrorEvent:
		case error:
			c.BaseExchangeClient.ErrorsChannel <- v
		case *bitfinex.WalletSnapshot:
			for _, w := range v.Snapshot {
				c.BaseExchangeClient.BalancesChannel <- c.NewBalanceEvent(w, domain.BalanceEventTypeSnapshot)
			}
		case *bitfinex.WalletUpdate:
			w := bitfinex.Wallet(*v)
			c.BaseExchangeClient.BalancesChannel <- c.NewBalanceEvent(&w, domain.BalanceEventTypeUpdate)
		case *bitfinex.PositionSnapshot:

			for _, p := range v.Snapshot {
				c.BaseExchangeClient.PositionsChannel <- c.NewPositionEvent(p, domain.PositionEventTypeSnapshot)
			}

		case *bitfinex.PositionNew:
			p := bitfinex.Position(*v)
			c.BaseExchangeClient.PositionsChannel <- c.NewPositionEvent(&p, domain.PositionEventTypeNew)
		case *bitfinex.PositionUpdate:
			p := bitfinex.Position(*v)
			c.BaseExchangeClient.PositionsChannel <- c.NewPositionEvent(&p, domain.PositionEventTypeUpdate)
		case *bitfinex.BookUpdate:

			pair := domain.NewCurrencyPairFrom2Currencies(domain.Currency(v.Symbol[0:3]), domain.Currency(v.Symbol[3:]))

			direction := domain.OrderDirectionBuy
			if v.Side == bitfinex.Short {
				direction = domain.OrderDirectionSell
			}

			c.BaseExchangeClient.OrderBooksChannel <- &domain.OrderBookEvent{
				OrderBook: domain.OrderBook{
					Pair:   pair,
					Price:  util.CurrencyFloat64ToInt64(v.Price, pair.GetCur1().String()),
					Count:  v.Count,
					Amount: util.CurrencyFloat64ToInt64(v.Amount, pair.GetCur1().String()),
					Side:   direction,
				},
			}
		case *bitfinex.OrderUpdate:
			o := bitfinex.Order(*v)
			c.BaseExchangeClient.OrdersChannel <- c.NewOrderEvent(&o, domain.OrderEventTypeUpdate)
		case *bitfinex.OrderNew:
			o := bitfinex.Order(*v)
			c.BaseExchangeClient.OrdersChannel <- c.NewOrderEvent(&o, domain.OrderEventTypeNew)
		case *bitfinex.OrderSnapshot:
			for _, o := range v.Snapshot {
				c.BaseExchangeClient.OrdersChannel <- c.NewOrderEvent(o, domain.OrderEventTypeSnapshot)
			}
		//case *bitfinex.Trade:
		case *bitfinex.Ticker:
			c.BaseExchangeClient.TickersChannel <- c.NewTickerEvent(v)
		default:
			fmt.Println(reflect.TypeOf(obj))
		}
	}
}

func (c *client) InitPositionsWatcher(ctx context.Context) error {
	err := c.connectPublicWS()
	if err != nil {
		return err
	}

	return nil
}

func (c *client) InitOrdersWatcher(ctx context.Context) error {
	err := c.connectPublicWS()
	if err != nil {
		return err
	}

	return nil
}

func (c *client) InitTradesWatcher(ctx context.Context, pairs []domain.CurrencyPair) error {
	err := c.connectPublicWS()
	if err != nil {
		return err
	}

	for _, pair := range pairs {
		_, _ = c.wsClient.SubscribeTrades(ctx, pair.String())
	}

	return nil
}

func (c *client) InitOrderBooksWatcher(ctx context.Context, pairs []domain.CurrencyPair) error {
	err := c.connectPublicWS()
	if err != nil {
		return err
	}

	for _, pair := range pairs {
		_, _ = c.wsClient.SubscribeBook(ctx, pair.String(), bitfinex.Precision0,
			bitfinex.FrequencyRealtime, 25)
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

func (c *client) Authenticate(accountID string, credentials map[string]string) error {
	c.Credentials = credentials
	c.AccountID = accountID
	c.wsClient = c.wsClient.Credentials(credentials["key"], credentials["secret"])

	return nil
}

func (c *client) SubmitOrder(ctx context.Context, order *domain.Order) error {
	amount := order.OriginalAmount
	if order.Direction == domain.OrderDirectionSell {
		amount = -1 * amount
	}

	var oType bitfinex.OrderType
	if order.Type == domain.OrderTypeMarket && order.Context == domain.OrderContextMargin {
		oType = bitfinex.OrderTypeMarket
	}
	if order.Type == domain.OrderTypeMarket && order.Context == domain.OrderContextExchange {
		oType = bitfinex.OrderTypeExchangeMarket
	}

	if order.Type == domain.OrderTypeLimit && order.Context == domain.OrderContextMargin {
		oType = bitfinex.OrderTypeLimit
	}
	if order.Type == domain.OrderTypeLimit && order.Context == domain.OrderContextExchange {
		oType = bitfinex.OrderTypeExchangeLimit
	}

	if order.Type == domain.OrderTypeStop && order.Context == domain.OrderContextMargin {
		oType = bitfinex.OrderTypeStop
	}
	if order.Type == domain.OrderTypeStop && order.Context == domain.OrderContextExchange {
		oType = bitfinex.OrderTypeExchangeStop
	}

	o := &bitfinex.OrderNewRequest{
		Type:   string(oType),
		Symbol: "t" + order.Pair.String(),
		Amount: money.New(amount, order.Pair.GetCur1().String()).AsMajorUnits(),
		Price:  money.New(order.Price, order.Pair.GetCur2().String()).AsMajorUnits(),
	}

	err := c.wsClient.SubmitOrder(ctx, o)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) UpdateOrder(ctx context.Context, orderID string, order *domain.Order) error {
	panic("implement me")
}

func (c *client) CancelOrder(ctx context.Context, orderID string) error {
	panic("implement me")
}

func (c *client) GetOrder(ctx context.Context, orderID string) error {
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

// New new bitfinex Client implementation
func New() exchange.Client {
	c := &client{
		BaseExchangeClient: &exchange.BaseExchangeClient{
			Name:              "bitfinex",
			OrdersChannel:     make(chan *domain.OrderEvent),
			TradesChannel:     make(chan *domain.TradeEvent),
			TickersChannel:    make(chan *domain.TickerEvent),
			OrderBooksChannel: make(chan *domain.OrderBookEvent),
			PositionsChannel:  make(chan *domain.PositionEvent),
			BalancesChannel:   make(chan *domain.BalanceEvent),
			ErrorsChannel:     make(chan error),
		},
		wsClient: nil,
	}

	p := websocket.NewDefaultParameters()
	p.ManageOrderbook = true
	c.wsClient = websocket.NewWithParams(p)

	return c
}
