package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TDD_UL_DL_SlotConfig_symbols_explicit struct {
	nrofDownlinkSymbols *int64 `lb:1,ub:maxNrofSymbols_1,optional`
	nrofUplinkSymbols   *int64 `lb:1,ub:maxNrofSymbols_1,optional`
}

func (ie *TDD_UL_DL_SlotConfig_symbols_explicit) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nrofDownlinkSymbols != nil, ie.nrofUplinkSymbols != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.nrofDownlinkSymbols != nil {
		if err = w.WriteInteger(*ie.nrofDownlinkSymbols, &uper.Constraint{Lb: 1, Ub: maxNrofSymbols_1}, false); err != nil {
			return utils.WrapError("Encode nrofDownlinkSymbols", err)
		}
	}
	if ie.nrofUplinkSymbols != nil {
		if err = w.WriteInteger(*ie.nrofUplinkSymbols, &uper.Constraint{Lb: 1, Ub: maxNrofSymbols_1}, false); err != nil {
			return utils.WrapError("Encode nrofUplinkSymbols", err)
		}
	}
	return nil
}

func (ie *TDD_UL_DL_SlotConfig_symbols_explicit) Decode(r *uper.UperReader) error {
	var err error
	var nrofDownlinkSymbolsPresent bool
	if nrofDownlinkSymbolsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nrofUplinkSymbolsPresent bool
	if nrofUplinkSymbolsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if nrofDownlinkSymbolsPresent {
		var tmp_int_nrofDownlinkSymbols int64
		if tmp_int_nrofDownlinkSymbols, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofSymbols_1}, false); err != nil {
			return utils.WrapError("Decode nrofDownlinkSymbols", err)
		}
		ie.nrofDownlinkSymbols = &tmp_int_nrofDownlinkSymbols
	}
	if nrofUplinkSymbolsPresent {
		var tmp_int_nrofUplinkSymbols int64
		if tmp_int_nrofUplinkSymbols, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofSymbols_1}, false); err != nil {
			return utils.WrapError("Decode nrofUplinkSymbols", err)
		}
		ie.nrofUplinkSymbols = &tmp_int_nrofUplinkSymbols
	}
	return nil
}
