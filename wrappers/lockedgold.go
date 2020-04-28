package wrappers

import (
	"math/big"

	"github.com/celo-org/kliento/client"
	"github.com/celo-org/kliento/contracts"
	"github.com/celo-org/kliento/utils/bn"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type LockedGoldWrapper struct {
	*contracts.LockedGold
}

func NewLockedGold(celoClient *client.CeloClient, registryWrapper *RegistryWrapper) (*LockedGoldWrapper, error) {
	lockedgold, err := registryWrapper.GetLockedGold(nil, celoClient.Eth)
	if err != nil {
		return nil, err
	}

	return &LockedGoldWrapper{
		lockedgold,
	}, nil
}

type PendingWithdrawal struct {
	Amount    *big.Int
	Timestamp *big.Int
}

type NonVotingLockedGold struct {
	Amount             *big.Int
	PendingWithdrawals []PendingWithdrawal
}

func (w *LockedGoldWrapper) GetTotalPendingWithdrawals(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	values, _, err := w.LockedGold.GetPendingWithdrawals(opts, account)
	if err != nil {
		return nil, err
	}

	return bn.Sum(values...), nil
}

func (w *LockedGoldWrapper) GetPendingWithdrawals(opts *bind.CallOpts, account common.Address) ([]PendingWithdrawal, error) {
	values, timestamps, err := w.LockedGold.GetPendingWithdrawals(opts, account)
	if err != nil {
		return nil, err
	}

	withdrawals := make([]PendingWithdrawal, len(values))
	for idx, val := range values {
		withdrawals[idx].Amount = val
		withdrawals[idx].Timestamp = timestamps[idx]
	}
	return withdrawals, nil
}
