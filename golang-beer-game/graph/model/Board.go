package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/LeonFelipeCordero/golang-beer-game/domain"
)

type Board struct {
	ID             string       `json:"id"`
	Name           string       `json:"name"`
	State          BoardState   `json:"state"`
	Full           bool         `json:"full"`
	Finished       bool         `json:"finished"`
	CreatedAt      time.Time    `json:"createdAt"`
	PlayersId      []string     `json:"playersId"`
	OrdersId       []string     `json:"ordersId"`
	AvailableRoles []BoardState `json:"availableRoles"`
}

type BoardState string

const (
	BoardStateCreated  BoardState = "CREATED"
	BoardStateRunning  BoardState = "RUNNING"
	BoardStateFinished BoardState = "FINISHED"
)

func (e BoardState) IsValid() bool {
	switch e {
	case BoardStateCreated, BoardStateRunning, BoardStateFinished:
		return true
	}
	return false
}

func (e BoardState) String() string {
	return string(e)
}

func (e *BoardState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BoardState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BardState", str)
	}
	return nil
}

func (e BoardState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func (e *Board) FromBoard(board domain.Board) {
	e.ID = board.Id
	e.Name = board.Name
	e.State = fromBoardSate(board.State)
	e.Full = board.Full
	e.Finished = board.Finished
	e.CreatedAt = board.CreatedAt
}

func fromBoardSate(state domain.State) BoardState {
	var result BoardState
	switch state {
	case domain.StateCreated:
		result = BoardStateCreated
	case domain.StateRunning:
		result = BoardStateRunning
	case domain.StateFinished:
		result = BoardStateFinished
	}
	return result
}
