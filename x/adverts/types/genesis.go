package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AdvertList: []Advert{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in advert
	advertIdMap := make(map[uint64]bool)
	advertCount := gs.GetAdvertCount()
	for _, elem := range gs.AdvertList {
		if _, ok := advertIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for advert")
		}
		if elem.Id >= advertCount {
			return fmt.Errorf("advert id should be lower or equal than the last id")
		}
		advertIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
