package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RLC_ChannelID_r17 struct {
	Value uint64 `lb:1,ub:maxSL_LCID_r16,madatory`
}

func (ie *SL_RLC_ChannelID_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false); err != nil {
		return utils.WrapError("Encode SL_RLC_ChannelID_r17", err)
	}
	return nil
}

func (ie *SL_RLC_ChannelID_r17) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false); err != nil {
		return utils.WrapError("Decode SL_RLC_ChannelID_r17", err)
	}
	ie.Value = uint64(v)
	return nil
}
