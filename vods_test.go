package owl

import (
	"testing"
)

func TestGetVods(t *testing.T) {
	c, err := NewClient(false)
	if err != nil {
		t.Error("Error creating client")
	}
	vods, err := c.GetVods()
	if err != nil {
		t.Error("Error contacting owl api for vods")
	}
	if len(vods.Data) == 0 {
		t.Error("No data for vods returned from owl api")
	}
}
