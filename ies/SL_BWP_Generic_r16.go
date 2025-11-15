package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_BWP_Generic_r16 struct {
	sl_BWP_r16                     *BWP                                     `optional`
	sl_LengthSymbols_r16           *SL_BWP_Generic_r16_sl_LengthSymbols_r16 `optional`
	sl_StartSymbol_r16             *SL_BWP_Generic_r16_sl_StartSymbol_r16   `optional`
	sl_PSBCH_Config_r16            *SL_PSBCH_Config_r16                     `optional,setuprelease`
	sl_TxDirectCurrentLocation_r16 *int64                                   `lb:0,ub:3301,optional`
}

func (ie *SL_BWP_Generic_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_BWP_r16 != nil, ie.sl_LengthSymbols_r16 != nil, ie.sl_StartSymbol_r16 != nil, ie.sl_PSBCH_Config_r16 != nil, ie.sl_TxDirectCurrentLocation_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_BWP_r16 != nil {
		if err = ie.sl_BWP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_BWP_r16", err)
		}
	}
	if ie.sl_LengthSymbols_r16 != nil {
		if err = ie.sl_LengthSymbols_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_LengthSymbols_r16", err)
		}
	}
	if ie.sl_StartSymbol_r16 != nil {
		if err = ie.sl_StartSymbol_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_StartSymbol_r16", err)
		}
	}
	if ie.sl_PSBCH_Config_r16 != nil {
		tmp_sl_PSBCH_Config_r16 := utils.SetupRelease[*SL_PSBCH_Config_r16]{
			Setup: ie.sl_PSBCH_Config_r16,
		}
		if err = tmp_sl_PSBCH_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PSBCH_Config_r16", err)
		}
	}
	if ie.sl_TxDirectCurrentLocation_r16 != nil {
		if err = w.WriteInteger(*ie.sl_TxDirectCurrentLocation_r16, &uper.Constraint{Lb: 0, Ub: 3301}, false); err != nil {
			return utils.WrapError("Encode sl_TxDirectCurrentLocation_r16", err)
		}
	}
	return nil
}

func (ie *SL_BWP_Generic_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_BWP_r16Present bool
	if sl_BWP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_LengthSymbols_r16Present bool
	if sl_LengthSymbols_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_StartSymbol_r16Present bool
	if sl_StartSymbol_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PSBCH_Config_r16Present bool
	if sl_PSBCH_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TxDirectCurrentLocation_r16Present bool
	if sl_TxDirectCurrentLocation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_BWP_r16Present {
		ie.sl_BWP_r16 = new(BWP)
		if err = ie.sl_BWP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_BWP_r16", err)
		}
	}
	if sl_LengthSymbols_r16Present {
		ie.sl_LengthSymbols_r16 = new(SL_BWP_Generic_r16_sl_LengthSymbols_r16)
		if err = ie.sl_LengthSymbols_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_LengthSymbols_r16", err)
		}
	}
	if sl_StartSymbol_r16Present {
		ie.sl_StartSymbol_r16 = new(SL_BWP_Generic_r16_sl_StartSymbol_r16)
		if err = ie.sl_StartSymbol_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_StartSymbol_r16", err)
		}
	}
	if sl_PSBCH_Config_r16Present {
		tmp_sl_PSBCH_Config_r16 := utils.SetupRelease[*SL_PSBCH_Config_r16]{}
		if err = tmp_sl_PSBCH_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PSBCH_Config_r16", err)
		}
		ie.sl_PSBCH_Config_r16 = tmp_sl_PSBCH_Config_r16.Setup
	}
	if sl_TxDirectCurrentLocation_r16Present {
		var tmp_int_sl_TxDirectCurrentLocation_r16 int64
		if tmp_int_sl_TxDirectCurrentLocation_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3301}, false); err != nil {
			return utils.WrapError("Decode sl_TxDirectCurrentLocation_r16", err)
		}
		ie.sl_TxDirectCurrentLocation_r16 = &tmp_int_sl_TxDirectCurrentLocation_r16
	}
	return nil
}
