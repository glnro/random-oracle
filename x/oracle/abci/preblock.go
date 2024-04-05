package abci

import (
	"cosmossdk.io/log"
	"fmt"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/glnro/random-oracle/x/oracle/types"
)

func PreBlocker(ctx sdk.Context, req *abci.RequestFinalizeBlock, ok types.OracleKeeper, log log.Logger) error {
	txsCount := len(req.Txs)
	if txsCount >= 1 {
		// unmarshall ve
		injVE, err := UnmarshalInjectedVE(req.Txs[0])
		if err != nil {
			log.Error("error unmarshalling vote extensions during preblock", "module", "oracle", "error", err)
			return err
		}

		//
		latestRandomRound, err := ReduceVe(injVE)
		if err != nil {
			log.Error("unable to calculate latest random beacon", "error", err)
			return err
		}

		err = ok.SaveRandomness(ctx, latestRandomRound)
		if err != nil {
			log.Error(fmt.Sprintf("error persisting random oracle update", "error", err))
		}
		return nil
	}
	return nil
}
