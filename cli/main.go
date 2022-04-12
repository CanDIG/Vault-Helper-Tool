package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"cli/cli/auth"
	"cli/cli/backend"
	"cli/cli/middleware"

	"github.com/hashicorp/vault/api"
	"github.com/urfave/cli/v2"
)

func connect() (*api.Client, error) {
	if err := auth.ValidateBearerToken(); err != nil {
		return nil, fmt.Errorf("token-setting error: %w", err)
	}
	token, err := auth.GetBearerToken()
	if err != nil {
		return nil, fmt.Errorf("token-setting error: %w", err)
	}

	tx, err := backend.Client(token)
	if err != nil {
		return nil, fmt.Errorf("connection error: %w", err)
	}

	return tx, nil
}

func resetToken(tx *api.Client) (string, error) {
	token, err := auth.ReadToken()
	if err != nil {
		return "", fmt.Errorf("token-setting error: %w", err)
	}
	if err = auth.ValidateToken(token); err != nil {
		return "", fmt.Errorf("token-setting error: %w", err)
	}
	tx.SetToken(token)

	return "Token reset successfully", nil
}

func main() {
	fmt.Println("Connecting to Vault using token in token.txt")
	tx, err := connect()
	if err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name:  "Vault Helper Tool CLI",
		Usage: "Lets user read, update and list user information. Modify the secretFile.txt with the token necessary. If entering interactive mode, enter q or to quit",
		Commands: []*cli.Command{
			{
				Name:    "write",
				Aliases: []string{"w"},
				Usage:   "update user information by overwriting (provide 1 argument - json file)",
				Action: func(c *cli.Context) error {
					jsonFile := c.Args().Get(0)

					response, err := middleware.Write(jsonFile, tx)
					if err != nil {
						return fmt.Errorf("middleware errored: %w", err)
					}

					fmt.Println(response)
					return nil
				},
			},
			{
				Name:    "read",
				Aliases: []string{"r"},
				Usage:   "read metadata for user (provide 1 argument - name of user)",
				Action: func(c *cli.Context) error {
					user := c.Args().Get(0)
					response, err := middleware.Read(user, tx)
					if err != nil {
						return fmt.Errorf("middleware errored: %w", err)
					}

					fmt.Println(response)
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list all users and their data (no arguments needed)",
				Action: func(c *cli.Context) error {
					response, err := middleware.List(tx)
					if err != nil {
						return fmt.Errorf("middleware errored: %w", err)
					}

					fmt.Println(response)
					return nil
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "delete user from vault (provide 1 argument - name of user)",
				Action: func(c *cli.Context) error {
					user := c.Args().Get(0)
					response, err := middleware.Delete(user, tx)
					if err != nil {
						return fmt.Errorf("middleware errored: %w", err)
					}

					fmt.Println(response)
					return nil
				},
			},
			{
				Name:    "set-token",
				Aliases: []string{"st"},
				Usage:   "sets user token with contents of the token.txt file",
				Action: func(c *cli.Context) error {
					response, err := resetToken(tx)
					if err != nil {
						return fmt.Errorf("reset token errored: %w", err)
					}
					fmt.Println(response)
					return nil
				},
			},
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	// used to run interactive mode, call ./cli to run this
	if len(os.Args) == 1 {
		interactiveApp(tx)
	}
}
