package main

import "testing"

var sent [][]string

func mockCallToAPI (to string, notification string) error {
	sent = make([][]string, 0, 0) // reset the slice
	sent = append(sent, []string{to, notification})
	return nil
}

func TestSendNotification(t *testing.T) {
	if err := SendNotification(mockCallToAPI, "Test Joe Blogs", "Test message"); err != nil {
		t.Fatalf("Something went wrong, unexpected error: %v", err)
	}
	if len(sent) != 1 {
		t.Errorf("expected 1 notification, got %d", len(sent))
	}
}