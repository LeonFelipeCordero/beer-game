package adapters

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/graph/model"
)

type PlayerApiAdapter struct {
	service ports.IPlayerService
}

func NewPlayerApiAdapter(service ports.IPlayerService) ports.IPlayerApi {
	return &PlayerApiAdapter{
		service: service,
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
	player, boardId, err := b.service.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	playerResponse := &model.Player{}
	playerResponse.FromPlayer(*player, *boardId)
	return playerResponse, nil
}
