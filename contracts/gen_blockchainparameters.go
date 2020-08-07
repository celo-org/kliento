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

// BlockchainParametersABI is the input ABI used to generate the binding from.
const BlockchainParametersABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blockGasLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"intrinsicGasForAlternativeFeeCurrency\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"major\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"minor\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"patch\",\"type\":\"uint256\"}],\"name\":\"MinimumClientVersionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"gas\",\"type\":\"uint256\"}],\"name\":\"IntrinsicGasForAlternativeFeeCurrencySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"BlockGasLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"major\",\"type\":\"uint256\"},{\"name\":\"minor\",\"type\":\"uint256\"},{\"name\":\"patch\",\"type\":\"uint256\"},{\"name\":\"_gasForNonGoldCurrencies\",\"type\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"major\",\"type\":\"uint256\"},{\"name\":\"minor\",\"type\":\"uint256\"},{\"name\":\"patch\",\"type\":\"uint256\"}],\"name\":\"setMinimumClientVersion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"setBlockGasLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"gas\",\"type\":\"uint256\"}],\"name\":\"setIntrinsicGasForAlternativeFeeCurrency\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMinimumClientVersion\",\"outputs\":[{\"name\":\"major\",\"type\":\"uint256\"},{\"name\":\"minor\",\"type\":\"uint256\"},{\"name\":\"patch\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BlockchainParameters is an auto generated Go binding around an Ethereum contract.
type BlockchainParameters struct {
	BlockchainParametersCaller     // Read-only binding to the contract
	BlockchainParametersTransactor // Write-only binding to the contract
	BlockchainParametersFilterer   // Log filterer for contract events
}

// BlockchainParametersCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockchainParametersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainParametersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockchainParametersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainParametersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockchainParametersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainParametersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockchainParametersSession struct {
	Contract     *BlockchainParameters // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BlockchainParametersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockchainParametersCallerSession struct {
	Contract *BlockchainParametersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// BlockchainParametersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockchainParametersTransactorSession struct {
	Contract     *BlockchainParametersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// BlockchainParametersRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockchainParametersRaw struct {
	Contract *BlockchainParameters // Generic contract binding to access the raw methods on
}

// BlockchainParametersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockchainParametersCallerRaw struct {
	Contract *BlockchainParametersCaller // Generic read-only contract binding to access the raw methods on
}

// BlockchainParametersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockchainParametersTransactorRaw struct {
	Contract *BlockchainParametersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockchainParameters creates a new instance of BlockchainParameters, bound to a specific deployed contract.
func NewBlockchainParameters(address common.Address, backend bind.ContractBackend) (*BlockchainParameters, error) {
	contract, err := bindBlockchainParameters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockchainParameters{BlockchainParametersCaller: BlockchainParametersCaller{contract: contract}, BlockchainParametersTransactor: BlockchainParametersTransactor{contract: contract}, BlockchainParametersFilterer: BlockchainParametersFilterer{contract: contract}}, nil
}

// NewBlockchainParametersCaller creates a new read-only instance of BlockchainParameters, bound to a specific deployed contract.
func NewBlockchainParametersCaller(address common.Address, caller bind.ContractCaller) (*BlockchainParametersCaller, error) {
	contract, err := bindBlockchainParameters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainParametersCaller{contract: contract}, nil
}

// NewBlockchainParametersTransactor creates a new write-only instance of BlockchainParameters, bound to a specific deployed contract.
func NewBlockchainParametersTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockchainParametersTransactor, error) {
	contract, err := bindBlockchainParameters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainParametersTransactor{contract: contract}, nil
}

// NewBlockchainParametersFilterer creates a new log filterer instance of BlockchainParameters, bound to a specific deployed contract.
func NewBlockchainParametersFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockchainParametersFilterer, error) {
	contract, err := bindBlockchainParameters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockchainParametersFilterer{contract: contract}, nil
}

// bindBlockchainParameters binds a generic wrapper to an already deployed contract.
func bindBlockchainParameters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockchainParametersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseBlockchainParametersABI parses the ABI
func ParseBlockchainParametersABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockchainParametersABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockchainParameters *BlockchainParametersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BlockchainParameters.Contract.BlockchainParametersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockchainParameters *BlockchainParametersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.BlockchainParametersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockchainParameters *BlockchainParametersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.BlockchainParametersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockchainParameters *BlockchainParametersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BlockchainParameters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockchainParameters *BlockchainParametersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockchainParameters *BlockchainParametersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.contract.Transact(opts, method, params...)
}

// BlockGasLimit is a free data retrieval call binding the contract method 0x7877a797.
//
// Solidity: function blockGasLimit() constant returns(uint256)
func (_BlockchainParameters *BlockchainParametersCaller) BlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BlockchainParameters.contract.Call(opts, out, "blockGasLimit")
	return *ret0, err
}

// BlockGasLimit is a free data retrieval call binding the contract method 0x7877a797.
//
// Solidity: function blockGasLimit() constant returns(uint256)
func (_BlockchainParameters *BlockchainParametersSession) BlockGasLimit() (*big.Int, error) {
	return _BlockchainParameters.Contract.BlockGasLimit(&_BlockchainParameters.CallOpts)
}

// BlockGasLimit is a free data retrieval call binding the contract method 0x7877a797.
//
// Solidity: function blockGasLimit() constant returns(uint256)
func (_BlockchainParameters *BlockchainParametersCallerSession) BlockGasLimit() (*big.Int, error) {
	return _BlockchainParameters.Contract.BlockGasLimit(&_BlockchainParameters.CallOpts)
}

// GetMinimumClientVersion is a free data retrieval call binding the contract method 0x25eb315d.
//
// Solidity: function getMinimumClientVersion() constant returns(uint256 major, uint256 minor, uint256 patch)
func (_BlockchainParameters *BlockchainParametersCaller) GetMinimumClientVersion(opts *bind.CallOpts) (struct {
	Major *big.Int
	Minor *big.Int
	Patch *big.Int
}, error) {
	ret := new(struct {
		Major *big.Int
		Minor *big.Int
		Patch *big.Int
	})
	out := ret
	err := _BlockchainParameters.contract.Call(opts, out, "getMinimumClientVersion")
	return *ret, err
}

// GetMinimumClientVersion is a free data retrieval call binding the contract method 0x25eb315d.
//
// Solidity: function getMinimumClientVersion() constant returns(uint256 major, uint256 minor, uint256 patch)
func (_BlockchainParameters *BlockchainParametersSession) GetMinimumClientVersion() (struct {
	Major *big.Int
	Minor *big.Int
	Patch *big.Int
}, error) {
	return _BlockchainParameters.Contract.GetMinimumClientVersion(&_BlockchainParameters.CallOpts)
}

// GetMinimumClientVersion is a free data retrieval call binding the contract method 0x25eb315d.
//
// Solidity: function getMinimumClientVersion() constant returns(uint256 major, uint256 minor, uint256 patch)
func (_BlockchainParameters *BlockchainParametersCallerSession) GetMinimumClientVersion() (struct {
	Major *big.Int
	Minor *big.Int
	Patch *big.Int
}, error) {
	return _BlockchainParameters.Contract.GetMinimumClientVersion(&_BlockchainParameters.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_BlockchainParameters *BlockchainParametersCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BlockchainParameters.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_BlockchainParameters *BlockchainParametersSession) Initialized() (bool, error) {
	return _BlockchainParameters.Contract.Initialized(&_BlockchainParameters.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_BlockchainParameters *BlockchainParametersCallerSession) Initialized() (bool, error) {
	return _BlockchainParameters.Contract.Initialized(&_BlockchainParameters.CallOpts)
}

// IntrinsicGasForAlternativeFeeCurrency is a free data retrieval call binding the contract method 0x808474f1.
//
// Solidity: function intrinsicGasForAlternativeFeeCurrency() constant returns(uint256)
func (_BlockchainParameters *BlockchainParametersCaller) IntrinsicGasForAlternativeFeeCurrency(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BlockchainParameters.contract.Call(opts, out, "intrinsicGasForAlternativeFeeCurrency")
	return *ret0, err
}

// IntrinsicGasForAlternativeFeeCurrency is a free data retrieval call binding the contract method 0x808474f1.
//
// Solidity: function intrinsicGasForAlternativeFeeCurrency() constant returns(uint256)
func (_BlockchainParameters *BlockchainParametersSession) IntrinsicGasForAlternativeFeeCurrency() (*big.Int, error) {
	return _BlockchainParameters.Contract.IntrinsicGasForAlternativeFeeCurrency(&_BlockchainParameters.CallOpts)
}

// IntrinsicGasForAlternativeFeeCurrency is a free data retrieval call binding the contract method 0x808474f1.
//
// Solidity: function intrinsicGasForAlternativeFeeCurrency() constant returns(uint256)
func (_BlockchainParameters *BlockchainParametersCallerSession) IntrinsicGasForAlternativeFeeCurrency() (*big.Int, error) {
	return _BlockchainParameters.Contract.IntrinsicGasForAlternativeFeeCurrency(&_BlockchainParameters.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BlockchainParameters *BlockchainParametersCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BlockchainParameters.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BlockchainParameters *BlockchainParametersSession) IsOwner() (bool, error) {
	return _BlockchainParameters.Contract.IsOwner(&_BlockchainParameters.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_BlockchainParameters *BlockchainParametersCallerSession) IsOwner() (bool, error) {
	return _BlockchainParameters.Contract.IsOwner(&_BlockchainParameters.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BlockchainParameters *BlockchainParametersCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BlockchainParameters.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BlockchainParameters *BlockchainParametersSession) Owner() (common.Address, error) {
	return _BlockchainParameters.Contract.Owner(&_BlockchainParameters.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BlockchainParameters *BlockchainParametersCallerSession) Owner() (common.Address, error) {
	return _BlockchainParameters.Contract.Owner(&_BlockchainParameters.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xdcbab608.
//
// Solidity: function initialize(uint256 major, uint256 minor, uint256 patch, uint256 _gasForNonGoldCurrencies, uint256 gasLimit) returns()
func (_BlockchainParameters *BlockchainParametersTransactor) Initialize(opts *bind.TransactOpts, major *big.Int, minor *big.Int, patch *big.Int, _gasForNonGoldCurrencies *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.contract.Transact(opts, "initialize", major, minor, patch, _gasForNonGoldCurrencies, gasLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0xdcbab608.
//
// Solidity: function initialize(uint256 major, uint256 minor, uint256 patch, uint256 _gasForNonGoldCurrencies, uint256 gasLimit) returns()
func (_BlockchainParameters *BlockchainParametersSession) Initialize(major *big.Int, minor *big.Int, patch *big.Int, _gasForNonGoldCurrencies *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.Initialize(&_BlockchainParameters.TransactOpts, major, minor, patch, _gasForNonGoldCurrencies, gasLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0xdcbab608.
//
// Solidity: function initialize(uint256 major, uint256 minor, uint256 patch, uint256 _gasForNonGoldCurrencies, uint256 gasLimit) returns()
func (_BlockchainParameters *BlockchainParametersTransactorSession) Initialize(major *big.Int, minor *big.Int, patch *big.Int, _gasForNonGoldCurrencies *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.Initialize(&_BlockchainParameters.TransactOpts, major, minor, patch, _gasForNonGoldCurrencies, gasLimit)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockchainParameters *BlockchainParametersTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockchainParameters.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockchainParameters *BlockchainParametersSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlockchainParameters.Contract.RenounceOwnership(&_BlockchainParameters.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockchainParameters *BlockchainParametersTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlockchainParameters.Contract.RenounceOwnership(&_BlockchainParameters.TransactOpts)
}

// SetBlockGasLimit is a paid mutator transaction binding the contract method 0xa69257f3.
//
// Solidity: function setBlockGasLimit(uint256 gasLimit) returns()
func (_BlockchainParameters *BlockchainParametersTransactor) SetBlockGasLimit(opts *bind.TransactOpts, gasLimit *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.contract.Transact(opts, "setBlockGasLimit", gasLimit)
}

// SetBlockGasLimit is a paid mutator transaction binding the contract method 0xa69257f3.
//
// Solidity: function setBlockGasLimit(uint256 gasLimit) returns()
func (_BlockchainParameters *BlockchainParametersSession) SetBlockGasLimit(gasLimit *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.SetBlockGasLimit(&_BlockchainParameters.TransactOpts, gasLimit)
}

// SetBlockGasLimit is a paid mutator transaction binding the contract method 0xa69257f3.
//
// Solidity: function setBlockGasLimit(uint256 gasLimit) returns()
func (_BlockchainParameters *BlockchainParametersTransactorSession) SetBlockGasLimit(gasLimit *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.SetBlockGasLimit(&_BlockchainParameters.TransactOpts, gasLimit)
}

// SetIntrinsicGasForAlternativeFeeCurrency is a paid mutator transaction binding the contract method 0xcb0ec628.
//
// Solidity: function setIntrinsicGasForAlternativeFeeCurrency(uint256 gas) returns()
func (_BlockchainParameters *BlockchainParametersTransactor) SetIntrinsicGasForAlternativeFeeCurrency(opts *bind.TransactOpts, gas *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.contract.Transact(opts, "setIntrinsicGasForAlternativeFeeCurrency", gas)
}

// SetIntrinsicGasForAlternativeFeeCurrency is a paid mutator transaction binding the contract method 0xcb0ec628.
//
// Solidity: function setIntrinsicGasForAlternativeFeeCurrency(uint256 gas) returns()
func (_BlockchainParameters *BlockchainParametersSession) SetIntrinsicGasForAlternativeFeeCurrency(gas *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.SetIntrinsicGasForAlternativeFeeCurrency(&_BlockchainParameters.TransactOpts, gas)
}

// SetIntrinsicGasForAlternativeFeeCurrency is a paid mutator transaction binding the contract method 0xcb0ec628.
//
// Solidity: function setIntrinsicGasForAlternativeFeeCurrency(uint256 gas) returns()
func (_BlockchainParameters *BlockchainParametersTransactorSession) SetIntrinsicGasForAlternativeFeeCurrency(gas *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.SetIntrinsicGasForAlternativeFeeCurrency(&_BlockchainParameters.TransactOpts, gas)
}

// SetMinimumClientVersion is a paid mutator transaction binding the contract method 0xbb3ff745.
//
// Solidity: function setMinimumClientVersion(uint256 major, uint256 minor, uint256 patch) returns()
func (_BlockchainParameters *BlockchainParametersTransactor) SetMinimumClientVersion(opts *bind.TransactOpts, major *big.Int, minor *big.Int, patch *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.contract.Transact(opts, "setMinimumClientVersion", major, minor, patch)
}

// SetMinimumClientVersion is a paid mutator transaction binding the contract method 0xbb3ff745.
//
// Solidity: function setMinimumClientVersion(uint256 major, uint256 minor, uint256 patch) returns()
func (_BlockchainParameters *BlockchainParametersSession) SetMinimumClientVersion(major *big.Int, minor *big.Int, patch *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.SetMinimumClientVersion(&_BlockchainParameters.TransactOpts, major, minor, patch)
}

// SetMinimumClientVersion is a paid mutator transaction binding the contract method 0xbb3ff745.
//
// Solidity: function setMinimumClientVersion(uint256 major, uint256 minor, uint256 patch) returns()
func (_BlockchainParameters *BlockchainParametersTransactorSession) SetMinimumClientVersion(major *big.Int, minor *big.Int, patch *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.SetMinimumClientVersion(&_BlockchainParameters.TransactOpts, major, minor, patch)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockchainParameters *BlockchainParametersTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BlockchainParameters.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockchainParameters *BlockchainParametersSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.TransferOwnership(&_BlockchainParameters.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockchainParameters *BlockchainParametersTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.TransferOwnership(&_BlockchainParameters.TransactOpts, newOwner)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_BlockchainParameters *BlockchainParametersFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _BlockchainParameters.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "BlockGasLimitSet":
		event, err = _BlockchainParameters.ParseBlockGasLimitSet(log)
	case "IntrinsicGasForAlternativeFeeCurrencySet":
		event, err = _BlockchainParameters.ParseIntrinsicGasForAlternativeFeeCurrencySet(log)
	case "MinimumClientVersionSet":
		event, err = _BlockchainParameters.ParseMinimumClientVersionSet(log)
	case "OwnershipTransferred":
		event, err = _BlockchainParameters.ParseOwnershipTransferred(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// BlockchainParametersBlockGasLimitSetIterator is returned from FilterBlockGasLimitSet and is used to iterate over the raw logs and unpacked data for BlockGasLimitSet events raised by the BlockchainParameters contract.
type BlockchainParametersBlockGasLimitSetIterator struct {
	Event *BlockchainParametersBlockGasLimitSet // Event containing the contract specifics and raw log

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
func (it *BlockchainParametersBlockGasLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainParametersBlockGasLimitSet)
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
		it.Event = new(BlockchainParametersBlockGasLimitSet)
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
func (it *BlockchainParametersBlockGasLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainParametersBlockGasLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainParametersBlockGasLimitSet represents a BlockGasLimitSet event raised by the BlockchainParameters contract.
type BlockchainParametersBlockGasLimitSet struct {
	Limit *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBlockGasLimitSet is a free log retrieval operation binding the contract event 0x55311ae9c14427b0863f38ed97a2a5944c50d824bbf692836246512e6822c3cf.
//
// Solidity: event BlockGasLimitSet(uint256 limit)
func (_BlockchainParameters *BlockchainParametersFilterer) FilterBlockGasLimitSet(opts *bind.FilterOpts) (*BlockchainParametersBlockGasLimitSetIterator, error) {

	logs, sub, err := _BlockchainParameters.contract.FilterLogs(opts, "BlockGasLimitSet")
	if err != nil {
		return nil, err
	}
	return &BlockchainParametersBlockGasLimitSetIterator{contract: _BlockchainParameters.contract, event: "BlockGasLimitSet", logs: logs, sub: sub}, nil
}

// WatchBlockGasLimitSet is a free log subscription operation binding the contract event 0x55311ae9c14427b0863f38ed97a2a5944c50d824bbf692836246512e6822c3cf.
//
// Solidity: event BlockGasLimitSet(uint256 limit)
func (_BlockchainParameters *BlockchainParametersFilterer) WatchBlockGasLimitSet(opts *bind.WatchOpts, sink chan<- *BlockchainParametersBlockGasLimitSet) (event.Subscription, error) {

	logs, sub, err := _BlockchainParameters.contract.WatchLogs(opts, "BlockGasLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainParametersBlockGasLimitSet)
				if err := _BlockchainParameters.contract.UnpackLog(event, "BlockGasLimitSet", log); err != nil {
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

// ParseBlockGasLimitSet is a log parse operation binding the contract event 0x55311ae9c14427b0863f38ed97a2a5944c50d824bbf692836246512e6822c3cf.
//
// Solidity: event BlockGasLimitSet(uint256 limit)
func (_BlockchainParameters *BlockchainParametersFilterer) ParseBlockGasLimitSet(log types.Log) (*BlockchainParametersBlockGasLimitSet, error) {
	event := new(BlockchainParametersBlockGasLimitSet)
	if err := _BlockchainParameters.contract.UnpackLog(event, "BlockGasLimitSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySetIterator is returned from FilterIntrinsicGasForAlternativeFeeCurrencySet and is used to iterate over the raw logs and unpacked data for IntrinsicGasForAlternativeFeeCurrencySet events raised by the BlockchainParameters contract.
type BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySetIterator struct {
	Event *BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySet // Event containing the contract specifics and raw log

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
func (it *BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySet)
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
		it.Event = new(BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySet)
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
func (it *BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySet represents a IntrinsicGasForAlternativeFeeCurrencySet event raised by the BlockchainParameters contract.
type BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySet struct {
	Gas *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterIntrinsicGasForAlternativeFeeCurrencySet is a free log retrieval operation binding the contract event 0xba9c6f28c7d9990745a5b5282dbee04706c28cae24a44736c3ba99b57c021f3e.
//
// Solidity: event IntrinsicGasForAlternativeFeeCurrencySet(uint256 gas)
func (_BlockchainParameters *BlockchainParametersFilterer) FilterIntrinsicGasForAlternativeFeeCurrencySet(opts *bind.FilterOpts) (*BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySetIterator, error) {

	logs, sub, err := _BlockchainParameters.contract.FilterLogs(opts, "IntrinsicGasForAlternativeFeeCurrencySet")
	if err != nil {
		return nil, err
	}
	return &BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySetIterator{contract: _BlockchainParameters.contract, event: "IntrinsicGasForAlternativeFeeCurrencySet", logs: logs, sub: sub}, nil
}

// WatchIntrinsicGasForAlternativeFeeCurrencySet is a free log subscription operation binding the contract event 0xba9c6f28c7d9990745a5b5282dbee04706c28cae24a44736c3ba99b57c021f3e.
//
// Solidity: event IntrinsicGasForAlternativeFeeCurrencySet(uint256 gas)
func (_BlockchainParameters *BlockchainParametersFilterer) WatchIntrinsicGasForAlternativeFeeCurrencySet(opts *bind.WatchOpts, sink chan<- *BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySet) (event.Subscription, error) {

	logs, sub, err := _BlockchainParameters.contract.WatchLogs(opts, "IntrinsicGasForAlternativeFeeCurrencySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySet)
				if err := _BlockchainParameters.contract.UnpackLog(event, "IntrinsicGasForAlternativeFeeCurrencySet", log); err != nil {
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

// ParseIntrinsicGasForAlternativeFeeCurrencySet is a log parse operation binding the contract event 0xba9c6f28c7d9990745a5b5282dbee04706c28cae24a44736c3ba99b57c021f3e.
//
// Solidity: event IntrinsicGasForAlternativeFeeCurrencySet(uint256 gas)
func (_BlockchainParameters *BlockchainParametersFilterer) ParseIntrinsicGasForAlternativeFeeCurrencySet(log types.Log) (*BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySet, error) {
	event := new(BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySet)
	if err := _BlockchainParameters.contract.UnpackLog(event, "IntrinsicGasForAlternativeFeeCurrencySet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BlockchainParametersMinimumClientVersionSetIterator is returned from FilterMinimumClientVersionSet and is used to iterate over the raw logs and unpacked data for MinimumClientVersionSet events raised by the BlockchainParameters contract.
type BlockchainParametersMinimumClientVersionSetIterator struct {
	Event *BlockchainParametersMinimumClientVersionSet // Event containing the contract specifics and raw log

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
func (it *BlockchainParametersMinimumClientVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainParametersMinimumClientVersionSet)
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
		it.Event = new(BlockchainParametersMinimumClientVersionSet)
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
func (it *BlockchainParametersMinimumClientVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainParametersMinimumClientVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainParametersMinimumClientVersionSet represents a MinimumClientVersionSet event raised by the BlockchainParameters contract.
type BlockchainParametersMinimumClientVersionSet struct {
	Major *big.Int
	Minor *big.Int
	Patch *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinimumClientVersionSet is a free log retrieval operation binding the contract event 0x809db05bd174a70ede53d18fc046c5ceb86ebffbb7746a0c8605772c97ef0d52.
//
// Solidity: event MinimumClientVersionSet(uint256 major, uint256 minor, uint256 patch)
func (_BlockchainParameters *BlockchainParametersFilterer) FilterMinimumClientVersionSet(opts *bind.FilterOpts) (*BlockchainParametersMinimumClientVersionSetIterator, error) {

	logs, sub, err := _BlockchainParameters.contract.FilterLogs(opts, "MinimumClientVersionSet")
	if err != nil {
		return nil, err
	}
	return &BlockchainParametersMinimumClientVersionSetIterator{contract: _BlockchainParameters.contract, event: "MinimumClientVersionSet", logs: logs, sub: sub}, nil
}

// WatchMinimumClientVersionSet is a free log subscription operation binding the contract event 0x809db05bd174a70ede53d18fc046c5ceb86ebffbb7746a0c8605772c97ef0d52.
//
// Solidity: event MinimumClientVersionSet(uint256 major, uint256 minor, uint256 patch)
func (_BlockchainParameters *BlockchainParametersFilterer) WatchMinimumClientVersionSet(opts *bind.WatchOpts, sink chan<- *BlockchainParametersMinimumClientVersionSet) (event.Subscription, error) {

	logs, sub, err := _BlockchainParameters.contract.WatchLogs(opts, "MinimumClientVersionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainParametersMinimumClientVersionSet)
				if err := _BlockchainParameters.contract.UnpackLog(event, "MinimumClientVersionSet", log); err != nil {
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

// ParseMinimumClientVersionSet is a log parse operation binding the contract event 0x809db05bd174a70ede53d18fc046c5ceb86ebffbb7746a0c8605772c97ef0d52.
//
// Solidity: event MinimumClientVersionSet(uint256 major, uint256 minor, uint256 patch)
func (_BlockchainParameters *BlockchainParametersFilterer) ParseMinimumClientVersionSet(log types.Log) (*BlockchainParametersMinimumClientVersionSet, error) {
	event := new(BlockchainParametersMinimumClientVersionSet)
	if err := _BlockchainParameters.contract.UnpackLog(event, "MinimumClientVersionSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BlockchainParametersOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BlockchainParameters contract.
type BlockchainParametersOwnershipTransferredIterator struct {
	Event *BlockchainParametersOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BlockchainParametersOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainParametersOwnershipTransferred)
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
		it.Event = new(BlockchainParametersOwnershipTransferred)
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
func (it *BlockchainParametersOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainParametersOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainParametersOwnershipTransferred represents a OwnershipTransferred event raised by the BlockchainParameters contract.
type BlockchainParametersOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlockchainParameters *BlockchainParametersFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BlockchainParametersOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlockchainParameters.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainParametersOwnershipTransferredIterator{contract: _BlockchainParameters.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlockchainParameters *BlockchainParametersFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlockchainParametersOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlockchainParameters.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainParametersOwnershipTransferred)
				if err := _BlockchainParameters.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BlockchainParameters *BlockchainParametersFilterer) ParseOwnershipTransferred(log types.Log) (*BlockchainParametersOwnershipTransferred, error) {
	event := new(BlockchainParametersOwnershipTransferred)
	if err := _BlockchainParameters.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
