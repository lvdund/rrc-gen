package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarConnEstFailReportList_r17 struct {
	connEstFailReportList_r17 []VarConnEstFailReport_r16 `lb:1,ub:maxCEFReport_r17,madatory`
}

func (ie *VarConnEstFailReportList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp_connEstFailReportList_r17 := utils.NewSequence[*VarConnEstFailReport_r16]([]*VarConnEstFailReport_r16{}, uper.Constraint{Lb: 1, Ub: maxCEFReport_r17}, false)
	for _, i := range ie.connEstFailReportList_r17 {
		tmp_connEstFailReportList_r17.Value = append(tmp_connEstFailReportList_r17.Value, &i)
	}
	if err = tmp_connEstFailReportList_r17.Encode(w); err != nil {
		return utils.WrapError("Encode connEstFailReportList_r17", err)
	}
	return nil
}

func (ie *VarConnEstFailReportList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp_connEstFailReportList_r17 := utils.NewSequence[*VarConnEstFailReport_r16]([]*VarConnEstFailReport_r16{}, uper.Constraint{Lb: 1, Ub: maxCEFReport_r17}, false)
	fn_connEstFailReportList_r17 := func() *VarConnEstFailReport_r16 {
		return new(VarConnEstFailReport_r16)
	}
	if err = tmp_connEstFailReportList_r17.Decode(r, fn_connEstFailReportList_r17); err != nil {
		return utils.WrapError("Decode connEstFailReportList_r17", err)
	}
	ie.connEstFailReportList_r17 = []VarConnEstFailReport_r16{}
	for _, i := range tmp_connEstFailReportList_r17.Value {
		ie.connEstFailReportList_r17 = append(ie.connEstFailReportList_r17, *i)
	}
	return nil
}
