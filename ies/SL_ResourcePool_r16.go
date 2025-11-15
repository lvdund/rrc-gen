package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_ResourcePool_r16 struct {
	sl_PSCCH_Config_r16               *SL_PSCCH_Config_r16                             `optional,setuprelease`
	sl_PSSCH_Config_r16               *SL_PSSCH_Config_r16                             `optional,setuprelease`
	sl_PSFCH_Config_r16               *SL_PSFCH_Config_r16                             `optional,setuprelease`
	sl_SyncAllowed_r16                *SL_SyncAllowed_r16                              `optional`
	sl_SubchannelSize_r16             *SL_ResourcePool_r16_sl_SubchannelSize_r16       `optional`
	dummy                             *int64                                           `lb:10,ub:160,optional`
	sl_StartRB_Subchannel_r16         *int64                                           `lb:0,ub:265,optional`
	sl_NumSubchannel_r16              *int64                                           `lb:1,ub:27,optional`
	sl_Additional_MCS_Table_r16       *SL_ResourcePool_r16_sl_Additional_MCS_Table_r16 `optional`
	sl_ThreshS_RSSI_CBR_r16           *int64                                           `lb:0,ub:45,optional`
	sl_TimeWindowSizeCBR_r16          *SL_ResourcePool_r16_sl_TimeWindowSizeCBR_r16    `optional`
	sl_TimeWindowSizeCR_r16           *SL_ResourcePool_r16_sl_TimeWindowSizeCR_r16     `optional`
	sl_PTRS_Config_r16                *SL_PTRS_Config_r16                              `optional`
	sl_UE_SelectedConfigRP_r16        *SL_UE_SelectedConfigRP_r16                      `optional`
	sl_RxParametersNcell_r16          *Sl_RxParametersNcell_r16                        `optional`
	sl_ZoneConfigMCR_List_r16         []SL_ZoneConfigMCR_r16                           `lb:16,ub:16,optional`
	sl_FilterCoefficient_r16          *FilterCoefficient                               `optional`
	sl_RB_Number_r16                  *int64                                           `lb:10,ub:275,optional`
	sl_PreemptionEnable_r16           *SL_ResourcePool_r16_sl_PreemptionEnable_r16     `optional`
	sl_PriorityThreshold_UL_URLLC_r16 *int64                                           `lb:1,ub:9,optional`
	sl_PriorityThreshold_r16          *int64                                           `lb:1,ub:9,optional`
	sl_X_Overhead_r16                 *SL_ResourcePool_r16_sl_X_Overhead_r16           `optional`
	sl_PowerControl_r16               *SL_PowerControl_r16                             `optional`
	sl_TxPercentageList_r16           *SL_TxPercentageList_r16                         `optional`
	sl_MinMaxMCS_List_r16             *SL_MinMaxMCS_List_r16                           `optional`
	sl_TimeResource_r16               *uper.BitString                                  `lb:10,ub:160,optional,ext-1`
	sl_PBPS_CPS_Config_r17            *SL_PBPS_CPS_Config_r17                          `optional,ext-2,setuprelease`
	sl_InterUE_CoordinationConfig_r17 *SL_InterUE_CoordinationConfig_r17               `optional,ext-2,setuprelease`
}

func (ie *SL_ResourcePool_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.sl_TimeResource_r16 != nil || ie.sl_PBPS_CPS_Config_r17 != nil || ie.sl_InterUE_CoordinationConfig_r17 != nil
	preambleBits := []bool{hasExtensions, ie.sl_PSCCH_Config_r16 != nil, ie.sl_PSSCH_Config_r16 != nil, ie.sl_PSFCH_Config_r16 != nil, ie.sl_SyncAllowed_r16 != nil, ie.sl_SubchannelSize_r16 != nil, ie.dummy != nil, ie.sl_StartRB_Subchannel_r16 != nil, ie.sl_NumSubchannel_r16 != nil, ie.sl_Additional_MCS_Table_r16 != nil, ie.sl_ThreshS_RSSI_CBR_r16 != nil, ie.sl_TimeWindowSizeCBR_r16 != nil, ie.sl_TimeWindowSizeCR_r16 != nil, ie.sl_PTRS_Config_r16 != nil, ie.sl_UE_SelectedConfigRP_r16 != nil, ie.sl_RxParametersNcell_r16 != nil, len(ie.sl_ZoneConfigMCR_List_r16) > 0, ie.sl_FilterCoefficient_r16 != nil, ie.sl_RB_Number_r16 != nil, ie.sl_PreemptionEnable_r16 != nil, ie.sl_PriorityThreshold_UL_URLLC_r16 != nil, ie.sl_PriorityThreshold_r16 != nil, ie.sl_X_Overhead_r16 != nil, ie.sl_PowerControl_r16 != nil, ie.sl_TxPercentageList_r16 != nil, ie.sl_MinMaxMCS_List_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_PSCCH_Config_r16 != nil {
		tmp_sl_PSCCH_Config_r16 := utils.SetupRelease[*SL_PSCCH_Config_r16]{
			Setup: ie.sl_PSCCH_Config_r16,
		}
		if err = tmp_sl_PSCCH_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PSCCH_Config_r16", err)
		}
	}
	if ie.sl_PSSCH_Config_r16 != nil {
		tmp_sl_PSSCH_Config_r16 := utils.SetupRelease[*SL_PSSCH_Config_r16]{
			Setup: ie.sl_PSSCH_Config_r16,
		}
		if err = tmp_sl_PSSCH_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PSSCH_Config_r16", err)
		}
	}
	if ie.sl_PSFCH_Config_r16 != nil {
		tmp_sl_PSFCH_Config_r16 := utils.SetupRelease[*SL_PSFCH_Config_r16]{
			Setup: ie.sl_PSFCH_Config_r16,
		}
		if err = tmp_sl_PSFCH_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PSFCH_Config_r16", err)
		}
	}
	if ie.sl_SyncAllowed_r16 != nil {
		if err = ie.sl_SyncAllowed_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_SyncAllowed_r16", err)
		}
	}
	if ie.sl_SubchannelSize_r16 != nil {
		if err = ie.sl_SubchannelSize_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_SubchannelSize_r16", err)
		}
	}
	if ie.dummy != nil {
		if err = w.WriteInteger(*ie.dummy, &uper.Constraint{Lb: 10, Ub: 160}, false); err != nil {
			return utils.WrapError("Encode dummy", err)
		}
	}
	if ie.sl_StartRB_Subchannel_r16 != nil {
		if err = w.WriteInteger(*ie.sl_StartRB_Subchannel_r16, &uper.Constraint{Lb: 0, Ub: 265}, false); err != nil {
			return utils.WrapError("Encode sl_StartRB_Subchannel_r16", err)
		}
	}
	if ie.sl_NumSubchannel_r16 != nil {
		if err = w.WriteInteger(*ie.sl_NumSubchannel_r16, &uper.Constraint{Lb: 1, Ub: 27}, false); err != nil {
			return utils.WrapError("Encode sl_NumSubchannel_r16", err)
		}
	}
	if ie.sl_Additional_MCS_Table_r16 != nil {
		if err = ie.sl_Additional_MCS_Table_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_Additional_MCS_Table_r16", err)
		}
	}
	if ie.sl_ThreshS_RSSI_CBR_r16 != nil {
		if err = w.WriteInteger(*ie.sl_ThreshS_RSSI_CBR_r16, &uper.Constraint{Lb: 0, Ub: 45}, false); err != nil {
			return utils.WrapError("Encode sl_ThreshS_RSSI_CBR_r16", err)
		}
	}
	if ie.sl_TimeWindowSizeCBR_r16 != nil {
		if err = ie.sl_TimeWindowSizeCBR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TimeWindowSizeCBR_r16", err)
		}
	}
	if ie.sl_TimeWindowSizeCR_r16 != nil {
		if err = ie.sl_TimeWindowSizeCR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TimeWindowSizeCR_r16", err)
		}
	}
	if ie.sl_PTRS_Config_r16 != nil {
		if err = ie.sl_PTRS_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PTRS_Config_r16", err)
		}
	}
	if ie.sl_UE_SelectedConfigRP_r16 != nil {
		if err = ie.sl_UE_SelectedConfigRP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_UE_SelectedConfigRP_r16", err)
		}
	}
	if ie.sl_RxParametersNcell_r16 != nil {
		if err = ie.sl_RxParametersNcell_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RxParametersNcell_r16", err)
		}
	}
	if len(ie.sl_ZoneConfigMCR_List_r16) > 0 {
		tmp_sl_ZoneConfigMCR_List_r16 := utils.NewSequence[*SL_ZoneConfigMCR_r16]([]*SL_ZoneConfigMCR_r16{}, uper.Constraint{Lb: 16, Ub: 16}, false)
		for _, i := range ie.sl_ZoneConfigMCR_List_r16 {
			tmp_sl_ZoneConfigMCR_List_r16.Value = append(tmp_sl_ZoneConfigMCR_List_r16.Value, &i)
		}
		if err = tmp_sl_ZoneConfigMCR_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ZoneConfigMCR_List_r16", err)
		}
	}
	if ie.sl_FilterCoefficient_r16 != nil {
		if err = ie.sl_FilterCoefficient_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_FilterCoefficient_r16", err)
		}
	}
	if ie.sl_RB_Number_r16 != nil {
		if err = w.WriteInteger(*ie.sl_RB_Number_r16, &uper.Constraint{Lb: 10, Ub: 275}, false); err != nil {
			return utils.WrapError("Encode sl_RB_Number_r16", err)
		}
	}
	if ie.sl_PreemptionEnable_r16 != nil {
		if err = ie.sl_PreemptionEnable_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PreemptionEnable_r16", err)
		}
	}
	if ie.sl_PriorityThreshold_UL_URLLC_r16 != nil {
		if err = w.WriteInteger(*ie.sl_PriorityThreshold_UL_URLLC_r16, &uper.Constraint{Lb: 1, Ub: 9}, false); err != nil {
			return utils.WrapError("Encode sl_PriorityThreshold_UL_URLLC_r16", err)
		}
	}
	if ie.sl_PriorityThreshold_r16 != nil {
		if err = w.WriteInteger(*ie.sl_PriorityThreshold_r16, &uper.Constraint{Lb: 1, Ub: 9}, false); err != nil {
			return utils.WrapError("Encode sl_PriorityThreshold_r16", err)
		}
	}
	if ie.sl_X_Overhead_r16 != nil {
		if err = ie.sl_X_Overhead_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_X_Overhead_r16", err)
		}
	}
	if ie.sl_PowerControl_r16 != nil {
		if err = ie.sl_PowerControl_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PowerControl_r16", err)
		}
	}
	if ie.sl_TxPercentageList_r16 != nil {
		if err = ie.sl_TxPercentageList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TxPercentageList_r16", err)
		}
	}
	if ie.sl_MinMaxMCS_List_r16 != nil {
		if err = ie.sl_MinMaxMCS_List_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MinMaxMCS_List_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.sl_TimeResource_r16 != nil, ie.sl_PBPS_CPS_Config_r17 != nil || ie.sl_InterUE_CoordinationConfig_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SL_ResourcePool_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.sl_TimeResource_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sl_TimeResource_r16 optional
			if ie.sl_TimeResource_r16 != nil {
				if err = extWriter.WriteBitString(ie.sl_TimeResource_r16.Bytes, uint(ie.sl_TimeResource_r16.NumBits), &uper.Constraint{Lb: 10, Ub: 160}, false); err != nil {
					return utils.WrapError("Encode sl_TimeResource_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.sl_PBPS_CPS_Config_r17 != nil, ie.sl_InterUE_CoordinationConfig_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sl_PBPS_CPS_Config_r17 optional
			if ie.sl_PBPS_CPS_Config_r17 != nil {
				tmp_sl_PBPS_CPS_Config_r17 := utils.SetupRelease[*SL_PBPS_CPS_Config_r17]{
					Setup: ie.sl_PBPS_CPS_Config_r17,
				}
				if err = tmp_sl_PBPS_CPS_Config_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_PBPS_CPS_Config_r17", err)
				}
			}
			// encode sl_InterUE_CoordinationConfig_r17 optional
			if ie.sl_InterUE_CoordinationConfig_r17 != nil {
				tmp_sl_InterUE_CoordinationConfig_r17 := utils.SetupRelease[*SL_InterUE_CoordinationConfig_r17]{
					Setup: ie.sl_InterUE_CoordinationConfig_r17,
				}
				if err = tmp_sl_InterUE_CoordinationConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_InterUE_CoordinationConfig_r17", err)
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

func (ie *SL_ResourcePool_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PSCCH_Config_r16Present bool
	if sl_PSCCH_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PSSCH_Config_r16Present bool
	if sl_PSSCH_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PSFCH_Config_r16Present bool
	if sl_PSFCH_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_SyncAllowed_r16Present bool
	if sl_SyncAllowed_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_SubchannelSize_r16Present bool
	if sl_SubchannelSize_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dummyPresent bool
	if dummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_StartRB_Subchannel_r16Present bool
	if sl_StartRB_Subchannel_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_NumSubchannel_r16Present bool
	if sl_NumSubchannel_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_Additional_MCS_Table_r16Present bool
	if sl_Additional_MCS_Table_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ThreshS_RSSI_CBR_r16Present bool
	if sl_ThreshS_RSSI_CBR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TimeWindowSizeCBR_r16Present bool
	if sl_TimeWindowSizeCBR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TimeWindowSizeCR_r16Present bool
	if sl_TimeWindowSizeCR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PTRS_Config_r16Present bool
	if sl_PTRS_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_UE_SelectedConfigRP_r16Present bool
	if sl_UE_SelectedConfigRP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RxParametersNcell_r16Present bool
	if sl_RxParametersNcell_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_ZoneConfigMCR_List_r16Present bool
	if sl_ZoneConfigMCR_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_FilterCoefficient_r16Present bool
	if sl_FilterCoefficient_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RB_Number_r16Present bool
	if sl_RB_Number_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PreemptionEnable_r16Present bool
	if sl_PreemptionEnable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PriorityThreshold_UL_URLLC_r16Present bool
	if sl_PriorityThreshold_UL_URLLC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PriorityThreshold_r16Present bool
	if sl_PriorityThreshold_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_X_Overhead_r16Present bool
	if sl_X_Overhead_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PowerControl_r16Present bool
	if sl_PowerControl_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TxPercentageList_r16Present bool
	if sl_TxPercentageList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MinMaxMCS_List_r16Present bool
	if sl_MinMaxMCS_List_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_PSCCH_Config_r16Present {
		tmp_sl_PSCCH_Config_r16 := utils.SetupRelease[*SL_PSCCH_Config_r16]{}
		if err = tmp_sl_PSCCH_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PSCCH_Config_r16", err)
		}
		ie.sl_PSCCH_Config_r16 = tmp_sl_PSCCH_Config_r16.Setup
	}
	if sl_PSSCH_Config_r16Present {
		tmp_sl_PSSCH_Config_r16 := utils.SetupRelease[*SL_PSSCH_Config_r16]{}
		if err = tmp_sl_PSSCH_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PSSCH_Config_r16", err)
		}
		ie.sl_PSSCH_Config_r16 = tmp_sl_PSSCH_Config_r16.Setup
	}
	if sl_PSFCH_Config_r16Present {
		tmp_sl_PSFCH_Config_r16 := utils.SetupRelease[*SL_PSFCH_Config_r16]{}
		if err = tmp_sl_PSFCH_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PSFCH_Config_r16", err)
		}
		ie.sl_PSFCH_Config_r16 = tmp_sl_PSFCH_Config_r16.Setup
	}
	if sl_SyncAllowed_r16Present {
		ie.sl_SyncAllowed_r16 = new(SL_SyncAllowed_r16)
		if err = ie.sl_SyncAllowed_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_SyncAllowed_r16", err)
		}
	}
	if sl_SubchannelSize_r16Present {
		ie.sl_SubchannelSize_r16 = new(SL_ResourcePool_r16_sl_SubchannelSize_r16)
		if err = ie.sl_SubchannelSize_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_SubchannelSize_r16", err)
		}
	}
	if dummyPresent {
		var tmp_int_dummy int64
		if tmp_int_dummy, err = r.ReadInteger(&uper.Constraint{Lb: 10, Ub: 160}, false); err != nil {
			return utils.WrapError("Decode dummy", err)
		}
		ie.dummy = &tmp_int_dummy
	}
	if sl_StartRB_Subchannel_r16Present {
		var tmp_int_sl_StartRB_Subchannel_r16 int64
		if tmp_int_sl_StartRB_Subchannel_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 265}, false); err != nil {
			return utils.WrapError("Decode sl_StartRB_Subchannel_r16", err)
		}
		ie.sl_StartRB_Subchannel_r16 = &tmp_int_sl_StartRB_Subchannel_r16
	}
	if sl_NumSubchannel_r16Present {
		var tmp_int_sl_NumSubchannel_r16 int64
		if tmp_int_sl_NumSubchannel_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 27}, false); err != nil {
			return utils.WrapError("Decode sl_NumSubchannel_r16", err)
		}
		ie.sl_NumSubchannel_r16 = &tmp_int_sl_NumSubchannel_r16
	}
	if sl_Additional_MCS_Table_r16Present {
		ie.sl_Additional_MCS_Table_r16 = new(SL_ResourcePool_r16_sl_Additional_MCS_Table_r16)
		if err = ie.sl_Additional_MCS_Table_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_Additional_MCS_Table_r16", err)
		}
	}
	if sl_ThreshS_RSSI_CBR_r16Present {
		var tmp_int_sl_ThreshS_RSSI_CBR_r16 int64
		if tmp_int_sl_ThreshS_RSSI_CBR_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 45}, false); err != nil {
			return utils.WrapError("Decode sl_ThreshS_RSSI_CBR_r16", err)
		}
		ie.sl_ThreshS_RSSI_CBR_r16 = &tmp_int_sl_ThreshS_RSSI_CBR_r16
	}
	if sl_TimeWindowSizeCBR_r16Present {
		ie.sl_TimeWindowSizeCBR_r16 = new(SL_ResourcePool_r16_sl_TimeWindowSizeCBR_r16)
		if err = ie.sl_TimeWindowSizeCBR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TimeWindowSizeCBR_r16", err)
		}
	}
	if sl_TimeWindowSizeCR_r16Present {
		ie.sl_TimeWindowSizeCR_r16 = new(SL_ResourcePool_r16_sl_TimeWindowSizeCR_r16)
		if err = ie.sl_TimeWindowSizeCR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TimeWindowSizeCR_r16", err)
		}
	}
	if sl_PTRS_Config_r16Present {
		ie.sl_PTRS_Config_r16 = new(SL_PTRS_Config_r16)
		if err = ie.sl_PTRS_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PTRS_Config_r16", err)
		}
	}
	if sl_UE_SelectedConfigRP_r16Present {
		ie.sl_UE_SelectedConfigRP_r16 = new(SL_UE_SelectedConfigRP_r16)
		if err = ie.sl_UE_SelectedConfigRP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_UE_SelectedConfigRP_r16", err)
		}
	}
	if sl_RxParametersNcell_r16Present {
		ie.sl_RxParametersNcell_r16 = new(Sl_RxParametersNcell_r16)
		if err = ie.sl_RxParametersNcell_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_RxParametersNcell_r16", err)
		}
	}
	if sl_ZoneConfigMCR_List_r16Present {
		tmp_sl_ZoneConfigMCR_List_r16 := utils.NewSequence[*SL_ZoneConfigMCR_r16]([]*SL_ZoneConfigMCR_r16{}, uper.Constraint{Lb: 16, Ub: 16}, false)
		fn_sl_ZoneConfigMCR_List_r16 := func() *SL_ZoneConfigMCR_r16 {
			return new(SL_ZoneConfigMCR_r16)
		}
		if err = tmp_sl_ZoneConfigMCR_List_r16.Decode(r, fn_sl_ZoneConfigMCR_List_r16); err != nil {
			return utils.WrapError("Decode sl_ZoneConfigMCR_List_r16", err)
		}
		ie.sl_ZoneConfigMCR_List_r16 = []SL_ZoneConfigMCR_r16{}
		for _, i := range tmp_sl_ZoneConfigMCR_List_r16.Value {
			ie.sl_ZoneConfigMCR_List_r16 = append(ie.sl_ZoneConfigMCR_List_r16, *i)
		}
	}
	if sl_FilterCoefficient_r16Present {
		ie.sl_FilterCoefficient_r16 = new(FilterCoefficient)
		if err = ie.sl_FilterCoefficient_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_FilterCoefficient_r16", err)
		}
	}
	if sl_RB_Number_r16Present {
		var tmp_int_sl_RB_Number_r16 int64
		if tmp_int_sl_RB_Number_r16, err = r.ReadInteger(&uper.Constraint{Lb: 10, Ub: 275}, false); err != nil {
			return utils.WrapError("Decode sl_RB_Number_r16", err)
		}
		ie.sl_RB_Number_r16 = &tmp_int_sl_RB_Number_r16
	}
	if sl_PreemptionEnable_r16Present {
		ie.sl_PreemptionEnable_r16 = new(SL_ResourcePool_r16_sl_PreemptionEnable_r16)
		if err = ie.sl_PreemptionEnable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PreemptionEnable_r16", err)
		}
	}
	if sl_PriorityThreshold_UL_URLLC_r16Present {
		var tmp_int_sl_PriorityThreshold_UL_URLLC_r16 int64
		if tmp_int_sl_PriorityThreshold_UL_URLLC_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode sl_PriorityThreshold_UL_URLLC_r16", err)
		}
		ie.sl_PriorityThreshold_UL_URLLC_r16 = &tmp_int_sl_PriorityThreshold_UL_URLLC_r16
	}
	if sl_PriorityThreshold_r16Present {
		var tmp_int_sl_PriorityThreshold_r16 int64
		if tmp_int_sl_PriorityThreshold_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 9}, false); err != nil {
			return utils.WrapError("Decode sl_PriorityThreshold_r16", err)
		}
		ie.sl_PriorityThreshold_r16 = &tmp_int_sl_PriorityThreshold_r16
	}
	if sl_X_Overhead_r16Present {
		ie.sl_X_Overhead_r16 = new(SL_ResourcePool_r16_sl_X_Overhead_r16)
		if err = ie.sl_X_Overhead_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_X_Overhead_r16", err)
		}
	}
	if sl_PowerControl_r16Present {
		ie.sl_PowerControl_r16 = new(SL_PowerControl_r16)
		if err = ie.sl_PowerControl_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PowerControl_r16", err)
		}
	}
	if sl_TxPercentageList_r16Present {
		ie.sl_TxPercentageList_r16 = new(SL_TxPercentageList_r16)
		if err = ie.sl_TxPercentageList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TxPercentageList_r16", err)
		}
	}
	if sl_MinMaxMCS_List_r16Present {
		ie.sl_MinMaxMCS_List_r16 = new(SL_MinMaxMCS_List_r16)
		if err = ie.sl_MinMaxMCS_List_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MinMaxMCS_List_r16", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			sl_TimeResource_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sl_TimeResource_r16 optional
			if sl_TimeResource_r16Present {
				var tmp_bs_sl_TimeResource_r16 []byte
				var n_sl_TimeResource_r16 uint
				if tmp_bs_sl_TimeResource_r16, n_sl_TimeResource_r16, err = extReader.ReadBitString(&uper.Constraint{Lb: 10, Ub: 160}, false); err != nil {
					return utils.WrapError("Decode sl_TimeResource_r16", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_sl_TimeResource_r16,
					NumBits: uint64(n_sl_TimeResource_r16),
				}
				ie.sl_TimeResource_r16 = &tmp_bitstring
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			sl_PBPS_CPS_Config_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sl_InterUE_CoordinationConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sl_PBPS_CPS_Config_r17 optional
			if sl_PBPS_CPS_Config_r17Present {
				tmp_sl_PBPS_CPS_Config_r17 := utils.SetupRelease[*SL_PBPS_CPS_Config_r17]{}
				if err = tmp_sl_PBPS_CPS_Config_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_PBPS_CPS_Config_r17", err)
				}
				ie.sl_PBPS_CPS_Config_r17 = tmp_sl_PBPS_CPS_Config_r17.Setup
			}
			// decode sl_InterUE_CoordinationConfig_r17 optional
			if sl_InterUE_CoordinationConfig_r17Present {
				tmp_sl_InterUE_CoordinationConfig_r17 := utils.SetupRelease[*SL_InterUE_CoordinationConfig_r17]{}
				if err = tmp_sl_InterUE_CoordinationConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_InterUE_CoordinationConfig_r17", err)
				}
				ie.sl_InterUE_CoordinationConfig_r17 = tmp_sl_InterUE_CoordinationConfig_r17.Setup
			}
		}
	}
	return nil
}
