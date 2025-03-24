package domain

import (
	"errors"
	"fmt"
	"time"
)

type Player struct {
	Id          string
	Name        string
	Role        Role
	Stock       int64
	Backlog     int64
	WeeklyOrder int64
	LastOrder   int64
	Cpu         bool
	Orders      []Order
	CreatedAt   time.Time
}

type Role string

const (
	RoleRetailer   Role = "RETAILER"
	RoleWholesaler Role = "WHOLESALER"
	RoleFactory    Role = "FACTORY"
)

func CreateNewPlayer(role Role) Player {
	player := &Player{
		Name:   string(role),
		Role:   role,
		Cpu:    false,
		Orders: []Order{},
    CreatedAt: time.Now(),
	}
	switch role {
	case RoleRetailer:
		player.Stock = 80
		player.Backlog = 80
		player.WeeklyOrder = 40
		player.LastOrder = 40
	case RoleWholesaler:
		player.Stock = 1200
		player.Backlog = 1200
		player.WeeklyOrder = 600
		player.LastOrder = 600
	case RoleFactory:
		player.Stock = 12000
		player.Backlog = 1200
		player.WeeklyOrder = 6000
		player.LastOrder = 6000
	}
	return *player
}

func GetRole(role string) (Role, error) {
	switch role {
	case string(RoleRetailer):
		return RoleRetailer, nil
	case string(RoleWholesaler):
		return RoleWholesaler, nil
	case string(RoleFactory):
		return RoleFactory, nil
	}
	return RoleRetailer, errors.New(fmt.Sprintf("Given role is not supported %s", role))
}

func (p *Player) AddOrder(order Order) {
	p.Orders = append(p.Orders, order)
}

func (p *Player) HasStock(amount int64) bool {
	return p.Stock >= amount
}
