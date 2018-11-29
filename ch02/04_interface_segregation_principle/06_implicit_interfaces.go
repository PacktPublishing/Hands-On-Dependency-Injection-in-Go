package isp

import (
	"fmt"
)

type Talker interface {
	SayHello() string
}

type Dog struct{}

// The method implicitly implements the Talker interface
func (d Dog) SayHello() string {
	return "Woof!"
}

func Speak() {
	var talker Talker
	talker = Dog{}

	fmt.Print(talker.SayHello())
}
