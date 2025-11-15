package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LogicalChannelIdentity struct {
	Value uint64 `lb:1,ub:maxLC_ID,madatory`
}

func (ie *LogicalChannelIdentity) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: maxLC_ID}, false); err != nil {
		return utils.WrapError("Encode LogicalChannelIdentity", err)
	}
	return nil
}

func (ie *LogicalChannelIdentity) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxLC_ID}, false); err != nil {
		return utils.WrapError("Decode LogicalChannelIdentity", err)
	}
	ie.Value = uint64(v)
	return nil
}
