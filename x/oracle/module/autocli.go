package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	oracle "github.com/glnro/random-oracle/api/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: oracle.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "RandomRound",
					Use:       "round",
					Short:     "Get latest random round",
				},
			},
		},
	}
}
