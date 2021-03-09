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

// TransferWhitelistABI is the input ABI used to generate the binding from.
const TransferWhitelistABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"contractIdentifier\",\"type\":\"bytes32\"}],\"name\":\"WhitelistedContractIdentifier\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelistedContractIdentifiers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"whitelistAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"removedAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"contractIdentifier\",\"type\":\"bytes32\"}],\"name\":\"whitelistRegisteredContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumberOfWhitelistedContractIdentifiers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_whitelist\",\"type\":\"address[]\"}],\"name\":\"setDirectlyWhitelistedAddresses\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_registeredContracts\",\"type\":\"bytes32[]\"}],\"name\":\"setWhitelistedContractIdentifiers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWhitelist\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"selfDestruct\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TransferWhitelist is an auto generated Go binding around an Ethereum contract.
type TransferWhitelist struct {
	TransferWhitelistCaller     // Read-only binding to the contract
	TransferWhitelistTransactor // Write-only binding to the contract
	TransferWhitelistFilterer   // Log filterer for contract events
}

// TransferWhitelistCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferWhitelistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferWhitelistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferWhitelistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferWhitelistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferWhitelistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferWhitelistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferWhitelistSession struct {
	Contract     *TransferWhitelist // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TransferWhitelistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferWhitelistCallerSession struct {
	Contract *TransferWhitelistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TransferWhitelistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferWhitelistTransactorSession struct {
	Contract     *TransferWhitelistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TransferWhitelistRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferWhitelistRaw struct {
	Contract *TransferWhitelist // Generic contract binding to access the raw methods on
}

// TransferWhitelistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferWhitelistCallerRaw struct {
	Contract *TransferWhitelistCaller // Generic read-only contract binding to access the raw methods on
}

// TransferWhitelistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferWhitelistTransactorRaw struct {
	Contract *TransferWhitelistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferWhitelist creates a new instance of TransferWhitelist, bound to a specific deployed contract.
func NewTransferWhitelist(address common.Address, backend bind.ContractBackend) (*TransferWhitelist, error) {
	contract, err := bindTransferWhitelist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransferWhitelist{TransferWhitelistCaller: TransferWhitelistCaller{contract: contract}, TransferWhitelistTransactor: TransferWhitelistTransactor{contract: contract}, TransferWhitelistFilterer: TransferWhitelistFilterer{contract: contract}}, nil
}

// NewTransferWhitelistCaller creates a new read-only instance of TransferWhitelist, bound to a specific deployed contract.
func NewTransferWhitelistCaller(address common.Address, caller bind.ContractCaller) (*TransferWhitelistCaller, error) {
	contract, err := bindTransferWhitelist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferWhitelistCaller{contract: contract}, nil
}

// NewTransferWhitelistTransactor creates a new write-only instance of TransferWhitelist, bound to a specific deployed contract.
func NewTransferWhitelistTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferWhitelistTransactor, error) {
	contract, err := bindTransferWhitelist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferWhitelistTransactor{contract: contract}, nil
}

// NewTransferWhitelistFilterer creates a new log filterer instance of TransferWhitelist, bound to a specific deployed contract.
func NewTransferWhitelistFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferWhitelistFilterer, error) {
	contract, err := bindTransferWhitelist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferWhitelistFilterer{contract: contract}, nil
}

// bindTransferWhitelist binds a generic wrapper to an already deployed contract.
func bindTransferWhitelist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferWhitelistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseTransferWhitelistABI parses the ABI
func ParseTransferWhitelistABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(TransferWhitelistABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferWhitelist *TransferWhitelistRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TransferWhitelist.Contract.TransferWhitelistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferWhitelist *TransferWhitelistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferWhitelist.Contract.TransferWhitelistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferWhitelist *TransferWhitelistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferWhitelist.Contract.TransferWhitelistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferWhitelist *TransferWhitelistCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TransferWhitelist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferWhitelist *TransferWhitelistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferWhitelist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferWhitelist *TransferWhitelistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferWhitelist.Contract.contract.Transact(opts, method, params...)
}

// GetNumberOfWhitelistedContractIdentifiers is a free data retrieval call binding the contract method 0x1dc34b89.
//
// Solidity: function getNumberOfWhitelistedContractIdentifiers() view returns(uint256 length)
func (_TransferWhitelist *TransferWhitelistCaller) GetNumberOfWhitelistedContractIdentifiers(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TransferWhitelist.contract.Call(opts, out, "getNumberOfWhitelistedContractIdentifiers")
	return *ret0, err
}

// GetNumberOfWhitelistedContractIdentifiers is a free data retrieval call binding the contract method 0x1dc34b89.
//
// Solidity: function getNumberOfWhitelistedContractIdentifiers() view returns(uint256 length)
func (_TransferWhitelist *TransferWhitelistSession) GetNumberOfWhitelistedContractIdentifiers() (*big.Int, error) {
	return _TransferWhitelist.Contract.GetNumberOfWhitelistedContractIdentifiers(&_TransferWhitelist.CallOpts)
}

// GetNumberOfWhitelistedContractIdentifiers is a free data retrieval call binding the contract method 0x1dc34b89.
//
// Solidity: function getNumberOfWhitelistedContractIdentifiers() view returns(uint256 length)
func (_TransferWhitelist *TransferWhitelistCallerSession) GetNumberOfWhitelistedContractIdentifiers() (*big.Int, error) {
	return _TransferWhitelist.Contract.GetNumberOfWhitelistedContractIdentifiers(&_TransferWhitelist.CallOpts)
}

// GetWhitelist is a free data retrieval call binding the contract method 0xd01f63f5.
//
// Solidity: function getWhitelist() view returns(address[])
func (_TransferWhitelist *TransferWhitelistCaller) GetWhitelist(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TransferWhitelist.contract.Call(opts, out, "getWhitelist")
	return *ret0, err
}

// GetWhitelist is a free data retrieval call binding the contract method 0xd01f63f5.
//
// Solidity: function getWhitelist() view returns(address[])
func (_TransferWhitelist *TransferWhitelistSession) GetWhitelist() ([]common.Address, error) {
	return _TransferWhitelist.Contract.GetWhitelist(&_TransferWhitelist.CallOpts)
}

// GetWhitelist is a free data retrieval call binding the contract method 0xd01f63f5.
//
// Solidity: function getWhitelist() view returns(address[])
func (_TransferWhitelist *TransferWhitelistCallerSession) GetWhitelist() ([]common.Address, error) {
	return _TransferWhitelist.Contract.GetWhitelist(&_TransferWhitelist.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_TransferWhitelist *TransferWhitelistCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TransferWhitelist.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_TransferWhitelist *TransferWhitelistSession) IsOwner() (bool, error) {
	return _TransferWhitelist.Contract.IsOwner(&_TransferWhitelist.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_TransferWhitelist *TransferWhitelistCallerSession) IsOwner() (bool, error) {
	return _TransferWhitelist.Contract.IsOwner(&_TransferWhitelist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TransferWhitelist *TransferWhitelistCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TransferWhitelist.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TransferWhitelist *TransferWhitelistSession) Owner() (common.Address, error) {
	return _TransferWhitelist.Contract.Owner(&_TransferWhitelist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TransferWhitelist *TransferWhitelistCallerSession) Owner() (common.Address, error) {
	return _TransferWhitelist.Contract.Owner(&_TransferWhitelist.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_TransferWhitelist *TransferWhitelistCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TransferWhitelist.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_TransferWhitelist *TransferWhitelistSession) Registry() (common.Address, error) {
	return _TransferWhitelist.Contract.Registry(&_TransferWhitelist.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_TransferWhitelist *TransferWhitelistCallerSession) Registry() (common.Address, error) {
	return _TransferWhitelist.Contract.Registry(&_TransferWhitelist.CallOpts)
}

// WhitelistedContractIdentifiers is a free data retrieval call binding the contract method 0x2ed1ce72.
//
// Solidity: function whitelistedContractIdentifiers(uint256 ) view returns(bytes32)
func (_TransferWhitelist *TransferWhitelistCaller) WhitelistedContractIdentifiers(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TransferWhitelist.contract.Call(opts, out, "whitelistedContractIdentifiers", arg0)
	return *ret0, err
}

// WhitelistedContractIdentifiers is a free data retrieval call binding the contract method 0x2ed1ce72.
//
// Solidity: function whitelistedContractIdentifiers(uint256 ) view returns(bytes32)
func (_TransferWhitelist *TransferWhitelistSession) WhitelistedContractIdentifiers(arg0 *big.Int) ([32]byte, error) {
	return _TransferWhitelist.Contract.WhitelistedContractIdentifiers(&_TransferWhitelist.CallOpts, arg0)
}

// WhitelistedContractIdentifiers is a free data retrieval call binding the contract method 0x2ed1ce72.
//
// Solidity: function whitelistedContractIdentifiers(uint256 ) view returns(bytes32)
func (_TransferWhitelist *TransferWhitelistCallerSession) WhitelistedContractIdentifiers(arg0 *big.Int) ([32]byte, error) {
	return _TransferWhitelist.Contract.WhitelistedContractIdentifiers(&_TransferWhitelist.CallOpts, arg0)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0x0662f008.
//
// Solidity: function removeAddress(address removedAddress, uint256 index) returns()
func (_TransferWhitelist *TransferWhitelistTransactor) RemoveAddress(opts *bind.TransactOpts, removedAddress common.Address, index *big.Int) (*types.Transaction, error) {
	return _TransferWhitelist.contract.Transact(opts, "removeAddress", removedAddress, index)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0x0662f008.
//
// Solidity: function removeAddress(address removedAddress, uint256 index) returns()
func (_TransferWhitelist *TransferWhitelistSession) RemoveAddress(removedAddress common.Address, index *big.Int) (*types.Transaction, error) {
	return _TransferWhitelist.Contract.RemoveAddress(&_TransferWhitelist.TransactOpts, removedAddress, index)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0x0662f008.
//
// Solidity: function removeAddress(address removedAddress, uint256 index) returns()
func (_TransferWhitelist *TransferWhitelistTransactorSession) RemoveAddress(removedAddress common.Address, index *big.Int) (*types.Transaction, error) {
	return _TransferWhitelist.Contract.RemoveAddress(&_TransferWhitelist.TransactOpts, removedAddress, index)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TransferWhitelist *TransferWhitelistTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferWhitelist.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TransferWhitelist *TransferWhitelistSession) RenounceOwnership() (*types.Transaction, error) {
	return _TransferWhitelist.Contract.RenounceOwnership(&_TransferWhitelist.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TransferWhitelist *TransferWhitelistTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TransferWhitelist.Contract.RenounceOwnership(&_TransferWhitelist.TransactOpts)
}

// SelfDestruct is a paid mutator transaction binding the contract method 0x9cb8a26a.
//
// Solidity: function selfDestruct() returns()
func (_TransferWhitelist *TransferWhitelistTransactor) SelfDestruct(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferWhitelist.contract.Transact(opts, "selfDestruct")
}

// SelfDestruct is a paid mutator transaction binding the contract method 0x9cb8a26a.
//
// Solidity: function selfDestruct() returns()
func (_TransferWhitelist *TransferWhitelistSession) SelfDestruct() (*types.Transaction, error) {
	return _TransferWhitelist.Contract.SelfDestruct(&_TransferWhitelist.TransactOpts)
}

// SelfDestruct is a paid mutator transaction binding the contract method 0x9cb8a26a.
//
// Solidity: function selfDestruct() returns()
func (_TransferWhitelist *TransferWhitelistTransactorSession) SelfDestruct() (*types.Transaction, error) {
	return _TransferWhitelist.Contract.SelfDestruct(&_TransferWhitelist.TransactOpts)
}

// SetDirectlyWhitelistedAddresses is a paid mutator transaction binding the contract method 0x879337f1.
//
// Solidity: function setDirectlyWhitelistedAddresses(address[] _whitelist) returns()
func (_TransferWhitelist *TransferWhitelistTransactor) SetDirectlyWhitelistedAddresses(opts *bind.TransactOpts, _whitelist []common.Address) (*types.Transaction, error) {
	return _TransferWhitelist.contract.Transact(opts, "setDirectlyWhitelistedAddresses", _whitelist)
}

// SetDirectlyWhitelistedAddresses is a paid mutator transaction binding the contract method 0x879337f1.
//
// Solidity: function setDirectlyWhitelistedAddresses(address[] _whitelist) returns()
func (_TransferWhitelist *TransferWhitelistSession) SetDirectlyWhitelistedAddresses(_whitelist []common.Address) (*types.Transaction, error) {
	return _TransferWhitelist.Contract.SetDirectlyWhitelistedAddresses(&_TransferWhitelist.TransactOpts, _whitelist)
}

// SetDirectlyWhitelistedAddresses is a paid mutator transaction binding the contract method 0x879337f1.
//
// Solidity: function setDirectlyWhitelistedAddresses(address[] _whitelist) returns()
func (_TransferWhitelist *TransferWhitelistTransactorSession) SetDirectlyWhitelistedAddresses(_whitelist []common.Address) (*types.Transaction, error) {
	return _TransferWhitelist.Contract.SetDirectlyWhitelistedAddresses(&_TransferWhitelist.TransactOpts, _whitelist)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_TransferWhitelist *TransferWhitelistTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _TransferWhitelist.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_TransferWhitelist *TransferWhitelistSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _TransferWhitelist.Contract.SetRegistry(&_TransferWhitelist.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_TransferWhitelist *TransferWhitelistTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _TransferWhitelist.Contract.SetRegistry(&_TransferWhitelist.TransactOpts, registryAddress)
}

// SetWhitelistedContractIdentifiers is a paid mutator transaction binding the contract method 0xc0d72486.
//
// Solidity: function setWhitelistedContractIdentifiers(bytes32[] _registeredContracts) returns()
func (_TransferWhitelist *TransferWhitelistTransactor) SetWhitelistedContractIdentifiers(opts *bind.TransactOpts, _registeredContracts [][32]byte) (*types.Transaction, error) {
	return _TransferWhitelist.contract.Transact(opts, "setWhitelistedContractIdentifiers", _registeredContracts)
}

// SetWhitelistedContractIdentifiers is a paid mutator transaction binding the contract method 0xc0d72486.
//
// Solidity: function setWhitelistedContractIdentifiers(bytes32[] _registeredContracts) returns()
func (_TransferWhitelist *TransferWhitelistSession) SetWhitelistedContractIdentifiers(_registeredContracts [][32]byte) (*types.Transaction, error) {
	return _TransferWhitelist.Contract.SetWhitelistedContractIdentifiers(&_TransferWhitelist.TransactOpts, _registeredContracts)
}

// SetWhitelistedContractIdentifiers is a paid mutator transaction binding the contract method 0xc0d72486.
//
// Solidity: function setWhitelistedContractIdentifiers(bytes32[] _registeredContracts) returns()
func (_TransferWhitelist *TransferWhitelistTransactorSession) SetWhitelistedContractIdentifiers(_registeredContracts [][32]byte) (*types.Transaction, error) {
	return _TransferWhitelist.Contract.SetWhitelistedContractIdentifiers(&_TransferWhitelist.TransactOpts, _registeredContracts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TransferWhitelist *TransferWhitelistTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TransferWhitelist.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TransferWhitelist *TransferWhitelistSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TransferWhitelist.Contract.TransferOwnership(&_TransferWhitelist.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TransferWhitelist *TransferWhitelistTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TransferWhitelist.Contract.TransferOwnership(&_TransferWhitelist.TransactOpts, newOwner)
}

// WhitelistAddress is a paid mutator transaction binding the contract method 0x41566585.
//
// Solidity: function whitelistAddress(address newAddress) returns()
func (_TransferWhitelist *TransferWhitelistTransactor) WhitelistAddress(opts *bind.TransactOpts, newAddress common.Address) (*types.Transaction, error) {
	return _TransferWhitelist.contract.Transact(opts, "whitelistAddress", newAddress)
}

// WhitelistAddress is a paid mutator transaction binding the contract method 0x41566585.
//
// Solidity: function whitelistAddress(address newAddress) returns()
func (_TransferWhitelist *TransferWhitelistSession) WhitelistAddress(newAddress common.Address) (*types.Transaction, error) {
	return _TransferWhitelist.Contract.WhitelistAddress(&_TransferWhitelist.TransactOpts, newAddress)
}

// WhitelistAddress is a paid mutator transaction binding the contract method 0x41566585.
//
// Solidity: function whitelistAddress(address newAddress) returns()
func (_TransferWhitelist *TransferWhitelistTransactorSession) WhitelistAddress(newAddress common.Address) (*types.Transaction, error) {
	return _TransferWhitelist.Contract.WhitelistAddress(&_TransferWhitelist.TransactOpts, newAddress)
}

// WhitelistRegisteredContract is a paid mutator transaction binding the contract method 0x23a39b54.
//
// Solidity: function whitelistRegisteredContract(bytes32 contractIdentifier) returns()
func (_TransferWhitelist *TransferWhitelistTransactor) WhitelistRegisteredContract(opts *bind.TransactOpts, contractIdentifier [32]byte) (*types.Transaction, error) {
	return _TransferWhitelist.contract.Transact(opts, "whitelistRegisteredContract", contractIdentifier)
}

// WhitelistRegisteredContract is a paid mutator transaction binding the contract method 0x23a39b54.
//
// Solidity: function whitelistRegisteredContract(bytes32 contractIdentifier) returns()
func (_TransferWhitelist *TransferWhitelistSession) WhitelistRegisteredContract(contractIdentifier [32]byte) (*types.Transaction, error) {
	return _TransferWhitelist.Contract.WhitelistRegisteredContract(&_TransferWhitelist.TransactOpts, contractIdentifier)
}

// WhitelistRegisteredContract is a paid mutator transaction binding the contract method 0x23a39b54.
//
// Solidity: function whitelistRegisteredContract(bytes32 contractIdentifier) returns()
func (_TransferWhitelist *TransferWhitelistTransactorSession) WhitelistRegisteredContract(contractIdentifier [32]byte) (*types.Transaction, error) {
	return _TransferWhitelist.Contract.WhitelistRegisteredContract(&_TransferWhitelist.TransactOpts, contractIdentifier)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_TransferWhitelist *TransferWhitelistFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _TransferWhitelist.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "OwnershipTransferred":
		event, err = _TransferWhitelist.ParseOwnershipTransferred(log)
	case "RegistrySet":
		event, err = _TransferWhitelist.ParseRegistrySet(log)
	case "WhitelistedAddress":
		event, err = _TransferWhitelist.ParseWhitelistedAddress(log)
	case "WhitelistedAddressRemoved":
		event, err = _TransferWhitelist.ParseWhitelistedAddressRemoved(log)
	case "WhitelistedContractIdentifier":
		event, err = _TransferWhitelist.ParseWhitelistedContractIdentifier(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// TransferWhitelistOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TransferWhitelist contract.
type TransferWhitelistOwnershipTransferredIterator struct {
	Event *TransferWhitelistOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TransferWhitelistOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferWhitelistOwnershipTransferred)
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
		it.Event = new(TransferWhitelistOwnershipTransferred)
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
func (it *TransferWhitelistOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferWhitelistOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferWhitelistOwnershipTransferred represents a OwnershipTransferred event raised by the TransferWhitelist contract.
type TransferWhitelistOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TransferWhitelist *TransferWhitelistFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TransferWhitelistOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TransferWhitelist.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TransferWhitelistOwnershipTransferredIterator{contract: _TransferWhitelist.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TransferWhitelist *TransferWhitelistFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TransferWhitelistOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TransferWhitelist.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferWhitelistOwnershipTransferred)
				if err := _TransferWhitelist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TransferWhitelist *TransferWhitelistFilterer) ParseOwnershipTransferred(log types.Log) (*TransferWhitelistOwnershipTransferred, error) {
	event := new(TransferWhitelistOwnershipTransferred)
	if err := _TransferWhitelist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransferWhitelistRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the TransferWhitelist contract.
type TransferWhitelistRegistrySetIterator struct {
	Event *TransferWhitelistRegistrySet // Event containing the contract specifics and raw log

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
func (it *TransferWhitelistRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferWhitelistRegistrySet)
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
		it.Event = new(TransferWhitelistRegistrySet)
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
func (it *TransferWhitelistRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferWhitelistRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferWhitelistRegistrySet represents a RegistrySet event raised by the TransferWhitelist contract.
type TransferWhitelistRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_TransferWhitelist *TransferWhitelistFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*TransferWhitelistRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _TransferWhitelist.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &TransferWhitelistRegistrySetIterator{contract: _TransferWhitelist.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_TransferWhitelist *TransferWhitelistFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *TransferWhitelistRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _TransferWhitelist.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferWhitelistRegistrySet)
				if err := _TransferWhitelist.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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

// ParseRegistrySet is a log parse operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_TransferWhitelist *TransferWhitelistFilterer) ParseRegistrySet(log types.Log) (*TransferWhitelistRegistrySet, error) {
	event := new(TransferWhitelistRegistrySet)
	if err := _TransferWhitelist.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransferWhitelistWhitelistedAddressIterator is returned from FilterWhitelistedAddress and is used to iterate over the raw logs and unpacked data for WhitelistedAddress events raised by the TransferWhitelist contract.
type TransferWhitelistWhitelistedAddressIterator struct {
	Event *TransferWhitelistWhitelistedAddress // Event containing the contract specifics and raw log

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
func (it *TransferWhitelistWhitelistedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferWhitelistWhitelistedAddress)
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
		it.Event = new(TransferWhitelistWhitelistedAddress)
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
func (it *TransferWhitelistWhitelistedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferWhitelistWhitelistedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferWhitelistWhitelistedAddress represents a WhitelistedAddress event raised by the TransferWhitelist contract.
type TransferWhitelistWhitelistedAddress struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAddress is a free log retrieval operation binding the contract event 0x9659de77bedae44137acf94b0b2a83c1a7e51fec986c5d6098bbde7e4f0d02e8.
//
// Solidity: event WhitelistedAddress(address indexed addr)
func (_TransferWhitelist *TransferWhitelistFilterer) FilterWhitelistedAddress(opts *bind.FilterOpts, addr []common.Address) (*TransferWhitelistWhitelistedAddressIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _TransferWhitelist.contract.FilterLogs(opts, "WhitelistedAddress", addrRule)
	if err != nil {
		return nil, err
	}
	return &TransferWhitelistWhitelistedAddressIterator{contract: _TransferWhitelist.contract, event: "WhitelistedAddress", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAddress is a free log subscription operation binding the contract event 0x9659de77bedae44137acf94b0b2a83c1a7e51fec986c5d6098bbde7e4f0d02e8.
//
// Solidity: event WhitelistedAddress(address indexed addr)
func (_TransferWhitelist *TransferWhitelistFilterer) WatchWhitelistedAddress(opts *bind.WatchOpts, sink chan<- *TransferWhitelistWhitelistedAddress, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _TransferWhitelist.contract.WatchLogs(opts, "WhitelistedAddress", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferWhitelistWhitelistedAddress)
				if err := _TransferWhitelist.contract.UnpackLog(event, "WhitelistedAddress", log); err != nil {
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

// ParseWhitelistedAddress is a log parse operation binding the contract event 0x9659de77bedae44137acf94b0b2a83c1a7e51fec986c5d6098bbde7e4f0d02e8.
//
// Solidity: event WhitelistedAddress(address indexed addr)
func (_TransferWhitelist *TransferWhitelistFilterer) ParseWhitelistedAddress(log types.Log) (*TransferWhitelistWhitelistedAddress, error) {
	event := new(TransferWhitelistWhitelistedAddress)
	if err := _TransferWhitelist.contract.UnpackLog(event, "WhitelistedAddress", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransferWhitelistWhitelistedAddressRemovedIterator is returned from FilterWhitelistedAddressRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedAddressRemoved events raised by the TransferWhitelist contract.
type TransferWhitelistWhitelistedAddressRemovedIterator struct {
	Event *TransferWhitelistWhitelistedAddressRemoved // Event containing the contract specifics and raw log

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
func (it *TransferWhitelistWhitelistedAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferWhitelistWhitelistedAddressRemoved)
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
		it.Event = new(TransferWhitelistWhitelistedAddressRemoved)
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
func (it *TransferWhitelistWhitelistedAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferWhitelistWhitelistedAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferWhitelistWhitelistedAddressRemoved represents a WhitelistedAddressRemoved event raised by the TransferWhitelist contract.
type TransferWhitelistWhitelistedAddressRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAddressRemoved is a free log retrieval operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: event WhitelistedAddressRemoved(address indexed addr)
func (_TransferWhitelist *TransferWhitelistFilterer) FilterWhitelistedAddressRemoved(opts *bind.FilterOpts, addr []common.Address) (*TransferWhitelistWhitelistedAddressRemovedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _TransferWhitelist.contract.FilterLogs(opts, "WhitelistedAddressRemoved", addrRule)
	if err != nil {
		return nil, err
	}
	return &TransferWhitelistWhitelistedAddressRemovedIterator{contract: _TransferWhitelist.contract, event: "WhitelistedAddressRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAddressRemoved is a free log subscription operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: event WhitelistedAddressRemoved(address indexed addr)
func (_TransferWhitelist *TransferWhitelistFilterer) WatchWhitelistedAddressRemoved(opts *bind.WatchOpts, sink chan<- *TransferWhitelistWhitelistedAddressRemoved, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _TransferWhitelist.contract.WatchLogs(opts, "WhitelistedAddressRemoved", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferWhitelistWhitelistedAddressRemoved)
				if err := _TransferWhitelist.contract.UnpackLog(event, "WhitelistedAddressRemoved", log); err != nil {
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

// ParseWhitelistedAddressRemoved is a log parse operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: event WhitelistedAddressRemoved(address indexed addr)
func (_TransferWhitelist *TransferWhitelistFilterer) ParseWhitelistedAddressRemoved(log types.Log) (*TransferWhitelistWhitelistedAddressRemoved, error) {
	event := new(TransferWhitelistWhitelistedAddressRemoved)
	if err := _TransferWhitelist.contract.UnpackLog(event, "WhitelistedAddressRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransferWhitelistWhitelistedContractIdentifierIterator is returned from FilterWhitelistedContractIdentifier and is used to iterate over the raw logs and unpacked data for WhitelistedContractIdentifier events raised by the TransferWhitelist contract.
type TransferWhitelistWhitelistedContractIdentifierIterator struct {
	Event *TransferWhitelistWhitelistedContractIdentifier // Event containing the contract specifics and raw log

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
func (it *TransferWhitelistWhitelistedContractIdentifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransferWhitelistWhitelistedContractIdentifier)
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
		it.Event = new(TransferWhitelistWhitelistedContractIdentifier)
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
func (it *TransferWhitelistWhitelistedContractIdentifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransferWhitelistWhitelistedContractIdentifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransferWhitelistWhitelistedContractIdentifier represents a WhitelistedContractIdentifier event raised by the TransferWhitelist contract.
type TransferWhitelistWhitelistedContractIdentifier struct {
	ContractIdentifier [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedContractIdentifier is a free log retrieval operation binding the contract event 0x4084a390e8b98bd7a32fac9847f0147f7f6773834e5738af4fe38598f80bd619.
//
// Solidity: event WhitelistedContractIdentifier(bytes32 indexed contractIdentifier)
func (_TransferWhitelist *TransferWhitelistFilterer) FilterWhitelistedContractIdentifier(opts *bind.FilterOpts, contractIdentifier [][32]byte) (*TransferWhitelistWhitelistedContractIdentifierIterator, error) {

	var contractIdentifierRule []interface{}
	for _, contractIdentifierItem := range contractIdentifier {
		contractIdentifierRule = append(contractIdentifierRule, contractIdentifierItem)
	}

	logs, sub, err := _TransferWhitelist.contract.FilterLogs(opts, "WhitelistedContractIdentifier", contractIdentifierRule)
	if err != nil {
		return nil, err
	}
	return &TransferWhitelistWhitelistedContractIdentifierIterator{contract: _TransferWhitelist.contract, event: "WhitelistedContractIdentifier", logs: logs, sub: sub}, nil
}

// WatchWhitelistedContractIdentifier is a free log subscription operation binding the contract event 0x4084a390e8b98bd7a32fac9847f0147f7f6773834e5738af4fe38598f80bd619.
//
// Solidity: event WhitelistedContractIdentifier(bytes32 indexed contractIdentifier)
func (_TransferWhitelist *TransferWhitelistFilterer) WatchWhitelistedContractIdentifier(opts *bind.WatchOpts, sink chan<- *TransferWhitelistWhitelistedContractIdentifier, contractIdentifier [][32]byte) (event.Subscription, error) {

	var contractIdentifierRule []interface{}
	for _, contractIdentifierItem := range contractIdentifier {
		contractIdentifierRule = append(contractIdentifierRule, contractIdentifierItem)
	}

	logs, sub, err := _TransferWhitelist.contract.WatchLogs(opts, "WhitelistedContractIdentifier", contractIdentifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransferWhitelistWhitelistedContractIdentifier)
				if err := _TransferWhitelist.contract.UnpackLog(event, "WhitelistedContractIdentifier", log); err != nil {
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

// ParseWhitelistedContractIdentifier is a log parse operation binding the contract event 0x4084a390e8b98bd7a32fac9847f0147f7f6773834e5738af4fe38598f80bd619.
//
// Solidity: event WhitelistedContractIdentifier(bytes32 indexed contractIdentifier)
func (_TransferWhitelist *TransferWhitelistFilterer) ParseWhitelistedContractIdentifier(log types.Log) (*TransferWhitelistWhitelistedContractIdentifier, error) {
	event := new(TransferWhitelistWhitelistedContractIdentifier)
	if err := _TransferWhitelist.contract.UnpackLog(event, "WhitelistedContractIdentifier", log); err != nil {
		return nil, err
	}
	return event, nil
}
