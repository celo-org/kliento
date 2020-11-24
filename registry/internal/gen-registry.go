package main

import (
	"bytes"
	"go/format"
	"os"
	"path"
	"text/template"
)

var contractsToGenerate = []string{
	"Accounts",
	"Attestations",
	"BlockchainParameters",
	"DoubleSigningSlasher",
	"DowntimeSlasher",
	"Election",
	"EpochRewards",
	"Escrow",
	"Exchange",
	"GasPriceMinimum",
	"GoldToken",
	"Governance",
	"GovernanceSlasher",
	"LockedGold",
	"Random",
	"Reserve",
	"SortedOracles",
	"StableToken",
	"Validators",
}

var templateStr = `
// Code generated - DO NOT EDIT.

package registry

import (
  	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/celo-org/kliento/contracts"
	"github.com/celo-org/kliento/contracts/helpers"
	"github.com/celo-org/kliento/client"
)

{{range .}}
// {{.}}ContractID is the registry identifier for '{{.}}' contract
var {{.}}ContractID ContractID = "{{.}}"
{{end}}

// RegisteredContractIDs are all (known) registered contract identifiers
var RegisteredContractIDs = []ContractID{
	{{range .}}
	{{.}}ContractID,
	{{end}}
}

type boundRegistry struct {
	{{range .}}
	_{{.}}Address *common.Address
	{{.}}Contract *contracts.{{.}}
	{{end}}
}

type generatedRegistry interface {
	{{range .}}
	Get{{.}}Contract(ctx context.Context, blockNumber *big.Int) (*contracts.{{.}}, error)
	{{end}}
}

{{range .}}
func (r *registryImpl) Get{{.}}Contract(ctx context.Context, blockNumber *big.Int) (*contracts.{{.}}, error) {
	address, err := r.GetAddressFor(ctx, blockNumber, {{.}}ContractID)
	if err != nil {
		return nil, err
	} else if address == *r._{{.}}Address {
		return r.{{.}}Contract, nil
	}
	
	contract, err := contracts.New{{.}}(address, r.cc.Eth)
	if err != nil {
		return nil, err
	}
	r._{{.}}Address, r.{{.}}Contract = &address, contract
	return contract, nil
}
{{end}}

// Hydrate populates contract bindings using the registry's mapping at the provided blockNumber 
// for all known and deployed contracts
func (r *registryImpl) Hydrate(ctx context.Context, blockNumber *big.Int) (error) {
	check := func(err error) (bool) {
		return err != nil && err != ErrRegistryNotDeployed && err != client.ErrContractNotDeployed
	}

	var err error
	{{range .}}
	_, err = r.Get{{.}}Contract(ctx, blockNumber)
	if check(err) {
		return err
	}
	{{end}}
	return nil
}

// ParseLog parses an event log using a "hydrated" registry
// Hydrate should be called at the desired block number prior to parsing for comprehensive results
func (r *registryImpl) ParseLog(eventLog types.Log) ([]interface{}) {
	buildSlice := func(contractName string, eventName string, event interface{}) []interface{} {
		slice := []interface{}{"contract", contractName, "event", eventName}
		eventSlice, _ := helpers.EventToSlice(event)
		return append(slice, eventSlice)
	}
	{{range .}}
	if (r.{{.}}Contract != nil) {
		eventName, event, ok, err := r.{{.}}Contract.TryParseLog(eventLog)
		if ok && err != nil {
			return buildSlice("{{.}}", eventName, event)
		}
	}
	{{end}}
	return nil
}
`

func main() {
	template := template.Must(template.New("registryTemplate").Parse(templateStr))

	var buf bytes.Buffer

	template.Execute(&buf, contractsToGenerate)
	p, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}

	f, err := os.Create(path.Join("registry/gen_registry.go"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(p)
}
