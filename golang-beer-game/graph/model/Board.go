package model

import (
	"fmt"
	"io"
	"strconv"
)

type Board struct {
	ID             string       `json:"id"`
	Name           string       `json:"name"`
	State          BoardState   `json:"state"`
	Full           bool         `json:"full"`
	Finished       bool         `json:"finished"`
	CreatedAt      string       `json:"createdAt"`
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
