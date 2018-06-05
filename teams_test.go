package owl

import (
	"testing"
)

func TestGetTeams(t *testing.T) {
	c, err := NewClient(false)
	if err != nil {
		t.Error("Error creating client")
	}
	teams, err := c.GetTeams()
	if err != nil {
		println(err.Error())
		t.Error("Error contacting owl api for teams")
	}
	if len(teams.Data) == 0 {
		t.Error("No data for teams returned from owl api")
	}
}
