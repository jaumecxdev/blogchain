package keeper

import (
	"blogchain/x/blogchain/types"
)

var _ types.QueryServer = Keeper{}
