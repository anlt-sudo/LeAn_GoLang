package main

import (
	"fmt"
	"go-interface/internal/notification"
	"go-interface/internal/notifiers"
)

//Dùng để test cho trường hợp không hợp lệ
type InvalidNotifier struct{}

func main() {
	emailNotifier := notifiers.EmailNotifier{Recipient: "test@example.com"}
	smsNotifier := notifiers.SMSNotifier{PhoneNumber: "123-456-7890"}
	slackNotifier := notifiers.SlackNotifier{ChannelName: "general"}

	notifiersSlice := []notification.Notifier{emailNotifier, smsNotifier, slackNotifier}

	message := "Your order has been shipped!"

	fmt.Println("--- Sending notifications via Notifier slice ---")
	for _, notifier := range notifiersSlice {
		notification.SendNotification(notifier, message)
	}

	mixedNotifiers := []interface{}{
		emailNotifier,
		123,
		smsNotifier,
		InvalidNotifier{},
		slackNotifier,
	}

	notification.ProcessNotifications(mixedNotifiers, "System will be down for maintenance.")
}