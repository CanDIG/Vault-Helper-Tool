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
		Usage: "Lets user read, update and list user information. If entering interactive mode, enter q to quit",
		Commands: []*cli.Command{
			{
				Name:    "write",
				Aliases: []string{"w"},s
				Usage:   "update user information by overwriting (provide 2 arguments - token, json file)",
				Action: func(c *cli.Context) error {
					cArg0 := c.Args().Get(0)
					cArg1 := c.Args().Get(1)
					callUpdate(cArg0, cArg1)
					return nil
				},
			},
			{
				Name:    "read",
				Aliases: []string{"r"},
				Usage:   "read metadata for user (provide 2 arguments - token, name of user)",
				Action: func(c *cli.Context) error {
					cArg0 := c.Args().Get(0)
					cArg1 := c.Args().Get(1)
					callRead(cArg0, cArg1)
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list all users and their data (no arguments needed)",
				Action: func(c *cli.Context) error {
					cArg0 := c.Args().Get(0)
					callList(cArg0)
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
	if len(os.Args) == 1 {
		readInput()
	}
}
