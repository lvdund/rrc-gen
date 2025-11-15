package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PSFCH_Config_r16 struct {
	sl_PSFCH_Period_r16                *SL_PSFCH_Config_r16_sl_PSFCH_Period_r16                `optional`
	sl_PSFCH_RB_Set_r16                *uper.BitString                                         `lb:10,ub:275,optional`
	sl_NumMuxCS_Pair_r16               *SL_PSFCH_Config_r16_sl_NumMuxCS_Pair_r16               `optional`
	sl_MinTimeGapPSFCH_r16             *SL_PSFCH_Config_r16_sl_MinTimeGapPSFCH_r16             `optional`
	sl_PSFCH_HopID_r16                 *int64                                                  `lb:0,ub:1023,optional`
	sl_PSFCH_CandidateResourceType_r16 *SL_PSFCH_Config_r16_sl_PSFCH_CandidateResourceType_r16 `optional`
}

func (ie *SL_PSFCH_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_PSFCH_Period_r16 != nil, ie.sl_PSFCH_RB_Set_r16 != nil, ie.sl_NumMuxCS_Pair_r16 != nil, ie.sl_MinTimeGapPSFCH_r16 != nil, ie.sl_PSFCH_HopID_r16 != nil, ie.sl_PSFCH_CandidateResourceType_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_PSFCH_Period_r16 != nil {
		if err = ie.sl_PSFCH_Period_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PSFCH_Period_r16", err)
		}
	}
	if ie.sl_PSFCH_RB_Set_r16 != nil {
		if err = w.WriteBitString(ie.sl_PSFCH_RB_Set_r16.Bytes, uint(ie.sl_PSFCH_RB_Set_r16.NumBits), &uper.Constraint{Lb: 10, Ub: 275}, false); err != nil {
			return utils.WrapError("Encode sl_PSFCH_RB_Set_r16", err)
		}
	}
	if ie.sl_NumMuxCS_Pair_r16 != nil {
		if err = ie.sl_NumMuxCS_Pair_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_NumMuxCS_Pair_r16", err)
		}
	}
	if ie.sl_MinTimeGapPSFCH_r16 != nil {
		if err = ie.sl_MinTimeGapPSFCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MinTimeGapPSFCH_r16", err)
		}
	}
	if ie.sl_PSFCH_HopID_r16 != nil {
		if err = w.WriteInteger(*ie.sl_PSFCH_HopID_r16, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode sl_PSFCH_HopID_r16", err)
		}
	}
	if ie.sl_PSFCH_CandidateResourceType_r16 != nil {
		if err = ie.sl_PSFCH_CandidateResourceType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PSFCH_CandidateResourceType_r16", err)
		}
	}
	return nil
}

func (ie *SL_PSFCH_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_PSFCH_Period_r16Present bool
	if sl_PSFCH_Period_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PSFCH_RB_Set_r16Present bool
	if sl_PSFCH_RB_Set_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_NumMuxCS_Pair_r16Present bool
	if sl_NumMuxCS_Pair_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MinTimeGapPSFCH_r16Present bool
	if sl_MinTimeGapPSFCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PSFCH_HopID_r16Present bool
	if sl_PSFCH_HopID_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PSFCH_CandidateResourceType_r16Present bool
	if sl_PSFCH_CandidateResourceType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_PSFCH_Period_r16Present {
		ie.sl_PSFCH_Period_r16 = new(SL_PSFCH_Config_r16_sl_PSFCH_Period_r16)
		if err = ie.sl_PSFCH_Period_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PSFCH_Period_r16", err)
		}
	}
	if sl_PSFCH_RB_Set_r16Present {
		var tmp_bs_sl_PSFCH_RB_Set_r16 []byte
		var n_sl_PSFCH_RB_Set_r16 uint
		if tmp_bs_sl_PSFCH_RB_Set_r16, n_sl_PSFCH_RB_Set_r16, err = r.ReadBitString(&uper.Constraint{Lb: 10, Ub: 275}, false); err != nil {
			return utils.WrapError("Decode sl_PSFCH_RB_Set_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_sl_PSFCH_RB_Set_r16,
			NumBits: uint64(n_sl_PSFCH_RB_Set_r16),
		}
		ie.sl_PSFCH_RB_Set_r16 = &tmp_bitstring
	}
	if sl_NumMuxCS_Pair_r16Present {
		ie.sl_NumMuxCS_Pair_r16 = new(SL_PSFCH_Config_r16_sl_NumMuxCS_Pair_r16)
		if err = ie.sl_NumMuxCS_Pair_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_NumMuxCS_Pair_r16", err)
		}
	}
	if sl_MinTimeGapPSFCH_r16Present {
		ie.sl_MinTimeGapPSFCH_r16 = new(SL_PSFCH_Config_r16_sl_MinTimeGapPSFCH_r16)
		if err = ie.sl_MinTimeGapPSFCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MinTimeGapPSFCH_r16", err)
		}
	}
	if sl_PSFCH_HopID_r16Present {
		var tmp_int_sl_PSFCH_HopID_r16 int64
		if tmp_int_sl_PSFCH_HopID_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode sl_PSFCH_HopID_r16", err)
		}
		ie.sl_PSFCH_HopID_r16 = &tmp_int_sl_PSFCH_HopID_r16
	}
	if sl_PSFCH_CandidateResourceType_r16Present {
		ie.sl_PSFCH_CandidateResourceType_r16 = new(SL_PSFCH_Config_r16_sl_PSFCH_CandidateResourceType_r16)
		if err = ie.sl_PSFCH_CandidateResourceType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PSFCH_CandidateResourceType_r16", err)
		}
	}
	return nil
}
