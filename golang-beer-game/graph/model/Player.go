package model

import (
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"io"
	"strconv"
)

type Player struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Role        Role     `json:"role"`
	Stock       int      `json:"stock"`
	Backlog     int      `json:"backlog"`
	WeeklyOrder int      `json:"weeklyOrder"`
	LastOrder   int      `json:"lastOrder"`
	CPU         bool     `json:"cpu"`
	BoardId     string   `json:"boardId"`
	OrdersId    []string `json:"ordersId"`
}

type Role string

const (
	RoleRetailer   Role = "RETAILER"
	RoleWholesaler Role = "WHOLESALER"
	RoleFactory    Role = "FACTORY"
)

func (p *Player) FromPlayer(player domain.Player, boardId string) {
	p.ID = player.Id
	p.Name = player.Name
	p.Role = FromPlayerRole(player.Role)
	p.Stock = player.Stock
	p.Backlog = player.Backlog
	p.WeeklyOrder = player.WeeklyOrder
	p.LastOrder = player.LastOrder
	p.CPU = player.Cpu
	p.BoardId = boardId
	p.OrdersId = []string{}
}

func FromPlayerRole(role domain.Role) Role {
	var result Role
	switch role {
	case domain.RoleRetailer:
		result = RoleRetailer
	case domain.RoleWholesaler:
		result = RoleWholesaler
	case domain.RoleFactory:
		result = RoleFactory
	}
	return result
}

func (e Role) IsValid() bool {
	switch e {
	case RoleRetailer, RoleWholesaler, RoleFactory:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
