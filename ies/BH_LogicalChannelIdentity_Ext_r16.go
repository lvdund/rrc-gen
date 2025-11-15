package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BH_LogicalChannelIdentity_Ext_r16 struct {
	Value uint64 `lb:320,ub:maxLC_ID_Iab_r16,madatory`
}

func (ie *BH_LogicalChannelIdentity_Ext_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 320, Ub: maxLC_ID_Iab_r16}, false); err != nil {
		return utils.WrapError("Encode BH_LogicalChannelIdentity_Ext_r16", err)
	}
	return nil
}

func (ie *BH_LogicalChannelIdentity_Ext_r16) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 320, Ub: maxLC_ID_Iab_r16}, false); err != nil {
		return utils.WrapError("Decode BH_LogicalChannelIdentity_Ext_r16", err)
	}
	ie.Value = uint64(v)
	return nil
}
