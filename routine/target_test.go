package routine_test

import (
	"testing"

	"github.com/gilang-anggara/testable-go/routine"
	"github.com/stretchr/testify/mock"
)

type expectedToBeCalled struct {
	mock.Mock
}

func (e *expectedToBeCalled) Run() {
	e.Called()
}

type manager struct{}

func (r *manager) Run(f func()) {
	f()
}

func Test_Target_ExampleFunction(t *testing.T) {
	e := &expectedToBeCalled{}

	e.On("Run")

	s := routine.Target{&manager{}, e}

	s.ExampleFunction()

	e.AssertExpectations(t)
}
