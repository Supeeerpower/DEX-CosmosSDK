package keeper

import (
	"mychain/x/dex/types"
)

var _ types.QueryServer = Keeper{}
