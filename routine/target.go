package routine

import (
	"time"

	"github.com/gilang-anggara/testable-go/routine/manager"
)

type Target struct {
	RoutineManager     manager.Manager
	ExpectedToBeCalled ExpectedToBeCalled
}

func (t *Target) ExampleFunction() {
	t.RoutineManager.Run(
		func() {
			time.Sleep(500 * time.Millisecond) // simulate long process
			t.ExpectedToBeCalled.Run()
		},
	)
}
