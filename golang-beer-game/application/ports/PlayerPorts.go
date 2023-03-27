package ports

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/application/events"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/LeonFelipeCordero/golang-beer-game/graph/model"
)

type IPlayerApi interface {
	AddPlayer(ctx context.Context, boardId string, role string) (*model.Player, error)
	Get(ctx context.Context, id string) (*model.Player, error)
	GetPlayersByBoard(ctx context.Context, boardId string) ([]*model.Player, error)
	UpdateWeeklyOrder(ctx context.Context, playerId string, amount int) (*model.Response, error)
	Subscribe(ctx context.Context, playerId string, streamers *events.Streamers) (chan *model.Player, error)
}

type IPlayerService interface {
	Save(ctx context.Context, player domain.Player) (*domain.Player, error)
	AddPlayer(ctx context.Context, boardId string, role string) (*domain.Player, error)
	Get(ctx context.Context, id string) (*domain.Player, error)
	GetPlayersByBoard(ctx context.Context, boardId string) ([]domain.Player, error)
	UpdateWeeklyOrder(ctx context.Context, playerId string, amount int) (*domain.Player, error)
	GetContraPart(ctx context.Context, player domain.Player) (*domain.Player, error)
}

type IPlayerRepository interface {
	Save(ctx context.Context, player domain.Player) (*domain.Player, error)
	AddPlayer(ctx context.Context, boardId string, player domain.Player) (*domain.Player, error)
	Get(ctx context.Context, id string) (*domain.Player, error)
	GetPlayersByBoard(ctx context.Context, boardId string) ([]domain.Player, error)
	DeleteAll(ctx context.Context)
}
