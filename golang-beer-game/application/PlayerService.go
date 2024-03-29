package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/application/events"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/google/uuid"
)

type PlayerService struct {
	repository   ports.IPlayerRepository
	boardService ports.IBoardService
	eventChan    chan events.Event
}

func NewPlayerService(
	repository ports.IPlayerRepository,
	boardService ports.IBoardService,
	eventChan chan events.Event,
) ports.IPlayerService {
	return &PlayerService{
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
		p.eventChan <- events.Event{
			Id:        uuid.NewString(),
			ObjectId:  board.Id,
			EventType: events.EventTypeUpdate,
			Object:    *board,
		}
	}
	p.eventChan <- events.Event{
		Id:        uuid.NewString(),
		ObjectId:  board.Id,
		EventType: events.EventTypeNew,
		Object:    *player,
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
	contraPart := getContraPart(*board, player)
	return &contraPart, nil
}

func getContraPart(board domain.Board, receiver domain.Player) domain.Player {
	var sender domain.Player
	switch receiver.Role {
	case domain.RoleRetailer:
		sender = board.GetPlayerByRole(domain.RoleWholesaler)
	case domain.RoleWholesaler:
		sender = board.GetPlayerByRole(domain.RoleFactory)
	}
	return sender
}
