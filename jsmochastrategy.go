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

func createJsMochaTestFile(fileInfo os.FileInfo) (err error) {
	tmp := strings.Split(fileInfo.Name(), ".")
	fileName := tmp[0]
	templateFile, err := template.ParseFiles("./testsjsmochabaseline.tmpl")
	if err != nil {
		fmt.Printf("Failed to open template file: %v", err)
		return
	}
	err = os.MkdirAll("tmp/tests", 0700)
	if err != nil {
		fmt.Println("Failed to create directory to put tests files:", err)
		return
	}
	file, err := os.Create("./tmp/tests/" + fileName + "_test.go")
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
	return
}

func runTestJsMochaFile(fileInfo os.FileInfo) (err error) {
	err = nil
	var fileName = "./tmp/tests/" + fileInfo.Name()
	var out, stderr bytes.Buffer

	cmd := exec.Command("go", "test", fileName)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Tests for file %s failed\n%s\n", fileName, out.String())
		return
	}
	fmt.Printf("Tests of file %s were executed:%s\n", fileName, out.String())
	return
}
