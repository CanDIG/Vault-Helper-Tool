package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// NOTE: this will not be useful once backend is implemented.

// userArray will store an array of all the users.
var userArray Users

// Users struct which contains an array of users
type Users struct {
	Users []User
}

// User struct which contains a name and metadata for this basic example
type User struct {
	Name     string                 `json:"name"`
	Metadata map[string]interface{} `json:"metadata"`
}

/* Used to write metadata to vault
 The API call needed to implement write is:
	_, err = client.Logical().Write("identity/entity", secretData)
	if err != nil {
		log.Fatalf("Unable to write secret: %v", err)
	}
*/
func updateUserInfo(jsonName string) {
	jsonFile, err := os.Open(jsonName + ".json")
	if err != nil {
		fmt.Println("File provided does not exist: ", err)
	}

	byteValue, parseErr := ioutil.ReadAll(jsonFile)
	if parseErr != nil {
		fmt.Println("Error parsing data: ", parseErr)
	}

	var value User
	marshErr := json.Unmarshal([]byte(byteValue), &value)
	if marshErr != nil {
		fmt.Println("Error using unmarshal: ", marshErr)
	}
	userArray.Users = append(userArray.Users, value)
	jsonFile.Close()
}

/* Used to read metadata from Vault
The API call needed to implement read is:
	secret, err := client.Logical().Read("identity/entity/name/user")
	if err != nil {
		log.Fatalf("Unable to read secret: %v", err)
	}
*/
func readUserInfo(name string, fromCli bool) {
	if !fromCli {
		inVault := false
		for _, v := range userArray.Users {
			if v.Name == name {
				inVault = true
				fmt.Println(v.Metadata)
			}
		}
		if !inVault {
			fmt.Println("User not in Vault")
		}
	} else {
		// this simply prints out sample user
		secretData := map[string]interface{}{
			"name": "user",
			"metadata": map[string]interface{}{
				"dataset123": 4,
			},
		}
		fmt.Println(secretData["metadata"])
	}
}

/* Used to list users in Vault
The API call needed to implement list is:
	listSecret, err := client.Logical().List("identity/entity/name")
	if err != nil {
		log.Fatalf("Unable to list secret: %v", err)
	}
*/
func listUserInfo(fromCli bool) {
	if !fromCli {
		fmt.Println(userArray.Users)
	} else {
		// this simply prints out sample user
		secretData := map[string]interface{}{
			"name": "user",
			"metadata": map[string]interface{}{
				"dataset123": 4,
			},
		}
		fmt.Println(secretData)
	}
}

// Used to mimic Vault functionality
func nextCommands() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result := scanner.Text()
		newRes := strings.Split(result, " ")
		command := newRes[0]
		if command == "write" {
			updateUserInfo(newRes[1])
		} else if command == "read" {
			readUserInfo(newRes[1], false)
		} else if command == "list" {
			listUserInfo(false)
		} else if command == "exit" || command == "q" {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
