package io

import (
	v "cli/validators"
	"fmt"
)

func Write(string user) {
	err := v.ValidateWrite(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = h.WriteUserInfo(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Secret written successfully.")
}

// TODO add Read(), List(), Delete() below as Write() above

/*
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
