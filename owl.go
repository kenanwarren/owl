package owl

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	client  *http.Client
	baseURL string
}

// NewClient returns new Client struct
func NewClient(useChinaAPI bool) (*Client, error) {
	baseURL := OverwatchAPIEndpoint
	if useChinaAPI {
		baseURL = ChineseOverwatchAPIEndpoint
	}

	return &Client{
		client:  &http.Client{},
		baseURL: baseURL,
	}, nil
}

func (c *Client) sendRequest(req *http.Request, v interface{}) (err error) {
	req.Header.Set("Accept", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return &ErrorResponse{
			StatusCode:   resp.StatusCode,
			RequestURL:   req.URL.String(),
			ResponseBody: string(body),
		}
	}

	err = json.Unmarshal(body, v)
	return
}
