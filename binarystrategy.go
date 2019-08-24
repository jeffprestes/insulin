package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

func createBinaryTestFile(fileInfo os.FileInfo) (err error) {
	tmp := strings.Split(fileInfo.Name(), ".")
	fileName := tmp[0]
	templateFile, err := template.ParseFiles("./testsbinarybaseline.tmpl")
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

func runTestBinaryFile(fileInfo os.FileInfo) (err error) {
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
