package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17_Enum_ms120  uper.Enumerated = 0
	RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17_Enum_ms240  uper.Enumerated = 1
	RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17_Enum_ms480  uper.Enumerated = 2
	RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17_Enum_ms640  uper.Enumerated = 3
	RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17_Enum_ms1024 uper.Enumerated = 4
)

type RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17 struct {
	Value uper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17", err)
	}
	return nil
}

func (ie *RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode RAN_VisibleParameters_r17_ran_VisiblePeriodicity_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
