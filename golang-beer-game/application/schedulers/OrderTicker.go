package schedulers

import (
	"context"
	"fmt"
	"github.com/LeonFelipeCordero/golang-beer-game/application/ports"
	"time"
)

type OrderScheduler struct {
	orderService ports.IOrderService
}

func NewOrderScheduler(service ports.IOrderService) OrderScheduler {
	return OrderScheduler{
		orderService: service,
	}
}

func (o *OrderScheduler) Start() {
	cpuTicker := time.NewTicker(10 * time.Second)
	factoryTicker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-cpuTicker.C:
				fmt.Printf("delivering cpu orders\n")
				o.orderService.CreateCpuOrders(context.Background())
			case <-factoryTicker.C:
				fmt.Printf("delivering factories batch\n")
				o.orderService.DeliverFactoryBatch(context.Background())
			}
		}
	}()
}
