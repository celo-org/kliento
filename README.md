# Kliento 

Kliento (/client in esperanto/) is a set of client libraries for Celo in Golang. It is a ported library from [ContractKit](https://github.com/celo-org/celo-monorepo/tree/master/packages/contractkit) and only supports a subset of what ContractKit does.

## Usage

Kliento has limited usage and capabilities and is not recommended for the majority of use cases. For most development, we recommend using [ContractKit](https://docs.celo.org/developer-guide/overview/introduction/contractkit).

## Development

Use `make gen-contracts` to update the abigen bindings to the build artifacts from local monorepo.

Use `make gen-registry` to update the generated registry bindings to a list of contracts specified in the [template script](`registry/internal/gen-registry.go`).
