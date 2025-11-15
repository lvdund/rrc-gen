package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MCCH_RepetitionPeriodAndOffset_r17_Choice_nothing uint64 = iota
	MCCH_RepetitionPeriodAndOffset_r17_Choice_rf1_r17
	MCCH_RepetitionPeriodAndOffset_r17_Choice_rf2_r17
	MCCH_RepetitionPeriodAndOffset_r17_Choice_rf4_r17
	MCCH_RepetitionPeriodAndOffset_r17_Choice_rf8_r17
	MCCH_RepetitionPeriodAndOffset_r17_Choice_rf16_r17
	MCCH_RepetitionPeriodAndOffset_r17_Choice_rf32_r17
	MCCH_RepetitionPeriodAndOffset_r17_Choice_rf64_r17
	MCCH_RepetitionPeriodAndOffset_r17_Choice_rf128_r17
	MCCH_RepetitionPeriodAndOffset_r17_Choice_rf256_r17
)

type MCCH_RepetitionPeriodAndOffset_r17 struct {
	Choice    uint64
	rf1_r17   int64 `lb:0,ub:0,madatory`
	rf2_r17   int64 `lb:0,ub:1,madatory`
	rf4_r17   int64 `lb:0,ub:3,madatory`
	rf8_r17   int64 `lb:0,ub:7,madatory`
	rf16_r17  int64 `lb:0,ub:15,madatory`
	rf32_r17  int64 `lb:0,ub:31,madatory`
	rf64_r17  int64 `lb:0,ub:63,madatory`
	rf128_r17 int64 `lb:0,ub:127,madatory`
	rf256_r17 int64 `lb:0,ub:255,madatory`
}

func (ie *MCCH_RepetitionPeriodAndOffset_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 9, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_rf1_r17:
		if err = w.WriteInteger(int64(ie.rf1_r17), &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			err = utils.WrapError("Encode rf1_r17", err)
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_rf2_r17:
		if err = w.WriteInteger(int64(ie.rf2_r17), &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			err = utils.WrapError("Encode rf2_r17", err)
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_rf4_r17:
		if err = w.WriteInteger(int64(ie.rf4_r17), &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode rf4_r17", err)
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_rf8_r17:
		if err = w.WriteInteger(int64(ie.rf8_r17), &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			err = utils.WrapError("Encode rf8_r17", err)
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_rf16_r17:
		if err = w.WriteInteger(int64(ie.rf16_r17), &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			err = utils.WrapError("Encode rf16_r17", err)
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_rf32_r17:
		if err = w.WriteInteger(int64(ie.rf32_r17), &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			err = utils.WrapError("Encode rf32_r17", err)
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_rf64_r17:
		if err = w.WriteInteger(int64(ie.rf64_r17), &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			err = utils.WrapError("Encode rf64_r17", err)
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_rf128_r17:
		if err = w.WriteInteger(int64(ie.rf128_r17), &uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			err = utils.WrapError("Encode rf128_r17", err)
		}
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_rf256_r17:
		if err = w.WriteInteger(int64(ie.rf256_r17), &uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			err = utils.WrapError("Encode rf256_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *MCCH_RepetitionPeriodAndOffset_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(9, false); err != nil {
		return err
	}
	switch ie.Choice {
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_rf1_r17:
		var tmp_int_rf1_r17 int64
		if tmp_int_rf1_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode rf1_r17", err)
		}
		ie.rf1_r17 = tmp_int_rf1_r17
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_rf2_r17:
		var tmp_int_rf2_r17 int64
		if tmp_int_rf2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode rf2_r17", err)
		}
		ie.rf2_r17 = tmp_int_rf2_r17
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_rf4_r17:
		var tmp_int_rf4_r17 int64
		if tmp_int_rf4_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode rf4_r17", err)
		}
		ie.rf4_r17 = tmp_int_rf4_r17
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_rf8_r17:
		var tmp_int_rf8_r17 int64
		if tmp_int_rf8_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode rf8_r17", err)
		}
		ie.rf8_r17 = tmp_int_rf8_r17
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_rf16_r17:
		var tmp_int_rf16_r17 int64
		if tmp_int_rf16_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode rf16_r17", err)
		}
		ie.rf16_r17 = tmp_int_rf16_r17
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_rf32_r17:
		var tmp_int_rf32_r17 int64
		if tmp_int_rf32_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
			return utils.WrapError("Decode rf32_r17", err)
		}
		ie.rf32_r17 = tmp_int_rf32_r17
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_rf64_r17:
		var tmp_int_rf64_r17 int64
		if tmp_int_rf64_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
			return utils.WrapError("Decode rf64_r17", err)
		}
		ie.rf64_r17 = tmp_int_rf64_r17
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_rf128_r17:
		var tmp_int_rf128_r17 int64
		if tmp_int_rf128_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode rf128_r17", err)
		}
		ie.rf128_r17 = tmp_int_rf128_r17
	case MCCH_RepetitionPeriodAndOffset_r17_Choice_rf256_r17:
		var tmp_int_rf256_r17 int64
		if tmp_int_rf256_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
			return utils.WrapError("Decode rf256_r17", err)
		}
		ie.rf256_r17 = tmp_int_rf256_r17
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
