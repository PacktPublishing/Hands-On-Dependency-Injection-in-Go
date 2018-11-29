package disadvantages

type MyPersonLoader interface {
	Load(ID int) (*Person, error)
}
