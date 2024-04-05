package module

import (
	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	modulev1 "github.com/glnro/random-oracle/api/module/v1"
	"github.com/glnro/random-oracle/provider"
	"github.com/glnro/random-oracle/x/oracle/keeper"
	oracletypes "github.com/glnro/random-oracle/x/oracle/types"
)

var _ appmodule.AppModule = AppModule{}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	Cdc          codec.Codec
	StoreService store.KVStoreService
	Logger       log.Logger

	RandProvider provider.RandProvider
	sk           oracletypes.StakingKeeper

	Config *modulev1.Module
}

type ModuleOutputs struct {
	depinject.Out

	Module appmodule.AppModule
	Keeper keeper.Keeper
}

func ProvideModule(in ModuleInputs) ModuleOutputs {

	k := keeper.NewKeeper(in.Cdc, in.StoreService, &in.RandProvider, in.sk, in.Logger)
	m := NewAppModule(in.Cdc, k)

	return ModuleOutputs{Module: m, Keeper: k}
}
