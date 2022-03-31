package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"cli/cli/io"

	"github.com/urfave/cli/v2"
)

func main() {
	if Err != nil {
		fmt.Println(Err)
	}
	app := &cli.App{
		Name:  "Vault Helper Tool CLI",
		Usage: "Lets user read, update and list user information. Modify the secretFile.txt with the token necessary. If entering interactive mode, enter q or to quit",
		Commands: []*cli.Command{
			{
				// TODO rewrite read, list, and delete to use io.Read() etc. (DONE)
				Name:    "write",
				Aliases: []string{"w"},
				Usage:   "update user information by overwriting (provide 1 argument - json file)",
				Action: func(c *cli.Context) error {
					jsonFile := c.Args().Get(0)
					io.Write(jsonFile)
					return nil
				},
			},
			{
				Name:    "read",
				Aliases: []string{"r"},
				Usage:   "read metadata for user (provide 1 argument - name of user)",
				Action: func(c *cli.Context) error {
					user := c.Args().Get(0)
					io.Read(user)
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list all users and their data (no arguments needed)",
				Action: func(c *cli.Context) error {
					/*	err := v.ValidateList(cs.TOKEN)
						if err != nil {
							fmt.Println(err)
						} */
					io.List()
					return nil
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "delete user from vault (provide 1 argument - name of user)",
				Action: func(c *cli.Context) error {
					user := c.Args().Get(0)
					io.Delete(user)
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
