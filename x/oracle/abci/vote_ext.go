package abci

import (
	"cosmossdk.io/log"
	"encoding/json"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	oracletypes "github.com/glnro/random-oracle/x/oracle/types"
)

func ExtendVote(ctx sdk.Context, req *abci.RequestExtendVote, ok oracletypes.OracleKeeper, log log.Logger) (*abci.ResponseExtendVote, error) {
	log.Info("submitting vote extensions", "height", req.Height)

	lr, err := ok.GetLatestRandomRound(ctx)
	if err != nil {
		log.Error("error fetching latest oracle update", "height", req.Height, "error", err)
		return &abci.ResponseExtendVote{VoteExtension: nil}, err
	}

	ve := oracletypes.VoteExtension{LatestRand: lr}
	log.Info("vote extension submitted", "value", ve.LatestRand, "height", req.Height)
	bz, err := json.Marshal(ve)
	if err != nil {
		log.Error("error marshalling latest oracle update", "height", req.Height, "error", err)
		return &abci.ResponseExtendVote{VoteExtension: nil}, err
	}

	return &abci.ResponseExtendVote{VoteExtension: bz}, nil
}

func VerifyVoteExt(ctx sdk.Context, req *abci.RequestVerifyVoteExtension, ok oracletypes.OracleKeeper, log log.Logger) (*abci.ResponseVerifyVoteExtension, error) {
	ve, err := UnmarshalVE(req.GetVoteExtension())
	if err != nil {
		log.Error("unable to unmarshal vote extensions", "height", req.Height, "error", err)
		return &abci.ResponseVerifyVoteExtension{Status: abci.ResponseVerifyVoteExtension_REJECT}, err
	}

	// Do additional verification
	err = ok.VerifyVoteExtension(ctx, ve)
	if err != nil {
		log.Error("invalid vote extensions", "height", req.Height, "error", err)
		return &abci.ResponseVerifyVoteExtension{Status: abci.ResponseVerifyVoteExtension_REJECT}, err
	}

	return &abci.ResponseVerifyVoteExtension{Status: abci.ResponseVerifyVoteExtension_ACCEPT}, nil
}
