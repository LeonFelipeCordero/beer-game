package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/application/events"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/google/uuid"
	"time"
)

type OrderService struct {
	repository    ports.IOrderRepository
	playerService ports.IPlayerService
	boardService  ports.IBoardService
	eventChan     chan events.Event
}

func NewOrderService(
	repository ports.IOrderRepository,
	playerService ports.IPlayerService,
	boardService ports.IBoardService,
	eventChan chan events.Event,
) ports.IOrderService {
	return &OrderService{
		repository:    repository,
		playerService: playerService,
		boardService:  boardService,
		eventChan:     eventChan,
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
	board, err := o.boardService.GetByPlayer(ctx, receiverId)
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
		CreatedAt:      time.Now().UTC(),
	}

	savedOrder, err := o.repository.Save(ctx, order)
	if err != nil {
		return nil, err
	}
	o.eventChan <- events.Event{
		Id:        uuid.NewString(),
		ObjectId:  board.Id,
		EventType: events.EventTypeNew,
		Object:    *savedOrder,
	}

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

	board, err := o.boardService.GetByPlayer(ctx, sender.Id)
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

	//todo do all saves in a single transaction
	if receiver != nil {
		o.playerService.Save(ctx, *receiver)

		o.eventChan <- events.Event{
			Id:        uuid.NewString(),
			ObjectId:  board.Id,
			EventType: events.EventTypeUpdate,
			Object:    *receiver,
		}
	}
	o.playerService.Save(ctx, *sender)
	savedOrder, err := o.repository.Save(ctx, *order)
	if err != nil {
		return nil, err
	}

	o.eventChan <- events.Event{
		Id:        uuid.NewString(),
		ObjectId:  board.Id,
		EventType: events.EventTypeUpdate,
		Object:    *savedOrder,
	}

	o.eventChan <- events.Event{
		Id:        uuid.NewString(),
		ObjectId:  board.Id,
		EventType: events.EventTypeUpdate,
		Object:    *sender,
	}

	return savedOrder, nil
}

func (o OrderService) Get(ctx context.Context, orderId string) (*domain.Order, error) {
	return o.repository.Get(ctx, orderId)
}

func (o OrderService) LoadByBoard(ctx context.Context, boardId string) ([]*domain.Order, error) {
	return o.repository.LoadByBoard(ctx, boardId)
}

func (o OrderService) LoadByPlayer(ctx context.Context, playerId string) ([]*domain.Order, error) {
	return o.repository.LoadByPlayer(ctx, playerId)
}

func (o OrderService) DeliverFactoryBatch(ctx context.Context) {
	boards, err := o.boardService.GetActiveBoards(ctx)
	if err != nil {
		panic(fmt.Sprintf("error getting active boards %s", err))
	}
	for _, board := range boards {
		factory := board.GetFactory()
		if factory != nil {
			factory.Stock += factory.WeeklyOrder
			factory.Backlog += factory.WeeklyOrder
			factory.LastOrder = factory.WeeklyOrder
			o.playerService.Save(ctx, *factory)
			o.eventChan <- events.Event{
				Id:        uuid.NewString(),
				ObjectId:  board.Id,
				EventType: events.EventTypeUpdate,
				Object:    *factory,
			}
		}
	}
}

func (o OrderService) CreateCpuOrders(ctx context.Context) {
	boards, err := o.boardService.GetActiveBoards(ctx)
	if err != nil {
		panic(fmt.Sprintf("error getting active boards %s", err))
	}
	for _, board := range boards {
		for _, player := range board.Players {
			order := domain.Order{
				Amount:         player.WeeklyOrder / 4,
				OriginalAmount: player.WeeklyOrder / 4,
				OrderType:      domain.OrderTypeCPUOrder,
				Status:         domain.StatusPending,
				Sender:         player.Id,
				CreatedAt:      time.Now().UTC(),
			}
			savedOrder, err2 := o.repository.Save(ctx, order)
			if err2 != nil {
				panic("Something went wrong creating cpu order")
			}
			o.eventChan <- events.Event{
				Id:        uuid.NewString(),
				ObjectId:  board.Id,
				EventType: events.EventTypeNew,
				Object:    *savedOrder,
			}
		}
	}
}
