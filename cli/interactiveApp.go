package main

import (
	"bufio"
	"cli/cli/middleware"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/vault/api"
)

// TODO Rewrite this function to act ONLY as an interface; no other logic.
// This function allows for increased functionality of the Vault helper tool
// Reads one line of user input at a time (DONE, but has errCommand to print out)
func interactiveApp(tx *api.Client) {
	inputPrompt := "# Enter command or enter q to quit: "
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(inputPrompt)
	for scanner.Scan() {
		// Reset the response and error values
		var response string
		var err error

		// Prompt the user for next command
		input := strings.Split(scanner.Text(), " ")
		command := input[0]
		args := input[1:]

		// TODO handle all commands like "w" block

		// Parse and fulfill command
		// TODO Technical Debt to consider during future refactors:
		// 		number-of-arguments validator should be shared between
		// 		interactiveApp (interactive mode) and main (single-command mode);
		//		refactor should modify the len(args)==n conditions below.
		if (command == "write" || command == "w") && len(args) == 1 {
			response, err = middleware.Write(args[0], tx)
		} else if (command == "read" || command == "r") && len(args) == 1 {
			response, err = middleware.Read(args[0], tx)
		} else if (command == "list" || command == "l") && len(args) == 0 {
			response, err = middleware.List(tx)
		} else if (command == "delete" || command == "d") && len(args) == 1 {
			response, err = middleware.Delete(args[0], tx)
		} else if command == "exit" || command == "q" {
			break
		} else {
			fmt.Println("Wrong command executed or wrong number of arguments. Use ./cli -help for more info")
		}

		// Respond to user
		if err != nil {
			fmt.Println(fmt.Errorf("middleware errored: %w", err))
			continue
		}
		fmt.Println(response)
		fmt.Print(inputPrompt)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
