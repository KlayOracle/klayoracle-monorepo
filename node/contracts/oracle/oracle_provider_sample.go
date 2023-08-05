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

// OracleProviderSampleMetaData contains all meta data concerning the OracleProviderSample contract.
var OracleProviderSampleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_adapterId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"NewOracleRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hashedMessage\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"VerifyKlaytnMessage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"roundAnswer\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"roundTime\",\"type\":\"uint256\"}],\"name\":\"_getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adapterId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fulfilledCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"answer\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"roundTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"callBackContract\",\"type\":\"address\"}],\"name\":\"newOracleRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"roundAnswer\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"newRoundData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"adapterId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"callBackContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rounds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"answer\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"roundTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"splitSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60c06040525f60015534801562000014575f80fd5b5060405162001b2538038062001b2583398181016040528101906200003a91906200011f565b81818173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508060a081815250505050505062000164565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620000b18262000086565b9050919050565b620000c381620000a5565b8114620000ce575f80fd5b50565b5f81519050620000e181620000b8565b92915050565b5f819050919050565b620000fb81620000e7565b811462000106575f80fd5b50565b5f815190506200011981620000f0565b92915050565b5f806040838503121562000138576200013762000082565b5b5f6200014785828601620000d1565b92505060206200015a8582860162000109565b9150509250929050565b60805160a051611975620001b05f395f81816104b80152818161079f015261084201525f81816102ab01528181610367015281816104940152818161077e015261080601526119755ff3fe608060405234801561000f575f80fd5b50600436106100b2575f3560e01c80639d8669851161006f5780639d866985146101a4578063a7bb5803146101da578063be2bb8cd1461020c578063c5c28fd11461023c578063ec2b83f11461025a578063ffa1ad741461028a576100b2565b806309c980c1146100b65780632aa8481f146100e65780635fb86b0114610104578063668a0f021461012257806377d3bfb9146101425780638c65c81f14610172575b5f80fd5b6100d060048036038101906100cb9190610de7565b6102a8565b6040516100dd9190610e6d565b60405180910390f35b6100ee610492565b6040516100fb9190610ec5565b60405180910390f35b61010c6104b6565b6040516101199190610eed565b60405180910390f35b61012a6104da565b60405161013993929190610f15565b60405180910390f35b61015c60048036038101906101579190610f4a565b6104f1565b6040516101699190610eed565b60405180910390f35b61018c60048036038101906101879190610f88565b610525565b60405161019b93929190610f15565b60405180910390f35b6101be60048036038101906101b99190610fb3565b61055a565b6040516101d19796959493929190611018565b60405180910390f35b6101f460048036038101906101ef9190611085565b6105e2565b604051610203939291906110e7565b60405180910390f35b61022660048036038101906102219190611146565b610666565b6040516102339190610ec5565b60405180910390f35b610244610729565b60405161025191906111aa565b60405180910390f35b610274600480360381019061026f9190611217565b61072f565b6040516102819190610e6d565b60405180910390f35b610292610948565b60405161029f91906112cf565b60405180910390f35b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610337576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032e90611339565b60405180910390fd5b5f61034284866104f1565b90505f805f610350866105e2565b9250925092505f61036385838686610666565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146103f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ea906113a1565b60405180910390fd5b5f60405180606001604052808a81526020018b8152602001428152509050600481908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f0155602082015181600101556040820151816002015550508060055f820151815f0155602082015181600101556040820151816002015590505088600281905550600196505050505050509392505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b6005805f0154908060010154908060020154905083565b5f33838360405160200161050793929190611444565b60405160208183030381529060405280519060200120905092915050565b60048181548110610534575f80fd5b905f5260205f2090600302015f91509050805f0154908060010154908060020154905083565b5f602052805f5260405f205f91509050805f015490806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806002015490806003015f9054906101000a900460e01b908060030160049054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154908060050154905087565b5f805f6041845114610629576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610620906114ca565b60405180910390fd5b602084015192506040840151915060608401515f1a90505f8160ff160361064f57601b90505b60018160ff160361065f57601c90505b9193909250565b5f806040518060400160405280601a81526020017f194b6c6179746e205369676e6564204d6573736167653a0a333200000000000081525090505f81876040516020016106b492919061152c565b6040516020818303038152906040528051906020012090505f6001828888886040515f81526020016040526040516106ef9493929190611553565b6020604051602081039080840390855afa15801561070f573d5f803e3d5ffd5b505050602060405103519050809350505050949350505050565b60015481565b5f61073933610981565b610778576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076f90611606565b60405180910390fd5b5f6003547f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008686600254426040516020016107da9796959493929190611644565b6040516020818303038152906040528051906020012090505f6040518060e001604052808381526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000008152602001867bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020018573ffffffffffffffffffffffffffffffffffffffff16815260200160055f015481526020014281525090507f7a6060117db4511b5ffd88545b9b43962c0ae441d1f07d05c86b869490d88369826040516108e69190610eed565b60405180910390a15f6108f983836109f9565b90508061093b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093290611734565b60405180910390fd5b6001935050505092915050565b6040518060400160405280600681526020017f76312e302e30000000000000000000000000000000000000000000000000000081525081565b5f8073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036109f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109e79061179c565b60405180910390fd5b60019050919050565b5f610a0382610be2565b610a42576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3990611804565b60405180910390fd5b815f808581526020019081526020015f205f820151815f01556020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201556060820151816003015f6101000a81548163ffffffff021916908360e01c021790555060808201518160030160046101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a0820151816004015560c0820151816005015590505060015f815480929190610b3e9061184f565b91905055505f826080015173ffffffffffffffffffffffffffffffffffffffff1683606001518460a00151604051602001610b7a929190611896565b604051602081830303815290604052604051610b9691906118c1565b5f604051808303815f865af19150503d805f8114610bcf576040519150601f19603f3d011682016040523d82523d5f602084013e610bd4565b606091505b505090508091505092915050565b5f805f1b8260a0015103610c2b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c2290611921565b60405180910390fd5b60019050919050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b610c5781610c45565b8114610c61575f80fd5b50565b5f81359050610c7281610c4e565b92915050565b5f819050919050565b610c8a81610c78565b8114610c94575f80fd5b50565b5f81359050610ca581610c81565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610cf982610cb3565b810181811067ffffffffffffffff82111715610d1857610d17610cc3565b5b80604052505050565b5f610d2a610c34565b9050610d368282610cf0565b919050565b5f67ffffffffffffffff821115610d5557610d54610cc3565b5b610d5e82610cb3565b9050602081019050919050565b828183375f83830152505050565b5f610d8b610d8684610d3b565b610d21565b905082815260208101848484011115610da757610da6610caf565b5b610db2848285610d6b565b509392505050565b5f82601f830112610dce57610dcd610cab565b5b8135610dde848260208601610d79565b91505092915050565b5f805f60608486031215610dfe57610dfd610c3d565b5b5f610e0b86828701610c64565b9350506020610e1c86828701610c97565b925050604084013567ffffffffffffffff811115610e3d57610e3c610c41565b5b610e4986828701610dba565b9150509250925092565b5f8115159050919050565b610e6781610e53565b82525050565b5f602082019050610e805f830184610e5e565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610eaf82610e86565b9050919050565b610ebf81610ea5565b82525050565b5f602082019050610ed85f830184610eb6565b92915050565b610ee781610c78565b82525050565b5f602082019050610f005f830184610ede565b92915050565b610f0f81610c45565b82525050565b5f606082019050610f285f830186610ede565b610f356020830185610f06565b610f426040830184610f06565b949350505050565b5f8060408385031215610f6057610f5f610c3d565b5b5f610f6d85828601610c97565b9250506020610f7e85828601610c64565b9150509250929050565b5f60208284031215610f9d57610f9c610c3d565b5b5f610faa84828501610c64565b91505092915050565b5f60208284031215610fc857610fc7610c3d565b5b5f610fd584828501610c97565b91505092915050565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61101281610fde565b82525050565b5f60e08201905061102b5f83018a610ede565b6110386020830189610eb6565b6110456040830188610ede565b6110526060830187611009565b61105f6080830186610eb6565b61106c60a0830185610ede565b61107960c0830184610f06565b98975050505050505050565b5f6020828403121561109a57611099610c3d565b5b5f82013567ffffffffffffffff8111156110b7576110b6610c41565b5b6110c384828501610dba565b91505092915050565b5f60ff82169050919050565b6110e1816110cc565b82525050565b5f6060820190506110fa5f830186610ede565b6111076020830185610ede565b61111460408301846110d8565b949350505050565b611125816110cc565b811461112f575f80fd5b50565b5f813590506111408161111c565b92915050565b5f805f806080858703121561115e5761115d610c3d565b5b5f61116b87828801610c97565b945050602061117c87828801611132565b935050604061118d87828801610c97565b925050606061119e87828801610c97565b91505092959194509250565b5f6020820190506111bd5f830184610f06565b92915050565b6111cc81610fde565b81146111d6575f80fd5b50565b5f813590506111e7816111c3565b92915050565b6111f681610ea5565b8114611200575f80fd5b50565b5f81359050611211816111ed565b92915050565b5f806040838503121561122d5761122c610c3d565b5b5f61123a858286016111d9565b925050602061124b85828601611203565b9150509250929050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561128c578082015181840152602081019050611271565b5f8484015250505050565b5f6112a182611255565b6112ab818561125f565b93506112bb81856020860161126f565b6112c481610cb3565b840191505092915050565b5f6020820190508181035f8301526112e78184611297565b905092915050565b7f4f7261636c653a206e6f646520756e6b6e6f776e0000000000000000000000005f82015250565b5f61132360148361125f565b915061132e826112ef565b602082019050919050565b5f6020820190508181035f83015261135081611317565b9050919050565b7f4f7261636c653a20496e76616c696420646174610000000000000000000000005f82015250565b5f61138b60148361125f565b915061139682611357565b602082019050919050565b5f6020820190508181035f8301526113b88161137f565b9050919050565b5f8160601b9050919050565b5f6113d5826113bf565b9050919050565b5f6113e6826113cb565b9050919050565b6113fe6113f982610ea5565b6113dc565b82525050565b5f819050919050565b61141e61141982610c78565b611404565b82525050565b5f819050919050565b61143e61143982610c45565b611424565b82525050565b5f61144f82866113ed565b60148201915061145f828561140d565b60208201915061146f828461142d565b602082019150819050949350505050565b7f696e76616c6964207369676e6174757265206c656e67746800000000000000005f82015250565b5f6114b460188361125f565b91506114bf82611480565b602082019050919050565b5f6020820190508181035f8301526114e1816114a8565b9050919050565b5f81519050919050565b5f81905092915050565b5f611506826114e8565b61151081856114f2565b935061152081856020860161126f565b80840191505092915050565b5f61153782856114fc565b9150611543828461140d565b6020820191508190509392505050565b5f6080820190506115665f830187610ede565b61157360208301866110d8565b6115806040830185610ede565b61158d6060830184610ede565b95945050505050565b7f4f7261636c653a20636f6e73756d6572206973206e6f742077686974656c69735f8201527f7465640000000000000000000000000000000000000000000000000000000000602082015250565b5f6115f060238361125f565b91506115fb82611596565b604082019050919050565b5f6020820190508181035f83015261161d816115e4565b9050919050565b5f819050919050565b61163e61163982610fde565b611624565b82525050565b5f61164f828a61140d565b60208201915061165f82896113ed565b60148201915061166f828861140d565b60208201915061167f828761162d565b60048201915061168f82866113ed565b60148201915061169f828561140d565b6020820191506116af828461142d565b60208201915081905098975050505050505050565b7f4b6c61794f7261636c653a2072657175657374206e6f742066756c66696c6c655f8201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b5f61171e60218361125f565b9150611729826116c4565b604082019050919050565b5f6020820190508181035f83015261174b81611712565b9050919050565b7f4b6c61794f7261636c653a20636f6e73756d657220697320696e76616c6964005f82015250565b5f611786601f8361125f565b915061179182611752565b602082019050919050565b5f6020820190508181035f8301526117b38161177a565b9050919050565b7f4b6c61794f7261636c653a2073756273637269626520746f20445000000000005f82015250565b5f6117ee601b8361125f565b91506117f9826117ba565b602082019050919050565b5f6020820190508181035f83015261181b816117e2565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61185982610c45565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361188b5761188a611822565b5b600182019050919050565b5f6118a1828561162d565b6004820191506118b1828461140d565b6020820191508190509392505050565b5f6118cc82846114fc565b915081905092915050565b7f4b6c61794f7261636c653a206461746120697320656d707479000000000000005f82015250565b5f61190b60198361125f565b9150611916826118d7565b602082019050919050565b5f6020820190508181035f830152611938816118ff565b905091905056fea2646970667358221220dbd761414e957d7ece233585475b35b637984535ab7b6cef173bc866f4989f2764736f6c63430008140033",
}

// OracleProviderSampleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleProviderSampleMetaData.ABI instead.
var OracleProviderSampleABI = OracleProviderSampleMetaData.ABI

// OracleProviderSampleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OracleProviderSampleMetaData.Bin instead.
var OracleProviderSampleBin = OracleProviderSampleMetaData.Bin

// DeployOracleProviderSample deploys a new Ethereum contract, binding an instance of OracleProviderSample to it.
func DeployOracleProviderSample(auth *bind.TransactOpts, backend bind.ContractBackend, _nodeAddress common.Address, _adapterId [32]byte) (common.Address, *types.Transaction, *OracleProviderSample, error) {
	parsed, err := OracleProviderSampleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OracleProviderSampleBin), backend, _nodeAddress, _adapterId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OracleProviderSample{OracleProviderSampleCaller: OracleProviderSampleCaller{contract: contract}, OracleProviderSampleTransactor: OracleProviderSampleTransactor{contract: contract}, OracleProviderSampleFilterer: OracleProviderSampleFilterer{contract: contract}}, nil
}

// OracleProviderSample is an auto generated Go binding around an Ethereum contract.
type OracleProviderSample struct {
	OracleProviderSampleCaller     // Read-only binding to the contract
	OracleProviderSampleTransactor // Write-only binding to the contract
	OracleProviderSampleFilterer   // Log filterer for contract events
}

// OracleProviderSampleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleProviderSampleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleProviderSampleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleProviderSampleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleProviderSampleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleProviderSampleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleProviderSampleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleProviderSampleSession struct {
	Contract     *OracleProviderSample // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OracleProviderSampleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleProviderSampleCallerSession struct {
	Contract *OracleProviderSampleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// OracleProviderSampleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleProviderSampleTransactorSession struct {
	Contract     *OracleProviderSampleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// OracleProviderSampleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleProviderSampleRaw struct {
	Contract *OracleProviderSample // Generic contract binding to access the raw methods on
}

// OracleProviderSampleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleProviderSampleCallerRaw struct {
	Contract *OracleProviderSampleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleProviderSampleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
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
	parsed, err := OracleProviderSampleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleProviderSample *OracleProviderSampleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_OracleProviderSample *OracleProviderSampleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
	var out []interface{}
	err := _OracleProviderSample.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

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

// VerifyKlaytnMessage is a free data retrieval call binding the contract method 0xbe2bb8cd.
//
// Solidity: function VerifyKlaytnMessage(bytes32 _hashedMessage, uint8 _v, bytes32 _r, bytes32 _s) pure returns(address)
func (_OracleProviderSample *OracleProviderSampleCaller) VerifyKlaytnMessage(opts *bind.CallOpts, _hashedMessage [32]byte, _v uint8, _r [32]byte, _s [32]byte) (common.Address, error) {
	var out []interface{}
	err := _OracleProviderSample.contract.Call(opts, &out, "VerifyKlaytnMessage", _hashedMessage, _v, _r, _s)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifyKlaytnMessage is a free data retrieval call binding the contract method 0xbe2bb8cd.
//
// Solidity: function VerifyKlaytnMessage(bytes32 _hashedMessage, uint8 _v, bytes32 _r, bytes32 _s) pure returns(address)
func (_OracleProviderSample *OracleProviderSampleSession) VerifyKlaytnMessage(_hashedMessage [32]byte, _v uint8, _r [32]byte, _s [32]byte) (common.Address, error) {
	return _OracleProviderSample.Contract.VerifyKlaytnMessage(&_OracleProviderSample.CallOpts, _hashedMessage, _v, _r, _s)
}

// VerifyKlaytnMessage is a free data retrieval call binding the contract method 0xbe2bb8cd.
//
// Solidity: function VerifyKlaytnMessage(bytes32 _hashedMessage, uint8 _v, bytes32 _r, bytes32 _s) pure returns(address)
func (_OracleProviderSample *OracleProviderSampleCallerSession) VerifyKlaytnMessage(_hashedMessage [32]byte, _v uint8, _r [32]byte, _s [32]byte) (common.Address, error) {
	return _OracleProviderSample.Contract.VerifyKlaytnMessage(&_OracleProviderSample.CallOpts, _hashedMessage, _v, _r, _s)
}

// GetHash is a free data retrieval call binding the contract method 0x77d3bfb9.
//
// Solidity: function _getHash(bytes32 roundAnswer, uint256 roundTime) view returns(bytes32)
func (_OracleProviderSample *OracleProviderSampleCaller) GetHash(opts *bind.CallOpts, roundAnswer [32]byte, roundTime *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _OracleProviderSample.contract.Call(opts, &out, "_getHash", roundAnswer, roundTime)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHash is a free data retrieval call binding the contract method 0x77d3bfb9.
//
// Solidity: function _getHash(bytes32 roundAnswer, uint256 roundTime) view returns(bytes32)
func (_OracleProviderSample *OracleProviderSampleSession) GetHash(roundAnswer [32]byte, roundTime *big.Int) ([32]byte, error) {
	return _OracleProviderSample.Contract.GetHash(&_OracleProviderSample.CallOpts, roundAnswer, roundTime)
}

// GetHash is a free data retrieval call binding the contract method 0x77d3bfb9.
//
// Solidity: function _getHash(bytes32 roundAnswer, uint256 roundTime) view returns(bytes32)
func (_OracleProviderSample *OracleProviderSampleCallerSession) GetHash(roundAnswer [32]byte, roundTime *big.Int) ([32]byte, error) {
	return _OracleProviderSample.Contract.GetHash(&_OracleProviderSample.CallOpts, roundAnswer, roundTime)
}

// AdapterId is a free data retrieval call binding the contract method 0x5fb86b01.
//
// Solidity: function adapterId() view returns(bytes32)
func (_OracleProviderSample *OracleProviderSampleCaller) AdapterId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OracleProviderSample.contract.Call(opts, &out, "adapterId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

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
	var out []interface{}
	err := _OracleProviderSample.contract.Call(opts, &out, "fulfilledCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

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
	var out []interface{}
	err := _OracleProviderSample.contract.Call(opts, &out, "latestRound")

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
	var out []interface{}
	err := _OracleProviderSample.contract.Call(opts, &out, "nodeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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
	var out []interface{}
	err := _OracleProviderSample.contract.Call(opts, &out, "requests", arg0)

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
	var out []interface{}
	err := _OracleProviderSample.contract.Call(opts, &out, "rounds", arg0)

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

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_OracleProviderSample *OracleProviderSampleCaller) SplitSignature(opts *bind.CallOpts, sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	var out []interface{}
	err := _OracleProviderSample.contract.Call(opts, &out, "splitSignature", sig)

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
func (_OracleProviderSample *OracleProviderSampleSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _OracleProviderSample.Contract.SplitSignature(&_OracleProviderSample.CallOpts, sig)
}

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_OracleProviderSample *OracleProviderSampleCallerSession) SplitSignature(sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	return _OracleProviderSample.Contract.SplitSignature(&_OracleProviderSample.CallOpts, sig)
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

// NewRoundData is a paid mutator transaction binding the contract method 0x09c980c1.
//
// Solidity: function newRoundData(uint256 roundTime, bytes32 roundAnswer, bytes signature) returns(bool)
func (_OracleProviderSample *OracleProviderSampleTransactor) NewRoundData(opts *bind.TransactOpts, roundTime *big.Int, roundAnswer [32]byte, signature []byte) (*types.Transaction, error) {
	return _OracleProviderSample.contract.Transact(opts, "newRoundData", roundTime, roundAnswer, signature)
}

// NewRoundData is a paid mutator transaction binding the contract method 0x09c980c1.
//
// Solidity: function newRoundData(uint256 roundTime, bytes32 roundAnswer, bytes signature) returns(bool)
func (_OracleProviderSample *OracleProviderSampleSession) NewRoundData(roundTime *big.Int, roundAnswer [32]byte, signature []byte) (*types.Transaction, error) {
	return _OracleProviderSample.Contract.NewRoundData(&_OracleProviderSample.TransactOpts, roundTime, roundAnswer, signature)
}

// NewRoundData is a paid mutator transaction binding the contract method 0x09c980c1.
//
// Solidity: function newRoundData(uint256 roundTime, bytes32 roundAnswer, bytes signature) returns(bool)
func (_OracleProviderSample *OracleProviderSampleTransactorSession) NewRoundData(roundTime *big.Int, roundAnswer [32]byte, signature []byte) (*types.Transaction, error) {
	return _OracleProviderSample.Contract.NewRoundData(&_OracleProviderSample.TransactOpts, roundTime, roundAnswer, signature)
}

// OracleProviderSampleNewOracleRequestIterator is returned from FilterNewOracleRequest and is used to iterate over the raw logs and unpacked data for NewOracleRequest events raised by the OracleProviderSample contract.
type OracleProviderSampleNewOracleRequestIterator struct {
	Event *OracleProviderSampleNewOracleRequest // Event containing the contract specifics and raw log

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
	event.Raw = log
	return event, nil
}
