package auth

import (
	"cli/cli/settings"
	"fmt"
	"io/ioutil"
	"strings"
)

// returns obtained token
func ReadToken() (string, error) {
	token, err := ioutil.ReadFile(settings.TOKEN_PATH) // just pass the file name
	if err != nil {
		return "", fmt.Errorf("reading token file errored. %w", err)
	}

	fmtToken := strings.TrimSuffix(string(token), "\n")

	return fmtToken, nil
}
