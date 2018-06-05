package owl

import (
	"testing"
)

func TestGetRanking(t *testing.T) {
	c, err := NewClient(false)
	if err != nil {
		t.Error("Error creating client")
	}
	ranking, err := c.GetRanking()
	if err != nil {
		t.Error("Error contacting owl api for ranking")
	}
	if len(ranking.Content) == 0 {
		t.Error("No ranking content returned from owl api")
	}
}
