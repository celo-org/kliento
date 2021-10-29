package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/celo-org/kliento/client"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
)

var ctx = context.Background()

func CeloClient(rpcEndpoint string) *client.CeloClient {
	celo, err := client.Dial(rpcEndpoint)
	if err != nil {
		panic(err)
	}
	return celo
}

func generateErrorReport(firstBlock, lastBlock int64, peers *[]p2p.PeerInfo, debug string) {
	report := VictoropsReport{
		MessageType:       "CRITICAL",
		EntityId:          "geth-multipart-test",
		EntityDisplayName: "Fail syncing chain",
		GethGitCommit:     getenv("GETH_COMMIT", "no commit detected"),
		StateMessage:      fmt.Sprintf("failed to sync mainnet at block %d", lastBlock),
		PodName:           getenv("HOSTNAME", "no hostname detected"),
		ClusterName:       "integration-tests",
		StartBlock:        int(firstBlock),
		EndBlock:          int(lastBlock),
		PeersCount:        len(*peers),
		Peers:             fmt.Sprintf("%#v", peers),
		DebugStacks:       debug,
	}
	reportVictorops(report)
}

func generateErrorReportInitial(err error) {
	log.Debug("Sending report to Victorops")
	report := VictoropsReport{
		MessageType:       "CRITICAL",
		EntityId:          "geth-multipart-test",
		EntityDisplayName: "Fail syncing chain",
		GethGitCommit:     getenv("GETH_COMMIT", "no commit detected"),
		StateMessage:      "failed to sync mainnet. Pod could not start",
		PodName:           getenv("HOSTNAME", "no hostname detected"),
		ClusterName:       "integration-tests",
		ErrorMessage:      err.Error(),
	}
	reportVictorops(report)
}

func generateSuccessReport(firstBlock, lastBlock int64) {
	log.Debug("Sending report to Slack")
	report := SlackReport{
		Text: fmt.Sprintf("Successfully mainnet sync %v from block %v to %v", getenv("HOSTNAME", "NO_HOSTNAME"), firstBlock, lastBlock),
	}
	err := reportSlack(report)
	if err != nil {
		log.Error("Error reporting to slack", "err", err)
		os.Exit(-1)
	}
}

func main() {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlDebug, log.StreamHandler(os.Stderr, log.TerminalFormat(true))))

	// os.Setenv("RPC_ENDPOINT", "https://forno.celo.org")
	// os.Setenv("SLACK_ENDPOINT", "https://hooks.slack.com/services/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")
	// os.Setenv("TARGET_BLOCK", "9553299")

	rpcEndPoint := getenv("RPC_ENDPOINT", "http://localhost:8545")
	targetBlock := new(big.Int)
	targetBlock, ok := targetBlock.SetString(getenv("TARGET_BLOCK", "0"), 10)
	if !ok {
		log.Crit("Cannot cast TARGET_BLOCK env var to big.Int")
	}
	cc := CeloClient(rpcEndPoint)
	_, err := waitGethStarted(ctx, cc)
	if err != nil {
		generateErrorReportInitial(err)
		os.Exit(-1)
	}
	oldLastBlock, _ := getLastBlockHeader(ctx, cc)
	firstBlock := oldLastBlock
	oldPeers, _ := getPeers(ctx, cc)
	oldDebugStacks, _ := getDebugStacks(ctx, cc)

	for {
		time.Sleep(10 * time.Second)
		isError := false
		lastBlock, err := getLastBlockHeader(ctx, cc)
		if lastBlock.Number.Cmp(oldLastBlock.Number) != 1 {
			// ERROR
			log.Warn("Node not syncing")
		} else if err != nil {
			isError = true
		} else {
			oldLastBlock = lastBlock
		}

		peers, _ := getPeers(ctx, cc)
		if err != nil {
			isError = true
		} else {
			oldPeers = peers
		}

		debugStacks, _ := getDebugStacks(ctx, cc)
		if err != nil {
			isError = true
		} else {
			oldDebugStacks = debugStacks
		}

		if isError {
			generateErrorReport(firstBlock.Number.Int64(), oldLastBlock.Number.Int64(), oldPeers, oldDebugStacks)
		}
		passed, err := targetBlockPassed(ctx, cc, *targetBlock)
		if err != nil {
			log.Error("Cannot check if target block is passed")
		} else if passed {
			log.Info("Sync success")
			generateSuccessReport(firstBlock.Number.Int64(), oldLastBlock.Number.Int64())
			os.Exit(0)
		}

		log.Info("Iteration info", "lastBlock", oldLastBlock.Number.Int64())
	}
}
