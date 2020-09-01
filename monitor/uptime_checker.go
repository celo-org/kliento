package monitor

import (
	"context"
	"fmt"
	"math/big"

	"github.com/celo-org/kliento/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/core/types"
)

const maxValsetSize = 100

func InitialValidators(ctx context.Context, cc *client.CeloClient) ([]common.Address, error) {
	firstBlock, err := cc.Eth.HeaderByNumber(ctx, big.NewInt(0))
	if err != nil {
		return nil, err
	}

	istExtra, err := types.ExtractIstanbulExtra(firstBlock)
	if err != nil {
		return nil, err
	}

	return istExtra.AddedValidators, nil
}

func getCurrentEpoch(ctx context.Context, epochSize uint64, cc *client.CeloClient) (uint64, error) {
	lastBlock, err := cc.Eth.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	return istanbul.GetEpochNumber(lastBlock.Number.Uint64(), epochSize), nil
}

func ApplyValidatorChanges(currentValset []common.Address, epochIstData *types.IstanbulExtra) (valset []common.Address, added int, removed int) {
	nextValidators := make([]common.Address, 0, maxValsetSize)

	// Remove validators we don't need
	removed = 0
	for i, v := range currentValset {
		if epochIstData.RemovedValidators.Bit(i) == 0 {
			nextValidators = append(nextValidators, v)
		} else {
			removed++
		}
	}
	nextValidators = append(nextValidators, epochIstData.AddedValidators...)

	return nextValidators, len(epochIstData.AddedValidators), removed
}

func EpochValidators(ctx context.Context, epochSize uint64, cc *client.CeloClient) ([][]common.Address, error) {
	currentEpoch, err := getCurrentEpoch(ctx, epochSize, cc)
	if err != nil {
		return nil, err
	}
	return EpochValidatorsAt(ctx, epochSize, currentEpoch, cc)
}

func EpochValidatorsAt(ctx context.Context, epochSize uint64, lastEpoch uint64, cc *client.CeloClient) ([][]common.Address, error) {
	genesisValidators, err := InitialValidators(ctx, cc)
	if err != nil {
		return nil, err
	}

	epochValidators := make([][]common.Address, 0)
	epochValidators = append(epochValidators, genesisValidators)
	fmt.Printf("Epoch %d #validators: %d\n", 1, len(genesisValidators))

	for epoch := uint64(1); epoch < lastEpoch; epoch++ {
		lastBlockNumber := istanbul.GetEpochLastBlockNumber(epoch, epochSize)
		lastBlockHeader, err := cc.Eth.HeaderByNumber(ctx, new(big.Int).SetUint64(lastBlockNumber))
		if err != nil {
			return nil, err
		}
		istExtra, err := types.ExtractIstanbulExtra(lastBlockHeader)

		nextValidators, removed, added := ApplyValidatorChanges(epochValidators[epoch-1], istExtra)

		fmt.Printf("Epoch %d #validators: %d new: %d removed: %d\n", epoch+1, len(nextValidators), added, removed)

		epochValidators = append(epochValidators, nextValidators)

	}
	return epochValidators, nil
}

func CheckSignatures(ctx context.Context, epochSize uint64, cc *client.CeloClient, lastBlock uint64) error {
	currentValidators, err := InitialValidators(ctx, cc)
	if err != nil {
		return err
	}

	// block, err := cc.Eth.HeaderByNumber(ctx, nil)
	// lastBlock := block.Number.Uint64()

	epoch := 1
	var missedBlocks map[common.Address]uint
	missedBlocks = make(map[common.Address]uint)

	for bn := uint64(2); bn <= lastBlock; bn++ {
		header, err := cc.Eth.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err != nil {
			return err
		}
		istExtra, err := types.ExtractIstanbulExtra(header)

		for i, v := range currentValidators {
			if (v.Hex() == "0x94306aBD8663E17B1125CAf11c53f392DB426586" || v.Hex() == "0x897c683206ddb5F134F451527C2b6f5e837407de") && istExtra.AggregatedSeal.Bitmap.Bit(i) == 0 {
				missedBlocks[v]++
				fmt.Printf("Missed Block: %d validator: %s\n", bn, v.Hex())
			}
		}

		if istanbul.IsLastBlockOfEpoch(bn, epochSize) {
			epoch++
			currentValidators, added, removed := ApplyValidatorChanges(currentValidators, istExtra)
			fmt.Printf("Starting Epoch %d #validators: %d new: %d removed: %d\n", epoch, len(currentValidators), added, removed)
		}

	}
	return nil
}
