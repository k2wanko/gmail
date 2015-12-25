package main

import (
	"fmt"

	"github.com/mitchellh/cli"
)

type AuthCommand struct {
	c *CLI
}

func (c *AuthCommand) Help() string {
	return "auth command"
}

func (c *AuthCommand) Run(args []string) int {
	m := c.c.newSubCommnad(Name + " auth")
	m.Args = args
	m.Commands = map[string]cli.CommandFactory{
		"login": func() (cli.Command, error) {
			return &LoginAuthCommand{c: c.c}, nil
		},
	}

	exitStatus, err := m.Run()
	if err != nil {
		fmt.Fprintf(c.c.err, errFormat, err)
	}

	return exitStatus
}

func (c *AuthCommand) Synopsis() string {
	return "Manage oauth2 credentials for Gmail Account"
}

type LoginAuthCommand struct {
	c *CLI
}

func (c *LoginAuthCommand) Help() string {
	return "auth login command"
}

func (c *LoginAuthCommand) Run(args []string) int {
	ui := c.c.newUi()

	errClose := func(typ string) int {
		ui.Error(fmt.Sprintf("Require: %s", typ))
		return ExitCodeError
	}

	email, _ := ui.Ask("Email:")
	if email == "" {
		return errClose("Email")
	}

	pass, _ := ui.AskSecret("Password:")
	if pass == "" {
		return errClose("Password")
	}

	// Add line break
	ui.Output("")

	fmt.Fprintf(c.c.out, "Email: %s Password: %s", email, pass)
	return ExitCodeOK
}

func (c *LoginAuthCommand) Synopsis() string {
	return "Login for Google Account"
}
