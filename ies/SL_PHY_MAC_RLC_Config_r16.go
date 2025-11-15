package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PHY_MAC_RLC_Config_r16 struct {
	sl_ScheduledConfig_r16         *SL_ScheduledConfig_r16                                `optional,setuprelease`
	sl_UE_SelectedConfig_r16       *SL_UE_SelectedConfig_r16                              `optional,setuprelease`
	sl_FreqInfoToReleaseList_r16   []SL_Freq_Id_r16                                       `lb:1,ub:maxNrofFreqSL_r16,optional`
	sl_FreqInfoToAddModList_r16    []SL_FreqConfig_r16                                    `lb:1,ub:maxNrofFreqSL_r16,optional`
	sl_RLC_BearerToReleaseList_r16 []SL_RLC_BearerConfigIndex_r16                         `lb:1,ub:maxSL_LCID_r16,optional`
	sl_RLC_BearerToAddModList_r16  []SL_RLC_BearerConfig_r16                              `lb:1,ub:maxSL_LCID_r16,optional`
	sl_MaxNumConsecutiveDTX_r16    *SL_PHY_MAC_RLC_Config_r16_sl_MaxNumConsecutiveDTX_r16 `optional`
	sl_CSI_Acquisition_r16         *SL_PHY_MAC_RLC_Config_r16_sl_CSI_Acquisition_r16      `optional`
	sl_CSI_SchedulingRequestId_r16 *SchedulingRequestId                                   `optional,setuprelease`
	sl_SSB_PriorityNR_r16          *int64                                                 `lb:1,ub:8,optional`
	networkControlledSyncTx_r16    *SL_PHY_MAC_RLC_Config_r16_networkControlledSyncTx_r16 `optional`
}

func (ie *SL_PHY_MAC_RLC_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_ScheduledConfig_r16 != nil, ie.sl_UE_SelectedConfig_r16 != nil, len(ie.sl_FreqInfoToReleaseList_r16) > 0, len(ie.sl_FreqInfoToAddModList_r16) > 0, len(ie.sl_RLC_BearerToReleaseList_r16) > 0, len(ie.sl_RLC_BearerToAddModList_r16) > 0, ie.sl_MaxNumConsecutiveDTX_r16 != nil, ie.sl_CSI_Acquisition_r16 != nil, ie.sl_CSI_SchedulingRequestId_r16 != nil, ie.sl_SSB_PriorityNR_r16 != nil, ie.networkControlledSyncTx_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_ScheduledConfig_r16 != nil {
		tmp_sl_ScheduledConfig_r16 := utils.SetupRelease[*SL_ScheduledConfig_r16]{
			Setup: ie.sl_ScheduledConfig_r16,
		}
		if err = tmp_sl_ScheduledConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ScheduledConfig_r16", err)
		}
	}
	if ie.sl_UE_SelectedConfig_r16 != nil {
		tmp_sl_UE_SelectedConfig_r16 := utils.SetupRelease[*SL_UE_SelectedConfig_r16]{
			Setup: ie.sl_UE_SelectedConfig_r16,
		}
		if err = tmp_sl_UE_SelectedConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_UE_SelectedConfig_r16", err)
		}
	}
	if len(ie.sl_FreqInfoToReleaseList_r16) > 0 {
		tmp_sl_FreqInfoToReleaseList_r16 := utils.NewSequence[*SL_Freq_Id_r16]([]*SL_Freq_Id_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		for _, i := range ie.sl_FreqInfoToReleaseList_r16 {
			tmp_sl_FreqInfoToReleaseList_r16.Value = append(tmp_sl_FreqInfoToReleaseList_r16.Value, &i)
		}
		if err = tmp_sl_FreqInfoToReleaseList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_FreqInfoToReleaseList_r16", err)
		}
	}
	if len(ie.sl_FreqInfoToAddModList_r16) > 0 {
		tmp_sl_FreqInfoToAddModList_r16 := utils.NewSequence[*SL_FreqConfig_r16]([]*SL_FreqConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		for _, i := range ie.sl_FreqInfoToAddModList_r16 {
			tmp_sl_FreqInfoToAddModList_r16.Value = append(tmp_sl_FreqInfoToAddModList_r16.Value, &i)
		}
		if err = tmp_sl_FreqInfoToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_FreqInfoToAddModList_r16", err)
		}
	}
	if len(ie.sl_RLC_BearerToReleaseList_r16) > 0 {
		tmp_sl_RLC_BearerToReleaseList_r16 := utils.NewSequence[*SL_RLC_BearerConfigIndex_r16]([]*SL_RLC_BearerConfigIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		for _, i := range ie.sl_RLC_BearerToReleaseList_r16 {
			tmp_sl_RLC_BearerToReleaseList_r16.Value = append(tmp_sl_RLC_BearerToReleaseList_r16.Value, &i)
		}
		if err = tmp_sl_RLC_BearerToReleaseList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RLC_BearerToReleaseList_r16", err)
		}
	}
	if len(ie.sl_RLC_BearerToAddModList_r16) > 0 {
		tmp_sl_RLC_BearerToAddModList_r16 := utils.NewSequence[*SL_RLC_BearerConfig_r16]([]*SL_RLC_BearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		for _, i := range ie.sl_RLC_BearerToAddModList_r16 {
			tmp_sl_RLC_BearerToAddModList_r16.Value = append(tmp_sl_RLC_BearerToAddModList_r16.Value, &i)
		}
		if err = tmp_sl_RLC_BearerToAddModList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RLC_BearerToAddModList_r16", err)
		}
	}
	if ie.sl_MaxNumConsecutiveDTX_r16 != nil {
		if err = ie.sl_MaxNumConsecutiveDTX_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MaxNumConsecutiveDTX_r16", err)
		}
	}
	if ie.sl_CSI_Acquisition_r16 != nil {
		if err = ie.sl_CSI_Acquisition_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_CSI_Acquisition_r16", err)
		}
	}
	if ie.sl_CSI_SchedulingRequestId_r16 != nil {
		tmp_sl_CSI_SchedulingRequestId_r16 := utils.SetupRelease[*SchedulingRequestId]{
			Setup: ie.sl_CSI_SchedulingRequestId_r16,
		}
		if err = tmp_sl_CSI_SchedulingRequestId_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_CSI_SchedulingRequestId_r16", err)
		}
	}
	if ie.sl_SSB_PriorityNR_r16 != nil {
		if err = w.WriteInteger(*ie.sl_SSB_PriorityNR_r16, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode sl_SSB_PriorityNR_r16", err)
		}
	}
	if ie.networkControlledSyncTx_r16 != nil {
		if err = ie.networkControlledSyncTx_r16.Encode(w); err != nil {
			return utils.WrapError("Encode networkControlledSyncTx_r16", err)
		}
	}
	return nil
}

func (ie *SL_PHY_MAC_RLC_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_ScheduledConfig_r16Present bool
	if sl_ScheduledConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_UE_SelectedConfig_r16Present bool
	if sl_UE_SelectedConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_FreqInfoToReleaseList_r16Present bool
	if sl_FreqInfoToReleaseList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_FreqInfoToAddModList_r16Present bool
	if sl_FreqInfoToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RLC_BearerToReleaseList_r16Present bool
	if sl_RLC_BearerToReleaseList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RLC_BearerToAddModList_r16Present bool
	if sl_RLC_BearerToAddModList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MaxNumConsecutiveDTX_r16Present bool
	if sl_MaxNumConsecutiveDTX_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_CSI_Acquisition_r16Present bool
	if sl_CSI_Acquisition_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_CSI_SchedulingRequestId_r16Present bool
	if sl_CSI_SchedulingRequestId_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_SSB_PriorityNR_r16Present bool
	if sl_SSB_PriorityNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var networkControlledSyncTx_r16Present bool
	if networkControlledSyncTx_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_ScheduledConfig_r16Present {
		tmp_sl_ScheduledConfig_r16 := utils.SetupRelease[*SL_ScheduledConfig_r16]{}
		if err = tmp_sl_ScheduledConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ScheduledConfig_r16", err)
		}
		ie.sl_ScheduledConfig_r16 = tmp_sl_ScheduledConfig_r16.Setup
	}
	if sl_UE_SelectedConfig_r16Present {
		tmp_sl_UE_SelectedConfig_r16 := utils.SetupRelease[*SL_UE_SelectedConfig_r16]{}
		if err = tmp_sl_UE_SelectedConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_UE_SelectedConfig_r16", err)
		}
		ie.sl_UE_SelectedConfig_r16 = tmp_sl_UE_SelectedConfig_r16.Setup
	}
	if sl_FreqInfoToReleaseList_r16Present {
		tmp_sl_FreqInfoToReleaseList_r16 := utils.NewSequence[*SL_Freq_Id_r16]([]*SL_Freq_Id_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		fn_sl_FreqInfoToReleaseList_r16 := func() *SL_Freq_Id_r16 {
			return new(SL_Freq_Id_r16)
		}
		if err = tmp_sl_FreqInfoToReleaseList_r16.Decode(r, fn_sl_FreqInfoToReleaseList_r16); err != nil {
			return utils.WrapError("Decode sl_FreqInfoToReleaseList_r16", err)
		}
		ie.sl_FreqInfoToReleaseList_r16 = []SL_Freq_Id_r16{}
		for _, i := range tmp_sl_FreqInfoToReleaseList_r16.Value {
			ie.sl_FreqInfoToReleaseList_r16 = append(ie.sl_FreqInfoToReleaseList_r16, *i)
		}
	}
	if sl_FreqInfoToAddModList_r16Present {
		tmp_sl_FreqInfoToAddModList_r16 := utils.NewSequence[*SL_FreqConfig_r16]([]*SL_FreqConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		fn_sl_FreqInfoToAddModList_r16 := func() *SL_FreqConfig_r16 {
			return new(SL_FreqConfig_r16)
		}
		if err = tmp_sl_FreqInfoToAddModList_r16.Decode(r, fn_sl_FreqInfoToAddModList_r16); err != nil {
			return utils.WrapError("Decode sl_FreqInfoToAddModList_r16", err)
		}
		ie.sl_FreqInfoToAddModList_r16 = []SL_FreqConfig_r16{}
		for _, i := range tmp_sl_FreqInfoToAddModList_r16.Value {
			ie.sl_FreqInfoToAddModList_r16 = append(ie.sl_FreqInfoToAddModList_r16, *i)
		}
	}
	if sl_RLC_BearerToReleaseList_r16Present {
		tmp_sl_RLC_BearerToReleaseList_r16 := utils.NewSequence[*SL_RLC_BearerConfigIndex_r16]([]*SL_RLC_BearerConfigIndex_r16{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		fn_sl_RLC_BearerToReleaseList_r16 := func() *SL_RLC_BearerConfigIndex_r16 {
			return new(SL_RLC_BearerConfigIndex_r16)
		}
		if err = tmp_sl_RLC_BearerToReleaseList_r16.Decode(r, fn_sl_RLC_BearerToReleaseList_r16); err != nil {
			return utils.WrapError("Decode sl_RLC_BearerToReleaseList_r16", err)
		}
		ie.sl_RLC_BearerToReleaseList_r16 = []SL_RLC_BearerConfigIndex_r16{}
		for _, i := range tmp_sl_RLC_BearerToReleaseList_r16.Value {
			ie.sl_RLC_BearerToReleaseList_r16 = append(ie.sl_RLC_BearerToReleaseList_r16, *i)
		}
	}
	if sl_RLC_BearerToAddModList_r16Present {
		tmp_sl_RLC_BearerToAddModList_r16 := utils.NewSequence[*SL_RLC_BearerConfig_r16]([]*SL_RLC_BearerConfig_r16{}, uper.Constraint{Lb: 1, Ub: maxSL_LCID_r16}, false)
		fn_sl_RLC_BearerToAddModList_r16 := func() *SL_RLC_BearerConfig_r16 {
			return new(SL_RLC_BearerConfig_r16)
		}
		if err = tmp_sl_RLC_BearerToAddModList_r16.Decode(r, fn_sl_RLC_BearerToAddModList_r16); err != nil {
			return utils.WrapError("Decode sl_RLC_BearerToAddModList_r16", err)
		}
		ie.sl_RLC_BearerToAddModList_r16 = []SL_RLC_BearerConfig_r16{}
		for _, i := range tmp_sl_RLC_BearerToAddModList_r16.Value {
			ie.sl_RLC_BearerToAddModList_r16 = append(ie.sl_RLC_BearerToAddModList_r16, *i)
		}
	}
	if sl_MaxNumConsecutiveDTX_r16Present {
		ie.sl_MaxNumConsecutiveDTX_r16 = new(SL_PHY_MAC_RLC_Config_r16_sl_MaxNumConsecutiveDTX_r16)
		if err = ie.sl_MaxNumConsecutiveDTX_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MaxNumConsecutiveDTX_r16", err)
		}
	}
	if sl_CSI_Acquisition_r16Present {
		ie.sl_CSI_Acquisition_r16 = new(SL_PHY_MAC_RLC_Config_r16_sl_CSI_Acquisition_r16)
		if err = ie.sl_CSI_Acquisition_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_CSI_Acquisition_r16", err)
		}
	}
	if sl_CSI_SchedulingRequestId_r16Present {
		tmp_sl_CSI_SchedulingRequestId_r16 := utils.SetupRelease[*SchedulingRequestId]{}
		if err = tmp_sl_CSI_SchedulingRequestId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_CSI_SchedulingRequestId_r16", err)
		}
		ie.sl_CSI_SchedulingRequestId_r16 = tmp_sl_CSI_SchedulingRequestId_r16.Setup
	}
	if sl_SSB_PriorityNR_r16Present {
		var tmp_int_sl_SSB_PriorityNR_r16 int64
		if tmp_int_sl_SSB_PriorityNR_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode sl_SSB_PriorityNR_r16", err)
		}
		ie.sl_SSB_PriorityNR_r16 = &tmp_int_sl_SSB_PriorityNR_r16
	}
	if networkControlledSyncTx_r16Present {
		ie.networkControlledSyncTx_r16 = new(SL_PHY_MAC_RLC_Config_r16_networkControlledSyncTx_r16)
		if err = ie.networkControlledSyncTx_r16.Decode(r); err != nil {
			return utils.WrapError("Decode networkControlledSyncTx_r16", err)
		}
	}
	return nil
}
