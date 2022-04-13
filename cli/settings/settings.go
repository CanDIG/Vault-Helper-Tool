package settings

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvironmentVariable(variableName string, defaultName string) string {
	envVar, ok := os.LookupEnv(variableName)
	if !ok {
		err := godotenv.Load(path + ".env")
		if err != nil {
			return defaultName
		}
		fmt.Println(os.Getenv(envVar))
		return os.Getenv(envVar)
	}
	return envVar
}

var path, _ = os.Getwd()
var DEFAULT_VAULT_ADDRESS = "http://127.0.0.1:8200"
var DEFAULT_TOKEN_PATH = path + "/Vault-Helper-Tool/token.txt"

var VAULT_ADDRESS = GetEnvironmentVariable("VAULT_SERVICE_PUBLIC_URL", DEFAULT_VAULT_ADDRESS)

var TOKEN_PATH = GetEnvironmentVariable("TOKEN_PATH", DEFAULT_TOKEN_PATH)

//var PROGRESS_FILE = goDotEnvVariable("PROGRESS_FILE")
