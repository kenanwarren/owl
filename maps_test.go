package owl

import (
	"testing"
)

func TestGetMaps(t *testing.T) {
	c, err := NewClient(false)
	if err != nil {
		t.Error("Error creating client")
	}
	maps, err := c.GetMaps()
	if err != nil {
		t.Error("Error contacting owl api for maps")
	}
	if maps == nil || len(*maps) == 0 {
		t.Error("No maps returned for matches from owl api")
	}
}
