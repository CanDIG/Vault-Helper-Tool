package configSettings

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// TODO have this happen very early, immediately before instantiation of Vault client
// All authentication with the Vault server, including the prerequisite token-reading from file,
// should happen before any commands are handled
// in the case of interactive mode, that should be before any input is read (DONE)

// returns obtained token
func GetToken() string {
	absPath, _ := filepath.Abs("../cli/configSettings/token.txt")
	token, _ := ioutil.ReadFile(absPath) // just pass the file name
	tokenStr := string(token)
	finalToken := strings.TrimSuffix(tokenStr, "\n")

	return finalToken

}
