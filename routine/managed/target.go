package managed

import (
	"time"

	"github.com/gilang-anggara/testable-go/routine"
)

type Target struct {
	RoutineManager     routine.RoutineManager
	ExpectedToBeCalled routine.ExpectedToBeCalled
}

func (t *Target) ExampleFunction() {
	t.ExpectedToBeCalled.Run()

	t.RoutineManager.Run(
		func() {
			time.Sleep(1 * time.Second)
			t.ExpectedToBeCalled.Run()
		},
	)
}
