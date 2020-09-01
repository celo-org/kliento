package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/celo-org/kliento/client"
	"github.com/celo-org/kliento/monitor"
	"github.com/ethereum/go-ethereum/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/core/types"
)

var ctx = context.Background()

func CeloClient() *client.CeloClient {
//	celo, err := client.Dial("http://localhost:8545/")
	celo, err := client.Dial("https://baklava-forno.celo-testnet.org/")
	if err != nil {
		panic(err)
	}
	return celo
}

func CheckSignatures(ctx context.Context, epochSize uint64, cc *client.CeloClient, firstBlock uint64, lastBlock uint64) error {
	firstEpoch := istanbul.GetEpochNumber(firstBlock, epochSize)
	epochValidators, err := monitor.EpochValidatorsAt(ctx, epochSize, firstEpoch, cc)
	currentValidators := epochValidators[firstEpoch-1]
	if err != nil {
		return err
	}

	epoch := firstEpoch
	var missedBlocks map[common.Address]uint
	missedBlocks = make(map[common.Address]uint)
	sequentialBlocks := make(map[common.Address]uint)

	totalMissed := 0
	totalSlots := 0
	totalTime := uint64(0)
	maxTime := uint64(0)
	maxMissed := uint(0)
	maxSeqMissed := uint(0)

	prevHeader, err := cc.Eth.HeaderByNumber(ctx, new(big.Int).SetUint64(firstBlock-1))
	if err != nil {
		return err
	}

	for bn := firstBlock; bn <= lastBlock; bn++ {
		header, err := cc.Eth.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err != nil {
			return err
		}
		istExtra, err := types.ExtractIstanbulExtra(header)

		blockTime := header.Time - prevHeader.Time
		if blockTime > maxTime {
			maxTime = blockTime
		}

		totalTime += blockTime

		for i, v := range currentValidators {
			if istExtra.ParentAggregatedSeal.Bitmap.Bit(i) == 0 {
				missedBlocks[v]++
				sequentialBlocks[v]++
				totalMissed++
				fmt.Printf("Missed Block: %d validator: %s\n", bn, v.Hex())
			} else {
				sequentialBlocks[v] = 0
			}
			totalSlots++
			if maxMissed < missedBlocks[v] {
				maxMissed = missedBlocks[v]
			}
			if maxSeqMissed < sequentialBlocks[v] {
				maxSeqMissed = sequentialBlocks[v]
			}
		}

		if istanbul.IsLastBlockOfEpoch(bn, epochSize) {
			epoch++
			currentValidators, added, removed := monitor.ApplyValidatorChanges(currentValidators, istExtra)
			fmt.Printf("Epoch %d #validators: %d new: %d removed: %d\n", epoch, len(currentValidators), added, removed)
		}

		prevHeader = header

	}
	fmt.Printf("Average misses: %v\n", float64(totalMissed) / float64(totalSlots))
	fmt.Printf("Average block time: %v\n", float64(totalTime) / float64(lastBlock-firstBlock+1))
	fmt.Printf("Max misses by validator: %v\n", maxMissed)
	fmt.Printf("Max sequential misses by validator: %v\n", maxSeqMissed)
	fmt.Printf("Max block time: %v\n", maxTime)
	return nil
}

func main() {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stderr, log.TerminalFormat(true))))

	// epochSize := 10
	epochSize := 17280

	cc := CeloClient()

	latest, err := cc.Eth.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Crit("Error Getting Latest Block", "err", err)
	}

	log.Info("At block", "num", latest.Number)

	if err := CheckSignatures(ctx, uint64(epochSize), cc, 20000, 20100); err != nil {
		log.Crit("Error Checking Uptime", "err", err)
	}
}
