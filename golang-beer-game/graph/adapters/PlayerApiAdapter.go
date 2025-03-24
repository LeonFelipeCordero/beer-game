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

type PlayerApiAdapter struct {
	service      application.PlayerService
	boardService application.BoardService
}

func NewPlayerApiAdapter(service application.PlayerService, boardService application.BoardService) PlayerApiAdapter {
	return PlayerApiAdapter{
		service:      service,
		boardService: boardService,
	}
}

func (b *PlayerApiAdapter) AddPlayer(ctx context.Context, boardId string, role string) (*model.Player, error) {
	player, err := b.service.AddPlayer(ctx, boardId, role)
	if err != nil {
		return nil, err
	}
	playerResponse := &model.Player{}
	playerResponse.FromPlayer(*player, boardId)
	return playerResponse, nil
}

func (b *PlayerApiAdapter) Get(ctx context.Context, id string) (*model.Player, error) {
	if id == "" {
		return nil, nil
	}

	player, err := b.service.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	board, err := b.boardService.GetByPlayer(ctx, id)
	if err != nil {
		return nil, err
	}

	playerResponse := &model.Player{}
	playerResponse.FromPlayer(*player, board.Id)
	return playerResponse, nil
}

func (b *PlayerApiAdapter) GetPlayersByBoard(ctx context.Context, boardId string) ([]*model.Player, error) {
	players, err := b.service.GetPlayersByBoard(ctx, boardId)
	if err != nil {
		return nil, err
	}

	var playersGraph []*model.Player
	for _, player := range players {
		playerResponse := model.Player{}
		playerResponse.FromPlayer(player, boardId)
		playersGraph = append(playersGraph, &playerResponse)
	}

	return playersGraph, nil
}

func (b *PlayerApiAdapter) UpdateWeeklyOrder(ctx context.Context, playerId string, amount int64) (*model.Response, error) {
	_, err := b.service.UpdateWeeklyOrder(ctx, playerId, amount)
	if err != nil {
		return nil, err
	}
	message := "weekly order updated"
	status := 200
	return &model.Response{
		Message: &message,
		Status:  &status,
	}, nil
}

func (b *PlayerApiAdapter) Subscribe(ctx context.Context, playerId string, streamers *events.Streamers) (chan *model.Player, error) {
	board, _ := b.boardService.GetByPlayer(ctx, playerId)

	eventChan := make(chan events.Event)
	responseChan := make(chan *model.Player)

	references := []events.Reference{
		{
			Object:    reflect.TypeOf(domain.Player{}).String(),
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

	go b.handleResponses(ctx, eventChan, responseChan, playerId)

	return responseChan, nil
}

func (b *PlayerApiAdapter) handleResponses(ctx context.Context, eventChan chan events.Event, responseChan chan *model.Player, id string) {
loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case event := <-eventChan:
			player := event.Object.(domain.Player)
			// Handle probably do this filtering by adding a filter by attribute in event
			if player.Id == id {
				response := &model.Player{}
				response.FromPlayer(player, event.ObjectId)
				responseChan <- response
			}
		}
	}
}
