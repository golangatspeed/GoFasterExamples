package service

// Notifier must be implemented for any service we use to send notifications.
type Notifier interface {
	Notify(to string, notification string) error
}
