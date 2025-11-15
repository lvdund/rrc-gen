package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReportConfigToAddMod struct {
	reportConfigId ReportConfigId                    `madatory`
	reportConfig   ReportConfigToAddMod_reportConfig `madatory`
}

func (ie *ReportConfigToAddMod) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.reportConfigId.Encode(w); err != nil {
		return utils.WrapError("Encode reportConfigId", err)
	}
	if err = ie.reportConfig.Encode(w); err != nil {
		return utils.WrapError("Encode reportConfig", err)
	}
	return nil
}

func (ie *ReportConfigToAddMod) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.reportConfigId.Decode(r); err != nil {
		return utils.WrapError("Decode reportConfigId", err)
	}
	if err = ie.reportConfig.Decode(r); err != nil {
		return utils.WrapError("Decode reportConfig", err)
	}
	return nil
}
