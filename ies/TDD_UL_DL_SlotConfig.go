package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TDD_UL_DL_SlotConfig struct {
	slotIndex TDD_UL_DL_SlotIndex           `madatory`
	symbols   *TDD_UL_DL_SlotConfig_symbols `lb:1,ub:maxNrofSymbols_1,optional`
}

func (ie *TDD_UL_DL_SlotConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.symbols != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.slotIndex.Encode(w); err != nil {
		return utils.WrapError("Encode slotIndex", err)
	}
	if ie.symbols != nil {
		if err = ie.symbols.Encode(w); err != nil {
			return utils.WrapError("Encode symbols", err)
		}
	}
	return nil
}

func (ie *TDD_UL_DL_SlotConfig) Decode(r *uper.UperReader) error {
	var err error
	var symbolsPresent bool
	if symbolsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.slotIndex.Decode(r); err != nil {
		return utils.WrapError("Decode slotIndex", err)
	}
	if symbolsPresent {
		ie.symbols = new(TDD_UL_DL_SlotConfig_symbols)
		if err = ie.symbols.Decode(r); err != nil {
			return utils.WrapError("Decode symbols", err)
		}
	}
	return nil
}
