package binance

import (
	"context"
	"fmt"
	"time"

	adp "github.com/matrixbotio/exchange-gates-lib/internal/adapters"
	"github.com/matrixbotio/exchange-gates-lib/internal/adapters/binance/helpers/errs"
	"github.com/matrixbotio/exchange-gates-lib/internal/adapters/binance/workers"
	"github.com/matrixbotio/exchange-gates-lib/internal/adapters/binance/wrapper"
	"github.com/matrixbotio/exchange-gates-lib/internal/consts"
	iWorkers "github.com/matrixbotio/exchange-gates-lib/internal/workers"
	pkgStructs "github.com/matrixbotio/exchange-gates-lib/pkg/structs"
)

const (
	adapterName = "Binance Spot"
	adapterTag  = "binance-spot"
)

type adapter struct {
	ExchangeID int
	Name       string
	Tag        string

	binanceAPI wrapper.BinanceAPIWrapper
}

func New(wrapper wrapper.BinanceAPIWrapper) adp.Adapter {
	return &adapter{
		ExchangeID: consts.ExchangeIDbinanceSpot,
		Name:       adapterName,
		Tag:        adapterTag,
		binanceAPI: wrapper,
	}
}

func (a *adapter) GetTag() string {
	return a.Tag
}

func (a *adapter) GetID() int {
	return a.ExchangeID
}

func (a *adapter) GetName() string {
	return a.Name
}

func (a *adapter) GetLimits() pkgStructs.ExchangeLimits {
	return pkgStructs.ExchangeLimits{
		MaxConnectionsPerBatch:   299,
		MaxConnectionsInDuration: 5 * time.Minute,
	}
}

func (a *adapter) Connect(credentials pkgStructs.APICredentials) error {
	if credentials.Type != pkgStructs.APICredentialsTypeKeypair {
		return errs.ErrInvalidCredentials
	}

	if err := a.binanceAPI.Connect(
		context.Background(),
		credentials.Keypair.Public,
		credentials.Keypair.Secret,
	); err != nil {
		return fmt.Errorf("binance adapter: connect: %w", err)
	}

	a.binanceAPI.Sync(context.Background())
	return nil
}

func (a *adapter) GetPriceWorker(callback iWorkers.PriceEventCallback) iWorkers.IPriceWorker {
	return workers.NewPriceWorker(
		a.GetTag(),
		a.binanceAPI,
		callback,
	)
}

func (a *adapter) GetTradeEventsWorker() iWorkers.ITradeEventWorker {
	return workers.NewTradeEventsWorker(
		a.GetTag(),
		a.binanceAPI,
	)
}
