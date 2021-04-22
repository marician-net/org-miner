// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ZapGettersABI is the input ABI used to generate the binding from.
const ZapGettersABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSymbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVariablesOnDeck\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_request\",\"type\":\"bytes32\"}],\"name\":\"getRequestIdByQueryHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"address[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestID\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentVariables\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256[9]\"},{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[51]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ZapGettersBin is the compiled bytecode used for deploying new contracts.
const ZapGettersBin = `0x608060405234801561001057600080fd5b50611a2a806100206000396000f3fe608060405234801561001057600080fd5b50600436106101f05760003560e01c806370a082311161010f578063af0b1327116100a2578063dd62ed3e11610071578063dd62ed3e146107fa578063e0ae93c114610828578063e1eee6d61461084b578063fc7cf0a014610962576101f0565b8063af0b1327146106f8578063b54130291461079c578063c775b542146107ba578063da379941146107dd576101f0565b806393fa4915116100de57806393fa4915146105da578063999cf26c146105fd578063a22e407a14610629578063a7c438bc146106cc576101f0565b806370a082311461052f578063733bdef01461055557806377fbb663146105945780637f6fd5d9146105b7576101f0565b80633180f8df11610187578063612c8f7f11610156578063612c8f7f146104a65780636173c0b8146104c357806363bb82ad146104e057806369026d631461050c576101f0565b80633180f8df146103f05780633df0777b1461042657806346eee1c41461045d5780634ee2cd7e1461047a576101f0565b806317d7de7c116101c357806317d7de7c1461033557806318160ddd1461033d57806319e8e03b146103455780631db842f0146103d3576101f0565b80630f0b424d146101f557806311c9851214610224578063133bee5e1461027f57806315070401146102b8575b600080fd5b6102126004803603602081101561020b57600080fd5b503561096a565b60408051918252519081900360200190f35b6102476004803603604081101561023a57600080fd5b5080359060200135610982565b604051808260a080838360005b8381101561026c578181015183820152602001610254565b5050505090500191505060405180910390f35b61029c6004803603602081101561029557600080fd5b50356109a3565b604080516001600160a01b039092168252519081900360200190f35b6102c06109b5565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102fa5781810151838201526020016102e2565b50505050905090810190601f1680156103275780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102c06109c6565b6102126109d2565b61034d6109de565b6040518084815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561039657818101518382015260200161037e565b50505050905090810190601f1680156103c35780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b610212600480360360208110156103e957600080fd5b50356109f8565b61040d6004803603602081101561040657600080fd5b5035610a0a565b6040805192835290151560208301528051918290030190f35b6104496004803603604081101561043c57600080fd5b5080359060200135610a26565b604080519115158252519081900360200190f35b6102126004803603602081101561047357600080fd5b5035610a39565b6102126004803603604081101561049057600080fd5b506001600160a01b038135169060200135610a4b565b610212600480360360208110156104bc57600080fd5b5035610ae8565b610212600480360360208110156104d957600080fd5b5035610afa565b610449600480360360408110156104f657600080fd5b50803590602001356001600160a01b0316610b0c565b6102476004803603604081101561052257600080fd5b5080359060200135610b1f565b6102126004803603602081101561054557600080fd5b50356001600160a01b0316610b39565b61057b6004803603602081101561056b57600080fd5b50356001600160a01b0316610bce565b6040805192835260208301919091528051918290030190f35b610212600480360360408110156105aa57600080fd5b5080359060200135610be1565b610212600480360360408110156105cd57600080fd5b5080359060200135610bf4565b610212600480360360408110156105f057600080fd5b5080359060200135610c07565b6104496004803603604081101561061357600080fd5b506001600160a01b038135169060200135610c1a565b610631610c84565b6040518087815260200186815260200185815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b8381101561068c578181015183820152602001610674565b50505050905090810190601f1680156106b95780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b610449600480360360408110156106e257600080fd5b50803590602001356001600160a01b0316610cab565b6107156004803603602081101561070e57600080fd5b5035610cbe565b604080518a815289151560208201528815159181019190915286151560608201526001600160a01b03808716608083015285811660a0830152841660c082015260e081018361012080838360005b8381101561077b578181015183820152602001610763565b50505050905001828152602001995050505050505050505060405180910390f35b6107a4610d02565b6040518151815280826106608083836020610254565b610212600480360360408110156107d057600080fd5b5080359060200135610d14565b610212600480360360208110156107f357600080fd5b5035610d27565b6102126004803603604081101561081057600080fd5b506001600160a01b0381358116916020013516610d39565b6102126004803603604081101561083e57600080fd5b5080359060200135610da4565b6108686004803603602081101561086157600080fd5b5035610db7565b604051808060200180602001878152602001868152602001858152602001848152602001838103835289818151815260200191508051906020019080838360005b838110156108c15781810151838201526020016108a9565b50505050905090810190601f1680156108ee5780820380516001836020036101000a031916815260200191505b5083810382528851815288516020918201918a019080838360005b83811015610921578181015183820152602001610909565b50505050905090810190601f16801561094e5780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b61040d610de3565b600061097c818363ffffffff610df816565b92915050565b61098a611999565b61099c6000848463ffffffff610e0e16565b9392505050565b600061097c818363ffffffff610e6816565b60606109c16000610e87565b905090565b60606109c16000610ea4565b60006109c16000610ece565b60008060606109ed6000610f01565b925092509250909192565b600061097c818363ffffffff610fed16565b600080610a1d818463ffffffff61100316565b91509150915091565b600061099c81848463ffffffff61106916565b600061097c818363ffffffff61109016565b60408051633f48b1ff60e01b81526000600482018190526001600160a01b038516602483015260448201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__91633f48b1ff916064808301926020929190829003018186803b158015610ab557600080fd5b505af4158015610ac9573d6000803e3d6000fd5b505050506040513d6020811015610adf57600080fd5b50519392505050565b600061097c818363ffffffff6110a916565b600061097c818363ffffffff6110bb16565b600061099c81848463ffffffff6110e216565b610b27611999565b61099c6000848463ffffffff61110f16565b604080516393b182b360e01b81526000600482018190526001600160a01b0384166024830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__916393b182b3916044808301926020929190829003018186803b158015610b9c57600080fd5b505af4158015610bb0573d6000803e3d6000fd5b505050506040513d6020811015610bc657600080fd5b505192915050565b600080610a1d818463ffffffff61117316565b600061099c81848463ffffffff61119a16565b600061099c81848463ffffffff6111cd16565b600061099c81848463ffffffff6111f116565b604080516356555cf160e11b81526000600482018190526001600160a01b038516602483015260448201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163acaab9e2916064808301926020929190829003018186803b158015610ab557600080fd5b60008060006060600080610c986000611215565b949b939a50919850965094509092509050565b600061099c81848463ffffffff6113d316565b6000806000806000806000610cd16119b7565b6000610ce3818b63ffffffff61140516565b9850985098509850985098509850985098509193959799909294969850565b610d0a6119d6565b6109c1600061161a565b600061099c81848463ffffffff61165a16565b600061097c818363ffffffff61167e16565b60408051632fcc801b60e21b81526000600482018190526001600160a01b03808616602484015284166044830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163bf32006c916064808301926020929190829003018186803b158015610ab557600080fd5b600061099c81848463ffffffff61169416565b6060806000808080610dcf818863ffffffff6116b816565b949c939b5091995097509550909350915050565b600080610df06000611893565b915091509091565b6000908152604291909101602052604090205490565b610e16611999565b6000838152604885016020908152604080832085845260090190915290819020815160a08101928390529160059082845b815481526020019060010190808311610e4757505050505090509392505050565b6000908152603f9190910160205260409020546001600160a01b031690565b50604080518082019091526002815261151560f21b602082015290565b5060408051808201909152600f81526e54656c6c6f7220547269627574657360881b602082015290565b604080516b746f74616c5f737570706c7960a01b8152815190819003600c01902060009081528183016020522054919050565b60008060606000610f1185611909565b600081815260488701602081815260408084208151670746f74616c5469760c41b8152825160089181900391909101812086526004820184528286205495879052938352805460026001821615610100026000190190911604601f810184900484028501840190925281845294955085949291839190830182828015610fd85780601f10610fad57610100808354040283529160200191610fd8565b820191906000526020600020905b815481529060010190602001808311610fbb57829003601f168201915b50505050509050935093509350509193909250565b6000908152604991909101602052604090205490565b600081815260488301602052604081206003810154829190156110595760038101805461104d918791879190600019810190811061103d57fe5b90600052602060002001546111f1565b60019250925050611062565b50600091508190505b9250929050565b60009182526048929092016020908152604080832093835260079093019052205460ff1690565b6000908152604891909101602052604090206003015490565b60009081526040918201602052205490565b600060328211156110cb57600080fd5b506000908152604391909101602052604090205490565b6000918252604192909201602090815260408083206001600160a01b039094168352929052205460ff1690565b611117611999565b6000838152604885016020908152604080832085845260080190915290819020815160a08101928390529160059082845b81546001600160a01b0316815260019091019060200180831161114857505050505090509392505050565b6001600160a01b031660009081526047919091016020526040902080546001909101549091565b600082815260488401602052604081206003018054839081106111b957fe5b906000526020600020015490509392505050565b60009182526044929092016020908152604080832093835260059093019052205490565b60009182526048929092016020908152604080832093835260069093019052205490565b8054604080516f18dd5c9c995b9d14995c5d595cdd125960821b80825282519182900360109081018320600090815284870160208181528683205469646966666963756c747960b01b8752875196879003600a01872084528282528784205486885288519788900386018820855283835288852054855260488b01808452898620888a528a51998a900388018a2087528585528a87205487528185528a87206a6772616e756c617269747960a81b8b528b519a8b9003600b018b208852600490810186528b882054998b528b519a8b90039098018a2087529484528986205486528352888520670746f74616c5469760c41b8952895198899003600801892086529095018252878420548354601f600260001961010060018516150201909216919091049081018490048402890184019099528888529398899889986060988a98899894979596909594919391929185918301828280156113b75780601f1061138c576101008083540402835291602001916113b7565b820191906000526020600020905b81548152906001019060200180831161139a57829003601f168201915b5050505050925095509550955095509550955091939550919395565b600082815260448401602090815260408083206001600160a01b038516845260060190915290205460ff169392505050565b60008060008060008060006114186119b7565b5050506000868152604488016020908152604080832080546002820154600380840154600485015486516101208101808952681c995c5d595cdd125960ba1b905287518082036101290190208a526005808801808b52898c20548352895168074696d657374616d760bc1b81528a519081900360099081019091208d52818c528a8d2054848d01528a516476616c756560d81b81528b51908190039093019092208c52808b52898c2054838b015289516f6d696e457865637574696f6e4461746560801b81528a519081900360100190208c52808b52898c2054606084015289516c6e756d6265724f66566f74657360981b81528a5190819003600d0190208c52808b52898c2054608084015289516a313637b1b5a73ab6b132b960a91b81528a5190819003600b0190208c52808b52898c205460a08401528951681b5a5b995c94db1bdd60ba1b81528a51908190039092019091208b52808a52888b205460c083015288516571756f72756d60d01b815289519081900360060190208b52808a52888b205460e083015288516266656560e81b81528951908190039095019094208a5292909752949096205461010087810191909152600190930154919a5060ff8082169a509281048316985062010000810490921696506001600160a01b0363010000009092048216955091811693921691909295985092959850929598565b6116226119d6565b6040805161066081019182905290600184019060339082845b81548152602001906001019080831161163b5750505050509050919050565b60009182526048929092016020908152604080832093835260059093019052205490565b6000908152604a91909101602052604090205490565b60009182526048929092016020908152604080832093835260049093019052205490565b6000818152604883016020908152604080832060028082015483516a6772616e756c617269747960a81b8152845190819003600b018120875260048401808752858820546f3932b8bab2b9ba28a837b9b4ba34b7b760811b83528651928390036010018320895281885286892054670746f74616c5469760c41b845287519384900360080184208a52918852868920548654601f6000196101006001848116159190910291909101909216979097049687018a90048a0285018a019098528584526060998a999098899889988998919788979388019692959294909188918301828280156117e75780601f106117bc576101008083540402835291602001916117e7565b820191906000526020600020905b8154815290600101906020018083116117ca57829003601f168201915b5050885460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959b508a9450925084019050828280156118755780601f1061184a57610100808354040283529160200191611875565b820191906000526020600020905b81548152906001019060200180831161185857829003601f168201915b50505050509450965096509650965096509650509295509295509295565b604080517174696d654f664c6173744e657756616c756560701b80825282516012928190038301812060009081528585016020818152868320548352604288018152868320549484528651938490039095019092208152925291812054909182916118ff9185916111f1565b9360019350915050565b604080516106608101918290526000918291829161194a9190600187019060339082845b81548152602001906001019080831161192d575050505050611965565b60009081526043909501602052505060409092205492915050565b60008060005b603381101561199357602081028401518084101561198a578093508192505b5060010161196b565b50915091565b6040518060a001604052806005906020820280388339509192915050565b6040518061012001604052806009906020820280388339509192915050565b604051806106600160405280603390602082028038833950919291505056fea265627a7a72305820cc7cdcaa2fea4ffcef4610ccbe2f219baa0c95dc886d7d996fc8219f373ffd7364736f6c634300050a0032`

// DeployZapGetters deploys a new Ethereum contract, binding an instance of ZapGetters to it.
func DeployZapGetters(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ZapGetters, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapGettersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapGettersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapGetters{ZapGettersCaller: ZapGettersCaller{contract: contract}, ZapGettersTransactor: ZapGettersTransactor{contract: contract}, ZapGettersFilterer: ZapGettersFilterer{contract: contract}}, nil
}

// ZapGetters is an auto generated Go binding around an Ethereum contract.
type ZapGetters struct {
	ZapGettersCaller     // Read-only binding to the contract
	ZapGettersTransactor // Write-only binding to the contract
	ZapGettersFilterer   // Log filterer for contract events
}

// ZapGettersCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapGettersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapGettersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapGettersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapGettersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapGettersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapGettersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapGettersSession struct {
	Contract     *ZapGetters       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapGettersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapGettersCallerSession struct {
	Contract *ZapGettersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZapGettersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapGettersTransactorSession struct {
	Contract     *ZapGettersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZapGettersRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapGettersRaw struct {
	Contract *ZapGetters // Generic contract binding to access the raw methods on
}

// ZapGettersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapGettersCallerRaw struct {
	Contract *ZapGettersCaller // Generic read-only contract binding to access the raw methods on
}

// ZapGettersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapGettersTransactorRaw struct {
	Contract *ZapGettersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapGetters creates a new instance of ZapGetters, bound to a specific deployed contract.
func NewZapGetters(address common.Address, backend bind.ContractBackend) (*ZapGetters, error) {
	contract, err := bindZapGetters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapGetters{ZapGettersCaller: ZapGettersCaller{contract: contract}, ZapGettersTransactor: ZapGettersTransactor{contract: contract}, ZapGettersFilterer: ZapGettersFilterer{contract: contract}}, nil
}

// NewZapGettersCaller creates a new read-only instance of ZapGetters, bound to a specific deployed contract.
func NewZapGettersCaller(address common.Address, caller bind.ContractCaller) (*ZapGettersCaller, error) {
	contract, err := bindZapGetters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapGettersCaller{contract: contract}, nil
}

// NewZapGettersTransactor creates a new write-only instance of ZapGetters, bound to a specific deployed contract.
func NewZapGettersTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapGettersTransactor, error) {
	contract, err := bindZapGetters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapGettersTransactor{contract: contract}, nil
}

// NewZapGettersFilterer creates a new log filterer instance of ZapGetters, bound to a specific deployed contract.
func NewZapGettersFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapGettersFilterer, error) {
	contract, err := bindZapGetters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapGettersFilterer{contract: contract}, nil
}

// bindZapGetters binds a generic wrapper to an already deployed contract.
func bindZapGetters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapGettersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapGetters *ZapGettersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZapGetters.Contract.ZapGettersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapGetters *ZapGettersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapGetters.Contract.ZapGettersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapGetters *ZapGettersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapGetters.Contract.ZapGettersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapGetters *ZapGettersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZapGetters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapGetters *ZapGettersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapGetters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapGetters *ZapGettersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapGetters.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_user address, _spender address) constant returns(uint256)
func (_ZapGetters *ZapGettersCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "allowance", _user, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_user address, _spender address) constant returns(uint256)
func (_ZapGetters *ZapGettersSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _ZapGetters.Contract.Allowance(&_ZapGetters.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_user address, _spender address) constant returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _ZapGetters.Contract.Allowance(&_ZapGetters.CallOpts, _user, _spender)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(_user address, _amount uint256) constant returns(bool)
func (_ZapGetters *ZapGettersCaller) AllowedToTrade(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "allowedToTrade", _user, _amount)
	return *ret0, err
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(_user address, _amount uint256) constant returns(bool)
func (_ZapGetters *ZapGettersSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _ZapGetters.Contract.AllowedToTrade(&_ZapGetters.CallOpts, _user, _amount)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(_user address, _amount uint256) constant returns(bool)
func (_ZapGetters *ZapGettersCallerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _ZapGetters.Contract.AllowedToTrade(&_ZapGetters.CallOpts, _user, _amount)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_user address) constant returns(uint256)
func (_ZapGetters *ZapGettersCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "balanceOf", _user)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_user address) constant returns(uint256)
func (_ZapGetters *ZapGettersSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _ZapGetters.Contract.BalanceOf(&_ZapGetters.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_user address) constant returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _ZapGetters.Contract.BalanceOf(&_ZapGetters.CallOpts, _user)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(_user address, _blockNumber uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersCaller) BalanceOfAt(opts *bind.CallOpts, _user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "balanceOfAt", _user, _blockNumber)
	return *ret0, err
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(_user address, _blockNumber uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.BalanceOfAt(&_ZapGetters.CallOpts, _user, _blockNumber)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(_user address, _blockNumber uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.BalanceOfAt(&_ZapGetters.CallOpts, _user, _blockNumber)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(_challenge bytes32, _miner address) constant returns(bool)
func (_ZapGetters *ZapGettersCaller) DidMine(opts *bind.CallOpts, _challenge [32]byte, _miner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "didMine", _challenge, _miner)
	return *ret0, err
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(_challenge bytes32, _miner address) constant returns(bool)
func (_ZapGetters *ZapGettersSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _ZapGetters.Contract.DidMine(&_ZapGetters.CallOpts, _challenge, _miner)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(_challenge bytes32, _miner address) constant returns(bool)
func (_ZapGetters *ZapGettersCallerSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _ZapGetters.Contract.DidMine(&_ZapGetters.CallOpts, _challenge, _miner)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(_disputeId uint256, _address address) constant returns(bool)
func (_ZapGetters *ZapGettersCaller) DidVote(opts *bind.CallOpts, _disputeId *big.Int, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "didVote", _disputeId, _address)
	return *ret0, err
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(_disputeId uint256, _address address) constant returns(bool)
func (_ZapGetters *ZapGettersSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _ZapGetters.Contract.DidVote(&_ZapGetters.CallOpts, _disputeId, _address)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(_disputeId uint256, _address address) constant returns(bool)
func (_ZapGetters *ZapGettersCallerSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _ZapGetters.Contract.DidVote(&_ZapGetters.CallOpts, _disputeId, _address)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(_data bytes32) constant returns(address)
func (_ZapGetters *ZapGettersCaller) GetAddressVars(opts *bind.CallOpts, _data [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getAddressVars", _data)
	return *ret0, err
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(_data bytes32) constant returns(address)
func (_ZapGetters *ZapGettersSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _ZapGetters.Contract.GetAddressVars(&_ZapGetters.CallOpts, _data)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(_data bytes32) constant returns(address)
func (_ZapGetters *ZapGettersCallerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _ZapGetters.Contract.GetAddressVars(&_ZapGetters.CallOpts, _data)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(_disputeId uint256) constant returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ZapGetters *ZapGettersCaller) GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
		ret2 = new(bool)
		ret3 = new(bool)
		ret4 = new(common.Address)
		ret5 = new(common.Address)
		ret6 = new(common.Address)
		ret7 = new([9]*big.Int)
		ret8 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
		ret7,
		ret8,
	}
	err := _ZapGetters.contract.Call(opts, out, "getAllDisputeVars", _disputeId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, *ret7, *ret8, err
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(_disputeId uint256) constant returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ZapGetters *ZapGettersSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetAllDisputeVars(&_ZapGetters.CallOpts, _disputeId)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(_disputeId uint256) constant returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ZapGetters *ZapGettersCallerSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetAllDisputeVars(&_ZapGetters.CallOpts, _disputeId)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() constant returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapGetters *ZapGettersCaller) GetCurrentVariables(opts *bind.CallOpts) ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(string)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _ZapGetters.contract.Call(opts, out, "getCurrentVariables")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() constant returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapGetters *ZapGettersSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetCurrentVariables(&_ZapGetters.CallOpts)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() constant returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapGetters *ZapGettersCallerSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetCurrentVariables(&_ZapGetters.CallOpts)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(_hash bytes32) constant returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetDisputeIdByDisputeHash(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getDisputeIdByDisputeHash", _hash)
	return *ret0, err
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(_hash bytes32) constant returns(uint256)
func (_ZapGetters *ZapGettersSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetDisputeIdByDisputeHash(&_ZapGetters.CallOpts, _hash)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(_hash bytes32) constant returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetDisputeIdByDisputeHash(&_ZapGetters.CallOpts, _hash)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(_disputeId uint256, _data bytes32) constant returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetDisputeUintVars(opts *bind.CallOpts, _disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getDisputeUintVars", _disputeId, _data)
	return *ret0, err
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(_disputeId uint256, _data bytes32) constant returns(uint256)
func (_ZapGetters *ZapGettersSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetDisputeUintVars(&_ZapGetters.CallOpts, _disputeId, _data)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(_disputeId uint256, _data bytes32) constant returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetDisputeUintVars(&_ZapGetters.CallOpts, _disputeId, _data)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() constant returns(uint256, bool)
func (_ZapGetters *ZapGettersCaller) GetLastNewValue(opts *bind.CallOpts) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ZapGetters.contract.Call(opts, out, "getLastNewValue")
	return *ret0, *ret1, err
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() constant returns(uint256, bool)
func (_ZapGetters *ZapGettersSession) GetLastNewValue() (*big.Int, bool, error) {
	return _ZapGetters.Contract.GetLastNewValue(&_ZapGetters.CallOpts)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() constant returns(uint256, bool)
func (_ZapGetters *ZapGettersCallerSession) GetLastNewValue() (*big.Int, bool, error) {
	return _ZapGetters.Contract.GetLastNewValue(&_ZapGetters.CallOpts)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(_requestId uint256) constant returns(uint256, bool)
func (_ZapGetters *ZapGettersCaller) GetLastNewValueById(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ZapGetters.contract.Call(opts, out, "getLastNewValueById", _requestId)
	return *ret0, *ret1, err
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(_requestId uint256) constant returns(uint256, bool)
func (_ZapGetters *ZapGettersSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _ZapGetters.Contract.GetLastNewValueById(&_ZapGetters.CallOpts, _requestId)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(_requestId uint256) constant returns(uint256, bool)
func (_ZapGetters *ZapGettersCallerSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _ZapGetters.Contract.GetLastNewValueById(&_ZapGetters.CallOpts, _requestId)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetMinedBlockNum(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getMinedBlockNum", _requestId, _timestamp)
	return *ret0, err
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetMinedBlockNum(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetMinedBlockNum(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(_requestId uint256, _timestamp uint256) constant returns(address[5])
func (_ZapGetters *ZapGettersCaller) GetMinersByRequestIdAndTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	var (
		ret0 = new([5]common.Address)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getMinersByRequestIdAndTimestamp", _requestId, _timestamp)
	return *ret0, err
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(_requestId uint256, _timestamp uint256) constant returns(address[5])
func (_ZapGetters *ZapGettersSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _ZapGetters.Contract.GetMinersByRequestIdAndTimestamp(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(_requestId uint256, _timestamp uint256) constant returns(address[5])
func (_ZapGetters *ZapGettersCallerSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _ZapGetters.Contract.GetMinersByRequestIdAndTimestamp(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_ZapGetters *ZapGettersCaller) GetName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getName")
	return *ret0, err
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_ZapGetters *ZapGettersSession) GetName() (string, error) {
	return _ZapGetters.Contract.GetName(&_ZapGetters.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_ZapGetters *ZapGettersCallerSession) GetName() (string, error) {
	return _ZapGetters.Contract.GetName(&_ZapGetters.CallOpts)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(_requestId uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getNewValueCountbyRequestId", _requestId)
	return *ret0, err
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(_requestId uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetNewValueCountbyRequestId(&_ZapGetters.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(_requestId uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetNewValueCountbyRequestId(&_ZapGetters.CallOpts, _requestId)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(_request bytes32) constant returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetRequestIdByQueryHash(opts *bind.CallOpts, _request [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getRequestIdByQueryHash", _request)
	return *ret0, err
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(_request bytes32) constant returns(uint256)
func (_ZapGetters *ZapGettersSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestIdByQueryHash(&_ZapGetters.CallOpts, _request)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(_request bytes32) constant returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestIdByQueryHash(&_ZapGetters.CallOpts, _request)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(_index uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetRequestIdByRequestQIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getRequestIdByRequestQIndex", _index)
	return *ret0, err
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(_index uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestIdByRequestQIndex(&_ZapGetters.CallOpts, _index)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(_index uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestIdByRequestQIndex(&_ZapGetters.CallOpts, _index)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(_timestamp uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetRequestIdByTimestamp(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getRequestIdByTimestamp", _timestamp)
	return *ret0, err
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(_timestamp uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestIdByTimestamp(&_ZapGetters.CallOpts, _timestamp)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(_timestamp uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestIdByTimestamp(&_ZapGetters.CallOpts, _timestamp)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() constant returns(uint256[51])
func (_ZapGetters *ZapGettersCaller) GetRequestQ(opts *bind.CallOpts) ([51]*big.Int, error) {
	var (
		ret0 = new([51]*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getRequestQ")
	return *ret0, err
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() constant returns(uint256[51])
func (_ZapGetters *ZapGettersSession) GetRequestQ() ([51]*big.Int, error) {
	return _ZapGetters.Contract.GetRequestQ(&_ZapGetters.CallOpts)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() constant returns(uint256[51])
func (_ZapGetters *ZapGettersCallerSession) GetRequestQ() ([51]*big.Int, error) {
	return _ZapGetters.Contract.GetRequestQ(&_ZapGetters.CallOpts)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(_requestId uint256, _data bytes32) constant returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetRequestUintVars(opts *bind.CallOpts, _requestId *big.Int, _data [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getRequestUintVars", _requestId, _data)
	return *ret0, err
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(_requestId uint256, _data bytes32) constant returns(uint256)
func (_ZapGetters *ZapGettersSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestUintVars(&_ZapGetters.CallOpts, _requestId, _data)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(_requestId uint256, _data bytes32) constant returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetRequestUintVars(&_ZapGetters.CallOpts, _requestId, _data)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(_requestId uint256) constant returns(string, string, bytes32, uint256, uint256, uint256)
func (_ZapGetters *ZapGettersCaller) GetRequestVars(opts *bind.CallOpts, _requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new([32]byte)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _ZapGetters.contract.Call(opts, out, "getRequestVars", _requestId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(_requestId uint256) constant returns(string, string, bytes32, uint256, uint256, uint256)
func (_ZapGetters *ZapGettersSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetRequestVars(&_ZapGetters.CallOpts, _requestId)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(_requestId uint256) constant returns(string, string, bytes32, uint256, uint256, uint256)
func (_ZapGetters *ZapGettersCallerSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetRequestVars(&_ZapGetters.CallOpts, _requestId)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(_staker address) constant returns(uint256, uint256)
func (_ZapGetters *ZapGettersCaller) GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ZapGetters.contract.Call(opts, out, "getStakerInfo", _staker)
	return *ret0, *ret1, err
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(_staker address) constant returns(uint256, uint256)
func (_ZapGetters *ZapGettersSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetStakerInfo(&_ZapGetters.CallOpts, _staker)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(_staker address) constant returns(uint256, uint256)
func (_ZapGetters *ZapGettersCallerSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _ZapGetters.Contract.GetStakerInfo(&_ZapGetters.CallOpts, _staker)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(_requestId uint256, _timestamp uint256) constant returns(uint256[5])
func (_ZapGetters *ZapGettersCaller) GetSubmissionsByTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	var (
		ret0 = new([5]*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getSubmissionsByTimestamp", _requestId, _timestamp)
	return *ret0, err
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(_requestId uint256, _timestamp uint256) constant returns(uint256[5])
func (_ZapGetters *ZapGettersSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _ZapGetters.Contract.GetSubmissionsByTimestamp(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(_requestId uint256, _timestamp uint256) constant returns(uint256[5])
func (_ZapGetters *ZapGettersCallerSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _ZapGetters.Contract.GetSubmissionsByTimestamp(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() constant returns(string)
func (_ZapGetters *ZapGettersCaller) GetSymbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getSymbol")
	return *ret0, err
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() constant returns(string)
func (_ZapGetters *ZapGettersSession) GetSymbol() (string, error) {
	return _ZapGetters.Contract.GetSymbol(&_ZapGetters.CallOpts)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() constant returns(string)
func (_ZapGetters *ZapGettersCallerSession) GetSymbol() (string, error) {
	return _ZapGetters.Contract.GetSymbol(&_ZapGetters.CallOpts)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(_requestID uint256, _index uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getTimestampbyRequestIDandIndex", _requestID, _index)
	return *ret0, err
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(_requestID uint256, _index uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetTimestampbyRequestIDandIndex(&_ZapGetters.CallOpts, _requestID, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(_requestID uint256, _index uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.GetTimestampbyRequestIDandIndex(&_ZapGetters.CallOpts, _requestID, _index)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(_data bytes32) constant returns(uint256)
func (_ZapGetters *ZapGettersCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "getUintVar", _data)
	return *ret0, err
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(_data bytes32) constant returns(uint256)
func (_ZapGetters *ZapGettersSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetUintVar(&_ZapGetters.CallOpts, _data)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(_data bytes32) constant returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _ZapGetters.Contract.GetUintVar(&_ZapGetters.CallOpts, _data)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() constant returns(uint256, uint256, string)
func (_ZapGetters *ZapGettersCaller) GetVariablesOnDeck(opts *bind.CallOpts) (*big.Int, *big.Int, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _ZapGetters.contract.Call(opts, out, "getVariablesOnDeck")
	return *ret0, *ret1, *ret2, err
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() constant returns(uint256, uint256, string)
func (_ZapGetters *ZapGettersSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _ZapGetters.Contract.GetVariablesOnDeck(&_ZapGetters.CallOpts)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() constant returns(uint256, uint256, string)
func (_ZapGetters *ZapGettersCallerSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _ZapGetters.Contract.GetVariablesOnDeck(&_ZapGetters.CallOpts)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(_requestId uint256, _timestamp uint256) constant returns(bool)
func (_ZapGetters *ZapGettersCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "isInDispute", _requestId, _timestamp)
	return *ret0, err
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(_requestId uint256, _timestamp uint256) constant returns(bool)
func (_ZapGetters *ZapGettersSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _ZapGetters.Contract.IsInDispute(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(_requestId uint256, _timestamp uint256) constant returns(bool)
func (_ZapGetters *ZapGettersCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _ZapGetters.Contract.IsInDispute(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "retrieveData", _requestId, _timestamp)
	return *ret0, err
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.RetrieveData(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapGetters.Contract.RetrieveData(&_ZapGetters.CallOpts, _requestId, _timestamp)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ZapGetters *ZapGettersCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapGetters.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ZapGetters *ZapGettersSession) TotalSupply() (*big.Int, error) {
	return _ZapGetters.Contract.TotalSupply(&_ZapGetters.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ZapGetters *ZapGettersCallerSession) TotalSupply() (*big.Int, error) {
	return _ZapGetters.Contract.TotalSupply(&_ZapGetters.CallOpts)
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

// ZapMasterABI is the input ABI used to generate the binding from.
const ZapMasterABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getSubmissionsByTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getAddressVars\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSymbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVariablesOnDeck\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_request\",\"type\":\"bytes32\"}],\"name\":\"getRequestIdByQueryHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getLastNewValueById\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDeity\",\"type\":\"address\"}],\"name\":\"changeDeity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balanceOfAt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getUintVar\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRequestIdByRequestQIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_challenge\",\"type\":\"bytes32\"},{\"name\":\"_miner\",\"type\":\"address\"}],\"name\":\"didMine\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinersByRequestIdAndTimestamp\",\"outputs\":[{\"name\":\"\",\"type\":\"address[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestID\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getDisputeUintVars\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"allowedToTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentVariables\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"didVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_zapContract\",\"type\":\"address\"}],\"name\":\"changeZapContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_disputeId\",\"type\":\"uint256\"}],\"name\":\"getAllDisputeVars\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256[9]\"},{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRequestQ\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[51]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinedBlockNum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"getDisputeIdByDisputeHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes32\"}],\"name\":\"getRequestUintVars\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getRequestVars\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastNewValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_zapContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_newZap\",\"type\":\"address\"}],\"name\":\"NewZapAddress\",\"type\":\"event\"}]"

// ZapMasterBin is the compiled bytecode used for deploying new contracts.
const ZapMasterBin = `0x608060405234801561001057600080fd5b506040516120123803806120128339818101604052602081101561003357600080fd5b5051604080517f4601f1cd000000000000000000000000000000000000000000000000000000008152600060048201819052915173__$799602413129f49037f52758954cf5aa52$__92634601f1cd9260248082019391829003018186803b15801561009e57600080fd5b505af41580156100b2573d6000803e3d6000fd5b5050604080517f5f6f776e657200000000000000000000000000000000000000000000000000008152815190819003600690810182206000908152603f602081815285832080546001600160a01b0319908116339081179092557f5f646569747900000000000000000000000000000000000000000000000000008752875196879003909501862084528282528684208054861690911790557f74656c6c6f72436f6e74726163740000000000000000000000000000000000008552855194859003600e01852083529081529084902080546001600160a01b03891693168317905590825291517fc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d9450908190039091019150a150611e3c806101d66000396000f3fe6080604052600436106101f95760003560e01c806370a082311161010d578063ae0a8279116100a0578063da3799411161006f578063da37994114610a3e578063dd62ed3e14610a68578063e0ae93c114610aa3578063e1eee6d614610ad3578063fc7cf0a014610bf7576101f9565b8063ae0a8279146108ff578063af0b132714610932578063b5413029146109e3578063c775b54214610a0e576101f9565b806393fa4915116100dc57806393fa4915146107ad578063999cf26c146107dd578063a22e407a14610816578063a7c438bc146108c6576101f9565b806370a08231146106ce578063733bdef01461070157806377fbb6631461074d5780637f6fd5d91461077d576101f9565b80633180f8df116101905780634ee2cd7e1161015f5780634ee2cd7e146105d8578063612c8f7f146106115780636173c0b81461063b57806363bb82ad1461066557806369026d631461069e576101f9565b80633180f8df146104f25780633df0777b1461053557806346eee1c41461057957806347abd7f1146105a3576101f9565b806317d7de7c116101cc57806317d7de7c1461040357806318160ddd1461041857806319e8e03b1461042d5780631db842f0146104c8576101f9565b80630f0b424d1461028f57806311c98512146102cb578063133bee5e146103335780631507040114610379575b604080516d1d195b1b1bdc90dbdb9d1c9858dd60921b8152815190819003600e0181206000908152603f602090815283822054601f369081018390048302850183019095528484526001600160a01b03169360609392918190840183828082843760009201829052508451949550938493509150506020840185600019f43d604051816000823e82801561028b578282f35b8282fd5b34801561029b57600080fd5b506102b9600480360360208110156102b257600080fd5b5035610c0c565b60408051918252519081900360200190f35b3480156102d757600080fd5b506102fb600480360360408110156102ee57600080fd5b5080359060200135610c24565b604051808260a080838360005b83811015610320578181015183820152602001610308565b5050505090500191505060405180910390f35b34801561033f57600080fd5b5061035d6004803603602081101561035657600080fd5b5035610c45565b604080516001600160a01b039092168252519081900360200190f35b34801561038557600080fd5b5061038e610c57565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103c85781810151838201526020016103b0565b50505050905090810190601f1680156103f55780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561040f57600080fd5b5061038e610c68565b34801561042457600080fd5b506102b9610c74565b34801561043957600080fd5b50610442610c80565b6040518084815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561048b578181015183820152602001610473565b50505050905090810190601f1680156104b85780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b3480156104d457600080fd5b506102b9600480360360208110156104eb57600080fd5b5035610c9a565b3480156104fe57600080fd5b5061051c6004803603602081101561051557600080fd5b5035610cac565b6040805192835290151560208301528051918290030190f35b34801561054157600080fd5b506105656004803603604081101561055857600080fd5b5080359060200135610cc8565b604080519115158252519081900360200190f35b34801561058557600080fd5b506102b96004803603602081101561059c57600080fd5b5035610cdb565b3480156105af57600080fd5b506105d6600480360360208110156105c657600080fd5b50356001600160a01b0316610ced565b005b3480156105e457600080fd5b506102b9600480360360408110156105fb57600080fd5b506001600160a01b038135169060200135610d01565b34801561061d57600080fd5b506102b96004803603602081101561063457600080fd5b5035610d9e565b34801561064757600080fd5b506102b96004803603602081101561065e57600080fd5b5035610db0565b34801561067157600080fd5b506105656004803603604081101561068857600080fd5b50803590602001356001600160a01b0316610dc2565b3480156106aa57600080fd5b506102fb600480360360408110156106c157600080fd5b5080359060200135610dd5565b3480156106da57600080fd5b506102b9600480360360208110156106f157600080fd5b50356001600160a01b0316610def565b34801561070d57600080fd5b506107346004803603602081101561072457600080fd5b50356001600160a01b0316610e84565b6040805192835260208301919091528051918290030190f35b34801561075957600080fd5b506102b96004803603604081101561077057600080fd5b5080359060200135610e97565b34801561078957600080fd5b506102b9600480360360408110156107a057600080fd5b5080359060200135610eaa565b3480156107b957600080fd5b506102b9600480360360408110156107d057600080fd5b5080359060200135610ebd565b3480156107e957600080fd5b506105656004803603604081101561080057600080fd5b506001600160a01b038135169060200135610ed0565b34801561082257600080fd5b5061082b610f3a565b6040518087815260200186815260200185815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b8381101561088657818101518382015260200161086e565b50505050905090810190601f1680156108b35780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b3480156108d257600080fd5b50610565600480360360408110156108e957600080fd5b50803590602001356001600160a01b0316610f61565b34801561090b57600080fd5b506105d66004803603602081101561092257600080fd5b50356001600160a01b0316610f74565b34801561093e57600080fd5b5061095c6004803603602081101561095557600080fd5b5035610f85565b604080518a815289151560208201528815159181019190915286151560608201526001600160a01b03808716608083015285811660a0830152841660c082015260e081018361012080838360005b838110156109c25781810151838201526020016109aa565b50505050905001828152602001995050505050505050505060405180910390f35b3480156109ef57600080fd5b506109f8610fc9565b6040518151815280826106608083836020610308565b348015610a1a57600080fd5b506102b960048036036040811015610a3157600080fd5b5080359060200135610fdb565b348015610a4a57600080fd5b506102b960048036036020811015610a6157600080fd5b5035610fee565b348015610a7457600080fd5b506102b960048036036040811015610a8b57600080fd5b506001600160a01b0381358116916020013516611000565b348015610aaf57600080fd5b506102b960048036036040811015610ac657600080fd5b508035906020013561106b565b348015610adf57600080fd5b50610afd60048036036020811015610af657600080fd5b503561107e565b604051808060200180602001878152602001868152602001858152602001848152602001838103835289818151815260200191508051906020019080838360005b83811015610b56578181015183820152602001610b3e565b50505050905090810190601f168015610b835780820380516001836020036101000a031916815260200191505b5083810382528851815288516020918201918a019080838360005b83811015610bb6578181015183820152602001610b9e565b50505050905090810190601f168015610be35780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b348015610c0357600080fd5b5061051c6110aa565b6000610c1e818363ffffffff6110bf16565b92915050565b610c2c611dab565b610c3e6000848463ffffffff6110d516565b9392505050565b6000610c1e818363ffffffff61112f16565b6060610c63600061114e565b905090565b6060610c63600061116b565b6000610c636000611195565b6000806060610c8f60006111c8565b925092509250909192565b6000610c1e818363ffffffff6112b416565b600080610cbf818463ffffffff6112ca16565b91509150915091565b6000610c3e81848463ffffffff61133016565b6000610c1e818363ffffffff61135716565b610cfe60008263ffffffff61137016565b50565b60408051633f48b1ff60e01b81526000600482018190526001600160a01b038516602483015260448201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__91633f48b1ff916064808301926020929190829003018186803b158015610d6b57600080fd5b505af4158015610d7f573d6000803e3d6000fd5b505050506040513d6020811015610d9557600080fd5b50519392505050565b6000610c1e818363ffffffff6113f916565b6000610c1e818363ffffffff61140b16565b6000610c3e81848463ffffffff61143216565b610ddd611dab565b610c3e6000848463ffffffff61145f16565b604080516393b182b360e01b81526000600482018190526001600160a01b0384166024830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__916393b182b3916044808301926020929190829003018186803b158015610e5257600080fd5b505af4158015610e66573d6000803e3d6000fd5b505050506040513d6020811015610e7c57600080fd5b505192915050565b600080610cbf818463ffffffff6114c316565b6000610c3e81848463ffffffff6114ea16565b6000610c3e81848463ffffffff61151d16565b6000610c3e81848463ffffffff61154116565b604080516356555cf160e11b81526000600482018190526001600160a01b038516602483015260448201849052915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163acaab9e2916064808301926020929190829003018186803b158015610d6b57600080fd5b60008060006060600080610f4e6000611565565b949b939a50919850965094509092509050565b6000610c3e81848463ffffffff61172316565b610cfe60008263ffffffff61175516565b6000806000806000806000610f98611dc9565b6000610faa818b63ffffffff61181716565b9850985098509850985098509850985098509193959799909294969850565b610fd1611de8565b610c636000611a2c565b6000610c3e81848463ffffffff611a6c16565b6000610c1e818363ffffffff611a9016565b60408051632fcc801b60e21b81526000600482018190526001600160a01b03808616602484015284166044830152915173__$ff8037bb7c49d17e1e79a1c3f1e9fdeb7a$__9163bf32006c916064808301926020929190829003018186803b158015610d6b57600080fd5b6000610c3e81848463ffffffff611aa616565b6060806000808080611096818863ffffffff611aca16565b949c939b5091995097509550909350915050565b6000806110b76000611ca5565b915091509091565b6000908152604291909101602052604090205490565b6110dd611dab565b6000838152604885016020908152604080832085845260090190915290819020815160a08101928390529160059082845b81548152602001906001019080831161110e57505050505090509392505050565b6000908152603f9190910160205260409020546001600160a01b031690565b50604080518082019091526002815261151560f21b602082015290565b5060408051808201909152600f81526e54656c6c6f7220547269627574657360881b602082015290565b604080516b746f74616c5f737570706c7960a01b8152815190819003600c01902060009081528183016020522054919050565b600080606060006111d885611d1b565b600081815260488701602081815260408084208151670746f74616c5469760c41b8152825160089181900391909101812086526004820184528286205495879052938352805460026001821615610100026000190190911604601f81018490048402850184019092528184529495508594929183919083018282801561129f5780601f106112745761010080835404028352916020019161129f565b820191906000526020600020905b81548152906001019060200180831161128257829003601f168201915b50505050509050935093509350509193909250565b6000908152604991909101602052604090205490565b6000818152604883016020526040812060038101548291901561132057600381018054611314918791879190600019810190811061130457fe5b9060005260206000200154611541565b60019250925050611329565b50600091508190505b9250929050565b60009182526048929092016020908152604080832093835260079093019052205460ff1690565b6000908152604891909101602052604090206003015490565b60408051655f646569747960d01b815281519081900360060190206000908152603f840160205220546001600160a01b031633146113ad57600080fd5b60408051655f646569747960d01b815281519081900360060190206000908152603f90930160205290912080546001600160a01b039092166001600160a01b0319909216919091179055565b60009081526040918201602052205490565b6000603282111561141b57600080fd5b506000908152604391909101602052604090205490565b6000918252604192909201602090815260408083206001600160a01b039094168352929052205460ff1690565b611467611dab565b6000838152604885016020908152604080832085845260080190915290819020815160a08101928390529160059082845b81546001600160a01b0316815260019091019060200180831161149857505050505090509392505050565b6001600160a01b031660009081526047919091016020526040902080546001909101549091565b6000828152604884016020526040812060030180548390811061150957fe5b906000526020600020015490509392505050565b60009182526044929092016020908152604080832093835260059093019052205490565b60009182526048929092016020908152604080832093835260069093019052205490565b8054604080516f18dd5c9c995b9d14995c5d595cdd125960821b80825282519182900360109081018320600090815284870160208181528683205469646966666963756c747960b01b8752875196879003600a01872084528282528784205486885288519788900386018820855283835288852054855260488b01808452898620888a528a51998a900388018a2087528585528a87205487528185528a87206a6772616e756c617269747960a81b8b528b519a8b9003600b018b208852600490810186528b882054998b528b519a8b90039098018a2087529484528986205486528352888520670746f74616c5469760c41b8952895198899003600801892086529095018252878420548354601f600260001961010060018516150201909216919091049081018490048402890184019099528888529398899889986060988a98899894979596909594919391929185918301828280156117075780601f106116dc57610100808354040283529160200191611707565b820191906000526020600020905b8154815290600101906020018083116116ea57829003601f168201915b5050505050925095509550955095509550955091939550919395565b600082815260448401602090815260408083206001600160a01b038516845260060190915290205460ff169392505050565b60408051655f646569747960d01b815281519081900360060190206000908152603f840160205220546001600160a01b0316331461179257600080fd5b604080516d1d195b1b1bdc90dbdb9d1c9858dd60921b8152815190819003600e0181206000908152603f850160209081529083902080546001600160a01b0386166001600160a01b03199091168117909155825291517fc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d929181900390910190a15050565b600080600080600080600061182a611dc9565b5050506000868152604488016020908152604080832080546002820154600380840154600485015486516101208101808952681c995c5d595cdd125960ba1b905287518082036101290190208a526005808801808b52898c20548352895168074696d657374616d760bc1b81528a519081900360099081019091208d52818c528a8d2054848d01528a516476616c756560d81b81528b51908190039093019092208c52808b52898c2054838b015289516f6d696e457865637574696f6e4461746560801b81528a519081900360100190208c52808b52898c2054606084015289516c6e756d6265724f66566f74657360981b81528a5190819003600d0190208c52808b52898c2054608084015289516a313637b1b5a73ab6b132b960a91b81528a5190819003600b0190208c52808b52898c205460a08401528951681b5a5b995c94db1bdd60ba1b81528a51908190039092019091208b52808a52888b205460c083015288516571756f72756d60d01b815289519081900360060190208b52808a52888b205460e083015288516266656560e81b81528951908190039095019094208a5292909752949096205461010087810191909152600190930154919a5060ff8082169a509281048316985062010000810490921696506001600160a01b0363010000009092048216955091811693921691909295985092959850929598565b611a34611de8565b6040805161066081019182905290600184019060339082845b815481526020019060010190808311611a4d5750505050509050919050565b60009182526048929092016020908152604080832093835260059093019052205490565b6000908152604a91909101602052604090205490565b60009182526048929092016020908152604080832093835260049093019052205490565b6000818152604883016020908152604080832060028082015483516a6772616e756c617269747960a81b8152845190819003600b018120875260048401808752858820546f3932b8bab2b9ba28a837b9b4ba34b7b760811b83528651928390036010018320895281885286892054670746f74616c5469760c41b845287519384900360080184208a52918852868920548654601f6000196101006001848116159190910291909101909216979097049687018a90048a0285018a019098528584526060998a99909889988998899891978897938801969295929490918891830182828015611bf95780601f10611bce57610100808354040283529160200191611bf9565b820191906000526020600020905b815481529060010190602001808311611bdc57829003601f168201915b5050885460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959b508a945092508401905082828015611c875780601f10611c5c57610100808354040283529160200191611c87565b820191906000526020600020905b815481529060010190602001808311611c6a57829003601f168201915b50505050509450965096509650965096509650509295509295509295565b604080517174696d654f664c6173744e657756616c756560701b8082528251601292819003830181206000908152858501602081815286832054835260428801815286832054948452865193849003909501909220815292529181205490918291611d11918591611541565b9360019350915050565b6040805161066081019182905260009182918291611d5c9190600187019060339082845b815481526020019060010190808311611d3f575050505050611d77565b60009081526043909501602052505060409092205492915050565b60008060005b6033811015611da5576020810284015180841015611d9c578093508192505b50600101611d7d565b50915091565b6040518060a001604052806005906020820280388339509192915050565b6040518061012001604052806009906020820280388339509192915050565b604051806106600160405280603390602082028038833950919291505056fea265627a7a723058208f825b37f40114c9f8689b2b8792c2aa6e313b233bb3f67a270a1197636999e964736f6c634300050a0032`

// DeployZapMaster deploys a new Ethereum contract, binding an instance of ZapMaster to it.
func DeployZapMaster(auth *bind.TransactOpts, backend bind.ContractBackend, _zapContract common.Address) (common.Address, *types.Transaction, *ZapMaster, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapMasterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ZapMasterBin), backend, _zapContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZapMaster{ZapMasterCaller: ZapMasterCaller{contract: contract}, ZapMasterTransactor: ZapMasterTransactor{contract: contract}, ZapMasterFilterer: ZapMasterFilterer{contract: contract}}, nil
}

// ZapMaster is an auto generated Go binding around an Ethereum contract.
type ZapMaster struct {
	ZapMasterCaller     // Read-only binding to the contract
	ZapMasterTransactor // Write-only binding to the contract
	ZapMasterFilterer   // Log filterer for contract events
}

// ZapMasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZapMasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapMasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZapMasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapMasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZapMasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZapMasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZapMasterSession struct {
	Contract     *ZapMaster        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZapMasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZapMasterCallerSession struct {
	Contract *ZapMasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ZapMasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZapMasterTransactorSession struct {
	Contract     *ZapMasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ZapMasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZapMasterRaw struct {
	Contract *ZapMaster // Generic contract binding to access the raw methods on
}

// ZapMasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZapMasterCallerRaw struct {
	Contract *ZapMasterCaller // Generic read-only contract binding to access the raw methods on
}

// ZapMasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZapMasterTransactorRaw struct {
	Contract *ZapMasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZapMaster creates a new instance of ZapMaster, bound to a specific deployed contract.
func NewZapMaster(address common.Address, backend bind.ContractBackend) (*ZapMaster, error) {
	contract, err := bindZapMaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZapMaster{ZapMasterCaller: ZapMasterCaller{contract: contract}, ZapMasterTransactor: ZapMasterTransactor{contract: contract}, ZapMasterFilterer: ZapMasterFilterer{contract: contract}}, nil
}

// NewZapMasterCaller creates a new read-only instance of ZapMaster, bound to a specific deployed contract.
func NewZapMasterCaller(address common.Address, caller bind.ContractCaller) (*ZapMasterCaller, error) {
	contract, err := bindZapMaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZapMasterCaller{contract: contract}, nil
}

// NewZapMasterTransactor creates a new write-only instance of ZapMaster, bound to a specific deployed contract.
func NewZapMasterTransactor(address common.Address, transactor bind.ContractTransactor) (*ZapMasterTransactor, error) {
	contract, err := bindZapMaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZapMasterTransactor{contract: contract}, nil
}

// NewZapMasterFilterer creates a new log filterer instance of ZapMaster, bound to a specific deployed contract.
func NewZapMasterFilterer(address common.Address, filterer bind.ContractFilterer) (*ZapMasterFilterer, error) {
	contract, err := bindZapMaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZapMasterFilterer{contract: contract}, nil
}

// bindZapMaster binds a generic wrapper to an already deployed contract.
func bindZapMaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZapMasterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapMaster *ZapMasterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZapMaster.Contract.ZapMasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapMaster *ZapMasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapMaster.Contract.ZapMasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapMaster *ZapMasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapMaster.Contract.ZapMasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZapMaster *ZapMasterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ZapMaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZapMaster *ZapMasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZapMaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZapMaster *ZapMasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZapMaster.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_user address, _spender address) constant returns(uint256)
func (_ZapMaster *ZapMasterCaller) Allowance(opts *bind.CallOpts, _user common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "allowance", _user, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_user address, _spender address) constant returns(uint256)
func (_ZapMaster *ZapMasterSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _ZapMaster.Contract.Allowance(&_ZapMaster.CallOpts, _user, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_user address, _spender address) constant returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) Allowance(_user common.Address, _spender common.Address) (*big.Int, error) {
	return _ZapMaster.Contract.Allowance(&_ZapMaster.CallOpts, _user, _spender)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(_user address, _amount uint256) constant returns(bool)
func (_ZapMaster *ZapMasterCaller) AllowedToTrade(opts *bind.CallOpts, _user common.Address, _amount *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "allowedToTrade", _user, _amount)
	return *ret0, err
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(_user address, _amount uint256) constant returns(bool)
func (_ZapMaster *ZapMasterSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _ZapMaster.Contract.AllowedToTrade(&_ZapMaster.CallOpts, _user, _amount)
}

// AllowedToTrade is a free data retrieval call binding the contract method 0x999cf26c.
//
// Solidity: function allowedToTrade(_user address, _amount uint256) constant returns(bool)
func (_ZapMaster *ZapMasterCallerSession) AllowedToTrade(_user common.Address, _amount *big.Int) (bool, error) {
	return _ZapMaster.Contract.AllowedToTrade(&_ZapMaster.CallOpts, _user, _amount)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_user address) constant returns(uint256)
func (_ZapMaster *ZapMasterCaller) BalanceOf(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "balanceOf", _user)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_user address) constant returns(uint256)
func (_ZapMaster *ZapMasterSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _ZapMaster.Contract.BalanceOf(&_ZapMaster.CallOpts, _user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_user address) constant returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) BalanceOf(_user common.Address) (*big.Int, error) {
	return _ZapMaster.Contract.BalanceOf(&_ZapMaster.CallOpts, _user)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(_user address, _blockNumber uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterCaller) BalanceOfAt(opts *bind.CallOpts, _user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "balanceOfAt", _user, _blockNumber)
	return *ret0, err
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(_user address, _blockNumber uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.BalanceOfAt(&_ZapMaster.CallOpts, _user, _blockNumber)
}

// BalanceOfAt is a free data retrieval call binding the contract method 0x4ee2cd7e.
//
// Solidity: function balanceOfAt(_user address, _blockNumber uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) BalanceOfAt(_user common.Address, _blockNumber *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.BalanceOfAt(&_ZapMaster.CallOpts, _user, _blockNumber)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(_challenge bytes32, _miner address) constant returns(bool)
func (_ZapMaster *ZapMasterCaller) DidMine(opts *bind.CallOpts, _challenge [32]byte, _miner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "didMine", _challenge, _miner)
	return *ret0, err
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(_challenge bytes32, _miner address) constant returns(bool)
func (_ZapMaster *ZapMasterSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _ZapMaster.Contract.DidMine(&_ZapMaster.CallOpts, _challenge, _miner)
}

// DidMine is a free data retrieval call binding the contract method 0x63bb82ad.
//
// Solidity: function didMine(_challenge bytes32, _miner address) constant returns(bool)
func (_ZapMaster *ZapMasterCallerSession) DidMine(_challenge [32]byte, _miner common.Address) (bool, error) {
	return _ZapMaster.Contract.DidMine(&_ZapMaster.CallOpts, _challenge, _miner)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(_disputeId uint256, _address address) constant returns(bool)
func (_ZapMaster *ZapMasterCaller) DidVote(opts *bind.CallOpts, _disputeId *big.Int, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "didVote", _disputeId, _address)
	return *ret0, err
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(_disputeId uint256, _address address) constant returns(bool)
func (_ZapMaster *ZapMasterSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _ZapMaster.Contract.DidVote(&_ZapMaster.CallOpts, _disputeId, _address)
}

// DidVote is a free data retrieval call binding the contract method 0xa7c438bc.
//
// Solidity: function didVote(_disputeId uint256, _address address) constant returns(bool)
func (_ZapMaster *ZapMasterCallerSession) DidVote(_disputeId *big.Int, _address common.Address) (bool, error) {
	return _ZapMaster.Contract.DidVote(&_ZapMaster.CallOpts, _disputeId, _address)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(_data bytes32) constant returns(address)
func (_ZapMaster *ZapMasterCaller) GetAddressVars(opts *bind.CallOpts, _data [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getAddressVars", _data)
	return *ret0, err
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(_data bytes32) constant returns(address)
func (_ZapMaster *ZapMasterSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _ZapMaster.Contract.GetAddressVars(&_ZapMaster.CallOpts, _data)
}

// GetAddressVars is a free data retrieval call binding the contract method 0x133bee5e.
//
// Solidity: function getAddressVars(_data bytes32) constant returns(address)
func (_ZapMaster *ZapMasterCallerSession) GetAddressVars(_data [32]byte) (common.Address, error) {
	return _ZapMaster.Contract.GetAddressVars(&_ZapMaster.CallOpts, _data)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(_disputeId uint256) constant returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ZapMaster *ZapMasterCaller) GetAllDisputeVars(opts *bind.CallOpts, _disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(bool)
		ret2 = new(bool)
		ret3 = new(bool)
		ret4 = new(common.Address)
		ret5 = new(common.Address)
		ret6 = new(common.Address)
		ret7 = new([9]*big.Int)
		ret8 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
		ret7,
		ret8,
	}
	err := _ZapMaster.contract.Call(opts, out, "getAllDisputeVars", _disputeId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, *ret7, *ret8, err
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(_disputeId uint256) constant returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ZapMaster *ZapMasterSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetAllDisputeVars(&_ZapMaster.CallOpts, _disputeId)
}

// GetAllDisputeVars is a free data retrieval call binding the contract method 0xaf0b1327.
//
// Solidity: function getAllDisputeVars(_disputeId uint256) constant returns(bytes32, bool, bool, bool, address, address, address, uint256[9], int256)
func (_ZapMaster *ZapMasterCallerSession) GetAllDisputeVars(_disputeId *big.Int) ([32]byte, bool, bool, bool, common.Address, common.Address, common.Address, [9]*big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetAllDisputeVars(&_ZapMaster.CallOpts, _disputeId)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() constant returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapMaster *ZapMasterCaller) GetCurrentVariables(opts *bind.CallOpts) ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(string)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _ZapMaster.contract.Call(opts, out, "getCurrentVariables")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() constant returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapMaster *ZapMasterSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetCurrentVariables(&_ZapMaster.CallOpts)
}

// GetCurrentVariables is a free data retrieval call binding the contract method 0xa22e407a.
//
// Solidity: function getCurrentVariables() constant returns(bytes32, uint256, uint256, string, uint256, uint256)
func (_ZapMaster *ZapMasterCallerSession) GetCurrentVariables() ([32]byte, *big.Int, *big.Int, string, *big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetCurrentVariables(&_ZapMaster.CallOpts)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(_hash bytes32) constant returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetDisputeIdByDisputeHash(opts *bind.CallOpts, _hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getDisputeIdByDisputeHash", _hash)
	return *ret0, err
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(_hash bytes32) constant returns(uint256)
func (_ZapMaster *ZapMasterSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetDisputeIdByDisputeHash(&_ZapMaster.CallOpts, _hash)
}

// GetDisputeIdByDisputeHash is a free data retrieval call binding the contract method 0xda379941.
//
// Solidity: function getDisputeIdByDisputeHash(_hash bytes32) constant returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetDisputeIdByDisputeHash(_hash [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetDisputeIdByDisputeHash(&_ZapMaster.CallOpts, _hash)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(_disputeId uint256, _data bytes32) constant returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetDisputeUintVars(opts *bind.CallOpts, _disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getDisputeUintVars", _disputeId, _data)
	return *ret0, err
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(_disputeId uint256, _data bytes32) constant returns(uint256)
func (_ZapMaster *ZapMasterSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetDisputeUintVars(&_ZapMaster.CallOpts, _disputeId, _data)
}

// GetDisputeUintVars is a free data retrieval call binding the contract method 0x7f6fd5d9.
//
// Solidity: function getDisputeUintVars(_disputeId uint256, _data bytes32) constant returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetDisputeUintVars(_disputeId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetDisputeUintVars(&_ZapMaster.CallOpts, _disputeId, _data)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() constant returns(uint256, bool)
func (_ZapMaster *ZapMasterCaller) GetLastNewValue(opts *bind.CallOpts) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ZapMaster.contract.Call(opts, out, "getLastNewValue")
	return *ret0, *ret1, err
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() constant returns(uint256, bool)
func (_ZapMaster *ZapMasterSession) GetLastNewValue() (*big.Int, bool, error) {
	return _ZapMaster.Contract.GetLastNewValue(&_ZapMaster.CallOpts)
}

// GetLastNewValue is a free data retrieval call binding the contract method 0xfc7cf0a0.
//
// Solidity: function getLastNewValue() constant returns(uint256, bool)
func (_ZapMaster *ZapMasterCallerSession) GetLastNewValue() (*big.Int, bool, error) {
	return _ZapMaster.Contract.GetLastNewValue(&_ZapMaster.CallOpts)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(_requestId uint256) constant returns(uint256, bool)
func (_ZapMaster *ZapMasterCaller) GetLastNewValueById(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ZapMaster.contract.Call(opts, out, "getLastNewValueById", _requestId)
	return *ret0, *ret1, err
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(_requestId uint256) constant returns(uint256, bool)
func (_ZapMaster *ZapMasterSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _ZapMaster.Contract.GetLastNewValueById(&_ZapMaster.CallOpts, _requestId)
}

// GetLastNewValueById is a free data retrieval call binding the contract method 0x3180f8df.
//
// Solidity: function getLastNewValueById(_requestId uint256) constant returns(uint256, bool)
func (_ZapMaster *ZapMasterCallerSession) GetLastNewValueById(_requestId *big.Int) (*big.Int, bool, error) {
	return _ZapMaster.Contract.GetLastNewValueById(&_ZapMaster.CallOpts, _requestId)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetMinedBlockNum(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getMinedBlockNum", _requestId, _timestamp)
	return *ret0, err
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetMinedBlockNum(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// GetMinedBlockNum is a free data retrieval call binding the contract method 0xc775b542.
//
// Solidity: function getMinedBlockNum(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetMinedBlockNum(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetMinedBlockNum(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(_requestId uint256, _timestamp uint256) constant returns(address[5])
func (_ZapMaster *ZapMasterCaller) GetMinersByRequestIdAndTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	var (
		ret0 = new([5]common.Address)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getMinersByRequestIdAndTimestamp", _requestId, _timestamp)
	return *ret0, err
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(_requestId uint256, _timestamp uint256) constant returns(address[5])
func (_ZapMaster *ZapMasterSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _ZapMaster.Contract.GetMinersByRequestIdAndTimestamp(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// GetMinersByRequestIdAndTimestamp is a free data retrieval call binding the contract method 0x69026d63.
//
// Solidity: function getMinersByRequestIdAndTimestamp(_requestId uint256, _timestamp uint256) constant returns(address[5])
func (_ZapMaster *ZapMasterCallerSession) GetMinersByRequestIdAndTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]common.Address, error) {
	return _ZapMaster.Contract.GetMinersByRequestIdAndTimestamp(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_ZapMaster *ZapMasterCaller) GetName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getName")
	return *ret0, err
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_ZapMaster *ZapMasterSession) GetName() (string, error) {
	return _ZapMaster.Contract.GetName(&_ZapMaster.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() constant returns(string)
func (_ZapMaster *ZapMasterCallerSession) GetName() (string, error) {
	return _ZapMaster.Contract.GetName(&_ZapMaster.CallOpts)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(_requestId uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getNewValueCountbyRequestId", _requestId)
	return *ret0, err
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(_requestId uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetNewValueCountbyRequestId(&_ZapMaster.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(_requestId uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetNewValueCountbyRequestId(&_ZapMaster.CallOpts, _requestId)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(_request bytes32) constant returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetRequestIdByQueryHash(opts *bind.CallOpts, _request [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getRequestIdByQueryHash", _request)
	return *ret0, err
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(_request bytes32) constant returns(uint256)
func (_ZapMaster *ZapMasterSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestIdByQueryHash(&_ZapMaster.CallOpts, _request)
}

// GetRequestIdByQueryHash is a free data retrieval call binding the contract method 0x1db842f0.
//
// Solidity: function getRequestIdByQueryHash(_request bytes32) constant returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetRequestIdByQueryHash(_request [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestIdByQueryHash(&_ZapMaster.CallOpts, _request)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(_index uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetRequestIdByRequestQIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getRequestIdByRequestQIndex", _index)
	return *ret0, err
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(_index uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestIdByRequestQIndex(&_ZapMaster.CallOpts, _index)
}

// GetRequestIdByRequestQIndex is a free data retrieval call binding the contract method 0x6173c0b8.
//
// Solidity: function getRequestIdByRequestQIndex(_index uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetRequestIdByRequestQIndex(_index *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestIdByRequestQIndex(&_ZapMaster.CallOpts, _index)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(_timestamp uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetRequestIdByTimestamp(opts *bind.CallOpts, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getRequestIdByTimestamp", _timestamp)
	return *ret0, err
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(_timestamp uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestIdByTimestamp(&_ZapMaster.CallOpts, _timestamp)
}

// GetRequestIdByTimestamp is a free data retrieval call binding the contract method 0x0f0b424d.
//
// Solidity: function getRequestIdByTimestamp(_timestamp uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetRequestIdByTimestamp(_timestamp *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestIdByTimestamp(&_ZapMaster.CallOpts, _timestamp)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() constant returns(uint256[51])
func (_ZapMaster *ZapMasterCaller) GetRequestQ(opts *bind.CallOpts) ([51]*big.Int, error) {
	var (
		ret0 = new([51]*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getRequestQ")
	return *ret0, err
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() constant returns(uint256[51])
func (_ZapMaster *ZapMasterSession) GetRequestQ() ([51]*big.Int, error) {
	return _ZapMaster.Contract.GetRequestQ(&_ZapMaster.CallOpts)
}

// GetRequestQ is a free data retrieval call binding the contract method 0xb5413029.
//
// Solidity: function getRequestQ() constant returns(uint256[51])
func (_ZapMaster *ZapMasterCallerSession) GetRequestQ() ([51]*big.Int, error) {
	return _ZapMaster.Contract.GetRequestQ(&_ZapMaster.CallOpts)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(_requestId uint256, _data bytes32) constant returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetRequestUintVars(opts *bind.CallOpts, _requestId *big.Int, _data [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getRequestUintVars", _requestId, _data)
	return *ret0, err
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(_requestId uint256, _data bytes32) constant returns(uint256)
func (_ZapMaster *ZapMasterSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestUintVars(&_ZapMaster.CallOpts, _requestId, _data)
}

// GetRequestUintVars is a free data retrieval call binding the contract method 0xe0ae93c1.
//
// Solidity: function getRequestUintVars(_requestId uint256, _data bytes32) constant returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetRequestUintVars(_requestId *big.Int, _data [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetRequestUintVars(&_ZapMaster.CallOpts, _requestId, _data)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(_requestId uint256) constant returns(string, string, bytes32, uint256, uint256, uint256)
func (_ZapMaster *ZapMasterCaller) GetRequestVars(opts *bind.CallOpts, _requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new([32]byte)
		ret3 = new(*big.Int)
		ret4 = new(*big.Int)
		ret5 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _ZapMaster.contract.Call(opts, out, "getRequestVars", _requestId)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(_requestId uint256) constant returns(string, string, bytes32, uint256, uint256, uint256)
func (_ZapMaster *ZapMasterSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetRequestVars(&_ZapMaster.CallOpts, _requestId)
}

// GetRequestVars is a free data retrieval call binding the contract method 0xe1eee6d6.
//
// Solidity: function getRequestVars(_requestId uint256) constant returns(string, string, bytes32, uint256, uint256, uint256)
func (_ZapMaster *ZapMasterCallerSession) GetRequestVars(_requestId *big.Int) (string, string, [32]byte, *big.Int, *big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetRequestVars(&_ZapMaster.CallOpts, _requestId)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(_staker address) constant returns(uint256, uint256)
func (_ZapMaster *ZapMasterCaller) GetStakerInfo(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ZapMaster.contract.Call(opts, out, "getStakerInfo", _staker)
	return *ret0, *ret1, err
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(_staker address) constant returns(uint256, uint256)
func (_ZapMaster *ZapMasterSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetStakerInfo(&_ZapMaster.CallOpts, _staker)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(_staker address) constant returns(uint256, uint256)
func (_ZapMaster *ZapMasterCallerSession) GetStakerInfo(_staker common.Address) (*big.Int, *big.Int, error) {
	return _ZapMaster.Contract.GetStakerInfo(&_ZapMaster.CallOpts, _staker)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(_requestId uint256, _timestamp uint256) constant returns(uint256[5])
func (_ZapMaster *ZapMasterCaller) GetSubmissionsByTimestamp(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	var (
		ret0 = new([5]*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getSubmissionsByTimestamp", _requestId, _timestamp)
	return *ret0, err
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(_requestId uint256, _timestamp uint256) constant returns(uint256[5])
func (_ZapMaster *ZapMasterSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _ZapMaster.Contract.GetSubmissionsByTimestamp(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// GetSubmissionsByTimestamp is a free data retrieval call binding the contract method 0x11c98512.
//
// Solidity: function getSubmissionsByTimestamp(_requestId uint256, _timestamp uint256) constant returns(uint256[5])
func (_ZapMaster *ZapMasterCallerSession) GetSubmissionsByTimestamp(_requestId *big.Int, _timestamp *big.Int) ([5]*big.Int, error) {
	return _ZapMaster.Contract.GetSubmissionsByTimestamp(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() constant returns(string)
func (_ZapMaster *ZapMasterCaller) GetSymbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getSymbol")
	return *ret0, err
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() constant returns(string)
func (_ZapMaster *ZapMasterSession) GetSymbol() (string, error) {
	return _ZapMaster.Contract.GetSymbol(&_ZapMaster.CallOpts)
}

// GetSymbol is a free data retrieval call binding the contract method 0x15070401.
//
// Solidity: function getSymbol() constant returns(string)
func (_ZapMaster *ZapMasterCallerSession) GetSymbol() (string, error) {
	return _ZapMaster.Contract.GetSymbol(&_ZapMaster.CallOpts)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(_requestID uint256, _index uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestID *big.Int, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getTimestampbyRequestIDandIndex", _requestID, _index)
	return *ret0, err
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(_requestID uint256, _index uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetTimestampbyRequestIDandIndex(&_ZapMaster.CallOpts, _requestID, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(_requestID uint256, _index uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetTimestampbyRequestIDandIndex(_requestID *big.Int, _index *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.GetTimestampbyRequestIDandIndex(&_ZapMaster.CallOpts, _requestID, _index)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(_data bytes32) constant returns(uint256)
func (_ZapMaster *ZapMasterCaller) GetUintVar(opts *bind.CallOpts, _data [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "getUintVar", _data)
	return *ret0, err
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(_data bytes32) constant returns(uint256)
func (_ZapMaster *ZapMasterSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetUintVar(&_ZapMaster.CallOpts, _data)
}

// GetUintVar is a free data retrieval call binding the contract method 0x612c8f7f.
//
// Solidity: function getUintVar(_data bytes32) constant returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) GetUintVar(_data [32]byte) (*big.Int, error) {
	return _ZapMaster.Contract.GetUintVar(&_ZapMaster.CallOpts, _data)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() constant returns(uint256, uint256, string)
func (_ZapMaster *ZapMasterCaller) GetVariablesOnDeck(opts *bind.CallOpts) (*big.Int, *big.Int, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _ZapMaster.contract.Call(opts, out, "getVariablesOnDeck")
	return *ret0, *ret1, *ret2, err
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() constant returns(uint256, uint256, string)
func (_ZapMaster *ZapMasterSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _ZapMaster.Contract.GetVariablesOnDeck(&_ZapMaster.CallOpts)
}

// GetVariablesOnDeck is a free data retrieval call binding the contract method 0x19e8e03b.
//
// Solidity: function getVariablesOnDeck() constant returns(uint256, uint256, string)
func (_ZapMaster *ZapMasterCallerSession) GetVariablesOnDeck() (*big.Int, *big.Int, string, error) {
	return _ZapMaster.Contract.GetVariablesOnDeck(&_ZapMaster.CallOpts)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(_requestId uint256, _timestamp uint256) constant returns(bool)
func (_ZapMaster *ZapMasterCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "isInDispute", _requestId, _timestamp)
	return *ret0, err
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(_requestId uint256, _timestamp uint256) constant returns(bool)
func (_ZapMaster *ZapMasterSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _ZapMaster.Contract.IsInDispute(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(_requestId uint256, _timestamp uint256) constant returns(bool)
func (_ZapMaster *ZapMasterCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _ZapMaster.Contract.IsInDispute(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "retrieveData", _requestId, _timestamp)
	return *ret0, err
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.RetrieveData(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(_requestId uint256, _timestamp uint256) constant returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _ZapMaster.Contract.RetrieveData(&_ZapMaster.CallOpts, _requestId, _timestamp)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ZapMaster *ZapMasterCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ZapMaster.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ZapMaster *ZapMasterSession) TotalSupply() (*big.Int, error) {
	return _ZapMaster.Contract.TotalSupply(&_ZapMaster.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ZapMaster *ZapMasterCallerSession) TotalSupply() (*big.Int, error) {
	return _ZapMaster.Contract.TotalSupply(&_ZapMaster.CallOpts)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(_newDeity address) returns()
func (_ZapMaster *ZapMasterTransactor) ChangeDeity(opts *bind.TransactOpts, _newDeity common.Address) (*types.Transaction, error) {
	return _ZapMaster.contract.Transact(opts, "changeDeity", _newDeity)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(_newDeity address) returns()
func (_ZapMaster *ZapMasterSession) ChangeDeity(_newDeity common.Address) (*types.Transaction, error) {
	return _ZapMaster.Contract.ChangeDeity(&_ZapMaster.TransactOpts, _newDeity)
}

// ChangeDeity is a paid mutator transaction binding the contract method 0x47abd7f1.
//
// Solidity: function changeDeity(_newDeity address) returns()
func (_ZapMaster *ZapMasterTransactorSession) ChangeDeity(_newDeity common.Address) (*types.Transaction, error) {
	return _ZapMaster.Contract.ChangeDeity(&_ZapMaster.TransactOpts, _newDeity)
}

// ChangeZapContract is a paid mutator transaction binding the contract method 0xae0a8279.
//
// Solidity: function changeZapContract(_zapContract address) returns()
func (_ZapMaster *ZapMasterTransactor) ChangeZapContract(opts *bind.TransactOpts, _zapContract common.Address) (*types.Transaction, error) {
	return _ZapMaster.contract.Transact(opts, "changeZapContract", _zapContract)
}

// ChangeZapContract is a paid mutator transaction binding the contract method 0xae0a8279.
//
// Solidity: function changeZapContract(_zapContract address) returns()
func (_ZapMaster *ZapMasterSession) ChangeZapContract(_zapContract common.Address) (*types.Transaction, error) {
	return _ZapMaster.Contract.ChangeZapContract(&_ZapMaster.TransactOpts, _zapContract)
}

// ChangeZapContract is a paid mutator transaction binding the contract method 0xae0a8279.
//
// Solidity: function changeZapContract(_zapContract address) returns()
func (_ZapMaster *ZapMasterTransactorSession) ChangeZapContract(_zapContract common.Address) (*types.Transaction, error) {
	return _ZapMaster.Contract.ChangeZapContract(&_ZapMaster.TransactOpts, _zapContract)
}

// ZapMasterNewZapAddressIterator is returned from FilterNewZapAddress and is used to iterate over the raw logs and unpacked data for NewZapAddress events raised by the ZapMaster contract.
type ZapMasterNewZapAddressIterator struct {
	Event *ZapMasterNewZapAddress // Event containing the contract specifics and raw log

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
func (it *ZapMasterNewZapAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZapMasterNewZapAddress)
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
		it.Event = new(ZapMasterNewZapAddress)
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
func (it *ZapMasterNewZapAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZapMasterNewZapAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZapMasterNewZapAddress represents a NewZapAddress event raised by the ZapMaster contract.
type ZapMasterNewZapAddress struct {
	NewZap common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewZapAddress is a free log retrieval operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: e NewZapAddress(_newZap address)
func (_ZapMaster *ZapMasterFilterer) FilterNewZapAddress(opts *bind.FilterOpts) (*ZapMasterNewZapAddressIterator, error) {

	logs, sub, err := _ZapMaster.contract.FilterLogs(opts, "NewZapAddress")
	if err != nil {
		return nil, err
	}
	return &ZapMasterNewZapAddressIterator{contract: _ZapMaster.contract, event: "NewZapAddress", logs: logs, sub: sub}, nil
}

// WatchNewZapAddress is a free log subscription operation binding the contract event 0xc2d1449eb0b6547aa426e09d9942a77fa4fc8cd3296305b3163e22452e0bcb8d.
//
// Solidity: e NewZapAddress(_newZap address)
func (_ZapMaster *ZapMasterFilterer) WatchNewZapAddress(opts *bind.WatchOpts, sink chan<- *ZapMasterNewZapAddress) (event.Subscription, error) {

	logs, sub, err := _ZapMaster.contract.WatchLogs(opts, "NewZapAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZapMasterNewZapAddress)
				if err := _ZapMaster.contract.UnpackLog(event, "NewZapAddress", log); err != nil {
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
