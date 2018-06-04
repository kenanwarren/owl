package owl

import (
	"fmt"
	"net/http"
)

type MapsResponse []struct {
	ID   string `json:"id"`
	Name struct {
		EnUS string `json:"en_US"`
		EsMX string `json:"es_MX"`
		PtBR string `json:"pt_BR"`
		DeDE string `json:"de_DE"`
		EnGB string `json:"en_GB"`
		EsES string `json:"es_ES"`
		FrFR string `json:"fr_FR"`
		ItIT string `json:"it_IT"`
		PlPL string `json:"pl_PL"`
		RuRU string `json:"ru_RU"`
		KoKR string `json:"ko_KR"`
		ZhTW string `json:"zh_TW"`
		ZhCN string `json:"zh_CN"`
		JaJP string `json:"ja_JP"`
	} `json:"name"`
	Background  string `json:"background"`
	Icon        string `json:"icon"`
	Type        string `json:"type"`
	Description struct {
	} `json:"description"`
	Thumbnail    string `json:"thumbnail"`
	EsportsAPIID string `json:"esportsApiId"`
}

// GetMaps gets list of competition maps from the OWL API
// Endpoint: GET /maps
func (c *Client) GetMaps() (*MapsResponse, error) {
	m := &MapsResponse{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/maps", c.baseURL), nil)
	if err != nil {
		return m, err
	}

	if err = c.sendRequest(req, m); err != nil {
		return m, err
	}

	return m, nil
}
