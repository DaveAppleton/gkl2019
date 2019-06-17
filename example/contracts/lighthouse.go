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

// LighthouseABI is the input ABI used to generate the binding from.
const LighthouseABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"DataValue\",\"type\":\"uint256\"}],\"name\":\"write\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"peekData\",\"outputs\":[{\"name\":\"v\",\"type\":\"uint128\"},{\"name\":\"b\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"read\",\"outputs\":[{\"name\":\"x\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"peek\",\"outputs\":[{\"name\":\"v\",\"type\":\"bytes32\"},{\"name\":\"ok\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newMaxAge\",\"type\":\"uint256\"}],\"name\":\"setMaxAge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"peekUpdated\",\"outputs\":[{\"name\":\"v\",\"type\":\"uint32\"},{\"name\":\"b\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"notTooLongSinceUpdated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LighthouseBin is the compiled bytecode used for deploying new contracts.
const LighthouseBin = `0x608060405234801561001057600080fd5b506103c6806100206000396000f3006080604052600436106100825763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632f048afa8114610087578063420b81f6146100a157806357de26a4146100e257806359e02dd7146101095780635ae28fc914610137578063bdf384a81461014f578063e2f9063214610184575b600080fd5b34801561009357600080fd5b5061009f6004356101ad565b005b3480156100ad57600080fd5b506100b6610247565b604080516fffffffffffffffffffffffffffffffff909316835290151560208301528051918290030190f35b3480156100ee57600080fd5b506100f7610267565b60408051918252519081900360200190f35b34801561011557600080fd5b5061011e610304565b6040805192835290151560208301528051918290030190f35b34801561014357600080fd5b5061009f600435610323565b34801561015b57600080fd5b50610164610328565b6040805163ffffffff909316835290151560208301528051918290030190f35b34801561019057600080fd5b5061019961035e565b604080519115158252519081900360200190f35b70010000000000000000000000000000000081041561022d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f56616c756520746f6f206c617267650000000000000000000000000000000000604482015290519081900360640190fd5b700100000000000000000000000000000000420201600055565b600080549061025461035e565b8015610261575060005415155b90509091565b600061027161035e565b801561027e575060005415155b15156102eb57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f496e76616c696420646174612073746f72656400000000000000000000000000604482015290519081900360640190fd5b506000546fffffffffffffffffffffffffffffffff1690565b600080546fffffffffffffffffffffffffffffffff169061025461035e565b600155565b600080547001000000000000000000000000000000009004908161034a61035e565b8015610357575060005415155b9150509091565b6000805460015470010000000000000000000000000000000090910467ffffffffffffffff164203908110806103945750600154155b915050905600a165627a7a7230582030ff1f87fea4f5965d9d2a119da90542701c9f17921ad2237642de91954226dd0029`

// DeployLighthouse deploys a new Ethereum contract, binding an instance of Lighthouse to it.
func DeployLighthouse(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Lighthouse, error) {
	parsed, err := abi.JSON(strings.NewReader(LighthouseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LighthouseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lighthouse{LighthouseCaller: LighthouseCaller{contract: contract}, LighthouseTransactor: LighthouseTransactor{contract: contract}, LighthouseFilterer: LighthouseFilterer{contract: contract}}, nil
}

// Lighthouse is an auto generated Go binding around an Ethereum contract.
type Lighthouse struct {
	LighthouseCaller     // Read-only binding to the contract
	LighthouseTransactor // Write-only binding to the contract
	LighthouseFilterer   // Log filterer for contract events
}

// LighthouseCaller is an auto generated read-only Go binding around an Ethereum contract.
type LighthouseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LighthouseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LighthouseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LighthouseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LighthouseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LighthouseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LighthouseSession struct {
	Contract     *Lighthouse       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LighthouseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LighthouseCallerSession struct {
	Contract *LighthouseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LighthouseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LighthouseTransactorSession struct {
	Contract     *LighthouseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LighthouseRaw is an auto generated low-level Go binding around an Ethereum contract.
type LighthouseRaw struct {
	Contract *Lighthouse // Generic contract binding to access the raw methods on
}

// LighthouseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LighthouseCallerRaw struct {
	Contract *LighthouseCaller // Generic read-only contract binding to access the raw methods on
}

// LighthouseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LighthouseTransactorRaw struct {
	Contract *LighthouseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLighthouse creates a new instance of Lighthouse, bound to a specific deployed contract.
func NewLighthouse(address common.Address, backend bind.ContractBackend) (*Lighthouse, error) {
	contract, err := bindLighthouse(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lighthouse{LighthouseCaller: LighthouseCaller{contract: contract}, LighthouseTransactor: LighthouseTransactor{contract: contract}, LighthouseFilterer: LighthouseFilterer{contract: contract}}, nil
}

// NewLighthouseCaller creates a new read-only instance of Lighthouse, bound to a specific deployed contract.
func NewLighthouseCaller(address common.Address, caller bind.ContractCaller) (*LighthouseCaller, error) {
	contract, err := bindLighthouse(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LighthouseCaller{contract: contract}, nil
}

// NewLighthouseTransactor creates a new write-only instance of Lighthouse, bound to a specific deployed contract.
func NewLighthouseTransactor(address common.Address, transactor bind.ContractTransactor) (*LighthouseTransactor, error) {
	contract, err := bindLighthouse(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LighthouseTransactor{contract: contract}, nil
}

// NewLighthouseFilterer creates a new log filterer instance of Lighthouse, bound to a specific deployed contract.
func NewLighthouseFilterer(address common.Address, filterer bind.ContractFilterer) (*LighthouseFilterer, error) {
	contract, err := bindLighthouse(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LighthouseFilterer{contract: contract}, nil
}

// bindLighthouse binds a generic wrapper to an already deployed contract.
func bindLighthouse(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LighthouseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lighthouse *LighthouseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lighthouse.Contract.LighthouseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lighthouse *LighthouseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lighthouse.Contract.LighthouseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lighthouse *LighthouseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lighthouse.Contract.LighthouseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lighthouse *LighthouseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lighthouse.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lighthouse *LighthouseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lighthouse.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lighthouse *LighthouseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lighthouse.Contract.contract.Transact(opts, method, params...)
}

// NotTooLongSinceUpdated is a free data retrieval call binding the contract method 0xe2f90632.
//
// Solidity: function notTooLongSinceUpdated() constant returns(bool)
func (_Lighthouse *LighthouseCaller) NotTooLongSinceUpdated(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Lighthouse.contract.Call(opts, out, "notTooLongSinceUpdated")
	return *ret0, err
}

// NotTooLongSinceUpdated is a free data retrieval call binding the contract method 0xe2f90632.
//
// Solidity: function notTooLongSinceUpdated() constant returns(bool)
func (_Lighthouse *LighthouseSession) NotTooLongSinceUpdated() (bool, error) {
	return _Lighthouse.Contract.NotTooLongSinceUpdated(&_Lighthouse.CallOpts)
}

// NotTooLongSinceUpdated is a free data retrieval call binding the contract method 0xe2f90632.
//
// Solidity: function notTooLongSinceUpdated() constant returns(bool)
func (_Lighthouse *LighthouseCallerSession) NotTooLongSinceUpdated() (bool, error) {
	return _Lighthouse.Contract.NotTooLongSinceUpdated(&_Lighthouse.CallOpts)
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() constant returns(bytes32 v, bool ok)
func (_Lighthouse *LighthouseCaller) Peek(opts *bind.CallOpts) (struct {
	V  [32]byte
	Ok bool
}, error) {
	ret := new(struct {
		V  [32]byte
		Ok bool
	})
	out := ret
	err := _Lighthouse.contract.Call(opts, out, "peek")
	return *ret, err
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() constant returns(bytes32 v, bool ok)
func (_Lighthouse *LighthouseSession) Peek() (struct {
	V  [32]byte
	Ok bool
}, error) {
	return _Lighthouse.Contract.Peek(&_Lighthouse.CallOpts)
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() constant returns(bytes32 v, bool ok)
func (_Lighthouse *LighthouseCallerSession) Peek() (struct {
	V  [32]byte
	Ok bool
}, error) {
	return _Lighthouse.Contract.Peek(&_Lighthouse.CallOpts)
}

// PeekData is a free data retrieval call binding the contract method 0x420b81f6.
//
// Solidity: function peekData() constant returns(uint128 v, bool b)
func (_Lighthouse *LighthouseCaller) PeekData(opts *bind.CallOpts) (struct {
	V *big.Int
	B bool
}, error) {
	ret := new(struct {
		V *big.Int
		B bool
	})
	out := ret
	err := _Lighthouse.contract.Call(opts, out, "peekData")
	return *ret, err
}

// PeekData is a free data retrieval call binding the contract method 0x420b81f6.
//
// Solidity: function peekData() constant returns(uint128 v, bool b)
func (_Lighthouse *LighthouseSession) PeekData() (struct {
	V *big.Int
	B bool
}, error) {
	return _Lighthouse.Contract.PeekData(&_Lighthouse.CallOpts)
}

// PeekData is a free data retrieval call binding the contract method 0x420b81f6.
//
// Solidity: function peekData() constant returns(uint128 v, bool b)
func (_Lighthouse *LighthouseCallerSession) PeekData() (struct {
	V *big.Int
	B bool
}, error) {
	return _Lighthouse.Contract.PeekData(&_Lighthouse.CallOpts)
}

// PeekUpdated is a free data retrieval call binding the contract method 0xbdf384a8.
//
// Solidity: function peekUpdated() constant returns(uint32 v, bool b)
func (_Lighthouse *LighthouseCaller) PeekUpdated(opts *bind.CallOpts) (struct {
	V uint32
	B bool
}, error) {
	ret := new(struct {
		V uint32
		B bool
	})
	out := ret
	err := _Lighthouse.contract.Call(opts, out, "peekUpdated")
	return *ret, err
}

// PeekUpdated is a free data retrieval call binding the contract method 0xbdf384a8.
//
// Solidity: function peekUpdated() constant returns(uint32 v, bool b)
func (_Lighthouse *LighthouseSession) PeekUpdated() (struct {
	V uint32
	B bool
}, error) {
	return _Lighthouse.Contract.PeekUpdated(&_Lighthouse.CallOpts)
}

// PeekUpdated is a free data retrieval call binding the contract method 0xbdf384a8.
//
// Solidity: function peekUpdated() constant returns(uint32 v, bool b)
func (_Lighthouse *LighthouseCallerSession) PeekUpdated() (struct {
	V uint32
	B bool
}, error) {
	return _Lighthouse.Contract.PeekUpdated(&_Lighthouse.CallOpts)
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() constant returns(bytes32 x)
func (_Lighthouse *LighthouseCaller) Read(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Lighthouse.contract.Call(opts, out, "read")
	return *ret0, err
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() constant returns(bytes32 x)
func (_Lighthouse *LighthouseSession) Read() ([32]byte, error) {
	return _Lighthouse.Contract.Read(&_Lighthouse.CallOpts)
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() constant returns(bytes32 x)
func (_Lighthouse *LighthouseCallerSession) Read() ([32]byte, error) {
	return _Lighthouse.Contract.Read(&_Lighthouse.CallOpts)
}

// SetMaxAge is a paid mutator transaction binding the contract method 0x5ae28fc9.
//
// Solidity: function setMaxAge(uint256 newMaxAge) returns()
func (_Lighthouse *LighthouseTransactor) SetMaxAge(opts *bind.TransactOpts, newMaxAge *big.Int) (*types.Transaction, error) {
	return _Lighthouse.contract.Transact(opts, "setMaxAge", newMaxAge)
}

// SetMaxAge is a paid mutator transaction binding the contract method 0x5ae28fc9.
//
// Solidity: function setMaxAge(uint256 newMaxAge) returns()
func (_Lighthouse *LighthouseSession) SetMaxAge(newMaxAge *big.Int) (*types.Transaction, error) {
	return _Lighthouse.Contract.SetMaxAge(&_Lighthouse.TransactOpts, newMaxAge)
}

// SetMaxAge is a paid mutator transaction binding the contract method 0x5ae28fc9.
//
// Solidity: function setMaxAge(uint256 newMaxAge) returns()
func (_Lighthouse *LighthouseTransactorSession) SetMaxAge(newMaxAge *big.Int) (*types.Transaction, error) {
	return _Lighthouse.Contract.SetMaxAge(&_Lighthouse.TransactOpts, newMaxAge)
}

// Write is a paid mutator transaction binding the contract method 0x2f048afa.
//
// Solidity: function write(uint256 DataValue) returns()
func (_Lighthouse *LighthouseTransactor) Write(opts *bind.TransactOpts, DataValue *big.Int) (*types.Transaction, error) {
	return _Lighthouse.contract.Transact(opts, "write", DataValue)
}

// Write is a paid mutator transaction binding the contract method 0x2f048afa.
//
// Solidity: function write(uint256 DataValue) returns()
func (_Lighthouse *LighthouseSession) Write(DataValue *big.Int) (*types.Transaction, error) {
	return _Lighthouse.Contract.Write(&_Lighthouse.TransactOpts, DataValue)
}

// Write is a paid mutator transaction binding the contract method 0x2f048afa.
//
// Solidity: function write(uint256 DataValue) returns()
func (_Lighthouse *LighthouseTransactorSession) Write(DataValue *big.Int) (*types.Transaction, error) {
	return _Lighthouse.Contract.Write(&_Lighthouse.TransactOpts, DataValue)
}

// SearcherABI is the input ABI used to generate the binding from.
const SearcherABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"poke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"identify\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SearcherBin is the compiled bytecode used for deploying new contracts.
const SearcherBin = `0x`

// DeploySearcher deploys a new Ethereum contract, binding an instance of Searcher to it.
func DeploySearcher(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Searcher, error) {
	parsed, err := abi.JSON(strings.NewReader(SearcherABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SearcherBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Searcher{SearcherCaller: SearcherCaller{contract: contract}, SearcherTransactor: SearcherTransactor{contract: contract}, SearcherFilterer: SearcherFilterer{contract: contract}}, nil
}

// Searcher is an auto generated Go binding around an Ethereum contract.
type Searcher struct {
	SearcherCaller     // Read-only binding to the contract
	SearcherTransactor // Write-only binding to the contract
	SearcherFilterer   // Log filterer for contract events
}

// SearcherCaller is an auto generated read-only Go binding around an Ethereum contract.
type SearcherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SearcherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SearcherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SearcherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SearcherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SearcherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SearcherSession struct {
	Contract     *Searcher         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SearcherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SearcherCallerSession struct {
	Contract *SearcherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SearcherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SearcherTransactorSession struct {
	Contract     *SearcherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SearcherRaw is an auto generated low-level Go binding around an Ethereum contract.
type SearcherRaw struct {
	Contract *Searcher // Generic contract binding to access the raw methods on
}

// SearcherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SearcherCallerRaw struct {
	Contract *SearcherCaller // Generic read-only contract binding to access the raw methods on
}

// SearcherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SearcherTransactorRaw struct {
	Contract *SearcherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSearcher creates a new instance of Searcher, bound to a specific deployed contract.
func NewSearcher(address common.Address, backend bind.ContractBackend) (*Searcher, error) {
	contract, err := bindSearcher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Searcher{SearcherCaller: SearcherCaller{contract: contract}, SearcherTransactor: SearcherTransactor{contract: contract}, SearcherFilterer: SearcherFilterer{contract: contract}}, nil
}

// NewSearcherCaller creates a new read-only instance of Searcher, bound to a specific deployed contract.
func NewSearcherCaller(address common.Address, caller bind.ContractCaller) (*SearcherCaller, error) {
	contract, err := bindSearcher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SearcherCaller{contract: contract}, nil
}

// NewSearcherTransactor creates a new write-only instance of Searcher, bound to a specific deployed contract.
func NewSearcherTransactor(address common.Address, transactor bind.ContractTransactor) (*SearcherTransactor, error) {
	contract, err := bindSearcher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SearcherTransactor{contract: contract}, nil
}

// NewSearcherFilterer creates a new log filterer instance of Searcher, bound to a specific deployed contract.
func NewSearcherFilterer(address common.Address, filterer bind.ContractFilterer) (*SearcherFilterer, error) {
	contract, err := bindSearcher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SearcherFilterer{contract: contract}, nil
}

// bindSearcher binds a generic wrapper to an already deployed contract.
func bindSearcher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SearcherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Searcher *SearcherRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Searcher.Contract.SearcherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Searcher *SearcherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Searcher.Contract.SearcherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Searcher *SearcherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Searcher.Contract.SearcherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Searcher *SearcherCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Searcher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Searcher *SearcherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Searcher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Searcher *SearcherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Searcher.Contract.contract.Transact(opts, method, params...)
}

// Identify is a free data retrieval call binding the contract method 0xeeb72866.
//
// Solidity: function identify() constant returns(uint256)
func (_Searcher *SearcherCaller) Identify(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Searcher.contract.Call(opts, out, "identify")
	return *ret0, err
}

// Identify is a free data retrieval call binding the contract method 0xeeb72866.
//
// Solidity: function identify() constant returns(uint256)
func (_Searcher *SearcherSession) Identify() (*big.Int, error) {
	return _Searcher.Contract.Identify(&_Searcher.CallOpts)
}

// Identify is a free data retrieval call binding the contract method 0xeeb72866.
//
// Solidity: function identify() constant returns(uint256)
func (_Searcher *SearcherCallerSession) Identify() (*big.Int, error) {
	return _Searcher.Contract.Identify(&_Searcher.CallOpts)
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_Searcher *SearcherTransactor) Poke(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Searcher.contract.Transact(opts, "poke")
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_Searcher *SearcherSession) Poke() (*types.Transaction, error) {
	return _Searcher.Contract.Poke(&_Searcher.TransactOpts)
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_Searcher *SearcherTransactorSession) Poke() (*types.Transaction, error) {
	return _Searcher.Contract.Poke(&_Searcher.TransactOpts)
}
