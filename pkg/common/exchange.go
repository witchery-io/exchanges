// Package common for helpers
package common

import (
	"errors"

	"github.com/witchery-io/go-exchanges/pkg/exchange"
	"github.com/witchery-io/go-exchanges/pkg/exchange/bitfinex"
)

// NewExchangeClientFromName helper function to init exchange client from name
// Example usage
// ```
// NewExchangeClientFromName("bitfinex")
// ```
func NewExchangeClientFromName(name string, options exchange.ClientOptions) (exchange.Client, error) {
	switch name {
	case "bitfinex":
		return bitfinex.New(options), nil
	default:
		return nil, errors.New("exchange not found")
	}
}
