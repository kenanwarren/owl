package owl

import "testing"

func TestNewClient(t *testing.T) {
	c, err := NewClient(false)
	if err != nil {
		t.Error("Error creating client.")
	}
	if c.baseURL != OverwatchAPIEndpoint {
		t.Error("Wrong api endpoint, expected nonchinese variant.")
	}
	if c.client == nil {
		t.Error("API HTTP client is nil")
	}
}

func TestNewChineseClient(t *testing.T) {
	c, err := NewClient(true)
	if err != nil {
		t.Error("Error creating client.")
	}
	if c.baseURL != ChineseOverwatchAPIEndpoint {
		t.Error("Wrong api endpoint, expected chinese variant.")
	}
	if c.client == nil {
		t.Error("API HTTP client is nil")
	}
}
