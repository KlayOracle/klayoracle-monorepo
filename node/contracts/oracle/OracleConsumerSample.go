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

// OracleConsumerSampleABI is the input ABI used to generate the binding from.
const OracleConsumerSampleABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracleAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"klayOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_klayOutput\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapEthtoKlay\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OracleConsumerSampleBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const OracleConsumerSampleBinRuntime = ``

// OracleConsumerSampleBin is the compiled bytecode used for deploying new contracts.
var OracleConsumerSampleBin = "0x60a060405234801561001057600080fd5b5060405161051f38038061051f83398181016040528101906100329190610108565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361006b57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff168152505050610135565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100d5826100aa565b9050919050565b6100e5816100ca565b81146100f057600080fd5b50565b600081519050610102816100dc565b92915050565b60006020828403121561011e5761011d6100a5565b5b600061012c848285016100f3565b91505092915050565b6080516103c96101566000396000818160cc015261018a01526103c96000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063760b78b01461005157806394b918de1461006f578063a89ae4ba1461008b578063d92a047c146100a9575b600080fd5b6100596100c7565b60405161006691906101cd565b60405180910390f35b61008960048036038101906100849190610223565b61017e565b005b610093610188565b6040516100a09190610291565b60405180910390f35b6100b16101ac565b6040516100be91906102bb565b60405180910390f35b6000807f0000000000000000000000000000000000000000000000000000000000000000905060008173ffffffffffffffffffffffffffffffffffffffff1663ec2b83f16394b918de60e01b306040518363ffffffff1660e01b8152600401610131929190610311565b6020604051808303816000875af1158015610150573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101749190610366565b9050809250505090565b8060008190555050565b7f000000000000000000000000000000000000000000000000000000000000000081565b60005481565b60008115159050919050565b6101c7816101b2565b82525050565b60006020820190506101e260008301846101be565b92915050565b600080fd5b6000819050919050565b610200816101ed565b811461020b57600080fd5b50565b60008135905061021d816101f7565b92915050565b600060208284031215610239576102386101e8565b5b60006102478482850161020e565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061027b82610250565b9050919050565b61028b81610270565b82525050565b60006020820190506102a66000830184610282565b92915050565b6102b5816101ed565b82525050565b60006020820190506102d060008301846102ac565b92915050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61030b816102d6565b82525050565b60006040820190506103266000830185610302565b6103336020830184610282565b9392505050565b610343816101b2565b811461034e57600080fd5b50565b6000815190506103608161033a565b92915050565b60006020828403121561037c5761037b6101e8565b5b600061038a84828501610351565b9150509291505056fea26469706673582212206cd850d5533d8fae277c2d521634c84ee372714ffcbf677bd16b1eedd75ad4c264736f6c63430008110033"

// DeployOracleConsumerSample deploys a new Klaytn contract, binding an instance of OracleConsumerSample to it.
func DeployOracleConsumerSample(auth *bind.TransactOpts, backend bind.ContractBackend, _oracleAddress common.Address) (common.Address, *types.Transaction, *OracleConsumerSample, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleConsumerSampleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OracleConsumerSampleBin), backend, _oracleAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OracleConsumerSample{OracleConsumerSampleCaller: OracleConsumerSampleCaller{contract: contract}, OracleConsumerSampleTransactor: OracleConsumerSampleTransactor{contract: contract}, OracleConsumerSampleFilterer: OracleConsumerSampleFilterer{contract: contract}}, nil
}

// OracleConsumerSample is an auto generated Go binding around a Klaytn contract.
type OracleConsumerSample struct {
	OracleConsumerSampleCaller     // Read-only binding to the contract
	OracleConsumerSampleTransactor // Write-only binding to the contract
	OracleConsumerSampleFilterer   // Log filterer for contract events
}

// OracleConsumerSampleCaller is an auto generated read-only Go binding around a Klaytn contract.
type OracleConsumerSampleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleConsumerSampleTransactor is an auto generated write-only Go binding around a Klaytn contract.
type OracleConsumerSampleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleConsumerSampleFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type OracleConsumerSampleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleConsumerSampleSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type OracleConsumerSampleSession struct {
	Contract     *OracleConsumerSample // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OracleConsumerSampleCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type OracleConsumerSampleCallerSession struct {
	Contract *OracleConsumerSampleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// OracleConsumerSampleTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type OracleConsumerSampleTransactorSession struct {
	Contract     *OracleConsumerSampleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// OracleConsumerSampleRaw is an auto generated low-level Go binding around a Klaytn contract.
type OracleConsumerSampleRaw struct {
	Contract *OracleConsumerSample // Generic contract binding to access the raw methods on
}

// OracleConsumerSampleCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type OracleConsumerSampleCallerRaw struct {
	Contract *OracleConsumerSampleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleConsumerSampleTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
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
	parsed, err := abi.JSON(strings.NewReader(OracleConsumerSampleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleConsumerSample *OracleConsumerSampleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
func (_OracleConsumerSample *OracleConsumerSampleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
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
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleConsumerSample.contract.Call(opts, out, "klayOutput")
	return *ret0, err
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
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OracleConsumerSample.contract.Call(opts, out, "oracleAddress")
	return *ret0, err
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

// SwapEthtoKlay is a paid mutator transaction binding the contract method 0x760b78b0.
//
// Solidity: function swapEthtoKlay() returns(bool)
func (_OracleConsumerSample *OracleConsumerSampleTransactor) SwapEthtoKlay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleConsumerSample.contract.Transact(opts, "swapEthtoKlay")
}

// SwapEthtoKlay is a paid mutator transaction binding the contract method 0x760b78b0.
//
// Solidity: function swapEthtoKlay() returns(bool)
func (_OracleConsumerSample *OracleConsumerSampleSession) SwapEthtoKlay() (*types.Transaction, error) {
	return _OracleConsumerSample.Contract.SwapEthtoKlay(&_OracleConsumerSample.TransactOpts)
}

// SwapEthtoKlay is a paid mutator transaction binding the contract method 0x760b78b0.
//
// Solidity: function swapEthtoKlay() returns(bool)
func (_OracleConsumerSample *OracleConsumerSampleTransactorSession) SwapEthtoKlay() (*types.Transaction, error) {
	return _OracleConsumerSample.Contract.SwapEthtoKlay(&_OracleConsumerSample.TransactOpts)
}
