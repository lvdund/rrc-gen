package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_TxPower_r16_Choice_nothing uint64 = iota
	SL_TxPower_r16_Choice_minusinfinity_r16
	SL_TxPower_r16_Choice_txPower_r16
)

type SL_TxPower_r16 struct {
	Choice            uint64
	minusinfinity_r16 uper.NULL `madatory`
	txPower_r16       int64     `lb:-30,ub:33,madatory`
}

func (ie *SL_TxPower_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_TxPower_r16_Choice_minusinfinity_r16:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode minusinfinity_r16", err)
		}
	case SL_TxPower_r16_Choice_txPower_r16:
		if err = w.WriteInteger(int64(ie.txPower_r16), &uper.Constraint{Lb: -30, Ub: 33}, false); err != nil {
			err = utils.WrapError("Encode txPower_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SL_TxPower_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SL_TxPower_r16_Choice_minusinfinity_r16:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode minusinfinity_r16", err)
		}
	case SL_TxPower_r16_Choice_txPower_r16:
		var tmp_int_txPower_r16 int64
		if tmp_int_txPower_r16, err = r.ReadInteger(&uper.Constraint{Lb: -30, Ub: 33}, false); err != nil {
			return utils.WrapError("Decode txPower_r16", err)
		}
		ie.txPower_r16 = tmp_int_txPower_r16
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
