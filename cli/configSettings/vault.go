package configSettings

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
	errToken := ValidateToken()
	if errToken != nil {
		return nil, errToken
	}
	client.SetToken(GetToken())
	return client, nil
}

// TODO verify if intialization here is good style
var VaultClient, Err = Client()
