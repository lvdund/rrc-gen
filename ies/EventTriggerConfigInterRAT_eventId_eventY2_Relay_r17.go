package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EventTriggerConfigInterRAT_eventId_eventY2_Relay_r17 struct {
	y2_Threshold_Relay_r17 SL_MeasTriggerQuantity_r16 `madatory`
	reportOnLeave_r17      bool                       `madatory`
	hysteresis_r17         Hysteresis                 `madatory`
	timeToTrigger_r17      TimeToTrigger              `madatory`
}

func (ie *EventTriggerConfigInterRAT_eventId_eventY2_Relay_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.y2_Threshold_Relay_r17.Encode(w); err != nil {
		return utils.WrapError("Encode y2_Threshold_Relay_r17", err)
	}
	if err = w.WriteBoolean(ie.reportOnLeave_r17); err != nil {
		return utils.WrapError("WriteBoolean reportOnLeave_r17", err)
	}
	if err = ie.hysteresis_r17.Encode(w); err != nil {
		return utils.WrapError("Encode hysteresis_r17", err)
	}
	if err = ie.timeToTrigger_r17.Encode(w); err != nil {
		return utils.WrapError("Encode timeToTrigger_r17", err)
	}
	return nil
}

func (ie *EventTriggerConfigInterRAT_eventId_eventY2_Relay_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.y2_Threshold_Relay_r17.Decode(r); err != nil {
		return utils.WrapError("Decode y2_Threshold_Relay_r17", err)
	}
	var tmp_bool_reportOnLeave_r17 bool
	if tmp_bool_reportOnLeave_r17, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean reportOnLeave_r17", err)
	}
	ie.reportOnLeave_r17 = tmp_bool_reportOnLeave_r17
	if err = ie.hysteresis_r17.Decode(r); err != nil {
		return utils.WrapError("Decode hysteresis_r17", err)
	}
	if err = ie.timeToTrigger_r17.Decode(r); err != nil {
		return utils.WrapError("Decode timeToTrigger_r17", err)
	}
	return nil
}
