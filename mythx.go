package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"

	cli "gopkg.in/urfave/cli.v1"
)

func mythX(c *cli.Context) {
	files, err := parseInput(c, "./contracts")
	if err != nil {
		return
	}

	for _, file := range files {
		r, _ := regexp.MatchString(".sol", file.Name())
		if r {
			data, err := ioutil.ReadFile("./contracts/" + file.Name())
			if err != nil {
				fmt.Println("Error reading extension", err)
				return
			}
			fmt.Println("Compiling contracts")
			err = compileForMythX(string(data), file)
			if err != nil {
				return
			}
		}
	}
	return
}

func compileForMythX(testSource string, fileInfo os.FileInfo) (err error) {
	err = nil
	fmt.Println("Compiling contracts from file: ", fileInfo.Name())

	//solc --pretty-json --combined-json ast,bin,bin-runtime,srcmap,srcmap-runtime ./contracts/registro.sol
	out, err := exec.Command("solc", "--pretty-json", "--combined-json", "ast,bin,bin-runtime,srcmap,srcmap-runtime", "./contracts/registro.sol").Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out) // SEND THIS TO MYTHX API

	return
}
