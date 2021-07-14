package matrixgates

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/adshao/go-binance/v2"
	"github.com/go-stack/stack"
	"github.com/matrixbotio/exchange-gates/workers"
)

const (
	pairDefaultMinQty    = 0.001
	pairDefaultMaxQty    = 99999.99
	pairDefaultMinPrice  = 0.000001
	pairDefaultQtyStep   = 0.001
	pairDefaultPriceStep = 0.000001

	candlesInterval = "1m"
)

//BinanceSpotAdapter - bot exchange adapter for BinanceSpot
type BinanceSpotAdapter struct {
	ExchangeAdapter
	binanceAPI *binance.Client
}

//NewBinanceSpotAdapter - create binance exchange adapter
func NewBinanceSpotAdapter() *BinanceSpotAdapter {
	stack.Caller(0)
	a := BinanceSpotAdapter{}
	a.Name = "Binance Spot"
	a.Tag = "binance-spot"
	return &a
}

//Connect to exchange
func (a *BinanceSpotAdapter) Connect(credentials APICredentials) error {
	switch credentials.Type {
	default:
		return errors.New("invalid credentials to connect to Binance")
	case APICredentialsTypeKeypair:
		a.binanceAPI = binance.NewClient(credentials.Keypair.Public, credentials.Keypair.Secret)
	}
	return a.ping()
}

//GetOrderData ..
func (a *BinanceSpotAdapter) GetOrderData(pairSymbol string, orderID int64) (*TradeEventData, error) {
	tradeData := TradeEventData{
		OrderID: orderID,
	}
	//order status: NEW, PARTIALLY_FILLED, FILLED, CANCELED, PENDING_CANCEL, REJECTED, EXPIRED
	orderResponse, err := a.binanceAPI.NewGetOrderService().Symbol(pairSymbol).
		OrderID(orderID).Do(context.Background())

	if err != nil {
		if strings.Contains(err.Error(), "Order does not exist") {
			log.Println("[DEBUG] CHECK. order " + strconv.FormatInt(orderID, 10) + " doesn't exists")
			tradeData.Status = "UNKNOWN"
			return &tradeData, nil
		}
		return nil, errors.New("service request failed: " + err.Error() + GetTrace())
	}
	orderFilledQty, convErr := strconv.ParseFloat(orderResponse.ExecutedQuantity, 64)
	if convErr != nil {
		return nil, errors.New("data handle error: failed to parse order filled qty: " + convErr.Error() + ", stack: " + GetTrace())
	}
	tradeData.OrderAwaitQty = orderFilledQty
	tradeData.Status = string(orderResponse.Status)
	return &tradeData, nil
}

//PlaceOrder - place order on exchange
func (a *BinanceSpotAdapter) PlaceOrder(order BotOrder, pairLimits ExchangePairData) (*CreateOrderResponse, error) {
	var orderSide binance.SideType
	{
		//move this block to another location?
		switch order.Type {
		default:
			return nil, errors.New("data invalid error: unknown strategy given for order, stack: " + GetTrace())
		case "buy":
			orderSide = binance.SideTypeBuy
		case "sell":
			orderSide = binance.SideTypeSell
		}
	}

	var quantityPrecision int = GetFloatPrecision(pairLimits.QtyStep)
	var quantityStr string = strconv.FormatFloat(order.Qty, 'f', quantityPrecision, 64)
	var ratePrecision int = GetFloatPrecision(pairLimits.PriceStep)
	var rateStr string = strconv.FormatFloat(order.Price, 'f', ratePrecision, 32)

	//check lot size
	if order.Qty < pairLimits.MinQty {
		fmt.Print("pair min qty: ", pairLimits.MinQty, ", order qty: ")
		fmt.Println(order.Qty)
		return nil, errors.New("bot order invalid error: insufficient amount to open an order in this pair, stack: " + GetTrace())
	}
	if order.Qty > pairLimits.MaxQty {
		return nil, errors.New("bot order invalid error: too much amount to open an order in this pair, stack: " + GetTrace())
	}
	if order.Price < pairLimits.MinPrice {
		return nil, errors.New("bot order invalid error: insufficient price to open an order in this pair, stack: " + GetTrace())
	}

	orderRes, err := a.binanceAPI.NewCreateOrderService().Symbol(order.PairSymbol).
		Side(orderSide).Type(binance.OrderTypeLimit).
		TimeInForce(binance.TimeInForceTypeGTC).Quantity(quantityStr).
		Price(rateStr).Do(context.Background())
	if err != nil {
		return nil, errors.New("service request failed: failed to create order, " + err.Error() + ", stack: " + GetTrace())
	}

	//parse qty & price from order response
	orderResOrigQty, convErr := strconv.ParseFloat(orderRes.OrigQuantity, 64)
	if convErr != nil {
		return nil, errors.New("data handle error: failed to parse order origQty, " + convErr.Error() + ", stack: " + GetTrace())
	}
	orderResPrice, convErr := strconv.ParseFloat(orderRes.Price, 64)
	if convErr != nil {
		return nil, errors.New("data handle error: failed to parse order price, " + convErr.Error() + ", stack: " + GetTrace())
	}

	return &CreateOrderResponse{
		OrderID:       orderRes.OrderID,
		ClientOrderID: orderRes.ClientOrderID,
		OrigQuantity:  orderResOrigQty,
		Price:         orderResPrice,
	}, nil
}

//GetAccountData - get account data ^ↀᴥↀ^
func (a *BinanceSpotAdapter) GetAccountData() (*AccountData, error) {
	binanceAccountData, clientErr := a.binanceAPI.NewGetAccountService().Do(context.Background())
	if clientErr != nil {
		return nil, errors.New("data invalid error: failed to send request to trade, " + clientErr.Error() + ", stack: " + GetTrace())
	}
	accountDataResult := AccountData{
		CanTrade: binanceAccountData.CanTrade,
	}
	balances := []Balance{}
	for _, binanceBalanceData := range binanceAccountData.Balances {
		//convert strings to float64
		balanceFree, convErr := strconv.ParseFloat(binanceBalanceData.Free, 64)
		if convErr != nil {
			balanceFree = 0
			//log.Println("failed to parse free balance: " + convErr.Error())
		}
		balanceLocked, convErr := strconv.ParseFloat(binanceBalanceData.Locked, 64)
		if convErr != nil {
			balanceLocked = 0
			//log.Println("failed to parse locked balance: " + convErr.Error())
		}
		balances = append(balances, Balance{
			Asset:  binanceBalanceData.Asset,
			Free:   balanceFree,
			Locked: balanceLocked,
		})
	}
	accountDataResult.Balances = balances
	return &accountDataResult, nil
}

//GetPairLastPrice - get pair last price ^ↀᴥↀ^
func (a *BinanceSpotAdapter) GetPairLastPrice(pairSymbol string) (float64, error) {
	tickerService := a.binanceAPI.NewListPricesService()
	prices, srvErr := tickerService.Symbol(pairSymbol).Do(context.Background())
	if srvErr != nil {
		return 0, errors.New("service request failed: failed to request last price, " +
			srvErr.Error() + ", stack: " + GetTrace())
	}
	//until just brute force. need to be done faster
	var price float64 = 0
	var parseErr error
	for _, p := range prices {
		if p.Symbol == pairSymbol {
			price, parseErr = strconv.ParseFloat(p.Price, 64)
			if parseErr != nil {
				return 0, errors.New("data handle error: failed to parse " + p.Price +
					" as float, stack: " + GetTrace())
			}
			break
		}
	}
	return price, nil
}

//CancelPairOrder - cancel one exchange pair order by ID
func (a *BinanceSpotAdapter) CancelPairOrder(pairSymbol string, orderID int64) error {
	_, clientErr := a.binanceAPI.NewCancelOrderService().Symbol(pairSymbol).
		OrderID(orderID).Do(context.Background())
	if clientErr != nil {
		return errors.New("service request failed: " + clientErr.Error() +
			", stack: " + GetTrace())
	}
	return nil
}

//CancelPairOrders - cancel pair all orders
func (a *BinanceSpotAdapter) CancelPairOrders(pairSymbol string) error {
	_, clientErr := a.binanceAPI.NewCancelOpenOrdersService().
		Symbol(pairSymbol).Do(context.Background())
	if clientErr != nil {
		//log.Println("failed to cancel all orders, " + clientErr.Error())
		//log.Println("let's try cancel orders manualy..")
		//handle error
		if strings.Contains(clientErr.Error(), "Unknown order sent") {
			/*canceling all orders failed,
			let's try to request a list of them and cancel them individually*/
			orders, err := a.GetPairOpenOrders(pairSymbol)
			if err != nil {
				// =(
				//log.Println("[DEBUG] error while b.getOpenOrders(): " + err.Error())
				return err
			}
			if len(orders) == 0 {
				//orders already cancelled
				return nil
			}
			for _, order := range orders {
				err := a.CancelPairOrder(pairSymbol, order.OrderID)
				if err != nil {
					return err
				}
			}
			return nil
		}
		//log.Println("[DEBUG] service error: " + clientErr.Error())
		return errors.New("service request failed: " + clientErr.Error() +
			", stack: " + GetTrace())
	}
	return nil
}

// GetPairData - get pair data & limits
func (a *BinanceSpotAdapter) GetPairData(pairSymbol string) (*ExchangePairData, error) {
	exchangeInfo, err := a.binanceAPI.NewExchangeInfoService().Symbol(pairSymbol).Do(context.Background())
	if err != nil {
		return nil, err
	}

	// find pairSymbol
	for _, symbolData := range exchangeInfo.Symbols {
		return a.getExchangePairData(symbolData)
	}

	return nil, errors.New("data for " + pairSymbol + " pair not found")
}

//GetPairOpenOrders - get open orders array
func (a *BinanceSpotAdapter) GetPairOpenOrders(pairSymbol string) ([]*Order, error) {
	ordersRaw, err := a.binanceAPI.NewListOpenOrdersService().Symbol(pairSymbol).Do(context.Background())
	if err != nil {
		return nil, err
	}

	orders := []*Order{}
	for _, orderRaw := range ordersRaw {
		orders = append(orders, &Order{
			OrderID:       orderRaw.OrderID,
			ClientOrderID: orderRaw.ClientOrderID,
			Status:        string(orderRaw.Status),
		})
	}

	return orders, nil
}

func (a *BinanceSpotAdapter) ping() error {
	err := a.binanceAPI.NewPingService().Do(context.Background())
	if err != nil {
		return errors.New("service disconnected: failed to connect to the exchange, please try again later, stack: " + GetTrace())
	}
	return nil
}

// VerifyAPIKeys - create new exchange client & attempt to get account data
func (a *BinanceSpotAdapter) VerifyAPIKeys(keyPublic, keySecret string) error {
	newClient := binance.NewClient(keyPublic, keySecret)
	accountService, err := newClient.NewGetAccountService().Do(context.Background())
	if err != nil {
		return errors.New("invalid api key")
	}
	if !accountService.CanTrade {
		return errors.New("service no access: Your API key does not have permission to trade, change its restrictions")
	}
	return nil
}

// convert binance.Symbol to ExchangePairData
func (a *BinanceSpotAdapter) getExchangePairData(symbolData binance.Symbol) (*ExchangePairData, error) {
	pairData := ExchangePairData{
		ExchangeID:     a.ExchangeID,
		BaseAsset:      symbolData.BaseAsset,
		BasePrecision:  symbolData.BaseAssetPrecision,
		QuoteAsset:     symbolData.QuoteAsset,
		QuotePrecision: symbolData.QuotePrecision,
		Status:         symbolData.Status,
		Symbol:         symbolData.Symbol,
		MinQty:         pairDefaultMinQty,
		MaxQty:         pairDefaultMaxQty,
		MinPrice:       pairDefaultMinPrice,
		QtyStep:        pairDefaultQtyStep,
		PriceStep:      pairDefaultPriceStep,
		AllowedMargin:  symbolData.IsMarginTradingAllowed,
		AllowedSpot:    symbolData.IsSpotTradingAllowed,
	}

	err := binanceParseLotSizeFilter(&symbolData, &pairData)
	if err != nil {
		return nil, err
	}

	err = binanceParsePriceFilter(&symbolData, &pairData)
	if err != nil {
		return nil, err
	}

	return &pairData, nil
}

func binanceParsePriceFilter(symbolData *binance.Symbol, pairData *ExchangePairData) error {
	var err error
	priceFilter := symbolData.PriceFilter()
	if priceFilter == nil {
		return errors.New("failed to get price filter for symbol data")
	}
	minPriceRaw := priceFilter.MinPrice
	pairData.MinPrice, err = strconv.ParseFloat(minPriceRaw, 64)
	if err != nil {
		return errors.New("data handle error: " + err.Error())
	}
	if pairData.MinPrice == 0 {
		pairData.MinPrice = pairDefaultMinPrice
	}
	priceStepRaw := symbolData.PriceFilter().TickSize
	pairData.PriceStep, err = strconv.ParseFloat(priceStepRaw, 64)
	if err != nil {
		return errors.New("data handle error: " + err.Error())
	}
	if pairData.PriceStep == 0 {
		pairData.PriceStep = pairData.MinPrice
	}
	return nil
}

func binanceParseLotSizeFilter(symbolData *binance.Symbol, pairData *ExchangePairData) error {
	marketLotSizeFilter := symbolData.MarketLotSizeFilter()
	if marketLotSizeFilter == nil {
		return errors.New("failed to get market lot size filter for symbol data")
	}
	minQtyRaw := marketLotSizeFilter.MinQuantity
	maxQtyRaw := marketLotSizeFilter.MaxQuantity
	minQty, err := strconv.ParseFloat(minQtyRaw, 64)
	if err != nil {
		return errors.New("data handle error: " + err.Error())
	}
	if minQty == 0 {
		minQty = pairDefaultMinQty
	}

	pairData.MaxQty, err = strconv.ParseFloat(maxQtyRaw, 64)
	if err != nil {
		return errors.New("data handle error" + err.Error())
	}

	qtyStepRaw := symbolData.MarketLotSizeFilter().StepSize
	pairData.QtyStep, err = strconv.ParseFloat(qtyStepRaw, 64)
	if err != nil {
		return errors.New("data handle error: " + err.Error())
	}
	if pairData.QtyStep == 0 {
		pairData.QtyStep = minQty
	}
	return nil
}

// GetPairs get all Binance pairs
func (a *BinanceSpotAdapter) GetPairs() ([]*ExchangePairData, error) {
	service := a.binanceAPI.NewExchangeInfoService()
	res, err := service.Do(context.Background())
	if err != nil {
		return nil, errors.New("service disconnected: error while connecting to ExchangeInfoService: " + err.Error() + ", stack: " + GetTrace())
	}

	pairs := []*ExchangePairData{}
	for _, symbolData := range res.Symbols {
		pairData, err := a.getExchangePairData(symbolData)
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, pairData)
	}
	return pairs, nil
}

// GetPairBalance - get pair balance: ticker, quote asset balance for pair symbol
func (a *BinanceSpotAdapter) GetPairBalance(pair PairSymbolData) (*PairBalance, error) {
	accountData, err := a.GetAccountData()
	if err != nil {
		return nil, err
	}

	pairBalanceData := &PairBalance{}
	for _, balanceData := range accountData.Balances {
		if balanceData.Asset == pair.BaseTicker || balanceData.Asset == pair.QuoteTicker {
			assetBalanceData := &AssetBalance{
				Ticker: balanceData.Asset,
				Free:   balanceData.Free,
				Locked: balanceData.Locked,
			}
			if balanceData.Asset == pair.BaseTicker {
				// base asset found
				pairBalanceData.BaseAsset = assetBalanceData
			}
			if balanceData.Asset == pair.QuoteTicker {
				// quote asset found
				pairBalanceData.QuoteAsset = assetBalanceData
			}
			if pairBalanceData.BaseAsset != nil && pairBalanceData.QuoteAsset != nil {
				// found
				break
			}
		}
	}
	if pairBalanceData.BaseAsset == nil {
		pairBalanceData.BaseAsset = &AssetBalance{
			Ticker: pair.BaseTicker,
			Free:   0,
			Locked: 0,
		}
	}
	if pairBalanceData.QuoteAsset == nil {
		pairBalanceData.QuoteAsset = &AssetBalance{
			Ticker: pair.QuoteTicker,
			Free:   0,
			Locked: 0,
		}
	}
	return pairBalanceData, nil
}

/*
                    _
                   | |
__      _____  _ __| | _____ _ __ ___
\ \ /\ / / _ \| '__| |/ / _ \ '__/ __|
 \ V  V / (_) | |  |   <  __/ |  \__ \
  \_/\_/ \___/|_|  |_|\_\___|_|  |___/

*/

//PriceWorkerBinance - MarketDataWorker for binance
type PriceWorkerBinance struct {
	workers.PriceWorker
}

//GetPriceWorker - create new market data worker
func (a *BinanceSpotAdapter) GetPriceWorker() workers.IPriceWorker {
	w := PriceWorkerBinance{}
	w.PriceWorker.ExchangeTag = a.Tag
	return &w
}

//SubscribeToPriceEvents - websocket subscription to change quotes and ask-, bid-qty on the exchange
func (w *PriceWorkerBinance) SubscribeToPriceEvents(
	eventCallback func(event workers.PriceEvent),
	errorHandler func(err error),
) error {
	wsBookHandler := func(event *binance.WsBookTickerEvent) {
		if event != nil {
			eventAsk, convErr := strconv.ParseFloat(event.BestAskPrice, 64)
			if convErr != nil {
				// ignore event
				log.Println(convErr)
				return
			}
			eventBid, convErr := strconv.ParseFloat(event.BestBidPrice, 64)
			if convErr != nil {
				// ignore event
				log.Println(convErr)
				return
			}
			wEvent := workers.PriceEvent{
				Symbol: event.Symbol,
				Ask:    eventAsk,
				Bid:    eventBid,
			}
			eventCallback(wEvent)
		}
	}
	wsErrHandler := func(err error) {
		errorHandler(errors.New("service request failed: " + err.Error()))
	}
	var openWsErr error
	w.WsChannels = new(workers.WorkerChannels)
	w.WsChannels.WsDone, w.WsChannels.WsStop, openWsErr = binance.WsAllBookTickerServe(wsBookHandler, wsErrHandler)
	if openWsErr != nil {
		return errors.New("service request failed: " + openWsErr.Error())
	}
	return nil
}

//CandleWorkerBinance - MarketDataWorker for binance
type CandleWorkerBinance struct {
	workers.CandleWorker
}

//GetCandleWorker - create new market candle worker
func (a *BinanceSpotAdapter) GetCandleWorker() workers.ICandleWorker {
	w := CandleWorkerBinance{}
	w.ExchangeTag = a.GetTag()
	return &w
}

//SubscribeToCandleEvents - websocket subscription to change trade candles on the exchange
func (w *CandleWorkerBinance) SubscribeToCandleEvents(
	pairSymbols []string,
	eventCallback func(event workers.CandleEvent),
	errorHandler func(err error),
) error {
	symbolIntervalsMap := make(map[string]string)
	for _, symbol := range pairSymbols {
		symbolIntervalsMap[symbol] = candlesInterval
	}

	wsCandleHandler := func(event *binance.WsKlineEvent) {
		if event != nil {
			// fix endTime
			if strings.HasSuffix(strconv.FormatInt(event.Kline.EndTime, 10), "999") {
				event.Kline.EndTime++
			}

			wEvent := workers.CandleEvent{
				Symbol: event.Symbol,
				Candle: workers.CandleData{
					StartTime: event.Kline.StartTime,
					EndTime:   event.Kline.EndTime,
					Interval:  event.Kline.Interval,
				},
			}

			errs := make([]error, 5)
			wEvent.Candle.Open, errs[0] = strconv.ParseFloat(event.Kline.Open, 64)
			wEvent.Candle.Close, errs[1] = strconv.ParseFloat(event.Kline.Close, 64)
			wEvent.Candle.High, errs[2] = strconv.ParseFloat(event.Kline.High, 64)
			wEvent.Candle.Low, errs[3] = strconv.ParseFloat(event.Kline.Low, 64)
			wEvent.Candle.Volume, errs[4] = strconv.ParseFloat(event.Kline.Volume, 64)
			if LogNotNilError(errs) {
				return
			}

			eventCallback(wEvent)
		}
	}
	wsErrHandler := func(err error) {
		errorHandler(errors.New("service request failed: " + err.Error()))
	}
	var openWsErr error
	w.WsChannels = new(workers.WorkerChannels)
	w.WsChannels.WsDone, w.WsChannels.WsStop, openWsErr = binance.WsCombinedKlineServe(symbolIntervalsMap, wsCandleHandler, wsErrHandler)
	if openWsErr != nil {
		return errors.New("service request failed: " + openWsErr.Error())
	}
	return nil
}

//TradeEventWorkerBinance - TradeEventWorker for binance
type TradeEventWorkerBinance struct {
	workers.TradeEventWorker
}

//GetTradeEventsWorker - create new market candle worker
func (a *BinanceSpotAdapter) GetTradeEventsWorker() workers.ITradeEventWorker {
	w := TradeEventWorkerBinance{}
	w.ExchangeTag = a.GetTag()
	return &w
}

//SubscribeToTradeEvents - websocket subscription to change trade candles on the exchange
func (w *TradeEventWorkerBinance) SubscribeToTradeEvents(
	symbol string,
	eventCallback func(event workers.TradeEvent),
	errorHandler func(err error),
) error {

	wsErrHandler := func(err error) {
		errorHandler(errors.New("service request failed: " + err.Error()))
	}

	wsTradeHandler := func(event *binance.WsTradeEvent) {
		if event != nil {
			// fix event.Time
			if strings.HasSuffix(strconv.FormatInt(event.Time, 10), "999") {
				event.Time++
			}
			wEvent := workers.TradeEvent{
				Time:          event.Time,
				Symbol:        event.Symbol,
				BuyerOrderID:  event.BuyerOrderID,
				SellerOrderID: event.SellerOrderID,
			}
			errs := make([]error, 2)
			wEvent.Price, errs[0] = strconv.ParseFloat(event.Price, 64)
			wEvent.Quantity, errs[0] = strconv.ParseFloat(event.Quantity, 64)
			if LogNotNilError(errs) {
				return
			}
			eventCallback(wEvent)
		}
	}

	var openWsErr error
	w.WsChannels = new(workers.WorkerChannels)
	w.WsChannels.WsDone, w.WsChannels.WsStop, openWsErr = binance.WsTradeServe(symbol, wsTradeHandler, wsErrHandler)
	if openWsErr != nil {
		return errors.New("failed to subscribe to trade events: " + openWsErr.Error())
	}
	return nil
}
