package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	EventTriggerConfig_eventId_Choice_nothing uint64 = iota
	EventTriggerConfig_eventId_Choice_eventA1
	EventTriggerConfig_eventId_Choice_eventA2
	EventTriggerConfig_eventId_Choice_eventA3
	EventTriggerConfig_eventId_Choice_eventA4
	EventTriggerConfig_eventId_Choice_eventA5
	EventTriggerConfig_eventId_Choice_eventA6
	EventTriggerConfig_eventId_Choice_eventX1_r17
	EventTriggerConfig_eventId_Choice_eventX2_r17
	EventTriggerConfig_eventId_Choice_eventD1_r17
)

type EventTriggerConfig_eventId struct {
	Choice      uint64
	eventA1     *EventTriggerConfig_eventId_eventA1
	eventA2     *EventTriggerConfig_eventId_eventA2
	eventA3     *EventTriggerConfig_eventId_eventA3
	eventA4     *EventTriggerConfig_eventId_eventA4
	eventA5     *EventTriggerConfig_eventId_eventA5
	eventA6     *EventTriggerConfig_eventId_eventA6
	eventX1_r17 *EventTriggerConfig_eventId_eventX1_r17
	eventX2_r17 *EventTriggerConfig_eventId_eventX2_r17
	eventD1_r17 *EventTriggerConfig_eventId_eventD1_r17
}

func (ie *EventTriggerConfig_eventId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 9, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EventTriggerConfig_eventId_Choice_eventA1:
		if err = ie.eventA1.Encode(w); err != nil {
			err = utils.WrapError("Encode eventA1", err)
		}
	case EventTriggerConfig_eventId_Choice_eventA2:
		if err = ie.eventA2.Encode(w); err != nil {
			err = utils.WrapError("Encode eventA2", err)
		}
	case EventTriggerConfig_eventId_Choice_eventA3:
		if err = ie.eventA3.Encode(w); err != nil {
			err = utils.WrapError("Encode eventA3", err)
		}
	case EventTriggerConfig_eventId_Choice_eventA4:
		if err = ie.eventA4.Encode(w); err != nil {
			err = utils.WrapError("Encode eventA4", err)
		}
	case EventTriggerConfig_eventId_Choice_eventA5:
		if err = ie.eventA5.Encode(w); err != nil {
			err = utils.WrapError("Encode eventA5", err)
		}
	case EventTriggerConfig_eventId_Choice_eventA6:
		if err = ie.eventA6.Encode(w); err != nil {
			err = utils.WrapError("Encode eventA6", err)
		}
	case EventTriggerConfig_eventId_Choice_eventX1_r17:
		if err = ie.eventX1_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode eventX1_r17", err)
		}
	case EventTriggerConfig_eventId_Choice_eventX2_r17:
		if err = ie.eventX2_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode eventX2_r17", err)
		}
	case EventTriggerConfig_eventId_Choice_eventD1_r17:
		if err = ie.eventD1_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode eventD1_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *EventTriggerConfig_eventId) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(9, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EventTriggerConfig_eventId_Choice_eventA1:
		ie.eventA1 = new(EventTriggerConfig_eventId_eventA1)
		if err = ie.eventA1.Decode(r); err != nil {
			return utils.WrapError("Decode eventA1", err)
		}
	case EventTriggerConfig_eventId_Choice_eventA2:
		ie.eventA2 = new(EventTriggerConfig_eventId_eventA2)
		if err = ie.eventA2.Decode(r); err != nil {
			return utils.WrapError("Decode eventA2", err)
		}
	case EventTriggerConfig_eventId_Choice_eventA3:
		ie.eventA3 = new(EventTriggerConfig_eventId_eventA3)
		if err = ie.eventA3.Decode(r); err != nil {
			return utils.WrapError("Decode eventA3", err)
		}
	case EventTriggerConfig_eventId_Choice_eventA4:
		ie.eventA4 = new(EventTriggerConfig_eventId_eventA4)
		if err = ie.eventA4.Decode(r); err != nil {
			return utils.WrapError("Decode eventA4", err)
		}
	case EventTriggerConfig_eventId_Choice_eventA5:
		ie.eventA5 = new(EventTriggerConfig_eventId_eventA5)
		if err = ie.eventA5.Decode(r); err != nil {
			return utils.WrapError("Decode eventA5", err)
		}
	case EventTriggerConfig_eventId_Choice_eventA6:
		ie.eventA6 = new(EventTriggerConfig_eventId_eventA6)
		if err = ie.eventA6.Decode(r); err != nil {
			return utils.WrapError("Decode eventA6", err)
		}
	case EventTriggerConfig_eventId_Choice_eventX1_r17:
		ie.eventX1_r17 = new(EventTriggerConfig_eventId_eventX1_r17)
		if err = ie.eventX1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode eventX1_r17", err)
		}
	case EventTriggerConfig_eventId_Choice_eventX2_r17:
		ie.eventX2_r17 = new(EventTriggerConfig_eventId_eventX2_r17)
		if err = ie.eventX2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode eventX2_r17", err)
		}
	case EventTriggerConfig_eventId_Choice_eventD1_r17:
		ie.eventD1_r17 = new(EventTriggerConfig_eventId_eventD1_r17)
		if err = ie.eventD1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode eventD1_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
