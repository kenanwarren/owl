package owl

import (
	"testing"
)

func TestGetAppPlaylist(t *testing.T) {
	c, err := NewClient(false)
	if err != nil {
		t.Error("Error creating client")
	}
	_, err = c.GetAppPlaylist()
	if err != nil {
		t.Error("Error contacting owl api for app playlist")
	}
}
