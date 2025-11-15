package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_ConfigCommonNR_r16 struct {
	sl_FreqInfoList_r16                []SL_FreqConfigCommon_r16                          `lb:1,ub:maxNrofFreqSL_r16,optional`
	sl_UE_SelectedConfig_r16           *SL_UE_SelectedConfig_r16                          `optional`
	sl_NR_AnchorCarrierFreqList_r16    *SL_NR_AnchorCarrierFreqList_r16                   `optional`
	sl_EUTRA_AnchorCarrierFreqList_r16 *SL_EUTRA_AnchorCarrierFreqList_r16                `optional`
	sl_RadioBearerConfigList_r16       []SL_RadioBearerConfig_r16                         `lb:1,ub:maxNrofSLRB_r16,optional`
	sl_RLC_BearerConfigList_r16        []SL_RLC_BearerConfig_r16                          `lb:1,ub:maxSL_LCID_r16,optional`
	sl_MeasConfigCommon_r16            *SL_MeasConfigCommon_r16                           `optional`
	sl_CSI_Acquisition_r16             *SL_ConfigCommonNR_r16_sl_CSI_Acquisition_r16      `optional`
	sl_OffsetDFN_r16                   *int64                                             `lb:1,ub:1000,optional`
	t400_r16                           *SL_ConfigCommonNR_r16_t400_r16                    `optional`
	sl_MaxNumConsecutiveDTX_r16        *SL_ConfigCommonNR_r16_sl_MaxNumConsecutiveDTX_r16 `optional`
	sl_SSB_PriorityNR_r16              *int64                                             `lb:1,ub:8,optional`
}

func (ie *SL_ConfigCommonNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.sl_FreqInfoList_r16) > 0, ie.sl_UE_SelectedConfig_r16 != nil, ie.sl_NR_AnchorCarrierFreqList_r16 != nil, ie.sl_EUTRA_AnchorCarrierFreqList_r16 != nil, len(ie.sl_RadioBearerConfigList_r16) > 0, len(ie.sl_RLC_BearerConfigList_r16) > 0, ie.sl_MeasConfigCommon_r16 != nil, ie.sl_CSI_Acquisition_r16 != nil, ie.sl_OffsetDFN_r16 != nil, ie.t400_r16 != nil, ie.sl_MaxNumConsecutiveDTX_r16 != nil, ie.sl_SSB_PriorityNR_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.sl_FreqInfoList_r16) > 0 {
		tmp_sl_FreqInfoList_r16 := utils.NewSequence[*SL_FreqConfigCommon_r16]([]*SL_FreqConfigCommon_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		for _, i := range ie.sl_FreqInfoList_r16 {
			tmp_sl_FreqInfoList_r16.Value = append(tmp_sl_FreqInfoList_r16.Value, &i)
		}
		if err = tmp_sl_FreqInfoList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_FreqInfoList_r16", err)
		}
	}
	if ie.sl_UE_SelectedConfig_r16 != nil {
		if err = ie.sl_UE_SelectedConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_UE_SelectedConfig_r16", err)
		}
	}
	if ie.sl_NR_AnchorCarrierFreqList_r16 != nil {
		if err = ie.sl_NR_AnchorCarrierFreqList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_NR_AnchorCarrierFreqList_r16", err)
		}
	}
	if ie.sl_EUTRA_AnchorCarrierFreqList_r16 != nil {
		if err = ie.sl_EUTRA_AnchorCarrierFreqList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_EUTRA_AnchorCarrierFreqList_r16", err)
		}
	}
	if len(ie.sl_RadioBearerConfigList_r16) > 0 {
		tmp_sl_RadioBearerConfigList_r16 := utils.NewSequence[*SL_RadioBearerConfig_r16]([]*SL_RadioBearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		for _, i := range ie.sl_RadioBearerConfigList_r16 {
			tmp_sl_RadioBearerConfigList_r16.Value = append(tmp_sl_RadioBearerConfigList_r16.Value, &i)
		}
		if err = tmp_sl_RadioBearerConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RadioBearerConfigList_r16", err)
		}
	}
	if len(ie.sl_RLC_BearerConfigList_r16) > 0 {
		tmp_sl_RLC_BearerConfigList_r16 := utils.NewSequence[*SL_RLC_BearerConfig_r16]([]*SL_RLC_BearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		for _, i := range ie.sl_RLC_BearerConfigList_r16 {
			tmp_sl_RLC_BearerConfigList_r16.Value = append(tmp_sl_RLC_BearerConfigList_r16.Value, &i)
		}
		if err = tmp_sl_RLC_BearerConfigList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RLC_BearerConfigList_r16", err)
		}
	}
	if ie.sl_MeasConfigCommon_r16 != nil {
		if err = ie.sl_MeasConfigCommon_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MeasConfigCommon_r16", err)
		}
	}
	if ie.sl_CSI_Acquisition_r16 != nil {
		if err = ie.sl_CSI_Acquisition_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_CSI_Acquisition_r16", err)
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
	return nil
}

func (ie *SL_ConfigCommonNR_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_FreqInfoList_r16Present bool
	if sl_FreqInfoList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_UE_SelectedConfig_r16Present bool
	if sl_UE_SelectedConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_NR_AnchorCarrierFreqList_r16Present bool
	if sl_NR_AnchorCarrierFreqList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_EUTRA_AnchorCarrierFreqList_r16Present bool
	if sl_EUTRA_AnchorCarrierFreqList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RadioBearerConfigList_r16Present bool
	if sl_RadioBearerConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RLC_BearerConfigList_r16Present bool
	if sl_RLC_BearerConfigList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MeasConfigCommon_r16Present bool
	if sl_MeasConfigCommon_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_CSI_Acquisition_r16Present bool
	if sl_CSI_Acquisition_r16Present, err = r.ReadBool(); err != nil {
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
	if sl_FreqInfoList_r16Present {
		tmp_sl_FreqInfoList_r16 := utils.NewSequence[*SL_FreqConfigCommon_r16]([]*SL_FreqConfigCommon_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		fn_sl_FreqInfoList_r16 := func() *SL_FreqConfigCommon_r16 {
			return new(SL_FreqConfigCommon_r16)
		}
		if err = tmp_sl_FreqInfoList_r16.Decode(r, fn_sl_FreqInfoList_r16); err != nil {
			return utils.WrapError("Decode sl_FreqInfoList_r16", err)
		}
		ie.sl_FreqInfoList_r16 = []SL_FreqConfigCommon_r16{}
		for _, i := range tmp_sl_FreqInfoList_r16.Value {
			ie.sl_FreqInfoList_r16 = append(ie.sl_FreqInfoList_r16, *i)
		}
	}
	if sl_UE_SelectedConfig_r16Present {
		ie.sl_UE_SelectedConfig_r16 = new(SL_UE_SelectedConfig_r16)
		if err = ie.sl_UE_SelectedConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_UE_SelectedConfig_r16", err)
		}
	}
	if sl_NR_AnchorCarrierFreqList_r16Present {
		ie.sl_NR_AnchorCarrierFreqList_r16 = new(SL_NR_AnchorCarrierFreqList_r16)
		if err = ie.sl_NR_AnchorCarrierFreqList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_NR_AnchorCarrierFreqList_r16", err)
		}
	}
	if sl_EUTRA_AnchorCarrierFreqList_r16Present {
		ie.sl_EUTRA_AnchorCarrierFreqList_r16 = new(SL_EUTRA_AnchorCarrierFreqList_r16)
		if err = ie.sl_EUTRA_AnchorCarrierFreqList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_EUTRA_AnchorCarrierFreqList_r16", err)
		}
	}
	if sl_RadioBearerConfigList_r16Present {
		tmp_sl_RadioBearerConfigList_r16 := utils.NewSequence[*SL_RadioBearerConfig_r16]([]*SL_RadioBearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSLRB_r16}, false)
		fn_sl_RadioBearerConfigList_r16 := func() *SL_RadioBearerConfig_r16 {
			return new(SL_RadioBearerConfig_r16)
		}
		if err = tmp_sl_RadioBearerConfigList_r16.Decode(r, fn_sl_RadioBearerConfigList_r16); err != nil {
			return utils.WrapError("Decode sl_RadioBearerConfigList_r16", err)
		}
		ie.sl_RadioBearerConfigList_r16 = []SL_RadioBearerConfig_r16{}
		for _, i := range tmp_sl_RadioBearerConfigList_r16.Value {
			ie.sl_RadioBearerConfigList_r16 = append(ie.sl_RadioBearerConfigList_r16, *i)
		}
	}
	if sl_RLC_BearerConfigList_r16Present {
		tmp_sl_RLC_BearerConfigList_r16 := utils.NewSequence[*SL_RLC_BearerConfig_r16]([]*SL_RLC_BearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		fn_sl_RLC_BearerConfigList_r16 := func() *SL_RLC_BearerConfig_r16 {
			return new(SL_RLC_BearerConfig_r16)
		}
		if err = tmp_sl_RLC_BearerConfigList_r16.Decode(r, fn_sl_RLC_BearerConfigList_r16); err != nil {
			return utils.WrapError("Decode sl_RLC_BearerConfigList_r16", err)
		}
		ie.sl_RLC_BearerConfigList_r16 = []SL_RLC_BearerConfig_r16{}
		for _, i := range tmp_sl_RLC_BearerConfigList_r16.Value {
			ie.sl_RLC_BearerConfigList_r16 = append(ie.sl_RLC_BearerConfigList_r16, *i)
		}
	}
	if sl_MeasConfigCommon_r16Present {
		ie.sl_MeasConfigCommon_r16 = new(SL_MeasConfigCommon_r16)
		if err = ie.sl_MeasConfigCommon_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MeasConfigCommon_r16", err)
		}
	}
	if sl_CSI_Acquisition_r16Present {
		ie.sl_CSI_Acquisition_r16 = new(SL_ConfigCommonNR_r16_sl_CSI_Acquisition_r16)
		if err = ie.sl_CSI_Acquisition_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_CSI_Acquisition_r16", err)
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
		ie.t400_r16 = new(SL_ConfigCommonNR_r16_t400_r16)
		if err = ie.t400_r16.Decode(r); err != nil {
			return utils.WrapError("Decode t400_r16", err)
		}
	}
	if sl_MaxNumConsecutiveDTX_r16Present {
		ie.sl_MaxNumConsecutiveDTX_r16 = new(SL_ConfigCommonNR_r16_sl_MaxNumConsecutiveDTX_r16)
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
	return nil
}
