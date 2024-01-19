package client

import (
	"fmt"
)

// APIError includes the response error from api
type APIError struct {
	StatusCode int    `json:"-"`
	ErrorCode  int    `json:"code"`
	Message    string `json:"message"`
}

func (ae APIError) Error() string {
	return fmt.Sprintf("Error - %d, message: %s", ae.ErrorCode, ae.Message)
}

func (ae APIError) Empty() bool {
	return ae.Message == ""
}

// HandleError returns any non-nil http-related error (creating the request,
// getting the response, decoding) if any. If the decoded apiError is non-zero
// the apiError is returned. Otherwise, no errors occurred, returns nil.
func HandleError(httpError error, apiError APIError) error {
	if httpError != nil {
		return httpError
	}

	if !apiError.Empty() {
		return apiError
	}

	return nil
}
