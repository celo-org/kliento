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
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/celo-org/kliento/contracts"
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
	{{.}}Contract *contracts.{{.}}
	{{.}}ContractProxy *contracts.Proxy
	{{end}}
}

type generatedRegistry interface {
	{{range .}}
	Get{{.}}Contract(ctx context.Context, blockNumber *big.Int) (*contracts.{{.}}, error)
	{{end}}
}

{{range .}}
func (r *registryImpl) Get{{.}}Contract(ctx context.Context, blockNumber *big.Int) (*contracts.{{.}}, error) {
	identifier := {{.}}ContractID.String()
	if (r.{{.}}Contract == nil || r.isCacheDirty(identifier)) {
		address, err := r.GetAddressFor(ctx, blockNumber, {{.}}ContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.New{{.}}(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.{{.}}Contract = contract
		contractProxy, err := contracts.NewProxy(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.{{.}}ContractProxy = contractProxy
	}
	return r.{{.}}Contract, nil
}
{{end}}

func (r *registryImpl) tryParseLogGenerated(ctx context.Context, eventLog *types.Log, blockNumber *big.Int) (*RegistryParsedLog, error) {
	var eventName string
	var event interface{}
	var ok bool
	var err error
	
	log := *eventLog
	{{range .}}
	_, err = r.Get{{.}}Contract(ctx, blockNumber)
	if err == nil {
		eventName, event, ok, err = r.{{.}}Contract.TryParseLog(log) // checks matching address
		if ok && err == nil {
			return &RegistryParsedLog{
				Contract: "{{.}}",
				Event: eventName,
				Log: event,
			}, nil
		}
		eventName, event, ok, err = r.{{.}}ContractProxy.TryParseLog(log) // checks matching address
		if ok && err == nil {
			return &RegistryParsedLog{
				Contract: "{{.}}Proxy",
				Event: eventName,
				Log: event,
			}, nil
		}
	} else {
		// skip deployed failures
		if !IsExpectedBeforeContractsDeployed(err) {
			return nil, fmt.Errorf("{{.}} %v", err)
		}
	}
	
	{{end}}
	return nil, nil
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
