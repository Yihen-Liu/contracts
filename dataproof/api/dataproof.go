// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proof

import (
	"math/big"
	"strings"
)

// ProofABI is the input ABI used to generate the binding from.
const ProofABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"file\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_nonce\",\"type\":\"bytes32\"}],\"name\":\"Challenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"ChallengeList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"file\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"FilePledge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetChallengeList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NodeLeave\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NodePledge\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"file\",\"type\":\"bytes32\"}],\"name\":\"PledgeForFile\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"PledgeForNode\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ProofList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"file\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"},{\"internalType\":\"enumProofDataPossession.ProofStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_file\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_nonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_proof\",\"type\":\"bytes32\"}],\"name\":\"ProofProvide\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"file\",\"type\":\"address\"}],\"name\":\"Punish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ProofBin is the compiled bytecode used for deploying new contracts.
const ProofBin = `60806040526103e8600460146101000a81548163ffffffff021916908363ffffffff1602179055503480156100345760006000fd5b505b33600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b61007d565b6110b68061008c6000396000f3fe6080604052600436106100c65760003560e01c80638da5cb5b1161007f578063b5453e3411610059578063b5453e3414610255578063b8aeb13614610271578063bde4efe4146102b2578063f8b21ae6146102ce576100c6565b80638da5cb5b146101e75780639bc49c4914610213578063a59909961461023d576100c6565b8063068fc622146100cc5780630bd8dbe41461010e5780630cc334981461012657806311ec5d80146101675780635a7067711461019157806383197ef0146101cf576100c6565b60006000fd5b3480156100d95760006000fd5b506100f460048036038101906100ef9190610bbe565b6102f8565b604051610105959493929190610dd1565b60405180910390f35b34801561011b5760006000fd5b5061012461039d565b005b3480156101335760006000fd5b5061014e60048036038101906101499190610abe565b6104a2565b60405161015e9493929190610e25565b60405180910390f35b3480156101745760006000fd5b5061018f600480360381019061018a9190610a7f565b610515565b005b34801561019e5760006000fd5b506101b960048036038101906101b49190610a54565b61051a565b6040516101c69190610e8c565b60405180910390f35b3480156101dc5760006000fd5b506101e5610542565b005b3480156101f45760006000fd5b506101fd6105ba565b60405161020a9190610d6f565b60405180910390f35b3480156102205760006000fd5b5061023b60048036038101906102369190610ae9565b6105e0565b005b34801561024a5760006000fd5b506102536106e1565b005b61026f600480360381019061026a9190610c55565b6106e4565b005b34801561027e5760006000fd5b5061029960048036038101906102949190610b7a565b610856565b6040516102a99493929190610d8b565b60405180910390f35b6102cc60048036038101906102c79190610c16565b6108d2565b005b3480156102db5760006000fd5b506102f660048036038101906102f19190610b28565b610999565b005b6003600050828051602081018201805184825260208301602085012081835280955050505050506000508181548110151561032f57fe5b906000526020600020906005020160005b91509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160005054908060020160005054908060030160005054908060040160009054906101000a900460ff16905085565b3373ffffffffffffffffffffffffffffffffffffffff166108fc600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff1663ffffffff169081150290604051600060405180830381858888f1935050505015801561043c573d600060003e3d6000fd5b506000600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff1602179055505b565b60016000506020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900463ffffffff16908060010160005054908060020160009054906101000a900460ff16905084565b5b5050565b600060005060205280600052604060002060009150909054906101000a900463ffffffff1681565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561059f5760006000fd5b3373ffffffffffffffffffffffffffffffffffffffff16ff5b565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606060405180608001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200184600019168152602001836000191681526020016001151581526020015060026000508260405161063a9190610d57565b908152602001604051809103902060005060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101600050906000191690556040820151816002016000509060001916905560608201518160030160006101000a81548160ff021916908315150217905550905050505b5050565b5b565b8163ffffffff163373ffffffffffffffffffffffffffffffffffffffff1631111515610745576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073c90610e6b565b60405180910390fd5b60405180608001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018363ffffffff168152602001826000191681526020016001151581526020015060016000506000836000191660001916815260200190815260200160002060005060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548163ffffffff021916908363ffffffff1602179055506040820151816001016000509060001916905560608201518160020160006101000a81548160ff0219169083151502179055509050505b5050565b6002600050818051602081018201805184825260208301602085012081835280955050505050506000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160005054908060020160005054908060030160009054906101000a900460ff16905084565b8163ffffffff163373ffffffffffffffffffffffffffffffffffffffff1631111515610933576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161092a90610e6b565b60405180910390fd5b81600060005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff1602179055505b5050565b5b5050505661107f565b6000813590506109b281611013565b5b92915050565b6000813590506109c88161102e565b5b92915050565b600082601f83011215156109e35760006000fd5b81356109f66109f182610ed7565b610ea8565b91508082526020830160208301858383011115610a135760006000fd5b610a1e838284610fbe565b5050505b92915050565b600081359050610a3781611049565b5b92915050565b600081359050610a4d81611064565b5b92915050565b600060208284031215610a675760006000fd5b6000610a75848285016109a3565b9150505b92915050565b6000600060408385031215610a945760006000fd5b6000610aa2858286016109a3565b9250506020610ab3858286016109a3565b9150505b9250929050565b600060208284031215610ad15760006000fd5b6000610adf848285016109b9565b9150505b92915050565b6000600060408385031215610afe5760006000fd5b6000610b0c858286016109b9565b9250506020610b1d858286016109b9565b9150505b9250929050565b60006000600060608486031215610b3f5760006000fd5b6000610b4d868287016109b9565b9350506020610b5e868287016109b9565b9250506040610b6f868287016109b9565b9150505b9250925092565b600060208284031215610b8d5760006000fd5b600082013567ffffffffffffffff811115610ba85760006000fd5b610bb4848285016109cf565b9150505b92915050565b6000600060408385031215610bd35760006000fd5b600083013567ffffffffffffffff811115610bee5760006000fd5b610bfa858286016109cf565b9250506020610c0b85828601610a28565b9150505b9250929050565b6000600060408385031215610c2b5760006000fd5b6000610c3985828601610a3e565b9250506020610c4a858286016109a3565b9150505b9250929050565b6000600060408385031215610c6a5760006000fd5b6000610c7885828601610a3e565b9250506020610c89858286016109b9565b9150505b9250929050565b610c9d81610f2f565b82525b5050565b610cad81610f42565b82525b5050565b610cbd81610f4f565b82525b5050565b610ccd81610fab565b82525b5050565b6000610cdf82610f05565b610ce98185610f23565b9350610cf9818560208601610fce565b8084019150505b92915050565b6000610d13601583610f11565b91507f62616c616e6365206973206e6f7420656e6f756768000000000000000000000060008301526020820190505b919050565b610d5081610f9a565b82525b5050565b6000610d638284610cd4565b91508190505b92915050565b6000602082019050610d846000830184610c94565b5b92915050565b6000608082019050610da06000830187610c94565b610dad6020830186610cb4565b610dba6040830185610cb4565b610dc76060830184610ca4565b5b95945050505050565b600060a082019050610de66000830188610c94565b610df36020830187610cb4565b610e006040830186610cb4565b610e0d6060830185610cb4565b610e1a6080830184610cc4565b5b9695505050505050565b6000608082019050610e3a6000830187610c94565b610e476020830186610d47565b610e546040830185610cb4565b610e616060830184610ca4565b5b95945050505050565b60006020820190508181036000830152610e8481610d06565b90505b919050565b6000602082019050610ea16000830184610d47565b5b92915050565b6000604051905081810181811067ffffffffffffffff82111715610ecc5760006000fd5b80604052505b919050565b600067ffffffffffffffff821115610eef5760006000fd5b601f19601f83011690506020810190505b919050565b6000815190505b919050565b60008282526020820190505b92915050565b60008190505b92915050565b6000610f3a82610f6e565b90505b919050565b600081151590505b919050565b60008190505b919050565b6000819050610f6882611003565b5b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b600063ffffffff821690505b919050565b6000610fb682610f5a565b90505b919050565b828183376000838301525b505050565b60005b83811015610fed5780820151818401525b602081019050610fd1565b83811115610ffc576000848401525b505b505050565b60048110151561100f57fe5b5b50565b61101c81610f2f565b8114151561102a5760006000fd5b5b50565b61103781610f4f565b811415156110455760006000fd5b5b50565b61105281610f8f565b811415156110605760006000fd5b5b50565b61106d81610f9a565b8114151561107b5760006000fd5b5b50565bfea26469706673582212204f66802a2069bc722290e5813526cad08c5dfac8e4c04489e006be7cf8a442e764736f6c63430006060033`

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

// Challenge is a paid mutator transaction binding the contract method 0x9bc49c49.
//
// Solidity: function Challenge(file bytes32, _nonce bytes32) returns()
func (_Proof *ProofTransactor) Challenge(opts *bind.TransactOpts, file [32]byte, _nonce [32]byte) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "Challenge", file, _nonce)
}

// Challenge is a paid mutator transaction binding the contract method 0x9bc49c49.
//
// Solidity: function Challenge(file bytes32, _nonce bytes32) returns()
func (_Proof *ProofSession) Challenge(file [32]byte, _nonce [32]byte) (*types.Transaction, error) {
	return _Proof.Contract.Challenge(&_Proof.TransactOpts, file, _nonce)
}

// Challenge is a paid mutator transaction binding the contract method 0x9bc49c49.
//
// Solidity: function Challenge(file bytes32, _nonce bytes32) returns()
func (_Proof *ProofTransactorSession) Challenge(file [32]byte, _nonce [32]byte) (*types.Transaction, error) {
	return _Proof.Contract.Challenge(&_Proof.TransactOpts, file, _nonce)
}

// ChallengeList is a paid mutator transaction binding the contract method 0xb8aeb136.
//
// Solidity: function ChallengeList( string) returns(owner address, file bytes32, nonce bytes32, valid bool)
func (_Proof *ProofTransactor) ChallengeList(opts *bind.TransactOpts, arg0 string) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "ChallengeList", arg0)
}

// ChallengeList is a paid mutator transaction binding the contract method 0xb8aeb136.
//
// Solidity: function ChallengeList( string) returns(owner address, file bytes32, nonce bytes32, valid bool)
func (_Proof *ProofSession) ChallengeList(arg0 string) (*types.Transaction, error) {
	return _Proof.Contract.ChallengeList(&_Proof.TransactOpts, arg0)
}

// ChallengeList is a paid mutator transaction binding the contract method 0xb8aeb136.
//
// Solidity: function ChallengeList( string) returns(owner address, file bytes32, nonce bytes32, valid bool)
func (_Proof *ProofTransactorSession) ChallengeList(arg0 string) (*types.Transaction, error) {
	return _Proof.Contract.ChallengeList(&_Proof.TransactOpts, arg0)
}

// FilePledge is a paid mutator transaction binding the contract method 0x0cc33498.
//
// Solidity: function FilePledge( bytes32) returns(owner address, amount uint32, hash bytes32, valid bool)
func (_Proof *ProofTransactor) FilePledge(opts *bind.TransactOpts, arg0 [32]byte) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "FilePledge", arg0)
}

// FilePledge is a paid mutator transaction binding the contract method 0x0cc33498.
//
// Solidity: function FilePledge( bytes32) returns(owner address, amount uint32, hash bytes32, valid bool)
func (_Proof *ProofSession) FilePledge(arg0 [32]byte) (*types.Transaction, error) {
	return _Proof.Contract.FilePledge(&_Proof.TransactOpts, arg0)
}

// FilePledge is a paid mutator transaction binding the contract method 0x0cc33498.
//
// Solidity: function FilePledge( bytes32) returns(owner address, amount uint32, hash bytes32, valid bool)
func (_Proof *ProofTransactorSession) FilePledge(arg0 [32]byte) (*types.Transaction, error) {
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

// PledgeForFile is a paid mutator transaction binding the contract method 0xb5453e34.
//
// Solidity: function PledgeForFile(amount uint32, file bytes32) returns()
func (_Proof *ProofTransactor) PledgeForFile(opts *bind.TransactOpts, amount uint32, file [32]byte) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "PledgeForFile", amount, file)
}

// PledgeForFile is a paid mutator transaction binding the contract method 0xb5453e34.
//
// Solidity: function PledgeForFile(amount uint32, file bytes32) returns()
func (_Proof *ProofSession) PledgeForFile(amount uint32, file [32]byte) (*types.Transaction, error) {
	return _Proof.Contract.PledgeForFile(&_Proof.TransactOpts, amount, file)
}

// PledgeForFile is a paid mutator transaction binding the contract method 0xb5453e34.
//
// Solidity: function PledgeForFile(amount uint32, file bytes32) returns()
func (_Proof *ProofTransactorSession) PledgeForFile(amount uint32, file [32]byte) (*types.Transaction, error) {
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

// ProofList is a paid mutator transaction binding the contract method 0x068fc622.
//
// Solidity: function ProofList( string,  uint256) returns(owner address, file bytes32, nonce bytes32, value bytes32, status uint8)
func (_Proof *ProofTransactor) ProofList(opts *bind.TransactOpts, arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "ProofList", arg0, arg1)
}

// ProofList is a paid mutator transaction binding the contract method 0x068fc622.
//
// Solidity: function ProofList( string,  uint256) returns(owner address, file bytes32, nonce bytes32, value bytes32, status uint8)
func (_Proof *ProofSession) ProofList(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Proof.Contract.ProofList(&_Proof.TransactOpts, arg0, arg1)
}

// ProofList is a paid mutator transaction binding the contract method 0x068fc622.
//
// Solidity: function ProofList( string,  uint256) returns(owner address, file bytes32, nonce bytes32, value bytes32, status uint8)
func (_Proof *ProofTransactorSession) ProofList(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Proof.Contract.ProofList(&_Proof.TransactOpts, arg0, arg1)
}

// ProofProvide is a paid mutator transaction binding the contract method 0xf8b21ae6.
//
// Solidity: function ProofProvide(_file bytes32, _nonce bytes32, _proof bytes32) returns()
func (_Proof *ProofTransactor) ProofProvide(opts *bind.TransactOpts, _file [32]byte, _nonce [32]byte, _proof [32]byte) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "ProofProvide", _file, _nonce, _proof)
}

// ProofProvide is a paid mutator transaction binding the contract method 0xf8b21ae6.
//
// Solidity: function ProofProvide(_file bytes32, _nonce bytes32, _proof bytes32) returns()
func (_Proof *ProofSession) ProofProvide(_file [32]byte, _nonce [32]byte, _proof [32]byte) (*types.Transaction, error) {
	return _Proof.Contract.ProofProvide(&_Proof.TransactOpts, _file, _nonce, _proof)
}

// ProofProvide is a paid mutator transaction binding the contract method 0xf8b21ae6.
//
// Solidity: function ProofProvide(_file bytes32, _nonce bytes32, _proof bytes32) returns()
func (_Proof *ProofTransactorSession) ProofProvide(_file [32]byte, _nonce [32]byte, _proof [32]byte) (*types.Transaction, error) {
	return _Proof.Contract.ProofProvide(&_Proof.TransactOpts, _file, _nonce, _proof)
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
