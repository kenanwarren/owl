package owl

import (
	"testing"
)

func TestGetNews(t *testing.T) {
	c, err := NewClient(false)
	if err != nil {
		t.Error("Error creating client")
	}
	news, err := c.GetNews()
	if err != nil {
		t.Error("Error contacting owl api for news")
	}
	if len(news.Blogs) == 0 {
		t.Error("No news posts returned from owl api")
	}
}
