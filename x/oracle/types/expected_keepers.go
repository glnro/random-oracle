package types

import (
	"context"
	abci "github.com/cometbft/cometbft/abci/types"
	cmtprotocrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci2 "github.com/glnro/random-oracle/abci"
	"github.com/glnro/random-oracle/provider"
)

type OracleKeeper interface {
	ValidateVoteExtension(ctx sdk.Context, height int64, chainId string, extCommitInfo abci.ExtendedCommitInfo) (bool, error)
	VerifyVoteExtension(ctx sdk.Context, ve *abci2.VoteExtension) error
	GetLatestRandomRound(ctx sdk.Context) (*provider.LatestRandomRound, error)
	SaveRandomness(ctx context.Context, result provider.LatestRandomRound) error
}

type StakingKeeper interface {
	GetPubKeyByConsAddr(context.Context, sdk.ConsAddress) (cmtprotocrypto.PublicKey, error)
}
