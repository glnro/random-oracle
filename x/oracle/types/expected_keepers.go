package types

import (
	"context"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/glnro/random-oracle/provider"
)

type OracleKeeper interface {
	ValidateVoteExtension(ctx sdk.Context, height int64, extCommitInfo abci.ExtendedCommitInfo) (bool, error)
	GetLatestRandomRound(ctx sdk.Context) (provider.LatestRandomRound, error)
	SaveRandomness(ctx context.Context, result provider.LatestRandomRound) error
}
