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
		// fmt.Println(n)
		userSecret, err := h.HandleRead(nStr, tx)
		if err != nil { // shouldn't really happen
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

// TODO reusing the printers code in the comment block at the bottom of the page,
// add RespondToRead(), RespondToList(), RespondToDelete() below as ResponToWrite() above

/* move functionality below into middleware and handlers packages

func List() {
	listSecret, errList := h.ListUserInfo()
	if errList != nil {
		fmt.Println(errList)
		return
	} else {
		datamap := listSecret.Data
		data := datamap["keys"].([]interface{})
		for _, n := range data {
			nStr := fmt.Sprint(n)
			fmt.Println(n)
			user, _ := h.ReadUserInfo(nStr)
			Read(user)
			fmt.Println("-------------------------") // just for legibility purposes
		}
	}
}

func Delete(user string) {
	err := v.ValidateDelete(user)
	if err != nil {
		fmt.Println(err)
	}
	errDelete := h.DeleteUserInfo(user)
	if errDelete != nil {
		fmt.Println(errDelete)
	} else {
		fmt.Println("User deleted successfully.")
	}
}
*/

/* move functionality below into responders package
import (
	h "cli/handlers"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/vault/api"
)

func PrintOutputRead(secret *api.Secret) {
	data, _ := secret.Data["metadata"].(map[string]interface{})
	jsonStr, _ := json.Marshal(data)
	fmt.Println(string(jsonStr))
}

func PrintOutputList(listSecret *api.Secret) {
	datamap := listSecret.Data
	data := datamap["keys"].([]interface{})
	for _, n := range data {
		nStr := fmt.Sprint(n)
		fmt.Println(n)
		user, _ := h.ReadUserInfo(nStr)
		PrintOutputRead(user)
		fmt.Println("-------------------------") // just for legibility purposes
	}
}

func PrintOuputDelete() {
	fmt.Println("User deleted successfully.")
}
*/
