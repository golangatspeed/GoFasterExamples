package notifications

// MockNotifyService provides for sending test notifications
type MockNotifyService struct {
	sent [][]string
}

// Notify mocks sending a notification
func (mn *MockNotifyService) Notify (to string, notification string) error {
	mn.sent = append(mn.sent, []string{to, notification})
	return nil
}