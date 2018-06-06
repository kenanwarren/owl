package owl

import (
	"fmt"
	"net/http"
)

type MatchesResponse struct {
	Content          []MatchResponse `json:"content"`
	Last             bool            `json:"last"`
	TotalElements    int             `json:"totalElements"`
	TotalPages       int             `json:"totalPages"`
	NumberOfElements int             `json:"numberOfElements"`
	Size             int             `json:"size"`
	Number           int             `json:"number"`
	Sort             interface{}     `json:"sort"`
	First            bool            `json:"first"`
}

// GetMatches gets all matches from the OWL API
// Endpoint: GET /matches
// TODO add query param size
func (c *Client) GetMatches() (*MatchesResponse, error) {
	m := &MatchesResponse{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/matches", c.baseURL), nil)
	if err != nil {
		return m, err
	}

	if err = c.sendRequest(req, m); err != nil {
		return m, err
	}

	return m, nil
}
