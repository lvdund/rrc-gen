package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_InterUE_CoordinationScheme2_r17 struct {
	sl_IUC_Scheme2_r17                *SL_InterUE_CoordinationScheme2_r17_sl_IUC_Scheme2_r17                `optional`
	sl_RB_SetPSFCH_r17                *uper.BitString                                                       `lb:10,ub:275,optional`
	sl_TypeUE_A_r17                   *SL_InterUE_CoordinationScheme2_r17_sl_TypeUE_A_r17                   `optional`
	sl_PSFCH_Occasion_r17             *int64                                                                `lb:0,ub:1,optional`
	sl_SlotLevelResourceExclusion_r17 *SL_InterUE_CoordinationScheme2_r17_sl_SlotLevelResourceExclusion_r17 `optional`
	sl_OptionForCondition2_A_1_r17    *int64                                                                `lb:0,ub:1,optional`
	sl_IndicationUE_B_r17             *SL_InterUE_CoordinationScheme2_r17_sl_IndicationUE_B_r17             `optional`
	sl_DeltaRSRP_Thresh_v1720         *int64                                                                `lb:-30,ub:30,optional,ext-1`
}

func (ie *SL_InterUE_CoordinationScheme2_r17) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.sl_DeltaRSRP_Thresh_v1720 != nil
	preambleBits := []bool{hasExtensions, ie.sl_IUC_Scheme2_r17 != nil, ie.sl_RB_SetPSFCH_r17 != nil, ie.sl_TypeUE_A_r17 != nil, ie.sl_PSFCH_Occasion_r17 != nil, ie.sl_SlotLevelResourceExclusion_r17 != nil, ie.sl_OptionForCondition2_A_1_r17 != nil, ie.sl_IndicationUE_B_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_IUC_Scheme2_r17 != nil {
		if err = ie.sl_IUC_Scheme2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_IUC_Scheme2_r17", err)
		}
	}
	if ie.sl_RB_SetPSFCH_r17 != nil {
		if err = w.WriteBitString(ie.sl_RB_SetPSFCH_r17.Bytes, uint(ie.sl_RB_SetPSFCH_r17.NumBits), &uper.Constraint{Lb: 10, Ub: 275}, false); err != nil {
			return utils.WrapError("Encode sl_RB_SetPSFCH_r17", err)
		}
	}
	if ie.sl_TypeUE_A_r17 != nil {
		if err = ie.sl_TypeUE_A_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TypeUE_A_r17", err)
		}
	}
	if ie.sl_PSFCH_Occasion_r17 != nil {
		if err = w.WriteInteger(*ie.sl_PSFCH_Occasion_r17, &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Encode sl_PSFCH_Occasion_r17", err)
		}
	}
	if ie.sl_SlotLevelResourceExclusion_r17 != nil {
		if err = ie.sl_SlotLevelResourceExclusion_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_SlotLevelResourceExclusion_r17", err)
		}
	}
	if ie.sl_OptionForCondition2_A_1_r17 != nil {
		if err = w.WriteInteger(*ie.sl_OptionForCondition2_A_1_r17, &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Encode sl_OptionForCondition2_A_1_r17", err)
		}
	}
	if ie.sl_IndicationUE_B_r17 != nil {
		if err = ie.sl_IndicationUE_B_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_IndicationUE_B_r17", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.sl_DeltaRSRP_Thresh_v1720 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SL_InterUE_CoordinationScheme2_r17", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.sl_DeltaRSRP_Thresh_v1720 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sl_DeltaRSRP_Thresh_v1720 optional
			if ie.sl_DeltaRSRP_Thresh_v1720 != nil {
				if err = extWriter.WriteInteger(*ie.sl_DeltaRSRP_Thresh_v1720, &uper.Constraint{Lb: -30, Ub: 30}, false); err != nil {
					return utils.WrapError("Encode sl_DeltaRSRP_Thresh_v1720", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *SL_InterUE_CoordinationScheme2_r17) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_IUC_Scheme2_r17Present bool
	if sl_IUC_Scheme2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RB_SetPSFCH_r17Present bool
	if sl_RB_SetPSFCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TypeUE_A_r17Present bool
	if sl_TypeUE_A_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PSFCH_Occasion_r17Present bool
	if sl_PSFCH_Occasion_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_SlotLevelResourceExclusion_r17Present bool
	if sl_SlotLevelResourceExclusion_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_OptionForCondition2_A_1_r17Present bool
	if sl_OptionForCondition2_A_1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_IndicationUE_B_r17Present bool
	if sl_IndicationUE_B_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_IUC_Scheme2_r17Present {
		ie.sl_IUC_Scheme2_r17 = new(SL_InterUE_CoordinationScheme2_r17_sl_IUC_Scheme2_r17)
		if err = ie.sl_IUC_Scheme2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_IUC_Scheme2_r17", err)
		}
	}
	if sl_RB_SetPSFCH_r17Present {
		var tmp_bs_sl_RB_SetPSFCH_r17 []byte
		var n_sl_RB_SetPSFCH_r17 uint
		if tmp_bs_sl_RB_SetPSFCH_r17, n_sl_RB_SetPSFCH_r17, err = r.ReadBitString(&uper.Constraint{Lb: 10, Ub: 275}, false); err != nil {
			return utils.WrapError("Decode sl_RB_SetPSFCH_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_sl_RB_SetPSFCH_r17,
			NumBits: uint64(n_sl_RB_SetPSFCH_r17),
		}
		ie.sl_RB_SetPSFCH_r17 = &tmp_bitstring
	}
	if sl_TypeUE_A_r17Present {
		ie.sl_TypeUE_A_r17 = new(SL_InterUE_CoordinationScheme2_r17_sl_TypeUE_A_r17)
		if err = ie.sl_TypeUE_A_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TypeUE_A_r17", err)
		}
	}
	if sl_PSFCH_Occasion_r17Present {
		var tmp_int_sl_PSFCH_Occasion_r17 int64
		if tmp_int_sl_PSFCH_Occasion_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode sl_PSFCH_Occasion_r17", err)
		}
		ie.sl_PSFCH_Occasion_r17 = &tmp_int_sl_PSFCH_Occasion_r17
	}
	if sl_SlotLevelResourceExclusion_r17Present {
		ie.sl_SlotLevelResourceExclusion_r17 = new(SL_InterUE_CoordinationScheme2_r17_sl_SlotLevelResourceExclusion_r17)
		if err = ie.sl_SlotLevelResourceExclusion_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_SlotLevelResourceExclusion_r17", err)
		}
	}
	if sl_OptionForCondition2_A_1_r17Present {
		var tmp_int_sl_OptionForCondition2_A_1_r17 int64
		if tmp_int_sl_OptionForCondition2_A_1_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode sl_OptionForCondition2_A_1_r17", err)
		}
		ie.sl_OptionForCondition2_A_1_r17 = &tmp_int_sl_OptionForCondition2_A_1_r17
	}
	if sl_IndicationUE_B_r17Present {
		ie.sl_IndicationUE_B_r17 = new(SL_InterUE_CoordinationScheme2_r17_sl_IndicationUE_B_r17)
		if err = ie.sl_IndicationUE_B_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_IndicationUE_B_r17", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			sl_DeltaRSRP_Thresh_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sl_DeltaRSRP_Thresh_v1720 optional
			if sl_DeltaRSRP_Thresh_v1720Present {
				var tmp_int_sl_DeltaRSRP_Thresh_v1720 int64
				if tmp_int_sl_DeltaRSRP_Thresh_v1720, err = extReader.ReadInteger(&uper.Constraint{Lb: -30, Ub: 30}, false); err != nil {
					return utils.WrapError("Decode sl_DeltaRSRP_Thresh_v1720", err)
				}
				ie.sl_DeltaRSRP_Thresh_v1720 = &tmp_int_sl_DeltaRSRP_Thresh_v1720
			}
		}
	}
	return nil
}
