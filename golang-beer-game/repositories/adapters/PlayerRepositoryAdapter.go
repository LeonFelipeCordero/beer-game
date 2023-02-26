package adapters

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/LeonFelipeCordero/golang-beer-game/repositories/neo4j"
)

type PlayerRepositoryAdapter struct {
	repository neo4j.IRepository
}

func NewPlayerRepository(repository neo4j.IRepository) *PlayerRepositoryAdapter {
	return &PlayerRepositoryAdapter{
		repository: repository,
	}
}

func (p PlayerRepositoryAdapter) AddPlayer(ctx context.Context, boardId string, player domain.Player) (*domain.Player, error) {
	//TODO implement me
	panic("implement me")
}

func (p PlayerRepositoryAdapter) Get(ctx context.Context, id string) (*domain.Player, error) {
	//TODO implement me
	panic("implement me")
}

func (p PlayerRepositoryAdapter) DeleteAll(ctx context.Context) {
	//TODO implement me
	panic("implement me")
}
