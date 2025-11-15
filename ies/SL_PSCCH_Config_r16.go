package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PSCCH_Config_r16 struct {
	sl_TimeResourcePSCCH_r16 *SL_PSCCH_Config_r16_sl_TimeResourcePSCCH_r16 `optional`
	sl_FreqResourcePSCCH_r16 *SL_PSCCH_Config_r16_sl_FreqResourcePSCCH_r16 `optional`
	sl_DMRS_ScrambleID_r16   *int64                                        `lb:0,ub:65535,optional`
	sl_NumReservedBits_r16   *int64                                        `lb:2,ub:4,optional`
}

func (ie *SL_PSCCH_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_TimeResourcePSCCH_r16 != nil, ie.sl_FreqResourcePSCCH_r16 != nil, ie.sl_DMRS_ScrambleID_r16 != nil, ie.sl_NumReservedBits_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_TimeResourcePSCCH_r16 != nil {
		if err = ie.sl_TimeResourcePSCCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TimeResourcePSCCH_r16", err)
		}
	}
	if ie.sl_FreqResourcePSCCH_r16 != nil {
		if err = ie.sl_FreqResourcePSCCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_FreqResourcePSCCH_r16", err)
		}
	}
	if ie.sl_DMRS_ScrambleID_r16 != nil {
		if err = w.WriteInteger(*ie.sl_DMRS_ScrambleID_r16, &uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode sl_DMRS_ScrambleID_r16", err)
		}
	}
	if ie.sl_NumReservedBits_r16 != nil {
		if err = w.WriteInteger(*ie.sl_NumReservedBits_r16, &uper.Constraint{Lb: 2, Ub: 4}, false); err != nil {
			return utils.WrapError("Encode sl_NumReservedBits_r16", err)
		}
	}
	return nil
}

func (ie *SL_PSCCH_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_TimeResourcePSCCH_r16Present bool
	if sl_TimeResourcePSCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_FreqResourcePSCCH_r16Present bool
	if sl_FreqResourcePSCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_DMRS_ScrambleID_r16Present bool
	if sl_DMRS_ScrambleID_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_NumReservedBits_r16Present bool
	if sl_NumReservedBits_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_TimeResourcePSCCH_r16Present {
		ie.sl_TimeResourcePSCCH_r16 = new(SL_PSCCH_Config_r16_sl_TimeResourcePSCCH_r16)
		if err = ie.sl_TimeResourcePSCCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TimeResourcePSCCH_r16", err)
		}
	}
	if sl_FreqResourcePSCCH_r16Present {
		ie.sl_FreqResourcePSCCH_r16 = new(SL_PSCCH_Config_r16_sl_FreqResourcePSCCH_r16)
		if err = ie.sl_FreqResourcePSCCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_FreqResourcePSCCH_r16", err)
		}
	}
	if sl_DMRS_ScrambleID_r16Present {
		var tmp_int_sl_DMRS_ScrambleID_r16 int64
		if tmp_int_sl_DMRS_ScrambleID_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode sl_DMRS_ScrambleID_r16", err)
		}
		ie.sl_DMRS_ScrambleID_r16 = &tmp_int_sl_DMRS_ScrambleID_r16
	}
	if sl_NumReservedBits_r16Present {
		var tmp_int_sl_NumReservedBits_r16 int64
		if tmp_int_sl_NumReservedBits_r16, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode sl_NumReservedBits_r16", err)
		}
		ie.sl_NumReservedBits_r16 = &tmp_int_sl_NumReservedBits_r16
	}
	return nil
}
