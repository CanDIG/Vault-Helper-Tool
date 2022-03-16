package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/vault/api"
)

// Connects to Vault server
func Client(token string) (*api.Client, error) {
	config := api.DefaultConfig()

	config.Address = "http://127.0.0.1:8200"

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}
	client.SetToken(token)
	return client, nil
}

// Used to write metadata to vault
func updateUserInfo(token string, jsonName string) {
	client, _ := Client(token)
	jsonFile, err := os.Open(jsonName)
	if err != nil {
		fmt.Println("File provided does not exist: ", err)
	}

	byteValue, parseErr := ioutil.ReadAll(jsonFile)
	if parseErr != nil {
		fmt.Println("Error parsing data: ", parseErr)
	}

	var value map[string]interface{}
	marshErr := json.Unmarshal([]byte(byteValue), &value)
	if marshErr != nil {
		fmt.Println("Error using unmarshal: ", marshErr)
	}
	_, err = client.Logical().Write("identity/entity", value)
	if err != nil {
		log.Fatalf("Unable to write secret: %v", err)
	}
	fmt.Println("Secret written successfully.")
	jsonFile.Close()
}

// Used to read metadata from Vault
func readUserInfo(token string, name string) {
	client, _ := Client(token)
	endpoint := "identity/entity/name/" + name
	secret, err := client.Logical().Read(endpoint)
	if err != nil {
		log.Fatalf("Unable to read secret: %v", err)
	}
	data, ok := secret.Data["metadata"].(map[string]interface{})
	if !ok {
		log.Fatalf("Data type assertion failed: %T %#v", secret.Data["metadata"], secret.Data["metadata"])
	}
	jsonStr, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	fmt.Println(string(jsonStr))
	//	fmt.Println("Secret read successfully.")
}

// Used to list users in Vault
func listUserInfo(token string) {
	client, _ := Client(token)
	listSecret, err := client.Logical().List("identity/entity/name")
	if err != nil {
		log.Fatalf("Unable to list secret: %v", err)
	}
	datamap := listSecret.Data
	data := datamap["keys"].([]interface{})
	for _, n := range data {
		nStr := fmt.Sprint(n)
		readUserInfo(token, nStr)
	}
	fmt.Println("Secret list accessed successfully.")
}
