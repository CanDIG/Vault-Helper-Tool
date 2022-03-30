package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"cli/cli/io"
	cs "cli/configSettings"
	h "cli/handlers"
	p "cli/printers"
	v "cli/validators"

	"github.com/urfave/cli/v2"
)

func main() {
	if cs.Err != nil {
		fmt.Println(cs.Err)
	}
	app := &cli.App{
		Name:  "Vault Helper Tool CLI",
		Usage: "Lets user read, update and list user information. Modify the secretFile.txt with the token necessary. If entering interactive mode, enter q or to quit",
		Commands: []*cli.Command{
			{
				// TODO rewrite read, list, and delete to use io.Read() etc.
				Name:    "write",
				Aliases: []string{"w"},
				Usage:   "update user information by overwriting (provide 1 argument - json file)",
				Action: func(c *cli.Context) error {
					user := c.Args().Get(0)
					io.Write(user)
					return nil
				},
			},
			{
				Name:    "read",
				Aliases: []string{"r"},
				Usage:   "read metadata for user (provide 1 argument - name of user)",
				Action: func(c *cli.Context) error {
					cArg0 := c.Args().Get(0)
					err := v.ValidateRead(cArg0)
					if err != nil {
						fmt.Println(err)
					}
					Secret, errRead := h.ReadUserInfo(cArg0)
					if errRead != nil {
						fmt.Println(errRead)
					} else {
						p.PrintOutputRead(Secret)
					}
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
					secretList, errList := h.ListUserInfo()
					if errList != nil {
						fmt.Println(errList)
					} else {
						p.PrintOutputList(secretList)
					}
					return nil
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "delete user from vault (provide 1 argument - name of user)",
				Action: func(c *cli.Context) error {
					cArg0 := c.Args().Get(0)
					err := v.ValidateDelete(cArg0)
					if err != nil {
						fmt.Println(err)
					}
					errDelete := h.DeleteUserInfo(cArg0)
					if errDelete != nil {
						fmt.Println(errDelete)
					} else {
						p.PrintOuputDelete()
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
