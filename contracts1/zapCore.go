// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts1

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72305820b0b25139b75e676b66895a7d1fb0a1f523d8c20a178651d70497b0640c4406c364736f6c634300050a0032`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// ZapABI is the input ABI used to generate the binding from.
const ZapABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_propNewZapAddress\",\"type\":\"address\"}],\"name\":\"proposeFork\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"requestStakingWithdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_c_sapi\",\"type\":\"string\"},{\"name\":\"_c_symbol\",\"type\":\"string\"},{\"name\":\"_granularity\",\"type\":\"uint256\"},{\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"requestData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"tallyVotes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nonce\",\"type\":\"string\"},{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"submitMiningSolution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"name\":\"_minerIndex\",\"type\":\"uint256\"}],\"name\":\"beginDispute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"theLazyCoon\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"name\":\"_supportsDispute\",\"type\":\"bool\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ZapBin is the compiled bytecode used for deploying new contracts.
const ZapBin = `0x608060405234801561001057600080fd5b50610c0b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806368c180d511610097578063b079f64a11610066578063b079f64a1461037b578063bed9d861146103a7578063c9d27afe146103af578063f2fde38b146103d4576100f5565b806368c180d51461028d578063752d49a1146103035780638581af1914610326578063a9059cbb1461034f576100f5565b806326b7d9f6116100d357806326b7d9f61461017a57806328449c3a146101a05780633fff2816146101a85780634d318b0e14610270576100f5565b8063095ea7b3146100fa5780630d2d76a21461013a57806323b872dd14610144575b600080fd5b6101266004803603604081101561011057600080fd5b506001600160a01b0381351690602001356103fa565b604080519115158252519081900360200190f35b610142610497565b005b6101266004803603606081101561015a57600080fd5b506001600160a01b03813581169160208101359091169060400135610501565b6101426004803603602081101561019057600080fd5b50356001600160a01b03166105a7565b610142610621565b610142600480360360808110156101be57600080fd5b8101906020810181356401000000008111156101d957600080fd5b8201836020820111156101eb57600080fd5b8035906020019184600183028401116401000000008311171561020d57600080fd5b91939092909160208101903564010000000081111561022b57600080fd5b82018360208201111561023d57600080fd5b8035906020019184600183028401116401000000008311171561025f57600080fd5b919350915080359060200135610671565b6101426004803603602081101561028657600080fd5b5035610753565b610142600480360360608110156102a357600080fd5b8101906020810181356401000000008111156102be57600080fd5b8201836020820111156102d057600080fd5b803590602001918460018302840111640100000000831117156102f257600080fd5b9193509150803590602001356107aa565b6101426004803603604081101561031957600080fd5b508035906020013561085e565b6101426004803603606081101561033c57600080fd5b50803590602081013590604001356108d8565b6101266004803603604081101561036557600080fd5b506001600160a01b03813516906020013561095a565b6101426004803603604081101561039157600080fd5b506001600160a01b0381351690602001356109c4565b610142610a2a565b610142600480360360408110156103c557600080fd5b50803590602001351515610a7a565b610142600480360360208110156103ea57600080fd5b50356001600160a01b0316610ad9565b60408051634286e61960e11b81526000600482018190526001600160a01b038516602483015260448201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163850dcc32916064808301926020929190829003018186803b15801561046457600080fd5b505af4158015610478573d6000803e3d6000fd5b505050506040513d602081101561048e57600080fd5b50519392505050565b6040805163410516b360e11b8152600060048201819052915173__$799602413129f49037f52758954cf5aa52$__9263820a2d669260248082019391829003018186803b1580156104e757600080fd5b505af41580156104fb573d6000803e3d6000fd5b50505050565b6040805163ca50189960e01b81526000600482018190526001600160a01b0380871660248401528516604483015260648201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163ca501899916084808301926020929190829003018186803b15801561057357600080fd5b505af4158015610587573d6000803e3d6000fd5b505050506040513d602081101561059d57600080fd5b5051949350505050565b6040805163694bf49f60e01b81526000600482018190526001600160a01b0384166024830152915173__$5c50ea773c3f1223822c80c8edbce14ea7$__9263694bf49f9260448082019391829003018186803b15801561060657600080fd5b505af415801561061a573d6000803e3d6000fd5b5050505050565b60408051633273d79360e21b8152600060048201819052915173__$799602413129f49037f52758954cf5aa52$__9263c9cf5e4c9260248082019391829003018186803b1580156104e757600080fd5b600073__$d2345629a19a8ba674464680be5d957238$__6385997bf690918888888888886040518863ffffffff1660e01b81526004018088815260200180602001806020018581526020018481526020018381038352898982818152602001925080828437600083820152601f01601f191690910184810383528781526020019050878780828437600081840152601f19601f820116905080830192505050995050505050505050505060006040518083038186803b15801561073357600080fd5b505af4158015610747573d6000803e3d6000fd5b50505050505050505050565b6040805163def6fac760e01b815260006004820181905260248201849052915173__$5c50ea773c3f1223822c80c8edbce14ea7$__9263def6fac79260448082019391829003018186803b15801561060657600080fd5b600073__$d2345629a19a8ba674464680be5d957238$__63a098b5b49091868686866040518663ffffffff1660e01b815260040180868152602001806020018481526020018381526020018281038252868682818152602001925080828437600081840152601f19601f820116905080830192505050965050505050505060006040518083038186803b15801561084057600080fd5b505af4158015610854573d6000803e3d6000fd5b5050505050505050565b604080516302e8f21b60e01b81526000600482018190526024820185905260448201849052915173__$d2345629a19a8ba674464680be5d957238$__926302e8f21b9260648082019391829003018186803b1580156108bc57600080fd5b505af41580156108d0573d6000803e3d6000fd5b505050505050565b6040805163ca9a4ea560e01b8152600060048201819052602482018690526044820185905260648201849052915173__$5c50ea773c3f1223822c80c8edbce14ea7$__9263ca9a4ea59260848082019391829003018186803b15801561093d57600080fd5b505af4158015610951573d6000803e3d6000fd5b50505050505050565b6040805163c84b96f560e01b81526000600482018190526001600160a01b038516602483015260448201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163c84b96f5916064808301926020929190829003018186803b15801561046457600080fd5b60408051632432b5d360e11b81526000600482018190526001600160a01b038516602483015260448201849052915173__$d2345629a19a8ba674464680be5d957238$__926348656ba69260648082019391829003018186803b1580156108bc57600080fd5b604080516344bacc4b60e01b8152600060048201819052915173__$799602413129f49037f52758954cf5aa52$__926344bacc4b9260248082019391829003018186803b1580156104e757600080fd5b604080516316d0383760e11b8152600060048201819052602482018590528315156044830152915173__$5c50ea773c3f1223822c80c8edbce14ea7$__92632da0706e9260648082019391829003018186803b1580156108bc57600080fd5b610aea60008263ffffffff610aed16565b50565b60408051652fb7bbb732b960d11b815281519081900360060190206000908152603f84016020522054336001600160a01b0390911614610b2c57600080fd5b60408051652fb7bbb732b960d11b815281519081900360060181206000908152603f8501602052918220546001600160a01b03808516939116917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a360408051652fb7bbb732b960d11b815281519081900360060190206000908152603f90930160205290912080546001600160a01b039092166001600160a01b031990921691909117905556fea265627a7a723058204ed69f8834274910d3d9493dcb33431ec6bbda4897487775254418b8ea1efcc464736f6c634300050a0032`

// DeployZap deploys a new Ethereum contract, binding an instance of Zap to it.
func DeployZap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Zap, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Zap{ZapCaller: ZapCaller{contract: contract}, ZapTransactor: ZapTransactor{contract: contract}, ZapFilterer: ZapFilterer{contract: contract}}, nil
}

// Zap is an auto generated Go binding around an Ethereum contract.
type Zap struct {
	ZapCaller     // Read-only binding to the contract
	ZapTransactor // Write-only binding to the contract
	ZapFilterer   // Log filterer for contract events
}

// ZapCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapSession struct {
	Contract     *Zap              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapCallerSession struct {
	Contract *ZapCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapTransactorSession struct {
	Contract     *ZapTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapRaw struct {
	Contract *Zap // Generic contract binding to access the raw methods on
}

// ZapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapCallerRaw struct {
	Contract *ZapCaller // Generic read-only contract binding to access the raw methods on
}

// ZapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapTransactorRaw struct {
	Contract *ZapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZap creates a new instance of Zap, bound to a specific deployed contract.
func NewZap(address common.Address, backend bind.ContractBackend) (*Zap, error) {
	contract, err := bindZap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zap{ZapCaller: ZapCaller{contract: contract}, ZapTransactor: ZapTransactor{contract: contract}, ZapFilterer: ZapFilterer{contract: contract}}, nil
}

// NewZapCaller creates a new read-only instance of Zap, bound to a specific deployed contract.
func NewZapCaller(address common.Address, caller bind.ContractCaller) (*ZapCaller, error) {
	contract, err := bindZap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapCaller{contract: contract}, nil
}

// NewZapTransactor creates a new write-only instance of Zap, bound to a specific deployed contract.
func NewZapTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapTransactor, error) {
	contract, err := bindZap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapTransactor{contract: contract}, nil
}

// NewZapFilterer creates a new log filterer instance of Zap, bound to a specific deployed contract.
func NewZapFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapFilterer, error) {
	contract, err := bindZap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapFilterer{contract: contract}, nil
}

// bindZap binds a generic wrapper to an already deployed contract.
func bindZap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zap *ZapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Zap.Contract.ZapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zap *ZapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zap.Contract.ZapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zap *ZapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zap.Contract.ZapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zap *ZapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Zap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zap *ZapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zap *ZapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zap.Contract.contract.Transact(opts, method, params...)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(_requestId uint256, _tip uint256) returns()
func (_Zap *ZapTransactor) AddTip(opts *bind.TransactOpts, _requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "addTip", _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(_requestId uint256, _tip uint256) returns()
func (_Zap *ZapSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.AddTip(&_Zap.TransactOpts, _requestId, _tip)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(_requestId uint256, _tip uint256) returns()
func (_Zap *ZapTransactorSession) AddTip(_requestId *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.AddTip(&_Zap.TransactOpts, _requestId, _tip)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _amount uint256) returns(bool)
func (_Zap *ZapTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _amount uint256) returns(bool)
func (_Zap *ZapSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.Approve(&_Zap.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _amount uint256) returns(bool)
func (_Zap *ZapTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.Approve(&_Zap.TransactOpts, _spender, _amount)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(_requestId uint256, _timestamp uint256, _minerIndex uint256) returns()
func (_Zap *ZapTransactor) BeginDispute(opts *bind.TransactOpts, _requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "beginDispute", _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(_requestId uint256, _timestamp uint256, _minerIndex uint256) returns()
func (_Zap *ZapSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.BeginDispute(&_Zap.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// BeginDispute is a paid mutator transaction binding the contract method 0x8581af19.
//
// Solidity: function beginDispute(_requestId uint256, _timestamp uint256, _minerIndex uint256) returns()
func (_Zap *ZapTransactorSession) BeginDispute(_requestId *big.Int, _timestamp *big.Int, _minerIndex *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.BeginDispute(&_Zap.TransactOpts, _requestId, _timestamp, _minerIndex)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Zap *ZapTransactor) DepositStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "depositStake")
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Zap *ZapSession) DepositStake() (*types.Transaction, error) {
	return _Zap.Contract.DepositStake(&_Zap.TransactOpts)
}

// DepositStake is a paid mutator transaction binding the contract method 0x0d2d76a2.
//
// Solidity: function depositStake() returns()
func (_Zap *ZapTransactorSession) DepositStake() (*types.Transaction, error) {
	return _Zap.Contract.DepositStake(&_Zap.TransactOpts)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(_propNewZapAddress address) returns()
func (_Zap *ZapTransactor) ProposeFork(opts *bind.TransactOpts, _propNewZapAddress common.Address) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "proposeFork", _propNewZapAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(_propNewZapAddress address) returns()
func (_Zap *ZapSession) ProposeFork(_propNewZapAddress common.Address) (*types.Transaction, error) {
	return _Zap.Contract.ProposeFork(&_Zap.TransactOpts, _propNewZapAddress)
}

// ProposeFork is a paid mutator transaction binding the contract method 0x26b7d9f6.
//
// Solidity: function proposeFork(_propNewZapAddress address) returns()
func (_Zap *ZapTransactorSession) ProposeFork(_propNewZapAddress common.Address) (*types.Transaction, error) {
	return _Zap.Contract.ProposeFork(&_Zap.TransactOpts, _propNewZapAddress)
}

// RequestData is a paid mutator transaction binding the contract method 0x3fff2816.
//
// Solidity: function requestData(_c_sapi string, _c_symbol string, _granularity uint256, _tip uint256) returns()
func (_Zap *ZapTransactor) RequestData(opts *bind.TransactOpts, _c_sapi string, _c_symbol string, _granularity *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "requestData", _c_sapi, _c_symbol, _granularity, _tip)
}

// RequestData is a paid mutator transaction binding the contract method 0x3fff2816.
//
// Solidity: function requestData(_c_sapi string, _c_symbol string, _granularity uint256, _tip uint256) returns()
func (_Zap *ZapSession) RequestData(_c_sapi string, _c_symbol string, _granularity *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.RequestData(&_Zap.TransactOpts, _c_sapi, _c_symbol, _granularity, _tip)
}

// RequestData is a paid mutator transaction binding the contract method 0x3fff2816.
//
// Solidity: function requestData(_c_sapi string, _c_symbol string, _granularity uint256, _tip uint256) returns()
func (_Zap *ZapTransactorSession) RequestData(_c_sapi string, _c_symbol string, _granularity *big.Int, _tip *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.RequestData(&_Zap.TransactOpts, _c_sapi, _c_symbol, _granularity, _tip)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Zap *ZapTransactor) RequestStakingWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "requestStakingWithdraw")
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Zap *ZapSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _Zap.Contract.RequestStakingWithdraw(&_Zap.TransactOpts)
}

// RequestStakingWithdraw is a paid mutator transaction binding the contract method 0x28449c3a.
//
// Solidity: function requestStakingWithdraw() returns()
func (_Zap *ZapTransactorSession) RequestStakingWithdraw() (*types.Transaction, error) {
	return _Zap.Contract.RequestStakingWithdraw(&_Zap.TransactOpts)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(_nonce string, _requestId uint256, _value uint256) returns()
func (_Zap *ZapTransactor) SubmitMiningSolution(opts *bind.TransactOpts, _nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "submitMiningSolution", _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(_nonce string, _requestId uint256, _value uint256) returns()
func (_Zap *ZapSession) SubmitMiningSolution(_nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.SubmitMiningSolution(&_Zap.TransactOpts, _nonce, _requestId, _value)
}

// SubmitMiningSolution is a paid mutator transaction binding the contract method 0x68c180d5.
//
// Solidity: function submitMiningSolution(_nonce string, _requestId uint256, _value uint256) returns()
func (_Zap *ZapTransactorSession) SubmitMiningSolution(_nonce string, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.SubmitMiningSolution(&_Zap.TransactOpts, _nonce, _requestId, _value)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(_disputeId uint256) returns()
func (_Zap *ZapTransactor) TallyVotes(opts *bind.TransactOpts, _disputeId *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "tallyVotes", _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(_disputeId uint256) returns()
func (_Zap *ZapSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.TallyVotes(&_Zap.TransactOpts, _disputeId)
}

// TallyVotes is a paid mutator transaction binding the contract method 0x4d318b0e.
//
// Solidity: function tallyVotes(_disputeId uint256) returns()
func (_Zap *ZapTransactorSession) TallyVotes(_disputeId *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.TallyVotes(&_Zap.TransactOpts, _disputeId)
}

// TheLazyCoon is a paid mutator transaction binding the contract method 0xb079f64a.
//
// Solidity: function theLazyCoon(_address address, _amount uint256) returns()
func (_Zap *ZapTransactor) TheLazyCoon(opts *bind.TransactOpts, _address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "theLazyCoon", _address, _amount)
}

// TheLazyCoon is a paid mutator transaction binding the contract method 0xb079f64a.
//
// Solidity: function theLazyCoon(_address address, _amount uint256) returns()
func (_Zap *ZapSession) TheLazyCoon(_address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.TheLazyCoon(&_Zap.TransactOpts, _address, _amount)
}

// TheLazyCoon is a paid mutator transaction binding the contract method 0xb079f64a.
//
// Solidity: function theLazyCoon(_address address, _amount uint256) returns()
func (_Zap *ZapTransactorSession) TheLazyCoon(_address common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.TheLazyCoon(&_Zap.TransactOpts, _address, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _amount uint256) returns(bool)
func (_Zap *ZapTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _amount uint256) returns(bool)
func (_Zap *ZapSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.Transfer(&_Zap.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _amount uint256) returns(bool)
func (_Zap *ZapTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.Transfer(&_Zap.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _amount uint256) returns(bool)
func (_Zap *ZapTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _amount uint256) returns(bool)
func (_Zap *ZapSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.TransferFrom(&_Zap.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _amount uint256) returns(bool)
func (_Zap *ZapTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Zap.Contract.TransferFrom(&_Zap.TransactOpts, _from, _to, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Zap *ZapTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Zap *ZapSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Zap.Contract.TransferOwnership(&_Zap.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Zap *ZapTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Zap.Contract.TransferOwnership(&_Zap.TransactOpts, _newOwner)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(_disputeId uint256, _supportsDispute bool) returns()
func (_Zap *ZapTransactor) Vote(opts *bind.TransactOpts, _disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "vote", _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(_disputeId uint256, _supportsDispute bool) returns()
func (_Zap *ZapSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Zap.Contract.Vote(&_Zap.TransactOpts, _disputeId, _supportsDispute)
}

// Vote is a paid mutator transaction binding the contract method 0xc9d27afe.
//
// Solidity: function vote(_disputeId uint256, _supportsDispute bool) returns()
func (_Zap *ZapTransactorSession) Vote(_disputeId *big.Int, _supportsDispute bool) (*types.Transaction, error) {
	return _Zap.Contract.Vote(&_Zap.TransactOpts, _disputeId, _supportsDispute)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Zap *ZapTransactor) WithdrawStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zap.contract.Transact(opts, "withdrawStake")
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Zap *ZapSession) WithdrawStake() (*types.Transaction, error) {
	return _Zap.Contract.WithdrawStake(&_Zap.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xbed9d861.
//
// Solidity: function withdrawStake() returns()
func (_Zap *ZapTransactorSession) WithdrawStake() (*types.Transaction, error) {
	return _Zap.Contract.WithdrawStake(&_Zap.TransactOpts)
}

// ZapDisputeABI is the input ABI used to generate the binding from.
const ZapDisputeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"NewDispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_position\",\"type\":\"bool\"},{\"indexed\":true,\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_disputeID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_result\",\"type\":\"int256\"},{\"indexed\":true,\"name\":\"_reportedMiner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_reportingParty\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"DisputeVoteTallied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_newZap\",\"type\":\"address\"}],\"name\":\"NewZapAddress\",\"type\":\"event\"}]"

// ZapDisputeBin is the compiled bytecode used for deploying new contracts.
const ZapDisputeBin = `0x6115f2610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100615760003560e01c80632da0706e14610066578063694bf49f146100a0578063ca9a4ea5146100d9578063def6fac714610115578063e15f6f7014610145575b600080fd5b81801561007257600080fd5b5061009e6004803603606081101561008957600080fd5b5080359060208101359060400135151561016f565b005b8180156100ac57600080fd5b5061009e600480360360408110156100c357600080fd5b50803590602001356001600160a01b0316610362565b8180156100e557600080fd5b5061009e600480360360808110156100fc57600080fd5b508035906020810135906040810135906060013561079d565b81801561012157600080fd5b5061009e6004803603604081101561013857600080fd5b5080359060200135610dda565b81801561015157600080fd5b5061009e6004803603602081101561016857600080fd5b50356113e7565b600082815260448085016020908152604080842081516a313637b1b5a73ab6b132b960a91b8152825190819003600b018120865260058201845282862054633f48b1ff60e01b8252600482018a905233602483015294810194909452905190939273__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__92633f48b1ff92606480840193829003018186803b15801561020657600080fd5b505af415801561021a573d6000803e3d6000fd5b505050506040513d602081101561023057600080fd5b505133600090815260068401602052604090205490915060ff1615156001141561025957600080fd5b6000811161026657600080fd5b3360009081526047860160205260409020546003141561028557600080fd5b336000908152600680840160209081526040808420805460ff1916600190811790915581516c6e756d6265724f66566f74657360981b8152825190819003600d01812086526005880180855283872080549093019092556571756f72756d60d01b81528251908190039094019093208452919052902080548201905582156103165760018201805482019055610322565b60018201805482900390555b6040805184151581529051339186917f86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae09181900360200190a35050505050565b604080516bffffffffffffffffffffffff19606084901b1660208083019190915282518083036014018152603490920183528151918101919091206000818152604a860190925291902054156103b757600080fd5b60408051696469737075746546656560b01b8152815190819003600a01812060009081528286016020528281205463c7bb46ad60e01b8352600483018790523360248401523060448401526064830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9263c7bb46ad9260848082019391829003018186803b15801561044057600080fd5b505af4158015610454573d6000803e3d6000fd5b5050505082604001600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c0190506040518091039020815260200190815260200160002060008154809291906001019190505550600083604001600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c019050604051809103902081526020019081526020016000205490508084604a0160008481526020019081526020016000208190555060405180610100016040528083815260200160008152602001600015158152602001600015158152602001600115158152602001336001600160a01b03168152602001336001600160a01b03168152602001846001600160a01b0316815250846044016000838152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548160ff02191690831515021790555060608201518160020160016101000a81548160ff02191690831515021790555060808201518160020160026101000a81548160ff02191690831515021790555060a08201518160020160036101000a8154816001600160a01b0302191690836001600160a01b0316021790555060c08201518160030160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060e08201518160040160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555090505043846044016000838152602001908152602001600020600501600060405180806a313637b1b5a73ab6b132b960a91b815250600b01905060405180910390208152602001908152602001600020819055508360400160006040518080696469737075746546656560b01b815250600a0190506040518091039020815260200190815260200160002054846044016000838152602001908152602001600020600501600060405180806266656560e81b815250600301905060405180910390208152602001908152602001600020819055504262093a8001846044016000838152602001908152602001600020600501600060405180806f6d696e457865637574696f6e4461746560801b8152506010019050604051809103902081526020019081526020016000208190555050505050565b6000838152604885016020908152604080832085845260058101909252909120546090439190910311156107d057600080fd5b60008381526005820160205260409020546107ea57600080fd5b600582106107f757600080fd5b60008381526008820160205260408120836005811061081257fe5b0154604080516bffffffffffffffffffffffff19606084901b1660208083019190915260348201899052605480830189905283518084039091018152607490920183528151918101919091206000818152604a8b01909252919020546001600160a01b039092169250901561088657600080fd5b60408051696469737075746546656560b01b8152815190819003600a0181206000908152828a016020528281205463c7bb46ad60e01b8352600483018b90523360248401523060448401526064830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9263c7bb46ad9260848082019391829003018186803b15801561090f57600080fd5b505af4158015610923573d6000803e3d6000fd5b5050505086604001600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c019050604051809103902081526020019081526020016000205460010187604001600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c0190506040518091039020815260200190815260200160002081905550600087604001600060405180806b191a5cdc1d5d1950dbdd5b9d60a21b815250600c019050604051809103902081526020019081526020016000205490508088604a0160008481526020019081526020016000208190555060405180610100016040528083815260200160008152602001600015158152602001600015158152602001600015158152602001846001600160a01b03168152602001336001600160a01b0316815260200160006001600160a01b0316815250886044016000838152602001908152602001600020600082015181600001556020820151816001015560408201518160020160006101000a81548160ff02191690831515021790555060608201518160020160016101000a81548160ff02191690831515021790555060808201518160020160026101000a81548160ff02191690831515021790555060a08201518160020160036101000a8154816001600160a01b0302191690836001600160a01b0316021790555060c08201518160030160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060e08201518160040160006101000a8154816001600160a01b0302191690836001600160a01b031602179055509050508688604401600083815260200190815260200160002060050160006040518080681c995c5d595cdd125960ba1b81525060090190506040518091039020815260200190815260200160002081905550858860440160008381526020019081526020016000206005016000604051808068074696d657374616d760bc1b815250600901905060405180910390208152602001908152602001600020819055508360090160008781526020019081526020016000208560058110610c2a57fe5b0154600082815260448a016020818152604080842081516476616c756560d81b8152825190819003600590810182208752909101808452828620969096558685528383526f6d696e457865637574696f6e4461746560801b81528151908190036010018120855285835281852062093a80420190558685528383526a313637b1b5a73ab6b132b960a91b8152815190819003600b0181208552858352818520439055868552838352681b5a5b995c94db1bdd60ba1b8152815190819003600901812085528583528185208b9055696469737075746546656560b01b8152815190819003600a0181208552818e018352818520548786529383526266656560e81b815281519081900360030190208452939052919020556002851415610d7357600087815260488901602090815260408083208984526007019091529020805460ff191660011790555b6001600160a01b038316600081815260478a016020908152604091829020600390558151898152908101929092528051899284927feceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da6492918290030190a35050505050505050565b600081815260448301602090815260408083208151681c995c5d595cdd125960ba1b81528251908190036009019020845260058101835281842054845260488601909252909120600282015460ff1615610e3357600080fd5b604080516f6d696e457865637574696f6e4461746560801b8152815190819003601001902060009081526005840160205220544211610e7157600080fd5b600282015462010000900460ff16611258576002820154630100000090046001600160a01b03166000908152604785016020526040812060018401549091121561110d576000815562015180420642036001820155604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b0190206000908152818701602052208054600019019055610f03856113e7565b60028301546003840154604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b0181206000908152828a016020528281205463c7bb46ad60e01b8352600483018b90526001600160a01b0363010000009096048616602484015293909416604482015260648101929092525173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9263c7bb46ad9260848082019391829003018186803b158015610fad57600080fd5b505af4158015610fc1573d6000803e3d6000fd5b505050600380850154604080516266656560e81b815281519081900390930183206000908152600588016020528181205463c7bb46ad60e01b8552600485018b90523060248601526001600160a01b03909316604485015260648401929092525173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__935063c7bb46ad926084808201939291829003018186803b15801561105b57600080fd5b505af415801561106f573d6000803e3d6000fd5b50505060028401805461ff001916610100179055506040805168074696d657374616d760bc1b815281519081900360090190206000908152600585016020908152828220548252600785019052205460ff16151560011415611108576040805168074696d657374616d760bc1b815281519081900360090190206000908152600585016020908152828220548252600685019052908120555b611252565b600181556002830154604080516266656560e81b815281519081900360030181206000908152600587016020528281205463c7bb46ad60e01b8352600483018a90523060248401526001600160a01b0363010000009095049490941660448301526064820193909352905173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9263c7bb46ad9260848082019391829003018186803b1580156111af57600080fd5b505af41580156111c3573d6000803e3d6000fd5b50506040805168074696d657374616d760bc1b815281519081900360090190206000908152600587016020908152828220548252600787019052205460ff1615156001141591506112529050576040805168074696d657374616d760bc1b81528151908190036009019020600090815260058501602090815282822054825260078501905220805460ff191690555b50611369565b60008260010154131561136957604080516b746f74616c5f737570706c7960a01b8152815190819003600c01902060009081528186016020522054606490601402604080516571756f72756d60d01b815281519081900360060190206000908152600586016020522054919004106112cf57600080fd5b600482018054604080516d1d195b1b1bdc90dbdb9d1c9858dd60921b8152815190819003600e0181206000908152603f890160209081529083902080546001600160a01b0319166001600160a01b0395861617905560028701805461ff00191661010017905593549092168252517fc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d929181900390910190a15b60028201805460ff19166001908117918290558301546003840154604080519283526001600160a01b039182166020840152610100840460ff16151583820152516301000000909304169185917f21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739919081900360600190a350505050565b604080516b7461726765744d696e65727360a01b8152815190819003600c01812060009081528284016020818152848320546a1cdd185ad95c90dbdd5b9d60aa1b8552855194859003600b0190942083525291909120546103e8919082028161144c57fe5b04101561154657604080516b7461726765744d696e65727360a01b8152815190819003600c01812060009081528284016020818152848320546a1cdd185ad95c90dbdd5b9d60aa1b8552855194859003600b0190942083525291909120546115159167d02ab486cedc0000916103e891611508918302816114c957fe5b604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b019020600090815281890160205220549190046103e80363ffffffff61158016565b8161150f57fe5b046115a7565b60408051696469737075746546656560b01b8152815190819003600a0190206000908152818401602052205561157d565b60408051696469737075746546656560b01b8152815190819003600a01902060009081528183016020522067d02ab486cedc000090555b50565b600082820283158061159a57508284828161159757fe5b04145b6115a057fe5b9392505050565b60008183116115b657816115a0565b509091905056fea265627a7a72305820dadb2b8460cb9a7746c8f349d0279e13a2fc595a14a9d398772ad34d573d450064736f6c634300050a0032`

// DeployZapDispute deploys a new Ethereum contract, binding an instance of ZapDispute to it.
func DeployZapDispute(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapDispute, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapDisputeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapDisputeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapDispute{ZapDisputeCaller: ZapDisputeCaller{contract: contract}, ZapDisputeTransactor: ZapDisputeTransactor{contract: contract}, ZapDisputeFilterer: ZapDisputeFilterer{contract: contract}}, nil
}

// ZapDispute is an auto generated Go binding around an Ethereum contract.
type ZapDispute struct {
	ZapDisputeCaller     // Read-only binding to the contract
	ZapDisputeTransactor // Write-only binding to the contract
	ZapDisputeFilterer   // Log filterer for contract events
}

// ZapDisputeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapDisputeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapDisputeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapDisputeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapDisputeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapDisputeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapDisputeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapDisputeSession struct {
	Contract     *ZapDispute       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapDisputeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapDisputeCallerSession struct {
	Contract *ZapDisputeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZapDisputeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapDisputeTransactorSession struct {
	Contract     *ZapDisputeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZapDisputeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapDisputeRaw struct {
	Contract *ZapDispute // Generic contract binding to access the raw methods on
}

// ZapDisputeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapDisputeCallerRaw struct {
	Contract *ZapDisputeCaller // Generic read-only contract binding to access the raw methods on
}

// ZapDisputeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapDisputeTransactorRaw struct {
	Contract *ZapDisputeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapDispute creates a new instance of ZapDispute, bound to a specific deployed contract.
func NewZapDispute(address common.Address, backend bind.ContractBackend) (*ZapDispute, error) {
	contract, err := bindZapDispute(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapDispute{ZapDisputeCaller: ZapDisputeCaller{contract: contract}, ZapDisputeTransactor: ZapDisputeTransactor{contract: contract}, ZapDisputeFilterer: ZapDisputeFilterer{contract: contract}}, nil
}

// NewZapDisputeCaller creates a new read-only instance of ZapDispute, bound to a specific deployed contract.
func NewZapDisputeCaller(address common.Address, caller bind.ContractCaller) (*ZapDisputeCaller, error) {
	contract, err := bindZapDispute(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapDisputeCaller{contract: contract}, nil
}

// NewZapDisputeTransactor creates a new write-only instance of ZapDispute, bound to a specific deployed contract.
func NewZapDisputeTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapDisputeTransactor, error) {
	contract, err := bindZapDispute(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapDisputeTransactor{contract: contract}, nil
}

// NewZapDisputeFilterer creates a new log filterer instance of ZapDispute, bound to a specific deployed contract.
func NewZapDisputeFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapDisputeFilterer, error) {
	contract, err := bindZapDispute(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapDisputeFilterer{contract: contract}, nil
}

// bindZapDispute binds a generic wrapper to an already deployed contract.
func bindZapDispute(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapDisputeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapDispute *ZapDisputeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZapDispute.Contract.ZapDisputeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapDispute *ZapDisputeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapDispute.Contract.ZapDisputeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapDispute *ZapDisputeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapDispute.Contract.ZapDisputeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapDispute *ZapDisputeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZapDispute.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapDispute *ZapDisputeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapDispute.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapDispute *ZapDisputeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapDispute.Contract.contract.Transact(opts, method, params...)
}

// ZapDisputeDisputeVoteTalliedIterator is returned from FilterDisputeVoteTallied and is used to iterate over the raw logs and unpacked data for DisputeVoteTallied events raised by the ZapDispute contract.
type ZapDisputeDisputeVoteTalliedIterator struct {
	Event *ZapDisputeDisputeVoteTallied // Event containing the contract specifics and raw log

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
func (it *ZapDisputeDisputeVoteTalliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapDisputeDisputeVoteTallied)
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
		it.Event = new(ZapDisputeDisputeVoteTallied)
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
func (it *ZapDisputeDisputeVoteTalliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapDisputeDisputeVoteTalliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapDisputeDisputeVoteTallied represents a DisputeVoteTallied event raised by the ZapDispute contract.
type ZapDisputeDisputeVoteTallied struct {
	DisputeID      *big.Int
	Result         *big.Int
	ReportedMiner  common.Address
	ReportingParty common.Address
	Active         bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDisputeVoteTallied is a free log retrieval operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: e DisputeVoteTallied(_disputeID indexed uint256, _result int256, _reportedMiner indexed address, _reportingParty address, _active bool)
func (_ZapDispute *ZapDisputeFilterer) FilterDisputeVoteTallied(opts *bind.FilterOpts, _disputeID []*big.Int, _reportedMiner []common.Address) (*ZapDisputeDisputeVoteTalliedIterator, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _reportedMinerRule []interface{}
	for _, _reportedMinerItem := range _reportedMiner {
		_reportedMinerRule = append(_reportedMinerRule, _reportedMinerItem)
	}

	logs, sub, err := _ZapDispute.contract.FilterLogs(opts, "DisputeVoteTallied", _disputeIDRule, _reportedMinerRule)
	if err != nil {
		return nil, err
	}
	return &ZapDisputeDisputeVoteTalliedIterator{contract: _ZapDispute.contract, event: "DisputeVoteTallied", logs: logs, sub: sub}, nil
}

// WatchDisputeVoteTallied is a free log subscription operation binding the contract event 0x21459c2f5447ebcf83a7f0a238c32c71076faef0d12295e771c0cb1e10434739.
//
// Solidity: e DisputeVoteTallied(_disputeID indexed uint256, _result int256, _reportedMiner indexed address, _reportingParty address, _active bool)
func (_ZapDispute *ZapDisputeFilterer) WatchDisputeVoteTallied(opts *bind.WatchOpts, sink chan<- *ZapDisputeDisputeVoteTallied, _disputeID []*big.Int, _reportedMiner []common.Address) (event.Subscription, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _reportedMinerRule []interface{}
	for _, _reportedMinerItem := range _reportedMiner {
		_reportedMinerRule = append(_reportedMinerRule, _reportedMinerItem)
	}

	logs, sub, err := _ZapDispute.contract.WatchLogs(opts, "DisputeVoteTallied", _disputeIDRule, _reportedMinerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapDisputeDisputeVoteTallied)
				if err := _ZapDispute.contract.UnpackLog(event, "DisputeVoteTallied", log); err != nil {
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

// ZapDisputeNewDisputeIterator is returned from FilterNewDispute and is used to iterate over the raw logs and unpacked data for NewDispute events raised by the ZapDispute contract.
type ZapDisputeNewDisputeIterator struct {
	Event *ZapDisputeNewDispute // Event containing the contract specifics and raw log

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
func (it *ZapDisputeNewDisputeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapDisputeNewDispute)
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
		it.Event = new(ZapDisputeNewDispute)
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
func (it *ZapDisputeNewDisputeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapDisputeNewDisputeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapDisputeNewDispute represents a NewDispute event raised by the ZapDispute contract.
type ZapDisputeNewDispute struct {
	DisputeId *big.Int
	RequestId *big.Int
	Timestamp *big.Int
	Miner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewDispute is a free log retrieval operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: e NewDispute(_disputeId indexed uint256, _requestId indexed uint256, _timestamp uint256, _miner address)
func (_ZapDispute *ZapDisputeFilterer) FilterNewDispute(opts *bind.FilterOpts, _disputeId []*big.Int, _requestId []*big.Int) (*ZapDisputeNewDisputeIterator, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapDispute.contract.FilterLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapDisputeNewDisputeIterator{contract: _ZapDispute.contract, event: "NewDispute", logs: logs, sub: sub}, nil
}

// WatchNewDispute is a free log subscription operation binding the contract event 0xeceec1aebf67772b2440120c4b4dc913a1fe1b865509219f9456785c23b9da64.
//
// Solidity: e NewDispute(_disputeId indexed uint256, _requestId indexed uint256, _timestamp uint256, _miner address)
func (_ZapDispute *ZapDisputeFilterer) WatchNewDispute(opts *bind.WatchOpts, sink chan<- *ZapDisputeNewDispute, _disputeId []*big.Int, _requestId []*big.Int) (event.Subscription, error) {

	var _disputeIdRule []interface{}
	for _, _disputeIdItem := range _disputeId {
		_disputeIdRule = append(_disputeIdRule, _disputeIdItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapDispute.contract.WatchLogs(opts, "NewDispute", _disputeIdRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapDisputeNewDispute)
				if err := _ZapDispute.contract.UnpackLog(event, "NewDispute", log); err != nil {
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

// ZapDisputeNewZapAddressIterator is returned from FilterNewZapAddress and is used to iterate over the raw logs and unpacked data for NewZapAddress events raised by the ZapDispute contract.
type ZapDisputeNewZapAddressIterator struct {
	Event *ZapDisputeNewZapAddress // Event containing the contract specifics and raw log

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
func (it *ZapDisputeNewZapAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapDisputeNewZapAddress)
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
		it.Event = new(ZapDisputeNewZapAddress)
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
func (it *ZapDisputeNewZapAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapDisputeNewZapAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapDisputeNewZapAddress represents a NewZapAddress event raised by the ZapDispute contract.
type ZapDisputeNewZapAddress struct {
	NewZap common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewZapAddress is a free log retrieval operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: e NewZapAddress(_newZap address)
func (_ZapDispute *ZapDisputeFilterer) FilterNewZapAddress(opts *bind.FilterOpts) (*ZapDisputeNewZapAddressIterator, error) {

	logs, sub, err := _ZapDispute.contract.FilterLogs(opts, "NewZapAddress")
	if err != nil {
		return nil, err
	}
	return &ZapDisputeNewZapAddressIterator{contract: _ZapDispute.contract, event: "NewZapAddress", logs: logs, sub: sub}, nil
}

// WatchNewZapAddress is a free log subscription operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: e NewZapAddress(_newZap address)
func (_ZapDispute *ZapDisputeFilterer) WatchNewZapAddress(opts *bind.WatchOpts, sink chan<- *ZapDisputeNewZapAddress) (event.Subscription, error) {

	logs, sub, err := _ZapDispute.contract.WatchLogs(opts, "NewZapAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapDisputeNewZapAddress)
				if err := _ZapDispute.contract.UnpackLog(event, "NewZapAddress", log); err != nil {
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

// ZapDisputeVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the ZapDispute contract.
type ZapDisputeVotedIterator struct {
	Event *ZapDisputeVoted // Event containing the contract specifics and raw log

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
func (it *ZapDisputeVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapDisputeVoted)
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
		it.Event = new(ZapDisputeVoted)
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
func (it *ZapDisputeVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapDisputeVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapDisputeVoted represents a Voted event raised by the ZapDispute contract.
type ZapDisputeVoted struct {
	DisputeID *big.Int
	Position  bool
	Voter     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: e Voted(_disputeID indexed uint256, _position bool, _voter indexed address)
func (_ZapDispute *ZapDisputeFilterer) FilterVoted(opts *bind.FilterOpts, _disputeID []*big.Int, _voter []common.Address) (*ZapDisputeVotedIterator, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}

	logs, sub, err := _ZapDispute.contract.FilterLogs(opts, "Voted", _disputeIDRule, _voterRule)
	if err != nil {
		return nil, err
	}
	return &ZapDisputeVotedIterator{contract: _ZapDispute.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x86abfce99b7dd908bec0169288797f85049ec73cbe046ed9de818fab3a497ae0.
//
// Solidity: e Voted(_disputeID indexed uint256, _position bool, _voter indexed address)
func (_ZapDispute *ZapDisputeFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *ZapDisputeVoted, _disputeID []*big.Int, _voter []common.Address) (event.Subscription, error) {

	var _disputeIDRule []interface{}
	for _, _disputeIDItem := range _disputeID {
		_disputeIDRule = append(_disputeIDRule, _disputeIDItem)
	}

	var _voterRule []interface{}
	for _, _voterItem := range _voter {
		_voterRule = append(_voterRule, _voterItem)
	}

	logs, sub, err := _ZapDispute.contract.WatchLogs(opts, "Voted", _disputeIDRule, _voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapDisputeVoted)
				if err := _ZapDispute.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ZapGettersLibraryABI is the input ABI used to generate the binding from.
const ZapGettersLibraryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_newZap\",\"type\":\"address\"}],\"name\":\"NewZapAddress\",\"type\":\"event\"}]"

// ZapGettersLibraryBin is the compiled bytecode used for deploying new contracts.
const ZapGettersLibraryBin = `0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72305820da6ebe9dfbd29ba4e44451f27443534bf73520041279e6b4bb3ce66caf83e1e364736f6c634300050a0032`

// DeployZapGettersLibrary deploys a new Ethereum contract, binding an instance of ZapGettersLibrary to it.
func DeployZapGettersLibrary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapGettersLibrary, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapGettersLibraryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapGettersLibraryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapGettersLibrary{ZapGettersLibraryCaller: ZapGettersLibraryCaller{contract: contract}, ZapGettersLibraryTransactor: ZapGettersLibraryTransactor{contract: contract}, ZapGettersLibraryFilterer: ZapGettersLibraryFilterer{contract: contract}}, nil
}

// ZapGettersLibrary is an auto generated Go binding around an Ethereum contract.
type ZapGettersLibrary struct {
	ZapGettersLibraryCaller     // Read-only binding to the contract
	ZapGettersLibraryTransactor // Write-only binding to the contract
	ZapGettersLibraryFilterer   // Log filterer for contract events
}

// ZapGettersLibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapGettersLibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapGettersLibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapGettersLibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapGettersLibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapGettersLibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapGettersLibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapGettersLibrarySession struct {
	Contract     *ZapGettersLibrary // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZapGettersLibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapGettersLibraryCallerSession struct {
	Contract *ZapGettersLibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ZapGettersLibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapGettersLibraryTransactorSession struct {
	Contract     *ZapGettersLibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ZapGettersLibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapGettersLibraryRaw struct {
	Contract *ZapGettersLibrary // Generic contract binding to access the raw methods on
}

// ZapGettersLibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapGettersLibraryCallerRaw struct {
	Contract *ZapGettersLibraryCaller // Generic read-only contract binding to access the raw methods on
}

// ZapGettersLibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapGettersLibraryTransactorRaw struct {
	Contract *ZapGettersLibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapGettersLibrary creates a new instance of ZapGettersLibrary, bound to a specific deployed contract.
func NewZapGettersLibrary(address common.Address, backend bind.ContractBackend) (*ZapGettersLibrary, error) {
	contract, err := bindZapGettersLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapGettersLibrary{ZapGettersLibraryCaller: ZapGettersLibraryCaller{contract: contract}, ZapGettersLibraryTransactor: ZapGettersLibraryTransactor{contract: contract}, ZapGettersLibraryFilterer: ZapGettersLibraryFilterer{contract: contract}}, nil
}

// NewZapGettersLibraryCaller creates a new read-only instance of ZapGettersLibrary, bound to a specific deployed contract.
func NewZapGettersLibraryCaller(address common.Address, caller bind.ContractCaller) (*ZapGettersLibraryCaller, error) {
	contract, err := bindZapGettersLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapGettersLibraryCaller{contract: contract}, nil
}

// NewZapGettersLibraryTransactor creates a new write-only instance of ZapGettersLibrary, bound to a specific deployed contract.
func NewZapGettersLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapGettersLibraryTransactor, error) {
	contract, err := bindZapGettersLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapGettersLibraryTransactor{contract: contract}, nil
}

// NewZapGettersLibraryFilterer creates a new log filterer instance of ZapGettersLibrary, bound to a specific deployed contract.
func NewZapGettersLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapGettersLibraryFilterer, error) {
	contract, err := bindZapGettersLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapGettersLibraryFilterer{contract: contract}, nil
}

// bindZapGettersLibrary binds a generic wrapper to an already deployed contract.
func bindZapGettersLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapGettersLibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapGettersLibrary *ZapGettersLibraryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZapGettersLibrary.Contract.ZapGettersLibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapGettersLibrary *ZapGettersLibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapGettersLibrary.Contract.ZapGettersLibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapGettersLibrary *ZapGettersLibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapGettersLibrary.Contract.ZapGettersLibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapGettersLibrary *ZapGettersLibraryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZapGettersLibrary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapGettersLibrary *ZapGettersLibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapGettersLibrary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapGettersLibrary *ZapGettersLibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapGettersLibrary.Contract.contract.Transact(opts, method, params...)
}

// ZapGettersLibraryNewZapAddressIterator is returned from FilterNewZapAddress and is used to iterate over the raw logs and unpacked data for NewZapAddress events raised by the ZapGettersLibrary contract.
type ZapGettersLibraryNewZapAddressIterator struct {
	Event *ZapGettersLibraryNewZapAddress // Event containing the contract specifics and raw log

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
func (it *ZapGettersLibraryNewZapAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapGettersLibraryNewZapAddress)
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
		it.Event = new(ZapGettersLibraryNewZapAddress)
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
func (it *ZapGettersLibraryNewZapAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapGettersLibraryNewZapAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapGettersLibraryNewZapAddress represents a NewZapAddress event raised by the ZapGettersLibrary contract.
type ZapGettersLibraryNewZapAddress struct {
	NewZap common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewZapAddress is a free log retrieval operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: e NewZapAddress(_newZap address)
func (_ZapGettersLibrary *ZapGettersLibraryFilterer) FilterNewZapAddress(opts *bind.FilterOpts) (*ZapGettersLibraryNewZapAddressIterator, error) {

	logs, sub, err := _ZapGettersLibrary.contract.FilterLogs(opts, "NewZapAddress")
	if err != nil {
		return nil, err
	}
	return &ZapGettersLibraryNewZapAddressIterator{contract: _ZapGettersLibrary.contract, event: "NewZapAddress", logs: logs, sub: sub}, nil
}

// WatchNewZapAddress is a free log subscription operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: e NewZapAddress(_newZap address)
func (_ZapGettersLibrary *ZapGettersLibraryFilterer) WatchNewZapAddress(opts *bind.WatchOpts, sink chan<- *ZapGettersLibraryNewZapAddress) (event.Subscription, error) {

	logs, sub, err := _ZapGettersLibrary.contract.WatchLogs(opts, "NewZapAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapGettersLibraryNewZapAddress)
				if err := _ZapGettersLibrary.contract.UnpackLog(event, "NewZapAddress", log); err != nil {
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

// ZapLibraryABI is the input ABI used to generate the binding from.
const ZapLibraryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_tip\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"TipAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_querySymbol\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_granularity\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"DataRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_currentChallenge\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_currentRequestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_multiplier\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_totalTips\",\"type\":\"uint256\"}],\"name\":\"NewChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_query\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_onDeckQueryHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_onDeckTotalTips\",\"type\":\"uint256\"}],\"name\":\"NewRequestOnDeck\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_time\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_totalTips\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NewValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_miner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_nonce\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_currentChallenge\",\"type\":\"bytes32\"}],\"name\":\"NonceSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ZapLibraryBin is the compiled bytecode used for deploying new contracts.
const ZapLibraryBin = `0x6126ec610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c806302e8f21b1461005b57806348656ba61461009357806385997bf6146100d2578063a098b5b414610218575b600080fd5b81801561006757600080fd5b506100916004803603606081101561007e57600080fd5b50803590602081013590604001356102d7565b005b81801561009f57600080fd5b50610091600480360360608110156100b657600080fd5b508035906001600160a01b0360208201351690604001356103ed565b8180156100de57600080fd5b50610091600480360360a08110156100f557600080fd5b8135919081019060408101602082013564010000000081111561011757600080fd5b82018360208201111561012957600080fd5b8035906020019184600183028401116401000000008311171561014b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561019e57600080fd5b8201836020820111156101b057600080fd5b803590602001918460018302840111640100000000831117156101d257600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356104a5565b81801561022457600080fd5b506100916004803603608081101561023b57600080fd5b8135919081019060408101602082013564010000000081111561025d57600080fd5b82018360208201111561026f57600080fd5b8035906020019184600183028401116401000000008311171561029157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610901565b600082116102e457600080fd5b8015610368576040805163c7bb46ad60e01b81526004810185905233602482015230604482015260648101839052905173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163c7bb46ad916084808301926000929190829003018186803b15801561034f57600080fd5b505af4158015610363573d6000803e3d6000fd5b505050505b610373838383610dc2565b600082815260488401602090815260408083208151670746f74616c5469760c41b815282519081900360080181208552600490910183529281902054848452918301919091528051849233927fd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e882092918290030190a3505050565b604080516b746f74616c5f737570706c7960a01b8152815190819003600c018120600090815282860160209081528382208054860190556001600160a01b0386168252604587019052828120631d6f7b8160e31b8352600483015260248201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9263eb7bdc089260448082019391829003018186803b15801561048857600080fd5b505af415801561049c573d6000803e3d6000fd5b50505050505050565b600082116104b257600080fd5b670de0b6b3a76400008211156104c757600080fd5b8351849084906104d657600080fd5b60408151106104e457600080fd5b600082856040516020018083805190602001908083835b6020831061051a5780518252601f1990920191602091820191016104fb565b51815160209384036101000a60001901801990921691161790529201938452506040805180850381529382018152835193820193909320600081815260498e019092529290205491935050151590506108da57604080516b1c995c5d595cdd10dbdd5b9d60a21b8082528251600c928190038301812060009081528c8501602081815286832080546001019055938352855192839003909401822081529282528383205460808201855287825281830187905281850186905284518481528084018652606083015280845260488d01835293909220825180519192610604928492909101906124f8565b50602082810151805161061d92600185019201906124f8565b506040820151600282015560608201518051610643916003840191602090910190612576565b505050600081815260488a016020818152604080842081516a6772616e756c617269747960a81b8152825190819003600b018120865260049091018084528286208c90558686528484526f3932b8bab2b9ba28a837b9b4ba34b7b760811b825282519182900360100182208652808452828620869055868652938352670746f74616c5469760c41b81528151908190036008019020845291815281832083905584835260498c01905290208190558415610775576040805163c7bb46ad60e01b8152600481018b905233602482015230604482015260648101879052905173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163c7bb46ad916084808301926000929190829003018186803b15801561075c57600080fd5b505af4158015610770573d6000803e3d6000fd5b505050505b610780898287610dc2565b600081815260488a016020908152604091829020825192830189905260608301889052608080845281546002600180831615610100026000190190921604918501829052859433947f6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc99493928401928d928d929091829182019060a08301908890801561084e5780601f106108235761010080835404028352916020019161084e565b820191906000526020600020905b81548152906001019060200180831161083157829003601f168201915b50508381038252865460026000196101006001841615020190911604808252602090910190879080156108c25780601f10610897576101008083540402835291602001916108c2565b820191906000526020600020905b8154815290600101906020018083116108a557829003601f168201915b5050965050505050505060405180910390a3506108f7565b60008181526049890160205260409020546108f7908990866102d7565b5050505050505050565b33600090815260478501602052604090205460011461091f57600080fd5b604080516f18dd5c9c995b9d14995c5d595cdd125960821b8152815190819003601001902060009081528186016020522054821461095c57600080fd5b836040016000604051808069646966666963756c747960b01b815250600a0190506040518091039020815260200190815260200160002054600260038660000154338760405160200180848152602001836001600160a01b03166001600160a01b031660601b815260140182805190602001908083835b602083106109f25780518252601f1990920191602091820191016109d3565b6001836020036101000a038019825116818451168082178552505050505050905001935050505060405160208183030381529060405280519060200120604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b60208310610a7d5780518252601f199092019160209182019101610a5e565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610abc573d6000803e3d6000fd5b5050506040515160601b60405160200180826bffffffffffffffffffffffff19166bffffffffffffffffffffffff191681526014019150506040516020818303038152906040526040518082805190602001908083835b60208310610b325780518252601f199092019160209182019101610b13565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015610b71573d6000803e3d6000fd5b5050506040513d6020811015610b8657600080fd5b505181610b8f57fe5b0615610b9a57600080fd5b83546000908152604185016020908152604080832033845290915290205460ff1615610bc557600080fd5b604080516b736c6f7450726f677265737360a01b8152815190819003600c019020600090815281860160205220548190603586019060058110610c0457fe5b6002020155604080516b736c6f7450726f677265737360a01b8152815190819003600c019020600090815281860160205220543390603586019060058110610c4857fe5b60020201600190810180546001600160a01b0319166001600160a01b039390931692909217909155604080516b736c6f7450726f677265737360a01b81528151600c91819003919091018120600090815287830160209081528382208054860190558854825260418901815283822033808452908252848320805460ff19169096179095558854818401879052938301849052606080845288519084015287518795947fe6d63a2aee0aaed2ab49702313ce54114f2145af219d7db30d6624af9e6dffc4948a9489949293919283926080840192918801918190849084905b83811015610d3f578181015183820152602001610d27565b50505050905090810190601f168015610d6c5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a3604080516b736c6f7450726f677265737360a01b8152815190819003600c0190206000908152818601602052205460051415610dbc57610dbc8484846114ad565b50505050565b6000828152604884016020526040812090610ddc85612421565b90508215610e4b5760408051670746f74616c5469760c41b815281519081900360080190206000908152600484016020522054610e1f908463ffffffff61247d16565b60408051670746f74616c5469760c41b8152815190819003600801902060009081526004850160205220555b60408051670746f74616c5469760c41b815281519081900360080181206000908152600485016020908152838220546f18dd5c9c995b9d14995c5d595cdd125960821b84528451938490036010019093208252888401905291909120546112235760008360040160006040518080670746f74616c5469760c41b815250600801905060405180910390208152602001908152602001600020819055508486604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020819055508086604001600060405180806f63757272656e74546f74616c5469707360801b81525060100190506040518091039020815260200190815260200160002081905550808660000154600143034060405160200180848152602001838152602001828152602001935050505060405160208183030381529060405280519060200120866000018190555085604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020547f9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f428760000154886040016000604051808069646966666963756c747960b01b815250600a01905060405180910390208152602001908152602001600020548960480160008b604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020548152602001908152602001600020600401600060405180806a6772616e756c617269747960a81b815250600b01905060405180910390208152602001908152602001600020548a60480160008c604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b8152506010019050604051809103902081526020019081526020016000205481526020019081526020016000206000018b604001600060405180806f63757272656e74546f74616c5469707360801b81525060100190506040518091039020815260200190815260200160002054604051808681526020018581526020018481526020018060200183815260200182810382528481815460018160011615610100020316600290048152602001915080546001816001161561010002031660029004801561120c5780601f106111e15761010080835404028352916020019161120c565b820191906000526020600020905b8154815290600101906020018083116111ef57829003601f168201915b5050965050505050505060405180910390a26114a5565b600082815260488701602090815260408083208151670746f74616c5469760c41b815282519081900360080190208452600401909152902054811180611267575081155b156113345760028084015460408051602081018390529081018490526060808252865460001961010060018316150201169390930492810183905287927f5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c928792909186919081906080820190869080156113235780601f106112f857610100808354040283529160200191611323565b820191906000526020600020905b81548152906001019060200180831161130657829003601f168201915b505094505050505060405180910390a25b604080516f3932b8bab2b9ba28a837b9b4ba34b7b760811b815281519081900360100190206000908152600485016020522054611452576040805161066081019182905260009182916113a99160018b019060339082845b81548152602001906001019080831161138c575050505050612493565b9092509050818311806113ba575081155b1561144b57828860010182603381106113cf57fe5b0155600081815260438901602081815260408084208054855260488d01835281852082516f3932b8bab2b9ba28a837b9b4ba34b7b760811b80825284519182900360109081018320895260049384018752858920899055898952968652928e905591825282519182900390940190208452918801905290208190555b50506114a5565b83156114a557604080516f3932b8bab2b9ba28a837b9b4ba34b7b760811b815281519081900360100190206000908152600485016020522054849060018801906033811061149c57fe5b01805490910190555b505050505050565b6000818152604884016020908152604080832081517174696d654f664c6173744e657756616c756560701b81528251601291819003919091018120855282880180855283862054691d1a5b5955185c99d95d60b21b83528451600a938190038401812088528287528588205469646966666963756c747960b01b808352875192839003860183208a52848952878a20549083528751928390039095019091208852919095529285205491946064429590950390930302929092059091019081136115a5576040805169646966666963756c747960b01b8152815190819003600a019020600090815281870160205220600190556115d4565b6040805169646966666963756c747960b01b8152815190819003600a0190206000908152818701602052208190555b603c42604080517174696d654f664c6173744e657756616c756560701b81528151908190036012019020600090815281890160205220919006420390556116196125b0565b6040805160a081019091526035870160056000835b8282101561166f5760408051808201909152600283028501805482526001908101546001600160a01b0316602080840191909152918352909201910161162e565b509293506001925050505b600581101561179957600082826005811061169157fe5b602002015151905060008383600581106116a757fe5b602002015160200151905060008390505b6000811180156116db57508460018203600581106116d257fe5b60200201515183105b1561174d578460018203600581106116ef57fe5b60200201515185826005811061170157fe5b6020020151528460001982016005811061171757fe5b60200201516020015185826005811061172c57fe5b602090810291909101516001600160a01b03909216910152600019016116b8565b8381101561178e578285826005811061176257fe5b6020020151528185826005811061177557fe5b602090810291909101516001600160a01b039092169101525b50505060010161167a565b5060005b60058110156118ba5773__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__63c7bb46ad88308585600581106117ce57fe5b60200201516020015160058c604001600060405180806f63757272656e74546f74616c5469707360801b815250601001905060405180910390208152602001908152602001600020548161181e57fe5b04674563918244f40000016040518563ffffffff1660e01b815260040180858152602001846001600160a01b03166001600160a01b03168152602001836001600160a01b03166001600160a01b0316815260200182815260200194505050505060006040518083038186803b15801561189657600080fd5b505af41580156118aa573d6000803e3d6000fd5b50506001909201915061179d9050565b604080517174696d654f664c6173744e657756616c756560701b81528151908190036012018120600090815289830160208181528483205487860151516f63757272656e74546f74616c5469707360801b80875287519687900360109081018820875285855288872054918852885197889003018720865293835293869020548d549186529185019390935260059091069003828401526060820152905186917fc509d04e0782579ee59912139246ccbdf6c36c43abd90591d912017f3467dd16916080918190039190910190a2604080516b746f74616c5f737570706c7960a01b8152815190819003600c0181206000908152828a016020908152838220805468017da3a04c7b3e0000019055652fb7bbb732b960d11b835283519283900360060183208252603f8b0190528281205463c7bb46ad60e01b8352600483018b90523060248401526001600160a01b031660448301526722b1c8c1227a00006064830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9263c7bb46ad9260848082019391829003018186803b158015611a5757600080fd5b505af4158015611a6b573d6000803e3d6000fd5b5050505081600260058110611a7c57fe5b6020908102919091015151604080517174696d654f664c6173744e657756616c756560701b80825282516012928190038301812060009081528d850180885285822054825260068c018852858220969096558282528451918290038401822081528587528481205460038c018054600181018255908352888320015560a08201855288518701516001600160a01b03908116835289880151880151811683890152898601518801518116838701526060808b01518901518216908401526080808b0151890151909116908301528451928352845192839003909301909120825292845281812054815260088801909352909120611b7a9160056125de565b506040805160a081018252835151815260208085015151818301528483015151828401526060808601515190830152608080860151519083015282517174696d654f664c6173744e657756616c756560701b8152835190819003601201902060009081528a840182528381205481526009880190915291909120611bff916005612632565b50604080517174696d654f664c6173744e657756616c756560701b808252825191829003601290810183206000908152848c01602081815286832054835260058b01815286832043905584865286519586900384018620835281815286832054835260428e0181528683208c9055938552855194859003909201842081528183528481205460348d01805460018101825590835284832001556b736c6f7450726f677265737360a01b8452845193849003600c019093208352905290812055611cc787612421565b604080516f18dd5c9c995b9d14995c5d595cdd125960821b80825282516010928190038301812060009081528c850160208181528683209790975592825284519182900390930190208252909252902054156123e057604080516f18dd5c9c995b9d14995c5d595cdd125960821b808252825191829003601090810183206000908152848c01602081815286832054835260488e01808252878420670746f74616c5469760c41b88528851978890036008018820855260049081018352888520546f63757272656e74546f74616c5469707360801b8952895198899003870189208652848452898620559587528751968790038501872084529181528683205483529081528582206f3932b8bab2b9ba28a837b9b4ba34b7b760811b86528651958690039093019094208152910190915290812054600189019060338110611e0b57fe5b018190555060008760430160008960480160008b604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020548152602001908152602001600020600401600060405180806f3932b8bab2b9ba28a837b9b4ba34b7b760811b81525060100190506040518091039020815260200190815260200160002054815260200190815260200160002081905550600087604801600089604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020548152602001908152602001600020600401600060405180806f3932b8bab2b9ba28a837b9b4ba34b7b760811b81525060100190506040518091039020815260200190815260200160002081905550600087604801600089604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b81525060100190506040518091039020815260200190815260200160002054815260200190815260200160002060040160006040518080670746f74616c5469760c41b815250600801905060405180910390208152602001908152602001600020819055506000611fe488612421565b905086886000015460014303406040516020018084805190602001908083835b602083106120235780518252601f199092019160209182019101612004565b6001836020036101000a038019825116818451168082178552505050505050905001838152602001828152602001935050505060405160208183030381529060405280519060200120886000018190555087604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020547f9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f4289600001548a6040016000604051808069646966666963756c747960b01b815250600a01905060405180910390208152602001908152602001600020548b60480160008d604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b815250601001905060405180910390208152602001908152602001600020548152602001908152602001600020600401600060405180806a6772616e756c617269747960a81b815250600b01905060405180910390208152602001908152602001600020548c60480160008e604001600060405180806f18dd5c9c995b9d14995c5d595cdd125960821b8152506010019050604051809103902081526020019081526020016000205481526020019081526020016000206000018d604001600060405180806f63757272656e74546f74616c5469707360801b8152506010019050604051809103902081526020019081526020016000205460405180868152602001858152602001848152602001806020018381526020018281038252848181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156122ce5780601f106122a3576101008083540402835291602001916122ce565b820191906000526020600020905b8154815290600101906020018083116122b157829003601f168201915b5050965050505050505060405180910390a2600081815260488901602090815260408083206002808201548351670746f74616c5469760c41b81528451908190036008018120875260048401865295849020549486018190529285018490526060808652825460001961010060018316150201169190910490850181905285947f5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c949293929181906080820190869080156123ca5780601f1061239f576101008083540402835291602001916123ca565b820191906000526020600020905b8154815290600101906020018083116123ad57829003601f168201915b505094505050505060405180910390a25061049c565b604080516f63757272656e74546f74616c5469707360801b815281519081900360100190206000908152818901602052908120819055875550505050505050565b60408051610660810191829052600091829182916124629190600187019060339082845b8154815260200190600101908083116124455750505050506124ca565b60009081526043909501602052505060409092205492915050565b60008282018381101561248c57fe5b9392505050565b610640810151603260315b80156124c45760208102840151838110156124ba578093508192505b506000190161249e565b50915091565b60008060005b60338110156124c45760208102840151808410156124ef578093508192505b506001016124d0565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061253957805160ff1916838001178555612566565b82800160010185558215612566579182015b8281111561256657825182559160200191906001019061254b565b5061257292915061265f565b5090565b828054828255906000526020600020908101928215612566579160200282018281111561256657825182559160200191906001019061254b565b6040518061014001604052806005905b6125c861267c565b8152602001906001900390816125c05790505090565b8260058101928215612626579160200282015b8281111561262657825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906125f1565b50612572929150612693565b8260058101928215612566579160200282018281111561256657825182559160200191906001019061254b565b61267991905b808211156125725760008155600101612665565b90565b604080518082019091526000808252602082015290565b61267991905b808211156125725780546001600160a01b031916815560010161269956fea265627a7a72305820cd24d65286aac042a0fcdf1c16229df2a2d28a01f35a638603abb9b2969c0fd164736f6c634300050a0032`

// DeployZapLibrary deploys a new Ethereum contract, binding an instance of ZapLibrary to it.
func DeployZapLibrary(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapLibrary, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapLibraryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapLibraryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapLibrary{ZapLibraryCaller: ZapLibraryCaller{contract: contract}, ZapLibraryTransactor: ZapLibraryTransactor{contract: contract}, ZapLibraryFilterer: ZapLibraryFilterer{contract: contract}}, nil
}

// ZapLibrary is an auto generated Go binding around an Ethereum contract.
type ZapLibrary struct {
	ZapLibraryCaller     // Read-only binding to the contract
	ZapLibraryTransactor // Write-only binding to the contract
	ZapLibraryFilterer   // Log filterer for contract events
}

// ZapLibraryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapLibraryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapLibraryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapLibraryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapLibraryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapLibraryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapLibrarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapLibrarySession struct {
	Contract     *ZapLibrary       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapLibraryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapLibraryCallerSession struct {
	Contract *ZapLibraryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZapLibraryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapLibraryTransactorSession struct {
	Contract     *ZapLibraryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZapLibraryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapLibraryRaw struct {
	Contract *ZapLibrary // Generic contract binding to access the raw methods on
}

// ZapLibraryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapLibraryCallerRaw struct {
	Contract *ZapLibraryCaller // Generic read-only contract binding to access the raw methods on
}

// ZapLibraryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapLibraryTransactorRaw struct {
	Contract *ZapLibraryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapLibrary creates a new instance of ZapLibrary, bound to a specific deployed contract.
func NewZapLibrary(address common.Address, backend bind.ContractBackend) (*ZapLibrary, error) {
	contract, err := bindZapLibrary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapLibrary{ZapLibraryCaller: ZapLibraryCaller{contract: contract}, ZapLibraryTransactor: ZapLibraryTransactor{contract: contract}, ZapLibraryFilterer: ZapLibraryFilterer{contract: contract}}, nil
}

// NewZapLibraryCaller creates a new read-only instance of ZapLibrary, bound to a specific deployed contract.
func NewZapLibraryCaller(address common.Address, caller bind.ContractCaller) (*ZapLibraryCaller, error) {
	contract, err := bindZapLibrary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryCaller{contract: contract}, nil
}

// NewZapLibraryTransactor creates a new write-only instance of ZapLibrary, bound to a specific deployed contract.
func NewZapLibraryTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapLibraryTransactor, error) {
	contract, err := bindZapLibrary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryTransactor{contract: contract}, nil
}

// NewZapLibraryFilterer creates a new log filterer instance of ZapLibrary, bound to a specific deployed contract.
func NewZapLibraryFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapLibraryFilterer, error) {
	contract, err := bindZapLibrary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryFilterer{contract: contract}, nil
}

// bindZapLibrary binds a generic wrapper to an already deployed contract.
func bindZapLibrary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapLibraryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapLibrary *ZapLibraryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZapLibrary.Contract.ZapLibraryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapLibrary *ZapLibraryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapLibrary.Contract.ZapLibraryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapLibrary *ZapLibraryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapLibrary.Contract.ZapLibraryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapLibrary *ZapLibraryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZapLibrary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapLibrary *ZapLibraryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapLibrary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapLibrary *ZapLibraryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapLibrary.Contract.contract.Transact(opts, method, params...)
}

// ZapLibraryDataRequestedIterator is returned from FilterDataRequested and is used to iterate over the raw logs and unpacked data for DataRequested events raised by the ZapLibrary contract.
type ZapLibraryDataRequestedIterator struct {
	Event *ZapLibraryDataRequested // Event containing the contract specifics and raw log

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
func (it *ZapLibraryDataRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryDataRequested)
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
		it.Event = new(ZapLibraryDataRequested)
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
func (it *ZapLibraryDataRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryDataRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryDataRequested represents a DataRequested event raised by the ZapLibrary contract.
type ZapLibraryDataRequested struct {
	Sender      common.Address
	Query       string
	QuerySymbol string
	Granularity *big.Int
	RequestId   *big.Int
	TotalTips   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDataRequested is a free log retrieval operation binding the contract event 0x6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9.
//
// Solidity: e DataRequested(_sender indexed address, _query string, _querySymbol string, _granularity uint256, _requestId indexed uint256, _totalTips uint256)
func (_ZapLibrary *ZapLibraryFilterer) FilterDataRequested(opts *bind.FilterOpts, _sender []common.Address, _requestId []*big.Int) (*ZapLibraryDataRequestedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "DataRequested", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryDataRequestedIterator{contract: _ZapLibrary.contract, event: "DataRequested", logs: logs, sub: sub}, nil
}

// WatchDataRequested is a free log subscription operation binding the contract event 0x6d7f869757848b19c8c2456e95cd98885bc6bed062fda4077bb07985e2f76cc9.
//
// Solidity: e DataRequested(_sender indexed address, _query string, _querySymbol string, _granularity uint256, _requestId indexed uint256, _totalTips uint256)
func (_ZapLibrary *ZapLibraryFilterer) WatchDataRequested(opts *bind.WatchOpts, sink chan<- *ZapLibraryDataRequested, _sender []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "DataRequested", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryDataRequested)
				if err := _ZapLibrary.contract.UnpackLog(event, "DataRequested", log); err != nil {
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

// ZapLibraryNewChallengeIterator is returned from FilterNewChallenge and is used to iterate over the raw logs and unpacked data for NewChallenge events raised by the ZapLibrary contract.
type ZapLibraryNewChallengeIterator struct {
	Event *ZapLibraryNewChallenge // Event containing the contract specifics and raw log

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
func (it *ZapLibraryNewChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryNewChallenge)
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
		it.Event = new(ZapLibraryNewChallenge)
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
func (it *ZapLibraryNewChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryNewChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryNewChallenge represents a NewChallenge event raised by the ZapLibrary contract.
type ZapLibraryNewChallenge struct {
	CurrentChallenge [32]byte
	CurrentRequestId *big.Int
	Difficulty       *big.Int
	Multiplier       *big.Int
	Query            string
	TotalTips        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewChallenge is a free log retrieval operation binding the contract event 0x9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42.
//
// Solidity: e NewChallenge(_currentChallenge bytes32, _currentRequestId indexed uint256, _difficulty uint256, _multiplier uint256, _query string, _totalTips uint256)
func (_ZapLibrary *ZapLibraryFilterer) FilterNewChallenge(opts *bind.FilterOpts, _currentRequestId []*big.Int) (*ZapLibraryNewChallengeIterator, error) {

	var _currentRequestIdRule []interface{}
	for _, _currentRequestIdItem := range _currentRequestId {
		_currentRequestIdRule = append(_currentRequestIdRule, _currentRequestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "NewChallenge", _currentRequestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryNewChallengeIterator{contract: _ZapLibrary.contract, event: "NewChallenge", logs: logs, sub: sub}, nil
}

// WatchNewChallenge is a free log subscription operation binding the contract event 0x9e959738f09af22daede9ff9563f8edf3f5cec8e17ce60287c911c95fa767f42.
//
// Solidity: e NewChallenge(_currentChallenge bytes32, _currentRequestId indexed uint256, _difficulty uint256, _multiplier uint256, _query string, _totalTips uint256)
func (_ZapLibrary *ZapLibraryFilterer) WatchNewChallenge(opts *bind.WatchOpts, sink chan<- *ZapLibraryNewChallenge, _currentRequestId []*big.Int) (event.Subscription, error) {

	var _currentRequestIdRule []interface{}
	for _, _currentRequestIdItem := range _currentRequestId {
		_currentRequestIdRule = append(_currentRequestIdRule, _currentRequestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "NewChallenge", _currentRequestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryNewChallenge)
				if err := _ZapLibrary.contract.UnpackLog(event, "NewChallenge", log); err != nil {
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

// ZapLibraryNewRequestOnDeckIterator is returned from FilterNewRequestOnDeck and is used to iterate over the raw logs and unpacked data for NewRequestOnDeck events raised by the ZapLibrary contract.
type ZapLibraryNewRequestOnDeckIterator struct {
	Event *ZapLibraryNewRequestOnDeck // Event containing the contract specifics and raw log

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
func (it *ZapLibraryNewRequestOnDeckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryNewRequestOnDeck)
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
		it.Event = new(ZapLibraryNewRequestOnDeck)
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
func (it *ZapLibraryNewRequestOnDeckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryNewRequestOnDeckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryNewRequestOnDeck represents a NewRequestOnDeck event raised by the ZapLibrary contract.
type ZapLibraryNewRequestOnDeck struct {
	RequestId       *big.Int
	Query           string
	OnDeckQueryHash [32]byte
	OnDeckTotalTips *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewRequestOnDeck is a free log retrieval operation binding the contract event 0x5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c.
//
// Solidity: e NewRequestOnDeck(_requestId indexed uint256, _query string, _onDeckQueryHash bytes32, _onDeckTotalTips uint256)
func (_ZapLibrary *ZapLibraryFilterer) FilterNewRequestOnDeck(opts *bind.FilterOpts, _requestId []*big.Int) (*ZapLibraryNewRequestOnDeckIterator, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "NewRequestOnDeck", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryNewRequestOnDeckIterator{contract: _ZapLibrary.contract, event: "NewRequestOnDeck", logs: logs, sub: sub}, nil
}

// WatchNewRequestOnDeck is a free log subscription operation binding the contract event 0x5ab383fe3e28263c839564bc7545520a7dd34b78cbd7ab21b2da5d2fd027bf6c.
//
// Solidity: e NewRequestOnDeck(_requestId indexed uint256, _query string, _onDeckQueryHash bytes32, _onDeckTotalTips uint256)
func (_ZapLibrary *ZapLibraryFilterer) WatchNewRequestOnDeck(opts *bind.WatchOpts, sink chan<- *ZapLibraryNewRequestOnDeck, _requestId []*big.Int) (event.Subscription, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "NewRequestOnDeck", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryNewRequestOnDeck)
				if err := _ZapLibrary.contract.UnpackLog(event, "NewRequestOnDeck", log); err != nil {
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

// ZapLibraryNewValueIterator is returned from FilterNewValue and is used to iterate over the raw logs and unpacked data for NewValue events raised by the ZapLibrary contract.
type ZapLibraryNewValueIterator struct {
	Event *ZapLibraryNewValue // Event containing the contract specifics and raw log

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
func (it *ZapLibraryNewValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryNewValue)
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
		it.Event = new(ZapLibraryNewValue)
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
func (it *ZapLibraryNewValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryNewValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryNewValue represents a NewValue event raised by the ZapLibrary contract.
type ZapLibraryNewValue struct {
	RequestId        *big.Int
	Time             *big.Int
	Value            *big.Int
	TotalTips        *big.Int
	CurrentChallenge [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewValue is a free log retrieval operation binding the contract event 0xc509d04e0782579ee59912139246ccbdf6c36c43abd90591d912017f3467dd16.
//
// Solidity: e NewValue(_requestId indexed uint256, _time uint256, _value uint256, _totalTips uint256, _currentChallenge bytes32)
func (_ZapLibrary *ZapLibraryFilterer) FilterNewValue(opts *bind.FilterOpts, _requestId []*big.Int) (*ZapLibraryNewValueIterator, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "NewValue", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryNewValueIterator{contract: _ZapLibrary.contract, event: "NewValue", logs: logs, sub: sub}, nil
}

// WatchNewValue is a free log subscription operation binding the contract event 0xc509d04e0782579ee59912139246ccbdf6c36c43abd90591d912017f3467dd16.
//
// Solidity: e NewValue(_requestId indexed uint256, _time uint256, _value uint256, _totalTips uint256, _currentChallenge bytes32)
func (_ZapLibrary *ZapLibraryFilterer) WatchNewValue(opts *bind.WatchOpts, sink chan<- *ZapLibraryNewValue, _requestId []*big.Int) (event.Subscription, error) {

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "NewValue", _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryNewValue)
				if err := _ZapLibrary.contract.UnpackLog(event, "NewValue", log); err != nil {
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

// ZapLibraryNonceSubmittedIterator is returned from FilterNonceSubmitted and is used to iterate over the raw logs and unpacked data for NonceSubmitted events raised by the ZapLibrary contract.
type ZapLibraryNonceSubmittedIterator struct {
	Event *ZapLibraryNonceSubmitted // Event containing the contract specifics and raw log

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
func (it *ZapLibraryNonceSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryNonceSubmitted)
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
		it.Event = new(ZapLibraryNonceSubmitted)
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
func (it *ZapLibraryNonceSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryNonceSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryNonceSubmitted represents a NonceSubmitted event raised by the ZapLibrary contract.
type ZapLibraryNonceSubmitted struct {
	Miner            common.Address
	Nonce            string
	RequestId        *big.Int
	Value            *big.Int
	CurrentChallenge [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNonceSubmitted is a free log retrieval operation binding the contract event 0xe6d63a2aee0aaed2ab49702313ce54114f2145af219d7db30d6624af9e6dffc4.
//
// Solidity: e NonceSubmitted(_miner indexed address, _nonce string, _requestId indexed uint256, _value uint256, _currentChallenge bytes32)
func (_ZapLibrary *ZapLibraryFilterer) FilterNonceSubmitted(opts *bind.FilterOpts, _miner []common.Address, _requestId []*big.Int) (*ZapLibraryNonceSubmittedIterator, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "NonceSubmitted", _minerRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryNonceSubmittedIterator{contract: _ZapLibrary.contract, event: "NonceSubmitted", logs: logs, sub: sub}, nil
}

// WatchNonceSubmitted is a free log subscription operation binding the contract event 0xe6d63a2aee0aaed2ab49702313ce54114f2145af219d7db30d6624af9e6dffc4.
//
// Solidity: e NonceSubmitted(_miner indexed address, _nonce string, _requestId indexed uint256, _value uint256, _currentChallenge bytes32)
func (_ZapLibrary *ZapLibraryFilterer) WatchNonceSubmitted(opts *bind.WatchOpts, sink chan<- *ZapLibraryNonceSubmitted, _miner []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _minerRule []interface{}
	for _, _minerItem := range _miner {
		_minerRule = append(_minerRule, _minerItem)
	}

	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "NonceSubmitted", _minerRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryNonceSubmitted)
				if err := _ZapLibrary.contract.UnpackLog(event, "NonceSubmitted", log); err != nil {
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

// ZapLibraryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZapLibrary contract.
type ZapLibraryOwnershipTransferredIterator struct {
	Event *ZapLibraryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZapLibraryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryOwnershipTransferred)
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
		it.Event = new(ZapLibraryOwnershipTransferred)
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
func (it *ZapLibraryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryOwnershipTransferred represents a OwnershipTransferred event raised by the ZapLibrary contract.
type ZapLibraryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(_previousOwner indexed address, _newOwner indexed address)
func (_ZapLibrary *ZapLibraryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _previousOwner []common.Address, _newOwner []common.Address) (*ZapLibraryOwnershipTransferredIterator, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryOwnershipTransferredIterator{contract: _ZapLibrary.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(_previousOwner indexed address, _newOwner indexed address)
func (_ZapLibrary *ZapLibraryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZapLibraryOwnershipTransferred, _previousOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _previousOwnerRule []interface{}
	for _, _previousOwnerItem := range _previousOwner {
		_previousOwnerRule = append(_previousOwnerRule, _previousOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "OwnershipTransferred", _previousOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryOwnershipTransferred)
				if err := _ZapLibrary.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ZapLibraryTipAddedIterator is returned from FilterTipAdded and is used to iterate over the raw logs and unpacked data for TipAdded events raised by the ZapLibrary contract.
type ZapLibraryTipAddedIterator struct {
	Event *ZapLibraryTipAdded // Event containing the contract specifics and raw log

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
func (it *ZapLibraryTipAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapLibraryTipAdded)
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
		it.Event = new(ZapLibraryTipAdded)
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
func (it *ZapLibraryTipAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapLibraryTipAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapLibraryTipAdded represents a TipAdded event raised by the ZapLibrary contract.
type ZapLibraryTipAdded struct {
	Sender    common.Address
	RequestId *big.Int
	Tip       *big.Int
	TotalTips *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTipAdded is a free log retrieval operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: e TipAdded(_sender indexed address, _requestId indexed uint256, _tip uint256, _totalTips uint256)
func (_ZapLibrary *ZapLibraryFilterer) FilterTipAdded(opts *bind.FilterOpts, _sender []common.Address, _requestId []*big.Int) (*ZapLibraryTipAddedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.FilterLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ZapLibraryTipAddedIterator{contract: _ZapLibrary.contract, event: "TipAdded", logs: logs, sub: sub}, nil
}

// WatchTipAdded is a free log subscription operation binding the contract event 0xd32134405b68f6f7220f9c38ae310df1b648d16188006768d45be2f7c24e8820.
//
// Solidity: e TipAdded(_sender indexed address, _requestId indexed uint256, _tip uint256, _totalTips uint256)
func (_ZapLibrary *ZapLibraryFilterer) WatchTipAdded(opts *bind.WatchOpts, sink chan<- *ZapLibraryTipAdded, _sender []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _ZapLibrary.contract.WatchLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapLibraryTipAdded)
				if err := _ZapLibrary.contract.UnpackLog(event, "TipAdded", log); err != nil {
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

// ZapStakeABI is the input ABI used to generate the binding from.
const ZapStakeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"NewStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"StakeWithdrawRequested\",\"type\":\"event\"}]"

// ZapStakeBin is the compiled bytecode used for deploying new contracts.
const ZapStakeBin = `0x6108c3610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100565760003560e01c806344bacc4b1461005b5780634601f1cd14610087578063820a2d66146100b1578063c9cf5e4c146100db575b600080fd5b81801561006757600080fd5b506100856004803603602081101561007e57600080fd5b5035610105565b005b81801561009357600080fd5b50610085600480360360208110156100aa57600080fd5b5035610173565b8180156100bd57600080fd5b50610085600480360360208110156100d457600080fd5b5035610559565b8180156100e757600080fd5b50610085600480360360208110156100fe57600080fd5b50356105cd565b3360009081526047820160205260409020600181015462093a8090620151804206420303101561013457600080fd5b805460021461014257600080fd5b600080825560405133917f4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec91a25050565b6040805167646563696d616c7360c01b8152815190819003600801902060009081528183016020522054156101a757600080fd5b3060009081526045820160205260408082208151631d6f7b8160e31b8152600481019190915269014542ba12a337c00000196024820152905173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9263eb7bdc089260448082019391829003018186803b15801561021757600080fd5b505af415801561022b573d6000803e3d6000fd5b50505050610237610870565b506040805160c08101825273e037ec8ec9ec423826750853899394de7f024fee815273cdd8fa31af8475574b8909f135d510579a8087d3602082015273b9dd5afd86547df817da2d0fb89334a6f8edd8919181019190915273230570cd052f40e14c14a81038c6f3aa685d712b6060820152733233afa02644ccd048587f8ba6e99b3c00a34dcc608082015273e010ac6e0248790e08f42d5f697160dedf97e02460a082015260005b60068110156103b95773__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__63eb7bdc0884604501600085856006811061031557fe5b60200201516001600160a01b03166001600160a01b03168152602001908152602001600020683635c9adc5dea000006040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561037f57600080fd5b505af4158015610393573d6000803e3d6000fd5b505050506103b1838383600681106103a757fe5b60200201516106c4565b6001016102e0565b50604080516b746f74616c5f737570706c7960a01b8152815190819003600c908101822060009081528386016020818152858320805469014542ba12a337c0000001905567646563696d616c7360c01b855285519485900360080185208352818152858320601290556b7461726765744d696e65727360a01b85528551948590039093018420825280835284822060c890556a1cdd185ad9505b5bdd5b9d60aa1b8452845193849003600b0184208252808352848220683635c9adc5dea000009055696469737075746546656560b01b8452845193849003600a908101852083528184528583206834957444b840e800009055691d1a5b5955185c99d95d60b21b80865286519586900382018620845282855286842061025890558552855194859003019093208152919052205442816104ef57fe5b604080517174696d654f664c6173744e657756616c756560701b815281519081900360120181206000908152958201602081815283882095909406420390945569646966666963756c747960b01b8152815190819003600a01902085529190529091206001905550565b61056381336106c4565b73__$5c50ea773c3f1223822c80c8edbce14ea7$__63e15f6f70826040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b1580156105b257600080fd5b505af41580156105c6573d6000803e3d6000fd5b5050505050565b336000908152604782016020526040902080546001146105ec57600080fd5b6002815562015180420642036001820155604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b01812060009081528285016020528281208054600019019055630e15f6f760e41b825260048201859052915173__$5c50ea773c3f1223822c80c8edbce14ea7$__9263e15f6f709260248082019391829003018186803b15801561067d57600080fd5b505af4158015610691573d6000803e3d6000fd5b50506040513392507f453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf9150600090a25050565b604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b0181206000908152828501602090815290839020546393b182b360e01b8352600483018690526001600160a01b0385166024840152925173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__926393b182b3926044808301939192829003018186803b15801561075057600080fd5b505af4158015610764573d6000803e3d6000fd5b505050506040513d602081101561077a57600080fd5b5051101561078757600080fd5b6001600160a01b038116600090815260478301602052604090205415806107c857506001600160a01b03811660009081526047830160205260409020546002145b6107d157600080fd5b604080516a1cdd185ad95c90dbdd5b9d60aa1b8152815190819003600b01812060009081528483016020908152838220805460019081019091558385018552808452620151804290810690038285019081526001600160a01b038716808552604789019093528584209451855551930192909255915190917f46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e291a25050565b6040518060c00160405280600690602082028038833950919291505056fea265627a7a723058203cc3787f4bf5992f9b9f3d798c1bfdb3f004f46517aa8f3cf7a95e43f392268564736f6c634300050a0032`

// DeployZapStake deploys a new Ethereum contract, binding an instance of ZapStake to it.
func DeployZapStake(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapStake, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapStakeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapStakeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapStake{ZapStakeCaller: ZapStakeCaller{contract: contract}, ZapStakeTransactor: ZapStakeTransactor{contract: contract}, ZapStakeFilterer: ZapStakeFilterer{contract: contract}}, nil
}

// ZapStake is an auto generated Go binding around an Ethereum contract.
type ZapStake struct {
	ZapStakeCaller     // Read-only binding to the contract
	ZapStakeTransactor // Write-only binding to the contract
	ZapStakeFilterer   // Log filterer for contract events
}

// ZapStakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapStakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapStakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapStakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapStakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapStakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapStakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapStakeSession struct {
	Contract     *ZapStake         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapStakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapStakeCallerSession struct {
	Contract *ZapStakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ZapStakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapStakeTransactorSession struct {
	Contract     *ZapStakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ZapStakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapStakeRaw struct {
	Contract *ZapStake // Generic contract binding to access the raw methods on
}

// ZapStakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapStakeCallerRaw struct {
	Contract *ZapStakeCaller // Generic read-only contract binding to access the raw methods on
}

// ZapStakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapStakeTransactorRaw struct {
	Contract *ZapStakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapStake creates a new instance of ZapStake, bound to a specific deployed contract.
func NewZapStake(address common.Address, backend bind.ContractBackend) (*ZapStake, error) {
	contract, err := bindZapStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapStake{ZapStakeCaller: ZapStakeCaller{contract: contract}, ZapStakeTransactor: ZapStakeTransactor{contract: contract}, ZapStakeFilterer: ZapStakeFilterer{contract: contract}}, nil
}

// NewZapStakeCaller creates a new read-only instance of ZapStake, bound to a specific deployed contract.
func NewZapStakeCaller(address common.Address, caller bind.ContractCaller) (*ZapStakeCaller, error) {
	contract, err := bindZapStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapStakeCaller{contract: contract}, nil
}

// NewZapStakeTransactor creates a new write-only instance of ZapStake, bound to a specific deployed contract.
func NewZapStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapStakeTransactor, error) {
	contract, err := bindZapStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapStakeTransactor{contract: contract}, nil
}

// NewZapStakeFilterer creates a new log filterer instance of ZapStake, bound to a specific deployed contract.
func NewZapStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapStakeFilterer, error) {
	contract, err := bindZapStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapStakeFilterer{contract: contract}, nil
}

// bindZapStake binds a generic wrapper to an already deployed contract.
func bindZapStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapStakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapStake *ZapStakeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZapStake.Contract.ZapStakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapStake *ZapStakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapStake.Contract.ZapStakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapStake *ZapStakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapStake.Contract.ZapStakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapStake *ZapStakeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZapStake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapStake *ZapStakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapStake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapStake *ZapStakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapStake.Contract.contract.Transact(opts, method, params...)
}

// ZapStakeNewStakeIterator is returned from FilterNewStake and is used to iterate over the raw logs and unpacked data for NewStake events raised by the ZapStake contract.
type ZapStakeNewStakeIterator struct {
	Event *ZapStakeNewStake // Event containing the contract specifics and raw log

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
func (it *ZapStakeNewStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapStakeNewStake)
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
		it.Event = new(ZapStakeNewStake)
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
func (it *ZapStakeNewStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapStakeNewStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapStakeNewStake represents a NewStake event raised by the ZapStake contract.
type ZapStakeNewStake struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewStake is a free log retrieval operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: e NewStake(_sender indexed address)
func (_ZapStake *ZapStakeFilterer) FilterNewStake(opts *bind.FilterOpts, _sender []common.Address) (*ZapStakeNewStakeIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ZapStake.contract.FilterLogs(opts, "NewStake", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ZapStakeNewStakeIterator{contract: _ZapStake.contract, event: "NewStake", logs: logs, sub: sub}, nil
}

// WatchNewStake is a free log subscription operation binding the contract event 0x46d8ab1385f70e5a3673e97c23c764f7600f7ed7a09b6687deae7131d51752e2.
//
// Solidity: e NewStake(_sender indexed address)
func (_ZapStake *ZapStakeFilterer) WatchNewStake(opts *bind.WatchOpts, sink chan<- *ZapStakeNewStake, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ZapStake.contract.WatchLogs(opts, "NewStake", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapStakeNewStake)
				if err := _ZapStake.contract.UnpackLog(event, "NewStake", log); err != nil {
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

// ZapStakeStakeWithdrawRequestedIterator is returned from FilterStakeWithdrawRequested and is used to iterate over the raw logs and unpacked data for StakeWithdrawRequested events raised by the ZapStake contract.
type ZapStakeStakeWithdrawRequestedIterator struct {
	Event *ZapStakeStakeWithdrawRequested // Event containing the contract specifics and raw log

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
func (it *ZapStakeStakeWithdrawRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapStakeStakeWithdrawRequested)
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
		it.Event = new(ZapStakeStakeWithdrawRequested)
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
func (it *ZapStakeStakeWithdrawRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapStakeStakeWithdrawRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapStakeStakeWithdrawRequested represents a StakeWithdrawRequested event raised by the ZapStake contract.
type ZapStakeStakeWithdrawRequested struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawRequested is a free log retrieval operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: e StakeWithdrawRequested(_sender indexed address)
func (_ZapStake *ZapStakeFilterer) FilterStakeWithdrawRequested(opts *bind.FilterOpts, _sender []common.Address) (*ZapStakeStakeWithdrawRequestedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ZapStake.contract.FilterLogs(opts, "StakeWithdrawRequested", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ZapStakeStakeWithdrawRequestedIterator{contract: _ZapStake.contract, event: "StakeWithdrawRequested", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawRequested is a free log subscription operation binding the contract event 0x453865710d0cb4b14ad25de371c860da196368895daa9662e5087711d14daecf.
//
// Solidity: e StakeWithdrawRequested(_sender indexed address)
func (_ZapStake *ZapStakeFilterer) WatchStakeWithdrawRequested(opts *bind.WatchOpts, sink chan<- *ZapStakeStakeWithdrawRequested, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ZapStake.contract.WatchLogs(opts, "StakeWithdrawRequested", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapStakeStakeWithdrawRequested)
				if err := _ZapStake.contract.UnpackLog(event, "StakeWithdrawRequested", log); err != nil {
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

// ZapStakeStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the ZapStake contract.
type ZapStakeStakeWithdrawnIterator struct {
	Event *ZapStakeStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *ZapStakeStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapStakeStakeWithdrawn)
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
		it.Event = new(ZapStakeStakeWithdrawn)
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
func (it *ZapStakeStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapStakeStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapStakeStakeWithdrawn represents a StakeWithdrawn event raised by the ZapStake contract.
type ZapStakeStakeWithdrawn struct {
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: e StakeWithdrawn(_sender indexed address)
func (_ZapStake *ZapStakeFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, _sender []common.Address) (*ZapStakeStakeWithdrawnIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ZapStake.contract.FilterLogs(opts, "StakeWithdrawn", _senderRule)
	if err != nil {
		return nil, err
	}
	return &ZapStakeStakeWithdrawnIterator{contract: _ZapStake.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x4a7934670bd8304e7da22378be1368f7c4fef17c5aee81804beda8638fe428ec.
//
// Solidity: e StakeWithdrawn(_sender indexed address)
func (_ZapStake *ZapStakeFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *ZapStakeStakeWithdrawn, _sender []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}

	logs, sub, err := _ZapStake.contract.WatchLogs(opts, "StakeWithdrawn", _senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapStakeStakeWithdrawn)
				if err := _ZapStake.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ZapStorageABI is the input ABI used to generate the binding from.
const ZapStorageABI = "[]"

// ZapStorageBin is the compiled bytecode used for deploying new contracts.
const ZapStorageBin = `0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7230582058ec2f862ab3051bed82bfb7dfd8444570e7f6447a286cf14fc693a0a697f7c664736f6c634300050a0032`

// DeployZapStorage deploys a new Ethereum contract, binding an instance of ZapStorage to it.
func DeployZapStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapStorage{ZapStorageCaller: ZapStorageCaller{contract: contract}, ZapStorageTransactor: ZapStorageTransactor{contract: contract}, ZapStorageFilterer: ZapStorageFilterer{contract: contract}}, nil
}

// ZapStorage is an auto generated Go binding around an Ethereum contract.
type ZapStorage struct {
	ZapStorageCaller     // Read-only binding to the contract
	ZapStorageTransactor // Write-only binding to the contract
	ZapStorageFilterer   // Log filterer for contract events
}

// ZapStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapStorageSession struct {
	Contract     *ZapStorage       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapStorageCallerSession struct {
	Contract *ZapStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZapStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapStorageTransactorSession struct {
	Contract     *ZapStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZapStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapStorageRaw struct {
	Contract *ZapStorage // Generic contract binding to access the raw methods on
}

// ZapStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapStorageCallerRaw struct {
	Contract *ZapStorageCaller // Generic read-only contract binding to access the raw methods on
}

// ZapStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapStorageTransactorRaw struct {
	Contract *ZapStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapStorage creates a new instance of ZapStorage, bound to a specific deployed contract.
func NewZapStorage(address common.Address, backend bind.ContractBackend) (*ZapStorage, error) {
	contract, err := bindZapStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapStorage{ZapStorageCaller: ZapStorageCaller{contract: contract}, ZapStorageTransactor: ZapStorageTransactor{contract: contract}, ZapStorageFilterer: ZapStorageFilterer{contract: contract}}, nil
}

// NewZapStorageCaller creates a new read-only instance of ZapStorage, bound to a specific deployed contract.
func NewZapStorageCaller(address common.Address, caller bind.ContractCaller) (*ZapStorageCaller, error) {
	contract, err := bindZapStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapStorageCaller{contract: contract}, nil
}

// NewZapStorageTransactor creates a new write-only instance of ZapStorage, bound to a specific deployed contract.
func NewZapStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapStorageTransactor, error) {
	contract, err := bindZapStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapStorageTransactor{contract: contract}, nil
}

// NewZapStorageFilterer creates a new log filterer instance of ZapStorage, bound to a specific deployed contract.
func NewZapStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapStorageFilterer, error) {
	contract, err := bindZapStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapStorageFilterer{contract: contract}, nil
}

// bindZapStorage binds a generic wrapper to an already deployed contract.
func bindZapStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapStorage *ZapStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZapStorage.Contract.ZapStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapStorage *ZapStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapStorage.Contract.ZapStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapStorage *ZapStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapStorage.Contract.ZapStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapStorage *ZapStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZapStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapStorage *ZapStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapStorage *ZapStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapStorage.Contract.contract.Transact(opts, method, params...)
}

// ZapTransferABI is the input ABI used to generate the binding from.
const ZapTransferABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ZapTransferBin is the compiled bytecode used for deploying new contracts.
const ZapTransferBin = `0x610936610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100a85760003560e01c8063bf32006c11610070578063bf32006c146101c5578063c7bb46ad146101f9578063c84b96f514610244578063ca50189914610283578063eb7bdc08146102cc576100a8565b80633f48b1ff146100ad578063850dcc32146100f157806393b182b3146101445780639be5647f14610170578063acaab9e214610193575b600080fd5b6100df600480360360608110156100c357600080fd5b508035906001600160a01b0360208201351690604001356102fc565b60408051918252519081900360200190f35b8180156100fd57600080fd5b506101306004803603606081101561011457600080fd5b508035906001600160a01b036020820135169060400135610395565b604080519115158252519081900360200190f35b6100df6004803603604081101561015a57600080fd5b50803590602001356001600160a01b0316610428565b6100df6004803603604081101561018657600080fd5b508035906020013561043e565b610130600480360360608110156101a957600080fd5b508035906001600160a01b03602082013516906040013561056e565b6100df600480360360608110156101db57600080fd5b508035906001600160a01b0360208201358116916040013516610614565b81801561020557600080fd5b506102426004803603608081101561021c57600080fd5b508035906001600160a01b03602082013581169160408101359091169060600135610641565b005b81801561025057600080fd5b506101306004803603606081101561026757600080fd5b508035906001600160a01b03602082013516906040013561073f565b81801561028f57600080fd5b50610130600480360360808110156102a657600080fd5b508035906001600160a01b03602082013581169160408101359091169060600135610757565b8180156102d857600080fd5b50610242600480360360408110156102ef57600080fd5b50803590602001356107cc565b6001600160a01b0382166000908152604584016020526040812054158061035a57506001600160a01b03831660009081526045850160205260408120805484929061034357fe5b6000918252602090912001546001600160801b0316115b156103675750600061038e565b6001600160a01b0383166000908152604585016020526040902061038b908361043e565b90505b9392505050565b60006103a284338461056e565b6103ab57600080fd5b6001600160a01b0383166103be57600080fd5b33600081815260468601602090815260408083206001600160a01b03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a35060019392505050565b60006104358383436102fc565b90505b92915050565b815460009061044f57506000610438565b82548390600019810190811061046157fe5b6000918252602090912001546001600160801b031682106104b15782548390600019810190811061048e57fe5b600091825260209091200154600160801b90046001600160801b03169050610438565b826000815481106104be57fe5b6000918252602090912001546001600160801b03168210156104e257506000610438565b8254600090600019015b8181111561053d57600060026001838501010490508486828154811061050e57fe5b6000918252602090912001546001600160801b03161161053057809250610537565b6001810391505b506104ec565b84828154811061054957fe5b600091825260209091200154600160801b90046001600160801b031695945050505050565b6001600160a01b0382166000908152604784016020526040812054156105ed57604080516a1cdd185ad9505b5bdd5b9d60aa1b8152815190819003600b0190206000908152818601602052908120546105db9084906105cf90818989610428565b9063ffffffff6108a516565b106105e85750600161038e565b61060a565b60006105fd836105cf8787610428565b1061060a5750600161038e565b5060009392505050565b6001600160a01b039182166000908152604693909301602090815260408085209290931684525290205490565b6000811161064e57600080fd5b6001600160a01b03821661066157600080fd5b61066c84848361056e565b61067557600080fd5b60006106828585436102fc565b6001600160a01b038516600090815260458701602052604090209091506106ab908383036107cc565b6106b68584436102fc565b90508082820110156106c757600080fd5b6001600160a01b038316600090815260458601602052604090206106ed908284016107cc565b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a35050505050565b600061074d84338585610641565b5060019392505050565b6001600160a01b0383166000908152604685016020908152604080832033845290915281205482111561078957600080fd5b6001600160a01b038416600090815260468601602090815260408083203384529091529020805483900390556107c185858585610641565b506001949350505050565b81541580610800575081544390839060001981019081106107e957fe5b6000918252602090912001546001600160801b0316105b15610867578154600090839061081982600183016108b7565b8154811061082357fe5b600091825260209091200180546001600160801b03848116600160801b024382166fffffffffffffffffffffffffffffffff199093169290921716179055506108a1565b81546000908390600019810190811061087c57fe5b600091825260209091200180546001600160801b03808516600160801b029116179055505b5050565b6000828211156108b157fe5b50900390565b8154818355818111156108db576000838152602090206108db9181019083016108e0565b505050565b6108fe91905b808211156108fa57600081556001016108e6565b5090565b9056fea265627a7a72305820782f81105098c84c9eb011f38e6970ea78434680be8cc6b5d842e4b5772271ea64736f6c634300050a0032`

// DeployZapTransfer deploys a new Ethereum contract, binding an instance of ZapTransfer to it.
func DeployZapTransfer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapTransfer, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapTransferABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapTransferBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapTransfer{ZapTransferCaller: ZapTransferCaller{contract: contract}, ZapTransferTransactor: ZapTransferTransactor{contract: contract}, ZapTransferFilterer: ZapTransferFilterer{contract: contract}}, nil
}

// ZapTransfer is an auto generated Go binding around an Ethereum contract.
type ZapTransfer struct {
	ZapTransferCaller     // Read-only binding to the contract
	ZapTransferTransactor // Write-only binding to the contract
	ZapTransferFilterer   // Log filterer for contract events
}

// ZapTransferCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapTransferCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapTransferTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapTransferTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapTransferFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapTransferFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapTransferSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapTransferSession struct {
	Contract     *ZapTransfer      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapTransferCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapTransferCallerSession struct {
	Contract *ZapTransferCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ZapTransferTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapTransferTransactorSession struct {
	Contract     *ZapTransferTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ZapTransferRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapTransferRaw struct {
	Contract *ZapTransfer // Generic contract binding to access the raw methods on
}

// ZapTransferCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapTransferCallerRaw struct {
	Contract *ZapTransferCaller // Generic read-only contract binding to access the raw methods on
}

// ZapTransferTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapTransferTransactorRaw struct {
	Contract *ZapTransferTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapTransfer creates a new instance of ZapTransfer, bound to a specific deployed contract.
func NewZapTransfer(address common.Address, backend bind.ContractBackend) (*ZapTransfer, error) {
	contract, err := bindZapTransfer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapTransfer{ZapTransferCaller: ZapTransferCaller{contract: contract}, ZapTransferTransactor: ZapTransferTransactor{contract: contract}, ZapTransferFilterer: ZapTransferFilterer{contract: contract}}, nil
}

// NewZapTransferCaller creates a new read-only instance of ZapTransfer, bound to a specific deployed contract.
func NewZapTransferCaller(address common.Address, caller bind.ContractCaller) (*ZapTransferCaller, error) {
	contract, err := bindZapTransfer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapTransferCaller{contract: contract}, nil
}

// NewZapTransferTransactor creates a new write-only instance of ZapTransfer, bound to a specific deployed contract.
func NewZapTransferTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapTransferTransactor, error) {
	contract, err := bindZapTransfer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapTransferTransactor{contract: contract}, nil
}

// NewZapTransferFilterer creates a new log filterer instance of ZapTransfer, bound to a specific deployed contract.
func NewZapTransferFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapTransferFilterer, error) {
	contract, err := bindZapTransfer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapTransferFilterer{contract: contract}, nil
}

// bindZapTransfer binds a generic wrapper to an already deployed contract.
func bindZapTransfer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapTransferABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapTransfer *ZapTransferRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZapTransfer.Contract.ZapTransferCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapTransfer *ZapTransferRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapTransfer.Contract.ZapTransferTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapTransfer *ZapTransferRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapTransfer.Contract.ZapTransferTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapTransfer *ZapTransferCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZapTransfer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapTransfer *ZapTransferTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapTransfer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapTransfer *ZapTransferTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapTransfer.Contract.contract.Transact(opts, method, params...)
}

// ZapTransferApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ZapTransfer contract.
type ZapTransferApprovalIterator struct {
	Event *ZapTransferApproval // Event containing the contract specifics and raw log

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
func (it *ZapTransferApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapTransferApproval)
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
		it.Event = new(ZapTransferApproval)
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
func (it *ZapTransferApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapTransferApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapTransferApproval represents a Approval event raised by the ZapTransfer contract.
type ZapTransferApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_ZapTransfer *ZapTransferFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*ZapTransferApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ZapTransfer.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &ZapTransferApprovalIterator{contract: _ZapTransfer.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_ZapTransfer *ZapTransferFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZapTransferApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ZapTransfer.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapTransferApproval)
				if err := _ZapTransfer.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ZapTransferTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ZapTransfer contract.
type ZapTransferTransferIterator struct {
	Event *ZapTransferTransfer // Event containing the contract specifics and raw log

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
func (it *ZapTransferTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapTransferTransfer)
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
		it.Event = new(ZapTransferTransfer)
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
func (it *ZapTransferTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapTransferTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapTransferTransfer represents a Transfer event raised by the ZapTransfer contract.
type ZapTransferTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_ZapTransfer *ZapTransferFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ZapTransferTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ZapTransfer.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ZapTransferTransferIterator{contract: _ZapTransfer.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_ZapTransfer *ZapTransferFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZapTransferTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ZapTransfer.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapTransferTransfer)
				if err := _ZapTransfer.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// UtilitiesABI is the input ABI used to generate the binding from.
const UtilitiesABI = "[]"

// UtilitiesBin is the compiled bytecode used for deploying new contracts.
const UtilitiesBin = `0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058203a5f66bda6f7a12c4d2010ffb4ddea6b3e3fea0390b238de6b2ad2897120313b64736f6c634300050a0032`

// DeployUtilities deploys a new Ethereum contract, binding an instance of Utilities to it.
func DeployUtilities(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Utilities, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilitiesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UtilitiesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Utilities{UtilitiesCaller: UtilitiesCaller{contract: contract}, UtilitiesTransactor: UtilitiesTransactor{contract: contract}, UtilitiesFilterer: UtilitiesFilterer{contract: contract}}, nil
}

// Utilities is an auto generated Go binding around an Ethereum contract.
type Utilities struct {
	UtilitiesCaller     // Read-only binding to the contract
	UtilitiesTransactor // Write-only binding to the contract
	UtilitiesFilterer   // Log filterer for contract events
}

// UtilitiesCaller is an auto generated read-only Go binding around an Ethereum contract.
type UtilitiesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilitiesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UtilitiesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilitiesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UtilitiesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UtilitiesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UtilitiesSession struct {
	Contract     *Utilities        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UtilitiesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UtilitiesCallerSession struct {
	Contract *UtilitiesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// UtilitiesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UtilitiesTransactorSession struct {
	Contract     *UtilitiesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UtilitiesRaw is an auto generated low-level Go binding around an Ethereum contract.
type UtilitiesRaw struct {
	Contract *Utilities // Generic contract binding to access the raw methods on
}

// UtilitiesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UtilitiesCallerRaw struct {
	Contract *UtilitiesCaller // Generic read-only contract binding to access the raw methods on
}

// UtilitiesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UtilitiesTransactorRaw struct {
	Contract *UtilitiesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUtilities creates a new instance of Utilities, bound to a specific deployed contract.
func NewUtilities(address common.Address, backend bind.ContractBackend) (*Utilities, error) {
	contract, err := bindUtilities(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Utilities{UtilitiesCaller: UtilitiesCaller{contract: contract}, UtilitiesTransactor: UtilitiesTransactor{contract: contract}, UtilitiesFilterer: UtilitiesFilterer{contract: contract}}, nil
}

// NewUtilitiesCaller creates a new read-only instance of Utilities, bound to a specific deployed contract.
func NewUtilitiesCaller(address common.Address, caller bind.ContractCaller) (*UtilitiesCaller, error) {
	contract, err := bindUtilities(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UtilitiesCaller{contract: contract}, nil
}

// NewUtilitiesTransactor creates a new write-only instance of Utilities, bound to a specific deployed contract.
func NewUtilitiesTransactor(address common.Address, transactor bind.ContractTransactor) (*UtilitiesTransactor, error) {
	contract, err := bindUtilities(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UtilitiesTransactor{contract: contract}, nil
}

// NewUtilitiesFilterer creates a new log filterer instance of Utilities, bound to a specific deployed contract.
func NewUtilitiesFilterer(address common.Address, filterer bind.ContractFilterer) (*UtilitiesFilterer, error) {
	contract, err := bindUtilities(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UtilitiesFilterer{contract: contract}, nil
}

// bindUtilities binds a generic wrapper to an already deployed contract.
func bindUtilities(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UtilitiesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utilities *UtilitiesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utilities.Contract.UtilitiesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utilities *UtilitiesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utilities.Contract.UtilitiesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utilities *UtilitiesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utilities.Contract.UtilitiesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Utilities *UtilitiesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Utilities.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Utilities *UtilitiesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Utilities.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Utilities *UtilitiesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Utilities.Contract.contract.Transact(opts, method, params...)
}
