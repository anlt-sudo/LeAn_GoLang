package notification

import "fmt"


type Notifier interface {
	Send(message string) error
}

func SendNotification(n Notifier, message string) {
	err := n.Send(message)
	if err != nil {
		fmt.Println("Error sending notification:", err)
	}
}


func ProcessNotifications(notifiers []interface{}, message string) {
	fmt.Println("\n--- Processing mixed notifications ---")
	for _, notifier := range notifiers {
		switch v := notifier.(type) {
		case Notifier:
			v.Send(message)
		case int:
			fmt.Printf("Skipping notification for number: %d\n", v)
		default:
			fmt.Printf("Unsupported notifier type: %T\n", v)
		}
	}
}