package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PosSchedulingInfo_r16_posSI_BroadcastStatus_r16_Enum_broadcasting    uper.Enumerated = 0
	PosSchedulingInfo_r16_posSI_BroadcastStatus_r16_Enum_notBroadcasting uper.Enumerated = 1
)

type PosSchedulingInfo_r16_posSI_BroadcastStatus_r16 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *PosSchedulingInfo_r16_posSI_BroadcastStatus_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PosSchedulingInfo_r16_posSI_BroadcastStatus_r16", err)
	}
	return nil
}

func (ie *PosSchedulingInfo_r16_posSI_BroadcastStatus_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PosSchedulingInfo_r16_posSI_BroadcastStatus_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
