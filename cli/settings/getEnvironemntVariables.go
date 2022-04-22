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
