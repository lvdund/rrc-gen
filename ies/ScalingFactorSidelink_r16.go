package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ScalingFactorSidelink_r16_Enum_f0p4  uper.Enumerated = 0
	ScalingFactorSidelink_r16_Enum_f0p75 uper.Enumerated = 1
	ScalingFactorSidelink_r16_Enum_f0p8  uper.Enumerated = 2
	ScalingFactorSidelink_r16_Enum_f1    uper.Enumerated = 3
)

type ScalingFactorSidelink_r16 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *ScalingFactorSidelink_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode ScalingFactorSidelink_r16", err)
	}
	return nil
}

func (ie *ScalingFactorSidelink_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode ScalingFactorSidelink_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
