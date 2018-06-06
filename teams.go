package owl

import (
	"fmt"
	"net/http"
)

type TeamsResponse struct {
	Data []Team `json:"data"`
	Meta struct {
		Count      int `json:"count"`
		Page       int `json:"page"`
		PageSize   int `json:"pageSize"`
		TotalPages int `json:"totalPages"`
	} `json:"meta"`
}

// GetTeams gets the list of teams from the OWL API
// Endpoint: GET /v2/teams
// TODO Add expand=team.content query params
func (c *Client) GetTeams() (*TeamsResponse, error) {
	t := &TeamsResponse{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v2/teams", c.baseURL), nil)
	if err != nil {
		return t, err
	}

	if err = c.sendRequest(req, t); err != nil {
		return t, err
	}

	return t, nil
}
