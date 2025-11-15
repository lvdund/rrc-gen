package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	EventType_r16_Choice_nothing uint64 = iota
	EventType_r16_Choice_outOfCoverage
	EventType_r16_Choice_eventL1
)

type EventType_r16 struct {
	Choice        uint64
	outOfCoverage uper.NULL `madatory`
	eventL1       *EventType_r16_eventL1
}

func (ie *EventType_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EventType_r16_Choice_outOfCoverage:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode outOfCoverage", err)
		}
	case EventType_r16_Choice_eventL1:
		if err = ie.eventL1.Encode(w); err != nil {
			err = utils.WrapError("Encode eventL1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *EventType_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case EventType_r16_Choice_outOfCoverage:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode outOfCoverage", err)
		}
	case EventType_r16_Choice_eventL1:
		ie.eventL1 = new(EventType_r16_eventL1)
		if err = ie.eventL1.Decode(r); err != nil {
			return utils.WrapError("Decode eventL1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
