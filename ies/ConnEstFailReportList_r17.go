package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConnEstFailReportList_r17 struct {
	Value []ConnEstFailReport_r16 `lb:1,ub:maxCEFReport_r17,madatory`
}

func (ie *ConnEstFailReportList_r17) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*ConnEstFailReport_r16]([]*ConnEstFailReport_r16{}, uper.Constraint{Lb: 1, Ub: maxCEFReport_r17}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode ConnEstFailReportList_r17", err)
	}
	return nil
}

func (ie *ConnEstFailReportList_r17) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*ConnEstFailReport_r16]([]*ConnEstFailReport_r16{}, uper.Constraint{Lb: 1, Ub: maxCEFReport_r17}, false)
	fn := func() *ConnEstFailReport_r16 {
		return new(ConnEstFailReport_r16)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode ConnEstFailReportList_r17", err)
	}
	ie.Value = []ConnEstFailReport_r16{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
