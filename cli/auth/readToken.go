package auth

import (
	"cli/cli/settings"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

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
