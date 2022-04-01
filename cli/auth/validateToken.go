package auth

import "errors"

// Does error handling
func ValidateToken(token string) error {
	if token == "" || token == "\n" {
		return errors.New("token is empty")
	}
	return nil
}
