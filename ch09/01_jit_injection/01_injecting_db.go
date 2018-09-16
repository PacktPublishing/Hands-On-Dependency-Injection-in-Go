package jit_injection

func NewMyLoadPersonLogic(ds DataSource) *MyLoadPersonLogic {
	return &MyLoadPersonLogic{
		dataSource: ds,
	}
}

type MyLoadPersonLogic struct {
	dataSource DataSource
}

// Load person by supplied ID
func (m *MyLoadPersonLogic) Load(ID int) (Person, error) {
	return m.dataSource.Load(ID)
}

type DataSource interface {
	// Load person by ID
	Load(ID int) (Person, error)
}

type Person struct {
	Name string
}
