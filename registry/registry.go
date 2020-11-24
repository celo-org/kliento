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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

// ContractID represents the ID of a contract according to the Registry
type ContractID string

func (cid ContractID) String() string { return string(cid) }

// RegistryAddress is the address of the registry which is the same on any celo network
var RegistryAddress = params.RegistrySmartContractAddress

// Registry defines an interface to access all Celo core Contracts
type Registry interface {
	GetAddressFor(ctx context.Context, blockNumber *big.Int, contractID ContractID) (common.Address, error)
	Hydrate(ctx context.Context, blockNumber *big.Int) error
	ParseLog(eventLog types.Log) []interface{}
	generatedRegistry
}

type registryImpl struct {
	cc               *client.CeloClient
	RegistryContract *contracts.Registry
	cache            map[string]common.Address
	boundRegistry
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
		cc:               cc,
		RegistryContract: registry,
		cache:            make(map[string]common.Address),
	}, err
}

func (r *registryImpl) GetAddressFor(ctx context.Context, blockNumber *big.Int, contractID ContractID) (common.Address, error) {
	// TODO: more sophisticated caching
	if address, ok := r.cache[contractID.String()]; ok {
		return address, nil
	}

	address, err := r.RegistryContract.GetAddressForString(&bind.CallOpts{BlockNumber: blockNumber, Context: ctx}, contractID.String())
	err = client.WrapRpcError(err)

	if err == client.ErrContractNotDeployed {
		return common.ZeroAddress, ErrRegistryNotDeployed
	} else if err != nil {
		return common.ZeroAddress, err
	} else if address == common.ZeroAddress {
		return common.ZeroAddress, client.ErrContractNotDeployed
	}

	r.cache[contractID.String()] = address
	return address, nil
}

func (r *registryImpl) UpdateCacheForBlocks(ctx context.Context, start uint64, end *uint64) error {
	opts := &bind.FilterOpts{
		Start:   start,
		End:     end,
		Context: ctx,
	}
	iter, err := r.RegistryContract.RegistryFilterer.FilterRegistryUpdated(opts, nil, nil)
	if err != nil {
		return err
	}
	for iter.Next() {
		event := iter.Event
		// TODO: add more sophisticated caching with event.raw.blocknumber
		r.cache[event.Identifier] = event.Addr
	}
	return nil
}
