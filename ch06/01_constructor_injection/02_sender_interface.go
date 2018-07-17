package constructor_injection

type Sender interface {
	Send(to string, body string) error
}

func NewWelcomeSenderV2(in Sender) *WelcomeSenderV2 {
	return &WelcomeSenderV2{
		sender: in,
	}
}

// WelcomeSenderV2 sends a Welcome email to new users
type WelcomeSenderV2 struct {
	sender Sender
}

func (w *WelcomeSenderV2) Send(to string) error {
	body := w.buildMessage()

	return w.sender.Send(to, body)
}

// build and return the message body
func (w *WelcomeSenderV2) buildMessage() string {
	return ""
}
