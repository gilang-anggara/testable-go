package managed_test

import (
	"testing"

	"github.com/gilang-anggara/testable-go/routine/managed"
	"github.com/stretchr/testify/mock"
)

type expectedToBeCalled struct {
	mock.Mock
}

func (e *expectedToBeCalled) Run() {
	e.Called()
}

type routineManager struct{}

func (r *routineManager) Run(f func()) {
	f()
}

func Test_ExampleFunction_WithMockRoutine(t *testing.T) {
	e := &expectedToBeCalled{}

	e.On("Run")

	s := managed.Target{&routineManager{}, e}

	s.ExampleFunction()

	e.AssertNumberOfCalls(t, "Run", 2)
}
