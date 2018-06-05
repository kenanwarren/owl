package owl

import "fmt"

type ErrorResponse struct {
	StatusCode   int
	RequestURL   string
	ResponseBody string
}

// Error method implementation for ErrorResponse struct
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("Response Error:\nRequest URL: %s\nStatus Code: %d\nResponse Body: %s\n", r.RequestURL, r.StatusCode, r.ResponseBody)
}
