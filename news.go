package owl

import (
	"fmt"
	"net/http"
)

type NewsResponse struct {
	TotalBlogs int                `json:"totalBlogs"`
	PageSize   int                `json:"pageSize"`
	Page       int                `json:"page"`
	TotalPages int                `json:"totalPages"`
	Blogs      []NewsPostResponse `json:"blogs"`
}

// GetNews gets all news from the OWL API
// Endpoint: GET /news
// TODO add query param pageSize=50&page=1
func (c *Client) GetNews() (*NewsResponse, error) {
	n := &NewsResponse{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/news", c.baseURL), nil)
	if err != nil {
		return n, err
	}

	if err = c.sendRequest(req, n); err != nil {
		return n, err
	}

	return n, nil
}
