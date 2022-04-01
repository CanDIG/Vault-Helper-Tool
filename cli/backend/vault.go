package backend

import (
	"cli/cli/settings"
	"fmt"

	"github.com/hashicorp/vault/api"
)

// Connects to Vault server
func Client(token string) (*api.Client, error) {
	config := api.DefaultConfig()

	config.Address = settings.VAULT_ADDRESS

	client, err := api.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("new Vault client creation errored: %w", err)
	}

	client.SetToken(token)
	return client, nil
}
