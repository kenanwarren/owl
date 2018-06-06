package owl

import (
	"fmt"
	"net/http"
)

type TeamResponse struct {
	Data Team `json:"data"`
	Meta struct {
		Strings struct {
			OwlTeamHeaderAtlantic                  string `json:"owl.team.header.atlantic"`
			OwlTeamHeaderPacific                   string `json:"owl.team.header.pacific"`
			OwlTeamViewAll                         string `json:"owl.team.view-all"`
			OwlTeamsRecentMatchVods                string `json:"owl.teams.recent-match-vods"`
			OwlTeamHeaderSocial                    string `json:"owl.team.header.social"`
			OwlTeamHeaderW                         string `json:"owl.team.header.w"`
			OwlTeamHeaderL                         string `json:"owl.team.header.l"`
			OwlStandingsHeaderAltDivision          string `json:"owl.standings.header.alt.division"`
			OwlTeamLeagueStanding                  string `json:"owl.team.league-standing"`
			OwlTeamLocation                        string `json:"owl.team.location"`
			OwcContentModulesRoundRobinMatchRecord string `json:"owc.content-modules.round-robin.match-record"`
			OwlTeamsRosterTeamWebsite              string `json:"owl.teams.roster.team-website"`
			OwlTeamUpcomingMatches                 string `json:"owl.team.upcoming-matches"`
			OwlMatchesGameRoster                   string `json:"owl.matches.game.roster"`
			OwlContentModulesRecentNewsTitle       string `json:"owl.content-modules.recent-news.title"`
			OwlPlayersRolesOffense                 string `json:"owl.players.roles.offense"`
			OwlPlayersRolesDefense                 string `json:"owl.players.roles.defense"`
			OwlPlayersRolesTank                    string `json:"owl.players.roles.tank"`
			OwlPlayersRolesSupport                 string `json:"owl.players.roles.support"`
			OwlPlayersRolesFlex                    string `json:"owl.players.roles.flex"`
			OwlScheduleToday                       string `json:"owl.schedule.today"`
			OwlScheduleMatchDetails                string `json:"owl.schedule.match-details"`
			OwlLabelsNews                          string `json:"owl.labels.news"`
			OwlLabelsFeature                       string `json:"owl.labels.feature"`
			OwlLabelsAnnouncement                  string `json:"owl.labels.announcement"`
			OwlLabelsTraining                      string `json:"owl.labels.training"`
			OwlLabelsPathToPro                     string `json:"owl.labels.path-to-pro"`
			OwlLabelsHighlight                     string `json:"owl.labels.highlight"`
			OwlLabelsAnalysis                      string `json:"owl.labels.analysis"`
		} `json:"strings"`
	} `json:"meta"`
}

// GetTeam gets a team from the OWL API
// Endpoint: GET /teams/:id
func (c *Client) GetTeam(teamID string) (*TeamResponse, error) {
	t := &TeamResponse{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v2/teams/%s", c.baseURL, teamID), nil)
	if err != nil {
		return t, err
	}

	if err = c.sendRequest(req, t); err != nil {
		return t, err
	}

	return t, nil
}
