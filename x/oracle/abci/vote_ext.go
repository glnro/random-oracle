package abci

import (
	"cosmossdk.io/log"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func ExtendVote(ctx sdk.Context, req *abci.RequestExtendVote, log log.Logger) (*abci.ResponseExtendVote, error) {
	log.Info("submitting vote extensions", "height", req.Height)

	//lr, err := ok.GetLatestRandomRound(ctx)

	return nil, nil
}

func VerifyVoteExt(ctx sdk.Context, req *abci.RequestVerifyVoteExtension, log log.Logger) (*abci.ResponseVerifyVoteExtension, error) {
	return nil, nil
}
