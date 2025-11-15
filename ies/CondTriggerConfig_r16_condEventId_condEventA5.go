package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CondTriggerConfig_r16_condEventId_condEventA5 struct {
	a5_Threshold1 MeasTriggerQuantity `madatory`
	a5_Threshold2 MeasTriggerQuantity `madatory`
	hysteresis    Hysteresis          `madatory`
	timeToTrigger TimeToTrigger       `madatory`
}

func (ie *CondTriggerConfig_r16_condEventId_condEventA5) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.a5_Threshold1.Encode(w); err != nil {
		return utils.WrapError("Encode a5_Threshold1", err)
	}
	if err = ie.a5_Threshold2.Encode(w); err != nil {
		return utils.WrapError("Encode a5_Threshold2", err)
	}
	if err = ie.hysteresis.Encode(w); err != nil {
		return utils.WrapError("Encode hysteresis", err)
	}
	if err = ie.timeToTrigger.Encode(w); err != nil {
		return utils.WrapError("Encode timeToTrigger", err)
	}
	return nil
}

func (ie *CondTriggerConfig_r16_condEventId_condEventA5) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.a5_Threshold1.Decode(r); err != nil {
		return utils.WrapError("Decode a5_Threshold1", err)
	}
	if err = ie.a5_Threshold2.Decode(r); err != nil {
		return utils.WrapError("Decode a5_Threshold2", err)
	}
	if err = ie.hysteresis.Decode(r); err != nil {
		return utils.WrapError("Decode hysteresis", err)
	}
	if err = ie.timeToTrigger.Decode(r); err != nil {
		return utils.WrapError("Decode timeToTrigger", err)
	}
	return nil
}
