package adapters

import (
	"context"
	"errors"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/google/uuid"
)

type BoardRepositoryAdapterFaker struct {
	boards map[string]domain.Board
}

func NewBoardRepositoryFaker() ports.IBoardRepository {
	return &BoardRepositoryAdapterFaker{
		boards: make(map[string]domain.Board),
	}
}

func (b *BoardRepositoryAdapterFaker) Save(ctx context.Context, board domain.Board) (*domain.Board, error) {
	existingBoard, err := b.Get(ctx, board.Id)
	if err == nil {
		delete(b.boards, existingBoard.Id)
		b.boards[existingBoard.Id] = board
		return &board, nil
	}

	id, _ := uuid.NewUUID()
	board.Id = id.String()
	b.boards[id.String()] = board
	return &board, nil
}

func (b *BoardRepositoryAdapterFaker) Get(ctx context.Context, id string) (*domain.Board, error) {
	for key, value := range b.boards {
		if key == id {
			return &value, nil
		}
	}
	return nil, errors.New("no board found")
}

func (b *BoardRepositoryAdapterFaker) GetByName(ctx context.Context, name string) (*domain.Board, error) {
	for _, value := range b.boards {
		if value.Name == name {
			return &value, nil
		}
	}
	return nil, errors.New("no board found")
}

func (b *BoardRepositoryAdapterFaker) Exist(ctx context.Context, name string) (bool, error) {
	_, err := b.GetByName(ctx, name)

	if err != nil && err.Error() == "no board found" {
		return false, nil
	} else if err != nil && err.Error() != "no board found" {
		return true, err
	}
	return true, nil
}

func (b *BoardRepositoryAdapterFaker) DeleteAll(ctx context.Context) {
	for k := range b.boards {
		delete(b.boards, k)
	}
}

func (b *BoardRepositoryAdapterFaker) GetByPlayer(ctx context.Context, id string) (*domain.Board, error) {
	for _, bv := range b.boards {
		for _, player := range bv.Players {
			if player.Id == id {
				return &bv, nil
			}
		}
	}
	return nil, errors.New("player is not assign to a board")
}
