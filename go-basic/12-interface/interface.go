package main

import "fmt"

// Notifier is an interface that defines a common method for sending notifications.
type Notifier interface {
	SendNotification(message string) error
}

// EmailNotifier is an implementation of the Notifier interface for sending email notifications.
type EmailNotifier struct {
	// Additional fields specific to email notifications can be added here
}

func (e *EmailNotifier) SendNotification(message string) error {
	// Implementation-specific logic to send an email
	fmt.Println("Sending email:", message)
	return nil
}

// SMSNotifier is an implementation of the Notifier interface for sending SMS notifications.
type SMSNotifier struct {
	// Additional fields specific to SMS notifications can be added here
}

func (s *SMSNotifier) SendNotification(message string) error {
	// Implementation-specific logic to send an SMS
	fmt.Println("Sending SMS:", message)
	return nil
}

func main() {
	// Depending on the scenario, you can switch between different notifiers
	emailNotifier := &EmailNotifier{}
	smsNotifier := &SMSNotifier{}

	// Send notifications using email
	sendNotification(emailNotifier, "Hello from Email!")

	// Send notifications using SMS
	sendNotification(smsNotifier, "Hello from SMS!")
}

// sendNotification is a generic function that takes any type implementing Notifier and sends a notification.
func sendNotification(notifier Notifier, message string) {
	if err := notifier.SendNotification(message); err != nil {
		fmt.Println("Error sending notification:", err)
	}
}
