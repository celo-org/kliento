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
)

type celoTokenInfo struct {
    contractID registry.ContractID
    isStableToken bool
    exchangeContractID registry.ContractID
}

var celoTokenInfos = map[CeloToken]celoTokenInfo{
    CELO: celoTokenInfo{
        contractID: registry.StableTokenContractID,
        isStableToken: false,
    },
    CUSD: celoTokenInfo{
        contractID: registry.StableTokenContractID,
        isStableToken: true,
        exchangeContractID: registry.ExchangeContractID,
    },
}

// var stableTokenExchangeContractIDs = map[CeloToken]registry.ContractID{
//     CUSD: registry.ExchangeContractID,
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
    tokenInfo, ok := celoTokenInfos[token]
    if !ok {
        return nil, fmt.Errorf("Token %w not found", token)
    }
    if !tokenInfo.isStableToken {
        return nil, fmt.Errorf("Token %w not a stable token", token)
    }
    contract, err := ct.registry.GetContractByID(ctx, string(tokenInfo.exchangeContractID), blockNumber)
    if err != nil {
        return nil, err
    }
    return contract.(*contracts.Exchange), nil
}

func (ct *CeloTokens) GetStableTokenContract(ctx context.Context, token CeloToken, blockNumber *big.Int) (*contracts.StableToken, error) {
    tokenInfo, ok := celoTokenInfos[token]
    if !ok {
        return nil, fmt.Errorf("Token %w not found", token)
    }
    if !tokenInfo.isStableToken {
        return nil, fmt.Errorf("Token %w not a stable token", token)
    }
    contract, err := ct.registry.GetContractByID(ctx, string(tokenInfo.contractID), blockNumber)
    if err != nil {
        return nil, err
    }
    return contract.(*contracts.StableToken), nil
}

func (ct *CeloTokens) GetContract(ctx context.Context, token CeloToken, blockNumber *big.Int) (interface {}, error) {
    tokenInfo, ok := celoTokenInfos[token]
    if !ok {
        return nil, fmt.Errorf("Token %w not found", token)
    }
    return ct.registry.GetContractByID(ctx, string(tokenInfo.contractID), blockNumber)
}
