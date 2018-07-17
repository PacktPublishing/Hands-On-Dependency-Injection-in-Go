package advantages_test

import (
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch06/02_advantages"
)

func ExampleWelcomeSenderV2_Send() {
	welcomeSender := advantages.NewWelcomeSenderV2(&advantages.Mailer{})
	welcomeSender.Send("me@home.com")
}
