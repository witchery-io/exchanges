package bitfinex

// @todo split file

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"time"

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
				currency := domain.Currency(w.Currency)

				c.BaseExchangeClient.BalancesChannel <- &domain.BalanceEvent{
					Type: domain.BalanceEventTypeSnapshot,
					Balance: domain.Balance{
						AccountID: c.AccountID,
						Currency:  currency,
						Name:      w.Type,
						Total:     util.CurrencyFloat64ToInt64(w.Balance, currency.String()),
						Exchange:  c.Name,
					},
				}
			}
		case *bitfinex.WalletUpdate:

			w := v
			currency := domain.Currency(w.Currency)

			c.BaseExchangeClient.BalancesChannel <- &domain.BalanceEvent{
				Type: domain.BalanceEventTypeUpdate,
				Balance: domain.Balance{
					AccountID: c.AccountID,
					Currency:  currency,
					Name:      w.Type,
					Total:     util.CurrencyFloat64ToInt64(w.Balance, currency.String()),
					Exchange:  c.Name,
				},
			}
		case *bitfinex.PositionSnapshot:

			for _, p := range v.Snapshot {
				pair := domain.NewCurrencyPairFrom2Currencies(domain.Currency(p.Symbol[1:4]), domain.Currency(p.Symbol[4:]))

				status := domain.PositionStatusActive
				if p.Status == bitfinex.PositionStatusClosed {
					status = domain.PositionStatusClosed
				}

				c.BaseExchangeClient.PositionsChannel <- &domain.PositionEvent{
					Type: domain.PositionEventTypeSnapshot,
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

		case *bitfinex.PositionNew:

			p := obj.(*bitfinex.PositionNew)

			pair := domain.NewCurrencyPairFrom2Currencies(domain.Currency(p.Symbol[1:4]), domain.Currency(p.Symbol[4:]))

			status := domain.PositionStatusActive
			if p.Status == bitfinex.PositionStatusClosed {
				status = domain.PositionStatusClosed
			}

			c.BaseExchangeClient.PositionsChannel <- &domain.PositionEvent{
				Type: domain.PositionEventTypeNew,
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
		case *bitfinex.PositionUpdate:

			p := obj.(*bitfinex.PositionUpdate)

			pair := domain.NewCurrencyPairFrom2Currencies(domain.Currency(p.Symbol[1:4]), domain.Currency(p.Symbol[4:]))

			status := domain.PositionStatusActive
			if p.Status == bitfinex.PositionStatusClosed {
				status = domain.PositionStatusClosed
			}

			c.BaseExchangeClient.PositionsChannel <- &domain.PositionEvent{
				Type: domain.PositionEventTypeUpdate,
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

			o := obj.(bitfinex.OrderUpdate)

			pair := domain.NewCurrencyPairFrom2Currencies(domain.Currency(o.Symbol[1:4]), domain.Currency(o.Symbol[4:]))
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
				Type: domain.OrderEventTypeUpdate,
				Order: domain.Order{
					OrderNumber:           strconv.Itoa(int(o.ID)),
					Direction:             direction,
					Context:               oContext,
					Type:                  oType,
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
		case *bitfinex.OrderNew:

			o := obj.(bitfinex.OrderNew)

			pair := domain.NewCurrencyPairFrom2Currencies(domain.Currency(o.Symbol[1:4]), domain.Currency(o.Symbol[4:]))
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
				Type: domain.OrderEventTypeNew,
				Order: domain.Order{
					OrderNumber:           strconv.Itoa(int(o.ID)),
					Direction:             direction,
					Context:               oContext,
					Type:                  oType,
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
		case *bitfinex.OrderSnapshot:

			for _, o := range v.Snapshot {
				pair := domain.NewCurrencyPairFrom2Currencies(domain.Currency(o.Symbol[1:4]), domain.Currency(o.Symbol[4:]))
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
		//case *bitfinex.Trade:
		case *bitfinex.Ticker:

			pair := domain.NewCurrencyPairFrom2Currencies(domain.Currency(v.Symbol[0:3]), domain.Currency(v.Symbol[3:]))

			c.BaseExchangeClient.TickersChannel <- &domain.TickerEvent{
				Ticker: domain.Ticker{
					Bid:                util.CurrencyFloat64ToInt64(v.Bid, pair.GetCur1().String()),
					Ask:                util.CurrencyFloat64ToInt64(v.Ask, pair.GetCur1().String()),
					Last:               util.CurrencyFloat64ToInt64(v.LastPrice, pair.GetCur1().String()),
					Volume:             util.CurrencyFloat64ToInt64(v.Volume, pair.GetCur1().String()),
					DailyChange:        util.CurrencyFloat64ToInt64(v.DailyChange, pair.GetCur1().String()),
					DailyChangePercent: v.DailyChangePerc,
					High:               util.CurrencyFloat64ToInt64(v.High, pair.GetCur1().String()),
					Low:                util.CurrencyFloat64ToInt64(v.Low, pair.GetCur1().String()),
					Pair:               pair,
					UpdatedAt:          time.Now(),
					Exchange:           c.GetName(),
				},
			}
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
