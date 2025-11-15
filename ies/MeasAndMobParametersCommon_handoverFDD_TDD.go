package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasAndMobParametersCommon_handoverFDD_TDD_Enum_supported uper.Enumerated = 0
)

type MeasAndMobParametersCommon_handoverFDD_TDD struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *MeasAndMobParametersCommon_handoverFDD_TDD) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersCommon_handoverFDD_TDD", err)
	}
	return nil
}

func (ie *MeasAndMobParametersCommon_handoverFDD_TDD) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersCommon_handoverFDD_TDD", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
