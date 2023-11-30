package types

import "cosmossdk.io/collections"

const (
	ModuleName = "random_oracle"
	StoreKey   = "random_oracle"
)

var (
	RoundKey         = collections.NewPrefix(0)
	RandomResultsKey = collections.NewPrefix(1)
)
