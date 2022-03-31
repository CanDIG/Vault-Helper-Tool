package auth

import "errors"

/*
TODO remove this comment
Getting the token and validating it for use are two seperate processes.
- Reintroduce error hadling for file IO in getToken; there are many reasons that a fileread can fail even if the token itself is formatted correctly
- Remove error handling for file IO in validateToken
Seq:
1. When app is launched, ReadToken (where? in main.go?)
2. In Client, ValidateToken then immediately after SetToken
*/
// Does error handling
func ValidateToken(token string) error {
	if token == "" || token == "\n" {
		return errors.New("token is empty")
	}
	return nil
}
