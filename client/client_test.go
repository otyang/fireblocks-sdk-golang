package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	t.Parallel()

	// Test basic client creation
	client := New("test-api-key", "test-private-key", "http://localhost:8080", ConfigDebugMode)
	assert.Equal(t, "test-api-key", client.apiKey)
	assert.Equal(t, "test-private-key", client.privateKey)
}

func TestMakeRequest_Success(t *testing.T) {
	t.Parallel()

	type APISuccessResponse struct {
		Message string
	}

	wantAPISuccessResponse := APISuccessResponse{
		Message: "this is an api error message",
	}

	// Set up a mock server for testing
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(wantAPISuccessResponse)
		if err != nil {
			panic(err)
		}
	}))
	defer ts.Close()

	var (
		path                  = "/test-path"
		body                  = map[string]any{"asset": "BTC"}
		gotAPISuccessResponse APISuccessResponse
		client                = New(ConfigApiKey, string(ConfigPrivateKey), ts.URL, ConfigDebugMode)
	)

	httpResponse, err := client.MakeRequest(context.Background(), "GET", path, body, &gotAPISuccessResponse)
	assert.NoError(t, err)
	assert.Equal(t, httpResponse.StatusCode(), http.StatusOK)
	assert.Equal(t, wantAPISuccessResponse, gotAPISuccessResponse)

	assert.Equal(t, httpResponse.Request.Method, "GET")
}

func TestMakeRequest_Error(t *testing.T) {
	t.Parallel()

	wantAPIError := APIError{
		StatusCode: 0,
		ErrorCode:  10001,
		Message:    "an error occured",
	}

	// Set up a mock server for testing
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(wantAPIError)
		if err != nil {
			panic(err)
		}
	}))
	defer ts.Close()

	var (
		path   = "/test-path"
		body   = map[string]any{"asset": "BTC"}
		client = New(ConfigApiKey, string(ConfigPrivateKey), ts.URL, ConfigDebugMode)
	)

	httpResponse, err := client.MakeRequest(context.Background(), "GET", path, body, nil)
	assert.Equal(t, httpResponse.StatusCode(), http.StatusBadRequest)
	assert.Error(t, err)
	assert.Equal(t, httpResponse.Request.Method, "GET")

	// ensure the error returned was rightly parsed + correct type
	gotAPIError, ok := err.(APIError)
	assert.True(t, ok)
	assert.Equal(t, wantAPIError, gotAPIError)

	// Test INVALID-METHOD request with body
	response, err := client.MakeRequest(context.Background(), "INVALID-METHOD", path, nil, nil)
	assert.Error(t, err)
	assert.Nil(t, response)
}

func TestMakeRequest_Live_Fireblocks_Sandbox_Server(t *testing.T) {
	t.Parallel()

	client := New(ConfigApiKey, ConfigPrivateKey, ConfigBaseURL, ConfigDebugMode)

	var (
		vaultAccountID = "1"
		path           = fmt.Sprintf("/v1/vault/accounts/%s", vaultAccountID)
		apiSuccess     any
	)

	_, err := client.MakeRequest(context.Background(), "get", path, nil, &apiSuccess)
	assert.NoError(t, err)
	assert.NotEmpty(t, apiSuccess)
}
