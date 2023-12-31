package abci

import (
	"cosmossdk.io/log"
	"encoding/json"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/glnro/random-oracle/provider"
)

func NewVoteExtensionHandler(l log.Logger, p *provider.RandProvider) *VoteExtHandler {
	return &VoteExtHandler{
		logger:       l,
		currentBlock: 0,
		randProvider: p,
	}
}

func (h *VoteExtHandler) ExtendVoteHandler() sdk.ExtendVoteHandler {
	return func(_ sdk.Context, _ *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {

		r, err := h.randProvider.GetLatestRandomness()
		if err != nil {
			return &abci.ResponseExtendVote{VoteExtension: []byte{}}, err
		}

		// Example only, use better encoding scheme
		bz, err := json.Marshal(r)

		return &abci.ResponseExtendVote{VoteExtension: bz}, nil
	}
}

func (h *VoteExtHandler) VerifyVoteExtensionHandler() sdk.VerifyVoteExtensionHandler {
	return func(ctx sdk.Context, req *abci.RequestVerifyVoteExtension) (*abci.ResponseVerifyVoteExtension, error) {
		// TODO: Implement
		return &abci.ResponseVerifyVoteExtension{Status: abci.ResponseVerifyVoteExtension_ACCEPT}, nil
	}
}
