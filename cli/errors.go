package main

import (
	"fmt"
)

// Basic error handling for number of arguments (update call)
func callUpdate(arg0 string, arg1 string) {
	if arg0 != "" && arg1 != "" && len(arg1) > 5 && arg1[len(arg1)-5:] == ".json" {
		updateUserInfo(arg0, arg1)
	} else if arg0 == "" && arg1 == "" {
		fmt.Println("No arguments provided, missing token and json file name")
	} else if arg0 == "" || arg1 == "" {
		fmt.Println("Only one argument provided")
	} else if len(arg1) <= 5 || arg1[len(arg1)-5:] != ".json" {
		fmt.Println("File provided is not a json file")
	}
}

// Basic error handling for number of arguments (read call)
func callRead(arg0 string, arg1 string) {
	if arg0 != "" && arg1 != "" {
		readUserInfo(arg0, arg1)
	} else if arg0 == "" && arg1 == "" {
		fmt.Println("No arguments provided, missing token and user's name")
	} else {
		fmt.Println("Only one argument provided")
	}
}

// Basic error handling for number of arguments (list call)
func callList(arg0 string) {
	if arg0 != "" {
		listUserInfo(arg0)
	} else {
		fmt.Println("No arguments provided, missing token")
	}
}
