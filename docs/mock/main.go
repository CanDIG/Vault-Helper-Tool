package main

import (
	"log"

	vault "github.com/hashicorp/vault/api"
)

func main() {
	// set up vault server
	config := vault.DefaultConfig()

	config.Address = "http://127.0.0.1:8200"

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	client.SetToken("dev-only-token")

	// vault data to write
	secretData := map[string]interface{}{
		"name": "user",
		"metadata": map[string]interface{}{
			"dataset123": 4,
		},
	}
	// Write a secret
	_, err = client.Logical().Write("identity/entity", secretData)
	if err != nil {
		log.Fatalf("Unable to write secret: %v", err)
	}

	log.Println("Secret written successfully.")

	// Read a secret
	secret, err := client.Logical().Read("identity/entity/name/user")
	if err != nil {
		log.Fatalf("Unable to read secret: %v", err)
	}
	log.Println(secret)
	log.Println("Secret read successfully.")

	// access certain fields of the secret - verify if stored/read data is correct
	data, ok := secret.Data["metadata"].(map[string]interface{})
	if !ok {
		log.Fatalf("Data type assertion failed: %T %#v", secret.Data["metadata"], secret.Data["metadata"])
	}
	log.Println(data)
	log.Println("Secret's field accessed successfully.")
	value, ok := data["dataset123"].(string)
	if !ok {
		log.Fatalf("Value type assertion failed: %T %#v", data["dataset123"], data["dataset123"])
	}

	if value != "4" {
		log.Fatalf("Unexpected password value %q retrieved from vault", value)
	}

	// list the secrets
	listSecret, err := client.Logical().List("identity/entity/name")
	if err != nil {
		log.Fatalf("Unable to list secret: %v", err)
	}
	log.Println(listSecret)
	log.Println("Secret list accessed successfully.")
	log.Println("Access granted!")
}
