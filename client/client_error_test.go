package client

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIError_Error(t *testing.T) {
	ae := APIError{
		ErrorCode: 10001, Message: "Invalid input data",
	}
	expected := "Error - 10001, message: Invalid input data"
	actual := ae.Error()
	assert.Equal(t, expected, actual)
}

func TestAPIError_Empty(t *testing.T) {
	AE := APIError{}
	assert.True(t, AE.Empty())

	AE.Message = "Error message"
	assert.False(t, AE.Empty())
}

func TestHandleError(t *testing.T) {
	httpError := fmt.Errorf("HTTP request failed")
	apiError := APIError{ErrorCode: 10001, Message: "Server error"}

	// HTTP error takes precedence
	actual := HandleError(httpError, apiError)
	if actual != httpError {
		t.Errorf("Expected HTTP error to be returned, got: %v", actual)
	}

	// API error returned if HTTP error is nil
	actual = HandleError(nil, apiError)
	if actual != apiError {
		t.Errorf("Expected API error to be returned, got: %v", actual)
	}

	// Nil returned if both errors are nil
	actual = HandleError(nil, APIError{})
	if actual != nil {
		t.Errorf("Expected nil error to be returned, got: %v", actual)
	}
}
