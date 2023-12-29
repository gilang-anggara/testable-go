package routine

type ExpectedToBeCalled interface {
	Run()
}

func NewExpectedToBeCalled() ExpectedToBeCalled {
	return &expectedToBeCalled{}
}

type expectedToBeCalled struct{}

func (*expectedToBeCalled) Run() {}
