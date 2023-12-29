package routine

type RoutineManager interface {
	Run(f func())
}

type routineManager struct{}

func NewRoutineManager() RoutineManager {
	return &routineManager{}
}

func (*routineManager) Run(f func()) {
	go f()
}
