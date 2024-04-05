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

	valid, err := ok.ValidateVoteExtension(ctx, req.Height, &extCommitInfo)
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

func ProcessProposal(ctx sdk.Context, req *abci.RequestProcessProposal, ok types.OracleKeeper, logger log.Logger) (*abci.ResponseProcessProposal, error) {
	totalCount := len(req.Txs)
	if totalCount == 0 {
		logger.Error("missing vote extensions", "height", req.Height)
		return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, fmt.Errorf("missing vote extensions")
	}

	if totalCount >= 1 {
		// Get Vote Extensions from proposal
		injVE, err := UnmarshalInjectedVE(req.Txs[0])
		if err != nil {
			logger.Error("unable to unmarshal vote extensions", "height", req.Height)
			return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, err
		}

		logger.Info("processing vote extensions for proposal", "height", req.Height, "vote-extensions", injVE.String())

		validated, err := ok.ValidateVoteExtension(ctx, req.Height, injVE)
		if err != nil {
			logger.Error("invalid vote extensions", "height", req.Height)
			return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, err
		}

		if !validated {
			logger.Error("invalid vote extension signatures", "height", req.Height)
			return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, err
		}

		//TODO: Additional validations on vote extensions here
		return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
	}

	return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
}
