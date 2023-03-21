package adapters

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/graph/model"
)

type OrderApiAdapter struct {
	service ports.IOrderService
}

func NewOrderApiAdapter(service ports.IOrderService) ports.IOrderApi {
	return &OrderApiAdapter{
		service: service,
	}
}

func (o OrderApiAdapter) CreateOrder(ctx context.Context, receiverId string) (*model.Order, error) {
	order, err := o.service.CreateOrder(ctx, receiverId)
	if err != nil {
		return nil, err
	}
	orderResponse := &model.Order{}
	orderResponse.FromOrder(*order)
	return orderResponse, nil
}

func (o OrderApiAdapter) DeliverOrder(ctx context.Context, orderId string, amount int) (*model.Response, error) {
	order, err := o.service.DeliverOrder(ctx, orderId, amount)
	if err != nil {
		return nil, err
	}
	orderResponse := &model.Order{}
	orderResponse.FromOrder(*order)
	message := "order delivered"
	status := 200
	return &model.Response{
		Message: &message,
		Status:  &status,
	}, nil
}

func (o OrderApiAdapter) LoadByBoard(ctx context.Context, boardId string) ([]*model.Order, error) {
	loadedOrders, err := o.service.LoadByBoard(ctx, boardId)
	if err != nil {
		return nil, err
	}
	var ordersResponse []*model.Order
	for _, order := range loadedOrders {
		orderResponse := &model.Order{}
		orderResponse.FromOrder(*order)
		ordersResponse = append(ordersResponse, orderResponse)
	}
	return ordersResponse, nil
}

func (o OrderApiAdapter) LoadByPlayer(ctx context.Context, playerId string) ([]*model.Order, error) {
	loadedOrders, err := o.service.LoadByBoard(ctx, playerId)
	if err != nil {
		return nil, err
	}
	var ordersResponse []*model.Order
	for _, order := range loadedOrders {
		orderResponse := &model.Order{}
		orderResponse.FromOrder(*order)
		ordersResponse = append(ordersResponse, orderResponse)
	}
	return ordersResponse, nil
}
