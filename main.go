package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/mitchellh/cli"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

const Name = "gmail"
const Version = "0.0.1"

const (
	ExitCodeOK = iota
	ExitCodeError
)

const (
	errFormat = "Error:\n%s"
)

type CLI struct {
	in       io.Reader
	out, err io.Writer
	Config   *Config
}

func (c *CLI) Run(args []string) int {
	m := cli.NewCLI(Name, Version)
	m.Args = args
	m.HelpWriter = c.err
	m.Commands = map[string]cli.CommandFactory{
		"auth": func() (cli.Command, error) {
			return &AuthCommand{c: c}, nil
		},
	}

	exitStatus, err := m.Run()
	if m.IsVersion() {
		exitStatus = ExitCodeOK
	}

	if err != nil {
		fmt.Fprintf(c.err, errFormat, err)
		return ExitCodeError
	}

	return exitStatus
}

func (c *CLI) newUi() cli.Ui {
	if c == nil {
		return nil
	}

	return &cli.BasicUi{
		Reader:      c.in,
		Writer:      c.out,
		ErrorWriter: c.err,
	}
}

func (c *CLI) newSubCommnad(name string) *cli.CLI {
	m := cli.NewCLI(name, Version)
	m.HelpWriter = c.err
	return m
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Error:\n%s", err)
			os.Exit(ExitCodeError)
		}
	}()

	cli := &CLI{
		out: os.Stdout,
		err: os.Stderr,
		in:  os.Stdin,
	}

	conf, err := loadConfig()
	if err != nil {
		panic(fmt.Sprintf("Error: Config load error\n", err))
		return
	}
	cli.Config = conf

	os.Exit(cli.Run(os.Args[1:]))
}

func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	return nil
}
