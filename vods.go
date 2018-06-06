package owl

import (
	"fmt"
	"net/http"
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
	Data []Video `json:"data"`
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
