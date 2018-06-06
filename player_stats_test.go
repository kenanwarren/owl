package owl

import (
	"testing"
)

func TestGetAllPlayerStats(t *testing.T) {
	c, err := NewClient(false)
	if err != nil {
		t.Error("Error creating client")
	}
	playerStats, err := c.GetAllPlayerStats()
	if err != nil {
		t.Error("Error contacting owl api for all player stats")
	}
	if len(playerStats.Data) == 0 {
		t.Error("No data for player stats returned from owl api")
	}
}
