package abci

import (
	"cosmossdk.io/log"
	"encoding/json"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewProposalHandler(
	lg log.Logger,
	txCg client.TxConfig,
) *ProposalHandler {
	return &ProposalHandler{
		logger:   lg,
		txConfig: txCg,
	}
}

func (h *ProposalHandler) PrepareProposalHandler() sdk.PrepareProposalHandler {
	return func(_ sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
		var txs [][]byte
		// Get VE from previous round
		ve := req.LocalLastCommit.Votes

		// Example encoding only
		bz, err := json.Marshal(ve)
		if err != nil {
			return &abci.ResponsePrepareProposal{Txs: req.Txs}, err
		}

		// Inject VE as first Tx in proposal
		txs = append(txs, bz)
		// Append remaining reaped Txs from Comet
		txs = append(txs, req.Txs...)

		return &abci.ResponsePrepareProposal{Txs: txs}, nil
	}
}

func (h *ProposalHandler) ProcessProposalHandler() sdk.ProcessProposalHandler {
	return func(_ sdk.Context, _ *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
		// TODO: Validate signatures on VE
		return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
	}
}

func (h *ProposalHandler) PreBlocker(ctx sdk.Context, req *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	res := &sdk.ResponsePreBlock{}
	if len(req.Txs) == 0 {
		return res, nil
	}

	return res, nil
}
