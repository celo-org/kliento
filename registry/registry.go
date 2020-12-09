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
	"fmt"
	"math/big"
	"strings"

	"github.com/celo-org/kliento/client"
	"github.com/celo-org/kliento/contracts"
	"github.com/celo-org/kliento/contracts/helpers"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	blockchainErrors "github.com/ethereum/go-ethereum/contract_comm/errors"
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
	TryParseLog(ctx context.Context, eventLog *types.Log, blockNumber *big.Int) ([]interface{}, error)
	generatedRegistry
}

type addressContainer struct {
	address common.Address
	dirty   bool
}

type registryImpl struct {
	cc                    *client.CeloClient
	RegistryContract      *contracts.Registry
	RegistryContractProxy *contracts.Proxy
	cache                 map[string]*addressContainer
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

	registryProxy, err := contracts.NewProxy(RegistryAddress, cc.Eth)
	err = client.WrapRpcError(err)

	return &registryImpl{
		cc:                    cc,
		RegistryContract:      registry,
		RegistryContractProxy: registryProxy,
		cache:                 make(map[string]*addressContainer),
	}, err
}

func (r *registryImpl) GetRegistryContract() *contracts.Registry {
	return r.RegistryContract
}

func (r *registryImpl) putCache(identifier string, address common.Address) {
	r.cache[identifier] = &addressContainer{
		address: address,
		dirty:   true,
	}
}

func (r *registryImpl) getCache(identifier string) *addressContainer {
	if addressContainer, ok := r.cache[identifier]; ok {
		return addressContainer
	}
	return nil
}

func (r *registryImpl) isCacheDirty(identifier string) bool {
	ac := r.getCache(identifier)
	if ac != nil {
		return ac.dirty
	}
	return false
}

func isErrSubset(err error, suberr error) bool {
	return strings.Contains(err.Error(), suberr.Error())
}

// IsExpectedBeforeContractsDeployed checks for expected errors when contracts are not deployed yet
func IsExpectedBeforeContractsDeployed(err error) bool {
	return isErrSubset(err, blockchainErrors.ErrRegistryContractNotDeployed) || isErrSubset(err, blockchainErrors.ErrSmartContractNotDeployed) || isErrSubset(err, client.ErrContractNotDeployed)
}

func (r *registryImpl) GetAddressFor(ctx context.Context, blockNumber *big.Int, contractID ContractID) (common.Address, error) {
	identifier := contractID.String()
	// retrieve cached result and mark clean
	ac := r.getCache(identifier)
	if ac != nil {
		ac.dirty = false
		return ac.address, nil
	}

	// if uncached, fetch from contract state
	address, err := r.RegistryContract.GetAddressForString(&bind.CallOpts{BlockNumber: blockNumber, Context: ctx}, identifier)
	err = client.WrapRpcError(err)
	if err != nil {
		return common.ZeroAddress, err
	} else if address == common.ZeroAddress {
		return address, client.ErrContractNotDeployed
	}

	r.putCache(identifier, address)
	return address, nil
}

// TryParseLog parses an event log using a fully hydrated registry
func (r *registryImpl) TryParseLog(ctx context.Context, eventLog *types.Log, blockNumber *big.Int) ([]interface{}, error) {
	if eventLog.Address == params.RegistrySmartContractAddress {
		log := *eventLog
		eventName, event, ok, err := r.RegistryContract.TryParseLog(log)
		if err != nil || !ok {
			proxyEventName, proxyEvent, proxyOk, proxyErr := r.RegistryContractProxy.TryParseLog(log)
			if proxyErr == nil && proxyOk {
				return helpers.BuildEventSlice("RegistryProxy", proxyEventName, proxyEvent)
			}
			return nil, fmt.Errorf("registry error %v, proxy error %v", err, proxyErr)
		}
		switch v := event.(type) {
		case contracts.RegistryRegistryUpdated:
			r.putCache(v.Identifier, v.Addr)
		}
		return helpers.BuildEventSlice("Registry", eventName, event)
	}

	return r.tryParseLogGenerated(ctx, eventLog, blockNumber)
}
