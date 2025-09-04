package notifiers

import "fmt"

type SlackNotifier struct {
	ChannelName string
}

func (s SlackNotifier) Send(message string) error {
	fmt.Printf("Sending Slack message to #%s: %s\n", s.ChannelName, message)
	return nil
}