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

// OracleProviderSampleABI is the input ABI used to generate the binding from.
const OracleProviderSampleABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_adapterId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"NewOracleRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adapterId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fulfilledCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"answer\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"roundTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"callBackContract\",\"type\":\"address\"}],\"name\":\"newOracleRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"newRoundData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"adapterId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"callBackContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rounds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"answer\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"roundTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OracleProviderSampleBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const OracleProviderSampleBinRuntime = ``

// OracleProviderSampleBin is the compiled bytecode used for deploying new contracts.
var OracleProviderSampleBin = "0x60c060405260006001553480156200001657600080fd5b5060405162001bda38038062001bda83398181016040528101906200003c919062000129565b81818173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508060a081815250505050505062000170565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000b68262000089565b9050919050565b620000c881620000a9565b8114620000d457600080fd5b50565b600081519050620000e881620000bd565b92915050565b6000819050919050565b6200010381620000ee565b81146200010f57600080fd5b50565b6000815190506200012381620000f8565b92915050565b6000806040838503121562000143576200014262000084565b5b60006200015385828601620000d7565b9250506020620001668582860162000112565b9150509250929050565b60805160a051611a1a620001c06000396000818161021e0152818161057201526106160152600081816101fa0152818161025c015281816103010152818161055101526105da0152611a1a6000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80638c65c81f116100665780638c65c81f146101245780639d86698514610156578063c5c28fd11461018c578063ec2b83f1146101aa578063ffa1ad74146101da57610093565b80632aa8481f146100985780635fb86b01146100b6578063668a0f02146100d457806377f8632f146100f4575b600080fd5b6100a06101f8565b6040516100ad9190610d1b565b60405180910390f35b6100be61021c565b6040516100cb9190610d4f565b60405180910390f35b6100dc610240565b6040516100eb93929190610d83565b60405180910390f35b61010e60048036038101906101099190610f6c565b610258565b60405161011b919061100a565b60405180910390f35b61013e60048036038101906101399190611025565b610431565b60405161014d93929190610d83565b60405180910390f35b610170600480360381019061016b9190611052565b61046b565b60405161018397969594939291906110ba565b60405180910390f35b6101946104fa565b6040516101a19190611129565b60405180910390f35b6101c460048036038101906101bf919061119c565b610500565b6040516101d1919061100a565b60405180910390f35b6101e261071e565b6040516101ef919061125b565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b60058060000154908060010154908060020154905083565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102df906112c9565b60405180910390fd5b60006102fd838561075790919063ffffffff16565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461038d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161038490611335565b60405180910390fd5b600060405180606001604052808781526020018881526020014281525090506004819080600181540180825580915050600190039060005260206000209060030201600090919091909150600082015181600001556020820151816001015560408201518160020155505080600560008201518160000155602082015181600101556040820151816002015590505085600281905550600192505050949350505050565b6004818154811061044157600080fd5b90600052602060002090600302016000915090508060000154908060010154908060020154905083565b60006020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030160009054906101000a900460e01b908060030160049054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154908060050154905087565b60015481565b600061050b3361077e565b61054a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610541906113c7565b60405180910390fd5b60006003547f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008686600254426040516020016105ad9796959493929190611492565b60405160208183030381529060405280519060200120905060006040518060e001604052808381526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000008152602001867bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020018573ffffffffffffffffffffffffffffffffffffffff16815260200160056000015481526020014281525090507f7a6060117db4511b5ffd88545b9b43962c0ae441d1f07d05c86b869490d88369826040516106bb9190610d4f565b60405180910390a160006106cf83836107f7565b905080610711576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070890611585565b60405180910390fd5b6001935050505092915050565b6040518060400160405280600681526020017f76312e302e30000000000000000000000000000000000000000000000000000081525081565b600080600061076685856109ed565b9150915061077381610a3e565b819250505092915050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036107ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107e5906115f1565b60405180910390fd5b60019050919050565b600061080282610ba4565b610841576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108389061165d565b60405180910390fd5b816000808581526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015560608201518160030160006101000a81548163ffffffff021916908360e01c021790555060808201518160030160046101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a0820151816004015560c0820151816005015590505060016000815480929190610944906116ac565b91905055506000826080015173ffffffffffffffffffffffffffffffffffffffff1683606001518460a001516040516020016109819291906116f4565b60405160208183030381529060405260405161099d9190611767565b6000604051808303816000865af19150503d80600081146109da576040519150601f19603f3d011682016040523d82523d6000602084013e6109df565b606091505b505090508091505092915050565b6000806041835103610a2e5760008060006020860151925060408601519150606086015160001a9050610a2287828585610bf8565b94509450505050610a37565b60006002915091505b9250929050565b60006004811115610a5257610a5161177e565b5b816004811115610a6557610a6461177e565b5b0315610ba15760016004811115610a7f57610a7e61177e565b5b816004811115610a9257610a9161177e565b5b03610ad2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac9906117f9565b60405180910390fd5b60026004811115610ae657610ae561177e565b5b816004811115610af957610af861177e565b5b03610b39576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3090611865565b60405180910390fd5b60036004811115610b4d57610b4c61177e565b5b816004811115610b6057610b5f61177e565b5b03610ba0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b97906118f7565b60405180910390fd5b5b50565b60008060001b8260a0015103610bef576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610be690611963565b60405180910390fd5b60019050919050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08360001c1115610c33576000600391509150610cd1565b600060018787878760405160008152602001604052604051610c58949392919061199f565b6020604051602081039080840390855afa158015610c7a573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610cc857600060019250925050610cd1565b80600092509250505b94509492505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610d0582610cda565b9050919050565b610d1581610cfa565b82525050565b6000602082019050610d306000830184610d0c565b92915050565b6000819050919050565b610d4981610d36565b82525050565b6000602082019050610d646000830184610d40565b92915050565b6000819050919050565b610d7d81610d6a565b82525050565b6000606082019050610d986000830186610d40565b610da56020830185610d74565b610db26040830184610d74565b949350505050565b6000604051905090565b600080fd5b600080fd5b610dd781610d6a565b8114610de257600080fd5b50565b600081359050610df481610dce565b92915050565b610e0381610d36565b8114610e0e57600080fd5b50565b600081359050610e2081610dfa565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610e7982610e30565b810181811067ffffffffffffffff82111715610e9857610e97610e41565b5b80604052505050565b6000610eab610dba565b9050610eb78282610e70565b919050565b600067ffffffffffffffff821115610ed757610ed6610e41565b5b610ee082610e30565b9050602081019050919050565b82818337600083830152505050565b6000610f0f610f0a84610ebc565b610ea1565b905082815260208101848484011115610f2b57610f2a610e2b565b5b610f36848285610eed565b509392505050565b600082601f830112610f5357610f52610e26565b5b8135610f63848260208601610efc565b91505092915050565b60008060008060808587031215610f8657610f85610dc4565b5b6000610f9487828801610de5565b9450506020610fa587828801610e11565b9350506040610fb687828801610e11565b925050606085013567ffffffffffffffff811115610fd757610fd6610dc9565b5b610fe387828801610f3e565b91505092959194509250565b60008115159050919050565b61100481610fef565b82525050565b600060208201905061101f6000830184610ffb565b92915050565b60006020828403121561103b5761103a610dc4565b5b600061104984828501610de5565b91505092915050565b60006020828403121561106857611067610dc4565b5b600061107684828501610e11565b91505092915050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6110b48161107f565b82525050565b600060e0820190506110cf600083018a610d40565b6110dc6020830189610d0c565b6110e96040830188610d40565b6110f660608301876110ab565b6111036080830186610d0c565b61111060a0830185610d40565b61111d60c0830184610d74565b98975050505050505050565b600060208201905061113e6000830184610d74565b92915050565b61114d8161107f565b811461115857600080fd5b50565b60008135905061116a81611144565b92915050565b61117981610cfa565b811461118457600080fd5b50565b60008135905061119681611170565b92915050565b600080604083850312156111b3576111b2610dc4565b5b60006111c18582860161115b565b92505060206111d285828601611187565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b838110156112165780820151818401526020810190506111fb565b60008484015250505050565b600061122d826111dc565b61123781856111e7565b93506112478185602086016111f8565b61125081610e30565b840191505092915050565b600060208201905081810360008301526112758184611222565b905092915050565b7f4f7261636c653a206e6f646520756e6b6e6f776e000000000000000000000000600082015250565b60006112b36014836111e7565b91506112be8261127d565b602082019050919050565b600060208201905081810360008301526112e2816112a6565b9050919050565b7f4f7261636c653a20496e76616c69642064617461000000000000000000000000600082015250565b600061131f6014836111e7565b915061132a826112e9565b602082019050919050565b6000602082019050818103600083015261134e81611312565b9050919050565b7f4f7261636c653a20636f6e73756d6572206973206e6f742077686974656c697360008201527f7465640000000000000000000000000000000000000000000000000000000000602082015250565b60006113b16023836111e7565b91506113bc82611355565b604082019050919050565b600060208201905081810360008301526113e0816113a4565b9050919050565b6000819050919050565b6114026113fd82610d36565b6113e7565b82525050565b60008160601b9050919050565b600061142082611408565b9050919050565b600061143282611415565b9050919050565b61144a61144582610cfa565b611427565b82525050565b6000819050919050565b61146b6114668261107f565b611450565b82525050565b6000819050919050565b61148c61148782610d6a565b611471565b82525050565b600061149e828a6113f1565b6020820191506114ae8289611439565b6014820191506114be82886113f1565b6020820191506114ce828761145a565b6004820191506114de8286611439565b6014820191506114ee82856113f1565b6020820191506114fe828461147b565b60208201915081905098975050505050505050565b7f4b6c61794f7261636c653a2072657175657374206e6f742066756c66696c6c6560008201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b600061156f6021836111e7565b915061157a82611513565b604082019050919050565b6000602082019050818103600083015261159e81611562565b9050919050565b7f4b6c61794f7261636c653a20636f6e73756d657220697320696e76616c696400600082015250565b60006115db601f836111e7565b91506115e6826115a5565b602082019050919050565b6000602082019050818103600083015261160a816115ce565b9050919050565b7f4b6c61794f7261636c653a2073756273637269626520746f2044500000000000600082015250565b6000611647601b836111e7565b915061165282611611565b602082019050919050565b600060208201905081810360008301526116768161163a565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006116b782610d6a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036116e9576116e861167d565b5b600182019050919050565b6000611700828561145a565b60048201915061171082846113f1565b6020820191508190509392505050565b600081519050919050565b600081905092915050565b600061174182611720565b61174b818561172b565b935061175b8185602086016111f8565b80840191505092915050565b60006117738284611736565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f45434453413a20696e76616c6964207369676e61747572650000000000000000600082015250565b60006117e36018836111e7565b91506117ee826117ad565b602082019050919050565b60006020820190508181036000830152611812816117d6565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265206c656e67746800600082015250565b600061184f601f836111e7565b915061185a82611819565b602082019050919050565b6000602082019050818103600083015261187e81611842565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b60006118e16022836111e7565b91506118ec82611885565b604082019050919050565b60006020820190508181036000830152611910816118d4565b9050919050565b7f4b6c61794f7261636c653a206461746120697320656d70747900000000000000600082015250565b600061194d6019836111e7565b915061195882611917565b602082019050919050565b6000602082019050818103600083015261197c81611940565b9050919050565b600060ff82169050919050565b61199981611983565b82525050565b60006080820190506119b46000830187610d40565b6119c16020830186611990565b6119ce6040830185610d40565b6119db6060830184610d40565b9594505050505056fea26469706673582212202dea01abb9678047156ddab32892f60b72e8c6c465f4e4f4ddc6122cbf29a7b664736f6c63430008110033"

// DeployOracleProviderSample deploys a new Klaytn contract, binding an instance of OracleProviderSample to it.
func DeployOracleProviderSample(auth *bind.TransactOpts, backend bind.ContractBackend, _nodeAddress common.Address, _adapterId [32]byte) (common.Address, *types.Transaction, *OracleProviderSample, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleProviderSampleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OracleProviderSampleBin), backend, _nodeAddress, _adapterId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OracleProviderSample{OracleProviderSampleCaller: OracleProviderSampleCaller{contract: contract}, OracleProviderSampleTransactor: OracleProviderSampleTransactor{contract: contract}, OracleProviderSampleFilterer: OracleProviderSampleFilterer{contract: contract}}, nil
}

// OracleProviderSample is an auto generated Go binding around a Klaytn contract.
type OracleProviderSample struct {
	OracleProviderSampleCaller     // Read-only binding to the contract
	OracleProviderSampleTransactor // Write-only binding to the contract
	OracleProviderSampleFilterer   // Log filterer for contract events
}

// OracleProviderSampleCaller is an auto generated read-only Go binding around a Klaytn contract.
type OracleProviderSampleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleProviderSampleTransactor is an auto generated write-only Go binding around a Klaytn contract.
type OracleProviderSampleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleProviderSampleFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type OracleProviderSampleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleProviderSampleSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type OracleProviderSampleSession struct {
	Contract     *OracleProviderSample // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OracleProviderSampleCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type OracleProviderSampleCallerSession struct {
	Contract *OracleProviderSampleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// OracleProviderSampleTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type OracleProviderSampleTransactorSession struct {
	Contract     *OracleProviderSampleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// OracleProviderSampleRaw is an auto generated low-level Go binding around a Klaytn contract.
type OracleProviderSampleRaw struct {
	Contract *OracleProviderSample // Generic contract binding to access the raw methods on
}

// OracleProviderSampleCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type OracleProviderSampleCallerRaw struct {
	Contract *OracleProviderSampleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleProviderSampleTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type OracleProviderSampleTransactorRaw struct {
	Contract *OracleProviderSampleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleProviderSample creates a new instance of OracleProviderSample, bound to a specific deployed contract.
func NewOracleProviderSample(address common.Address, backend bind.ContractBackend) (*OracleProviderSample, error) {
	contract, err := bindOracleProviderSample(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleProviderSample{OracleProviderSampleCaller: OracleProviderSampleCaller{contract: contract}, OracleProviderSampleTransactor: OracleProviderSampleTransactor{contract: contract}, OracleProviderSampleFilterer: OracleProviderSampleFilterer{contract: contract}}, nil
}

// NewOracleProviderSampleCaller creates a new read-only instance of OracleProviderSample, bound to a specific deployed contract.
func NewOracleProviderSampleCaller(address common.Address, caller bind.ContractCaller) (*OracleProviderSampleCaller, error) {
	contract, err := bindOracleProviderSample(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleProviderSampleCaller{contract: contract}, nil
}

// NewOracleProviderSampleTransactor creates a new write-only instance of OracleProviderSample, bound to a specific deployed contract.
func NewOracleProviderSampleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleProviderSampleTransactor, error) {
	contract, err := bindOracleProviderSample(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleProviderSampleTransactor{contract: contract}, nil
}

// NewOracleProviderSampleFilterer creates a new log filterer instance of OracleProviderSample, bound to a specific deployed contract.
func NewOracleProviderSampleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleProviderSampleFilterer, error) {
	contract, err := bindOracleProviderSample(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleProviderSampleFilterer{contract: contract}, nil
}

// bindOracleProviderSample binds a generic wrapper to an already deployed contract.
func bindOracleProviderSample(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleProviderSampleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleProviderSample *OracleProviderSampleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OracleProviderSample.Contract.OracleProviderSampleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleProviderSample *OracleProviderSampleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleProviderSample.Contract.OracleProviderSampleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleProviderSample *OracleProviderSampleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleProviderSample.Contract.OracleProviderSampleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleProviderSample *OracleProviderSampleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OracleProviderSample.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleProviderSample *OracleProviderSampleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleProviderSample.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleProviderSample *OracleProviderSampleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleProviderSample.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_OracleProviderSample *OracleProviderSampleCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _OracleProviderSample.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_OracleProviderSample *OracleProviderSampleSession) VERSION() (string, error) {
	return _OracleProviderSample.Contract.VERSION(&_OracleProviderSample.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_OracleProviderSample *OracleProviderSampleCallerSession) VERSION() (string, error) {
	return _OracleProviderSample.Contract.VERSION(&_OracleProviderSample.CallOpts)
}

// AdapterId is a free data retrieval call binding the contract method 0x5fb86b01.
//
// Solidity: function adapterId() view returns(bytes32)
func (_OracleProviderSample *OracleProviderSampleCaller) AdapterId(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OracleProviderSample.contract.Call(opts, out, "adapterId")
	return *ret0, err
}

// AdapterId is a free data retrieval call binding the contract method 0x5fb86b01.
//
// Solidity: function adapterId() view returns(bytes32)
func (_OracleProviderSample *OracleProviderSampleSession) AdapterId() ([32]byte, error) {
	return _OracleProviderSample.Contract.AdapterId(&_OracleProviderSample.CallOpts)
}

// AdapterId is a free data retrieval call binding the contract method 0x5fb86b01.
//
// Solidity: function adapterId() view returns(bytes32)
func (_OracleProviderSample *OracleProviderSampleCallerSession) AdapterId() ([32]byte, error) {
	return _OracleProviderSample.Contract.AdapterId(&_OracleProviderSample.CallOpts)
}

// FulfilledCount is a free data retrieval call binding the contract method 0xc5c28fd1.
//
// Solidity: function fulfilledCount() view returns(uint256)
func (_OracleProviderSample *OracleProviderSampleCaller) FulfilledCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OracleProviderSample.contract.Call(opts, out, "fulfilledCount")
	return *ret0, err
}

// FulfilledCount is a free data retrieval call binding the contract method 0xc5c28fd1.
//
// Solidity: function fulfilledCount() view returns(uint256)
func (_OracleProviderSample *OracleProviderSampleSession) FulfilledCount() (*big.Int, error) {
	return _OracleProviderSample.Contract.FulfilledCount(&_OracleProviderSample.CallOpts)
}

// FulfilledCount is a free data retrieval call binding the contract method 0xc5c28fd1.
//
// Solidity: function fulfilledCount() view returns(uint256)
func (_OracleProviderSample *OracleProviderSampleCallerSession) FulfilledCount() (*big.Int, error) {
	return _OracleProviderSample.Contract.FulfilledCount(&_OracleProviderSample.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(bytes32 answer, uint256 roundTime, uint256 timestamp)
func (_OracleProviderSample *OracleProviderSampleCaller) LatestRound(opts *bind.CallOpts) (struct {
	Answer    [32]byte
	RoundTime *big.Int
	Timestamp *big.Int
}, error) {
	ret := new(struct {
		Answer    [32]byte
		RoundTime *big.Int
		Timestamp *big.Int
	})
	out := ret
	err := _OracleProviderSample.contract.Call(opts, out, "latestRound")
	return *ret, err
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(bytes32 answer, uint256 roundTime, uint256 timestamp)
func (_OracleProviderSample *OracleProviderSampleSession) LatestRound() (struct {
	Answer    [32]byte
	RoundTime *big.Int
	Timestamp *big.Int
}, error) {
	return _OracleProviderSample.Contract.LatestRound(&_OracleProviderSample.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(bytes32 answer, uint256 roundTime, uint256 timestamp)
func (_OracleProviderSample *OracleProviderSampleCallerSession) LatestRound() (struct {
	Answer    [32]byte
	RoundTime *big.Int
	Timestamp *big.Int
}, error) {
	return _OracleProviderSample.Contract.LatestRound(&_OracleProviderSample.CallOpts)
}

// NodeAddress is a free data retrieval call binding the contract method 0x2aa8481f.
//
// Solidity: function nodeAddress() view returns(address)
func (_OracleProviderSample *OracleProviderSampleCaller) NodeAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OracleProviderSample.contract.Call(opts, out, "nodeAddress")
	return *ret0, err
}

// NodeAddress is a free data retrieval call binding the contract method 0x2aa8481f.
//
// Solidity: function nodeAddress() view returns(address)
func (_OracleProviderSample *OracleProviderSampleSession) NodeAddress() (common.Address, error) {
	return _OracleProviderSample.Contract.NodeAddress(&_OracleProviderSample.CallOpts)
}

// NodeAddress is a free data retrieval call binding the contract method 0x2aa8481f.
//
// Solidity: function nodeAddress() view returns(address)
func (_OracleProviderSample *OracleProviderSampleCallerSession) NodeAddress() (common.Address, error) {
	return _OracleProviderSample.Contract.NodeAddress(&_OracleProviderSample.CallOpts)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(bytes32 requestId, address nodeAddress, bytes32 adapterId, bytes4 callbackFunctionId, address callBackContract, bytes32 data, uint256 timestamp)
func (_OracleProviderSample *OracleProviderSampleCaller) Requests(opts *bind.CallOpts, arg0 [32]byte) (struct {
	RequestId          [32]byte
	NodeAddress        common.Address
	AdapterId          [32]byte
	CallbackFunctionId [4]byte
	CallBackContract   common.Address
	Data               [32]byte
	Timestamp          *big.Int
}, error) {
	ret := new(struct {
		RequestId          [32]byte
		NodeAddress        common.Address
		AdapterId          [32]byte
		CallbackFunctionId [4]byte
		CallBackContract   common.Address
		Data               [32]byte
		Timestamp          *big.Int
	})
	out := ret
	err := _OracleProviderSample.contract.Call(opts, out, "requests", arg0)
	return *ret, err
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(bytes32 requestId, address nodeAddress, bytes32 adapterId, bytes4 callbackFunctionId, address callBackContract, bytes32 data, uint256 timestamp)
func (_OracleProviderSample *OracleProviderSampleSession) Requests(arg0 [32]byte) (struct {
	RequestId          [32]byte
	NodeAddress        common.Address
	AdapterId          [32]byte
	CallbackFunctionId [4]byte
	CallBackContract   common.Address
	Data               [32]byte
	Timestamp          *big.Int
}, error) {
	return _OracleProviderSample.Contract.Requests(&_OracleProviderSample.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) view returns(bytes32 requestId, address nodeAddress, bytes32 adapterId, bytes4 callbackFunctionId, address callBackContract, bytes32 data, uint256 timestamp)
func (_OracleProviderSample *OracleProviderSampleCallerSession) Requests(arg0 [32]byte) (struct {
	RequestId          [32]byte
	NodeAddress        common.Address
	AdapterId          [32]byte
	CallbackFunctionId [4]byte
	CallBackContract   common.Address
	Data               [32]byte
	Timestamp          *big.Int
}, error) {
	return _OracleProviderSample.Contract.Requests(&_OracleProviderSample.CallOpts, arg0)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(bytes32 answer, uint256 roundTime, uint256 timestamp)
func (_OracleProviderSample *OracleProviderSampleCaller) Rounds(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Answer    [32]byte
	RoundTime *big.Int
	Timestamp *big.Int
}, error) {
	ret := new(struct {
		Answer    [32]byte
		RoundTime *big.Int
		Timestamp *big.Int
	})
	out := ret
	err := _OracleProviderSample.contract.Call(opts, out, "rounds", arg0)
	return *ret, err
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(bytes32 answer, uint256 roundTime, uint256 timestamp)
func (_OracleProviderSample *OracleProviderSampleSession) Rounds(arg0 *big.Int) (struct {
	Answer    [32]byte
	RoundTime *big.Int
	Timestamp *big.Int
}, error) {
	return _OracleProviderSample.Contract.Rounds(&_OracleProviderSample.CallOpts, arg0)
}

// Rounds is a free data retrieval call binding the contract method 0x8c65c81f.
//
// Solidity: function rounds(uint256 ) view returns(bytes32 answer, uint256 roundTime, uint256 timestamp)
func (_OracleProviderSample *OracleProviderSampleCallerSession) Rounds(arg0 *big.Int) (struct {
	Answer    [32]byte
	RoundTime *big.Int
	Timestamp *big.Int
}, error) {
	return _OracleProviderSample.Contract.Rounds(&_OracleProviderSample.CallOpts, arg0)
}

// NewOracleRequest is a paid mutator transaction binding the contract method 0xec2b83f1.
//
// Solidity: function newOracleRequest(bytes4 callbackFunctionId, address callBackContract) returns(bool)
func (_OracleProviderSample *OracleProviderSampleTransactor) NewOracleRequest(opts *bind.TransactOpts, callbackFunctionId [4]byte, callBackContract common.Address) (*types.Transaction, error) {
	return _OracleProviderSample.contract.Transact(opts, "newOracleRequest", callbackFunctionId, callBackContract)
}

// NewOracleRequest is a paid mutator transaction binding the contract method 0xec2b83f1.
//
// Solidity: function newOracleRequest(bytes4 callbackFunctionId, address callBackContract) returns(bool)
func (_OracleProviderSample *OracleProviderSampleSession) NewOracleRequest(callbackFunctionId [4]byte, callBackContract common.Address) (*types.Transaction, error) {
	return _OracleProviderSample.Contract.NewOracleRequest(&_OracleProviderSample.TransactOpts, callbackFunctionId, callBackContract)
}

// NewOracleRequest is a paid mutator transaction binding the contract method 0xec2b83f1.
//
// Solidity: function newOracleRequest(bytes4 callbackFunctionId, address callBackContract) returns(bool)
func (_OracleProviderSample *OracleProviderSampleTransactorSession) NewOracleRequest(callbackFunctionId [4]byte, callBackContract common.Address) (*types.Transaction, error) {
	return _OracleProviderSample.Contract.NewOracleRequest(&_OracleProviderSample.TransactOpts, callbackFunctionId, callBackContract)
}

// NewRoundData is a paid mutator transaction binding the contract method 0x77f8632f.
//
// Solidity: function newRoundData(uint256 roundTime, bytes32 data, bytes32 dataHash, bytes signature) returns(bool)
func (_OracleProviderSample *OracleProviderSampleTransactor) NewRoundData(opts *bind.TransactOpts, roundTime *big.Int, data [32]byte, dataHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _OracleProviderSample.contract.Transact(opts, "newRoundData", roundTime, data, dataHash, signature)
}

// NewRoundData is a paid mutator transaction binding the contract method 0x77f8632f.
//
// Solidity: function newRoundData(uint256 roundTime, bytes32 data, bytes32 dataHash, bytes signature) returns(bool)
func (_OracleProviderSample *OracleProviderSampleSession) NewRoundData(roundTime *big.Int, data [32]byte, dataHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _OracleProviderSample.Contract.NewRoundData(&_OracleProviderSample.TransactOpts, roundTime, data, dataHash, signature)
}

// NewRoundData is a paid mutator transaction binding the contract method 0x77f8632f.
//
// Solidity: function newRoundData(uint256 roundTime, bytes32 data, bytes32 dataHash, bytes signature) returns(bool)
func (_OracleProviderSample *OracleProviderSampleTransactorSession) NewRoundData(roundTime *big.Int, data [32]byte, dataHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _OracleProviderSample.Contract.NewRoundData(&_OracleProviderSample.TransactOpts, roundTime, data, dataHash, signature)
}

// OracleProviderSampleNewOracleRequestIterator is returned from FilterNewOracleRequest and is used to iterate over the raw logs and unpacked data for NewOracleRequest events raised by the OracleProviderSample contract.
type OracleProviderSampleNewOracleRequestIterator struct {
	Event *OracleProviderSampleNewOracleRequest // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OracleProviderSampleNewOracleRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProviderSampleNewOracleRequest)
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
		it.Event = new(OracleProviderSampleNewOracleRequest)
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
func (it *OracleProviderSampleNewOracleRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProviderSampleNewOracleRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProviderSampleNewOracleRequest represents a NewOracleRequest event raised by the OracleProviderSample contract.
type OracleProviderSampleNewOracleRequest struct {
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewOracleRequest is a free log retrieval operation binding the contract event 0x7a6060117db4511b5ffd88545b9b43962c0ae441d1f07d05c86b869490d88369.
//
// Solidity: event NewOracleRequest(bytes32 requestId)
func (_OracleProviderSample *OracleProviderSampleFilterer) FilterNewOracleRequest(opts *bind.FilterOpts) (*OracleProviderSampleNewOracleRequestIterator, error) {

	logs, sub, err := _OracleProviderSample.contract.FilterLogs(opts, "NewOracleRequest")
	if err != nil {
		return nil, err
	}
	return &OracleProviderSampleNewOracleRequestIterator{contract: _OracleProviderSample.contract, event: "NewOracleRequest", logs: logs, sub: sub}, nil
}

// WatchNewOracleRequest is a free log subscription operation binding the contract event 0x7a6060117db4511b5ffd88545b9b43962c0ae441d1f07d05c86b869490d88369.
//
// Solidity: event NewOracleRequest(bytes32 requestId)
func (_OracleProviderSample *OracleProviderSampleFilterer) WatchNewOracleRequest(opts *bind.WatchOpts, sink chan<- *OracleProviderSampleNewOracleRequest) (event.Subscription, error) {

	logs, sub, err := _OracleProviderSample.contract.WatchLogs(opts, "NewOracleRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProviderSampleNewOracleRequest)
				if err := _OracleProviderSample.contract.UnpackLog(event, "NewOracleRequest", log); err != nil {
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
func (_OracleProviderSample *OracleProviderSampleFilterer) ParseNewOracleRequest(log types.Log) (*OracleProviderSampleNewOracleRequest, error) {
	event := new(OracleProviderSampleNewOracleRequest)
	if err := _OracleProviderSample.contract.UnpackLog(event, "NewOracleRequest", log); err != nil {
		return nil, err
	}
	return event, nil
}
