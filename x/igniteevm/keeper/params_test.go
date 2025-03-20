package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "igniteevm/testutil/keeper"
	"igniteevm/x/igniteevm/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.IgniteevmKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
