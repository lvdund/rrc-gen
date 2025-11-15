package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CondTriggerConfig_r16_condEventId_condEventA3 struct {
	a3_Offset     MeasTriggerQuantityOffset `madatory`
	hysteresis    Hysteresis                `madatory`
	timeToTrigger TimeToTrigger             `madatory`
}

func (ie *CondTriggerConfig_r16_condEventId_condEventA3) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.a3_Offset.Encode(w); err != nil {
		return utils.WrapError("Encode a3_Offset", err)
	}
	if err = ie.hysteresis.Encode(w); err != nil {
		return utils.WrapError("Encode hysteresis", err)
	}
	if err = ie.timeToTrigger.Encode(w); err != nil {
		return utils.WrapError("Encode timeToTrigger", err)
	}
	return nil
}

func (ie *CondTriggerConfig_r16_condEventId_condEventA3) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.a3_Offset.Decode(r); err != nil {
		return utils.WrapError("Decode a3_Offset", err)
	}
	if err = ie.hysteresis.Decode(r); err != nil {
		return utils.WrapError("Decode hysteresis", err)
	}
	if err = ie.timeToTrigger.Decode(r); err != nil {
		return utils.WrapError("Decode timeToTrigger", err)
	}
	return nil
}
