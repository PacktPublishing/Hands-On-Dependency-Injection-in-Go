package advantages_test

import (
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/02_advantages"
)

func ExampleWelcomeSender_Send() {
	welcomeSender := &advantages.WelcomeSender{
		Mailer: &advantages.Mailer{},
	}
	welcomeSender.Send("me@home.com")
}
