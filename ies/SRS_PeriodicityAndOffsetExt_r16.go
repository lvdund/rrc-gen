package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_PeriodicityAndOffsetExt_r16_Choice_nothing uint64 = iota
	SRS_PeriodicityAndOffsetExt_r16_Choice_sl128
	SRS_PeriodicityAndOffsetExt_r16_Choice_sl256
	SRS_PeriodicityAndOffsetExt_r16_Choice_sl512
	SRS_PeriodicityAndOffsetExt_r16_Choice_sl20480
)

type SRS_PeriodicityAndOffsetExt_r16 struct {
	Choice  uint64
	sl128   int64 `lb:0,ub:127,madatory`
	sl256   int64 `lb:0,ub:255,madatory`
	sl512   int64 `lb:0,ub:511,madatory`
	sl20480 int64 `lb:0,ub:20479,madatory`
}

func (ie *SRS_PeriodicityAndOffsetExt_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_PeriodicityAndOffsetExt_r16_Choice_sl128:
		if err = w.WriteInteger(int64(ie.sl128), &uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			err = utils.WrapError("Encode sl128", err)
		}
	case SRS_PeriodicityAndOffsetExt_r16_Choice_sl256:
		if err = w.WriteInteger(int64(ie.sl256), &uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			err = utils.WrapError("Encode sl256", err)
		}
	case SRS_PeriodicityAndOffsetExt_r16_Choice_sl512:
		if err = w.WriteInteger(int64(ie.sl512), &uper.Constraint{Lb: 0, Ub: 511}, false); err != nil {
			err = utils.WrapError("Encode sl512", err)
		}
	case SRS_PeriodicityAndOffsetExt_r16_Choice_sl20480:
		if err = w.WriteInteger(int64(ie.sl20480), &uper.Constraint{Lb: 0, Ub: 20479}, false); err != nil {
			err = utils.WrapError("Encode sl20480", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_PeriodicityAndOffsetExt_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_PeriodicityAndOffsetExt_r16_Choice_sl128:
		var tmp_int_sl128 int64
		if tmp_int_sl128, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode sl128", err)
		}
		ie.sl128 = tmp_int_sl128
	case SRS_PeriodicityAndOffsetExt_r16_Choice_sl256:
		var tmp_int_sl256 int64
		if tmp_int_sl256, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return utils.WrapError("Decode sl256", err)
		}
		ie.sl256 = tmp_int_sl256
	case SRS_PeriodicityAndOffsetExt_r16_Choice_sl512:
		var tmp_int_sl512 int64
		if tmp_int_sl512, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 511}, false); err != nil {
			return utils.WrapError("Decode sl512", err)
		}
		ie.sl512 = tmp_int_sl512
	case SRS_PeriodicityAndOffsetExt_r16_Choice_sl20480:
		var tmp_int_sl20480 int64
		if tmp_int_sl20480, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 20479}, false); err != nil {
			return utils.WrapError("Decode sl20480", err)
		}
		ie.sl20480 = tmp_int_sl20480
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
