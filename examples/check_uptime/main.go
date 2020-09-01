package main

import (
	"context"
	"os"

	"github.com/celo-org/kliento/client"
	"github.com/celo-org/kliento/monitor"
	"github.com/ethereum/go-ethereum/log"
)

var ctx = context.Background()

func CeloClient() *client.CeloClient {
	celo, err := client.Dial("https://rc1-forno.celo-testnet.org/")
	if err != nil {
		panic(err)
	}
	return celo
}

func main() {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stderr, log.TerminalFormat(true))))

	epochSize := 17280

	cc := CeloClient()

	latest, err := cc.Eth.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Crit("Error Getting Latest Block", "err", err)
	}

	if err := monitor.CheckSignatures(ctx, uint64(epochSize), cc, latest.Number.Uint64()); err != nil {
		log.Crit("Error Checking Uptime", "err", err)
	}
}