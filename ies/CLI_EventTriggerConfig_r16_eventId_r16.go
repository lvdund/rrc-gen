package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CLI_EventTriggerConfig_r16_eventId_r16_Choice_nothing uint64 = iota
	CLI_EventTriggerConfig_r16_eventId_r16_Choice_eventI1_r16
)

type CLI_EventTriggerConfig_r16_eventId_r16 struct {
	Choice      uint64
	eventI1_r16 *CLI_EventTriggerConfig_r16_eventId_r16_eventI1_r16
}

func (ie *CLI_EventTriggerConfig_r16_eventId_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CLI_EventTriggerConfig_r16_eventId_r16_Choice_eventI1_r16:
		if err = ie.eventI1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode eventI1_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CLI_EventTriggerConfig_r16_eventId_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CLI_EventTriggerConfig_r16_eventId_r16_Choice_eventI1_r16:
		ie.eventI1_r16 = new(CLI_EventTriggerConfig_r16_eventId_r16_eventI1_r16)
		if err = ie.eventI1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode eventI1_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
