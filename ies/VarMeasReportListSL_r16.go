package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarMeasReportListSL_r16 struct {
	Value []VarMeasReportSL_r16 `lb:1,ub:maxNrofSL_MeasId_r16,madatory`
}

func (ie *VarMeasReportListSL_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*VarMeasReportSL_r16]([]*VarMeasReportSL_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_MeasId_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode VarMeasReportListSL_r16", err)
	}
	return nil
}

func (ie *VarMeasReportListSL_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*VarMeasReportSL_r16]([]*VarMeasReportSL_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSL_MeasId_r16}, false)
	fn := func() *VarMeasReportSL_r16 {
		return new(VarMeasReportSL_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode VarMeasReportListSL_r16", err)
	}
	ie.Value = []VarMeasReportSL_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
