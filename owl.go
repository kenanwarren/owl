package owl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	client  *http.Client
	baseURL string
}

// NewClient returns new Client struct
func NewClient(useChinaAPI bool) (*Client, error) {
	baseURL := "https://api.overwatchleague.com"
	if useChinaAPI {
		baseURL = "https://api.overwatchleague.cn"
	}

	return &Client{
		client:  &http.Client{},
		baseURL: baseURL,
	}, nil
}

func (c *Client) sendRequest(req *http.Request, v interface{}) (err error) {
	var (
		resp *http.Response
		data []byte
	)
	req.Header.Set("Accept", "application/json")

	resp, err = c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		errResp := &ErrorResponse{
			StatusCode: resp.StatusCode,
			RequestURL: req.URL.String(),
		}
		data, err = ioutil.ReadAll(resp.Body)

		if err == nil && len(data) > 0 {
			errResp.ResponseBody = string(data)
		}
		return errResp
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, v)
	return
}

type ErrorResponse struct {
	StatusCode   int
	RequestURL   string
	ResponseBody string
}

// Error method implementation for ErrorResponse struct
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("Response Error:\nRequest URL: %s\nStatus Code: %i\nResponse Body: %s\n", r.StatusCode, r.ResponseBody)
}
