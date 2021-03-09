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

// EscrowABI is the input ABI used to generate the binding from.
const EscrowABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentId\",\"type\":\"address\"}],\"name\":\"Revocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minAttestations\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentId\",\"type\":\"address\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"escrowedPayments\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"recipientIdentifier\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sentIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAttestations\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"receivedPaymentIds\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sentPaymentIds\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVersionNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirySeconds\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAttestations\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"paymentId\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"paymentId\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"}],\"name\":\"getReceivedPaymentIds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getSentPaymentIds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Escrow is an auto generated Go binding around an Ethereum contract.
type Escrow struct {
	EscrowCaller     // Read-only binding to the contract
	EscrowTransactor // Write-only binding to the contract
	EscrowFilterer   // Log filterer for contract events
}

// EscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type EscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EscrowSession struct {
	Contract     *Escrow           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EscrowCallerSession struct {
	Contract *EscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EscrowTransactorSession struct {
	Contract     *EscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type EscrowRaw struct {
	Contract *Escrow // Generic contract binding to access the raw methods on
}

// EscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EscrowCallerRaw struct {
	Contract *EscrowCaller // Generic read-only contract binding to access the raw methods on
}

// EscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EscrowTransactorRaw struct {
	Contract *EscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEscrow creates a new instance of Escrow, bound to a specific deployed contract.
func NewEscrow(address common.Address, backend bind.ContractBackend) (*Escrow, error) {
	contract, err := bindEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Escrow{EscrowCaller: EscrowCaller{contract: contract}, EscrowTransactor: EscrowTransactor{contract: contract}, EscrowFilterer: EscrowFilterer{contract: contract}}, nil
}

// NewEscrowCaller creates a new read-only instance of Escrow, bound to a specific deployed contract.
func NewEscrowCaller(address common.Address, caller bind.ContractCaller) (*EscrowCaller, error) {
	contract, err := bindEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowCaller{contract: contract}, nil
}

// NewEscrowTransactor creates a new write-only instance of Escrow, bound to a specific deployed contract.
func NewEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*EscrowTransactor, error) {
	contract, err := bindEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowTransactor{contract: contract}, nil
}

// NewEscrowFilterer creates a new log filterer instance of Escrow, bound to a specific deployed contract.
func NewEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*EscrowFilterer, error) {
	contract, err := bindEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EscrowFilterer{contract: contract}, nil
}

// bindEscrow binds a generic wrapper to an already deployed contract.
func bindEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EscrowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseEscrowABI parses the ABI
func ParseEscrowABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(EscrowABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrow *EscrowRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Escrow.Contract.EscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrow *EscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.Contract.EscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrow *EscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrow.Contract.EscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrow *EscrowCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Escrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrow *EscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrow *EscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrow.Contract.contract.Transact(opts, method, params...)
}

// EscrowedPayments is a free data retrieval call binding the contract method 0x680d782c.
//
// Solidity: function escrowedPayments(address ) view returns(bytes32 recipientIdentifier, address sender, address token, uint256 value, uint256 sentIndex, uint256 receivedIndex, uint256 timestamp, uint256 expirySeconds, uint256 minAttestations)
func (_Escrow *EscrowCaller) EscrowedPayments(opts *bind.CallOpts, arg0 common.Address) (struct {
	RecipientIdentifier [32]byte
	Sender              common.Address
	Token               common.Address
	Value               *big.Int
	SentIndex           *big.Int
	ReceivedIndex       *big.Int
	Timestamp           *big.Int
	ExpirySeconds       *big.Int
	MinAttestations     *big.Int
}, error) {
	ret := new(struct {
		RecipientIdentifier [32]byte
		Sender              common.Address
		Token               common.Address
		Value               *big.Int
		SentIndex           *big.Int
		ReceivedIndex       *big.Int
		Timestamp           *big.Int
		ExpirySeconds       *big.Int
		MinAttestations     *big.Int
	})
	out := ret
	err := _Escrow.contract.Call(opts, out, "escrowedPayments", arg0)
	return *ret, err
}

// EscrowedPayments is a free data retrieval call binding the contract method 0x680d782c.
//
// Solidity: function escrowedPayments(address ) view returns(bytes32 recipientIdentifier, address sender, address token, uint256 value, uint256 sentIndex, uint256 receivedIndex, uint256 timestamp, uint256 expirySeconds, uint256 minAttestations)
func (_Escrow *EscrowSession) EscrowedPayments(arg0 common.Address) (struct {
	RecipientIdentifier [32]byte
	Sender              common.Address
	Token               common.Address
	Value               *big.Int
	SentIndex           *big.Int
	ReceivedIndex       *big.Int
	Timestamp           *big.Int
	ExpirySeconds       *big.Int
	MinAttestations     *big.Int
}, error) {
	return _Escrow.Contract.EscrowedPayments(&_Escrow.CallOpts, arg0)
}

// EscrowedPayments is a free data retrieval call binding the contract method 0x680d782c.
//
// Solidity: function escrowedPayments(address ) view returns(bytes32 recipientIdentifier, address sender, address token, uint256 value, uint256 sentIndex, uint256 receivedIndex, uint256 timestamp, uint256 expirySeconds, uint256 minAttestations)
func (_Escrow *EscrowCallerSession) EscrowedPayments(arg0 common.Address) (struct {
	RecipientIdentifier [32]byte
	Sender              common.Address
	Token               common.Address
	Value               *big.Int
	SentIndex           *big.Int
	ReceivedIndex       *big.Int
	Timestamp           *big.Int
	ExpirySeconds       *big.Int
	MinAttestations     *big.Int
}, error) {
	return _Escrow.Contract.EscrowedPayments(&_Escrow.CallOpts, arg0)
}

// GetReceivedPaymentIds is a free data retrieval call binding the contract method 0x5b57b65b.
//
// Solidity: function getReceivedPaymentIds(bytes32 identifier) view returns(address[])
func (_Escrow *EscrowCaller) GetReceivedPaymentIds(opts *bind.CallOpts, identifier [32]byte) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Escrow.contract.Call(opts, out, "getReceivedPaymentIds", identifier)
	return *ret0, err
}

// GetReceivedPaymentIds is a free data retrieval call binding the contract method 0x5b57b65b.
//
// Solidity: function getReceivedPaymentIds(bytes32 identifier) view returns(address[])
func (_Escrow *EscrowSession) GetReceivedPaymentIds(identifier [32]byte) ([]common.Address, error) {
	return _Escrow.Contract.GetReceivedPaymentIds(&_Escrow.CallOpts, identifier)
}

// GetReceivedPaymentIds is a free data retrieval call binding the contract method 0x5b57b65b.
//
// Solidity: function getReceivedPaymentIds(bytes32 identifier) view returns(address[])
func (_Escrow *EscrowCallerSession) GetReceivedPaymentIds(identifier [32]byte) ([]common.Address, error) {
	return _Escrow.Contract.GetReceivedPaymentIds(&_Escrow.CallOpts, identifier)
}

// GetSentPaymentIds is a free data retrieval call binding the contract method 0x18d46532.
//
// Solidity: function getSentPaymentIds(address sender) view returns(address[])
func (_Escrow *EscrowCaller) GetSentPaymentIds(opts *bind.CallOpts, sender common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Escrow.contract.Call(opts, out, "getSentPaymentIds", sender)
	return *ret0, err
}

// GetSentPaymentIds is a free data retrieval call binding the contract method 0x18d46532.
//
// Solidity: function getSentPaymentIds(address sender) view returns(address[])
func (_Escrow *EscrowSession) GetSentPaymentIds(sender common.Address) ([]common.Address, error) {
	return _Escrow.Contract.GetSentPaymentIds(&_Escrow.CallOpts, sender)
}

// GetSentPaymentIds is a free data retrieval call binding the contract method 0x18d46532.
//
// Solidity: function getSentPaymentIds(address sender) view returns(address[])
func (_Escrow *EscrowCallerSession) GetSentPaymentIds(sender common.Address) ([]common.Address, error) {
	return _Escrow.Contract.GetSentPaymentIds(&_Escrow.CallOpts, sender)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_Escrow *EscrowCaller) GetVersionNumber(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Escrow.contract.Call(opts, out, "getVersionNumber")
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_Escrow *EscrowSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Escrow.Contract.GetVersionNumber(&_Escrow.CallOpts)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_Escrow *EscrowCallerSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Escrow.Contract.GetVersionNumber(&_Escrow.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Escrow *EscrowCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Escrow.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Escrow *EscrowSession) Initialized() (bool, error) {
	return _Escrow.Contract.Initialized(&_Escrow.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Escrow *EscrowCallerSession) Initialized() (bool, error) {
	return _Escrow.Contract.Initialized(&_Escrow.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Escrow *EscrowCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Escrow.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Escrow *EscrowSession) IsOwner() (bool, error) {
	return _Escrow.Contract.IsOwner(&_Escrow.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Escrow *EscrowCallerSession) IsOwner() (bool, error) {
	return _Escrow.Contract.IsOwner(&_Escrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Escrow *EscrowCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Escrow.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Escrow *EscrowSession) Owner() (common.Address, error) {
	return _Escrow.Contract.Owner(&_Escrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Escrow *EscrowCallerSession) Owner() (common.Address, error) {
	return _Escrow.Contract.Owner(&_Escrow.CallOpts)
}

// ReceivedPaymentIds is a free data retrieval call binding the contract method 0x8f80c33e.
//
// Solidity: function receivedPaymentIds(bytes32 , uint256 ) view returns(address)
func (_Escrow *EscrowCaller) ReceivedPaymentIds(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Escrow.contract.Call(opts, out, "receivedPaymentIds", arg0, arg1)
	return *ret0, err
}

// ReceivedPaymentIds is a free data retrieval call binding the contract method 0x8f80c33e.
//
// Solidity: function receivedPaymentIds(bytes32 , uint256 ) view returns(address)
func (_Escrow *EscrowSession) ReceivedPaymentIds(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _Escrow.Contract.ReceivedPaymentIds(&_Escrow.CallOpts, arg0, arg1)
}

// ReceivedPaymentIds is a free data retrieval call binding the contract method 0x8f80c33e.
//
// Solidity: function receivedPaymentIds(bytes32 , uint256 ) view returns(address)
func (_Escrow *EscrowCallerSession) ReceivedPaymentIds(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _Escrow.Contract.ReceivedPaymentIds(&_Escrow.CallOpts, arg0, arg1)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Escrow *EscrowCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Escrow.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Escrow *EscrowSession) Registry() (common.Address, error) {
	return _Escrow.Contract.Registry(&_Escrow.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Escrow *EscrowCallerSession) Registry() (common.Address, error) {
	return _Escrow.Contract.Registry(&_Escrow.CallOpts)
}

// SentPaymentIds is a free data retrieval call binding the contract method 0xe1d9a080.
//
// Solidity: function sentPaymentIds(address , uint256 ) view returns(address)
func (_Escrow *EscrowCaller) SentPaymentIds(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Escrow.contract.Call(opts, out, "sentPaymentIds", arg0, arg1)
	return *ret0, err
}

// SentPaymentIds is a free data retrieval call binding the contract method 0xe1d9a080.
//
// Solidity: function sentPaymentIds(address , uint256 ) view returns(address)
func (_Escrow *EscrowSession) SentPaymentIds(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Escrow.Contract.SentPaymentIds(&_Escrow.CallOpts, arg0, arg1)
}

// SentPaymentIds is a free data retrieval call binding the contract method 0xe1d9a080.
//
// Solidity: function sentPaymentIds(address , uint256 ) view returns(address)
func (_Escrow *EscrowCallerSession) SentPaymentIds(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Escrow.Contract.SentPaymentIds(&_Escrow.CallOpts, arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address registryAddress) returns()
func (_Escrow *EscrowTransactor) Initialize(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "initialize", registryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address registryAddress) returns()
func (_Escrow *EscrowSession) Initialize(registryAddress common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.Initialize(&_Escrow.TransactOpts, registryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address registryAddress) returns()
func (_Escrow *EscrowTransactorSession) Initialize(registryAddress common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.Initialize(&_Escrow.TransactOpts, registryAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Escrow *EscrowTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Escrow *EscrowSession) RenounceOwnership() (*types.Transaction, error) {
	return _Escrow.Contract.RenounceOwnership(&_Escrow.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Escrow *EscrowTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Escrow.Contract.RenounceOwnership(&_Escrow.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address paymentId) returns(bool)
func (_Escrow *EscrowTransactor) Revoke(opts *bind.TransactOpts, paymentId common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "revoke", paymentId)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address paymentId) returns(bool)
func (_Escrow *EscrowSession) Revoke(paymentId common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.Revoke(&_Escrow.TransactOpts, paymentId)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address paymentId) returns(bool)
func (_Escrow *EscrowTransactorSession) Revoke(paymentId common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.Revoke(&_Escrow.TransactOpts, paymentId)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Escrow *EscrowTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Escrow *EscrowSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.SetRegistry(&_Escrow.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Escrow *EscrowTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.SetRegistry(&_Escrow.TransactOpts, registryAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0x702cb75d.
//
// Solidity: function transfer(bytes32 identifier, address token, uint256 value, uint256 expirySeconds, address paymentId, uint256 minAttestations) returns(bool)
func (_Escrow *EscrowTransactor) Transfer(opts *bind.TransactOpts, identifier [32]byte, token common.Address, value *big.Int, expirySeconds *big.Int, paymentId common.Address, minAttestations *big.Int) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "transfer", identifier, token, value, expirySeconds, paymentId, minAttestations)
}

// Transfer is a paid mutator transaction binding the contract method 0x702cb75d.
//
// Solidity: function transfer(bytes32 identifier, address token, uint256 value, uint256 expirySeconds, address paymentId, uint256 minAttestations) returns(bool)
func (_Escrow *EscrowSession) Transfer(identifier [32]byte, token common.Address, value *big.Int, expirySeconds *big.Int, paymentId common.Address, minAttestations *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.Transfer(&_Escrow.TransactOpts, identifier, token, value, expirySeconds, paymentId, minAttestations)
}

// Transfer is a paid mutator transaction binding the contract method 0x702cb75d.
//
// Solidity: function transfer(bytes32 identifier, address token, uint256 value, uint256 expirySeconds, address paymentId, uint256 minAttestations) returns(bool)
func (_Escrow *EscrowTransactorSession) Transfer(identifier [32]byte, token common.Address, value *big.Int, expirySeconds *big.Int, paymentId common.Address, minAttestations *big.Int) (*types.Transaction, error) {
	return _Escrow.Contract.Transfer(&_Escrow.TransactOpts, identifier, token, value, expirySeconds, paymentId, minAttestations)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Escrow *EscrowTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Escrow *EscrowSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.TransferOwnership(&_Escrow.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Escrow *EscrowTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.TransferOwnership(&_Escrow.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3e68d5d7.
//
// Solidity: function withdraw(address paymentId, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Escrow *EscrowTransactor) Withdraw(opts *bind.TransactOpts, paymentId common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "withdraw", paymentId, v, r, s)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3e68d5d7.
//
// Solidity: function withdraw(address paymentId, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Escrow *EscrowSession) Withdraw(paymentId common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Escrow.Contract.Withdraw(&_Escrow.TransactOpts, paymentId, v, r, s)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3e68d5d7.
//
// Solidity: function withdraw(address paymentId, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Escrow *EscrowTransactorSession) Withdraw(paymentId common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Escrow.Contract.Withdraw(&_Escrow.TransactOpts, paymentId, v, r, s)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_Escrow *EscrowFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _Escrow.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "OwnershipTransferred":
		event, err = _Escrow.ParseOwnershipTransferred(log)
	case "RegistrySet":
		event, err = _Escrow.ParseRegistrySet(log)
	case "Revocation":
		event, err = _Escrow.ParseRevocation(log)
	case "Transfer":
		event, err = _Escrow.ParseTransfer(log)
	case "Withdrawal":
		event, err = _Escrow.ParseWithdrawal(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// EscrowOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Escrow contract.
type EscrowOwnershipTransferredIterator struct {
	Event *EscrowOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EscrowOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowOwnershipTransferred)
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
		it.Event = new(EscrowOwnershipTransferred)
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
func (it *EscrowOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowOwnershipTransferred represents a OwnershipTransferred event raised by the Escrow contract.
type EscrowOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Escrow *EscrowFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EscrowOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EscrowOwnershipTransferredIterator{contract: _Escrow.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Escrow *EscrowFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EscrowOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowOwnershipTransferred)
				if err := _Escrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Escrow *EscrowFilterer) ParseOwnershipTransferred(log types.Log) (*EscrowOwnershipTransferred, error) {
	event := new(EscrowOwnershipTransferred)
	if err := _Escrow.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EscrowRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the Escrow contract.
type EscrowRegistrySetIterator struct {
	Event *EscrowRegistrySet // Event containing the contract specifics and raw log

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
func (it *EscrowRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowRegistrySet)
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
		it.Event = new(EscrowRegistrySet)
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
func (it *EscrowRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowRegistrySet represents a RegistrySet event raised by the Escrow contract.
type EscrowRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Escrow *EscrowFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*EscrowRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &EscrowRegistrySetIterator{contract: _Escrow.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Escrow *EscrowFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *EscrowRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowRegistrySet)
				if err := _Escrow.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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
func (_Escrow *EscrowFilterer) ParseRegistrySet(log types.Log) (*EscrowRegistrySet, error) {
	event := new(EscrowRegistrySet)
	if err := _Escrow.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EscrowRevocationIterator is returned from FilterRevocation and is used to iterate over the raw logs and unpacked data for Revocation events raised by the Escrow contract.
type EscrowRevocationIterator struct {
	Event *EscrowRevocation // Event containing the contract specifics and raw log

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
func (it *EscrowRevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowRevocation)
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
		it.Event = new(EscrowRevocation)
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
func (it *EscrowRevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowRevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowRevocation represents a Revocation event raised by the Escrow contract.
type EscrowRevocation struct {
	Identifier [32]byte
	By         common.Address
	Token      common.Address
	Value      *big.Int
	PaymentId  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRevocation is a free log retrieval operation binding the contract event 0x6c464fad8039e6f09ec3a57a29f132cf2573d166833256960e2407eefff8f592.
//
// Solidity: event Revocation(bytes32 indexed identifier, address indexed by, address indexed token, uint256 value, address paymentId)
func (_Escrow *EscrowFilterer) FilterRevocation(opts *bind.FilterOpts, identifier [][32]byte, by []common.Address, token []common.Address) (*EscrowRevocationIterator, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "Revocation", identifierRule, byRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &EscrowRevocationIterator{contract: _Escrow.contract, event: "Revocation", logs: logs, sub: sub}, nil
}

// WatchRevocation is a free log subscription operation binding the contract event 0x6c464fad8039e6f09ec3a57a29f132cf2573d166833256960e2407eefff8f592.
//
// Solidity: event Revocation(bytes32 indexed identifier, address indexed by, address indexed token, uint256 value, address paymentId)
func (_Escrow *EscrowFilterer) WatchRevocation(opts *bind.WatchOpts, sink chan<- *EscrowRevocation, identifier [][32]byte, by []common.Address, token []common.Address) (event.Subscription, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "Revocation", identifierRule, byRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowRevocation)
				if err := _Escrow.contract.UnpackLog(event, "Revocation", log); err != nil {
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

// ParseRevocation is a log parse operation binding the contract event 0x6c464fad8039e6f09ec3a57a29f132cf2573d166833256960e2407eefff8f592.
//
// Solidity: event Revocation(bytes32 indexed identifier, address indexed by, address indexed token, uint256 value, address paymentId)
func (_Escrow *EscrowFilterer) ParseRevocation(log types.Log) (*EscrowRevocation, error) {
	event := new(EscrowRevocation)
	if err := _Escrow.contract.UnpackLog(event, "Revocation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EscrowTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Escrow contract.
type EscrowTransferIterator struct {
	Event *EscrowTransfer // Event containing the contract specifics and raw log

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
func (it *EscrowTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowTransfer)
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
		it.Event = new(EscrowTransfer)
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
func (it *EscrowTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowTransfer represents a Transfer event raised by the Escrow contract.
type EscrowTransfer struct {
	From            common.Address
	Identifier      [32]byte
	Token           common.Address
	Value           *big.Int
	PaymentId       common.Address
	MinAttestations *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x0fc2463e82c3b8a7868e75b68a76a144816d772687e5b09f45c02db37eedf4f6.
//
// Solidity: event Transfer(address indexed from, bytes32 indexed identifier, address indexed token, uint256 value, address paymentId, uint256 minAttestations)
func (_Escrow *EscrowFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, identifier [][32]byte, token []common.Address) (*EscrowTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "Transfer", fromRule, identifierRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &EscrowTransferIterator{contract: _Escrow.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x0fc2463e82c3b8a7868e75b68a76a144816d772687e5b09f45c02db37eedf4f6.
//
// Solidity: event Transfer(address indexed from, bytes32 indexed identifier, address indexed token, uint256 value, address paymentId, uint256 minAttestations)
func (_Escrow *EscrowFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EscrowTransfer, from []common.Address, identifier [][32]byte, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "Transfer", fromRule, identifierRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowTransfer)
				if err := _Escrow.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x0fc2463e82c3b8a7868e75b68a76a144816d772687e5b09f45c02db37eedf4f6.
//
// Solidity: event Transfer(address indexed from, bytes32 indexed identifier, address indexed token, uint256 value, address paymentId, uint256 minAttestations)
func (_Escrow *EscrowFilterer) ParseTransfer(log types.Log) (*EscrowTransfer, error) {
	event := new(EscrowTransfer)
	if err := _Escrow.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EscrowWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the Escrow contract.
type EscrowWithdrawalIterator struct {
	Event *EscrowWithdrawal // Event containing the contract specifics and raw log

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
func (it *EscrowWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowWithdrawal)
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
		it.Event = new(EscrowWithdrawal)
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
func (it *EscrowWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowWithdrawal represents a Withdrawal event raised by the Escrow contract.
type EscrowWithdrawal struct {
	Identifier [32]byte
	To         common.Address
	Token      common.Address
	Value      *big.Int
	PaymentId  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xab4f92d461fdbd1af5db2375223d65edb43bcb99129b19ab4954004883e52025.
//
// Solidity: event Withdrawal(bytes32 indexed identifier, address indexed to, address indexed token, uint256 value, address paymentId)
func (_Escrow *EscrowFilterer) FilterWithdrawal(opts *bind.FilterOpts, identifier [][32]byte, to []common.Address, token []common.Address) (*EscrowWithdrawalIterator, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "Withdrawal", identifierRule, toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &EscrowWithdrawalIterator{contract: _Escrow.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xab4f92d461fdbd1af5db2375223d65edb43bcb99129b19ab4954004883e52025.
//
// Solidity: event Withdrawal(bytes32 indexed identifier, address indexed to, address indexed token, uint256 value, address paymentId)
func (_Escrow *EscrowFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *EscrowWithdrawal, identifier [][32]byte, to []common.Address, token []common.Address) (event.Subscription, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "Withdrawal", identifierRule, toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowWithdrawal)
				if err := _Escrow.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0xab4f92d461fdbd1af5db2375223d65edb43bcb99129b19ab4954004883e52025.
//
// Solidity: event Withdrawal(bytes32 indexed identifier, address indexed to, address indexed token, uint256 value, address paymentId)
func (_Escrow *EscrowFilterer) ParseWithdrawal(log types.Log) (*EscrowWithdrawal, error) {
	event := new(EscrowWithdrawal)
	if err := _Escrow.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	return event, nil
}
