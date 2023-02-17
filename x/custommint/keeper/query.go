package keeper

import (
	"testmint/x/custommint/types"
)

var _ types.QueryServer = Keeper{}
