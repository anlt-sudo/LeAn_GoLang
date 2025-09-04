package main

import (
	"go-interface/internal/notification"
	"go-interface/internal/notifiers"
	"fmt"
)

// Struct không liên quan để kiểm thử phần nâng cao
type InvalidNotifier struct{}

func main() {
	// Tạo các instance từ package 'notifiers'
	emailNotifier := notifiers.EmailNotifier{Recipient: "test@example.com"}
	smsNotifier := notifiers.SMSNotifier{PhoneNumber: "123-456-7890"}
	slackNotifier := notifiers.SlackNotifier{ChannelName: "general"}

	// Tạo slice Notifier. Kiểu Notifier được import từ package 'notification'.
	notifiersSlice := []notification.Notifier{emailNotifier, smsNotifier, slackNotifier}

	message := "Your order has been shipped!"

	// Gửi thông báo qua tất cả các kênh
	fmt.Println("--- Sending notifications via Notifier slice ---")
	for _, notifier := range notifiersSlice {
		// Gọi hàm xử lý từ package 'notification'
		notification.SendNotification(notifier, message)
	}

	// Kiểm thử phần nâng cao với interface{}
	mixedNotifiers := []interface{}{
		emailNotifier,
		123, // Thêm một số int
		smsNotifier,
		InvalidNotifier{}, // Thêm một struct không liên quan
		slackNotifier,
	}

	// Gọi hàm xử lý từ package 'notification'
	notification.ProcessNotifications(mixedNotifiers, "System will be down for maintenance.")
}