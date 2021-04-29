package matrixgates

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"./workers"
	"github.com/adshao/go-binance/v2"
	sharederrs "github.com/matrixbotio/shared-errors"
)

//BinanceSpotAdapter - bot exchange adapter for BinanceSpot
type BinanceSpotAdapter struct {
	ExchangeAdapter
	binanceAPI *binance.Client
}

//NewBinanceSpotAdapter - create binance exchange adapter
func NewBinanceSpotAdapter() *BinanceSpotAdapter {
	return &BinanceSpotAdapter{}
}

//Connect to exchange
func (a *BinanceSpotAdapter) Connect(credentials APICredentials) *sharederrs.APIError {
	switch credentials.Type {
	default:
		return sharederrs.DataInvalidErr
	case APICredentialsTypeKeypair:
		a.binanceAPI = binance.NewClient(credentials.Keypair.Public, credentials.Keypair.Secret)
		break
	}
	return a.ping()
}

//GetOrderData ..
func (a *BinanceSpotAdapter) GetOrderData(pairSymbol string, orderID int64) (*TradeEventData, *sharederrs.APIError) {
	tradeData := TradeEventData{
		OrderID: orderID,
	}
	//order status: NEW, PARTIALLY_FILLED, FILLED, CANCELED, PENDING_CANCEL, REJECTED, EXPIRED
	orderResponse, err := a.binanceAPI.NewGetOrderService().Symbol(pairSymbol).
		OrderID(orderID).Do(context.Background())

	if err != nil {
		if strings.Contains(err.Error(), "Order does not exist") {
			fmt.Println("[CHECK] order " + strconv.FormatInt(orderID, 10) + " doesn't exists")
			tradeData.Status = "UNKNOWN"
			return &tradeData, nil
		}
		return nil, sharederrs.ServiceReqFailedErr.SetMessage(err.Error()).SetTrace()
	}
	orderFilledQty, convErr := strconv.ParseFloat(orderResponse.ExecutedQuantity, 64)
	if convErr != nil {
		return nil, sharederrs.DataHandleErr.M("failed to parse order filled qty: " + convErr.Error()).SetTrace()
	}
	tradeData.OrderAwaitQty = orderFilledQty
	tradeData.Status = string(orderResponse.Status)
	return &tradeData, nil
}

//PlaceOrder - place order on exchange
func (a *BinanceSpotAdapter) PlaceOrder(order BotOrder, pairLimits ExchangePairData) (*CreateOrderResponse, *sharederrs.APIError) {
	var orderSide binance.SideType
	{
		//move this block to another location?
		switch order.Type {
		default:
			return nil, sharederrs.DataInvalidErr.M("unknown strategy given for order").SetTrace()
		case "buy":
			orderSide = binance.SideTypeBuy
			break
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
		return nil, sharederrs.BotOrderInvalidErr.
			M("insufficient amount to open an order in this pair").SetTrace()
	}
	if order.Qty > pairLimits.MaxQty {
		return nil, sharederrs.BotOrderInvalidErr.
			M("too much amount to open an order in this pair").SetTrace()
	}
	if order.Price < pairLimits.MinPrice {
		return nil, sharederrs.BotOrderInvalidErr.
			M("insufficient price to open an order in this pair").SetTrace()
	}

	orderRes, err := a.binanceAPI.NewCreateOrderService().Symbol(order.PairSymbol).
		Side(orderSide).Type(binance.OrderTypeLimit).
		TimeInForce(binance.TimeInForceTypeGTC).Quantity(quantityStr).
		Price(rateStr).Do(context.Background())
	if err != nil {
		return nil, sharederrs.ServiceReqFailedErr.M("failed to create order, " + err.Error()).SetTrace()
	}

	//parse qty & price from order response
	orderResOrigQty, convErr := strconv.ParseFloat(orderRes.OrigQuantity, 64)
	if convErr != nil {
		return nil, sharederrs.DataHandleErr.
			M("failed to parse order origQty, " + convErr.Error()).SetTrace()
	}
	orderResPrice, convErr := strconv.ParseFloat(orderRes.Price, 64)
	if convErr != nil {
		return nil, sharederrs.DataHandleErr.
			M("failed to parse order price, " + convErr.Error()).SetTrace()
	}

	return &CreateOrderResponse{
		OrderID:       orderRes.OrderID,
		ClientOrderID: orderRes.ClientOrderID,
		OrigQuantity:  orderResOrigQty,
		Price:         orderResPrice,
	}, nil
}

//GetAccountData - get account data ^ↀᴥↀ^
func (a *BinanceSpotAdapter) GetAccountData() (*AccountData, *sharederrs.APIError) {
	binanceAccountData, clientErr := a.binanceAPI.NewGetAccountService().Do(context.Background())
	if clientErr != nil {
		return nil, sharederrs.DataInvalidErr.
			M("failed to send request to trade, " + clientErr.Error()).SetTrace()
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
func (a *BinanceSpotAdapter) GetPairLastPrice(pairSymbol string) (float64, *sharederrs.APIError) {
	tickerService := a.binanceAPI.NewListPricesService()
	prices, srvErr := tickerService.Symbol(pairSymbol).Do(context.Background())
	if srvErr != nil {
		return 0, sharederrs.ServiceReqFailedErr.
			SetMessage("failed to request last price, " + srvErr.Error()).SetTrace()
	}
	//until just brute force. need to be done faster
	var price float64 = 0
	var parseErr error
	for _, p := range prices {
		if p.Symbol == pairSymbol {
			price, parseErr = strconv.ParseFloat(p.Price, 64)
			if parseErr != nil {
				return 0, sharederrs.DataHandleErr.
					M("failed to parse " + p.Price + " as float").SetTrace()
			}
			break
		}
	}
	return price, nil
}

//CancelPairOrder - cancel one exchange pair order by ID
func (a *BinanceSpotAdapter) CancelPairOrder(pairSymbol string, orderID int64) *sharederrs.APIError {
	_, clientErr := a.binanceAPI.NewCancelOrderService().Symbol(pairSymbol).
		OrderID(orderID).Do(context.Background())
	if clientErr != nil {
		return sharederrs.ServiceReqFailedErr.M(clientErr.Error()).SetTrace()
	}
	return nil
}

//CancelPairOrders - cancel pair all orders
func (a *BinanceSpotAdapter) CancelPairOrders(pairSymbol string) *sharederrs.APIError {
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
				fmt.Println("[DEBUG] error while b.getOpenOrders(): " + err.Message)
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
		return sharederrs.ServiceReqFailedErr.SetMessage(clientErr.Error()).SetTrace()
	}
	return nil
}

//GetPairOpenOrders ..
func (a *BinanceSpotAdapter) GetPairOpenOrders(pairSymbol string) ([]*Order, *sharederrs.APIError) {
	//TODO
	return nil, nil
}

func (a *BinanceSpotAdapter) ping() *sharederrs.APIError {
	err := a.binanceAPI.NewPingService().Do(context.Background())
	if err != nil {
		return sharederrs.ServiceDisconnectedErr.
			SetMessage("failed to connect to the trade, please try again later").SetTrace()
	}
	return nil
}

//VerifyAPIKeys ..
func (a *BinanceSpotAdapter) VerifyAPIKeys(keyPublic, keySecret string) *sharederrs.APIError {
	accountService, err := a.binanceAPI.NewGetAccountService().Do(context.Background())
	if err != nil {
		return sharederrs.ServiceReqFailedErr.
			SetMessage(err.Error()).SetTrace()
	}
	if !accountService.CanTrade {
		return sharederrs.ServiceNoAccess.
			SetMessage("Your API key does not have permission to trade, change its restrictions").SetTrace()
	}
	return a.ping()
}

//GetPairs get all Binance pairs
func (a *BinanceSpotAdapter) GetPairs() ([]*ExchangePairData, *sharederrs.APIError) {
	service := a.binanceAPI.NewExchangeInfoService()
	res, err := service.Do(context.Background())
	if err != nil {
		return nil, sharederrs.ServiceDisconnectedErr.
			M("error while connecting to ExchangeInfoService: " + err.Error()).SetTrace()
	}

	pairs := []*ExchangePairData{}
	for _, symbolData := range res.Symbols {
		pairData := ExchangePairData{
			ExchangeID:     a.ExchangeID,
			BaseAsset:      symbolData.BaseAsset,
			BasePrecision:  symbolData.BaseAssetPrecision,
			QuoteAsset:     symbolData.QuoteAsset,
			QuotePrecision: symbolData.QuotePrecision,
			Status:         symbolData.Status,
			Symbol:         symbolData.Symbol,
			MinQty:         0.001,
			MaxQty:         99999.99,
			MinPrice:       0.000001,
			QtyStep:        0.001,
			PriceStep:      0.000001,
			AllowedMargin:  symbolData.IsMarginTradingAllowed,
			AllowedSpot:    symbolData.IsSpotTradingAllowed,
		}

		marketLotSizeFilter := symbolData.MarketLotSizeFilter()
		if marketLotSizeFilter != nil {
			minQtyRaw := marketLotSizeFilter.MinQuantity
			maxQtyRaw := marketLotSizeFilter.MaxQuantity
			minQty, err := strconv.ParseFloat(minQtyRaw, 64)
			if err != nil {
				return nil, sharederrs.DataHandleErr.SetMessage(err.Error())
			}
			if minQty == 0 {
				minQty = 0.001
			}

			pairData.MaxQty, err = strconv.ParseFloat(maxQtyRaw, 64)
			if err != nil {
				return nil, sharederrs.DataHandleErr.SetMessage(err.Error())
			}

			qtyStepRaw := symbolData.MarketLotSizeFilter().StepSize
			pairData.QtyStep, err = strconv.ParseFloat(qtyStepRaw, 64)
			if err != nil {
				return nil, sharederrs.DataHandleErr.SetMessage(err.Error())
			}
			if pairData.QtyStep == 0 {
				pairData.QtyStep = minQty
			}
		}

		priceFilter := symbolData.PriceFilter()
		if priceFilter != nil {
			//add max price?
			minPriceRaw := priceFilter.MinPrice
			pairData.MinPrice, err = strconv.ParseFloat(minPriceRaw, 64)
			if err != nil {
				return nil, sharederrs.DataHandleErr.SetMessage(err.Error())
			}
			if pairData.MinPrice == 0 {
				pairData.MinPrice = 0.000001
			}
			priceStepRaw := symbolData.PriceFilter().TickSize
			pairData.PriceStep, err = strconv.ParseFloat(priceStepRaw, 64)
			if err != nil {
				return nil, sharederrs.DataHandleErr.SetMessage(err.Error())
			}
			if pairData.PriceStep == 0 {
				pairData.PriceStep = pairData.MinPrice
			}
		}
		pairs = append(pairs, &pairData)
	}
	return pairs, nil
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
	return &PriceWorkerBinance{}
}

func newBinancePriceWorker() *PriceWorkerBinance {
	return &PriceWorkerBinance{}
}

//SubscribeToPriceEvents - websocket subscription to change quotes and ask-, bid-qty on the exchange
func (w *PriceWorkerBinance) SubscribeToPriceEvents(
	eventCallback func(event workers.PriceEvent),
	errorHandler func(err *sharederrs.APIError),
) *sharederrs.APIError {
	wsBookHandler := func(event *binance.WsBookTickerEvent) {
		if event != nil {
			wEvent := workers.PriceEvent{
				Symbol: event.Symbol,
				Ask:    event.BestAskPrice,
				Bid:    event.BestBidPrice,
			}
			eventCallback(wEvent)
		}
	}
	wsErrHandler := func(err error) {
		errorHandler(sharederrs.ServiceReqFailedErr.M(err.Error()))
	}
	var openWsErr error
	w.WsChannels = new(workers.WorkerChannels)
	w.WsChannels.WsDone, w.WsChannels.WsStop, openWsErr = binance.WsAllBookTickerServe(wsBookHandler, wsErrHandler)
	if openWsErr != nil {
		return sharederrs.ServiceReqFailedErr.M(openWsErr.Error())
	}
	return nil
}

//CandleWorkerBinance - MarketDataWorker for binance
type CandleWorkerBinance struct {
	workers.CandleWorker
}

//GetCandleWorker - create new market candle worker
func (a *BinanceSpotAdapter) GetCandleWorker() workers.ICandleWorker {
	return &CandleWorkerBinance{}
}

func newBinanceCandleWorker() *CandleWorkerBinance {
	return &CandleWorkerBinance{}
}

//SubscribeToCandleEvents - websocket subscription to change trade candles on the exchange
func (w *CandleWorkerBinance) SubscribeToCandleEvents(
	pairSymbols []string,
	eventCallback func(event workers.CandleEvent),
	errorHandler func(err *sharederrs.APIError),
) *sharederrs.APIError {
	timeInterval := "1m"
	symbolIntervalsMap := make(map[string]string)
	for _, symbol := range pairSymbols {
		symbolIntervalsMap[symbol] = timeInterval
	}

	wsCandleHandler := func(event *binance.WsKlineEvent) {
		if event != nil {
			wEvent := workers.CandleEvent{
				Symbol: event.Symbol,
				Candle: workers.CandleData{
					StartTime: event.Kline.StartTime,
					EndTime:   event.Kline.EndTime,
					Interval:  event.Kline.Interval,
					Open:      event.Kline.Open,
					Close:     event.Kline.Close,
					High:      event.Kline.High,
					Low:       event.Kline.Low,
				},
			}
			eventCallback(wEvent)
		}
	}
	wsErrHandler := func(err error) {
		errorHandler(sharederrs.ServiceReqFailedErr.M(err.Error()))
	}
	var openWsErr error
	w.WsChannels = new(workers.WorkerChannels)
	w.WsChannels.WsDone, w.WsChannels.WsStop, openWsErr = binance.WsCombinedKlineServe(symbolIntervalsMap, wsCandleHandler, wsErrHandler)
	if openWsErr != nil {
		return sharederrs.ServiceReqFailedErr.M(openWsErr.Error())
	}
	return nil
}
