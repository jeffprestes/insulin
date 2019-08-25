package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

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
		fmt.Println("Error ", err)
		return
	}

	err = ioutil.WriteFile("./artifacts/"+destFileName, out, 0644)
	if err != nil {
		fmt.Printf("compileForMythX - Failed to write file to disk: %v", err)
		return err
	}

	contractCompiled := SolidityContract{}
	err = json.NewDecoder(strings.NewReader(string(out))).Decode(&contractCompiled)
	if err != nil {
		fmt.Printf("Error decoding SolC results to a struct: %s\n", err.Error())
		return
	}

	fmt.Println("Analysis file generated")

	authCredentials := MythXCredentialsRequest{}
	authCredentials.Password = "Insulin4!"
	authCredentials.Username = "0x230e277b1a6b36d56da0f143fe73abda7a926dbb"
	authCredentials.EthAddress = authCredentials.Username
	authCredentials.UserID = "fabiohildebrand@gmail.com"
	authCredentials.JwtLifetimes.Access = "10 mins"
	authCredentials.JwtLifetimes.Refresh = "3 days"
	authCredentials.Permissions = []string{"ANALYSIS_ALLOWANCE_MINIMAL"}

	requestBody, err := json.Marshal(authCredentials)
	if err != nil {
		fmt.Println("Error marsheling data", err.Error())
		return
	}

	resp, err := http.Post("https://api.mythx.io/v1/auth/login", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error calling MythX authentication data", err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		text, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Printf("Error decoding Mythx authentication response: %s\n", err2.Error())
			return
		}
		fmt.Println("Error returned from MythX authentication data", resp.Status, string(text))
		return
	}
	credentials := MythXCredentialsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&credentials)
	if err != nil {
		fmt.Printf("Error decoding Mythx authentication response: %s\n", err.Error())
		return
	}
	fmt.Println("Success! MythX Credentials Response", credentials)

	contract := MythXSCAnalysisRequest{}
	contract.ClientToolName = "Insulin"
	contract.Data.Bytecode = contractCompiled.Contracts.ContractsRegistroSolRegistro.Bin
	contract.Data.MainSource = contractCompiled.SourceList[0]
	contract.Data.Sources.Contract.Source = `pragma solidity ^0.5.0;

	contract Registro {
			string public mensagem;
	
			constructor () public {
					mensagem = "Uma boa e pacifica morte para todos...";
			}
	
			function RegistrarMensagem(string memory _mensagem) public {
					mensagem = _mensagem;
					emit NovaMensagem(mensagem);
			}
	
			event NovaMensagem(
					string _novaMsg
			);
	}`
	ast, err := json.Marshal(contractCompiled.Sources.ContractsRegistroSol.AST)
	if err != nil {
		fmt.Println("Error marsheling AST compiled to string", err)
		return
	}
	contract.Data.Sources.Contract.Ast = string(ast)

	requestBody, err = json.Marshal(contract)
	if err != nil {
		fmt.Println("Error marsheling contract data", err.Error())
		return
	}

	var httpClient = http.Client{
		Timeout: time.Duration(45 * time.Second),
	}
	req, err := http.NewRequest("POST", "https://api.mythx.io/v1/analyses", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error calling MythX SC analysis data", err.Error())
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+credentials.JwtTokens.Access)
	//req.PostForm = form
	resp, err = httpClient.Do(req)
	if err != nil {
		log.Printf("Error requesting MythX SC Analysis: [%+v]\nErro: %s", req, err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		text, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Printf("Error decoding Mythx SC analysis  response: %s\n", err2.Error())
			return
		}
		fmt.Println("Error returned from MythX SC analysis data", resp.Status, string(text))
		return
	}
	analysis := MythXSCAnalysisResponse{}
	err = json.NewDecoder(resp.Body).Decode(&analysis)
	if err != nil {
		fmt.Printf("Error decoding Mythx SC analysis response: %s\n", err.Error())
		return
	}
	fmt.Println("Success! MythX Response", analysis)
	return
}
