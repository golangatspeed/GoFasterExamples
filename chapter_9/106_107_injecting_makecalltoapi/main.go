package main

import "fmt"

// SendNotification sends a notification using the external notification API.
func SendNotification(provider func(string, string) error, to, notification string) error {
	if err := provider(to, notification); err != nil {
		return err
	}
	return nil
}

// helper to make the REST call to API to send notification
func makeCallToAPI(to, notification string) error {
	fmt.Println(to, notification)
	return nil
}

func main() {
	_ = SendNotification(makeCallToAPI, "Joe Blogs", "Hello there!")
}
