package types

// NewGenesisState - Create a new genesis state
func NewGenesisState(params Params, tokenInfos []TokenInfo, tokenTargetInfos []TokenTargetInfo) GenesisState {
	return GenesisState{
		Params:       params,
		TokenInfos:   tokenInfos,
		TokenTargets: tokenTargetInfos,
	}
}

// DefaultGenesisState - Return a default genesis state
func DefaultGenesisState() GenesisState {
	return NewGenesisState(
		DefaultParams(),
		[]TokenInfo{},
		[]TokenTargetInfo{},
	)
}

// ValidateGenesis performs basic validation of auth genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error {
	return nil
}
