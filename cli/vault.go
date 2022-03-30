package cli

import (
	"github.com/hashicorp/vault/api"
)

// Connects to Vault server
func Client() (*api.Client, error) {
	config := DEFAULT_CONFIG

	config.Address = VAULT_ADDRESS

	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	// TODO need to declare token somewhere
	err := ValidateToken(token)
	if err != nil {
		return nil, err
	}
	client.SetToken(token)
	return client, nil
}

// TODO verify if intialization here is good style
var VaultClient, Err = Client()
