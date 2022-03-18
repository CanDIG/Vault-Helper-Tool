package configSettings

import (
	"log"

	"github.com/hashicorp/vault/api"
)

// Connects to Vault server
func Client(token string) (*api.Client, error) {
	config := DEFAULT_CONFIG

	config.Address = VAULT_ADDRESS

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}
	client.SetToken(token)
	return client, nil
}
