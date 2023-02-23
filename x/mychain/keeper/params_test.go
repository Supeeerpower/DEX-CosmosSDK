package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "mychain/testutil/keeper"
	"mychain/x/mychain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MychainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
