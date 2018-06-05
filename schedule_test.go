package owl

import (
	"testing"
)

func TestGetSchedule(t *testing.T) {
	c, err := NewClient(false)
	if err != nil {
		t.Error("Error creating client")
	}
	schedule, err := c.GetSchedule()
	if err != nil {
		t.Error("Error contacting owl api for schedule")
	}
	if len(schedule.Data.Stages) == 0 {
		t.Error("No data for schedule stages returned from owl api")
	}
}
