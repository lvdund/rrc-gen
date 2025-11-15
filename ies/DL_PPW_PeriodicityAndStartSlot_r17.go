package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DL_PPW_PeriodicityAndStartSlot_r17_Choice_nothing uint64 = iota
	DL_PPW_PeriodicityAndStartSlot_r17_Choice_scs15
	DL_PPW_PeriodicityAndStartSlot_r17_Choice_scs30
	DL_PPW_PeriodicityAndStartSlot_r17_Choice_scs60
	DL_PPW_PeriodicityAndStartSlot_r17_Choice_scs120
)

type DL_PPW_PeriodicityAndStartSlot_r17 struct {
	Choice uint64
	scs15  *DL_PPW_PeriodicityAndStartSlot_r17_scs15
	scs30  *DL_PPW_PeriodicityAndStartSlot_r17_scs30
	scs60  *DL_PPW_PeriodicityAndStartSlot_r17_scs60
	scs120 *DL_PPW_PeriodicityAndStartSlot_r17_scs120
}

func (ie *DL_PPW_PeriodicityAndStartSlot_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_PPW_PeriodicityAndStartSlot_r17_Choice_scs15:
		if err = ie.scs15.Encode(w); err != nil {
			err = utils.WrapError("Encode scs15", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_Choice_scs30:
		if err = ie.scs30.Encode(w); err != nil {
			err = utils.WrapError("Encode scs30", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_Choice_scs60:
		if err = ie.scs60.Encode(w); err != nil {
			err = utils.WrapError("Encode scs60", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_Choice_scs120:
		if err = ie.scs120.Encode(w); err != nil {
			err = utils.WrapError("Encode scs120", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *DL_PPW_PeriodicityAndStartSlot_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case DL_PPW_PeriodicityAndStartSlot_r17_Choice_scs15:
		ie.scs15 = new(DL_PPW_PeriodicityAndStartSlot_r17_scs15)
		if err = ie.scs15.Decode(r); err != nil {
			return utils.WrapError("Decode scs15", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_Choice_scs30:
		ie.scs30 = new(DL_PPW_PeriodicityAndStartSlot_r17_scs30)
		if err = ie.scs30.Decode(r); err != nil {
			return utils.WrapError("Decode scs30", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_Choice_scs60:
		ie.scs60 = new(DL_PPW_PeriodicityAndStartSlot_r17_scs60)
		if err = ie.scs60.Decode(r); err != nil {
			return utils.WrapError("Decode scs60", err)
		}
	case DL_PPW_PeriodicityAndStartSlot_r17_Choice_scs120:
		ie.scs120 = new(DL_PPW_PeriodicityAndStartSlot_r17_scs120)
		if err = ie.scs120.Decode(r); err != nil {
			return utils.WrapError("Decode scs120", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
