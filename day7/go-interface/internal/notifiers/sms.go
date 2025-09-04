package notifiers

import "fmt"

// SMSNotifier là một triển khai của Notifier.
type SMSNotifier struct {
	PhoneNumber string
}

// Send triển khai phương thức của interface Notifier.
func (s SMSNotifier) Send(message string) error {
	fmt.Printf("Sending SMS to %s: %s\n", s.PhoneNumber, message)
	return nil
}