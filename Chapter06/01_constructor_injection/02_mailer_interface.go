package constructor_injection

// Mailer sends and receives emails
type MailerInterface interface {
	Send(to string, body string) error
	Receive(address string) (string, error)
}
