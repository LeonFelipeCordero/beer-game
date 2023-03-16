package adapters

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/graph/model"
)

type OrderApiAdapter struct {
	service       ports.IOrderService
	boardService  ports.IBoardService
	playerService ports.IPlayerService
}

func NewOrderApiAdapter(service ports.IOrderService, boardService ports.IBoardService, playerService ports.IPlayerService) ports.IOrderApi {
	return &OrderApiAdapter{
		service:       service,
		boardService:  boardService,
		playerService: playerService,
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
