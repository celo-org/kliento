package main

import (
	"context"
	"encoding/json"
	"math/big"
	"strconv"
	"time"

	"github.com/celo-org/kliento/client"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
)

func getLastBlockHeader(ctx context.Context, cc *client.CeloClient) (*types.Header, error) {
	lastBlockHeader, err := cc.Eth.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Crit("Error Getting Latest Block", "err", err)
		return nil, err
	}
	return lastBlockHeader, nil
}

//checkLastBlockAge checks the age of the head block. If the head block is older than blockMaxAgeSeconds,
// returns false, elsewhere returns true
func checkLastBlockAge(ctx context.Context, cc *client.CeloClient, blockMaxAgeSeconds int) (bool, error) {
	lastBlockHeader, err := getLastBlockHeader(ctx, cc)
	if err != nil {
		return false, err
	}
	lastBlockTimestamp := parseTimestamp(lastBlockHeader.Time)
	if lastBlockTimestamp.Add(time.Second * time.Duration(blockMaxAgeSeconds)).After(time.Now()) {
		return true, nil
	}
	return false, nil
}

// targetBlockPassed checks the current head block and returns true if the head block is higher
func getConnectedPeersCount(ctx context.Context, cc *client.CeloClient) (*big.Int, error) {
	var result hexutil.Big
	err := cc.Rpc.CallContext(ctx, &result, "net_peerCount")
	if err != nil {
		log.Error("Error getting peerCount", "err", err)
		return nil, err
	}
	log.Debug("Connected peers", "peerCount", (*big.Int)(&result))
	return (*big.Int)(&result), nil
}

// targetBlockPassed checks the current head block and returns true if the head block is higher
// or equal than a given target. If target==0, it checks the age of the block to check if it is
// recent
func targetBlockPassed(ctx context.Context, cc *client.CeloClient, target big.Int) (bool, error) {
	lastBlockHeader, err := getLastBlockHeader(ctx, cc)
	if err != nil {
		return false, nil
	}
	zero := big.NewInt(0)
	if target.Cmp(zero) <= 0 {
		passed, err := checkLastBlockAge(ctx, cc, 10)
		if err != nil {
			return false, err
		} else if passed {
			return true, nil
		}

	} else if lastBlockHeader.Number.Cmp(&target) != -1 {
		return true, nil
	}

	return false, nil
}

// isSyncing returns true if there's async currently running,
// if not it returns false.
func isSyncing(ctx context.Context, cc *client.CeloClient) (bool, error) {
	syncProgress, err := cc.Eth.SyncProgress(ctx)
	if err != nil {
		log.Error("Error checking syncProcess", "err", err)
		return false, err
	}
	if syncProgress == nil {
		return false, nil
	} else {
		return true, nil
	}
}

// parsePeers calls admin_peers method from api and returns
func getPeers(ctx context.Context, cc *client.CeloClient) (*[]p2p.PeerInfo, error) {
	var peers []p2p.PeerInfo
	err := cc.Rpc.CallContext(ctx, &peers, "admin_peers")
	if err != nil {
		log.Error("Error getting peers", "err", err)
		return nil, err
	}
	jsonData, _ := json.MarshalIndent(peers, "", "  ")
	log.Trace("Read admin_peers", "peers", string(jsonData))
	return &peers, nil
}

func getDebugStacks(ctx context.Context, cc *client.CeloClient) (string, error) {
	var stack string
	err := cc.Rpc.CallContext(ctx, &stack, "debug_stacks")
	if err != nil {
		log.Error("Error getting debug stack", "err", err)
		return "", err
	}
	log.Trace("Read debug_stacks", "stacks", stack)
	return stack, nil
}

func waitGethStarted(ctx context.Context, cc *client.CeloClient) (bool, error) {
	retries, err := strconv.Atoi(getenv("GETH_CONNECT_RETRIES", "10"))
	sleepSeconds := 5

	if err != nil {
		log.Crit("Cannot read GETH_TIMEOUT", "err", err)
		return false, err
	}

	for i := 0; i <= retries; i++ {
		listening, err := cc.Eth.NetworkListening(ctx)
		if err != nil {
			log.Debug("Cannot connect to RPC endpoint", "err", err)
			time.Sleep(time.Duration(sleepSeconds) * time.Second)
			continue
		}
		if listening {
			return true, nil
		}
	}
	log.Error("Cannot connect to Geth", "err", err)
	_, err = cc.Eth.NetworkListening(ctx)
	return false, err
}
