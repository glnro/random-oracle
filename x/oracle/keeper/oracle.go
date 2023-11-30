package keeper

import (
	"context"
	"fmt"
	"github.com/glnro/oracle/provider"
	"github.com/glnro/oracle/types"
)

func (k *Keeper) SaveRandomness(goCtx context.Context, result provider.LatestRandomRound) error {
	lastRound, err := k.Rounds.Peek(goCtx)
	if err != nil {
		return err
	}

	if lastRound < result.Round {
		_, err = k.Rounds.Next(goCtx)
		if err != nil {
			return fmt.Errorf("unable to set new round :: %s", err.Error())
		}

		newRand := types.RandomResult{
			Round:      result.Round,
			Randomness: result.Randomness,
			Signature:  result.Signature,
		}

		err = k.RandomResults.Set(goCtx, result.Round, newRand)
		if err != nil {
			return fmt.Errorf("unable to set new random result :: %s", err.Error())
		}
	}
	return nil
}
