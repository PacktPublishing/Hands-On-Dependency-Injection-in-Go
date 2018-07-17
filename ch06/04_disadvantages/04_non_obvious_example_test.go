package disadvantages_test

func Example() {

}

// StubClient is a stub implementation of disadvantages.Client interface
type StubClient struct{}

// DoSomethingUseful implements disadvantages.Client
func (s *StubClient) DoSomethingUseful() (bool, error) {
	return true, nil
}
