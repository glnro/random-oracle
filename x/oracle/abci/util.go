package abci

import (
	"encoding/json"
	abci "github.com/cometbft/cometbft/abci/types"
	abci2 "github.com/glnro/random-oracle/abci"
	"github.com/glnro/random-oracle/provider"
)

func UnmarshalInjectedVE(ve []byte) (*abci.ExtendedCommitInfo, error) {
	var extCommInfo abci.ExtendedCommitInfo
	err := json.Unmarshal(ve, &extCommInfo)
	if err != nil {
		return nil, err
	}
	return &extCommInfo, nil
}

func UnmarshalVE(ve []byte) (*abci2.VoteExtension, error) {
	var voteExt abci2.VoteExtension
	err := json.Unmarshal(ve, &voteExt)
	if err != nil {
		return nil, err
	}
	return &voteExt, nil
}

func ReduceVe(ve *abci.ExtendedCommitInfo) (provider.LatestRandomRound, error) {
	var totalVP int64
	reducedVE := make(map[string]int64)

	// Reduce Vote Extensions to count by Validator Power
	for _, evi := range ve.GetVotes() {
		val := evi.GetValidator()
		v, err := UnmarshalVE(evi.GetVoteExtension())
		if err != nil {
			return provider.LatestRandomRound{}, err
		}
		reducedVE = GetOrCreate(reducedVE, v.LatestRand, val)
		totalVP += val.GetPower()
	}

	return provider.LatestRandomRound{}, nil
}

func GetOrCreate(idx map[string]int64, round provider.LatestRandomRound, validator abci.Validator) map[string]int64 {
	_, ok := idx[round.Randomness]

	if !ok {
		idx[round.Randomness] = int64(0)
	}
	idx[round.Randomness] += validator.GetPower()

	return idx
}
