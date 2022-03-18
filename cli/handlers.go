package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Used to write metadata to vault
func writeUserInfo(token string, jsonName string) {
	errOpening := false
	client, _ := Client(token)
	jsonFile, err := os.Open(jsonName)
	if err != nil {
		errOpening = true
		log.Println("File provided does not exist: ", err)
	}

	byteValue, parseErr := ioutil.ReadAll(jsonFile)
	if parseErr != nil {
		errOpening = true
		log.Println("Error parsing data: ", parseErr)
	}

	var value map[string]interface{}
	marshErr := json.Unmarshal([]byte(byteValue), &value)
	if marshErr != nil {
		errOpening = true
		log.Println("Error using unmarshal: ", marshErr)
	}
	if !errOpening {
		_, err = client.Logical().Write("identity/entity", value)
		if err != nil {
			log.Fatalf("Unable to write secret: %v", err)
		}
		fmt.Println("Secret written successfully.")
	}
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
	if secret != nil {
		data, ok := secret.Data["metadata"].(map[string]interface{})
		if !ok {
			log.Fatalf("Data type assertion failed: %T %#v", secret.Data["metadata"], secret.Data["metadata"])
		}
		jsonStr, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
		fmt.Println(string(jsonStr))
	} else {
		errMsg := name + " does not exist in Vault."
		fmt.Println(errMsg)
	}
}

// Used to list users + metadata in Vault
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
		fmt.Println(n)
		readUserInfo(token, nStr)
		fmt.Println("-------------------------") // just for legibility purposes
	}
}
