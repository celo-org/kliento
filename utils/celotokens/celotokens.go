package celotokens

import (
    "context"
    "math/big"
    "fmt"

    "github.com/celo-org/kliento/contracts"
    "github.com/celo-org/kliento/registry"
)

// CeloToken is an int representation of a native token on Celo, like CELO
// or stable tokens
type CeloToken string

const (
    CELO CeloToken = "CELO"
    CUSD CeloToken = "cUSD"
    CEUR CeloToken = "cEUR"
)

type stableTokenInfo struct {
    contractID registry.ContractID
    exchangeContractID registry.ContractID
}

var stableTokenInfos = map[CeloToken]stableTokenInfo{
	CUSD: stableTokenInfo{
        contractID: registry.StableTokenContractID,
        exchangeContractID: registry.ExchangeContractID,
    },
}

// const celoTokenInfos map[CeloToken]CeloTokenInfo{
// 	CELO: CeloTokenInfo{
//         contractID: registry.GoldTokenContractID
//     }
// }

type CeloTokens struct {
    registry registry.Registry
}

func New(reg registry.Registry) (*CeloTokens) {
    return &CeloTokens{
        registry: reg,
    }
}

func (ct *CeloTokens) GetExchangeContract(ctx context.Context, token CeloToken, blockNumber *big.Int) (*contracts.Exchange, error) {
    tokenInfo, ok := stableTokenInfos[token]
    if !ok {
        return nil, fmt.Errorf("Token %w not found in stable token infos", stableTokenInfos)
    }
    contract, err := ct.registry.GetContractByID(ctx, string(tokenInfo.exchangeContractID), blockNumber)
    if (err != nil) {
        return nil, err
    }
    return contract.(*contracts.Exchange), nil
}

func (ct *CeloTokens) GetStableTokenContract(ctx context.Context, token CeloToken, blockNumber *big.Int) (*contracts.StableToken, error) {
    tokenInfo, ok := stableTokenInfos[token]
    if !ok {
        return nil, fmt.Errorf("Token %w not found in stable token infos", stableTokenInfos)
    }
    contract, err := ct.registry.GetContractByID(ctx, string(tokenInfo.contractID), blockNumber)
    if (err != nil) {
        return nil, err
    }
    return contract.(*contracts.StableToken), nil
}
