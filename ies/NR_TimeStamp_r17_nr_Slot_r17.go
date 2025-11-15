package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NR_TimeStamp_r17_nr_Slot_r17_Choice_nothing uint64 = iota
	NR_TimeStamp_r17_nr_Slot_r17_Choice_scs15_r17
	NR_TimeStamp_r17_nr_Slot_r17_Choice_scs30_r17
	NR_TimeStamp_r17_nr_Slot_r17_Choice_scs60_r17
	NR_TimeStamp_r17_nr_Slot_r17_Choice_scs120_r17
)

type NR_TimeStamp_r17_nr_Slot_r17 struct {
	Choice     uint64
	scs15_r17  int64 `lb:0,ub:9,madatory`
	scs30_r17  int64 `lb:0,ub:19,madatory`
	scs60_r17  int64 `lb:0,ub:39,madatory`
	scs120_r17 int64 `lb:0,ub:79,madatory`
}

func (ie *NR_TimeStamp_r17_nr_Slot_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_TimeStamp_r17_nr_Slot_r17_Choice_scs15_r17:
		if err = w.WriteInteger(int64(ie.scs15_r17), &uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			err = utils.WrapError("Encode scs15_r17", err)
		}
	case NR_TimeStamp_r17_nr_Slot_r17_Choice_scs30_r17:
		if err = w.WriteInteger(int64(ie.scs30_r17), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode scs30_r17", err)
		}
	case NR_TimeStamp_r17_nr_Slot_r17_Choice_scs60_r17:
		if err = w.WriteInteger(int64(ie.scs60_r17), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode scs60_r17", err)
		}
	case NR_TimeStamp_r17_nr_Slot_r17_Choice_scs120_r17:
		if err = w.WriteInteger(int64(ie.scs120_r17), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode scs120_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *NR_TimeStamp_r17_nr_Slot_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_TimeStamp_r17_nr_Slot_r17_Choice_scs15_r17:
		var tmp_int_scs15_r17 int64
		if tmp_int_scs15_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode scs15_r17", err)
		}
		ie.scs15_r17 = tmp_int_scs15_r17
	case NR_TimeStamp_r17_nr_Slot_r17_Choice_scs30_r17:
		var tmp_int_scs30_r17 int64
		if tmp_int_scs30_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode scs30_r17", err)
		}
		ie.scs30_r17 = tmp_int_scs30_r17
	case NR_TimeStamp_r17_nr_Slot_r17_Choice_scs60_r17:
		var tmp_int_scs60_r17 int64
		if tmp_int_scs60_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode scs60_r17", err)
		}
		ie.scs60_r17 = tmp_int_scs60_r17
	case NR_TimeStamp_r17_nr_Slot_r17_Choice_scs120_r17:
		var tmp_int_scs120_r17 int64
		if tmp_int_scs120_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode scs120_r17", err)
		}
		ie.scs120_r17 = tmp_int_scs120_r17
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
