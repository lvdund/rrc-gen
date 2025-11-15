package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_ms192  uper.Enumerated = 0
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_ms256  uper.Enumerated = 1
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_ms320  uper.Enumerated = 2
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_ms384  uper.Enumerated = 3
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_ms448  uper.Enumerated = 4
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_ms512  uper.Enumerated = 5
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_ms576  uper.Enumerated = 6
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_ms640  uper.Enumerated = 7
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_ms1082 uper.Enumerated = 8
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_spare7 uper.Enumerated = 9
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_spare6 uper.Enumerated = 10
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_spare5 uper.Enumerated = 11
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_spare4 uper.Enumerated = 12
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_spare3 uper.Enumerated = 13
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_spare2 uper.Enumerated = 14
	SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700_Enum_spare1 uper.Enumerated = 15
)

type SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700", err)
	}
	return nil
}

func (ie *SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SchedulingRequestToAddModExt_v1700_sr_ProhibitTimer_v1700", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
