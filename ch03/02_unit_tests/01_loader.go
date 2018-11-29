package unit_tests

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Loader interface {
	Load(ID int) (*Pet, error)
}

func TestLoadAndPrint_happyPath(t *testing.T) {
	result := &bytes.Buffer{}
	LoadAndPrint(&happyPathLoader{}, 1, result)
	assert.Contains(t, result.String(), "Pet named")
}

func TestLoadAndPrint_notFound(t *testing.T) {
	result := &bytes.Buffer{}
	LoadAndPrint(&missingLoader{}, 1, result)
	assert.Contains(t, result.String(), "no such pet")
}

func TestLoadAndPrint_error(t *testing.T) {
	result := &bytes.Buffer{}
	LoadAndPrint(&errorLoader{}, 1, result)
	assert.Contains(t, result.String(), "failed to load")
}

func LoadAndPrint(loader Loader, ID int, dest io.Writer) {
	loadedPet, err := loader.Load(ID)
	if err != nil {
		fmt.Fprintf(dest, "failed to load pet with ID %d with error: %s", ID, err)
		return
	}

	if loadedPet == nil {
		fmt.Fprintf(dest, "no such pet found")
		return
	}

	fmt.Fprintf(dest, "Pet named %s loaded", loadedPet.Name)
}

// implements Loader
type happyPathLoader struct {
}

func (l *happyPathLoader) Load(ID int) (*Pet, error) {
	return &Pet{}, nil
}

// implements Loader
type missingLoader struct {
}

func (l *missingLoader) Load(ID int) (*Pet, error) {
	return nil, nil
}

// implements Loader
type errorLoader struct {
}

func (l *errorLoader) Load(ID int) (*Pet, error) {
	return nil, errors.New("failed")
}
