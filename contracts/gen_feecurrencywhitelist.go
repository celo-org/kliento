// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/celo-org/celo-blockchain"
	"github.com/celo-org/celo-blockchain/accounts/abi"
	"github.com/celo-org/celo-blockchain/accounts/abi/bind"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/celo-blockchain/event"
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

// FeeCurrencyWhitelistABI is the input ABI used to generate the binding from.
const FeeCurrencyWhitelistABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWhitelist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FeeCurrencyWhitelist is an auto generated Go binding around an Ethereum contract.
type FeeCurrencyWhitelist struct {
	FeeCurrencyWhitelistCaller     // Read-only binding to the contract
	FeeCurrencyWhitelistTransactor // Write-only binding to the contract
	FeeCurrencyWhitelistFilterer   // Log filterer for contract events
}

// FeeCurrencyWhitelistCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeCurrencyWhitelistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeCurrencyWhitelistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeCurrencyWhitelistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeCurrencyWhitelistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeCurrencyWhitelistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeCurrencyWhitelistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeCurrencyWhitelistSession struct {
	Contract     *FeeCurrencyWhitelist // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FeeCurrencyWhitelistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeCurrencyWhitelistCallerSession struct {
	Contract *FeeCurrencyWhitelistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// FeeCurrencyWhitelistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeCurrencyWhitelistTransactorSession struct {
	Contract     *FeeCurrencyWhitelistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// FeeCurrencyWhitelistRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeCurrencyWhitelistRaw struct {
	Contract *FeeCurrencyWhitelist // Generic contract binding to access the raw methods on
}

// FeeCurrencyWhitelistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeCurrencyWhitelistCallerRaw struct {
	Contract *FeeCurrencyWhitelistCaller // Generic read-only contract binding to access the raw methods on
}

// FeeCurrencyWhitelistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeCurrencyWhitelistTransactorRaw struct {
	Contract *FeeCurrencyWhitelistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeCurrencyWhitelist creates a new instance of FeeCurrencyWhitelist, bound to a specific deployed contract.
func NewFeeCurrencyWhitelist(address common.Address, backend bind.ContractBackend) (*FeeCurrencyWhitelist, error) {
	contract, err := bindFeeCurrencyWhitelist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeCurrencyWhitelist{FeeCurrencyWhitelistCaller: FeeCurrencyWhitelistCaller{contract: contract}, FeeCurrencyWhitelistTransactor: FeeCurrencyWhitelistTransactor{contract: contract}, FeeCurrencyWhitelistFilterer: FeeCurrencyWhitelistFilterer{contract: contract}}, nil
}

// NewFeeCurrencyWhitelistCaller creates a new read-only instance of FeeCurrencyWhitelist, bound to a specific deployed contract.
func NewFeeCurrencyWhitelistCaller(address common.Address, caller bind.ContractCaller) (*FeeCurrencyWhitelistCaller, error) {
	contract, err := bindFeeCurrencyWhitelist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeCurrencyWhitelistCaller{contract: contract}, nil
}

// NewFeeCurrencyWhitelistTransactor creates a new write-only instance of FeeCurrencyWhitelist, bound to a specific deployed contract.
func NewFeeCurrencyWhitelistTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeCurrencyWhitelistTransactor, error) {
	contract, err := bindFeeCurrencyWhitelist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeCurrencyWhitelistTransactor{contract: contract}, nil
}

// NewFeeCurrencyWhitelistFilterer creates a new log filterer instance of FeeCurrencyWhitelist, bound to a specific deployed contract.
func NewFeeCurrencyWhitelistFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeCurrencyWhitelistFilterer, error) {
	contract, err := bindFeeCurrencyWhitelist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeCurrencyWhitelistFilterer{contract: contract}, nil
}

// bindFeeCurrencyWhitelist binds a generic wrapper to an already deployed contract.
func bindFeeCurrencyWhitelist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeCurrencyWhitelistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseFeeCurrencyWhitelistABI parses the ABI
func ParseFeeCurrencyWhitelistABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeCurrencyWhitelistABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FeeCurrencyWhitelist.Contract.FeeCurrencyWhitelistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeCurrencyWhitelist.Contract.FeeCurrencyWhitelistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeCurrencyWhitelist.Contract.FeeCurrencyWhitelistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FeeCurrencyWhitelist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeCurrencyWhitelist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeCurrencyWhitelist.Contract.contract.Transact(opts, method, params...)
}

// GetWhitelist is a free data retrieval call binding the contract method 0xd01f63f5.
//
// Solidity: function getWhitelist() view returns(address[])
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistCaller) GetWhitelist(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _FeeCurrencyWhitelist.contract.Call(opts, out, "getWhitelist")
	return *ret0, err
}

// GetWhitelist is a free data retrieval call binding the contract method 0xd01f63f5.
//
// Solidity: function getWhitelist() view returns(address[])
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistSession) GetWhitelist() ([]common.Address, error) {
	return _FeeCurrencyWhitelist.Contract.GetWhitelist(&_FeeCurrencyWhitelist.CallOpts)
}

// GetWhitelist is a free data retrieval call binding the contract method 0xd01f63f5.
//
// Solidity: function getWhitelist() view returns(address[])
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistCallerSession) GetWhitelist() ([]common.Address, error) {
	return _FeeCurrencyWhitelist.Contract.GetWhitelist(&_FeeCurrencyWhitelist.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FeeCurrencyWhitelist.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistSession) Initialized() (bool, error) {
	return _FeeCurrencyWhitelist.Contract.Initialized(&_FeeCurrencyWhitelist.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistCallerSession) Initialized() (bool, error) {
	return _FeeCurrencyWhitelist.Contract.Initialized(&_FeeCurrencyWhitelist.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FeeCurrencyWhitelist.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistSession) IsOwner() (bool, error) {
	return _FeeCurrencyWhitelist.Contract.IsOwner(&_FeeCurrencyWhitelist.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistCallerSession) IsOwner() (bool, error) {
	return _FeeCurrencyWhitelist.Contract.IsOwner(&_FeeCurrencyWhitelist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FeeCurrencyWhitelist.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistSession) Owner() (common.Address, error) {
	return _FeeCurrencyWhitelist.Contract.Owner(&_FeeCurrencyWhitelist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistCallerSession) Owner() (common.Address, error) {
	return _FeeCurrencyWhitelist.Contract.Owner(&_FeeCurrencyWhitelist.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist(uint256 ) view returns(address)
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistCaller) Whitelist(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FeeCurrencyWhitelist.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist(uint256 ) view returns(address)
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistSession) Whitelist(arg0 *big.Int) (common.Address, error) {
	return _FeeCurrencyWhitelist.Contract.Whitelist(&_FeeCurrencyWhitelist.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x7ebd1b30.
//
// Solidity: function whitelist(uint256 ) view returns(address)
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistCallerSession) Whitelist(arg0 *big.Int) (common.Address, error) {
	return _FeeCurrencyWhitelist.Contract.Whitelist(&_FeeCurrencyWhitelist.CallOpts, arg0)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address tokenAddress) returns()
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistTransactor) AddToken(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeCurrencyWhitelist.contract.Transact(opts, "addToken", tokenAddress)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address tokenAddress) returns()
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistSession) AddToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeCurrencyWhitelist.Contract.AddToken(&_FeeCurrencyWhitelist.TransactOpts, tokenAddress)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address tokenAddress) returns()
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistTransactorSession) AddToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeCurrencyWhitelist.Contract.AddToken(&_FeeCurrencyWhitelist.TransactOpts, tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeCurrencyWhitelist.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistSession) Initialize() (*types.Transaction, error) {
	return _FeeCurrencyWhitelist.Contract.Initialize(&_FeeCurrencyWhitelist.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistTransactorSession) Initialize() (*types.Transaction, error) {
	return _FeeCurrencyWhitelist.Contract.Initialize(&_FeeCurrencyWhitelist.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeCurrencyWhitelist.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeeCurrencyWhitelist.Contract.RenounceOwnership(&_FeeCurrencyWhitelist.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeeCurrencyWhitelist.Contract.RenounceOwnership(&_FeeCurrencyWhitelist.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FeeCurrencyWhitelist.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeCurrencyWhitelist.Contract.TransferOwnership(&_FeeCurrencyWhitelist.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeCurrencyWhitelist.Contract.TransferOwnership(&_FeeCurrencyWhitelist.TransactOpts, newOwner)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _FeeCurrencyWhitelist.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "OwnershipTransferred":
		event, err = _FeeCurrencyWhitelist.ParseOwnershipTransferred(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// FeeCurrencyWhitelistOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FeeCurrencyWhitelist contract.
type FeeCurrencyWhitelistOwnershipTransferredIterator struct {
	Event *FeeCurrencyWhitelistOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FeeCurrencyWhitelistOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeCurrencyWhitelistOwnershipTransferred)
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
		it.Event = new(FeeCurrencyWhitelistOwnershipTransferred)
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
func (it *FeeCurrencyWhitelistOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeCurrencyWhitelistOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeCurrencyWhitelistOwnershipTransferred represents a OwnershipTransferred event raised by the FeeCurrencyWhitelist contract.
type FeeCurrencyWhitelistOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FeeCurrencyWhitelistOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeCurrencyWhitelist.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeeCurrencyWhitelistOwnershipTransferredIterator{contract: _FeeCurrencyWhitelist.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeCurrencyWhitelistOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeCurrencyWhitelist.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeCurrencyWhitelistOwnershipTransferred)
				if err := _FeeCurrencyWhitelist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeCurrencyWhitelist *FeeCurrencyWhitelistFilterer) ParseOwnershipTransferred(log types.Log) (*FeeCurrencyWhitelistOwnershipTransferred, error) {
	event := new(FeeCurrencyWhitelistOwnershipTransferred)
	if err := _FeeCurrencyWhitelist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
