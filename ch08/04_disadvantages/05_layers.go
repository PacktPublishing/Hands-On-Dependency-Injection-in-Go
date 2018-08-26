package disadvantages

func NewLayer1Object(config Layer1Config) *Layer1Object {
	return &Layer1Object{
		MyConfig:     config,
		MyDependency: NewLayer2Object(config),
	}
}

// Configuration for the Layer 1 Object
type Layer1Config interface {
	Logger() Logger
}

// Layer 1 Object
type Layer1Object struct {
	MyConfig     Layer1Config
	MyDependency *Layer2Object
}

// Configuration for the Layer 2 Object
type Layer2Config interface {
	Logger() Logger
}

// Layer 2 Object
type Layer2Object struct {
	MyConfig Layer2Config
}

func NewLayer2Object(config Layer2Config) *Layer2Object {
	return &Layer2Object{
		MyConfig: config,
	}
}

// Stub implementation to make the compiler happy
type Logger interface {
}
