package module

import (
	"context"
	"cosmossdk.io/core/appmodule"
	"encoding/json"
	"fmt"
	oracletypes "github.com/glnro/random-oracle/x/oracle/types"

	abci "github.com/cometbft/cometbft/abci/types"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/glnro/random-oracle/x/oracle/keeper"
)

const ConsensusVersion = 1

var (
	_ module.AppModuleBasic = AppModule{}
	//_ module.HasGenesis     = AppModule{}
	_ module.HasServices = AppModule{}

	_ appmodule.AppModule = AppModule{}
)

type AppModule struct {
	cdc    codec.Codec
	keeper keeper.Keeper
}

func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
	return AppModule{
		cdc:    cdc,
		keeper: keeper,
	}
}

func NewAppModuleBasic(m AppModule) module.AppModuleBasic {
	return module.CoreAppModuleBasicAdaptor(m.Name(), m)
}

func (AppModule) Name() string { return oracletypes.ModuleName }

func (AppModule) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	oracletypes.RegisterLegacyAminoCodec(cdc)
}

func (AppModule) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *gwruntime.ServeMux) {
	if err := oracletypes.RegisterQueryHandlerClient(context.Background(), mux, oracletypes.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

func (AppModule) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	oracletypes.RegisterInterfaces(registry)
}

func (AppModule) ConsensusVersion() uint64 { return ConsensusVersion }

func (am AppModule) RegisterServices(cfg module.Configurator) {
	oracletypes.RegisterQueryServer(cfg.QueryServer(), keeper.NewQueryServerImpl(am.keeper))

	// Register in place module state migration migrations
	// m := keeper.NewMigrator(am.keeper)
	// if err := cfg.RegisterMigration(ns.ModuleName, 1, m.Migrate1to2); err != nil {
	// 	panic(fmt.Sprintf("failed to migrate x/%s from version 1 to 2: %v", ns.ModuleName, err))
	// }
}

func (AppModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(oracletypes.NewGenesisState())
}

func (AppModule) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var data oracletypes.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", oracletypes.ModuleName, err)
	}

	return data.Validate()
}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState oracletypes.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)

	if err := am.keeper.InitGenesis(ctx, &genesisState); err != nil {
		panic(fmt.Sprintf("failed to initialize %s genesis state: %v", oracletypes.ModuleName, err))
	}

	return nil
}

func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs, err := am.keeper.ExportGenesis(ctx)
	if err != nil {
		panic(fmt.Sprintf("failed to export %s genesis state: %v", oracletypes.ModuleName, err))
	}

	return cdc.MustMarshalJSON(gs)
}
