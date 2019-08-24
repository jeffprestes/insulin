package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common/compiler"
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
	fmt.Printf("\nCompiled Contract %#v\n", c)
	file, _ := json.MarshalIndent(c, "", " ")
	_ = ioutil.WriteFile("testaaa.json", file, 0644)
}
