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

// OracleConsumerSampleMetaData contains all meta data concerning the OracleConsumerSample contract.
var OracleConsumerSampleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracleAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"klayOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_klayOutput\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapUsdtoKlay\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b506040516105ff3803806105ff83398181016040528101906100319190610100565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610068575f80fd5b8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250505061012b565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100cf826100a6565b9050919050565b6100df816100c5565b81146100e9575f80fd5b50565b5f815190506100fa816100d6565b92915050565b5f60208284031215610115576101146100a2565b5b5f610122848285016100ec565b91505092915050565b6080516104af6101505f395f818160c801528181610178015261020f01526104af5ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c80638fe9e70f1461004e57806394b918de1461006c578063a89ae4ba14610088578063d92a047c146100a6575b5f80fd5b6100566100c4565b6040516100639190610250565b60405180910390f35b610086600480360381019061008191906102a0565b610176565b005b61009061020d565b60405161009d919061030a565b60405180910390f35b6100ae610231565b6040516100bb9190610332565b60405180910390f35b5f807f000000000000000000000000000000000000000000000000000000000000000090505f8173ffffffffffffffffffffffffffffffffffffffff1663ec2b83f16394b918de60e01b306040518363ffffffff1660e01b815260040161012c929190610385565b6020604051808303815f875af1158015610148573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061016c91906103d6565b9050809250505090565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610204576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101fb9061045b565b60405180910390fd5b805f8190555050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f5481565b5f8115159050919050565b61024a81610236565b82525050565b5f6020820190506102635f830184610241565b92915050565b5f80fd5b5f819050919050565b61027f8161026d565b8114610289575f80fd5b50565b5f8135905061029a81610276565b92915050565b5f602082840312156102b5576102b4610269565b5b5f6102c28482850161028c565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102f4826102cb565b9050919050565b610304816102ea565b82525050565b5f60208201905061031d5f8301846102fb565b92915050565b61032c8161026d565b82525050565b5f6020820190506103455f830184610323565b92915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61037f8161034b565b82525050565b5f6040820190506103985f830185610376565b6103a560208301846102fb565b9392505050565b6103b581610236565b81146103bf575f80fd5b50565b5f815190506103d0816103ac565b92915050565b5f602082840312156103eb576103ea610269565b5b5f6103f8848285016103c2565b91505092915050565b5f82825260208201905092915050565b7f6e6f7420616c6c6f7765640000000000000000000000000000000000000000005f82015250565b5f610445600b83610401565b915061045082610411565b602082019050919050565b5f6020820190508181035f83015261047281610439565b905091905056fea26469706673582212205626fcc8cfddbf341d83fbfe5805f47f62c9fa58e976ae28ea538fed4410273c64736f6c63430008140033",
}

// OracleConsumerSampleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleConsumerSampleMetaData.ABI instead.
var OracleConsumerSampleABI = OracleConsumerSampleMetaData.ABI

// OracleConsumerSampleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OracleConsumerSampleMetaData.Bin instead.
var OracleConsumerSampleBin = OracleConsumerSampleMetaData.Bin

// DeployOracleConsumerSample deploys a new Ethereum contract, binding an instance of OracleConsumerSample to it.
func DeployOracleConsumerSample(auth *bind.TransactOpts, backend bind.ContractBackend, _oracleAddress common.Address) (common.Address, *types.Transaction, *OracleConsumerSample, error) {
	parsed, err := OracleConsumerSampleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OracleConsumerSampleBin), backend, _oracleAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OracleConsumerSample{OracleConsumerSampleCaller: OracleConsumerSampleCaller{contract: contract}, OracleConsumerSampleTransactor: OracleConsumerSampleTransactor{contract: contract}, OracleConsumerSampleFilterer: OracleConsumerSampleFilterer{contract: contract}}, nil
}

// OracleConsumerSample is an auto generated Go binding around an Ethereum contract.
type OracleConsumerSample struct {
	OracleConsumerSampleCaller     // Read-only binding to the contract
	OracleConsumerSampleTransactor // Write-only binding to the contract
	OracleConsumerSampleFilterer   // Log filterer for contract events
}

// OracleConsumerSampleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleConsumerSampleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleConsumerSampleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleConsumerSampleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleConsumerSampleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleConsumerSampleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleConsumerSampleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleConsumerSampleSession struct {
	Contract     *OracleConsumerSample // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OracleConsumerSampleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleConsumerSampleCallerSession struct {
	Contract *OracleConsumerSampleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// OracleConsumerSampleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleConsumerSampleTransactorSession struct {
	Contract     *OracleConsumerSampleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// OracleConsumerSampleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleConsumerSampleRaw struct {
	Contract *OracleConsumerSample // Generic contract binding to access the raw methods on
}

// OracleConsumerSampleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleConsumerSampleCallerRaw struct {
	Contract *OracleConsumerSampleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleConsumerSampleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleConsumerSampleTransactorRaw struct {
	Contract *OracleConsumerSampleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleConsumerSample creates a new instance of OracleConsumerSample, bound to a specific deployed contract.
func NewOracleConsumerSample(address common.Address, backend bind.ContractBackend) (*OracleConsumerSample, error) {
	contract, err := bindOracleConsumerSample(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleConsumerSample{OracleConsumerSampleCaller: OracleConsumerSampleCaller{contract: contract}, OracleConsumerSampleTransactor: OracleConsumerSampleTransactor{contract: contract}, OracleConsumerSampleFilterer: OracleConsumerSampleFilterer{contract: contract}}, nil
}

// NewOracleConsumerSampleCaller creates a new read-only instance of OracleConsumerSample, bound to a specific deployed contract.
func NewOracleConsumerSampleCaller(address common.Address, caller bind.ContractCaller) (*OracleConsumerSampleCaller, error) {
	contract, err := bindOracleConsumerSample(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleConsumerSampleCaller{contract: contract}, nil
}

// NewOracleConsumerSampleTransactor creates a new write-only instance of OracleConsumerSample, bound to a specific deployed contract.
func NewOracleConsumerSampleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleConsumerSampleTransactor, error) {
	contract, err := bindOracleConsumerSample(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleConsumerSampleTransactor{contract: contract}, nil
}

// NewOracleConsumerSampleFilterer creates a new log filterer instance of OracleConsumerSample, bound to a specific deployed contract.
func NewOracleConsumerSampleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleConsumerSampleFilterer, error) {
	contract, err := bindOracleConsumerSample(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleConsumerSampleFilterer{contract: contract}, nil
}

// bindOracleConsumerSample binds a generic wrapper to an already deployed contract.
func bindOracleConsumerSample(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleConsumerSampleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleConsumerSample *OracleConsumerSampleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleConsumerSample.Contract.OracleConsumerSampleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleConsumerSample *OracleConsumerSampleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleConsumerSample.Contract.OracleConsumerSampleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleConsumerSample *OracleConsumerSampleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleConsumerSample.Contract.OracleConsumerSampleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleConsumerSample *OracleConsumerSampleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleConsumerSample.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleConsumerSample *OracleConsumerSampleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleConsumerSample.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleConsumerSample *OracleConsumerSampleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleConsumerSample.Contract.contract.Transact(opts, method, params...)
}

// KlayOutput is a free data retrieval call binding the contract method 0xd92a047c.
//
// Solidity: function klayOutput() view returns(uint256)
func (_OracleConsumerSample *OracleConsumerSampleCaller) KlayOutput(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleConsumerSample.contract.Call(opts, &out, "klayOutput")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KlayOutput is a free data retrieval call binding the contract method 0xd92a047c.
//
// Solidity: function klayOutput() view returns(uint256)
func (_OracleConsumerSample *OracleConsumerSampleSession) KlayOutput() (*big.Int, error) {
	return _OracleConsumerSample.Contract.KlayOutput(&_OracleConsumerSample.CallOpts)
}

// KlayOutput is a free data retrieval call binding the contract method 0xd92a047c.
//
// Solidity: function klayOutput() view returns(uint256)
func (_OracleConsumerSample *OracleConsumerSampleCallerSession) KlayOutput() (*big.Int, error) {
	return _OracleConsumerSample.Contract.KlayOutput(&_OracleConsumerSample.CallOpts)
}

// OracleAddress is a free data retrieval call binding the contract method 0xa89ae4ba.
//
// Solidity: function oracleAddress() view returns(address)
func (_OracleConsumerSample *OracleConsumerSampleCaller) OracleAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleConsumerSample.contract.Call(opts, &out, "oracleAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleAddress is a free data retrieval call binding the contract method 0xa89ae4ba.
//
// Solidity: function oracleAddress() view returns(address)
func (_OracleConsumerSample *OracleConsumerSampleSession) OracleAddress() (common.Address, error) {
	return _OracleConsumerSample.Contract.OracleAddress(&_OracleConsumerSample.CallOpts)
}

// OracleAddress is a free data retrieval call binding the contract method 0xa89ae4ba.
//
// Solidity: function oracleAddress() view returns(address)
func (_OracleConsumerSample *OracleConsumerSampleCallerSession) OracleAddress() (common.Address, error) {
	return _OracleConsumerSample.Contract.OracleAddress(&_OracleConsumerSample.CallOpts)
}

// Swap is a paid mutator transaction binding the contract method 0x94b918de.
//
// Solidity: function swap(uint256 _klayOutput) returns()
func (_OracleConsumerSample *OracleConsumerSampleTransactor) Swap(opts *bind.TransactOpts, _klayOutput *big.Int) (*types.Transaction, error) {
	return _OracleConsumerSample.contract.Transact(opts, "swap", _klayOutput)
}

// Swap is a paid mutator transaction binding the contract method 0x94b918de.
//
// Solidity: function swap(uint256 _klayOutput) returns()
func (_OracleConsumerSample *OracleConsumerSampleSession) Swap(_klayOutput *big.Int) (*types.Transaction, error) {
	return _OracleConsumerSample.Contract.Swap(&_OracleConsumerSample.TransactOpts, _klayOutput)
}

// Swap is a paid mutator transaction binding the contract method 0x94b918de.
//
// Solidity: function swap(uint256 _klayOutput) returns()
func (_OracleConsumerSample *OracleConsumerSampleTransactorSession) Swap(_klayOutput *big.Int) (*types.Transaction, error) {
	return _OracleConsumerSample.Contract.Swap(&_OracleConsumerSample.TransactOpts, _klayOutput)
}

// SwapUsdtoKlay is a paid mutator transaction binding the contract method 0x8fe9e70f.
//
// Solidity: function swapUsdtoKlay() returns(bool)
func (_OracleConsumerSample *OracleConsumerSampleTransactor) SwapUsdtoKlay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleConsumerSample.contract.Transact(opts, "swapUsdtoKlay")
}

// SwapUsdtoKlay is a paid mutator transaction binding the contract method 0x8fe9e70f.
//
// Solidity: function swapUsdtoKlay() returns(bool)
func (_OracleConsumerSample *OracleConsumerSampleSession) SwapUsdtoKlay() (*types.Transaction, error) {
	return _OracleConsumerSample.Contract.SwapUsdtoKlay(&_OracleConsumerSample.TransactOpts)
}

// SwapUsdtoKlay is a paid mutator transaction binding the contract method 0x8fe9e70f.
//
// Solidity: function swapUsdtoKlay() returns(bool)
func (_OracleConsumerSample *OracleConsumerSampleTransactorSession) SwapUsdtoKlay() (*types.Transaction, error) {
	return _OracleConsumerSample.Contract.SwapUsdtoKlay(&_OracleConsumerSample.TransactOpts)
}
