// Code generated - DO NOT EDIT.

package registry

import (
	"context"
	"fmt"
	"math/big"

	"github.com/celo-org/kliento/contracts"
	"github.com/celo-org/kliento/contracts/helpers"
	blockchainErrors "github.com/ethereum/go-ethereum/contract_comm/errors"
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
	AccountsContract      *contracts.Accounts
	AccountsContractProxy *contracts.Proxy

	AttestationsContract      *contracts.Attestations
	AttestationsContractProxy *contracts.Proxy

	BlockchainParametersContract      *contracts.BlockchainParameters
	BlockchainParametersContractProxy *contracts.Proxy

	DoubleSigningSlasherContract      *contracts.DoubleSigningSlasher
	DoubleSigningSlasherContractProxy *contracts.Proxy

	DowntimeSlasherContract      *contracts.DowntimeSlasher
	DowntimeSlasherContractProxy *contracts.Proxy

	ElectionContract      *contracts.Election
	ElectionContractProxy *contracts.Proxy

	EpochRewardsContract      *contracts.EpochRewards
	EpochRewardsContractProxy *contracts.Proxy

	EscrowContract      *contracts.Escrow
	EscrowContractProxy *contracts.Proxy

	ExchangeContract      *contracts.Exchange
	ExchangeContractProxy *contracts.Proxy

	GasPriceMinimumContract      *contracts.GasPriceMinimum
	GasPriceMinimumContractProxy *contracts.Proxy

	GoldTokenContract      *contracts.GoldToken
	GoldTokenContractProxy *contracts.Proxy

	GovernanceContract      *contracts.Governance
	GovernanceContractProxy *contracts.Proxy

	GovernanceSlasherContract      *contracts.GovernanceSlasher
	GovernanceSlasherContractProxy *contracts.Proxy

	LockedGoldContract      *contracts.LockedGold
	LockedGoldContractProxy *contracts.Proxy

	RandomContract      *contracts.Random
	RandomContractProxy *contracts.Proxy

	ReserveContract      *contracts.Reserve
	ReserveContractProxy *contracts.Proxy

	SortedOraclesContract      *contracts.SortedOracles
	SortedOraclesContractProxy *contracts.Proxy

	StableTokenContract      *contracts.StableToken
	StableTokenContractProxy *contracts.Proxy

	ValidatorsContract      *contracts.Validators
	ValidatorsContractProxy *contracts.Proxy
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
	identifier := AccountsContractID.String()
	if r.AccountsContract == nil || r.isCacheDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, AccountsContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewAccounts(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.AccountsContract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.AccountsContractProxy = contractProxy
	}
	return r.AccountsContract, nil
}

func (r *registryImpl) GetAttestationsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Attestations, error) {
	identifier := AttestationsContractID.String()
	if r.AttestationsContract == nil || r.isCacheDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, AttestationsContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewAttestations(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.AttestationsContract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.AttestationsContractProxy = contractProxy
	}
	return r.AttestationsContract, nil
}

func (r *registryImpl) GetBlockchainParametersContract(ctx context.Context, blockNumber *big.Int) (*contracts.BlockchainParameters, error) {
	identifier := BlockchainParametersContractID.String()
	if r.BlockchainParametersContract == nil || r.isCacheDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, BlockchainParametersContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewBlockchainParameters(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.BlockchainParametersContract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.BlockchainParametersContractProxy = contractProxy
	}
	return r.BlockchainParametersContract, nil
}

func (r *registryImpl) GetDoubleSigningSlasherContract(ctx context.Context, blockNumber *big.Int) (*contracts.DoubleSigningSlasher, error) {
	identifier := DoubleSigningSlasherContractID.String()
	if r.DoubleSigningSlasherContract == nil || r.isCacheDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, DoubleSigningSlasherContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewDoubleSigningSlasher(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.DoubleSigningSlasherContract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.DoubleSigningSlasherContractProxy = contractProxy
	}
	return r.DoubleSigningSlasherContract, nil
}

func (r *registryImpl) GetDowntimeSlasherContract(ctx context.Context, blockNumber *big.Int) (*contracts.DowntimeSlasher, error) {
	identifier := DowntimeSlasherContractID.String()
	if r.DowntimeSlasherContract == nil || r.isCacheDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, DowntimeSlasherContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewDowntimeSlasher(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.DowntimeSlasherContract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.DowntimeSlasherContractProxy = contractProxy
	}
	return r.DowntimeSlasherContract, nil
}

func (r *registryImpl) GetElectionContract(ctx context.Context, blockNumber *big.Int) (*contracts.Election, error) {
	identifier := ElectionContractID.String()
	if r.ElectionContract == nil || r.isCacheDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, ElectionContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewElection(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.ElectionContract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.ElectionContractProxy = contractProxy
	}
	return r.ElectionContract, nil
}

func (r *registryImpl) GetEpochRewardsContract(ctx context.Context, blockNumber *big.Int) (*contracts.EpochRewards, error) {
	identifier := EpochRewardsContractID.String()
	if r.EpochRewardsContract == nil || r.isCacheDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, EpochRewardsContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewEpochRewards(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.EpochRewardsContract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.EpochRewardsContractProxy = contractProxy
	}
	return r.EpochRewardsContract, nil
}

func (r *registryImpl) GetEscrowContract(ctx context.Context, blockNumber *big.Int) (*contracts.Escrow, error) {
	identifier := EscrowContractID.String()
	if r.EscrowContract == nil || r.isCacheDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, EscrowContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewEscrow(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.EscrowContract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.EscrowContractProxy = contractProxy
	}
	return r.EscrowContract, nil
}

func (r *registryImpl) GetExchangeContract(ctx context.Context, blockNumber *big.Int) (*contracts.Exchange, error) {
	identifier := ExchangeContractID.String()
	if r.ExchangeContract == nil || r.isCacheDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, ExchangeContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewExchange(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.ExchangeContract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.ExchangeContractProxy = contractProxy
	}
	return r.ExchangeContract, nil
}

func (r *registryImpl) GetGasPriceMinimumContract(ctx context.Context, blockNumber *big.Int) (*contracts.GasPriceMinimum, error) {
	identifier := GasPriceMinimumContractID.String()
	if r.GasPriceMinimumContract == nil || r.isCacheDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, GasPriceMinimumContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewGasPriceMinimum(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.GasPriceMinimumContract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.GasPriceMinimumContractProxy = contractProxy
	}
	return r.GasPriceMinimumContract, nil
}

func (r *registryImpl) GetGoldTokenContract(ctx context.Context, blockNumber *big.Int) (*contracts.GoldToken, error) {
	identifier := GoldTokenContractID.String()
	if r.GoldTokenContract == nil || r.isCacheDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, GoldTokenContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewGoldToken(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.GoldTokenContract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.GoldTokenContractProxy = contractProxy
	}
	return r.GoldTokenContract, nil
}

func (r *registryImpl) GetGovernanceContract(ctx context.Context, blockNumber *big.Int) (*contracts.Governance, error) {
	identifier := GovernanceContractID.String()
	if r.GovernanceContract == nil || r.isCacheDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, GovernanceContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewGovernance(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.GovernanceContract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.GovernanceContractProxy = contractProxy
	}
	return r.GovernanceContract, nil
}

func (r *registryImpl) GetGovernanceSlasherContract(ctx context.Context, blockNumber *big.Int) (*contracts.GovernanceSlasher, error) {
	identifier := GovernanceSlasherContractID.String()
	if r.GovernanceSlasherContract == nil || r.isCacheDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, GovernanceSlasherContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewGovernanceSlasher(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.GovernanceSlasherContract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.GovernanceSlasherContractProxy = contractProxy
	}
	return r.GovernanceSlasherContract, nil
}

func (r *registryImpl) GetLockedGoldContract(ctx context.Context, blockNumber *big.Int) (*contracts.LockedGold, error) {
	identifier := LockedGoldContractID.String()
	if r.LockedGoldContract == nil || r.isCacheDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, LockedGoldContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewLockedGold(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.LockedGoldContract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.LockedGoldContractProxy = contractProxy
	}
	return r.LockedGoldContract, nil
}

func (r *registryImpl) GetRandomContract(ctx context.Context, blockNumber *big.Int) (*contracts.Random, error) {
	identifier := RandomContractID.String()
	if r.RandomContract == nil || r.isCacheDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, RandomContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewRandom(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.RandomContract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.RandomContractProxy = contractProxy
	}
	return r.RandomContract, nil
}

func (r *registryImpl) GetReserveContract(ctx context.Context, blockNumber *big.Int) (*contracts.Reserve, error) {
	identifier := ReserveContractID.String()
	if r.ReserveContract == nil || r.isCacheDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, ReserveContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewReserve(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.ReserveContract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.ReserveContractProxy = contractProxy
	}
	return r.ReserveContract, nil
}

func (r *registryImpl) GetSortedOraclesContract(ctx context.Context, blockNumber *big.Int) (*contracts.SortedOracles, error) {
	identifier := SortedOraclesContractID.String()
	if r.SortedOraclesContract == nil || r.isCacheDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, SortedOraclesContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewSortedOracles(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.SortedOraclesContract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.SortedOraclesContractProxy = contractProxy
	}
	return r.SortedOraclesContract, nil
}

func (r *registryImpl) GetStableTokenContract(ctx context.Context, blockNumber *big.Int) (*contracts.StableToken, error) {
	identifier := StableTokenContractID.String()
	if r.StableTokenContract == nil || r.isCacheDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, StableTokenContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewStableToken(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.StableTokenContract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.StableTokenContractProxy = contractProxy
	}
	return r.StableTokenContract, nil
}

func (r *registryImpl) GetValidatorsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Validators, error) {
	identifier := ValidatorsContractID.String()
	if r.ValidatorsContract == nil || r.isCacheDirty(identifier) {
		address, err := r.GetAddressFor(ctx, blockNumber, ValidatorsContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.NewValidators(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.ValidatorsContract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.ValidatorsContractProxy = contractProxy
	}
	return r.ValidatorsContract, nil
}

func (r *registryImpl) tryParseLogGenerated(ctx context.Context, eventLog *types.Log, blockNumber *big.Int) ([]interface{}, error) {
	var eventName string
	var event interface{}
	var ok bool
	var err error

	_, err = r.GetAccountsContract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.AccountsContract.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("Accounts", eventName, event)
		}
		eventName, event, ok, err = r.AccountsContractProxy.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("AccountsProxy", eventName, event)
		}
	} else {
		// skip deployed failures
		if err != blockchainErrors.ErrSmartContractNotDeployed {
			return nil, fmt.Errorf("Accounts %v", err)
		}
	}

	_, err = r.GetAttestationsContract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.AttestationsContract.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("Attestations", eventName, event)
		}
		eventName, event, ok, err = r.AttestationsContractProxy.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("AttestationsProxy", eventName, event)
		}
	} else {
		// skip deployed failures
		if err != blockchainErrors.ErrSmartContractNotDeployed {
			return nil, fmt.Errorf("Attestations %v", err)
		}
	}

	_, err = r.GetBlockchainParametersContract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.BlockchainParametersContract.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("BlockchainParameters", eventName, event)
		}
		eventName, event, ok, err = r.BlockchainParametersContractProxy.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("BlockchainParametersProxy", eventName, event)
		}
	} else {
		// skip deployed failures
		if err != blockchainErrors.ErrSmartContractNotDeployed {
			return nil, fmt.Errorf("BlockchainParameters %v", err)
		}
	}

	_, err = r.GetDoubleSigningSlasherContract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.DoubleSigningSlasherContract.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("DoubleSigningSlasher", eventName, event)
		}
		eventName, event, ok, err = r.DoubleSigningSlasherContractProxy.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("DoubleSigningSlasherProxy", eventName, event)
		}
	} else {
		// skip deployed failures
		if err != blockchainErrors.ErrSmartContractNotDeployed {
			return nil, fmt.Errorf("DoubleSigningSlasher %v", err)
		}
	}

	_, err = r.GetDowntimeSlasherContract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.DowntimeSlasherContract.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("DowntimeSlasher", eventName, event)
		}
		eventName, event, ok, err = r.DowntimeSlasherContractProxy.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("DowntimeSlasherProxy", eventName, event)
		}
	} else {
		// skip deployed failures
		if err != blockchainErrors.ErrSmartContractNotDeployed {
			return nil, fmt.Errorf("DowntimeSlasher %v", err)
		}
	}

	_, err = r.GetElectionContract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.ElectionContract.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("Election", eventName, event)
		}
		eventName, event, ok, err = r.ElectionContractProxy.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("ElectionProxy", eventName, event)
		}
	} else {
		// skip deployed failures
		if err != blockchainErrors.ErrSmartContractNotDeployed {
			return nil, fmt.Errorf("Election %v", err)
		}
	}

	_, err = r.GetEpochRewardsContract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.EpochRewardsContract.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("EpochRewards", eventName, event)
		}
		eventName, event, ok, err = r.EpochRewardsContractProxy.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("EpochRewardsProxy", eventName, event)
		}
	} else {
		// skip deployed failures
		if err != blockchainErrors.ErrSmartContractNotDeployed {
			return nil, fmt.Errorf("EpochRewards %v", err)
		}
	}

	_, err = r.GetEscrowContract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.EscrowContract.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("Escrow", eventName, event)
		}
		eventName, event, ok, err = r.EscrowContractProxy.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("EscrowProxy", eventName, event)
		}
	} else {
		// skip deployed failures
		if err != blockchainErrors.ErrSmartContractNotDeployed {
			return nil, fmt.Errorf("Escrow %v", err)
		}
	}

	_, err = r.GetExchangeContract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.ExchangeContract.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("Exchange", eventName, event)
		}
		eventName, event, ok, err = r.ExchangeContractProxy.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("ExchangeProxy", eventName, event)
		}
	} else {
		// skip deployed failures
		if err != blockchainErrors.ErrSmartContractNotDeployed {
			return nil, fmt.Errorf("Exchange %v", err)
		}
	}

	_, err = r.GetGasPriceMinimumContract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.GasPriceMinimumContract.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("GasPriceMinimum", eventName, event)
		}
		eventName, event, ok, err = r.GasPriceMinimumContractProxy.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("GasPriceMinimumProxy", eventName, event)
		}
	} else {
		// skip deployed failures
		if err != blockchainErrors.ErrSmartContractNotDeployed {
			return nil, fmt.Errorf("GasPriceMinimum %v", err)
		}
	}

	_, err = r.GetGoldTokenContract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.GoldTokenContract.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("GoldToken", eventName, event)
		}
		eventName, event, ok, err = r.GoldTokenContractProxy.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("GoldTokenProxy", eventName, event)
		}
	} else {
		// skip deployed failures
		if err != blockchainErrors.ErrSmartContractNotDeployed {
			return nil, fmt.Errorf("GoldToken %v", err)
		}
	}

	_, err = r.GetGovernanceContract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.GovernanceContract.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("Governance", eventName, event)
		}
		eventName, event, ok, err = r.GovernanceContractProxy.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("GovernanceProxy", eventName, event)
		}
	} else {
		// skip deployed failures
		if err != blockchainErrors.ErrSmartContractNotDeployed {
			return nil, fmt.Errorf("Governance %v", err)
		}
	}

	_, err = r.GetGovernanceSlasherContract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.GovernanceSlasherContract.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("GovernanceSlasher", eventName, event)
		}
		eventName, event, ok, err = r.GovernanceSlasherContractProxy.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("GovernanceSlasherProxy", eventName, event)
		}
	} else {
		// skip deployed failures
		if err != blockchainErrors.ErrSmartContractNotDeployed {
			return nil, fmt.Errorf("GovernanceSlasher %v", err)
		}
	}

	_, err = r.GetLockedGoldContract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.LockedGoldContract.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("LockedGold", eventName, event)
		}
		eventName, event, ok, err = r.LockedGoldContractProxy.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("LockedGoldProxy", eventName, event)
		}
	} else {
		// skip deployed failures
		if err != blockchainErrors.ErrSmartContractNotDeployed {
			return nil, fmt.Errorf("LockedGold %v", err)
		}
	}

	_, err = r.GetRandomContract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.RandomContract.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("Random", eventName, event)
		}
		eventName, event, ok, err = r.RandomContractProxy.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("RandomProxy", eventName, event)
		}
	} else {
		// skip deployed failures
		if err != blockchainErrors.ErrSmartContractNotDeployed {
			return nil, fmt.Errorf("Random %v", err)
		}
	}

	_, err = r.GetReserveContract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.ReserveContract.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("Reserve", eventName, event)
		}
		eventName, event, ok, err = r.ReserveContractProxy.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("ReserveProxy", eventName, event)
		}
	} else {
		// skip deployed failures
		if err != blockchainErrors.ErrSmartContractNotDeployed {
			return nil, fmt.Errorf("Reserve %v", err)
		}
	}

	_, err = r.GetSortedOraclesContract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.SortedOraclesContract.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("SortedOracles", eventName, event)
		}
		eventName, event, ok, err = r.SortedOraclesContractProxy.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("SortedOraclesProxy", eventName, event)
		}
	} else {
		// skip deployed failures
		if err != blockchainErrors.ErrSmartContractNotDeployed {
			return nil, fmt.Errorf("SortedOracles %v", err)
		}
	}

	_, err = r.GetStableTokenContract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.StableTokenContract.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("StableToken", eventName, event)
		}
		eventName, event, ok, err = r.StableTokenContractProxy.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("StableTokenProxy", eventName, event)
		}
	} else {
		// skip deployed failures
		if err != blockchainErrors.ErrSmartContractNotDeployed {
			return nil, fmt.Errorf("StableToken %v", err)
		}
	}

	_, err = r.GetValidatorsContract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.ValidatorsContract.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("Validators", eventName, event)
		}
		eventName, event, ok, err = r.ValidatorsContractProxy.TryParseLog(*eventLog) // checks matching address
		if ok && err == nil {
			return helpers.BuildEventSlice("ValidatorsProxy", eventName, event)
		}
	} else {
		// skip deployed failures
		if err != blockchainErrors.ErrSmartContractNotDeployed {
			return nil, fmt.Errorf("Validators %v", err)
		}
	}

	return nil, nil
}
