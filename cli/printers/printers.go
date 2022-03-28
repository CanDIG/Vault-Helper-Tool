package handlers

import (
	cs "cli/configSettings"
	"encoding/json"
	"fmt"
)

func PrintOuputWrite() {
	fmt.Println("Secret written successfully.")
}

func PrintOutputRead(name string) {
	endpoint := "identity/entity/name/" + name
	secret, _ := cs.VaultClient.Logical().Read(endpoint)
	data, _ := secret.Data["metadata"].(map[string]interface{})
	jsonStr, _ := json.Marshal(data)
	fmt.Println(string(jsonStr))
}

func PrintOutputList() {
	listSecret, _ := cs.VaultClient.Logical().List("identity/entity/name")
	datamap := listSecret.Data
	data := datamap["keys"].([]interface{})
	for _, n := range data {
		nStr := fmt.Sprint(n)
		fmt.Println(n)
		PrintOutputRead(nStr)
		fmt.Println("-------------------------") // just for legibility purposes
	}
}

func PrintOuputDelete() {
	fmt.Println("User deleted successfully.")
}
