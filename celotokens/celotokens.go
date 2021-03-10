package celotokens

import (
    "context"
    "math/big"
    "fmt"

    "github.com/celo-org/kliento/contracts"
    "github.com/celo-org/kliento/registry"
    "github.com/celo-org/celo-blockchain/common"
)

// CeloToken is an int representation of a native token on Celo, like CELO
// or stable tokens
type CeloToken string

const (
    // CELO - previously known as cGLD
    CELO CeloToken = "CELO"
    // CUSD - Celo Dollar
    CUSD CeloToken = "cUSD"
)

type CeloTokenInfo struct {
    contractID registry.ContractID
    isStableToken bool
    exchangeContractID registry.ContractID
}

// CeloTokenInfos contains a CeloTokenInfo entry for each CeloToken key
var CeloTokenInfos = map[CeloToken]CeloTokenInfo{
    CELO: CeloTokenInfo{
        contractID: registry.StableTokenContractID,
        isStableToken: false,
    },
    CUSD: CeloTokenInfo{
        contractID: registry.StableTokenContractID,
        isStableToken: true,
        exchangeContractID: registry.ExchangeContractID,
    },
}

type CeloTokens struct {
    registry registry.Registry
}

func New(reg registry.Registry) (*CeloTokens) {
    return &CeloTokens{
        registry: reg,
    }
}

func (ct *CeloTokens) IsStableToken(token CeloToken) bool {
    tokenInfo, ok := CeloTokenInfos[token]
    if !ok {
        return false
    }
    return tokenInfo.isStableToken
}

func (ct *CeloTokens) GetExchangeContract(ctx context.Context, token CeloToken, blockNumber *big.Int) (*contracts.Exchange, error) {
    tokenInfo, ok := CeloTokenInfos[token]
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
    tokenInfo, ok := CeloTokenInfos[token]
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

func (ct *CeloTokens) GetContract(ctx context.Context, token CeloToken, blockNumber *big.Int) (contracts.CeloTokenContract, error) {
    tokenInfo, ok := CeloTokenInfos[token]
    if !ok {
        return nil, fmt.Errorf("Token %w not found", token)
    }
    contract, err := ct.registry.GetContractByID(ctx, string(tokenInfo.contractID), blockNumber)
    if err != nil {
        return nil, err
    }
    return contract.(contracts.CeloTokenContract), nil
}

func (ct *CeloTokens) GetExchangeContracts(ctx context.Context, blockNumber *big.Int) (map[CeloToken]*contracts.Exchange, error) {
    exchangeContracts := make(map[CeloToken]*contracts.Exchange)
    for token, tokenInfo := range CeloTokenInfos {
        // Only stable tokens have a corresponding exchange contract
        if !tokenInfo.isStableToken {
            continue
        }
        exchangeContract, err := ct.GetExchangeContract(ctx, token, blockNumber)
        if err != nil {
            return nil, err
        }
        exchangeContracts[token] = exchangeContract
    }
    return exchangeContracts, nil
}

func (ct *CeloTokens) GetContracts(ctx context.Context, blockNumber *big.Int, onlyStables bool) (map[CeloToken]contracts.CeloTokenContract, error) {
    contracts := make(map[CeloToken]contracts.CeloTokenContract)
    for token, tokenInfo := range CeloTokenInfos {
        if onlyStables && !tokenInfo.isStableToken {
            continue
        }
        contract, err := ct.GetContract(ctx, token, blockNumber)
        if err != nil {
            return nil, err
        }
        contracts[token] = contract
    }
    return contracts, nil
}

func (ct *CeloTokens) GetAddresses(ctx context.Context, blockNumber *big.Int, onlyStables bool) (map[CeloToken]common.Address, error) {
    addresses := make(map[CeloToken]common.Address)
    for token, tokenInfo := range CeloTokenInfos {
        if onlyStables && !tokenInfo.isStableToken {
            continue
        }
        address, err := ct.registry.GetAddressFor(ctx, blockNumber, tokenInfo.contractID)
        if err != nil {
            return nil, err
        }
        addresses[token] = address
    }
    return addresses, nil
}
