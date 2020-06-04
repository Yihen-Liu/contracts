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
const ProofABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"}],\"name\":\"Challenge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"ChallengeList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"file\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"FilePledge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetChallengeList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"}],\"name\":\"GetChallengeResponse\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NodeLeave\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NodePledge\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"file\",\"type\":\"string\"}],\"name\":\"PledgeForFile\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"PledgeForNode\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"ProofList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"file\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"exist\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_proof\",\"type\":\"string\"}],\"name\":\"ProofProvide\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"file\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"name\":\"PunishOrReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ProofBin is the compiled bytecode used for deploying new contracts.
const ProofBin = `6080604052600061200460146101000a81548163ffffffff021916908363ffffffff1602179055503480156100345760006000fd5b505b3361200460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b61007e565b611bbf8061008d6000396000f3fe6080604052600436106100e15760003560e01c80638da5cb5b1161007f578063b8aeb13611610059578063b8aeb136146102be578063bde4efe4146102fe578063d66a31211461031a578063e5fe23351461035b576100e1565b80638da5cb5b1461023c578063a2047fff14610268578063a599099614610292576100e1565b806360dadb4e116100bb57806360dadb4e1461016757806381a53877146101a857806383197ef0146101e65780638337846a146101fe576100e1565b80630bd8dbe4146100e75780630d9cb378146100ff5780635a70677114610129576100e1565b60006000fd5b3480156100f45760006000fd5b506100fd610377565b005b34801561010c5760006000fd5b5061012760048036038101906101229190611687565b61047c565b005b3480156101365760006000fd5b50610151600480360381019061014c91906115ad565b610622565b60405161015e91906119bd565b60405180910390f35b3480156101745760006000fd5b5061018f600480360381019061018a9190611643565b61064a565b60405161019f9493929190611933565b60405180910390f35b3480156101b55760006000fd5b506101d060048036038101906101cb9190611643565b61076b565b6040516101dd9190611980565b60405180910390f35b3480156101f35760006000fd5b506101fc610901565b005b34801561020b5760006000fd5b5061022660048036038101906102219190611643565b61097a565b6040516102339190611884565b60405180910390f35b3480156102495760006000fd5b50610252610a1f565b60405161025f9190611884565b60405180910390f35b3480156102755760006000fd5b50610290600480360381019061028b91906115d8565b610a46565b005b34801561029f5760006000fd5b506102a8610db6565b6040516102b59190611884565b60405180910390f35b3480156102cb5760006000fd5b506102e660048036038101906102e19190611643565b610e36565b6040516102f5939291906118a0565b60405180910390f35b610318600480360381019061031391906116f8565b610f41565b005b3480156103275760006000fd5b50610342600480360381019061033d9190611643565b611008565b60405161035294939291906118df565b60405180910390f35b61037560048036038101906103709190611737565b6111b7565b005b3373ffffffffffffffffffffffffffffffffffffffff166108fc600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900463ffffffff1663ffffffff169081150290604051600060405180830381858888f19350505050158015610416573d600060003e3d6000fd5b506000600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff1602179055505b565b6001600360005083604051610491919061186c565b908152602001604051809103902060005060030160009054906101000a900463ffffffff1663ffffffff1614156104c75761061e565b60006002600050836040516104dc919061186c565b908152602001604051809103902060005060020160006101000a81548160ff02191690831515021790555060405180608001604052803373ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001600163ffffffff1681526020015060036000508360405161055b919061186c565b908152602001604051809103902060005060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160005090805190602001906105d2929190611400565b5060408201518160020160005090805190602001906105f2929190611400565b5060608201518160030160006101000a81548163ffffffff021916908363ffffffff1602179055509050505b5050565b600060005060205280600052604060002060009150909054906101000a900463ffffffff1681565b6001600050818051602081018201805184825260208301602085012081835280955050505050506000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900463ffffffff1690806001016000508054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561074e5780601f106107235761010080835404028352916020019161074e565b820191906000526020600020905b81548152906001019060200180831161073157829003601f168201915b5050505050908060020160009054906101000a900460ff16905084565b6000611000600460005050611000111561078857600090506108fc565b6040518060400160405280838152602001600063ffffffff16815260200150600460005042611000811015156107ba57fe5b90906002020160005b5060008201518160000160005090805190602001906107e3929190611400565b5060208201518160010160006101000a81548163ffffffff021916908363ffffffff16021790555090505060405180606001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200183815260200160011515815260200150600260005083604051610858919061186c565b908152602001604051809103902060005060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160005090805190602001906108cf929190611400565b5060408201518160020160006101000a81548160ff021916908315150217905550905050600190506108fc565b919050565b61200460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561095f5760006000fd5b3373ffffffffffffffffffffffffffffffffffffffff16ff5b565b60006000600360005083604051610991919061186c565b908152602001604051809103902060005060030160009054906101000a900463ffffffff1663ffffffff1614156109cb5760009050610a1a565b6003600050826040516109de919061186c565b908152602001604051809103902060005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050610a1a565b919050565b61200460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600115158115151415610c66576000600090505b6004600050506110008163ffffffff161015610c6057610b418360046000508363ffffffff1661100081101515610a8d57fe5b90906002020160005b506000016000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b315780601f10610b0657610100808354040283529160200191610b31565b820191906000526020600020905b815481529060010190602001808311610b1457829003601f168201915b505050505061133563ffffffff16565b15610c52578373ffffffffffffffffffffffffffffffffffffffff166108fc6103e863ffffffff169081150290604051600060405180830381858888f19350505050158015610b95573d600060003e3d6000fd5b5060046000508163ffffffff1661100081101515610baf57fe5b90906002020160005b600082016000610bc89190611485565b6001820160006101000a81549063ffffffff02191690555050600260005083604051610bf4919061186c565b908152602001604051809103902060006000820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182016000610c3b9190611485565b6002820160006101000a81549060ff021916905550505b5b8080600101915050610a5a565b50610db0565b6000600090505b6004600050506110008163ffffffff161015610dae57610d548360046000508363ffffffff1661100081101515610ca057fe5b90906002020160005b506000016000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d445780601f10610d1957610100808354040283529160200191610d44565b820191906000526020600020905b815481529060010190602001808311610d2757829003601f168201915b505050505061133563ffffffff16565b15610da057600060046000508263ffffffff1661100081101515610d7457fe5b90906002020160005b5060010160006101000a81548163ffffffff021916908363ffffffff1602179055505b5b8080600101915050610c6d565b505b5b505050565b60006000600090505b6004600050506110008163ffffffff161015610e2957600060046000508263ffffffff1661100081101515610df057fe5b90906002020160005b5060010160009054906101000a900463ffffffff1663ffffffff161415610e1b575b5b8080600101915050610dbf565b5060009050610e33565b90565b6002600050818051602081018201805184825260208301602085012081835280955050505050506000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001016000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f245780601f10610ef957610100808354040283529160200191610f24565b820191906000526020600020905b815481529060010190602001808311610f0757829003601f168201915b5050505050908060020160009054906101000a900460ff16905083565b8163ffffffff163373ffffffffffffffffffffffffffffffffffffffff1631111515610fa2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f999061199c565b60405180910390fd5b81600060005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908363ffffffff1602179055505b5050565b6003600050818051602081018201805184825260208301602085012081835280955050505050506000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001016000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110f65780601f106110cb576101008083540402835291602001916110f6565b820191906000526020600020905b8154815290600101906020018083116110d957829003601f168201915b505050505090806002016000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111975780601f1061116c57610100808354040283529160200191611197565b820191906000526020600020905b81548152906001019060200180831161117a57829003601f168201915b5050505050908060030160009054906101000a900463ffffffff16905084565b8163ffffffff163373ffffffffffffffffffffffffffffffffffffffff1631111515611218576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161120f9061199c565b60405180910390fd5b60405180608001604052803373ffffffffffffffffffffffffffffffffffffffff1681526020018363ffffffff1681526020018281526020016001151581526020015060016000508260405161126e919061186c565b908152602001604051809103902060005060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548163ffffffff021916908363ffffffff160217905550604082015181600101600050908051906020019061130c929190611400565b5060608201518160020160006101000a81548160ff0219169083151502179055509050505b5050565b60008151835114151561134b57600090506113fa565b6000600090505b83518110156113f057828181518110151561136957fe5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191684828151811015156113a457fe5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161415156113e25760009150506113fa565b5b8080600101915050611352565b50600190506113fa565b92915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061144157805160ff1916838001178555611474565b82800160010185558215611474579182015b828111156114735782518260005090905591602001919060010190611453565b5b50905061148191906114cd565b5090565b50805460018160011615610100020316600290046000825580601f106114ab57506114ca565b601f0160209004906000526020600020908101906114c991906114cd565b5b50565b6114f591906114d7565b808211156114f157600081815060009055506001016114d7565b5090565b9056611b88565b60008135905061150b81611b1c565b5b92915050565b60008135905061152181611b37565b5b92915050565b60008135905061153781611b52565b5b92915050565b600082601f83011215156115525760006000fd5b813561156561156082611a08565b6119d9565b915080825260208301602083018583830111156115825760006000fd5b61158d838284611ac5565b5050505b92915050565b6000813590506115a681611b6d565b5b92915050565b6000602082840312156115c05760006000fd5b60006115ce848285016114fc565b9150505b92915050565b600060006000606084860312156115ef5760006000fd5b60006115fd86828701611512565b935050602084013567ffffffffffffffff81111561161b5760006000fd5b6116278682870161153e565b925050604061163886828701611528565b9150505b9250925092565b6000602082840312156116565760006000fd5b600082013567ffffffffffffffff8111156116715760006000fd5b61167d8482850161153e565b9150505b92915050565b600060006040838503121561169c5760006000fd5b600083013567ffffffffffffffff8111156116b75760006000fd5b6116c38582860161153e565b925050602083013567ffffffffffffffff8111156116e15760006000fd5b6116ed8582860161153e565b9150505b9250929050565b600060006040838503121561170d5760006000fd5b600061171b85828601611597565b925050602061172c858286016114fc565b9150505b9250929050565b600060006040838503121561174c5760006000fd5b600061175a85828601611597565b925050602083013567ffffffffffffffff8111156117785760006000fd5b6117848582860161153e565b9150505b9250929050565b61179881611a60565b82525b5050565b6117a881611a86565b82525b5050565b60006117ba82611a36565b6117c48185611a42565b93506117d4818560208601611ad5565b6117dd81611b0a565b84019150505b92915050565b60006117f482611a36565b6117fe8185611a54565b935061180e818560208601611ad5565b8084019150505b92915050565b6000611828601583611a42565b91507f62616c616e6365206973206e6f7420656e6f756768000000000000000000000060008301526020820190505b919050565b61186581611ab4565b82525b5050565b600061187882846117e9565b91508190505b92915050565b6000602082019050611899600083018461178f565b5b92915050565b60006060820190506118b5600083018661178f565b81810360208301526118c781856117af565b90506118d6604083018461179f565b5b949350505050565b60006080820190506118f4600083018761178f565b818103602083015261190681866117af565b9050818103604083015261191a81856117af565b9050611929606083018461185c565b5b95945050505050565b6000608082019050611948600083018761178f565b611955602083018661185c565b818103604083015261196781856117af565b9050611976606083018461179f565b5b95945050505050565b6000602082019050611995600083018461179f565b5b92915050565b600060208201905081810360008301526119b58161181b565b90505b919050565b60006020820190506119d2600083018461185c565b5b92915050565b6000604051905081810181811067ffffffffffffffff821117156119fd5760006000fd5b80604052505b919050565b600067ffffffffffffffff821115611a205760006000fd5b601f19601f83011690506020810190505b919050565b6000815190505b919050565b60008282526020820190505b92915050565b60008190505b92915050565b6000611a6b82611a93565b90505b919050565b6000611a7e82611a93565b90505b919050565b600081151590505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b600063ffffffff821690505b919050565b828183376000838301525b505050565b60005b83811015611af45780820151818401525b602081019050611ad8565b83811115611b03576000848401525b505b505050565b6000601f19601f83011690505b919050565b611b2581611a60565b81141515611b335760006000fd5b5b50565b611b4081611a73565b81141515611b4e5760006000fd5b5b50565b611b5b81611a86565b81141515611b695760006000fd5b5b50565b611b7681611ab4565b81141515611b845760006000fd5b5b50565bfea264697066735822122089eb5bcd3ff2c1929d0893b7f9e68b66bd6be8e5ec75b47cf5abe3f562e422a864736f6c63430006060033`

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
// Solidity: function Challenge(_file string) returns(bool)
func (_Proof *ProofTransactor) Challenge(opts *bind.TransactOpts, _file string) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "Challenge", _file)
}

// Challenge is a paid mutator transaction binding the contract method 0x81a53877.
//
// Solidity: function Challenge(_file string) returns(bool)
func (_Proof *ProofSession) Challenge(_file string) (*types.Transaction, error) {
	return _Proof.Contract.Challenge(&_Proof.TransactOpts, _file)
}

// Challenge is a paid mutator transaction binding the contract method 0x81a53877.
//
// Solidity: function Challenge(_file string) returns(bool)
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
// Solidity: function GetChallengeList() returns(address)
func (_Proof *ProofTransactor) GetChallengeList(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "GetChallengeList")
}

// GetChallengeList is a paid mutator transaction binding the contract method 0xa5990996.
//
// Solidity: function GetChallengeList() returns(address)
func (_Proof *ProofSession) GetChallengeList() (*types.Transaction, error) {
	return _Proof.Contract.GetChallengeList(&_Proof.TransactOpts)
}

// GetChallengeList is a paid mutator transaction binding the contract method 0xa5990996.
//
// Solidity: function GetChallengeList() returns(address)
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

// PunishOrReward is a paid mutator transaction binding the contract method 0xa2047fff.
//
// Solidity: function PunishOrReward(node address, file string, valid bool) returns()
func (_Proof *ProofTransactor) PunishOrReward(opts *bind.TransactOpts, node common.Address, file string, valid bool) (*types.Transaction, error) {
	return _Proof.contract.Transact(opts, "PunishOrReward", node, file, valid)
}

// PunishOrReward is a paid mutator transaction binding the contract method 0xa2047fff.
//
// Solidity: function PunishOrReward(node address, file string, valid bool) returns()
func (_Proof *ProofSession) PunishOrReward(node common.Address, file string, valid bool) (*types.Transaction, error) {
	return _Proof.Contract.PunishOrReward(&_Proof.TransactOpts, node, file, valid)
}

// PunishOrReward is a paid mutator transaction binding the contract method 0xa2047fff.
//
// Solidity: function PunishOrReward(node address, file string, valid bool) returns()
func (_Proof *ProofTransactorSession) PunishOrReward(node common.Address, file string, valid bool) (*types.Transaction, error) {
	return _Proof.Contract.PunishOrReward(&_Proof.TransactOpts, node, file, valid)
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
