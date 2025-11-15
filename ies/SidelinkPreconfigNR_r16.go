package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SidelinkPreconfigNR_r16 struct {
	sl_PreconfigFreqInfoList_r16                []SL_FreqConfigCommon_r16                            `lb:1,ub:maxNrofFreqSL_r16,optional`
	sl_PreconfigNR_AnchorCarrierFreqList_r16    *SL_NR_AnchorCarrierFreqList_r16                     `optional`
	sl_PreconfigEUTRA_AnchorCarrierFreqList_r16 *SL_EUTRA_AnchorCarrierFreqList_r16                  `optional`
	sl_RadioBearerPreConfigList_r16             []SL_RadioBearerConfig_r16                           `lb:1,ub:maxNrofSLRB_r16,optional`
	sl_RLC_BearerPreConfigList_r16              []SL_RLC_BearerConfig_r16                            `lb:1,ub:maxSL_LCID_r16,optional`
	sl_MeasPreConfig_r16                        *SL_MeasConfigCommon_r16                             `optional`
	sl_OffsetDFN_r16                            *int64                                               `lb:1,ub:1000,optional`
	t400_r16                                    *SidelinkPreconfigNR_r16_t400_r16                    `optional`
	sl_MaxNumConsecutiveDTX_r16                 *SidelinkPreconfigNR_r16_sl_MaxNumConsecutiveDTX_r16 `optional`
	sl_SSB_PriorityNR_r16                       *int64                                               `lb:1,ub:8,optional`
	sl_PreconfigGeneral_r16                     *SL_PreconfigGeneral_r16                             `optional`
	sl_UE_SelectedPreConfig_r16                 *SL_UE_SelectedConfig_r16                            `optional`
	sl_CSI_Acquisition_r16                      *SidelinkPreconfigNR_r16_sl_CSI_Acquisition_r16      `optional`
	sl_RoHC_Profiles_r16                        *SL_RoHC_Profiles_r16                                `optional`
	sl_MaxCID_r16                               int64                                                `lb:1,ub:16383,madatory`
	sl_DRX_PreConfigGC_BC_r17                   *SL_DRX_ConfigGC_BC_r17                              `optional,ext-1`
	sl_TxProfileList_r17                        *SL_TxProfileList_r17                                `optional,ext-1`
	sl_PreconfigDiscConfig_r17                  *SL_RemoteUE_Config_r17                              `optional,ext-1`
}

func (ie *SidelinkPreconfigNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.sl_DRX_PreConfigGC_BC_r17 != nil || ie.sl_TxProfileList_r17 != nil || ie.sl_PreconfigDiscConfig_r17 != nil
	preambleBits := []bool{hasExtensions, len(ie.sl_PreconfigFreqInfoList_r16) > 0, ie.sl_PreconfigNR_AnchorCarrierFreqList_r16 != nil, ie.sl_PreconfigEUTRA_AnchorCarrierFreqList_r16 != nil, len(ie.sl_RadioBearerPreConfigList_r16) > 0, len(ie.sl_RLC_BearerPreConfigList_r16) > 0, ie.sl_MeasPreConfig_r16 != nil, ie.sl_OffsetDFN_r16 != nil, ie.t400_r16 != nil, ie.sl_MaxNumConsecutiveDTX_r16 != nil, ie.sl_SSB_PriorityNR_r16 != nil, ie.sl_PreconfigGeneral_r16 != nil, ie.sl_UE_SelectedPreConfig_r16 != nil, ie.sl_CSI_Acquisition_r16 != nil, ie.sl_RoHC_Profiles_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.sl_PreconfigFreqInfoList_r16) > 0 {
		tmp_sl_PreconfigFreqInfoList_r16 := utils.NewSequence[*SL_FreqConfigCommon_r16]([]*SL_FreqConfigCommon_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		for _, i := range ie.sl_PreconfigFreqInfoList_r16 {
			tmp_sl_PreconfigFreqInfoList_r16.Value = append(tmp_sl_PreconfigFreqInfoList_r16.Value, &i)
		}
		if err = tmp_sl_PreconfigFreqInfoList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PreconfigFreqInfoList_r16", err)
		}
	}
	if ie.sl_PreconfigNR_AnchorCarrierFreqList_r16 != nil {
		if err = ie.sl_PreconfigNR_AnchorCarrierFreqList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PreconfigNR_AnchorCarrierFreqList_r16", err)
		}
	}
	if ie.sl_PreconfigEUTRA_AnchorCarrierFreqList_r16 != nil {
		if err = ie.sl_PreconfigEUTRA_AnchorCarrierFreqList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PreconfigEUTRA_AnchorCarrierFreqList_r16", err)
		}
	}
	if len(ie.sl_RadioBearerPreConfigList_r16) > 0 {
		tmp_sl_RadioBearerPreConfigList_r16 := utils.NewSequence[*SL_RadioBearerConfig_r16]([]*SL_RadioBearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		for _, i := range ie.sl_RadioBearerPreConfigList_r16 {
			tmp_sl_RadioBearerPreConfigList_r16.Value = append(tmp_sl_RadioBearerPreConfigList_r16.Value, &i)
		}
		if err = tmp_sl_RadioBearerPreConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RadioBearerPreConfigList_r16", err)
		}
	}
	if len(ie.sl_RLC_BearerPreConfigList_r16) > 0 {
		tmp_sl_RLC_BearerPreConfigList_r16 := utils.NewSequence[*SL_RLC_BearerConfig_r16]([]*SL_RLC_BearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		for _, i := range ie.sl_RLC_BearerPreConfigList_r16 {
			tmp_sl_RLC_BearerPreConfigList_r16.Value = append(tmp_sl_RLC_BearerPreConfigList_r16.Value, &i)
		}
		if err = tmp_sl_RLC_BearerPreConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RLC_BearerPreConfigList_r16", err)
		}
	}
	if ie.sl_MeasPreConfig_r16 != nil {
		if err = ie.sl_MeasPreConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MeasPreConfig_r16", err)
		}
	}
	if ie.sl_OffsetDFN_r16 != nil {
		if err = w.WriteInteger(*ie.sl_OffsetDFN_r16, &uper.Constraint{Lb: 1, Ub: 1000}, false); err != nil {
			return utils.WrapError("Encode sl_OffsetDFN_r16", err)
		}
	}
	if ie.t400_r16 != nil {
		if err = ie.t400_r16.Encode(w); err != nil {
			return utils.WrapError("Encode t400_r16", err)
		}
	}
	if ie.sl_MaxNumConsecutiveDTX_r16 != nil {
		if err = ie.sl_MaxNumConsecutiveDTX_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MaxNumConsecutiveDTX_r16", err)
		}
	}
	if ie.sl_SSB_PriorityNR_r16 != nil {
		if err = w.WriteInteger(*ie.sl_SSB_PriorityNR_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode sl_SSB_PriorityNR_r16", err)
		}
	}
	if ie.sl_PreconfigGeneral_r16 != nil {
		if err = ie.sl_PreconfigGeneral_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PreconfigGeneral_r16", err)
		}
	}
	if ie.sl_UE_SelectedPreConfig_r16 != nil {
		if err = ie.sl_UE_SelectedPreConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_UE_SelectedPreConfig_r16", err)
		}
	}
	if ie.sl_CSI_Acquisition_r16 != nil {
		if err = ie.sl_CSI_Acquisition_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_CSI_Acquisition_r16", err)
		}
	}
	if ie.sl_RoHC_Profiles_r16 != nil {
		if err = ie.sl_RoHC_Profiles_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RoHC_Profiles_r16", err)
		}
	}
	if err = w.WriteInteger(ie.sl_MaxCID_r16, &uper.Constraint{Lb: 1, Ub: 16383}, false); err != nil {
		return utils.WrapError("WriteInteger sl_MaxCID_r16", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.sl_DRX_PreConfigGC_BC_r17 != nil || ie.sl_TxProfileList_r17 != nil || ie.sl_PreconfigDiscConfig_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SidelinkPreconfigNR_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.sl_DRX_PreConfigGC_BC_r17 != nil, ie.sl_TxProfileList_r17 != nil, ie.sl_PreconfigDiscConfig_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sl_DRX_PreConfigGC_BC_r17 optional
			if ie.sl_DRX_PreConfigGC_BC_r17 != nil {
				if err = ie.sl_DRX_PreConfigGC_BC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_DRX_PreConfigGC_BC_r17", err)
				}
			}
			// encode sl_TxProfileList_r17 optional
			if ie.sl_TxProfileList_r17 != nil {
				if err = ie.sl_TxProfileList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_TxProfileList_r17", err)
				}
			}
			// encode sl_PreconfigDiscConfig_r17 optional
			if ie.sl_PreconfigDiscConfig_r17 != nil {
				if err = ie.sl_PreconfigDiscConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_PreconfigDiscConfig_r17", err)
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

func (ie *SidelinkPreconfigNR_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PreconfigFreqInfoList_r16Present bool
	if sl_PreconfigFreqInfoList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PreconfigNR_AnchorCarrierFreqList_r16Present bool
	if sl_PreconfigNR_AnchorCarrierFreqList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PreconfigEUTRA_AnchorCarrierFreqList_r16Present bool
	if sl_PreconfigEUTRA_AnchorCarrierFreqList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RadioBearerPreConfigList_r16Present bool
	if sl_RadioBearerPreConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RLC_BearerPreConfigList_r16Present bool
	if sl_RLC_BearerPreConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MeasPreConfig_r16Present bool
	if sl_MeasPreConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_OffsetDFN_r16Present bool
	if sl_OffsetDFN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var t400_r16Present bool
	if t400_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MaxNumConsecutiveDTX_r16Present bool
	if sl_MaxNumConsecutiveDTX_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_SSB_PriorityNR_r16Present bool
	if sl_SSB_PriorityNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PreconfigGeneral_r16Present bool
	if sl_PreconfigGeneral_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_UE_SelectedPreConfig_r16Present bool
	if sl_UE_SelectedPreConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_CSI_Acquisition_r16Present bool
	if sl_CSI_Acquisition_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RoHC_Profiles_r16Present bool
	if sl_RoHC_Profiles_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_PreconfigFreqInfoList_r16Present {
		tmp_sl_PreconfigFreqInfoList_r16 := utils.NewSequence[*SL_FreqConfigCommon_r16]([]*SL_FreqConfigCommon_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		fn_sl_PreconfigFreqInfoList_r16 := func() *SL_FreqConfigCommon_r16 {
			return new(SL_FreqConfigCommon_r16)
		}
		if err = tmp_sl_PreconfigFreqInfoList_r16.Decode(r, fn_sl_PreconfigFreqInfoList_r16); err != nil {
			return utils.WrapError("Decode sl_PreconfigFreqInfoList_r16", err)
		}
		ie.sl_PreconfigFreqInfoList_r16 = []SL_FreqConfigCommon_r16{}
		for _, i := range tmp_sl_PreconfigFreqInfoList_r16.Value {
			ie.sl_PreconfigFreqInfoList_r16 = append(ie.sl_PreconfigFreqInfoList_r16, *i)
		}
	}
	if sl_PreconfigNR_AnchorCarrierFreqList_r16Present {
		ie.sl_PreconfigNR_AnchorCarrierFreqList_r16 = new(SL_NR_AnchorCarrierFreqList_r16)
		if err = ie.sl_PreconfigNR_AnchorCarrierFreqList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PreconfigNR_AnchorCarrierFreqList_r16", err)
		}
	}
	if sl_PreconfigEUTRA_AnchorCarrierFreqList_r16Present {
		ie.sl_PreconfigEUTRA_AnchorCarrierFreqList_r16 = new(SL_EUTRA_AnchorCarrierFreqList_r16)
		if err = ie.sl_PreconfigEUTRA_AnchorCarrierFreqList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PreconfigEUTRA_AnchorCarrierFreqList_r16", err)
		}
	}
	if sl_RadioBearerPreConfigList_r16Present {
		tmp_sl_RadioBearerPreConfigList_r16 := utils.NewSequence[*SL_RadioBearerConfig_r16]([]*SL_RadioBearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		fn_sl_RadioBearerPreConfigList_r16 := func() *SL_RadioBearerConfig_r16 {
			return new(SL_RadioBearerConfig_r16)
		}
		if err = tmp_sl_RadioBearerPreConfigList_r16.Decode(r, fn_sl_RadioBearerPreConfigList_r16); err != nil {
			return utils.WrapError("Decode sl_RadioBearerPreConfigList_r16", err)
		}
		ie.sl_RadioBearerPreConfigList_r16 = []SL_RadioBearerConfig_r16{}
		for _, i := range tmp_sl_RadioBearerPreConfigList_r16.Value {
			ie.sl_RadioBearerPreConfigList_r16 = append(ie.sl_RadioBearerPreConfigList_r16, *i)
		}
	}
	if sl_RLC_BearerPreConfigList_r16Present {
		tmp_sl_RLC_BearerPreConfigList_r16 := utils.NewSequence[*SL_RLC_BearerConfig_r16]([]*SL_RLC_BearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		fn_sl_RLC_BearerPreConfigList_r16 := func() *SL_RLC_BearerConfig_r16 {
			return new(SL_RLC_BearerConfig_r16)
		}
		if err = tmp_sl_RLC_BearerPreConfigList_r16.Decode(r, fn_sl_RLC_BearerPreConfigList_r16); err != nil {
			return utils.WrapError("Decode sl_RLC_BearerPreConfigList_r16", err)
		}
		ie.sl_RLC_BearerPreConfigList_r16 = []SL_RLC_BearerConfig_r16{}
		for _, i := range tmp_sl_RLC_BearerPreConfigList_r16.Value {
			ie.sl_RLC_BearerPreConfigList_r16 = append(ie.sl_RLC_BearerPreConfigList_r16, *i)
		}
	}
	if sl_MeasPreConfig_r16Present {
		ie.sl_MeasPreConfig_r16 = new(SL_MeasConfigCommon_r16)
		if err = ie.sl_MeasPreConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MeasPreConfig_r16", err)
		}
	}
	if sl_OffsetDFN_r16Present {
		var tmp_int_sl_OffsetDFN_r16 int64
		if tmp_int_sl_OffsetDFN_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 1000}, false); err != nil {
			return utils.WrapError("Decode sl_OffsetDFN_r16", err)
		}
		ie.sl_OffsetDFN_r16 = &tmp_int_sl_OffsetDFN_r16
	}
	if t400_r16Present {
		ie.t400_r16 = new(SidelinkPreconfigNR_r16_t400_r16)
		if err = ie.t400_r16.Decode(r); err != nil {
			return utils.WrapError("Decode t400_r16", err)
		}
	}
	if sl_MaxNumConsecutiveDTX_r16Present {
		ie.sl_MaxNumConsecutiveDTX_r16 = new(SidelinkPreconfigNR_r16_sl_MaxNumConsecutiveDTX_r16)
		if err = ie.sl_MaxNumConsecutiveDTX_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MaxNumConsecutiveDTX_r16", err)
		}
	}
	if sl_SSB_PriorityNR_r16Present {
		var tmp_int_sl_SSB_PriorityNR_r16 int64
		if tmp_int_sl_SSB_PriorityNR_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode sl_SSB_PriorityNR_r16", err)
		}
		ie.sl_SSB_PriorityNR_r16 = &tmp_int_sl_SSB_PriorityNR_r16
	}
	if sl_PreconfigGeneral_r16Present {
		ie.sl_PreconfigGeneral_r16 = new(SL_PreconfigGeneral_r16)
		if err = ie.sl_PreconfigGeneral_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PreconfigGeneral_r16", err)
		}
	}
	if sl_UE_SelectedPreConfig_r16Present {
		ie.sl_UE_SelectedPreConfig_r16 = new(SL_UE_SelectedConfig_r16)
		if err = ie.sl_UE_SelectedPreConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_UE_SelectedPreConfig_r16", err)
		}
	}
	if sl_CSI_Acquisition_r16Present {
		ie.sl_CSI_Acquisition_r16 = new(SidelinkPreconfigNR_r16_sl_CSI_Acquisition_r16)
		if err = ie.sl_CSI_Acquisition_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_CSI_Acquisition_r16", err)
		}
	}
	if sl_RoHC_Profiles_r16Present {
		ie.sl_RoHC_Profiles_r16 = new(SL_RoHC_Profiles_r16)
		if err = ie.sl_RoHC_Profiles_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_RoHC_Profiles_r16", err)
		}
	}
	var tmp_int_sl_MaxCID_r16 int64
	if tmp_int_sl_MaxCID_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16383}, false); err != nil {
		return utils.WrapError("ReadInteger sl_MaxCID_r16", err)
	}
	ie.sl_MaxCID_r16 = tmp_int_sl_MaxCID_r16

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

			sl_DRX_PreConfigGC_BC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sl_TxProfileList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sl_PreconfigDiscConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sl_DRX_PreConfigGC_BC_r17 optional
			if sl_DRX_PreConfigGC_BC_r17Present {
				ie.sl_DRX_PreConfigGC_BC_r17 = new(SL_DRX_ConfigGC_BC_r17)
				if err = ie.sl_DRX_PreConfigGC_BC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_DRX_PreConfigGC_BC_r17", err)
				}
			}
			// decode sl_TxProfileList_r17 optional
			if sl_TxProfileList_r17Present {
				ie.sl_TxProfileList_r17 = new(SL_TxProfileList_r17)
				if err = ie.sl_TxProfileList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_TxProfileList_r17", err)
				}
			}
			// decode sl_PreconfigDiscConfig_r17 optional
			if sl_PreconfigDiscConfig_r17Present {
				ie.sl_PreconfigDiscConfig_r17 = new(SL_RemoteUE_Config_r17)
				if err = ie.sl_PreconfigDiscConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_PreconfigDiscConfig_r17", err)
				}
			}
		}
	}
	return nil
}
