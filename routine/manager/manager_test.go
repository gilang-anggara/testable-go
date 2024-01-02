package manager_test

import (
	"testing"
	"time"

	"github.com/gilang-anggara/testable-go/routine/manager"
	"github.com/stretchr/testify/assert"
)

func Test_RoutineManager_Run(t *testing.T) {
	manager := manager.New()

	count := 0

	manager.Run(func() {
		time.Sleep(500 * time.Millisecond) // simulate long process
		count += 1
	})

	assert.EventuallyWithT(t, func(collect *assert.CollectT) {
		assert.Equal(collect, 1, count)
	}, 10*time.Second, 100*time.Millisecond)
}
