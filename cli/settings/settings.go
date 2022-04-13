package settings

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvironmentVariable(variableName string, defaultName string) string {
	envVar, ok := os.LookupEnv(variableName)
	if !ok {
		err := godotenv.Load(path + "/.env")
		if err != nil {
			return defaultName
		}
		return os.Getenv(variableName)
	}
	return envVar
}

var path, _ = os.Getwd()
var DEFAULT_VAULT_ADDRESS = "http://127.0.0.1:8200"
var DEFAULT_TOKEN_PATH = "/Vault-Helper-Tool/token.txt"

var VAULT_ADDRESS = GetEnvironmentVariable("VAULT_SERVICE_PUBLIC_URL", DEFAULT_VAULT_ADDRESS)

var TOKEN_PATH = path + GetEnvironmentVariable("TOKEN_PATH", DEFAULT_TOKEN_PATH)

//var PROGRESS_FILE = goDotEnvVariable("PROGRESS_FILE")
