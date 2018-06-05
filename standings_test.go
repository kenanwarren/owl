package owl

import (
	"testing"
)

func TestGetStandings(t *testing.T) {
	c, err := NewClient(false)
	if err != nil {
		t.Error("Error creating client")
	}
	standings, err := c.GetStandings()
	if err != nil {
		t.Error("Error contacting owl api for standings")
	}
	if len(standings.Data) == 0 {
		t.Error("No data for standings returned from owl api")
	}
}
