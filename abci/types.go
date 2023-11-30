package abci

import (
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/glnro/random-oracle/provider"
)

type ProposalHandler struct {
	logger   log.Logger
	txConfig client.TxConfig
}

type VoteExtHandler struct {
	logger       log.Logger
	currentBlock int64
	randProvider *provider.RandProvider
}
