package main

// SendNotification sends a notification using the external notification API.
func SendNotification(to, notification string) error {
	if err := makeCallToAPI(to, notification); err != nil {
		return err
	}
	return nil
}

// helper to make the REST call to API to send notification
func makeCallToAPI(to, notification string) error {
	return nil
}

func main() {
	_ = SendNotification("Joe Blogs", "Hello there!")
}
