package owl

import (
	"strconv"
	"time"
)

type Game struct {
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
		Player    Player `json:"player"`
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
}

type Bracket struct {
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
}

type Player struct {
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
}

type Competitor struct {
	ID                 int      `json:"id"`
	AvailableLanguages []string `json:"availableLanguages"`
	Handle             string   `json:"handle"`
	Name               string   `json:"name"`
	HomeLocation       string   `json:"homeLocation"`
	PrimaryColor       string   `json:"primaryColor"`
	SecondaryColor     string   `json:"secondaryColor"`
	Game               string   `json:"game"`
	Divisions          []struct {
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
}

type Video struct {
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
	Tags             []struct {
		ID        int       `json:"id"`
		TagTypeID int       `json:"tag_type_id"`
		TagType   string    `json:"tag_type"`
		TagValue  string    `json:"tag_value"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"tags"`
	Thumbnails []struct {
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
	VideoAssets []struct {
		ID         int         `json:"id"`
		UniqueID   string      `json:"unique_id"`
		VideoID    int         `json:"video_id"`
		Source     string      `json:"source"`
		SourceID   string      `json:"source_id"`
		SourceURL  string      `json:"source_url"`
		SourceHost string      `json:"source_host"`
		SourcePath string      `json:"source_path"`
		SourceData string      `json:"source_data"`
		Format     string      `json:"format"`
		AssetType  string      `json:"asset_type"`
		Status     string      `json:"status"`
		Available  interface{} `json:"available"`
		CreatedAt  time.Time   `json:"created_at"`
		UpdatedAt  time.Time   `json:"updated_at"`
	}
}

type Team struct {
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
	HasFallback bool     `json:"hasFallback"`
	Location    string   `json:"location"`
	Players     []Player `json:"players"`
	Colors      struct {
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
	Website   string `json:"website"`
	Placement int    `json:"placement"`
	Advantage int    `json:"advantage"`
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
}

type FlexibleTime struct {
	time.Time
}

func (f *FlexibleTime) UnmarshalJSON(in []byte) error {
	if err := (&f.Time).UnmarshalJSON(in); err == nil {
		return err
	}
	if asInt, err := strconv.Atoi(string(in)); err != nil {
		return err
	} else {
		s, n := asInt/1000, asInt%1000
		f.Time = time.Unix(int64(s), int64(n))
	}
	return nil
}
