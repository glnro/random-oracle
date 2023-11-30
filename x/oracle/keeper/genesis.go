package keeper

import (
	"context"
	oracletypes "github.com/glnro/random-oracle/x/oracle/types"
)

func (k *Keeper) InitGenesis(ctx context.Context, data *oracletypes.GenesisState) error {
	// TODO: Implement
	return nil
}

func (k *Keeper) ExportGenesis(ctx context.Context) (*oracletypes.GenesisState, error) {
	// TODO: Implement
	return &oracletypes.GenesisState{}, nil
}
