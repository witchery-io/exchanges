package common

import (
	"errors"
	"github.com/witchery-io/go-exchanges/pkg/exchange"
	"github.com/witchery-io/go-exchanges/pkg/exchange/bitfinex"
)

func NewMultipleExchangeEventSubscriber() {

}

func NewExchangeClientFromName(name string) (exchange.Client, error) {
	switch name {
	case "bitfinex":
		return bitfinex.New(), nil
	default:
		return nil, errors.New("exchange not found")
	}
}