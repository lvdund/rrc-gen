package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_explicit_r16 struct {
	nrofDownlinkSymbols_r16 *int64 `lb:1,ub:maxNrofSymbols_1,optional`
	nrofUplinkSymbols_r16   *int64 `lb:1,ub:maxNrofSymbols_1,optional`
}

func (ie *TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_explicit_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.nrofDownlinkSymbols_r16 != nil, ie.nrofUplinkSymbols_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.nrofDownlinkSymbols_r16 != nil {
		if err = w.WriteInteger(*ie.nrofDownlinkSymbols_r16, &uper.Constraint{Lb: 1, Ub: maxNrofSymbols_1}, false); err != nil {
			return utils.WrapError("Encode nrofDownlinkSymbols_r16", err)
		}
	}
	if ie.nrofUplinkSymbols_r16 != nil {
		if err = w.WriteInteger(*ie.nrofUplinkSymbols_r16, &uper.Constraint{Lb: 1, Ub: maxNrofSymbols_1}, false); err != nil {
			return utils.WrapError("Encode nrofUplinkSymbols_r16", err)
		}
	}
	return nil
}

func (ie *TDD_UL_DL_SlotConfig_IAB_MT_r16_symbols_IAB_MT_r16_explicit_r16) Decode(r *uper.UperReader) error {
	var err error
	var nrofDownlinkSymbols_r16Present bool
	if nrofDownlinkSymbols_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nrofUplinkSymbols_r16Present bool
	if nrofUplinkSymbols_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if nrofDownlinkSymbols_r16Present {
		var tmp_int_nrofDownlinkSymbols_r16 int64
		if tmp_int_nrofDownlinkSymbols_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofSymbols_1}, false); err != nil {
			return utils.WrapError("Decode nrofDownlinkSymbols_r16", err)
		}
		ie.nrofDownlinkSymbols_r16 = &tmp_int_nrofDownlinkSymbols_r16
	}
	if nrofUplinkSymbols_r16Present {
		var tmp_int_nrofUplinkSymbols_r16 int64
		if tmp_int_nrofUplinkSymbols_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofSymbols_1}, false); err != nil {
			return utils.WrapError("Decode nrofUplinkSymbols_r16", err)
		}
		ie.nrofUplinkSymbols_r16 = &tmp_int_nrofUplinkSymbols_r16
	}
	return nil
}
