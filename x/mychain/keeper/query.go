package keeper

import (
	"mychain/x/mychain/types"
)

var _ types.QueryServer = Keeper{}
