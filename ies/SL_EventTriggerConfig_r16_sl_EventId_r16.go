package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_EventTriggerConfig_r16_sl_EventId_r16_Choice_nothing uint64 = iota
	SL_EventTriggerConfig_r16_sl_EventId_r16_Choice_eventS1_r16
	SL_EventTriggerConfig_r16_sl_EventId_r16_Choice_eventS2_r16
)

type SL_EventTriggerConfig_r16_sl_EventId_r16 struct {
	Choice      uint64
	eventS1_r16 *SL_EventTriggerConfig_r16_sl_EventId_r16_eventS1_r16
	eventS2_r16 *SL_EventTriggerConfig_r16_sl_EventId_r16_eventS2_r16
}

func (ie *SL_EventTriggerConfig_r16_sl_EventId_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_EventTriggerConfig_r16_sl_EventId_r16_Choice_eventS1_r16:
		if err = ie.eventS1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode eventS1_r16", err)
		}
	case SL_EventTriggerConfig_r16_sl_EventId_r16_Choice_eventS2_r16:
		if err = ie.eventS2_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode eventS2_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SL_EventTriggerConfig_r16_sl_EventId_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_EventTriggerConfig_r16_sl_EventId_r16_Choice_eventS1_r16:
		ie.eventS1_r16 = new(SL_EventTriggerConfig_r16_sl_EventId_r16_eventS1_r16)
		if err = ie.eventS1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode eventS1_r16", err)
		}
	case SL_EventTriggerConfig_r16_sl_EventId_r16_Choice_eventS2_r16:
		ie.eventS2_r16 = new(SL_EventTriggerConfig_r16_sl_EventId_r16_eventS2_r16)
		if err = ie.eventS2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode eventS2_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
