package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasAndMobParametersFRX_Diff_handoverLTE_EPC_Enum_supported uper.Enumerated = 0
)

type MeasAndMobParametersFRX_Diff_handoverLTE_EPC struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *MeasAndMobParametersFRX_Diff_handoverLTE_EPC) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersFRX_Diff_handoverLTE_EPC", err)
	}
	return nil
}

func (ie *MeasAndMobParametersFRX_Diff_handoverLTE_EPC) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersFRX_Diff_handoverLTE_EPC", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
