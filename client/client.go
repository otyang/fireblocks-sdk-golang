package client

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

// Client represents a client for interacting with an API
type Client struct {
	restyClient *resty.Client
	apiKey      string // API key for authentication
	privateKey  string // Private key for signing JWT tokens
	debugMode   bool
}

// New creates a new Client instance, configuring the base URL, headers, and API token.
func New(apiKey, privateKey, baseURL string, debugMode bool) *Client {
	// Create a new Resty client with the specified configurations.
	clnt := resty.New().
		SetBaseURL(baseURL).
		SetHeader("X-API-Key", apiKey).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json")

	// Return a Client instance using the configured Resty client.
	return &Client{
		restyClient: clnt,
		apiKey:      apiKey,
		privateKey:  privateKey,
		debugMode:   debugMode,
	}
}

// MakeRequest makes an API request with the specified method, path, and body.
// It handles different HTTP methods (GET, POST, PUT, DELETE) and potential errors.
func (x *Client) MakeRequest(ctx context.Context, method string, path string, body any, apiSuccess any) (*resty.Response, error) {
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
	// Create a request object with debug mode enabled.
	rClient := x.restyClient.
		SetHeader("Authorization", fmt.Sprintf("Bearer %v", token)).
		SetDebug(x.debugMode).R()

	// Variables to store errors and responses.
	var (
		apiError     APIError
		httpResponse *resty.Response
	)

	// Perform the appropriate request based on the HTTP method.
	switch strings.ToUpper(method) {
	case "DELETE":
		httpResponse, err = rClient.SetContext(ctx).ForceContentType("application/json; charset=utf-8").
			SetBody(body).SetResult(apiSuccess).SetError(&apiError).Delete(path)
	case "GET":
		httpResponse, err = rClient.SetContext(ctx).SetBody(body).SetResult(apiSuccess).SetError(&apiError).Get(path)
	case "POST":
		httpResponse, err = rClient.SetContext(ctx).SetBody(body).SetResult(apiSuccess).SetError(&apiError).Post(path)
	case "PUT":
		httpResponse, err = rClient.SetContext(ctx).SetBody(body).SetResult(apiSuccess).SetError(&apiError).Put(path)
	default:
		httpResponse, err = nil, errors.New("undefined method")
	}

	return httpResponse, HandleError(err, apiError)
}
