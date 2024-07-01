package adverts

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"remap/testutil/sample"
	advertssimulation "remap/x/adverts/simulation"
	"remap/x/adverts/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = advertssimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateAdvert = "op_weight_msg_advert"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAdvert int = 100

	opWeightMsgUpdateAdvert = "op_weight_msg_advert"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAdvert int = 100

	opWeightMsgDeleteAdvert = "op_weight_msg_advert"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAdvert int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	advertsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		AdvertList: []types.Advert{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		AdvertCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&advertsGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateAdvert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAdvert, &weightMsgCreateAdvert, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAdvert = defaultWeightMsgCreateAdvert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAdvert,
		advertssimulation.SimulateMsgCreateAdvert(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateAdvert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateAdvert, &weightMsgUpdateAdvert, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateAdvert = defaultWeightMsgUpdateAdvert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateAdvert,
		advertssimulation.SimulateMsgUpdateAdvert(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteAdvert int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteAdvert, &weightMsgDeleteAdvert, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAdvert = defaultWeightMsgDeleteAdvert
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteAdvert,
		advertssimulation.SimulateMsgDeleteAdvert(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
