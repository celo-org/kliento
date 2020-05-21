// Copyright 2020 Celo Org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package registry

import (
	"context"
	"errors"
	"math/big"

	"github.com/celo-org/kliento/client"
	"github.com/celo-org/kliento/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

// RegistryAddress is the address of the registry which is the same on any celo network
var RegistryAddress = params.RegistrySmartContractAddress

// Registry defines an interface to access all Celo core Contracts
type Registry interface {
	GetAddressFor(ctx context.Context, blockNumber *big.Int, contractID ContractID) (common.Address, error)

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
	GetFreezerContract(ctx context.Context, blockNumber *big.Int) (*contracts.Freezer, error)
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
	GetTransferWhitelistContract(ctx context.Context, blockNumber *big.Int) (*contracts.TransferWhitelist, error)
	GetValidatorsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Validators, error)
}

type registryImpl struct {
	cc       *client.CeloClient
	contract *contracts.Registry
}

var (
	// ErrRegistryNotDeployed sigals that registry is not yet deployed and
	// thus no contract address can be found
	ErrRegistryNotDeployed = errors.New("Registry Not Deployed")
)

// New creates a new contracts Registry
func New(cc *client.CeloClient) (Registry, error) {
	registry, err := contracts.NewRegistry(RegistryAddress, cc.Eth)
	err = client.WrapRpcError(err)
	return &registryImpl{
		cc:       cc,
		contract: registry,
	}, err
}

func (r *registryImpl) GetAddressFor(ctx context.Context, blockNumber *big.Int, contractID ContractID) (common.Address, error) {
	address, err := r.contract.GetAddressForString(&bind.CallOpts{BlockNumber: blockNumber, Context: ctx}, contractID.String())

	err = client.WrapRpcError(err)

	if err != nil {
		return common.ZeroAddress, err
	} else if err == client.ErrContractNotDeployed {
		return common.ZeroAddress, ErrRegistryNotDeployed
	}

	if address == common.ZeroAddress {
		return common.ZeroAddress, client.ErrContractNotDeployed
	}

	return address, nil
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

func (r *registryImpl) GetFreezerContract(ctx context.Context, blockNumber *big.Int) (*contracts.Freezer, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, FreezerContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewFreezer(address, r.cc.Eth)
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

func (r *registryImpl) GetTransferWhitelistContract(ctx context.Context, blockNumber *big.Int) (*contracts.TransferWhitelist, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, TransferWhitelistContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewTransferWhitelist(address, r.cc.Eth)
}

func (r *registryImpl) GetValidatorsContract(ctx context.Context, blockNumber *big.Int) (*contracts.Validators, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, ValidatorsContractID)
	if err != nil {
		return nil, err
	}
	return contracts.NewValidators(address, r.cc.Eth)
}
