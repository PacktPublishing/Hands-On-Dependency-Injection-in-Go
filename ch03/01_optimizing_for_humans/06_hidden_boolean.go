package humans

const (
	isDog = true
	isCat = false
)

func NewDog(name string) Pet {
	return NewPet(name, isDog)
}

func NewCat(name string) Pet {
	return NewPet(name, isCat)
}

func CreatePetsV2() {
	NewDog("Fido")
}
