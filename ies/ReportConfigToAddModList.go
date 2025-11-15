package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReportConfigToAddModList struct {
	Value []ReportConfigToAddMod `lb:1,ub:maxReportConfigId,madatory`
}

func (ie *ReportConfigToAddModList) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*ReportConfigToAddMod]([]*ReportConfigToAddMod{}, uper.Constraint{Lb: 1, Ub: maxReportConfigId}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode ReportConfigToAddModList", err)
	}
	return nil
}

func (ie *ReportConfigToAddModList) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*ReportConfigToAddMod]([]*ReportConfigToAddMod{}, uper.Constraint{Lb: 1, Ub: maxReportConfigId}, false)
	fn := func() *ReportConfigToAddMod {
		return new(ReportConfigToAddMod)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode ReportConfigToAddModList", err)
	}
	ie.Value = []ReportConfigToAddMod{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
