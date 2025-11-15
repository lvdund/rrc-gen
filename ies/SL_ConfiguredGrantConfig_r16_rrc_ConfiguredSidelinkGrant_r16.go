package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_ConfiguredGrantConfig_r16_rrc_ConfiguredSidelinkGrant_r16 struct {
	sl_TimeResourceCG_Type1_r16    *int64                                                                                      `lb:0,ub:496,optional`
	sl_StartSubchannelCG_Type1_r16 *int64                                                                                      `lb:0,ub:26,optional`
	sl_FreqResourceCG_Type1_r16    *int64                                                                                      `lb:0,ub:6929,optional`
	sl_TimeOffsetCG_Type1_r16      *int64                                                                                      `lb:0,ub:7999,optional`
	sl_N1PUCCH_AN_r16              *PUCCH_ResourceId                                                                           `optional`
	sl_PSFCH_ToPUCCH_CG_Type1_r16  *int64                                                                                      `lb:0,ub:15,optional`
	sl_ResourcePoolID_r16          *SL_ResourcePoolID_r16                                                                      `optional`
	sl_TimeReferenceSFN_Type1_r16  *SL_ConfiguredGrantConfig_r16_rrc_ConfiguredSidelinkGrant_r16_sl_TimeReferenceSFN_Type1_r16 `optional`
}

func (ie *SL_ConfiguredGrantConfig_r16_rrc_ConfiguredSidelinkGrant_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_TimeResourceCG_Type1_r16 != nil, ie.sl_StartSubchannelCG_Type1_r16 != nil, ie.sl_FreqResourceCG_Type1_r16 != nil, ie.sl_TimeOffsetCG_Type1_r16 != nil, ie.sl_N1PUCCH_AN_r16 != nil, ie.sl_PSFCH_ToPUCCH_CG_Type1_r16 != nil, ie.sl_ResourcePoolID_r16 != nil, ie.sl_TimeReferenceSFN_Type1_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_TimeResourceCG_Type1_r16 != nil {
		if err = w.WriteInteger(*ie.sl_TimeResourceCG_Type1_r16, &uper.Constraint{Lb: 0, Ub: 496}, false); err != nil {
			return utils.WrapError("Encode sl_TimeResourceCG_Type1_r16", err)
		}
	}
	if ie.sl_StartSubchannelCG_Type1_r16 != nil {
		if err = w.WriteInteger(*ie.sl_StartSubchannelCG_Type1_r16, &uper.Constraint{Lb: 0, Ub: 26}, false); err != nil {
			return utils.WrapError("Encode sl_StartSubchannelCG_Type1_r16", err)
		}
	}
	if ie.sl_FreqResourceCG_Type1_r16 != nil {
		if err = w.WriteInteger(*ie.sl_FreqResourceCG_Type1_r16, &uper.Constraint{Lb: 0, Ub: 6929}, false); err != nil {
			return utils.WrapError("Encode sl_FreqResourceCG_Type1_r16", err)
		}
	}
	if ie.sl_TimeOffsetCG_Type1_r16 != nil {
		if err = w.WriteInteger(*ie.sl_TimeOffsetCG_Type1_r16, &uper.Constraint{Lb: 0, Ub: 7999}, false); err != nil {
			return utils.WrapError("Encode sl_TimeOffsetCG_Type1_r16", err)
		}
	}
	if ie.sl_N1PUCCH_AN_r16 != nil {
		if err = ie.sl_N1PUCCH_AN_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_N1PUCCH_AN_r16", err)
		}
	}
	if ie.sl_PSFCH_ToPUCCH_CG_Type1_r16 != nil {
		if err = w.WriteInteger(*ie.sl_PSFCH_ToPUCCH_CG_Type1_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode sl_PSFCH_ToPUCCH_CG_Type1_r16", err)
		}
	}
	if ie.sl_ResourcePoolID_r16 != nil {
		if err = ie.sl_ResourcePoolID_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ResourcePoolID_r16", err)
		}
	}
	if ie.sl_TimeReferenceSFN_Type1_r16 != nil {
		if err = ie.sl_TimeReferenceSFN_Type1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TimeReferenceSFN_Type1_r16", err)
		}
	}
	return nil
}

func (ie *SL_ConfiguredGrantConfig_r16_rrc_ConfiguredSidelinkGrant_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_TimeResourceCG_Type1_r16Present bool
	if sl_TimeResourceCG_Type1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_StartSubchannelCG_Type1_r16Present bool
	if sl_StartSubchannelCG_Type1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_FreqResourceCG_Type1_r16Present bool
	if sl_FreqResourceCG_Type1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TimeOffsetCG_Type1_r16Present bool
	if sl_TimeOffsetCG_Type1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_N1PUCCH_AN_r16Present bool
	if sl_N1PUCCH_AN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PSFCH_ToPUCCH_CG_Type1_r16Present bool
	if sl_PSFCH_ToPUCCH_CG_Type1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ResourcePoolID_r16Present bool
	if sl_ResourcePoolID_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TimeReferenceSFN_Type1_r16Present bool
	if sl_TimeReferenceSFN_Type1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_TimeResourceCG_Type1_r16Present {
		var tmp_int_sl_TimeResourceCG_Type1_r16 int64
		if tmp_int_sl_TimeResourceCG_Type1_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 496}, false); err != nil {
			return utils.WrapError("Decode sl_TimeResourceCG_Type1_r16", err)
		}
		ie.sl_TimeResourceCG_Type1_r16 = &tmp_int_sl_TimeResourceCG_Type1_r16
	}
	if sl_StartSubchannelCG_Type1_r16Present {
		var tmp_int_sl_StartSubchannelCG_Type1_r16 int64
		if tmp_int_sl_StartSubchannelCG_Type1_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 26}, false); err != nil {
			return utils.WrapError("Decode sl_StartSubchannelCG_Type1_r16", err)
		}
		ie.sl_StartSubchannelCG_Type1_r16 = &tmp_int_sl_StartSubchannelCG_Type1_r16
	}
	if sl_FreqResourceCG_Type1_r16Present {
		var tmp_int_sl_FreqResourceCG_Type1_r16 int64
		if tmp_int_sl_FreqResourceCG_Type1_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 6929}, false); err != nil {
			return utils.WrapError("Decode sl_FreqResourceCG_Type1_r16", err)
		}
		ie.sl_FreqResourceCG_Type1_r16 = &tmp_int_sl_FreqResourceCG_Type1_r16
	}
	if sl_TimeOffsetCG_Type1_r16Present {
		var tmp_int_sl_TimeOffsetCG_Type1_r16 int64
		if tmp_int_sl_TimeOffsetCG_Type1_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7999}, false); err != nil {
			return utils.WrapError("Decode sl_TimeOffsetCG_Type1_r16", err)
		}
		ie.sl_TimeOffsetCG_Type1_r16 = &tmp_int_sl_TimeOffsetCG_Type1_r16
	}
	if sl_N1PUCCH_AN_r16Present {
		ie.sl_N1PUCCH_AN_r16 = new(PUCCH_ResourceId)
		if err = ie.sl_N1PUCCH_AN_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_N1PUCCH_AN_r16", err)
		}
	}
	if sl_PSFCH_ToPUCCH_CG_Type1_r16Present {
		var tmp_int_sl_PSFCH_ToPUCCH_CG_Type1_r16 int64
		if tmp_int_sl_PSFCH_ToPUCCH_CG_Type1_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode sl_PSFCH_ToPUCCH_CG_Type1_r16", err)
		}
		ie.sl_PSFCH_ToPUCCH_CG_Type1_r16 = &tmp_int_sl_PSFCH_ToPUCCH_CG_Type1_r16
	}
	if sl_ResourcePoolID_r16Present {
		ie.sl_ResourcePoolID_r16 = new(SL_ResourcePoolID_r16)
		if err = ie.sl_ResourcePoolID_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ResourcePoolID_r16", err)
		}
	}
	if sl_TimeReferenceSFN_Type1_r16Present {
		ie.sl_TimeReferenceSFN_Type1_r16 = new(SL_ConfiguredGrantConfig_r16_rrc_ConfiguredSidelinkGrant_r16_sl_TimeReferenceSFN_Type1_r16)
		if err = ie.sl_TimeReferenceSFN_Type1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TimeReferenceSFN_Type1_r16", err)
		}
	}
	return nil
}
