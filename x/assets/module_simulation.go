package assets

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"remap/testutil/sample"
	assetssimulation "remap/x/assets/simulation"
	"remap/x/assets/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = assetssimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateAsset = "op_weight_msg_asset"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAsset int = 100

	opWeightMsgUpdateAsset = "op_weight_msg_asset"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAsset int = 100

	opWeightMsgDeleteAsset = "op_weight_msg_asset"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAsset int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	assetsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		AssetList: []types.Asset{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		AssetCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&assetsGenesis)
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

	var weightMsgCreateAsset int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAsset, &weightMsgCreateAsset, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAsset = defaultWeightMsgCreateAsset
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAsset,
		assetssimulation.SimulateMsgCreateAsset(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateAsset int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateAsset, &weightMsgUpdateAsset, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateAsset = defaultWeightMsgUpdateAsset
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateAsset,
		assetssimulation.SimulateMsgUpdateAsset(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteAsset int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteAsset, &weightMsgDeleteAsset, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAsset = defaultWeightMsgDeleteAsset
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteAsset,
		assetssimulation.SimulateMsgDeleteAsset(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
