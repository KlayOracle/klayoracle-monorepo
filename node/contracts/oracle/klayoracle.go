// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracle

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
	_ = abi.ConvertType
)

// KlayOracleMetaData contains all meta data concerning the KlayOracle contract.
var KlayOracleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"NewOracleRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hashedMessage\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"VerifyKlaytnMessage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"roundAnswer\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"roundTime\",\"type\":\"uint256\"}],\"name\":\"_getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adapterId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fulfilledCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"answer\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"roundTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"callBackContract\",\"type\":\"address\"}],\"name\":\"newOracleRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"roundAnswer\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"newRoundData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"adapterId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"callBackContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rounds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"answer\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"roundTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"splitSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// KlayOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use KlayOracleMetaData.ABI instead.
var KlayOracleABI = KlayOracleMetaData.ABI

// KlayOracle is an auto generated Go binding around an Ethereum contract.
type KlayOracle struct {
	KlayOracleCaller     // Read-only binding to the contract
	KlayOracleTransactor // Write-only binding to the contract
	KlayOracleFilterer   // Log filterer for contract events
}

// KlayOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type KlayOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KlayOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KlayOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KlayOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KlayOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KlayOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KlayOracleSession struct {
	Contract     *KlayOracle       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KlayOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KlayOracleCallerSession struct {
	Contract *KlayOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// KlayOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KlayOracleTransactorSession struct {
	Contract     *KlayOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// KlayOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type KlayOracleRaw struct {
	Contract *KlayOracle // Generic contract binding to access the raw methods on
}

// KlayOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KlayOracleCallerRaw struct {
	Contract *KlayOracleCaller // Generic read-only contract binding to access the raw methods on
}

// KlayOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KlayOracleTransactorRaw struct {
	Contract *KlayOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKlayOracle creates a new instance of KlayOracle, bound to a specific deployed contract.
func NewKlayOracle(address common.Address, backend bind.ContractBackend) (*KlayOracle, error) {
	contract, err := bindKlayOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KlayOracle{KlayOracleCaller: KlayOracleCaller{contract: contract}, KlayOracleTransactor: KlayOracleTransactor{contract: contract}, KlayOracleFilterer: KlayOracleFilterer{contract: contract}}, nil
}

// NewKlayOracleCaller creates a new read-only instance of KlayOracle, bound to a specific deployed contract.
func NewKlayOracleCaller(address common.Address, caller bind.ContractCaller) (*KlayOracleCaller, error) {
	contract, err := bindKlayOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KlayOracleCaller{contract: contract}, nil
}

// NewKlayOracleTransactor creates a new write-only instance of KlayOracle, bound to a specific deployed contract.
func NewKlayOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*KlayOracleTransactor, error) {
	contract, err := bindKlayOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KlayOracleTransactor{contract: contract}, nil
}

// NewKlayOracleFilterer creates a new log filterer instance of KlayOracle, bound to a specific deployed contract.
func NewKlayOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*KlayOracleFilterer, error) {
	contract, err := bindKlayOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KlayOracleFilterer{contract: contract}, nil
}

// bindKlayOracle binds a generic wrapper to an already deployed contract.
func bindKlayOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KlayOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KlayOracle *KlayOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KlayOracle.Contract.KlayOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KlayOracle *KlayOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KlayOracle.Contract.KlayOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KlayOracle *KlayOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KlayOracle.Contract.KlayOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KlayOracle *KlayOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KlayOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KlayOracle *KlayOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KlayOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KlayOracle *KlayOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KlayOracle.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_KlayOracle *KlayOracleCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KlayOracle.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_KlayOracle *KlayOracleSession) VERSION() (string, error) {
	return _KlayOracle.Contract.VERSION(&_KlayOracle.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_KlayOracle *KlayOracleCallerSession) VERSION() (string, error) {
	return _KlayOracle.Contract.VERSION(&_KlayOracle.CallOpts)
}

// VerifyKlaytnMessage is a free data retrieval call binding the contract method 0xbe2bb8cd.
//
// Solidity: function VerifyKlaytnMessage(bytes32 _hashedMessage, uint8 _v, bytes32 _r, bytes32 _s) pure returns(address)
func (_KlayOracle *KlayOracleCaller) VerifyKlaytnMessage(opts *bind.CallOpts, _hashedMessage [32]byte, _v uint8, _r [32]byte, _s [32]byte) (common.Address, error) {
	var out []interface{}
	err := _KlayOracle.contract.Call(opts, &out, "VerifyKlaytnMessage", _hashedMessage, _v, _r, _s)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifyKlaytnMessage is a free data retrieval call binding the contract method 0xbe2bb8cd.
//
// Solidity: function VerifyKlaytnMessage(bytes32 _hashedMessage, uint8 _v, bytes32 _r, bytes32 _s) pure returns(address)
func (_KlayOracle *KlayOracleSession) VerifyKlaytnMessage(_hashedMessage [32]byte, _v uint8, _r [32]byte, _s [32]byte) (common.Address, error) {
	return _KlayOracle.Contract.VerifyKlaytnMessage(&_KlayOracle.CallOpts, _hashedMessage, _v, _r, _s)
}

// VerifyKlaytnMessage is a free data retrieval call binding the contract method 0xbe2bb8cd.
//
// Solidity: function VerifyKlaytnMessage(bytes32 _hashedMessage, uint8 _v, bytes32 _r, bytes32 _s) pure returns(address)
func (_KlayOracle *KlayOracleCallerSession) VerifyKlaytnMessage(_hashedMessage [32]byte, _v uint8, _r [32]byte, _s [32]byte) (common.Address, error) {
	return _KlayOracle.Contract.VerifyKlaytnMessage(&_KlayOracle.CallOpts, _hashedMessage, _v, _r, _s)
}

// GetHash is a free data retrieval call binding the contract method 0x77d3bfb9.
//
// Solidity: function _getHash(bytes32 roundAnswer, uint256 roundTime) view returns(bytes32)
func (_KlayOracle *KlayOracleCaller) GetHash(opts *bind.CallOpts, roundAnswer [32]byte, roundTime *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _KlayOracle.contract.Call(opts, &out, "_getHash", roundAnswer, roundTime)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHash is a free data retrieval call binding the contract method 0x77d3bfb9.
//
// Solidity: function _getHash(bytes32 roundAnswer, uint256 roundTime) view returns(bytes32)
func (_KlayOracle *KlayOracleSession) GetHash(roundAnswer [32]byte, roundTime *big.Int) ([32]byte, error) {
	return _KlayOracle.Contract.GetHash(&_KlayOracle.CallOpts, roundAnswer, roundTime)
}

// GetHash is a free data retrieval call binding the contract method 0x77d3bfb9.
//
// Solidity: function _getHash(bytes32 roundAnswer, uint256 roundTime) view returns(bytes32)
func (_KlayOracle *KlayOracleCallerSession) GetHash(roundAnswer [32]byte, roundTime *big.Int) ([32]byte, error) {
	return _KlayOracle.Contract.GetHash(&_KlayOracle.CallOpts, roundAnswer, roundTime)
}

// AdapterId is a free data retrieval call binding the contract method 0x5fb86b01.
//
// Solidity: function adapterId() view returns(bytes32)
func (_KlayOracle *KlayOracleCaller) AdapterId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KlayOracle.contract.Call(opts, &out, "adapterId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AdapterId is a free data retrieval call binding the contract method 0x5fb86b01.
//
// Solidity: function adapterId() view returns(bytes32)
func (_KlayOracle *KlayOracleSession) AdapterId() ([32]byte, error) {
	return _KlayOracle.Contract.AdapterId(&_KlayOracle.CallOpts)
}

// AdapterId is a free data retrieval call binding the contract method 0x5fb86b01.
//
// Solidity: function adapterId() view returns(bytes32)
func (_KlayOracle *KlayOracleCallerSession) AdapterId() ([32]byte, error) {
	return _KlayOracle.Contract.AdapterId(&_KlayOracle.CallOpts)
}

// FulfilledCount is a free data retrieval call binding the contract method 0xc5c28fd1.
//
// Solidity: function fulfilledCount() view returns(uint256)
func (_KlayOracle *KlayOracleCaller) FulfilledCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KlayOracle.contract.Call(opts, &out, "fulfilledCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FulfilledCount is a free data retrieval call binding the contract method 0xc5c28fd1.
//
// Solidity: function fulfilledCount() view returns(uint256)
func (_KlayOracle *KlayOracleSession) FulfilledCount() (*big.Int, error) {
	return _KlayOracle.Contract.FulfilledCount(&_KlayOracle.CallOpts)
}

// FulfilledCount is a free data retrieval call binding the contract method 0xc5c28fd1.
//
// Solidity: function fulfilledCount() view returns(uint256)
func (_KlayOracle *KlayOracleCallerSession) FulfilledCount() (*big.Int, error) {
	return _KlayOracle.Contract.FulfilledCount(&_KlayOracle.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(bytes32 answer, uint256 roundTime, uint256 timestamp)
func (_KlayOracle *KlayOracleCaller) LatestRound(opts *bind.CallOpts) (struct {
	Answer    [32]byte
	RoundTime *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _KlayOracle.contract.Call(opts, &out, "latestRound")

	outstruct := new(struct {
		Answer    [32]byte
		RoundTime *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Answer = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.RoundTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(bytes32 answer, uint256 roundTime, uint256 timestamp)
func (_KlayOracle *KlayOracleSession) LatestRound() (struct {
	Answer    [32]byte
	RoundTime *big.Int
	Timestamp *big.Int
}, error) {
	return _KlayOracle.Contract.LatestRound(&_KlayOracle.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(bytes32 answer, uint256 roundTime, uint256 timestamp)
func (_KlayOracle *KlayOracleCallerSession) LatestRound() (struct {
	Answer    [32]byte
	RoundTime *big.Int
	Timestamp *big.Int
}, error) {
	return _KlayOracle.Contract.LatestRound(&_KlayOracle.CallOpts)
}

// NodeAddress is a free data retrieval call binding the contract method 0x2aa8481f.
//
// Solidity: function nodeAddress() view returns(address)
func (_KlayOracle *KlayOracleCaller) NodeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KlayOracle.contract.Call(opts, &out, "nodeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeAddress is a free data retrieval call binding the contract method 0x2aa8481f.
//
// Solidity: function nodeAddress() view returns(address)
func (_KlayOracle *KlayOracleSession) NodeAddress() (common.Address, error) {
	return _KlayOracle.Contract.NodeAddress(&_KlayOracle.CallOpts)
}

// NodeAddress is a free data retrieval call binding the contract method 0x2aa8481f.
//
// Solidity: function nodeAddress() view returns(address)
func (_KlayOracle *KlayOracleCallerSession) NodeAddress() (common.Address, error) {
	return _KlayOracle.Contract.NodeAddress(&_KlayOracle.CallOpts)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(bytes32 requestId, address nodeAddress, bytes32 adapterId, bytes4 callbackFunctionId, address callBackContract, bytes32 data, uint256 timestamp)
func (_KlayOracle *KlayOracleCaller) Requests(opts *bind.CallOpts, arg0 [32]byte) (struct {
	RequestId          [32]byte
	NodeAddress        common.Address
	AdapterId          [32]byte
	CallbackFunctionId [4]byte
	CallBackContract   common.Address
	Data               [32]byte
	Timestamp          *big.Int
}, error) {
	var out []interface{}
	err := _KlayOracle.contract.Call(opts, &out, "requests", arg0)

	outstruct := new(struct {
		RequestId          [32]byte
		NodeAddress        common.Address
		AdapterId          [32]byte
		CallbackFunctionId [4]byte
		CallBackContract   common.Address
		Data               [32]byte
		Timestamp          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.NodeAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AdapterId = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.CallbackFunctionId = *abi.ConvertType(out[3], new([4]byte)).(*[4]byte)
	outstruct.CallBackContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Data = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Timestamp = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(bytes32 requestId, address nodeAddress, bytes32 adapterId, bytes4 callbackFunctionId, address callBackContract, bytes32 data, uint256 timestamp)
func (_KlayOracle *KlayOracleSession) Requests(arg0 [32]byte) (struct {
	RequestId          [32]byte
	NodeAddress        common.Address
	AdapterId          [32]byte
	CallbackFunctionId [4]byte
	CallBackContract   common.Address
	Data               [32]byte
	Timestamp          *big.Int
}, error) {
	return _KlayOracle.Contract.Requests(&_KlayOracle.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(bytes32 requestId, address nodeAddress, bytes32 adapterId, bytes4 callbackFunctionId, address callBackContract, bytes32 data, uint256 timestamp)
func (_KlayOracle *KlayOracleCallerSession) Requests(arg0 [32]byte) (struct {
	RequestId          [32]byte
	NodeAddress        common.Address
	AdapterId          [32]byte
	CallbackFunctionId [4]byte
	CallBackContract   common.Address
	Data               [32]byte
	Timestamp          *big.Int
}, error) {
	return _KlayOracle.Contract.Requests(&_KlayOracle.CallOpts, arg0)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(bytes32 answer, uint256 roundTime, uint256 timestamp)
func (_KlayOracle *KlayOracleCaller) Rounds(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Answer    [32]byte
	RoundTime *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _KlayOracle.contract.Call(opts, &out, "rounds", arg0)

	outstruct := new(struct {
		Answer    [32]byte
		RoundTime *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Answer = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.RoundTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(bytes32 answer, uint256 roundTime, uint256 timestamp)
func (_KlayOracle *KlayOracleSession) Rounds(arg0 *big.Int) (struct {
	Answer    [32]byte
	RoundTime *big.Int
	Timestamp *big.Int
}, error) {
	return _KlayOracle.Contract.Rounds(&_KlayOracle.CallOpts, arg0)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(bytes32 answer, uint256 roundTime, uint256 timestamp)
func (_KlayOracle *KlayOracleCallerSession) Rounds(arg0 *big.Int) (struct {
	Answer    [32]byte
	RoundTime *big.Int
	Timestamp *big.Int
}, error) {
	return _KlayOracle.Contract.Rounds(&_KlayOracle.CallOpts, arg0)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_KlayOracle *KlayOracleCaller) SplitSignature(opts *bind.CallOpts, sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	var out []interface{}
	err := _KlayOracle.contract.Call(opts, &out, "splitSignature", sig)

	outstruct := new(struct {
		R [32]byte
		S [32]byte
		V uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.R = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.V = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_KlayOracle *KlayOracleSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _KlayOracle.Contract.SplitSignature(&_KlayOracle.CallOpts, sig)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_KlayOracle *KlayOracleCallerSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _KlayOracle.Contract.SplitSignature(&_KlayOracle.CallOpts, sig)
}

// NewOracleRequest is a paid mutator transaction binding the contract method 0xec2b83f1.
//
// Solidity: function newOracleRequest(bytes4 callbackFunctionId, address callBackContract) returns(bool)
func (_KlayOracle *KlayOracleTransactor) NewOracleRequest(opts *bind.TransactOpts, callbackFunctionId [4]byte, callBackContract common.Address) (*types.Transaction, error) {
	return _KlayOracle.contract.Transact(opts, "newOracleRequest", callbackFunctionId, callBackContract)
}

// NewOracleRequest is a paid mutator transaction binding the contract method 0xec2b83f1.
//
// Solidity: function newOracleRequest(bytes4 callbackFunctionId, address callBackContract) returns(bool)
func (_KlayOracle *KlayOracleSession) NewOracleRequest(callbackFunctionId [4]byte, callBackContract common.Address) (*types.Transaction, error) {
	return _KlayOracle.Contract.NewOracleRequest(&_KlayOracle.TransactOpts, callbackFunctionId, callBackContract)
}

// NewOracleRequest is a paid mutator transaction binding the contract method 0xec2b83f1.
//
// Solidity: function newOracleRequest(bytes4 callbackFunctionId, address callBackContract) returns(bool)
func (_KlayOracle *KlayOracleTransactorSession) NewOracleRequest(callbackFunctionId [4]byte, callBackContract common.Address) (*types.Transaction, error) {
	return _KlayOracle.Contract.NewOracleRequest(&_KlayOracle.TransactOpts, callbackFunctionId, callBackContract)
}

// NewRoundData is a paid mutator transaction binding the contract method 0x09c980c1.
//
// Solidity: function newRoundData(uint256 roundTime, bytes32 roundAnswer, bytes signature) returns(bool)
func (_KlayOracle *KlayOracleTransactor) NewRoundData(opts *bind.TransactOpts, roundTime *big.Int, roundAnswer [32]byte, signature []byte) (*types.Transaction, error) {
	return _KlayOracle.contract.Transact(opts, "newRoundData", roundTime, roundAnswer, signature)
}

// NewRoundData is a paid mutator transaction binding the contract method 0x09c980c1.
//
// Solidity: function newRoundData(uint256 roundTime, bytes32 roundAnswer, bytes signature) returns(bool)
func (_KlayOracle *KlayOracleSession) NewRoundData(roundTime *big.Int, roundAnswer [32]byte, signature []byte) (*types.Transaction, error) {
	return _KlayOracle.Contract.NewRoundData(&_KlayOracle.TransactOpts, roundTime, roundAnswer, signature)
}

// NewRoundData is a paid mutator transaction binding the contract method 0x09c980c1.
//
// Solidity: function newRoundData(uint256 roundTime, bytes32 roundAnswer, bytes signature) returns(bool)
func (_KlayOracle *KlayOracleTransactorSession) NewRoundData(roundTime *big.Int, roundAnswer [32]byte, signature []byte) (*types.Transaction, error) {
	return _KlayOracle.Contract.NewRoundData(&_KlayOracle.TransactOpts, roundTime, roundAnswer, signature)
}

// KlayOracleNewOracleRequestIterator is returned from FilterNewOracleRequest and is used to iterate over the raw logs and unpacked data for NewOracleRequest events raised by the KlayOracle contract.
type KlayOracleNewOracleRequestIterator struct {
	Event *KlayOracleNewOracleRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KlayOracleNewOracleRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KlayOracleNewOracleRequest)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KlayOracleNewOracleRequest)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KlayOracleNewOracleRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KlayOracleNewOracleRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KlayOracleNewOracleRequest represents a NewOracleRequest event raised by the KlayOracle contract.
type KlayOracleNewOracleRequest struct {
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewOracleRequest is a free log retrieval operation binding the contract event 0x7a6060117db4511b5ffd88545b9b43962c0ae441d1f07d05c86b869490d88369.
//
// Solidity: event NewOracleRequest(bytes32 requestId)
func (_KlayOracle *KlayOracleFilterer) FilterNewOracleRequest(opts *bind.FilterOpts) (*KlayOracleNewOracleRequestIterator, error) {

	logs, sub, err := _KlayOracle.contract.FilterLogs(opts, "NewOracleRequest")
	if err != nil {
		return nil, err
	}
	return &KlayOracleNewOracleRequestIterator{contract: _KlayOracle.contract, event: "NewOracleRequest", logs: logs, sub: sub}, nil
}

// WatchNewOracleRequest is a free log subscription operation binding the contract event 0x7a6060117db4511b5ffd88545b9b43962c0ae441d1f07d05c86b869490d88369.
//
// Solidity: event NewOracleRequest(bytes32 requestId)
func (_KlayOracle *KlayOracleFilterer) WatchNewOracleRequest(opts *bind.WatchOpts, sink chan<- *KlayOracleNewOracleRequest) (event.Subscription, error) {

	logs, sub, err := _KlayOracle.contract.WatchLogs(opts, "NewOracleRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KlayOracleNewOracleRequest)
				if err := _KlayOracle.contract.UnpackLog(event, "NewOracleRequest", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewOracleRequest is a log parse operation binding the contract event 0x7a6060117db4511b5ffd88545b9b43962c0ae441d1f07d05c86b869490d88369.
//
// Solidity: event NewOracleRequest(bytes32 requestId)
func (_KlayOracle *KlayOracleFilterer) ParseNewOracleRequest(log types.Log) (*KlayOracleNewOracleRequest, error) {
	event := new(KlayOracleNewOracleRequest)
	if err := _KlayOracle.contract.UnpackLog(event, "NewOracleRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
