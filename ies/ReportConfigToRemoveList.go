package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReportConfigToRemoveList struct {
	Value []ReportConfigId `lb:1,ub:maxReportConfigId,madatory`
}

func (ie *ReportConfigToRemoveList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*ReportConfigId]([]*ReportConfigId{}, uper.Constraint{Lb: 1, Ub: maxReportConfigId}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode ReportConfigToRemoveList", err)
	}
	return nil
}

func (ie *ReportConfigToRemoveList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*ReportConfigId]([]*ReportConfigId{}, uper.Constraint{Lb: 1, Ub: maxReportConfigId}, false)
	fn := func() *ReportConfigId {
		return new(ReportConfigId)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode ReportConfigToRemoveList", err)
	}
	ie.Value = []ReportConfigId{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
