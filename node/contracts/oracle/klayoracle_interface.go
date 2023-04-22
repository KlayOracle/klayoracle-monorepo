// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracle

import (
	"math/big"
	"strings"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = klaytn.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// KlayOracleInterfaceABI is the input ABI used to generate the binding from.
const KlayOracleInterfaceABI = "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"callBackContract\",\"type\":\"address\"}],\"name\":\"newOracleRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"roundAnswer\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"newRoundData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KlayOracleInterfaceBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const KlayOracleInterfaceBinRuntime = ``

// KlayOracleInterface is an auto generated Go binding around a Klaytn contract.
type KlayOracleInterface struct {
	KlayOracleInterfaceCaller     // Read-only binding to the contract
	KlayOracleInterfaceTransactor // Write-only binding to the contract
	KlayOracleInterfaceFilterer   // Log filterer for contract events
}

// KlayOracleInterfaceCaller is an auto generated read-only Go binding around a Klaytn contract.
type KlayOracleInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KlayOracleInterfaceTransactor is an auto generated write-only Go binding around a Klaytn contract.
type KlayOracleInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KlayOracleInterfaceFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type KlayOracleInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KlayOracleInterfaceSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type KlayOracleInterfaceSession struct {
	Contract     *KlayOracleInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// KlayOracleInterfaceCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type KlayOracleInterfaceCallerSession struct {
	Contract *KlayOracleInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// KlayOracleInterfaceTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type KlayOracleInterfaceTransactorSession struct {
	Contract     *KlayOracleInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// KlayOracleInterfaceRaw is an auto generated low-level Go binding around a Klaytn contract.
type KlayOracleInterfaceRaw struct {
	Contract *KlayOracleInterface // Generic contract binding to access the raw methods on
}

// KlayOracleInterfaceCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type KlayOracleInterfaceCallerRaw struct {
	Contract *KlayOracleInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// KlayOracleInterfaceTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type KlayOracleInterfaceTransactorRaw struct {
	Contract *KlayOracleInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKlayOracleInterface creates a new instance of KlayOracleInterface, bound to a specific deployed contract.
func NewKlayOracleInterface(address common.Address, backend bind.ContractBackend) (*KlayOracleInterface, error) {
	contract, err := bindKlayOracleInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KlayOracleInterface{KlayOracleInterfaceCaller: KlayOracleInterfaceCaller{contract: contract}, KlayOracleInterfaceTransactor: KlayOracleInterfaceTransactor{contract: contract}, KlayOracleInterfaceFilterer: KlayOracleInterfaceFilterer{contract: contract}}, nil
}

// NewKlayOracleInterfaceCaller creates a new read-only instance of KlayOracleInterface, bound to a specific deployed contract.
func NewKlayOracleInterfaceCaller(address common.Address, caller bind.ContractCaller) (*KlayOracleInterfaceCaller, error) {
	contract, err := bindKlayOracleInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KlayOracleInterfaceCaller{contract: contract}, nil
}

// NewKlayOracleInterfaceTransactor creates a new write-only instance of KlayOracleInterface, bound to a specific deployed contract.
func NewKlayOracleInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*KlayOracleInterfaceTransactor, error) {
	contract, err := bindKlayOracleInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KlayOracleInterfaceTransactor{contract: contract}, nil
}

// NewKlayOracleInterfaceFilterer creates a new log filterer instance of KlayOracleInterface, bound to a specific deployed contract.
func NewKlayOracleInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*KlayOracleInterfaceFilterer, error) {
	contract, err := bindKlayOracleInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KlayOracleInterfaceFilterer{contract: contract}, nil
}

// bindKlayOracleInterface binds a generic wrapper to an already deployed contract.
func bindKlayOracleInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KlayOracleInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KlayOracleInterface *KlayOracleInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KlayOracleInterface.Contract.KlayOracleInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KlayOracleInterface *KlayOracleInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KlayOracleInterface.Contract.KlayOracleInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KlayOracleInterface *KlayOracleInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KlayOracleInterface.Contract.KlayOracleInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KlayOracleInterface *KlayOracleInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KlayOracleInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KlayOracleInterface *KlayOracleInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KlayOracleInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KlayOracleInterface *KlayOracleInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KlayOracleInterface.Contract.contract.Transact(opts, method, params...)
}

// NewOracleRequest is a paid mutator transaction binding the contract method 0xec2b83f1.
//
// Solidity: function newOracleRequest(bytes4 callbackFunctionId, address callBackContract) returns(bool)
func (_KlayOracleInterface *KlayOracleInterfaceTransactor) NewOracleRequest(opts *bind.TransactOpts, callbackFunctionId [4]byte, callBackContract common.Address) (*types.Transaction, error) {
	return _KlayOracleInterface.contract.Transact(opts, "newOracleRequest", callbackFunctionId, callBackContract)
}

// NewOracleRequest is a paid mutator transaction binding the contract method 0xec2b83f1.
//
// Solidity: function newOracleRequest(bytes4 callbackFunctionId, address callBackContract) returns(bool)
func (_KlayOracleInterface *KlayOracleInterfaceSession) NewOracleRequest(callbackFunctionId [4]byte, callBackContract common.Address) (*types.Transaction, error) {
	return _KlayOracleInterface.Contract.NewOracleRequest(&_KlayOracleInterface.TransactOpts, callbackFunctionId, callBackContract)
}

// NewOracleRequest is a paid mutator transaction binding the contract method 0xec2b83f1.
//
// Solidity: function newOracleRequest(bytes4 callbackFunctionId, address callBackContract) returns(bool)
func (_KlayOracleInterface *KlayOracleInterfaceTransactorSession) NewOracleRequest(callbackFunctionId [4]byte, callBackContract common.Address) (*types.Transaction, error) {
	return _KlayOracleInterface.Contract.NewOracleRequest(&_KlayOracleInterface.TransactOpts, callbackFunctionId, callBackContract)
}

// NewRoundData is a paid mutator transaction binding the contract method 0x09c980c1.
//
// Solidity: function newRoundData(uint256 roundTime, bytes32 roundAnswer, bytes signature) returns(bool)
func (_KlayOracleInterface *KlayOracleInterfaceTransactor) NewRoundData(opts *bind.TransactOpts, roundTime *big.Int, roundAnswer [32]byte, signature []byte) (*types.Transaction, error) {
	return _KlayOracleInterface.contract.Transact(opts, "newRoundData", roundTime, roundAnswer, signature)
}

// NewRoundData is a paid mutator transaction binding the contract method 0x09c980c1.
//
// Solidity: function newRoundData(uint256 roundTime, bytes32 roundAnswer, bytes signature) returns(bool)
func (_KlayOracleInterface *KlayOracleInterfaceSession) NewRoundData(roundTime *big.Int, roundAnswer [32]byte, signature []byte) (*types.Transaction, error) {
	return _KlayOracleInterface.Contract.NewRoundData(&_KlayOracleInterface.TransactOpts, roundTime, roundAnswer, signature)
}

// NewRoundData is a paid mutator transaction binding the contract method 0x09c980c1.
//
// Solidity: function newRoundData(uint256 roundTime, bytes32 roundAnswer, bytes signature) returns(bool)
func (_KlayOracleInterface *KlayOracleInterfaceTransactorSession) NewRoundData(roundTime *big.Int, roundAnswer [32]byte, signature []byte) (*types.Transaction, error) {
	return _KlayOracleInterface.Contract.NewRoundData(&_KlayOracleInterface.TransactOpts, roundTime, roundAnswer, signature)
}
