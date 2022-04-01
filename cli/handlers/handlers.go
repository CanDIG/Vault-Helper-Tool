package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/vault/api"
)

// Used to write metadata to vault
func HandleWrite(jsonName string, tx *api.Client) error {
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

	_, err = tx.Logical().Write("identity/entity", value)
	if err != nil {
		return fmt.Errorf("unable to write secret: %w", err)
	}

	jsonFile.Close()
	return nil
}

func HandleRead(name string, tx *api.Client) (*api.Secret, error) {
	endpoint := "identity/entity/name/" + name
	secret, err := tx.Logical().Read(endpoint)
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
	} else {
		err := name + " does not exist in Vault."
		return nil, fmt.Errorf(err)
	}
	return secret, nil
}

func HandleList(tx *api.Client) (*api.Secret, error) {
	listSecret, err := tx.Logical().List("identity/entity/name")
	if err != nil {
		return nil, fmt.Errorf("unable to list secret: %v", err)
	}
	if listSecret == nil {
		return nil, fmt.Errorf("no users in vault")
	}
	return listSecret, nil
}

func HandleDelete(name string, tx *api.Client) error {
	endpoint := "identity/entity/name/" + name
	secret, err := tx.Logical().Delete(endpoint)
	if err != nil {
		return fmt.Errorf("unable to delete secret: %v", err)
	}
	if secret != nil {
		err := name + " does not exist in Vault."
		return fmt.Errorf(err)
	}
	return nil
}
