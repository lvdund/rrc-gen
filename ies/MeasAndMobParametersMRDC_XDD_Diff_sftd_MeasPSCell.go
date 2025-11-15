package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasPSCell_Enum_supported uper.Enumerated = 0
)

type MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasPSCell struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasPSCell) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasPSCell", err)
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasPSCell) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersMRDC_XDD_Diff_sftd_MeasPSCell", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
