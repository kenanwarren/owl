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
		Videos         []struct {
			UniqueID         string      `json:"unique_id"`
			UserID           int         `json:"user_id"`
			OrganizationID   string      `json:"organization_id"`
			VideoType        string      `json:"video_type"`
			Title            string      `json:"title"`
			Description      interface{} `json:"description"`
			Length           int         `json:"length"`
			Privacy          interface{} `json:"privacy"`
			Status           string      `json:"status"`
			Available        int         `json:"available"`
			Thumbnail        string      `json:"thumbnail"`
			VideoThumbnailID int         `json:"video_thumbnail_id"`
			Location         interface{} `json:"location"`
			EveSlug          interface{} `json:"eve_slug"`
			StreamName       interface{} `json:"stream_name"`
			ChannelSlug      interface{} `json:"channel_slug"`
			ChannelID        interface{} `json:"channel_id"`
			Views            interface{} `json:"views"`
			Likes            interface{} `json:"likes"`
			Dislikes         interface{} `json:"dislikes"`
			Favorites        interface{} `json:"favorites"`
			PublishType      string      `json:"publish_type"`
			StatsAt          interface{} `json:"stats_at"`
			RecordedAt       interface{} `json:"recorded_at"`
			BroadcastedAt    interface{} `json:"broadcasted_at"`
			AvailableAt      time.Time   `json:"available_at"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			ShareURL         string      `json:"share_url"`
			Tags             []struct {
				ID        int       `json:"id"`
				TagTypeID int       `json:"tag_type_id"`
				TagType   string    `json:"tag_type"`
				TagValue  string    `json:"tag_value"`
				CreatedAt time.Time `json:"created_at"`
			} `json:"tags"`
			Thumbnails []struct {
				ID                int         `json:"id"`
				VideoID           int         `json:"video_id"`
				SourceThumbnailID interface{} `json:"source_thumbnail_id"`
				ThumbnailType     string      `json:"thumbnail_type"`
				ThumbnailURL      string      `json:"thumbnail_url"`
				Size              string      `json:"size"`
				Width             interface{} `json:"width"`
				Height            interface{} `json:"height"`
				TimeAt            interface{} `json:"time_at"`
				CreatedAt         time.Time   `json:"created_at"`
				UpdatedAt         time.Time   `json:"updated_at"`
			} `json:"thumbnails"`
			VideoAssets []struct {
				ID         int         `json:"id"`
				UniqueID   string      `json:"unique_id"`
				VideoID    int         `json:"video_id"`
				Source     string      `json:"source"`
				SourceID   string      `json:"source_id"`
				SourceURL  string      `json:"source_url"`
				SourceHost string      `json:"source_host"`
				SourcePath string      `json:"source_path"`
				SourceData string      `json:"source_data"`
				Format     string      `json:"format"`
				AssetType  string      `json:"asset_type"`
				Status     string      `json:"status"`
				Available  interface{} `json:"available"`
				CreatedAt  time.Time   `json:"created_at"`
				UpdatedAt  time.Time   `json:"updated_at"`
			} `json:"video_assets"`
			Embed string `json:"embed"`
		} `json:"videos"`
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
