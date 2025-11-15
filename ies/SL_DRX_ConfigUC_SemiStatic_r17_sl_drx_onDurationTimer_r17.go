package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_onDurationTimer_r17_Choice_nothing uint64 = iota
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_onDurationTimer_r17_Choice_subMilliSeconds
	SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_onDurationTimer_r17_Choice_milliSeconds
)

type SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_onDurationTimer_r17 struct {
	Choice          uint64
	subMilliSeconds int64 `lb:1,ub:31,madatory`
	milliSeconds    *SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_onDurationTimer_r17_milliSeconds
}

func (ie *SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_onDurationTimer_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_onDurationTimer_r17_Choice_subMilliSeconds:
		if err = w.WriteInteger(int64(ie.subMilliSeconds), &uper.Constraint{Lb: 1, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode subMilliSeconds", err)
		}
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_onDurationTimer_r17_Choice_milliSeconds:
		if err = ie.milliSeconds.Encode(w); err != nil {
			err = utils.WrapError("Encode milliSeconds", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_onDurationTimer_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_onDurationTimer_r17_Choice_subMilliSeconds:
		var tmp_int_subMilliSeconds int64
		if tmp_int_subMilliSeconds, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode subMilliSeconds", err)
		}
		ie.subMilliSeconds = tmp_int_subMilliSeconds
	case SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_onDurationTimer_r17_Choice_milliSeconds:
		ie.milliSeconds = new(SL_DRX_ConfigUC_SemiStatic_r17_sl_drx_onDurationTimer_r17_milliSeconds)
		if err = ie.milliSeconds.Decode(r); err != nil {
			return utils.WrapError("Decode milliSeconds", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
