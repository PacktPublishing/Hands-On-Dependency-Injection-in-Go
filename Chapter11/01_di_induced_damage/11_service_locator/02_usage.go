package service_locator

func Example() {
	locator := buildServiceLocator()
	useServiceLocator(locator)
}

func buildServiceLocator() *ServiceLocator {
	// build a service locator
	locator := NewServiceLocator()

	// load the dependency mappings
	locator.Store("logger", &myLogger{})
	locator.Store("converter", &myConverter{})

	return locator
}

func useServiceLocator(locator *ServiceLocator) {
	// use the locators to get the logger
	logger := locator.Get("logger").(Logger)

	// use the logger
	logger.Info("Hello World!")
}

func useServiceLocatorExtended(locator *ServiceLocator) {
	// use the locators to get the logger
	loggerRetrieved := locator.Get("logger")
	if loggerRetrieved == nil {
		return
	}
	logger, ok := loggerRetrieved.(Logger)
	if !ok {
		return
	}

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

type Converter interface {
	Convert(int float64) (float64, error)
}

type myConverter struct{}

func (m *myConverter) Convert(in float64) (float64, error)  {
	// not implemented
	return 0, nil
}
