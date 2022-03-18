package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// This function allows for increased functionality of the Vault helper tool
// Reads one line of user input at a time
func interactiveApp() {
	inputPrompt := "# Enter command or enter q to quit: "
	fmt.Print(inputPrompt)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result := scanner.Text()
		newRes := strings.Split(result, " ")
		command := newRes[0]
		if (command == "write" || command == "w") && len(newRes) >= 2 {
			rightInput := validateWrite(TOKEN, newRes[1])
			if rightInput {
				writeUserInfo(TOKEN, newRes[1])
			}
		} else if (command == "read" || command == "r") && len(newRes) >= 2 {
			rightInput := validateRead(TOKEN, newRes[1])
			if rightInput {
				readUserInfo(TOKEN, newRes[1])
			}
		} else if (command == "list" || command == "l") && len(newRes) >= 1 {
			rightInput := validateList(TOKEN)
			if rightInput {
				listUserInfo(TOKEN)
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
