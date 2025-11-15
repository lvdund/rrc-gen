package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CLI_PeriodicalReportConfig_r16 struct {
	reportInterval_r16    ReportInterval                                  `madatory`
	reportAmount_r16      CLI_PeriodicalReportConfig_r16_reportAmount_r16 `madatory`
	reportQuantityCLI_r16 MeasReportQuantityCLI_r16                       `madatory`
	maxReportCLI_r16      int64                                           `lb:1,ub:maxCLI_Report_r16,madatory`
}

func (ie *CLI_PeriodicalReportConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.reportInterval_r16.Encode(w); err != nil {
		return utils.WrapError("Encode reportInterval_r16", err)
	}
	if err = ie.reportAmount_r16.Encode(w); err != nil {
		return utils.WrapError("Encode reportAmount_r16", err)
	}
	if err = ie.reportQuantityCLI_r16.Encode(w); err != nil {
		return utils.WrapError("Encode reportQuantityCLI_r16", err)
	}
	if err = w.WriteInteger(ie.maxReportCLI_r16, &uper.Constraint{Lb: 1, Ub: maxCLI_Report_r16}, false); err != nil {
		return utils.WrapError("WriteInteger maxReportCLI_r16", err)
	}
	return nil
}

func (ie *CLI_PeriodicalReportConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.reportInterval_r16.Decode(r); err != nil {
		return utils.WrapError("Decode reportInterval_r16", err)
	}
	if err = ie.reportAmount_r16.Decode(r); err != nil {
		return utils.WrapError("Decode reportAmount_r16", err)
	}
	if err = ie.reportQuantityCLI_r16.Decode(r); err != nil {
		return utils.WrapError("Decode reportQuantityCLI_r16", err)
	}
	var tmp_int_maxReportCLI_r16 int64
	if tmp_int_maxReportCLI_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxCLI_Report_r16}, false); err != nil {
		return utils.WrapError("ReadInteger maxReportCLI_r16", err)
	}
	ie.maxReportCLI_r16 = tmp_int_maxReportCLI_r16
	return nil
}
