module github.com/celo-org/kliento

go 1.13

require (
	github.com/ethereum/go-ethereum v1.9.8
	github.com/onsi/gomega v1.9.0
)

replace github.com/celo-org/bls-zexe/go => ./external/bls-zexe/go

// DO NOT CHANGE DIRECTORY: Create a symlink so this works
// replace github.com/ethereum/go-ethereum => ../celo-blockchain

// Use this to use external build
replace github.com/ethereum/go-ethereum => github.com/celo-org/celo-blockchain v0.0.0-20200422165556-dee2db520e23
