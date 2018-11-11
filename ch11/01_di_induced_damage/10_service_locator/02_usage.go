package service_locator

func Example() {
	locator := buildServiceLocator()
	useServiceLocator(locator)
}

func buildServiceLocator() *ServiceLocator {
	// start of the application or test
	locator := NewServiceLocator()

	// load the dependency mappings
	locator.Store("logger", &myLogger{})

	return locator
}

func useServiceLocator(locator *ServiceLocator) {
	// use the locators to get the logger
	logger := locator.Get("logger").(Logger)

	// use the logger
	logger.Info("Hello World!")
}

type Logger interface {
	Info(message string, args ...interface{})
}

type myLogger struct{}

func (m *myLogger) Info(message string, args ...interface{}) {
	// not implemented
}
