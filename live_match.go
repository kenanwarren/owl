package owl

import (
	"fmt"
	"net/http"
	"time"
)

// Need to document https://api.overwatchleague.com/live-match
// query params expand=team.content&locale=en-us
type LiveMatchResponse struct {
	Data struct {
		LiveMatch struct {
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
				Content            struct {
					Nid             int    `json:"nid"`
					Name            string `json:"name"`
					AbbreviatedName string `json:"abbreviatedName"`
					ID              int    `json:"id"`
					Description     string `json:"description"`
					Location        string `json:"location"`
					LogoLockup      string `json:"logoLockup"`
					TeamWebsite     string `json:"teamWebsite"`
					Colors          []struct {
						Name  string `json:"name"`
						Usage string `json:"usage"`
						Color struct {
							Color   string `json:"color"`
							Opacity int    `json:"opacity"`
						} `json:"color"`
					} `json:"colors"`
					Icons []struct {
						Png   string `json:"png"`
						Svg   string `json:"svg"`
						Usage string `json:"usage"`
					} `json:"icons"`
				} `json:"content"`
			} `json:"competitors"`
			Scores []struct {
				Value int `json:"value"`
			} `json:"scores"`
			ConclusionValue    int    `json:"conclusionValue"`
			ConclusionStrategy string `json:"conclusionStrategy"`
			State              string `json:"state"`
			Status             string `json:"status"`
			StatusReason       string `json:"statusReason"`
			Attributes         struct {
			} `json:"attributes"`
			Games []struct {
				ID         int   `json:"id"`
				Number     int   `json:"number"`
				Points     []int `json:"points,omitempty"`
				Attributes struct {
					InstanceID string `json:"instanceID"`
					Map        string `json:"map"`
					MapGUID    string `json:"mapGuid"`
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
				Handle            string      `json:"handle,omitempty"`
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
						Featured           bool     `json:"featured"`
						Draft              bool     `json:"draft"`
						Handle             string   `json:"handle"`
						Title              string   `json:"title"`
						Series             struct {
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
			TimeToMatch int    `json:"timeToMatch"`
			LiveStatus  string `json:"liveStatus"`
		} `json:"liveMatch"`
		NextMatch struct {
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
				Content            struct {
					Nid             int    `json:"nid"`
					Name            string `json:"name"`
					AbbreviatedName string `json:"abbreviatedName"`
					ID              int    `json:"id"`
					Description     string `json:"description"`
					Location        string `json:"location"`
					LogoLockup      string `json:"logoLockup"`
					TeamWebsite     string `json:"teamWebsite"`
					Colors          []struct {
						Name  string `json:"name"`
						Usage string `json:"usage"`
						Color struct {
							Color   string `json:"color"`
							Opacity int    `json:"opacity"`
						} `json:"color"`
					} `json:"colors"`
					Icons []struct {
						Png   string `json:"png"`
						Svg   string `json:"svg"`
						Usage string `json:"usage"`
					} `json:"icons"`
				} `json:"content"`
			} `json:"competitors"`
			Scores []struct {
				Value int `json:"value"`
			} `json:"scores"`
			ConclusionValue    int    `json:"conclusionValue"`
			ConclusionStrategy string `json:"conclusionStrategy"`
			State              string `json:"state"`
			Status             string `json:"status"`
			StatusReason       string `json:"statusReason"`
			Attributes         struct {
			} `json:"attributes"`
			Games []struct {
				ID         int `json:"id"`
				Number     int `json:"number"`
				Attributes struct {
					Map     string `json:"map"`
					MapGUID string `json:"mapGuid"`
				} `json:"attributes"`
				AttributesVersion string      `json:"attributesVersion"`
				State             string      `json:"state"`
				Status            string      `json:"status"`
				StatusReason      string      `json:"statusReason"`
				Stats             interface{} `json:"stats"`
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
						Featured           bool     `json:"featured"`
						Draft              bool     `json:"draft"`
						Handle             string   `json:"handle"`
						Title              string   `json:"title"`
						Series             struct {
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
			Videos        []interface{} `json:"videos"`
			TimeToMatch   int           `json:"timeToMatch"`
			LiveStatus    string        `json:"liveStatus"`
		} `json:"nextMatch"`
	} `json:"data"`
	Meta struct {
		Strings struct {
			OwlLiveMatchUpcoming              string `json:"owl.live-match.upcoming"`
			OwlLiveMatchNow                   string `json:"owl.live-match.now"`
			OwlLiveStreamGame                 string `json:"owl.live-stream.game"`
			OwlLiveStreamGameOf               string `json:"owl.live-stream.game-of"`
			OwlLiveStreamResults              string `json:"owl.live-stream.results"`
			OwlLiveStreamViewMore             string `json:"owl.live-stream.view-more"`
			OwlLiveStreamViewLess             string `json:"owl.live-stream.view-less"`
			OwlLiveStreamPhoenixDisclaimer    string `json:"owl.live-stream.phoenix-disclaimer"`
			OwlRewardsInfoDescription         string `json:"owl.rewards.info.description"`
			OwlRewardsInfoDescriptionLoggedIn string `json:"owl.rewards.info.description.logged-in"`
		} `json:"strings"`
	} `json:"meta"`
}

// GetLiveMatch gets the current live match if one is in progress
// from the owl api.
// Endpoint: GET /live-match
// query params expand=team.content&locale=en-us
func (c *Client) GetLiveMatch() (*LiveMatchResponse, error) {
	l := &LiveMatchResponse{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/live-match", c.baseURL), nil)
	if err != nil {
		return l, err
	}

	if err = c.sendRequest(req, l); err != nil {
		return l, err
	}

	return l, nil
}
