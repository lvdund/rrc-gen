package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PosSchedulingInfo_r16_posSI_Periodicity_r16_Enum_rf8   uper.Enumerated = 0
	PosSchedulingInfo_r16_posSI_Periodicity_r16_Enum_rf16  uper.Enumerated = 1
	PosSchedulingInfo_r16_posSI_Periodicity_r16_Enum_rf32  uper.Enumerated = 2
	PosSchedulingInfo_r16_posSI_Periodicity_r16_Enum_rf64  uper.Enumerated = 3
	PosSchedulingInfo_r16_posSI_Periodicity_r16_Enum_rf128 uper.Enumerated = 4
	PosSchedulingInfo_r16_posSI_Periodicity_r16_Enum_rf256 uper.Enumerated = 5
	PosSchedulingInfo_r16_posSI_Periodicity_r16_Enum_rf512 uper.Enumerated = 6
)

type PosSchedulingInfo_r16_posSI_Periodicity_r16 struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *PosSchedulingInfo_r16_posSI_Periodicity_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode PosSchedulingInfo_r16_posSI_Periodicity_r16", err)
	}
	return nil
}

func (ie *PosSchedulingInfo_r16_posSI_Periodicity_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode PosSchedulingInfo_r16_posSI_Periodicity_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
