package owl

import (
	"fmt"
	"net/http"
	"time"
)

type StreamsResponse struct {
	Platform   string `json:"platform"`
	EmbedURL   string `json:"embedUrl"`
	Status     int    `json:"status"`
	Attributes struct {
		Stream struct {
			ID          int64     `json:"_id"`
			Game        string    `json:"game"`
			Viewers     int       `json:"viewers"`
			VideoHeight int       `json:"video_height"`
			AverageFps  float64   `json:"average_fps"`
			Delay       int       `json:"delay"`
			CreatedAt   time.Time `json:"created_at"`
			IsPlaylist  bool      `json:"is_playlist"`
			StreamType  string    `json:"stream_type"`
			Preview     struct {
				Small    string `json:"small"`
				Medium   string `json:"medium"`
				Large    string `json:"large"`
				Template string `json:"template"`
			} `json:"preview"`
			Channel struct {
				Mature                       bool        `json:"mature"`
				Partner                      bool        `json:"partner"`
				Status                       string      `json:"status"`
				BroadcasterLanguage          string      `json:"broadcaster_language"`
				DisplayName                  string      `json:"display_name"`
				Game                         string      `json:"game"`
				Language                     string      `json:"language"`
				ID                           int         `json:"_id"`
				Name                         string      `json:"name"`
				CreatedAt                    time.Time   `json:"created_at"`
				UpdatedAt                    time.Time   `json:"updated_at"`
				Delay                        interface{} `json:"delay"`
				Logo                         string      `json:"logo"`
				Banner                       interface{} `json:"banner"`
				VideoBanner                  string      `json:"video_banner"`
				Background                   interface{} `json:"background"`
				ProfileBanner                string      `json:"profile_banner"`
				ProfileBannerBackgroundColor string      `json:"profile_banner_background_color"`
				URL                          string      `json:"url"`
				Views                        int         `json:"views"`
				Followers                    int         `json:"followers"`
				Links                        struct {
					Self          string `json:"self"`
					Follows       string `json:"follows"`
					Commercial    string `json:"commercial"`
					StreamKey     string `json:"stream_key"`
					Chat          string `json:"chat"`
					Features      string `json:"features"`
					Subscriptions string `json:"subscriptions"`
					Editors       string `json:"editors"`
					Teams         string `json:"teams"`
					Videos        string `json:"videos"`
				} `json:"_links"`
			} `json:"channel"`
			Links struct {
				Self string `json:"self"`
			} `json:"_links"`
		} `json:"stream"`
		Links struct {
			Self    string `json:"self"`
			Channel string `json:"channel"`
		} `json:"_links"`
		ChannelName string `json:"channelName"`
		StreamID    int    `json:"streamId"`
	} `json:"attributes"`
}

// GetStreams gets the list of current owl streams from the OWL API
// Endpoint: GET /v2/streams
func (c *Client) GetStreams() (*StreamsResponse, error) {
	s := &StreamsResponse{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v2/streams", c.baseURL), nil)
	if err != nil {
		return s, err
	}

	if err = c.sendRequest(req, s); err != nil {
		return s, err
	}

	return s, nil
}
