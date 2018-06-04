package owl

import (
	"fmt"
	"net/http"
)

// https://api.overwatchleague.com/v2/standings
// also a normal /standings with less information
// query param locale=en_US
type StandingsResponse struct {
	Data []struct {
		ID              int    `json:"id"`
		DivisionID      int    `json:"divisionId"`
		Name            string `json:"name"`
		AbbreviatedName string `json:"abbreviatedName"`
		Logo            struct {
			Alt struct {
				Svg string `json:"svg"`
				Png string `json:"png"`
			} `json:"alt"`
			Main struct {
				Svg string `json:"svg"`
				Png string `json:"png"`
			} `json:"main"`
			MainName struct {
				Svg string `json:"svg"`
				Png string `json:"png"`
			} `json:"mainName"`
		} `json:"logo"`
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
		League struct {
			MatchWin          int `json:"matchWin"`
			MatchLoss         int `json:"matchLoss"`
			MatchDraw         int `json:"matchDraw"`
			MatchBye          int `json:"matchBye"`
			GameWin           int `json:"gameWin"`
			GameLoss          int `json:"gameLoss"`
			GameTie           int `json:"gameTie"`
			GamePointsFor     int `json:"gamePointsFor"`
			GamePointsAgainst int `json:"gamePointsAgainst"`
			Placement         int `json:"placement"`
		} `json:"league"`
		Stages struct {
			Stage1 struct {
				MatchWin          int  `json:"matchWin"`
				MatchLoss         int  `json:"matchLoss"`
				MatchDraw         int  `json:"matchDraw"`
				MatchBye          int  `json:"matchBye"`
				GameWin           int  `json:"gameWin"`
				GameLoss          int  `json:"gameLoss"`
				GameTie           int  `json:"gameTie"`
				GamePointsFor     int  `json:"gamePointsFor"`
				GamePointsAgainst int  `json:"gamePointsAgainst"`
				Placement         int  `json:"placement"`
				IsPlayoffWinner   bool `json:"isPlayoffWinner"`
			} `json:"stage1"`
			Stage2 struct {
				MatchWin          int  `json:"matchWin"`
				MatchLoss         int  `json:"matchLoss"`
				MatchDraw         int  `json:"matchDraw"`
				MatchBye          int  `json:"matchBye"`
				GameWin           int  `json:"gameWin"`
				GameLoss          int  `json:"gameLoss"`
				GameTie           int  `json:"gameTie"`
				GamePointsFor     int  `json:"gamePointsFor"`
				GamePointsAgainst int  `json:"gamePointsAgainst"`
				Placement         int  `json:"placement"`
				IsPlayoffWinner   bool `json:"isPlayoffWinner"`
			} `json:"stage2"`
			Stage3 struct {
				MatchWin          int  `json:"matchWin"`
				MatchLoss         int  `json:"matchLoss"`
				MatchDraw         int  `json:"matchDraw"`
				MatchBye          int  `json:"matchBye"`
				GameWin           int  `json:"gameWin"`
				GameLoss          int  `json:"gameLoss"`
				GameTie           int  `json:"gameTie"`
				GamePointsFor     int  `json:"gamePointsFor"`
				GamePointsAgainst int  `json:"gamePointsAgainst"`
				Placement         int  `json:"placement"`
				IsPlayoffWinner   bool `json:"isPlayoffWinner"`
			} `json:"stage3"`
			Stage4 struct {
				MatchWin          int  `json:"matchWin"`
				MatchLoss         int  `json:"matchLoss"`
				MatchDraw         int  `json:"matchDraw"`
				MatchBye          int  `json:"matchBye"`
				GameWin           int  `json:"gameWin"`
				GameLoss          int  `json:"gameLoss"`
				GameTie           int  `json:"gameTie"`
				GamePointsFor     int  `json:"gamePointsFor"`
				GamePointsAgainst int  `json:"gamePointsAgainst"`
				Placement         int  `json:"placement"`
				IsPlayoffWinner   bool `json:"isPlayoffWinner"`
			} `json:"stage4"`
		} `json:"stages"`
		Preseason struct {
			MatchWin          int `json:"matchWin"`
			MatchLoss         int `json:"matchLoss"`
			MatchDraw         int `json:"matchDraw"`
			MatchBye          int `json:"matchBye"`
			GameWin           int `json:"gameWin"`
			GameLoss          int `json:"gameLoss"`
			GameTie           int `json:"gameTie"`
			GamePointsFor     int `json:"gamePointsFor"`
			GamePointsAgainst int `json:"gamePointsAgainst"`
			Placement         int `json:"placement"`
		} `json:"preseason"`
	} `json:"data"`
	Meta struct {
		Divisions []struct {
			ID     string `json:"id"`
			String string `json:"string"`
			Name   string `json:"name"`
			Abbrev string `json:"abbrev"`
		} `json:"divisions"`
		Strings struct {
			OwlStandingsTitle                               string `json:"owl.standings.title"`
			OwlStandingsDescription                         string `json:"owl.standings.description"`
			OwlStandingsLegendKey                           string `json:"owl.standings.legend.key"`
			OwlStandingsHeaderDivision                      string `json:"owl.standings.header.division"`
			OwlStandingsHeaderWon                           string `json:"owl.standings.header.won"`
			OwlStandingsHeaderLost                          string `json:"owl.standings.header.lost"`
			OwlStandingsHeaderMatchesPlayed                 string `json:"owl.standings.header.matches-played"`
			OwlStandingsHeaderMapRecord                     string `json:"owl.standings.header.map-record"`
			OwlStandingsHeaderMapDiff                       string `json:"owl.standings.header.map-diff"`
			OwlStandingsHeaderAltDivision                   string `json:"owl.standings.header.alt.division"`
			OwlStandingsHeaderAltMatchesPlayed              string `json:"owl.standings.header.alt.matches-played"`
			OwlStandingsHeaderAltMatchesWon                 string `json:"owl.standings.header.alt.matches-won"`
			OwlStandingsHeaderAltMatchesLost                string `json:"owl.standings.header.alt.matches-lost"`
			OwlStandingsHeaderAltMapRecord                  string `json:"owl.standings.header.alt.map-record"`
			OwlStandingsHeaderAltMatchesDiff                string `json:"owl.standings.header.alt.matches-diff"`
			OwlStandingsHeaderTitleStage                    string `json:"owl.standings.header.title.stage"`
			OwlScheduleStageStage1                          string `json:"owl.schedule.stage.stage1"`
			OwlScheduleStageStage2                          string `json:"owl.schedule.stage.stage2"`
			OwlScheduleStageStage3                          string `json:"owl.schedule.stage.stage3"`
			OwlScheduleStageStage4                          string `json:"owl.schedule.stage.stage4"`
			OwlStandingsHeaderTitleLeague                   string `json:"owl.standings.header.title.league"`
			OwlStandingsAbbreviatedDivisionAtlantic         string `json:"owl.standings.abbreviated-division.atlantic"`
			OwlStandingsAbbreviatedDivisionPacific          string `json:"owl.standings.abbreviated-division.pacific"`
			OwlStandingsStagePlayoffsCutoff                 string `json:"owl.standings.stage-playoffs-cutoff"`
			OwlStandingsLegendTitle                         string `json:"owl.standings.legend.title"`
			OwlStandingsLegendClinchPostseasonPlayoffsKey   string `json:"owl.standings.legend.clinch-postseason-playoffs.key"`
			OwlStandingsLegendClinchPostseasonPlayoffsValue string `json:"owl.standings.legend.clinch-postseason-playoffs.value"`
			OwlStandingsLegendClinchAtlanticKey             string `json:"owl.standings.legend.clinch-atlantic.key"`
			OwlStandingsLegendClinchAtlanticValue           string `json:"owl.standings.legend.clinch-atlantic.value"`
			OwlStandingsLegendClinchPacificKey              string `json:"owl.standings.legend.clinch-pacific.key"`
			OwlStandingsLegendClinchPacificValue            string `json:"owl.standings.legend.clinch-pacific.value"`
			OwlStandingsLegendClinchStagePlayoffsKey        string `json:"owl.standings.legend.clinch-stage-playoffs.key"`
			OwlStandingsLegendClinchStagePlayoffsValue      string `json:"owl.standings.legend.clinch-stage-playoffs.value"`
			OwlStandingsLegendStagePlayoffsWinnerKey        string `json:"owl.standings.legend.stage-playoffs-winner.key"`
			OwlStandingsLegendStagePlayoffsWinnerValue      string `json:"owl.standings.legend.stage-playoffs-winner.value"`
			OwlStandingsLegendPostseasonPlayoffsChampKey    string `json:"owl.standings.legend.postseason-playoffs-champ.key"`
			OwlStandingsLegendPostseasonPlayoffsChampValue  string `json:"owl.standings.legend.postseason-playoffs-champ.value"`
			OwlStandingsLegendTiedKey                       string `json:"owl.standings.legend.tied.key"`
			OwlStandingsLegendTiedValue                     string `json:"owl.standings.legend.tied.value"`
			OwlStandingsPlayoffFormatTitle                  string `json:"owl.standings.playoff-format.title"`
			OwlStandingsPlayoffFormatStageTitle             string `json:"owl.standings.playoff-format.stage.title"`
			OwlStandingsPlayoffFormatStageDescription       string `json:"owl.standings.playoff-format.stage.description"`
			OwlStandingsPlayoffFormatPostseasonTitle        string `json:"owl.standings.playoff-format.postseason.title"`
			OwlStandingsPlayoffFormatPostseasonDescription  string `json:"owl.standings.playoff-format.postseason.description"`
			OwlStandingsTieBreakingTitle                    string `json:"owl.standings.tie-breaking.title"`
			OwlStandingsTieBreakingDescription              string `json:"owl.standings.tie-breaking.description"`
		} `json:"strings"`
		Separators struct {
			Stages struct {
				Stage1 int `json:"stage1"`
				Stage2 int `json:"stage2"`
				Stage3 int `json:"stage3"`
				Stage4 int `json:"stage4"`
			} `json:"stages"`
			Stage    int `json:"stage"`
			Wildcard int `json:"wildcard"`
		} `json:"separators"`
	} `json:"meta"`
}

// GetStandings gets the current owl standings from the OWL API
// Endpoint: GET /v2/standings
func (c *Client) GetStandings() (*StandingsResponse, error) {
	s := &StandingsResponse{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v2/standings", c.baseURL), nil)
	if err != nil {
		return s, err
	}

	if err = c.sendRequest(req, s); err != nil {
		return s, err
	}

	return s, nil
}
