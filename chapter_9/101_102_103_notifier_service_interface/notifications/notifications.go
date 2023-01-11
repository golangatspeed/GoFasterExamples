package notifications

// NotifyService provides for sending notifications using external API
type NotifyService struct {}

// Notify sends notification
func (n *NotifyService) Notify (to, notification string) error {
	if err := makeCallToAPI(to, notification); err != nil {
		return err
	}
	return nil
}

// helper to make the REST call to API to send notification
func makeCallToAPI(to, notification string) error {
	return nil
}
