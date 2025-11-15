package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_QoS_FlowIdentity_r16 struct {
	Value uint64 `lb:1,ub:maxNrofSL_QFIs_r16,madatory`
}

func (ie *SL_QoS_FlowIdentity_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: maxNrofSL_QFIs_r16}, false); err != nil {
		return utils.WrapError("Encode SL_QoS_FlowIdentity_r16", err)
	}
	return nil
}

func (ie *SL_QoS_FlowIdentity_r16) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofSL_QFIs_r16}, false); err != nil {
		return utils.WrapError("Decode SL_QoS_FlowIdentity_r16", err)
	}
	ie.Value = uint64(v)
	return nil
}
