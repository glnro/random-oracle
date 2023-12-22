package test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/glnro/random-oracle/x/oracle/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInterface(t *testing.T) {
	bs := banktypes.NewMsgSend(sdk.AccAddress("sender"), sdk.AccAddress("recipient"), sdk.NewCoins(sdk.NewInt64Coin("stake", int64(5))))

	sideTxA := types.MsgNewSideTx{}
	err := sideTxA.SetSideTransaction(&types.SideTransactionA{
		Name: "A",
	})
	require.NoError(t, err)

	sideTxB := types.MsgNewSideTx{}
	err = sideTxB.SetSideTransaction(&types.SideTransactionB{
		Title: "B",
	})
	require.NoError(t, err)

	msgs := []sdk.Msg{
		bs,
		&sideTxA,
		&sideTxB,
	}

	for _, msg := range msgs {
		switch m := msg.(type) {
		case *types.MsgNewSideTx:
			switch m.GetSideTransaction().GetTypeUrl() {
			case sdk.MsgTypeURL(&types.SideTransactionA{}):
				require.Equal(t, msg, &sideTxA)
			case sdk.MsgTypeURL(&types.SideTransactionB{}):
				require.Equal(t, msg, &sideTxB)
			}
		}
	}
}
