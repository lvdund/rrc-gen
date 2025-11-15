package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CondTriggerConfig_r16_condEventId_Choice_nothing uint64 = iota
	CondTriggerConfig_r16_condEventId_Choice_condEventA3
	CondTriggerConfig_r16_condEventId_Choice_condEventA5
	CondTriggerConfig_r16_condEventId_Choice_condEventA4_r17
	CondTriggerConfig_r16_condEventId_Choice_condEventD1_r17
	CondTriggerConfig_r16_condEventId_Choice_condEventT1_r17
)

type CondTriggerConfig_r16_condEventId struct {
	Choice          uint64
	condEventA3     *CondTriggerConfig_r16_condEventId_condEventA3
	condEventA5     *CondTriggerConfig_r16_condEventId_condEventA5
	condEventA4_r17 *CondTriggerConfig_r16_condEventId_condEventA4_r17
	condEventD1_r17 *CondTriggerConfig_r16_condEventId_condEventD1_r17
	condEventT1_r17 *CondTriggerConfig_r16_condEventId_condEventT1_r17
}

func (ie *CondTriggerConfig_r16_condEventId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 5, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CondTriggerConfig_r16_condEventId_Choice_condEventA3:
		if err = ie.condEventA3.Encode(w); err != nil {
			err = utils.WrapError("Encode condEventA3", err)
		}
	case CondTriggerConfig_r16_condEventId_Choice_condEventA5:
		if err = ie.condEventA5.Encode(w); err != nil {
			err = utils.WrapError("Encode condEventA5", err)
		}
	case CondTriggerConfig_r16_condEventId_Choice_condEventA4_r17:
		if err = ie.condEventA4_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode condEventA4_r17", err)
		}
	case CondTriggerConfig_r16_condEventId_Choice_condEventD1_r17:
		if err = ie.condEventD1_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode condEventD1_r17", err)
		}
	case CondTriggerConfig_r16_condEventId_Choice_condEventT1_r17:
		if err = ie.condEventT1_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode condEventT1_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CondTriggerConfig_r16_condEventId) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(5, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CondTriggerConfig_r16_condEventId_Choice_condEventA3:
		ie.condEventA3 = new(CondTriggerConfig_r16_condEventId_condEventA3)
		if err = ie.condEventA3.Decode(r); err != nil {
			return utils.WrapError("Decode condEventA3", err)
		}
	case CondTriggerConfig_r16_condEventId_Choice_condEventA5:
		ie.condEventA5 = new(CondTriggerConfig_r16_condEventId_condEventA5)
		if err = ie.condEventA5.Decode(r); err != nil {
			return utils.WrapError("Decode condEventA5", err)
		}
	case CondTriggerConfig_r16_condEventId_Choice_condEventA4_r17:
		ie.condEventA4_r17 = new(CondTriggerConfig_r16_condEventId_condEventA4_r17)
		if err = ie.condEventA4_r17.Decode(r); err != nil {
			return utils.WrapError("Decode condEventA4_r17", err)
		}
	case CondTriggerConfig_r16_condEventId_Choice_condEventD1_r17:
		ie.condEventD1_r17 = new(CondTriggerConfig_r16_condEventId_condEventD1_r17)
		if err = ie.condEventD1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode condEventD1_r17", err)
		}
	case CondTriggerConfig_r16_condEventId_Choice_condEventT1_r17:
		ie.condEventT1_r17 = new(CondTriggerConfig_r16_condEventId_condEventT1_r17)
		if err = ie.condEventT1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode condEventT1_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
