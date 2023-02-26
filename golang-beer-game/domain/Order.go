package domain

import "time"

type Order struct {
	id             string
	amount         int32
	originalAmount int32
	state          Status
	orderType      OrderType
	sender         string
	receiver       *string
	createdAt      time.Time
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
