package model

import (
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/domain"
	"io"
	"strconv"
)

type Order struct {
	ID             string     `json:"id"`
	Amount         int        `json:"amount"`
	OriginalAmount int        `json:"originalAmount"`
	State          OrderState `json:"state"`
	Type           OrderType  `json:"type"`
	SenderId       string     `json:"senderId"`
	ReceiverId     string     `json:"receiverId"`
	BoardId        string     `json:"boardId"`
	CreatedAt      string     `json:"createdAt"`
}

func (o *Order) FromOrder(order domain.Order) {
	// todo find the way to pass the board here if need for graphql
	o.ID = order.Id
	o.Amount = order.Amount
	o.OriginalAmount = order.OriginalAmount
	o.State = FromOrderStatus(order.Status)
	o.Type = FromOrderType(order.OrderType)
	o.SenderId = order.Sender
	o.ReceiverId = order.Receiver
	o.CreatedAt = order.CreatedAt.String()
}

func FromOrderStatus(status domain.Status) OrderState {
	var result OrderState
	switch status {
	case domain.StatusPending:
		result = OrderStatePending
	case domain.StatusDelivered:
		result = OrderStateDelivered
	}
	return result
}

func FromOrderType(orderType domain.OrderType) OrderType {
	var result OrderType
	switch orderType {
	case domain.OrderTypePlayerOrder:
		result = OrderTypePlayerOrder
	case domain.OrderTypeCPUOrder:
		result = OrderTypeCPUOrder
	}
	return result
}

type OrderState string

const (
	OrderStatePending   OrderState = "PENDING"
	OrderStateDelivered OrderState = "DELIVERED"
)

var AllOrderState = []OrderState{
	OrderStatePending,
	OrderStateDelivered,
}

func (e OrderState) IsValid() bool {
	switch e {
	case OrderStatePending, OrderStateDelivered:
		return true
	}
	return false
}

func (e OrderState) String() string {
	return string(e)
}

func (e *OrderState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderState", str)
	}
	return nil
}

func (e OrderState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderType string

const (
	OrderTypePlayerOrder OrderType = "PLAYER_ORDER"
	OrderTypeCPUOrder    OrderType = "CPU_ORDER"
)

var AllOrderType = []OrderType{
	OrderTypePlayerOrder,
	OrderTypeCPUOrder,
}

func (e OrderType) IsValid() bool {
	switch e {
	case OrderTypePlayerOrder, OrderTypeCPUOrder:
		return true
	}
	return false
}

func (e OrderType) String() string {
	return string(e)
}

func (e *OrderType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderType", str)
	}
	return nil
}

func (e OrderType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
