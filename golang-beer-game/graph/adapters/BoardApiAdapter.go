package adapters

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/application/events"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/LeonFelipeCordero/golang-beer-game/graph/model"
	"github.com/google/uuid"
	"reflect"
)

type BoardApiAdapter struct {
	service ports.IBoardService
}

func NewBoardApiAdapter(service ports.IBoardService) ports.IBoardApi {
	return &BoardApiAdapter{
		service: service,
	}
}

func (b *BoardApiAdapter) Create(ctx context.Context, name string) (*model.Board, error) {
	board, err := b.service.Create(ctx, name)
	if err != nil {
		return nil, err
	}
	boardResponse := &model.Board{}
	boardResponse.FromBoard(*board)
	return boardResponse, nil
}

func (b *BoardApiAdapter) Get(ctx context.Context, id string) (*model.Board, error) {
	board, err := b.service.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	boardResponse := &model.Board{}
	boardResponse.FromBoard(*board)
	return boardResponse, nil
}

func (b *BoardApiAdapter) GetByName(ctx context.Context, name string) (*model.Board, error) {
	board, err := b.service.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}
	boardResponse := &model.Board{}
	boardResponse.FromBoard(*board)
	return boardResponse, nil
}

func (b *BoardApiAdapter) GetAvailableRoles(ctx context.Context, id string) ([]*model.Role, error) {
	roles, err := b.service.GetAvailableRoles(ctx, id)
	if err != nil {
		return nil, err
	}

	var rolesResponse []*model.Role
	for _, role := range roles {
		roleResponse := model.FromPlayerRole(role)
		rolesResponse = append(rolesResponse, &roleResponse)
	}

	return rolesResponse, nil
}

func (b *BoardApiAdapter) GetByPlayer(ctx context.Context, playerId string) (*model.Board, error) {
	board, err := b.service.GetByPlayer(ctx, playerId)
	if err != nil {
		return nil, err
	}

	boardResponse := &model.Board{}
	boardResponse.FromBoard(*board)
	return boardResponse, nil
}

func (b *BoardApiAdapter) Subscribe(ctx context.Context, boardId string, streamers *events.Streamers) (chan *model.Board, error) {
	eventChan := make(chan events.Event)
	responseChan := make(chan *model.Board)

	references := []events.Reference{
		{
			Object:    reflect.TypeOf(domain.Board{}).String(),
			ObjectId:  boardId,
			EventType: events.EventTypeUpdate,
		},
		{
			Object:    reflect.TypeOf(domain.Player{}).String(),
			ObjectId:  boardId,
			EventType: events.EventTypeNew,
		},
	}

	streamer := events.Streamer{
		Id:         uuid.NewString(),
		References: references,
		Chan:       eventChan,
	}
	streamers.Register(ctx, streamer)

	go b.handleResponses(ctx, eventChan, responseChan)

	return responseChan, nil
}

func (b *BoardApiAdapter) handleResponses(ctx context.Context, eventChan chan events.Event, responseChan chan *model.Board) {
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case event := <-eventChan:
			var board domain.Board
			if event.Object == reflect.TypeOf(domain.Board{}).String() {
				board = event.Object.(domain.Board)
			} else {
				ref, _ := b.service.Get(ctx, event.Id)
				board = *ref
			}
			response := &model.Board{}
			response.FromBoard(board)
			responseChan <- response
		}
	}
}
