package custommint_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "testmint/testutil/keeper"
	"testmint/testutil/nullify"
	"testmint/x/custommint"
	"testmint/x/custommint/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CustommintKeeper(t)
	custommint.InitGenesis(ctx, *k, genesisState)
	got := custommint.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
