package owl

import (
	"testing"
)

func TestGetMatches(t *testing.T) {
	c, err := NewClient(false)
	if err != nil {
		t.Error("Error creating client")
	}
	matches, err := c.GetMatches()
	if err != nil {
		t.Log(err)
		t.Error("Error contacting owl api for matches")
	}
	if len(matches.Content) == 0 {
		t.Error("No content returned for matches from owl api")
	}

	t.Log(matches.Content[0].StartDate)
}
