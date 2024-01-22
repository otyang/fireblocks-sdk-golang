# Fireblocks SDK in Go (golang)

This SDK provides Golang bindings for interacting with the Fireblocks API, enabling you to manage vault accounts, transactions, webhooks, and more. This isnt meant to be a full sdk, but a streeamlined one with only function needed at this point


## Available Actions at this point 

- **VaultService:**
    - Create vault accounts
    - Create asset wallets and addresses
    - Find vault accounts by ID
- **TransactionService:**
    - Estimate transaction fees
    - Create transactions
    - Find transactions by ID or external transaction ID
- **WebhookService:**
    - Resend failed webhooks
    - Verify webhook transactions
- **CustomFlowSVC:**
    - Streamlined methods for address creation and sending transactions


## Installation

```bash
go get github.com/otyang/fireblocks-sdk-golang.git
```

## Usage

1. **Import the SDK:**

```go
import (
    "github.com/otyang/fireblocks/client"
    "github.com/otyang/fireblocks/transaction"
    "github.com/otyang/fireblocks/vault"
    "github.com/otyang/fireblocks/webhook"
)
```

2. **Create a Fireblocks client:**

```go 
fireblocks, err := New(client.ConfigApiKey,  client.ConfigPrivateKey, client.ConfigBaseURL, true)
```

Usage:
```go
 
func main() {
    client.ConfigApiKey = "api-key"
    client.ConfigPrivateKey = "private-key"
    client.ConfigBaseURL = "https://sandbox-api.fireblocks.io"
    client.ConfigDebugMode =  true

    fireblocks, err := New(client.ConfigApiKey,  client.ConfigPrivateKey, client.ConfigBaseURL, client.ConfigDebugMode) 
    if err != nil {
        return err
    }

    // Create a vault account
	vault, err := fireblocks.Vault.FindVaultAccountByID(context.Background(), "1")
    if err != nil {
        return err
    }
    // .....

    feeResponse, err := client.Transaction.EstimateFeeForSendingToExternalAddress(
        ctx context.Context, assetID, amount, treatAsGrossAmount,
    )

    // ... (other examples for transactions, webhooks, etc.)
}
```


## Documentation

- **API Reference:** [https://developers.fireblocks.com/reference](https://developers.fireblocks.com/reference) 

## Contributing

Contributions are welcome! Please follow the contributing guidelines.

## License

This SDK is licensed under the MIT License.




 

 

 


 