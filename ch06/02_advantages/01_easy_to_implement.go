package advantages

// WelcomeSender sends a Welcome email to new users
type WelcomeSender struct {
	Mailer *Mailer
}

func (w *WelcomeSender) Send(to string) error {
	body := w.buildMessage()

	return w.Mailer.Send(to, body)
}

// build and return the message body
func (w *WelcomeSender) buildMessage() string {
	return ""
}

// Mailer will send an email
type Mailer struct{}

func (m *Mailer) Send(to string, body string) error {
	// send email
	return nil
}
