package main

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
			Srcmap        string `json:"srcmap"`
			SrcmapRuntime string `json:"srcmap-runtime"`
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
