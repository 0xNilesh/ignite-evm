package igniteevm_test

import (
	"testing"

	keepertest "igniteevm/testutil/keeper"
	"igniteevm/testutil/nullify"
	igniteevm "igniteevm/x/igniteevm/module"
	"igniteevm/x/igniteevm/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IgniteevmKeeper(t)
	igniteevm.InitGenesis(ctx, k, genesisState)
	got := igniteevm.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
