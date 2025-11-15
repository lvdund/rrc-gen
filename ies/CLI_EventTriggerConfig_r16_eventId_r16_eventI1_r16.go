package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CLI_EventTriggerConfig_r16_eventId_r16_eventI1_r16 struct {
	i1_Threshold_r16  MeasTriggerQuantityCLI_r16 `madatory`
	reportOnLeave_r16 bool                       `madatory`
	hysteresis_r16    Hysteresis                 `madatory`
	timeToTrigger_r16 TimeToTrigger              `madatory`
}

func (ie *CLI_EventTriggerConfig_r16_eventId_r16_eventI1_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.i1_Threshold_r16.Encode(w); err != nil {
		return utils.WrapError("Encode i1_Threshold_r16", err)
	}
	if err = w.WriteBoolean(ie.reportOnLeave_r16); err != nil {
		return utils.WrapError("WriteBoolean reportOnLeave_r16", err)
	}
	if err = ie.hysteresis_r16.Encode(w); err != nil {
		return utils.WrapError("Encode hysteresis_r16", err)
	}
	if err = ie.timeToTrigger_r16.Encode(w); err != nil {
		return utils.WrapError("Encode timeToTrigger_r16", err)
	}
	return nil
}

func (ie *CLI_EventTriggerConfig_r16_eventId_r16_eventI1_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.i1_Threshold_r16.Decode(r); err != nil {
		return utils.WrapError("Decode i1_Threshold_r16", err)
	}
	var tmp_bool_reportOnLeave_r16 bool
	if tmp_bool_reportOnLeave_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean reportOnLeave_r16", err)
	}
	ie.reportOnLeave_r16 = tmp_bool_reportOnLeave_r16
	if err = ie.hysteresis_r16.Decode(r); err != nil {
		return utils.WrapError("Decode hysteresis_r16", err)
	}
	if err = ie.timeToTrigger_r16.Decode(r); err != nil {
		return utils.WrapError("Decode timeToTrigger_r16", err)
	}
	return nil
}
