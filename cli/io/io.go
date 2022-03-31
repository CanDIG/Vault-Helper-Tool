package io

import (
	h "cli/cli/handlers"
	v "cli/cli/validators"
	"encoding/json"
	"fmt"
)

func Write(jsonFile string) {
	err := v.ValidateWrite(jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = h.WriteUserInfo(jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Secret written successfully.")
}

// TODO add Read(), List(), Delete() below as Write() above

func Read(user string) {
	err := v.ValidateRead(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	Secret, err := h.ReadUserInfo(user)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		data, _ := Secret.Data["metadata"].(map[string]interface{})
		jsonStr, _ := json.Marshal(data)
		fmt.Println(string(jsonStr))
	}
}

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
