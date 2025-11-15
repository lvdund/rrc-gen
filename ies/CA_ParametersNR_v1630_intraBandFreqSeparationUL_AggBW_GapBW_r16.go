package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CA_ParametersNR_v1630_intraBandFreqSeparationUL_AggBW_GapBW_r16_Enum_classI   uper.Enumerated = 0
	CA_ParametersNR_v1630_intraBandFreqSeparationUL_AggBW_GapBW_r16_Enum_classII  uper.Enumerated = 1
	CA_ParametersNR_v1630_intraBandFreqSeparationUL_AggBW_GapBW_r16_Enum_classIII uper.Enumerated = 2
)

type CA_ParametersNR_v1630_intraBandFreqSeparationUL_AggBW_GapBW_r16 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *CA_ParametersNR_v1630_intraBandFreqSeparationUL_AggBW_GapBW_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode CA_ParametersNR_v1630_intraBandFreqSeparationUL_AggBW_GapBW_r16", err)
	}
	return nil
}

func (ie *CA_ParametersNR_v1630_intraBandFreqSeparationUL_AggBW_GapBW_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode CA_ParametersNR_v1630_intraBandFreqSeparationUL_AggBW_GapBW_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
