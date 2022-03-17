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
func readInput() {
	fmt.Print("# Enter command or press q to quit: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result := scanner.Text()
		newRes := strings.Split(result, " ")
		command := newRes[0]
		if (command == "write" || command == "w") && len(newRes) >= 3 {
			callUpdate(newRes[1], newRes[2])
		} else if (command == "read" || command == "r") && len(newRes) >= 3 {
			callRead(newRes[1], newRes[2])
		} else if (command == "list" || command == "l") && len(newRes) >= 2 {
			callList(newRes[1])
		} else if command == "exit" || command == "q" {
			break
		} else {
			fmt.Println("Wrong command executed or wrong number of arguments. Use ./cli -help for more info")
		}
		fmt.Print("# Enter command or enter q to quit: ")
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
