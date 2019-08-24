// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package artifacts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TestaaaABI is the input ABI used to generate the binding from.
const TestaaaABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"multiply\",\"outputs\":[{\"name\":\"d\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TestaaaFuncSigs maps the 4-byte function signature to its string representation.
var TestaaaFuncSigs = map[string]string{
	"c6888fa1": "multiply(uint256)",
}

// TestaaaBin is the compiled bytecode used for deploying new contracts.
var TestaaaBin = "0x6080604052348015600f57600080fd5b50609c8061001e6000396000f300608060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663c6888fa181146043575b600080fd5b348015604e57600080fd5b506058600435606a565b60408051918252519081900360200190f35b600702905600a165627a7a72305820daf188452fd1ba786d97ae48518ec19fd9b90d800e2d65d2200caa40aef0beff0029"

// DeployTestaaa deploys a new Ethereum contract, binding an instance of Testaaa to it.
func DeployTestaaa(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Testaaa, error) {
	parsed, err := abi.JSON(strings.NewReader(TestaaaABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestaaaBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Testaaa{TestaaaCaller: TestaaaCaller{contract: contract}, TestaaaTransactor: TestaaaTransactor{contract: contract}, TestaaaFilterer: TestaaaFilterer{contract: contract}}, nil
}

// Testaaa is an auto generated Go binding around an Ethereum contract.
type Testaaa struct {
	TestaaaCaller     // Read-only binding to the contract
	TestaaaTransactor // Write-only binding to the contract
	TestaaaFilterer   // Log filterer for contract events
}

// TestaaaCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestaaaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestaaaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestaaaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestaaaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestaaaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestaaaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestaaaSession struct {
	Contract     *Testaaa          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestaaaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestaaaCallerSession struct {
	Contract *TestaaaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TestaaaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestaaaTransactorSession struct {
	Contract     *TestaaaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TestaaaRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestaaaRaw struct {
	Contract *Testaaa // Generic contract binding to access the raw methods on
}

// TestaaaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestaaaCallerRaw struct {
	Contract *TestaaaCaller // Generic read-only contract binding to access the raw methods on
}

// TestaaaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestaaaTransactorRaw struct {
	Contract *TestaaaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestaaa creates a new instance of Testaaa, bound to a specific deployed contract.
func NewTestaaa(address common.Address, backend bind.ContractBackend) (*Testaaa, error) {
	contract, err := bindTestaaa(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Testaaa{TestaaaCaller: TestaaaCaller{contract: contract}, TestaaaTransactor: TestaaaTransactor{contract: contract}, TestaaaFilterer: TestaaaFilterer{contract: contract}}, nil
}

// NewTestaaaCaller creates a new read-only instance of Testaaa, bound to a specific deployed contract.
func NewTestaaaCaller(address common.Address, caller bind.ContractCaller) (*TestaaaCaller, error) {
	contract, err := bindTestaaa(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestaaaCaller{contract: contract}, nil
}

// NewTestaaaTransactor creates a new write-only instance of Testaaa, bound to a specific deployed contract.
func NewTestaaaTransactor(address common.Address, transactor bind.ContractTransactor) (*TestaaaTransactor, error) {
	contract, err := bindTestaaa(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestaaaTransactor{contract: contract}, nil
}

// NewTestaaaFilterer creates a new log filterer instance of Testaaa, bound to a specific deployed contract.
func NewTestaaaFilterer(address common.Address, filterer bind.ContractFilterer) (*TestaaaFilterer, error) {
	contract, err := bindTestaaa(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestaaaFilterer{contract: contract}, nil
}

// bindTestaaa binds a generic wrapper to an already deployed contract.
func bindTestaaa(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestaaaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testaaa *TestaaaRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Testaaa.Contract.TestaaaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testaaa *TestaaaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testaaa.Contract.TestaaaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testaaa *TestaaaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Testaaa.Contract.TestaaaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Testaaa *TestaaaCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Testaaa.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Testaaa *TestaaaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Testaaa.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Testaaa *TestaaaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Testaaa.Contract.contract.Transact(opts, method, params...)
}

// Multiply is a paid mutator transaction binding the contract method 0xc6888fa1.
//
// Solidity: function multiply(uint256 a) returns(uint256 d)
func (_Testaaa *TestaaaTransactor) Multiply(opts *bind.TransactOpts, a *big.Int) (*types.Transaction, error) {
	return _Testaaa.contract.Transact(opts, "multiply", a)
}

// Multiply is a paid mutator transaction binding the contract method 0xc6888fa1.
//
// Solidity: function multiply(uint256 a) returns(uint256 d)
func (_Testaaa *TestaaaSession) Multiply(a *big.Int) (*types.Transaction, error) {
	return _Testaaa.Contract.Multiply(&_Testaaa.TransactOpts, a)
}

// Multiply is a paid mutator transaction binding the contract method 0xc6888fa1.
//
// Solidity: function multiply(uint256 a) returns(uint256 d)
func (_Testaaa *TestaaaTransactorSession) Multiply(a *big.Int) (*types.Transaction, error) {
	return _Testaaa.Contract.Multiply(&_Testaaa.TransactOpts, a)
}
