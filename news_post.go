package owl

import (
	"fmt"
	"net/http"
)

type NewsPostResponse struct {
	BlogID     int      `json:"blogId"`
	Created    int64    `json:"created"`
	Updated    int64    `json:"updated"`
	Publish    int64    `json:"publish"`
	Title      string   `json:"title"`
	Author     string   `json:"author"`
	Locale     string   `json:"locale"`
	Keywords   []string `json:"keywords"`
	Content    string   `json:"content"`
	Summary    string   `json:"summary"`
	CommentKey string   `json:"commentKey"`
	Status     string   `json:"status"`
	Thumbnail  struct {
		MediaID          int    `json:"mediaId"`
		URL              string `json:"url"`
		MimeType         string `json:"mimeType"`
		Type             string `json:"type"`
		Size             int    `json:"size"`
		Width            int    `json:"width"`
		Height           int    `json:"height"`
		OriginalFileName string `json:"originalFileName"`
		MediaType        int    `json:"mediaType"`
	} `json:"thumbnail"`
	Header struct {
		MediaID          int    `json:"mediaId"`
		URL              string `json:"url"`
		MimeType         string `json:"mimeType"`
		Type             string `json:"type"`
		Size             int    `json:"size"`
		Width            int    `json:"width"`
		Height           int    `json:"height"`
		OriginalFileName string `json:"originalFileName"`
		MediaType        int    `json:"mediaType"`
	} `json:"header"`
	DefaultCommunity string `json:"defaultCommunity"`
	DefaultURL       string `json:"defaultUrl"`
	CommentsEnabled  bool   `json:"commentsEnabled"`
	PollAttached     bool   `json:"pollAttached"`
	CommunityBlogs   []struct {
		CommunityID int  `json:"communityId"`
		Sticky      bool `json:"sticky"`
		Defaulted   bool `json:"defaulted"`
	} `json:"communityBlogs"`
	LocalizationPublish int64    `json:"localizationPublish"`
	SiteCategory        string   `json:"siteCategory"`
	Draft               bool     `json:"draft"`
	ShowPtr             bool     `json:"showPtr"`
	FeaturedMap         bool     `json:"featuredMap"`
	ShowRating          bool     `json:"showRating"`
	Spotlight           bool     `json:"spotlight"`
	CustomPublish       bool     `json:"customPublish"`
	Tags                []string `json:"tags"`
	StickyForThisLocale bool     `json:"stickyForThisLocale"`
}

// GetNewsPost gets a news post from the OWL API
// Endpoint: GET /news/:id
func (c *Client) GetNewsPost(newsPostID string) (*NewsPostResponse, error) {
	n := &NewsPostResponse{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/news/%s", c.baseURL, newsPostID), nil)
	if err != nil {
		return n, err
	}

	if err = c.sendRequest(req, n); err != nil {
		return n, err
	}

	return n, nil
}
