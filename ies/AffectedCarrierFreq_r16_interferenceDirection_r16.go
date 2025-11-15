package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	AffectedCarrierFreq_r16_interferenceDirection_r16_Enum_nr    uper.Enumerated = 0
	AffectedCarrierFreq_r16_interferenceDirection_r16_Enum_other uper.Enumerated = 1
	AffectedCarrierFreq_r16_interferenceDirection_r16_Enum_both  uper.Enumerated = 2
	AffectedCarrierFreq_r16_interferenceDirection_r16_Enum_spare uper.Enumerated = 3
)

type AffectedCarrierFreq_r16_interferenceDirection_r16 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *AffectedCarrierFreq_r16_interferenceDirection_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode AffectedCarrierFreq_r16_interferenceDirection_r16", err)
	}
	return nil
}

func (ie *AffectedCarrierFreq_r16_interferenceDirection_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode AffectedCarrierFreq_r16_interferenceDirection_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
