package owl

import (
	"fmt"
	"net/http"
	"time"
)

type PlayerResponse struct {
	Data struct {
		Player struct {
			ID                 int           `json:"id"`
			AvailableLanguages []string      `json:"availableLanguages"`
			Handle             string        `json:"handle"`
			Name               string        `json:"name"`
			HomeLocation       string        `json:"homeLocation"`
			Accounts           []interface{} `json:"accounts"`
			Game               string        `json:"game"`
			Attributes         struct {
				Heroes        []string `json:"heroes"`
				PlayerNumber  int      `json:"player_number"`
				PreferredSlot string   `json:"preferred_slot"`
				Role          string   `json:"role"`
			} `json:"attributes"`
			AttributesVersion string `json:"attributesVersion"`
			FamilyName        string `json:"familyName"`
			GivenName         string `json:"givenName"`
			Nationality       string `json:"nationality"`
			Headshot          string `json:"headshot"`
			Teams             []struct {
				Team struct {
					ID                 int      `json:"id"`
					AvailableLanguages []string `json:"availableLanguages"`
					Handle             string   `json:"handle"`
					Name               string   `json:"name"`
					HomeLocation       string   `json:"homeLocation"`
					PrimaryColor       string   `json:"primaryColor"`
					SecondaryColor     string   `json:"secondaryColor"`
					Accounts           []struct {
						ID           int    `json:"id"`
						CompetitorID int    `json:"competitorId"`
						Value        string `json:"value"`
						AccountType  string `json:"accountType"`
						IsPublic     bool   `json:"isPublic"`
					} `json:"accounts"`
					Game       string `json:"game"`
					Attributes struct {
						City      interface{} `json:"city"`
						HeroImage interface{} `json:"hero_image"`
						Manager   interface{} `json:"manager"`
						TeamGUID  string      `json:"team_guid"`
					} `json:"attributes"`
					AttributesVersion string `json:"attributesVersion"`
					Divisions         []struct {
						Competitor struct {
							ID int `json:"id"`
						} `json:"competitor"`
						Division struct {
							ID int `json:"id"`
						} `json:"division"`
					} `json:"divisions"`
					AbbreviatedName string `json:"abbreviatedName"`
					AddressCountry  string `json:"addressCountry"`
					Logo            string `json:"logo"`
					Icon            string `json:"icon"`
					SecondaryPhoto  string `json:"secondaryPhoto"`
					Type            string `json:"type"`
				} `json:"team"`
				Player struct {
					ID   int    `json:"id"`
					Type string `json:"type"`
				} `json:"player"`
				Flags []interface{} `json:"flags"`
			} `json:"teams"`
			Type string `json:"type"`
		} `json:"player"`
		SimilarPlayers []struct {
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
			Teams             []struct {
				Team struct {
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
					Divisions         []struct {
						Competitor struct {
							ID int `json:"id"`
						} `json:"competitor"`
						Division struct {
							ID int `json:"id"`
						} `json:"division"`
					} `json:"divisions"`
					AbbreviatedName string `json:"abbreviatedName"`
					AddressCountry  string `json:"addressCountry"`
					Logo            string `json:"logo"`
					Icon            string `json:"icon"`
					SecondaryPhoto  string `json:"secondaryPhoto"`
					Type            string `json:"type"`
				} `json:"team"`
				Player struct {
					ID   int    `json:"id"`
					Type string `json:"type"`
				} `json:"player"`
				Flags []interface{} `json:"flags"`
			} `json:"teams"`
			Type       string `json:"type"`
			Similarity struct {
				Score         int      `json:"score"`
				SameTopHeroes []string `json:"sameTopHeroes"`
			} `json:"similarity"`
		} `json:"similarPlayers"`
		Stats struct {
			All struct {
				EliminationsAvgPer10M    float64 `json:"eliminations_avg_per_10m"`
				DeathsAvgPer10M          float64 `json:"deaths_avg_per_10m"`
				DamageAvgPer10M          float64 `json:"damage_avg_per_10m"`
				HealingAvgPer10M         float64 `json:"healing_avg_per_10m"`
				UltimatesEarnedAvgPer10M float64 `json:"ultimates_earned_avg_per_10m"`
				FinalBlowsAvgPer10M      float64 `json:"final_blows_avg_per_10m"`
				TimePlayedTotal          float64 `json:"time_played_total"`
			} `json:"all"`
			Heroes []struct {
				HeroID string `json:"hero_id"`
				Name   string `json:"name"`
				Stats  struct {
					EliminationsAvgPer10M    float64 `json:"eliminations_avg_per_10m"`
					DeathsAvgPer10M          float64 `json:"deaths_avg_per_10m"`
					DamageAvgPer10M          float64 `json:"damage_avg_per_10m"`
					HealingAvgPer10M         float64 `json:"healing_avg_per_10m"`
					UltimatesEarnedAvgPer10M float64 `json:"ultimates_earned_avg_per_10m"`
					FinalBlowsAvgPer10M      float64 `json:"final_blows_avg_per_10m"`
					TimePlayedTotal          float64 `json:"time_played_total"`
				} `json:"stats"`
			} `json:"heroes"`
		} `json:"stats"`
		StatRanks struct {
			EliminationsAvgPer10M    int `json:"eliminations_avg_per_10m"`
			DeathsAvgPer10M          int `json:"deaths_avg_per_10m"`
			DamageAvgPer10M          int `json:"damage_avg_per_10m"`
			HealingAvgPer10M         int `json:"healing_avg_per_10m"`
			UltimatesEarnedAvgPer10M int `json:"ultimates_earned_avg_per_10m"`
			FinalBlowsAvgPer10M      int `json:"final_blows_avg_per_10m"`
		} `json:"statRanks"`
		Vods []struct {
			UniqueID         string      `json:"unique_id"`
			UserID           int         `json:"user_id"`
			OrganizationID   string      `json:"organization_id"`
			VideoType        string      `json:"video_type"`
			Title            string      `json:"title"`
			Description      interface{} `json:"description"`
			Length           int         `json:"length"`
			Privacy          interface{} `json:"privacy"`
			Status           string      `json:"status"`
			Available        int         `json:"available"`
			Thumbnail        string      `json:"thumbnail"`
			VideoThumbnailID int         `json:"video_thumbnail_id"`
			Location         interface{} `json:"location"`
			EveSlug          interface{} `json:"eve_slug"`
			StreamName       interface{} `json:"stream_name"`
			ChannelSlug      interface{} `json:"channel_slug"`
			ChannelID        interface{} `json:"channel_id"`
			Views            interface{} `json:"views"`
			Likes            interface{} `json:"likes"`
			Dislikes         interface{} `json:"dislikes"`
			Favorites        interface{} `json:"favorites"`
			PublishType      string      `json:"publish_type"`
			StatsAt          interface{} `json:"stats_at"`
			RecordedAt       interface{} `json:"recorded_at"`
			BroadcastedAt    interface{} `json:"broadcasted_at"`
			AvailableAt      time.Time   `json:"available_at"`
			CreatedAt        time.Time   `json:"created_at"`
			UpdatedAt        time.Time   `json:"updated_at"`
			ShareURL         string      `json:"share_url"`
			Thumbnails       []struct {
				ID                int         `json:"id"`
				VideoID           int         `json:"video_id"`
				SourceThumbnailID interface{} `json:"source_thumbnail_id"`
				ThumbnailType     string      `json:"thumbnail_type"`
				ThumbnailURL      string      `json:"thumbnail_url"`
				Size              string      `json:"size"`
				Width             interface{} `json:"width"`
				Height            interface{} `json:"height"`
				TimeAt            interface{} `json:"time_at"`
				CreatedAt         time.Time   `json:"created_at"`
				UpdatedAt         time.Time   `json:"updated_at"`
			} `json:"thumbnails"`
			Tags []struct {
				ID        int       `json:"id"`
				TagTypeID int       `json:"tag_type_id"`
				TagType   string    `json:"tag_type"`
				TagValue  string    `json:"tag_value"`
				CreatedAt time.Time `json:"created_at"`
			} `json:"tags"`
			VideoAssets []struct {
				ID         int         `json:"id"`
				UniqueID   string      `json:"unique_id"`
				VideoID    int         `json:"video_id"`
				Source     string      `json:"source"`
				SourceID   interface{} `json:"source_id"`
				SourceURL  string      `json:"source_url"`
				SourceHost string      `json:"source_host"`
				SourcePath string      `json:"source_path"`
				SourceData interface{} `json:"source_data"`
				Format     string      `json:"format"`
				AssetType  string      `json:"asset_type"`
				Status     string      `json:"status"`
				Available  interface{} `json:"available"`
				CreatedAt  time.Time   `json:"created_at"`
				UpdatedAt  time.Time   `json:"updated_at"`
			} `json:"video_assets"`
		} `json:"vods"`
		Team struct {
			Data struct {
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
					Alt struct {
						Svg string `json:"svg"`
						Png string `json:"png"`
					} `json:"alt"`
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
					Number       int    `json:"number"`
					Headshot     string `json:"headshot"`
					HomeLocation string `json:"homeLocation,omitempty"`
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
				Matches struct {
					Recent []struct {
						ID          int    `json:"id"`
						Handle      string `json:"handle"`
						Competitors []struct {
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
								Alt struct {
									Svg string `json:"svg"`
									Png string `json:"png"`
								} `json:"alt"`
								MainName struct {
									Svg string `json:"svg"`
									Png string `json:"png"`
								} `json:"mainName"`
							} `json:"logo"`
							HasFallback bool   `json:"hasFallback"`
							Location    string `json:"location"`
							Players     []struct {
								ID           int           `json:"id"`
								Handle       string        `json:"handle"`
								Name         string        `json:"name"`
								FullName     string        `json:"fullName"`
								Role         string        `json:"role"`
								Accounts     []interface{} `json:"accounts"`
								Number       int           `json:"number"`
								Headshot     string        `json:"headshot"`
								HomeLocation string        `json:"homeLocation,omitempty"`
							} `json:"players"`
							Colors struct {
								Secondary struct {
									Color   string `json:"color"`
									Opacity int    `json:"opacity"`
								} `json:"secondary"`
								Primary struct {
									Color   string `json:"color"`
									Opacity int    `json:"opacity"`
								} `json:"primary"`
								Tertiary struct {
									Color   string `json:"color"`
									Opacity int    `json:"opacity"`
								} `json:"tertiary"`
							} `json:"colors"`
							Accounts []interface{} `json:"accounts"`
							Website  string        `json:"website"`
						} `json:"competitors"`
						Scores []struct {
							Value int `json:"value"`
						} `json:"scores"`
						StartDate int64  `json:"startDate"`
						EndDate   int64  `json:"endDate"`
						State     string `json:"state"`
						Games     []struct {
							ID            int    `json:"id"`
							Number        int    `json:"number"`
							State         string `json:"state"`
							Map           string `json:"map"`
							DeprecatedMap string `json:"deprecatedMap"`
							Players       []struct {
								ID           int           `json:"id"`
								Handle       string        `json:"handle"`
								Name         string        `json:"name"`
								FullName     string        `json:"fullName"`
								Role         string        `json:"role"`
								Accounts     []interface{} `json:"accounts"`
								Number       int           `json:"number"`
								Headshot     string        `json:"headshot"`
								HomeLocation string        `json:"homeLocation,omitempty"`
							} `json:"players"`
						} `json:"games"`
					} `json:"recent"`
				} `json:"matches"`
				Placement int `json:"placement"`
				Advantage int `json:"advantage"`
				Records   struct {
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
			} `json:"data"`
			Meta struct {
			} `json:"meta"`
		} `json:"team"`
	} `json:"data"`
	Meta struct {
	} `json:"meta"`
}

// GetPlayer gets the player info given an id from the OWL API
// Endpoint: GET /players/:id
// TODO Add locale=en-us&expand=article,vods,stats,stat.ranks,team,team.matches.recent,similarPlayers
func (c *Client) GetPlayer(playerID string) (*PlayerResponse, error) {
	p := &PlayerResponse{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/players/%s", c.baseURL, playerID), nil)
	if err != nil {
		return p, err
	}

	if err = c.sendRequest(req, p); err != nil {
		return p, err
	}

	return p, nil
}
