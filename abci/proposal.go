package abci

import (
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (h *PrepareProposalHandler) PrepareProposalHandler() sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
		var totalTxBytes int64
		var txs [][]byte

		// Process & validate VE for Oracle
		veBz, err := h.oraclePrepareProposal(ctx, req, h.ok, h.logger)
		if err != nil {
			return &abci.ResponsePrepareProposal{Txs: req.Txs}, err
		}

		// Inject VE as first Tx in proposal
		if veBz != nil {
			totalTxBytes = int64(len(veBz))
			txs = append(txs, veBz)
		}

		// Append transactions until max bytes reached
		for _, tx := range req.Txs {
			txSize := int64(len(tx))
			totalTxBytes += txSize
			if totalTxBytes <= req.MaxTxBytes {
				txs = append(txs, tx)
			} else {
				break
			}
		}

		return &abci.ResponsePrepareProposal{Txs: txs}, nil
	}
}

func (h *PrepareProposalHandler) ProcessProposalHandler() sdk.ProcessProposalHandler {
	return func(_ sdk.Context, _ *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
		// TODO: Validate signatures on VE
		return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
	}
}

func (h *PrepareProposalHandler) PreBlocker(ctx sdk.Context, req *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
	res := &sdk.ResponsePreBlock{}
	if len(req.Txs) == 0 {
		return res, nil
	}

	return res, nil
}
