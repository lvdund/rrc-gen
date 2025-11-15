package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	EventTriggerConfigNR_SL_r16_reportAmount_r16_Enum_r1       uper.Enumerated = 0
	EventTriggerConfigNR_SL_r16_reportAmount_r16_Enum_r2       uper.Enumerated = 1
	EventTriggerConfigNR_SL_r16_reportAmount_r16_Enum_r4       uper.Enumerated = 2
	EventTriggerConfigNR_SL_r16_reportAmount_r16_Enum_r8       uper.Enumerated = 3
	EventTriggerConfigNR_SL_r16_reportAmount_r16_Enum_r16      uper.Enumerated = 4
	EventTriggerConfigNR_SL_r16_reportAmount_r16_Enum_r32      uper.Enumerated = 5
	EventTriggerConfigNR_SL_r16_reportAmount_r16_Enum_r64      uper.Enumerated = 6
	EventTriggerConfigNR_SL_r16_reportAmount_r16_Enum_infinity uper.Enumerated = 7
)

type EventTriggerConfigNR_SL_r16_reportAmount_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *EventTriggerConfigNR_SL_r16_reportAmount_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode EventTriggerConfigNR_SL_r16_reportAmount_r16", err)
	}
	return nil
}

func (ie *EventTriggerConfigNR_SL_r16_reportAmount_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode EventTriggerConfigNR_SL_r16_reportAmount_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
