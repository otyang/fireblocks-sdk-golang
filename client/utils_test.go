package client

import (
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestHashBody(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"}, // Empty string
		{"string without special chars", "Hello, world!", "315f5bdb76d078c43b8ac0064e4a0164612b1fce77c869345bfc94c75894edd3"},
		{"string with special chars", "Test with special characters: !@#$%^&*()_+", "1f99021668a80f1f240c04fd7573154c9a24a8d37adcfa0754853820cfdaf12a"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := hashBody(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestParsePrivateKey(t *testing.T) {
	t.Parallel()

	prvtKey, err := ParsePrivateKey(ConfigPrivateKey)
	assert.NoError(t, err)
	assert.NotEmpty(t, prvtKey)
}

func TestJsonify(t *testing.T) {
	t.Parallel()

	// Test cases
	tests := []struct {
		name     string
		input    any
		expected string
		wantErr  bool
	}{
		{name: "Nil input", input: nil, expected: "", wantErr: false},
		{name: "Basic string", input: "hello", expected: "\"hello\"", wantErr: false},
		{name: "Struct", input: struct{ Name string }{Name: "Alice"}, expected: "{\"Name\":\"Alice\"}", wantErr: false},
		{name: "Map", input: map[string]int{"age": 30}, expected: "{\"age\":30}", wantErr: false},
		{name: "Marshal error", input: func() {}, expected: "", wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := jsonify(tc.input)
			assert.Equal(t, tc.wantErr, (err != nil))
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestCreateAndSignJWTToken(t *testing.T) {
	t.Parallel()

	testData := []struct {
		name     string
		uri      string
		bodyJson string
	}{
		{
			name:     "With Body",
			uri:      "v1/hello/hi",
			bodyJson: `{"body":"hello"}`,
		},
		{
			name:     "Without Body",
			uri:      "v1/hello/hi",
			bodyJson: "",
		},
		{
			name:     "With Body and Query Params",
			uri:      "v1/hello/hi?name=John",
			bodyJson: `{"body":"hello hey"}`,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			prvtKey, err := ParsePrivateKey(ConfigPrivateKey)
			assert.NoError(t, err)

			signedToken, err := createAndSignJWTToken(prvtKey, ConfigApiKey, tt.uri, tt.bodyJson)
			assert.NoError(t, err)

			// validate the signed token
			parsedToken, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
				return prvtKey.Public(), nil
			})
			assert.NoError(t, err)
			assert.Truef(t, parsedToken.Valid, "The signed token is not valid.")

			// validate token has the correct claims
			claims, ok := parsedToken.Claims.(jwt.MapClaims)
			assert.Truef(t, ok, "The signed token does not have the correct claims.")

			// validate claims
			assert.Equal(t, hashBody(tt.bodyJson), claims["bodyHash"])
			assert.Equal(t, tt.uri, claims["uri"])
			assert.Equal(t, ConfigApiKey, claims["sub"])

			t.Log("Signed token:", signedToken)
		})
	}
}
