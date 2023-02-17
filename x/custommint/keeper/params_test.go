package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "testmint/testutil/keeper"
	"testmint/x/custommint/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CustommintKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
