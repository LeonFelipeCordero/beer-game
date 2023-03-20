package application

import (
	"context"
	"errors"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"time"
)

type OrderService struct {
	repository    ports.IOrderRepository
	boardService  ports.IBoardService
	playerService ports.IPlayerService
}

func NewOrderService(
	repository ports.IOrderRepository,
	boardService ports.IBoardService,
	playerService ports.IPlayerService,
) ports.IOrderService {
	return &OrderService{
		repository:    repository,
		boardService:  boardService,
		playerService: playerService,
	}
}

func (o OrderService) CreateOrder(ctx context.Context, receiverId string) (*domain.Order, error) {
	receiver, err := o.playerService.Get(ctx, receiverId)
	if err != nil {
		return nil, err
	}
	sender, err := o.playerService.GetContraPart(ctx, *receiver)
	if err != nil {
		return nil, err
	}

	order := domain.Order{
		Amount:         receiver.WeeklyOrder,
		OriginalAmount: receiver.WeeklyOrder,
		OrderType:      domain.OrderTypePlayerOrder,
		Status:         domain.StatusPending,
		Sender:         sender.Id,
		Receiver:       receiverId,
		CreatedAt:      time.Now(),
	}

	savedOrder, err := o.repository.Save(ctx, order)
	if err != nil {
		return nil, err
	}

	sender.AddOrder(*savedOrder)
	receiver.AddOrder(*savedOrder)
	o.playerService.Save(ctx, *sender)
	o.playerService.Save(ctx, *receiver)

	return savedOrder, nil
}

func (o OrderService) DeliverOrder(ctx context.Context, orderId string, amount int) (*domain.Order, error) {
	order, err := o.repository.Get(ctx, orderId)

	var receiver *domain.Player
	if order.Receiver != "" {
		receiver, err = o.playerService.Get(ctx, order.Receiver)
		if err != nil {
			return nil, err
		}
	}

	sender, err := o.playerService.Get(ctx, order.Sender)
	if err != nil {
		return nil, err
	}

	order.Amount = amount

	if !order.ValidOrderAmount() {
		return nil, errors.New("the new value can't be bigger than the original one")
	} else if !sender.HasStock(order.Amount) {
		return nil, errors.New("the Sender don't have enough stock to deliver this order")
	}

	// todo last order?
	sender.Stock -= order.Amount
	if order.Receiver != "" {
		receiver.Stock += order.Amount
	}
	order.Status = domain.StatusDelivered
	o.playerService.Save(ctx, *receiver)
	o.playerService.Save(ctx, *sender)
	return o.repository.Save(ctx, *order)
}

func (o OrderService) Get(ctx context.Context, orderId string) (*domain.Order, error) {
	return o.repository.Get(ctx, orderId)
}
