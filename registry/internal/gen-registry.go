package main

import (
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
	"FeeCurrencyWhitelist",
	"GasPriceMinimum",
	"GoldToken",
	"Governance",
	"GovernanceApproverMultiSig",
	"LockedGold",
	"MultiSig",
	"Proxy",
	"Random",
	"Registry",
	"ReleaseGold",
	"Reserve",
	"ReserveSpenderMultiSig",
	"SortedOracles",
	"StableToken",
	"Validators",
}

var templateStr = `
package registry

import (
  "context"
	"math/big"

	"github.com/celo-org/kliento/contracts"
)

{{range .}}
// {{.}}ContractID is the registry Id for '{{.}}' contract
var {{.}}ContractID ContractID = "{{.}}"
{{end}}

type GeneratedRegistry interface {
{{range .}}
  Get{{.}}Contract(ctx context.Context, blockNumber *big.Int) (*contracts.{{.}}, error)
{{end}}
}

{{range .}}
func (r *registryImpl) Get{{.}}Contract(ctx context.Context, blockNumber *big.Int) (*contracts.{{.}}, error) {
  address, err := r.GetAddressFor(ctx, blockNumber, {{.}}ContractID)
 	if err != nil {
 		return nil, err
 	}
 	return contracts.New{{.}}(address, r.cc.Eth)
}
{{end}}
`

func main() {
	template := template.Must(template.New("registryTemplate").Parse(templateStr))

	f, err := os.Create(path.Join("registry/gen_registry.go"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	template.Execute(f, contractsToGenerate)
}
