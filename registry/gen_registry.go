package registry

import (
	"context"
	"math/big"

	"github.com/celo-org/kliento/contracts"
)

// AccountsContractID is the registry Id for 'Accounts' contract
var AccountsContractID ContractID = "Accounts"

// AttestationsContractID is the registry Id for 'Attestations' contract
var AttestationsContractID ContractID = "Attestations"

// BlockchainParametersContractID is the registry Id for 'BlockchainParameters' contract
var BlockchainParametersContractID ContractID = "BlockchainParameters"

// DoubleSigningSlasherContractID is the registry Id for 'DoubleSigningSlasher' contract
var DoubleSigningSlasherContractID ContractID = "DoubleSigningSlasher"

// DowntimeSlasherContractID is the registry Id for 'DowntimeSlasher' contract
var DowntimeSlasherContractID ContractID = "DowntimeSlasher"

// ElectionContractID is the registry Id for 'Election' contract
var ElectionContractID ContractID = "Election"

// EpochRewardsContractID is the registry Id for 'EpochRewards' contract
var EpochRewardsContractID ContractID = "EpochRewards"

// EscrowContractID is the registry Id for 'Escrow' contract
var EscrowContractID ContractID = "Escrow"

// ExchangeContractID is the registry Id for 'Exchange' contract
var ExchangeContractID ContractID = "Exchange"

// FeeCurrencyWhitelistContractID is the registry Id for 'FeeCurrencyWhitelist' contract
var FeeCurrencyWhitelistContractID ContractID = "FeeCurrencyWhitelist"

// GasPriceMinimumContractID is the registry Id for 'GasPriceMinimum' contract
var GasPriceMinimumContractID ContractID = "GasPriceMinimum"

// GoldTokenContractID is the registry Id for 'GoldToken' contract
var GoldTokenContractID ContractID = "GoldToken"

// GovernanceContractID is the registry Id for 'Governance' contract
var GovernanceContractID ContractID = "Governance"

// GovernanceApproverMultiSigContractID is the registry Id for 'GovernanceApproverMultiSig' contract
var GovernanceApproverMultiSigContractID ContractID = "GovernanceApproverMultiSig"

// LockedGoldContractID is the registry Id for 'LockedGold' contract
var LockedGoldContractID ContractID = "LockedGold"

// MultiSigContractID is the registry Id for 'MultiSig' contract
var MultiSigContractID ContractID = "MultiSig"

// ProxyContractID is the registry Id for 'Proxy' contract
var ProxyContractID ContractID = "Proxy"

// RandomContractID is the registry Id for 'Random' contract
var RandomContractID ContractID = "Random"

// RegistryContractID is the registry Id for 'Registry' contract
var RegistryContractID ContractID = "Registry"

// ReleaseGoldContractID is the registry Id for 'ReleaseGold' contract
var ReleaseGoldContractID ContractID = "ReleaseGold"

// ReserveContractID is the registry Id for 'Reserve' contract
var ReserveContractID ContractID = "Reserve"

// ReserveSpenderMultiSigContractID is the registry Id for 'ReserveSpenderMultiSig' contract
var ReserveSpenderMultiSigContractID ContractID = "ReserveSpenderMultiSig"

// SortedOraclesContractID is the registry Id for 'SortedOracles' contract
var SortedOraclesContractID ContractID = "SortedOracles"

// StableTokenContractID is the registry Id for 'StableToken' contract
var StableTokenContractID ContractID = "StableToken"

// ValidatorsContractID is the registry Id for 'Validators' contract
var ValidatorsContractID ContractID = "Validators"

type GeneratedRegistry interface {
	GetAccountsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Accounts, error)

	GetAttestationsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Attestations, error)

	GetBlockchainParametersContract(ctx context.Context, blockNumber *big.Int) (*contracts.BlockchainParameters, error)

	GetDoubleSigningSlasherContract(ctx context.Context, blockNumber *big.Int) (*contracts.DoubleSigningSlasher, error)

	GetDowntimeSlasherContract(ctx context.Context, blockNumber *big.Int) (*contracts.DowntimeSlasher, error)

	GetElectionContract(ctx context.Context, blockNumber *big.Int) (*contracts.Election, error)

	GetEpochRewardsContract(ctx context.Context, blockNumber *big.Int) (*contracts.EpochRewards, error)

	GetEscrowContract(ctx context.Context, blockNumber *big.Int) (*contracts.Escrow, error)

	GetExchangeContract(ctx context.Context, blockNumber *big.Int) (*contracts.Exchange, error)

	GetFeeCurrencyWhitelistContract(ctx context.Context, blockNumber *big.Int) (*contracts.FeeCurrencyWhitelist, error)

	GetGasPriceMinimumContract(ctx context.Context, blockNumber *big.Int) (*contracts.GasPriceMinimum, error)

	GetGoldTokenContract(ctx context.Context, blockNumber *big.Int) (*contracts.GoldToken, error)

	GetGovernanceContract(ctx context.Context, blockNumber *big.Int) (*contracts.Governance, error)

	GetGovernanceApproverMultiSigContract(ctx context.Context, blockNumber *big.Int) (*contracts.GovernanceApproverMultiSig, error)

	GetLockedGoldContract(ctx context.Context, blockNumber *big.Int) (*contracts.LockedGold, error)

	GetMultiSigContract(ctx context.Context, blockNumber *big.Int) (*contracts.MultiSig, error)

	GetProxyContract(ctx context.Context, blockNumber *big.Int) (*contracts.Proxy, error)

	GetRandomContract(ctx context.Context, blockNumber *big.Int) (*contracts.Random, error)

	GetRegistryContract(ctx context.Context, blockNumber *big.Int) (*contracts.Registry, error)

	GetReleaseGoldContract(ctx context.Context, blockNumber *big.Int) (*contracts.ReleaseGold, error)

	GetReserveContract(ctx context.Context, blockNumber *big.Int) (*contracts.Reserve, error)

	GetReserveSpenderMultiSigContract(ctx context.Context, blockNumber *big.Int) (*contracts.ReserveSpenderMultiSig, error)

	GetSortedOraclesContract(ctx context.Context, blockNumber *big.Int) (*contracts.SortedOracles, error)

	GetStableTokenContract(ctx context.Context, blockNumber *big.Int) (*contracts.StableToken, error)

	GetValidatorsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Validators, error)
}

func (r *registryImpl) GetAccountsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Accounts, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, AccountsContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewAccounts(address, r.cc.Eth)
}

func (r *registryImpl) GetAttestationsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Attestations, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, AttestationsContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewAttestations(address, r.cc.Eth)
}

func (r *registryImpl) GetBlockchainParametersContract(ctx context.Context, blockNumber *big.Int) (*contracts.BlockchainParameters, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, BlockchainParametersContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewBlockchainParameters(address, r.cc.Eth)
}

func (r *registryImpl) GetDoubleSigningSlasherContract(ctx context.Context, blockNumber *big.Int) (*contracts.DoubleSigningSlasher, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, DoubleSigningSlasherContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewDoubleSigningSlasher(address, r.cc.Eth)
}

func (r *registryImpl) GetDowntimeSlasherContract(ctx context.Context, blockNumber *big.Int) (*contracts.DowntimeSlasher, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, DowntimeSlasherContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewDowntimeSlasher(address, r.cc.Eth)
}

func (r *registryImpl) GetElectionContract(ctx context.Context, blockNumber *big.Int) (*contracts.Election, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, ElectionContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewElection(address, r.cc.Eth)
}

func (r *registryImpl) GetEpochRewardsContract(ctx context.Context, blockNumber *big.Int) (*contracts.EpochRewards, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, EpochRewardsContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewEpochRewards(address, r.cc.Eth)
}

func (r *registryImpl) GetEscrowContract(ctx context.Context, blockNumber *big.Int) (*contracts.Escrow, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, EscrowContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewEscrow(address, r.cc.Eth)
}

func (r *registryImpl) GetExchangeContract(ctx context.Context, blockNumber *big.Int) (*contracts.Exchange, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, ExchangeContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewExchange(address, r.cc.Eth)
}

func (r *registryImpl) GetFeeCurrencyWhitelistContract(ctx context.Context, blockNumber *big.Int) (*contracts.FeeCurrencyWhitelist, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, FeeCurrencyWhitelistContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewFeeCurrencyWhitelist(address, r.cc.Eth)
}

func (r *registryImpl) GetGasPriceMinimumContract(ctx context.Context, blockNumber *big.Int) (*contracts.GasPriceMinimum, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, GasPriceMinimumContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewGasPriceMinimum(address, r.cc.Eth)
}

func (r *registryImpl) GetGoldTokenContract(ctx context.Context, blockNumber *big.Int) (*contracts.GoldToken, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, GoldTokenContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewGoldToken(address, r.cc.Eth)
}

func (r *registryImpl) GetGovernanceContract(ctx context.Context, blockNumber *big.Int) (*contracts.Governance, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, GovernanceContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewGovernance(address, r.cc.Eth)
}

func (r *registryImpl) GetGovernanceApproverMultiSigContract(ctx context.Context, blockNumber *big.Int) (*contracts.GovernanceApproverMultiSig, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, GovernanceApproverMultiSigContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewGovernanceApproverMultiSig(address, r.cc.Eth)
}

func (r *registryImpl) GetLockedGoldContract(ctx context.Context, blockNumber *big.Int) (*contracts.LockedGold, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, LockedGoldContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewLockedGold(address, r.cc.Eth)
}

func (r *registryImpl) GetMultiSigContract(ctx context.Context, blockNumber *big.Int) (*contracts.MultiSig, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, MultiSigContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewMultiSig(address, r.cc.Eth)
}

func (r *registryImpl) GetProxyContract(ctx context.Context, blockNumber *big.Int) (*contracts.Proxy, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, ProxyContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewProxy(address, r.cc.Eth)
}

func (r *registryImpl) GetRandomContract(ctx context.Context, blockNumber *big.Int) (*contracts.Random, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, RandomContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewRandom(address, r.cc.Eth)
}

func (r *registryImpl) GetRegistryContract(ctx context.Context, blockNumber *big.Int) (*contracts.Registry, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, RegistryContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewRegistry(address, r.cc.Eth)
}

func (r *registryImpl) GetReleaseGoldContract(ctx context.Context, blockNumber *big.Int) (*contracts.ReleaseGold, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, ReleaseGoldContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewReleaseGold(address, r.cc.Eth)
}

func (r *registryImpl) GetReserveContract(ctx context.Context, blockNumber *big.Int) (*contracts.Reserve, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, ReserveContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewReserve(address, r.cc.Eth)
}

func (r *registryImpl) GetReserveSpenderMultiSigContract(ctx context.Context, blockNumber *big.Int) (*contracts.ReserveSpenderMultiSig, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, ReserveSpenderMultiSigContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewReserveSpenderMultiSig(address, r.cc.Eth)
}

func (r *registryImpl) GetSortedOraclesContract(ctx context.Context, blockNumber *big.Int) (*contracts.SortedOracles, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, SortedOraclesContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewSortedOracles(address, r.cc.Eth)
}

func (r *registryImpl) GetStableTokenContract(ctx context.Context, blockNumber *big.Int) (*contracts.StableToken, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, StableTokenContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewStableToken(address, r.cc.Eth)
}

func (r *registryImpl) GetValidatorsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Validators, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, ValidatorsContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewValidators(address, r.cc.Eth)
}
