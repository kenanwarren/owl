package owl

import (
	"fmt"
	"net/http"
)

type TeamsResponse struct {
	Data []struct {
		ID              int    `json:"id"`
		DivisionID      int    `json:"divisionId"`
		Handle          string `json:"handle"`
		Name            string `json:"name"`
		AbbreviatedName string `json:"abbreviatedName"`
		Logo            struct {
			Main struct {
				Svg string `json:"svg"`
				Png string `json:"png"`
			} `json:"main"`
			MainName struct {
				Svg string `json:"svg"`
				Png string `json:"png"`
			} `json:"mainName"`
		} `json:"logo"`
		HasFallback bool   `json:"hasFallback"`
		Location    string `json:"location"`
		Players     []struct {
			ID       int    `json:"id"`
			Handle   string `json:"handle"`
			Name     string `json:"name"`
			FullName string `json:"fullName"`
			Role     string `json:"role"`
			Accounts []struct {
				ID   int    `json:"id"`
				Type string `json:"type"`
				URL  string `json:"url"`
			} `json:"accounts"`
			Number       FlexibleInteger `json:"number"` // Empty is "" ?!? use custom unmarshaler to handle
			Headshot     string          `json:"headshot"`
			HomeLocation string          `json:"homeLocation,omitempty"`
		} `json:"players"`
		Colors struct {
			Primary struct {
				Color   string `json:"color"`
				Opacity int    `json:"opacity"`
			} `json:"primary"`
			Secondary struct {
				Color   string `json:"color"`
				Opacity int    `json:"opacity"`
			} `json:"secondary"`
			Tertiary struct {
				Color   string `json:"color"`
				Opacity int    `json:"opacity"`
			} `json:"tertiary"`
		} `json:"colors"`
		Accounts []struct {
			ID   int    `json:"id"`
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"accounts"`
		Website string `json:"website"`
	} `json:"data"`
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
