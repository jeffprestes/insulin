package main

import "time"

//MythXSCAnalysisResponse response from MythX analysis
type MythXSCAnalysisResponse struct {
	APIVersion     string    `json:"apiVersion"`
	HarveyVersion  string    `json:"harveyVersion"`
	MaruVersion    string    `json:"maruVersion"`
	MythrilVersion string    `json:"mythrilVersion"`
	QueueTime      int       `json:"queueTime"`
	Status         string    `json:"status"`
	SubmittedAt    time.Time `json:"submittedAt"`
	SubmittedBy    string    `json:"submittedBy"`
	UUID           string    `json:"uuid"`
}

//MythXSCAnalysisRequest request to MythX SC Analysis
type MythXSCAnalysisRequest struct {
	ClientToolName string `json:"clientToolName"`
	NoCacheLookup  bool   `json:"noCacheLookup"`
	Data           struct {
		Bytecode   string `json:"bytecode"`
		MainSource string `json:"mainSource"`
		Sources    struct {
			Contract struct {
				Source string `json:"source"`
				Ast    string `json:"ast"`
			} `json:"Registro.sol"`
		} `json:"sources"`
	} `json:"data"`
}

//MythXCredentialsResponse return from Mythx
type MythXCredentialsResponse struct {
	JwtTokens struct {
		Access  string `json:"access"`
		Refresh string `json:"refresh"`
	} `json:"jwtTokens"`
	Permissions struct {
		Owned     []string `json:"owned"`
		Requested []string `json:"requested"`
		Granted   []string `json:"granted"`
	} `json:"permissions"`
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

//MythXCredentialsRequest submit to MythX
type MythXCredentialsRequest struct {
	Password     string `json:"password"`
	Username     string `json:"username"`
	JwtLifetimes struct {
		Access  string `json:"access"`
		Refresh string `json:"refresh"`
	} `json:"jwtLifetimes"`
	Permissions []string `json:"permissions"`
	EthAddress  string   `json:"ethAddress"`
	UserID      string   `json:"userId"`
}

//SolidityContract This reflects a structure of solc compilation parameters ast,bin,bin-runtime,srcmap,srcmap-runtime,abi,userdoc,devdoc,metadata,hashes
type SolidityContract struct {
	Contracts struct {
		ContractsRegistroSolRegistro struct {
			Abi        string `json:"abi"`
			Bin        string `json:"bin"`
			BinRuntime string `json:"bin-runtime"`
			Devdoc     string `json:"devdoc"`
			Hashes     struct {
				RegistrarMensagemString string `json:"RegistrarMensagem(string)"`
				Mensagem                string `json:"mensagem()"`
			} `json:"hashes"`
			Metadata      string `json:"metadata"`
			Srcmap        string `json:"sourceMap"`
			SrcmapRuntime string `json:"deployedSourceMap"`
			Userdoc       string `json:"userdoc"`
		} `json:"./contracts/registro.sol:Registro"`
	} `json:"contracts"`
	SourceList []string `json:"sourceList"`
	Sources    struct {
		ContractsRegistroSol struct {
			AST struct {
				Attributes struct {
					AbsolutePath    string `json:"absolutePath"`
					ExportedSymbols struct {
						Registro []int `json:"Registro"`
					} `json:"exportedSymbols"`
				} `json:"attributes"`
				Children []struct {
					Attributes struct {
						Literals []string `json:"literals"`
					} `json:"attributes"`
					ID       int    `json:"id"`
					Name     string `json:"name"`
					Src      string `json:"src"`
					Children []struct {
						Attributes struct {
							Constant        bool        `json:"constant"`
							Name            string      `json:"name"`
							Scope           int         `json:"scope"`
							StateVariable   bool        `json:"stateVariable"`
							StorageLocation string      `json:"storageLocation"`
							Type            string      `json:"type"`
							Value           interface{} `json:"value"`
							Visibility      string      `json:"visibility"`
						} `json:"attributes"`
						Children []struct {
							Attributes struct {
								Name string `json:"name"`
								Type string `json:"type"`
							} `json:"attributes"`
							ID   int    `json:"id"`
							Name string `json:"name"`
							Src  string `json:"src"`
						} `json:"children"`
						ID   int    `json:"id"`
						Name string `json:"name"`
						Src  string `json:"src"`
					} `json:"children,omitempty"`
				} `json:"children"`
				ID   int    `json:"id"`
				Name string `json:"name"`
				Src  string `json:"src"`
			} `json:"AST"`
		} `json:"./contracts/registro.sol"`
	} `json:"sources"`
	Version string `json:"version"`
}
