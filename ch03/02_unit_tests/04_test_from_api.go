package unit_tests

import (
	"database/sql"
)

type PetSaver struct{}

// save the supplied pet and return the ID
func (p PetSaver) Save(pet Pet) (int, error) {
	err := p.validate(pet)
	if err != nil {
		return 0, err
	}

	result, err := p.save(pet)
	if err != nil {
		return 0, err
	}

	return p.extractID(result)
}

// ensure the pet record is complete
func (p PetSaver) validate(pet Pet) error {
	return nil
}

// save to the datastore
func (p PetSaver) save(pet Pet) (sql.Result, error) {
	return nil, nil
}

// extract the ID from the result
func (p PetSaver) extractID(result sql.Result) (int, error) {
	return 0, nil
}
