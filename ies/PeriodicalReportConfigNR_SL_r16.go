package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PeriodicalReportConfigNR_SL_r16 struct {
	reportInterval_r16 ReportInterval                                   `madatory`
	reportAmount_r16   PeriodicalReportConfigNR_SL_r16_reportAmount_r16 `madatory`
	reportQuantity_r16 MeasReportQuantity_r16                           `madatory`
}

func (ie *PeriodicalReportConfigNR_SL_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.reportInterval_r16.Encode(w); err != nil {
		return utils.WrapError("Encode reportInterval_r16", err)
	}
	if err = ie.reportAmount_r16.Encode(w); err != nil {
		return utils.WrapError("Encode reportAmount_r16", err)
	}
	if err = ie.reportQuantity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode reportQuantity_r16", err)
	}
	return nil
}

func (ie *PeriodicalReportConfigNR_SL_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.reportInterval_r16.Decode(r); err != nil {
		return utils.WrapError("Decode reportInterval_r16", err)
	}
	if err = ie.reportAmount_r16.Decode(r); err != nil {
		return utils.WrapError("Decode reportAmount_r16", err)
	}
	if err = ie.reportQuantity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode reportQuantity_r16", err)
	}
	return nil
}
