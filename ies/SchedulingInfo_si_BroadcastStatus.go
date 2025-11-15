package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SchedulingInfo_si_BroadcastStatus_Enum_broadcasting    uper.Enumerated = 0
	SchedulingInfo_si_BroadcastStatus_Enum_notBroadcasting uper.Enumerated = 1
)

type SchedulingInfo_si_BroadcastStatus struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *SchedulingInfo_si_BroadcastStatus) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode SchedulingInfo_si_BroadcastStatus", err)
	}
	return nil
}

func (ie *SchedulingInfo_si_BroadcastStatus) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode SchedulingInfo_si_BroadcastStatus", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
