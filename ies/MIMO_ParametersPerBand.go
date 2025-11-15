package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand struct {
	tci_StatePDSCH                               *MIMO_ParametersPerBand_tci_StatePDSCH                               `optional`
	additionalActiveTCI_StatePDCCH               *MIMO_ParametersPerBand_additionalActiveTCI_StatePDCCH               `optional`
	pusch_TransCoherence                         *MIMO_ParametersPerBand_pusch_TransCoherence                         `optional`
	beamCorrespondenceWithoutUL_BeamSweeping     *MIMO_ParametersPerBand_beamCorrespondenceWithoutUL_BeamSweeping     `optional`
	periodicBeamReport                           *MIMO_ParametersPerBand_periodicBeamReport                           `optional`
	aperiodicBeamReport                          *MIMO_ParametersPerBand_aperiodicBeamReport                          `optional`
	sp_BeamReportPUCCH                           *MIMO_ParametersPerBand_sp_BeamReportPUCCH                           `optional`
	sp_BeamReportPUSCH                           *MIMO_ParametersPerBand_sp_BeamReportPUSCH                           `optional`
	dummy1                                       *DummyG                                                              `optional`
	maxNumberRxBeam                              *int64                                                               `lb:2,ub:8,optional`
	maxNumberRxTxBeamSwitchDL                    *MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL                    `optional`
	maxNumberNonGroupBeamReporting               *MIMO_ParametersPerBand_maxNumberNonGroupBeamReporting               `optional`
	groupBeamReporting                           *MIMO_ParametersPerBand_groupBeamReporting                           `optional`
	uplinkBeamManagement                         *MIMO_ParametersPerBand_uplinkBeamManagement                         `lb:1,ub:8,optional`
	maxNumberCSI_RS_BFD                          *int64                                                               `lb:1,ub:64,optional`
	maxNumberSSB_BFD                             *int64                                                               `lb:1,ub:64,optional`
	maxNumberCSI_RS_SSB_CBD                      *int64                                                               `lb:1,ub:256,optional`
	dummy2                                       *MIMO_ParametersPerBand_dummy2                                       `optional`
	twoPortsPTRS_UL                              *MIMO_ParametersPerBand_twoPortsPTRS_UL                              `optional`
	dummy5                                       *SRS_Resources                                                       `optional`
	dummy3                                       *int64                                                               `lb:1,ub:4,optional`
	beamReportTiming                             *MIMO_ParametersPerBand_beamReportTiming                             `optional`
	ptrs_DensityRecommendationSetDL              *MIMO_ParametersPerBand_ptrs_DensityRecommendationSetDL              `optional`
	ptrs_DensityRecommendationSetUL              *MIMO_ParametersPerBand_ptrs_DensityRecommendationSetUL              `optional`
	dummy4                                       *DummyH                                                              `optional`
	aperiodicTRS                                 *MIMO_ParametersPerBand_aperiodicTRS                                 `optional`
	dummy6                                       *MIMO_ParametersPerBand_dummy6                                       `optional,ext-1`
	beamManagementSSB_CSI_RS                     *BeamManagementSSB_CSI_RS                                            `optional,ext-1`
	beamSwitchTiming                             *MIMO_ParametersPerBand_beamSwitchTiming                             `optional,ext-1`
	codebookParameters                           *CodebookParameters                                                  `optional,ext-1`
	csi_RS_IM_ReceptionForFeedback               *CSI_RS_IM_ReceptionForFeedback                                      `optional,ext-1`
	csi_RS_ProcFrameworkForSRS                   *CSI_RS_ProcFrameworkForSRS                                          `optional,ext-1`
	csi_ReportFramework                          *CSI_ReportFramework                                                 `optional,ext-1`
	csi_RS_ForTracking                           *CSI_RS_ForTracking                                                  `optional,ext-1`
	srs_AssocCSI_RS                              []SupportedCSI_RS_Resource                                           `lb:1,ub:maxNrofCSI_RS_Resources,optional,ext-1`
	spatialRelations                             *SpatialRelations                                                    `optional,ext-1`
	defaultQCL_TwoTCI_r16                        *MIMO_ParametersPerBand_defaultQCL_TwoTCI_r16                        `optional,ext-2`
	codebookParametersPerBand_r16                *CodebookParameters_v1610                                            `optional,ext-2`
	simul_SpatialRelationUpdatePUCCHResGroup_r16 *MIMO_ParametersPerBand_simul_SpatialRelationUpdatePUCCHResGroup_r16 `optional,ext-2`
	maxNumberSCellBFR_r16                        *MIMO_ParametersPerBand_maxNumberSCellBFR_r16                        `optional,ext-2`
	simultaneousReceptionDiffTypeD_r16           *MIMO_ParametersPerBand_simultaneousReceptionDiffTypeD_r16           `optional,ext-2`
	ssb_csirs_SINR_measurement_r16               *MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16               `optional,ext-2`
	nonGroupSINR_reporting_r16                   *MIMO_ParametersPerBand_nonGroupSINR_reporting_r16                   `optional,ext-2`
	groupSINR_reporting_r16                      *MIMO_ParametersPerBand_groupSINR_reporting_r16                      `optional,ext-2`
	multiDCI_multiTRP_Parameters_r16             *MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16             `lb:1,ub:2,optional,ext-2`
	singleDCI_SDM_scheme_Parameters_r16          *MIMO_ParametersPerBand_singleDCI_SDM_scheme_Parameters_r16          `optional,ext-2`
	supportFDM_SchemeA_r16                       *MIMO_ParametersPerBand_supportFDM_SchemeA_r16                       `optional,ext-2`
	supportCodeWordSoftCombining_r16             *MIMO_ParametersPerBand_supportCodeWordSoftCombining_r16             `optional,ext-2`
	supportTDM_SchemeA_r16                       *MIMO_ParametersPerBand_supportTDM_SchemeA_r16                       `optional,ext-2`
	supportInter_slotTDM_r16                     *MIMO_ParametersPerBand_supportInter_slotTDM_r16                     `lb:1,ub:2,optional,ext-2`
	lowPAPR_DMRS_PDSCH_r16                       *MIMO_ParametersPerBand_lowPAPR_DMRS_PDSCH_r16                       `optional,ext-2`
	lowPAPR_DMRS_PUSCHwithoutPrecoding_r16       *MIMO_ParametersPerBand_lowPAPR_DMRS_PUSCHwithoutPrecoding_r16       `optional,ext-2`
	lowPAPR_DMRS_PUCCH_r16                       *MIMO_ParametersPerBand_lowPAPR_DMRS_PUCCH_r16                       `optional,ext-2`
	lowPAPR_DMRS_PUSCHwithPrecoding_r16          *MIMO_ParametersPerBand_lowPAPR_DMRS_PUSCHwithPrecoding_r16          `optional,ext-2`
	csi_ReportFrameworkExt_r16                   *CSI_ReportFrameworkExt_r16                                          `optional,ext-2`
	codebookParametersAddition_r16               *CodebookParametersAddition_r16                                      `optional,ext-2`
	codebookComboParametersAddition_r16          *CodebookComboParametersAddition_r16                                 `optional,ext-2`
	beamCorrespondenceSSB_based_r16              *MIMO_ParametersPerBand_beamCorrespondenceSSB_based_r16              `optional,ext-2`
	beamCorrespondenceCSI_RS_based_r16           *MIMO_ParametersPerBand_beamCorrespondenceCSI_RS_based_r16           `optional,ext-2`
	beamSwitchTiming_r16                         *MIMO_ParametersPerBand_beamSwitchTiming_r16                         `optional,ext-2`
	semi_PersistentL1_SINR_Report_PUCCH_r16      *MIMO_ParametersPerBand_semi_PersistentL1_SINR_Report_PUCCH_r16      `optional,ext-3`
	semi_PersistentL1_SINR_Report_PUSCH_r16      *MIMO_ParametersPerBand_semi_PersistentL1_SINR_Report_PUSCH_r16      `optional,ext-3`
	spatialRelations_v1640                       *MIMO_ParametersPerBand_spatialRelations_v1640                       `optional,ext-4`
	support64CandidateBeamRS_BFR_r16             *MIMO_ParametersPerBand_support64CandidateBeamRS_BFR_r16             `optional,ext-4`
	maxMIMO_LayersForMulti_DCI_mTRP_r16          *MIMO_ParametersPerBand_maxMIMO_LayersForMulti_DCI_mTRP_r16          `optional,ext-5`
	supportedSINR_meas_v1670                     *uper.BitString                                                      `lb:4,ub:4,optional,ext-6`
	srs_increasedRepetition_r17                  *MIMO_ParametersPerBand_srs_increasedRepetition_r17                  `optional,ext-7`
	srs_partialFrequencySounding_r17             *MIMO_ParametersPerBand_srs_partialFrequencySounding_r17             `optional,ext-7`
	srs_startRB_locationHoppingPartial_r17       *MIMO_ParametersPerBand_srs_startRB_locationHoppingPartial_r17       `optional,ext-7`
	srs_combEight_r17                            *MIMO_ParametersPerBand_srs_combEight_r17                            `optional,ext-7`
	codebookParametersfetype2_r17                *CodebookParametersfetype2_r17                                       `optional,ext-7`
	mTRP_PUSCH_twoCSI_RS_r17                     *MIMO_ParametersPerBand_mTRP_PUSCH_twoCSI_RS_r17                     `optional,ext-7`
	mTRP_PUCCH_InterSlot_r17                     *MIMO_ParametersPerBand_mTRP_PUCCH_InterSlot_r17                     `optional,ext-7`
	mTRP_PUCCH_CyclicMapping_r17                 *MIMO_ParametersPerBand_mTRP_PUCCH_CyclicMapping_r17                 `optional,ext-7`
	mTRP_PUCCH_SecondTPC_r17                     *MIMO_ParametersPerBand_mTRP_PUCCH_SecondTPC_r17                     `optional,ext-7`
	mTRP_BFR_twoBFD_RS_Set_r17                   *MIMO_ParametersPerBand_mTRP_BFR_twoBFD_RS_Set_r17                   `lb:1,ub:9,optional,ext-7`
	mTRP_BFR_PUCCH_SR_perCG_r17                  *MIMO_ParametersPerBand_mTRP_BFR_PUCCH_SR_perCG_r17                  `optional,ext-7`
	mTRP_BFR_association_PUCCH_SR_r17            *MIMO_ParametersPerBand_mTRP_BFR_association_PUCCH_SR_r17            `optional,ext-7`
	sfn_SimulTwoTCI_AcrossMultiCC_r17            *MIMO_ParametersPerBand_sfn_SimulTwoTCI_AcrossMultiCC_r17            `optional,ext-7`
	sfn_DefaultDL_BeamSetup_r17                  *MIMO_ParametersPerBand_sfn_DefaultDL_BeamSetup_r17                  `optional,ext-7`
	sfn_DefaultUL_BeamSetup_r17                  *MIMO_ParametersPerBand_sfn_DefaultUL_BeamSetup_r17                  `optional,ext-7`
	srs_TriggeringOffset_r17                     *MIMO_ParametersPerBand_srs_TriggeringOffset_r17                     `optional,ext-7`
	srs_TriggeringDCI_r17                        *MIMO_ParametersPerBand_srs_TriggeringDCI_r17                        `optional,ext-7`
	codebookComboParameterMixedType_r17          *CodebookComboParameterMixedType_r17                                 `optional,ext-7`
	unifiedJointTCI_r17                          *MIMO_ParametersPerBand_unifiedJointTCI_r17                          `optional,ext-7`
	unifiedJointTCI_multiMAC_CE_r17              *MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17              `optional,ext-7`
	unifiedJointTCI_perBWP_CA_r17                *MIMO_ParametersPerBand_unifiedJointTCI_perBWP_CA_r17                `optional,ext-7`
	unifiedJointTCI_ListSharingCA_r17            *MIMO_ParametersPerBand_unifiedJointTCI_ListSharingCA_r17            `optional,ext-7`
	unifiedJointTCI_commonMultiCC_r17            *MIMO_ParametersPerBand_unifiedJointTCI_commonMultiCC_r17            `optional,ext-7`
	unifiedJointTCI_BeamAlignDLRS_r17            *MIMO_ParametersPerBand_unifiedJointTCI_BeamAlignDLRS_r17            `optional,ext-7`
	unifiedJointTCI_PC_association_r17           *MIMO_ParametersPerBand_unifiedJointTCI_PC_association_r17           `optional,ext-7`
	unifiedJointTCI_Legacy_r17                   *MIMO_ParametersPerBand_unifiedJointTCI_Legacy_r17                   `optional,ext-7`
	unifiedJointTCI_Legacy_SRS_r17               *MIMO_ParametersPerBand_unifiedJointTCI_Legacy_SRS_r17               `optional,ext-7`
	unifiedJointTCI_Legacy_CORESET0_r17          *MIMO_ParametersPerBand_unifiedJointTCI_Legacy_CORESET0_r17          `optional,ext-7`
	unifiedJointTCI_SCellBFR_r17                 *MIMO_ParametersPerBand_unifiedJointTCI_SCellBFR_r17                 `optional,ext-7`
	unifiedJointTCI_InterCell_r17                *MIMO_ParametersPerBand_unifiedJointTCI_InterCell_r17                `optional,ext-7`
	unifiedSeparateTCI_r17                       *MIMO_ParametersPerBand_unifiedSeparateTCI_r17                       `optional,ext-7`
	unifiedSeparateTCI_multiMAC_CE_r17           *MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17           `lb:2,ub:8,optional,ext-7`
	unifiedSeparateTCI_perBWP_CA_r17             *MIMO_ParametersPerBand_unifiedSeparateTCI_perBWP_CA_r17             `optional,ext-7`
	unifiedSeparateTCI_ListSharingCA_r17         *MIMO_ParametersPerBand_unifiedSeparateTCI_ListSharingCA_r17         `optional,ext-7`
	unifiedSeparateTCI_commonMultiCC_r17         *MIMO_ParametersPerBand_unifiedSeparateTCI_commonMultiCC_r17         `optional,ext-7`
	unifiedSeparateTCI_InterCell_r17             *MIMO_ParametersPerBand_unifiedSeparateTCI_InterCell_r17             `optional,ext-7`
	unifiedJointTCI_mTRP_InterCell_BM_r17        *MIMO_ParametersPerBand_unifiedJointTCI_mTRP_InterCell_BM_r17        `lb:1,ub:7,optional,ext-7`
	mpe_Mitigation_r17                           *MIMO_ParametersPerBand_mpe_Mitigation_r17                           `lb:1,ub:4,optional,ext-7`
	srs_PortReport_r17                           *MIMO_ParametersPerBand_srs_PortReport_r17                           `optional,ext-7`
	mTRP_PDCCH_individual_r17                    *MIMO_ParametersPerBand_mTRP_PDCCH_individual_r17                    `optional,ext-7`
	mTRP_PDCCH_anySpan_3Symbols_r17              *MIMO_ParametersPerBand_mTRP_PDCCH_anySpan_3Symbols_r17              `optional,ext-7`
	mTRP_PDCCH_TwoQCL_TypeD_r17                  *MIMO_ParametersPerBand_mTRP_PDCCH_TwoQCL_TypeD_r17                  `optional,ext-7`
	mTRP_PUSCH_CSI_RS_r17                        *MIMO_ParametersPerBand_mTRP_PUSCH_CSI_RS_r17                        `lb:1,ub:8,optional,ext-7`
	mTRP_PUSCH_cyclicMapping_r17                 *MIMO_ParametersPerBand_mTRP_PUSCH_cyclicMapping_r17                 `optional,ext-7`
	mTRP_PUSCH_secondTPC_r17                     *MIMO_ParametersPerBand_mTRP_PUSCH_secondTPC_r17                     `optional,ext-7`
	mTRP_PUSCH_twoPHR_Reporting_r17              *MIMO_ParametersPerBand_mTRP_PUSCH_twoPHR_Reporting_r17              `optional,ext-7`
	mTRP_PUSCH_A_CSI_r17                         *MIMO_ParametersPerBand_mTRP_PUSCH_A_CSI_r17                         `optional,ext-7`
	mTRP_PUSCH_SP_CSI_r17                        *MIMO_ParametersPerBand_mTRP_PUSCH_SP_CSI_r17                        `optional,ext-7`
	mTRP_PUSCH_CG_r17                            *MIMO_ParametersPerBand_mTRP_PUSCH_CG_r17                            `optional,ext-7`
	mTRP_PUCCH_MAC_CE_r17                        *MIMO_ParametersPerBand_mTRP_PUCCH_MAC_CE_r17                        `optional,ext-7`
	mTRP_PUCCH_maxNum_PC_FR1_r17                 *int64                                                               `lb:3,ub:8,optional,ext-7`
	mTRP_inter_Cell_r17                          *MIMO_ParametersPerBand_mTRP_inter_Cell_r17                          `lb:1,ub:7,optional,ext-7`
	mTRP_GroupBasedL1_RSRP_r17                   *MIMO_ParametersPerBand_mTRP_GroupBasedL1_RSRP_r17                   `lb:1,ub:4,optional,ext-7`
	mTRP_BFD_RS_MAC_CE_r17                       *MIMO_ParametersPerBand_mTRP_BFD_RS_MAC_CE_r17                       `optional,ext-7`
	mTRP_CSI_EnhancementPerBand_r17              *MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17              `lb:1,ub:16,optional,ext-7`
	codebookComboParameterMultiTRP_r17           *CodebookComboParameterMultiTRP_r17                                  `optional,ext-7`
	mTRP_CSI_additionalCSI_r17                   *MIMO_ParametersPerBand_mTRP_CSI_additionalCSI_r17                   `optional,ext-7`
	mTRP_CSI_N_Max2_r17                          *MIMO_ParametersPerBand_mTRP_CSI_N_Max2_r17                          `optional,ext-7`
	mTRP_CSI_CMR_r17                             *MIMO_ParametersPerBand_mTRP_CSI_CMR_r17                             `optional,ext-7`
	srs_partialFreqSounding_r17                  *MIMO_ParametersPerBand_srs_partialFreqSounding_r17                  `optional,ext-7`
	beamSwitchTiming_v1710                       *MIMO_ParametersPerBand_beamSwitchTiming_v1710                       `optional,ext-7`
	beamSwitchTiming_r17                         *MIMO_ParametersPerBand_beamSwitchTiming_r17                         `optional,ext-7`
	beamReportTiming_v1710                       *MIMO_ParametersPerBand_beamReportTiming_v1710                       `optional,ext-7`
	maxNumberRxTxBeamSwitchDL_v1710              *MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_v1710              `optional,ext-7`
	srs_PortReportSP_AP_r17                      *MIMO_ParametersPerBand_srs_PortReportSP_AP_r17                      `optional,ext-8`
	maxNumberRxBeam_v1720                        *int64                                                               `lb:9,ub:12,optional,ext-8`
	sfn_ImplicitRS_twoTCI_r17                    *MIMO_ParametersPerBand_sfn_ImplicitRS_twoTCI_r17                    `optional,ext-8`
	sfn_QCL_TypeD_Collision_twoTCI_r17           *MIMO_ParametersPerBand_sfn_QCL_TypeD_Collision_twoTCI_r17           `optional,ext-8`
	mTRP_CSI_numCPU_r17                          *MIMO_ParametersPerBand_mTRP_CSI_numCPU_r17                          `optional,ext-8`
	supportRepNumPDSCH_TDRA_DCI_1_2_r17          *MIMO_ParametersPerBand_supportRepNumPDSCH_TDRA_DCI_1_2_r17          `optional,ext-9`
}

func (ie *MIMO_ParametersPerBand) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.dummy6 != nil || ie.beamManagementSSB_CSI_RS != nil || ie.beamSwitchTiming != nil || ie.codebookParameters != nil || ie.csi_RS_IM_ReceptionForFeedback != nil || ie.csi_RS_ProcFrameworkForSRS != nil || ie.csi_ReportFramework != nil || ie.csi_RS_ForTracking != nil || len(ie.srs_AssocCSI_RS) > 0 || ie.spatialRelations != nil || ie.defaultQCL_TwoTCI_r16 != nil || ie.codebookParametersPerBand_r16 != nil || ie.simul_SpatialRelationUpdatePUCCHResGroup_r16 != nil || ie.maxNumberSCellBFR_r16 != nil || ie.simultaneousReceptionDiffTypeD_r16 != nil || ie.ssb_csirs_SINR_measurement_r16 != nil || ie.nonGroupSINR_reporting_r16 != nil || ie.groupSINR_reporting_r16 != nil || ie.multiDCI_multiTRP_Parameters_r16 != nil || ie.singleDCI_SDM_scheme_Parameters_r16 != nil || ie.supportFDM_SchemeA_r16 != nil || ie.supportCodeWordSoftCombining_r16 != nil || ie.supportTDM_SchemeA_r16 != nil || ie.supportInter_slotTDM_r16 != nil || ie.lowPAPR_DMRS_PDSCH_r16 != nil || ie.lowPAPR_DMRS_PUSCHwithoutPrecoding_r16 != nil || ie.lowPAPR_DMRS_PUCCH_r16 != nil || ie.lowPAPR_DMRS_PUSCHwithPrecoding_r16 != nil || ie.csi_ReportFrameworkExt_r16 != nil || ie.codebookParametersAddition_r16 != nil || ie.codebookComboParametersAddition_r16 != nil || ie.beamCorrespondenceSSB_based_r16 != nil || ie.beamCorrespondenceCSI_RS_based_r16 != nil || ie.beamSwitchTiming_r16 != nil || ie.semi_PersistentL1_SINR_Report_PUCCH_r16 != nil || ie.semi_PersistentL1_SINR_Report_PUSCH_r16 != nil || ie.spatialRelations_v1640 != nil || ie.support64CandidateBeamRS_BFR_r16 != nil || ie.maxMIMO_LayersForMulti_DCI_mTRP_r16 != nil || ie.supportedSINR_meas_v1670 != nil || ie.srs_increasedRepetition_r17 != nil || ie.srs_partialFrequencySounding_r17 != nil || ie.srs_startRB_locationHoppingPartial_r17 != nil || ie.srs_combEight_r17 != nil || ie.codebookParametersfetype2_r17 != nil || ie.mTRP_PUSCH_twoCSI_RS_r17 != nil || ie.mTRP_PUCCH_InterSlot_r17 != nil || ie.mTRP_PUCCH_CyclicMapping_r17 != nil || ie.mTRP_PUCCH_SecondTPC_r17 != nil || ie.mTRP_BFR_twoBFD_RS_Set_r17 != nil || ie.mTRP_BFR_PUCCH_SR_perCG_r17 != nil || ie.mTRP_BFR_association_PUCCH_SR_r17 != nil || ie.sfn_SimulTwoTCI_AcrossMultiCC_r17 != nil || ie.sfn_DefaultDL_BeamSetup_r17 != nil || ie.sfn_DefaultUL_BeamSetup_r17 != nil || ie.srs_TriggeringOffset_r17 != nil || ie.srs_TriggeringDCI_r17 != nil || ie.codebookComboParameterMixedType_r17 != nil || ie.unifiedJointTCI_r17 != nil || ie.unifiedJointTCI_multiMAC_CE_r17 != nil || ie.unifiedJointTCI_perBWP_CA_r17 != nil || ie.unifiedJointTCI_ListSharingCA_r17 != nil || ie.unifiedJointTCI_commonMultiCC_r17 != nil || ie.unifiedJointTCI_BeamAlignDLRS_r17 != nil || ie.unifiedJointTCI_PC_association_r17 != nil || ie.unifiedJointTCI_Legacy_r17 != nil || ie.unifiedJointTCI_Legacy_SRS_r17 != nil || ie.unifiedJointTCI_Legacy_CORESET0_r17 != nil || ie.unifiedJointTCI_SCellBFR_r17 != nil || ie.unifiedJointTCI_InterCell_r17 != nil || ie.unifiedSeparateTCI_r17 != nil || ie.unifiedSeparateTCI_multiMAC_CE_r17 != nil || ie.unifiedSeparateTCI_perBWP_CA_r17 != nil || ie.unifiedSeparateTCI_ListSharingCA_r17 != nil || ie.unifiedSeparateTCI_commonMultiCC_r17 != nil || ie.unifiedSeparateTCI_InterCell_r17 != nil || ie.unifiedJointTCI_mTRP_InterCell_BM_r17 != nil || ie.mpe_Mitigation_r17 != nil || ie.srs_PortReport_r17 != nil || ie.mTRP_PDCCH_individual_r17 != nil || ie.mTRP_PDCCH_anySpan_3Symbols_r17 != nil || ie.mTRP_PDCCH_TwoQCL_TypeD_r17 != nil || ie.mTRP_PUSCH_CSI_RS_r17 != nil || ie.mTRP_PUSCH_cyclicMapping_r17 != nil || ie.mTRP_PUSCH_secondTPC_r17 != nil || ie.mTRP_PUSCH_twoPHR_Reporting_r17 != nil || ie.mTRP_PUSCH_A_CSI_r17 != nil || ie.mTRP_PUSCH_SP_CSI_r17 != nil || ie.mTRP_PUSCH_CG_r17 != nil || ie.mTRP_PUCCH_MAC_CE_r17 != nil || ie.mTRP_PUCCH_maxNum_PC_FR1_r17 != nil || ie.mTRP_inter_Cell_r17 != nil || ie.mTRP_GroupBasedL1_RSRP_r17 != nil || ie.mTRP_BFD_RS_MAC_CE_r17 != nil || ie.mTRP_CSI_EnhancementPerBand_r17 != nil || ie.codebookComboParameterMultiTRP_r17 != nil || ie.mTRP_CSI_additionalCSI_r17 != nil || ie.mTRP_CSI_N_Max2_r17 != nil || ie.mTRP_CSI_CMR_r17 != nil || ie.srs_partialFreqSounding_r17 != nil || ie.beamSwitchTiming_v1710 != nil || ie.beamSwitchTiming_r17 != nil || ie.beamReportTiming_v1710 != nil || ie.maxNumberRxTxBeamSwitchDL_v1710 != nil || ie.srs_PortReportSP_AP_r17 != nil || ie.maxNumberRxBeam_v1720 != nil || ie.sfn_ImplicitRS_twoTCI_r17 != nil || ie.sfn_QCL_TypeD_Collision_twoTCI_r17 != nil || ie.mTRP_CSI_numCPU_r17 != nil || ie.supportRepNumPDSCH_TDRA_DCI_1_2_r17 != nil
	preambleBits := []bool{hasExtensions, ie.tci_StatePDSCH != nil, ie.additionalActiveTCI_StatePDCCH != nil, ie.pusch_TransCoherence != nil, ie.beamCorrespondenceWithoutUL_BeamSweeping != nil, ie.periodicBeamReport != nil, ie.aperiodicBeamReport != nil, ie.sp_BeamReportPUCCH != nil, ie.sp_BeamReportPUSCH != nil, ie.dummy1 != nil, ie.maxNumberRxBeam != nil, ie.maxNumberRxTxBeamSwitchDL != nil, ie.maxNumberNonGroupBeamReporting != nil, ie.groupBeamReporting != nil, ie.uplinkBeamManagement != nil, ie.maxNumberCSI_RS_BFD != nil, ie.maxNumberSSB_BFD != nil, ie.maxNumberCSI_RS_SSB_CBD != nil, ie.dummy2 != nil, ie.twoPortsPTRS_UL != nil, ie.dummy5 != nil, ie.dummy3 != nil, ie.beamReportTiming != nil, ie.ptrs_DensityRecommendationSetDL != nil, ie.ptrs_DensityRecommendationSetUL != nil, ie.dummy4 != nil, ie.aperiodicTRS != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.tci_StatePDSCH != nil {
		if err = ie.tci_StatePDSCH.Encode(w); err != nil {
			return utils.WrapError("Encode tci_StatePDSCH", err)
		}
	}
	if ie.additionalActiveTCI_StatePDCCH != nil {
		if err = ie.additionalActiveTCI_StatePDCCH.Encode(w); err != nil {
			return utils.WrapError("Encode additionalActiveTCI_StatePDCCH", err)
		}
	}
	if ie.pusch_TransCoherence != nil {
		if err = ie.pusch_TransCoherence.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_TransCoherence", err)
		}
	}
	if ie.beamCorrespondenceWithoutUL_BeamSweeping != nil {
		if err = ie.beamCorrespondenceWithoutUL_BeamSweeping.Encode(w); err != nil {
			return utils.WrapError("Encode beamCorrespondenceWithoutUL_BeamSweeping", err)
		}
	}
	if ie.periodicBeamReport != nil {
		if err = ie.periodicBeamReport.Encode(w); err != nil {
			return utils.WrapError("Encode periodicBeamReport", err)
		}
	}
	if ie.aperiodicBeamReport != nil {
		if err = ie.aperiodicBeamReport.Encode(w); err != nil {
			return utils.WrapError("Encode aperiodicBeamReport", err)
		}
	}
	if ie.sp_BeamReportPUCCH != nil {
		if err = ie.sp_BeamReportPUCCH.Encode(w); err != nil {
			return utils.WrapError("Encode sp_BeamReportPUCCH", err)
		}
	}
	if ie.sp_BeamReportPUSCH != nil {
		if err = ie.sp_BeamReportPUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode sp_BeamReportPUSCH", err)
		}
	}
	if ie.dummy1 != nil {
		if err = ie.dummy1.Encode(w); err != nil {
			return utils.WrapError("Encode dummy1", err)
		}
	}
	if ie.maxNumberRxBeam != nil {
		if err = w.WriteInteger(*ie.maxNumberRxBeam, &uper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode maxNumberRxBeam", err)
		}
	}
	if ie.maxNumberRxTxBeamSwitchDL != nil {
		if err = ie.maxNumberRxTxBeamSwitchDL.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumberRxTxBeamSwitchDL", err)
		}
	}
	if ie.maxNumberNonGroupBeamReporting != nil {
		if err = ie.maxNumberNonGroupBeamReporting.Encode(w); err != nil {
			return utils.WrapError("Encode maxNumberNonGroupBeamReporting", err)
		}
	}
	if ie.groupBeamReporting != nil {
		if err = ie.groupBeamReporting.Encode(w); err != nil {
			return utils.WrapError("Encode groupBeamReporting", err)
		}
	}
	if ie.uplinkBeamManagement != nil {
		if err = ie.uplinkBeamManagement.Encode(w); err != nil {
			return utils.WrapError("Encode uplinkBeamManagement", err)
		}
	}
	if ie.maxNumberCSI_RS_BFD != nil {
		if err = w.WriteInteger(*ie.maxNumberCSI_RS_BFD, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Encode maxNumberCSI_RS_BFD", err)
		}
	}
	if ie.maxNumberSSB_BFD != nil {
		if err = w.WriteInteger(*ie.maxNumberSSB_BFD, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Encode maxNumberSSB_BFD", err)
		}
	}
	if ie.maxNumberCSI_RS_SSB_CBD != nil {
		if err = w.WriteInteger(*ie.maxNumberCSI_RS_SSB_CBD, &uper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
			return utils.WrapError("Encode maxNumberCSI_RS_SSB_CBD", err)
		}
	}
	if ie.dummy2 != nil {
		if err = ie.dummy2.Encode(w); err != nil {
			return utils.WrapError("Encode dummy2", err)
		}
	}
	if ie.twoPortsPTRS_UL != nil {
		if err = ie.twoPortsPTRS_UL.Encode(w); err != nil {
			return utils.WrapError("Encode twoPortsPTRS_UL", err)
		}
	}
	if ie.dummy5 != nil {
		if err = ie.dummy5.Encode(w); err != nil {
			return utils.WrapError("Encode dummy5", err)
		}
	}
	if ie.dummy3 != nil {
		if err = w.WriteInteger(*ie.dummy3, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
			return utils.WrapError("Encode dummy3", err)
		}
	}
	if ie.beamReportTiming != nil {
		if err = ie.beamReportTiming.Encode(w); err != nil {
			return utils.WrapError("Encode beamReportTiming", err)
		}
	}
	if ie.ptrs_DensityRecommendationSetDL != nil {
		if err = ie.ptrs_DensityRecommendationSetDL.Encode(w); err != nil {
			return utils.WrapError("Encode ptrs_DensityRecommendationSetDL", err)
		}
	}
	if ie.ptrs_DensityRecommendationSetUL != nil {
		if err = ie.ptrs_DensityRecommendationSetUL.Encode(w); err != nil {
			return utils.WrapError("Encode ptrs_DensityRecommendationSetUL", err)
		}
	}
	if ie.dummy4 != nil {
		if err = ie.dummy4.Encode(w); err != nil {
			return utils.WrapError("Encode dummy4", err)
		}
	}
	if ie.aperiodicTRS != nil {
		if err = ie.aperiodicTRS.Encode(w); err != nil {
			return utils.WrapError("Encode aperiodicTRS", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 9 bits for 9 extension groups
		extBitmap := []bool{ie.dummy6 != nil || ie.beamManagementSSB_CSI_RS != nil || ie.beamSwitchTiming != nil || ie.codebookParameters != nil || ie.csi_RS_IM_ReceptionForFeedback != nil || ie.csi_RS_ProcFrameworkForSRS != nil || ie.csi_ReportFramework != nil || ie.csi_RS_ForTracking != nil || len(ie.srs_AssocCSI_RS) > 0 || ie.spatialRelations != nil, ie.defaultQCL_TwoTCI_r16 != nil || ie.codebookParametersPerBand_r16 != nil || ie.simul_SpatialRelationUpdatePUCCHResGroup_r16 != nil || ie.maxNumberSCellBFR_r16 != nil || ie.simultaneousReceptionDiffTypeD_r16 != nil || ie.ssb_csirs_SINR_measurement_r16 != nil || ie.nonGroupSINR_reporting_r16 != nil || ie.groupSINR_reporting_r16 != nil || ie.multiDCI_multiTRP_Parameters_r16 != nil || ie.singleDCI_SDM_scheme_Parameters_r16 != nil || ie.supportFDM_SchemeA_r16 != nil || ie.supportCodeWordSoftCombining_r16 != nil || ie.supportTDM_SchemeA_r16 != nil || ie.supportInter_slotTDM_r16 != nil || ie.lowPAPR_DMRS_PDSCH_r16 != nil || ie.lowPAPR_DMRS_PUSCHwithoutPrecoding_r16 != nil || ie.lowPAPR_DMRS_PUCCH_r16 != nil || ie.lowPAPR_DMRS_PUSCHwithPrecoding_r16 != nil || ie.csi_ReportFrameworkExt_r16 != nil || ie.codebookParametersAddition_r16 != nil || ie.codebookComboParametersAddition_r16 != nil || ie.beamCorrespondenceSSB_based_r16 != nil || ie.beamCorrespondenceCSI_RS_based_r16 != nil || ie.beamSwitchTiming_r16 != nil, ie.semi_PersistentL1_SINR_Report_PUCCH_r16 != nil || ie.semi_PersistentL1_SINR_Report_PUSCH_r16 != nil, ie.spatialRelations_v1640 != nil || ie.support64CandidateBeamRS_BFR_r16 != nil, ie.maxMIMO_LayersForMulti_DCI_mTRP_r16 != nil, ie.supportedSINR_meas_v1670 != nil, ie.srs_increasedRepetition_r17 != nil || ie.srs_partialFrequencySounding_r17 != nil || ie.srs_startRB_locationHoppingPartial_r17 != nil || ie.srs_combEight_r17 != nil || ie.codebookParametersfetype2_r17 != nil || ie.mTRP_PUSCH_twoCSI_RS_r17 != nil || ie.mTRP_PUCCH_InterSlot_r17 != nil || ie.mTRP_PUCCH_CyclicMapping_r17 != nil || ie.mTRP_PUCCH_SecondTPC_r17 != nil || ie.mTRP_BFR_twoBFD_RS_Set_r17 != nil || ie.mTRP_BFR_PUCCH_SR_perCG_r17 != nil || ie.mTRP_BFR_association_PUCCH_SR_r17 != nil || ie.sfn_SimulTwoTCI_AcrossMultiCC_r17 != nil || ie.sfn_DefaultDL_BeamSetup_r17 != nil || ie.sfn_DefaultUL_BeamSetup_r17 != nil || ie.srs_TriggeringOffset_r17 != nil || ie.srs_TriggeringDCI_r17 != nil || ie.codebookComboParameterMixedType_r17 != nil || ie.unifiedJointTCI_r17 != nil || ie.unifiedJointTCI_multiMAC_CE_r17 != nil || ie.unifiedJointTCI_perBWP_CA_r17 != nil || ie.unifiedJointTCI_ListSharingCA_r17 != nil || ie.unifiedJointTCI_commonMultiCC_r17 != nil || ie.unifiedJointTCI_BeamAlignDLRS_r17 != nil || ie.unifiedJointTCI_PC_association_r17 != nil || ie.unifiedJointTCI_Legacy_r17 != nil || ie.unifiedJointTCI_Legacy_SRS_r17 != nil || ie.unifiedJointTCI_Legacy_CORESET0_r17 != nil || ie.unifiedJointTCI_SCellBFR_r17 != nil || ie.unifiedJointTCI_InterCell_r17 != nil || ie.unifiedSeparateTCI_r17 != nil || ie.unifiedSeparateTCI_multiMAC_CE_r17 != nil || ie.unifiedSeparateTCI_perBWP_CA_r17 != nil || ie.unifiedSeparateTCI_ListSharingCA_r17 != nil || ie.unifiedSeparateTCI_commonMultiCC_r17 != nil || ie.unifiedSeparateTCI_InterCell_r17 != nil || ie.unifiedJointTCI_mTRP_InterCell_BM_r17 != nil || ie.mpe_Mitigation_r17 != nil || ie.srs_PortReport_r17 != nil || ie.mTRP_PDCCH_individual_r17 != nil || ie.mTRP_PDCCH_anySpan_3Symbols_r17 != nil || ie.mTRP_PDCCH_TwoQCL_TypeD_r17 != nil || ie.mTRP_PUSCH_CSI_RS_r17 != nil || ie.mTRP_PUSCH_cyclicMapping_r17 != nil || ie.mTRP_PUSCH_secondTPC_r17 != nil || ie.mTRP_PUSCH_twoPHR_Reporting_r17 != nil || ie.mTRP_PUSCH_A_CSI_r17 != nil || ie.mTRP_PUSCH_SP_CSI_r17 != nil || ie.mTRP_PUSCH_CG_r17 != nil || ie.mTRP_PUCCH_MAC_CE_r17 != nil || ie.mTRP_PUCCH_maxNum_PC_FR1_r17 != nil || ie.mTRP_inter_Cell_r17 != nil || ie.mTRP_GroupBasedL1_RSRP_r17 != nil || ie.mTRP_BFD_RS_MAC_CE_r17 != nil || ie.mTRP_CSI_EnhancementPerBand_r17 != nil || ie.codebookComboParameterMultiTRP_r17 != nil || ie.mTRP_CSI_additionalCSI_r17 != nil || ie.mTRP_CSI_N_Max2_r17 != nil || ie.mTRP_CSI_CMR_r17 != nil || ie.srs_partialFreqSounding_r17 != nil || ie.beamSwitchTiming_v1710 != nil || ie.beamSwitchTiming_r17 != nil || ie.beamReportTiming_v1710 != nil || ie.maxNumberRxTxBeamSwitchDL_v1710 != nil, ie.srs_PortReportSP_AP_r17 != nil || ie.maxNumberRxBeam_v1720 != nil || ie.sfn_ImplicitRS_twoTCI_r17 != nil || ie.sfn_QCL_TypeD_Collision_twoTCI_r17 != nil || ie.mTRP_CSI_numCPU_r17 != nil, ie.supportRepNumPDSCH_TDRA_DCI_1_2_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MIMO_ParametersPerBand", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.dummy6 != nil, ie.beamManagementSSB_CSI_RS != nil, ie.beamSwitchTiming != nil, ie.codebookParameters != nil, ie.csi_RS_IM_ReceptionForFeedback != nil, ie.csi_RS_ProcFrameworkForSRS != nil, ie.csi_ReportFramework != nil, ie.csi_RS_ForTracking != nil, len(ie.srs_AssocCSI_RS) > 0, ie.spatialRelations != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode dummy6 optional
			if ie.dummy6 != nil {
				if err = ie.dummy6.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dummy6", err)
				}
			}
			// encode beamManagementSSB_CSI_RS optional
			if ie.beamManagementSSB_CSI_RS != nil {
				if err = ie.beamManagementSSB_CSI_RS.Encode(extWriter); err != nil {
					return utils.WrapError("Encode beamManagementSSB_CSI_RS", err)
				}
			}
			// encode beamSwitchTiming optional
			if ie.beamSwitchTiming != nil {
				if err = ie.beamSwitchTiming.Encode(extWriter); err != nil {
					return utils.WrapError("Encode beamSwitchTiming", err)
				}
			}
			// encode codebookParameters optional
			if ie.codebookParameters != nil {
				if err = ie.codebookParameters.Encode(extWriter); err != nil {
					return utils.WrapError("Encode codebookParameters", err)
				}
			}
			// encode csi_RS_IM_ReceptionForFeedback optional
			if ie.csi_RS_IM_ReceptionForFeedback != nil {
				if err = ie.csi_RS_IM_ReceptionForFeedback.Encode(extWriter); err != nil {
					return utils.WrapError("Encode csi_RS_IM_ReceptionForFeedback", err)
				}
			}
			// encode csi_RS_ProcFrameworkForSRS optional
			if ie.csi_RS_ProcFrameworkForSRS != nil {
				if err = ie.csi_RS_ProcFrameworkForSRS.Encode(extWriter); err != nil {
					return utils.WrapError("Encode csi_RS_ProcFrameworkForSRS", err)
				}
			}
			// encode csi_ReportFramework optional
			if ie.csi_ReportFramework != nil {
				if err = ie.csi_ReportFramework.Encode(extWriter); err != nil {
					return utils.WrapError("Encode csi_ReportFramework", err)
				}
			}
			// encode csi_RS_ForTracking optional
			if ie.csi_RS_ForTracking != nil {
				if err = ie.csi_RS_ForTracking.Encode(extWriter); err != nil {
					return utils.WrapError("Encode csi_RS_ForTracking", err)
				}
			}
			// encode srs_AssocCSI_RS optional
			if len(ie.srs_AssocCSI_RS) > 0 {
				tmp_srs_AssocCSI_RS := utils.NewSequence[*SupportedCSI_RS_Resource]([]*SupportedCSI_RS_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_Resources}, false)
				for _, i := range ie.srs_AssocCSI_RS {
					tmp_srs_AssocCSI_RS.Value = append(tmp_srs_AssocCSI_RS.Value, &i)
				}
				if err = tmp_srs_AssocCSI_RS.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_AssocCSI_RS", err)
				}
			}
			// encode spatialRelations optional
			if ie.spatialRelations != nil {
				if err = ie.spatialRelations.Encode(extWriter); err != nil {
					return utils.WrapError("Encode spatialRelations", err)
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
			optionals_ext_2 := []bool{ie.defaultQCL_TwoTCI_r16 != nil, ie.codebookParametersPerBand_r16 != nil, ie.simul_SpatialRelationUpdatePUCCHResGroup_r16 != nil, ie.maxNumberSCellBFR_r16 != nil, ie.simultaneousReceptionDiffTypeD_r16 != nil, ie.ssb_csirs_SINR_measurement_r16 != nil, ie.nonGroupSINR_reporting_r16 != nil, ie.groupSINR_reporting_r16 != nil, ie.multiDCI_multiTRP_Parameters_r16 != nil, ie.singleDCI_SDM_scheme_Parameters_r16 != nil, ie.supportFDM_SchemeA_r16 != nil, ie.supportCodeWordSoftCombining_r16 != nil, ie.supportTDM_SchemeA_r16 != nil, ie.supportInter_slotTDM_r16 != nil, ie.lowPAPR_DMRS_PDSCH_r16 != nil, ie.lowPAPR_DMRS_PUSCHwithoutPrecoding_r16 != nil, ie.lowPAPR_DMRS_PUCCH_r16 != nil, ie.lowPAPR_DMRS_PUSCHwithPrecoding_r16 != nil, ie.csi_ReportFrameworkExt_r16 != nil, ie.codebookParametersAddition_r16 != nil, ie.codebookComboParametersAddition_r16 != nil, ie.beamCorrespondenceSSB_based_r16 != nil, ie.beamCorrespondenceCSI_RS_based_r16 != nil, ie.beamSwitchTiming_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode defaultQCL_TwoTCI_r16 optional
			if ie.defaultQCL_TwoTCI_r16 != nil {
				if err = ie.defaultQCL_TwoTCI_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode defaultQCL_TwoTCI_r16", err)
				}
			}
			// encode codebookParametersPerBand_r16 optional
			if ie.codebookParametersPerBand_r16 != nil {
				if err = ie.codebookParametersPerBand_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode codebookParametersPerBand_r16", err)
				}
			}
			// encode simul_SpatialRelationUpdatePUCCHResGroup_r16 optional
			if ie.simul_SpatialRelationUpdatePUCCHResGroup_r16 != nil {
				if err = ie.simul_SpatialRelationUpdatePUCCHResGroup_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simul_SpatialRelationUpdatePUCCHResGroup_r16", err)
				}
			}
			// encode maxNumberSCellBFR_r16 optional
			if ie.maxNumberSCellBFR_r16 != nil {
				if err = ie.maxNumberSCellBFR_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxNumberSCellBFR_r16", err)
				}
			}
			// encode simultaneousReceptionDiffTypeD_r16 optional
			if ie.simultaneousReceptionDiffTypeD_r16 != nil {
				if err = ie.simultaneousReceptionDiffTypeD_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simultaneousReceptionDiffTypeD_r16", err)
				}
			}
			// encode ssb_csirs_SINR_measurement_r16 optional
			if ie.ssb_csirs_SINR_measurement_r16 != nil {
				if err = ie.ssb_csirs_SINR_measurement_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ssb_csirs_SINR_measurement_r16", err)
				}
			}
			// encode nonGroupSINR_reporting_r16 optional
			if ie.nonGroupSINR_reporting_r16 != nil {
				if err = ie.nonGroupSINR_reporting_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nonGroupSINR_reporting_r16", err)
				}
			}
			// encode groupSINR_reporting_r16 optional
			if ie.groupSINR_reporting_r16 != nil {
				if err = ie.groupSINR_reporting_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode groupSINR_reporting_r16", err)
				}
			}
			// encode multiDCI_multiTRP_Parameters_r16 optional
			if ie.multiDCI_multiTRP_Parameters_r16 != nil {
				if err = ie.multiDCI_multiTRP_Parameters_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode multiDCI_multiTRP_Parameters_r16", err)
				}
			}
			// encode singleDCI_SDM_scheme_Parameters_r16 optional
			if ie.singleDCI_SDM_scheme_Parameters_r16 != nil {
				if err = ie.singleDCI_SDM_scheme_Parameters_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode singleDCI_SDM_scheme_Parameters_r16", err)
				}
			}
			// encode supportFDM_SchemeA_r16 optional
			if ie.supportFDM_SchemeA_r16 != nil {
				if err = ie.supportFDM_SchemeA_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportFDM_SchemeA_r16", err)
				}
			}
			// encode supportCodeWordSoftCombining_r16 optional
			if ie.supportCodeWordSoftCombining_r16 != nil {
				if err = ie.supportCodeWordSoftCombining_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportCodeWordSoftCombining_r16", err)
				}
			}
			// encode supportTDM_SchemeA_r16 optional
			if ie.supportTDM_SchemeA_r16 != nil {
				if err = ie.supportTDM_SchemeA_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportTDM_SchemeA_r16", err)
				}
			}
			// encode supportInter_slotTDM_r16 optional
			if ie.supportInter_slotTDM_r16 != nil {
				if err = ie.supportInter_slotTDM_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportInter_slotTDM_r16", err)
				}
			}
			// encode lowPAPR_DMRS_PDSCH_r16 optional
			if ie.lowPAPR_DMRS_PDSCH_r16 != nil {
				if err = ie.lowPAPR_DMRS_PDSCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode lowPAPR_DMRS_PDSCH_r16", err)
				}
			}
			// encode lowPAPR_DMRS_PUSCHwithoutPrecoding_r16 optional
			if ie.lowPAPR_DMRS_PUSCHwithoutPrecoding_r16 != nil {
				if err = ie.lowPAPR_DMRS_PUSCHwithoutPrecoding_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode lowPAPR_DMRS_PUSCHwithoutPrecoding_r16", err)
				}
			}
			// encode lowPAPR_DMRS_PUCCH_r16 optional
			if ie.lowPAPR_DMRS_PUCCH_r16 != nil {
				if err = ie.lowPAPR_DMRS_PUCCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode lowPAPR_DMRS_PUCCH_r16", err)
				}
			}
			// encode lowPAPR_DMRS_PUSCHwithPrecoding_r16 optional
			if ie.lowPAPR_DMRS_PUSCHwithPrecoding_r16 != nil {
				if err = ie.lowPAPR_DMRS_PUSCHwithPrecoding_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode lowPAPR_DMRS_PUSCHwithPrecoding_r16", err)
				}
			}
			// encode csi_ReportFrameworkExt_r16 optional
			if ie.csi_ReportFrameworkExt_r16 != nil {
				if err = ie.csi_ReportFrameworkExt_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode csi_ReportFrameworkExt_r16", err)
				}
			}
			// encode codebookParametersAddition_r16 optional
			if ie.codebookParametersAddition_r16 != nil {
				if err = ie.codebookParametersAddition_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode codebookParametersAddition_r16", err)
				}
			}
			// encode codebookComboParametersAddition_r16 optional
			if ie.codebookComboParametersAddition_r16 != nil {
				if err = ie.codebookComboParametersAddition_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode codebookComboParametersAddition_r16", err)
				}
			}
			// encode beamCorrespondenceSSB_based_r16 optional
			if ie.beamCorrespondenceSSB_based_r16 != nil {
				if err = ie.beamCorrespondenceSSB_based_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode beamCorrespondenceSSB_based_r16", err)
				}
			}
			// encode beamCorrespondenceCSI_RS_based_r16 optional
			if ie.beamCorrespondenceCSI_RS_based_r16 != nil {
				if err = ie.beamCorrespondenceCSI_RS_based_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode beamCorrespondenceCSI_RS_based_r16", err)
				}
			}
			// encode beamSwitchTiming_r16 optional
			if ie.beamSwitchTiming_r16 != nil {
				if err = ie.beamSwitchTiming_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode beamSwitchTiming_r16", err)
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
			optionals_ext_3 := []bool{ie.semi_PersistentL1_SINR_Report_PUCCH_r16 != nil, ie.semi_PersistentL1_SINR_Report_PUSCH_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode semi_PersistentL1_SINR_Report_PUCCH_r16 optional
			if ie.semi_PersistentL1_SINR_Report_PUCCH_r16 != nil {
				if err = ie.semi_PersistentL1_SINR_Report_PUCCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode semi_PersistentL1_SINR_Report_PUCCH_r16", err)
				}
			}
			// encode semi_PersistentL1_SINR_Report_PUSCH_r16 optional
			if ie.semi_PersistentL1_SINR_Report_PUSCH_r16 != nil {
				if err = ie.semi_PersistentL1_SINR_Report_PUSCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode semi_PersistentL1_SINR_Report_PUSCH_r16", err)
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
			optionals_ext_4 := []bool{ie.spatialRelations_v1640 != nil, ie.support64CandidateBeamRS_BFR_r16 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode spatialRelations_v1640 optional
			if ie.spatialRelations_v1640 != nil {
				if err = ie.spatialRelations_v1640.Encode(extWriter); err != nil {
					return utils.WrapError("Encode spatialRelations_v1640", err)
				}
			}
			// encode support64CandidateBeamRS_BFR_r16 optional
			if ie.support64CandidateBeamRS_BFR_r16 != nil {
				if err = ie.support64CandidateBeamRS_BFR_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode support64CandidateBeamRS_BFR_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 5
		if extBitmap[4] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 5
			optionals_ext_5 := []bool{ie.maxMIMO_LayersForMulti_DCI_mTRP_r16 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode maxMIMO_LayersForMulti_DCI_mTRP_r16 optional
			if ie.maxMIMO_LayersForMulti_DCI_mTRP_r16 != nil {
				if err = ie.maxMIMO_LayersForMulti_DCI_mTRP_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxMIMO_LayersForMulti_DCI_mTRP_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 6
		if extBitmap[5] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 6
			optionals_ext_6 := []bool{ie.supportedSINR_meas_v1670 != nil}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode supportedSINR_meas_v1670 optional
			if ie.supportedSINR_meas_v1670 != nil {
				if err = extWriter.WriteBitString(ie.supportedSINR_meas_v1670.Bytes, uint(ie.supportedSINR_meas_v1670.NumBits), &uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
					return utils.WrapError("Encode supportedSINR_meas_v1670", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 7
		if extBitmap[6] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 7
			optionals_ext_7 := []bool{ie.srs_increasedRepetition_r17 != nil, ie.srs_partialFrequencySounding_r17 != nil, ie.srs_startRB_locationHoppingPartial_r17 != nil, ie.srs_combEight_r17 != nil, ie.codebookParametersfetype2_r17 != nil, ie.mTRP_PUSCH_twoCSI_RS_r17 != nil, ie.mTRP_PUCCH_InterSlot_r17 != nil, ie.mTRP_PUCCH_CyclicMapping_r17 != nil, ie.mTRP_PUCCH_SecondTPC_r17 != nil, ie.mTRP_BFR_twoBFD_RS_Set_r17 != nil, ie.mTRP_BFR_PUCCH_SR_perCG_r17 != nil, ie.mTRP_BFR_association_PUCCH_SR_r17 != nil, ie.sfn_SimulTwoTCI_AcrossMultiCC_r17 != nil, ie.sfn_DefaultDL_BeamSetup_r17 != nil, ie.sfn_DefaultUL_BeamSetup_r17 != nil, ie.srs_TriggeringOffset_r17 != nil, ie.srs_TriggeringDCI_r17 != nil, ie.codebookComboParameterMixedType_r17 != nil, ie.unifiedJointTCI_r17 != nil, ie.unifiedJointTCI_multiMAC_CE_r17 != nil, ie.unifiedJointTCI_perBWP_CA_r17 != nil, ie.unifiedJointTCI_ListSharingCA_r17 != nil, ie.unifiedJointTCI_commonMultiCC_r17 != nil, ie.unifiedJointTCI_BeamAlignDLRS_r17 != nil, ie.unifiedJointTCI_PC_association_r17 != nil, ie.unifiedJointTCI_Legacy_r17 != nil, ie.unifiedJointTCI_Legacy_SRS_r17 != nil, ie.unifiedJointTCI_Legacy_CORESET0_r17 != nil, ie.unifiedJointTCI_SCellBFR_r17 != nil, ie.unifiedJointTCI_InterCell_r17 != nil, ie.unifiedSeparateTCI_r17 != nil, ie.unifiedSeparateTCI_multiMAC_CE_r17 != nil, ie.unifiedSeparateTCI_perBWP_CA_r17 != nil, ie.unifiedSeparateTCI_ListSharingCA_r17 != nil, ie.unifiedSeparateTCI_commonMultiCC_r17 != nil, ie.unifiedSeparateTCI_InterCell_r17 != nil, ie.unifiedJointTCI_mTRP_InterCell_BM_r17 != nil, ie.mpe_Mitigation_r17 != nil, ie.srs_PortReport_r17 != nil, ie.mTRP_PDCCH_individual_r17 != nil, ie.mTRP_PDCCH_anySpan_3Symbols_r17 != nil, ie.mTRP_PDCCH_TwoQCL_TypeD_r17 != nil, ie.mTRP_PUSCH_CSI_RS_r17 != nil, ie.mTRP_PUSCH_cyclicMapping_r17 != nil, ie.mTRP_PUSCH_secondTPC_r17 != nil, ie.mTRP_PUSCH_twoPHR_Reporting_r17 != nil, ie.mTRP_PUSCH_A_CSI_r17 != nil, ie.mTRP_PUSCH_SP_CSI_r17 != nil, ie.mTRP_PUSCH_CG_r17 != nil, ie.mTRP_PUCCH_MAC_CE_r17 != nil, ie.mTRP_PUCCH_maxNum_PC_FR1_r17 != nil, ie.mTRP_inter_Cell_r17 != nil, ie.mTRP_GroupBasedL1_RSRP_r17 != nil, ie.mTRP_BFD_RS_MAC_CE_r17 != nil, ie.mTRP_CSI_EnhancementPerBand_r17 != nil, ie.codebookComboParameterMultiTRP_r17 != nil, ie.mTRP_CSI_additionalCSI_r17 != nil, ie.mTRP_CSI_N_Max2_r17 != nil, ie.mTRP_CSI_CMR_r17 != nil, ie.srs_partialFreqSounding_r17 != nil, ie.beamSwitchTiming_v1710 != nil, ie.beamSwitchTiming_r17 != nil, ie.beamReportTiming_v1710 != nil, ie.maxNumberRxTxBeamSwitchDL_v1710 != nil}
			for _, bit := range optionals_ext_7 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode srs_increasedRepetition_r17 optional
			if ie.srs_increasedRepetition_r17 != nil {
				if err = ie.srs_increasedRepetition_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_increasedRepetition_r17", err)
				}
			}
			// encode srs_partialFrequencySounding_r17 optional
			if ie.srs_partialFrequencySounding_r17 != nil {
				if err = ie.srs_partialFrequencySounding_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_partialFrequencySounding_r17", err)
				}
			}
			// encode srs_startRB_locationHoppingPartial_r17 optional
			if ie.srs_startRB_locationHoppingPartial_r17 != nil {
				if err = ie.srs_startRB_locationHoppingPartial_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_startRB_locationHoppingPartial_r17", err)
				}
			}
			// encode srs_combEight_r17 optional
			if ie.srs_combEight_r17 != nil {
				if err = ie.srs_combEight_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_combEight_r17", err)
				}
			}
			// encode codebookParametersfetype2_r17 optional
			if ie.codebookParametersfetype2_r17 != nil {
				if err = ie.codebookParametersfetype2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode codebookParametersfetype2_r17", err)
				}
			}
			// encode mTRP_PUSCH_twoCSI_RS_r17 optional
			if ie.mTRP_PUSCH_twoCSI_RS_r17 != nil {
				if err = ie.mTRP_PUSCH_twoCSI_RS_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_PUSCH_twoCSI_RS_r17", err)
				}
			}
			// encode mTRP_PUCCH_InterSlot_r17 optional
			if ie.mTRP_PUCCH_InterSlot_r17 != nil {
				if err = ie.mTRP_PUCCH_InterSlot_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_PUCCH_InterSlot_r17", err)
				}
			}
			// encode mTRP_PUCCH_CyclicMapping_r17 optional
			if ie.mTRP_PUCCH_CyclicMapping_r17 != nil {
				if err = ie.mTRP_PUCCH_CyclicMapping_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_PUCCH_CyclicMapping_r17", err)
				}
			}
			// encode mTRP_PUCCH_SecondTPC_r17 optional
			if ie.mTRP_PUCCH_SecondTPC_r17 != nil {
				if err = ie.mTRP_PUCCH_SecondTPC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_PUCCH_SecondTPC_r17", err)
				}
			}
			// encode mTRP_BFR_twoBFD_RS_Set_r17 optional
			if ie.mTRP_BFR_twoBFD_RS_Set_r17 != nil {
				if err = ie.mTRP_BFR_twoBFD_RS_Set_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_BFR_twoBFD_RS_Set_r17", err)
				}
			}
			// encode mTRP_BFR_PUCCH_SR_perCG_r17 optional
			if ie.mTRP_BFR_PUCCH_SR_perCG_r17 != nil {
				if err = ie.mTRP_BFR_PUCCH_SR_perCG_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_BFR_PUCCH_SR_perCG_r17", err)
				}
			}
			// encode mTRP_BFR_association_PUCCH_SR_r17 optional
			if ie.mTRP_BFR_association_PUCCH_SR_r17 != nil {
				if err = ie.mTRP_BFR_association_PUCCH_SR_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_BFR_association_PUCCH_SR_r17", err)
				}
			}
			// encode sfn_SimulTwoTCI_AcrossMultiCC_r17 optional
			if ie.sfn_SimulTwoTCI_AcrossMultiCC_r17 != nil {
				if err = ie.sfn_SimulTwoTCI_AcrossMultiCC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sfn_SimulTwoTCI_AcrossMultiCC_r17", err)
				}
			}
			// encode sfn_DefaultDL_BeamSetup_r17 optional
			if ie.sfn_DefaultDL_BeamSetup_r17 != nil {
				if err = ie.sfn_DefaultDL_BeamSetup_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sfn_DefaultDL_BeamSetup_r17", err)
				}
			}
			// encode sfn_DefaultUL_BeamSetup_r17 optional
			if ie.sfn_DefaultUL_BeamSetup_r17 != nil {
				if err = ie.sfn_DefaultUL_BeamSetup_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sfn_DefaultUL_BeamSetup_r17", err)
				}
			}
			// encode srs_TriggeringOffset_r17 optional
			if ie.srs_TriggeringOffset_r17 != nil {
				if err = ie.srs_TriggeringOffset_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_TriggeringOffset_r17", err)
				}
			}
			// encode srs_TriggeringDCI_r17 optional
			if ie.srs_TriggeringDCI_r17 != nil {
				if err = ie.srs_TriggeringDCI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_TriggeringDCI_r17", err)
				}
			}
			// encode codebookComboParameterMixedType_r17 optional
			if ie.codebookComboParameterMixedType_r17 != nil {
				if err = ie.codebookComboParameterMixedType_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode codebookComboParameterMixedType_r17", err)
				}
			}
			// encode unifiedJointTCI_r17 optional
			if ie.unifiedJointTCI_r17 != nil {
				if err = ie.unifiedJointTCI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode unifiedJointTCI_r17", err)
				}
			}
			// encode unifiedJointTCI_multiMAC_CE_r17 optional
			if ie.unifiedJointTCI_multiMAC_CE_r17 != nil {
				if err = ie.unifiedJointTCI_multiMAC_CE_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode unifiedJointTCI_multiMAC_CE_r17", err)
				}
			}
			// encode unifiedJointTCI_perBWP_CA_r17 optional
			if ie.unifiedJointTCI_perBWP_CA_r17 != nil {
				if err = ie.unifiedJointTCI_perBWP_CA_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode unifiedJointTCI_perBWP_CA_r17", err)
				}
			}
			// encode unifiedJointTCI_ListSharingCA_r17 optional
			if ie.unifiedJointTCI_ListSharingCA_r17 != nil {
				if err = ie.unifiedJointTCI_ListSharingCA_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode unifiedJointTCI_ListSharingCA_r17", err)
				}
			}
			// encode unifiedJointTCI_commonMultiCC_r17 optional
			if ie.unifiedJointTCI_commonMultiCC_r17 != nil {
				if err = ie.unifiedJointTCI_commonMultiCC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode unifiedJointTCI_commonMultiCC_r17", err)
				}
			}
			// encode unifiedJointTCI_BeamAlignDLRS_r17 optional
			if ie.unifiedJointTCI_BeamAlignDLRS_r17 != nil {
				if err = ie.unifiedJointTCI_BeamAlignDLRS_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode unifiedJointTCI_BeamAlignDLRS_r17", err)
				}
			}
			// encode unifiedJointTCI_PC_association_r17 optional
			if ie.unifiedJointTCI_PC_association_r17 != nil {
				if err = ie.unifiedJointTCI_PC_association_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode unifiedJointTCI_PC_association_r17", err)
				}
			}
			// encode unifiedJointTCI_Legacy_r17 optional
			if ie.unifiedJointTCI_Legacy_r17 != nil {
				if err = ie.unifiedJointTCI_Legacy_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode unifiedJointTCI_Legacy_r17", err)
				}
			}
			// encode unifiedJointTCI_Legacy_SRS_r17 optional
			if ie.unifiedJointTCI_Legacy_SRS_r17 != nil {
				if err = ie.unifiedJointTCI_Legacy_SRS_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode unifiedJointTCI_Legacy_SRS_r17", err)
				}
			}
			// encode unifiedJointTCI_Legacy_CORESET0_r17 optional
			if ie.unifiedJointTCI_Legacy_CORESET0_r17 != nil {
				if err = ie.unifiedJointTCI_Legacy_CORESET0_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode unifiedJointTCI_Legacy_CORESET0_r17", err)
				}
			}
			// encode unifiedJointTCI_SCellBFR_r17 optional
			if ie.unifiedJointTCI_SCellBFR_r17 != nil {
				if err = ie.unifiedJointTCI_SCellBFR_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode unifiedJointTCI_SCellBFR_r17", err)
				}
			}
			// encode unifiedJointTCI_InterCell_r17 optional
			if ie.unifiedJointTCI_InterCell_r17 != nil {
				if err = ie.unifiedJointTCI_InterCell_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode unifiedJointTCI_InterCell_r17", err)
				}
			}
			// encode unifiedSeparateTCI_r17 optional
			if ie.unifiedSeparateTCI_r17 != nil {
				if err = ie.unifiedSeparateTCI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode unifiedSeparateTCI_r17", err)
				}
			}
			// encode unifiedSeparateTCI_multiMAC_CE_r17 optional
			if ie.unifiedSeparateTCI_multiMAC_CE_r17 != nil {
				if err = ie.unifiedSeparateTCI_multiMAC_CE_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode unifiedSeparateTCI_multiMAC_CE_r17", err)
				}
			}
			// encode unifiedSeparateTCI_perBWP_CA_r17 optional
			if ie.unifiedSeparateTCI_perBWP_CA_r17 != nil {
				if err = ie.unifiedSeparateTCI_perBWP_CA_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode unifiedSeparateTCI_perBWP_CA_r17", err)
				}
			}
			// encode unifiedSeparateTCI_ListSharingCA_r17 optional
			if ie.unifiedSeparateTCI_ListSharingCA_r17 != nil {
				if err = ie.unifiedSeparateTCI_ListSharingCA_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode unifiedSeparateTCI_ListSharingCA_r17", err)
				}
			}
			// encode unifiedSeparateTCI_commonMultiCC_r17 optional
			if ie.unifiedSeparateTCI_commonMultiCC_r17 != nil {
				if err = ie.unifiedSeparateTCI_commonMultiCC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode unifiedSeparateTCI_commonMultiCC_r17", err)
				}
			}
			// encode unifiedSeparateTCI_InterCell_r17 optional
			if ie.unifiedSeparateTCI_InterCell_r17 != nil {
				if err = ie.unifiedSeparateTCI_InterCell_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode unifiedSeparateTCI_InterCell_r17", err)
				}
			}
			// encode unifiedJointTCI_mTRP_InterCell_BM_r17 optional
			if ie.unifiedJointTCI_mTRP_InterCell_BM_r17 != nil {
				if err = ie.unifiedJointTCI_mTRP_InterCell_BM_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode unifiedJointTCI_mTRP_InterCell_BM_r17", err)
				}
			}
			// encode mpe_Mitigation_r17 optional
			if ie.mpe_Mitigation_r17 != nil {
				if err = ie.mpe_Mitigation_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mpe_Mitigation_r17", err)
				}
			}
			// encode srs_PortReport_r17 optional
			if ie.srs_PortReport_r17 != nil {
				if err = ie.srs_PortReport_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_PortReport_r17", err)
				}
			}
			// encode mTRP_PDCCH_individual_r17 optional
			if ie.mTRP_PDCCH_individual_r17 != nil {
				if err = ie.mTRP_PDCCH_individual_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_PDCCH_individual_r17", err)
				}
			}
			// encode mTRP_PDCCH_anySpan_3Symbols_r17 optional
			if ie.mTRP_PDCCH_anySpan_3Symbols_r17 != nil {
				if err = ie.mTRP_PDCCH_anySpan_3Symbols_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_PDCCH_anySpan_3Symbols_r17", err)
				}
			}
			// encode mTRP_PDCCH_TwoQCL_TypeD_r17 optional
			if ie.mTRP_PDCCH_TwoQCL_TypeD_r17 != nil {
				if err = ie.mTRP_PDCCH_TwoQCL_TypeD_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_PDCCH_TwoQCL_TypeD_r17", err)
				}
			}
			// encode mTRP_PUSCH_CSI_RS_r17 optional
			if ie.mTRP_PUSCH_CSI_RS_r17 != nil {
				if err = ie.mTRP_PUSCH_CSI_RS_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_PUSCH_CSI_RS_r17", err)
				}
			}
			// encode mTRP_PUSCH_cyclicMapping_r17 optional
			if ie.mTRP_PUSCH_cyclicMapping_r17 != nil {
				if err = ie.mTRP_PUSCH_cyclicMapping_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_PUSCH_cyclicMapping_r17", err)
				}
			}
			// encode mTRP_PUSCH_secondTPC_r17 optional
			if ie.mTRP_PUSCH_secondTPC_r17 != nil {
				if err = ie.mTRP_PUSCH_secondTPC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_PUSCH_secondTPC_r17", err)
				}
			}
			// encode mTRP_PUSCH_twoPHR_Reporting_r17 optional
			if ie.mTRP_PUSCH_twoPHR_Reporting_r17 != nil {
				if err = ie.mTRP_PUSCH_twoPHR_Reporting_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_PUSCH_twoPHR_Reporting_r17", err)
				}
			}
			// encode mTRP_PUSCH_A_CSI_r17 optional
			if ie.mTRP_PUSCH_A_CSI_r17 != nil {
				if err = ie.mTRP_PUSCH_A_CSI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_PUSCH_A_CSI_r17", err)
				}
			}
			// encode mTRP_PUSCH_SP_CSI_r17 optional
			if ie.mTRP_PUSCH_SP_CSI_r17 != nil {
				if err = ie.mTRP_PUSCH_SP_CSI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_PUSCH_SP_CSI_r17", err)
				}
			}
			// encode mTRP_PUSCH_CG_r17 optional
			if ie.mTRP_PUSCH_CG_r17 != nil {
				if err = ie.mTRP_PUSCH_CG_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_PUSCH_CG_r17", err)
				}
			}
			// encode mTRP_PUCCH_MAC_CE_r17 optional
			if ie.mTRP_PUCCH_MAC_CE_r17 != nil {
				if err = ie.mTRP_PUCCH_MAC_CE_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_PUCCH_MAC_CE_r17", err)
				}
			}
			// encode mTRP_PUCCH_maxNum_PC_FR1_r17 optional
			if ie.mTRP_PUCCH_maxNum_PC_FR1_r17 != nil {
				if err = extWriter.WriteInteger(*ie.mTRP_PUCCH_maxNum_PC_FR1_r17, &uper.Constraint{Lb: 3, Ub: 8}, false); err != nil {
					return utils.WrapError("Encode mTRP_PUCCH_maxNum_PC_FR1_r17", err)
				}
			}
			// encode mTRP_inter_Cell_r17 optional
			if ie.mTRP_inter_Cell_r17 != nil {
				if err = ie.mTRP_inter_Cell_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_inter_Cell_r17", err)
				}
			}
			// encode mTRP_GroupBasedL1_RSRP_r17 optional
			if ie.mTRP_GroupBasedL1_RSRP_r17 != nil {
				if err = ie.mTRP_GroupBasedL1_RSRP_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_GroupBasedL1_RSRP_r17", err)
				}
			}
			// encode mTRP_BFD_RS_MAC_CE_r17 optional
			if ie.mTRP_BFD_RS_MAC_CE_r17 != nil {
				if err = ie.mTRP_BFD_RS_MAC_CE_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_BFD_RS_MAC_CE_r17", err)
				}
			}
			// encode mTRP_CSI_EnhancementPerBand_r17 optional
			if ie.mTRP_CSI_EnhancementPerBand_r17 != nil {
				if err = ie.mTRP_CSI_EnhancementPerBand_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_CSI_EnhancementPerBand_r17", err)
				}
			}
			// encode codebookComboParameterMultiTRP_r17 optional
			if ie.codebookComboParameterMultiTRP_r17 != nil {
				if err = ie.codebookComboParameterMultiTRP_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode codebookComboParameterMultiTRP_r17", err)
				}
			}
			// encode mTRP_CSI_additionalCSI_r17 optional
			if ie.mTRP_CSI_additionalCSI_r17 != nil {
				if err = ie.mTRP_CSI_additionalCSI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_CSI_additionalCSI_r17", err)
				}
			}
			// encode mTRP_CSI_N_Max2_r17 optional
			if ie.mTRP_CSI_N_Max2_r17 != nil {
				if err = ie.mTRP_CSI_N_Max2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_CSI_N_Max2_r17", err)
				}
			}
			// encode mTRP_CSI_CMR_r17 optional
			if ie.mTRP_CSI_CMR_r17 != nil {
				if err = ie.mTRP_CSI_CMR_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_CSI_CMR_r17", err)
				}
			}
			// encode srs_partialFreqSounding_r17 optional
			if ie.srs_partialFreqSounding_r17 != nil {
				if err = ie.srs_partialFreqSounding_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_partialFreqSounding_r17", err)
				}
			}
			// encode beamSwitchTiming_v1710 optional
			if ie.beamSwitchTiming_v1710 != nil {
				if err = ie.beamSwitchTiming_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode beamSwitchTiming_v1710", err)
				}
			}
			// encode beamSwitchTiming_r17 optional
			if ie.beamSwitchTiming_r17 != nil {
				if err = ie.beamSwitchTiming_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode beamSwitchTiming_r17", err)
				}
			}
			// encode beamReportTiming_v1710 optional
			if ie.beamReportTiming_v1710 != nil {
				if err = ie.beamReportTiming_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode beamReportTiming_v1710", err)
				}
			}
			// encode maxNumberRxTxBeamSwitchDL_v1710 optional
			if ie.maxNumberRxTxBeamSwitchDL_v1710 != nil {
				if err = ie.maxNumberRxTxBeamSwitchDL_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxNumberRxTxBeamSwitchDL_v1710", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 8
		if extBitmap[7] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 8
			optionals_ext_8 := []bool{ie.srs_PortReportSP_AP_r17 != nil, ie.maxNumberRxBeam_v1720 != nil, ie.sfn_ImplicitRS_twoTCI_r17 != nil, ie.sfn_QCL_TypeD_Collision_twoTCI_r17 != nil, ie.mTRP_CSI_numCPU_r17 != nil}
			for _, bit := range optionals_ext_8 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode srs_PortReportSP_AP_r17 optional
			if ie.srs_PortReportSP_AP_r17 != nil {
				if err = ie.srs_PortReportSP_AP_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_PortReportSP_AP_r17", err)
				}
			}
			// encode maxNumberRxBeam_v1720 optional
			if ie.maxNumberRxBeam_v1720 != nil {
				if err = extWriter.WriteInteger(*ie.maxNumberRxBeam_v1720, &uper.Constraint{Lb: 9, Ub: 12}, false); err != nil {
					return utils.WrapError("Encode maxNumberRxBeam_v1720", err)
				}
			}
			// encode sfn_ImplicitRS_twoTCI_r17 optional
			if ie.sfn_ImplicitRS_twoTCI_r17 != nil {
				if err = ie.sfn_ImplicitRS_twoTCI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sfn_ImplicitRS_twoTCI_r17", err)
				}
			}
			// encode sfn_QCL_TypeD_Collision_twoTCI_r17 optional
			if ie.sfn_QCL_TypeD_Collision_twoTCI_r17 != nil {
				if err = ie.sfn_QCL_TypeD_Collision_twoTCI_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sfn_QCL_TypeD_Collision_twoTCI_r17", err)
				}
			}
			// encode mTRP_CSI_numCPU_r17 optional
			if ie.mTRP_CSI_numCPU_r17 != nil {
				if err = ie.mTRP_CSI_numCPU_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_CSI_numCPU_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 9
		if extBitmap[8] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 9
			optionals_ext_9 := []bool{ie.supportRepNumPDSCH_TDRA_DCI_1_2_r17 != nil}
			for _, bit := range optionals_ext_9 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode supportRepNumPDSCH_TDRA_DCI_1_2_r17 optional
			if ie.supportRepNumPDSCH_TDRA_DCI_1_2_r17 != nil {
				if err = ie.supportRepNumPDSCH_TDRA_DCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportRepNumPDSCH_TDRA_DCI_1_2_r17", err)
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

func (ie *MIMO_ParametersPerBand) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var tci_StatePDSCHPresent bool
	if tci_StatePDSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var additionalActiveTCI_StatePDCCHPresent bool
	if additionalActiveTCI_StatePDCCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_TransCoherencePresent bool
	if pusch_TransCoherencePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var beamCorrespondenceWithoutUL_BeamSweepingPresent bool
	if beamCorrespondenceWithoutUL_BeamSweepingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var periodicBeamReportPresent bool
	if periodicBeamReportPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var aperiodicBeamReportPresent bool
	if aperiodicBeamReportPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sp_BeamReportPUCCHPresent bool
	if sp_BeamReportPUCCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sp_BeamReportPUSCHPresent bool
	if sp_BeamReportPUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy1Present bool
	if dummy1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumberRxBeamPresent bool
	if maxNumberRxBeamPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumberRxTxBeamSwitchDLPresent bool
	if maxNumberRxTxBeamSwitchDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumberNonGroupBeamReportingPresent bool
	if maxNumberNonGroupBeamReportingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var groupBeamReportingPresent bool
	if groupBeamReportingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var uplinkBeamManagementPresent bool
	if uplinkBeamManagementPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumberCSI_RS_BFDPresent bool
	if maxNumberCSI_RS_BFDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumberSSB_BFDPresent bool
	if maxNumberSSB_BFDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var maxNumberCSI_RS_SSB_CBDPresent bool
	if maxNumberCSI_RS_SSB_CBDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy2Present bool
	if dummy2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPortsPTRS_ULPresent bool
	if twoPortsPTRS_ULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy5Present bool
	if dummy5Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy3Present bool
	if dummy3Present, err = r.ReadBool(); err != nil {
		return err
	}
	var beamReportTimingPresent bool
	if beamReportTimingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ptrs_DensityRecommendationSetDLPresent bool
	if ptrs_DensityRecommendationSetDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ptrs_DensityRecommendationSetULPresent bool
	if ptrs_DensityRecommendationSetULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy4Present bool
	if dummy4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var aperiodicTRSPresent bool
	if aperiodicTRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if tci_StatePDSCHPresent {
		ie.tci_StatePDSCH = new(MIMO_ParametersPerBand_tci_StatePDSCH)
		if err = ie.tci_StatePDSCH.Decode(r); err != nil {
			return utils.WrapError("Decode tci_StatePDSCH", err)
		}
	}
	if additionalActiveTCI_StatePDCCHPresent {
		ie.additionalActiveTCI_StatePDCCH = new(MIMO_ParametersPerBand_additionalActiveTCI_StatePDCCH)
		if err = ie.additionalActiveTCI_StatePDCCH.Decode(r); err != nil {
			return utils.WrapError("Decode additionalActiveTCI_StatePDCCH", err)
		}
	}
	if pusch_TransCoherencePresent {
		ie.pusch_TransCoherence = new(MIMO_ParametersPerBand_pusch_TransCoherence)
		if err = ie.pusch_TransCoherence.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_TransCoherence", err)
		}
	}
	if beamCorrespondenceWithoutUL_BeamSweepingPresent {
		ie.beamCorrespondenceWithoutUL_BeamSweeping = new(MIMO_ParametersPerBand_beamCorrespondenceWithoutUL_BeamSweeping)
		if err = ie.beamCorrespondenceWithoutUL_BeamSweeping.Decode(r); err != nil {
			return utils.WrapError("Decode beamCorrespondenceWithoutUL_BeamSweeping", err)
		}
	}
	if periodicBeamReportPresent {
		ie.periodicBeamReport = new(MIMO_ParametersPerBand_periodicBeamReport)
		if err = ie.periodicBeamReport.Decode(r); err != nil {
			return utils.WrapError("Decode periodicBeamReport", err)
		}
	}
	if aperiodicBeamReportPresent {
		ie.aperiodicBeamReport = new(MIMO_ParametersPerBand_aperiodicBeamReport)
		if err = ie.aperiodicBeamReport.Decode(r); err != nil {
			return utils.WrapError("Decode aperiodicBeamReport", err)
		}
	}
	if sp_BeamReportPUCCHPresent {
		ie.sp_BeamReportPUCCH = new(MIMO_ParametersPerBand_sp_BeamReportPUCCH)
		if err = ie.sp_BeamReportPUCCH.Decode(r); err != nil {
			return utils.WrapError("Decode sp_BeamReportPUCCH", err)
		}
	}
	if sp_BeamReportPUSCHPresent {
		ie.sp_BeamReportPUSCH = new(MIMO_ParametersPerBand_sp_BeamReportPUSCH)
		if err = ie.sp_BeamReportPUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode sp_BeamReportPUSCH", err)
		}
	}
	if dummy1Present {
		ie.dummy1 = new(DummyG)
		if err = ie.dummy1.Decode(r); err != nil {
			return utils.WrapError("Decode dummy1", err)
		}
	}
	if maxNumberRxBeamPresent {
		var tmp_int_maxNumberRxBeam int64
		if tmp_int_maxNumberRxBeam, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode maxNumberRxBeam", err)
		}
		ie.maxNumberRxBeam = &tmp_int_maxNumberRxBeam
	}
	if maxNumberRxTxBeamSwitchDLPresent {
		ie.maxNumberRxTxBeamSwitchDL = new(MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL)
		if err = ie.maxNumberRxTxBeamSwitchDL.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumberRxTxBeamSwitchDL", err)
		}
	}
	if maxNumberNonGroupBeamReportingPresent {
		ie.maxNumberNonGroupBeamReporting = new(MIMO_ParametersPerBand_maxNumberNonGroupBeamReporting)
		if err = ie.maxNumberNonGroupBeamReporting.Decode(r); err != nil {
			return utils.WrapError("Decode maxNumberNonGroupBeamReporting", err)
		}
	}
	if groupBeamReportingPresent {
		ie.groupBeamReporting = new(MIMO_ParametersPerBand_groupBeamReporting)
		if err = ie.groupBeamReporting.Decode(r); err != nil {
			return utils.WrapError("Decode groupBeamReporting", err)
		}
	}
	if uplinkBeamManagementPresent {
		ie.uplinkBeamManagement = new(MIMO_ParametersPerBand_uplinkBeamManagement)
		if err = ie.uplinkBeamManagement.Decode(r); err != nil {
			return utils.WrapError("Decode uplinkBeamManagement", err)
		}
	}
	if maxNumberCSI_RS_BFDPresent {
		var tmp_int_maxNumberCSI_RS_BFD int64
		if tmp_int_maxNumberCSI_RS_BFD, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode maxNumberCSI_RS_BFD", err)
		}
		ie.maxNumberCSI_RS_BFD = &tmp_int_maxNumberCSI_RS_BFD
	}
	if maxNumberSSB_BFDPresent {
		var tmp_int_maxNumberSSB_BFD int64
		if tmp_int_maxNumberSSB_BFD, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode maxNumberSSB_BFD", err)
		}
		ie.maxNumberSSB_BFD = &tmp_int_maxNumberSSB_BFD
	}
	if maxNumberCSI_RS_SSB_CBDPresent {
		var tmp_int_maxNumberCSI_RS_SSB_CBD int64
		if tmp_int_maxNumberCSI_RS_SSB_CBD, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
			return utils.WrapError("Decode maxNumberCSI_RS_SSB_CBD", err)
		}
		ie.maxNumberCSI_RS_SSB_CBD = &tmp_int_maxNumberCSI_RS_SSB_CBD
	}
	if dummy2Present {
		ie.dummy2 = new(MIMO_ParametersPerBand_dummy2)
		if err = ie.dummy2.Decode(r); err != nil {
			return utils.WrapError("Decode dummy2", err)
		}
	}
	if twoPortsPTRS_ULPresent {
		ie.twoPortsPTRS_UL = new(MIMO_ParametersPerBand_twoPortsPTRS_UL)
		if err = ie.twoPortsPTRS_UL.Decode(r); err != nil {
			return utils.WrapError("Decode twoPortsPTRS_UL", err)
		}
	}
	if dummy5Present {
		ie.dummy5 = new(SRS_Resources)
		if err = ie.dummy5.Decode(r); err != nil {
			return utils.WrapError("Decode dummy5", err)
		}
	}
	if dummy3Present {
		var tmp_int_dummy3 int64
		if tmp_int_dummy3, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode dummy3", err)
		}
		ie.dummy3 = &tmp_int_dummy3
	}
	if beamReportTimingPresent {
		ie.beamReportTiming = new(MIMO_ParametersPerBand_beamReportTiming)
		if err = ie.beamReportTiming.Decode(r); err != nil {
			return utils.WrapError("Decode beamReportTiming", err)
		}
	}
	if ptrs_DensityRecommendationSetDLPresent {
		ie.ptrs_DensityRecommendationSetDL = new(MIMO_ParametersPerBand_ptrs_DensityRecommendationSetDL)
		if err = ie.ptrs_DensityRecommendationSetDL.Decode(r); err != nil {
			return utils.WrapError("Decode ptrs_DensityRecommendationSetDL", err)
		}
	}
	if ptrs_DensityRecommendationSetULPresent {
		ie.ptrs_DensityRecommendationSetUL = new(MIMO_ParametersPerBand_ptrs_DensityRecommendationSetUL)
		if err = ie.ptrs_DensityRecommendationSetUL.Decode(r); err != nil {
			return utils.WrapError("Decode ptrs_DensityRecommendationSetUL", err)
		}
	}
	if dummy4Present {
		ie.dummy4 = new(DummyH)
		if err = ie.dummy4.Decode(r); err != nil {
			return utils.WrapError("Decode dummy4", err)
		}
	}
	if aperiodicTRSPresent {
		ie.aperiodicTRS = new(MIMO_ParametersPerBand_aperiodicTRS)
		if err = ie.aperiodicTRS.Decode(r); err != nil {
			return utils.WrapError("Decode aperiodicTRS", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 9 bits for 9 extension groups
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

			dummy6Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			beamManagementSSB_CSI_RSPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			beamSwitchTimingPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			codebookParametersPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			csi_RS_IM_ReceptionForFeedbackPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			csi_RS_ProcFrameworkForSRSPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			csi_ReportFrameworkPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			csi_RS_ForTrackingPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_AssocCSI_RSPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			spatialRelationsPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode dummy6 optional
			if dummy6Present {
				ie.dummy6 = new(MIMO_ParametersPerBand_dummy6)
				if err = ie.dummy6.Decode(extReader); err != nil {
					return utils.WrapError("Decode dummy6", err)
				}
			}
			// decode beamManagementSSB_CSI_RS optional
			if beamManagementSSB_CSI_RSPresent {
				ie.beamManagementSSB_CSI_RS = new(BeamManagementSSB_CSI_RS)
				if err = ie.beamManagementSSB_CSI_RS.Decode(extReader); err != nil {
					return utils.WrapError("Decode beamManagementSSB_CSI_RS", err)
				}
			}
			// decode beamSwitchTiming optional
			if beamSwitchTimingPresent {
				ie.beamSwitchTiming = new(MIMO_ParametersPerBand_beamSwitchTiming)
				if err = ie.beamSwitchTiming.Decode(extReader); err != nil {
					return utils.WrapError("Decode beamSwitchTiming", err)
				}
			}
			// decode codebookParameters optional
			if codebookParametersPresent {
				ie.codebookParameters = new(CodebookParameters)
				if err = ie.codebookParameters.Decode(extReader); err != nil {
					return utils.WrapError("Decode codebookParameters", err)
				}
			}
			// decode csi_RS_IM_ReceptionForFeedback optional
			if csi_RS_IM_ReceptionForFeedbackPresent {
				ie.csi_RS_IM_ReceptionForFeedback = new(CSI_RS_IM_ReceptionForFeedback)
				if err = ie.csi_RS_IM_ReceptionForFeedback.Decode(extReader); err != nil {
					return utils.WrapError("Decode csi_RS_IM_ReceptionForFeedback", err)
				}
			}
			// decode csi_RS_ProcFrameworkForSRS optional
			if csi_RS_ProcFrameworkForSRSPresent {
				ie.csi_RS_ProcFrameworkForSRS = new(CSI_RS_ProcFrameworkForSRS)
				if err = ie.csi_RS_ProcFrameworkForSRS.Decode(extReader); err != nil {
					return utils.WrapError("Decode csi_RS_ProcFrameworkForSRS", err)
				}
			}
			// decode csi_ReportFramework optional
			if csi_ReportFrameworkPresent {
				ie.csi_ReportFramework = new(CSI_ReportFramework)
				if err = ie.csi_ReportFramework.Decode(extReader); err != nil {
					return utils.WrapError("Decode csi_ReportFramework", err)
				}
			}
			// decode csi_RS_ForTracking optional
			if csi_RS_ForTrackingPresent {
				ie.csi_RS_ForTracking = new(CSI_RS_ForTracking)
				if err = ie.csi_RS_ForTracking.Decode(extReader); err != nil {
					return utils.WrapError("Decode csi_RS_ForTracking", err)
				}
			}
			// decode srs_AssocCSI_RS optional
			if srs_AssocCSI_RSPresent {
				tmp_srs_AssocCSI_RS := utils.NewSequence[*SupportedCSI_RS_Resource]([]*SupportedCSI_RS_Resource{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_Resources}, false)
				fn_srs_AssocCSI_RS := func() *SupportedCSI_RS_Resource {
					return new(SupportedCSI_RS_Resource)
				}
				if err = tmp_srs_AssocCSI_RS.Decode(extReader, fn_srs_AssocCSI_RS); err != nil {
					return utils.WrapError("Decode srs_AssocCSI_RS", err)
				}
				ie.srs_AssocCSI_RS = []SupportedCSI_RS_Resource{}
				for _, i := range tmp_srs_AssocCSI_RS.Value {
					ie.srs_AssocCSI_RS = append(ie.srs_AssocCSI_RS, *i)
				}
			}
			// decode spatialRelations optional
			if spatialRelationsPresent {
				ie.spatialRelations = new(SpatialRelations)
				if err = ie.spatialRelations.Decode(extReader); err != nil {
					return utils.WrapError("Decode spatialRelations", err)
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

			defaultQCL_TwoTCI_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			codebookParametersPerBand_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			simul_SpatialRelationUpdatePUCCHResGroup_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumberSCellBFR_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			simultaneousReceptionDiffTypeD_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ssb_csirs_SINR_measurement_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nonGroupSINR_reporting_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			groupSINR_reporting_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			multiDCI_multiTRP_Parameters_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			singleDCI_SDM_scheme_Parameters_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportFDM_SchemeA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportCodeWordSoftCombining_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportTDM_SchemeA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportInter_slotTDM_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			lowPAPR_DMRS_PDSCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			lowPAPR_DMRS_PUSCHwithoutPrecoding_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			lowPAPR_DMRS_PUCCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			lowPAPR_DMRS_PUSCHwithPrecoding_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			csi_ReportFrameworkExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			codebookParametersAddition_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			codebookComboParametersAddition_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			beamCorrespondenceSSB_based_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			beamCorrespondenceCSI_RS_based_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			beamSwitchTiming_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode defaultQCL_TwoTCI_r16 optional
			if defaultQCL_TwoTCI_r16Present {
				ie.defaultQCL_TwoTCI_r16 = new(MIMO_ParametersPerBand_defaultQCL_TwoTCI_r16)
				if err = ie.defaultQCL_TwoTCI_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode defaultQCL_TwoTCI_r16", err)
				}
			}
			// decode codebookParametersPerBand_r16 optional
			if codebookParametersPerBand_r16Present {
				ie.codebookParametersPerBand_r16 = new(CodebookParameters_v1610)
				if err = ie.codebookParametersPerBand_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode codebookParametersPerBand_r16", err)
				}
			}
			// decode simul_SpatialRelationUpdatePUCCHResGroup_r16 optional
			if simul_SpatialRelationUpdatePUCCHResGroup_r16Present {
				ie.simul_SpatialRelationUpdatePUCCHResGroup_r16 = new(MIMO_ParametersPerBand_simul_SpatialRelationUpdatePUCCHResGroup_r16)
				if err = ie.simul_SpatialRelationUpdatePUCCHResGroup_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode simul_SpatialRelationUpdatePUCCHResGroup_r16", err)
				}
			}
			// decode maxNumberSCellBFR_r16 optional
			if maxNumberSCellBFR_r16Present {
				ie.maxNumberSCellBFR_r16 = new(MIMO_ParametersPerBand_maxNumberSCellBFR_r16)
				if err = ie.maxNumberSCellBFR_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxNumberSCellBFR_r16", err)
				}
			}
			// decode simultaneousReceptionDiffTypeD_r16 optional
			if simultaneousReceptionDiffTypeD_r16Present {
				ie.simultaneousReceptionDiffTypeD_r16 = new(MIMO_ParametersPerBand_simultaneousReceptionDiffTypeD_r16)
				if err = ie.simultaneousReceptionDiffTypeD_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode simultaneousReceptionDiffTypeD_r16", err)
				}
			}
			// decode ssb_csirs_SINR_measurement_r16 optional
			if ssb_csirs_SINR_measurement_r16Present {
				ie.ssb_csirs_SINR_measurement_r16 = new(MIMO_ParametersPerBand_ssb_csirs_SINR_measurement_r16)
				if err = ie.ssb_csirs_SINR_measurement_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ssb_csirs_SINR_measurement_r16", err)
				}
			}
			// decode nonGroupSINR_reporting_r16 optional
			if nonGroupSINR_reporting_r16Present {
				ie.nonGroupSINR_reporting_r16 = new(MIMO_ParametersPerBand_nonGroupSINR_reporting_r16)
				if err = ie.nonGroupSINR_reporting_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode nonGroupSINR_reporting_r16", err)
				}
			}
			// decode groupSINR_reporting_r16 optional
			if groupSINR_reporting_r16Present {
				ie.groupSINR_reporting_r16 = new(MIMO_ParametersPerBand_groupSINR_reporting_r16)
				if err = ie.groupSINR_reporting_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode groupSINR_reporting_r16", err)
				}
			}
			// decode multiDCI_multiTRP_Parameters_r16 optional
			if multiDCI_multiTRP_Parameters_r16Present {
				ie.multiDCI_multiTRP_Parameters_r16 = new(MIMO_ParametersPerBand_multiDCI_multiTRP_Parameters_r16)
				if err = ie.multiDCI_multiTRP_Parameters_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode multiDCI_multiTRP_Parameters_r16", err)
				}
			}
			// decode singleDCI_SDM_scheme_Parameters_r16 optional
			if singleDCI_SDM_scheme_Parameters_r16Present {
				ie.singleDCI_SDM_scheme_Parameters_r16 = new(MIMO_ParametersPerBand_singleDCI_SDM_scheme_Parameters_r16)
				if err = ie.singleDCI_SDM_scheme_Parameters_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode singleDCI_SDM_scheme_Parameters_r16", err)
				}
			}
			// decode supportFDM_SchemeA_r16 optional
			if supportFDM_SchemeA_r16Present {
				ie.supportFDM_SchemeA_r16 = new(MIMO_ParametersPerBand_supportFDM_SchemeA_r16)
				if err = ie.supportFDM_SchemeA_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportFDM_SchemeA_r16", err)
				}
			}
			// decode supportCodeWordSoftCombining_r16 optional
			if supportCodeWordSoftCombining_r16Present {
				ie.supportCodeWordSoftCombining_r16 = new(MIMO_ParametersPerBand_supportCodeWordSoftCombining_r16)
				if err = ie.supportCodeWordSoftCombining_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportCodeWordSoftCombining_r16", err)
				}
			}
			// decode supportTDM_SchemeA_r16 optional
			if supportTDM_SchemeA_r16Present {
				ie.supportTDM_SchemeA_r16 = new(MIMO_ParametersPerBand_supportTDM_SchemeA_r16)
				if err = ie.supportTDM_SchemeA_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportTDM_SchemeA_r16", err)
				}
			}
			// decode supportInter_slotTDM_r16 optional
			if supportInter_slotTDM_r16Present {
				ie.supportInter_slotTDM_r16 = new(MIMO_ParametersPerBand_supportInter_slotTDM_r16)
				if err = ie.supportInter_slotTDM_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportInter_slotTDM_r16", err)
				}
			}
			// decode lowPAPR_DMRS_PDSCH_r16 optional
			if lowPAPR_DMRS_PDSCH_r16Present {
				ie.lowPAPR_DMRS_PDSCH_r16 = new(MIMO_ParametersPerBand_lowPAPR_DMRS_PDSCH_r16)
				if err = ie.lowPAPR_DMRS_PDSCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode lowPAPR_DMRS_PDSCH_r16", err)
				}
			}
			// decode lowPAPR_DMRS_PUSCHwithoutPrecoding_r16 optional
			if lowPAPR_DMRS_PUSCHwithoutPrecoding_r16Present {
				ie.lowPAPR_DMRS_PUSCHwithoutPrecoding_r16 = new(MIMO_ParametersPerBand_lowPAPR_DMRS_PUSCHwithoutPrecoding_r16)
				if err = ie.lowPAPR_DMRS_PUSCHwithoutPrecoding_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode lowPAPR_DMRS_PUSCHwithoutPrecoding_r16", err)
				}
			}
			// decode lowPAPR_DMRS_PUCCH_r16 optional
			if lowPAPR_DMRS_PUCCH_r16Present {
				ie.lowPAPR_DMRS_PUCCH_r16 = new(MIMO_ParametersPerBand_lowPAPR_DMRS_PUCCH_r16)
				if err = ie.lowPAPR_DMRS_PUCCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode lowPAPR_DMRS_PUCCH_r16", err)
				}
			}
			// decode lowPAPR_DMRS_PUSCHwithPrecoding_r16 optional
			if lowPAPR_DMRS_PUSCHwithPrecoding_r16Present {
				ie.lowPAPR_DMRS_PUSCHwithPrecoding_r16 = new(MIMO_ParametersPerBand_lowPAPR_DMRS_PUSCHwithPrecoding_r16)
				if err = ie.lowPAPR_DMRS_PUSCHwithPrecoding_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode lowPAPR_DMRS_PUSCHwithPrecoding_r16", err)
				}
			}
			// decode csi_ReportFrameworkExt_r16 optional
			if csi_ReportFrameworkExt_r16Present {
				ie.csi_ReportFrameworkExt_r16 = new(CSI_ReportFrameworkExt_r16)
				if err = ie.csi_ReportFrameworkExt_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode csi_ReportFrameworkExt_r16", err)
				}
			}
			// decode codebookParametersAddition_r16 optional
			if codebookParametersAddition_r16Present {
				ie.codebookParametersAddition_r16 = new(CodebookParametersAddition_r16)
				if err = ie.codebookParametersAddition_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode codebookParametersAddition_r16", err)
				}
			}
			// decode codebookComboParametersAddition_r16 optional
			if codebookComboParametersAddition_r16Present {
				ie.codebookComboParametersAddition_r16 = new(CodebookComboParametersAddition_r16)
				if err = ie.codebookComboParametersAddition_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode codebookComboParametersAddition_r16", err)
				}
			}
			// decode beamCorrespondenceSSB_based_r16 optional
			if beamCorrespondenceSSB_based_r16Present {
				ie.beamCorrespondenceSSB_based_r16 = new(MIMO_ParametersPerBand_beamCorrespondenceSSB_based_r16)
				if err = ie.beamCorrespondenceSSB_based_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode beamCorrespondenceSSB_based_r16", err)
				}
			}
			// decode beamCorrespondenceCSI_RS_based_r16 optional
			if beamCorrespondenceCSI_RS_based_r16Present {
				ie.beamCorrespondenceCSI_RS_based_r16 = new(MIMO_ParametersPerBand_beamCorrespondenceCSI_RS_based_r16)
				if err = ie.beamCorrespondenceCSI_RS_based_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode beamCorrespondenceCSI_RS_based_r16", err)
				}
			}
			// decode beamSwitchTiming_r16 optional
			if beamSwitchTiming_r16Present {
				ie.beamSwitchTiming_r16 = new(MIMO_ParametersPerBand_beamSwitchTiming_r16)
				if err = ie.beamSwitchTiming_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode beamSwitchTiming_r16", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			semi_PersistentL1_SINR_Report_PUCCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			semi_PersistentL1_SINR_Report_PUSCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode semi_PersistentL1_SINR_Report_PUCCH_r16 optional
			if semi_PersistentL1_SINR_Report_PUCCH_r16Present {
				ie.semi_PersistentL1_SINR_Report_PUCCH_r16 = new(MIMO_ParametersPerBand_semi_PersistentL1_SINR_Report_PUCCH_r16)
				if err = ie.semi_PersistentL1_SINR_Report_PUCCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode semi_PersistentL1_SINR_Report_PUCCH_r16", err)
				}
			}
			// decode semi_PersistentL1_SINR_Report_PUSCH_r16 optional
			if semi_PersistentL1_SINR_Report_PUSCH_r16Present {
				ie.semi_PersistentL1_SINR_Report_PUSCH_r16 = new(MIMO_ParametersPerBand_semi_PersistentL1_SINR_Report_PUSCH_r16)
				if err = ie.semi_PersistentL1_SINR_Report_PUSCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode semi_PersistentL1_SINR_Report_PUSCH_r16", err)
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			spatialRelations_v1640Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			support64CandidateBeamRS_BFR_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode spatialRelations_v1640 optional
			if spatialRelations_v1640Present {
				ie.spatialRelations_v1640 = new(MIMO_ParametersPerBand_spatialRelations_v1640)
				if err = ie.spatialRelations_v1640.Decode(extReader); err != nil {
					return utils.WrapError("Decode spatialRelations_v1640", err)
				}
			}
			// decode support64CandidateBeamRS_BFR_r16 optional
			if support64CandidateBeamRS_BFR_r16Present {
				ie.support64CandidateBeamRS_BFR_r16 = new(MIMO_ParametersPerBand_support64CandidateBeamRS_BFR_r16)
				if err = ie.support64CandidateBeamRS_BFR_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode support64CandidateBeamRS_BFR_r16", err)
				}
			}
		}
		// decode extension group 5
		if len(extBitmap) > 4 && extBitmap[4] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			maxMIMO_LayersForMulti_DCI_mTRP_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode maxMIMO_LayersForMulti_DCI_mTRP_r16 optional
			if maxMIMO_LayersForMulti_DCI_mTRP_r16Present {
				ie.maxMIMO_LayersForMulti_DCI_mTRP_r16 = new(MIMO_ParametersPerBand_maxMIMO_LayersForMulti_DCI_mTRP_r16)
				if err = ie.maxMIMO_LayersForMulti_DCI_mTRP_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxMIMO_LayersForMulti_DCI_mTRP_r16", err)
				}
			}
		}
		// decode extension group 6
		if len(extBitmap) > 5 && extBitmap[5] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			supportedSINR_meas_v1670Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode supportedSINR_meas_v1670 optional
			if supportedSINR_meas_v1670Present {
				var tmp_bs_supportedSINR_meas_v1670 []byte
				var n_supportedSINR_meas_v1670 uint
				if tmp_bs_supportedSINR_meas_v1670, n_supportedSINR_meas_v1670, err = extReader.ReadBitString(&uper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
					return utils.WrapError("Decode supportedSINR_meas_v1670", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_supportedSINR_meas_v1670,
					NumBits: uint64(n_supportedSINR_meas_v1670),
				}
				ie.supportedSINR_meas_v1670 = &tmp_bitstring
			}
		}
		// decode extension group 7
		if len(extBitmap) > 6 && extBitmap[6] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			srs_increasedRepetition_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_partialFrequencySounding_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_startRB_locationHoppingPartial_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_combEight_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			codebookParametersfetype2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_PUSCH_twoCSI_RS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_PUCCH_InterSlot_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_PUCCH_CyclicMapping_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_PUCCH_SecondTPC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_BFR_twoBFD_RS_Set_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_BFR_PUCCH_SR_perCG_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_BFR_association_PUCCH_SR_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sfn_SimulTwoTCI_AcrossMultiCC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sfn_DefaultDL_BeamSetup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sfn_DefaultUL_BeamSetup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_TriggeringOffset_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_TriggeringDCI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			codebookComboParameterMixedType_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedJointTCI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedJointTCI_multiMAC_CE_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedJointTCI_perBWP_CA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedJointTCI_ListSharingCA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedJointTCI_commonMultiCC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedJointTCI_BeamAlignDLRS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedJointTCI_PC_association_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedJointTCI_Legacy_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedJointTCI_Legacy_SRS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedJointTCI_Legacy_CORESET0_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedJointTCI_SCellBFR_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedJointTCI_InterCell_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedSeparateTCI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedSeparateTCI_multiMAC_CE_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedSeparateTCI_perBWP_CA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedSeparateTCI_ListSharingCA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedSeparateTCI_commonMultiCC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedSeparateTCI_InterCell_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedJointTCI_mTRP_InterCell_BM_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mpe_Mitigation_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_PortReport_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_PDCCH_individual_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_PDCCH_anySpan_3Symbols_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_PDCCH_TwoQCL_TypeD_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_PUSCH_CSI_RS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_PUSCH_cyclicMapping_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_PUSCH_secondTPC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_PUSCH_twoPHR_Reporting_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_PUSCH_A_CSI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_PUSCH_SP_CSI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_PUSCH_CG_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_PUCCH_MAC_CE_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_PUCCH_maxNum_PC_FR1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_inter_Cell_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_GroupBasedL1_RSRP_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_BFD_RS_MAC_CE_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_CSI_EnhancementPerBand_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			codebookComboParameterMultiTRP_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_CSI_additionalCSI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_CSI_N_Max2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_CSI_CMR_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_partialFreqSounding_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			beamSwitchTiming_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			beamSwitchTiming_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			beamReportTiming_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumberRxTxBeamSwitchDL_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode srs_increasedRepetition_r17 optional
			if srs_increasedRepetition_r17Present {
				ie.srs_increasedRepetition_r17 = new(MIMO_ParametersPerBand_srs_increasedRepetition_r17)
				if err = ie.srs_increasedRepetition_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode srs_increasedRepetition_r17", err)
				}
			}
			// decode srs_partialFrequencySounding_r17 optional
			if srs_partialFrequencySounding_r17Present {
				ie.srs_partialFrequencySounding_r17 = new(MIMO_ParametersPerBand_srs_partialFrequencySounding_r17)
				if err = ie.srs_partialFrequencySounding_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode srs_partialFrequencySounding_r17", err)
				}
			}
			// decode srs_startRB_locationHoppingPartial_r17 optional
			if srs_startRB_locationHoppingPartial_r17Present {
				ie.srs_startRB_locationHoppingPartial_r17 = new(MIMO_ParametersPerBand_srs_startRB_locationHoppingPartial_r17)
				if err = ie.srs_startRB_locationHoppingPartial_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode srs_startRB_locationHoppingPartial_r17", err)
				}
			}
			// decode srs_combEight_r17 optional
			if srs_combEight_r17Present {
				ie.srs_combEight_r17 = new(MIMO_ParametersPerBand_srs_combEight_r17)
				if err = ie.srs_combEight_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode srs_combEight_r17", err)
				}
			}
			// decode codebookParametersfetype2_r17 optional
			if codebookParametersfetype2_r17Present {
				ie.codebookParametersfetype2_r17 = new(CodebookParametersfetype2_r17)
				if err = ie.codebookParametersfetype2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode codebookParametersfetype2_r17", err)
				}
			}
			// decode mTRP_PUSCH_twoCSI_RS_r17 optional
			if mTRP_PUSCH_twoCSI_RS_r17Present {
				ie.mTRP_PUSCH_twoCSI_RS_r17 = new(MIMO_ParametersPerBand_mTRP_PUSCH_twoCSI_RS_r17)
				if err = ie.mTRP_PUSCH_twoCSI_RS_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_PUSCH_twoCSI_RS_r17", err)
				}
			}
			// decode mTRP_PUCCH_InterSlot_r17 optional
			if mTRP_PUCCH_InterSlot_r17Present {
				ie.mTRP_PUCCH_InterSlot_r17 = new(MIMO_ParametersPerBand_mTRP_PUCCH_InterSlot_r17)
				if err = ie.mTRP_PUCCH_InterSlot_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_PUCCH_InterSlot_r17", err)
				}
			}
			// decode mTRP_PUCCH_CyclicMapping_r17 optional
			if mTRP_PUCCH_CyclicMapping_r17Present {
				ie.mTRP_PUCCH_CyclicMapping_r17 = new(MIMO_ParametersPerBand_mTRP_PUCCH_CyclicMapping_r17)
				if err = ie.mTRP_PUCCH_CyclicMapping_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_PUCCH_CyclicMapping_r17", err)
				}
			}
			// decode mTRP_PUCCH_SecondTPC_r17 optional
			if mTRP_PUCCH_SecondTPC_r17Present {
				ie.mTRP_PUCCH_SecondTPC_r17 = new(MIMO_ParametersPerBand_mTRP_PUCCH_SecondTPC_r17)
				if err = ie.mTRP_PUCCH_SecondTPC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_PUCCH_SecondTPC_r17", err)
				}
			}
			// decode mTRP_BFR_twoBFD_RS_Set_r17 optional
			if mTRP_BFR_twoBFD_RS_Set_r17Present {
				ie.mTRP_BFR_twoBFD_RS_Set_r17 = new(MIMO_ParametersPerBand_mTRP_BFR_twoBFD_RS_Set_r17)
				if err = ie.mTRP_BFR_twoBFD_RS_Set_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_BFR_twoBFD_RS_Set_r17", err)
				}
			}
			// decode mTRP_BFR_PUCCH_SR_perCG_r17 optional
			if mTRP_BFR_PUCCH_SR_perCG_r17Present {
				ie.mTRP_BFR_PUCCH_SR_perCG_r17 = new(MIMO_ParametersPerBand_mTRP_BFR_PUCCH_SR_perCG_r17)
				if err = ie.mTRP_BFR_PUCCH_SR_perCG_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_BFR_PUCCH_SR_perCG_r17", err)
				}
			}
			// decode mTRP_BFR_association_PUCCH_SR_r17 optional
			if mTRP_BFR_association_PUCCH_SR_r17Present {
				ie.mTRP_BFR_association_PUCCH_SR_r17 = new(MIMO_ParametersPerBand_mTRP_BFR_association_PUCCH_SR_r17)
				if err = ie.mTRP_BFR_association_PUCCH_SR_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_BFR_association_PUCCH_SR_r17", err)
				}
			}
			// decode sfn_SimulTwoTCI_AcrossMultiCC_r17 optional
			if sfn_SimulTwoTCI_AcrossMultiCC_r17Present {
				ie.sfn_SimulTwoTCI_AcrossMultiCC_r17 = new(MIMO_ParametersPerBand_sfn_SimulTwoTCI_AcrossMultiCC_r17)
				if err = ie.sfn_SimulTwoTCI_AcrossMultiCC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sfn_SimulTwoTCI_AcrossMultiCC_r17", err)
				}
			}
			// decode sfn_DefaultDL_BeamSetup_r17 optional
			if sfn_DefaultDL_BeamSetup_r17Present {
				ie.sfn_DefaultDL_BeamSetup_r17 = new(MIMO_ParametersPerBand_sfn_DefaultDL_BeamSetup_r17)
				if err = ie.sfn_DefaultDL_BeamSetup_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sfn_DefaultDL_BeamSetup_r17", err)
				}
			}
			// decode sfn_DefaultUL_BeamSetup_r17 optional
			if sfn_DefaultUL_BeamSetup_r17Present {
				ie.sfn_DefaultUL_BeamSetup_r17 = new(MIMO_ParametersPerBand_sfn_DefaultUL_BeamSetup_r17)
				if err = ie.sfn_DefaultUL_BeamSetup_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sfn_DefaultUL_BeamSetup_r17", err)
				}
			}
			// decode srs_TriggeringOffset_r17 optional
			if srs_TriggeringOffset_r17Present {
				ie.srs_TriggeringOffset_r17 = new(MIMO_ParametersPerBand_srs_TriggeringOffset_r17)
				if err = ie.srs_TriggeringOffset_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode srs_TriggeringOffset_r17", err)
				}
			}
			// decode srs_TriggeringDCI_r17 optional
			if srs_TriggeringDCI_r17Present {
				ie.srs_TriggeringDCI_r17 = new(MIMO_ParametersPerBand_srs_TriggeringDCI_r17)
				if err = ie.srs_TriggeringDCI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode srs_TriggeringDCI_r17", err)
				}
			}
			// decode codebookComboParameterMixedType_r17 optional
			if codebookComboParameterMixedType_r17Present {
				ie.codebookComboParameterMixedType_r17 = new(CodebookComboParameterMixedType_r17)
				if err = ie.codebookComboParameterMixedType_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode codebookComboParameterMixedType_r17", err)
				}
			}
			// decode unifiedJointTCI_r17 optional
			if unifiedJointTCI_r17Present {
				ie.unifiedJointTCI_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_r17)
				if err = ie.unifiedJointTCI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode unifiedJointTCI_r17", err)
				}
			}
			// decode unifiedJointTCI_multiMAC_CE_r17 optional
			if unifiedJointTCI_multiMAC_CE_r17Present {
				ie.unifiedJointTCI_multiMAC_CE_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_multiMAC_CE_r17)
				if err = ie.unifiedJointTCI_multiMAC_CE_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode unifiedJointTCI_multiMAC_CE_r17", err)
				}
			}
			// decode unifiedJointTCI_perBWP_CA_r17 optional
			if unifiedJointTCI_perBWP_CA_r17Present {
				ie.unifiedJointTCI_perBWP_CA_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_perBWP_CA_r17)
				if err = ie.unifiedJointTCI_perBWP_CA_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode unifiedJointTCI_perBWP_CA_r17", err)
				}
			}
			// decode unifiedJointTCI_ListSharingCA_r17 optional
			if unifiedJointTCI_ListSharingCA_r17Present {
				ie.unifiedJointTCI_ListSharingCA_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_ListSharingCA_r17)
				if err = ie.unifiedJointTCI_ListSharingCA_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode unifiedJointTCI_ListSharingCA_r17", err)
				}
			}
			// decode unifiedJointTCI_commonMultiCC_r17 optional
			if unifiedJointTCI_commonMultiCC_r17Present {
				ie.unifiedJointTCI_commonMultiCC_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_commonMultiCC_r17)
				if err = ie.unifiedJointTCI_commonMultiCC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode unifiedJointTCI_commonMultiCC_r17", err)
				}
			}
			// decode unifiedJointTCI_BeamAlignDLRS_r17 optional
			if unifiedJointTCI_BeamAlignDLRS_r17Present {
				ie.unifiedJointTCI_BeamAlignDLRS_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_BeamAlignDLRS_r17)
				if err = ie.unifiedJointTCI_BeamAlignDLRS_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode unifiedJointTCI_BeamAlignDLRS_r17", err)
				}
			}
			// decode unifiedJointTCI_PC_association_r17 optional
			if unifiedJointTCI_PC_association_r17Present {
				ie.unifiedJointTCI_PC_association_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_PC_association_r17)
				if err = ie.unifiedJointTCI_PC_association_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode unifiedJointTCI_PC_association_r17", err)
				}
			}
			// decode unifiedJointTCI_Legacy_r17 optional
			if unifiedJointTCI_Legacy_r17Present {
				ie.unifiedJointTCI_Legacy_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_Legacy_r17)
				if err = ie.unifiedJointTCI_Legacy_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode unifiedJointTCI_Legacy_r17", err)
				}
			}
			// decode unifiedJointTCI_Legacy_SRS_r17 optional
			if unifiedJointTCI_Legacy_SRS_r17Present {
				ie.unifiedJointTCI_Legacy_SRS_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_Legacy_SRS_r17)
				if err = ie.unifiedJointTCI_Legacy_SRS_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode unifiedJointTCI_Legacy_SRS_r17", err)
				}
			}
			// decode unifiedJointTCI_Legacy_CORESET0_r17 optional
			if unifiedJointTCI_Legacy_CORESET0_r17Present {
				ie.unifiedJointTCI_Legacy_CORESET0_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_Legacy_CORESET0_r17)
				if err = ie.unifiedJointTCI_Legacy_CORESET0_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode unifiedJointTCI_Legacy_CORESET0_r17", err)
				}
			}
			// decode unifiedJointTCI_SCellBFR_r17 optional
			if unifiedJointTCI_SCellBFR_r17Present {
				ie.unifiedJointTCI_SCellBFR_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_SCellBFR_r17)
				if err = ie.unifiedJointTCI_SCellBFR_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode unifiedJointTCI_SCellBFR_r17", err)
				}
			}
			// decode unifiedJointTCI_InterCell_r17 optional
			if unifiedJointTCI_InterCell_r17Present {
				ie.unifiedJointTCI_InterCell_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_InterCell_r17)
				if err = ie.unifiedJointTCI_InterCell_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode unifiedJointTCI_InterCell_r17", err)
				}
			}
			// decode unifiedSeparateTCI_r17 optional
			if unifiedSeparateTCI_r17Present {
				ie.unifiedSeparateTCI_r17 = new(MIMO_ParametersPerBand_unifiedSeparateTCI_r17)
				if err = ie.unifiedSeparateTCI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode unifiedSeparateTCI_r17", err)
				}
			}
			// decode unifiedSeparateTCI_multiMAC_CE_r17 optional
			if unifiedSeparateTCI_multiMAC_CE_r17Present {
				ie.unifiedSeparateTCI_multiMAC_CE_r17 = new(MIMO_ParametersPerBand_unifiedSeparateTCI_multiMAC_CE_r17)
				if err = ie.unifiedSeparateTCI_multiMAC_CE_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode unifiedSeparateTCI_multiMAC_CE_r17", err)
				}
			}
			// decode unifiedSeparateTCI_perBWP_CA_r17 optional
			if unifiedSeparateTCI_perBWP_CA_r17Present {
				ie.unifiedSeparateTCI_perBWP_CA_r17 = new(MIMO_ParametersPerBand_unifiedSeparateTCI_perBWP_CA_r17)
				if err = ie.unifiedSeparateTCI_perBWP_CA_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode unifiedSeparateTCI_perBWP_CA_r17", err)
				}
			}
			// decode unifiedSeparateTCI_ListSharingCA_r17 optional
			if unifiedSeparateTCI_ListSharingCA_r17Present {
				ie.unifiedSeparateTCI_ListSharingCA_r17 = new(MIMO_ParametersPerBand_unifiedSeparateTCI_ListSharingCA_r17)
				if err = ie.unifiedSeparateTCI_ListSharingCA_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode unifiedSeparateTCI_ListSharingCA_r17", err)
				}
			}
			// decode unifiedSeparateTCI_commonMultiCC_r17 optional
			if unifiedSeparateTCI_commonMultiCC_r17Present {
				ie.unifiedSeparateTCI_commonMultiCC_r17 = new(MIMO_ParametersPerBand_unifiedSeparateTCI_commonMultiCC_r17)
				if err = ie.unifiedSeparateTCI_commonMultiCC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode unifiedSeparateTCI_commonMultiCC_r17", err)
				}
			}
			// decode unifiedSeparateTCI_InterCell_r17 optional
			if unifiedSeparateTCI_InterCell_r17Present {
				ie.unifiedSeparateTCI_InterCell_r17 = new(MIMO_ParametersPerBand_unifiedSeparateTCI_InterCell_r17)
				if err = ie.unifiedSeparateTCI_InterCell_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode unifiedSeparateTCI_InterCell_r17", err)
				}
			}
			// decode unifiedJointTCI_mTRP_InterCell_BM_r17 optional
			if unifiedJointTCI_mTRP_InterCell_BM_r17Present {
				ie.unifiedJointTCI_mTRP_InterCell_BM_r17 = new(MIMO_ParametersPerBand_unifiedJointTCI_mTRP_InterCell_BM_r17)
				if err = ie.unifiedJointTCI_mTRP_InterCell_BM_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode unifiedJointTCI_mTRP_InterCell_BM_r17", err)
				}
			}
			// decode mpe_Mitigation_r17 optional
			if mpe_Mitigation_r17Present {
				ie.mpe_Mitigation_r17 = new(MIMO_ParametersPerBand_mpe_Mitigation_r17)
				if err = ie.mpe_Mitigation_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mpe_Mitigation_r17", err)
				}
			}
			// decode srs_PortReport_r17 optional
			if srs_PortReport_r17Present {
				ie.srs_PortReport_r17 = new(MIMO_ParametersPerBand_srs_PortReport_r17)
				if err = ie.srs_PortReport_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode srs_PortReport_r17", err)
				}
			}
			// decode mTRP_PDCCH_individual_r17 optional
			if mTRP_PDCCH_individual_r17Present {
				ie.mTRP_PDCCH_individual_r17 = new(MIMO_ParametersPerBand_mTRP_PDCCH_individual_r17)
				if err = ie.mTRP_PDCCH_individual_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_PDCCH_individual_r17", err)
				}
			}
			// decode mTRP_PDCCH_anySpan_3Symbols_r17 optional
			if mTRP_PDCCH_anySpan_3Symbols_r17Present {
				ie.mTRP_PDCCH_anySpan_3Symbols_r17 = new(MIMO_ParametersPerBand_mTRP_PDCCH_anySpan_3Symbols_r17)
				if err = ie.mTRP_PDCCH_anySpan_3Symbols_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_PDCCH_anySpan_3Symbols_r17", err)
				}
			}
			// decode mTRP_PDCCH_TwoQCL_TypeD_r17 optional
			if mTRP_PDCCH_TwoQCL_TypeD_r17Present {
				ie.mTRP_PDCCH_TwoQCL_TypeD_r17 = new(MIMO_ParametersPerBand_mTRP_PDCCH_TwoQCL_TypeD_r17)
				if err = ie.mTRP_PDCCH_TwoQCL_TypeD_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_PDCCH_TwoQCL_TypeD_r17", err)
				}
			}
			// decode mTRP_PUSCH_CSI_RS_r17 optional
			if mTRP_PUSCH_CSI_RS_r17Present {
				ie.mTRP_PUSCH_CSI_RS_r17 = new(MIMO_ParametersPerBand_mTRP_PUSCH_CSI_RS_r17)
				if err = ie.mTRP_PUSCH_CSI_RS_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_PUSCH_CSI_RS_r17", err)
				}
			}
			// decode mTRP_PUSCH_cyclicMapping_r17 optional
			if mTRP_PUSCH_cyclicMapping_r17Present {
				ie.mTRP_PUSCH_cyclicMapping_r17 = new(MIMO_ParametersPerBand_mTRP_PUSCH_cyclicMapping_r17)
				if err = ie.mTRP_PUSCH_cyclicMapping_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_PUSCH_cyclicMapping_r17", err)
				}
			}
			// decode mTRP_PUSCH_secondTPC_r17 optional
			if mTRP_PUSCH_secondTPC_r17Present {
				ie.mTRP_PUSCH_secondTPC_r17 = new(MIMO_ParametersPerBand_mTRP_PUSCH_secondTPC_r17)
				if err = ie.mTRP_PUSCH_secondTPC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_PUSCH_secondTPC_r17", err)
				}
			}
			// decode mTRP_PUSCH_twoPHR_Reporting_r17 optional
			if mTRP_PUSCH_twoPHR_Reporting_r17Present {
				ie.mTRP_PUSCH_twoPHR_Reporting_r17 = new(MIMO_ParametersPerBand_mTRP_PUSCH_twoPHR_Reporting_r17)
				if err = ie.mTRP_PUSCH_twoPHR_Reporting_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_PUSCH_twoPHR_Reporting_r17", err)
				}
			}
			// decode mTRP_PUSCH_A_CSI_r17 optional
			if mTRP_PUSCH_A_CSI_r17Present {
				ie.mTRP_PUSCH_A_CSI_r17 = new(MIMO_ParametersPerBand_mTRP_PUSCH_A_CSI_r17)
				if err = ie.mTRP_PUSCH_A_CSI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_PUSCH_A_CSI_r17", err)
				}
			}
			// decode mTRP_PUSCH_SP_CSI_r17 optional
			if mTRP_PUSCH_SP_CSI_r17Present {
				ie.mTRP_PUSCH_SP_CSI_r17 = new(MIMO_ParametersPerBand_mTRP_PUSCH_SP_CSI_r17)
				if err = ie.mTRP_PUSCH_SP_CSI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_PUSCH_SP_CSI_r17", err)
				}
			}
			// decode mTRP_PUSCH_CG_r17 optional
			if mTRP_PUSCH_CG_r17Present {
				ie.mTRP_PUSCH_CG_r17 = new(MIMO_ParametersPerBand_mTRP_PUSCH_CG_r17)
				if err = ie.mTRP_PUSCH_CG_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_PUSCH_CG_r17", err)
				}
			}
			// decode mTRP_PUCCH_MAC_CE_r17 optional
			if mTRP_PUCCH_MAC_CE_r17Present {
				ie.mTRP_PUCCH_MAC_CE_r17 = new(MIMO_ParametersPerBand_mTRP_PUCCH_MAC_CE_r17)
				if err = ie.mTRP_PUCCH_MAC_CE_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_PUCCH_MAC_CE_r17", err)
				}
			}
			// decode mTRP_PUCCH_maxNum_PC_FR1_r17 optional
			if mTRP_PUCCH_maxNum_PC_FR1_r17Present {
				var tmp_int_mTRP_PUCCH_maxNum_PC_FR1_r17 int64
				if tmp_int_mTRP_PUCCH_maxNum_PC_FR1_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 3, Ub: 8}, false); err != nil {
					return utils.WrapError("Decode mTRP_PUCCH_maxNum_PC_FR1_r17", err)
				}
				ie.mTRP_PUCCH_maxNum_PC_FR1_r17 = &tmp_int_mTRP_PUCCH_maxNum_PC_FR1_r17
			}
			// decode mTRP_inter_Cell_r17 optional
			if mTRP_inter_Cell_r17Present {
				ie.mTRP_inter_Cell_r17 = new(MIMO_ParametersPerBand_mTRP_inter_Cell_r17)
				if err = ie.mTRP_inter_Cell_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_inter_Cell_r17", err)
				}
			}
			// decode mTRP_GroupBasedL1_RSRP_r17 optional
			if mTRP_GroupBasedL1_RSRP_r17Present {
				ie.mTRP_GroupBasedL1_RSRP_r17 = new(MIMO_ParametersPerBand_mTRP_GroupBasedL1_RSRP_r17)
				if err = ie.mTRP_GroupBasedL1_RSRP_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_GroupBasedL1_RSRP_r17", err)
				}
			}
			// decode mTRP_BFD_RS_MAC_CE_r17 optional
			if mTRP_BFD_RS_MAC_CE_r17Present {
				ie.mTRP_BFD_RS_MAC_CE_r17 = new(MIMO_ParametersPerBand_mTRP_BFD_RS_MAC_CE_r17)
				if err = ie.mTRP_BFD_RS_MAC_CE_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_BFD_RS_MAC_CE_r17", err)
				}
			}
			// decode mTRP_CSI_EnhancementPerBand_r17 optional
			if mTRP_CSI_EnhancementPerBand_r17Present {
				ie.mTRP_CSI_EnhancementPerBand_r17 = new(MIMO_ParametersPerBand_mTRP_CSI_EnhancementPerBand_r17)
				if err = ie.mTRP_CSI_EnhancementPerBand_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_CSI_EnhancementPerBand_r17", err)
				}
			}
			// decode codebookComboParameterMultiTRP_r17 optional
			if codebookComboParameterMultiTRP_r17Present {
				ie.codebookComboParameterMultiTRP_r17 = new(CodebookComboParameterMultiTRP_r17)
				if err = ie.codebookComboParameterMultiTRP_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode codebookComboParameterMultiTRP_r17", err)
				}
			}
			// decode mTRP_CSI_additionalCSI_r17 optional
			if mTRP_CSI_additionalCSI_r17Present {
				ie.mTRP_CSI_additionalCSI_r17 = new(MIMO_ParametersPerBand_mTRP_CSI_additionalCSI_r17)
				if err = ie.mTRP_CSI_additionalCSI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_CSI_additionalCSI_r17", err)
				}
			}
			// decode mTRP_CSI_N_Max2_r17 optional
			if mTRP_CSI_N_Max2_r17Present {
				ie.mTRP_CSI_N_Max2_r17 = new(MIMO_ParametersPerBand_mTRP_CSI_N_Max2_r17)
				if err = ie.mTRP_CSI_N_Max2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_CSI_N_Max2_r17", err)
				}
			}
			// decode mTRP_CSI_CMR_r17 optional
			if mTRP_CSI_CMR_r17Present {
				ie.mTRP_CSI_CMR_r17 = new(MIMO_ParametersPerBand_mTRP_CSI_CMR_r17)
				if err = ie.mTRP_CSI_CMR_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_CSI_CMR_r17", err)
				}
			}
			// decode srs_partialFreqSounding_r17 optional
			if srs_partialFreqSounding_r17Present {
				ie.srs_partialFreqSounding_r17 = new(MIMO_ParametersPerBand_srs_partialFreqSounding_r17)
				if err = ie.srs_partialFreqSounding_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode srs_partialFreqSounding_r17", err)
				}
			}
			// decode beamSwitchTiming_v1710 optional
			if beamSwitchTiming_v1710Present {
				ie.beamSwitchTiming_v1710 = new(MIMO_ParametersPerBand_beamSwitchTiming_v1710)
				if err = ie.beamSwitchTiming_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode beamSwitchTiming_v1710", err)
				}
			}
			// decode beamSwitchTiming_r17 optional
			if beamSwitchTiming_r17Present {
				ie.beamSwitchTiming_r17 = new(MIMO_ParametersPerBand_beamSwitchTiming_r17)
				if err = ie.beamSwitchTiming_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode beamSwitchTiming_r17", err)
				}
			}
			// decode beamReportTiming_v1710 optional
			if beamReportTiming_v1710Present {
				ie.beamReportTiming_v1710 = new(MIMO_ParametersPerBand_beamReportTiming_v1710)
				if err = ie.beamReportTiming_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode beamReportTiming_v1710", err)
				}
			}
			// decode maxNumberRxTxBeamSwitchDL_v1710 optional
			if maxNumberRxTxBeamSwitchDL_v1710Present {
				ie.maxNumberRxTxBeamSwitchDL_v1710 = new(MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_v1710)
				if err = ie.maxNumberRxTxBeamSwitchDL_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxNumberRxTxBeamSwitchDL_v1710", err)
				}
			}
		}
		// decode extension group 8
		if len(extBitmap) > 7 && extBitmap[7] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			srs_PortReportSP_AP_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumberRxBeam_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sfn_ImplicitRS_twoTCI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sfn_QCL_TypeD_Collision_twoTCI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_CSI_numCPU_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode srs_PortReportSP_AP_r17 optional
			if srs_PortReportSP_AP_r17Present {
				ie.srs_PortReportSP_AP_r17 = new(MIMO_ParametersPerBand_srs_PortReportSP_AP_r17)
				if err = ie.srs_PortReportSP_AP_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode srs_PortReportSP_AP_r17", err)
				}
			}
			// decode maxNumberRxBeam_v1720 optional
			if maxNumberRxBeam_v1720Present {
				var tmp_int_maxNumberRxBeam_v1720 int64
				if tmp_int_maxNumberRxBeam_v1720, err = extReader.ReadInteger(&uper.Constraint{Lb: 9, Ub: 12}, false); err != nil {
					return utils.WrapError("Decode maxNumberRxBeam_v1720", err)
				}
				ie.maxNumberRxBeam_v1720 = &tmp_int_maxNumberRxBeam_v1720
			}
			// decode sfn_ImplicitRS_twoTCI_r17 optional
			if sfn_ImplicitRS_twoTCI_r17Present {
				ie.sfn_ImplicitRS_twoTCI_r17 = new(MIMO_ParametersPerBand_sfn_ImplicitRS_twoTCI_r17)
				if err = ie.sfn_ImplicitRS_twoTCI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sfn_ImplicitRS_twoTCI_r17", err)
				}
			}
			// decode sfn_QCL_TypeD_Collision_twoTCI_r17 optional
			if sfn_QCL_TypeD_Collision_twoTCI_r17Present {
				ie.sfn_QCL_TypeD_Collision_twoTCI_r17 = new(MIMO_ParametersPerBand_sfn_QCL_TypeD_Collision_twoTCI_r17)
				if err = ie.sfn_QCL_TypeD_Collision_twoTCI_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sfn_QCL_TypeD_Collision_twoTCI_r17", err)
				}
			}
			// decode mTRP_CSI_numCPU_r17 optional
			if mTRP_CSI_numCPU_r17Present {
				ie.mTRP_CSI_numCPU_r17 = new(MIMO_ParametersPerBand_mTRP_CSI_numCPU_r17)
				if err = ie.mTRP_CSI_numCPU_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_CSI_numCPU_r17", err)
				}
			}
		}
		// decode extension group 9
		if len(extBitmap) > 8 && extBitmap[8] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			supportRepNumPDSCH_TDRA_DCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode supportRepNumPDSCH_TDRA_DCI_1_2_r17 optional
			if supportRepNumPDSCH_TDRA_DCI_1_2_r17Present {
				ie.supportRepNumPDSCH_TDRA_DCI_1_2_r17 = new(MIMO_ParametersPerBand_supportRepNumPDSCH_TDRA_DCI_1_2_r17)
				if err = ie.supportRepNumPDSCH_TDRA_DCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportRepNumPDSCH_TDRA_DCI_1_2_r17", err)
				}
			}
		}
	}
	return nil
}
