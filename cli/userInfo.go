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
var userArray Users

// Users struct which contains
// an array of users
type Users struct {
	Users []User
}

// User struct which contains a name and metadata for this basic example
type User struct {
	Name     string                 `json:"name"`
	Metadata map[string]interface{} `json:"metadata"`
}

// Used to write metadata to vault
func updateUserInfo(jsonName string) {
	jsonFile, err := os.Open(jsonName)
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
	fmt.Printf("Content of %s added to Vault's identity/entity secret engine sucessfully!\n", jsonName)
	jsonFile.Close()
}

// Used to read metadata from Vault
func readUserInfo(name string, fromCli bool) {
	if !fromCli {
		inVault := false
		for _, v := range userArray.Users {
			if v.Name == name {
				inVault = true
				jsonStr, err := json.Marshal(v.Metadata)
				if err != nil {
					fmt.Printf("Error: %s", err.Error())
				} else {
					fmt.Println(string(jsonStr))
				}
			}
		}
		if !inVault {
			fmt.Println("User not in Vault's identity/secret engine")
		}
	} else {
		// this simply prints out sample user
		secretData := map[string]interface{}{
			"name": "user",
			"metadata": map[string]interface{}{
				"dataset123": 4,
			},
		}
		jsonStr, err := json.Marshal(secretData)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		} else {
			fmt.Println(string(jsonStr))
		}
	}
}

// Used to list users in Vault
func listUserInfo(fromCli bool) {
	if !fromCli {
		jsonStr, err := json.Marshal(userArray.Users)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		} else {
			fmt.Println(string(jsonStr))
		}
	} else {
		// this simply prints out sample user
		secretData := map[string]interface{}{
			"name": "user",
			"metadata": map[string]interface{}{
				"dataset123": 4,
			},
		}
		jsonStr, err := json.Marshal(secretData)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		} else {
			fmt.Println(string(jsonStr))
		}
	}
}

// Used to mimic Vault functionality
// This is primarily useful for the mock, as it stores, reads and lists information
// not entirely necessary for cli
func readInput() {
	fmt.Print("# ")
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
		fmt.Print("# ")
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
