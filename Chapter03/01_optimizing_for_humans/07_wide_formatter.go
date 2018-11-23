package humans

type WideFormatter interface {
	ToCSV(pets []Pet) ([]byte, error)
	ToGOB(pets []Pet) ([]byte, error)
	ToJSON(pets []Pet) ([]byte, error)
}
