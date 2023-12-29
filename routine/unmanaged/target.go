package unmanaged

import (
	"time"

	"github.com/gilang-anggara/testable-go/routine"
)

type Target struct {
	ExpectedToBeCalled routine.ExpectedToBeCalled
}

func (t *Target) ExampleFunction() {
	t.ExpectedToBeCalled.Run()

	go func() {
		time.Sleep(1 * time.Second)
		t.ExpectedToBeCalled.Run()
	}()
}
