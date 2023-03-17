package entities

import "github.com/mindstand/gogm/v2"

type PlayerNode struct {
	gogm.BaseUUIDNode

	Name           string       `gogm:"name=name"`
	Role           string       `gogm:"name=role"`
	Stock          int          `gogm:"name=stock"`
	Backlog        int          `gogm:"name=backlog"`
	WeeklyOrder    int          `gogm:"name=weekly_order"`
	LastOrder      int          `gogm:"name=last_order"`
	Cpu            bool         `gogm:"name=cpu"`
	Board          *BoardNode   `gogm:"name=board;direction=outgoing;relationship=plays_in"`
	OutgoingOrders []*OrderNode `gogm:"name=outgoing_orders;direction=outgoing;relationship=received"`
	IncomingOrders []*OrderNode `gogm:"name=incoming_orders;direction=incoming;relationship=ordered"`
}
