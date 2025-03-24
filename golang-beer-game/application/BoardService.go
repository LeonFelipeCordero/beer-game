package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/application/events"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/LeonFelipeCordero/golang-beer-game/repositories/adapters"
	"time"
)

type BoardService struct {
	repository adapters.BoardRepositoryAdapter
	eventChan  chan events.Event
}

func NewBoardService(repository adapters.BoardRepositoryAdapter, eventChan chan events.Event) BoardService {
	return BoardService{
		repository: repository,
		eventChan:  eventChan,
	}
}

func (s *BoardService) Create(ctx context.Context, name string) (*domain.Board, error) {
	existingBoard, err := s.repository.Exist(ctx, name)
	if err != nil {
		return nil, err
	}
	if existingBoard {
		return nil, errors.New(fmt.Sprintf("Board name %s already exit", name))
	}

	board := domain.Board{
		Name:      name,
		State:     domain.StateCreated,
		Full:      false,
		Finished:  false,
		CreatedAt: time.Now().UTC(),
		Players:   []domain.Player{},
	}

	savedBoard, err := s.repository.Save(ctx, board)
	if err != nil {
		return nil, err
	}

	return savedBoard, nil
}

func (s *BoardService) GetByName(ctx context.Context, name string) (*domain.Board, error) {
	return s.repository.GetByName(ctx, name)
}

func (s *BoardService) Get(ctx context.Context, id string) (*domain.Board, error) {
	return s.repository.Get(ctx, id)
}

func (s *BoardService) GetByPlayer(ctx context.Context, playerId string) (*domain.Board, error) {
	return s.repository.GetByPlayer(ctx, playerId)
}

func (s *BoardService) StartBoard(ctx context.Context, id string) error {
	board, err := s.repository.Get(ctx, id)
	if err != nil {
		return err
	}

	availableRoles, err := s.repository.GetAvailableRoles(ctx, board.Id)
	if len(availableRoles) == 0 {
		return fmt.Errorf("not possible to start a game without all roles being selected %s", board.Id)
	}

	err = s.repository.StartBoard(ctx, board.Id)

	return nil
}

func (s *BoardService) GetAvailableRoles(ctx context.Context, id string) ([]domain.Role, error) {
	roles, err := s.repository.GetAvailableRoles(ctx, id)
	if err != nil {
		return nil, err
	}

	availableRoles := []domain.Role{domain.RoleRetailer, domain.RoleWholesaler, domain.RoleFactory}
	for _, role := range roles {
		for index, availableRole := range availableRoles {
			if availableRole == domain.Role(role) {
				availableRoles = append(availableRoles[:index], availableRoles[index+1:]...)
			}
		}
	}

	return availableRoles, nil
}

func (s *BoardService) GetActiveBoards(ctx context.Context) ([]domain.Board, error) {
	return s.repository.GetActiveBoards(ctx)
}
