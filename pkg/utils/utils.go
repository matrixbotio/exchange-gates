package utils

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/matrixbotio/exchange-gates-lib/internal/consts"
	"github.com/matrixbotio/exchange-gates-lib/internal/structs"
	"github.com/matrixbotio/exchange-gates-lib/internal/workers"
	pkgStructs "github.com/matrixbotio/exchange-gates-lib/pkg/structs"
	"github.com/shopspring/decimal"
)

// GetFloatPrecision returns the number of decimal places in a float
func GetFloatPrecision(value float64) int {
	// if you put 15, then the test will fall,
	// because Float is rounded incorrectly
	maxPrecision := 14
	valueFormated := strconv.FormatFloat(math.Abs(value), 'f', maxPrecision, 64)
	valueParts := strings.Split(valueFormated, ".")
	if len(valueParts) <= 1 {
		return 0
	}
	valueLastPartTrimmed := strings.TrimRight(valueParts[1], "0")
	return len(valueLastPartTrimmed)
}

// LogNotNilError logs an array of errors and returns true if an error is found
func LogNotNilError(errs []error) bool {
	for _, err := range errs {
		if err != nil {
			log.Println(err)
			return true
		}
	}
	return false
}

// OrderResponseToBotOrder - convert raw order response to bot order
func OrderResponseToBotOrder(response structs.CreateOrderResponse) pkgStructs.BotOrder {
	return pkgStructs.BotOrder{
		PairSymbol:    response.Symbol,
		Type:          response.Type,
		Qty:           response.OrigQuantity,
		Price:         response.Price,
		Deposit:       response.OrigQuantity * response.Price,
		ClientOrderID: response.ClientOrderID,
	}
}

// OrderDataToBotOrder - convert order data to bot order
func OrderDataToBotOrder(order structs.OrderData) pkgStructs.BotOrder {
	return pkgStructs.BotOrder{
		PairSymbol:    order.Symbol,
		Type:          order.Type,
		Qty:           order.AwaitQty,
		Price:         order.Price,
		Deposit:       order.AwaitQty * order.Price,
		ClientOrderID: order.ClientOrderID,
	}
}

func RoundFloatFloor(val float64, precision int) (float64, error) {
	if math.IsNaN(val) {
		return 0, errors.New("value is NaN")
	}
	if math.IsInf(val, 0) {
		return 0, errors.New("value is Inf")
	}

	f, _ := decimal.NewFromFloat(val).RoundFloor(int32(precision)).Float64()
	return f, nil
}

func formatFloatFloor(val float64, precision int) (string, error) {
	if precision == 0 {
		return strconv.FormatFloat(val, 'f', 0, 64), nil
	}

	valRounded, err := RoundFloatFloor(val, precision)
	if err != nil {
		return "", fmt.Errorf("round value: %w", err)
	}

	v := strings.TrimRight(
		strconv.FormatFloat(valRounded, 'f', precision, 64),
		".0",
	)
	if v == "" {
		return "0", nil
	}
	return v, nil
}

/*
RoundPairOrderValues - adjusts the order values in accordance
with the trading pair parameters
*/
func RoundPairOrderValues(
	order pkgStructs.BotOrder,
	pairLimits structs.ExchangePairData,
) (structs.BotOrderAdjusted, error) {
	result := structs.BotOrderAdjusted{
		PairSymbol:       order.PairSymbol,
		Type:             order.Type,
		ClientOrderID:    order.ClientOrderID,
		MinQty:           pairLimits.MinQty,
		MinQtyPassed:     true, // by default
		MinDeposit:       pairLimits.OriginalMinDeposit,
		MinDepositPassed: true, // by default
	}

	if order.Qty == 0 {
		return structs.BotOrderAdjusted{}, errors.New("order qty is not set")
	}
	if order.Qty < pairLimits.MinQty {
		result.MinQtyPassed = false
	}
	if order.Qty > pairLimits.MaxQty {
		return structs.BotOrderAdjusted{}, errors.New("too much amount to open an order in this pair. " +
			"order qty: " + strconv.FormatFloat(order.Qty, 'f', 8, 32) +
			" max: " + strconv.FormatFloat(pairLimits.MaxQty, 'f', 8, 32))
	}

	if order.Price == 0 {
		return structs.BotOrderAdjusted{}, errors.New("order price is not set")
	}
	if order.Price < pairLimits.MinPrice {
		return structs.BotOrderAdjusted{}, errors.New("insufficient price to open an order in this pair. " +
			"order price: " + strconv.FormatFloat(order.Price, 'f', 8, 32) +
			" min: " + strconv.FormatFloat(pairLimits.MinPrice, 'f', 8, 32))
	}

	// check min deposit
	orderDeposit := order.Qty * order.Price
	if orderDeposit < pairLimits.OriginalMinDeposit {
		result.MinDepositPassed = false
	}

	// round order values
	var err error
	result.Qty, err = formatFloatFloor(order.Qty, GetFloatPrecision(pairLimits.QtyStep))
	if err != nil {
		return structs.BotOrderAdjusted{}, fmt.Errorf("format qty: %w", err)
	}
	result.Price, err = formatFloatFloor(order.Price, GetFloatPrecision(pairLimits.PriceStep))
	if err != nil {
		return structs.BotOrderAdjusted{}, fmt.Errorf("format price: %w", err)
	}

	qtyRounded, err := strconv.ParseFloat(result.Qty, 64)
	if err != nil {
		return result, fmt.Errorf("parse order qty: %w", err)
	}
	priceRounded, err := strconv.ParseFloat(result.Price, 64)
	if err != nil {
		return structs.BotOrderAdjusted{}, fmt.Errorf("parse order price: %w", err)
	}

	depositRounded := qtyRounded * priceRounded
	result.Deposit = strconv.FormatFloat(depositRounded, 'f', GetFloatPrecision(orderDeposit), 64)
	return result, nil
}

// RoundDeposit - round deposit for grid by pair limits
func RoundDeposit(deposit float64, pairLimits structs.ExchangePairData) (float64, error) {
	depositStep := pairLimits.PriceStep * pairLimits.QtyStep
	depositRoundedStr := strconv.FormatFloat(deposit, 'f', GetFloatPrecision(depositStep), 64)
	depositRounded, err := strconv.ParseFloat(depositRoundedStr, 64)
	if err != nil {
		return 0, fmt.Errorf("round deposit: %w", err)
	}
	return depositRounded, nil
}

// ParseAdjustedOrder - parse rounded order to bot order
func ParseAdjustedOrder(order structs.BotOrderAdjusted) (pkgStructs.BotOrder, error) {
	resultOrder := pkgStructs.BotOrder{
		PairSymbol: order.PairSymbol,
		Type:       order.Type,
	}
	// parse qty
	var err error
	resultOrder.Qty, err = strconv.ParseFloat(order.Qty, 64)
	if err != nil {
		return resultOrder, fmt.Errorf("parse order qty: %w", err)
	}
	// parse price
	resultOrder.Price, err = strconv.ParseFloat(order.Price, 64)
	if err != nil {
		return resultOrder, fmt.Errorf("parse order price: %w", err)
	}
	// parse deposit
	resultOrder.Deposit, err = strconv.ParseFloat(order.Deposit, 64)
	if err != nil {
		return resultOrder, fmt.Errorf("parse order deposit: %w", err)
	}
	return resultOrder, nil
}

// GetDefaultPairData !
func GetDefaultPairData() structs.ExchangePairData {
	return structs.ExchangePairData{
		ExchangeID: consts.PairDefaultExchangeID,
		BaseAsset:  consts.PairDefaultBaseAsset,
		QuoteAsset: consts.PairDefaultQuoteAsset,
		MinQty:     consts.PairDefaultMinQty,
		MaxQty:     consts.PairDefaultMaxQty,
		MinDeposit: consts.PairMinDeposit,
		MinPrice:   consts.PairDefaultMinPrice,
		QtyStep:    consts.PairDefaultQtyStep,
	}
}

// OrderDataToTradeEvent data
type TradeOrderConvertTask struct {
	Order       structs.OrderData
	ExchangeTag string
}

// OrderDataToTradeEvent - convert order data into a trade event.
func OrderDataToTradeEvent(task TradeOrderConvertTask) workers.TradeEvent {
	e := workers.TradeEvent{
		ID:          0,
		Time:        task.Order.UpdatedTime,
		Symbol:      task.Order.Symbol,
		Price:       task.Order.Price,
		Quantity:    task.Order.FilledQty,
		ExchangeTag: task.ExchangeTag,
	}

	if task.Order.Type == pkgStructs.OrderTypeBuy {
		e.BuyerOrderID = task.Order.OrderID
	} else {
		e.SellerOrderID = task.Order.OrderID
	}

	return e
}

func CalcTPOrder(
	strategy pkgStructs.BotStrategy,
	coinsQty float64,
	profit float64,
	depositSpent float64,
	pairSymbol string,
) pkgStructs.BotOrder {
	if strategy == pkgStructs.BotStrategyShort {
		return calcShortTPOrder(coinsQty, profit, depositSpent, pairSymbol)
	}
	return calcLongOrder(coinsQty, profit, depositSpent, pairSymbol)
}

func GetTPOrderType(strategy pkgStructs.BotStrategy) string {
	if strategy == pkgStructs.BotStrategyLong {
		return pkgStructs.OrderTypeSell
	}
	return pkgStructs.OrderTypeBuy
}

func calcShortTPOrder(
	coinsQty float64,
	profit float64,
	depositSpent float64,
	pairSymbol string,
) pkgStructs.BotOrder {
	orderType := GetTPOrderType(pkgStructs.BotStrategyShort)
	tpQty := (1 + profit/100) * coinsQty
	tpPrice := depositSpent / tpQty

	return pkgStructs.BotOrder{
		PairSymbol: pairSymbol,
		Type:       orderType,
		Qty:        tpQty,
		Price:      tpPrice,
		Deposit:    depositSpent,
	}
}

func calcLongOrder(
	coinsQty float64,
	profit float64,
	depositSpent float64,
	pairSymbol string,
) pkgStructs.BotOrder {
	orderType := GetTPOrderType(pkgStructs.BotStrategyLong)
	tpDeposit := (1 + profit/100) * depositSpent
	tpPrice := tpDeposit / coinsQty

	return pkgStructs.BotOrder{
		PairSymbol: pairSymbol,
		Type:       orderType,
		Qty:        coinsQty,
		Price:      tpPrice,
		Deposit:    tpDeposit,
	}
}
