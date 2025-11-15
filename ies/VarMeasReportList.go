package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarMeasReportList struct {
	Value []VarMeasReport `lb:1,ub:maxNrofMeasId,madatory`
}

func (ie *VarMeasReportList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*VarMeasReport]([]*VarMeasReport{}, uper.Constraint{Lb: 1, Ub: maxNrofMeasId}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode VarMeasReportList", err)
	}
	return nil
}

func (ie *VarMeasReportList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*VarMeasReport]([]*VarMeasReport{}, uper.Constraint{Lb: 1, Ub: maxNrofMeasId}, false)
	fn := func() *VarMeasReport {
		return new(VarMeasReport)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode VarMeasReportList", err)
	}
	ie.Value = []VarMeasReport{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
