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
const OracleProviderSampleABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_adapterId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"NewOracleRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hashedMessage\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"VerifyKlaytnMessage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"roundAnswer\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"roundTime\",\"type\":\"uint256\"}],\"name\":\"_getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adapterId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fulfilledCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"answer\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"roundTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"callBackContract\",\"type\":\"address\"}],\"name\":\"newOracleRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"roundTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"roundAnswer\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"newRoundData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"nodeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"adapterId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes4\",\"name\":\"callbackFunctionId\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"callBackContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rounds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"answer\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"roundTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"splitSignature\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// OracleProviderSampleBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const OracleProviderSampleBinRuntime = ``

// OracleProviderSampleBin is the compiled bytecode used for deploying new contracts.
var OracleProviderSampleBin = "0x60c060405260006001553480156200001657600080fd5b5060405162001be838038062001be883398181016040528101906200003c919062000129565b81818173ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508060a081815250505050505062000170565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000b68262000089565b9050919050565b620000c881620000a9565b8114620000d457600080fd5b50565b600081519050620000e881620000bd565b92915050565b6000819050919050565b6200010381620000ee565b81146200010f57600080fd5b50565b6000815190506200012381620000f8565b92915050565b6000806040838503121562000143576200014262000084565b5b60006200015385828601620000d7565b9250506020620001668582860162000112565b9150509250929050565b60805160a051611a28620001c0600039600081816104c8015281816107c9015261086d0152600081816102af0152818161036f015281816104a4015281816107a801526108310152611a286000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80639d866985116100715780639d866985146101a7578063a7bb5803146101dd578063be2bb8cd1461020f578063c5c28fd11461023f578063ec2b83f11461025d578063ffa1ad741461028d576100b4565b806309c980c1146100b95780632aa8481f146100e95780635fb86b0114610107578063668a0f021461012557806377d3bfb9146101455780638c65c81f14610175575b600080fd5b6100d360048036038101906100ce9190610e37565b6102ab565b6040516100e09190610ec1565b60405180910390f35b6100f16104a2565b6040516100fe9190610f1d565b60405180910390f35b61010f6104c6565b60405161011c9190610f47565b60405180910390f35b61012d6104ea565b60405161013c93929190610f71565b60405180910390f35b61015f600480360381019061015a9190610fa8565b610502565b60405161016c9190610f47565b60405180910390f35b61018f600480360381019061018a9190610fe8565b610537565b60405161019e93929190610f71565b60405180910390f35b6101c160048036038101906101bc9190611015565b610571565b6040516101d4979695949392919061107d565b60405180910390f35b6101f760048036038101906101f291906110ec565b610600565b60405161020693929190611151565b60405180910390f35b610229600480360381019061022491906111b4565b610688565b6040516102369190610f1d565b60405180910390f35b610247610751565b604051610254919061121b565b60405180910390f35b6102776004803603810190610272919061128e565b610757565b6040516102849190610ec1565b60405180910390f35b610295610975565b6040516102a2919061134d565b60405180910390f35b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461033b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610332906113bb565b60405180910390fd5b60006103478486610502565b9050600080600061035786610600565b925092509250600061036b85838686610688565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146103fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f290611427565b60405180910390fd5b600060405180606001604052808a81526020018b81526020014281525090506004819080600181540180825580915050600190039060005260206000209060030201600090919091909150600082015181600001556020820151816001015560408201518160020155505080600560008201518160000155602082015181600101556040820151816002015590505088600281905550600196505050505050509392505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b60058060000154908060010154908060020154905083565b6000338383604051602001610519939291906114d1565b60405160208183030381529060405280519060200120905092915050565b6004818154811061054757600080fd5b90600052602060002090600302016000915090508060000154908060010154908060020154905083565b60006020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154908060030160009054906101000a900460e01b908060030160049054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154908060050154905087565b60008060006041845114610649576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106409061155a565b60405180910390fd5b6020840151925060408401519150606084015160001a905060008160ff160361067157601b90505b60018160ff160361068157601c90505b9193909250565b6000806040518060400160405280601a81526020017f194b6c6179746e205369676e6564204d6573736167653a0a33320000000000008152509050600081876040516020016106d89291906115c1565b60405160208183030381529060405280519060200120905060006001828888886040516000815260200160405260405161071594939291906115e9565b6020604051602081039080840390855afa158015610737573d6000803e3d6000fd5b505050602060405103519050809350505050949350505050565b60015481565b6000610762336109ae565b6107a1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610798906116a0565b60405180910390fd5b60006003547f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000086866002544260405160200161080497969594939291906116e1565b60405160208183030381529060405280519060200120905060006040518060e001604052808381526020017f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000008152602001867bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681526020018573ffffffffffffffffffffffffffffffffffffffff16815260200160056000015481526020014281525090507f7a6060117db4511b5ffd88545b9b43962c0ae441d1f07d05c86b869490d88369826040516109129190610f47565b60405180910390a160006109268383610a27565b905080610968576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095f906117d4565b60405180910390fd5b6001935050505092915050565b6040518060400160405280600681526020017f76312e302e30000000000000000000000000000000000000000000000000000081525081565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610a1e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1590611840565b60405180910390fd5b60019050919050565b6000610a3282610c1d565b610a71576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a68906118ac565b60405180910390fd5b816000808581526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015560608201518160030160006101000a81548163ffffffff021916908360e01c021790555060808201518160030160046101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a0820151816004015560c0820151816005015590505060016000815480929190610b74906118fb565b91905055506000826080015173ffffffffffffffffffffffffffffffffffffffff1683606001518460a00151604051602001610bb1929190611943565b604051602081830303815290604052604051610bcd919061196f565b6000604051808303816000865af19150503d8060008114610c0a576040519150601f19603f3d011682016040523d82523d6000602084013e610c0f565b606091505b505090508091505092915050565b60008060001b8260a0015103610c68576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c5f906119d2565b60405180910390fd5b60019050919050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610c9881610c85565b8114610ca357600080fd5b50565b600081359050610cb581610c8f565b92915050565b6000819050919050565b610cce81610cbb565b8114610cd957600080fd5b50565b600081359050610ceb81610cc5565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610d4482610cfb565b810181811067ffffffffffffffff82111715610d6357610d62610d0c565b5b80604052505050565b6000610d76610c71565b9050610d828282610d3b565b919050565b600067ffffffffffffffff821115610da257610da1610d0c565b5b610dab82610cfb565b9050602081019050919050565b82818337600083830152505050565b6000610dda610dd584610d87565b610d6c565b905082815260208101848484011115610df657610df5610cf6565b5b610e01848285610db8565b509392505050565b600082601f830112610e1e57610e1d610cf1565b5b8135610e2e848260208601610dc7565b91505092915050565b600080600060608486031215610e5057610e4f610c7b565b5b6000610e5e86828701610ca6565b9350506020610e6f86828701610cdc565b925050604084013567ffffffffffffffff811115610e9057610e8f610c80565b5b610e9c86828701610e09565b9150509250925092565b60008115159050919050565b610ebb81610ea6565b82525050565b6000602082019050610ed66000830184610eb2565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610f0782610edc565b9050919050565b610f1781610efc565b82525050565b6000602082019050610f326000830184610f0e565b92915050565b610f4181610cbb565b82525050565b6000602082019050610f5c6000830184610f38565b92915050565b610f6b81610c85565b82525050565b6000606082019050610f866000830186610f38565b610f936020830185610f62565b610fa06040830184610f62565b949350505050565b60008060408385031215610fbf57610fbe610c7b565b5b6000610fcd85828601610cdc565b9250506020610fde85828601610ca6565b9150509250929050565b600060208284031215610ffe57610ffd610c7b565b5b600061100c84828501610ca6565b91505092915050565b60006020828403121561102b5761102a610c7b565b5b600061103984828501610cdc565b91505092915050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61107781611042565b82525050565b600060e082019050611092600083018a610f38565b61109f6020830189610f0e565b6110ac6040830188610f38565b6110b9606083018761106e565b6110c66080830186610f0e565b6110d360a0830185610f38565b6110e060c0830184610f62565b98975050505050505050565b60006020828403121561110257611101610c7b565b5b600082013567ffffffffffffffff8111156111205761111f610c80565b5b61112c84828501610e09565b91505092915050565b600060ff82169050919050565b61114b81611135565b82525050565b60006060820190506111666000830186610f38565b6111736020830185610f38565b6111806040830184611142565b949350505050565b61119181611135565b811461119c57600080fd5b50565b6000813590506111ae81611188565b92915050565b600080600080608085870312156111ce576111cd610c7b565b5b60006111dc87828801610cdc565b94505060206111ed8782880161119f565b93505060406111fe87828801610cdc565b925050606061120f87828801610cdc565b91505092959194509250565b60006020820190506112306000830184610f62565b92915050565b61123f81611042565b811461124a57600080fd5b50565b60008135905061125c81611236565b92915050565b61126b81610efc565b811461127657600080fd5b50565b60008135905061128881611262565b92915050565b600080604083850312156112a5576112a4610c7b565b5b60006112b38582860161124d565b92505060206112c485828601611279565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b838110156113085780820151818401526020810190506112ed565b60008484015250505050565b600061131f826112ce565b61132981856112d9565b93506113398185602086016112ea565b61134281610cfb565b840191505092915050565b600060208201905081810360008301526113678184611314565b905092915050565b7f4f7261636c653a206e6f646520756e6b6e6f776e000000000000000000000000600082015250565b60006113a56014836112d9565b91506113b08261136f565b602082019050919050565b600060208201905081810360008301526113d481611398565b9050919050565b7f4f7261636c653a20496e76616c69642064617461000000000000000000000000600082015250565b60006114116014836112d9565b915061141c826113db565b602082019050919050565b6000602082019050818103600083015261144081611404565b9050919050565b60008160601b9050919050565b600061145f82611447565b9050919050565b600061147182611454565b9050919050565b61148961148482610efc565b611466565b82525050565b6000819050919050565b6114aa6114a582610cbb565b61148f565b82525050565b6000819050919050565b6114cb6114c682610c85565b6114b0565b82525050565b60006114dd8286611478565b6014820191506114ed8285611499565b6020820191506114fd82846114ba565b602082019150819050949350505050565b7f696e76616c6964207369676e6174757265206c656e6774680000000000000000600082015250565b60006115446018836112d9565b915061154f8261150e565b602082019050919050565b6000602082019050818103600083015261157381611537565b9050919050565b600081519050919050565b600081905092915050565b600061159b8261157a565b6115a58185611585565b93506115b58185602086016112ea565b80840191505092915050565b60006115cd8285611590565b91506115d98284611499565b6020820191508190509392505050565b60006080820190506115fe6000830187610f38565b61160b6020830186611142565b6116186040830185610f38565b6116256060830184610f38565b95945050505050565b7f4f7261636c653a20636f6e73756d6572206973206e6f742077686974656c697360008201527f7465640000000000000000000000000000000000000000000000000000000000602082015250565b600061168a6023836112d9565b91506116958261162e565b604082019050919050565b600060208201905081810360008301526116b98161167d565b9050919050565b6000819050919050565b6116db6116d682611042565b6116c0565b82525050565b60006116ed828a611499565b6020820191506116fd8289611478565b60148201915061170d8288611499565b60208201915061171d82876116ca565b60048201915061172d8286611478565b60148201915061173d8285611499565b60208201915061174d82846114ba565b60208201915081905098975050505050505050565b7f4b6c61794f7261636c653a2072657175657374206e6f742066756c66696c6c6560008201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b60006117be6021836112d9565b91506117c982611762565b604082019050919050565b600060208201905081810360008301526117ed816117b1565b9050919050565b7f4b6c61794f7261636c653a20636f6e73756d657220697320696e76616c696400600082015250565b600061182a601f836112d9565b9150611835826117f4565b602082019050919050565b600060208201905081810360008301526118598161181d565b9050919050565b7f4b6c61794f7261636c653a2073756273637269626520746f2044500000000000600082015250565b6000611896601b836112d9565b91506118a182611860565b602082019050919050565b600060208201905081810360008301526118c581611889565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061190682610c85565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611938576119376118cc565b5b600182019050919050565b600061194f82856116ca565b60048201915061195f8284611499565b6020820191508190509392505050565b600061197b8284611590565b915081905092915050565b7f4b6c61794f7261636c653a206461746120697320656d70747900000000000000600082015250565b60006119bc6019836112d9565b91506119c782611986565b602082019050919050565b600060208201905081810360008301526119eb816119af565b905091905056fea26469706673582212206a8ea0ba3282141d5de7ba1603115a5e2d8f70e02f5f9b8d95fee5df79a0449764736f6c63430008110033"

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

// VerifyKlaytnMessage is a free data retrieval call binding the contract method 0xbe2bb8cd.
//
// Solidity: function VerifyKlaytnMessage(bytes32 _hashedMessage, uint8 _v, bytes32 _r, bytes32 _s) pure returns(address)
func (_OracleProviderSample *OracleProviderSampleCaller) VerifyKlaytnMessage(opts *bind.CallOpts, _hashedMessage [32]byte, _v uint8, _r [32]byte, _s [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _OracleProviderSample.contract.Call(opts, out, "VerifyKlaytnMessage", _hashedMessage, _v, _r, _s)
	return *ret0, err
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
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OracleProviderSample.contract.Call(opts, out, "_getHash", roundAnswer, roundTime)
	return *ret0, err
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

// SplitSignature is a free data retrieval call binding the contract method 0xa7bb5803.
//
// Solidity: function splitSignature(bytes sig) pure returns(bytes32 r, bytes32 s, uint8 v)
func (_OracleProviderSample *OracleProviderSampleCaller) SplitSignature(opts *bind.CallOpts, sig []byte) (struct {
	R [32]byte
	S [32]byte
	V uint8
}, error) {
	ret := new(struct {
		R [32]byte
		S [32]byte
		V uint8
	})
	out := ret
	err := _OracleProviderSample.contract.Call(opts, out, "splitSignature", sig)
	return *ret, err
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
