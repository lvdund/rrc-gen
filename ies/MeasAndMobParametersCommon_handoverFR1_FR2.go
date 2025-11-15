package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasAndMobParametersCommon_handoverFR1_FR2_Enum_supported uper.Enumerated = 0
)

type MeasAndMobParametersCommon_handoverFR1_FR2 struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *MeasAndMobParametersCommon_handoverFR1_FR2) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersCommon_handoverFR1_FR2", err)
	}
	return nil
}

func (ie *MeasAndMobParametersCommon_handoverFR1_FR2) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersCommon_handoverFR1_FR2", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
