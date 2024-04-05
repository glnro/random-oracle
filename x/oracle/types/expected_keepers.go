package types

import (
	"context"
	abci "github.com/cometbft/cometbft/abci/types"
	cmtprotocrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/glnro/random-oracle/provider"
)

type OracleKeeper interface {
	ValidateVoteExtension(ctx sdk.Context, height int64, extCommitInfo *abci.ExtendedCommitInfo) (bool, error)
	VerifyVoteExtension(ctx sdk.Context, ve *VoteExtension) error
	GetLatestRandomRound(ctx sdk.Context) (provider.LatestRandomRound, error)
	SaveRandomness(ctx context.Context, result provider.LatestRandomRound) error
}

type StakingKeeper interface {
	GetValidatorByConsAddr(context.Context, sdk.ConsAddress) (stakingtypes.Validator, error)
	GetPubKeyByConsAddr(ctx context.Context, addr sdk.ConsAddress) (cmtprotocrypto.PublicKey, error)
}
