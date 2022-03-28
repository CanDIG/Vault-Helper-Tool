package main

import (
	"bufio"
	h "cli/handlers"
	p "cli/printers"
	v "cli/validators"
	"fmt"
	"log"
	"os"
	"strings"
)

// TODO Rewrite this function to act ONLY as an interface; no other logic.
// This function allows for increased functionality of the Vault helper tool
// Reads one line of user input at a time (DONE, but has errCommand to print out)
func interactiveApp() {
	inputPrompt := "# Enter command or enter q to quit: "
	fmt.Print(inputPrompt)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result := scanner.Text()
		newRes := strings.Split(result, " ")
		command := newRes[0]

		// TODO len(newRes) >= validates that command + arg is given; refactor
		// s.t. easier to read, correct assignment of responsibilities
		// b/w interface, validator, and handler (NOT DONE)
		// TODO validation of the token is a generic step; shouldn't be in
		// command-specific validators (DONE)

		// TODO handle all commands like "w" block (refactor error handling) (DONE, confirm print)
		if (command == "write" || command == "w") && len(newRes) >= 2 {
			err := v.ValidateWrite(newRes[1])
			if err != nil {
				fmt.Println(err)
			}
			errWrite := h.WriteUserInfo(newRes[1])
			if errWrite != nil {
				fmt.Println(errWrite)
			} else {
				p.PrintOuputWrite()
			}
		} else if (command == "read" || command == "r") && len(newRes) >= 2 {
			err := v.ValidateRead(newRes[1])
			if err != nil {
				fmt.Println(err)
			}
			errRead := h.ReadUserInfo(newRes[1])
			if errRead != nil {
				fmt.Println(errRead)
			} else {
				p.PrintOutputRead(newRes[1])
			}
		} else if (command == "list" || command == "l") && len(newRes) >= 1 {
			/*	err := v.ValidateList()
				if err != nil {
					fmt.Println(err)
				} */
			errList := h.ListUserInfo()
			if errList != nil {
				fmt.Println(errList)
			} else {
				p.PrintOutputList()
			}
		} else if (command == "delete" || command == "d") && len(newRes) >= 2 {
			err := v.ValidateDelete(newRes[1])
			if err != nil {
				fmt.Println(err)
			}
			errDelete := h.DeleteUserInfo(newRes[1])
			if errDelete != nil {
				fmt.Println(errDelete)
			} else {
				p.PrintOuputDelete()
			}
		} else if command == "exit" || command == "q" {
			break
		} else {
			fmt.Println("Wrong command executed or wrong number of arguments. Use ./cli -help for more info")
		}
		fmt.Print(inputPrompt)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
