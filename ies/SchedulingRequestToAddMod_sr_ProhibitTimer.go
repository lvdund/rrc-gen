package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SchedulingRequestToAddMod_sr_ProhibitTimer_Enum_ms1   uper.Enumerated = 0
	SchedulingRequestToAddMod_sr_ProhibitTimer_Enum_ms2   uper.Enumerated = 1
	SchedulingRequestToAddMod_sr_ProhibitTimer_Enum_ms4   uper.Enumerated = 2
	SchedulingRequestToAddMod_sr_ProhibitTimer_Enum_ms8   uper.Enumerated = 3
	SchedulingRequestToAddMod_sr_ProhibitTimer_Enum_ms16  uper.Enumerated = 4
	SchedulingRequestToAddMod_sr_ProhibitTimer_Enum_ms32  uper.Enumerated = 5
	SchedulingRequestToAddMod_sr_ProhibitTimer_Enum_ms64  uper.Enumerated = 6
	SchedulingRequestToAddMod_sr_ProhibitTimer_Enum_ms128 uper.Enumerated = 7
)

type SchedulingRequestToAddMod_sr_ProhibitTimer struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SchedulingRequestToAddMod_sr_ProhibitTimer) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SchedulingRequestToAddMod_sr_ProhibitTimer", err)
	}
	return nil
}

func (ie *SchedulingRequestToAddMod_sr_ProhibitTimer) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SchedulingRequestToAddMod_sr_ProhibitTimer", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
