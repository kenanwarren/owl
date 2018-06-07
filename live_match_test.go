package owl

import (
	"testing"
)

func TestGetLiveMatch(t *testing.T) {
	c, err := NewClient(false)
	if err != nil {
		t.Error("Error creating client")
	}
	match, err := c.GetLiveMatch()
	if err != nil {
		t.Error("Error contacting owl api for live match")
	}
	if len(match.Data.LiveMatch.LiveStatus) == 0 {
		t.Error("No live match status from owl api")
	}
}
