// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SmartcontractMetaData contains all meta data concerning the Smartcontract contract.
var SmartcontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_wallet\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"gettasknr\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_tasknr\",\"type\":\"int256\"}],\"name\":\"settasknr\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161018b38038061018b83398101604081905261002f91610060565b600080546001600160a01b039092166001600160a01b03199283161781556001556002805490911633179055610090565b60006020828403121561007257600080fd5b81516001600160a01b038116811461008957600080fd5b9392505050565b60ed8061009e6000396000f3fe60806040526004361060265760003560e01c8063c50b299014602b578063de55bdcd14603c575b600080fd5b603a6036366004609f565b605d565b005b348015604757600080fd5b5060015460405190815260200160405180910390f35b6001819055600080546040516001600160a01b03909116913480156108fc02929091818181858888f19350505050158015609b573d6000803e3d6000fd5b5050565b60006020828403121560b057600080fd5b503591905056fea264697066735822122029cc843a8d59b7d241d1bbfa97584f36270495506c1a250f3703773878477b3a64736f6c634300080d0033",
}

// SmartcontractABI is the input ABI used to generate the binding from.
// Deprecated: Use SmartcontractMetaData.ABI instead.
var SmartcontractABI = SmartcontractMetaData.ABI

// SmartcontractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SmartcontractMetaData.Bin instead.
var SmartcontractBin = SmartcontractMetaData.Bin

// DeploySmartcontract deploys a new Ethereum contract, binding an instance of Smartcontract to it.
func DeploySmartcontract(auth *bind.TransactOpts, backend bind.ContractBackend, _wallet common.Address) (common.Address, *types.Transaction, *Smartcontract, error) {
	parsed, err := SmartcontractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SmartcontractBin), backend, _wallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Smartcontract{SmartcontractCaller: SmartcontractCaller{contract: contract}, SmartcontractTransactor: SmartcontractTransactor{contract: contract}, SmartcontractFilterer: SmartcontractFilterer{contract: contract}}, nil
}

// Smartcontract is an auto generated Go binding around an Ethereum contract.
type Smartcontract struct {
	SmartcontractCaller     // Read-only binding to the contract
	SmartcontractTransactor // Write-only binding to the contract
	SmartcontractFilterer   // Log filterer for contract events
}

// SmartcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmartcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartcontractSession struct {
	Contract     *Smartcontract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmartcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartcontractCallerSession struct {
	Contract *SmartcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SmartcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartcontractTransactorSession struct {
	Contract     *SmartcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SmartcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartcontractRaw struct {
	Contract *Smartcontract // Generic contract binding to access the raw methods on
}

// SmartcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartcontractCallerRaw struct {
	Contract *SmartcontractCaller // Generic read-only contract binding to access the raw methods on
}

// SmartcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartcontractTransactorRaw struct {
	Contract *SmartcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartcontract creates a new instance of Smartcontract, bound to a specific deployed contract.
func NewSmartcontract(address common.Address, backend bind.ContractBackend) (*Smartcontract, error) {
	contract, err := bindSmartcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Smartcontract{SmartcontractCaller: SmartcontractCaller{contract: contract}, SmartcontractTransactor: SmartcontractTransactor{contract: contract}, SmartcontractFilterer: SmartcontractFilterer{contract: contract}}, nil
}

// NewSmartcontractCaller creates a new read-only instance of Smartcontract, bound to a specific deployed contract.
func NewSmartcontractCaller(address common.Address, caller bind.ContractCaller) (*SmartcontractCaller, error) {
	contract, err := bindSmartcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmartcontractCaller{contract: contract}, nil
}

// NewSmartcontractTransactor creates a new write-only instance of Smartcontract, bound to a specific deployed contract.
func NewSmartcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartcontractTransactor, error) {
	contract, err := bindSmartcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmartcontractTransactor{contract: contract}, nil
}

// NewSmartcontractFilterer creates a new log filterer instance of Smartcontract, bound to a specific deployed contract.
func NewSmartcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*SmartcontractFilterer, error) {
	contract, err := bindSmartcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmartcontractFilterer{contract: contract}, nil
}

// bindSmartcontract binds a generic wrapper to an already deployed contract.
func bindSmartcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SmartcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smartcontract *SmartcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smartcontract.Contract.SmartcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smartcontract *SmartcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartcontract.Contract.SmartcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smartcontract *SmartcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smartcontract.Contract.SmartcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smartcontract *SmartcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smartcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smartcontract *SmartcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smartcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smartcontract *SmartcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smartcontract.Contract.contract.Transact(opts, method, params...)
}

// Gettasknr is a free data retrieval call binding the contract method 0xde55bdcd.
//
// Solidity: function gettasknr() view returns(int256)
func (_Smartcontract *SmartcontractCaller) Gettasknr(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Smartcontract.contract.Call(opts, &out, "gettasknr")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gettasknr is a free data retrieval call binding the contract method 0xde55bdcd.
//
// Solidity: function gettasknr() view returns(int256)
func (_Smartcontract *SmartcontractSession) Gettasknr() (*big.Int, error) {
	return _Smartcontract.Contract.Gettasknr(&_Smartcontract.CallOpts)
}

// Gettasknr is a free data retrieval call binding the contract method 0xde55bdcd.
//
// Solidity: function gettasknr() view returns(int256)
func (_Smartcontract *SmartcontractCallerSession) Gettasknr() (*big.Int, error) {
	return _Smartcontract.Contract.Gettasknr(&_Smartcontract.CallOpts)
}

// Settasknr is a paid mutator transaction binding the contract method 0xc50b2990.
//
// Solidity: function settasknr(int256 _tasknr) payable returns()
func (_Smartcontract *SmartcontractTransactor) Settasknr(opts *bind.TransactOpts, _tasknr *big.Int) (*types.Transaction, error) {
	return _Smartcontract.contract.Transact(opts, "settasknr", _tasknr)
}

// Settasknr is a paid mutator transaction binding the contract method 0xc50b2990.
//
// Solidity: function settasknr(int256 _tasknr) payable returns()
func (_Smartcontract *SmartcontractSession) Settasknr(_tasknr *big.Int) (*types.Transaction, error) {
	return _Smartcontract.Contract.Settasknr(&_Smartcontract.TransactOpts, _tasknr)
}

// Settasknr is a paid mutator transaction binding the contract method 0xc50b2990.
//
// Solidity: function settasknr(int256 _tasknr) payable returns()
func (_Smartcontract *SmartcontractTransactorSession) Settasknr(_tasknr *big.Int) (*types.Transaction, error) {
	return _Smartcontract.Contract.Settasknr(&_Smartcontract.TransactOpts, _tasknr)
}
