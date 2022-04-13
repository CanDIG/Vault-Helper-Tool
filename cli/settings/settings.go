package settings

import (
	"os"
)

var path, _ = os.Getwd()

// Hardcoded variables
var DEFAULT_VAULT_ADDRESS = "http://127.0.0.1:8200"
var DEFAULT_TOKEN_PATH = "/Vault-Helper-Tool/token.txt"

// Environment variables
var VAULT_ADDRESS = GetEnvironmentVariable("VAULT_SERVICE_PUBLIC_URL", DEFAULT_VAULT_ADDRESS)
var TOKEN_PATH = path + GetEnvironmentVariable("TOKEN_PATH", DEFAULT_TOKEN_PATH)

//var PROGRESS_FILE = goDotEnvVariable("PROGRESS_FILE")
