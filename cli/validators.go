package main

import (
	"fmt"
)

// Basic error handling for number of arguments (update call)
func validateWrite(arg0 string, arg1 string) bool {
	if arg0 != "" && arg1 != "" && len(arg1) > 5 && arg1[len(arg1)-5:] == ".json" {
		return true
	} else if arg0 == "" && arg1 == "" {
		fmt.Println("No arguments provided, missing token and json file name")
		return false
	} else if arg1 == "" {
		fmt.Println("Only one argument provided")
		return false
	} else if len(arg1) <= 5 || arg1[len(arg1)-5:] != ".json" {
		fmt.Println("File provided is not a json file")
		return false
	} else {
		return false
	}
}

// Basic error handling for number of arguments (read call)
func validateRead(arg0 string, arg1 string) bool {
	if arg0 != "" && arg1 != "" {
		return true
	} else if arg0 == "" && arg1 == "" {
		fmt.Println("No arguments provided, missing token and user's name")
		return false
	} else {
		fmt.Println("Only one argument provided")
		return false
	}
}

// Basic error handling for number of arguments (list call)
func validateList(arg0 string) bool {
	if arg0 != "" {
		return true
	} else {
		fmt.Println("No arguments provided, missing token")
		return false
	}
}
