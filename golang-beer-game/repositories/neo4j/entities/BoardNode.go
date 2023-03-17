package entities

import (
	"github.com/mindstand/gogm/v2"
	"time"
)

type BoardNode struct {
	gogm.BaseUUIDNode

	Name      string        `gogm:"name=name;unique"`
	State     string        `gogm:"name=state"`
	Full      bool          `gogm:"name=full"`
	Finished  bool          `gogm:"name=finished"`
	CreatedAt time.Time     `gogm:"name=created_at"`
	Players   []*PlayerNode `gogm:"name=players;direction=incoming;relationship=plays_in"`
}
