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
go get github.com/otyang/fireblocks
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
client, err := client.New(clientID, clientSecret, apiKey, apiEndpoint)
```

Usage:
```go
 
func main() {
    client := fireblocks.New(apiKey, privateKey, baseURL)
    vaultService := vault.New(client)

    // Create a vault account
    vaultAccount, err := client.Vault.CreateVaultAccount(context.Background(), vault.CreateVaultParams{
        Name: "My Vault",
    })

    feeResponse, err := client.Transaction.EstimateFeeForSendingToExternalAddress(
        ctx context.Context, assetID, amount, treatAsGrossAmount,
    )

    // ... (other examples for transactions, webhooks, etc.)
}
```



## Examples

See the `examples` directory for complete code examples.

## Documentation

- **API Reference:** [https://developers.fireblocks.com/reference](https://developers.fireblocks.com/reference) 

## Contributing

Contributions are welcome! Please follow the contributing guidelines.

## License

This SDK is licensed under the MIT License.




 

 

 


 