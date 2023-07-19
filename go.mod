module github.com/celo-org/kliento

go 1.15

require (
	// github.com/btcsuite/btcutil v1.0.2 // indirect
    // TODO EN: this currently uses master;
    // needs to be updated to the next blockchain release as soon as it's available
	github.com/celo-org/celo-blockchain v0.0.0-20230713152423-4a1e3fc3b943
	// github.com/gogo/protobuf v1.3.2 // indirect
)

// DO NOT CHANGE DIRECTORY: Create a symlink so this works
// replace github.com/celo-org/celo-blockchain => ../celo-blockchain
