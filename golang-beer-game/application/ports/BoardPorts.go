package ports

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/LeonFelipeCordero/golang-beer-game/graph/model"
)

type IBoardApi interface {
	Create(ctx context.Context, name string) (*model.Board, error)
	Get(ctx context.Context, id string) (*model.Board, error)
	GetByName(ctx context.Context, name string) (*model.Board, error)
}

type IBoardService interface {
	Create(ctx context.Context, name string) (*domain.Board, error)
	Get(ctx context.Context, id string) (*domain.Board, error)
	GetByName(ctx context.Context, name string) (*domain.Board, error)
	CompleteBoard(ctx context.Context, id string) error
}

type IBoardRepository interface {
	Save(ctx context.Context, board domain.Board) (*domain.Board, error)
	Get(ctx context.Context, id string) (*domain.Board, error)
	GetByName(ctx context.Context, name string) (*domain.Board, error)
	GetByPlayer(ctx context.Context, id string) (*domain.Board, error)
	Exist(ctx context.Context, name string) (bool, error)
	DeleteAll(ctx context.Context)
}
