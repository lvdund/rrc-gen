package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_RLC_Config_r16_Choice_nothing uint64 = iota
	SL_RLC_Config_r16_Choice_sl_AM_RLC_r16
	SL_RLC_Config_r16_Choice_sl_UM_RLC_r16
)

type SL_RLC_Config_r16 struct {
	Choice        uint64
	sl_AM_RLC_r16 *SL_RLC_Config_r16_sl_AM_RLC_r16
	sl_UM_RLC_r16 *SL_RLC_Config_r16_sl_UM_RLC_r16
}

func (ie *SL_RLC_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_RLC_Config_r16_Choice_sl_AM_RLC_r16:
		if err = ie.sl_AM_RLC_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode sl_AM_RLC_r16", err)
		}
	case SL_RLC_Config_r16_Choice_sl_UM_RLC_r16:
		if err = ie.sl_UM_RLC_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode sl_UM_RLC_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SL_RLC_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_RLC_Config_r16_Choice_sl_AM_RLC_r16:
		ie.sl_AM_RLC_r16 = new(SL_RLC_Config_r16_sl_AM_RLC_r16)
		if err = ie.sl_AM_RLC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_AM_RLC_r16", err)
		}
	case SL_RLC_Config_r16_Choice_sl_UM_RLC_r16:
		ie.sl_UM_RLC_r16 = new(SL_RLC_Config_r16_sl_UM_RLC_r16)
		if err = ie.sl_UM_RLC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_UM_RLC_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
