package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_PeriodCG_r16_Choice_nothing uint64 = iota
	SL_PeriodCG_r16_Choice_sl_PeriodCG1_r16
	SL_PeriodCG_r16_Choice_sl_PeriodCG2_r16
)

type SL_PeriodCG_r16 struct {
	Choice           uint64
	sl_PeriodCG1_r16 *SL_PeriodCG_r16_sl_PeriodCG1_r16
	sl_PeriodCG2_r16 int64 `lb:1,ub:99,madatory`
}

func (ie *SL_PeriodCG_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_PeriodCG_r16_Choice_sl_PeriodCG1_r16:
		if err = ie.sl_PeriodCG1_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode sl_PeriodCG1_r16", err)
		}
	case SL_PeriodCG_r16_Choice_sl_PeriodCG2_r16:
		if err = w.WriteInteger(int64(ie.sl_PeriodCG2_r16), &uper.Constraint{Lb: 1, Ub: 99}, false); err != nil {
			err = utils.WrapError("Encode sl_PeriodCG2_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SL_PeriodCG_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_PeriodCG_r16_Choice_sl_PeriodCG1_r16:
		ie.sl_PeriodCG1_r16 = new(SL_PeriodCG_r16_sl_PeriodCG1_r16)
		if err = ie.sl_PeriodCG1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PeriodCG1_r16", err)
		}
	case SL_PeriodCG_r16_Choice_sl_PeriodCG2_r16:
		var tmp_int_sl_PeriodCG2_r16 int64
		if tmp_int_sl_PeriodCG2_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 99}, false); err != nil {
			return utils.WrapError("Decode sl_PeriodCG2_r16", err)
		}
		ie.sl_PeriodCG2_r16 = tmp_int_sl_PeriodCG2_r16
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
