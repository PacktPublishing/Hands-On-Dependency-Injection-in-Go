package constructor_injection

import (
	"errors"
)

func NewWelcomeSender(in *Mailer) (*WelcomeSender, error) {
	// guard clause
	if in == nil {
		return nil, errors.New("programmer error: mailer must not provided")
	}

	return &WelcomeSender{
		mailer: in,
	}, nil
}

func NewWelcomeSenderNoGuard(in *Mailer) *WelcomeSender {
	return &WelcomeSender{
		mailer: in,
	}
}

// WelcomeSender sends a Welcome email to new users
type WelcomeSender struct {
	mailer *Mailer
}

func (w *WelcomeSender) Send(to string) error {
	body := w.buildMessage()

	return w.mailer.Send(to, body)
}

// build and return the message body
func (w *WelcomeSender) buildMessage() string {
	return ""
}

// Mailer sends and receives emails
type Mailer struct {
	Host     string
	Port     string
	Username string
	Password string
}

func (m *Mailer) Send(to string, body string) error {
	// send email
	return nil
}

func (m *Mailer) Receive(address string) (string, error) {
	// receive email
	return "", nil
}
