package configSettings

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func GetToken() string {
	absPath, _ := filepath.Abs("../cli/configSettings/token.txt")
	token, err := ioutil.ReadFile(absPath) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	tokenStr := string(token)

	return tokenStr
}
