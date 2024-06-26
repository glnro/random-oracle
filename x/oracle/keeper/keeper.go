package keeper

import (
	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/glnro/random-oracle/provider"
	oracletypes "github.com/glnro/random-oracle/x/oracle/types"
)

type Keeper struct {
	cdc    codec.BinaryCodec
	logger log.Logger

	randProvider provider.RandomnessProvider
	sk           oracletypes.StakingKeeper

	// state management
	Schema        collections.Schema
	Rounds        collections.Sequence
	RandomResults collections.Map[uint64, oracletypes.RandomResult]
}

func NewKeeper(cdc codec.BinaryCodec, storeService storetypes.KVStoreService, randProvider provider.RandomnessProvider, sk oracletypes.StakingKeeper, logger log.Logger) Keeper {
	sb := collections.NewSchemaBuilder(storeService)
	rnds := collections.NewSequence(sb, oracletypes.RoundKey, "rounds")
	rndRes := collections.NewMap(sb, oracletypes.RandomResultsKey, "randomres", collections.Uint64Key, codec.CollValue[oracletypes.RandomResult](cdc))

	k := Keeper{
		cdc:          cdc,
		logger:       logger,
		randProvider: randProvider,
		sk:           sk,
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema
	k.Rounds = rnds
	k.RandomResults = rndRes

	return k
}

func (k *Keeper) GetLatestRandomRound(ctx sdk.Context) (*provider.LatestRandomRound, error) {
	k.logger.Info("fetching latest random value", "height", ctx.BlockHeight())
	lr, err := k.randProvider.GetLatestRandomness(ctx)

	if err != nil {
		k.logger.Error("error fetching latest random value from random provider", "err", err)
		return &provider.LatestRandomRound{}, err
	}

	return lr, nil
}

func (k *Keeper) VerifyVoteExtension(ctx sdk.Context, ve *oracletypes.VoteExtension) error {
	//TODO: Perform any additional checks on vote extensions here
	return nil
}

func (k Keeper) ValidateVoteExtension(ctx sdk.Context, height int64, chainId string, extCommitInfo abci.ExtendedCommitInfo) (bool, error) {
	if err := baseapp.ValidateVoteExtensions(
		ctx,
		k.sk,
		height,
		chainId,
		extCommitInfo,
	); err != nil {
		return false, err
	}
	return true, nil
}
