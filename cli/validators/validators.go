package validators

import (
	"errors"
)

// Basic error handling for number of arguments (write call)
func ValidateWrite(arg1 string) error {
	if arg1 == "" { // this case exclusively happens in the cli mode
		return errors.New("file name not provided")
	} else if len(arg1) <= 5 || arg1[len(arg1)-5:] != ".json" {
		return errors.New("file provided is not a json file")
	}

	return nil
}

// Basic error handling for number of arguments (read call)
func ValidateRead(arg1 string) error {
	if arg1 == "" {
		return errors.New("no arguments provided, missing user's name")
	}
	return nil
}

/*
// Basic error handling for number of arguments (list call)
// no error checking necessary for now
func ValidateList(token string) error {
	if token != "" {
		return errors.New("no arguments provided, missing token")
	}
	return nil
} */

// Basic error handling for number of arguments (delelte call)
func ValidateDelete(arg1 string) error {
	if arg1 == "" {
		return errors.New("no arguments provided, missing user's name")
	}
	return nil
}

func ValidateUpdateRole(arg1 string, arg2 string) error {
	if arg1 == "" {
		return errors.New("no arguments provided, missing filename")
	}
	if arg2 == "" {
		return errors.New("no arguments provided, missing role's name")
	}
	return nil
}
