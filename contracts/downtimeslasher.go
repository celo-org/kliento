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

// DowntimeSlasherABI is the input ABI used to generate the binding from.
const DowntimeSlasherABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"slashingIncentives\",\"outputs\":[{\"name\":\"penalty\",\"type\":\"uint256\"},{\"name\":\"reward\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromCurrentSet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"blsKey\",\"type\":\"bytes\"},{\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"checkProofOfPossession\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEpochNumberOfBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slashableDowntime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getVerifiedSealBitmapFromHeader\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromSet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"hashHeader\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minQuorumSizeInCurrentSet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberValidatorsInCurrentSet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"groupMembershipHistoryIndex\",\"type\":\"uint256\"}],\"name\":\"groupMembershipAtBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getBlockNumberFromHeader\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"numberValidatorsInSet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"penalty\",\"type\":\"uint256\"},{\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"setSlashingIncentives\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"minQuorumSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"aNumerator\",\"type\":\"uint256\"},{\"name\":\"aDenominator\",\"type\":\"uint256\"},{\"name\":\"bNumerator\",\"type\":\"uint256\"},{\"name\":\"bDenominator\",\"type\":\"uint256\"},{\"name\":\"exponent\",\"type\":\"uint256\"},{\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"fractionMulExp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getParentSealBitmap\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"interval\",\"type\":\"uint256\"}],\"name\":\"SlashableDowntimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"startBlock\",\"type\":\"uint256\"}],\"name\":\"DowntimeSlashPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"penalty\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"SlashingIncentivesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"},{\"name\":\"_penalty\",\"type\":\"uint256\"},{\"name\":\"_reward\",\"type\":\"uint256\"},{\"name\":\"_slashableDowntime\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"interval\",\"type\":\"uint256\"}],\"name\":\"setSlashableDowntime\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"startBlock\",\"type\":\"uint256\"},{\"name\":\"startSignerIndex\",\"type\":\"uint256\"},{\"name\":\"endSignerIndex\",\"type\":\"uint256\"}],\"name\":\"isDown\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"startBlock\",\"type\":\"uint256\"},{\"name\":\"startSignerIndex\",\"type\":\"uint256\"},{\"name\":\"endSignerIndex\",\"type\":\"uint256\"},{\"name\":\"groupMembershipHistoryIndex\",\"type\":\"uint256\"},{\"name\":\"validatorElectionLessers\",\"type\":\"address[]\"},{\"name\":\"validatorElectionGreaters\",\"type\":\"address[]\"},{\"name\":\"validatorElectionIndices\",\"type\":\"uint256[]\"},{\"name\":\"groupElectionLessers\",\"type\":\"address[]\"},{\"name\":\"groupElectionGreaters\",\"type\":\"address[]\"},{\"name\":\"groupElectionIndices\",\"type\":\"uint256[]\"}],\"name\":\"slash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DowntimeSlasher is an auto generated Go binding around an Ethereum contract.
type DowntimeSlasher struct {
	DowntimeSlasherCaller     // Read-only binding to the contract
	DowntimeSlasherTransactor // Write-only binding to the contract
	DowntimeSlasherFilterer   // Log filterer for contract events
}

// DowntimeSlasherCaller is an auto generated read-only Go binding around an Ethereum contract.
type DowntimeSlasherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DowntimeSlasherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DowntimeSlasherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DowntimeSlasherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DowntimeSlasherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DowntimeSlasherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DowntimeSlasherSession struct {
	Contract     *DowntimeSlasher  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DowntimeSlasherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DowntimeSlasherCallerSession struct {
	Contract *DowntimeSlasherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DowntimeSlasherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DowntimeSlasherTransactorSession struct {
	Contract     *DowntimeSlasherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DowntimeSlasherRaw is an auto generated low-level Go binding around an Ethereum contract.
type DowntimeSlasherRaw struct {
	Contract *DowntimeSlasher // Generic contract binding to access the raw methods on
}

// DowntimeSlasherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DowntimeSlasherCallerRaw struct {
	Contract *DowntimeSlasherCaller // Generic read-only contract binding to access the raw methods on
}

// DowntimeSlasherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DowntimeSlasherTransactorRaw struct {
	Contract *DowntimeSlasherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDowntimeSlasher creates a new instance of DowntimeSlasher, bound to a specific deployed contract.
func NewDowntimeSlasher(address common.Address, backend bind.ContractBackend) (*DowntimeSlasher, error) {
	contract, err := bindDowntimeSlasher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DowntimeSlasher{DowntimeSlasherCaller: DowntimeSlasherCaller{contract: contract}, DowntimeSlasherTransactor: DowntimeSlasherTransactor{contract: contract}, DowntimeSlasherFilterer: DowntimeSlasherFilterer{contract: contract}}, nil
}

// NewDowntimeSlasherCaller creates a new read-only instance of DowntimeSlasher, bound to a specific deployed contract.
func NewDowntimeSlasherCaller(address common.Address, caller bind.ContractCaller) (*DowntimeSlasherCaller, error) {
	contract, err := bindDowntimeSlasher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DowntimeSlasherCaller{contract: contract}, nil
}

// NewDowntimeSlasherTransactor creates a new write-only instance of DowntimeSlasher, bound to a specific deployed contract.
func NewDowntimeSlasherTransactor(address common.Address, transactor bind.ContractTransactor) (*DowntimeSlasherTransactor, error) {
	contract, err := bindDowntimeSlasher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DowntimeSlasherTransactor{contract: contract}, nil
}

// NewDowntimeSlasherFilterer creates a new log filterer instance of DowntimeSlasher, bound to a specific deployed contract.
func NewDowntimeSlasherFilterer(address common.Address, filterer bind.ContractFilterer) (*DowntimeSlasherFilterer, error) {
	contract, err := bindDowntimeSlasher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DowntimeSlasherFilterer{contract: contract}, nil
}

// bindDowntimeSlasher binds a generic wrapper to an already deployed contract.
func bindDowntimeSlasher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DowntimeSlasherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseDowntimeSlasherABI parses the ABI
func ParseDowntimeSlasherABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(DowntimeSlasherABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DowntimeSlasher *DowntimeSlasherRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DowntimeSlasher.Contract.DowntimeSlasherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DowntimeSlasher *DowntimeSlasherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.DowntimeSlasherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DowntimeSlasher *DowntimeSlasherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.DowntimeSlasherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DowntimeSlasher *DowntimeSlasherCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DowntimeSlasher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DowntimeSlasher *DowntimeSlasherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DowntimeSlasher *DowntimeSlasherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.contract.Transact(opts, method, params...)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) constant returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCaller) CheckProofOfPossession(opts *bind.CallOpts, sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "checkProofOfPossession", sender, blsKey, blsPop)
	return *ret0, err
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) constant returns(bool)
func (_DowntimeSlasher *DowntimeSlasherSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _DowntimeSlasher.Contract.CheckProofOfPossession(&_DowntimeSlasher.CallOpts, sender, blsKey, blsPop)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) constant returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _DowntimeSlasher.Contract.CheckProofOfPossession(&_DowntimeSlasher.CallOpts, sender, blsKey, blsPop)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) constant returns(uint256, uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) FractionMulExp(opts *bind.CallOpts, aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DowntimeSlasher.contract.Call(opts, out, "fractionMulExp", aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
	return *ret0, *ret1, err
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) constant returns(uint256, uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _DowntimeSlasher.Contract.FractionMulExp(&_DowntimeSlasher.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) constant returns(uint256, uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _DowntimeSlasher.Contract.FractionMulExp(&_DowntimeSlasher.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) GetBlockNumberFromHeader(opts *bind.CallOpts, header []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "getBlockNumberFromHeader", header)
	return *ret0, err
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _DowntimeSlasher.Contract.GetBlockNumberFromHeader(&_DowntimeSlasher.CallOpts, header)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _DowntimeSlasher.Contract.GetBlockNumberFromHeader(&_DowntimeSlasher.CallOpts, header)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) GetEpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "getEpochNumber")
	return *ret0, err
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) GetEpochNumber() (*big.Int, error) {
	return _DowntimeSlasher.Contract.GetEpochNumber(&_DowntimeSlasher.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) GetEpochNumber() (*big.Int, error) {
	return _DowntimeSlasher.Contract.GetEpochNumber(&_DowntimeSlasher.CallOpts)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) GetEpochNumberOfBlock(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "getEpochNumberOfBlock", blockNumber)
	return *ret0, err
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _DowntimeSlasher.Contract.GetEpochNumberOfBlock(&_DowntimeSlasher.CallOpts, blockNumber)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _DowntimeSlasher.Contract.GetEpochNumberOfBlock(&_DowntimeSlasher.CallOpts, blockNumber)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) GetEpochSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "getEpochSize")
	return *ret0, err
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) GetEpochSize() (*big.Int, error) {
	return _DowntimeSlasher.Contract.GetEpochSize(&_DowntimeSlasher.CallOpts)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) GetEpochSize() (*big.Int, error) {
	return _DowntimeSlasher.Contract.GetEpochSize(&_DowntimeSlasher.CallOpts)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) constant returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherCaller) GetParentSealBitmap(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "getParentSealBitmap", blockNumber)
	return *ret0, err
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) constant returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _DowntimeSlasher.Contract.GetParentSealBitmap(&_DowntimeSlasher.CallOpts, blockNumber)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) constant returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _DowntimeSlasher.Contract.GetParentSealBitmap(&_DowntimeSlasher.CallOpts, blockNumber)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) constant returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherCaller) GetVerifiedSealBitmapFromHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "getVerifiedSealBitmapFromHeader", header)
	return *ret0, err
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) constant returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _DowntimeSlasher.Contract.GetVerifiedSealBitmapFromHeader(&_DowntimeSlasher.CallOpts, header)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) constant returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _DowntimeSlasher.Contract.GetVerifiedSealBitmapFromHeader(&_DowntimeSlasher.CallOpts, header)
}

// GroupMembershipAtBlock is a free data retrieval call binding the contract method 0x88498aaf.
//
// Solidity: function groupMembershipAtBlock(address validator, uint256 blockNumber, uint256 groupMembershipHistoryIndex) constant returns(address)
func (_DowntimeSlasher *DowntimeSlasherCaller) GroupMembershipAtBlock(opts *bind.CallOpts, validator common.Address, blockNumber *big.Int, groupMembershipHistoryIndex *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "groupMembershipAtBlock", validator, blockNumber, groupMembershipHistoryIndex)
	return *ret0, err
}

// GroupMembershipAtBlock is a free data retrieval call binding the contract method 0x88498aaf.
//
// Solidity: function groupMembershipAtBlock(address validator, uint256 blockNumber, uint256 groupMembershipHistoryIndex) constant returns(address)
func (_DowntimeSlasher *DowntimeSlasherSession) GroupMembershipAtBlock(validator common.Address, blockNumber *big.Int, groupMembershipHistoryIndex *big.Int) (common.Address, error) {
	return _DowntimeSlasher.Contract.GroupMembershipAtBlock(&_DowntimeSlasher.CallOpts, validator, blockNumber, groupMembershipHistoryIndex)
}

// GroupMembershipAtBlock is a free data retrieval call binding the contract method 0x88498aaf.
//
// Solidity: function groupMembershipAtBlock(address validator, uint256 blockNumber, uint256 groupMembershipHistoryIndex) constant returns(address)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) GroupMembershipAtBlock(validator common.Address, blockNumber *big.Int, groupMembershipHistoryIndex *big.Int) (common.Address, error) {
	return _DowntimeSlasher.Contract.GroupMembershipAtBlock(&_DowntimeSlasher.CallOpts, validator, blockNumber, groupMembershipHistoryIndex)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) constant returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherCaller) HashHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "hashHeader", header)
	return *ret0, err
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) constant returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherSession) HashHeader(header []byte) ([32]byte, error) {
	return _DowntimeSlasher.Contract.HashHeader(&_DowntimeSlasher.CallOpts, header)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) constant returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) HashHeader(header []byte) ([32]byte, error) {
	return _DowntimeSlasher.Contract.HashHeader(&_DowntimeSlasher.CallOpts, header)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_DowntimeSlasher *DowntimeSlasherSession) Initialized() (bool, error) {
	return _DowntimeSlasher.Contract.Initialized(&_DowntimeSlasher.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) Initialized() (bool, error) {
	return _DowntimeSlasher.Contract.Initialized(&_DowntimeSlasher.CallOpts)
}

// IsDown is a free data retrieval call binding the contract method 0x86bcf1e6.
//
// Solidity: function isDown(uint256 startBlock, uint256 startSignerIndex, uint256 endSignerIndex) constant returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCaller) IsDown(opts *bind.CallOpts, startBlock *big.Int, startSignerIndex *big.Int, endSignerIndex *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "isDown", startBlock, startSignerIndex, endSignerIndex)
	return *ret0, err
}

// IsDown is a free data retrieval call binding the contract method 0x86bcf1e6.
//
// Solidity: function isDown(uint256 startBlock, uint256 startSignerIndex, uint256 endSignerIndex) constant returns(bool)
func (_DowntimeSlasher *DowntimeSlasherSession) IsDown(startBlock *big.Int, startSignerIndex *big.Int, endSignerIndex *big.Int) (bool, error) {
	return _DowntimeSlasher.Contract.IsDown(&_DowntimeSlasher.CallOpts, startBlock, startSignerIndex, endSignerIndex)
}

// IsDown is a free data retrieval call binding the contract method 0x86bcf1e6.
//
// Solidity: function isDown(uint256 startBlock, uint256 startSignerIndex, uint256 endSignerIndex) constant returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) IsDown(startBlock *big.Int, startSignerIndex *big.Int, endSignerIndex *big.Int) (bool, error) {
	return _DowntimeSlasher.Contract.IsDown(&_DowntimeSlasher.CallOpts, startBlock, startSignerIndex, endSignerIndex)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DowntimeSlasher *DowntimeSlasherSession) IsOwner() (bool, error) {
	return _DowntimeSlasher.Contract.IsOwner(&_DowntimeSlasher.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) IsOwner() (bool, error) {
	return _DowntimeSlasher.Contract.IsOwner(&_DowntimeSlasher.CallOpts)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) MinQuorumSize(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "minQuorumSize", blockNumber)
	return *ret0, err
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _DowntimeSlasher.Contract.MinQuorumSize(&_DowntimeSlasher.CallOpts, blockNumber)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _DowntimeSlasher.Contract.MinQuorumSize(&_DowntimeSlasher.CallOpts, blockNumber)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) MinQuorumSizeInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "minQuorumSizeInCurrentSet")
	return *ret0, err
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _DowntimeSlasher.Contract.MinQuorumSizeInCurrentSet(&_DowntimeSlasher.CallOpts)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _DowntimeSlasher.Contract.MinQuorumSizeInCurrentSet(&_DowntimeSlasher.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) NumberValidatorsInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "numberValidatorsInCurrentSet")
	return *ret0, err
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _DowntimeSlasher.Contract.NumberValidatorsInCurrentSet(&_DowntimeSlasher.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _DowntimeSlasher.Contract.NumberValidatorsInCurrentSet(&_DowntimeSlasher.CallOpts)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) NumberValidatorsInSet(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "numberValidatorsInSet", blockNumber)
	return *ret0, err
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _DowntimeSlasher.Contract.NumberValidatorsInSet(&_DowntimeSlasher.CallOpts, blockNumber)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _DowntimeSlasher.Contract.NumberValidatorsInSet(&_DowntimeSlasher.CallOpts, blockNumber)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DowntimeSlasher *DowntimeSlasherCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DowntimeSlasher *DowntimeSlasherSession) Owner() (common.Address, error) {
	return _DowntimeSlasher.Contract.Owner(&_DowntimeSlasher.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) Owner() (common.Address, error) {
	return _DowntimeSlasher.Contract.Owner(&_DowntimeSlasher.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_DowntimeSlasher *DowntimeSlasherCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_DowntimeSlasher *DowntimeSlasherSession) Registry() (common.Address, error) {
	return _DowntimeSlasher.Contract.Registry(&_DowntimeSlasher.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) Registry() (common.Address, error) {
	return _DowntimeSlasher.Contract.Registry(&_DowntimeSlasher.CallOpts)
}

// SlashableDowntime is a free data retrieval call binding the contract method 0x4227d971.
//
// Solidity: function slashableDowntime() constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) SlashableDowntime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "slashableDowntime")
	return *ret0, err
}

// SlashableDowntime is a free data retrieval call binding the contract method 0x4227d971.
//
// Solidity: function slashableDowntime() constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) SlashableDowntime() (*big.Int, error) {
	return _DowntimeSlasher.Contract.SlashableDowntime(&_DowntimeSlasher.CallOpts)
}

// SlashableDowntime is a free data retrieval call binding the contract method 0x4227d971.
//
// Solidity: function slashableDowntime() constant returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) SlashableDowntime() (*big.Int, error) {
	return _DowntimeSlasher.Contract.SlashableDowntime(&_DowntimeSlasher.CallOpts)
}

// SlashingIncentives is a free data retrieval call binding the contract method 0x0a05cd84.
//
// Solidity: function slashingIncentives() constant returns(uint256 penalty, uint256 reward)
func (_DowntimeSlasher *DowntimeSlasherCaller) SlashingIncentives(opts *bind.CallOpts) (struct {
	Penalty *big.Int
	Reward  *big.Int
}, error) {
	ret := new(struct {
		Penalty *big.Int
		Reward  *big.Int
	})
	out := ret
	err := _DowntimeSlasher.contract.Call(opts, out, "slashingIncentives")
	return *ret, err
}

// SlashingIncentives is a free data retrieval call binding the contract method 0x0a05cd84.
//
// Solidity: function slashingIncentives() constant returns(uint256 penalty, uint256 reward)
func (_DowntimeSlasher *DowntimeSlasherSession) SlashingIncentives() (struct {
	Penalty *big.Int
	Reward  *big.Int
}, error) {
	return _DowntimeSlasher.Contract.SlashingIncentives(&_DowntimeSlasher.CallOpts)
}

// SlashingIncentives is a free data retrieval call binding the contract method 0x0a05cd84.
//
// Solidity: function slashingIncentives() constant returns(uint256 penalty, uint256 reward)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) SlashingIncentives() (struct {
	Penalty *big.Int
	Reward  *big.Int
}, error) {
	return _DowntimeSlasher.Contract.SlashingIncentives(&_DowntimeSlasher.CallOpts)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) constant returns(address)
func (_DowntimeSlasher *DowntimeSlasherCaller) ValidatorSignerAddressFromCurrentSet(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "validatorSignerAddressFromCurrentSet", index)
	return *ret0, err
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) constant returns(address)
func (_DowntimeSlasher *DowntimeSlasherSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _DowntimeSlasher.Contract.ValidatorSignerAddressFromCurrentSet(&_DowntimeSlasher.CallOpts, index)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) constant returns(address)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _DowntimeSlasher.Contract.ValidatorSignerAddressFromCurrentSet(&_DowntimeSlasher.CallOpts, index)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) constant returns(address)
func (_DowntimeSlasher *DowntimeSlasherCaller) ValidatorSignerAddressFromSet(opts *bind.CallOpts, index *big.Int, blockNumber *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DowntimeSlasher.contract.Call(opts, out, "validatorSignerAddressFromSet", index, blockNumber)
	return *ret0, err
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) constant returns(address)
func (_DowntimeSlasher *DowntimeSlasherSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _DowntimeSlasher.Contract.ValidatorSignerAddressFromSet(&_DowntimeSlasher.CallOpts, index, blockNumber)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) constant returns(address)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _DowntimeSlasher.Contract.ValidatorSignerAddressFromSet(&_DowntimeSlasher.CallOpts, index, blockNumber)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address registryAddress, uint256 _penalty, uint256 _reward, uint256 _slashableDowntime) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactor) Initialize(opts *bind.TransactOpts, registryAddress common.Address, _penalty *big.Int, _reward *big.Int, _slashableDowntime *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.contract.Transact(opts, "initialize", registryAddress, _penalty, _reward, _slashableDowntime)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address registryAddress, uint256 _penalty, uint256 _reward, uint256 _slashableDowntime) returns()
func (_DowntimeSlasher *DowntimeSlasherSession) Initialize(registryAddress common.Address, _penalty *big.Int, _reward *big.Int, _slashableDowntime *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.Initialize(&_DowntimeSlasher.TransactOpts, registryAddress, _penalty, _reward, _slashableDowntime)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address registryAddress, uint256 _penalty, uint256 _reward, uint256 _slashableDowntime) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactorSession) Initialize(registryAddress common.Address, _penalty *big.Int, _reward *big.Int, _slashableDowntime *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.Initialize(&_DowntimeSlasher.TransactOpts, registryAddress, _penalty, _reward, _slashableDowntime)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DowntimeSlasher *DowntimeSlasherTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DowntimeSlasher.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DowntimeSlasher *DowntimeSlasherSession) RenounceOwnership() (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.RenounceOwnership(&_DowntimeSlasher.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DowntimeSlasher *DowntimeSlasherTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.RenounceOwnership(&_DowntimeSlasher.TransactOpts)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _DowntimeSlasher.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_DowntimeSlasher *DowntimeSlasherSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.SetRegistry(&_DowntimeSlasher.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.SetRegistry(&_DowntimeSlasher.TransactOpts, registryAddress)
}

// SetSlashableDowntime is a paid mutator transaction binding the contract method 0x4d643e17.
//
// Solidity: function setSlashableDowntime(uint256 interval) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactor) SetSlashableDowntime(opts *bind.TransactOpts, interval *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.contract.Transact(opts, "setSlashableDowntime", interval)
}

// SetSlashableDowntime is a paid mutator transaction binding the contract method 0x4d643e17.
//
// Solidity: function setSlashableDowntime(uint256 interval) returns()
func (_DowntimeSlasher *DowntimeSlasherSession) SetSlashableDowntime(interval *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.SetSlashableDowntime(&_DowntimeSlasher.TransactOpts, interval)
}

// SetSlashableDowntime is a paid mutator transaction binding the contract method 0x4d643e17.
//
// Solidity: function setSlashableDowntime(uint256 interval) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactorSession) SetSlashableDowntime(interval *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.SetSlashableDowntime(&_DowntimeSlasher.TransactOpts, interval)
}

// SetSlashingIncentives is a paid mutator transaction binding the contract method 0xbd0d9979.
//
// Solidity: function setSlashingIncentives(uint256 penalty, uint256 reward) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactor) SetSlashingIncentives(opts *bind.TransactOpts, penalty *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.contract.Transact(opts, "setSlashingIncentives", penalty, reward)
}

// SetSlashingIncentives is a paid mutator transaction binding the contract method 0xbd0d9979.
//
// Solidity: function setSlashingIncentives(uint256 penalty, uint256 reward) returns()
func (_DowntimeSlasher *DowntimeSlasherSession) SetSlashingIncentives(penalty *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.SetSlashingIncentives(&_DowntimeSlasher.TransactOpts, penalty, reward)
}

// SetSlashingIncentives is a paid mutator transaction binding the contract method 0xbd0d9979.
//
// Solidity: function setSlashingIncentives(uint256 penalty, uint256 reward) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactorSession) SetSlashingIncentives(penalty *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.SetSlashingIncentives(&_DowntimeSlasher.TransactOpts, penalty, reward)
}

// Slash is a paid mutator transaction binding the contract method 0xf7499291.
//
// Solidity: function slash(uint256 startBlock, uint256 startSignerIndex, uint256 endSignerIndex, uint256 groupMembershipHistoryIndex, address[] validatorElectionLessers, address[] validatorElectionGreaters, uint256[] validatorElectionIndices, address[] groupElectionLessers, address[] groupElectionGreaters, uint256[] groupElectionIndices) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactor) Slash(opts *bind.TransactOpts, startBlock *big.Int, startSignerIndex *big.Int, endSignerIndex *big.Int, groupMembershipHistoryIndex *big.Int, validatorElectionLessers []common.Address, validatorElectionGreaters []common.Address, validatorElectionIndices []*big.Int, groupElectionLessers []common.Address, groupElectionGreaters []common.Address, groupElectionIndices []*big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.contract.Transact(opts, "slash", startBlock, startSignerIndex, endSignerIndex, groupMembershipHistoryIndex, validatorElectionLessers, validatorElectionGreaters, validatorElectionIndices, groupElectionLessers, groupElectionGreaters, groupElectionIndices)
}

// Slash is a paid mutator transaction binding the contract method 0xf7499291.
//
// Solidity: function slash(uint256 startBlock, uint256 startSignerIndex, uint256 endSignerIndex, uint256 groupMembershipHistoryIndex, address[] validatorElectionLessers, address[] validatorElectionGreaters, uint256[] validatorElectionIndices, address[] groupElectionLessers, address[] groupElectionGreaters, uint256[] groupElectionIndices) returns()
func (_DowntimeSlasher *DowntimeSlasherSession) Slash(startBlock *big.Int, startSignerIndex *big.Int, endSignerIndex *big.Int, groupMembershipHistoryIndex *big.Int, validatorElectionLessers []common.Address, validatorElectionGreaters []common.Address, validatorElectionIndices []*big.Int, groupElectionLessers []common.Address, groupElectionGreaters []common.Address, groupElectionIndices []*big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.Slash(&_DowntimeSlasher.TransactOpts, startBlock, startSignerIndex, endSignerIndex, groupMembershipHistoryIndex, validatorElectionLessers, validatorElectionGreaters, validatorElectionIndices, groupElectionLessers, groupElectionGreaters, groupElectionIndices)
}

// Slash is a paid mutator transaction binding the contract method 0xf7499291.
//
// Solidity: function slash(uint256 startBlock, uint256 startSignerIndex, uint256 endSignerIndex, uint256 groupMembershipHistoryIndex, address[] validatorElectionLessers, address[] validatorElectionGreaters, uint256[] validatorElectionIndices, address[] groupElectionLessers, address[] groupElectionGreaters, uint256[] groupElectionIndices) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactorSession) Slash(startBlock *big.Int, startSignerIndex *big.Int, endSignerIndex *big.Int, groupMembershipHistoryIndex *big.Int, validatorElectionLessers []common.Address, validatorElectionGreaters []common.Address, validatorElectionIndices []*big.Int, groupElectionLessers []common.Address, groupElectionGreaters []common.Address, groupElectionIndices []*big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.Slash(&_DowntimeSlasher.TransactOpts, startBlock, startSignerIndex, endSignerIndex, groupMembershipHistoryIndex, validatorElectionLessers, validatorElectionGreaters, validatorElectionIndices, groupElectionLessers, groupElectionGreaters, groupElectionIndices)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DowntimeSlasher.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DowntimeSlasher *DowntimeSlasherSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.TransferOwnership(&_DowntimeSlasher.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.TransferOwnership(&_DowntimeSlasher.TransactOpts, newOwner)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_DowntimeSlasher *DowntimeSlasherFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _DowntimeSlasher.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "DowntimeSlashPerformed":
		event, err = _DowntimeSlasher.ParseDowntimeSlashPerformed(log)
	case "OwnershipTransferred":
		event, err = _DowntimeSlasher.ParseOwnershipTransferred(log)
	case "RegistrySet":
		event, err = _DowntimeSlasher.ParseRegistrySet(log)
	case "SlashableDowntimeSet":
		event, err = _DowntimeSlasher.ParseSlashableDowntimeSet(log)
	case "SlashingIncentivesSet":
		event, err = _DowntimeSlasher.ParseSlashingIncentivesSet(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// DowntimeSlasherDowntimeSlashPerformedIterator is returned from FilterDowntimeSlashPerformed and is used to iterate over the raw logs and unpacked data for DowntimeSlashPerformed events raised by the DowntimeSlasher contract.
type DowntimeSlasherDowntimeSlashPerformedIterator struct {
	Event *DowntimeSlasherDowntimeSlashPerformed // Event containing the contract specifics and raw log

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
func (it *DowntimeSlasherDowntimeSlashPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DowntimeSlasherDowntimeSlashPerformed)
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
		it.Event = new(DowntimeSlasherDowntimeSlashPerformed)
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
func (it *DowntimeSlasherDowntimeSlashPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DowntimeSlasherDowntimeSlashPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DowntimeSlasherDowntimeSlashPerformed represents a DowntimeSlashPerformed event raised by the DowntimeSlasher contract.
type DowntimeSlasherDowntimeSlashPerformed struct {
	Validator  common.Address
	StartBlock *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDowntimeSlashPerformed is a free log retrieval operation binding the contract event 0x8c83b12f71cc3bbf5922a9600018df883cfe2e1f40bfbc3856fb794d69b2ae54.
//
// Solidity: event DowntimeSlashPerformed(address indexed validator, uint256 indexed startBlock)
func (_DowntimeSlasher *DowntimeSlasherFilterer) FilterDowntimeSlashPerformed(opts *bind.FilterOpts, validator []common.Address, startBlock []*big.Int) (*DowntimeSlasherDowntimeSlashPerformedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var startBlockRule []interface{}
	for _, startBlockItem := range startBlock {
		startBlockRule = append(startBlockRule, startBlockItem)
	}

	logs, sub, err := _DowntimeSlasher.contract.FilterLogs(opts, "DowntimeSlashPerformed", validatorRule, startBlockRule)
	if err != nil {
		return nil, err
	}
	return &DowntimeSlasherDowntimeSlashPerformedIterator{contract: _DowntimeSlasher.contract, event: "DowntimeSlashPerformed", logs: logs, sub: sub}, nil
}

// WatchDowntimeSlashPerformed is a free log subscription operation binding the contract event 0x8c83b12f71cc3bbf5922a9600018df883cfe2e1f40bfbc3856fb794d69b2ae54.
//
// Solidity: event DowntimeSlashPerformed(address indexed validator, uint256 indexed startBlock)
func (_DowntimeSlasher *DowntimeSlasherFilterer) WatchDowntimeSlashPerformed(opts *bind.WatchOpts, sink chan<- *DowntimeSlasherDowntimeSlashPerformed, validator []common.Address, startBlock []*big.Int) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var startBlockRule []interface{}
	for _, startBlockItem := range startBlock {
		startBlockRule = append(startBlockRule, startBlockItem)
	}

	logs, sub, err := _DowntimeSlasher.contract.WatchLogs(opts, "DowntimeSlashPerformed", validatorRule, startBlockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DowntimeSlasherDowntimeSlashPerformed)
				if err := _DowntimeSlasher.contract.UnpackLog(event, "DowntimeSlashPerformed", log); err != nil {
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

// ParseDowntimeSlashPerformed is a log parse operation binding the contract event 0x8c83b12f71cc3bbf5922a9600018df883cfe2e1f40bfbc3856fb794d69b2ae54.
//
// Solidity: event DowntimeSlashPerformed(address indexed validator, uint256 indexed startBlock)
func (_DowntimeSlasher *DowntimeSlasherFilterer) ParseDowntimeSlashPerformed(log types.Log) (*DowntimeSlasherDowntimeSlashPerformed, error) {
	event := new(DowntimeSlasherDowntimeSlashPerformed)
	if err := _DowntimeSlasher.contract.UnpackLog(event, "DowntimeSlashPerformed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DowntimeSlasherOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DowntimeSlasher contract.
type DowntimeSlasherOwnershipTransferredIterator struct {
	Event *DowntimeSlasherOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DowntimeSlasherOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DowntimeSlasherOwnershipTransferred)
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
		it.Event = new(DowntimeSlasherOwnershipTransferred)
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
func (it *DowntimeSlasherOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DowntimeSlasherOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DowntimeSlasherOwnershipTransferred represents a OwnershipTransferred event raised by the DowntimeSlasher contract.
type DowntimeSlasherOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DowntimeSlasher *DowntimeSlasherFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DowntimeSlasherOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DowntimeSlasher.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DowntimeSlasherOwnershipTransferredIterator{contract: _DowntimeSlasher.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DowntimeSlasher *DowntimeSlasherFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DowntimeSlasherOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DowntimeSlasher.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DowntimeSlasherOwnershipTransferred)
				if err := _DowntimeSlasher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DowntimeSlasher *DowntimeSlasherFilterer) ParseOwnershipTransferred(log types.Log) (*DowntimeSlasherOwnershipTransferred, error) {
	event := new(DowntimeSlasherOwnershipTransferred)
	if err := _DowntimeSlasher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DowntimeSlasherRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the DowntimeSlasher contract.
type DowntimeSlasherRegistrySetIterator struct {
	Event *DowntimeSlasherRegistrySet // Event containing the contract specifics and raw log

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
func (it *DowntimeSlasherRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DowntimeSlasherRegistrySet)
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
		it.Event = new(DowntimeSlasherRegistrySet)
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
func (it *DowntimeSlasherRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DowntimeSlasherRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DowntimeSlasherRegistrySet represents a RegistrySet event raised by the DowntimeSlasher contract.
type DowntimeSlasherRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_DowntimeSlasher *DowntimeSlasherFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*DowntimeSlasherRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _DowntimeSlasher.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &DowntimeSlasherRegistrySetIterator{contract: _DowntimeSlasher.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_DowntimeSlasher *DowntimeSlasherFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *DowntimeSlasherRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _DowntimeSlasher.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DowntimeSlasherRegistrySet)
				if err := _DowntimeSlasher.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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
func (_DowntimeSlasher *DowntimeSlasherFilterer) ParseRegistrySet(log types.Log) (*DowntimeSlasherRegistrySet, error) {
	event := new(DowntimeSlasherRegistrySet)
	if err := _DowntimeSlasher.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DowntimeSlasherSlashableDowntimeSetIterator is returned from FilterSlashableDowntimeSet and is used to iterate over the raw logs and unpacked data for SlashableDowntimeSet events raised by the DowntimeSlasher contract.
type DowntimeSlasherSlashableDowntimeSetIterator struct {
	Event *DowntimeSlasherSlashableDowntimeSet // Event containing the contract specifics and raw log

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
func (it *DowntimeSlasherSlashableDowntimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DowntimeSlasherSlashableDowntimeSet)
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
		it.Event = new(DowntimeSlasherSlashableDowntimeSet)
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
func (it *DowntimeSlasherSlashableDowntimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DowntimeSlasherSlashableDowntimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DowntimeSlasherSlashableDowntimeSet represents a SlashableDowntimeSet event raised by the DowntimeSlasher contract.
type DowntimeSlasherSlashableDowntimeSet struct {
	Interval *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSlashableDowntimeSet is a free log retrieval operation binding the contract event 0xc3293b70d45615822039f6f13747ece88efbbb4e645c42070413a6c3fd21d771.
//
// Solidity: event SlashableDowntimeSet(uint256 interval)
func (_DowntimeSlasher *DowntimeSlasherFilterer) FilterSlashableDowntimeSet(opts *bind.FilterOpts) (*DowntimeSlasherSlashableDowntimeSetIterator, error) {

	logs, sub, err := _DowntimeSlasher.contract.FilterLogs(opts, "SlashableDowntimeSet")
	if err != nil {
		return nil, err
	}
	return &DowntimeSlasherSlashableDowntimeSetIterator{contract: _DowntimeSlasher.contract, event: "SlashableDowntimeSet", logs: logs, sub: sub}, nil
}

// WatchSlashableDowntimeSet is a free log subscription operation binding the contract event 0xc3293b70d45615822039f6f13747ece88efbbb4e645c42070413a6c3fd21d771.
//
// Solidity: event SlashableDowntimeSet(uint256 interval)
func (_DowntimeSlasher *DowntimeSlasherFilterer) WatchSlashableDowntimeSet(opts *bind.WatchOpts, sink chan<- *DowntimeSlasherSlashableDowntimeSet) (event.Subscription, error) {

	logs, sub, err := _DowntimeSlasher.contract.WatchLogs(opts, "SlashableDowntimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DowntimeSlasherSlashableDowntimeSet)
				if err := _DowntimeSlasher.contract.UnpackLog(event, "SlashableDowntimeSet", log); err != nil {
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

// ParseSlashableDowntimeSet is a log parse operation binding the contract event 0xc3293b70d45615822039f6f13747ece88efbbb4e645c42070413a6c3fd21d771.
//
// Solidity: event SlashableDowntimeSet(uint256 interval)
func (_DowntimeSlasher *DowntimeSlasherFilterer) ParseSlashableDowntimeSet(log types.Log) (*DowntimeSlasherSlashableDowntimeSet, error) {
	event := new(DowntimeSlasherSlashableDowntimeSet)
	if err := _DowntimeSlasher.contract.UnpackLog(event, "SlashableDowntimeSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DowntimeSlasherSlashingIncentivesSetIterator is returned from FilterSlashingIncentivesSet and is used to iterate over the raw logs and unpacked data for SlashingIncentivesSet events raised by the DowntimeSlasher contract.
type DowntimeSlasherSlashingIncentivesSetIterator struct {
	Event *DowntimeSlasherSlashingIncentivesSet // Event containing the contract specifics and raw log

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
func (it *DowntimeSlasherSlashingIncentivesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DowntimeSlasherSlashingIncentivesSet)
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
		it.Event = new(DowntimeSlasherSlashingIncentivesSet)
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
func (it *DowntimeSlasherSlashingIncentivesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DowntimeSlasherSlashingIncentivesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DowntimeSlasherSlashingIncentivesSet represents a SlashingIncentivesSet event raised by the DowntimeSlasher contract.
type DowntimeSlasherSlashingIncentivesSet struct {
	Penalty *big.Int
	Reward  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSlashingIncentivesSet is a free log retrieval operation binding the contract event 0x716dc7c34384df36c6ccc5a2949f2ce9b019f5d4075ef39139a80038a4fdd1c3.
//
// Solidity: event SlashingIncentivesSet(uint256 penalty, uint256 reward)
func (_DowntimeSlasher *DowntimeSlasherFilterer) FilterSlashingIncentivesSet(opts *bind.FilterOpts) (*DowntimeSlasherSlashingIncentivesSetIterator, error) {

	logs, sub, err := _DowntimeSlasher.contract.FilterLogs(opts, "SlashingIncentivesSet")
	if err != nil {
		return nil, err
	}
	return &DowntimeSlasherSlashingIncentivesSetIterator{contract: _DowntimeSlasher.contract, event: "SlashingIncentivesSet", logs: logs, sub: sub}, nil
}

// WatchSlashingIncentivesSet is a free log subscription operation binding the contract event 0x716dc7c34384df36c6ccc5a2949f2ce9b019f5d4075ef39139a80038a4fdd1c3.
//
// Solidity: event SlashingIncentivesSet(uint256 penalty, uint256 reward)
func (_DowntimeSlasher *DowntimeSlasherFilterer) WatchSlashingIncentivesSet(opts *bind.WatchOpts, sink chan<- *DowntimeSlasherSlashingIncentivesSet) (event.Subscription, error) {

	logs, sub, err := _DowntimeSlasher.contract.WatchLogs(opts, "SlashingIncentivesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DowntimeSlasherSlashingIncentivesSet)
				if err := _DowntimeSlasher.contract.UnpackLog(event, "SlashingIncentivesSet", log); err != nil {
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

// ParseSlashingIncentivesSet is a log parse operation binding the contract event 0x716dc7c34384df36c6ccc5a2949f2ce9b019f5d4075ef39139a80038a4fdd1c3.
//
// Solidity: event SlashingIncentivesSet(uint256 penalty, uint256 reward)
func (_DowntimeSlasher *DowntimeSlasherFilterer) ParseSlashingIncentivesSet(log types.Log) (*DowntimeSlasherSlashingIncentivesSet, error) {
	event := new(DowntimeSlasherSlashingIncentivesSet)
	if err := _DowntimeSlasher.contract.UnpackLog(event, "SlashingIncentivesSet", log); err != nil {
		return nil, err
	}
	return event, nil
}
