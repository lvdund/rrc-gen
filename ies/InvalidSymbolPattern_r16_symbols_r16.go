package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	InvalidSymbolPattern_r16_symbols_r16_Choice_nothing uint64 = iota
	InvalidSymbolPattern_r16_symbols_r16_Choice_oneSlot
	InvalidSymbolPattern_r16_symbols_r16_Choice_twoSlots
)

type InvalidSymbolPattern_r16_symbols_r16 struct {
	Choice   uint64
	oneSlot  uper.BitString `lb:14,ub:14,madatory`
	twoSlots uper.BitString `lb:28,ub:28,madatory`
}

func (ie *InvalidSymbolPattern_r16_symbols_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case InvalidSymbolPattern_r16_symbols_r16_Choice_oneSlot:
		if err = w.WriteBitString(ie.oneSlot.Bytes, uint(ie.oneSlot.NumBits), &uper.Constraint{Lb: 14, Ub: 14}, false); err != nil {
			err = utils.WrapError("Encode oneSlot", err)
		}
	case InvalidSymbolPattern_r16_symbols_r16_Choice_twoSlots:
		if err = w.WriteBitString(ie.twoSlots.Bytes, uint(ie.twoSlots.NumBits), &uper.Constraint{Lb: 28, Ub: 28}, false); err != nil {
			err = utils.WrapError("Encode twoSlots", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *InvalidSymbolPattern_r16_symbols_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case InvalidSymbolPattern_r16_symbols_r16_Choice_oneSlot:
		var tmp_bs_oneSlot []byte
		var n_oneSlot uint
		if tmp_bs_oneSlot, n_oneSlot, err = r.ReadBitString(&uper.Constraint{Lb: 14, Ub: 14}, false); err != nil {
			return utils.WrapError("Decode oneSlot", err)
		}
		ie.oneSlot = uper.BitString{
			Bytes:   tmp_bs_oneSlot,
			NumBits: uint64(n_oneSlot),
		}
	case InvalidSymbolPattern_r16_symbols_r16_Choice_twoSlots:
		var tmp_bs_twoSlots []byte
		var n_twoSlots uint
		if tmp_bs_twoSlots, n_twoSlots, err = r.ReadBitString(&uper.Constraint{Lb: 28, Ub: 28}, false); err != nil {
			return utils.WrapError("Decode twoSlots", err)
		}
		ie.twoSlots = uper.BitString{
			Bytes:   tmp_bs_twoSlots,
			NumBits: uint64(n_twoSlots),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
