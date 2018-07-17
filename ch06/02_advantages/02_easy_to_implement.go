package advantages

func NewWelcomeSenderV2(mailer *Mailer) *WelcomeSenderV2 {
	return &WelcomeSenderV2{
		mailer: mailer,
	}
}

// WelcomeSenderV2 sends a Welcome email to new users
type WelcomeSenderV2 struct {
	mailer *Mailer
}

func (w *WelcomeSenderV2) Send(to string) error {
	body := w.buildMessage()

	return w.mailer.Send(to, body)
}

// build and return the message body
func (w *WelcomeSenderV2) buildMessage() string {
	return ""
}
