package registry

// ContractID represents the ID of a contract according to the Registry
type ContractID string

func (cid ContractID) String() string { return string(cid) }

var (
	// AccountsContractID is the registry Id for 'Accounts' contract
	AccountsContractID ContractID = "Accounts"

	// AttestationsContractID is the registry Id for 'Attestations' contract
	AttestationsContractID ContractID = "Attestations"

	// BlockchainParametersContractID is the registry Id for 'BlockchainParameters' contract
	BlockchainParametersContractID ContractID = "BlockchainParameters"

	// DoubleSigningSlasherContractID is the registry Id for 'DoubleSigningSlasher' contract
	DoubleSigningSlasherContractID ContractID = "DoubleSigningSlasher"

	// DowntimeSlasherContractID is the registry Id for 'DowntimeSlasher' contract
	DowntimeSlasherContractID ContractID = "DowntimeSlasher"

	// ElectionContractID is the registry Id for 'Election' contract
	ElectionContractID ContractID = "Election"

	// EpochRewardsContractID is the registry Id for 'EpochRewards' contract
	EpochRewardsContractID ContractID = "EpochRewards"

	// EscrowContractID is the registry Id for 'Escrow' contract
	EscrowContractID ContractID = "Escrow"

	// ExchangeContractID is the registry Id for 'Exchange' contract
	ExchangeContractID ContractID = "Exchange"

	// FeeCurrencyWhitelistContractID is the registry Id for 'FeeCurrencyWhitelist' contract
	FeeCurrencyWhitelistContractID ContractID = "FeeCurrencyWhitelist"

	// FreezerContractID is the registry Id for 'Freezer' contract
	FreezerContractID ContractID = "Freezer"

	// GasPriceMinimumContractID is the registry Id for 'GasPriceMinimum' contract
	GasPriceMinimumContractID ContractID = "GasPriceMinimum"

	// GoldTokenContractID is the registry Id for 'GoldToken' contract
	GoldTokenContractID ContractID = "GoldToken"

	// GovernanceContractID is the registry Id for 'Governance' contract
	GovernanceContractID ContractID = "Governance"

	// GovernanceApproverMultiSigContractID is the registry Id for 'GovernanceApproverMultiSig' contract
	GovernanceApproverMultiSigContractID ContractID = "GovernanceApproverMultiSig"

	// LockedGoldContractID is the registry Id for 'LockedGold' contract
	LockedGoldContractID ContractID = "LockedGold"

	// MultiSigContractID is the registry Id for 'MultiSig' contract
	MultiSigContractID ContractID = "MultiSig"

	// ProxyContractID is the registry Id for 'Proxy' contract
	ProxyContractID ContractID = "Proxy"

	// RandomContractID is the registry Id for 'Random' contract
	RandomContractID ContractID = "Random"

	// RegistryContractID is the registry Id for 'Registry' contract
	RegistryContractID ContractID = "Registry"

	// ReleaseGoldContractID is the registry Id for 'ReleaseGold' contract
	ReleaseGoldContractID ContractID = "ReleaseGold"

	// ReserveContractID is the registry Id for 'Reserve' contract
	ReserveContractID ContractID = "Reserve"

	// ReserveSpenderMultiSigContractID is the registry Id for 'ReserveSpenderMultiSig' contract
	ReserveSpenderMultiSigContractID ContractID = "ReserveSpenderMultiSig"

	// SortedOraclesContractID is the registry Id for 'SortedOracles' contract
	SortedOraclesContractID ContractID = "SortedOracles"

	// StableTokenContractID is the registry Id for 'StableToken' contract
	StableTokenContractID ContractID = "StableToken"

	// TransferWhitelistContractID is the registry Id for 'TransferWhitelist' contract
	TransferWhitelistContractID ContractID = "TransferWhitelist"

	// ValidatorsContractID is the registry Id for 'Validators' contract
	ValidatorsContractID ContractID = "Validators"
)
