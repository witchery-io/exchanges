package exchange

import (
	"context"
	"log"

	"github.com/witchery-io/go-exchanges/pkg/domain"
)

// BaseExchangeClient abstract Client interface implementation
type BaseExchangeClient struct {
	Name            string
	AccountID       string
	isAuthenticated bool
	Credentials     map[string]string

	Currencies    []domain.Currency
	CurrencyPairs []domain.CurrencyPair

	OrdersChannel     chan *domain.OrderEvent
	TradesChannel     chan *domain.TradeEvent
	TickersChannel    chan *domain.TickerEvent
	OrderBooksChannel chan *domain.OrderBookEvent
	PositionsChannel  chan *domain.PositionEvent
	BalancesChannel   chan *domain.BalanceEvent
	ErrorsChannel     chan error
}

type ClientOptions struct {
	Logger     *log.Logger
	Debug      bool
	// @todo rid off generic
	Additional interface{}
}

// Client interface for all exchanges
type Client interface {
	// GetName return exchange name
	GetName() string
	// IsAuthenticated checks if current client is authenticated
	IsAuthenticated() bool

	// Authenticate attempt to authenticate current client
	Authenticate(accountID string, credentials map[string]string) error

	// InitOrdersWatcher prepare client to start watching orders
	InitOrdersWatcher(ctx context.Context) error
	// OrderEvents get order updates channel
	OrderEvents(ctx context.Context) <-chan *domain.OrderEvent

	// InitPositionsWatcher prepare client to start watching positions
	InitPositionsWatcher(ctx context.Context) error
	// PositionEvents get position updates channel
	PositionEvents(ctx context.Context) <-chan *domain.PositionEvent

	// SubmitOrder submit order to exchange
	SubmitOrder(ctx context.Context, order *domain.Order) error
	// UpdateOrder submit order update to exchange
	UpdateOrder(ctx context.Context, orderID string, order *domain.Order) error
	// CancelOrder submit order cancel to exchange
	CancelOrder(ctx context.Context, orderID string) error
	// GetOrder get order with given ID from exchange
	GetOrder(ctx context.Context, orderID string) error
	// GetOrders get all active orders from exchange
	GetOrders(ctx context.Context) ([]*domain.Order, error)

	// InitTradesWatcher prepare client to start watching trades
	InitTradesWatcher(ctx context.Context, pairs []domain.CurrencyPair) error
	// TradeEvents get trade updates channel
	TradeEvents(ctx context.Context) <-chan *domain.TradeEvent

	// InitTickersWatcher prepare client to start watching tickers
	InitTickersWatcher(ctx context.Context, pairs []domain.CurrencyPair) error
	// TickerEvents get ticker updates channel
	TickerEvents(ctx context.Context) <-chan *domain.TickerEvent

	// InitBalancesWatcher prepare client to start watching balances
	InitBalancesWatcher(ctx context.Context) error
	// BalanceEvents get balance updates channel
	BalanceEvents(ctx context.Context) <-chan *domain.BalanceEvent

	// InitOrderBooksWatcher prepare client to start watching order books
	InitOrderBooksWatcher(ctx context.Context, pairs []domain.CurrencyPair) error
	// OrderBookEvents get order book updates channel
	OrderBookEvents(ctx context.Context) <-chan *domain.OrderBookEvent

	// Start listening exchange events
	Start()
	// ErrorEvents get errors updates channel
	ErrorEvents() <-chan error
}

// GetName satisfies Client interface
func (c *BaseExchangeClient) GetName() string {
	return c.Name
}

// IsAuthenticated satisfies Client interface
func (c *BaseExchangeClient) IsAuthenticated() bool {
	return c.isAuthenticated
}

// OrderEvents satisfies Client interface
func (c *BaseExchangeClient) OrderEvents(ctx context.Context) <-chan *domain.OrderEvent {
	return c.OrdersChannel
}

// TradeEvents satisfies Client interface
func (c *BaseExchangeClient) TradeEvents(ctx context.Context) <-chan *domain.TradeEvent {
	return c.TradesChannel
}

// TickerEvents satisfies Client interface
func (c *BaseExchangeClient) TickerEvents(ctx context.Context) <-chan *domain.TickerEvent {
	return c.TickersChannel
}

// OrderBookEvents satisfies Client interface
func (c *BaseExchangeClient) OrderBookEvents(ctx context.Context) <-chan *domain.OrderBookEvent {
	return c.OrderBooksChannel
}

// PositionEvents satisfies Client interface
func (c *BaseExchangeClient) PositionEvents(ctx context.Context) <-chan *domain.PositionEvent {
	return c.PositionsChannel
}

// BalanceEvents satisfies Client interface
func (c *BaseExchangeClient) BalanceEvents(ctx context.Context) <-chan *domain.BalanceEvent {
	return c.BalancesChannel
}

// ErrorEvents satisfies Client interface
func (c *BaseExchangeClient) ErrorEvents() <-chan error {
	return c.ErrorsChannel
}
