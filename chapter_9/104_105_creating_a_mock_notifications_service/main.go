package main

import (
	"service/notifications"
	"service/service"
)

// SendNotification sends a notification using a Notifier service
func SendNotification(svc service.Notifier, to, notification string) error {
	return svc.Notify(to, notification)
}

func main() {
	// create a service
	notifySvc := &notifications.NotifyService{}
	_ = SendNotification(notifySvc, "Joe Blogs", "Hello there!")
}g