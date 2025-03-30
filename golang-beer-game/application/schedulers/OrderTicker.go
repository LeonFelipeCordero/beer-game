package schedulers

import (
	"github.com/LeonFelipeCordero/golang-beer-game/application"
	"time"
)

type OrderScheduler struct {
	orderService application.OrderService
}

func NewOrderScheduler(service application.OrderService) OrderScheduler {
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
				// fmt.Printf("delivering cpu orders\n")
				// o.orderService.CreateCpuOrders(context.Background())
			case <-factoryTicker.C:
				// fmt.Printf("delivering factories batch\n")
				// o.orderService.DeliverFactoryBatch(context.Background())
			}
		}
	}()
}
