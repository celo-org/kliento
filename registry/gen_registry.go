// Code generated - DO NOT EDIT.

package registry

import (
	"context"
	"math/big"

	"github.com/celo-org/kliento/client"
	"github.com/celo-org/kliento/contracts"
	"github.com/celo-org/kliento/contracts/helpers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// AccountsContractID is the registry identifier for 'Accounts' contract
var AccountsContractID ContractID = "Accounts"

// AttestationsContractID is the registry identifier for 'Attestations' contract
var AttestationsContractID ContractID = "Attestations"

// BlockchainParametersContractID is the registry identifier for 'BlockchainParameters' contract
var BlockchainParametersContractID ContractID = "BlockchainParameters"

// DoubleSigningSlasherContractID is the registry identifier for 'DoubleSigningSlasher' contract
var DoubleSigningSlasherContractID ContractID = "DoubleSigningSlasher"

// DowntimeSlasherContractID is the registry identifier for 'DowntimeSlasher' contract
var DowntimeSlasherContractID ContractID = "DowntimeSlasher"

// ElectionContractID is the registry identifier for 'Election' contract
var ElectionContractID ContractID = "Election"

// EpochRewardsContractID is the registry identifier for 'EpochRewards' contract
var EpochRewardsContractID ContractID = "EpochRewards"

// EscrowContractID is the registry identifier for 'Escrow' contract
var EscrowContractID ContractID = "Escrow"

// ExchangeContractID is the registry identifier for 'Exchange' contract
var ExchangeContractID ContractID = "Exchange"

// GasPriceMinimumContractID is the registry identifier for 'GasPriceMinimum' contract
var GasPriceMinimumContractID ContractID = "GasPriceMinimum"

// GoldTokenContractID is the registry identifier for 'GoldToken' contract
var GoldTokenContractID ContractID = "GoldToken"

// GovernanceContractID is the registry identifier for 'Governance' contract
var GovernanceContractID ContractID = "Governance"

// GovernanceSlasherContractID is the registry identifier for 'GovernanceSlasher' contract
var GovernanceSlasherContractID ContractID = "GovernanceSlasher"

// LockedGoldContractID is the registry identifier for 'LockedGold' contract
var LockedGoldContractID ContractID = "LockedGold"

// RandomContractID is the registry identifier for 'Random' contract
var RandomContractID ContractID = "Random"

// ReserveContractID is the registry identifier for 'Reserve' contract
var ReserveContractID ContractID = "Reserve"

// SortedOraclesContractID is the registry identifier for 'SortedOracles' contract
var SortedOraclesContractID ContractID = "SortedOracles"

// StableTokenContractID is the registry identifier for 'StableToken' contract
var StableTokenContractID ContractID = "StableToken"

// ValidatorsContractID is the registry identifier for 'Validators' contract
var ValidatorsContractID ContractID = "Validators"

// RegisteredContractIDs are all (known) registered contract identifiers
var RegisteredContractIDs = []ContractID{

	AccountsContractID,

	AttestationsContractID,

	BlockchainParametersContractID,

	DoubleSigningSlasherContractID,

	DowntimeSlasherContractID,

	ElectionContractID,

	EpochRewardsContractID,

	EscrowContractID,

	ExchangeContractID,

	GasPriceMinimumContractID,

	GoldTokenContractID,

	GovernanceContractID,

	GovernanceSlasherContractID,

	LockedGoldContractID,

	RandomContractID,

	ReserveContractID,

	SortedOraclesContractID,

	StableTokenContractID,

	ValidatorsContractID,
}

type boundRegistry struct {
	_AccountsAddress *common.Address
	AccountsContract *contracts.Accounts

	_AttestationsAddress *common.Address
	AttestationsContract *contracts.Attestations

	_BlockchainParametersAddress *common.Address
	BlockchainParametersContract *contracts.BlockchainParameters

	_DoubleSigningSlasherAddress *common.Address
	DoubleSigningSlasherContract *contracts.DoubleSigningSlasher

	_DowntimeSlasherAddress *common.Address
	DowntimeSlasherContract *contracts.DowntimeSlasher

	_ElectionAddress *common.Address
	ElectionContract *contracts.Election

	_EpochRewardsAddress *common.Address
	EpochRewardsContract *contracts.EpochRewards

	_EscrowAddress *common.Address
	EscrowContract *contracts.Escrow

	_ExchangeAddress *common.Address
	ExchangeContract *contracts.Exchange

	_GasPriceMinimumAddress *common.Address
	GasPriceMinimumContract *contracts.GasPriceMinimum

	_GoldTokenAddress *common.Address
	GoldTokenContract *contracts.GoldToken

	_GovernanceAddress *common.Address
	GovernanceContract *contracts.Governance

	_GovernanceSlasherAddress *common.Address
	GovernanceSlasherContract *contracts.GovernanceSlasher

	_LockedGoldAddress *common.Address
	LockedGoldContract *contracts.LockedGold

	_RandomAddress *common.Address
	RandomContract *contracts.Random

	_ReserveAddress *common.Address
	ReserveContract *contracts.Reserve

	_SortedOraclesAddress *common.Address
	SortedOraclesContract *contracts.SortedOracles

	_StableTokenAddress *common.Address
	StableTokenContract *contracts.StableToken

	_ValidatorsAddress *common.Address
	ValidatorsContract *contracts.Validators
}

type generatedRegistry interface {
	GetAccountsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Accounts, error)

	GetAttestationsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Attestations, error)

	GetBlockchainParametersContract(ctx context.Context, blockNumber *big.Int) (*contracts.BlockchainParameters, error)

	GetDoubleSigningSlasherContract(ctx context.Context, blockNumber *big.Int) (*contracts.DoubleSigningSlasher, error)

	GetDowntimeSlasherContract(ctx context.Context, blockNumber *big.Int) (*contracts.DowntimeSlasher, error)

	GetElectionContract(ctx context.Context, blockNumber *big.Int) (*contracts.Election, error)

	GetEpochRewardsContract(ctx context.Context, blockNumber *big.Int) (*contracts.EpochRewards, error)

	GetEscrowContract(ctx context.Context, blockNumber *big.Int) (*contracts.Escrow, error)

	GetExchangeContract(ctx context.Context, blockNumber *big.Int) (*contracts.Exchange, error)

	GetGasPriceMinimumContract(ctx context.Context, blockNumber *big.Int) (*contracts.GasPriceMinimum, error)

	GetGoldTokenContract(ctx context.Context, blockNumber *big.Int) (*contracts.GoldToken, error)

	GetGovernanceContract(ctx context.Context, blockNumber *big.Int) (*contracts.Governance, error)

	GetGovernanceSlasherContract(ctx context.Context, blockNumber *big.Int) (*contracts.GovernanceSlasher, error)

	GetLockedGoldContract(ctx context.Context, blockNumber *big.Int) (*contracts.LockedGold, error)

	GetRandomContract(ctx context.Context, blockNumber *big.Int) (*contracts.Random, error)

	GetReserveContract(ctx context.Context, blockNumber *big.Int) (*contracts.Reserve, error)

	GetSortedOraclesContract(ctx context.Context, blockNumber *big.Int) (*contracts.SortedOracles, error)

	GetStableTokenContract(ctx context.Context, blockNumber *big.Int) (*contracts.StableToken, error)

	GetValidatorsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Validators, error)
}

func (r *registryImpl) GetAccountsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Accounts, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, AccountsContractID)
	if err != nil {
		return nil, err
	} else if address == *r._AccountsAddress {
		return r.AccountsContract, nil
	}

	contract, err := contracts.NewAccounts(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._AccountsAddress, r.AccountsContract = &address, contract
	return contract, nil
}

func (r *registryImpl) GetAttestationsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Attestations, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, AttestationsContractID)
	if err != nil {
		return nil, err
	} else if address == *r._AttestationsAddress {
		return r.AttestationsContract, nil
	}

	contract, err := contracts.NewAttestations(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._AttestationsAddress, r.AttestationsContract = &address, contract
	return contract, nil
}

func (r *registryImpl) GetBlockchainParametersContract(ctx context.Context, blockNumber *big.Int) (*contracts.BlockchainParameters, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, BlockchainParametersContractID)
	if err != nil {
		return nil, err
	} else if address == *r._BlockchainParametersAddress {
		return r.BlockchainParametersContract, nil
	}

	contract, err := contracts.NewBlockchainParameters(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._BlockchainParametersAddress, r.BlockchainParametersContract = &address, contract
	return contract, nil
}

func (r *registryImpl) GetDoubleSigningSlasherContract(ctx context.Context, blockNumber *big.Int) (*contracts.DoubleSigningSlasher, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, DoubleSigningSlasherContractID)
	if err != nil {
		return nil, err
	} else if address == *r._DoubleSigningSlasherAddress {
		return r.DoubleSigningSlasherContract, nil
	}

	contract, err := contracts.NewDoubleSigningSlasher(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._DoubleSigningSlasherAddress, r.DoubleSigningSlasherContract = &address, contract
	return contract, nil
}

func (r *registryImpl) GetDowntimeSlasherContract(ctx context.Context, blockNumber *big.Int) (*contracts.DowntimeSlasher, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, DowntimeSlasherContractID)
	if err != nil {
		return nil, err
	} else if address == *r._DowntimeSlasherAddress {
		return r.DowntimeSlasherContract, nil
	}

	contract, err := contracts.NewDowntimeSlasher(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._DowntimeSlasherAddress, r.DowntimeSlasherContract = &address, contract
	return contract, nil
}

func (r *registryImpl) GetElectionContract(ctx context.Context, blockNumber *big.Int) (*contracts.Election, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, ElectionContractID)
	if err != nil {
		return nil, err
	} else if address == *r._ElectionAddress {
		return r.ElectionContract, nil
	}

	contract, err := contracts.NewElection(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._ElectionAddress, r.ElectionContract = &address, contract
	return contract, nil
}

func (r *registryImpl) GetEpochRewardsContract(ctx context.Context, blockNumber *big.Int) (*contracts.EpochRewards, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, EpochRewardsContractID)
	if err != nil {
		return nil, err
	} else if address == *r._EpochRewardsAddress {
		return r.EpochRewardsContract, nil
	}

	contract, err := contracts.NewEpochRewards(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._EpochRewardsAddress, r.EpochRewardsContract = &address, contract
	return contract, nil
}

func (r *registryImpl) GetEscrowContract(ctx context.Context, blockNumber *big.Int) (*contracts.Escrow, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, EscrowContractID)
	if err != nil {
		return nil, err
	} else if address == *r._EscrowAddress {
		return r.EscrowContract, nil
	}

	contract, err := contracts.NewEscrow(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._EscrowAddress, r.EscrowContract = &address, contract
	return contract, nil
}

func (r *registryImpl) GetExchangeContract(ctx context.Context, blockNumber *big.Int) (*contracts.Exchange, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, ExchangeContractID)
	if err != nil {
		return nil, err
	} else if address == *r._ExchangeAddress {
		return r.ExchangeContract, nil
	}

	contract, err := contracts.NewExchange(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._ExchangeAddress, r.ExchangeContract = &address, contract
	return contract, nil
}

func (r *registryImpl) GetGasPriceMinimumContract(ctx context.Context, blockNumber *big.Int) (*contracts.GasPriceMinimum, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, GasPriceMinimumContractID)
	if err != nil {
		return nil, err
	} else if address == *r._GasPriceMinimumAddress {
		return r.GasPriceMinimumContract, nil
	}

	contract, err := contracts.NewGasPriceMinimum(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._GasPriceMinimumAddress, r.GasPriceMinimumContract = &address, contract
	return contract, nil
}

func (r *registryImpl) GetGoldTokenContract(ctx context.Context, blockNumber *big.Int) (*contracts.GoldToken, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, GoldTokenContractID)
	if err != nil {
		return nil, err
	} else if address == *r._GoldTokenAddress {
		return r.GoldTokenContract, nil
	}

	contract, err := contracts.NewGoldToken(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._GoldTokenAddress, r.GoldTokenContract = &address, contract
	return contract, nil
}

func (r *registryImpl) GetGovernanceContract(ctx context.Context, blockNumber *big.Int) (*contracts.Governance, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, GovernanceContractID)
	if err != nil {
		return nil, err
	} else if address == *r._GovernanceAddress {
		return r.GovernanceContract, nil
	}

	contract, err := contracts.NewGovernance(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._GovernanceAddress, r.GovernanceContract = &address, contract
	return contract, nil
}

func (r *registryImpl) GetGovernanceSlasherContract(ctx context.Context, blockNumber *big.Int) (*contracts.GovernanceSlasher, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, GovernanceSlasherContractID)
	if err != nil {
		return nil, err
	} else if address == *r._GovernanceSlasherAddress {
		return r.GovernanceSlasherContract, nil
	}

	contract, err := contracts.NewGovernanceSlasher(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._GovernanceSlasherAddress, r.GovernanceSlasherContract = &address, contract
	return contract, nil
}

func (r *registryImpl) GetLockedGoldContract(ctx context.Context, blockNumber *big.Int) (*contracts.LockedGold, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, LockedGoldContractID)
	if err != nil {
		return nil, err
	} else if address == *r._LockedGoldAddress {
		return r.LockedGoldContract, nil
	}

	contract, err := contracts.NewLockedGold(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._LockedGoldAddress, r.LockedGoldContract = &address, contract
	return contract, nil
}

func (r *registryImpl) GetRandomContract(ctx context.Context, blockNumber *big.Int) (*contracts.Random, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, RandomContractID)
	if err != nil {
		return nil, err
	} else if address == *r._RandomAddress {
		return r.RandomContract, nil
	}

	contract, err := contracts.NewRandom(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._RandomAddress, r.RandomContract = &address, contract
	return contract, nil
}

func (r *registryImpl) GetReserveContract(ctx context.Context, blockNumber *big.Int) (*contracts.Reserve, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, ReserveContractID)
	if err != nil {
		return nil, err
	} else if address == *r._ReserveAddress {
		return r.ReserveContract, nil
	}

	contract, err := contracts.NewReserve(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._ReserveAddress, r.ReserveContract = &address, contract
	return contract, nil
}

func (r *registryImpl) GetSortedOraclesContract(ctx context.Context, blockNumber *big.Int) (*contracts.SortedOracles, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, SortedOraclesContractID)
	if err != nil {
		return nil, err
	} else if address == *r._SortedOraclesAddress {
		return r.SortedOraclesContract, nil
	}

	contract, err := contracts.NewSortedOracles(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._SortedOraclesAddress, r.SortedOraclesContract = &address, contract
	return contract, nil
}

func (r *registryImpl) GetStableTokenContract(ctx context.Context, blockNumber *big.Int) (*contracts.StableToken, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, StableTokenContractID)
	if err != nil {
		return nil, err
	} else if address == *r._StableTokenAddress {
		return r.StableTokenContract, nil
	}

	contract, err := contracts.NewStableToken(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._StableTokenAddress, r.StableTokenContract = &address, contract
	return contract, nil
}

func (r *registryImpl) GetValidatorsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Validators, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, ValidatorsContractID)
	if err != nil {
		return nil, err
	} else if address == *r._ValidatorsAddress {
		return r.ValidatorsContract, nil
	}

	contract, err := contracts.NewValidators(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._ValidatorsAddress, r.ValidatorsContract = &address, contract
	return contract, nil
}

// Hydrate populates contract bindings using the registry's mapping at the provided blockNumber
// for all known and deployed contracts
func (r *registryImpl) Hydrate(ctx context.Context, blockNumber *big.Int) error {
	check := func(err error) bool {
		return err != nil && err != ErrRegistryNotDeployed && err != client.ErrContractNotDeployed
	}

	var err error

	_, err = r.GetAccountsContract(ctx, blockNumber)
	if check(err) {
		return err
	}

	_, err = r.GetAttestationsContract(ctx, blockNumber)
	if check(err) {
		return err
	}

	_, err = r.GetBlockchainParametersContract(ctx, blockNumber)
	if check(err) {
		return err
	}

	_, err = r.GetDoubleSigningSlasherContract(ctx, blockNumber)
	if check(err) {
		return err
	}

	_, err = r.GetDowntimeSlasherContract(ctx, blockNumber)
	if check(err) {
		return err
	}

	_, err = r.GetElectionContract(ctx, blockNumber)
	if check(err) {
		return err
	}

	_, err = r.GetEpochRewardsContract(ctx, blockNumber)
	if check(err) {
		return err
	}

	_, err = r.GetEscrowContract(ctx, blockNumber)
	if check(err) {
		return err
	}

	_, err = r.GetExchangeContract(ctx, blockNumber)
	if check(err) {
		return err
	}

	_, err = r.GetGasPriceMinimumContract(ctx, blockNumber)
	if check(err) {
		return err
	}

	_, err = r.GetGoldTokenContract(ctx, blockNumber)
	if check(err) {
		return err
	}

	_, err = r.GetGovernanceContract(ctx, blockNumber)
	if check(err) {
		return err
	}

	_, err = r.GetGovernanceSlasherContract(ctx, blockNumber)
	if check(err) {
		return err
	}

	_, err = r.GetLockedGoldContract(ctx, blockNumber)
	if check(err) {
		return err
	}

	_, err = r.GetRandomContract(ctx, blockNumber)
	if check(err) {
		return err
	}

	_, err = r.GetReserveContract(ctx, blockNumber)
	if check(err) {
		return err
	}

	_, err = r.GetSortedOraclesContract(ctx, blockNumber)
	if check(err) {
		return err
	}

	_, err = r.GetStableTokenContract(ctx, blockNumber)
	if check(err) {
		return err
	}

	_, err = r.GetValidatorsContract(ctx, blockNumber)
	if check(err) {
		return err
	}

	return nil
}

// ParseLog parses an event log using a "hydrated" registry
// Hydrate should be called at the desired block number prior to parsing for comprehensive results
func (r *registryImpl) ParseLog(eventLog types.Log) []interface{} {
	buildSlice := func(contractName string, eventName string, event interface{}) []interface{} {
		slice := []interface{}{"contract", contractName, "event", eventName}
		eventSlice, _ := helpers.EventToSlice(event)
		return append(slice, eventSlice)
	}

	if r.AccountsContract != nil {
		eventName, event, ok, err := r.AccountsContract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("Accounts", eventName, event)
		}
	}

	if r.AttestationsContract != nil {
		eventName, event, ok, err := r.AttestationsContract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("Attestations", eventName, event)
		}
	}

	if r.BlockchainParametersContract != nil {
		eventName, event, ok, err := r.BlockchainParametersContract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("BlockchainParameters", eventName, event)
		}
	}

	if r.DoubleSigningSlasherContract != nil {
		eventName, event, ok, err := r.DoubleSigningSlasherContract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("DoubleSigningSlasher", eventName, event)
		}
	}

	if r.DowntimeSlasherContract != nil {
		eventName, event, ok, err := r.DowntimeSlasherContract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("DowntimeSlasher", eventName, event)
		}
	}

	if r.ElectionContract != nil {
		eventName, event, ok, err := r.ElectionContract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("Election", eventName, event)
		}
	}

	if r.EpochRewardsContract != nil {
		eventName, event, ok, err := r.EpochRewardsContract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("EpochRewards", eventName, event)
		}
	}

	if r.EscrowContract != nil {
		eventName, event, ok, err := r.EscrowContract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("Escrow", eventName, event)
		}
	}

	if r.ExchangeContract != nil {
		eventName, event, ok, err := r.ExchangeContract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("Exchange", eventName, event)
		}
	}

	if r.GasPriceMinimumContract != nil {
		eventName, event, ok, err := r.GasPriceMinimumContract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("GasPriceMinimum", eventName, event)
		}
	}

	if r.GoldTokenContract != nil {
		eventName, event, ok, err := r.GoldTokenContract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("GoldToken", eventName, event)
		}
	}

	if r.GovernanceContract != nil {
		eventName, event, ok, err := r.GovernanceContract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("Governance", eventName, event)
		}
	}

	if r.GovernanceSlasherContract != nil {
		eventName, event, ok, err := r.GovernanceSlasherContract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("GovernanceSlasher", eventName, event)
		}
	}

	if r.LockedGoldContract != nil {
		eventName, event, ok, err := r.LockedGoldContract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("LockedGold", eventName, event)
		}
	}

	if r.RandomContract != nil {
		eventName, event, ok, err := r.RandomContract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("Random", eventName, event)
		}
	}

	if r.ReserveContract != nil {
		eventName, event, ok, err := r.ReserveContract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("Reserve", eventName, event)
		}
	}

	if r.SortedOraclesContract != nil {
		eventName, event, ok, err := r.SortedOraclesContract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("SortedOracles", eventName, event)
		}
	}

	if r.StableTokenContract != nil {
		eventName, event, ok, err := r.StableTokenContract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("StableToken", eventName, event)
		}
	}

	if r.ValidatorsContract != nil {
		eventName, event, ok, err := r.ValidatorsContract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("Validators", eventName, event)
		}
	}

	eventName, event, ok, err := r.RegistryContract.TryParseLog(eventLog)
	if ok && err != nil {
		return buildSlice("Registry", eventName, event)
	}
	return nil
}
