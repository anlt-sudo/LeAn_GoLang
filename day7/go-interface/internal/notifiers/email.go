package notifiers

import "fmt"

type EmailNotifier struct {
	Recipient string
}

func (e EmailNotifier) Send(message string) error {
	fmt.Printf("Sending email to %s: %s\n", e.Recipient, message)
	return nil
}