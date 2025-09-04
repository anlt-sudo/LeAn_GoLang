package notifiers

import "fmt"

// SlackNotifier là một triển khai của Notifier.
type SlackNotifier struct {
	ChannelName string
}

// Send triển khai phương thức của interface Notifier.
func (s SlackNotifier) Send(message string) error {
	fmt.Printf("Sending Slack message to #%s: %s\n", s.ChannelName, message)
	return nil
}