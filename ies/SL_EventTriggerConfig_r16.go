package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_EventTriggerConfig_r16 struct {
	sl_EventId_r16        SL_EventTriggerConfig_r16_sl_EventId_r16      `madatory`
	sl_ReportInterval_r16 ReportInterval                                `madatory,ext`
	sl_ReportAmount_r16   SL_EventTriggerConfig_r16_sl_ReportAmount_r16 `madatory,ext`
	sl_ReportQuantity_r16 SL_MeasReportQuantity_r16                     `madatory,ext`
	sl_RS_Type_r16        SL_RS_Type_r16                                `madatory,ext`
}

func (ie *SL_EventTriggerConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.sl_EventId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_EventId_r16", err)
	}
	return nil
}

func (ie *SL_EventTriggerConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.sl_EventId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_EventId_r16", err)
	}
	return nil
}
