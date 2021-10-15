// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fisco

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/event"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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

// ComputingABI is the input ABI used to generate the binding from.
const ComputingABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"datashare\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initledger\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"paramAddr\",\"type\":\"string\"},{\"name\":\"paramAbstract\",\"type\":\"string\"},{\"name\":\"correctRate\",\"type\":\"string\"}],\"name\":\"computingsharecallback\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"model\",\"type\":\"string\"},{\"name\":\"dataset\",\"type\":\"string\"},{\"name\":\"round\",\"type\":\"string\"},{\"name\":\"epoch\",\"type\":\"string\"}],\"name\":\"computingshare\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"_abstract\",\"type\":\"string\"},{\"name\":\"_source\",\"type\":\"string\"},{\"name\":\"_type\",\"type\":\"uint256\"}],\"name\":\"addmetadata\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"deletemetadata\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dataAbstract\",\"type\":\"string\"},{\"name\":\"source\",\"type\":\"string\"},{\"name\":\"dataType\",\"type\":\"uint256\"}],\"name\":\"createmetadata\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"_result\",\"type\":\"string\"}],\"name\":\"datasharecallback\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getresult\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getmetadata\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"dataAbstract\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"model\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"dataset\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"round\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"epoch\",\"type\":\"string\"}],\"name\":\"computingShareEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"source\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"dataAbstract\",\"type\":\"string\"}],\"name\":\"dataShareEvent\",\"type\":\"event\"}]"

// ComputingBin is the compiled bytecode used for deploying new contracts.
var ComputingBin = "0x608060405234801561001057600080fd5b506000600381905550611fd2806100286000396000f3006080604052600436106100a4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063025849e9146100a9578063213e42e61461014f57806334f614dc146101df57806343cadf4d146103575780635408ffe9146105155780635a52d17314610651578063a6eeb39b14610696578063b328e04514610763578063d1b878e91461084f578063ef41ff8514610a39575b600080fd5b3480156100b557600080fd5b506100d460048036038101908080359060200190929190505050610b5d565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101145780820151818401526020810190506100f9565b50505050905090810190601f1680156101415780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561015b57600080fd5b50610164610dfa565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101a4578082015181840152602081019050610189565b50505050905090810190601f1680156101d15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101eb57600080fd5b506102dc60048036038101908080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610e3f565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561031c578082015181840152602081019050610301565b50505050905090810190601f1680156103495780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561036357600080fd5b5061049a60048036038101908080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611077565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104da5780820151818401526020810190506104bf565b50505050905090810190601f1680156105075780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561052157600080fd5b506105d660048036038101908080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192908035906020019092919050505061142c565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106165780820151818401526020810190506105fb565b50505050905090810190601f1680156106435780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561065d57600080fd5b5061067c60048036038101908080359060200190929190505050611666565b604051808215151515815260200191505060405180910390f35b3480156106a257600080fd5b5061074d600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929080359060200190929190505050611765565b6040518082815260200191505060405180910390f35b34801561076f57600080fd5b506107d460048036038101908080359060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050611797565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156108145780820151818401526020810190506107f9565b50505050905090810190601f1680156108415780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561085b57600080fd5b5061087a600480360381019080803590602001909291905050506119ee565b6040518080602001806020018060200180602001858103855289818151815260200191508051906020019080838360005b838110156108c65780820151818401526020810190506108ab565b50505050905090810190601f1680156108f35780820380516001836020036101000a031916815260200191505b50858103845288818151815260200191508051906020019080838360005b8381101561092c578082015181840152602081019050610911565b50505050905090810190601f1680156109595780820380516001836020036101000a031916815260200191505b50858103835287818151815260200191508051906020019080838360005b83811015610992578082015181840152602081019050610977565b50505050905090810190601f1680156109bf5780820380516001836020036101000a031916815260200191505b50858103825286818151815260200191508051906020019080838360005b838110156109f85780820151818401526020810190506109dd565b50505050905090810190601f168015610a255780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b348015610a4557600080fd5b50610a6460048036038101908080359060200190929190505050611ccb565b60405180806020018581526020018060200184151515158152602001838103835287818151815260200191508051906020019080838360005b83811015610ab8578082015181840152602081019050610a9d565b50505050905090810190601f168015610ae55780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b83811015610b1e578082015181840152602081019050610b03565b50505050905090810190601f168015610b4b5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b6060600015156001600084815260200190815260200160002060030160009054906101000a900460ff16151514158015610bad575060026001600084815260200190815260200160002060020154145b1515610c21576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f676976656e20696420696e636f7272656374000000000000000000000000000081525060200191505060405180910390fd5b7fb73929caa2cdb2bbe2022b1b84ae0243553999a2dbd627193217719c1d02d1ab826001600085815260200190815260200160002060000160016000868152602001908152602001600020600101604051808481526020018060200180602001838103835285818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015610d025780601f10610cd757610100808354040283529160200191610d02565b820191906000526020600020905b815481529060010190602001808311610ce557829003601f168201915b5050838103825284818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015610d855780601f10610d5a57610100808354040283529160200191610d85565b820191906000526020600020905b815481529060010190602001808311610d6857829003601f168201915b50509550505050505060405180910390a1606060405190810160405280602181526020017f656d697420646174617368617265206576656e74207375636365737366756c6c81526020017f79000000000000000000000000000000000000000000000000000000000000008152509050919050565b606060006003819055506040805190810160405280601281526020017f696e69746c656467657220737563636573730000000000000000000000000000815250905090565b6060600015156001600087815260200190815260200160002060030160009054906101000a900460ff16151514158015610e8e5750600180600087815260200190815260200160002060020154145b1515610f02576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f676976656e20696420696e636f7272656374000000000000000000000000000081525060200191505060405180910390fd5b6080604051908101604052808581526020018481526020018381526020016020604051908101604052806000815250815250600260008781526020019081526020016000206000820151816000019080519060200190610f63929190611e81565b506020820151816001019080519060200190610f80929190611e81565b506040820151816002019080519060200190610f9d929190611e81565b506060820151816003019080519060200190610fba929190611e81565b50905050600260008681526020019081526020016000206002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110685780601f1061103d57610100808354040283529160200191611068565b820191906000526020600020905b81548152906001019060200180831161104b57829003601f168201915b50505050509050949350505050565b6060600015156001600088815260200190815260200160002060030160009054906101000a900460ff161515141580156110c65750600180600088815260200190815260200160002060020154145b151561113a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f676976656e20696420696e636f7272656374000000000000000000000000000081525060200191505060405180910390fd5b7fc2e5588cc05e73bc08ec881c55c11bb9a6f95a5c3a0d4293fdf745ab97c34c0e86600160008981526020019081526020016000206001018787878760405180878152602001806020018060200180602001806020018060200186810386528b8181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156112155780601f106111ea57610100808354040283529160200191611215565b820191906000526020600020905b8154815290600101906020018083116111f857829003601f168201915b505086810385528a818151815260200191508051906020019080838360005b8381101561124f578082015181840152602081019050611234565b50505050905090810190601f16801561127c5780820380516001836020036101000a031916815260200191505b50868103845289818151815260200191508051906020019080838360005b838110156112b557808201518184015260208101905061129a565b50505050905090810190601f1680156112e25780820380516001836020036101000a031916815260200191505b50868103835288818151815260200191508051906020019080838360005b8381101561131b578082015181840152602081019050611300565b50505050905090810190601f1680156113485780820380516001836020036101000a031916815260200191505b50868103825287818151815260200191508051906020019080838360005b83811015611381578082015181840152602081019050611366565b50505050905090810190601f1680156113ae5780820380516001836020036101000a031916815260200191505b509b50505050505050505050505060405180910390a1606060405190810160405280602681526020017f656d697420636f6d707574696e677368617265206576656e742073756363657381526020017f7366756c6c790000000000000000000000000000000000000000000000000000815250905095945050505050565b60606001600086815260200190815260200160002060030160009054906101000a900460ff161561150d57600160008681526020019081526020016000206001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156115015780601f106114d657610100808354040283529160200191611501565b820191906000526020600020905b8154815290600101906020018083116114e457829003601f168201915b5050505050905061165e565b83600160008781526020019081526020016000206001019080519060200190611537929190611f01565b508160016000878152602001908152602001600020600201819055508260016000878152602001908152602001600020600001908051906020019061157d929190611f01565b50600180600087815260200190815260200160002060030160006101000a81548160ff021916908315150217905550600160008681526020019081526020016000206001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156116565780601f1061162b57610100808354040283529160200191611656565b820191906000526020600020905b81548152906001019060200180831161163957829003601f168201915b505050505090505b949350505050565b60006001600083815260200190815260200160002060030160009054906101000a900460ff161561175b5760206040519081016040528060008152506001600084815260200190815260200160002060010190805190602001906116cb929190611f01565b506020604051908101604052806000815250600160008481526020019081526020016000206000019080519060200190611706929190611f01565b506000600160008481526020019081526020016000206002018190555060006001600084815260200190815260200160002060030160006101000a81548160ff02191690831515021790555060019050611760565b600090505b919050565b600061177560035485858561142c565b5060036000815480929190600101919050555060016003540390509392505050565b6060600015156001600085815260200190815260200160002060030160009054906101000a900460ff161515141580156117e7575060026001600085815260200190815260200160002060020154145b151561185b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f676976656e20696420696e636f7272656374000000000000000000000000000081525060200191505060405180910390fd5b608060405190810160405280602060405190810160405280600081525081526020016020604051908101604052806000815250815260200160206040519081016040528060008152508152602001838152506002600085815260200190815260200160002060008201518160000190805190602001906118dc929190611e81565b5060208201518160010190805190602001906118f9929190611e81565b506040820151816002019080519060200190611916929190611e81565b506060820151816003019080519060200190611933929190611e81565b50905050600260008481526020019081526020016000206003018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156119e15780601f106119b6576101008083540402835291602001916119e1565b820191906000526020600020905b8154815290600101906020018083116119c457829003601f168201915b5050505050905092915050565b60608060608060026000868152602001908152602001600020600001600260008781526020019081526020016000206001016002600088815260200190815260200160002060020160026000898152602001908152602001600020600301838054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611ae15780601f10611ab657610100808354040283529160200191611ae1565b820191906000526020600020905b815481529060010190602001808311611ac457829003601f168201915b50505050509350828054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611b7d5780601f10611b5257610100808354040283529160200191611b7d565b820191906000526020600020905b815481529060010190602001808311611b6057829003601f168201915b50505050509250818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611c195780601f10611bee57610100808354040283529160200191611c19565b820191906000526020600020905b815481529060010190602001808311611bfc57829003601f168201915b50505050509150808054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611cb55780601f10611c8a57610100808354040283529160200191611cb5565b820191906000526020600020905b815481529060010190602001808311611c9857829003601f168201915b5050505050905093509350935093509193509193565b6060600060606000600160008681526020019081526020016000206001016001600087815260200190815260200160002060020154600160008881526020019081526020016000206000016001600089815260200190815260200160002060030160009054906101000a900460ff16838054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611dcf5780601f10611da457610100808354040283529160200191611dcf565b820191906000526020600020905b815481529060010190602001808311611db257829003601f168201915b50505050509350818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611e6b5780601f10611e4057610100808354040283529160200191611e6b565b820191906000526020600020905b815481529060010190602001808311611e4e57829003601f168201915b5050505050915093509350935093509193509193565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611ec257805160ff1916838001178555611ef0565b82800160010185558215611ef0579182015b82811115611eef578251825591602001919060010190611ed4565b5b509050611efd9190611f81565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611f4257805160ff1916838001178555611f70565b82800160010185558215611f70579182015b82811115611f6f578251825591602001919060010190611f54565b5b509050611f7d9190611f81565b5090565b611fa391905b80821115611f9f576000816000905550600101611f87565b5090565b905600a165627a7a723058203b7f86cf6c55899599da05ec3d328f7a5ab96fbd59d819b71f4b74d21b20b6d60029"

// DeployComputing deploys a new contract, binding an instance of Computing to it.
func DeployComputing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Computing, error) {
	parsed, err := abi.JSON(strings.NewReader(ComputingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ComputingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Computing{ComputingCaller: ComputingCaller{contract: contract}, ComputingTransactor: ComputingTransactor{contract: contract}, ComputingFilterer: ComputingFilterer{contract: contract}}, nil
}

func AsyncDeployComputing(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(ComputingABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(ComputingBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Computing is an auto generated Go binding around a Solidity contract.
type Computing struct {
	ComputingCaller     // Read-only binding to the contract
	ComputingTransactor // Write-only binding to the contract
	ComputingFilterer   // Log filterer for contract events
}

// ComputingCaller is an auto generated read-only Go binding around a Solidity contract.
type ComputingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComputingTransactor is an auto generated write-only Go binding around a Solidity contract.
type ComputingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComputingFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ComputingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComputingSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ComputingSession struct {
	Contract     *Computing        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ComputingCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ComputingCallerSession struct {
	Contract *ComputingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ComputingTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ComputingTransactorSession struct {
	Contract     *ComputingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ComputingRaw is an auto generated low-level Go binding around a Solidity contract.
type ComputingRaw struct {
	Contract *Computing // Generic contract binding to access the raw methods on
}

// ComputingCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ComputingCallerRaw struct {
	Contract *ComputingCaller // Generic read-only contract binding to access the raw methods on
}

// ComputingTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ComputingTransactorRaw struct {
	Contract *ComputingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComputing creates a new instance of Computing, bound to a specific deployed contract.
func NewComputing(address common.Address, backend bind.ContractBackend) (*Computing, error) {
	contract, err := bindComputing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Computing{ComputingCaller: ComputingCaller{contract: contract}, ComputingTransactor: ComputingTransactor{contract: contract}, ComputingFilterer: ComputingFilterer{contract: contract}}, nil
}

// NewComputingCaller creates a new read-only instance of Computing, bound to a specific deployed contract.
func NewComputingCaller(address common.Address, caller bind.ContractCaller) (*ComputingCaller, error) {
	contract, err := bindComputing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComputingCaller{contract: contract}, nil
}

// NewComputingTransactor creates a new write-only instance of Computing, bound to a specific deployed contract.
func NewComputingTransactor(address common.Address, transactor bind.ContractTransactor) (*ComputingTransactor, error) {
	contract, err := bindComputing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComputingTransactor{contract: contract}, nil
}

// NewComputingFilterer creates a new log filterer instance of Computing, bound to a specific deployed contract.
func NewComputingFilterer(address common.Address, filterer bind.ContractFilterer) (*ComputingFilterer, error) {
	contract, err := bindComputing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComputingFilterer{contract: contract}, nil
}

// bindComputing binds a generic wrapper to an already deployed contract.
func bindComputing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ComputingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Computing *ComputingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Computing.Contract.ComputingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Computing *ComputingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.ComputingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Computing *ComputingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.ComputingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Computing *ComputingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Computing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Computing *ComputingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Computing *ComputingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.contract.Transact(opts, method, params...)
}

// Addmetadata is a paid mutator transaction binding the contract method 0x5408ffe9.
//
// Solidity: function addmetadata(uint256 id, string _abstract, string _source, uint256 _type) returns(string)
func (_Computing *ComputingTransactor) Addmetadata(opts *bind.TransactOpts, id *big.Int, _abstract string, _source string, _type *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Computing.contract.Transact(opts, "addmetadata", id, _abstract, _source, _type)
}

func (_Computing *ComputingTransactor) AsyncAddmetadata(handler func(*types.Receipt, error), opts *bind.TransactOpts, id *big.Int, _abstract string, _source string, _type *big.Int) (*types.Transaction, error) {
	return _Computing.contract.AsyncTransact(opts, handler, "addmetadata", id, _abstract, _source, _type)
}

// Addmetadata is a paid mutator transaction binding the contract method 0x5408ffe9.
//
// Solidity: function addmetadata(uint256 id, string _abstract, string _source, uint256 _type) returns(string)
func (_Computing *ComputingSession) Addmetadata(id *big.Int, _abstract string, _source string, _type *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Addmetadata(&_Computing.TransactOpts, id, _abstract, _source, _type)
}

func (_Computing *ComputingSession) AsyncAddmetadata(handler func(*types.Receipt, error), id *big.Int, _abstract string, _source string, _type *big.Int) (*types.Transaction, error) {
	return _Computing.Contract.AsyncAddmetadata(handler, &_Computing.TransactOpts, id, _abstract, _source, _type)
}

// Addmetadata is a paid mutator transaction binding the contract method 0x5408ffe9.
//
// Solidity: function addmetadata(uint256 id, string _abstract, string _source, uint256 _type) returns(string)
func (_Computing *ComputingTransactorSession) Addmetadata(id *big.Int, _abstract string, _source string, _type *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Addmetadata(&_Computing.TransactOpts, id, _abstract, _source, _type)
}

func (_Computing *ComputingTransactorSession) AsyncAddmetadata(handler func(*types.Receipt, error), id *big.Int, _abstract string, _source string, _type *big.Int) (*types.Transaction, error) {
	return _Computing.Contract.AsyncAddmetadata(handler, &_Computing.TransactOpts, id, _abstract, _source, _type)
}

// Computingshare is a paid mutator transaction binding the contract method 0x43cadf4d.
//
// Solidity: function computingshare(uint256 id, string model, string dataset, string round, string epoch) returns(string)
func (_Computing *ComputingTransactor) Computingshare(opts *bind.TransactOpts, id *big.Int, model string, dataset string, round string, epoch string) (*types.Transaction, *types.Receipt, error) {
	return _Computing.contract.Transact(opts, "computingshare", id, model, dataset, round, epoch)
}

func (_Computing *ComputingTransactor) AsyncComputingshare(handler func(*types.Receipt, error), opts *bind.TransactOpts, id *big.Int, model string, dataset string, round string, epoch string) (*types.Transaction, error) {
	return _Computing.contract.AsyncTransact(opts, handler, "computingshare", id, model, dataset, round, epoch)
}

// Computingshare is a paid mutator transaction binding the contract method 0x43cadf4d.
//
// Solidity: function computingshare(uint256 id, string model, string dataset, string round, string epoch) returns(string)
func (_Computing *ComputingSession) Computingshare(id *big.Int, model string, dataset string, round string, epoch string) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Computingshare(&_Computing.TransactOpts, id, model, dataset, round, epoch)
}

func (_Computing *ComputingSession) AsyncComputingshare(handler func(*types.Receipt, error), id *big.Int, model string, dataset string, round string, epoch string) (*types.Transaction, error) {
	return _Computing.Contract.AsyncComputingshare(handler, &_Computing.TransactOpts, id, model, dataset, round, epoch)
}

// Computingshare is a paid mutator transaction binding the contract method 0x43cadf4d.
//
// Solidity: function computingshare(uint256 id, string model, string dataset, string round, string epoch) returns(string)
func (_Computing *ComputingTransactorSession) Computingshare(id *big.Int, model string, dataset string, round string, epoch string) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Computingshare(&_Computing.TransactOpts, id, model, dataset, round, epoch)
}

func (_Computing *ComputingTransactorSession) AsyncComputingshare(handler func(*types.Receipt, error), id *big.Int, model string, dataset string, round string, epoch string) (*types.Transaction, error) {
	return _Computing.Contract.AsyncComputingshare(handler, &_Computing.TransactOpts, id, model, dataset, round, epoch)
}

// Computingsharecallback is a paid mutator transaction binding the contract method 0x34f614dc.
//
// Solidity: function computingsharecallback(uint256 id, string paramAddr, string paramAbstract, string correctRate) returns(string)
func (_Computing *ComputingTransactor) Computingsharecallback(opts *bind.TransactOpts, id *big.Int, paramAddr string, paramAbstract string, correctRate string) (*types.Transaction, *types.Receipt, error) {
	return _Computing.contract.Transact(opts, "computingsharecallback", id, paramAddr, paramAbstract, correctRate)
}

func (_Computing *ComputingTransactor) AsyncComputingsharecallback(handler func(*types.Receipt, error), opts *bind.TransactOpts, id *big.Int, paramAddr string, paramAbstract string, correctRate string) (*types.Transaction, error) {
	return _Computing.contract.AsyncTransact(opts, handler, "computingsharecallback", id, paramAddr, paramAbstract, correctRate)
}

// Computingsharecallback is a paid mutator transaction binding the contract method 0x34f614dc.
//
// Solidity: function computingsharecallback(uint256 id, string paramAddr, string paramAbstract, string correctRate) returns(string)
func (_Computing *ComputingSession) Computingsharecallback(id *big.Int, paramAddr string, paramAbstract string, correctRate string) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Computingsharecallback(&_Computing.TransactOpts, id, paramAddr, paramAbstract, correctRate)
}

func (_Computing *ComputingSession) AsyncComputingsharecallback(handler func(*types.Receipt, error), id *big.Int, paramAddr string, paramAbstract string, correctRate string) (*types.Transaction, error) {
	return _Computing.Contract.AsyncComputingsharecallback(handler, &_Computing.TransactOpts, id, paramAddr, paramAbstract, correctRate)
}

// Computingsharecallback is a paid mutator transaction binding the contract method 0x34f614dc.
//
// Solidity: function computingsharecallback(uint256 id, string paramAddr, string paramAbstract, string correctRate) returns(string)
func (_Computing *ComputingTransactorSession) Computingsharecallback(id *big.Int, paramAddr string, paramAbstract string, correctRate string) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Computingsharecallback(&_Computing.TransactOpts, id, paramAddr, paramAbstract, correctRate)
}

func (_Computing *ComputingTransactorSession) AsyncComputingsharecallback(handler func(*types.Receipt, error), id *big.Int, paramAddr string, paramAbstract string, correctRate string) (*types.Transaction, error) {
	return _Computing.Contract.AsyncComputingsharecallback(handler, &_Computing.TransactOpts, id, paramAddr, paramAbstract, correctRate)
}

// Createmetadata is a paid mutator transaction binding the contract method 0xa6eeb39b.
//
// Solidity: function createmetadata(string dataAbstract, string source, uint256 dataType) returns(uint256)
func (_Computing *ComputingTransactor) Createmetadata(opts *bind.TransactOpts, dataAbstract string, source string, dataType *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Computing.contract.Transact(opts, "createmetadata", dataAbstract, source, dataType)
}

func (_Computing *ComputingTransactor) AsyncCreatemetadata(handler func(*types.Receipt, error), opts *bind.TransactOpts, dataAbstract string, source string, dataType *big.Int) (*types.Transaction, error) {
	return _Computing.contract.AsyncTransact(opts, handler, "createmetadata", dataAbstract, source, dataType)
}

// Createmetadata is a paid mutator transaction binding the contract method 0xa6eeb39b.
//
// Solidity: function createmetadata(string dataAbstract, string source, uint256 dataType) returns(uint256)
func (_Computing *ComputingSession) Createmetadata(dataAbstract string, source string, dataType *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Createmetadata(&_Computing.TransactOpts, dataAbstract, source, dataType)
}

func (_Computing *ComputingSession) AsyncCreatemetadata(handler func(*types.Receipt, error), dataAbstract string, source string, dataType *big.Int) (*types.Transaction, error) {
	return _Computing.Contract.AsyncCreatemetadata(handler, &_Computing.TransactOpts, dataAbstract, source, dataType)
}

// Createmetadata is a paid mutator transaction binding the contract method 0xa6eeb39b.
//
// Solidity: function createmetadata(string dataAbstract, string source, uint256 dataType) returns(uint256)
func (_Computing *ComputingTransactorSession) Createmetadata(dataAbstract string, source string, dataType *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Createmetadata(&_Computing.TransactOpts, dataAbstract, source, dataType)
}

func (_Computing *ComputingTransactorSession) AsyncCreatemetadata(handler func(*types.Receipt, error), dataAbstract string, source string, dataType *big.Int) (*types.Transaction, error) {
	return _Computing.Contract.AsyncCreatemetadata(handler, &_Computing.TransactOpts, dataAbstract, source, dataType)
}

// Datashare is a paid mutator transaction binding the contract method 0x025849e9.
//
// Solidity: function datashare(uint256 id) returns(string)
func (_Computing *ComputingTransactor) Datashare(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Computing.contract.Transact(opts, "datashare", id)
}

func (_Computing *ComputingTransactor) AsyncDatashare(handler func(*types.Receipt, error), opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Computing.contract.AsyncTransact(opts, handler, "datashare", id)
}

// Datashare is a paid mutator transaction binding the contract method 0x025849e9.
//
// Solidity: function datashare(uint256 id) returns(string)
func (_Computing *ComputingSession) Datashare(id *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Datashare(&_Computing.TransactOpts, id)
}

func (_Computing *ComputingSession) AsyncDatashare(handler func(*types.Receipt, error), id *big.Int) (*types.Transaction, error) {
	return _Computing.Contract.AsyncDatashare(handler, &_Computing.TransactOpts, id)
}

// Datashare is a paid mutator transaction binding the contract method 0x025849e9.
//
// Solidity: function datashare(uint256 id) returns(string)
func (_Computing *ComputingTransactorSession) Datashare(id *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Datashare(&_Computing.TransactOpts, id)
}

func (_Computing *ComputingTransactorSession) AsyncDatashare(handler func(*types.Receipt, error), id *big.Int) (*types.Transaction, error) {
	return _Computing.Contract.AsyncDatashare(handler, &_Computing.TransactOpts, id)
}

// Datasharecallback is a paid mutator transaction binding the contract method 0xb328e045.
//
// Solidity: function datasharecallback(uint256 id, string _result) returns(string)
func (_Computing *ComputingTransactor) Datasharecallback(opts *bind.TransactOpts, id *big.Int, _result string) (*types.Transaction, *types.Receipt, error) {
	return _Computing.contract.Transact(opts, "datasharecallback", id, _result)
}

func (_Computing *ComputingTransactor) AsyncDatasharecallback(handler func(*types.Receipt, error), opts *bind.TransactOpts, id *big.Int, _result string) (*types.Transaction, error) {
	return _Computing.contract.AsyncTransact(opts, handler, "datasharecallback", id, _result)
}

// Datasharecallback is a paid mutator transaction binding the contract method 0xb328e045.
//
// Solidity: function datasharecallback(uint256 id, string _result) returns(string)
func (_Computing *ComputingSession) Datasharecallback(id *big.Int, _result string) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Datasharecallback(&_Computing.TransactOpts, id, _result)
}

func (_Computing *ComputingSession) AsyncDatasharecallback(handler func(*types.Receipt, error), id *big.Int, _result string) (*types.Transaction, error) {
	return _Computing.Contract.AsyncDatasharecallback(handler, &_Computing.TransactOpts, id, _result)
}

// Datasharecallback is a paid mutator transaction binding the contract method 0xb328e045.
//
// Solidity: function datasharecallback(uint256 id, string _result) returns(string)
func (_Computing *ComputingTransactorSession) Datasharecallback(id *big.Int, _result string) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Datasharecallback(&_Computing.TransactOpts, id, _result)
}

func (_Computing *ComputingTransactorSession) AsyncDatasharecallback(handler func(*types.Receipt, error), id *big.Int, _result string) (*types.Transaction, error) {
	return _Computing.Contract.AsyncDatasharecallback(handler, &_Computing.TransactOpts, id, _result)
}

// Deletemetadata is a paid mutator transaction binding the contract method 0x5a52d173.
//
// Solidity: function deletemetadata(uint256 id) returns(bool)
func (_Computing *ComputingTransactor) Deletemetadata(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Computing.contract.Transact(opts, "deletemetadata", id)
}

func (_Computing *ComputingTransactor) AsyncDeletemetadata(handler func(*types.Receipt, error), opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Computing.contract.AsyncTransact(opts, handler, "deletemetadata", id)
}

// Deletemetadata is a paid mutator transaction binding the contract method 0x5a52d173.
//
// Solidity: function deletemetadata(uint256 id) returns(bool)
func (_Computing *ComputingSession) Deletemetadata(id *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Deletemetadata(&_Computing.TransactOpts, id)
}

func (_Computing *ComputingSession) AsyncDeletemetadata(handler func(*types.Receipt, error), id *big.Int) (*types.Transaction, error) {
	return _Computing.Contract.AsyncDeletemetadata(handler, &_Computing.TransactOpts, id)
}

// Deletemetadata is a paid mutator transaction binding the contract method 0x5a52d173.
//
// Solidity: function deletemetadata(uint256 id) returns(bool)
func (_Computing *ComputingTransactorSession) Deletemetadata(id *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Deletemetadata(&_Computing.TransactOpts, id)
}

func (_Computing *ComputingTransactorSession) AsyncDeletemetadata(handler func(*types.Receipt, error), id *big.Int) (*types.Transaction, error) {
	return _Computing.Contract.AsyncDeletemetadata(handler, &_Computing.TransactOpts, id)
}

// Getmetadata is a paid mutator transaction binding the contract method 0xef41ff85.
//
// Solidity: function getmetadata(uint256 id) returns(string, uint256, string, bool)
func (_Computing *ComputingTransactor) Getmetadata(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Computing.contract.Transact(opts, "getmetadata", id)
}

func (_Computing *ComputingTransactor) AsyncGetmetadata(handler func(*types.Receipt, error), opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Computing.contract.AsyncTransact(opts, handler, "getmetadata", id)
}

// Getmetadata is a paid mutator transaction binding the contract method 0xef41ff85.
//
// Solidity: function getmetadata(uint256 id) returns(string, uint256, string, bool)
func (_Computing *ComputingSession) Getmetadata(id *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Getmetadata(&_Computing.TransactOpts, id)
}

func (_Computing *ComputingSession) AsyncGetmetadata(handler func(*types.Receipt, error), id *big.Int) (*types.Transaction, error) {
	return _Computing.Contract.AsyncGetmetadata(handler, &_Computing.TransactOpts, id)
}

// Getmetadata is a paid mutator transaction binding the contract method 0xef41ff85.
//
// Solidity: function getmetadata(uint256 id) returns(string, uint256, string, bool)
func (_Computing *ComputingTransactorSession) Getmetadata(id *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Getmetadata(&_Computing.TransactOpts, id)
}

func (_Computing *ComputingTransactorSession) AsyncGetmetadata(handler func(*types.Receipt, error), id *big.Int) (*types.Transaction, error) {
	return _Computing.Contract.AsyncGetmetadata(handler, &_Computing.TransactOpts, id)
}

// Getresult is a paid mutator transaction binding the contract method 0xd1b878e9.
//
// Solidity: function getresult(uint256 id) returns(string, string, string, string)
func (_Computing *ComputingTransactor) Getresult(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Computing.contract.Transact(opts, "getresult", id)
}

func (_Computing *ComputingTransactor) AsyncGetresult(handler func(*types.Receipt, error), opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Computing.contract.AsyncTransact(opts, handler, "getresult", id)
}

// Getresult is a paid mutator transaction binding the contract method 0xd1b878e9.
//
// Solidity: function getresult(uint256 id) returns(string, string, string, string)
func (_Computing *ComputingSession) Getresult(id *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Getresult(&_Computing.TransactOpts, id)
}

func (_Computing *ComputingSession) AsyncGetresult(handler func(*types.Receipt, error), id *big.Int) (*types.Transaction, error) {
	return _Computing.Contract.AsyncGetresult(handler, &_Computing.TransactOpts, id)
}

// Getresult is a paid mutator transaction binding the contract method 0xd1b878e9.
//
// Solidity: function getresult(uint256 id) returns(string, string, string, string)
func (_Computing *ComputingTransactorSession) Getresult(id *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Getresult(&_Computing.TransactOpts, id)
}

func (_Computing *ComputingTransactorSession) AsyncGetresult(handler func(*types.Receipt, error), id *big.Int) (*types.Transaction, error) {
	return _Computing.Contract.AsyncGetresult(handler, &_Computing.TransactOpts, id)
}

// Initledger is a paid mutator transaction binding the contract method 0x213e42e6.
//
// Solidity: function initledger() returns(string)
func (_Computing *ComputingTransactor) Initledger(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Computing.contract.Transact(opts, "initledger")
}

func (_Computing *ComputingTransactor) AsyncInitledger(handler func(*types.Receipt, error), opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Computing.contract.AsyncTransact(opts, handler, "initledger")
}

// Initledger is a paid mutator transaction binding the contract method 0x213e42e6.
//
// Solidity: function initledger() returns(string)
func (_Computing *ComputingSession) Initledger() (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Initledger(&_Computing.TransactOpts)
}

func (_Computing *ComputingSession) AsyncInitledger(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Computing.Contract.AsyncInitledger(handler, &_Computing.TransactOpts)
}

// Initledger is a paid mutator transaction binding the contract method 0x213e42e6.
//
// Solidity: function initledger() returns(string)
func (_Computing *ComputingTransactorSession) Initledger() (*types.Transaction, *types.Receipt, error) {
	return _Computing.Contract.Initledger(&_Computing.TransactOpts)
}

func (_Computing *ComputingTransactorSession) AsyncInitledger(handler func(*types.Receipt, error)) (*types.Transaction, error) {
	return _Computing.Contract.AsyncInitledger(handler, &_Computing.TransactOpts)
}

// ComputingComputingShareEventIterator is returned from FilterComputingShareEvent and is used to iterate over the raw logs and unpacked data for ComputingShareEvent events raised by the Computing contract.
type ComputingComputingShareEventIterator struct {
	Event *ComputingComputingShareEvent // Event containing the contract specifics and raw log

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
func (it *ComputingComputingShareEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComputingComputingShareEvent)
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
		it.Event = new(ComputingComputingShareEvent)
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
func (it *ComputingComputingShareEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComputingComputingShareEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComputingComputingShareEvent represents a ComputingShareEvent event raised by the Computing contract.
type ComputingComputingShareEvent struct {
	Id           *big.Int
	DataAbstract string
	Model        string
	Dataset      string
	Round        string
	Epoch        string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterComputingShareEvent is a free log retrieval operation binding the contract event 0xc2e5588cc05e73bc08ec881c55c11bb9a6f95a5c3a0d4293fdf745ab97c34c0e.
//
// Solidity: event computingShareEvent(uint256 id, string dataAbstract, string model, string dataset, string round, string epoch)
func (_Computing *ComputingFilterer) FilterComputingShareEvent(opts *bind.FilterOpts) (*ComputingComputingShareEventIterator, error) {

	logs, sub, err := _Computing.contract.FilterLogs(opts, "computingShareEvent")
	if err != nil {
		return nil, err
	}
	return &ComputingComputingShareEventIterator{contract: _Computing.contract, event: "computingShareEvent", logs: logs, sub: sub}, nil
}

// WatchComputingShareEvent is a free log subscription operation binding the contract event 0xc2e5588cc05e73bc08ec881c55c11bb9a6f95a5c3a0d4293fdf745ab97c34c0e.
//
// Solidity: event computingShareEvent(uint256 id, string dataAbstract, string model, string dataset, string round, string epoch)
func (_Computing *ComputingFilterer) WatchComputingShareEvent(opts *bind.WatchOpts, sink chan<- *ComputingComputingShareEvent) (event.Subscription, error) {

	logs, sub, err := _Computing.contract.WatchLogs(opts, "computingShareEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComputingComputingShareEvent)
				if err := _Computing.contract.UnpackLog(event, "computingShareEvent", log); err != nil {
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

// ParseComputingShareEvent is a log parse operation binding the contract event 0xc2e5588cc05e73bc08ec881c55c11bb9a6f95a5c3a0d4293fdf745ab97c34c0e.
//
// Solidity: event computingShareEvent(uint256 id, string dataAbstract, string model, string dataset, string round, string epoch)
func (_Computing *ComputingFilterer) ParseComputingShareEvent(log types.Log) (*ComputingComputingShareEvent, error) {
	event := new(ComputingComputingShareEvent)
	if err := _Computing.contract.UnpackLog(event, "computingShareEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ComputingDataShareEventIterator is returned from FilterDataShareEvent and is used to iterate over the raw logs and unpacked data for DataShareEvent events raised by the Computing contract.
type ComputingDataShareEventIterator struct {
	Event *ComputingDataShareEvent // Event containing the contract specifics and raw log

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
func (it *ComputingDataShareEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComputingDataShareEvent)
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
		it.Event = new(ComputingDataShareEvent)
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
func (it *ComputingDataShareEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComputingDataShareEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComputingDataShareEvent represents a DataShareEvent event raised by the Computing contract.
type ComputingDataShareEvent struct {
	Id           *big.Int
	Source       string
	DataAbstract string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDataShareEvent is a free log retrieval operation binding the contract event 0xb73929caa2cdb2bbe2022b1b84ae0243553999a2dbd627193217719c1d02d1ab.
//
// Solidity: event dataShareEvent(uint256 id, string source, string dataAbstract)
func (_Computing *ComputingFilterer) FilterDataShareEvent(opts *bind.FilterOpts) (*ComputingDataShareEventIterator, error) {

	logs, sub, err := _Computing.contract.FilterLogs(opts, "dataShareEvent")
	if err != nil {
		return nil, err
	}
	return &ComputingDataShareEventIterator{contract: _Computing.contract, event: "dataShareEvent", logs: logs, sub: sub}, nil
}

// WatchDataShareEvent is a free log subscription operation binding the contract event 0xb73929caa2cdb2bbe2022b1b84ae0243553999a2dbd627193217719c1d02d1ab.
//
// Solidity: event dataShareEvent(uint256 id, string source, string dataAbstract)
func (_Computing *ComputingFilterer) WatchDataShareEvent(opts *bind.WatchOpts, sink chan<- *ComputingDataShareEvent) (event.Subscription, error) {

	logs, sub, err := _Computing.contract.WatchLogs(opts, "dataShareEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComputingDataShareEvent)
				if err := _Computing.contract.UnpackLog(event, "dataShareEvent", log); err != nil {
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

// ParseDataShareEvent is a log parse operation binding the contract event 0xb73929caa2cdb2bbe2022b1b84ae0243553999a2dbd627193217719c1d02d1ab.
//
// Solidity: event dataShareEvent(uint256 id, string source, string dataAbstract)
func (_Computing *ComputingFilterer) ParseDataShareEvent(log types.Log) (*ComputingDataShareEvent, error) {
	event := new(ComputingDataShareEvent)
	if err := _Computing.contract.UnpackLog(event, "dataShareEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
