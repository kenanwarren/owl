package owl

import (
	"testing"
)

func TestGetPlayer(t *testing.T) {
	c, err := NewClient(false)
	if err != nil {
		t.Error("Error creating client")
	}
	player, err := c.GetPlayer("5197")
	if err != nil {
		t.Error("Error contacting owl api for player")
	}
	if len(player.Data.Player.Name) == 0 {
		t.Error("No data for player name returned from owl api")
	}
}
