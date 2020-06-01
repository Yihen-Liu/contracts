// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proof

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ProofABI is the input ABI used to generate the binding from.
const ProofABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"}],\"name\":\"Challenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"ChallengeList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"file\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"FilePledge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetChallengeList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"}],\"name\":\"GetChallengeResponse\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NodeLeave\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NodePledge\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"file\",\"type\":\"string\"}],\"name\":\"PledgeForFile\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"PledgeForNode\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"ProofList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"file\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"exist\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_proof\",\"type\":\"string\"}],\"name\":\"ProofProvide\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"file\",\"type\":\"address\"}],\"name\":\"Punish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ProofBin is the compiled bytecode used for deploying new contracts.
const ProofBin = `60806040526103e8600460146101000a81548163ffffffff021916908363ffffffff1602179055503480156100345760006000fd5b505b33600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b61007d565b6115308061008c6000396000f3fe6080604052600436106100e15760003560e01c80638337846a1161007f578063b8aeb13611610059578063b8aeb13614610296578063bde4efe4146102d6578063d66a3121146102f2578063e5fe233514610333576100e1565b80638337846a146102145780638da5cb5b14610252578063a59909961461027e576100e1565b80635a706771116100bb5780635a7067711461015357806360dadb4e1461019157806381a53877146101d257806383197ef0146101fc576100e1565b80630bd8dbe4146100e75780630d9cb378146100ff57806311ec5d8014610129576100e1565b60006000fd5b3480156100f45760006000fd5b506100fd61034f565b005b34801561010c5760006000fd5b506101276004803603810190610122919061105d565b610454565b005b3480156101365760006000fd5b50610151600480360381019061014c9190610fda565b6105fa565b005b3480156101605760006000fd5b5061017b60048036038101906101769190610faf565b6105ff565b6040516101889190611377565b60405180910390f35b34801561019e5760006000fd5b506101b960048036038101906101b49190611019565b610627565b6040516101c99493929190611309565b60405180910390f35b3480156101df5760006000fd5b506101fa60048036038101906101f59190611019565b610748565b005b3480156102095760006000fd5b50610212610831565b005b3480156102215760006000fd5b5061023c60048036038101906102379190611019565b6108a9565b604051610249919061125a565b60405180910390f35b34801561025f5760006000fd5b5061026861094e565b604051610275919061125a565b60405180910390f35b34801561028b5760006000fd5b50610294610974565b005b3480156102a35760006000fd5b506102be60048036038101906102b99190611019565b610977565b6040516102cd93929190611276565b60405180910390f35b6102f060048036038101906102eb91906110ce565b610a82565b005b3480156102ff5760006000fd5b5061031a60048036038101906103159190611019565b610b49565b60405161032a94939291906112b5565b60405180910390f35b61034d6004803603810190610348919061110d565b610cf8565b005b3373ffffffffffffffffffffffffffffffffffffffff166108fc600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff1663ffffffff169081150290604051600060405180830381858888f193505050501580156103ee573d600060003e3d6000fd5b506000600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff1602179055505b565b60016003600050836040516104699190611242565b908152602001604051809103902060005060030160009054906101000a900463ffffffff1663ffffffff16141561049f576105f6565b60006002600050836040516104b49190611242565b908152602001604051809103902060005060020160006101000a81548160ff02191690831515021790555060405180608001604052803373ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001600163ffffffff168152602001506003600050836040516105339190611242565b908152602001604051809103902060005060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160005090805190602001906105aa929190610e76565b5060408201518160020160005090805190602001906105ca929190610e76565b5060608201518160030160006101000a81548163ffffffff021916908363ffffffff1602179055509050505b5050565b5b5050565b600060005060205280600052604060002060009150909054906101000a900463ffffffff1681565b6001600050818051602081018201805184825260208301602085012081835280955050505050506000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900463ffffffff1690806001016000508054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561072b5780601f106107005761010080835404028352916020019161072b565b820191906000526020600020905b81548152906001019060200180831161070e57829003601f168201915b5050505050908060020160009054906101000a900460ff16905084565b60405180606001604052803373ffffffffffffffffffffffffffffffffffffffff168152602001828152602001600115158152602001506002600050826040516107929190611242565b908152602001604051809103902060005060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001016000509080519060200190610809929190610e76565b5060408201518160020160006101000a81548160ff0219169083151502179055509050505b50565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561088e5760006000fd5b3373ffffffffffffffffffffffffffffffffffffffff16ff5b565b600060006003600050836040516108c09190611242565b908152602001604051809103902060005060030160009054906101000a900463ffffffff1663ffffffff1614156108fa5760009050610949565b60036000508260405161090d9190611242565b908152602001604051809103902060005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050610949565b919050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5b565b6002600050818051602081018201805184825260208301602085012081835280955050505050506000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001016000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a655780601f10610a3a57610100808354040283529160200191610a65565b820191906000526020600020905b815481529060010190602001808311610a4857829003601f168201915b5050505050908060020160009054906101000a900460ff16905083565b8163ffffffff163373ffffffffffffffffffffffffffffffffffffffff1631111515610ae3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ada90611356565b60405180910390fd5b81600060005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff1602179055505b5050565b6003600050818051602081018201805184825260208301602085012081835280955050505050506000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001016000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c375780601f10610c0c57610100808354040283529160200191610c37565b820191906000526020600020905b815481529060010190602001808311610c1a57829003601f168201915b505050505090806002016000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610cd85780601f10610cad57610100808354040283529160200191610cd8565b820191906000526020600020905b815481529060010190602001808311610cbb57829003601f168201915b5050505050908060030160009054906101000a900463ffffffff16905084565b8163ffffffff163373ffffffffffffffffffffffffffffffffffffffff1631111515610d59576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d5090611356565b60405180910390fd5b60405180608001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018363ffffffff16815260200182815260200160011515815260200150600160005082604051610daf9190611242565b908152602001604051809103902060005060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548163ffffffff021916908363ffffffff1602179055506040820151816001016000509080519060200190610e4d929190610e76565b5060608201518160020160006101000a81548160ff0219169083151502179055509050505b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610eb757805160ff1916838001178555610eea565b82800160010185558215610eea579182015b82811115610ee95782518260005090905591602001919060010190610ec9565b5b509050610ef79190610efb565b5090565b610f239190610f05565b80821115610f1f5760008181506000905550600101610f05565b5090565b90566114f9565b600081359050610f39816114c3565b5b92915050565b600082601f8301121515610f545760006000fd5b8135610f67610f62826113c2565b611393565b91508082526020830160208301858383011115610f845760006000fd5b610f8f83828461146c565b5050505b92915050565b600081359050610fa8816114de565b5b92915050565b600060208284031215610fc25760006000fd5b6000610fd084828501610f2a565b9150505b92915050565b6000600060408385031215610fef5760006000fd5b6000610ffd85828601610f2a565b925050602061100e85828601610f2a565b9150505b9250929050565b60006020828403121561102c5760006000fd5b600082013567ffffffffffffffff8111156110475760006000fd5b61105384828501610f40565b9150505b92915050565b60006000604083850312156110725760006000fd5b600083013567ffffffffffffffff81111561108d5760006000fd5b61109985828601610f40565b925050602083013567ffffffffffffffff8111156110b75760006000fd5b6110c385828601610f40565b9150505b9250929050565b60006000604083850312156110e35760006000fd5b60006110f185828601610f99565b925050602061110285828601610f2a565b9150505b9250929050565b60006000604083850312156111225760006000fd5b600061113085828601610f99565b925050602083013567ffffffffffffffff81111561114e5760006000fd5b61115a85828601610f40565b9150505b9250929050565b61116e8161141a565b82525b5050565b61117e8161142d565b82525b5050565b6000611190826113f0565b61119a81856113fc565b93506111aa81856020860161147c565b6111b3816114b1565b84019150505b92915050565b60006111ca826113f0565b6111d4818561140e565b93506111e481856020860161147c565b8084019150505b92915050565b60006111fe6015836113fc565b91507f62616c616e6365206973206e6f7420656e6f756768000000000000000000000060008301526020820190505b919050565b61123b8161145b565b82525b5050565b600061124e82846111bf565b91508190505b92915050565b600060208201905061126f6000830184611165565b5b92915050565b600060608201905061128b6000830186611165565b818103602083015261129d8185611185565b90506112ac6040830184611175565b5b949350505050565b60006080820190506112ca6000830187611165565b81810360208301526112dc8186611185565b905081810360408301526112f08185611185565b90506112ff6060830184611232565b5b95945050505050565b600060808201905061131e6000830187611165565b61132b6020830186611232565b818103604083015261133d8185611185565b905061134c6060830184611175565b5b95945050505050565b6000602082019050818103600083015261136f816111f1565b90505b919050565b600060208201905061138c6000830184611232565b5b92915050565b6000604051905081810181811067ffffffffffffffff821117156113b75760006000fd5b80604052505b919050565b600067ffffffffffffffff8211156113da5760006000fd5b601f19601f83011690506020810190505b919050565b6000815190505b919050565b60008282526020820190505b92915050565b60008190505b92915050565b60006114258261143a565b90505b919050565b600081151590505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b600063ffffffff821690505b919050565b828183376000838301525b505050565b60005b8381101561149b5780820151818401525b60208101905061147f565b838111156114aa576000848401525b505b505050565b6000601f19601f83011690505b919050565b6114cc8161141a565b811415156114da5760006000fd5b5b50565b6114e78161145b565b811415156114f55760006000fd5b5b50565bfea26469706673582212205ac6cb171dfadd74545dfc9285e7a3b351c4cc03caf52b439f58ed407a38155564736f6c63430006060033`

// DeployProof deploys a new Ethereum contract, binding an instance of Proof to it.
func DeployProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Proof, error) {
	parsed, err := abi.JSON(strings.NewReader(ProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Proof{ProofCaller: ProofCaller{contract: contract}, ProofTransactor: ProofTransactor{contract: contract}, ProofFilterer: ProofFilterer{contract: contract}}, nil
}

// Proof is an auto generated Go binding around an Ethereum contract.
type Proof struct {
	ProofCaller     // Read-only binding to the contract
	ProofTransactor // Write-only binding to the contract
	ProofFilterer   // Log filterer for contract events
}

// ProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProofSession struct {
	Contract     *Proof            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProofCallerSession struct {
	Contract *ProofCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProofTransactorSession struct {
	Contract     *ProofTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProofRaw struct {
	Contract *Proof // Generic contract binding to access the raw methods on
}

// ProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProofCallerRaw struct {
	Contract *ProofCaller // Generic read-only contract binding to access the raw methods on
}

// ProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProofTransactorRaw struct {
	Contract *ProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProof creates a new instance of Proof, bound to a specific deployed contract.
func NewProof(address common.Address, backend bind.ContractBackend) (*Proof, error) {
	contract, err := bindProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proof{ProofCaller: ProofCaller{contract: contract}, ProofTransactor: ProofTransactor{contract: contract}, ProofFilterer: ProofFilterer{contract: contract}}, nil
}

// NewProofCaller creates a new read-only instance of Proof, bound to a specific deployed contract.
func NewProofCaller(address common.Address, caller bind.ContractCaller) (*ProofCaller, error) {
	contract, err := bindProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProofCaller{contract: contract}, nil
}

// NewProofTransactor creates a new write-only instance of Proof, bound to a specific deployed contract.
func NewProofTransactor(address common.Address, transactor bind.ContractTransactor) (*ProofTransactor, error) {
	contract, err := bindProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProofTransactor{contract: contract}, nil
}

// NewProofFilterer creates a new log filterer instance of Proof, bound to a specific deployed contract.
func NewProofFilterer(address common.Address, filterer bind.ContractFilterer) (*ProofFilterer, error) {
	contract, err := bindProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProofFilterer{contract: contract}, nil
}

// bindProof binds a generic wrapper to an already deployed contract.
func bindProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proof *ProofRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Proof.Contract.ProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proof *ProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proof.Contract.ProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proof *ProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proof.Contract.ProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proof *ProofCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Proof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proof *ProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proof *ProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proof.Contract.contract.Transact(opts, method, params...)
}

// Challenge is a paid mutator transaction binding the contract method 0x81a53877.
//
// Solidity: function Challenge(_file string) returns()
func (_Proof *ProofTransactor) Challenge(opts *bind.TransactOpts, _file string) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "Challenge", _file)
}

// Challenge is a paid mutator transaction binding the contract method 0x81a53877.
//
// Solidity: function Challenge(_file string) returns()
func (_Proof *ProofSession) Challenge(_file string) (*types.Transaction, error) {
	return _Proof.Contract.Challenge(&_Proof.TransactOpts, _file)
}

// Challenge is a paid mutator transaction binding the contract method 0x81a53877.
//
// Solidity: function Challenge(_file string) returns()
func (_Proof *ProofTransactorSession) Challenge(_file string) (*types.Transaction, error) {
	return _Proof.Contract.Challenge(&_Proof.TransactOpts, _file)
}

// ChallengeList is a paid mutator transaction binding the contract method 0xb8aeb136.
//
// Solidity: function ChallengeList( string) returns(owner address, file string, valid bool)
func (_Proof *ProofTransactor) ChallengeList(opts *bind.TransactOpts, arg0 string) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "ChallengeList", arg0)
}

// ChallengeList is a paid mutator transaction binding the contract method 0xb8aeb136.
//
// Solidity: function ChallengeList( string) returns(owner address, file string, valid bool)
func (_Proof *ProofSession) ChallengeList(arg0 string) (*types.Transaction, error) {
	return _Proof.Contract.ChallengeList(&_Proof.TransactOpts, arg0)
}

// ChallengeList is a paid mutator transaction binding the contract method 0xb8aeb136.
//
// Solidity: function ChallengeList( string) returns(owner address, file string, valid bool)
func (_Proof *ProofTransactorSession) ChallengeList(arg0 string) (*types.Transaction, error) {
	return _Proof.Contract.ChallengeList(&_Proof.TransactOpts, arg0)
}

// FilePledge is a paid mutator transaction binding the contract method 0x60dadb4e.
//
// Solidity: function FilePledge( string) returns(owner address, amount uint32, hash string, valid bool)
func (_Proof *ProofTransactor) FilePledge(opts *bind.TransactOpts, arg0 string) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "FilePledge", arg0)
}

// FilePledge is a paid mutator transaction binding the contract method 0x60dadb4e.
//
// Solidity: function FilePledge( string) returns(owner address, amount uint32, hash string, valid bool)
func (_Proof *ProofSession) FilePledge(arg0 string) (*types.Transaction, error) {
	return _Proof.Contract.FilePledge(&_Proof.TransactOpts, arg0)
}

// FilePledge is a paid mutator transaction binding the contract method 0x60dadb4e.
//
// Solidity: function FilePledge( string) returns(owner address, amount uint32, hash string, valid bool)
func (_Proof *ProofTransactorSession) FilePledge(arg0 string) (*types.Transaction, error) {
	return _Proof.Contract.FilePledge(&_Proof.TransactOpts, arg0)
}

// GetChallengeList is a paid mutator transaction binding the contract method 0xa5990996.
//
// Solidity: function GetChallengeList() returns()
func (_Proof *ProofTransactor) GetChallengeList(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "GetChallengeList")
}

// GetChallengeList is a paid mutator transaction binding the contract method 0xa5990996.
//
// Solidity: function GetChallengeList() returns()
func (_Proof *ProofSession) GetChallengeList() (*types.Transaction, error) {
	return _Proof.Contract.GetChallengeList(&_Proof.TransactOpts)
}

// GetChallengeList is a paid mutator transaction binding the contract method 0xa5990996.
//
// Solidity: function GetChallengeList() returns()
func (_Proof *ProofTransactorSession) GetChallengeList() (*types.Transaction, error) {
	return _Proof.Contract.GetChallengeList(&_Proof.TransactOpts)
}

// GetChallengeResponse is a paid mutator transaction binding the contract method 0x8337846a.
//
// Solidity: function GetChallengeResponse(_file string) returns(address)
func (_Proof *ProofTransactor) GetChallengeResponse(opts *bind.TransactOpts, _file string) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "GetChallengeResponse", _file)
}

// GetChallengeResponse is a paid mutator transaction binding the contract method 0x8337846a.
//
// Solidity: function GetChallengeResponse(_file string) returns(address)
func (_Proof *ProofSession) GetChallengeResponse(_file string) (*types.Transaction, error) {
	return _Proof.Contract.GetChallengeResponse(&_Proof.TransactOpts, _file)
}

// GetChallengeResponse is a paid mutator transaction binding the contract method 0x8337846a.
//
// Solidity: function GetChallengeResponse(_file string) returns(address)
func (_Proof *ProofTransactorSession) GetChallengeResponse(_file string) (*types.Transaction, error) {
	return _Proof.Contract.GetChallengeResponse(&_Proof.TransactOpts, _file)
}

// NodeLeave is a paid mutator transaction binding the contract method 0x0bd8dbe4.
//
// Solidity: function NodeLeave() returns()
func (_Proof *ProofTransactor) NodeLeave(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "NodeLeave")
}

// NodeLeave is a paid mutator transaction binding the contract method 0x0bd8dbe4.
//
// Solidity: function NodeLeave() returns()
func (_Proof *ProofSession) NodeLeave() (*types.Transaction, error) {
	return _Proof.Contract.NodeLeave(&_Proof.TransactOpts)
}

// NodeLeave is a paid mutator transaction binding the contract method 0x0bd8dbe4.
//
// Solidity: function NodeLeave() returns()
func (_Proof *ProofTransactorSession) NodeLeave() (*types.Transaction, error) {
	return _Proof.Contract.NodeLeave(&_Proof.TransactOpts)
}

// NodePledge is a paid mutator transaction binding the contract method 0x5a706771.
//
// Solidity: function NodePledge( address) returns(uint32)
func (_Proof *ProofTransactor) NodePledge(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "NodePledge", arg0)
}

// NodePledge is a paid mutator transaction binding the contract method 0x5a706771.
//
// Solidity: function NodePledge( address) returns(uint32)
func (_Proof *ProofSession) NodePledge(arg0 common.Address) (*types.Transaction, error) {
	return _Proof.Contract.NodePledge(&_Proof.TransactOpts, arg0)
}

// NodePledge is a paid mutator transaction binding the contract method 0x5a706771.
//
// Solidity: function NodePledge( address) returns(uint32)
func (_Proof *ProofTransactorSession) NodePledge(arg0 common.Address) (*types.Transaction, error) {
	return _Proof.Contract.NodePledge(&_Proof.TransactOpts, arg0)
}

// PledgeForFile is a paid mutator transaction binding the contract method 0xe5fe2335.
//
// Solidity: function PledgeForFile(amount uint32, file string) returns()
func (_Proof *ProofTransactor) PledgeForFile(opts *bind.TransactOpts, amount uint32, file string) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "PledgeForFile", amount, file)
}

// PledgeForFile is a paid mutator transaction binding the contract method 0xe5fe2335.
//
// Solidity: function PledgeForFile(amount uint32, file string) returns()
func (_Proof *ProofSession) PledgeForFile(amount uint32, file string) (*types.Transaction, error) {
	return _Proof.Contract.PledgeForFile(&_Proof.TransactOpts, amount, file)
}

// PledgeForFile is a paid mutator transaction binding the contract method 0xe5fe2335.
//
// Solidity: function PledgeForFile(amount uint32, file string) returns()
func (_Proof *ProofTransactorSession) PledgeForFile(amount uint32, file string) (*types.Transaction, error) {
	return _Proof.Contract.PledgeForFile(&_Proof.TransactOpts, amount, file)
}

// PledgeForNode is a paid mutator transaction binding the contract method 0xbde4efe4.
//
// Solidity: function PledgeForNode(amount uint32, node address) returns()
func (_Proof *ProofTransactor) PledgeForNode(opts *bind.TransactOpts, amount uint32, node common.Address) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "PledgeForNode", amount, node)
}

// PledgeForNode is a paid mutator transaction binding the contract method 0xbde4efe4.
//
// Solidity: function PledgeForNode(amount uint32, node address) returns()
func (_Proof *ProofSession) PledgeForNode(amount uint32, node common.Address) (*types.Transaction, error) {
	return _Proof.Contract.PledgeForNode(&_Proof.TransactOpts, amount, node)
}

// PledgeForNode is a paid mutator transaction binding the contract method 0xbde4efe4.
//
// Solidity: function PledgeForNode(amount uint32, node address) returns()
func (_Proof *ProofTransactorSession) PledgeForNode(amount uint32, node common.Address) (*types.Transaction, error) {
	return _Proof.Contract.PledgeForNode(&_Proof.TransactOpts, amount, node)
}

// ProofList is a paid mutator transaction binding the contract method 0xd66a3121.
//
// Solidity: function ProofList( string) returns(owner address, file string, value string, exist uint32)
func (_Proof *ProofTransactor) ProofList(opts *bind.TransactOpts, arg0 string) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "ProofList", arg0)
}

// ProofList is a paid mutator transaction binding the contract method 0xd66a3121.
//
// Solidity: function ProofList( string) returns(owner address, file string, value string, exist uint32)
func (_Proof *ProofSession) ProofList(arg0 string) (*types.Transaction, error) {
	return _Proof.Contract.ProofList(&_Proof.TransactOpts, arg0)
}

// ProofList is a paid mutator transaction binding the contract method 0xd66a3121.
//
// Solidity: function ProofList( string) returns(owner address, file string, value string, exist uint32)
func (_Proof *ProofTransactorSession) ProofList(arg0 string) (*types.Transaction, error) {
	return _Proof.Contract.ProofList(&_Proof.TransactOpts, arg0)
}

// ProofProvide is a paid mutator transaction binding the contract method 0x0d9cb378.
//
// Solidity: function ProofProvide(_file string, _proof string) returns()
func (_Proof *ProofTransactor) ProofProvide(opts *bind.TransactOpts, _file string, _proof string) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "ProofProvide", _file, _proof)
}

// ProofProvide is a paid mutator transaction binding the contract method 0x0d9cb378.
//
// Solidity: function ProofProvide(_file string, _proof string) returns()
func (_Proof *ProofSession) ProofProvide(_file string, _proof string) (*types.Transaction, error) {
	return _Proof.Contract.ProofProvide(&_Proof.TransactOpts, _file, _proof)
}

// ProofProvide is a paid mutator transaction binding the contract method 0x0d9cb378.
//
// Solidity: function ProofProvide(_file string, _proof string) returns()
func (_Proof *ProofTransactorSession) ProofProvide(_file string, _proof string) (*types.Transaction, error) {
	return _Proof.Contract.ProofProvide(&_Proof.TransactOpts, _file, _proof)
}

// Punish is a paid mutator transaction binding the contract method 0x11ec5d80.
//
// Solidity: function Punish(node address, file address) returns()
func (_Proof *ProofTransactor) Punish(opts *bind.TransactOpts, node common.Address, file common.Address) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "Punish", node, file)
}

// Punish is a paid mutator transaction binding the contract method 0x11ec5d80.
//
// Solidity: function Punish(node address, file address) returns()
func (_Proof *ProofSession) Punish(node common.Address, file common.Address) (*types.Transaction, error) {
	return _Proof.Contract.Punish(&_Proof.TransactOpts, node, file)
}

// Punish is a paid mutator transaction binding the contract method 0x11ec5d80.
//
// Solidity: function Punish(node address, file address) returns()
func (_Proof *ProofTransactorSession) Punish(node common.Address, file common.Address) (*types.Transaction, error) {
	return _Proof.Contract.Punish(&_Proof.TransactOpts, node, file)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Proof *ProofTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Proof *ProofSession) Destroy() (*types.Transaction, error) {
	return _Proof.Contract.Destroy(&_Proof.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Proof *ProofTransactorSession) Destroy() (*types.Transaction, error) {
	return _Proof.Contract.Destroy(&_Proof.TransactOpts)
}

// Owner is a paid mutator transaction binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Proof *ProofTransactor) Owner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "owner")
}

// Owner is a paid mutator transaction binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Proof *ProofSession) Owner() (*types.Transaction, error) {
	return _Proof.Contract.Owner(&_Proof.TransactOpts)
}

// Owner is a paid mutator transaction binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() returns(address)
func (_Proof *ProofTransactorSession) Owner() (*types.Transaction, error) {
	return _Proof.Contract.Owner(&_Proof.TransactOpts)
}
