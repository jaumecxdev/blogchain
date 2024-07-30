package blogchain_test

import (
	"testing"

	keepertest "blogchain/testutil/keeper"
	"blogchain/testutil/nullify"
	blogchain "blogchain/x/blogchain/module"
	"blogchain/x/blogchain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlogchainKeeper(t)
	blogchain.InitGenesis(ctx, k, genesisState)
	got := blogchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
