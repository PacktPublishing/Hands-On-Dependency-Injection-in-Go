package humans

import (
	"time"
)

type Pet struct {
	Name string
	Dog  bool
	Born time.Time
}

func NewPet(name string, isDog bool) Pet {
	return Pet{
		Name: name,
		Dog:  isDog,
		Born: time.Now(),
	}
}

func CreatePetsV1() {
	NewPet("Fido", true)
}
