package keeper

import (
	"github.com/neutron-org/neutron/v7/x/cron/types"
)

var _ types.QueryServer = Keeper{}
