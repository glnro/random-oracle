package keeper

import (
	"context"
	oracletypes "github.com/glnro/random-oracle/x/oracle/types"
)

var _ oracletypes.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the module QueryServer.
func NewQueryServerImpl(k Keeper) oracletypes.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

func (q queryServer) RandomRound(ctx context.Context, request *oracletypes.QueryRandomRoundRequest) (*oracletypes.QueryRandomRoundResponse, error) {
	//TODO implement me
	panic("implement me")
}
