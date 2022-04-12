package auth

import (
	"cli/cli/settings"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// read contents of a file
func readFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("reading tmp/secrets file errored. %w", err)
	}
	filedata := string(data)
	return filedata, nil
}

// returns bearer token
func GetBearerToken() (string, error) {
	keycloakClientSecret, _ := readFile(settings.KEYCLOAK_CLIENT_SECRET)
	testUsername, _ := readFile(settings.KEYCLOAK_TEST_USERNAME)
	testPassword, _ := readFile(settings.KEYCLOAK_TEST_PASSWORD)

	// set payload values
	payload := url.Values{}
	payload.Set("client_id", "local_candig")
	payload.Set("client_secret", keycloakClientSecret)
	payload.Set("grant_type", "password")
	payload.Set("username", testUsername)
	payload.Set("password", testPassword)
	payload.Set("scope", "openid")

	// curl keycloak
	req, err := http.NewRequest("POST", settings.KEYCLOAK_ACCESS_TOKEN_URL, strings.NewReader(payload.Encode()))
	if err != nil {
		return "", fmt.Errorf("post request to keycloak errored. %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("reading test user's password errored. %w", err)
	}

	jsonDataFromHttp, err := json.Marshal(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error. %w", err)
	}
	token := string(jsonDataFromHttp)
	defer resp.Body.Close()

	return token, nil
}
