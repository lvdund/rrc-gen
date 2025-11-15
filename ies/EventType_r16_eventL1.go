package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type EventType_r16_eventL1 struct {
	l1_Threshold  MeasTriggerQuantity `madatory`
	hysteresis    Hysteresis          `madatory`
	timeToTrigger TimeToTrigger       `madatory`
}

func (ie *EventType_r16_eventL1) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.l1_Threshold.Encode(w); err != nil {
		return utils.WrapError("Encode l1_Threshold", err)
	}
	if err = ie.hysteresis.Encode(w); err != nil {
		return utils.WrapError("Encode hysteresis", err)
	}
	if err = ie.timeToTrigger.Encode(w); err != nil {
		return utils.WrapError("Encode timeToTrigger", err)
	}
	return nil
}

func (ie *EventType_r16_eventL1) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.l1_Threshold.Decode(r); err != nil {
		return utils.WrapError("Decode l1_Threshold", err)
	}
	if err = ie.hysteresis.Decode(r); err != nil {
		return utils.WrapError("Decode hysteresis", err)
	}
	if err = ie.timeToTrigger.Decode(r); err != nil {
		return utils.WrapError("Decode timeToTrigger", err)
	}
	return nil
}
