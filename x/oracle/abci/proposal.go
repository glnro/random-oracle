package abci

import (
	"cosmossdk.io/log"
	"encoding/json"
	"fmt"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/glnro/random-oracle/x/oracle/types"
)

func PrepareProposal(ctx sdk.Context, req *abci.RequestPrepareProposal, ok types.OracleKeeper, logger log.Logger) ([]byte, error) {
	logger.Info("preparing vote extensions", "height", req.Height, "proposer", req.ProposerAddress)
	// Get VE from previous round
	extCommitInfo := req.LocalLastCommit

	valid, err := ok.ValidateVoteExtension(ctx, req.Height, extCommitInfo)
	if !valid {
		return nil, fmt.Errorf("invalid vote extensions: %w", err)
	}

	// Encode VE to be injected
	bz, err := json.Marshal(extCommitInfo.Votes)
	if err != nil {
		return nil, fmt.Errorf("error marshalling vote extensions: %w", err)
	}

	return bz, nil
}

func ProcessProposal(ctx sdk.Context, req *abci.RequestProcessProposal, logger log.Logger) (*abci.ResponseProcessProposal, error) {
	return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
}
