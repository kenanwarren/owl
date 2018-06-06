package owl

import (
	"fmt"
	"net/http"
)

type MatchResponse struct {
	ID          int          `json:"id"`
	Competitors []Competitor `json:"competitors"`
	Scores      []struct {
		Value int `json:"value"`
	} `json:"scores"`
	ConclusionValue    int    `json:"conclusionValue"`
	ConclusionStrategy string `json:"conclusionStrategy"`
	Winner             Player `json:"winner"`
	State              string `json:"state"`
	Status             string `json:"status"`
	StatusReason       string `json:"statusReason"`
	Attributes         struct {
	} `json:"attributes"`
	Games         []Game        `json:"games"`
	ClientHints   []interface{} `json:"clientHints"`
	Bracket       Bracket       `json:"bracket"`
	DateCreated   int64         `json:"dateCreated"`
	Flags         []interface{} `json:"flags"`
	Handle        string        `json:"handle"`
	StartDate     FlexibleTime  `json:"startDate"`
	EndDate       FlexibleTime  `json:"endDate"`
	ShowStartTime bool          `json:"showStartTime"`
	ShowEndTime   bool          `json:"showEndTime"`
	Rankings      []interface{} `json:"rankings"`
	Meta          struct {
		Strings struct {
			OwlMatchesCountdownUnitDays    string `json:"owl.matches.countdown.unit.days"`
			OwlMatchesCountdownUnitHours   string `json:"owl.matches.countdown.unit.hours"`
			OwlMatchesCountdownUnitMinutes string `json:"owl.matches.countdown.unit.minutes"`
			OwlMatchesMapsTitleUpcoming    string `json:"owl.matches.maps.title.upcoming"`
			OwlMatchesMapsTitleFinished    string `json:"owl.matches.maps.title.finished"`
			OwlMatchesGame                 string `json:"owl.matches.game"`
			OwlMatchesGameFinal            string `json:"owl.matches.game.final"`
			OwlMatchesGameFinalOt          string `json:"owl.matches.game.final-ot"`
			OwlMatchesGameOverview         string `json:"owl.matches.game.overview"`
			OwlMatchesGameType             string `json:"owl.matches.game.type"`
			OwlMatchesGameTypeMinimum      string `json:"owl.matches.game.type.minimum"`
			OwlMatchesGameTypeBestOf       string `json:"owl.matches.game.type.best-of"`
			OwlMatchesGameResult           string `json:"owl.matches.game.result"`
			OwlMatchesGameRoster           string `json:"owl.matches.game.roster"`
			OwlMatchesPaginationPrevious   string `json:"owl.matches.pagination.previous"`
			OwlMatchesPaginationVersus     string `json:"owl.matches.pagination.versus"`
			OwlMatchesPaginationNext       string `json:"owl.matches.pagination.next"`
			OwlPlayersRolesTank            string `json:"owl.players.roles.tank"`
			OwlPlayersRolesOffense         string `json:"owl.players.roles.offense"`
			OwlPlayersRolesDefense         string `json:"owl.players.roles.defense"`
			OwlPlayersRolesSupport         string `json:"owl.players.roles.support"`
			OwlPlayersRolesFlex            string `json:"owl.players.roles.flex"`
		} `json:"strings"`
	} `json:"meta"`
}

// GetMatch gets a match from the OWL API
// Endpoint: GET /matches/:id
func (c *Client) GetMatch(matchID string) (*MatchResponse, error) {
	m := &MatchResponse{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/matches/%s", c.baseURL, matchID), nil)
	if err != nil {
		return m, err
	}

	if err = c.sendRequest(req, m); err != nil {
		return m, err
	}

	return m, nil
}
