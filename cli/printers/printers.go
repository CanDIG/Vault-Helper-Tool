package handlers

import (
	h "cli/handlers"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/vault/api"
)

func PrintOuputWrite() {
	fmt.Println("Secret written successfully.")
}

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
