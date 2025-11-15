package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RepFactorAndTimeGap_r17_repetitionFactor_r17_Enum_n2     uper.Enumerated = 0
	RepFactorAndTimeGap_r17_repetitionFactor_r17_Enum_n4     uper.Enumerated = 1
	RepFactorAndTimeGap_r17_repetitionFactor_r17_Enum_n6     uper.Enumerated = 2
	RepFactorAndTimeGap_r17_repetitionFactor_r17_Enum_n8     uper.Enumerated = 3
	RepFactorAndTimeGap_r17_repetitionFactor_r17_Enum_n16    uper.Enumerated = 4
	RepFactorAndTimeGap_r17_repetitionFactor_r17_Enum_n32    uper.Enumerated = 5
	RepFactorAndTimeGap_r17_repetitionFactor_r17_Enum_spare2 uper.Enumerated = 6
	RepFactorAndTimeGap_r17_repetitionFactor_r17_Enum_spare1 uper.Enumerated = 7
)

type RepFactorAndTimeGap_r17_repetitionFactor_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RepFactorAndTimeGap_r17_repetitionFactor_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RepFactorAndTimeGap_r17_repetitionFactor_r17", err)
	}
	return nil
}

func (ie *RepFactorAndTimeGap_r17_repetitionFactor_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RepFactorAndTimeGap_r17_repetitionFactor_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
