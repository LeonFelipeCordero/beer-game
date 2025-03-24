package adapters

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/application"
	"github.com/LeonFelipeCordero/golang-beer-game/application/events"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/LeonFelipeCordero/golang-beer-game/graph/model"
	"github.com/google/uuid"
	"reflect"
)

type OrderApiAdapter struct {
	service      application.OrderService
	boardService application.BoardService
}

func NewOrderApiAdapter(service application.OrderService, boardService application.BoardService) OrderApiAdapter {
	return OrderApiAdapter{
		service:      service,
		boardService: boardService,
	}
}

func (o *OrderApiAdapter) CreateOrder(ctx context.Context, receiverId string) (*model.Order, error) {
	order, err := o.service.CreateOrder(ctx, receiverId)
	if err != nil {
		return nil, err
	}
	orderResponse := &model.Order{}
	orderResponse.FromOrder(*order)
	return orderResponse, nil
}

func (o *OrderApiAdapter) DeliverOrder(ctx context.Context, orderId string, amount int64) (*model.Response, error) {
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

func (o *OrderApiAdapter) LoadByBoard(ctx context.Context, boardId string) ([]*model.Order, error) {
	loadedOrders, err := o.service.LoadByBoard(ctx, boardId)
	if err != nil {
		return nil, err
	}
	var ordersResponse []*model.Order
	for _, order := range loadedOrders {
		orderResponse := model.Order{}
		orderResponse.FromOrder(order)
		ordersResponse = append(ordersResponse, &orderResponse)
	}
	return ordersResponse, nil
}

func (o *OrderApiAdapter) LoadByPlayer(ctx context.Context, playerId string) ([]*model.Order, error) {
	loadedOrders, err := o.service.LoadByPlayer(ctx, playerId)
	if err != nil {
		return nil, err
	}
	var ordersResponse []*model.Order
	for _, order := range loadedOrders {
		orderResponse := model.Order{}
		orderResponse.FromOrder(order)
		ordersResponse = append(ordersResponse, &orderResponse)
	}
	return ordersResponse, nil
}

func (o *OrderApiAdapter) NewOrderSubscription(ctx context.Context, playerId string, streamers *events.Streamers) (chan *model.Order, error) {
	board, _ := o.boardService.GetByPlayer(ctx, playerId)

	eventChan := make(chan events.Event)
	responseChan := make(chan *model.Order)

	references := []events.Reference{
		{
			Object:    reflect.TypeOf(domain.Order{}).String(),
			ObjectId:  board.Id,
			EventType: events.EventTypeNew,
		},
	}

	streamer := events.Streamer{
		Id:         uuid.NewString(),
		References: references,
		Chan:       eventChan,
	}
	streamers.Register(ctx, streamer)

	go o.handleResponses(ctx, eventChan, responseChan, playerId)

	return responseChan, nil
}

func (o *OrderApiAdapter) OrderDeliveredSubscription(ctx context.Context, playerId string, streamers *events.Streamers) (chan *model.Order, error) {
	board, _ := o.boardService.GetByPlayer(ctx, playerId)

	eventChan := make(chan events.Event)
	responseChan := make(chan *model.Order)

	references := []events.Reference{
		{
			Object:    reflect.TypeOf(domain.Order{}).String(),
			ObjectId:  board.Id,
			EventType: events.EventTypeUpdate,
		},
	}

	streamer := events.Streamer{
		Id:         uuid.NewString(),
		References: references,
		Chan:       eventChan,
	}
	streamers.Register(ctx, streamer)

	go o.handleResponses(ctx, eventChan, responseChan, playerId)

	return responseChan, nil
}

func (o *OrderApiAdapter) handleResponses(ctx context.Context, eventChan chan events.Event, responseChan chan *model.Order, id string) {
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case event := <-eventChan:
			order := event.Object.(domain.Order)
			// Handle probably do this filtering by adding a filter by attribute in event
			if order.Receiver == id || order.Sender == id {
				response := &model.Order{}
				response.FromOrder(order)
				responseChan <- response
			}
		}
	}
}
