package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EventTriggerConfig_eventId_eventA1 struct {
	a1_Threshold  MeasTriggerQuantity `madatory`
	reportOnLeave bool                `madatory`
	hysteresis    Hysteresis          `madatory`
	timeToTrigger TimeToTrigger       `madatory`
}

func (ie *EventTriggerConfig_eventId_eventA1) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.a1_Threshold.Encode(w); err != nil {
		return utils.WrapError("Encode a1_Threshold", err)
	}
	if err = w.WriteBoolean(ie.reportOnLeave); err != nil {
		return utils.WrapError("WriteBoolean reportOnLeave", err)
	}
	if err = ie.hysteresis.Encode(w); err != nil {
		return utils.WrapError("Encode hysteresis", err)
	}
	if err = ie.timeToTrigger.Encode(w); err != nil {
		return utils.WrapError("Encode timeToTrigger", err)
	}
	return nil
}

func (ie *EventTriggerConfig_eventId_eventA1) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.a1_Threshold.Decode(r); err != nil {
		return utils.WrapError("Decode a1_Threshold", err)
	}
	var tmp_bool_reportOnLeave bool
	if tmp_bool_reportOnLeave, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean reportOnLeave", err)
	}
	ie.reportOnLeave = tmp_bool_reportOnLeave
	if err = ie.hysteresis.Decode(r); err != nil {
		return utils.WrapError("Decode hysteresis", err)
	}
	if err = ie.timeToTrigger.Decode(r); err != nil {
		return utils.WrapError("Decode timeToTrigger", err)
	}
	return nil
}
