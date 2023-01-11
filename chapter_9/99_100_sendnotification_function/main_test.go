package main

import "testing"

func TestSendNotification(t *testing.T) {
	err := SendNotification("Test Joe Blogs", "Test message")
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}
