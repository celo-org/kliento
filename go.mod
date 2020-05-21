module github.com/celo-org/kliento

go 1.13

require (
	github.com/celo-org/bls-zexe/go v0.0.0-20200502082044-230d2f4866d5 // indirect
	github.com/ethereum/go-ethereum v1.9.8
)

// DO NOT CHANGE DIRECTORY: Create a symlink so this works
// replace github.com/ethereum/go-ethereum => ../celo-blockchain

// Use this to use external build
replace github.com/ethereum/go-ethereum => github.com/celo-org/celo-blockchain v0.0.0-20200519153823-adbdc7f8c27e
