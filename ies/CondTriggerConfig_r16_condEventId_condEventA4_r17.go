package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CondTriggerConfig_r16_condEventId_condEventA4_r17 struct {
	a4_Threshold_r17  MeasTriggerQuantity `madatory`
	hysteresis_r17    Hysteresis          `madatory`
	timeToTrigger_r17 TimeToTrigger       `madatory`
}

func (ie *CondTriggerConfig_r16_condEventId_condEventA4_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.a4_Threshold_r17.Encode(w); err != nil {
		return utils.WrapError("Encode a4_Threshold_r17", err)
	}
	if err = ie.hysteresis_r17.Encode(w); err != nil {
		return utils.WrapError("Encode hysteresis_r17", err)
	}
	if err = ie.timeToTrigger_r17.Encode(w); err != nil {
		return utils.WrapError("Encode timeToTrigger_r17", err)
	}
	return nil
}

func (ie *CondTriggerConfig_r16_condEventId_condEventA4_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.a4_Threshold_r17.Decode(r); err != nil {
		return utils.WrapError("Decode a4_Threshold_r17", err)
	}
	if err = ie.hysteresis_r17.Decode(r); err != nil {
		return utils.WrapError("Decode hysteresis_r17", err)
	}
	if err = ie.timeToTrigger_r17.Decode(r); err != nil {
		return utils.WrapError("Decode timeToTrigger_r17", err)
	}
	return nil
}
