package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_Config struct {
	dataScramblingIdentityPDSCH                             *int64                                                      `lb:0,ub:1023,optional`
	dmrs_DownlinkForPDSCH_MappingTypeA                      *DMRS_DownlinkConfig                                        `optional,setuprelease`
	dmrs_DownlinkForPDSCH_MappingTypeB                      *DMRS_DownlinkConfig                                        `optional,setuprelease`
	tci_StatesToAddModList                                  []TCI_State                                                 `lb:1,ub:maxNrofTCI_States,optional`
	tci_StatesToReleaseList                                 []TCI_StateId                                               `lb:1,ub:maxNrofTCI_States,optional`
	vrb_ToPRB_Interleaver                                   *PDSCH_Config_vrb_ToPRB_Interleaver                         `optional`
	resourceAllocation                                      PDSCH_Config_resourceAllocation                             `madatory`
	pdsch_TimeDomainAllocationList                          *PDSCH_TimeDomainResourceAllocationList                     `optional,setuprelease`
	pdsch_AggregationFactor                                 *PDSCH_Config_pdsch_AggregationFactor                       `optional`
	rateMatchPatternToAddModList                            []RateMatchPattern                                          `lb:1,ub:maxNrofRateMatchPatterns,optional`
	rateMatchPatternToReleaseList                           []RateMatchPatternId                                        `lb:1,ub:maxNrofRateMatchPatterns,optional`
	rateMatchPatternGroup1                                  *RateMatchPatternGroup                                      `optional`
	rateMatchPatternGroup2                                  *RateMatchPatternGroup                                      `optional`
	rbg_Size                                                PDSCH_Config_rbg_Size                                       `madatory`
	mcs_Table                                               *PDSCH_Config_mcs_Table                                     `optional`
	maxNrofCodeWordsScheduledByDCI                          *PDSCH_Config_maxNrofCodeWordsScheduledByDCI                `optional`
	prb_BundlingType                                        *PDSCH_Config_prb_BundlingType                              `optional`
	zp_CSI_RS_ResourceToAddModList                          []ZP_CSI_RS_Resource                                        `lb:1,ub:maxNrofZP_CSI_RS_Resources,optional`
	zp_CSI_RS_ResourceToReleaseList                         []ZP_CSI_RS_ResourceId                                      `lb:1,ub:maxNrofZP_CSI_RS_Resources,optional`
	aperiodic_ZP_CSI_RS_ResourceSetsToAddModList            []ZP_CSI_RS_ResourceSet                                     `lb:1,ub:maxNrofZP_CSI_RS_ResourceSets,optional`
	aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList           []ZP_CSI_RS_ResourceSetId                                   `lb:1,ub:maxNrofZP_CSI_RS_ResourceSets,optional`
	sp_ZP_CSI_RS_ResourceSetsToAddModList                   []ZP_CSI_RS_ResourceSet                                     `lb:1,ub:maxNrofZP_CSI_RS_ResourceSets,optional`
	sp_ZP_CSI_RS_ResourceSetsToReleaseList                  []ZP_CSI_RS_ResourceSetId                                   `lb:1,ub:maxNrofZP_CSI_RS_ResourceSets,optional`
	p_ZP_CSI_RS_ResourceSet                                 *ZP_CSI_RS_ResourceSet                                      `optional,setuprelease`
	maxMIMO_Layers_r16                                      *MaxMIMO_LayersDL_r16                                       `optional,ext-1,setuprelease`
	minimumSchedulingOffsetK0_r16                           *MinSchedulingOffsetK0_Values_r16                           `optional,ext-1,setuprelease`
	antennaPortsFieldPresenceDCI_1_2_r16                    *PDSCH_Config_antennaPortsFieldPresenceDCI_1_2_r16          `optional,ext-1`
	aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16  []ZP_CSI_RS_ResourceSet                                     `lb:1,ub:maxNrofZP_CSI_RS_ResourceSets,optional,ext-1`
	aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 []ZP_CSI_RS_ResourceSetId                                   `lb:1,ub:maxNrofZP_CSI_RS_ResourceSets,optional,ext-1`
	dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16          *DMRS_DownlinkConfig                                        `optional,ext-1,setuprelease`
	dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16          *DMRS_DownlinkConfig                                        `optional,ext-1,setuprelease`
	dmrs_SequenceInitializationDCI_1_2_r16                  *PDSCH_Config_dmrs_SequenceInitializationDCI_1_2_r16        `optional,ext-1`
	harq_ProcessNumberSizeDCI_1_2_r16                       *int64                                                      `lb:0,ub:4,optional,ext-1`
	mcs_TableDCI_1_2_r16                                    *PDSCH_Config_mcs_TableDCI_1_2_r16                          `optional,ext-1`
	numberOfBitsForRV_DCI_1_2_r16                           *int64                                                      `lb:0,ub:2,optional,ext-1`
	pdsch_TimeDomainAllocationListDCI_1_2_r16               *PDSCH_TimeDomainResourceAllocationList_r16                 `optional,ext-1,setuprelease`
	prb_BundlingTypeDCI_1_2_r16                             *PDSCH_Config_prb_BundlingTypeDCI_1_2_r16                   `optional,ext-1`
	priorityIndicatorDCI_1_2_r16                            *PDSCH_Config_priorityIndicatorDCI_1_2_r16                  `optional,ext-1`
	rateMatchPatternGroup1DCI_1_2_r16                       *RateMatchPatternGroup                                      `optional,ext-1`
	rateMatchPatternGroup2DCI_1_2_r16                       *RateMatchPatternGroup                                      `optional,ext-1`
	resourceAllocationType1GranularityDCI_1_2_r16           *PDSCH_Config_resourceAllocationType1GranularityDCI_1_2_r16 `optional,ext-1`
	vrb_ToPRB_InterleaverDCI_1_2_r16                        *PDSCH_Config_vrb_ToPRB_InterleaverDCI_1_2_r16              `optional,ext-1`
	referenceOfSLIVDCI_1_2_r16                              *PDSCH_Config_referenceOfSLIVDCI_1_2_r16                    `optional,ext-1`
	resourceAllocationDCI_1_2_r16                           *PDSCH_Config_resourceAllocationDCI_1_2_r16                 `optional,ext-1`
	priorityIndicatorDCI_1_1_r16                            *PDSCH_Config_priorityIndicatorDCI_1_1_r16                  `optional,ext-1`
	dataScramblingIdentityPDSCH2_r16                        *int64                                                      `lb:0,ub:1023,optional,ext-1`
	pdsch_TimeDomainAllocationList_r16                      *PDSCH_TimeDomainResourceAllocationList_r16                 `optional,ext-1,setuprelease`
	repetitionSchemeConfig_r16                              *RepetitionSchemeConfig_r16                                 `optional,ext-1,setuprelease`
	repetitionSchemeConfig_v1630                            *RepetitionSchemeConfig_v1630                               `optional,ext-2,setuprelease`
	pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17               *PDSCH_Config_pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17     `optional,ext-3`
	pdsch_HARQ_ACK_EnhType3DCI_1_2_r17                      *PDSCH_Config_pdsch_HARQ_ACK_EnhType3DCI_1_2_r17            `optional,ext-3`
	pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17                *PDSCH_Config_pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17      `optional,ext-3`
	pdsch_HARQ_ACK_RetxDCI_1_2_r17                          *PDSCH_Config_pdsch_HARQ_ACK_RetxDCI_1_2_r17                `optional,ext-3`
	pucch_sSCellDynDCI_1_2_r17                              *PDSCH_Config_pucch_sSCellDynDCI_1_2_r17                    `optional,ext-3`
	dl_OrJointTCI_StateList_r17                             *PDSCH_Config_dl_OrJointTCI_StateList_r17                   `lb:1,ub:maxNrofTCI_States,optional,ext-3`
	beamAppTime_r17                                         *PDSCH_Config_beamAppTime_r17                               `optional,ext-3`
	dummy                                                   *Dummy_TDRA_List                                            `optional,ext-3,setuprelease`
	dmrs_FD_OCC_DisabledForRank1_PDSCH_r17                  *PDSCH_Config_dmrs_FD_OCC_DisabledForRank1_PDSCH_r17        `optional,ext-3`
	minimumSchedulingOffsetK0_r17                           *MinSchedulingOffsetK0_Values_r17                           `optional,ext-3,setuprelease`
	harq_ProcessNumberSizeDCI_1_2_v1700                     *int64                                                      `lb:0,ub:5,optional,ext-3`
	harq_ProcessNumberSizeDCI_1_1_r17                       *int64                                                      `lb:5,ub:5,optional,ext-3`
	mcs_Table_r17                                           *PDSCH_Config_mcs_Table_r17                                 `optional,ext-3`
	mcs_TableDCI_1_2_r17                                    *PDSCH_Config_mcs_TableDCI_1_2_r17                          `optional,ext-3`
	xOverheadMulticast_r17                                  *PDSCH_Config_xOverheadMulticast_r17                        `optional,ext-3`
	priorityIndicatorDCI_4_2_r17                            *PDSCH_Config_priorityIndicatorDCI_4_2_r17                  `optional,ext-3`
	sizeDCI_4_2_r17                                         *int64                                                      `lb:20,ub:maxDCI_4_2_Size_r17,optional,ext-3`
	pdsch_TimeDomainAllocationListForMultiPDSCH_r17         *MultiPDSCH_TDRA_List_r17                                   `optional,ext-4,setuprelease`
}

func (ie *PDSCH_Config) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.maxMIMO_Layers_r16 != nil || ie.minimumSchedulingOffsetK0_r16 != nil || ie.antennaPortsFieldPresenceDCI_1_2_r16 != nil || len(ie.aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16) > 0 || len(ie.aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16) > 0 || ie.dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 != nil || ie.dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 != nil || ie.dmrs_SequenceInitializationDCI_1_2_r16 != nil || ie.harq_ProcessNumberSizeDCI_1_2_r16 != nil || ie.mcs_TableDCI_1_2_r16 != nil || ie.numberOfBitsForRV_DCI_1_2_r16 != nil || ie.pdsch_TimeDomainAllocationListDCI_1_2_r16 != nil || ie.prb_BundlingTypeDCI_1_2_r16 != nil || ie.priorityIndicatorDCI_1_2_r16 != nil || ie.rateMatchPatternGroup1DCI_1_2_r16 != nil || ie.rateMatchPatternGroup2DCI_1_2_r16 != nil || ie.resourceAllocationType1GranularityDCI_1_2_r16 != nil || ie.vrb_ToPRB_InterleaverDCI_1_2_r16 != nil || ie.referenceOfSLIVDCI_1_2_r16 != nil || ie.resourceAllocationDCI_1_2_r16 != nil || ie.priorityIndicatorDCI_1_1_r16 != nil || ie.dataScramblingIdentityPDSCH2_r16 != nil || ie.pdsch_TimeDomainAllocationList_r16 != nil || ie.repetitionSchemeConfig_r16 != nil || ie.repetitionSchemeConfig_v1630 != nil || ie.pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17 != nil || ie.pdsch_HARQ_ACK_EnhType3DCI_1_2_r17 != nil || ie.pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17 != nil || ie.pdsch_HARQ_ACK_RetxDCI_1_2_r17 != nil || ie.pucch_sSCellDynDCI_1_2_r17 != nil || ie.dl_OrJointTCI_StateList_r17 != nil || ie.beamAppTime_r17 != nil || ie.dummy != nil || ie.dmrs_FD_OCC_DisabledForRank1_PDSCH_r17 != nil || ie.minimumSchedulingOffsetK0_r17 != nil || ie.harq_ProcessNumberSizeDCI_1_2_v1700 != nil || ie.harq_ProcessNumberSizeDCI_1_1_r17 != nil || ie.mcs_Table_r17 != nil || ie.mcs_TableDCI_1_2_r17 != nil || ie.xOverheadMulticast_r17 != nil || ie.priorityIndicatorDCI_4_2_r17 != nil || ie.sizeDCI_4_2_r17 != nil || ie.pdsch_TimeDomainAllocationListForMultiPDSCH_r17 != nil
	preambleBits := []bool{hasExtensions, ie.dataScramblingIdentityPDSCH != nil, ie.dmrs_DownlinkForPDSCH_MappingTypeA != nil, ie.dmrs_DownlinkForPDSCH_MappingTypeB != nil, len(ie.tci_StatesToAddModList) > 0, len(ie.tci_StatesToReleaseList) > 0, ie.vrb_ToPRB_Interleaver != nil, ie.pdsch_TimeDomainAllocationList != nil, ie.pdsch_AggregationFactor != nil, len(ie.rateMatchPatternToAddModList) > 0, len(ie.rateMatchPatternToReleaseList) > 0, ie.rateMatchPatternGroup1 != nil, ie.rateMatchPatternGroup2 != nil, ie.mcs_Table != nil, ie.maxNrofCodeWordsScheduledByDCI != nil, ie.prb_BundlingType != nil, len(ie.zp_CSI_RS_ResourceToAddModList) > 0, len(ie.zp_CSI_RS_ResourceToReleaseList) > 0, len(ie.aperiodic_ZP_CSI_RS_ResourceSetsToAddModList) > 0, len(ie.aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList) > 0, len(ie.sp_ZP_CSI_RS_ResourceSetsToAddModList) > 0, len(ie.sp_ZP_CSI_RS_ResourceSetsToReleaseList) > 0, ie.p_ZP_CSI_RS_ResourceSet != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dataScramblingIdentityPDSCH != nil {
		if err = w.WriteInteger(*ie.dataScramblingIdentityPDSCH, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode dataScramblingIdentityPDSCH", err)
		}
	}
	if ie.dmrs_DownlinkForPDSCH_MappingTypeA != nil {
		tmp_dmrs_DownlinkForPDSCH_MappingTypeA := utils.SetupRelease[*DMRS_DownlinkConfig]{
			Setup: ie.dmrs_DownlinkForPDSCH_MappingTypeA,
		}
		if err = tmp_dmrs_DownlinkForPDSCH_MappingTypeA.Encode(w); err != nil {
			return utils.WrapError("Encode dmrs_DownlinkForPDSCH_MappingTypeA", err)
		}
	}
	if ie.dmrs_DownlinkForPDSCH_MappingTypeB != nil {
		tmp_dmrs_DownlinkForPDSCH_MappingTypeB := utils.SetupRelease[*DMRS_DownlinkConfig]{
			Setup: ie.dmrs_DownlinkForPDSCH_MappingTypeB,
		}
		if err = tmp_dmrs_DownlinkForPDSCH_MappingTypeB.Encode(w); err != nil {
			return utils.WrapError("Encode dmrs_DownlinkForPDSCH_MappingTypeB", err)
		}
	}
	if len(ie.tci_StatesToAddModList) > 0 {
		tmp_tci_StatesToAddModList := utils.NewSequence[*TCI_State]([]*TCI_State{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false)
		for _, i := range ie.tci_StatesToAddModList {
			tmp_tci_StatesToAddModList.Value = append(tmp_tci_StatesToAddModList.Value, &i)
		}
		if err = tmp_tci_StatesToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode tci_StatesToAddModList", err)
		}
	}
	if len(ie.tci_StatesToReleaseList) > 0 {
		tmp_tci_StatesToReleaseList := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false)
		for _, i := range ie.tci_StatesToReleaseList {
			tmp_tci_StatesToReleaseList.Value = append(tmp_tci_StatesToReleaseList.Value, &i)
		}
		if err = tmp_tci_StatesToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode tci_StatesToReleaseList", err)
		}
	}
	if ie.vrb_ToPRB_Interleaver != nil {
		if err = ie.vrb_ToPRB_Interleaver.Encode(w); err != nil {
			return utils.WrapError("Encode vrb_ToPRB_Interleaver", err)
		}
	}
	if err = ie.resourceAllocation.Encode(w); err != nil {
		return utils.WrapError("Encode resourceAllocation", err)
	}
	if ie.pdsch_TimeDomainAllocationList != nil {
		tmp_pdsch_TimeDomainAllocationList := utils.SetupRelease[*PDSCH_TimeDomainResourceAllocationList]{
			Setup: ie.pdsch_TimeDomainAllocationList,
		}
		if err = tmp_pdsch_TimeDomainAllocationList.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_TimeDomainAllocationList", err)
		}
	}
	if ie.pdsch_AggregationFactor != nil {
		if err = ie.pdsch_AggregationFactor.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_AggregationFactor", err)
		}
	}
	if len(ie.rateMatchPatternToAddModList) > 0 {
		tmp_rateMatchPatternToAddModList := utils.NewSequence[*RateMatchPattern]([]*RateMatchPattern{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		for _, i := range ie.rateMatchPatternToAddModList {
			tmp_rateMatchPatternToAddModList.Value = append(tmp_rateMatchPatternToAddModList.Value, &i)
		}
		if err = tmp_rateMatchPatternToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode rateMatchPatternToAddModList", err)
		}
	}
	if len(ie.rateMatchPatternToReleaseList) > 0 {
		tmp_rateMatchPatternToReleaseList := utils.NewSequence[*RateMatchPatternId]([]*RateMatchPatternId{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		for _, i := range ie.rateMatchPatternToReleaseList {
			tmp_rateMatchPatternToReleaseList.Value = append(tmp_rateMatchPatternToReleaseList.Value, &i)
		}
		if err = tmp_rateMatchPatternToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode rateMatchPatternToReleaseList", err)
		}
	}
	if ie.rateMatchPatternGroup1 != nil {
		if err = ie.rateMatchPatternGroup1.Encode(w); err != nil {
			return utils.WrapError("Encode rateMatchPatternGroup1", err)
		}
	}
	if ie.rateMatchPatternGroup2 != nil {
		if err = ie.rateMatchPatternGroup2.Encode(w); err != nil {
			return utils.WrapError("Encode rateMatchPatternGroup2", err)
		}
	}
	if err = ie.rbg_Size.Encode(w); err != nil {
		return utils.WrapError("Encode rbg_Size", err)
	}
	if ie.mcs_Table != nil {
		if err = ie.mcs_Table.Encode(w); err != nil {
			return utils.WrapError("Encode mcs_Table", err)
		}
	}
	if ie.maxNrofCodeWordsScheduledByDCI != nil {
		if err = ie.maxNrofCodeWordsScheduledByDCI.Encode(w); err != nil {
			return utils.WrapError("Encode maxNrofCodeWordsScheduledByDCI", err)
		}
	}
	if ie.prb_BundlingType != nil {
		if err = ie.prb_BundlingType.Encode(w); err != nil {
			return utils.WrapError("Encode prb_BundlingType", err)
		}
	}
	if len(ie.zp_CSI_RS_ResourceToAddModList) > 0 {
		tmp_zp_CSI_RS_ResourceToAddModList := utils.NewSequence[*ZP_CSI_RS_Resource]([]*ZP_CSI_RS_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_Resources}, false)
		for _, i := range ie.zp_CSI_RS_ResourceToAddModList {
			tmp_zp_CSI_RS_ResourceToAddModList.Value = append(tmp_zp_CSI_RS_ResourceToAddModList.Value, &i)
		}
		if err = tmp_zp_CSI_RS_ResourceToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode zp_CSI_RS_ResourceToAddModList", err)
		}
	}
	if len(ie.zp_CSI_RS_ResourceToReleaseList) > 0 {
		tmp_zp_CSI_RS_ResourceToReleaseList := utils.NewSequence[*ZP_CSI_RS_ResourceId]([]*ZP_CSI_RS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_Resources}, false)
		for _, i := range ie.zp_CSI_RS_ResourceToReleaseList {
			tmp_zp_CSI_RS_ResourceToReleaseList.Value = append(tmp_zp_CSI_RS_ResourceToReleaseList.Value, &i)
		}
		if err = tmp_zp_CSI_RS_ResourceToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode zp_CSI_RS_ResourceToReleaseList", err)
		}
	}
	if len(ie.aperiodic_ZP_CSI_RS_ResourceSetsToAddModList) > 0 {
		tmp_aperiodic_ZP_CSI_RS_ResourceSetsToAddModList := utils.NewSequence[*ZP_CSI_RS_ResourceSet]([]*ZP_CSI_RS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
		for _, i := range ie.aperiodic_ZP_CSI_RS_ResourceSetsToAddModList {
			tmp_aperiodic_ZP_CSI_RS_ResourceSetsToAddModList.Value = append(tmp_aperiodic_ZP_CSI_RS_ResourceSetsToAddModList.Value, &i)
		}
		if err = tmp_aperiodic_ZP_CSI_RS_ResourceSetsToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode aperiodic_ZP_CSI_RS_ResourceSetsToAddModList", err)
		}
	}
	if len(ie.aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList) > 0 {
		tmp_aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList := utils.NewSequence[*ZP_CSI_RS_ResourceSetId]([]*ZP_CSI_RS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
		for _, i := range ie.aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList {
			tmp_aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList.Value = append(tmp_aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList.Value, &i)
		}
		if err = tmp_aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList", err)
		}
	}
	if len(ie.sp_ZP_CSI_RS_ResourceSetsToAddModList) > 0 {
		tmp_sp_ZP_CSI_RS_ResourceSetsToAddModList := utils.NewSequence[*ZP_CSI_RS_ResourceSet]([]*ZP_CSI_RS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
		for _, i := range ie.sp_ZP_CSI_RS_ResourceSetsToAddModList {
			tmp_sp_ZP_CSI_RS_ResourceSetsToAddModList.Value = append(tmp_sp_ZP_CSI_RS_ResourceSetsToAddModList.Value, &i)
		}
		if err = tmp_sp_ZP_CSI_RS_ResourceSetsToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode sp_ZP_CSI_RS_ResourceSetsToAddModList", err)
		}
	}
	if len(ie.sp_ZP_CSI_RS_ResourceSetsToReleaseList) > 0 {
		tmp_sp_ZP_CSI_RS_ResourceSetsToReleaseList := utils.NewSequence[*ZP_CSI_RS_ResourceSetId]([]*ZP_CSI_RS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
		for _, i := range ie.sp_ZP_CSI_RS_ResourceSetsToReleaseList {
			tmp_sp_ZP_CSI_RS_ResourceSetsToReleaseList.Value = append(tmp_sp_ZP_CSI_RS_ResourceSetsToReleaseList.Value, &i)
		}
		if err = tmp_sp_ZP_CSI_RS_ResourceSetsToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode sp_ZP_CSI_RS_ResourceSetsToReleaseList", err)
		}
	}
	if ie.p_ZP_CSI_RS_ResourceSet != nil {
		tmp_p_ZP_CSI_RS_ResourceSet := utils.SetupRelease[*ZP_CSI_RS_ResourceSet]{
			Setup: ie.p_ZP_CSI_RS_ResourceSet,
		}
		if err = tmp_p_ZP_CSI_RS_ResourceSet.Encode(w); err != nil {
			return utils.WrapError("Encode p_ZP_CSI_RS_ResourceSet", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 4 bits for 4 extension groups
		extBitmap := []bool{ie.maxMIMO_Layers_r16 != nil || ie.minimumSchedulingOffsetK0_r16 != nil || ie.antennaPortsFieldPresenceDCI_1_2_r16 != nil || len(ie.aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16) > 0 || len(ie.aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16) > 0 || ie.dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 != nil || ie.dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 != nil || ie.dmrs_SequenceInitializationDCI_1_2_r16 != nil || ie.harq_ProcessNumberSizeDCI_1_2_r16 != nil || ie.mcs_TableDCI_1_2_r16 != nil || ie.numberOfBitsForRV_DCI_1_2_r16 != nil || ie.pdsch_TimeDomainAllocationListDCI_1_2_r16 != nil || ie.prb_BundlingTypeDCI_1_2_r16 != nil || ie.priorityIndicatorDCI_1_2_r16 != nil || ie.rateMatchPatternGroup1DCI_1_2_r16 != nil || ie.rateMatchPatternGroup2DCI_1_2_r16 != nil || ie.resourceAllocationType1GranularityDCI_1_2_r16 != nil || ie.vrb_ToPRB_InterleaverDCI_1_2_r16 != nil || ie.referenceOfSLIVDCI_1_2_r16 != nil || ie.resourceAllocationDCI_1_2_r16 != nil || ie.priorityIndicatorDCI_1_1_r16 != nil || ie.dataScramblingIdentityPDSCH2_r16 != nil || ie.pdsch_TimeDomainAllocationList_r16 != nil || ie.repetitionSchemeConfig_r16 != nil, ie.repetitionSchemeConfig_v1630 != nil, ie.pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17 != nil || ie.pdsch_HARQ_ACK_EnhType3DCI_1_2_r17 != nil || ie.pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17 != nil || ie.pdsch_HARQ_ACK_RetxDCI_1_2_r17 != nil || ie.pucch_sSCellDynDCI_1_2_r17 != nil || ie.dl_OrJointTCI_StateList_r17 != nil || ie.beamAppTime_r17 != nil || ie.dummy != nil || ie.dmrs_FD_OCC_DisabledForRank1_PDSCH_r17 != nil || ie.minimumSchedulingOffsetK0_r17 != nil || ie.harq_ProcessNumberSizeDCI_1_2_v1700 != nil || ie.harq_ProcessNumberSizeDCI_1_1_r17 != nil || ie.mcs_Table_r17 != nil || ie.mcs_TableDCI_1_2_r17 != nil || ie.xOverheadMulticast_r17 != nil || ie.priorityIndicatorDCI_4_2_r17 != nil || ie.sizeDCI_4_2_r17 != nil, ie.pdsch_TimeDomainAllocationListForMultiPDSCH_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PDSCH_Config", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.maxMIMO_Layers_r16 != nil, ie.minimumSchedulingOffsetK0_r16 != nil, ie.antennaPortsFieldPresenceDCI_1_2_r16 != nil, len(ie.aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16) > 0, len(ie.aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16) > 0, ie.dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 != nil, ie.dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 != nil, ie.dmrs_SequenceInitializationDCI_1_2_r16 != nil, ie.harq_ProcessNumberSizeDCI_1_2_r16 != nil, ie.mcs_TableDCI_1_2_r16 != nil, ie.numberOfBitsForRV_DCI_1_2_r16 != nil, ie.pdsch_TimeDomainAllocationListDCI_1_2_r16 != nil, ie.prb_BundlingTypeDCI_1_2_r16 != nil, ie.priorityIndicatorDCI_1_2_r16 != nil, ie.rateMatchPatternGroup1DCI_1_2_r16 != nil, ie.rateMatchPatternGroup2DCI_1_2_r16 != nil, ie.resourceAllocationType1GranularityDCI_1_2_r16 != nil, ie.vrb_ToPRB_InterleaverDCI_1_2_r16 != nil, ie.referenceOfSLIVDCI_1_2_r16 != nil, ie.resourceAllocationDCI_1_2_r16 != nil, ie.priorityIndicatorDCI_1_1_r16 != nil, ie.dataScramblingIdentityPDSCH2_r16 != nil, ie.pdsch_TimeDomainAllocationList_r16 != nil, ie.repetitionSchemeConfig_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode maxMIMO_Layers_r16 optional
			if ie.maxMIMO_Layers_r16 != nil {
				tmp_maxMIMO_Layers_r16 := utils.SetupRelease[*MaxMIMO_LayersDL_r16]{
					Setup: ie.maxMIMO_Layers_r16,
				}
				if err = tmp_maxMIMO_Layers_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxMIMO_Layers_r16", err)
				}
			}
			// encode minimumSchedulingOffsetK0_r16 optional
			if ie.minimumSchedulingOffsetK0_r16 != nil {
				tmp_minimumSchedulingOffsetK0_r16 := utils.SetupRelease[*MinSchedulingOffsetK0_Values_r16]{
					Setup: ie.minimumSchedulingOffsetK0_r16,
				}
				if err = tmp_minimumSchedulingOffsetK0_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode minimumSchedulingOffsetK0_r16", err)
				}
			}
			// encode antennaPortsFieldPresenceDCI_1_2_r16 optional
			if ie.antennaPortsFieldPresenceDCI_1_2_r16 != nil {
				if err = ie.antennaPortsFieldPresenceDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode antennaPortsFieldPresenceDCI_1_2_r16", err)
				}
			}
			// encode aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 optional
			if len(ie.aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16) > 0 {
				tmp_aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 := utils.NewSequence[*ZP_CSI_RS_ResourceSet]([]*ZP_CSI_RS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
				for _, i := range ie.aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 {
					tmp_aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16.Value = append(tmp_aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16.Value, &i)
				}
				if err = tmp_aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16", err)
				}
			}
			// encode aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 optional
			if len(ie.aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16) > 0 {
				tmp_aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 := utils.NewSequence[*ZP_CSI_RS_ResourceSetId]([]*ZP_CSI_RS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
				for _, i := range ie.aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 {
					tmp_aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16.Value = append(tmp_aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16.Value, &i)
				}
				if err = tmp_aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16", err)
				}
			}
			// encode dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 optional
			if ie.dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 != nil {
				tmp_dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 := utils.SetupRelease[*DMRS_DownlinkConfig]{
					Setup: ie.dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16,
				}
				if err = tmp_dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16", err)
				}
			}
			// encode dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 optional
			if ie.dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 != nil {
				tmp_dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 := utils.SetupRelease[*DMRS_DownlinkConfig]{
					Setup: ie.dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16,
				}
				if err = tmp_dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16", err)
				}
			}
			// encode dmrs_SequenceInitializationDCI_1_2_r16 optional
			if ie.dmrs_SequenceInitializationDCI_1_2_r16 != nil {
				if err = ie.dmrs_SequenceInitializationDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dmrs_SequenceInitializationDCI_1_2_r16", err)
				}
			}
			// encode harq_ProcessNumberSizeDCI_1_2_r16 optional
			if ie.harq_ProcessNumberSizeDCI_1_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.harq_ProcessNumberSizeDCI_1_2_r16, &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
					return utils.WrapError("Encode harq_ProcessNumberSizeDCI_1_2_r16", err)
				}
			}
			// encode mcs_TableDCI_1_2_r16 optional
			if ie.mcs_TableDCI_1_2_r16 != nil {
				if err = ie.mcs_TableDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mcs_TableDCI_1_2_r16", err)
				}
			}
			// encode numberOfBitsForRV_DCI_1_2_r16 optional
			if ie.numberOfBitsForRV_DCI_1_2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.numberOfBitsForRV_DCI_1_2_r16, &uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
					return utils.WrapError("Encode numberOfBitsForRV_DCI_1_2_r16", err)
				}
			}
			// encode pdsch_TimeDomainAllocationListDCI_1_2_r16 optional
			if ie.pdsch_TimeDomainAllocationListDCI_1_2_r16 != nil {
				tmp_pdsch_TimeDomainAllocationListDCI_1_2_r16 := utils.SetupRelease[*PDSCH_TimeDomainResourceAllocationList_r16]{
					Setup: ie.pdsch_TimeDomainAllocationListDCI_1_2_r16,
				}
				if err = tmp_pdsch_TimeDomainAllocationListDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_TimeDomainAllocationListDCI_1_2_r16", err)
				}
			}
			// encode prb_BundlingTypeDCI_1_2_r16 optional
			if ie.prb_BundlingTypeDCI_1_2_r16 != nil {
				if err = ie.prb_BundlingTypeDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode prb_BundlingTypeDCI_1_2_r16", err)
				}
			}
			// encode priorityIndicatorDCI_1_2_r16 optional
			if ie.priorityIndicatorDCI_1_2_r16 != nil {
				if err = ie.priorityIndicatorDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode priorityIndicatorDCI_1_2_r16", err)
				}
			}
			// encode rateMatchPatternGroup1DCI_1_2_r16 optional
			if ie.rateMatchPatternGroup1DCI_1_2_r16 != nil {
				if err = ie.rateMatchPatternGroup1DCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rateMatchPatternGroup1DCI_1_2_r16", err)
				}
			}
			// encode rateMatchPatternGroup2DCI_1_2_r16 optional
			if ie.rateMatchPatternGroup2DCI_1_2_r16 != nil {
				if err = ie.rateMatchPatternGroup2DCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rateMatchPatternGroup2DCI_1_2_r16", err)
				}
			}
			// encode resourceAllocationType1GranularityDCI_1_2_r16 optional
			if ie.resourceAllocationType1GranularityDCI_1_2_r16 != nil {
				if err = ie.resourceAllocationType1GranularityDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode resourceAllocationType1GranularityDCI_1_2_r16", err)
				}
			}
			// encode vrb_ToPRB_InterleaverDCI_1_2_r16 optional
			if ie.vrb_ToPRB_InterleaverDCI_1_2_r16 != nil {
				if err = ie.vrb_ToPRB_InterleaverDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode vrb_ToPRB_InterleaverDCI_1_2_r16", err)
				}
			}
			// encode referenceOfSLIVDCI_1_2_r16 optional
			if ie.referenceOfSLIVDCI_1_2_r16 != nil {
				if err = ie.referenceOfSLIVDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode referenceOfSLIVDCI_1_2_r16", err)
				}
			}
			// encode resourceAllocationDCI_1_2_r16 optional
			if ie.resourceAllocationDCI_1_2_r16 != nil {
				if err = ie.resourceAllocationDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode resourceAllocationDCI_1_2_r16", err)
				}
			}
			// encode priorityIndicatorDCI_1_1_r16 optional
			if ie.priorityIndicatorDCI_1_1_r16 != nil {
				if err = ie.priorityIndicatorDCI_1_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode priorityIndicatorDCI_1_1_r16", err)
				}
			}
			// encode dataScramblingIdentityPDSCH2_r16 optional
			if ie.dataScramblingIdentityPDSCH2_r16 != nil {
				if err = extWriter.WriteInteger(*ie.dataScramblingIdentityPDSCH2_r16, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
					return utils.WrapError("Encode dataScramblingIdentityPDSCH2_r16", err)
				}
			}
			// encode pdsch_TimeDomainAllocationList_r16 optional
			if ie.pdsch_TimeDomainAllocationList_r16 != nil {
				tmp_pdsch_TimeDomainAllocationList_r16 := utils.SetupRelease[*PDSCH_TimeDomainResourceAllocationList_r16]{
					Setup: ie.pdsch_TimeDomainAllocationList_r16,
				}
				if err = tmp_pdsch_TimeDomainAllocationList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_TimeDomainAllocationList_r16", err)
				}
			}
			// encode repetitionSchemeConfig_r16 optional
			if ie.repetitionSchemeConfig_r16 != nil {
				tmp_repetitionSchemeConfig_r16 := utils.SetupRelease[*RepetitionSchemeConfig_r16]{
					Setup: ie.repetitionSchemeConfig_r16,
				}
				if err = tmp_repetitionSchemeConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode repetitionSchemeConfig_r16", err)
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
			optionals_ext_2 := []bool{ie.repetitionSchemeConfig_v1630 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode repetitionSchemeConfig_v1630 optional
			if ie.repetitionSchemeConfig_v1630 != nil {
				tmp_repetitionSchemeConfig_v1630 := utils.SetupRelease[*RepetitionSchemeConfig_v1630]{
					Setup: ie.repetitionSchemeConfig_v1630,
				}
				if err = tmp_repetitionSchemeConfig_v1630.Encode(extWriter); err != nil {
					return utils.WrapError("Encode repetitionSchemeConfig_v1630", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17 != nil, ie.pdsch_HARQ_ACK_EnhType3DCI_1_2_r17 != nil, ie.pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17 != nil, ie.pdsch_HARQ_ACK_RetxDCI_1_2_r17 != nil, ie.pucch_sSCellDynDCI_1_2_r17 != nil, ie.dl_OrJointTCI_StateList_r17 != nil, ie.beamAppTime_r17 != nil, ie.dummy != nil, ie.dmrs_FD_OCC_DisabledForRank1_PDSCH_r17 != nil, ie.minimumSchedulingOffsetK0_r17 != nil, ie.harq_ProcessNumberSizeDCI_1_2_v1700 != nil, ie.harq_ProcessNumberSizeDCI_1_1_r17 != nil, ie.mcs_Table_r17 != nil, ie.mcs_TableDCI_1_2_r17 != nil, ie.xOverheadMulticast_r17 != nil, ie.priorityIndicatorDCI_4_2_r17 != nil, ie.sizeDCI_4_2_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17 optional
			if ie.pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17 != nil {
				if err = ie.pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17", err)
				}
			}
			// encode pdsch_HARQ_ACK_EnhType3DCI_1_2_r17 optional
			if ie.pdsch_HARQ_ACK_EnhType3DCI_1_2_r17 != nil {
				if err = ie.pdsch_HARQ_ACK_EnhType3DCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_HARQ_ACK_EnhType3DCI_1_2_r17", err)
				}
			}
			// encode pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17 optional
			if ie.pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17 != nil {
				if err = ie.pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17", err)
				}
			}
			// encode pdsch_HARQ_ACK_RetxDCI_1_2_r17 optional
			if ie.pdsch_HARQ_ACK_RetxDCI_1_2_r17 != nil {
				if err = ie.pdsch_HARQ_ACK_RetxDCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_HARQ_ACK_RetxDCI_1_2_r17", err)
				}
			}
			// encode pucch_sSCellDynDCI_1_2_r17 optional
			if ie.pucch_sSCellDynDCI_1_2_r17 != nil {
				if err = ie.pucch_sSCellDynDCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pucch_sSCellDynDCI_1_2_r17", err)
				}
			}
			// encode dl_OrJointTCI_StateList_r17 optional
			if ie.dl_OrJointTCI_StateList_r17 != nil {
				if err = ie.dl_OrJointTCI_StateList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dl_OrJointTCI_StateList_r17", err)
				}
			}
			// encode beamAppTime_r17 optional
			if ie.beamAppTime_r17 != nil {
				if err = ie.beamAppTime_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode beamAppTime_r17", err)
				}
			}
			// encode dummy optional
			if ie.dummy != nil {
				tmp_dummy := utils.SetupRelease[*Dummy_TDRA_List]{
					Setup: ie.dummy,
				}
				if err = tmp_dummy.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dummy", err)
				}
			}
			// encode dmrs_FD_OCC_DisabledForRank1_PDSCH_r17 optional
			if ie.dmrs_FD_OCC_DisabledForRank1_PDSCH_r17 != nil {
				if err = ie.dmrs_FD_OCC_DisabledForRank1_PDSCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dmrs_FD_OCC_DisabledForRank1_PDSCH_r17", err)
				}
			}
			// encode minimumSchedulingOffsetK0_r17 optional
			if ie.minimumSchedulingOffsetK0_r17 != nil {
				tmp_minimumSchedulingOffsetK0_r17 := utils.SetupRelease[*MinSchedulingOffsetK0_Values_r17]{
					Setup: ie.minimumSchedulingOffsetK0_r17,
				}
				if err = tmp_minimumSchedulingOffsetK0_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode minimumSchedulingOffsetK0_r17", err)
				}
			}
			// encode harq_ProcessNumberSizeDCI_1_2_v1700 optional
			if ie.harq_ProcessNumberSizeDCI_1_2_v1700 != nil {
				if err = extWriter.WriteInteger(*ie.harq_ProcessNumberSizeDCI_1_2_v1700, &uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
					return utils.WrapError("Encode harq_ProcessNumberSizeDCI_1_2_v1700", err)
				}
			}
			// encode harq_ProcessNumberSizeDCI_1_1_r17 optional
			if ie.harq_ProcessNumberSizeDCI_1_1_r17 != nil {
				if err = extWriter.WriteInteger(*ie.harq_ProcessNumberSizeDCI_1_1_r17, &uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
					return utils.WrapError("Encode harq_ProcessNumberSizeDCI_1_1_r17", err)
				}
			}
			// encode mcs_Table_r17 optional
			if ie.mcs_Table_r17 != nil {
				if err = ie.mcs_Table_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mcs_Table_r17", err)
				}
			}
			// encode mcs_TableDCI_1_2_r17 optional
			if ie.mcs_TableDCI_1_2_r17 != nil {
				if err = ie.mcs_TableDCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mcs_TableDCI_1_2_r17", err)
				}
			}
			// encode xOverheadMulticast_r17 optional
			if ie.xOverheadMulticast_r17 != nil {
				if err = ie.xOverheadMulticast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode xOverheadMulticast_r17", err)
				}
			}
			// encode priorityIndicatorDCI_4_2_r17 optional
			if ie.priorityIndicatorDCI_4_2_r17 != nil {
				if err = ie.priorityIndicatorDCI_4_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode priorityIndicatorDCI_4_2_r17", err)
				}
			}
			// encode sizeDCI_4_2_r17 optional
			if ie.sizeDCI_4_2_r17 != nil {
				if err = extWriter.WriteInteger(*ie.sizeDCI_4_2_r17, &uper.Constraint{Lb: 20, Ub: maxDCI_4_2_Size_r17}, false); err != nil {
					return utils.WrapError("Encode sizeDCI_4_2_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 4
		if extBitmap[3] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{ie.pdsch_TimeDomainAllocationListForMultiPDSCH_r17 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode pdsch_TimeDomainAllocationListForMultiPDSCH_r17 optional
			if ie.pdsch_TimeDomainAllocationListForMultiPDSCH_r17 != nil {
				tmp_pdsch_TimeDomainAllocationListForMultiPDSCH_r17 := utils.SetupRelease[*MultiPDSCH_TDRA_List_r17]{
					Setup: ie.pdsch_TimeDomainAllocationListForMultiPDSCH_r17,
				}
				if err = tmp_pdsch_TimeDomainAllocationListForMultiPDSCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_TimeDomainAllocationListForMultiPDSCH_r17", err)
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

func (ie *PDSCH_Config) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var dataScramblingIdentityPDSCHPresent bool
	if dataScramblingIdentityPDSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dmrs_DownlinkForPDSCH_MappingTypeAPresent bool
	if dmrs_DownlinkForPDSCH_MappingTypeAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dmrs_DownlinkForPDSCH_MappingTypeBPresent bool
	if dmrs_DownlinkForPDSCH_MappingTypeBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tci_StatesToAddModListPresent bool
	if tci_StatesToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tci_StatesToReleaseListPresent bool
	if tci_StatesToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var vrb_ToPRB_InterleaverPresent bool
	if vrb_ToPRB_InterleaverPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_TimeDomainAllocationListPresent bool
	if pdsch_TimeDomainAllocationListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_AggregationFactorPresent bool
	if pdsch_AggregationFactorPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rateMatchPatternToAddModListPresent bool
	if rateMatchPatternToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rateMatchPatternToReleaseListPresent bool
	if rateMatchPatternToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rateMatchPatternGroup1Present bool
	if rateMatchPatternGroup1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rateMatchPatternGroup2Present bool
	if rateMatchPatternGroup2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mcs_TablePresent bool
	if mcs_TablePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNrofCodeWordsScheduledByDCIPresent bool
	if maxNrofCodeWordsScheduledByDCIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var prb_BundlingTypePresent bool
	if prb_BundlingTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var zp_CSI_RS_ResourceToAddModListPresent bool
	if zp_CSI_RS_ResourceToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var zp_CSI_RS_ResourceToReleaseListPresent bool
	if zp_CSI_RS_ResourceToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var aperiodic_ZP_CSI_RS_ResourceSetsToAddModListPresent bool
	if aperiodic_ZP_CSI_RS_ResourceSetsToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var aperiodic_ZP_CSI_RS_ResourceSetsToReleaseListPresent bool
	if aperiodic_ZP_CSI_RS_ResourceSetsToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sp_ZP_CSI_RS_ResourceSetsToAddModListPresent bool
	if sp_ZP_CSI_RS_ResourceSetsToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sp_ZP_CSI_RS_ResourceSetsToReleaseListPresent bool
	if sp_ZP_CSI_RS_ResourceSetsToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var p_ZP_CSI_RS_ResourceSetPresent bool
	if p_ZP_CSI_RS_ResourceSetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if dataScramblingIdentityPDSCHPresent {
		var tmp_int_dataScramblingIdentityPDSCH int64
		if tmp_int_dataScramblingIdentityPDSCH, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode dataScramblingIdentityPDSCH", err)
		}
		ie.dataScramblingIdentityPDSCH = &tmp_int_dataScramblingIdentityPDSCH
	}
	if dmrs_DownlinkForPDSCH_MappingTypeAPresent {
		tmp_dmrs_DownlinkForPDSCH_MappingTypeA := utils.SetupRelease[*DMRS_DownlinkConfig]{}
		if err = tmp_dmrs_DownlinkForPDSCH_MappingTypeA.Decode(r); err != nil {
			return utils.WrapError("Decode dmrs_DownlinkForPDSCH_MappingTypeA", err)
		}
		ie.dmrs_DownlinkForPDSCH_MappingTypeA = tmp_dmrs_DownlinkForPDSCH_MappingTypeA.Setup
	}
	if dmrs_DownlinkForPDSCH_MappingTypeBPresent {
		tmp_dmrs_DownlinkForPDSCH_MappingTypeB := utils.SetupRelease[*DMRS_DownlinkConfig]{}
		if err = tmp_dmrs_DownlinkForPDSCH_MappingTypeB.Decode(r); err != nil {
			return utils.WrapError("Decode dmrs_DownlinkForPDSCH_MappingTypeB", err)
		}
		ie.dmrs_DownlinkForPDSCH_MappingTypeB = tmp_dmrs_DownlinkForPDSCH_MappingTypeB.Setup
	}
	if tci_StatesToAddModListPresent {
		tmp_tci_StatesToAddModList := utils.NewSequence[*TCI_State]([]*TCI_State{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false)
		fn_tci_StatesToAddModList := func() *TCI_State {
			return new(TCI_State)
		}
		if err = tmp_tci_StatesToAddModList.Decode(r, fn_tci_StatesToAddModList); err != nil {
			return utils.WrapError("Decode tci_StatesToAddModList", err)
		}
		ie.tci_StatesToAddModList = []TCI_State{}
		for _, i := range tmp_tci_StatesToAddModList.Value {
			ie.tci_StatesToAddModList = append(ie.tci_StatesToAddModList, *i)
		}
	}
	if tci_StatesToReleaseListPresent {
		tmp_tci_StatesToReleaseList := utils.NewSequence[*TCI_StateId]([]*TCI_StateId{}, uper.Constraint{Lb: 1, Ub: maxNrofTCI_States}, false)
		fn_tci_StatesToReleaseList := func() *TCI_StateId {
			return new(TCI_StateId)
		}
		if err = tmp_tci_StatesToReleaseList.Decode(r, fn_tci_StatesToReleaseList); err != nil {
			return utils.WrapError("Decode tci_StatesToReleaseList", err)
		}
		ie.tci_StatesToReleaseList = []TCI_StateId{}
		for _, i := range tmp_tci_StatesToReleaseList.Value {
			ie.tci_StatesToReleaseList = append(ie.tci_StatesToReleaseList, *i)
		}
	}
	if vrb_ToPRB_InterleaverPresent {
		ie.vrb_ToPRB_Interleaver = new(PDSCH_Config_vrb_ToPRB_Interleaver)
		if err = ie.vrb_ToPRB_Interleaver.Decode(r); err != nil {
			return utils.WrapError("Decode vrb_ToPRB_Interleaver", err)
		}
	}
	if err = ie.resourceAllocation.Decode(r); err != nil {
		return utils.WrapError("Decode resourceAllocation", err)
	}
	if pdsch_TimeDomainAllocationListPresent {
		tmp_pdsch_TimeDomainAllocationList := utils.SetupRelease[*PDSCH_TimeDomainResourceAllocationList]{}
		if err = tmp_pdsch_TimeDomainAllocationList.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_TimeDomainAllocationList", err)
		}
		ie.pdsch_TimeDomainAllocationList = tmp_pdsch_TimeDomainAllocationList.Setup
	}
	if pdsch_AggregationFactorPresent {
		ie.pdsch_AggregationFactor = new(PDSCH_Config_pdsch_AggregationFactor)
		if err = ie.pdsch_AggregationFactor.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_AggregationFactor", err)
		}
	}
	if rateMatchPatternToAddModListPresent {
		tmp_rateMatchPatternToAddModList := utils.NewSequence[*RateMatchPattern]([]*RateMatchPattern{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		fn_rateMatchPatternToAddModList := func() *RateMatchPattern {
			return new(RateMatchPattern)
		}
		if err = tmp_rateMatchPatternToAddModList.Decode(r, fn_rateMatchPatternToAddModList); err != nil {
			return utils.WrapError("Decode rateMatchPatternToAddModList", err)
		}
		ie.rateMatchPatternToAddModList = []RateMatchPattern{}
		for _, i := range tmp_rateMatchPatternToAddModList.Value {
			ie.rateMatchPatternToAddModList = append(ie.rateMatchPatternToAddModList, *i)
		}
	}
	if rateMatchPatternToReleaseListPresent {
		tmp_rateMatchPatternToReleaseList := utils.NewSequence[*RateMatchPatternId]([]*RateMatchPatternId{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatterns}, false)
		fn_rateMatchPatternToReleaseList := func() *RateMatchPatternId {
			return new(RateMatchPatternId)
		}
		if err = tmp_rateMatchPatternToReleaseList.Decode(r, fn_rateMatchPatternToReleaseList); err != nil {
			return utils.WrapError("Decode rateMatchPatternToReleaseList", err)
		}
		ie.rateMatchPatternToReleaseList = []RateMatchPatternId{}
		for _, i := range tmp_rateMatchPatternToReleaseList.Value {
			ie.rateMatchPatternToReleaseList = append(ie.rateMatchPatternToReleaseList, *i)
		}
	}
	if rateMatchPatternGroup1Present {
		ie.rateMatchPatternGroup1 = new(RateMatchPatternGroup)
		if err = ie.rateMatchPatternGroup1.Decode(r); err != nil {
			return utils.WrapError("Decode rateMatchPatternGroup1", err)
		}
	}
	if rateMatchPatternGroup2Present {
		ie.rateMatchPatternGroup2 = new(RateMatchPatternGroup)
		if err = ie.rateMatchPatternGroup2.Decode(r); err != nil {
			return utils.WrapError("Decode rateMatchPatternGroup2", err)
		}
	}
	if err = ie.rbg_Size.Decode(r); err != nil {
		return utils.WrapError("Decode rbg_Size", err)
	}
	if mcs_TablePresent {
		ie.mcs_Table = new(PDSCH_Config_mcs_Table)
		if err = ie.mcs_Table.Decode(r); err != nil {
			return utils.WrapError("Decode mcs_Table", err)
		}
	}
	if maxNrofCodeWordsScheduledByDCIPresent {
		ie.maxNrofCodeWordsScheduledByDCI = new(PDSCH_Config_maxNrofCodeWordsScheduledByDCI)
		if err = ie.maxNrofCodeWordsScheduledByDCI.Decode(r); err != nil {
			return utils.WrapError("Decode maxNrofCodeWordsScheduledByDCI", err)
		}
	}
	if prb_BundlingTypePresent {
		ie.prb_BundlingType = new(PDSCH_Config_prb_BundlingType)
		if err = ie.prb_BundlingType.Decode(r); err != nil {
			return utils.WrapError("Decode prb_BundlingType", err)
		}
	}
	if zp_CSI_RS_ResourceToAddModListPresent {
		tmp_zp_CSI_RS_ResourceToAddModList := utils.NewSequence[*ZP_CSI_RS_Resource]([]*ZP_CSI_RS_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_Resources}, false)
		fn_zp_CSI_RS_ResourceToAddModList := func() *ZP_CSI_RS_Resource {
			return new(ZP_CSI_RS_Resource)
		}
		if err = tmp_zp_CSI_RS_ResourceToAddModList.Decode(r, fn_zp_CSI_RS_ResourceToAddModList); err != nil {
			return utils.WrapError("Decode zp_CSI_RS_ResourceToAddModList", err)
		}
		ie.zp_CSI_RS_ResourceToAddModList = []ZP_CSI_RS_Resource{}
		for _, i := range tmp_zp_CSI_RS_ResourceToAddModList.Value {
			ie.zp_CSI_RS_ResourceToAddModList = append(ie.zp_CSI_RS_ResourceToAddModList, *i)
		}
	}
	if zp_CSI_RS_ResourceToReleaseListPresent {
		tmp_zp_CSI_RS_ResourceToReleaseList := utils.NewSequence[*ZP_CSI_RS_ResourceId]([]*ZP_CSI_RS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_Resources}, false)
		fn_zp_CSI_RS_ResourceToReleaseList := func() *ZP_CSI_RS_ResourceId {
			return new(ZP_CSI_RS_ResourceId)
		}
		if err = tmp_zp_CSI_RS_ResourceToReleaseList.Decode(r, fn_zp_CSI_RS_ResourceToReleaseList); err != nil {
			return utils.WrapError("Decode zp_CSI_RS_ResourceToReleaseList", err)
		}
		ie.zp_CSI_RS_ResourceToReleaseList = []ZP_CSI_RS_ResourceId{}
		for _, i := range tmp_zp_CSI_RS_ResourceToReleaseList.Value {
			ie.zp_CSI_RS_ResourceToReleaseList = append(ie.zp_CSI_RS_ResourceToReleaseList, *i)
		}
	}
	if aperiodic_ZP_CSI_RS_ResourceSetsToAddModListPresent {
		tmp_aperiodic_ZP_CSI_RS_ResourceSetsToAddModList := utils.NewSequence[*ZP_CSI_RS_ResourceSet]([]*ZP_CSI_RS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
		fn_aperiodic_ZP_CSI_RS_ResourceSetsToAddModList := func() *ZP_CSI_RS_ResourceSet {
			return new(ZP_CSI_RS_ResourceSet)
		}
		if err = tmp_aperiodic_ZP_CSI_RS_ResourceSetsToAddModList.Decode(r, fn_aperiodic_ZP_CSI_RS_ResourceSetsToAddModList); err != nil {
			return utils.WrapError("Decode aperiodic_ZP_CSI_RS_ResourceSetsToAddModList", err)
		}
		ie.aperiodic_ZP_CSI_RS_ResourceSetsToAddModList = []ZP_CSI_RS_ResourceSet{}
		for _, i := range tmp_aperiodic_ZP_CSI_RS_ResourceSetsToAddModList.Value {
			ie.aperiodic_ZP_CSI_RS_ResourceSetsToAddModList = append(ie.aperiodic_ZP_CSI_RS_ResourceSetsToAddModList, *i)
		}
	}
	if aperiodic_ZP_CSI_RS_ResourceSetsToReleaseListPresent {
		tmp_aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList := utils.NewSequence[*ZP_CSI_RS_ResourceSetId]([]*ZP_CSI_RS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
		fn_aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList := func() *ZP_CSI_RS_ResourceSetId {
			return new(ZP_CSI_RS_ResourceSetId)
		}
		if err = tmp_aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList.Decode(r, fn_aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList); err != nil {
			return utils.WrapError("Decode aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList", err)
		}
		ie.aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList = []ZP_CSI_RS_ResourceSetId{}
		for _, i := range tmp_aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList.Value {
			ie.aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList = append(ie.aperiodic_ZP_CSI_RS_ResourceSetsToReleaseList, *i)
		}
	}
	if sp_ZP_CSI_RS_ResourceSetsToAddModListPresent {
		tmp_sp_ZP_CSI_RS_ResourceSetsToAddModList := utils.NewSequence[*ZP_CSI_RS_ResourceSet]([]*ZP_CSI_RS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
		fn_sp_ZP_CSI_RS_ResourceSetsToAddModList := func() *ZP_CSI_RS_ResourceSet {
			return new(ZP_CSI_RS_ResourceSet)
		}
		if err = tmp_sp_ZP_CSI_RS_ResourceSetsToAddModList.Decode(r, fn_sp_ZP_CSI_RS_ResourceSetsToAddModList); err != nil {
			return utils.WrapError("Decode sp_ZP_CSI_RS_ResourceSetsToAddModList", err)
		}
		ie.sp_ZP_CSI_RS_ResourceSetsToAddModList = []ZP_CSI_RS_ResourceSet{}
		for _, i := range tmp_sp_ZP_CSI_RS_ResourceSetsToAddModList.Value {
			ie.sp_ZP_CSI_RS_ResourceSetsToAddModList = append(ie.sp_ZP_CSI_RS_ResourceSetsToAddModList, *i)
		}
	}
	if sp_ZP_CSI_RS_ResourceSetsToReleaseListPresent {
		tmp_sp_ZP_CSI_RS_ResourceSetsToReleaseList := utils.NewSequence[*ZP_CSI_RS_ResourceSetId]([]*ZP_CSI_RS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
		fn_sp_ZP_CSI_RS_ResourceSetsToReleaseList := func() *ZP_CSI_RS_ResourceSetId {
			return new(ZP_CSI_RS_ResourceSetId)
		}
		if err = tmp_sp_ZP_CSI_RS_ResourceSetsToReleaseList.Decode(r, fn_sp_ZP_CSI_RS_ResourceSetsToReleaseList); err != nil {
			return utils.WrapError("Decode sp_ZP_CSI_RS_ResourceSetsToReleaseList", err)
		}
		ie.sp_ZP_CSI_RS_ResourceSetsToReleaseList = []ZP_CSI_RS_ResourceSetId{}
		for _, i := range tmp_sp_ZP_CSI_RS_ResourceSetsToReleaseList.Value {
			ie.sp_ZP_CSI_RS_ResourceSetsToReleaseList = append(ie.sp_ZP_CSI_RS_ResourceSetsToReleaseList, *i)
		}
	}
	if p_ZP_CSI_RS_ResourceSetPresent {
		tmp_p_ZP_CSI_RS_ResourceSet := utils.SetupRelease[*ZP_CSI_RS_ResourceSet]{}
		if err = tmp_p_ZP_CSI_RS_ResourceSet.Decode(r); err != nil {
			return utils.WrapError("Decode p_ZP_CSI_RS_ResourceSet", err)
		}
		ie.p_ZP_CSI_RS_ResourceSet = tmp_p_ZP_CSI_RS_ResourceSet.Setup
	}

	if extensionBit {
		// Read extension bitmap: 4 bits for 4 extension groups
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

			maxMIMO_Layers_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			minimumSchedulingOffsetK0_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			antennaPortsFieldPresenceDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dmrs_SequenceInitializationDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			harq_ProcessNumberSizeDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mcs_TableDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			numberOfBitsForRV_DCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_TimeDomainAllocationListDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			prb_BundlingTypeDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			priorityIndicatorDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			rateMatchPatternGroup1DCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			rateMatchPatternGroup2DCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			resourceAllocationType1GranularityDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			vrb_ToPRB_InterleaverDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			referenceOfSLIVDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			resourceAllocationDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			priorityIndicatorDCI_1_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dataScramblingIdentityPDSCH2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_TimeDomainAllocationList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			repetitionSchemeConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode maxMIMO_Layers_r16 optional
			if maxMIMO_Layers_r16Present {
				tmp_maxMIMO_Layers_r16 := utils.SetupRelease[*MaxMIMO_LayersDL_r16]{}
				if err = tmp_maxMIMO_Layers_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxMIMO_Layers_r16", err)
				}
				ie.maxMIMO_Layers_r16 = tmp_maxMIMO_Layers_r16.Setup
			}
			// decode minimumSchedulingOffsetK0_r16 optional
			if minimumSchedulingOffsetK0_r16Present {
				tmp_minimumSchedulingOffsetK0_r16 := utils.SetupRelease[*MinSchedulingOffsetK0_Values_r16]{}
				if err = tmp_minimumSchedulingOffsetK0_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode minimumSchedulingOffsetK0_r16", err)
				}
				ie.minimumSchedulingOffsetK0_r16 = tmp_minimumSchedulingOffsetK0_r16.Setup
			}
			// decode antennaPortsFieldPresenceDCI_1_2_r16 optional
			if antennaPortsFieldPresenceDCI_1_2_r16Present {
				ie.antennaPortsFieldPresenceDCI_1_2_r16 = new(PDSCH_Config_antennaPortsFieldPresenceDCI_1_2_r16)
				if err = ie.antennaPortsFieldPresenceDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode antennaPortsFieldPresenceDCI_1_2_r16", err)
				}
			}
			// decode aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 optional
			if aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16Present {
				tmp_aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 := utils.NewSequence[*ZP_CSI_RS_ResourceSet]([]*ZP_CSI_RS_ResourceSet{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
				fn_aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 := func() *ZP_CSI_RS_ResourceSet {
					return new(ZP_CSI_RS_ResourceSet)
				}
				if err = tmp_aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16.Decode(extReader, fn_aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16); err != nil {
					return utils.WrapError("Decode aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16", err)
				}
				ie.aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 = []ZP_CSI_RS_ResourceSet{}
				for _, i := range tmp_aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16.Value {
					ie.aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16 = append(ie.aperiodicZP_CSI_RS_ResourceSetsToAddModListDCI_1_2_r16, *i)
				}
			}
			// decode aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 optional
			if aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16Present {
				tmp_aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 := utils.NewSequence[*ZP_CSI_RS_ResourceSetId]([]*ZP_CSI_RS_ResourceSetId{}, uper.Constraint{Lb: 1, Ub: maxNrofZP_CSI_RS_ResourceSets}, false)
				fn_aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 := func() *ZP_CSI_RS_ResourceSetId {
					return new(ZP_CSI_RS_ResourceSetId)
				}
				if err = tmp_aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16.Decode(extReader, fn_aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16); err != nil {
					return utils.WrapError("Decode aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16", err)
				}
				ie.aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 = []ZP_CSI_RS_ResourceSetId{}
				for _, i := range tmp_aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16.Value {
					ie.aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16 = append(ie.aperiodicZP_CSI_RS_ResourceSetsToReleaseListDCI_1_2_r16, *i)
				}
			}
			// decode dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 optional
			if dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16Present {
				tmp_dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 := utils.SetupRelease[*DMRS_DownlinkConfig]{}
				if err = tmp_dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16", err)
				}
				ie.dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16 = tmp_dmrs_DownlinkForPDSCH_MappingTypeA_DCI_1_2_r16.Setup
			}
			// decode dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 optional
			if dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16Present {
				tmp_dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 := utils.SetupRelease[*DMRS_DownlinkConfig]{}
				if err = tmp_dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16", err)
				}
				ie.dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16 = tmp_dmrs_DownlinkForPDSCH_MappingTypeB_DCI_1_2_r16.Setup
			}
			// decode dmrs_SequenceInitializationDCI_1_2_r16 optional
			if dmrs_SequenceInitializationDCI_1_2_r16Present {
				ie.dmrs_SequenceInitializationDCI_1_2_r16 = new(PDSCH_Config_dmrs_SequenceInitializationDCI_1_2_r16)
				if err = ie.dmrs_SequenceInitializationDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode dmrs_SequenceInitializationDCI_1_2_r16", err)
				}
			}
			// decode harq_ProcessNumberSizeDCI_1_2_r16 optional
			if harq_ProcessNumberSizeDCI_1_2_r16Present {
				var tmp_int_harq_ProcessNumberSizeDCI_1_2_r16 int64
				if tmp_int_harq_ProcessNumberSizeDCI_1_2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
					return utils.WrapError("Decode harq_ProcessNumberSizeDCI_1_2_r16", err)
				}
				ie.harq_ProcessNumberSizeDCI_1_2_r16 = &tmp_int_harq_ProcessNumberSizeDCI_1_2_r16
			}
			// decode mcs_TableDCI_1_2_r16 optional
			if mcs_TableDCI_1_2_r16Present {
				ie.mcs_TableDCI_1_2_r16 = new(PDSCH_Config_mcs_TableDCI_1_2_r16)
				if err = ie.mcs_TableDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode mcs_TableDCI_1_2_r16", err)
				}
			}
			// decode numberOfBitsForRV_DCI_1_2_r16 optional
			if numberOfBitsForRV_DCI_1_2_r16Present {
				var tmp_int_numberOfBitsForRV_DCI_1_2_r16 int64
				if tmp_int_numberOfBitsForRV_DCI_1_2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
					return utils.WrapError("Decode numberOfBitsForRV_DCI_1_2_r16", err)
				}
				ie.numberOfBitsForRV_DCI_1_2_r16 = &tmp_int_numberOfBitsForRV_DCI_1_2_r16
			}
			// decode pdsch_TimeDomainAllocationListDCI_1_2_r16 optional
			if pdsch_TimeDomainAllocationListDCI_1_2_r16Present {
				tmp_pdsch_TimeDomainAllocationListDCI_1_2_r16 := utils.SetupRelease[*PDSCH_TimeDomainResourceAllocationList_r16]{}
				if err = tmp_pdsch_TimeDomainAllocationListDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_TimeDomainAllocationListDCI_1_2_r16", err)
				}
				ie.pdsch_TimeDomainAllocationListDCI_1_2_r16 = tmp_pdsch_TimeDomainAllocationListDCI_1_2_r16.Setup
			}
			// decode prb_BundlingTypeDCI_1_2_r16 optional
			if prb_BundlingTypeDCI_1_2_r16Present {
				ie.prb_BundlingTypeDCI_1_2_r16 = new(PDSCH_Config_prb_BundlingTypeDCI_1_2_r16)
				if err = ie.prb_BundlingTypeDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode prb_BundlingTypeDCI_1_2_r16", err)
				}
			}
			// decode priorityIndicatorDCI_1_2_r16 optional
			if priorityIndicatorDCI_1_2_r16Present {
				ie.priorityIndicatorDCI_1_2_r16 = new(PDSCH_Config_priorityIndicatorDCI_1_2_r16)
				if err = ie.priorityIndicatorDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode priorityIndicatorDCI_1_2_r16", err)
				}
			}
			// decode rateMatchPatternGroup1DCI_1_2_r16 optional
			if rateMatchPatternGroup1DCI_1_2_r16Present {
				ie.rateMatchPatternGroup1DCI_1_2_r16 = new(RateMatchPatternGroup)
				if err = ie.rateMatchPatternGroup1DCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode rateMatchPatternGroup1DCI_1_2_r16", err)
				}
			}
			// decode rateMatchPatternGroup2DCI_1_2_r16 optional
			if rateMatchPatternGroup2DCI_1_2_r16Present {
				ie.rateMatchPatternGroup2DCI_1_2_r16 = new(RateMatchPatternGroup)
				if err = ie.rateMatchPatternGroup2DCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode rateMatchPatternGroup2DCI_1_2_r16", err)
				}
			}
			// decode resourceAllocationType1GranularityDCI_1_2_r16 optional
			if resourceAllocationType1GranularityDCI_1_2_r16Present {
				ie.resourceAllocationType1GranularityDCI_1_2_r16 = new(PDSCH_Config_resourceAllocationType1GranularityDCI_1_2_r16)
				if err = ie.resourceAllocationType1GranularityDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode resourceAllocationType1GranularityDCI_1_2_r16", err)
				}
			}
			// decode vrb_ToPRB_InterleaverDCI_1_2_r16 optional
			if vrb_ToPRB_InterleaverDCI_1_2_r16Present {
				ie.vrb_ToPRB_InterleaverDCI_1_2_r16 = new(PDSCH_Config_vrb_ToPRB_InterleaverDCI_1_2_r16)
				if err = ie.vrb_ToPRB_InterleaverDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode vrb_ToPRB_InterleaverDCI_1_2_r16", err)
				}
			}
			// decode referenceOfSLIVDCI_1_2_r16 optional
			if referenceOfSLIVDCI_1_2_r16Present {
				ie.referenceOfSLIVDCI_1_2_r16 = new(PDSCH_Config_referenceOfSLIVDCI_1_2_r16)
				if err = ie.referenceOfSLIVDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode referenceOfSLIVDCI_1_2_r16", err)
				}
			}
			// decode resourceAllocationDCI_1_2_r16 optional
			if resourceAllocationDCI_1_2_r16Present {
				ie.resourceAllocationDCI_1_2_r16 = new(PDSCH_Config_resourceAllocationDCI_1_2_r16)
				if err = ie.resourceAllocationDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode resourceAllocationDCI_1_2_r16", err)
				}
			}
			// decode priorityIndicatorDCI_1_1_r16 optional
			if priorityIndicatorDCI_1_1_r16Present {
				ie.priorityIndicatorDCI_1_1_r16 = new(PDSCH_Config_priorityIndicatorDCI_1_1_r16)
				if err = ie.priorityIndicatorDCI_1_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode priorityIndicatorDCI_1_1_r16", err)
				}
			}
			// decode dataScramblingIdentityPDSCH2_r16 optional
			if dataScramblingIdentityPDSCH2_r16Present {
				var tmp_int_dataScramblingIdentityPDSCH2_r16 int64
				if tmp_int_dataScramblingIdentityPDSCH2_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
					return utils.WrapError("Decode dataScramblingIdentityPDSCH2_r16", err)
				}
				ie.dataScramblingIdentityPDSCH2_r16 = &tmp_int_dataScramblingIdentityPDSCH2_r16
			}
			// decode pdsch_TimeDomainAllocationList_r16 optional
			if pdsch_TimeDomainAllocationList_r16Present {
				tmp_pdsch_TimeDomainAllocationList_r16 := utils.SetupRelease[*PDSCH_TimeDomainResourceAllocationList_r16]{}
				if err = tmp_pdsch_TimeDomainAllocationList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_TimeDomainAllocationList_r16", err)
				}
				ie.pdsch_TimeDomainAllocationList_r16 = tmp_pdsch_TimeDomainAllocationList_r16.Setup
			}
			// decode repetitionSchemeConfig_r16 optional
			if repetitionSchemeConfig_r16Present {
				tmp_repetitionSchemeConfig_r16 := utils.SetupRelease[*RepetitionSchemeConfig_r16]{}
				if err = tmp_repetitionSchemeConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode repetitionSchemeConfig_r16", err)
				}
				ie.repetitionSchemeConfig_r16 = tmp_repetitionSchemeConfig_r16.Setup
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			repetitionSchemeConfig_v1630Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode repetitionSchemeConfig_v1630 optional
			if repetitionSchemeConfig_v1630Present {
				tmp_repetitionSchemeConfig_v1630 := utils.SetupRelease[*RepetitionSchemeConfig_v1630]{}
				if err = tmp_repetitionSchemeConfig_v1630.Decode(extReader); err != nil {
					return utils.WrapError("Decode repetitionSchemeConfig_v1630", err)
				}
				ie.repetitionSchemeConfig_v1630 = tmp_repetitionSchemeConfig_v1630.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_HARQ_ACK_EnhType3DCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_HARQ_ACK_RetxDCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pucch_sSCellDynDCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dl_OrJointTCI_StateList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			beamAppTime_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dummyPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dmrs_FD_OCC_DisabledForRank1_PDSCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			minimumSchedulingOffsetK0_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			harq_ProcessNumberSizeDCI_1_2_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			harq_ProcessNumberSizeDCI_1_1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mcs_Table_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mcs_TableDCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			xOverheadMulticast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			priorityIndicatorDCI_4_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sizeDCI_4_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17 optional
			if pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17Present {
				ie.pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17 = new(PDSCH_Config_pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17)
				if err = ie.pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_HARQ_ACK_OneShotFeedbackDCI_1_2_r17", err)
				}
			}
			// decode pdsch_HARQ_ACK_EnhType3DCI_1_2_r17 optional
			if pdsch_HARQ_ACK_EnhType3DCI_1_2_r17Present {
				ie.pdsch_HARQ_ACK_EnhType3DCI_1_2_r17 = new(PDSCH_Config_pdsch_HARQ_ACK_EnhType3DCI_1_2_r17)
				if err = ie.pdsch_HARQ_ACK_EnhType3DCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_HARQ_ACK_EnhType3DCI_1_2_r17", err)
				}
			}
			// decode pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17 optional
			if pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17Present {
				ie.pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17 = new(PDSCH_Config_pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17)
				if err = ie.pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_HARQ_ACK_EnhType3DCI_Field_1_2_r17", err)
				}
			}
			// decode pdsch_HARQ_ACK_RetxDCI_1_2_r17 optional
			if pdsch_HARQ_ACK_RetxDCI_1_2_r17Present {
				ie.pdsch_HARQ_ACK_RetxDCI_1_2_r17 = new(PDSCH_Config_pdsch_HARQ_ACK_RetxDCI_1_2_r17)
				if err = ie.pdsch_HARQ_ACK_RetxDCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_HARQ_ACK_RetxDCI_1_2_r17", err)
				}
			}
			// decode pucch_sSCellDynDCI_1_2_r17 optional
			if pucch_sSCellDynDCI_1_2_r17Present {
				ie.pucch_sSCellDynDCI_1_2_r17 = new(PDSCH_Config_pucch_sSCellDynDCI_1_2_r17)
				if err = ie.pucch_sSCellDynDCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pucch_sSCellDynDCI_1_2_r17", err)
				}
			}
			// decode dl_OrJointTCI_StateList_r17 optional
			if dl_OrJointTCI_StateList_r17Present {
				ie.dl_OrJointTCI_StateList_r17 = new(PDSCH_Config_dl_OrJointTCI_StateList_r17)
				if err = ie.dl_OrJointTCI_StateList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dl_OrJointTCI_StateList_r17", err)
				}
			}
			// decode beamAppTime_r17 optional
			if beamAppTime_r17Present {
				ie.beamAppTime_r17 = new(PDSCH_Config_beamAppTime_r17)
				if err = ie.beamAppTime_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode beamAppTime_r17", err)
				}
			}
			// decode dummy optional
			if dummyPresent {
				tmp_dummy := utils.SetupRelease[*Dummy_TDRA_List]{}
				if err = tmp_dummy.Decode(extReader); err != nil {
					return utils.WrapError("Decode dummy", err)
				}
				ie.dummy = tmp_dummy.Setup
			}
			// decode dmrs_FD_OCC_DisabledForRank1_PDSCH_r17 optional
			if dmrs_FD_OCC_DisabledForRank1_PDSCH_r17Present {
				ie.dmrs_FD_OCC_DisabledForRank1_PDSCH_r17 = new(PDSCH_Config_dmrs_FD_OCC_DisabledForRank1_PDSCH_r17)
				if err = ie.dmrs_FD_OCC_DisabledForRank1_PDSCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dmrs_FD_OCC_DisabledForRank1_PDSCH_r17", err)
				}
			}
			// decode minimumSchedulingOffsetK0_r17 optional
			if minimumSchedulingOffsetK0_r17Present {
				tmp_minimumSchedulingOffsetK0_r17 := utils.SetupRelease[*MinSchedulingOffsetK0_Values_r17]{}
				if err = tmp_minimumSchedulingOffsetK0_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode minimumSchedulingOffsetK0_r17", err)
				}
				ie.minimumSchedulingOffsetK0_r17 = tmp_minimumSchedulingOffsetK0_r17.Setup
			}
			// decode harq_ProcessNumberSizeDCI_1_2_v1700 optional
			if harq_ProcessNumberSizeDCI_1_2_v1700Present {
				var tmp_int_harq_ProcessNumberSizeDCI_1_2_v1700 int64
				if tmp_int_harq_ProcessNumberSizeDCI_1_2_v1700, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
					return utils.WrapError("Decode harq_ProcessNumberSizeDCI_1_2_v1700", err)
				}
				ie.harq_ProcessNumberSizeDCI_1_2_v1700 = &tmp_int_harq_ProcessNumberSizeDCI_1_2_v1700
			}
			// decode harq_ProcessNumberSizeDCI_1_1_r17 optional
			if harq_ProcessNumberSizeDCI_1_1_r17Present {
				var tmp_int_harq_ProcessNumberSizeDCI_1_1_r17 int64
				if tmp_int_harq_ProcessNumberSizeDCI_1_1_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
					return utils.WrapError("Decode harq_ProcessNumberSizeDCI_1_1_r17", err)
				}
				ie.harq_ProcessNumberSizeDCI_1_1_r17 = &tmp_int_harq_ProcessNumberSizeDCI_1_1_r17
			}
			// decode mcs_Table_r17 optional
			if mcs_Table_r17Present {
				ie.mcs_Table_r17 = new(PDSCH_Config_mcs_Table_r17)
				if err = ie.mcs_Table_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mcs_Table_r17", err)
				}
			}
			// decode mcs_TableDCI_1_2_r17 optional
			if mcs_TableDCI_1_2_r17Present {
				ie.mcs_TableDCI_1_2_r17 = new(PDSCH_Config_mcs_TableDCI_1_2_r17)
				if err = ie.mcs_TableDCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mcs_TableDCI_1_2_r17", err)
				}
			}
			// decode xOverheadMulticast_r17 optional
			if xOverheadMulticast_r17Present {
				ie.xOverheadMulticast_r17 = new(PDSCH_Config_xOverheadMulticast_r17)
				if err = ie.xOverheadMulticast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode xOverheadMulticast_r17", err)
				}
			}
			// decode priorityIndicatorDCI_4_2_r17 optional
			if priorityIndicatorDCI_4_2_r17Present {
				ie.priorityIndicatorDCI_4_2_r17 = new(PDSCH_Config_priorityIndicatorDCI_4_2_r17)
				if err = ie.priorityIndicatorDCI_4_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode priorityIndicatorDCI_4_2_r17", err)
				}
			}
			// decode sizeDCI_4_2_r17 optional
			if sizeDCI_4_2_r17Present {
				var tmp_int_sizeDCI_4_2_r17 int64
				if tmp_int_sizeDCI_4_2_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 20, Ub: maxDCI_4_2_Size_r17}, false); err != nil {
					return utils.WrapError("Decode sizeDCI_4_2_r17", err)
				}
				ie.sizeDCI_4_2_r17 = &tmp_int_sizeDCI_4_2_r17
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			pdsch_TimeDomainAllocationListForMultiPDSCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode pdsch_TimeDomainAllocationListForMultiPDSCH_r17 optional
			if pdsch_TimeDomainAllocationListForMultiPDSCH_r17Present {
				tmp_pdsch_TimeDomainAllocationListForMultiPDSCH_r17 := utils.SetupRelease[*MultiPDSCH_TDRA_List_r17]{}
				if err = tmp_pdsch_TimeDomainAllocationListForMultiPDSCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_TimeDomainAllocationListForMultiPDSCH_r17", err)
				}
				ie.pdsch_TimeDomainAllocationListForMultiPDSCH_r17 = tmp_pdsch_TimeDomainAllocationListForMultiPDSCH_r17.Setup
			}
		}
	}
	return nil
}
