package unmanaged_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gilang-anggara/testable-go/routine/unmanaged"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type expectedToBeCalled struct {
	mock.Mock
}

func (e *expectedToBeCalled) Run() {
	e.Called()
}

func Test_ExampleFunction(t *testing.T) {
	e := &expectedToBeCalled{}

	e.On("Run")

	s := unmanaged.Target{e}

	s.ExampleFunction()

	assert.EventuallyWithT(t, func(c *assert.CollectT) {
		fmt.Println(e.Calls)
		assert.True(c, len(e.Calls) == 2)
	}, 10*time.Second, 1*time.Second)
}
