package entities

import (
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"github.com/mindstand/gogm/v2"
	"strconv"
	"time"
)

type OrderNode struct {
	gogm.BaseNode

	Amount         int         `gogm:"name=name"`
	OriginalAmount int         `gogm:"name=original_amount"`
	State          string      `gogm:"name=stat"`
	OrderType      string      `gogm:"name=order_type"`
	CreatedAt      time.Time   `gogm:"name=created_at"`
	Sender         *PlayerNode `gogm:"direction=incoming;relationship=received"`
	Receiver       *PlayerNode `gogm:"direction=outgoing;relationship=ordered"`
}

func (o *OrderNode) FromOrder(order domain.Order) {
	if order.Id != "" {
		id, _ := strconv.ParseInt(order.Id, 0, 64)
		o.Id = &id
	}
	o.Amount = order.Amount
	o.OriginalAmount = order.OriginalAmount
	o.State = string(order.State)
	o.OrderType = string(order.OrderType)
	o.CreatedAt = order.CreatedAt
	// todo convert this
	//o.Sender = order.Id
	//o.Receiver = order.Id
}

func (o *OrderNode) ToOrder() domain.Order {
	return domain.Order{
		Id:             strconv.FormatInt(*o.Id, 10),
		Amount:         o.Amount,
		OriginalAmount: o.OriginalAmount,
		State:          toStatus(o.State),
		OrderType:      toType(o.OrderType),
		CreatedAt:      o.CreatedAt,
		Sender:         strconv.FormatInt(*o.Sender.Id, 10),
		Receiver:       strconv.FormatInt(*o.Receiver.Id, 10),
	}
}

func toStatus(status string) domain.Status {
	var result domain.Status
	switch status {
	case string(domain.StatePending):
		result = domain.StatePending
	case string(domain.StateDelivered):
		result = domain.StateDelivered
	}
	return result
}

func toType(orderType string) domain.OrderType {
	var result domain.OrderType
	switch orderType {
	case string(domain.OrderTypeCPUOrder):
		result = domain.OrderTypeCPUOrder
	case string(domain.OrderTypePlayerOrder):
		result = domain.OrderTypePlayerOrder
	}
	return result
}
