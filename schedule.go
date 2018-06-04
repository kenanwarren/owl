package owl

import (
	"fmt"
	"net/http"
	"time"
)

type ScheduleResponse struct {
	Data struct {
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		Stages    []struct {
			ID          int    `json:"id"`
			Slug        string `json:"slug"`
			Enabled     bool   `json:"enabled"`
			Name        string `json:"name"`
			Tournaments []struct {
				ID   int    `json:"id"`
				Type string `json:"type"`
			} `json:"tournaments"`
			Matches []struct {
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
					AbbreviatedName    string   `json:"abbreviatedName"`
					AddressCountry     string   `json:"addressCountry"`
					Logo               string   `json:"logo"`
					Icon               string   `json:"icon"`
					SecondaryPhoto     string   `json:"secondaryPhoto"`
					Type               string   `json:"type"`
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
					AbbreviatedName    string   `json:"abbreviatedName"`
					AddressCountry     string   `json:"addressCountry"`
					Logo               string   `json:"logo"`
					Icon               string   `json:"icon"`
					SecondaryPhoto     string   `json:"secondaryPhoto"`
					Type               string   `json:"type"`
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
					AttributesVersion string      `json:"attributesVersion"`
					State             string      `json:"state"`
					Status            string      `json:"status"`
					StatusReason      string      `json:"statusReason"`
					Stats             interface{} `json:"stats"`
					Handle            string      `json:"handle"`
				} `json:"games"`
				ClientHints   []interface{} `json:"clientHints"`
				DateCreated   int64         `json:"dateCreated"`
				Flags         []interface{} `json:"flags"`
				Handle        string        `json:"handle"`
				StartDate     time.Time     `json:"startDate"`
				EndDate       time.Time     `json:"endDate"`
				ShowStartTime bool          `json:"showStartTime"`
				ShowEndTime   bool          `json:"showEndTime"`
				StartDateTS   int64         `json:"startDateTS"`
				EndDateTS     int64         `json:"endDateTS"`
				YoutubeID     string        `json:"youtubeId"`
				Wins          []int         `json:"wins"`
				Ties          []int         `json:"ties"`
				Losses        []int         `json:"losses"`
				Videos        []struct {
					Name         int    `json:"name"`
					Description  string `json:"description"`
					VodLink      string `json:"vodLink"`
					YoutubeID    string `json:"youtubeId"`
					ThumbnailURL string `json:"thumbnailUrl"`
				} `json:"videos"`
				Tournament struct {
					ID   int    `json:"id"`
					Type string `json:"type"`
				} `json:"tournament"`
			} `json:"matches"`
			Weeks []struct {
				ID        int   `json:"id"`
				StartDate int64 `json:"startDate"`
				EndDate   int64 `json:"endDate"`
				Matches   []struct {
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
						AbbreviatedName    string   `json:"abbreviatedName"`
						AddressCountry     string   `json:"addressCountry"`
						Logo               string   `json:"logo"`
						Icon               string   `json:"icon"`
						SecondaryPhoto     string   `json:"secondaryPhoto"`
						Type               string   `json:"type"`
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
						AbbreviatedName    string   `json:"abbreviatedName"`
						AddressCountry     string   `json:"addressCountry"`
						Logo               string   `json:"logo"`
						Icon               string   `json:"icon"`
						SecondaryPhoto     string   `json:"secondaryPhoto"`
						Type               string   `json:"type"`
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
						AttributesVersion string      `json:"attributesVersion"`
						State             string      `json:"state"`
						Status            string      `json:"status"`
						StatusReason      string      `json:"statusReason"`
						Stats             interface{} `json:"stats"`
						Handle            string      `json:"handle"`
					} `json:"games"`
					ClientHints   []interface{} `json:"clientHints"`
					DateCreated   int64         `json:"dateCreated"`
					Flags         []interface{} `json:"flags"`
					Handle        string        `json:"handle"`
					StartDate     time.Time     `json:"startDate"`
					EndDate       time.Time     `json:"endDate"`
					ShowStartTime bool          `json:"showStartTime"`
					ShowEndTime   bool          `json:"showEndTime"`
					StartDateTS   int64         `json:"startDateTS"`
					EndDateTS     int64         `json:"endDateTS"`
					YoutubeID     string        `json:"youtubeId"`
					Wins          []int         `json:"wins"`
					Ties          []int         `json:"ties"`
					Losses        []int         `json:"losses"`
					Videos        []struct {
						Name         int    `json:"name"`
						Description  string `json:"description"`
						VodLink      string `json:"vodLink"`
						YoutubeID    string `json:"youtubeId"`
						ThumbnailURL string `json:"thumbnailUrl"`
					} `json:"videos"`
					Tournament struct {
						ID   int    `json:"id"`
						Type string `json:"type"`
					} `json:"tournament"`
				} `json:"matches"`
				Name string `json:"name"`
			} `json:"weeks"`
		} `json:"stages"`
	} `json:"data"`
	Meta struct {
		Strings struct {
			OwlScheduleTitle          string `json:"owl.schedule.title"`
			OwlScheduleDropdownTeams  string `json:"owl.schedule.dropdown.teams"`
			OwlScheduleStagePreseason string `json:"owl.schedule.stage.preseason"`
			OwlScheduleStageStage1    string `json:"owl.schedule.stage.stage1"`
			OwlScheduleStageStage2    string `json:"owl.schedule.stage.stage2"`
			OwlScheduleStageStage3    string `json:"owl.schedule.stage.stage3"`
			OwlScheduleStageStage4    string `json:"owl.schedule.stage.stage4"`
			OwlScheduleStagePlayoffs  string `json:"owl.schedule.stage.playoffs"`
			OwlScheduleStageFinals    string `json:"owl.schedule.stage.finals"`
			OwlScheduleStageAllStar   string `json:"owl.schedule.stage.all-star"`
			OwlScheduleDisclaimer1    string `json:"owl.schedule.disclaimer1"`
			OwlScheduleDisclaimer2    string `json:"owl.schedule.disclaimer2"`
			OwlScheduleSpoilers       string `json:"owl.schedule.spoilers"`
			OwlScheduleLive           string `json:"owl.schedule.live"`
			OwlScheduleMatchDetails   string `json:"owl.schedule.match-details"`
			OwlScheduleStageEnd       string `json:"owl.schedule.stage-end"`
			OwlScheduleToday          string `json:"owl.schedule.today"`
			OwlSchedulePrevWeek       string `json:"owl.schedule.prev-week"`
			OwlScheduleNextWeek       string `json:"owl.schedule.next-week"`
			OwlScheduleTitleMatch     string `json:"owl.schedule.title-match"`
			OwlSchedulePrizePool      string `json:"owl.schedule.prize-pool"`
			OwlScheduleWeek           string `json:"owl.schedule.week"`
		} `json:"strings"`
	} `json:"meta"`
}

// GetSchedule gets the owl schedule from the OWL API
// Endpoint: GET /schedule
func (c *Client) GetSchedule() (*ScheduleResponse, error) {
	s := &ScheduleResponse{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/schedule", c.baseURL), nil)
	if err != nil {
		return s, err
	}

	if err = c.sendRequest(req, s); err != nil {
		return s, err
	}

	return s, nil
}
