package client

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dghubble/sling"
)

// Client represents a client for interacting with an API
type Client struct {
	slingClient *sling.Sling // Sling client for making requests
	apiKey      string       // API key for authentication
	privateKey  string       // Private key for signing JWT tokens
}

func New(apiKey, privateKey, baseURL string, httpClient *http.Client) *Client {
	s := sling.
		New().
		Client(httpClient).
		Base(baseURL).
		Set("X-API-Key", apiKey).
		Set("Accept", "application/json").
		Set("Content-Type", "application/json")

	return &Client{slingClient: s, apiKey: apiKey, privateKey: privateKey}
}

// MakeRequest makes an API request with the specified method, path, and body
func (x *Client) MakeRequest(method string, path string, body any, apiSuccess any) (*http.Response, error) {
	jsonBody, err := jsonify(body)
	if err != nil {
		return nil, err
	}

	privKey, err := ParsePrivateKey(x.privateKey)
	if err != nil {
		return nil, err
	}

	token, err := createAndSignJWTToken(privKey, x.apiKey, path, jsonBody)
	if err != nil {
		return nil, errors.New("error signing JWT token" + err.Error())
	}

	// Create a new Sling client with the Authorization header set
	sClient := x.slingClient.Set("Authorization", fmt.Sprintf("Bearer %v", token)).New()

	var (
		apiError     APIError
		httpResponse *http.Response
	)

	switch strings.ToUpper(method) {
	case "DELETE":
		httpResponse, err = sClient.Delete(path).Receive(apiSuccess, &apiError)
	case "GET":
		httpResponse, err = sClient.Get(path).Receive(apiSuccess, &apiError)
	case "POST":
		httpResponse, err = sClient.Post(path).Receive(apiSuccess, &apiError)
	case "PUT":
		httpResponse, err = sClient.Put(path).Receive(apiSuccess, &apiError)
	default:
		httpResponse, err = nil, errors.New("undefined method")
	}

	return httpResponse, HandleError(err, apiError)
}
