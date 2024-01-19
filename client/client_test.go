package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	// Test basic client creation
	client := New("test-api-key", "test-private-key", "http://localhost:8080", nil)
	assert.Equal(t, "test-api-key", client.apiKey)
	assert.Equal(t, "test-private-key", client.privateKey)
}

func TestMakeRequest_Success(t *testing.T) {
	type APISuccessResponse struct {
		Message string
	}

	wantAPISuccessResponse := APISuccessResponse{
		Message: "this is an api error message",
	}

	// Set up a mock server for testing
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(wantAPISuccessResponse)
		w.Write(b)
	}))
	defer ts.Close()

	var gotAPISuccessResponse APISuccessResponse

	client := New(ConfigApiKey, string(ConfigPrivateKey), ts.URL, nil)
	httpResponse, err := client.MakeRequest("GET", "/test-path", nil, &gotAPISuccessResponse)
	assert.NoError(t, err)
	assert.Equal(t, httpResponse.StatusCode, http.StatusOK)
	assert.Equal(t, wantAPISuccessResponse, gotAPISuccessResponse)
}

func TestMakeRequest_Error(t *testing.T) {
	wantAPIError := APIError{
		StatusCode: 0,
		ErrorCode:  10001,
		Message:    "an error occured",
	}

	// Set up a mock server for testing
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		b, _ := json.Marshal(wantAPIError)
		w.Write(b)
	}))
	defer ts.Close()

	client := New(ConfigApiKey, string(ConfigPrivateKey), ts.URL, nil)

	httpResponse, err := client.MakeRequest("GET", "/test-path", nil, nil)
	assert.Equal(t, httpResponse.StatusCode, http.StatusBadRequest)
	assert.Error(t, err)

	// ensure the error returned was rightly parsed + correct type
	gotAPIError, ok := err.(APIError)
	assert.True(t, ok)
	assert.Equal(t, wantAPIError, gotAPIError)

	// Test INVALID-METHOD request with body
	response, err := client.MakeRequest("INVALID-METHOD", "/test-path", nil, nil)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestMakeRequest_Live_Fireblocks_Sandbox_Server(t *testing.T) {
	client := New(ConfigApiKey, string(ConfigPrivateKey), ConfigBaseURL, nil)

	var (
		vaultAccountID = "1"
		path           = fmt.Sprintf("/v1/vault/accounts/%s", vaultAccountID)
		apiSuccess     any
	)

	_, err := client.MakeRequest("get", path, nil, &apiSuccess)
	assert.NoError(t, err)

	assert.NotEmpty(t, apiSuccess)
}
