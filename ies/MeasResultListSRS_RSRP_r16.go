package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultListSRS_RSRP_r16 struct {
	Value []MeasResultSRS_RSRP_r16 `lb:1,ub:maxCLI_Report_r16,madatory`
}

func (ie *MeasResultListSRS_RSRP_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*MeasResultSRS_RSRP_r16]([]*MeasResultSRS_RSRP_r16{}, uper.Constraint{Lb: 1, Ub: maxCLI_Report_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultListSRS_RSRP_r16", err)
	}
	return nil
}

func (ie *MeasResultListSRS_RSRP_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*MeasResultSRS_RSRP_r16]([]*MeasResultSRS_RSRP_r16{}, uper.Constraint{Lb: 1, Ub: maxCLI_Report_r16}, false)
	fn := func() *MeasResultSRS_RSRP_r16 {
		return new(MeasResultSRS_RSRP_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode MeasResultListSRS_RSRP_r16", err)
	}
	ie.Value = []MeasResultSRS_RSRP_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
