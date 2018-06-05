package owl

import (
	"testing"
)

func TestGetPost(t *testing.T) {
	c, err := NewClient(false)
	if err != nil {
		t.Error("Error creating client")
	}
	post, err := c.GetNewsPost("21838051")
	if err != nil {
		t.Error("Error contacting owl api for post")
	}
	if len(post.Content) == 0 {
		t.Error("No post content returned from owl api")
	}
}
