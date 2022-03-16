package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "Vault Helper Tool CLI",
		Usage: "Lets user read, update and list user information.",
		Commands: []*cli.Command{
			{
				Name:    "write",
				Aliases: []string{"w"},
				Usage:   "update user information by overwriting (provide 1 argument - name of json file)",
				Action: func(c *cli.Context) error {
					arg0 := c.Args().Get(0)
					arg1 := c.Args().Get(1)
					if arg0 != "" && arg1 != "" {
						updateUserInfo(arg0, arg1)
					} else {
						fmt.Println("Missing json file name")
					}
					return nil
				},
			},
			{
				Name:    "read",
				Aliases: []string{"r"},
				Usage:   "read information for user (provide 1 argument - name of user to update)",
				Action: func(c *cli.Context) error {
					arg0 := c.Args().Get(0)
					arg1 := c.Args().Get(1)
					if arg0 != "" && arg1 != "" {
						readUserInfo(arg0, arg1)
					} else {
						fmt.Println("Missing name of user")
					}
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list all users and their data (no arguments needed)",
				Action: func(c *cli.Context) error {
					arg0 := c.Args().Get(0)
					if arg0 != "" {
						listUserInfo(arg0)
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
}
