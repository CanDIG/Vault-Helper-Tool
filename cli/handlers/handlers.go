package handlers

import (
	cs "cli/configSettings"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/vault/api"
)

// TODO rewrite all error to resemble (DONE)
/* 	jsonFile, err := os.Open(jsonName)
if err != nil {
	return fmt.Errorf("Could not open file. %w", err)
}
*/

// TODO make sure that none of the handlers print output directly;
// The only code doing printing should be in the "interface" source code,
// ie. main.go and interactiveApp.go (DONE: Added another print function for each)
// (DONE: Confirm about printers)

// Used to write metadata to vault
func WriteUserInfo(jsonName string) error {
	jsonFile, err := os.Open(jsonName)
	if err != nil {
		return fmt.Errorf("could not open file. %w", err)
	}

	byteValue, parseErr := ioutil.ReadAll(jsonFile)
	if parseErr != nil {
		return fmt.Errorf("error parsing data: %w", parseErr)
	}

	var value map[string]interface{}
	marshErr := json.Unmarshal([]byte(byteValue), &value)
	if marshErr != nil {
		return fmt.Errorf("error using unmarshal: %w", marshErr)
	}

	_, err = cs.VaultClient.Logical().Write("identity/entity", value)
	if err != nil {
		return fmt.Errorf("unable to write secret: %w", err)
	}

	jsonFile.Close()
	return nil
}

// Used to read metadata from Vault
func ReadUserInfo(name string) (*api.Secret, error) {
	endpoint := "identity/entity/name/" + name
	secret, err := cs.VaultClient.Logical().Read(endpoint)
	if err != nil {
		return nil, fmt.Errorf("unable to read secret: %w", err)
	}
	if secret != nil { // if doesn't exist
		data, ok := secret.Data["metadata"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("data type assertion failed: %T %#v", secret.Data["metadata"], secret.Data["metadata"])
		}
		_, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("error: %s", err.Error())
		}
		//	fmt.Println(string(jsonStr))
	} else {
		err := name + " does not exist in Vault."
		return nil, fmt.Errorf(err)
	}
	return secret, nil
}

// Used to list users + metadata in Vault
func ListUserInfo() (*api.Secret, error) {
	listSecret, err := cs.VaultClient.Logical().List("identity/entity/name")
	if err != nil {
		return nil, fmt.Errorf("unable to list secret: %v", err)
	}
	if listSecret == nil {
		return nil, fmt.Errorf("no users in vault")
	}
	return listSecret, nil
}

// Used to read metadata from Vault
func DeleteUserInfo(name string) error {
	endpoint := "identity/entity/name/" + name
	secret, err := cs.VaultClient.Logical().Delete(endpoint)
	if err != nil {
		return fmt.Errorf("unable to delete secret: %v", err)
	}
	if secret != nil {
		err := name + " does not exist in Vault."
		return fmt.Errorf(err)
	}
	return nil
}
