package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PeriodicalReportConfig_r16 struct {
	sl_ReportInterval_r16 ReportInterval                                    `madatory`
	sl_ReportAmount_r16   SL_PeriodicalReportConfig_r16_sl_ReportAmount_r16 `madatory`
	sl_ReportQuantity_r16 SL_MeasReportQuantity_r16                         `madatory`
	sl_RS_Type_r16        SL_RS_Type_r16                                    `madatory`
}

func (ie *SL_PeriodicalReportConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.sl_ReportInterval_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_ReportInterval_r16", err)
	}
	if err = ie.sl_ReportAmount_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_ReportAmount_r16", err)
	}
	if err = ie.sl_ReportQuantity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_ReportQuantity_r16", err)
	}
	if err = ie.sl_RS_Type_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_RS_Type_r16", err)
	}
	return nil
}

func (ie *SL_PeriodicalReportConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.sl_ReportInterval_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_ReportInterval_r16", err)
	}
	if err = ie.sl_ReportAmount_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_ReportAmount_r16", err)
	}
	if err = ie.sl_ReportQuantity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_ReportQuantity_r16", err)
	}
	if err = ie.sl_RS_Type_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_RS_Type_r16", err)
	}
	return nil
}
