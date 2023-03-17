package entities

import (
	"github.com/mindstand/gogm/v2"
	"time"
)

type OrderNode struct {
	gogm.BaseUUIDNode

	Amount         int           `gogm:"name=name"`
	OriginalAmount int           `gogm:"name=original_amount"`
	State          string        `gogm:"name=stat"`
	OrderType      string        `gogm:"name=order_type"`
	CreatedAt      time.Time     `gogm:"name=created_at"`
	Sender         []*PlayerNode `gogm:"name=sender;direction=outgoing;relationship=received"`
	Receiver       []*PlayerNode `gogm:"name=receiver;direction=incoming;relationship=ordered"`
}
