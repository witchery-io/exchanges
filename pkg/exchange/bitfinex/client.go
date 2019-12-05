package bitfinex

// @todo split file

import (
	"context"
	"fmt"
	"github.com/Rhymond/go-money"
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
		case *websocket.ErrorEvent:
			fmt.Println(v)
		case error:
			fmt.Println(v)
			c.BaseExchangeClient.ErrorsChannel <- v
		case *bitfinex.PositionSnapshot:

			for _, p := range v.Snapshot {

				pair := domain.NewCurrencyPairFromString(p.Symbol[1:])

				status := domain.PositionStatusActive
				if p.Status == bitfinex.PositionStatusClosed {
					status = domain.PositionStatusClosed
				}

				c.BaseExchangeClient.PositionsChannel <- &domain.PositionEvent{
					Type: domain.PositionEventTypeSnapshot,
					Position: domain.Position{
						Id:                   strconv.Itoa(int(p.Id)),
						Pair:                 pair,
						Status:               status,
						Amount:               util.CurrencyFloat64ToInt64(p.Amount, pair.Currency1.String()),
						BasePrice:            util.CurrencyFloat64ToInt64(p.BasePrice, pair.Currency2.String()),
						MarginFunding:        p.MarginFunding,
						MarginFundingType:    p.MarginFundingType,
						ProfitLoss:           util.CurrencyFloat64ToInt64(p.ProfitLoss, pair.Currency2.String()),
						ProfitLossPercentage: p.ProfitLossPercentage,
						LiquidationPrice:     util.CurrencyFloat64ToInt64(p.LiquidationPrice, pair.Currency2.String()),
						Exchange:             c.GetName(),
						AccountId:            c.AccountId,
					},
				}

			}

		case *bitfinex.PositionNew:

			p := obj.(*bitfinex.PositionNew)

			pair := domain.NewCurrencyPairFromString(p.Symbol[1:])

			status := domain.PositionStatusActive
			if p.Status == bitfinex.PositionStatusClosed {
				status = domain.PositionStatusClosed
			}

			c.BaseExchangeClient.PositionsChannel <- &domain.PositionEvent{
				Type: domain.PositionEventTypeNew,
				Position: domain.Position{
					Id:                   strconv.Itoa(int(p.Id)),
					Pair:                 pair,
					Status:               status,
					Amount:               util.CurrencyFloat64ToInt64(p.Amount, pair.Currency1.String()),
					BasePrice:            util.CurrencyFloat64ToInt64(p.BasePrice, pair.Currency2.String()),
					MarginFunding:        p.MarginFunding,
					MarginFundingType:    p.MarginFundingType,
					ProfitLoss:           util.CurrencyFloat64ToInt64(p.ProfitLoss, pair.Currency2.String()),
					ProfitLossPercentage: p.ProfitLossPercentage,
					LiquidationPrice:     util.CurrencyFloat64ToInt64(p.LiquidationPrice, pair.Currency2.String()),
					Exchange:             c.GetName(),
					AccountId:            c.AccountId,
				},
			}
		case *bitfinex.PositionUpdate:

			p := obj.(*bitfinex.PositionUpdate)

			pair := domain.NewCurrencyPairFromString(p.Symbol[1:])

			status := domain.PositionStatusActive
			if p.Status == bitfinex.PositionStatusClosed {
				status = domain.PositionStatusClosed
			}

			c.BaseExchangeClient.PositionsChannel <- &domain.PositionEvent{
				Type: domain.PositionEventTypeUpdate,
				Position: domain.Position{
					Id:                   strconv.Itoa(int(p.Id)),
					Pair:                 pair,
					Status:               status,
					Amount:               util.CurrencyFloat64ToInt64(p.Amount, pair.Currency1.String()),
					BasePrice:            util.CurrencyFloat64ToInt64(p.BasePrice, pair.Currency2.String()),
					MarginFunding:        p.MarginFunding,
					MarginFundingType:    p.MarginFundingType,
					ProfitLoss:           util.CurrencyFloat64ToInt64(p.ProfitLoss, pair.Currency2.String()),
					ProfitLossPercentage: p.ProfitLossPercentage,
					LiquidationPrice:     util.CurrencyFloat64ToInt64(p.LiquidationPrice, pair.Currency2.String()),
					Exchange:             c.GetName(),
					AccountId:            c.AccountId,
				},
			}
		case *bitfinex.BookUpdate:

			pair := domain.NewCurrencyPairFromString(v.Symbol)

			direction := domain.OrderDirectionBuy
			if v.Side == bitfinex.Short {
				direction = domain.OrderDirectionSell
			}

			c.BaseExchangeClient.OrderBooksChannel <- &domain.OrderBookEvent{
				OrderBook: domain.OrderBook{
					Pair:   pair,
					Price:  util.CurrencyFloat64ToInt64(v.Price, pair.Currency1.String()),
					Count:  v.Count,
					Amount: util.CurrencyFloat64ToInt64(v.Amount, pair.Currency1.String()),
					Side:   direction,
				},
			}
		case *bitfinex.OrderUpdate:

			o := obj.(bitfinex.OrderUpdate)

			pair := domain.NewCurrencyPairFromString(o.Symbol[1:])
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
					OriginalAmount:        util.CurrencyFloat64ToInt64(o.AmountOrig, pair.Currency1.String()),
					RemainingAmount:       util.CurrencyFloat64ToInt64(o.Amount, pair.Currency1.String()),
					Price:                 util.CurrencyFloat64ToInt64(o.Price, pair.Currency2.String()),
					AverageExecutionPrice: util.CurrencyFloat64ToInt64(o.PriceAvg, pair.Currency2.String()),
					OpenedAt:              time.Unix(o.MTSCreated, 0),
					UpdatedAt:             time.Unix(o.MTSUpdated, 0),
					AccountId:             c.AccountId,
					Status:                statuses[o.Status],
					Exchange:              c.GetName(),
				},
			}
		case *bitfinex.OrderNew:

			o := obj.(bitfinex.OrderNew)

			pair := domain.NewCurrencyPairFromString(o.Symbol[1:])
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
					OriginalAmount:        util.CurrencyFloat64ToInt64(o.AmountOrig, pair.Currency1.String()),
					RemainingAmount:       util.CurrencyFloat64ToInt64(o.Amount, pair.Currency1.String()),
					Price:                 util.CurrencyFloat64ToInt64(o.Price, pair.Currency2.String()),
					AverageExecutionPrice: util.CurrencyFloat64ToInt64(o.PriceAvg, pair.Currency2.String()),
					OpenedAt:              time.Unix(o.MTSCreated, 0),
					UpdatedAt:             time.Unix(o.MTSUpdated, 0),
					AccountId:             c.AccountId,
					Status:                statuses[o.Status],
					Exchange:              c.GetName(),
				},
			}
		case *bitfinex.OrderSnapshot:

			for _, o := range v.Snapshot {
				pair := domain.NewCurrencyPairFromString(o.Symbol[1:])
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
						AccountId:             c.AccountId,
						Status:                statuses[o.Status],
						Exchange:              c.GetName(),
					},
				}
			}
		//case *bitfinex.Trade:
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

func (c *client) Authenticate(accountId string, credentials map[string]string) error {

	c.AccountId = accountId
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
		Amount: money.New(amount, order.Pair.Currency1.String()).AsMajorUnits(),
		Price:  money.New(order.Price, order.Pair.Currency2.String()).AsMajorUnits(),
	}

	err := c.wsClient.SubmitOrder(ctx, o)
	if err != nil {
		return err
	}

	return nil

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
			Name:              "bitfinex",
			OrdersChannel:     make(chan *domain.OrderEvent, 0),
			TradesChannel:     make(chan *domain.TradeEvent, 0),
			TickersChannel:    make(chan *domain.TickerEvent, 0),
			OrderBooksChannel: make(chan *domain.OrderBookEvent, 0),
			PositionsChannel:  make(chan *domain.PositionEvent, 0),
			ErrorsChannel:     make(chan error, 0),
		},
		wsClient: nil,
	}

	p := websocket.NewDefaultParameters()
	p.ManageOrderbook = true
	c.wsClient = websocket.NewWithParams(p)

	return c
}
