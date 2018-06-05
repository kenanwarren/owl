package owl

import (
	"testing"
)

func TestGetStreams(t *testing.T) {
	c, err := NewClient(false)
	if err != nil {
		t.Error("Error creating client")
	}
	streams, err := c.GetStreams()
	if err != nil {
		t.Error("Error contacting owl api for streams")
	}
	if len(streams.EmbedURL) == 0 {
		t.Error("No embed url for streams returned from owl api")
	}
}
