package celotokens

import (
    "context"
    "math/big"
    "fmt"

    "github.com/celo-org/kliento/contracts"
    "github.com/celo-org/kliento/registry"
    "github.com/celo-org/celo-blockchain/common"
)

// CeloToken is a native token, eg CELO and stable tokens
type CeloToken string

const (
    // CELO - previously known as cGLD
    CELO CeloToken = "CELO"
    // CUSD - Celo Dollar
    CUSD CeloToken = "cUSD"
)

type CeloTokenInfo struct {
    contractID registry.ContractID
    exchangeContractID registry.ContractID
    isStableToken bool
}

// CeloTokenInfos contains a CeloTokenInfo entry for each CeloToken key
var CeloTokenInfos = map[CeloToken]CeloTokenInfo{
    CELO: CeloTokenInfo{
        contractID: registry.StableTokenContractID,
        isStableToken: false,
    },
    CUSD: CeloTokenInfo{
        contractID: registry.StableTokenContractID,
        exchangeContractID: registry.ExchangeContractID,
        isStableToken: true,
    },
}

// CeloTokens provides a friendly interface for interacting with Celo tokens
type CeloTokens struct {
    registry registry.Registry
}

// New creates a new pointer to a CeloTokens
func New(reg registry.Registry) (*CeloTokens) {
    return &CeloTokens{
        registry: reg,
    }
}

// IsStableToken returns whether the provided token is a stable token
func (ct *CeloTokens) IsStableToken(token CeloToken) bool {
    tokenInfo, ok := CeloTokenInfos[token]
    if !ok {
        return false
    }
    return tokenInfo.isStableToken
}

// GetExchangeContract gets the exchange contract for a provided stable token
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

// GetStableTokenContract gets the stabletoken contract for a provided stable token
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

// GetContract gets the contract for a provided celo token
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

// GetExchangeContracts gets the exchange contracts for all stable tokens
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

// GetContracts gets the contracts for all tokens. If onlyStables is true, contracts
// for only stable tokens are provided.
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

// GetAddresses gets the addresses for all token contracts. If onlyStables is true, addresses
// for only stable tokens are provided.
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
