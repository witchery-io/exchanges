package exchange

import (
	"context"
	"github.com/witchery-io/go-exchanges/pkg/domain"
)

type BaseExchangeClient struct {
	Name            string
	AccountId       string
	isAuthenticated bool
	Credentials     map[string]string

	OrdersChannel     chan *domain.OrderEvent
	TradesChannel     chan *domain.TradeEvent
	TickersChannel    chan *domain.TickerEvent
	OrderBooksChannel chan *domain.OrderBookEvent
	PositionsChannel  chan *domain.PositionEvent
	BalancesChannel   chan *domain.BalanceEvent
	ErrorsChannel     chan error
}

type Client interface {
	GetName() string
	IsAuthenticated() bool

	Authenticate(accountId string, credentials map[string]string) error

	InitOrdersWatcher(ctx context.Context) error
	OrderEvents(ctx context.Context) <-chan *domain.OrderEvent

	InitPositionsWatcher(ctx context.Context) error
	PositionEvents(ctx context.Context) <-chan *domain.PositionEvent

	SubmitOrder(ctx context.Context, order *domain.Order) error
	UpdateOrder(ctx context.Context, orderId string, order *domain.Order) error
	CancelOrder(ctx context.Context, orderId string) error
	GetOrder(ctx context.Context, orderId string) error
	GetOrders(ctx context.Context) ([]*domain.Order, error)

	InitTradesWatcher(ctx context.Context, pairs []domain.CurrencyPair) error
	TradeEvents(ctx context.Context) <-chan *domain.TradeEvent

	InitTickersWatcher(ctx context.Context, pairs []domain.CurrencyPair) error
	TickerEvents(ctx context.Context) <-chan *domain.TickerEvent

	InitBalancesWatcher(ctx context.Context) error
	BalanceEvents(ctx context.Context) <-chan *domain.BalanceEvent

	InitOrderBooksWatcher(ctx context.Context, pairs []domain.CurrencyPair) error
	OrderBookEvents(ctx context.Context) <-chan *domain.OrderBookEvent

	Start()
	ErrorEvents() <-chan error
}

func (c *BaseExchangeClient) GetName() string {
	return c.Name
}

func (c *BaseExchangeClient) IsAuthenticated() bool {
	return c.isAuthenticated
}

func (c *BaseExchangeClient) OrderEvents(ctx context.Context) <-chan *domain.OrderEvent {
	return c.OrdersChannel
}

func (c *BaseExchangeClient) TradeEvents(ctx context.Context) <-chan *domain.TradeEvent {
	return c.TradesChannel
}

func (c *BaseExchangeClient) TickerEvents(ctx context.Context) <-chan *domain.TickerEvent {
	return c.TickersChannel
}

func (c *BaseExchangeClient) OrderBookEvents(ctx context.Context) <-chan *domain.OrderBookEvent {
	return c.OrderBooksChannel
}

func (c *BaseExchangeClient) PositionEvents(ctx context.Context) <-chan *domain.PositionEvent {
	return c.PositionsChannel
}

func (c *BaseExchangeClient) BalanceEvents(ctx context.Context) <-chan *domain.BalanceEvent {
	return c.BalancesChannel
}

func (c *BaseExchangeClient) ErrorEvents() <-chan error {
	return c.ErrorsChannel
}
