package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"text/template"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/crypto"

	"gopkg.in/urfave/cli.v1"
)

func compile(testSource string, fileInfo os.FileInfo) {
	fmt.Println("Compiling contracts...")
	contracts, err := compiler.CompileSolidityString("", testSource)
	if err != nil {
		fmt.Printf("error compiling source. result %v: %v", contracts, err)
		return
	}

	var (
		abis  []string
		bins  []string
		types []string
		sigs  []map[string]string
		libs  = make(map[string]string)
	)

	for name, contract := range contracts {
		c, ok := contracts[name]
		if !ok {
			c, ok = contracts["<stdin>:"+name]
			if !ok {
				fmt.Println("\nCompilation failed :(")
				return
			}
		}
		fmt.Printf("\nCompiled Contract\n")

		// writing to d
		file, err := json.MarshalIndent(c, "", " ")
		if err != nil {
			fmt.Printf("Failed to parse ABIs from compiler output: %v", err)
		}

		os.Mkdir("artifacts", 0700)
		err = ioutil.WriteFile("./artifacts/"+strings.Split(fileInfo.Name(), ".")[0]+".json", file, 0644)
		if err != nil {
			fmt.Printf("Failed to write file to disk: %v", err)
		}

		abi, err := json.Marshal(contract.Info.AbiDefinition) // Flatten the compiler parse
		if err != nil {
			fmt.Printf("Failed to parse ABIs from compiler output: %v", err)
		}
		abis = append(abis, string(abi))
		bins = append(bins, contract.Code)
		sigs = append(sigs, contract.Hashes)
		nameParts := strings.Split(name, ":")
		types = append(types, nameParts[len(nameParts)-1])

		libPattern := crypto.Keccak256Hash([]byte(name)).String()[2:36]
		libs[libPattern] = nameParts[len(nameParts)-1]
	}

	//fmt.Println("abis: ", abis, "bins: ", bins, "sigs: ", sigs, "types: ", types, "libs: ", libs)

	code, err := bind.Bind(types, abis, bins, sigs, "proxycontracts", bind.LangGo, libs)
	if err != nil {
		fmt.Println("Failed to generate ABI binding:", err)
		return
	}

	if err := os.MkdirAll("tmp/proxycontracts", 0700); err != nil {
		fmt.Println("Failed to create directory:", err)
		return
	}

	if err := ioutil.WriteFile("./tmp/proxycontracts/"+strings.Split(fileInfo.Name(), ".")[0]+".go", []byte(code), 0600); err != nil {
		fmt.Println("Failed to write ABI binding", err)
	}
}

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

func callComp(c *cli.Context) (err error) {
	solidityPath := c.Args().Get(0)
	files, err := ioutil.ReadDir(solidityPath)
	if err != nil {
		fmt.Println("callComp error: ", err)
		return
	}

	for _, file := range files {

		r, err := regexp.MatchString(".sol", file.Name())
		if r {
			data, err := ioutil.ReadFile(file.Name())
			if err != nil {
				fmt.Println("Error reading extension", err)
				return err
			}
			compile(string(data), file)
			createTestFile(file)
			runTestFile(file)
		}
		if err != nil {
			fmt.Println("File reading error", err)
			return err
		}

	}
	return
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func createTestFile(fileInfo os.FileInfo) (err error) {
	tmp := strings.Split(fileInfo.Name(), ".")
	fileName := tmp[0]
	templateFile, err := template.ParseFiles("./testsbaseline.tmpl")
	if err != nil {
		fmt.Printf("Failed to open template file: %v", err)
		return
	}
	file, err := os.Create("./tmp/" + fileName + ".go")
	if err != nil {
		fmt.Printf("Failed to create test file: %v", err)
		return
	}
	err = templateFile.Execute(file, "")
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	file.Close()
	var out, stderr bytes.Buffer
	cmd := exec.Command("go", "build", "-o", "./tmp/"+fileName, "./tmp/"+fileName+".go")
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to compile Go code: %v\n%v", err, stderr.String())
		return
	}
	fmt.Println("Binary generated: ", out.String())
	return
}

func runTestFile(fileInfo os.FileInfo) (err error) {
	tmp := strings.Split(fileInfo.Name(), ".")
	fileName := tmp[0]

	var out, stderr bytes.Buffer

	cmd := exec.Command("./tmp/" + fileName)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to execute Go code dinamically generated: %v\n%v", err, stderr.String())
		return
	}
	fmt.Println("Binary executed: ", out.String())
	return
}
