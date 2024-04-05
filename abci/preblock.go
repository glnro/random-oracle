package abci

import (
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (h *PreBlockHandler) PreBlocker() sdk.PreBlocker {
	return func(ctx sdk.Context, req *abci.RequestFinalizeBlock) (*sdk.ResponsePreBlock, error) {
		err := h.oraclePreBlock(ctx, req, h.ok, h.logger)
		if err != nil {
			return &sdk.ResponsePreBlock{}, err
		}

		return h.mm.PreBlock(ctx)
	}
}
