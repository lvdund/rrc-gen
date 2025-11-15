package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	EventTriggerConfigNR_SL_r16_eventId_r16_Choice_nothing uint64 = iota
	EventTriggerConfigNR_SL_r16_eventId_r16_Choice_eventC1
	EventTriggerConfigNR_SL_r16_eventId_r16_Choice_eventC2_r16
)

type EventTriggerConfigNR_SL_r16_eventId_r16 struct {
	Choice      uint64
	eventC1     *EventTriggerConfigNR_SL_r16_eventId_r16_eventC1
	eventC2_r16 *EventTriggerConfigNR_SL_r16_eventId_r16_eventC2_r16
}

func (ie *EventTriggerConfigNR_SL_r16_eventId_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EventTriggerConfigNR_SL_r16_eventId_r16_Choice_eventC1:
		if err = ie.eventC1.Encode(w); err != nil {
			err = utils.WrapError("Encode eventC1", err)
		}
	case EventTriggerConfigNR_SL_r16_eventId_r16_Choice_eventC2_r16:
		if err = ie.eventC2_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode eventC2_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *EventTriggerConfigNR_SL_r16_eventId_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EventTriggerConfigNR_SL_r16_eventId_r16_Choice_eventC1:
		ie.eventC1 = new(EventTriggerConfigNR_SL_r16_eventId_r16_eventC1)
		if err = ie.eventC1.Decode(r); err != nil {
			return utils.WrapError("Decode eventC1", err)
		}
	case EventTriggerConfigNR_SL_r16_eventId_r16_Choice_eventC2_r16:
		ie.eventC2_r16 = new(EventTriggerConfigNR_SL_r16_eventId_r16_eventC2_r16)
		if err = ie.eventC2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode eventC2_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
