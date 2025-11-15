package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh_DRX_Enum_supported uper.Enumerated = 0
)

type MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh_DRX struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh_DRX) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh_DRX", err)
	}
	return nil
}

func (ie *MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh_DRX) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode MeasAndMobParametersXDD_Diff_sftd_MeasNR_Neigh_DRX", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
