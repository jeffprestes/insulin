package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	testSource = `
pragma solidity >0.0.0;
contract testaaa {
   /// @notice Will multiply ` + "`a`" + ` by 7.
   function multiply(uint a) public returns(uint d) {
       return a * 7;
   }
}
`
)

func main() {
	fmt.Println("Compiling contracts...")
	contracts, err := compiler.CompileSolidityString("", testSource)
	if err != nil {
		fmt.Printf("error compiling source. result %v: %v", contracts, err)
	}
	c, ok := contracts["testaaa"]
	if !ok {
		c, ok = contracts["<stdin>:testaaa"]
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
	err = ioutil.WriteFile("./artifacts/testaaa.json", file, 0644)
	if err != nil {
		fmt.Printf("Failed to write file to disk: %v", err)
	}

	var (
		abis  []string
		bins  []string
		types []string
		sigs  []map[string]string
		libs  = make(map[string]string)
	)

	// Gather all non-excluded contract for binding
	for name, contract := range contracts {
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

	code, err := bind.Bind(types, abis, bins, sigs, "artifacts", bind.LangGo, libs)
	if err != nil {
		fmt.Println("Failed to generate ABI binding: %v", err)
	}

	os.MkdirAll("tmp/proxycontracts", 0700)
	if err := ioutil.WriteFile("./tmp/proxycontracts/testaaa.go", []byte(code), 0600); err != nil {
		fmt.Println("Failed to write ABI binding: %v", err)
	}
}
