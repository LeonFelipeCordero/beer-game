package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
)

type PlayerService struct {
	repository   ports.IPlayerRepository
	boardService ports.IBoardService
}

func NewPlayerService(
	repository ports.IPlayerRepository,
	boardService ports.IBoardService) ports.IPlayerService {
	return &PlayerService{
		repository:   repository,
		boardService: boardService,
	}
}

func (p *PlayerService) AddPlayer(ctx context.Context, boardId string, role string) (*domain.Player, error) {
	selectedRole, err := domain.GetRole(role)
	if err != nil {
		return nil, err
	}

	board, err := p.boardService.Get(ctx, boardId)
	if !board.HasRoleAvailable(selectedRole) {
		return nil, errors.New(fmt.Sprintf("Role %s is already selected in the board", role))
	}

	player, err := p.repository.AddPlayer(ctx, boardId, domain.CreateNewPlayer(selectedRole))
	if err != nil {
		return nil, err
	}

	if len(board.Players) == 2 {
		p.boardService.CompleteBoard(ctx, boardId)
	}

	return player, nil
}

func (p *PlayerService) Get(ctx context.Context, id string) (*domain.Player, error) {
	return p.repository.Get(ctx, id)
}

func (p *PlayerService) GetPlayersByBoard(ctx context.Context, boardId string) ([]domain.Player, error) {
	return p.repository.GetPlayersByBoard(ctx, boardId)
}

func (p *PlayerService) UpdateWeeklyOrder(ctx context.Context, playerId string, amount int) (*domain.Player, error) {
	player, err := p.repository.Get(ctx, playerId)
	if err != nil {
		return nil, err
	}

	player.WeeklyOrder = amount

	return p.repository.Save(ctx, *player)
}
