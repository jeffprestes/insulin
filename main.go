package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

const (
	commandHelperTemplate = `{{.Name}}{{if .Subcommands}} command{{end}}{{if .Flags}} [command options]{{end}} [arguments...]
{{if .Description}}{{.Description}}
{{end}}{{if .Subcommands}}
SUBCOMMANDS:
	{{range .Subcommands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
	{{end}}{{end}}{{if .Flags}}
OPTIONS:
{{range $.Flags}}{{"\t"}}{{.}}
{{end}}
{{end}}`
)

var (
	app *cli.App
)

func init() {
	app = cli.NewApp()
	app.Name = "Insulin"
	app.Commands = []cli.Command{
		{
			Name:    "compile",
			Aliases: []string{"c"},
			Usage:   "compile Solidity contract",
			Action:  callComp,
		},
	}
	cli.CommandHelpTemplate = commandHelperTemplate
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
