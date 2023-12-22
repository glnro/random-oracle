package types

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/gogoproto/proto"
)

type SideTransaction interface {
	proto.Message
}

var _ SideTransaction = &SideTransactionA{}
var _ SideTransaction = &SideTransactionB{}

func (m *MsgNewSideTx) SetSideTransaction(sidetx SideTransaction) error {
	st, err := types.NewAnyWithValue(sidetx)
	if err != nil {
		return err
	}
	m.SideTransaction = st
	return nil
}
