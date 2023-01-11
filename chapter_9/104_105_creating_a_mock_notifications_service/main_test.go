package main

import (
	"service/notifications"
	"testing"
)

func TestSendNotification(t *testing.T) {
	tstSvc := &notifications.MockNotifyService{}
	if err := SendNotification(tstSvc, "Test Joe Blogs", "Test message"); err != nil {
		t.Fatalf("Something went wrong, unexpected error: %v", err)
	}
	if len(tstSvc.sent) != 1 {
		t.Errorf("expected 1 notification, got %d", len(tstSvc.sent))
	}

	// check that we stored the expected data
}


