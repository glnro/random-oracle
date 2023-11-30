package types

import "cosmossdk.io/collections"

const (
	ModuleName = "oracle"
	StoreKey   = "oracle"
)

var (
	RoundKey         = collections.NewPrefix(0)
	RandomResultsKey = collections.NewPrefix(1)
)
