package adapters

import (
	"context"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/LeonFelipeCordero/golang-beer-game/graph/model"
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
	fromBoard(*board, boardResponse)
	return boardResponse, nil
}

func (b *BoardApiAdapter) Get(ctx context.Context, id string) (*model.Board, error) {
	board, err := b.service.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	boardResponse := &model.Board{}
	fromBoard(*board, boardResponse)
	return boardResponse, nil
}

func (b *BoardApiAdapter) GetByName(ctx context.Context, name string) (*model.Board, error) {
	board, err := b.service.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}
	boardResponse := &model.Board{}
	fromBoard(*board, boardResponse)
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
	fromBoard(*board, boardResponse)
	return boardResponse, nil
}

func fromBoard(board domain.Board, target *model.Board) {
	target.ID = board.Id
	target.Name = board.Name
	target.State = fromBoardSate(board.State)
	target.Full = board.Full
	target.Finished = board.Finished
	target.CreatedAt = board.CreatedAt.String()
}

func fromBoardSate(state domain.State) model.BoardState {
	var result model.BoardState
	switch state {
	case domain.StateCreated:
		result = model.BoardStateCreated
	case domain.StateRunning:
		result = model.BoardStateRunning
	case domain.StateFinished:
		result = model.BoardStateFinished
	}
	return result
}
