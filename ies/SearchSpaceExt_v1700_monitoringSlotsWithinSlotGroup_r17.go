package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17_Choice_nothing uint64 = iota
	SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17_Choice_slotGroupLength4_r17
	SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17_Choice_slotGroupLength8_r17
)

type SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17 struct {
	Choice               uint64
	slotGroupLength4_r17 uper.BitString `lb:4,ub:4,madatory`
	slotGroupLength8_r17 uper.BitString `lb:8,ub:8,madatory`
}

func (ie *SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17_Choice_slotGroupLength4_r17:
		if err = w.WriteBitString(ie.slotGroupLength4_r17.Bytes, uint(ie.slotGroupLength4_r17.NumBits), &uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode slotGroupLength4_r17", err)
		}
	case SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17_Choice_slotGroupLength8_r17:
		if err = w.WriteBitString(ie.slotGroupLength8_r17.Bytes, uint(ie.slotGroupLength8_r17.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode slotGroupLength8_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17_Choice_slotGroupLength4_r17:
		var tmp_bs_slotGroupLength4_r17 []byte
		var n_slotGroupLength4_r17 uint
		if tmp_bs_slotGroupLength4_r17, n_slotGroupLength4_r17, err = r.ReadBitString(&uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode slotGroupLength4_r17", err)
		}
		ie.slotGroupLength4_r17 = uper.BitString{
			Bytes:   tmp_bs_slotGroupLength4_r17,
			NumBits: uint64(n_slotGroupLength4_r17),
		}
	case SearchSpaceExt_v1700_monitoringSlotsWithinSlotGroup_r17_Choice_slotGroupLength8_r17:
		var tmp_bs_slotGroupLength8_r17 []byte
		var n_slotGroupLength8_r17 uint
		if tmp_bs_slotGroupLength8_r17, n_slotGroupLength8_r17, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode slotGroupLength8_r17", err)
		}
		ie.slotGroupLength8_r17 = uper.BitString{
			Bytes:   tmp_bs_slotGroupLength8_r17,
			NumBits: uint64(n_slotGroupLength8_r17),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
