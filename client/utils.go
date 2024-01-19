package client

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// createAndSignJWTToken generates and signs a JWT token with specified claims and private key.
//
// Args:
//   - privateKey: *rsa.PrivateKey - The private key used for signing the token.
//   - apiKey: string - The API key to be included in the token's claims.
//   - path: string - The URI path to be included in the token's claims.
//   - bodyJSON: string - The JSON body to be hashed and included in the token's claims.
//
// Returns:
//   - string: The signed JWT token.
//   - error: An error if any occurred during token creation or signing.
func createAndSignJWTToken(privateKey *rsa.PrivateKey, apiKey, path, bodyJSON string) (string, error) {
	// timestamp :=
	nonce := make([]byte, 8)
	if _, err := rand.Read(nonce); err != nil {
		return "", fmt.Errorf("error generating nonce: %w", err)
	}

	token := jwt.MapClaims{
		"uri":      path,                                    // The requested URI path.
		"nonce":    hex.EncodeToString(nonce),               // A unique nonce for security.
		"iat":      time.Now().Unix(),                       // Issued at timestamp.
		"exp":      time.Now().Add(time.Second * 30).Unix(), // Expiration time (30 seconds from now).
		"sub":      apiKey,                                  // The API key associated with the request.
		"bodyHash": hashBody(bodyJSON),                      // Hash of the request body for integrity.
	}

	// Sign the token using RS256 algorithm and the provided private key.
	return jwt.NewWithClaims(jwt.SigningMethodRS256, token).SignedString(privateKey)
}

// hashBody generates a SHA-256 hash of the given string data
// and returns the hash as a hexadecimal string.
func hashBody(bodyJSON string) string {
	h := sha256.New()
	h.Write([]byte(bodyJSON))
	return hex.EncodeToString(h.Sum(nil))
}

// ParsePrivateKey parses a PEM-encoded RSA private key and returns an *rsa.PrivateKey object.
func ParsePrivateKey(privateKeyPEM string) (*rsa.PrivateKey, error) {
	return jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyPEM))
}

// jsonify marshals the given value v into a JSON string.
// It handles errors gracefully and returns an empty string if v is nil.
func jsonify(v any) (string, error) {
	if v == nil {
		return "", nil
	}

	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
