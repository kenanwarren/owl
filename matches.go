package owl

import (
	"fmt"
	"net/http"
)

type MatchesResponse struct {
	Content []struct {
		ID          int `json:"id"`
		Competitors []struct {
			ID                 int      `json:"id"`
			AvailableLanguages []string `json:"availableLanguages"`
			Handle             string   `json:"handle"`
			Name               string   `json:"name"`
			HomeLocation       string   `json:"homeLocation"`
			PrimaryColor       string   `json:"primaryColor"`
			SecondaryColor     string   `json:"secondaryColor"`
			Game               string   `json:"game"`
			Attributes         struct {
				City      interface{} `json:"city"`
				HeroImage interface{} `json:"hero_image"`
				Manager   interface{} `json:"manager"`
				TeamGUID  string      `json:"team_guid"`
			} `json:"attributes"`
			AttributesVersion string `json:"attributesVersion"`
			AbbreviatedName   string `json:"abbreviatedName"`
			AddressCountry    string `json:"addressCountry"`
			Logo              string `json:"logo"`
			Icon              string `json:"icon"`
			Players           []struct {
				Team struct {
					ID   int    `json:"id"`
					Type string `json:"type"`
				} `json:"team"`
				Player struct {
					ID                 int      `json:"id"`
					AvailableLanguages []string `json:"availableLanguages"`
					Handle             string   `json:"handle"`
					Name               string   `json:"name"`
					HomeLocation       string   `json:"homeLocation"`
					Game               string   `json:"game"`
					Attributes         struct {
						HeroImage     interface{} `json:"hero_image"`
						Heroes        []string    `json:"heroes"`
						Hometown      string      `json:"hometown"`
						PlayerNumber  int         `json:"player_number"`
						PreferredSlot string      `json:"preferred_slot"`
						Role          string      `json:"role"`
					} `json:"attributes"`
					AttributesVersion string `json:"attributesVersion"`
					FamilyName        string `json:"familyName"`
					GivenName         string `json:"givenName"`
					Nationality       string `json:"nationality"`
					Headshot          string `json:"headshot"`
					Type              string `json:"type"`
				} `json:"player"`
				Flags []interface{} `json:"flags"`
			} `json:"players"`
			SecondaryPhoto string `json:"secondaryPhoto"`
			Type           string `json:"type"`
		} `json:"competitors"`
		Scores []struct {
			Value int `json:"value"`
		} `json:"scores"`
		ConclusionValue    int    `json:"conclusionValue"`
		ConclusionStrategy string `json:"conclusionStrategy"`
		Winner             struct {
			ID                 int      `json:"id"`
			AvailableLanguages []string `json:"availableLanguages"`
			Handle             string   `json:"handle"`
			Name               string   `json:"name"`
			HomeLocation       string   `json:"homeLocation"`
			PrimaryColor       string   `json:"primaryColor"`
			SecondaryColor     string   `json:"secondaryColor"`
			Game               string   `json:"game"`
			Attributes         struct {
				City      interface{} `json:"city"`
				HeroImage interface{} `json:"hero_image"`
				Manager   interface{} `json:"manager"`
				TeamGUID  string      `json:"team_guid"`
			} `json:"attributes"`
			AttributesVersion string `json:"attributesVersion"`
			AbbreviatedName   string `json:"abbreviatedName"`
			AddressCountry    string `json:"addressCountry"`
			Logo              string `json:"logo"`
			Icon              string `json:"icon"`
			Players           []struct {
				Team struct {
					ID   int    `json:"id"`
					Type string `json:"type"`
				} `json:"team"`
				Player struct {
					ID                 int      `json:"id"`
					AvailableLanguages []string `json:"availableLanguages"`
					Handle             string   `json:"handle"`
					Name               string   `json:"name"`
					HomeLocation       string   `json:"homeLocation"`
					Game               string   `json:"game"`
					Attributes         struct {
						HeroImage     interface{} `json:"hero_image"`
						Heroes        []string    `json:"heroes"`
						Hometown      string      `json:"hometown"`
						PlayerNumber  int         `json:"player_number"`
						PreferredSlot string      `json:"preferred_slot"`
						Role          string      `json:"role"`
					} `json:"attributes"`
					AttributesVersion string `json:"attributesVersion"`
					FamilyName        string `json:"familyName"`
					GivenName         string `json:"givenName"`
					Nationality       string `json:"nationality"`
					Headshot          string `json:"headshot"`
					Type              string `json:"type"`
				} `json:"player"`
				Flags []interface{} `json:"flags"`
			} `json:"players"`
			SecondaryPhoto string `json:"secondaryPhoto"`
			Type           string `json:"type"`
		} `json:"winner"`
		State        string `json:"state"`
		Status       string `json:"status"`
		StatusReason string `json:"statusReason"`
		Attributes   struct {
		} `json:"attributes"`
		Games []struct {
			ID         int   `json:"id"`
			Number     int   `json:"number"`
			Points     []int `json:"points"`
			Attributes struct {
				InstanceID string `json:"instanceID"`
				Map        string `json:"map"`
				MapScore   struct {
					Team1 int `json:"team1"`
					Team2 int `json:"team2"`
				} `json:"mapScore"`
			} `json:"attributes"`
			AttributesVersion string `json:"attributesVersion"`
			Players           []struct {
				Team struct {
					ID   int    `json:"id"`
					Type string `json:"type"`
				} `json:"team"`
				Player struct {
					ID                 int      `json:"id"`
					AvailableLanguages []string `json:"availableLanguages"`
					Handle             string   `json:"handle"`
					Name               string   `json:"name"`
					HomeLocation       string   `json:"homeLocation"`
					Game               string   `json:"game"`
					Attributes         struct {
						HeroImage     interface{} `json:"hero_image"`
						Heroes        []string    `json:"heroes"`
						Hometown      interface{} `json:"hometown"`
						PlayerNumber  int         `json:"player_number"`
						PreferredSlot string      `json:"preferred_slot"`
						Role          string      `json:"role"`
					} `json:"attributes"`
					AttributesVersion string `json:"attributesVersion"`
					FamilyName        string `json:"familyName"`
					GivenName         string `json:"givenName"`
					Nationality       string `json:"nationality"`
					Headshot          string `json:"headshot"`
					Type              string `json:"type"`
				} `json:"player"`
				MatchGame struct {
					ID    int         `json:"id"`
					Stats interface{} `json:"stats"`
				} `json:"matchGame"`
			} `json:"players"`
			State        string      `json:"state"`
			Status       string      `json:"status"`
			StatusReason string      `json:"statusReason"`
			Stats        interface{} `json:"stats"`
			Handle       string      `json:"handle"`
		} `json:"games"`
		ClientHints []interface{} `json:"clientHints"`
		Bracket     struct {
			ID                      int    `json:"id"`
			MatchConclusionValue    int    `json:"matchConclusionValue"`
			MatchConclusionStrategy string `json:"matchConclusionStrategy"`
			Winners                 int    `json:"winners"`
			TeamSize                int    `json:"teamSize"`
			RepeatableMatchUps      int    `json:"repeatableMatchUps"`
			Stage                   struct {
				ID                 int      `json:"id"`
				AvailableLanguages []string `json:"availableLanguages"`
				Title              string   `json:"title"`
				Tournament         struct {
					ID                 int      `json:"id"`
					AvailableLanguages []string `json:"availableLanguages"`
					Game               string   `json:"game"`
					Location           string   `json:"location"`
					Featured           bool     `json:"featured"`
					Draft              bool     `json:"draft"`
					Handle             string   `json:"handle"`
					Title              string   `json:"title"`
					Attributes         struct {
						Program struct {
							Environment string `json:"environment"`
							Phase       string `json:"phase"`
							SeasonID    string `json:"season_id"`
							Stage       struct {
								Format string `json:"format"`
								Number int    `json:"number"`
							} `json:"stage"`
							Type string `json:"type"`
						} `json:"program"`
					} `json:"attributes"`
					AttributesVersion string `json:"attributesVersion"`
					SubEvents         []struct {
						ID                 int      `json:"id"`
						AvailableLanguages []string `json:"availableLanguages"`
						Name               string   `json:"name"`
						StartDate          int64    `json:"startDate"`
						EndDate            int64    `json:"endDate"`
						ShowStartTime      bool     `json:"showStartTime"`
						ShowEndTime        bool     `json:"showEndTime"`
					} `json:"subEvents"`
					Series struct {
						ID int `json:"id"`
					} `json:"series"`
				} `json:"tournament"`
			} `json:"stage"`
			Type               string        `json:"type"`
			ClientHints        []interface{} `json:"clientHints"`
			AdvantageComparing string        `json:"advantageComparing"`
			ThirdPlaceMatch    bool          `json:"thirdPlaceMatch"`
			AllowDraw          bool          `json:"allowDraw"`
		} `json:"bracket"`
		DateCreated      int64         `json:"dateCreated"`
		Flags            []interface{} `json:"flags"`
		Handle           string        `json:"handle"`
		StartDate        int64         `json:"startDate"`
		EndDate          int64         `json:"endDate"`
		ShowStartTime    bool          `json:"showStartTime"`
		ShowEndTime      bool          `json:"showEndTime"`
		Round            int           `json:"round,omitempty"`
		Ordinal          int           `json:"ordinal,omitempty"`
		WinnersNextMatch int           `json:"winnersNextMatch,omitempty"`
		WinnerRound      int           `json:"winnerRound,omitempty"`
		WinnerOrdinal    int           `json:"winnerOrdinal,omitempty"`
		BestOf           int           `json:"bestOf,omitempty"`
	} `json:"content"`
	Last             bool        `json:"last"`
	TotalElements    int         `json:"totalElements"`
	TotalPages       int         `json:"totalPages"`
	NumberOfElements int         `json:"numberOfElements"`
	Size             int         `json:"size"`
	Number           int         `json:"number"`
	Sort             interface{} `json:"sort"`
	First            bool        `json:"first"`
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
