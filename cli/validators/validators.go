package validators

import (
	"errors"
)

// TODO Make all validators return error (or nil if succesful)

// Basic error handling for number of arguments (update call)
// TODO arg1 == "" will never happen if the length of the input is checked in InteractiveApp (SOLVED: for cli)
// TODO check how token validation differs between interactive and detached modes
func ValidateWrite(arg1 string) error {
	// if arg1 != "" && len(arg1) > 5 && arg1[len(arg1)-5:] == ".json" {
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
// no error checking necessary
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
