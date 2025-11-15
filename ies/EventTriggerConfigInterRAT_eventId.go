package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	EventTriggerConfigInterRAT_eventId_Choice_nothing uint64 = iota
	EventTriggerConfigInterRAT_eventId_Choice_eventB1
	EventTriggerConfigInterRAT_eventId_Choice_eventB2
	EventTriggerConfigInterRAT_eventId_Choice_eventB1_UTRA_FDD_r16
	EventTriggerConfigInterRAT_eventId_Choice_eventB2_UTRA_FDD_r16
	EventTriggerConfigInterRAT_eventId_Choice_eventY1_Relay_r17
	EventTriggerConfigInterRAT_eventId_Choice_eventY2_Relay_r17
)

type EventTriggerConfigInterRAT_eventId struct {
	Choice               uint64
	eventB1              *EventTriggerConfigInterRAT_eventId_eventB1
	eventB2              *EventTriggerConfigInterRAT_eventId_eventB2
	eventB1_UTRA_FDD_r16 *EventTriggerConfigInterRAT_eventId_eventB1_UTRA_FDD_r16
	eventB2_UTRA_FDD_r16 *EventTriggerConfigInterRAT_eventId_eventB2_UTRA_FDD_r16
	eventY1_Relay_r17    *EventTriggerConfigInterRAT_eventId_eventY1_Relay_r17
	eventY2_Relay_r17    *EventTriggerConfigInterRAT_eventId_eventY2_Relay_r17
}

func (ie *EventTriggerConfigInterRAT_eventId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 6, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EventTriggerConfigInterRAT_eventId_Choice_eventB1:
		if err = ie.eventB1.Encode(w); err != nil {
			err = utils.WrapError("Encode eventB1", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_eventB2:
		if err = ie.eventB2.Encode(w); err != nil {
			err = utils.WrapError("Encode eventB2", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_eventB1_UTRA_FDD_r16:
		if err = ie.eventB1_UTRA_FDD_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode eventB1_UTRA_FDD_r16", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_eventB2_UTRA_FDD_r16:
		if err = ie.eventB2_UTRA_FDD_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode eventB2_UTRA_FDD_r16", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_eventY1_Relay_r17:
		if err = ie.eventY1_Relay_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode eventY1_Relay_r17", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_eventY2_Relay_r17:
		if err = ie.eventY2_Relay_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode eventY2_Relay_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *EventTriggerConfigInterRAT_eventId) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(6, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EventTriggerConfigInterRAT_eventId_Choice_eventB1:
		ie.eventB1 = new(EventTriggerConfigInterRAT_eventId_eventB1)
		if err = ie.eventB1.Decode(r); err != nil {
			return utils.WrapError("Decode eventB1", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_eventB2:
		ie.eventB2 = new(EventTriggerConfigInterRAT_eventId_eventB2)
		if err = ie.eventB2.Decode(r); err != nil {
			return utils.WrapError("Decode eventB2", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_eventB1_UTRA_FDD_r16:
		ie.eventB1_UTRA_FDD_r16 = new(EventTriggerConfigInterRAT_eventId_eventB1_UTRA_FDD_r16)
		if err = ie.eventB1_UTRA_FDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode eventB1_UTRA_FDD_r16", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_eventB2_UTRA_FDD_r16:
		ie.eventB2_UTRA_FDD_r16 = new(EventTriggerConfigInterRAT_eventId_eventB2_UTRA_FDD_r16)
		if err = ie.eventB2_UTRA_FDD_r16.Decode(r); err != nil {
			return utils.WrapError("Decode eventB2_UTRA_FDD_r16", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_eventY1_Relay_r17:
		ie.eventY1_Relay_r17 = new(EventTriggerConfigInterRAT_eventId_eventY1_Relay_r17)
		if err = ie.eventY1_Relay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode eventY1_Relay_r17", err)
		}
	case EventTriggerConfigInterRAT_eventId_Choice_eventY2_Relay_r17:
		ie.eventY2_Relay_r17 = new(EventTriggerConfigInterRAT_eventId_eventY2_Relay_r17)
		if err = ie.eventY2_Relay_r17.Decode(r); err != nil {
			return utils.WrapError("Decode eventY2_Relay_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
