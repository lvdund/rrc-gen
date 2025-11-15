package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SchedulingInfo2_r17_si_Periodicity_r17_Enum_rf8   uper.Enumerated = 0
	SchedulingInfo2_r17_si_Periodicity_r17_Enum_rf16  uper.Enumerated = 1
	SchedulingInfo2_r17_si_Periodicity_r17_Enum_rf32  uper.Enumerated = 2
	SchedulingInfo2_r17_si_Periodicity_r17_Enum_rf64  uper.Enumerated = 3
	SchedulingInfo2_r17_si_Periodicity_r17_Enum_rf128 uper.Enumerated = 4
	SchedulingInfo2_r17_si_Periodicity_r17_Enum_rf256 uper.Enumerated = 5
	SchedulingInfo2_r17_si_Periodicity_r17_Enum_rf512 uper.Enumerated = 6
)

type SchedulingInfo2_r17_si_Periodicity_r17 struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *SchedulingInfo2_r17_si_Periodicity_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode SchedulingInfo2_r17_si_Periodicity_r17", err)
	}
	return nil
}

func (ie *SchedulingInfo2_r17_si_Periodicity_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode SchedulingInfo2_r17_si_Periodicity_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
