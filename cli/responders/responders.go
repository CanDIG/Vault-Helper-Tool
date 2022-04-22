package responders

import (
	h "cli/cli/handlers"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/vault/api"
)

func RespondToWrite() (string, error) {
	return "Secret written successfully.", nil
}

func RespondToRead(secret *api.Secret) (string, error) {
	data, _ := secret.Data["metadata"].(map[string]interface{})
	jsonStr, _ := json.Marshal(data)
	return string(jsonStr), nil
}

func RespondToList(listSecret *api.Secret, tx *api.Client) (string, error) {
	var userList string
	datamap := listSecret.Data
	data := datamap["keys"].([]interface{})
	for _, n := range data {
		nStr := fmt.Sprint(n)
		userSecret, err := h.HandleRead(nStr, tx)
		if err != nil {
			return "", nil
		}
		user, _ := RespondToRead(userSecret)
		userList += nStr
		userList += "\n"
		userList += user
		userList += "\n"
		userList += "-------------------------" // just for legibility purposes
		userList += "\n"
	}
	return userList, nil
}

func RespondToDelete() (string, error) {
	return ("User deleted successfully."), nil
}

func RespondToUpdateRole() (string, error) {
	return "Role updated successfully.", nil
}
