package disadvantages

type InnerService struct {
	innerDep Dependency
}

func NewInnerService(innerDep Dependency) *InnerService {
	return &InnerService{
		innerDep: innerDep,
	}
}

type OuterService struct {
	// composition
	innerService *InnerService

	outerDep Dependency
}

func NewOuterService(outerDep Dependency, innerDep Dependency) *OuterService {
	return &OuterService{
		innerService: NewInnerService(innerDep),
		outerDep:     outerDep,
	}
}

// fake type to satisfy the compiler
type Dependency interface {
}
