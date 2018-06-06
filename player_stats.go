package owl

import (
	"fmt"
	"net/http"
)

type AllPlayerStatsResponse struct {
	Data []struct {
		PlayerID                 int     `json:"playerId"`
		TeamID                   int     `json:"teamId"`
		Role                     string  `json:"role"`
		Name                     string  `json:"name"`
		Team                     string  `json:"team"`
		EliminationsAvgPer10M    float64 `json:"eliminations_avg_per_10m"`
		DeathsAvgPer10M          float64 `json:"deaths_avg_per_10m"`
		DamageAvgPer10M          float64 `json:"damage_avg_per_10m"`
		HealingAvgPer10M         float64 `json:"healing_avg_per_10m"`
		UltimatesEarnedAvgPer10M float64 `json:"ultimates_earned_avg_per_10m"`
		FinalBlowsAvgPer10M      float64 `json:"final_blows_avg_per_10m"`
		TimePlayedTotal          float64 `json:"time_played_total"`
	} `json:"data"`
}

// GetAllPlayerStats gets the list of player stats from the OWL API
// Endpoint: GET /stats/players
// TODO Add local=en-US query params
func (c *Client) GetAllPlayerStats() (*AllPlayerStatsResponse, error) {
	a := &AllPlayerStatsResponse{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/stats/players", c.baseURL), nil)
	if err != nil {
		return a, err
	}

	if err = c.sendRequest(req, a); err != nil {
		return a, err
	}

	return a, nil
}
