package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/urfave/cli.v1"
)

const (
	testSource = `
pragma solidity >0.0.0;
contract test {
   /// @notice Will multiply ` + "`a`" + ` by 7.
   function multiply(uint a) public returns(uint d) {
       return a * 7;
   }
}
`
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
			Action:  comp,
		},
	}
	cli.CommandHelpTemplate = commandHelperTemplate
}

func comp(c *cli.Context) error {
	SolidityPath := c.Args().Get(0)
	data, err := ioutil.ReadFile(SolidityPath)
	if err != nil {
		fmt.Println("File reading error", err)
	}
	fmt.Printf(string(data))
	return nil
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
