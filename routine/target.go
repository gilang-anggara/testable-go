package routine

import "time"

type Target struct {
	RoutineManager     RoutineManager
	ExpectedToBeCalled ExpectedToBeCalled
}

func (s *Target) ExampleFunction() {
	s.ExpectedToBeCalled.Run()

	s.RoutineManager.Run(
		func() {
			time.Sleep(1000)
			s.ExpectedToBeCalled.Run()
		},
	)
}
