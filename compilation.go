package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/crypto"
	"gopkg.in/urfave/cli.v1"
)

func callComp(c *cli.Context) {
	solidityPath := c.Args().Get(0)
	files, err := ioutil.ReadDir(solidityPath)
	if err != nil {
		fmt.Println("callComp error: ", err)
		return
	}

	for _, file := range files {
		r, _ := regexp.MatchString(".sol", file.Name())
		if r {
			data, err := ioutil.ReadFile(file.Name())
			if err != nil {
				fmt.Println("Error reading extension", err)
				return
			}
			err = compile(string(data), file)
			if err != nil {
				return
			}
			createJsMochaTestFile(file)
			if err != nil {
				return
			}
			runTestJsMochaFile(file)
			if err != nil {
				return
			}
		}
	}
	return
}

func compile(testSource string, fileInfo os.FileInfo) (err error) {
	err = nil
	fmt.Println("Compiling contracts from file: ", fileInfo.Name())
	contracts, err := compiler.CompileSolidityString("", testSource)
	if err != nil {
		fmt.Printf("error compiling source from file: %s. Result %#v: %s\n", fileInfo.Name(), contracts, err.Error())
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
			return err
		}

		os.Mkdir("artifacts", 0700)
		err = ioutil.WriteFile("./artifacts/"+strings.Split(fileInfo.Name(), ".")[0]+".json", file, 0644)
		if err != nil {
			fmt.Printf("Failed to write file to disk: %v", err)
			return err
		}

		abi, err := json.Marshal(contract.Info.AbiDefinition) // Flatten the compiler parse
		if err != nil {
			fmt.Printf("Failed to parse ABIs from compiler output: %v", err)
			return err
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

	err = os.MkdirAll("tmp/proxycontracts", 0700)
	if err != nil {
		fmt.Println("Failed to create directory to contract in Go:", err)
		return
	}

	err = ioutil.WriteFile("./tmp/proxycontracts/"+strings.Split(fileInfo.Name(), ".")[0]+".go", []byte(code), 0600)
	if err != nil {
		fmt.Println("Failed to write ABI binding", err)
		return
	}
	return
}
