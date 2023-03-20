package entities

import (
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/mindstand/gogm/v2"
	"strconv"
)

type PlayerNode struct {
	gogm.BaseNode

	Name           string       `gogm:"name=name"`
	Role           string       `gogm:"name=role"`
	Stock          int          `gogm:"name=stock"`
	Backlog        int          `gogm:"name=backlog"`
	WeeklyOrder    int          `gogm:"name=weekly_order"`
	LastOrder      int          `gogm:"name=last_order"`
	Cpu            bool         `gogm:"name=cpu"`
	Board          *BoardNode   `gogm:"direction=outgoing;relationship=plays_in"`
	OutgoingOrders []*OrderNode `gogm:"direction=outgoing;relationship=received"`
	IncomingOrders []*OrderNode `gogm:"direction=incoming;relationship=ordered"`
}

func (p *PlayerNode) FromPlayer(player domain.Player) {
	if player.Id != "" {
		id, _ := strconv.ParseInt(player.Id, 0, 64)
		p.Id = &id
	}
	p.Name = player.Name
	p.Role = string(player.Role)
	p.Stock = player.Stock
	p.Backlog = player.Backlog
	p.WeeklyOrder = player.WeeklyOrder
	p.LastOrder = player.LastOrder
	p.Cpu = player.Cpu

	outgoingOrders, incomingOrders := classifyOrders(player.Id, player.Orders)
	p.OutgoingOrders = outgoingOrders
	p.IncomingOrders = incomingOrders
	// todo convert this
	//p.Board = player.Name
}

func (p *PlayerNode) ToPlayer() domain.Player {
	return domain.Player{
		Id:          strconv.FormatInt(*p.Id, 10),
		Name:        p.Name,
		Role:        toRole(p.Role),
		Stock:       p.Stock,
		Backlog:     p.Backlog,
		WeeklyOrder: p.WeeklyOrder,
		LastOrder:   p.LastOrder,
		Cpu:         p.Cpu,
		Orders:      toOrders(p.IncomingOrders, p.OutgoingOrders),
	}
}

func classifyOrders(playerId string, orders []domain.Order) ([]*OrderNode, []*OrderNode) {
	var outgoingOrder []*OrderNode
	var incomingOrder []*OrderNode
	for _, order := range orders {
		orderNode := OrderNode{}
		orderNode.FromOrder(order)
		if order.Sender == playerId {
			outgoingOrder = append(outgoingOrder, &orderNode)
		} else {
			incomingOrder = append(incomingOrder, &orderNode)
		}
	}
	return outgoingOrder, incomingOrder
}

func toRole(role string) domain.Role {
	var result domain.Role
	switch role {
	case string(domain.RoleRetailer):
		result = domain.RoleRetailer
	case string(domain.RoleWholesaler):
		result = domain.RoleWholesaler
	case string(domain.RoleFactory):
		result = domain.RoleFactory
	}
	return result
}

func toOrders(incoming []*OrderNode, outgoing []*OrderNode) []domain.Order {
	var orders []domain.Order
	for _, order := range incoming {
		orders = append(orders, order.ToOrder())
	}
	for _, order := range outgoing {
		orders = append(orders, order.ToOrder())
	}
	return orders
}
