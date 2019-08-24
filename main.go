package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/ethereum/go-ethereum/common/compiler"
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
	testGoSource = `package main
	import "fmt"
	func main() {
		fmt.Println("I was compiled dinamically")
	}
	`
)

func main() {
	fmt.Println("O Rio de Janeiro continua lindo...")
	contracts, err := compiler.CompileSolidityString("", testSource)
	if err != nil {
		fmt.Printf("error compiling source. result %v: %v", contracts, err)
		return
	}
	if len(contracts) != 1 {
		fmt.Printf("one contract expected, got %d", len(contracts))
		return
	}
	c, ok := contracts["test"]
	if !ok {
		c, ok = contracts["<stdin>:test"]
		if !ok {
			fmt.Println("info for contract 'test' not present in result")
			return
		}
	}
	if c.Code == "" {
		fmt.Println("empty code")
		return
	}
	if c.Info.Source != testSource {
		fmt.Println("wrong source")
		return
	}
	if c.Info.CompilerVersion == "" {
		fmt.Println("empty version")
		return
	}
	fmt.Printf("\nCompiled Contract %#v\n\n", contracts["<stdin>:test"])
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
