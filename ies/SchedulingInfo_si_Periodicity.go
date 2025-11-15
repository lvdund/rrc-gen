package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SchedulingInfo_si_Periodicity_Enum_rf8   uper.Enumerated = 0
	SchedulingInfo_si_Periodicity_Enum_rf16  uper.Enumerated = 1
	SchedulingInfo_si_Periodicity_Enum_rf32  uper.Enumerated = 2
	SchedulingInfo_si_Periodicity_Enum_rf64  uper.Enumerated = 3
	SchedulingInfo_si_Periodicity_Enum_rf128 uper.Enumerated = 4
	SchedulingInfo_si_Periodicity_Enum_rf256 uper.Enumerated = 5
	SchedulingInfo_si_Periodicity_Enum_rf512 uper.Enumerated = 6
)

type SchedulingInfo_si_Periodicity struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *SchedulingInfo_si_Periodicity) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode SchedulingInfo_si_Periodicity", err)
	}
	return nil
}

func (ie *SchedulingInfo_si_Periodicity) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode SchedulingInfo_si_Periodicity", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
