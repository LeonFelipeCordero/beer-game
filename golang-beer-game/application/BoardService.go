package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"time"
)

type BoardService struct {
	repository ports.IBoardRepository
}

func NewBoardService(repository ports.IBoardRepository) ports.IBoardService {
	return &BoardService{repository: repository}
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

func (s *BoardService) CompleteBoard(ctx context.Context, id string) error {
	board, err := s.repository.Get(ctx, id)
	if err != nil {
		return err
	}

	if len(board.Players) != 3 {
		return err
	}

	board.Start()
	board, err = s.repository.Save(ctx, *board)

	return nil
}

func (s *BoardService) GetAvailableRoles(ctx context.Context, id string) ([]domain.Role, error) {
	board, err := s.repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return board.AvailableRoles(), nil
}
