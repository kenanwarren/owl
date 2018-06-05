package owl

import (
	"testing"
)

func TestGetTeam(t *testing.T) {
	c, err := NewClient(false)
	if err != nil {
		t.Error("Error creating client")
	}
	team, err := c.GetTeam("4523")
	if err != nil {
		t.Error("Error contacting owl api for team")
	}
	if len(team.Data.Name) == 0 {
		t.Error("No data for team name returned from owl api")
	}
}
