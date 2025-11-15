package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TDD_UL_DL_SlotConfig_IAB_MT_r16 struct {
	slotIndex_r16      TDD_UL_DL_SlotIndex                                 `madatory`
	symbols_IAB_MT_r16 *TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16 `lb:1,ub:maxNrofSymbols_1,optional`
}

func (ie *TDD_UL_DL_SlotConfig_IAB_MT_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.symbols_IAB_MT_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.slotIndex_r16.Encode(w); err != nil {
		return utils.WrapError("Encode slotIndex_r16", err)
	}
	if ie.symbols_IAB_MT_r16 != nil {
		if err = ie.symbols_IAB_MT_r16.Encode(w); err != nil {
			return utils.WrapError("Encode symbols_IAB_MT_r16", err)
		}
	}
	return nil
}

func (ie *TDD_UL_DL_SlotConfig_IAB_MT_r16) Decode(r *uper.UperReader) error {
	var err error
	var symbols_IAB_MT_r16Present bool
	if symbols_IAB_MT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.slotIndex_r16.Decode(r); err != nil {
		return utils.WrapError("Decode slotIndex_r16", err)
	}
	if symbols_IAB_MT_r16Present {
		ie.symbols_IAB_MT_r16 = new(TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16)
		if err = ie.symbols_IAB_MT_r16.Decode(r); err != nil {
			return utils.WrapError("Decode symbols_IAB_MT_r16", err)
		}
	}
	return nil
}
