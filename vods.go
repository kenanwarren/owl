package owl

import (
	"fmt"
	"net/http"
	"time"
)

type VodsResponse struct {
	Status   string `json:"status"`
	Code     int    `json:"code"`
	Metadata struct {
		TotalCount int    `json:"total_count"`
		Page       string `json:"page"`
		PerPage    string `json:"per_page"`
		PageCount  int    `json:"page_count"`
	} `json:"_metadata"`
	Data []struct {
		UniqueID         string      `json:"unique_id"`
		UserID           int         `json:"user_id"`
		OrganizationID   string      `json:"organization_id"`
		VideoType        string      `json:"video_type"`
		Title            string      `json:"title"`
		Description      string      `json:"description"`
		Length           int         `json:"length"`
		Privacy          interface{} `json:"privacy"`
		Status           string      `json:"status"`
		Available        int         `json:"available"`
		Thumbnail        string      `json:"thumbnail"`
		VideoThumbnailID int         `json:"video_thumbnail_id"`
		Location         string      `json:"location"`
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
		VideoAssets      []struct {
			ID         int         `json:"id"`
			UniqueID   string      `json:"unique_id"`
			VideoID    int         `json:"video_id"`
			Source     string      `json:"source"`
			SourceID   string      `json:"source_id"`
			SourceURL  string      `json:"source_url"`
			SourceHost string      `json:"source_host"`
			SourcePath string      `json:"source_path"`
			SourceData interface{} `json:"source_data"`
			Format     string      `json:"format"`
			AssetType  string      `json:"asset_type"`
			Status     string      `json:"status"`
			Available  interface{} `json:"available"`
			CreatedAt  time.Time   `json:"created_at"`
			UpdatedAt  time.Time   `json:"updated_at"`
		} `json:"video_assets"`
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
		Tags []struct {
			ID        int       `json:"id"`
			TagTypeID int       `json:"tag_type_id"`
			TagType   string    `json:"tag_type"`
			TagValue  string    `json:"tag_value"`
			CreatedAt time.Time `json:"created_at"`
		} `json:"tags"`
		Embed string `json:"embed"`
	} `json:"data"`
}

// GetVods gets all VODs from the OWL API
// Endpoint: GET /vods
// TODO Add Query Params
// per_page=50
// page=1
// tag=esports-match-10223
// locale=en-us
func (c *Client) GetVods() (*VodsResponse, error) {
	v := &VodsResponse{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/vods", c.baseURL), nil)
	if err != nil {
		return v, err
	}

	if err = c.sendRequest(req, v); err != nil {
		return v, err
	}

	return v, nil
}
