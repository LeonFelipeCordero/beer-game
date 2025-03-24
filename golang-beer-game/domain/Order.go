package domain

import "time"

type Order struct {
	Id             string
	Amount         int64
	OriginalAmount int64
	Status         Status
	OrderType      OrderType
	Sender         string
	Receiver       string
	CreatedAt      time.Time
}

type Status string

const (
	StatusPending   Status = "PENDING"
	StatusDelivered Status = "DELIVERED"
)

type OrderType string

const (
	OrderTypePlayerOrder OrderType = "PLAYER_ORDER"
	OrderTypeCPUOrder    OrderType = "CPU_ORDER"
)

func (o Order) ValidOrderAmount() bool {
	return o.Amount <= o.OriginalAmount
}
