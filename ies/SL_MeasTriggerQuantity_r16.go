package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_MeasTriggerQuantity_r16_Choice_nothing uint64 = iota
	SL_MeasTriggerQuantity_r16_Choice_sl_RSRP_r16
)

type SL_MeasTriggerQuantity_r16 struct {
	Choice      uint64
	sl_RSRP_r16 *RSRP_Range
}

func (ie *SL_MeasTriggerQuantity_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 1, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_MeasTriggerQuantity_r16_Choice_sl_RSRP_r16:
		if err = ie.sl_RSRP_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode sl_RSRP_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SL_MeasTriggerQuantity_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(1, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_MeasTriggerQuantity_r16_Choice_sl_RSRP_r16:
		ie.sl_RSRP_r16 = new(RSRP_Range)
		if err = ie.sl_RSRP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_RSRP_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
