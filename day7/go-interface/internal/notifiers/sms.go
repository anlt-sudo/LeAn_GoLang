package notifiers

import "fmt"

type SMSNotifier struct {
	PhoneNumber string
}

func (s SMSNotifier) Send(message string) error {
	fmt.Printf("Sending SMS to %s: %s\n", s.PhoneNumber, message)
	return nil
}