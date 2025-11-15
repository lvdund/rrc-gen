package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SchedulingRequestToAddMod_sr_TransMax_Enum_n4     uper.Enumerated = 0
	SchedulingRequestToAddMod_sr_TransMax_Enum_n8     uper.Enumerated = 1
	SchedulingRequestToAddMod_sr_TransMax_Enum_n16    uper.Enumerated = 2
	SchedulingRequestToAddMod_sr_TransMax_Enum_n32    uper.Enumerated = 3
	SchedulingRequestToAddMod_sr_TransMax_Enum_n64    uper.Enumerated = 4
	SchedulingRequestToAddMod_sr_TransMax_Enum_spare3 uper.Enumerated = 5
	SchedulingRequestToAddMod_sr_TransMax_Enum_spare2 uper.Enumerated = 6
	SchedulingRequestToAddMod_sr_TransMax_Enum_spare1 uper.Enumerated = 7
)

type SchedulingRequestToAddMod_sr_TransMax struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SchedulingRequestToAddMod_sr_TransMax) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SchedulingRequestToAddMod_sr_TransMax", err)
	}
	return nil
}

func (ie *SchedulingRequestToAddMod_sr_TransMax) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SchedulingRequestToAddMod_sr_TransMax", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
