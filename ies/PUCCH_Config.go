package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_Config struct {
	resourceSetToAddModList                           []PUCCH_ResourceSet                                  `lb:1,ub:maxNrofPUCCH_ResourceSets,optional`
	resourceSetToReleaseList                          []PUCCH_ResourceSetId                                `lb:1,ub:maxNrofPUCCH_ResourceSets,optional`
	resourceToAddModList                              []PUCCH_Resource                                     `lb:1,ub:maxNrofPUCCH_Resources,optional`
	resourceToReleaseList                             []PUCCH_ResourceId                                   `lb:1,ub:maxNrofPUCCH_Resources,optional`
	format1                                           *PUCCH_FormatConfig                                  `optional,setuprelease`
	format2                                           *PUCCH_FormatConfig                                  `optional,setuprelease`
	format3                                           *PUCCH_FormatConfig                                  `optional,setuprelease`
	format4                                           *PUCCH_FormatConfig                                  `optional,setuprelease`
	schedulingRequestResourceToAddModList             []SchedulingRequestResourceConfig                    `lb:1,ub:maxNrofSR_Resources,optional`
	schedulingRequestResourceToReleaseList            []SchedulingRequestResourceId                        `lb:1,ub:maxNrofSR_Resources,optional`
	multi_CSI_PUCCH_ResourceList                      []PUCCH_ResourceId                                   `lb:1,ub:2,optional`
	dl_DataToUL_ACK                                   []int64                                              `lb:1,ub:8,e_lb:0,e_ub:15,optional`
	spatialRelationInfoToAddModList                   []PUCCH_SpatialRelationInfo                          `lb:1,ub:maxNrofSpatialRelationInfos,optional`
	spatialRelationInfoToReleaseList                  []PUCCH_SpatialRelationInfoId                        `lb:1,ub:maxNrofSpatialRelationInfos,optional`
	pucch_PowerControl                                *PUCCH_PowerControl                                  `optional`
	resourceToAddModListExt_v1610                     []PUCCH_ResourceExt_v1610                            `lb:1,ub:maxNrofPUCCH_Resources,optional,ext-1`
	dl_DataToUL_ACK_r16                               *DL_DataToUL_ACK_r16                                 `optional,ext-1,setuprelease`
	ul_AccessConfigListDCI_1_1_r16                    *UL_AccessConfigListDCI_1_1_r16                      `optional,ext-1,setuprelease`
	subslotLengthForPUCCH_r16                         *PUCCH_Config_subslotLengthForPUCCH_r16              `optional,ext-1`
	dl_DataToUL_ACK_DCI_1_2_r16                       *DL_DataToUL_ACK_DCI_1_2_r16                         `optional,ext-1,setuprelease`
	numberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 *int64                                               `lb:0,ub:3,optional,ext-1`
	dmrs_UplinkTransformPrecodingPUCCH_r16            *PUCCH_Config_dmrs_UplinkTransformPrecodingPUCCH_r16 `optional,ext-1`
	spatialRelationInfoToAddModListSizeExt_v1610      []PUCCH_SpatialRelationInfo                          `lb:1,ub:maxNrofSpatialRelationInfosDiff_r16,optional,ext-1`
	spatialRelationInfoToReleaseListSizeExt_v1610     []PUCCH_SpatialRelationInfoId                        `lb:1,ub:maxNrofSpatialRelationInfosDiff_r16,optional,ext-1`
	spatialRelationInfoToAddModListExt_v1610          []PUCCH_SpatialRelationInfoExt_r16                   `lb:1,ub:maxNrofSpatialRelationInfos_r16,optional,ext-1`
	spatialRelationInfoToReleaseListExt_v1610         []PUCCH_SpatialRelationInfoId_r16                    `lb:1,ub:maxNrofSpatialRelationInfos_r16,optional,ext-1`
	resourceGroupToAddModList_r16                     []PUCCH_ResourceGroup_r16                            `lb:1,ub:maxNrofPUCCH_ResourceGroups_r16,optional,ext-1`
	resourceGroupToReleaseList_r16                    []PUCCH_ResourceGroupId_r16                          `lb:1,ub:maxNrofPUCCH_ResourceGroups_r16,optional,ext-1`
	sps_PUCCH_AN_List_r16                             *SPS_PUCCH_AN_List_r16                               `optional,ext-1,setuprelease`
	schedulingRequestResourceToAddModListExt_v1610    []SchedulingRequestResourceConfigExt_v1610           `lb:1,ub:maxNrofSR_Resources,optional,ext-1`
	format0_r17                                       *PUCCH_FormatConfig                                  `optional,ext-2,setuprelease`
	format2Ext_r17                                    *PUCCH_FormatConfigExt_r17                           `optional,ext-2,setuprelease`
	format3Ext_r17                                    *PUCCH_FormatConfigExt_r17                           `optional,ext-2,setuprelease`
	format4Ext_r17                                    *PUCCH_FormatConfigExt_r17                           `optional,ext-2,setuprelease`
	ul_AccessConfigListDCI_1_2_r17                    *UL_AccessConfigListDCI_1_2_r17                      `optional,ext-2,setuprelease`
	mappingPattern_r17                                *PUCCH_Config_mappingPattern_r17                     `optional,ext-2`
	powerControlSetInfoToAddModList_r17               []PUCCH_PowerControlSetInfo_r17                      `lb:1,ub:maxNrofPowerControlSetInfos_r17,optional,ext-2`
	powerControlSetInfoToReleaseList_r17              []PUCCH_PowerControlSetInfoId_r17                    `lb:1,ub:maxNrofPowerControlSetInfos_r17,optional,ext-2`
	secondTPCFieldDCI_1_1_r17                         *PUCCH_Config_secondTPCFieldDCI_1_1_r17              `optional,ext-2`
	secondTPCFieldDCI_1_2_r17                         *PUCCH_Config_secondTPCFieldDCI_1_2_r17              `optional,ext-2`
	dl_DataToUL_ACK_r17                               *DL_DataToUL_ACK_r17                                 `optional,ext-2,setuprelease`
	dl_DataToUL_ACK_DCI_1_2_r17                       *DL_DataToUL_ACK_DCI_1_2_r17                         `optional,ext-2,setuprelease`
	ul_AccessConfigListDCI_1_1_r17                    *UL_AccessConfigListDCI_1_1_r17                      `optional,ext-2,setuprelease`
	schedulingRequestResourceToAddModListExt_v1700    []SchedulingRequestResourceConfigExt_v1700           `lb:1,ub:maxNrofSR_Resources,optional,ext-2`
	dmrs_BundlingPUCCH_Config_r17                     *DMRS_BundlingPUCCH_Config_r17                       `optional,ext-2,setuprelease`
	dl_DataToUL_ACK_v1700                             *DL_DataToUL_ACK_v1700                               `optional,ext-2,setuprelease`
	dl_DataToUL_ACK_MulticastDCI_Format4_1_r17        *DL_DataToUL_ACK_MulticastDCI_Format4_1_r17          `optional,ext-2,setuprelease`
	sps_PUCCH_AN_ListMulticast_r17                    *SPS_PUCCH_AN_List_r16                               `optional,ext-2,setuprelease`
}

func (ie *PUCCH_Config) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := len(ie.resourceToAddModListExt_v1610) > 0 || ie.dl_DataToUL_ACK_r16 != nil || ie.ul_AccessConfigListDCI_1_1_r16 != nil || ie.subslotLengthForPUCCH_r16 != nil || ie.dl_DataToUL_ACK_DCI_1_2_r16 != nil || ie.numberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 != nil || ie.dmrs_UplinkTransformPrecodingPUCCH_r16 != nil || len(ie.spatialRelationInfoToAddModListSizeExt_v1610) > 0 || len(ie.spatialRelationInfoToReleaseListSizeExt_v1610) > 0 || len(ie.spatialRelationInfoToAddModListExt_v1610) > 0 || len(ie.spatialRelationInfoToReleaseListExt_v1610) > 0 || len(ie.resourceGroupToAddModList_r16) > 0 || len(ie.resourceGroupToReleaseList_r16) > 0 || ie.sps_PUCCH_AN_List_r16 != nil || len(ie.schedulingRequestResourceToAddModListExt_v1610) > 0 || ie.format0_r17 != nil || ie.format2Ext_r17 != nil || ie.format3Ext_r17 != nil || ie.format4Ext_r17 != nil || ie.ul_AccessConfigListDCI_1_2_r17 != nil || ie.mappingPattern_r17 != nil || len(ie.powerControlSetInfoToAddModList_r17) > 0 || len(ie.powerControlSetInfoToReleaseList_r17) > 0 || ie.secondTPCFieldDCI_1_1_r17 != nil || ie.secondTPCFieldDCI_1_2_r17 != nil || ie.dl_DataToUL_ACK_r17 != nil || ie.dl_DataToUL_ACK_DCI_1_2_r17 != nil || ie.ul_AccessConfigListDCI_1_1_r17 != nil || len(ie.schedulingRequestResourceToAddModListExt_v1700) > 0 || ie.dmrs_BundlingPUCCH_Config_r17 != nil || ie.dl_DataToUL_ACK_v1700 != nil || ie.dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 != nil || ie.sps_PUCCH_AN_ListMulticast_r17 != nil
	preambleBits := []bool{hasExtensions, len(ie.resourceSetToAddModList) > 0, len(ie.resourceSetToReleaseList) > 0, len(ie.resourceToAddModList) > 0, len(ie.resourceToReleaseList) > 0, ie.format1 != nil, ie.format2 != nil, ie.format3 != nil, ie.format4 != nil, len(ie.schedulingRequestResourceToAddModList) > 0, len(ie.schedulingRequestResourceToReleaseList) > 0, len(ie.multi_CSI_PUCCH_ResourceList) > 0, len(ie.dl_DataToUL_ACK) > 0, len(ie.spatialRelationInfoToAddModList) > 0, len(ie.spatialRelationInfoToReleaseList) > 0, ie.pucch_PowerControl != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.resourceSetToAddModList) > 0 {
		tmp_resourceSetToAddModList := utils.NewSequence[*PUCCH_ResourceSet]([]*PUCCH_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourceSets}, false)
		for _, i := range ie.resourceSetToAddModList {
			tmp_resourceSetToAddModList.Value = append(tmp_resourceSetToAddModList.Value, &i)
		}
		if err = tmp_resourceSetToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode resourceSetToAddModList", err)
		}
	}
	if len(ie.resourceSetToReleaseList) > 0 {
		tmp_resourceSetToReleaseList := utils.NewSequence[*PUCCH_ResourceSetId]([]*PUCCH_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourceSets}, false)
		for _, i := range ie.resourceSetToReleaseList {
			tmp_resourceSetToReleaseList.Value = append(tmp_resourceSetToReleaseList.Value, &i)
		}
		if err = tmp_resourceSetToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode resourceSetToReleaseList", err)
		}
	}
	if len(ie.resourceToAddModList) > 0 {
		tmp_resourceToAddModList := utils.NewSequence[*PUCCH_Resource]([]*PUCCH_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_Resources}, false)
		for _, i := range ie.resourceToAddModList {
			tmp_resourceToAddModList.Value = append(tmp_resourceToAddModList.Value, &i)
		}
		if err = tmp_resourceToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode resourceToAddModList", err)
		}
	}
	if len(ie.resourceToReleaseList) > 0 {
		tmp_resourceToReleaseList := utils.NewSequence[*PUCCH_ResourceId]([]*PUCCH_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_Resources}, false)
		for _, i := range ie.resourceToReleaseList {
			tmp_resourceToReleaseList.Value = append(tmp_resourceToReleaseList.Value, &i)
		}
		if err = tmp_resourceToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode resourceToReleaseList", err)
		}
	}
	if ie.format1 != nil {
		tmp_format1 := utils.SetupRelease[*PUCCH_FormatConfig]{
			Setup: ie.format1,
		}
		if err = tmp_format1.Encode(w); err != nil {
			return utils.WrapError("Encode format1", err)
		}
	}
	if ie.format2 != nil {
		tmp_format2 := utils.SetupRelease[*PUCCH_FormatConfig]{
			Setup: ie.format2,
		}
		if err = tmp_format2.Encode(w); err != nil {
			return utils.WrapError("Encode format2", err)
		}
	}
	if ie.format3 != nil {
		tmp_format3 := utils.SetupRelease[*PUCCH_FormatConfig]{
			Setup: ie.format3,
		}
		if err = tmp_format3.Encode(w); err != nil {
			return utils.WrapError("Encode format3", err)
		}
	}
	if ie.format4 != nil {
		tmp_format4 := utils.SetupRelease[*PUCCH_FormatConfig]{
			Setup: ie.format4,
		}
		if err = tmp_format4.Encode(w); err != nil {
			return utils.WrapError("Encode format4", err)
		}
	}
	if len(ie.schedulingRequestResourceToAddModList) > 0 {
		tmp_schedulingRequestResourceToAddModList := utils.NewSequence[*SchedulingRequestResourceConfig]([]*SchedulingRequestResourceConfig{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_Resources}, false)
		for _, i := range ie.schedulingRequestResourceToAddModList {
			tmp_schedulingRequestResourceToAddModList.Value = append(tmp_schedulingRequestResourceToAddModList.Value, &i)
		}
		if err = tmp_schedulingRequestResourceToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode schedulingRequestResourceToAddModList", err)
		}
	}
	if len(ie.schedulingRequestResourceToReleaseList) > 0 {
		tmp_schedulingRequestResourceToReleaseList := utils.NewSequence[*SchedulingRequestResourceId]([]*SchedulingRequestResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_Resources}, false)
		for _, i := range ie.schedulingRequestResourceToReleaseList {
			tmp_schedulingRequestResourceToReleaseList.Value = append(tmp_schedulingRequestResourceToReleaseList.Value, &i)
		}
		if err = tmp_schedulingRequestResourceToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode schedulingRequestResourceToReleaseList", err)
		}
	}
	if len(ie.multi_CSI_PUCCH_ResourceList) > 0 {
		tmp_multi_CSI_PUCCH_ResourceList := utils.NewSequence[*PUCCH_ResourceId]([]*PUCCH_ResourceId{}, uper.Constraint{Lb: 1, Ub: 2}, false)
		for _, i := range ie.multi_CSI_PUCCH_ResourceList {
			tmp_multi_CSI_PUCCH_ResourceList.Value = append(tmp_multi_CSI_PUCCH_ResourceList.Value, &i)
		}
		if err = tmp_multi_CSI_PUCCH_ResourceList.Encode(w); err != nil {
			return utils.WrapError("Encode multi_CSI_PUCCH_ResourceList", err)
		}
	}
	if len(ie.dl_DataToUL_ACK) > 0 {
		tmp_dl_DataToUL_ACK := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		for _, i := range ie.dl_DataToUL_ACK {
			tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 15}, false)
			tmp_dl_DataToUL_ACK.Value = append(tmp_dl_DataToUL_ACK.Value, &tmp_ie)
		}
		if err = tmp_dl_DataToUL_ACK.Encode(w); err != nil {
			return utils.WrapError("Encode dl_DataToUL_ACK", err)
		}
	}
	if len(ie.spatialRelationInfoToAddModList) > 0 {
		tmp_spatialRelationInfoToAddModList := utils.NewSequence[*PUCCH_SpatialRelationInfo]([]*PUCCH_SpatialRelationInfo{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfos}, false)
		for _, i := range ie.spatialRelationInfoToAddModList {
			tmp_spatialRelationInfoToAddModList.Value = append(tmp_spatialRelationInfoToAddModList.Value, &i)
		}
		if err = tmp_spatialRelationInfoToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode spatialRelationInfoToAddModList", err)
		}
	}
	if len(ie.spatialRelationInfoToReleaseList) > 0 {
		tmp_spatialRelationInfoToReleaseList := utils.NewSequence[*PUCCH_SpatialRelationInfoId]([]*PUCCH_SpatialRelationInfoId{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfos}, false)
		for _, i := range ie.spatialRelationInfoToReleaseList {
			tmp_spatialRelationInfoToReleaseList.Value = append(tmp_spatialRelationInfoToReleaseList.Value, &i)
		}
		if err = tmp_spatialRelationInfoToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode spatialRelationInfoToReleaseList", err)
		}
	}
	if ie.pucch_PowerControl != nil {
		if err = ie.pucch_PowerControl.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_PowerControl", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{len(ie.resourceToAddModListExt_v1610) > 0 || ie.dl_DataToUL_ACK_r16 != nil || ie.ul_AccessConfigListDCI_1_1_r16 != nil || ie.subslotLengthForPUCCH_r16 != nil || ie.dl_DataToUL_ACK_DCI_1_2_r16 != nil || ie.numberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 != nil || ie.dmrs_UplinkTransformPrecodingPUCCH_r16 != nil || len(ie.spatialRelationInfoToAddModListSizeExt_v1610) > 0 || len(ie.spatialRelationInfoToReleaseListSizeExt_v1610) > 0 || len(ie.spatialRelationInfoToAddModListExt_v1610) > 0 || len(ie.spatialRelationInfoToReleaseListExt_v1610) > 0 || len(ie.resourceGroupToAddModList_r16) > 0 || len(ie.resourceGroupToReleaseList_r16) > 0 || ie.sps_PUCCH_AN_List_r16 != nil || len(ie.schedulingRequestResourceToAddModListExt_v1610) > 0, ie.format0_r17 != nil || ie.format2Ext_r17 != nil || ie.format3Ext_r17 != nil || ie.format4Ext_r17 != nil || ie.ul_AccessConfigListDCI_1_2_r17 != nil || ie.mappingPattern_r17 != nil || len(ie.powerControlSetInfoToAddModList_r17) > 0 || len(ie.powerControlSetInfoToReleaseList_r17) > 0 || ie.secondTPCFieldDCI_1_1_r17 != nil || ie.secondTPCFieldDCI_1_2_r17 != nil || ie.dl_DataToUL_ACK_r17 != nil || ie.dl_DataToUL_ACK_DCI_1_2_r17 != nil || ie.ul_AccessConfigListDCI_1_1_r17 != nil || len(ie.schedulingRequestResourceToAddModListExt_v1700) > 0 || ie.dmrs_BundlingPUCCH_Config_r17 != nil || ie.dl_DataToUL_ACK_v1700 != nil || ie.dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 != nil || ie.sps_PUCCH_AN_ListMulticast_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PUCCH_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.resourceToAddModListExt_v1610) > 0, ie.dl_DataToUL_ACK_r16 != nil, ie.ul_AccessConfigListDCI_1_1_r16 != nil, ie.subslotLengthForPUCCH_r16 != nil, ie.dl_DataToUL_ACK_DCI_1_2_r16 != nil, ie.numberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 != nil, ie.dmrs_UplinkTransformPrecodingPUCCH_r16 != nil, len(ie.spatialRelationInfoToAddModListSizeExt_v1610) > 0, len(ie.spatialRelationInfoToReleaseListSizeExt_v1610) > 0, len(ie.spatialRelationInfoToAddModListExt_v1610) > 0, len(ie.spatialRelationInfoToReleaseListExt_v1610) > 0, len(ie.resourceGroupToAddModList_r16) > 0, len(ie.resourceGroupToReleaseList_r16) > 0, ie.sps_PUCCH_AN_List_r16 != nil, len(ie.schedulingRequestResourceToAddModListExt_v1610) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode resourceToAddModListExt_v1610 optional
			if len(ie.resourceToAddModListExt_v1610) > 0 {
				tmp_resourceToAddModListExt_v1610 := utils.NewSequence[*PUCCH_ResourceExt_v1610]([]*PUCCH_ResourceExt_v1610{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_Resources}, false)
				for _, i := range ie.resourceToAddModListExt_v1610 {
					tmp_resourceToAddModListExt_v1610.Value = append(tmp_resourceToAddModListExt_v1610.Value, &i)
				}
				if err = tmp_resourceToAddModListExt_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode resourceToAddModListExt_v1610", err)
				}
			}
			// encode dl_DataToUL_ACK_r16 optional
			if ie.dl_DataToUL_ACK_r16 != nil {
				tmp_dl_DataToUL_ACK_r16 := utils.SetupRelease[*DL_DataToUL_ACK_r16]{
					Setup: ie.dl_DataToUL_ACK_r16,
				}
				if err = tmp_dl_DataToUL_ACK_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dl_DataToUL_ACK_r16", err)
				}
			}
			// encode ul_AccessConfigListDCI_1_1_r16 optional
			if ie.ul_AccessConfigListDCI_1_1_r16 != nil {
				tmp_ul_AccessConfigListDCI_1_1_r16 := utils.SetupRelease[*UL_AccessConfigListDCI_1_1_r16]{
					Setup: ie.ul_AccessConfigListDCI_1_1_r16,
				}
				if err = tmp_ul_AccessConfigListDCI_1_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_AccessConfigListDCI_1_1_r16", err)
				}
			}
			// encode subslotLengthForPUCCH_r16 optional
			if ie.subslotLengthForPUCCH_r16 != nil {
				if err = ie.subslotLengthForPUCCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode subslotLengthForPUCCH_r16", err)
				}
			}
			// encode dl_DataToUL_ACK_DCI_1_2_r16 optional
			if ie.dl_DataToUL_ACK_DCI_1_2_r16 != nil {
				tmp_dl_DataToUL_ACK_DCI_1_2_r16 := utils.SetupRelease[*DL_DataToUL_ACK_DCI_1_2_r16]{
					Setup: ie.dl_DataToUL_ACK_DCI_1_2_r16,
				}
				if err = tmp_dl_DataToUL_ACK_DCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dl_DataToUL_ACK_DCI_1_2_r16", err)
				}
			}
			// encode numberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 optional
			if ie.numberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.numberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16, &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
					return utils.WrapError("Encode numberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16", err)
				}
			}
			// encode dmrs_UplinkTransformPrecodingPUCCH_r16 optional
			if ie.dmrs_UplinkTransformPrecodingPUCCH_r16 != nil {
				if err = ie.dmrs_UplinkTransformPrecodingPUCCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dmrs_UplinkTransformPrecodingPUCCH_r16", err)
				}
			}
			// encode spatialRelationInfoToAddModListSizeExt_v1610 optional
			if len(ie.spatialRelationInfoToAddModListSizeExt_v1610) > 0 {
				tmp_spatialRelationInfoToAddModListSizeExt_v1610 := utils.NewSequence[*PUCCH_SpatialRelationInfo]([]*PUCCH_SpatialRelationInfo{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfosDiff_r16}, false)
				for _, i := range ie.spatialRelationInfoToAddModListSizeExt_v1610 {
					tmp_spatialRelationInfoToAddModListSizeExt_v1610.Value = append(tmp_spatialRelationInfoToAddModListSizeExt_v1610.Value, &i)
				}
				if err = tmp_spatialRelationInfoToAddModListSizeExt_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode spatialRelationInfoToAddModListSizeExt_v1610", err)
				}
			}
			// encode spatialRelationInfoToReleaseListSizeExt_v1610 optional
			if len(ie.spatialRelationInfoToReleaseListSizeExt_v1610) > 0 {
				tmp_spatialRelationInfoToReleaseListSizeExt_v1610 := utils.NewSequence[*PUCCH_SpatialRelationInfoId]([]*PUCCH_SpatialRelationInfoId{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfosDiff_r16}, false)
				for _, i := range ie.spatialRelationInfoToReleaseListSizeExt_v1610 {
					tmp_spatialRelationInfoToReleaseListSizeExt_v1610.Value = append(tmp_spatialRelationInfoToReleaseListSizeExt_v1610.Value, &i)
				}
				if err = tmp_spatialRelationInfoToReleaseListSizeExt_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode spatialRelationInfoToReleaseListSizeExt_v1610", err)
				}
			}
			// encode spatialRelationInfoToAddModListExt_v1610 optional
			if len(ie.spatialRelationInfoToAddModListExt_v1610) > 0 {
				tmp_spatialRelationInfoToAddModListExt_v1610 := utils.NewSequence[*PUCCH_SpatialRelationInfoExt_r16]([]*PUCCH_SpatialRelationInfoExt_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfos_r16}, false)
				for _, i := range ie.spatialRelationInfoToAddModListExt_v1610 {
					tmp_spatialRelationInfoToAddModListExt_v1610.Value = append(tmp_spatialRelationInfoToAddModListExt_v1610.Value, &i)
				}
				if err = tmp_spatialRelationInfoToAddModListExt_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode spatialRelationInfoToAddModListExt_v1610", err)
				}
			}
			// encode spatialRelationInfoToReleaseListExt_v1610 optional
			if len(ie.spatialRelationInfoToReleaseListExt_v1610) > 0 {
				tmp_spatialRelationInfoToReleaseListExt_v1610 := utils.NewSequence[*PUCCH_SpatialRelationInfoId_r16]([]*PUCCH_SpatialRelationInfoId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfos_r16}, false)
				for _, i := range ie.spatialRelationInfoToReleaseListExt_v1610 {
					tmp_spatialRelationInfoToReleaseListExt_v1610.Value = append(tmp_spatialRelationInfoToReleaseListExt_v1610.Value, &i)
				}
				if err = tmp_spatialRelationInfoToReleaseListExt_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode spatialRelationInfoToReleaseListExt_v1610", err)
				}
			}
			// encode resourceGroupToAddModList_r16 optional
			if len(ie.resourceGroupToAddModList_r16) > 0 {
				tmp_resourceGroupToAddModList_r16 := utils.NewSequence[*PUCCH_ResourceGroup_r16]([]*PUCCH_ResourceGroup_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourceGroups_r16}, false)
				for _, i := range ie.resourceGroupToAddModList_r16 {
					tmp_resourceGroupToAddModList_r16.Value = append(tmp_resourceGroupToAddModList_r16.Value, &i)
				}
				if err = tmp_resourceGroupToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode resourceGroupToAddModList_r16", err)
				}
			}
			// encode resourceGroupToReleaseList_r16 optional
			if len(ie.resourceGroupToReleaseList_r16) > 0 {
				tmp_resourceGroupToReleaseList_r16 := utils.NewSequence[*PUCCH_ResourceGroupId_r16]([]*PUCCH_ResourceGroupId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourceGroups_r16}, false)
				for _, i := range ie.resourceGroupToReleaseList_r16 {
					tmp_resourceGroupToReleaseList_r16.Value = append(tmp_resourceGroupToReleaseList_r16.Value, &i)
				}
				if err = tmp_resourceGroupToReleaseList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode resourceGroupToReleaseList_r16", err)
				}
			}
			// encode sps_PUCCH_AN_List_r16 optional
			if ie.sps_PUCCH_AN_List_r16 != nil {
				tmp_sps_PUCCH_AN_List_r16 := utils.SetupRelease[*SPS_PUCCH_AN_List_r16]{
					Setup: ie.sps_PUCCH_AN_List_r16,
				}
				if err = tmp_sps_PUCCH_AN_List_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sps_PUCCH_AN_List_r16", err)
				}
			}
			// encode schedulingRequestResourceToAddModListExt_v1610 optional
			if len(ie.schedulingRequestResourceToAddModListExt_v1610) > 0 {
				tmp_schedulingRequestResourceToAddModListExt_v1610 := utils.NewSequence[*SchedulingRequestResourceConfigExt_v1610]([]*SchedulingRequestResourceConfigExt_v1610{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_Resources}, false)
				for _, i := range ie.schedulingRequestResourceToAddModListExt_v1610 {
					tmp_schedulingRequestResourceToAddModListExt_v1610.Value = append(tmp_schedulingRequestResourceToAddModListExt_v1610.Value, &i)
				}
				if err = tmp_schedulingRequestResourceToAddModListExt_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode schedulingRequestResourceToAddModListExt_v1610", err)
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
			optionals_ext_2 := []bool{ie.format0_r17 != nil, ie.format2Ext_r17 != nil, ie.format3Ext_r17 != nil, ie.format4Ext_r17 != nil, ie.ul_AccessConfigListDCI_1_2_r17 != nil, ie.mappingPattern_r17 != nil, len(ie.powerControlSetInfoToAddModList_r17) > 0, len(ie.powerControlSetInfoToReleaseList_r17) > 0, ie.secondTPCFieldDCI_1_1_r17 != nil, ie.secondTPCFieldDCI_1_2_r17 != nil, ie.dl_DataToUL_ACK_r17 != nil, ie.dl_DataToUL_ACK_DCI_1_2_r17 != nil, ie.ul_AccessConfigListDCI_1_1_r17 != nil, len(ie.schedulingRequestResourceToAddModListExt_v1700) > 0, ie.dmrs_BundlingPUCCH_Config_r17 != nil, ie.dl_DataToUL_ACK_v1700 != nil, ie.dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 != nil, ie.sps_PUCCH_AN_ListMulticast_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode format0_r17 optional
			if ie.format0_r17 != nil {
				tmp_format0_r17 := utils.SetupRelease[*PUCCH_FormatConfig]{
					Setup: ie.format0_r17,
				}
				if err = tmp_format0_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode format0_r17", err)
				}
			}
			// encode format2Ext_r17 optional
			if ie.format2Ext_r17 != nil {
				tmp_format2Ext_r17 := utils.SetupRelease[*PUCCH_FormatConfigExt_r17]{
					Setup: ie.format2Ext_r17,
				}
				if err = tmp_format2Ext_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode format2Ext_r17", err)
				}
			}
			// encode format3Ext_r17 optional
			if ie.format3Ext_r17 != nil {
				tmp_format3Ext_r17 := utils.SetupRelease[*PUCCH_FormatConfigExt_r17]{
					Setup: ie.format3Ext_r17,
				}
				if err = tmp_format3Ext_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode format3Ext_r17", err)
				}
			}
			// encode format4Ext_r17 optional
			if ie.format4Ext_r17 != nil {
				tmp_format4Ext_r17 := utils.SetupRelease[*PUCCH_FormatConfigExt_r17]{
					Setup: ie.format4Ext_r17,
				}
				if err = tmp_format4Ext_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode format4Ext_r17", err)
				}
			}
			// encode ul_AccessConfigListDCI_1_2_r17 optional
			if ie.ul_AccessConfigListDCI_1_2_r17 != nil {
				tmp_ul_AccessConfigListDCI_1_2_r17 := utils.SetupRelease[*UL_AccessConfigListDCI_1_2_r17]{
					Setup: ie.ul_AccessConfigListDCI_1_2_r17,
				}
				if err = tmp_ul_AccessConfigListDCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_AccessConfigListDCI_1_2_r17", err)
				}
			}
			// encode mappingPattern_r17 optional
			if ie.mappingPattern_r17 != nil {
				if err = ie.mappingPattern_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mappingPattern_r17", err)
				}
			}
			// encode powerControlSetInfoToAddModList_r17 optional
			if len(ie.powerControlSetInfoToAddModList_r17) > 0 {
				tmp_powerControlSetInfoToAddModList_r17 := utils.NewSequence[*PUCCH_PowerControlSetInfo_r17]([]*PUCCH_PowerControlSetInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPowerControlSetInfos_r17}, false)
				for _, i := range ie.powerControlSetInfoToAddModList_r17 {
					tmp_powerControlSetInfoToAddModList_r17.Value = append(tmp_powerControlSetInfoToAddModList_r17.Value, &i)
				}
				if err = tmp_powerControlSetInfoToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode powerControlSetInfoToAddModList_r17", err)
				}
			}
			// encode powerControlSetInfoToReleaseList_r17 optional
			if len(ie.powerControlSetInfoToReleaseList_r17) > 0 {
				tmp_powerControlSetInfoToReleaseList_r17 := utils.NewSequence[*PUCCH_PowerControlSetInfoId_r17]([]*PUCCH_PowerControlSetInfoId_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPowerControlSetInfos_r17}, false)
				for _, i := range ie.powerControlSetInfoToReleaseList_r17 {
					tmp_powerControlSetInfoToReleaseList_r17.Value = append(tmp_powerControlSetInfoToReleaseList_r17.Value, &i)
				}
				if err = tmp_powerControlSetInfoToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode powerControlSetInfoToReleaseList_r17", err)
				}
			}
			// encode secondTPCFieldDCI_1_1_r17 optional
			if ie.secondTPCFieldDCI_1_1_r17 != nil {
				if err = ie.secondTPCFieldDCI_1_1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode secondTPCFieldDCI_1_1_r17", err)
				}
			}
			// encode secondTPCFieldDCI_1_2_r17 optional
			if ie.secondTPCFieldDCI_1_2_r17 != nil {
				if err = ie.secondTPCFieldDCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode secondTPCFieldDCI_1_2_r17", err)
				}
			}
			// encode dl_DataToUL_ACK_r17 optional
			if ie.dl_DataToUL_ACK_r17 != nil {
				tmp_dl_DataToUL_ACK_r17 := utils.SetupRelease[*DL_DataToUL_ACK_r17]{
					Setup: ie.dl_DataToUL_ACK_r17,
				}
				if err = tmp_dl_DataToUL_ACK_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dl_DataToUL_ACK_r17", err)
				}
			}
			// encode dl_DataToUL_ACK_DCI_1_2_r17 optional
			if ie.dl_DataToUL_ACK_DCI_1_2_r17 != nil {
				tmp_dl_DataToUL_ACK_DCI_1_2_r17 := utils.SetupRelease[*DL_DataToUL_ACK_DCI_1_2_r17]{
					Setup: ie.dl_DataToUL_ACK_DCI_1_2_r17,
				}
				if err = tmp_dl_DataToUL_ACK_DCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dl_DataToUL_ACK_DCI_1_2_r17", err)
				}
			}
			// encode ul_AccessConfigListDCI_1_1_r17 optional
			if ie.ul_AccessConfigListDCI_1_1_r17 != nil {
				tmp_ul_AccessConfigListDCI_1_1_r17 := utils.SetupRelease[*UL_AccessConfigListDCI_1_1_r17]{
					Setup: ie.ul_AccessConfigListDCI_1_1_r17,
				}
				if err = tmp_ul_AccessConfigListDCI_1_1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_AccessConfigListDCI_1_1_r17", err)
				}
			}
			// encode schedulingRequestResourceToAddModListExt_v1700 optional
			if len(ie.schedulingRequestResourceToAddModListExt_v1700) > 0 {
				tmp_schedulingRequestResourceToAddModListExt_v1700 := utils.NewSequence[*SchedulingRequestResourceConfigExt_v1700]([]*SchedulingRequestResourceConfigExt_v1700{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_Resources}, false)
				for _, i := range ie.schedulingRequestResourceToAddModListExt_v1700 {
					tmp_schedulingRequestResourceToAddModListExt_v1700.Value = append(tmp_schedulingRequestResourceToAddModListExt_v1700.Value, &i)
				}
				if err = tmp_schedulingRequestResourceToAddModListExt_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode schedulingRequestResourceToAddModListExt_v1700", err)
				}
			}
			// encode dmrs_BundlingPUCCH_Config_r17 optional
			if ie.dmrs_BundlingPUCCH_Config_r17 != nil {
				tmp_dmrs_BundlingPUCCH_Config_r17 := utils.SetupRelease[*DMRS_BundlingPUCCH_Config_r17]{
					Setup: ie.dmrs_BundlingPUCCH_Config_r17,
				}
				if err = tmp_dmrs_BundlingPUCCH_Config_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dmrs_BundlingPUCCH_Config_r17", err)
				}
			}
			// encode dl_DataToUL_ACK_v1700 optional
			if ie.dl_DataToUL_ACK_v1700 != nil {
				tmp_dl_DataToUL_ACK_v1700 := utils.SetupRelease[*DL_DataToUL_ACK_v1700]{
					Setup: ie.dl_DataToUL_ACK_v1700,
				}
				if err = tmp_dl_DataToUL_ACK_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dl_DataToUL_ACK_v1700", err)
				}
			}
			// encode dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 optional
			if ie.dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 != nil {
				tmp_dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 := utils.SetupRelease[*DL_DataToUL_ACK_MulticastDCI_Format4_1_r17]{
					Setup: ie.dl_DataToUL_ACK_MulticastDCI_Format4_1_r17,
				}
				if err = tmp_dl_DataToUL_ACK_MulticastDCI_Format4_1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dl_DataToUL_ACK_MulticastDCI_Format4_1_r17", err)
				}
			}
			// encode sps_PUCCH_AN_ListMulticast_r17 optional
			if ie.sps_PUCCH_AN_ListMulticast_r17 != nil {
				tmp_sps_PUCCH_AN_ListMulticast_r17 := utils.SetupRelease[*SPS_PUCCH_AN_List_r16]{
					Setup: ie.sps_PUCCH_AN_ListMulticast_r17,
				}
				if err = tmp_sps_PUCCH_AN_ListMulticast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sps_PUCCH_AN_ListMulticast_r17", err)
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

func (ie *PUCCH_Config) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var resourceSetToAddModListPresent bool
	if resourceSetToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var resourceSetToReleaseListPresent bool
	if resourceSetToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var resourceToAddModListPresent bool
	if resourceToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var resourceToReleaseListPresent bool
	if resourceToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var format1Present bool
	if format1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var format2Present bool
	if format2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var format3Present bool
	if format3Present, err = r.ReadBool(); err != nil {
		return err
	}
	var format4Present bool
	if format4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var schedulingRequestResourceToAddModListPresent bool
	if schedulingRequestResourceToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var schedulingRequestResourceToReleaseListPresent bool
	if schedulingRequestResourceToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var multi_CSI_PUCCH_ResourceListPresent bool
	if multi_CSI_PUCCH_ResourceListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_DataToUL_ACKPresent bool
	if dl_DataToUL_ACKPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var spatialRelationInfoToAddModListPresent bool
	if spatialRelationInfoToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var spatialRelationInfoToReleaseListPresent bool
	if spatialRelationInfoToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_PowerControlPresent bool
	if pucch_PowerControlPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if resourceSetToAddModListPresent {
		tmp_resourceSetToAddModList := utils.NewSequence[*PUCCH_ResourceSet]([]*PUCCH_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourceSets}, false)
		fn_resourceSetToAddModList := func() *PUCCH_ResourceSet {
			return new(PUCCH_ResourceSet)
		}
		if err = tmp_resourceSetToAddModList.Decode(r, fn_resourceSetToAddModList); err != nil {
			return utils.WrapError("Decode resourceSetToAddModList", err)
		}
		ie.resourceSetToAddModList = []PUCCH_ResourceSet{}
		for _, i := range tmp_resourceSetToAddModList.Value {
			ie.resourceSetToAddModList = append(ie.resourceSetToAddModList, *i)
		}
	}
	if resourceSetToReleaseListPresent {
		tmp_resourceSetToReleaseList := utils.NewSequence[*PUCCH_ResourceSetId]([]*PUCCH_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourceSets}, false)
		fn_resourceSetToReleaseList := func() *PUCCH_ResourceSetId {
			return new(PUCCH_ResourceSetId)
		}
		if err = tmp_resourceSetToReleaseList.Decode(r, fn_resourceSetToReleaseList); err != nil {
			return utils.WrapError("Decode resourceSetToReleaseList", err)
		}
		ie.resourceSetToReleaseList = []PUCCH_ResourceSetId{}
		for _, i := range tmp_resourceSetToReleaseList.Value {
			ie.resourceSetToReleaseList = append(ie.resourceSetToReleaseList, *i)
		}
	}
	if resourceToAddModListPresent {
		tmp_resourceToAddModList := utils.NewSequence[*PUCCH_Resource]([]*PUCCH_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_Resources}, false)
		fn_resourceToAddModList := func() *PUCCH_Resource {
			return new(PUCCH_Resource)
		}
		if err = tmp_resourceToAddModList.Decode(r, fn_resourceToAddModList); err != nil {
			return utils.WrapError("Decode resourceToAddModList", err)
		}
		ie.resourceToAddModList = []PUCCH_Resource{}
		for _, i := range tmp_resourceToAddModList.Value {
			ie.resourceToAddModList = append(ie.resourceToAddModList, *i)
		}
	}
	if resourceToReleaseListPresent {
		tmp_resourceToReleaseList := utils.NewSequence[*PUCCH_ResourceId]([]*PUCCH_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_Resources}, false)
		fn_resourceToReleaseList := func() *PUCCH_ResourceId {
			return new(PUCCH_ResourceId)
		}
		if err = tmp_resourceToReleaseList.Decode(r, fn_resourceToReleaseList); err != nil {
			return utils.WrapError("Decode resourceToReleaseList", err)
		}
		ie.resourceToReleaseList = []PUCCH_ResourceId{}
		for _, i := range tmp_resourceToReleaseList.Value {
			ie.resourceToReleaseList = append(ie.resourceToReleaseList, *i)
		}
	}
	if format1Present {
		tmp_format1 := utils.SetupRelease[*PUCCH_FormatConfig]{}
		if err = tmp_format1.Decode(r); err != nil {
			return utils.WrapError("Decode format1", err)
		}
		ie.format1 = tmp_format1.Setup
	}
	if format2Present {
		tmp_format2 := utils.SetupRelease[*PUCCH_FormatConfig]{}
		if err = tmp_format2.Decode(r); err != nil {
			return utils.WrapError("Decode format2", err)
		}
		ie.format2 = tmp_format2.Setup
	}
	if format3Present {
		tmp_format3 := utils.SetupRelease[*PUCCH_FormatConfig]{}
		if err = tmp_format3.Decode(r); err != nil {
			return utils.WrapError("Decode format3", err)
		}
		ie.format3 = tmp_format3.Setup
	}
	if format4Present {
		tmp_format4 := utils.SetupRelease[*PUCCH_FormatConfig]{}
		if err = tmp_format4.Decode(r); err != nil {
			return utils.WrapError("Decode format4", err)
		}
		ie.format4 = tmp_format4.Setup
	}
	if schedulingRequestResourceToAddModListPresent {
		tmp_schedulingRequestResourceToAddModList := utils.NewSequence[*SchedulingRequestResourceConfig]([]*SchedulingRequestResourceConfig{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_Resources}, false)
		fn_schedulingRequestResourceToAddModList := func() *SchedulingRequestResourceConfig {
			return new(SchedulingRequestResourceConfig)
		}
		if err = tmp_schedulingRequestResourceToAddModList.Decode(r, fn_schedulingRequestResourceToAddModList); err != nil {
			return utils.WrapError("Decode schedulingRequestResourceToAddModList", err)
		}
		ie.schedulingRequestResourceToAddModList = []SchedulingRequestResourceConfig{}
		for _, i := range tmp_schedulingRequestResourceToAddModList.Value {
			ie.schedulingRequestResourceToAddModList = append(ie.schedulingRequestResourceToAddModList, *i)
		}
	}
	if schedulingRequestResourceToReleaseListPresent {
		tmp_schedulingRequestResourceToReleaseList := utils.NewSequence[*SchedulingRequestResourceId]([]*SchedulingRequestResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_Resources}, false)
		fn_schedulingRequestResourceToReleaseList := func() *SchedulingRequestResourceId {
			return new(SchedulingRequestResourceId)
		}
		if err = tmp_schedulingRequestResourceToReleaseList.Decode(r, fn_schedulingRequestResourceToReleaseList); err != nil {
			return utils.WrapError("Decode schedulingRequestResourceToReleaseList", err)
		}
		ie.schedulingRequestResourceToReleaseList = []SchedulingRequestResourceId{}
		for _, i := range tmp_schedulingRequestResourceToReleaseList.Value {
			ie.schedulingRequestResourceToReleaseList = append(ie.schedulingRequestResourceToReleaseList, *i)
		}
	}
	if multi_CSI_PUCCH_ResourceListPresent {
		tmp_multi_CSI_PUCCH_ResourceList := utils.NewSequence[*PUCCH_ResourceId]([]*PUCCH_ResourceId{}, uper.Constraint{Lb: 1, Ub: 2}, false)
		fn_multi_CSI_PUCCH_ResourceList := func() *PUCCH_ResourceId {
			return new(PUCCH_ResourceId)
		}
		if err = tmp_multi_CSI_PUCCH_ResourceList.Decode(r, fn_multi_CSI_PUCCH_ResourceList); err != nil {
			return utils.WrapError("Decode multi_CSI_PUCCH_ResourceList", err)
		}
		ie.multi_CSI_PUCCH_ResourceList = []PUCCH_ResourceId{}
		for _, i := range tmp_multi_CSI_PUCCH_ResourceList.Value {
			ie.multi_CSI_PUCCH_ResourceList = append(ie.multi_CSI_PUCCH_ResourceList, *i)
		}
	}
	if dl_DataToUL_ACKPresent {
		tmp_dl_DataToUL_ACK := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: 8}, false)
		fn_dl_DataToUL_ACK := func() *utils.INTEGER {
			ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 15}, false)
			return &ie
		}
		if err = tmp_dl_DataToUL_ACK.Decode(r, fn_dl_DataToUL_ACK); err != nil {
			return utils.WrapError("Decode dl_DataToUL_ACK", err)
		}
		ie.dl_DataToUL_ACK = []int64{}
		for _, i := range tmp_dl_DataToUL_ACK.Value {
			ie.dl_DataToUL_ACK = append(ie.dl_DataToUL_ACK, int64(i.Value))
		}
	}
	if spatialRelationInfoToAddModListPresent {
		tmp_spatialRelationInfoToAddModList := utils.NewSequence[*PUCCH_SpatialRelationInfo]([]*PUCCH_SpatialRelationInfo{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfos}, false)
		fn_spatialRelationInfoToAddModList := func() *PUCCH_SpatialRelationInfo {
			return new(PUCCH_SpatialRelationInfo)
		}
		if err = tmp_spatialRelationInfoToAddModList.Decode(r, fn_spatialRelationInfoToAddModList); err != nil {
			return utils.WrapError("Decode spatialRelationInfoToAddModList", err)
		}
		ie.spatialRelationInfoToAddModList = []PUCCH_SpatialRelationInfo{}
		for _, i := range tmp_spatialRelationInfoToAddModList.Value {
			ie.spatialRelationInfoToAddModList = append(ie.spatialRelationInfoToAddModList, *i)
		}
	}
	if spatialRelationInfoToReleaseListPresent {
		tmp_spatialRelationInfoToReleaseList := utils.NewSequence[*PUCCH_SpatialRelationInfoId]([]*PUCCH_SpatialRelationInfoId{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfos}, false)
		fn_spatialRelationInfoToReleaseList := func() *PUCCH_SpatialRelationInfoId {
			return new(PUCCH_SpatialRelationInfoId)
		}
		if err = tmp_spatialRelationInfoToReleaseList.Decode(r, fn_spatialRelationInfoToReleaseList); err != nil {
			return utils.WrapError("Decode spatialRelationInfoToReleaseList", err)
		}
		ie.spatialRelationInfoToReleaseList = []PUCCH_SpatialRelationInfoId{}
		for _, i := range tmp_spatialRelationInfoToReleaseList.Value {
			ie.spatialRelationInfoToReleaseList = append(ie.spatialRelationInfoToReleaseList, *i)
		}
	}
	if pucch_PowerControlPresent {
		ie.pucch_PowerControl = new(PUCCH_PowerControl)
		if err = ie.pucch_PowerControl.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_PowerControl", err)
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

			resourceToAddModListExt_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dl_DataToUL_ACK_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_AccessConfigListDCI_1_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			subslotLengthForPUCCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dl_DataToUL_ACK_DCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			numberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dmrs_UplinkTransformPrecodingPUCCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			spatialRelationInfoToAddModListSizeExt_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			spatialRelationInfoToReleaseListSizeExt_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			spatialRelationInfoToAddModListExt_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			spatialRelationInfoToReleaseListExt_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			resourceGroupToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			resourceGroupToReleaseList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sps_PUCCH_AN_List_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			schedulingRequestResourceToAddModListExt_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode resourceToAddModListExt_v1610 optional
			if resourceToAddModListExt_v1610Present {
				tmp_resourceToAddModListExt_v1610 := utils.NewSequence[*PUCCH_ResourceExt_v1610]([]*PUCCH_ResourceExt_v1610{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_Resources}, false)
				fn_resourceToAddModListExt_v1610 := func() *PUCCH_ResourceExt_v1610 {
					return new(PUCCH_ResourceExt_v1610)
				}
				if err = tmp_resourceToAddModListExt_v1610.Decode(extReader, fn_resourceToAddModListExt_v1610); err != nil {
					return utils.WrapError("Decode resourceToAddModListExt_v1610", err)
				}
				ie.resourceToAddModListExt_v1610 = []PUCCH_ResourceExt_v1610{}
				for _, i := range tmp_resourceToAddModListExt_v1610.Value {
					ie.resourceToAddModListExt_v1610 = append(ie.resourceToAddModListExt_v1610, *i)
				}
			}
			// decode dl_DataToUL_ACK_r16 optional
			if dl_DataToUL_ACK_r16Present {
				tmp_dl_DataToUL_ACK_r16 := utils.SetupRelease[*DL_DataToUL_ACK_r16]{}
				if err = tmp_dl_DataToUL_ACK_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode dl_DataToUL_ACK_r16", err)
				}
				ie.dl_DataToUL_ACK_r16 = tmp_dl_DataToUL_ACK_r16.Setup
			}
			// decode ul_AccessConfigListDCI_1_1_r16 optional
			if ul_AccessConfigListDCI_1_1_r16Present {
				tmp_ul_AccessConfigListDCI_1_1_r16 := utils.SetupRelease[*UL_AccessConfigListDCI_1_1_r16]{}
				if err = tmp_ul_AccessConfigListDCI_1_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_AccessConfigListDCI_1_1_r16", err)
				}
				ie.ul_AccessConfigListDCI_1_1_r16 = tmp_ul_AccessConfigListDCI_1_1_r16.Setup
			}
			// decode subslotLengthForPUCCH_r16 optional
			if subslotLengthForPUCCH_r16Present {
				ie.subslotLengthForPUCCH_r16 = new(PUCCH_Config_subslotLengthForPUCCH_r16)
				if err = ie.subslotLengthForPUCCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode subslotLengthForPUCCH_r16", err)
				}
			}
			// decode dl_DataToUL_ACK_DCI_1_2_r16 optional
			if dl_DataToUL_ACK_DCI_1_2_r16Present {
				tmp_dl_DataToUL_ACK_DCI_1_2_r16 := utils.SetupRelease[*DL_DataToUL_ACK_DCI_1_2_r16]{}
				if err = tmp_dl_DataToUL_ACK_DCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode dl_DataToUL_ACK_DCI_1_2_r16", err)
				}
				ie.dl_DataToUL_ACK_DCI_1_2_r16 = tmp_dl_DataToUL_ACK_DCI_1_2_r16.Setup
			}
			// decode numberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 optional
			if numberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16Present {
				var tmp_int_numberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 int64
				if tmp_int_numberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
					return utils.WrapError("Decode numberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16", err)
				}
				ie.numberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16 = &tmp_int_numberOfBitsForPUCCH_ResourceIndicatorDCI_1_2_r16
			}
			// decode dmrs_UplinkTransformPrecodingPUCCH_r16 optional
			if dmrs_UplinkTransformPrecodingPUCCH_r16Present {
				ie.dmrs_UplinkTransformPrecodingPUCCH_r16 = new(PUCCH_Config_dmrs_UplinkTransformPrecodingPUCCH_r16)
				if err = ie.dmrs_UplinkTransformPrecodingPUCCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode dmrs_UplinkTransformPrecodingPUCCH_r16", err)
				}
			}
			// decode spatialRelationInfoToAddModListSizeExt_v1610 optional
			if spatialRelationInfoToAddModListSizeExt_v1610Present {
				tmp_spatialRelationInfoToAddModListSizeExt_v1610 := utils.NewSequence[*PUCCH_SpatialRelationInfo]([]*PUCCH_SpatialRelationInfo{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfosDiff_r16}, false)
				fn_spatialRelationInfoToAddModListSizeExt_v1610 := func() *PUCCH_SpatialRelationInfo {
					return new(PUCCH_SpatialRelationInfo)
				}
				if err = tmp_spatialRelationInfoToAddModListSizeExt_v1610.Decode(extReader, fn_spatialRelationInfoToAddModListSizeExt_v1610); err != nil {
					return utils.WrapError("Decode spatialRelationInfoToAddModListSizeExt_v1610", err)
				}
				ie.spatialRelationInfoToAddModListSizeExt_v1610 = []PUCCH_SpatialRelationInfo{}
				for _, i := range tmp_spatialRelationInfoToAddModListSizeExt_v1610.Value {
					ie.spatialRelationInfoToAddModListSizeExt_v1610 = append(ie.spatialRelationInfoToAddModListSizeExt_v1610, *i)
				}
			}
			// decode spatialRelationInfoToReleaseListSizeExt_v1610 optional
			if spatialRelationInfoToReleaseListSizeExt_v1610Present {
				tmp_spatialRelationInfoToReleaseListSizeExt_v1610 := utils.NewSequence[*PUCCH_SpatialRelationInfoId]([]*PUCCH_SpatialRelationInfoId{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfosDiff_r16}, false)
				fn_spatialRelationInfoToReleaseListSizeExt_v1610 := func() *PUCCH_SpatialRelationInfoId {
					return new(PUCCH_SpatialRelationInfoId)
				}
				if err = tmp_spatialRelationInfoToReleaseListSizeExt_v1610.Decode(extReader, fn_spatialRelationInfoToReleaseListSizeExt_v1610); err != nil {
					return utils.WrapError("Decode spatialRelationInfoToReleaseListSizeExt_v1610", err)
				}
				ie.spatialRelationInfoToReleaseListSizeExt_v1610 = []PUCCH_SpatialRelationInfoId{}
				for _, i := range tmp_spatialRelationInfoToReleaseListSizeExt_v1610.Value {
					ie.spatialRelationInfoToReleaseListSizeExt_v1610 = append(ie.spatialRelationInfoToReleaseListSizeExt_v1610, *i)
				}
			}
			// decode spatialRelationInfoToAddModListExt_v1610 optional
			if spatialRelationInfoToAddModListExt_v1610Present {
				tmp_spatialRelationInfoToAddModListExt_v1610 := utils.NewSequence[*PUCCH_SpatialRelationInfoExt_r16]([]*PUCCH_SpatialRelationInfoExt_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfos_r16}, false)
				fn_spatialRelationInfoToAddModListExt_v1610 := func() *PUCCH_SpatialRelationInfoExt_r16 {
					return new(PUCCH_SpatialRelationInfoExt_r16)
				}
				if err = tmp_spatialRelationInfoToAddModListExt_v1610.Decode(extReader, fn_spatialRelationInfoToAddModListExt_v1610); err != nil {
					return utils.WrapError("Decode spatialRelationInfoToAddModListExt_v1610", err)
				}
				ie.spatialRelationInfoToAddModListExt_v1610 = []PUCCH_SpatialRelationInfoExt_r16{}
				for _, i := range tmp_spatialRelationInfoToAddModListExt_v1610.Value {
					ie.spatialRelationInfoToAddModListExt_v1610 = append(ie.spatialRelationInfoToAddModListExt_v1610, *i)
				}
			}
			// decode spatialRelationInfoToReleaseListExt_v1610 optional
			if spatialRelationInfoToReleaseListExt_v1610Present {
				tmp_spatialRelationInfoToReleaseListExt_v1610 := utils.NewSequence[*PUCCH_SpatialRelationInfoId_r16]([]*PUCCH_SpatialRelationInfoId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofSpatialRelationInfos_r16}, false)
				fn_spatialRelationInfoToReleaseListExt_v1610 := func() *PUCCH_SpatialRelationInfoId_r16 {
					return new(PUCCH_SpatialRelationInfoId_r16)
				}
				if err = tmp_spatialRelationInfoToReleaseListExt_v1610.Decode(extReader, fn_spatialRelationInfoToReleaseListExt_v1610); err != nil {
					return utils.WrapError("Decode spatialRelationInfoToReleaseListExt_v1610", err)
				}
				ie.spatialRelationInfoToReleaseListExt_v1610 = []PUCCH_SpatialRelationInfoId_r16{}
				for _, i := range tmp_spatialRelationInfoToReleaseListExt_v1610.Value {
					ie.spatialRelationInfoToReleaseListExt_v1610 = append(ie.spatialRelationInfoToReleaseListExt_v1610, *i)
				}
			}
			// decode resourceGroupToAddModList_r16 optional
			if resourceGroupToAddModList_r16Present {
				tmp_resourceGroupToAddModList_r16 := utils.NewSequence[*PUCCH_ResourceGroup_r16]([]*PUCCH_ResourceGroup_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourceGroups_r16}, false)
				fn_resourceGroupToAddModList_r16 := func() *PUCCH_ResourceGroup_r16 {
					return new(PUCCH_ResourceGroup_r16)
				}
				if err = tmp_resourceGroupToAddModList_r16.Decode(extReader, fn_resourceGroupToAddModList_r16); err != nil {
					return utils.WrapError("Decode resourceGroupToAddModList_r16", err)
				}
				ie.resourceGroupToAddModList_r16 = []PUCCH_ResourceGroup_r16{}
				for _, i := range tmp_resourceGroupToAddModList_r16.Value {
					ie.resourceGroupToAddModList_r16 = append(ie.resourceGroupToAddModList_r16, *i)
				}
			}
			// decode resourceGroupToReleaseList_r16 optional
			if resourceGroupToReleaseList_r16Present {
				tmp_resourceGroupToReleaseList_r16 := utils.NewSequence[*PUCCH_ResourceGroupId_r16]([]*PUCCH_ResourceGroupId_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourceGroups_r16}, false)
				fn_resourceGroupToReleaseList_r16 := func() *PUCCH_ResourceGroupId_r16 {
					return new(PUCCH_ResourceGroupId_r16)
				}
				if err = tmp_resourceGroupToReleaseList_r16.Decode(extReader, fn_resourceGroupToReleaseList_r16); err != nil {
					return utils.WrapError("Decode resourceGroupToReleaseList_r16", err)
				}
				ie.resourceGroupToReleaseList_r16 = []PUCCH_ResourceGroupId_r16{}
				for _, i := range tmp_resourceGroupToReleaseList_r16.Value {
					ie.resourceGroupToReleaseList_r16 = append(ie.resourceGroupToReleaseList_r16, *i)
				}
			}
			// decode sps_PUCCH_AN_List_r16 optional
			if sps_PUCCH_AN_List_r16Present {
				tmp_sps_PUCCH_AN_List_r16 := utils.SetupRelease[*SPS_PUCCH_AN_List_r16]{}
				if err = tmp_sps_PUCCH_AN_List_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sps_PUCCH_AN_List_r16", err)
				}
				ie.sps_PUCCH_AN_List_r16 = tmp_sps_PUCCH_AN_List_r16.Setup
			}
			// decode schedulingRequestResourceToAddModListExt_v1610 optional
			if schedulingRequestResourceToAddModListExt_v1610Present {
				tmp_schedulingRequestResourceToAddModListExt_v1610 := utils.NewSequence[*SchedulingRequestResourceConfigExt_v1610]([]*SchedulingRequestResourceConfigExt_v1610{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_Resources}, false)
				fn_schedulingRequestResourceToAddModListExt_v1610 := func() *SchedulingRequestResourceConfigExt_v1610 {
					return new(SchedulingRequestResourceConfigExt_v1610)
				}
				if err = tmp_schedulingRequestResourceToAddModListExt_v1610.Decode(extReader, fn_schedulingRequestResourceToAddModListExt_v1610); err != nil {
					return utils.WrapError("Decode schedulingRequestResourceToAddModListExt_v1610", err)
				}
				ie.schedulingRequestResourceToAddModListExt_v1610 = []SchedulingRequestResourceConfigExt_v1610{}
				for _, i := range tmp_schedulingRequestResourceToAddModListExt_v1610.Value {
					ie.schedulingRequestResourceToAddModListExt_v1610 = append(ie.schedulingRequestResourceToAddModListExt_v1610, *i)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			format0_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			format2Ext_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			format3Ext_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			format4Ext_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_AccessConfigListDCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mappingPattern_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			powerControlSetInfoToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			powerControlSetInfoToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			secondTPCFieldDCI_1_1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			secondTPCFieldDCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dl_DataToUL_ACK_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dl_DataToUL_ACK_DCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_AccessConfigListDCI_1_1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			schedulingRequestResourceToAddModListExt_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dmrs_BundlingPUCCH_Config_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dl_DataToUL_ACK_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dl_DataToUL_ACK_MulticastDCI_Format4_1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sps_PUCCH_AN_ListMulticast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode format0_r17 optional
			if format0_r17Present {
				tmp_format0_r17 := utils.SetupRelease[*PUCCH_FormatConfig]{}
				if err = tmp_format0_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode format0_r17", err)
				}
				ie.format0_r17 = tmp_format0_r17.Setup
			}
			// decode format2Ext_r17 optional
			if format2Ext_r17Present {
				tmp_format2Ext_r17 := utils.SetupRelease[*PUCCH_FormatConfigExt_r17]{}
				if err = tmp_format2Ext_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode format2Ext_r17", err)
				}
				ie.format2Ext_r17 = tmp_format2Ext_r17.Setup
			}
			// decode format3Ext_r17 optional
			if format3Ext_r17Present {
				tmp_format3Ext_r17 := utils.SetupRelease[*PUCCH_FormatConfigExt_r17]{}
				if err = tmp_format3Ext_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode format3Ext_r17", err)
				}
				ie.format3Ext_r17 = tmp_format3Ext_r17.Setup
			}
			// decode format4Ext_r17 optional
			if format4Ext_r17Present {
				tmp_format4Ext_r17 := utils.SetupRelease[*PUCCH_FormatConfigExt_r17]{}
				if err = tmp_format4Ext_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode format4Ext_r17", err)
				}
				ie.format4Ext_r17 = tmp_format4Ext_r17.Setup
			}
			// decode ul_AccessConfigListDCI_1_2_r17 optional
			if ul_AccessConfigListDCI_1_2_r17Present {
				tmp_ul_AccessConfigListDCI_1_2_r17 := utils.SetupRelease[*UL_AccessConfigListDCI_1_2_r17]{}
				if err = tmp_ul_AccessConfigListDCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_AccessConfigListDCI_1_2_r17", err)
				}
				ie.ul_AccessConfigListDCI_1_2_r17 = tmp_ul_AccessConfigListDCI_1_2_r17.Setup
			}
			// decode mappingPattern_r17 optional
			if mappingPattern_r17Present {
				ie.mappingPattern_r17 = new(PUCCH_Config_mappingPattern_r17)
				if err = ie.mappingPattern_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mappingPattern_r17", err)
				}
			}
			// decode powerControlSetInfoToAddModList_r17 optional
			if powerControlSetInfoToAddModList_r17Present {
				tmp_powerControlSetInfoToAddModList_r17 := utils.NewSequence[*PUCCH_PowerControlSetInfo_r17]([]*PUCCH_PowerControlSetInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPowerControlSetInfos_r17}, false)
				fn_powerControlSetInfoToAddModList_r17 := func() *PUCCH_PowerControlSetInfo_r17 {
					return new(PUCCH_PowerControlSetInfo_r17)
				}
				if err = tmp_powerControlSetInfoToAddModList_r17.Decode(extReader, fn_powerControlSetInfoToAddModList_r17); err != nil {
					return utils.WrapError("Decode powerControlSetInfoToAddModList_r17", err)
				}
				ie.powerControlSetInfoToAddModList_r17 = []PUCCH_PowerControlSetInfo_r17{}
				for _, i := range tmp_powerControlSetInfoToAddModList_r17.Value {
					ie.powerControlSetInfoToAddModList_r17 = append(ie.powerControlSetInfoToAddModList_r17, *i)
				}
			}
			// decode powerControlSetInfoToReleaseList_r17 optional
			if powerControlSetInfoToReleaseList_r17Present {
				tmp_powerControlSetInfoToReleaseList_r17 := utils.NewSequence[*PUCCH_PowerControlSetInfoId_r17]([]*PUCCH_PowerControlSetInfoId_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPowerControlSetInfos_r17}, false)
				fn_powerControlSetInfoToReleaseList_r17 := func() *PUCCH_PowerControlSetInfoId_r17 {
					return new(PUCCH_PowerControlSetInfoId_r17)
				}
				if err = tmp_powerControlSetInfoToReleaseList_r17.Decode(extReader, fn_powerControlSetInfoToReleaseList_r17); err != nil {
					return utils.WrapError("Decode powerControlSetInfoToReleaseList_r17", err)
				}
				ie.powerControlSetInfoToReleaseList_r17 = []PUCCH_PowerControlSetInfoId_r17{}
				for _, i := range tmp_powerControlSetInfoToReleaseList_r17.Value {
					ie.powerControlSetInfoToReleaseList_r17 = append(ie.powerControlSetInfoToReleaseList_r17, *i)
				}
			}
			// decode secondTPCFieldDCI_1_1_r17 optional
			if secondTPCFieldDCI_1_1_r17Present {
				ie.secondTPCFieldDCI_1_1_r17 = new(PUCCH_Config_secondTPCFieldDCI_1_1_r17)
				if err = ie.secondTPCFieldDCI_1_1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode secondTPCFieldDCI_1_1_r17", err)
				}
			}
			// decode secondTPCFieldDCI_1_2_r17 optional
			if secondTPCFieldDCI_1_2_r17Present {
				ie.secondTPCFieldDCI_1_2_r17 = new(PUCCH_Config_secondTPCFieldDCI_1_2_r17)
				if err = ie.secondTPCFieldDCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode secondTPCFieldDCI_1_2_r17", err)
				}
			}
			// decode dl_DataToUL_ACK_r17 optional
			if dl_DataToUL_ACK_r17Present {
				tmp_dl_DataToUL_ACK_r17 := utils.SetupRelease[*DL_DataToUL_ACK_r17]{}
				if err = tmp_dl_DataToUL_ACK_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dl_DataToUL_ACK_r17", err)
				}
				ie.dl_DataToUL_ACK_r17 = tmp_dl_DataToUL_ACK_r17.Setup
			}
			// decode dl_DataToUL_ACK_DCI_1_2_r17 optional
			if dl_DataToUL_ACK_DCI_1_2_r17Present {
				tmp_dl_DataToUL_ACK_DCI_1_2_r17 := utils.SetupRelease[*DL_DataToUL_ACK_DCI_1_2_r17]{}
				if err = tmp_dl_DataToUL_ACK_DCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dl_DataToUL_ACK_DCI_1_2_r17", err)
				}
				ie.dl_DataToUL_ACK_DCI_1_2_r17 = tmp_dl_DataToUL_ACK_DCI_1_2_r17.Setup
			}
			// decode ul_AccessConfigListDCI_1_1_r17 optional
			if ul_AccessConfigListDCI_1_1_r17Present {
				tmp_ul_AccessConfigListDCI_1_1_r17 := utils.SetupRelease[*UL_AccessConfigListDCI_1_1_r17]{}
				if err = tmp_ul_AccessConfigListDCI_1_1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_AccessConfigListDCI_1_1_r17", err)
				}
				ie.ul_AccessConfigListDCI_1_1_r17 = tmp_ul_AccessConfigListDCI_1_1_r17.Setup
			}
			// decode schedulingRequestResourceToAddModListExt_v1700 optional
			if schedulingRequestResourceToAddModListExt_v1700Present {
				tmp_schedulingRequestResourceToAddModListExt_v1700 := utils.NewSequence[*SchedulingRequestResourceConfigExt_v1700]([]*SchedulingRequestResourceConfigExt_v1700{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_Resources}, false)
				fn_schedulingRequestResourceToAddModListExt_v1700 := func() *SchedulingRequestResourceConfigExt_v1700 {
					return new(SchedulingRequestResourceConfigExt_v1700)
				}
				if err = tmp_schedulingRequestResourceToAddModListExt_v1700.Decode(extReader, fn_schedulingRequestResourceToAddModListExt_v1700); err != nil {
					return utils.WrapError("Decode schedulingRequestResourceToAddModListExt_v1700", err)
				}
				ie.schedulingRequestResourceToAddModListExt_v1700 = []SchedulingRequestResourceConfigExt_v1700{}
				for _, i := range tmp_schedulingRequestResourceToAddModListExt_v1700.Value {
					ie.schedulingRequestResourceToAddModListExt_v1700 = append(ie.schedulingRequestResourceToAddModListExt_v1700, *i)
				}
			}
			// decode dmrs_BundlingPUCCH_Config_r17 optional
			if dmrs_BundlingPUCCH_Config_r17Present {
				tmp_dmrs_BundlingPUCCH_Config_r17 := utils.SetupRelease[*DMRS_BundlingPUCCH_Config_r17]{}
				if err = tmp_dmrs_BundlingPUCCH_Config_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dmrs_BundlingPUCCH_Config_r17", err)
				}
				ie.dmrs_BundlingPUCCH_Config_r17 = tmp_dmrs_BundlingPUCCH_Config_r17.Setup
			}
			// decode dl_DataToUL_ACK_v1700 optional
			if dl_DataToUL_ACK_v1700Present {
				tmp_dl_DataToUL_ACK_v1700 := utils.SetupRelease[*DL_DataToUL_ACK_v1700]{}
				if err = tmp_dl_DataToUL_ACK_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode dl_DataToUL_ACK_v1700", err)
				}
				ie.dl_DataToUL_ACK_v1700 = tmp_dl_DataToUL_ACK_v1700.Setup
			}
			// decode dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 optional
			if dl_DataToUL_ACK_MulticastDCI_Format4_1_r17Present {
				tmp_dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 := utils.SetupRelease[*DL_DataToUL_ACK_MulticastDCI_Format4_1_r17]{}
				if err = tmp_dl_DataToUL_ACK_MulticastDCI_Format4_1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dl_DataToUL_ACK_MulticastDCI_Format4_1_r17", err)
				}
				ie.dl_DataToUL_ACK_MulticastDCI_Format4_1_r17 = tmp_dl_DataToUL_ACK_MulticastDCI_Format4_1_r17.Setup
			}
			// decode sps_PUCCH_AN_ListMulticast_r17 optional
			if sps_PUCCH_AN_ListMulticast_r17Present {
				tmp_sps_PUCCH_AN_ListMulticast_r17 := utils.SetupRelease[*SPS_PUCCH_AN_List_r16]{}
				if err = tmp_sps_PUCCH_AN_ListMulticast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sps_PUCCH_AN_ListMulticast_r17", err)
				}
				ie.sps_PUCCH_AN_ListMulticast_r17 = tmp_sps_PUCCH_AN_ListMulticast_r17.Setup
			}
		}
	}
	return nil
}
