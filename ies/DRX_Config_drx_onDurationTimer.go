package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DRX_Config_drx_onDurationTimer_Choice_nothing uint64 = iota
	DRX_Config_drx_onDurationTimer_Choice_subMilliSeconds
	DRX_Config_drx_onDurationTimer_Choice_milliSeconds
)

type DRX_Config_drx_onDurationTimer struct {
	Choice          uint64
	subMilliSeconds int64 `lb:1,ub:31,madatory`
	milliSeconds    *DRX_Config_drx_onDurationTimer_milliSeconds
}

func (ie *DRX_Config_drx_onDurationTimer) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DRX_Config_drx_onDurationTimer_Choice_subMilliSeconds:
		if err = w.WriteInteger(int64(ie.subMilliSeconds), &uper.Constraint{Lb: 1, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode subMilliSeconds", err)
		}
	case DRX_Config_drx_onDurationTimer_Choice_milliSeconds:
		if err = ie.milliSeconds.Encode(w); err != nil {
			err = utils.WrapError("Encode milliSeconds", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DRX_Config_drx_onDurationTimer) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DRX_Config_drx_onDurationTimer_Choice_subMilliSeconds:
		var tmp_int_subMilliSeconds int64
		if tmp_int_subMilliSeconds, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode subMilliSeconds", err)
		}
		ie.subMilliSeconds = tmp_int_subMilliSeconds
	case DRX_Config_drx_onDurationTimer_Choice_milliSeconds:
		ie.milliSeconds = new(DRX_Config_drx_onDurationTimer_milliSeconds)
		if err = ie.milliSeconds.Decode(r); err != nil {
			return utils.WrapError("Decode milliSeconds", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
