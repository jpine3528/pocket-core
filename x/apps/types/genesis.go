package types

import (
	sdk "github.com/pokt-network/pocket-core/types"
)

// GenesisState - all staking state that must be provided at genesis
type GenesisState struct {
	Params       Params       `json:"params" yaml:"params"`
	Applications Applications `json:"applications" yaml:"applications"`
	Exported     bool         `json:"exported" yaml:"exported"`
}

// PrevState application power, needed for application set update logic
type PrevStatePowerMapping struct {
	Address sdk.Address
	Power   int64
}

// get raw genesis raw message for testing
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Params:       DefaultParams(),
		Applications: make(Applications, 0),
	}
}
