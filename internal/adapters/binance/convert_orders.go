package binance

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/adshao/go-binance/v2"

	"github.com/matrixbotio/exchange-gates-lib/internal/structs"
	pkgStructs "github.com/matrixbotio/exchange-gates-lib/pkg/structs"
)

// from binance format to our bot order type format
func convertOrderSide(orderSide binance.SideType) (string, error) {
	switch orderSide {
	default:
		return "", errors.New("unknown order type: " + string(orderSide))
	case binance.SideTypeBuy:
		return pkgStructs.OrderTypeBuy, nil
	case binance.SideTypeSell:
		return pkgStructs.OrderTypeSell, nil
	}
}

// converting the order from binance to our format
func convertOrder(orderRaw *binance.Order) (structs.OrderData, error) {
	r := structs.OrderData{}
	awaitQty, err := strconv.ParseFloat(orderRaw.OrigQuantity, 64)
	if err != nil {
		return r, fmt.Errorf("parse await qty: %w", err)
	}

	filledQty, err := strconv.ParseFloat(orderRaw.ExecutedQuantity, 64)
	if err != nil {
		return r, fmt.Errorf("parse executed qty: %w", err)
	}

	price, err := strconv.ParseFloat(orderRaw.Price, 64)
	if err != nil {
		return r, fmt.Errorf("parse price: %w", err)
	}

	orderType, err := convertOrderSide(orderRaw.Side)
	if err != nil {
		return r, err
	}

	r = structs.OrderData{
		OrderID:       orderRaw.OrderID,
		ClientOrderID: orderRaw.ClientOrderID,
		Status:        string(orderRaw.Status),
		AwaitQty:      awaitQty,
		FilledQty:     filledQty,
		Price:         price,
		Symbol:        orderRaw.Symbol,
		Type:          orderType,
		CreatedTime:   orderRaw.Time,
		UpdatedTime:   orderRaw.UpdateTime,
	}
	return r, nil
}

func convertOrders(ordersRaw []*binance.Order) ([]structs.OrderData, error) {
	orders := []structs.OrderData{}
	for _, orderRaw := range ordersRaw {
		order, err := convertOrder(orderRaw)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}
	return orders, nil
}
