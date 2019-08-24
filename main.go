package main

import (
	"fmt"

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
)

func main() {
	fmt.Println("O Rio de Janeiro continua lindo...")
	contracts, err := compiler.CompileSolidityString("", testSource)
	if err != nil {
		fmt.Printf("error compiling source. result %v: %v", contracts, err)
	}
	if len(contracts) != 1 {
		fmt.Printf("one contract expected, got %d", len(contracts))
	}
	c, ok := contracts["test"]
	if !ok {
		c, ok = contracts["<stdin>:test"]
		if !ok {
			fmt.Println("info for contract 'test' not present in result")
		}
	}
	if c.Code == "" {
		fmt.Println("empty code")
	}
	if c.Info.Source != testSource {
		fmt.Println("wrong source")
	}
	if c.Info.CompilerVersion == "" {
		fmt.Println("empty version")
	}
	fmt.Printf("\nCompiled Contract %#v\n", contracts["<stdin>:test"])
}
