package auth

import (
	"cli/cli/settings"
	"fmt"
)

func ValidateBearerToken() error {
	/*	path, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("error getting working directory. %w", err)
		} */
	_, err := readFile(settings.KEYCLOAK_CLIENT_SECRET)
	if err != nil {
		return fmt.Errorf("reading keycloak client secret errored. %w", err)
	}
	_, err = readFile(settings.KEYCLOAK_TEST_USERNAME)
	if err != nil {
		return fmt.Errorf("reading test username errored. %w", err)
	}
	_, err = readFile(settings.KEYCLOAK_TEST_PASSWORD)
	if err != nil {
		return fmt.Errorf("reading test user's password errored. %w", err)
	}
	return nil
}
