package manager

type Manager interface {
	Run(f func())
}

type manager struct{}

func New() Manager {
	return &manager{}
}

func (*manager) Run(f func()) {
	go f()
}
