package owl

import (
	"fmt"
	"net/http"
)

type RankingResponse struct {
	Content []struct {
		Competitor Competitor `json:"competitor"`
		Placement  int        `json:"placement"`
		Advantage  int        `json:"advantage"`
		Records    []struct {
			MatchWin          int `json:"matchWin"`
			MatchLoss         int `json:"matchLoss"`
			MatchDraw         int `json:"matchDraw"`
			MatchBye          int `json:"matchBye"`
			GameWin           int `json:"gameWin"`
			GameLoss          int `json:"gameLoss"`
			GameTie           int `json:"gameTie"`
			GamePointsFor     int `json:"gamePointsFor"`
			GamePointsAgainst int `json:"gamePointsAgainst"`
		} `json:"records"`
	} `json:"content"`
	TotalMatches     int `json:"totalMatches"`
	MatchesConcluded int `json:"matchesConcluded"`
	PlayoffCutoff    int `json:"playoffCutoff"`
}

// GetRanking gets the current ranking from the OWL API
// Endpoint: GET /ranking
func (c *Client) GetRanking() (*RankingResponse, error) {
	r := &RankingResponse{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/ranking", c.baseURL), nil)
	if err != nil {
		return r, err
	}

	if err = c.sendRequest(req, r); err != nil {
		return r, err
	}

	return r, nil
}
