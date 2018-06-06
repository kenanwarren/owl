package owl

import (
	"fmt"
	"net/http"
	"time"
)

type AppPlaylistResponse struct {
	Status   string `json:"status"`
	Code     int    `json:"code"`
	Metadata struct {
		TotalCount int `json:"total_count"`
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		PageCount  int `json:"page_count"`
	} `json:"_metadata"`
	Data struct {
		ID             int         `json:"id"`
		UserID         int         `json:"user_id"`
		OrganizationID string      `json:"organization_id"`
		Name           string      `json:"name"`
		DisplayName    string      `json:"display_name"`
		PlaylistType   interface{} `json:"playlist_type"`
		PlayAs         string      `json:"play_as"`
		DisplayAs      string      `json:"display_as"`
		Thumbnail      interface{} `json:"thumbnail"`
		Available      int         `json:"available"`
		CreatedAt      time.Time   `json:"created_at"`
		UpdatedAt      time.Time   `json:"updated_at"`
		Videos         []Video     `json:"videos"`
	} `json:"data"`
}

// GetAppPlaylist gets the app video playlist from them OWL API
// Endpoint: GET /playlist/owl-app-playlist
// TODO Add query params per_page=50&page=1
func (c *Client) GetAppPlaylist() (*AppPlaylistResponse, error) {
	a := &AppPlaylistResponse{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/playlist/owl-app-playlist", c.baseURL), nil)
	if err != nil {
		return a, err
	}

	if err = c.sendRequest(req, a); err != nil {
		return a, err
	}

	return a, nil
}
