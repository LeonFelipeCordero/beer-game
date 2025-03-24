package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/application/events"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/LeonFelipeCordero/golang-beer-game/repositories/adapters"
	"github.com/google/uuid"
)

type PlayerService struct {
	repository   adapters.PlayerRepositoryAdapter
	boardService BoardService
	eventChan    chan events.Event
}

func NewPlayerService(
	repository adapters.PlayerRepositoryAdapter,
	boardService BoardService,
	eventChan chan events.Event,
) PlayerService {
	return PlayerService{
		repository:   repository,
		boardService: boardService,
		eventChan:    eventChan,
	}
}

func (p *PlayerService) Save(ctx context.Context, player domain.Player) (*domain.Player, error) {
	return p.repository.Save(ctx, player)
}

func (p *PlayerService) AddPlayer(ctx context.Context, boardId string, role string) (*domain.Player, error) {
	selectedRole, err := domain.GetRole(role)
	if err != nil {
		return nil, err
	}

	availableRoles, err := p.boardService.GetAvailableRoles(ctx, boardId)
	var isAvailable = false
	for _, availableRole := range availableRoles {
		if availableRole == selectedRole {
			isAvailable = true
		}
	}
	if !isAvailable {
		return nil, errors.New(fmt.Sprintf("Role %s is already selected in the board", role))
	}

	player, err := p.repository.AddPlayer(ctx, boardId, domain.CreateNewPlayer(selectedRole))
	if err != nil {
		return nil, err
	}
	p.eventChan <- events.Event{
		Id:        uuid.NewString(),
		ObjectId:  boardId,
		EventType: events.EventTypeNew,
		Object:    *player,
	}

	if len(availableRoles) == 1 {
		board, _ := p.boardService.Get(ctx, boardId)
		p.boardService.StartBoard(ctx, boardId)
		p.eventChan <- events.Event{
			Id:        uuid.NewString(),
			ObjectId:  boardId,
			EventType: events.EventTypeUpdate,
			Object:    *board,
		}
	}

	return player, nil
}

func (p *PlayerService) Get(ctx context.Context, id string) (*domain.Player, error) {
	return p.repository.Get(ctx, id)
}

func (p *PlayerService) GetPlayersByBoard(ctx context.Context, boardId string) ([]domain.Player, error) {
	return p.repository.GetPlayersByBoard(ctx, boardId)
}

func (p *PlayerService) GetPlayersByRoleAndBoard(ctx context.Context, boardId string, role string) (*domain.Player, error) {
	return p.repository.GetByRoleAndBoardId(ctx, boardId, role)
}

func (p *PlayerService) UpdateWeeklyOrder(ctx context.Context, playerId string, amount int64) (*domain.Player, error) {
	player, err := p.repository.Get(ctx, playerId)
	if err != nil {
		return nil, err
	}

	player.WeeklyOrder = amount

	response, err := p.repository.Save(ctx, *player)

	board, _ := p.boardService.GetByPlayer(ctx, playerId)
	p.eventChan <- events.Event{
		Id:        uuid.NewString(),
		ObjectId:  board.Id,
		EventType: events.EventTypeUpdate,
		Object:    *player,
	}

	return response, err
}

func (p *PlayerService) GetContraPart(ctx context.Context, player domain.Player) (*domain.Player, error) {
	board, err := p.boardService.GetByPlayer(ctx, player.Id)
	if err != nil {
		return nil, err
	}
	contraPartRole := getContraPart(player)
	contraPart, err := p.repository.GetByRoleAndBoardId(ctx, board.Id, string(contraPartRole))
	if err != nil {
		return nil, err
	}
	return contraPart, nil
}

func getContraPart(receiver domain.Player) domain.Role {
	switch receiver.Role {
	case domain.RoleRetailer:
		return domain.RoleWholesaler
	case domain.RoleWholesaler:
		return domain.RoleFactory
	}
	panic("Role not found in board")
}
