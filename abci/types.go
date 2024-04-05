package abci

import (
	"cosmossdk.io/log"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/glnro/random-oracle/provider"
	oabci "github.com/glnro/random-oracle/x/oracle/abci"
	oracletypes "github.com/glnro/random-oracle/x/oracle/types"
)

type (
	VoteExtension struct {
		LatestRand provider.LatestRandomRound
	}

	PrepareProposalHandler struct {
		logger                log.Logger
		ok                    oracletypes.OracleKeeper
		oraclePrepareProposal OraclePrepareProposal
	}

	ProcessProposalHandler struct {
		logger                log.Logger
		ok                    oracletypes.OracleKeeper
		oracleProcessProposal OracleProcessProposal
	}

	VoteExtHandler struct {
		logger           log.Logger
		ok               oracletypes.OracleKeeper
		oracleExtendVote OracleExtendVote
	}

	VerifyVoteExtHandler struct {
		logger           log.Logger
		ok               oracletypes.OracleKeeper
		oracleVerifyVote OracleVerifyVote
	}

	PreBlockHandler struct {
		logger         log.Logger
		ok             oracletypes.OracleKeeper
		oraclePreBlock OraclePreBlocker
		mm             *module.Manager
	}

	OraclePrepareProposal func(ctx sdk.Context, req *abci.RequestPrepareProposal, ok oracletypes.OracleKeeper, logger log.Logger) ([]byte, error)
	OracleProcessProposal func(ctx sdk.Context, req *abci.RequestProcessProposal, ok oracletypes.OracleKeeper, log log.Logger) (*abci.ResponseProcessProposal, error)
	OracleExtendVote      func(ctx sdk.Context, req *abci.RequestExtendVote, ok oracletypes.OracleKeeper, log log.Logger) (*abci.ResponseExtendVote, error)
	OracleVerifyVote      func(ctx sdk.Context, req *abci.RequestVerifyVoteExtension, ok oracletypes.OracleKeeper, log log.Logger) (*abci.ResponseVerifyVoteExtension, error)
	OraclePreBlocker      func(ctx sdk.Context, req *abci.RequestFinalizeBlock, ok oracletypes.OracleKeeper, log log.Logger) error
)

func NewPrepareProposalHandler(log log.Logger, cdc codec.BinaryCodec, k oracletypes.OracleKeeper) *PrepareProposalHandler {
	return &PrepareProposalHandler{
		log,
		k,
		oabci.PrepareProposal,
	}
}

func NewProcessProposalHandler(log log.Logger, k oracletypes.OracleKeeper) *ProcessProposalHandler {
	return &ProcessProposalHandler{
		log,
		k,
		oabci.ProcessProposal,
	}
}

func NewExtendVoteHandler(log log.Logger, k oracletypes.OracleKeeper) *VoteExtHandler {
	return &VoteExtHandler{
		log,
		k,
		oabci.ExtendVote,
	}
}

func NewVerifyVoteExtHandler(log log.Logger, k oracletypes.OracleKeeper) *VerifyVoteExtHandler {
	return &VerifyVoteExtHandler{
		log,
		k,
		oabci.VerifyVoteExt,
	}
}

func NewPreBlockHandler(log log.Logger, k oracletypes.OracleKeeper, mm *module.Manager) *PreBlockHandler {
	return &PreBlockHandler{
		log,
		k,
		oabci.PreBlocker,
		mm,
	}
}
