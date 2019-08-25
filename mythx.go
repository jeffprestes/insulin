package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

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
	fmt.Println("Compiling for MythX contracts from file: ", fileInfo.Name())
	destFileName := strings.Split(fileInfo.Name(), ".")[0] + ".json"

	//solc --pretty-json --combined-json ast,bin,bin-runtime,srcmap,srcmap-runtime ./contracts/registro.sol
	out, err := exec.Command("solc", "--pretty-json", "--combined-json", "ast,bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes", "./contracts/"+fileInfo.Name()).Output()
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("./artifacts/"+destFileName, out, 0644)
	if err != nil {
		fmt.Printf("compileForMythX - Failed to write file to disk: %v", err)
		return err
	}

	fmt.Println("Analysis file generated")
	return
}
