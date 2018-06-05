package humans

type ThinFormatter interface {
	Format(pets []Pet) ([]byte, error)
}

type CSVFormatter struct{}

func (f CSVFormatter) Format(pets []Pet) ([]byte, error) {
	// convert slice of pets to CSV
	return nil, nil
}

type GOBFormatter struct{}

func (f GOBFormatter) Format(pets []Pet) ([]byte, error) {
	// convert slice of pets to GOB
	return nil, nil
}

type JSONFormatter struct{}

func (f JSONFormatter) Format(pets []Pet) ([]byte, error) {
	// convert slice of pets to JSON
	return nil, nil
}
