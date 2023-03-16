package domain

import "time"

type Order struct {
	Id             string
	Amount         int
	OriginalAmount int
	State          Status
	OrderType      OrderType
	Sender         string
	Receiver       string
	CreatedAt      time.Time
}

type Status string

const (
	StatePending   Status = "PENDING"
	StateDelivered Status = "DELIVERED"
)

type OrderType string

const (
	OrderTypePlayerOrder OrderType = "PLAYER_ORDER"
	OrderTypeCPUOrder    OrderType = "CPU_ORDER"
)

func (o Order) ValidOrderAmount() bool {
	return o.Amount <= o.OriginalAmount
}
