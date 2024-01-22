package fireblocks

import (
	"context"
	"testing"

	"github.com/otyang/fireblocks/client"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// client.ConfigApiKey = "api-key"
	// client.ConfigPrivateKey = "private-key"
	// client.ConfigBaseURL = "https://sandbox-api.fireblocks.io"
	// client.ConfigDebugMode = true

	fireblocks, err := New(client.ConfigApiKey, client.ConfigPrivateKey, client.ConfigBaseURL, true)

	assert.NoError(t, err)

	vault, err := fireblocks.Vault.FindVaultAccountByID(context.Background(), "1")

	assert.NoError(t, err)
	assert.NotEmpty(t, vault)
}
