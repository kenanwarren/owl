package owl

import (
	"testing"
)

func TestGetMatch(t *testing.T) {
	c, err := NewClient(false)
	if err != nil {
		t.Error("Error creating client")
	}
	match, err := c.GetMatch("10229")
	if err != nil {
		t.Error("Error contacting owl api for match")
	}
	if len(match.Competitors) == 0 {
		t.Error("No competitors in match returned from owl api")
	}
}
