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

type boundContracts struct {
	{{range .}}
	{{.}}Contract *contracts.{{.}}
	{{end}}
}

type generatedRegistry interface {
	GetContractByID(ctx context.Context, identifier string, blockNumber *big.Int) (interface{}, error)
	{{range .}}
	Get{{.}}Contract(ctx context.Context, blockNumber *big.Int) (*contracts.{{.}}, error)
	{{end}}
}

func (r *registryImpl) GetContractByID(ctx context.Context, identifier string, blockNumber *big.Int) (interface{}, error) {
	switch identifier {
		{{range .}}
		case {{.}}ContractID.String():
			return r.Get{{.}}Contract(ctx, blockNumber)
		{{end}}
	}
	return nil, fmt.Errorf("identifier '%s' not found", identifier)
}

{{range .}}
func (r *registryImpl) Get{{.}}Contract(ctx context.Context, blockNumber *big.Int) (*contracts.{{.}}, error) {
	identifier := {{.}}ContractID.String()
	if (r.{{.}}Contract == nil || r.cache.isDirty(identifier)) {
		address, err := r.GetAddressFor(ctx, blockNumber, {{.}}ContractID)
		if err != nil {
			return nil, err
		}
		contract, err := contracts.New{{.}}(address, r.cc.Eth)
		if err != nil {
			return nil, err
		}
		r.{{.}}Contract = contract
	}
	return r.{{.}}Contract, nil
}
{{end}}
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
