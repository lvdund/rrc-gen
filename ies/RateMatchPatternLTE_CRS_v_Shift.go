package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RateMatchPatternLTE_CRS_v_Shift_Enum_n0 uper.Enumerated = 0
	RateMatchPatternLTE_CRS_v_Shift_Enum_n1 uper.Enumerated = 1
	RateMatchPatternLTE_CRS_v_Shift_Enum_n2 uper.Enumerated = 2
	RateMatchPatternLTE_CRS_v_Shift_Enum_n3 uper.Enumerated = 3
	RateMatchPatternLTE_CRS_v_Shift_Enum_n4 uper.Enumerated = 4
	RateMatchPatternLTE_CRS_v_Shift_Enum_n5 uper.Enumerated = 5
)

type RateMatchPatternLTE_CRS_v_Shift struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *RateMatchPatternLTE_CRS_v_Shift) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode RateMatchPatternLTE_CRS_v_Shift", err)
	}
	return nil
}

func (ie *RateMatchPatternLTE_CRS_v_Shift) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode RateMatchPatternLTE_CRS_v_Shift", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
