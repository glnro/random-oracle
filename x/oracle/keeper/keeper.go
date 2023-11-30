package keeper

import (
	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	oracletypes "github.com/glnro/random-oracle/x/oracle/types"
)

type Keeper struct {
	cdc    codec.BinaryCodec
	logger log.Logger

	// state management
	Schema        collections.Schema
	Rounds        collections.Sequence
	RandomResults collections.Map[uint64, oracletypes.RandomResult]
}

func NewKeeper(cdc codec.BinaryCodec, storeService storetypes.KVStoreService, logger log.Logger) Keeper {
	sb := collections.NewSchemaBuilder(storeService)
	rnds := collections.NewSequence(sb, oracletypes.RoundKey, "rounds")
	rndRes := collections.NewMap(sb, oracletypes.RandomResultsKey, "randomres", collections.Uint64Key, codec.CollValue[oracletypes.RandomResult](cdc))

	k := Keeper{
		cdc:    cdc,
		logger: logger,
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
