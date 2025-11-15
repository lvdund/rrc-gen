package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NeedForGapsNR_r16_gapIndication_r16_Enum_gap    uper.Enumerated = 0
	NeedForGapsNR_r16_gapIndication_r16_Enum_no_gap uper.Enumerated = 1
)

type NeedForGapsNR_r16_gapIndication_r16 struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *NeedForGapsNR_r16_gapIndication_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode NeedForGapsNR_r16_gapIndication_r16", err)
	}
	return nil
}

func (ie *NeedForGapsNR_r16_gapIndication_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode NeedForGapsNR_r16_gapIndication_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
