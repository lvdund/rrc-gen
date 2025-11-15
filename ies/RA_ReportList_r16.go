package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RA_ReportList_r16 struct {
	Value []RA_Report_r16 `lb:1,ub:maxRAReport_r16,madatory`
}

func (ie *RA_ReportList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*RA_Report_r16]([]*RA_Report_r16{}, uper.Constraint{Lb: 1, Ub: maxRAReport_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode RA_ReportList_r16", err)
	}
	return nil
}

func (ie *RA_ReportList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*RA_Report_r16]([]*RA_Report_r16{}, uper.Constraint{Lb: 1, Ub: maxRAReport_r16}, false)
	fn := func() *RA_Report_r16 {
		return new(RA_Report_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode RA_ReportList_r16", err)
	}
	ie.Value = []RA_Report_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
