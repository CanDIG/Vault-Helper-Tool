package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "Vault Helper Tool CLI",
		Usage: "Lets user read, update and list user information. Modify the secretFile.txt with the token necessary. If entering interactive mode, enter q or to quit",
		Commands: []*cli.Command{
			{
				Name:    "write",
				Aliases: []string{"w"},
				Usage:   "update user information by overwriting (provide 1 argument - json file)",
				Action: func(c *cli.Context) error {
					cArg0 := c.Args().Get(0)
					rightInput := validateWrite(TOKEN, cArg0)
					if rightInput {
						writeUserInfo(TOKEN, cArg0)
					}
					return nil
				},
			},
			{
				Name:    "read",
				Aliases: []string{"r"},
				Usage:   "read metadata for user (provide 1 argument - name of user)",
				Action: func(c *cli.Context) error {
					cArg0 := c.Args().Get(0)
					rightInput := validateRead(TOKEN, cArg0)
					if rightInput {
						readUserInfo(TOKEN, cArg0)
					}
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list all users and their data (no arguments needed)",
				Action: func(c *cli.Context) error {
					rightInput := validateList(TOKEN)
					if rightInput {
						listUserInfo(TOKEN)
					}
					return nil
				},
			},
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	// used to run interactive mode, call ./cli to run this
	if len(os.Args) == 1 {
		interactiveApp()
	}
}
