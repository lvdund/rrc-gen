package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_BWP_Generic_r16_sl_LengthSymbols_r16_Enum_sym7  uper.Enumerated = 0
	SL_BWP_Generic_r16_sl_LengthSymbols_r16_Enum_sym8  uper.Enumerated = 1
	SL_BWP_Generic_r16_sl_LengthSymbols_r16_Enum_sym9  uper.Enumerated = 2
	SL_BWP_Generic_r16_sl_LengthSymbols_r16_Enum_sym10 uper.Enumerated = 3
	SL_BWP_Generic_r16_sl_LengthSymbols_r16_Enum_sym11 uper.Enumerated = 4
	SL_BWP_Generic_r16_sl_LengthSymbols_r16_Enum_sym12 uper.Enumerated = 5
	SL_BWP_Generic_r16_sl_LengthSymbols_r16_Enum_sym13 uper.Enumerated = 6
	SL_BWP_Generic_r16_sl_LengthSymbols_r16_Enum_sym14 uper.Enumerated = 7
)

type SL_BWP_Generic_r16_sl_LengthSymbols_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SL_BWP_Generic_r16_sl_LengthSymbols_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SL_BWP_Generic_r16_sl_LengthSymbols_r16", err)
	}
	return nil
}

func (ie *SL_BWP_Generic_r16_sl_LengthSymbols_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SL_BWP_Generic_r16_sl_LengthSymbols_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
