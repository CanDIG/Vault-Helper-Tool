package auth

import (
	"cli/cli/settings"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// TODO have this happen very early, immediately before instantiation of Vault client
// All authentication with the Vault server, including the prerequisite token-reading from file,
// should happen before any commands are handled
// in the case of interactive mode, that should be before any input is read (DONE)

// returns obtained token
func ReadToken() (string, error) {
	absPath, _ := filepath.Abs(settings.TOKEN_PATH)
	token, err := ioutil.ReadFile(absPath) // just pass the file name
	if err != nil {
		return "", fmt.Errorf("reading token file errored. %w", err)
	}

	fmtToken := strings.TrimSuffix(string(token), "\n")

	return fmtToken, nil
}
