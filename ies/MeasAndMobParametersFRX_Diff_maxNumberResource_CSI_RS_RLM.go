package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM_Enum_n2 uper.Enumerated = 0
	MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM_Enum_n4 uper.Enumerated = 1
	MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM_Enum_n6 uper.Enumerated = 2
	MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM_Enum_n8 uper.Enumerated = 3
)

type MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM", err)
	}
	return nil
}

func (ie *MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
