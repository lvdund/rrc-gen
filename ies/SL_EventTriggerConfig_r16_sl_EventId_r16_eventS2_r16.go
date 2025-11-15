package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_EventTriggerConfig_r16_sl_EventId_r16_eventS2_r16 struct {
	s2_Threshold_r16     SL_MeasTriggerQuantity_r16 `madatory`
	sl_ReportOnLeave_r16 bool                       `madatory`
	sl_Hysteresis_r16    Hysteresis                 `madatory`
	sl_TimeToTrigger_r16 TimeToTrigger              `madatory`
}

func (ie *SL_EventTriggerConfig_r16_sl_EventId_r16_eventS2_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.s2_Threshold_r16.Encode(w); err != nil {
		return utils.WrapError("Encode s2_Threshold_r16", err)
	}
	if err = w.WriteBoolean(ie.sl_ReportOnLeave_r16); err != nil {
		return utils.WrapError("WriteBoolean sl_ReportOnLeave_r16", err)
	}
	if err = ie.sl_Hysteresis_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_Hysteresis_r16", err)
	}
	if err = ie.sl_TimeToTrigger_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_TimeToTrigger_r16", err)
	}
	return nil
}

func (ie *SL_EventTriggerConfig_r16_sl_EventId_r16_eventS2_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.s2_Threshold_r16.Decode(r); err != nil {
		return utils.WrapError("Decode s2_Threshold_r16", err)
	}
	var tmp_bool_sl_ReportOnLeave_r16 bool
	if tmp_bool_sl_ReportOnLeave_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean sl_ReportOnLeave_r16", err)
	}
	ie.sl_ReportOnLeave_r16 = tmp_bool_sl_ReportOnLeave_r16
	if err = ie.sl_Hysteresis_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_Hysteresis_r16", err)
	}
	if err = ie.sl_TimeToTrigger_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_TimeToTrigger_r16", err)
	}
	return nil
}
