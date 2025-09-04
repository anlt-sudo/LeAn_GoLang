package notifiers

import "fmt"

// EmailNotifier là một triển khai của Notifier.
type EmailNotifier struct {
	Recipient string
}

// Send triển khai phương thức của interface Notifier.
func (e EmailNotifier) Send(message string) error {
	fmt.Printf("Sending email to %s: %s\n", e.Recipient, message)
	return nil
}