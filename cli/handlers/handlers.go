package handlers

import (
	cs "cli/configSettings"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Used to write metadata to vault
func WriteUserInfo(token string, jsonName string) {
	errOpening := false
	client, _ := cs.Client(token)
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
func ReadUserInfo(token string, name string) {
	client, _ := cs.Client(token)
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
func ListUserInfo(token string) {
	client, _ := cs.Client(token)
	listSecret, err := client.Logical().List("identity/entity/name")
	if err != nil {
		log.Fatalf("Unable to list secret: %v", err)
	}
	if listSecret != nil {
		datamap := listSecret.Data
		data := datamap["keys"].([]interface{})
		for _, n := range data {
			nStr := fmt.Sprint(n)
			fmt.Println(n)
			ReadUserInfo(token, nStr)
			fmt.Println("-------------------------") // just for legibility purposes
		}
	}
}

// Used to read metadata from Vault
func DeleteUserInfo(token string, name string) {
	client, _ := cs.Client(token)
	endpoint := "identity/entity/name/" + name
	secret, err := client.Logical().Delete(endpoint)
	if err != nil {
		log.Fatalf("Unable to delete secret: %v", err)
	}
	if secret == nil {
		fmt.Println("User sucessfully deleted from Vault.")
	} else {
		errMsg := name + " does not exist in Vault."
		fmt.Println(errMsg)
	}
}
