package main

import (
	"log"

	cs "cli/configSettings"

	"github.com/hashicorp/vault/api"
)

// Connects to Vault server
func Client(token string) (*api.Client, error) {
	config := cs.DEFAULT_CONFIG

	config.Address = cs.VAULT_ADDRESS

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}
	client.SetToken(token)
	return client, nil
}
