package tests

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/jeffprestes/goethereumhelper"
	"github.com/jeffprestes/insulin/tmp/proxycontracts"

	. "github.com/franela/goblin"
)

var auth *bind.TransactOpts
var backend *backends.SimulatedBackend
var contractAddress common.Address
var trx *types.Transaction
var contractInstance *proxycontracts.Registro

func Test(t *testing.T) {
	g := Goblin(t)

	g.Describe("Registro must be deployed and return the message", func() {
		// Passing Test
		g.It("Should be deployed ", func() {
			var err error
			auth, backend, _ = goethereumhelper.GetMockBlockchain()
			contractAddress, trx, contractInstance, err = proxycontracts.DeployRegistro(auth, backend)
			if err != nil {
				g.Fail(fmt.Sprintf("error when trying to compile contract: %v\n", err))
			}
			backend.Commit()
			g.Assert(err).Equal(nil)

		})
		// OK Test
		g.It("Should bring message", func() {
			var err error
			msg, err := contractInstance.Mensagem(nil)
			if err != nil {
				g.Fail(fmt.Sprintf("It was not possible to bring the message: %v\n", err))
			}
			g.Assert(msg).Equal("Uma boa e pacifica morte para todos...")
			fmt.Println("Contract Address: ", contractAddress.String(), " - Transaction: ", trx.Hash().String(), " - Message stored: ", msg)
		})
		// Failing Test
		g.It("Should not bring message", func() {
			var err error
			msg, err := contractInstance.Mensagem(nil)
			if err != nil {
				g.Fail(fmt.Sprintf("It was not possible to bring the message: %v\n", err))
			}
			g.Assert(msg).Equal("Umasss boa e pacifica morte para todos...")
			fmt.Println("Contract Address: ", contractAddress.String(), " - Transaction: ", trx.Hash().String(), " - Message stored: ", msg)
		})
	})
}
