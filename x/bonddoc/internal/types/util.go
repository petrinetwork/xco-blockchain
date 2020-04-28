package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
)

func NewMsgCreateBond(bondDoc BondDoc, bondDid sovrin.SovrinDid) MsgCreateBond {
	return MsgCreateBond{
		SignBytes: "",
		TxHash:    "",
		SenderDid: "",
		BondDid:   bondDid.Did,
		PubKey:    bondDid.VerifyKey,
		Data:      bondDoc,
	}
}

func NewMsgUpdateBondStatus(senderDid string, updateBondStatusDoc UpdateBondStatusDoc, bondDid sovrin.SovrinDid) MsgUpdateBondStatus {
	return MsgUpdateBondStatus{
		SignBytes: "",
		SenderDid: senderDid,
		BondDid:   bondDid.Did,
		Data:      updateBondStatusDoc,
	}
}

func CheckNotEmpty(value string, name string) (valid bool, err sdk.Error) {
	if len(value) == 0 {
		return false, sdk.ErrUnknownRequest(name + " is empty.")
	} else {
		return true, nil
	}
}
