package constructor_injection

import "fmt"

type Talker interface {
	Speak() string
	Shout() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Shout() string {
	return "WOOF!"
}

func SpeakExample() {
	var talker Talker
	talker = Dog{}

	fmt.Print(talker.Speak())
}
