package adapters

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/graph/model"
)

type PlayerApiAdapter struct {
	service      ports.IPlayerService
	boardService ports.IBoardService
}

func NewPlayerApiAdapter(service ports.IPlayerService, boardService ports.IBoardService) ports.IPlayerApi {
	return &PlayerApiAdapter{
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
		playerResponse := &model.Player{}
		playerResponse.FromPlayer(player, boardId)
		playersGraph = append(playersGraph, playerResponse)
	}

	return playersGraph, nil
}

func (b *PlayerApiAdapter) UpdateWeeklyOrder(ctx context.Context, playerId string, amount int) (*model.Response, error) {
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
