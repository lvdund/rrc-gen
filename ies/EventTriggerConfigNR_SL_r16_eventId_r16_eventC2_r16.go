package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EventTriggerConfigNR_SL_r16_eventId_r16_eventC2_r16 struct {
	c2_Threshold_r16  SL_CBR_r16    `madatory`
	hysteresis_r16    Hysteresis    `madatory`
	timeToTrigger_r16 TimeToTrigger `madatory`
}

func (ie *EventTriggerConfigNR_SL_r16_eventId_r16_eventC2_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.c2_Threshold_r16.Encode(w); err != nil {
		return utils.WrapError("Encode c2_Threshold_r16", err)
	}
	if err = ie.hysteresis_r16.Encode(w); err != nil {
		return utils.WrapError("Encode hysteresis_r16", err)
	}
	if err = ie.timeToTrigger_r16.Encode(w); err != nil {
		return utils.WrapError("Encode timeToTrigger_r16", err)
	}
	return nil
}

func (ie *EventTriggerConfigNR_SL_r16_eventId_r16_eventC2_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.c2_Threshold_r16.Decode(r); err != nil {
		return utils.WrapError("Decode c2_Threshold_r16", err)
	}
	if err = ie.hysteresis_r16.Decode(r); err != nil {
		return utils.WrapError("Decode hysteresis_r16", err)
	}
	if err = ie.timeToTrigger_r16.Decode(r); err != nil {
		return utils.WrapError("Decode timeToTrigger_r16", err)
	}
	return nil
}
