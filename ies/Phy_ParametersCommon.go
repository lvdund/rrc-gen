package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersCommon struct {
	csi_RS_CFRA_ForHO                                          *Phy_ParametersCommon_csi_RS_CFRA_ForHO                                          `optional`
	dynamicPRB_BundlingDL                                      *Phy_ParametersCommon_dynamicPRB_BundlingDL                                      `optional`
	sp_CSI_ReportPUCCH                                         *Phy_ParametersCommon_sp_CSI_ReportPUCCH                                         `optional`
	sp_CSI_ReportPUSCH                                         *Phy_ParametersCommon_sp_CSI_ReportPUSCH                                         `optional`
	nzp_CSI_RS_IntefMgmt                                       *Phy_ParametersCommon_nzp_CSI_RS_IntefMgmt                                       `optional`
	type2_SP_CSI_Feedback_LongPUCCH                            *Phy_ParametersCommon_type2_SP_CSI_Feedback_LongPUCCH                            `optional`
	precoderGranularityCORESET                                 *Phy_ParametersCommon_precoderGranularityCORESET                                 `optional`
	dynamicHARQ_ACK_Codebook                                   *Phy_ParametersCommon_dynamicHARQ_ACK_Codebook                                   `optional`
	semiStaticHARQ_ACK_Codebook                                *Phy_ParametersCommon_semiStaticHARQ_ACK_Codebook                                `optional`
	spatialBundlingHARQ_ACK                                    *Phy_ParametersCommon_spatialBundlingHARQ_ACK                                    `optional`
	dynamicBetaOffsetInd_HARQ_ACK_CSI                          *Phy_ParametersCommon_dynamicBetaOffsetInd_HARQ_ACK_CSI                          `optional`
	pucch_Repetition_F1_3_4                                    *Phy_ParametersCommon_pucch_Repetition_F1_3_4                                    `optional`
	ra_Type0_PUSCH                                             *Phy_ParametersCommon_ra_Type0_PUSCH                                             `optional`
	dynamicSwitchRA_Type0_1_PDSCH                              *Phy_ParametersCommon_dynamicSwitchRA_Type0_1_PDSCH                              `optional`
	dynamicSwitchRA_Type0_1_PUSCH                              *Phy_ParametersCommon_dynamicSwitchRA_Type0_1_PUSCH                              `optional`
	pdsch_MappingTypeA                                         *Phy_ParametersCommon_pdsch_MappingTypeA                                         `optional`
	pdsch_MappingTypeB                                         *Phy_ParametersCommon_pdsch_MappingTypeB                                         `optional`
	interleavingVRB_ToPRB_PDSCH                                *Phy_ParametersCommon_interleavingVRB_ToPRB_PDSCH                                `optional`
	interSlotFreqHopping_PUSCH                                 *Phy_ParametersCommon_interSlotFreqHopping_PUSCH                                 `optional`
	type1_PUSCH_RepetitionMultiSlots                           *Phy_ParametersCommon_type1_PUSCH_RepetitionMultiSlots                           `optional`
	type2_PUSCH_RepetitionMultiSlots                           *Phy_ParametersCommon_type2_PUSCH_RepetitionMultiSlots                           `optional`
	pusch_RepetitionMultiSlots                                 *Phy_ParametersCommon_pusch_RepetitionMultiSlots                                 `optional`
	pdsch_RepetitionMultiSlots                                 *Phy_ParametersCommon_pdsch_RepetitionMultiSlots                                 `optional`
	downlinkSPS                                                *Phy_ParametersCommon_downlinkSPS                                                `optional`
	configuredUL_GrantType1                                    *Phy_ParametersCommon_configuredUL_GrantType1                                    `optional`
	configuredUL_GrantType2                                    *Phy_ParametersCommon_configuredUL_GrantType2                                    `optional`
	pre_EmptIndication_DL                                      *Phy_ParametersCommon_pre_EmptIndication_DL                                      `optional`
	cbg_TransIndication_DL                                     *Phy_ParametersCommon_cbg_TransIndication_DL                                     `optional`
	cbg_TransIndication_UL                                     *Phy_ParametersCommon_cbg_TransIndication_UL                                     `optional`
	cbg_FlushIndication_DL                                     *Phy_ParametersCommon_cbg_FlushIndication_DL                                     `optional`
	dynamicHARQ_ACK_CodeB_CBG_Retx_DL                          *Phy_ParametersCommon_dynamicHARQ_ACK_CodeB_CBG_Retx_DL                          `optional`
	rateMatchingResrcSetSemi_Static                            *Phy_ParametersCommon_rateMatchingResrcSetSemi_Static                            `optional`
	rateMatchingResrcSetDynamic                                *Phy_ParametersCommon_rateMatchingResrcSetDynamic                                `optional`
	bwp_SwitchingDelay                                         *Phy_ParametersCommon_bwp_SwitchingDelay                                         `optional`
	dummy                                                      *Phy_ParametersCommon_dummy                                                      `optional,ext-1`
	maxNumberSearchSpaces                                      *Phy_ParametersCommon_maxNumberSearchSpaces                                      `optional,ext-2`
	rateMatchingCtrlResrcSetDynamic                            *Phy_ParametersCommon_rateMatchingCtrlResrcSetDynamic                            `optional,ext-2`
	maxLayersMIMO_Indication                                   *Phy_ParametersCommon_maxLayersMIMO_Indication                                   `optional,ext-2`
	spCellPlacement                                            *CarrierAggregationVariant                                                       `optional,ext-3`
	twoStepRACH_r16                                            *Phy_ParametersCommon_twoStepRACH_r16                                            `optional,ext-4`
	dci_Format1_2And0_2_r16                                    *Phy_ParametersCommon_dci_Format1_2And0_2_r16                                    `optional,ext-4`
	monitoringDCI_SameSearchSpace_r16                          *Phy_ParametersCommon_monitoringDCI_SameSearchSpace_r16                          `optional,ext-4`
	type2_CG_ReleaseDCI_0_1_r16                                *Phy_ParametersCommon_type2_CG_ReleaseDCI_0_1_r16                                `optional,ext-4`
	type2_CG_ReleaseDCI_0_2_r16                                *Phy_ParametersCommon_type2_CG_ReleaseDCI_0_2_r16                                `optional,ext-4`
	sps_ReleaseDCI_1_1_r16                                     *Phy_ParametersCommon_sps_ReleaseDCI_1_1_r16                                     `optional,ext-4`
	sps_ReleaseDCI_1_2_r16                                     *Phy_ParametersCommon_sps_ReleaseDCI_1_2_r16                                     `optional,ext-4`
	csi_TriggerStateNon_ActiveBWP_r16                          *Phy_ParametersCommon_csi_TriggerStateNon_ActiveBWP_r16                          `optional,ext-4`
	separateSMTC_InterIAB_Support_r16                          *Phy_ParametersCommon_separateSMTC_InterIAB_Support_r16                          `optional,ext-4`
	separateRACH_IAB_Support_r16                               *Phy_ParametersCommon_separateRACH_IAB_Support_r16                               `optional,ext-4`
	ul_flexibleDL_SlotFormatSemiStatic_IAB_r16                 *Phy_ParametersCommon_ul_flexibleDL_SlotFormatSemiStatic_IAB_r16                 `optional,ext-4`
	ul_flexibleDL_SlotFormatDynamics_IAB_r16                   *Phy_ParametersCommon_ul_flexibleDL_SlotFormatDynamics_IAB_r16                   `optional,ext-4`
	dft_S_OFDM_WaveformUL_IAB_r16                              *Phy_ParametersCommon_dft_S_OFDM_WaveformUL_IAB_r16                              `optional,ext-4`
	dci_25_AI_RNTI_Support_IAB_r16                             *Phy_ParametersCommon_dci_25_AI_RNTI_Support_IAB_r16                             `optional,ext-4`
	t_DeltaReceptionSupport_IAB_r16                            *Phy_ParametersCommon_t_DeltaReceptionSupport_IAB_r16                            `optional,ext-4`
	guardSymbolReportReception_IAB_r16                         *Phy_ParametersCommon_guardSymbolReportReception_IAB_r16                         `optional,ext-4`
	harqACK_CB_SpatialBundlingPUCCH_Group_r16                  *Phy_ParametersCommon_harqACK_CB_SpatialBundlingPUCCH_Group_r16                  `optional,ext-4`
	crossSlotScheduling_r16                                    *Phy_ParametersCommon_crossSlotScheduling_r16                                    `optional,ext-4`
	maxNumberSRS_PosPathLossEstimateAllServingCells_r16        *Phy_ParametersCommon_maxNumberSRS_PosPathLossEstimateAllServingCells_r16        `optional,ext-4`
	extendedCG_Periodicities_r16                               *Phy_ParametersCommon_extendedCG_Periodicities_r16                               `optional,ext-4`
	extendedSPS_Periodicities_r16                              *Phy_ParametersCommon_extendedSPS_Periodicities_r16                              `optional,ext-4`
	codebookVariantsList_r16                                   *CodebookVariantsList_r16                                                        `optional,ext-4`
	pusch_RepetitionTypeA_r16                                  *Phy_ParametersCommon_pusch_RepetitionTypeA_r16                                  `optional,ext-4`
	dci_DL_PriorityIndicator_r16                               *Phy_ParametersCommon_dci_DL_PriorityIndicator_r16                               `optional,ext-4`
	dci_UL_PriorityIndicator_r16                               *Phy_ParametersCommon_dci_UL_PriorityIndicator_r16                               `optional,ext-4`
	maxNumberPathlossRS_Update_r16                             *Phy_ParametersCommon_maxNumberPathlossRS_Update_r16                             `optional,ext-4`
	type2_HARQ_ACK_Codebook_r16                                *Phy_ParametersCommon_type2_HARQ_ACK_Codebook_r16                                `optional,ext-4`
	maxTotalResourcesForAcrossFreqRanges_r16                   *Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16                   `optional,ext-4`
	harqACK_separateMultiDCI_MultiTRP_r16                      *Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16                      `optional,ext-4`
	harqACK_jointMultiDCI_MultiTRP_r16                         *Phy_ParametersCommon_harqACK_jointMultiDCI_MultiTRP_r16                         `optional,ext-4`
	bwp_SwitchingMultiCCs_r16                                  *Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16                                  `optional,ext-4`
	targetSMTC_SCG_r16                                         *Phy_ParametersCommon_targetSMTC_SCG_r16                                         `optional,ext-5`
	supportRepetitionZeroOffsetRV_r16                          *Phy_ParametersCommon_supportRepetitionZeroOffsetRV_r16                          `optional,ext-5`
	cbg_TransInOrderPUSCH_UL_r16                               *Phy_ParametersCommon_cbg_TransInOrderPUSCH_UL_r16                               `optional,ext-5`
	bwp_SwitchingMultiDormancyCCs_r16                          *Phy_ParametersCommon_bwp_SwitchingMultiDormancyCCs_r16                          `optional,ext-6`
	supportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16             *Phy_ParametersCommon_supportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16             `optional,ext-6`
	pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 *Phy_ParametersCommon_pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 `optional,ext-6`
	newBeamIdentifications2PortCSI_RS_r16                      *Phy_ParametersCommon_newBeamIdentifications2PortCSI_RS_r16                      `optional,ext-7`
	pathlossEstimation2PortCSI_RS_r16                          *Phy_ParametersCommon_pathlossEstimation2PortCSI_RS_r16                          `optional,ext-7`
	mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16                      *Phy_ParametersCommon_mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16                      `optional,ext-8`
	guardSymbolReportReception_IAB_r17                         *Phy_ParametersCommon_guardSymbolReportReception_IAB_r17                         `optional,ext-9`
	restricted_IAB_DU_BeamReception_r17                        *Phy_ParametersCommon_restricted_IAB_DU_BeamReception_r17                        `optional,ext-9`
	recommended_IAB_MT_BeamTransmission_r17                    *Phy_ParametersCommon_recommended_IAB_MT_BeamTransmission_r17                    `optional,ext-9`
	case6_TimingAlignmentReception_IAB_r17                     *Phy_ParametersCommon_case6_TimingAlignmentReception_IAB_r17                     `optional,ext-9`
	case7_TimingAlignmentReception_IAB_r17                     *Phy_ParametersCommon_case7_TimingAlignmentReception_IAB_r17                     `optional,ext-9`
	dl_tx_PowerAdjustment_IAB_r17                              *Phy_ParametersCommon_dl_tx_PowerAdjustment_IAB_r17                              `optional,ext-9`
	desired_ul_tx_PowerAdjustment_r17                          *Phy_ParametersCommon_desired_ul_tx_PowerAdjustment_r17                          `optional,ext-9`
	fdm_SoftResourceAvailability_DynamicIndication_r17         *Phy_ParametersCommon_fdm_SoftResourceAvailability_DynamicIndication_r17         `optional,ext-9`
	updated_T_DeltaRangeRecption_r17                           *Phy_ParametersCommon_updated_T_DeltaRangeRecption_r17                           `optional,ext-9`
	slotBasedDynamicPUCCH_Rep_r17                              *Phy_ParametersCommon_slotBasedDynamicPUCCH_Rep_r17                              `optional,ext-9`
	sps_HARQ_ACK_Deferral_r17                                  *Phy_ParametersCommon_sps_HARQ_ACK_Deferral_r17                                  `optional,ext-9`
	unifiedJointTCI_commonUpdate_r17                           *int64                                                                           `lb:1,ub:4,optional,ext-9`
	mTRP_PDCCH_singleSpan_r17                                  *Phy_ParametersCommon_mTRP_PDCCH_singleSpan_r17                                  `optional,ext-9`
	supportedActivatedPRS_ProcessingWindow_r17                 *Phy_ParametersCommon_supportedActivatedPRS_ProcessingWindow_r17                 `optional,ext-9`
	cg_TimeDomainAllocationExtension_r17                       *Phy_ParametersCommon_cg_TimeDomainAllocationExtension_r17                       `optional,ext-9`
	ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17               *Phy_ParametersCommon_ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17               `optional,ext-10`
	directionalCollisionDC_IAB_r17                             *Phy_ParametersCommon_directionalCollisionDC_IAB_r17                             `optional,ext-10`
	priorityIndicatorInDCI_Multicast_r17                       *Phy_ParametersCommon_priorityIndicatorInDCI_Multicast_r17                       `optional,ext-11`
	priorityIndicatorInDCI_SPS_Multicast_r17                   *Phy_ParametersCommon_priorityIndicatorInDCI_SPS_Multicast_r17                   `optional,ext-11`
	twoHARQ_ACK_CodebookForUnicastAndMulticast_r17             *Phy_ParametersCommon_twoHARQ_ACK_CodebookForUnicastAndMulticast_r17             `optional,ext-11`
	multiPUCCH_HARQ_ACK_ForMulticastUnicast_r17                *Phy_ParametersCommon_multiPUCCH_HARQ_ACK_ForMulticastUnicast_r17                `optional,ext-11`
	srs_AdditionalRepetition_r17                               *Phy_ParametersCommon_srs_AdditionalRepetition_r17                               `optional,ext-11`
	pusch_Repetition_CG_SDT_r17                                *Phy_ParametersCommon_pusch_Repetition_CG_SDT_r17                                `optional,ext-11`
}

func (ie *Phy_ParametersCommon) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.dummy != nil || ie.maxNumberSearchSpaces != nil || ie.rateMatchingCtrlResrcSetDynamic != nil || ie.maxLayersMIMO_Indication != nil || ie.spCellPlacement != nil || ie.twoStepRACH_r16 != nil || ie.dci_Format1_2And0_2_r16 != nil || ie.monitoringDCI_SameSearchSpace_r16 != nil || ie.type2_CG_ReleaseDCI_0_1_r16 != nil || ie.type2_CG_ReleaseDCI_0_2_r16 != nil || ie.sps_ReleaseDCI_1_1_r16 != nil || ie.sps_ReleaseDCI_1_2_r16 != nil || ie.csi_TriggerStateNon_ActiveBWP_r16 != nil || ie.separateSMTC_InterIAB_Support_r16 != nil || ie.separateRACH_IAB_Support_r16 != nil || ie.ul_flexibleDL_SlotFormatSemiStatic_IAB_r16 != nil || ie.ul_flexibleDL_SlotFormatDynamics_IAB_r16 != nil || ie.dft_S_OFDM_WaveformUL_IAB_r16 != nil || ie.dci_25_AI_RNTI_Support_IAB_r16 != nil || ie.t_DeltaReceptionSupport_IAB_r16 != nil || ie.guardSymbolReportReception_IAB_r16 != nil || ie.harqACK_CB_SpatialBundlingPUCCH_Group_r16 != nil || ie.crossSlotScheduling_r16 != nil || ie.maxNumberSRS_PosPathLossEstimateAllServingCells_r16 != nil || ie.extendedCG_Periodicities_r16 != nil || ie.extendedSPS_Periodicities_r16 != nil || ie.codebookVariantsList_r16 != nil || ie.pusch_RepetitionTypeA_r16 != nil || ie.dci_DL_PriorityIndicator_r16 != nil || ie.dci_UL_PriorityIndicator_r16 != nil || ie.maxNumberPathlossRS_Update_r16 != nil || ie.type2_HARQ_ACK_Codebook_r16 != nil || ie.maxTotalResourcesForAcrossFreqRanges_r16 != nil || ie.harqACK_separateMultiDCI_MultiTRP_r16 != nil || ie.harqACK_jointMultiDCI_MultiTRP_r16 != nil || ie.bwp_SwitchingMultiCCs_r16 != nil || ie.targetSMTC_SCG_r16 != nil || ie.supportRepetitionZeroOffsetRV_r16 != nil || ie.cbg_TransInOrderPUSCH_UL_r16 != nil || ie.bwp_SwitchingMultiDormancyCCs_r16 != nil || ie.supportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16 != nil || ie.pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 != nil || ie.newBeamIdentifications2PortCSI_RS_r16 != nil || ie.pathlossEstimation2PortCSI_RS_r16 != nil || ie.mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16 != nil || ie.guardSymbolReportReception_IAB_r17 != nil || ie.restricted_IAB_DU_BeamReception_r17 != nil || ie.recommended_IAB_MT_BeamTransmission_r17 != nil || ie.case6_TimingAlignmentReception_IAB_r17 != nil || ie.case7_TimingAlignmentReception_IAB_r17 != nil || ie.dl_tx_PowerAdjustment_IAB_r17 != nil || ie.desired_ul_tx_PowerAdjustment_r17 != nil || ie.fdm_SoftResourceAvailability_DynamicIndication_r17 != nil || ie.updated_T_DeltaRangeRecption_r17 != nil || ie.slotBasedDynamicPUCCH_Rep_r17 != nil || ie.sps_HARQ_ACK_Deferral_r17 != nil || ie.unifiedJointTCI_commonUpdate_r17 != nil || ie.mTRP_PDCCH_singleSpan_r17 != nil || ie.supportedActivatedPRS_ProcessingWindow_r17 != nil || ie.cg_TimeDomainAllocationExtension_r17 != nil || ie.ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17 != nil || ie.directionalCollisionDC_IAB_r17 != nil || ie.priorityIndicatorInDCI_Multicast_r17 != nil || ie.priorityIndicatorInDCI_SPS_Multicast_r17 != nil || ie.twoHARQ_ACK_CodebookForUnicastAndMulticast_r17 != nil || ie.multiPUCCH_HARQ_ACK_ForMulticastUnicast_r17 != nil || ie.srs_AdditionalRepetition_r17 != nil || ie.pusch_Repetition_CG_SDT_r17 != nil
	preambleBits := []bool{hasExtensions, ie.csi_RS_CFRA_ForHO != nil, ie.dynamicPRB_BundlingDL != nil, ie.sp_CSI_ReportPUCCH != nil, ie.sp_CSI_ReportPUSCH != nil, ie.nzp_CSI_RS_IntefMgmt != nil, ie.type2_SP_CSI_Feedback_LongPUCCH != nil, ie.precoderGranularityCORESET != nil, ie.dynamicHARQ_ACK_Codebook != nil, ie.semiStaticHARQ_ACK_Codebook != nil, ie.spatialBundlingHARQ_ACK != nil, ie.dynamicBetaOffsetInd_HARQ_ACK_CSI != nil, ie.pucch_Repetition_F1_3_4 != nil, ie.ra_Type0_PUSCH != nil, ie.dynamicSwitchRA_Type0_1_PDSCH != nil, ie.dynamicSwitchRA_Type0_1_PUSCH != nil, ie.pdsch_MappingTypeA != nil, ie.pdsch_MappingTypeB != nil, ie.interleavingVRB_ToPRB_PDSCH != nil, ie.interSlotFreqHopping_PUSCH != nil, ie.type1_PUSCH_RepetitionMultiSlots != nil, ie.type2_PUSCH_RepetitionMultiSlots != nil, ie.pusch_RepetitionMultiSlots != nil, ie.pdsch_RepetitionMultiSlots != nil, ie.downlinkSPS != nil, ie.configuredUL_GrantType1 != nil, ie.configuredUL_GrantType2 != nil, ie.pre_EmptIndication_DL != nil, ie.cbg_TransIndication_DL != nil, ie.cbg_TransIndication_UL != nil, ie.cbg_FlushIndication_DL != nil, ie.dynamicHARQ_ACK_CodeB_CBG_Retx_DL != nil, ie.rateMatchingResrcSetSemi_Static != nil, ie.rateMatchingResrcSetDynamic != nil, ie.bwp_SwitchingDelay != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.csi_RS_CFRA_ForHO != nil {
		if err = ie.csi_RS_CFRA_ForHO.Encode(w); err != nil {
			return utils.WrapError("Encode csi_RS_CFRA_ForHO", err)
		}
	}
	if ie.dynamicPRB_BundlingDL != nil {
		if err = ie.dynamicPRB_BundlingDL.Encode(w); err != nil {
			return utils.WrapError("Encode dynamicPRB_BundlingDL", err)
		}
	}
	if ie.sp_CSI_ReportPUCCH != nil {
		if err = ie.sp_CSI_ReportPUCCH.Encode(w); err != nil {
			return utils.WrapError("Encode sp_CSI_ReportPUCCH", err)
		}
	}
	if ie.sp_CSI_ReportPUSCH != nil {
		if err = ie.sp_CSI_ReportPUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode sp_CSI_ReportPUSCH", err)
		}
	}
	if ie.nzp_CSI_RS_IntefMgmt != nil {
		if err = ie.nzp_CSI_RS_IntefMgmt.Encode(w); err != nil {
			return utils.WrapError("Encode nzp_CSI_RS_IntefMgmt", err)
		}
	}
	if ie.type2_SP_CSI_Feedback_LongPUCCH != nil {
		if err = ie.type2_SP_CSI_Feedback_LongPUCCH.Encode(w); err != nil {
			return utils.WrapError("Encode type2_SP_CSI_Feedback_LongPUCCH", err)
		}
	}
	if ie.precoderGranularityCORESET != nil {
		if err = ie.precoderGranularityCORESET.Encode(w); err != nil {
			return utils.WrapError("Encode precoderGranularityCORESET", err)
		}
	}
	if ie.dynamicHARQ_ACK_Codebook != nil {
		if err = ie.dynamicHARQ_ACK_Codebook.Encode(w); err != nil {
			return utils.WrapError("Encode dynamicHARQ_ACK_Codebook", err)
		}
	}
	if ie.semiStaticHARQ_ACK_Codebook != nil {
		if err = ie.semiStaticHARQ_ACK_Codebook.Encode(w); err != nil {
			return utils.WrapError("Encode semiStaticHARQ_ACK_Codebook", err)
		}
	}
	if ie.spatialBundlingHARQ_ACK != nil {
		if err = ie.spatialBundlingHARQ_ACK.Encode(w); err != nil {
			return utils.WrapError("Encode spatialBundlingHARQ_ACK", err)
		}
	}
	if ie.dynamicBetaOffsetInd_HARQ_ACK_CSI != nil {
		if err = ie.dynamicBetaOffsetInd_HARQ_ACK_CSI.Encode(w); err != nil {
			return utils.WrapError("Encode dynamicBetaOffsetInd_HARQ_ACK_CSI", err)
		}
	}
	if ie.pucch_Repetition_F1_3_4 != nil {
		if err = ie.pucch_Repetition_F1_3_4.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_Repetition_F1_3_4", err)
		}
	}
	if ie.ra_Type0_PUSCH != nil {
		if err = ie.ra_Type0_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode ra_Type0_PUSCH", err)
		}
	}
	if ie.dynamicSwitchRA_Type0_1_PDSCH != nil {
		if err = ie.dynamicSwitchRA_Type0_1_PDSCH.Encode(w); err != nil {
			return utils.WrapError("Encode dynamicSwitchRA_Type0_1_PDSCH", err)
		}
	}
	if ie.dynamicSwitchRA_Type0_1_PUSCH != nil {
		if err = ie.dynamicSwitchRA_Type0_1_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode dynamicSwitchRA_Type0_1_PUSCH", err)
		}
	}
	if ie.pdsch_MappingTypeA != nil {
		if err = ie.pdsch_MappingTypeA.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_MappingTypeA", err)
		}
	}
	if ie.pdsch_MappingTypeB != nil {
		if err = ie.pdsch_MappingTypeB.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_MappingTypeB", err)
		}
	}
	if ie.interleavingVRB_ToPRB_PDSCH != nil {
		if err = ie.interleavingVRB_ToPRB_PDSCH.Encode(w); err != nil {
			return utils.WrapError("Encode interleavingVRB_ToPRB_PDSCH", err)
		}
	}
	if ie.interSlotFreqHopping_PUSCH != nil {
		if err = ie.interSlotFreqHopping_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode interSlotFreqHopping_PUSCH", err)
		}
	}
	if ie.type1_PUSCH_RepetitionMultiSlots != nil {
		if err = ie.type1_PUSCH_RepetitionMultiSlots.Encode(w); err != nil {
			return utils.WrapError("Encode type1_PUSCH_RepetitionMultiSlots", err)
		}
	}
	if ie.type2_PUSCH_RepetitionMultiSlots != nil {
		if err = ie.type2_PUSCH_RepetitionMultiSlots.Encode(w); err != nil {
			return utils.WrapError("Encode type2_PUSCH_RepetitionMultiSlots", err)
		}
	}
	if ie.pusch_RepetitionMultiSlots != nil {
		if err = ie.pusch_RepetitionMultiSlots.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_RepetitionMultiSlots", err)
		}
	}
	if ie.pdsch_RepetitionMultiSlots != nil {
		if err = ie.pdsch_RepetitionMultiSlots.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_RepetitionMultiSlots", err)
		}
	}
	if ie.downlinkSPS != nil {
		if err = ie.downlinkSPS.Encode(w); err != nil {
			return utils.WrapError("Encode downlinkSPS", err)
		}
	}
	if ie.configuredUL_GrantType1 != nil {
		if err = ie.configuredUL_GrantType1.Encode(w); err != nil {
			return utils.WrapError("Encode configuredUL_GrantType1", err)
		}
	}
	if ie.configuredUL_GrantType2 != nil {
		if err = ie.configuredUL_GrantType2.Encode(w); err != nil {
			return utils.WrapError("Encode configuredUL_GrantType2", err)
		}
	}
	if ie.pre_EmptIndication_DL != nil {
		if err = ie.pre_EmptIndication_DL.Encode(w); err != nil {
			return utils.WrapError("Encode pre_EmptIndication_DL", err)
		}
	}
	if ie.cbg_TransIndication_DL != nil {
		if err = ie.cbg_TransIndication_DL.Encode(w); err != nil {
			return utils.WrapError("Encode cbg_TransIndication_DL", err)
		}
	}
	if ie.cbg_TransIndication_UL != nil {
		if err = ie.cbg_TransIndication_UL.Encode(w); err != nil {
			return utils.WrapError("Encode cbg_TransIndication_UL", err)
		}
	}
	if ie.cbg_FlushIndication_DL != nil {
		if err = ie.cbg_FlushIndication_DL.Encode(w); err != nil {
			return utils.WrapError("Encode cbg_FlushIndication_DL", err)
		}
	}
	if ie.dynamicHARQ_ACK_CodeB_CBG_Retx_DL != nil {
		if err = ie.dynamicHARQ_ACK_CodeB_CBG_Retx_DL.Encode(w); err != nil {
			return utils.WrapError("Encode dynamicHARQ_ACK_CodeB_CBG_Retx_DL", err)
		}
	}
	if ie.rateMatchingResrcSetSemi_Static != nil {
		if err = ie.rateMatchingResrcSetSemi_Static.Encode(w); err != nil {
			return utils.WrapError("Encode rateMatchingResrcSetSemi_Static", err)
		}
	}
	if ie.rateMatchingResrcSetDynamic != nil {
		if err = ie.rateMatchingResrcSetDynamic.Encode(w); err != nil {
			return utils.WrapError("Encode rateMatchingResrcSetDynamic", err)
		}
	}
	if ie.bwp_SwitchingDelay != nil {
		if err = ie.bwp_SwitchingDelay.Encode(w); err != nil {
			return utils.WrapError("Encode bwp_SwitchingDelay", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 11 bits for 11 extension groups
		extBitmap := []bool{ie.dummy != nil, ie.maxNumberSearchSpaces != nil || ie.rateMatchingCtrlResrcSetDynamic != nil || ie.maxLayersMIMO_Indication != nil, ie.spCellPlacement != nil, ie.twoStepRACH_r16 != nil || ie.dci_Format1_2And0_2_r16 != nil || ie.monitoringDCI_SameSearchSpace_r16 != nil || ie.type2_CG_ReleaseDCI_0_1_r16 != nil || ie.type2_CG_ReleaseDCI_0_2_r16 != nil || ie.sps_ReleaseDCI_1_1_r16 != nil || ie.sps_ReleaseDCI_1_2_r16 != nil || ie.csi_TriggerStateNon_ActiveBWP_r16 != nil || ie.separateSMTC_InterIAB_Support_r16 != nil || ie.separateRACH_IAB_Support_r16 != nil || ie.ul_flexibleDL_SlotFormatSemiStatic_IAB_r16 != nil || ie.ul_flexibleDL_SlotFormatDynamics_IAB_r16 != nil || ie.dft_S_OFDM_WaveformUL_IAB_r16 != nil || ie.dci_25_AI_RNTI_Support_IAB_r16 != nil || ie.t_DeltaReceptionSupport_IAB_r16 != nil || ie.guardSymbolReportReception_IAB_r16 != nil || ie.harqACK_CB_SpatialBundlingPUCCH_Group_r16 != nil || ie.crossSlotScheduling_r16 != nil || ie.maxNumberSRS_PosPathLossEstimateAllServingCells_r16 != nil || ie.extendedCG_Periodicities_r16 != nil || ie.extendedSPS_Periodicities_r16 != nil || ie.codebookVariantsList_r16 != nil || ie.pusch_RepetitionTypeA_r16 != nil || ie.dci_DL_PriorityIndicator_r16 != nil || ie.dci_UL_PriorityIndicator_r16 != nil || ie.maxNumberPathlossRS_Update_r16 != nil || ie.type2_HARQ_ACK_Codebook_r16 != nil || ie.maxTotalResourcesForAcrossFreqRanges_r16 != nil || ie.harqACK_separateMultiDCI_MultiTRP_r16 != nil || ie.harqACK_jointMultiDCI_MultiTRP_r16 != nil || ie.bwp_SwitchingMultiCCs_r16 != nil, ie.targetSMTC_SCG_r16 != nil || ie.supportRepetitionZeroOffsetRV_r16 != nil || ie.cbg_TransInOrderPUSCH_UL_r16 != nil, ie.bwp_SwitchingMultiDormancyCCs_r16 != nil || ie.supportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16 != nil || ie.pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 != nil, ie.newBeamIdentifications2PortCSI_RS_r16 != nil || ie.pathlossEstimation2PortCSI_RS_r16 != nil, ie.mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16 != nil, ie.guardSymbolReportReception_IAB_r17 != nil || ie.restricted_IAB_DU_BeamReception_r17 != nil || ie.recommended_IAB_MT_BeamTransmission_r17 != nil || ie.case6_TimingAlignmentReception_IAB_r17 != nil || ie.case7_TimingAlignmentReception_IAB_r17 != nil || ie.dl_tx_PowerAdjustment_IAB_r17 != nil || ie.desired_ul_tx_PowerAdjustment_r17 != nil || ie.fdm_SoftResourceAvailability_DynamicIndication_r17 != nil || ie.updated_T_DeltaRangeRecption_r17 != nil || ie.slotBasedDynamicPUCCH_Rep_r17 != nil || ie.sps_HARQ_ACK_Deferral_r17 != nil || ie.unifiedJointTCI_commonUpdate_r17 != nil || ie.mTRP_PDCCH_singleSpan_r17 != nil || ie.supportedActivatedPRS_ProcessingWindow_r17 != nil || ie.cg_TimeDomainAllocationExtension_r17 != nil, ie.ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17 != nil || ie.directionalCollisionDC_IAB_r17 != nil, ie.priorityIndicatorInDCI_Multicast_r17 != nil || ie.priorityIndicatorInDCI_SPS_Multicast_r17 != nil || ie.twoHARQ_ACK_CodebookForUnicastAndMulticast_r17 != nil || ie.multiPUCCH_HARQ_ACK_ForMulticastUnicast_r17 != nil || ie.srs_AdditionalRepetition_r17 != nil || ie.pusch_Repetition_CG_SDT_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap Phy_ParametersCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.dummy != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode dummy optional
			if ie.dummy != nil {
				if err = ie.dummy.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dummy", err)
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
			optionals_ext_2 := []bool{ie.maxNumberSearchSpaces != nil, ie.rateMatchingCtrlResrcSetDynamic != nil, ie.maxLayersMIMO_Indication != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode maxNumberSearchSpaces optional
			if ie.maxNumberSearchSpaces != nil {
				if err = ie.maxNumberSearchSpaces.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxNumberSearchSpaces", err)
				}
			}
			// encode rateMatchingCtrlResrcSetDynamic optional
			if ie.rateMatchingCtrlResrcSetDynamic != nil {
				if err = ie.rateMatchingCtrlResrcSetDynamic.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rateMatchingCtrlResrcSetDynamic", err)
				}
			}
			// encode maxLayersMIMO_Indication optional
			if ie.maxLayersMIMO_Indication != nil {
				if err = ie.maxLayersMIMO_Indication.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxLayersMIMO_Indication", err)
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
			optionals_ext_3 := []bool{ie.spCellPlacement != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode spCellPlacement optional
			if ie.spCellPlacement != nil {
				if err = ie.spCellPlacement.Encode(extWriter); err != nil {
					return utils.WrapError("Encode spCellPlacement", err)
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
			optionals_ext_4 := []bool{ie.twoStepRACH_r16 != nil, ie.dci_Format1_2And0_2_r16 != nil, ie.monitoringDCI_SameSearchSpace_r16 != nil, ie.type2_CG_ReleaseDCI_0_1_r16 != nil, ie.type2_CG_ReleaseDCI_0_2_r16 != nil, ie.sps_ReleaseDCI_1_1_r16 != nil, ie.sps_ReleaseDCI_1_2_r16 != nil, ie.csi_TriggerStateNon_ActiveBWP_r16 != nil, ie.separateSMTC_InterIAB_Support_r16 != nil, ie.separateRACH_IAB_Support_r16 != nil, ie.ul_flexibleDL_SlotFormatSemiStatic_IAB_r16 != nil, ie.ul_flexibleDL_SlotFormatDynamics_IAB_r16 != nil, ie.dft_S_OFDM_WaveformUL_IAB_r16 != nil, ie.dci_25_AI_RNTI_Support_IAB_r16 != nil, ie.t_DeltaReceptionSupport_IAB_r16 != nil, ie.guardSymbolReportReception_IAB_r16 != nil, ie.harqACK_CB_SpatialBundlingPUCCH_Group_r16 != nil, ie.crossSlotScheduling_r16 != nil, ie.maxNumberSRS_PosPathLossEstimateAllServingCells_r16 != nil, ie.extendedCG_Periodicities_r16 != nil, ie.extendedSPS_Periodicities_r16 != nil, ie.codebookVariantsList_r16 != nil, ie.pusch_RepetitionTypeA_r16 != nil, ie.dci_DL_PriorityIndicator_r16 != nil, ie.dci_UL_PriorityIndicator_r16 != nil, ie.maxNumberPathlossRS_Update_r16 != nil, ie.type2_HARQ_ACK_Codebook_r16 != nil, ie.maxTotalResourcesForAcrossFreqRanges_r16 != nil, ie.harqACK_separateMultiDCI_MultiTRP_r16 != nil, ie.harqACK_jointMultiDCI_MultiTRP_r16 != nil, ie.bwp_SwitchingMultiCCs_r16 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode twoStepRACH_r16 optional
			if ie.twoStepRACH_r16 != nil {
				if err = ie.twoStepRACH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode twoStepRACH_r16", err)
				}
			}
			// encode dci_Format1_2And0_2_r16 optional
			if ie.dci_Format1_2And0_2_r16 != nil {
				if err = ie.dci_Format1_2And0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dci_Format1_2And0_2_r16", err)
				}
			}
			// encode monitoringDCI_SameSearchSpace_r16 optional
			if ie.monitoringDCI_SameSearchSpace_r16 != nil {
				if err = ie.monitoringDCI_SameSearchSpace_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode monitoringDCI_SameSearchSpace_r16", err)
				}
			}
			// encode type2_CG_ReleaseDCI_0_1_r16 optional
			if ie.type2_CG_ReleaseDCI_0_1_r16 != nil {
				if err = ie.type2_CG_ReleaseDCI_0_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode type2_CG_ReleaseDCI_0_1_r16", err)
				}
			}
			// encode type2_CG_ReleaseDCI_0_2_r16 optional
			if ie.type2_CG_ReleaseDCI_0_2_r16 != nil {
				if err = ie.type2_CG_ReleaseDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode type2_CG_ReleaseDCI_0_2_r16", err)
				}
			}
			// encode sps_ReleaseDCI_1_1_r16 optional
			if ie.sps_ReleaseDCI_1_1_r16 != nil {
				if err = ie.sps_ReleaseDCI_1_1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sps_ReleaseDCI_1_1_r16", err)
				}
			}
			// encode sps_ReleaseDCI_1_2_r16 optional
			if ie.sps_ReleaseDCI_1_2_r16 != nil {
				if err = ie.sps_ReleaseDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sps_ReleaseDCI_1_2_r16", err)
				}
			}
			// encode csi_TriggerStateNon_ActiveBWP_r16 optional
			if ie.csi_TriggerStateNon_ActiveBWP_r16 != nil {
				if err = ie.csi_TriggerStateNon_ActiveBWP_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode csi_TriggerStateNon_ActiveBWP_r16", err)
				}
			}
			// encode separateSMTC_InterIAB_Support_r16 optional
			if ie.separateSMTC_InterIAB_Support_r16 != nil {
				if err = ie.separateSMTC_InterIAB_Support_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode separateSMTC_InterIAB_Support_r16", err)
				}
			}
			// encode separateRACH_IAB_Support_r16 optional
			if ie.separateRACH_IAB_Support_r16 != nil {
				if err = ie.separateRACH_IAB_Support_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode separateRACH_IAB_Support_r16", err)
				}
			}
			// encode ul_flexibleDL_SlotFormatSemiStatic_IAB_r16 optional
			if ie.ul_flexibleDL_SlotFormatSemiStatic_IAB_r16 != nil {
				if err = ie.ul_flexibleDL_SlotFormatSemiStatic_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_flexibleDL_SlotFormatSemiStatic_IAB_r16", err)
				}
			}
			// encode ul_flexibleDL_SlotFormatDynamics_IAB_r16 optional
			if ie.ul_flexibleDL_SlotFormatDynamics_IAB_r16 != nil {
				if err = ie.ul_flexibleDL_SlotFormatDynamics_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_flexibleDL_SlotFormatDynamics_IAB_r16", err)
				}
			}
			// encode dft_S_OFDM_WaveformUL_IAB_r16 optional
			if ie.dft_S_OFDM_WaveformUL_IAB_r16 != nil {
				if err = ie.dft_S_OFDM_WaveformUL_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dft_S_OFDM_WaveformUL_IAB_r16", err)
				}
			}
			// encode dci_25_AI_RNTI_Support_IAB_r16 optional
			if ie.dci_25_AI_RNTI_Support_IAB_r16 != nil {
				if err = ie.dci_25_AI_RNTI_Support_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dci_25_AI_RNTI_Support_IAB_r16", err)
				}
			}
			// encode t_DeltaReceptionSupport_IAB_r16 optional
			if ie.t_DeltaReceptionSupport_IAB_r16 != nil {
				if err = ie.t_DeltaReceptionSupport_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode t_DeltaReceptionSupport_IAB_r16", err)
				}
			}
			// encode guardSymbolReportReception_IAB_r16 optional
			if ie.guardSymbolReportReception_IAB_r16 != nil {
				if err = ie.guardSymbolReportReception_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode guardSymbolReportReception_IAB_r16", err)
				}
			}
			// encode harqACK_CB_SpatialBundlingPUCCH_Group_r16 optional
			if ie.harqACK_CB_SpatialBundlingPUCCH_Group_r16 != nil {
				if err = ie.harqACK_CB_SpatialBundlingPUCCH_Group_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode harqACK_CB_SpatialBundlingPUCCH_Group_r16", err)
				}
			}
			// encode crossSlotScheduling_r16 optional
			if ie.crossSlotScheduling_r16 != nil {
				if err = ie.crossSlotScheduling_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode crossSlotScheduling_r16", err)
				}
			}
			// encode maxNumberSRS_PosPathLossEstimateAllServingCells_r16 optional
			if ie.maxNumberSRS_PosPathLossEstimateAllServingCells_r16 != nil {
				if err = ie.maxNumberSRS_PosPathLossEstimateAllServingCells_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxNumberSRS_PosPathLossEstimateAllServingCells_r16", err)
				}
			}
			// encode extendedCG_Periodicities_r16 optional
			if ie.extendedCG_Periodicities_r16 != nil {
				if err = ie.extendedCG_Periodicities_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode extendedCG_Periodicities_r16", err)
				}
			}
			// encode extendedSPS_Periodicities_r16 optional
			if ie.extendedSPS_Periodicities_r16 != nil {
				if err = ie.extendedSPS_Periodicities_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode extendedSPS_Periodicities_r16", err)
				}
			}
			// encode codebookVariantsList_r16 optional
			if ie.codebookVariantsList_r16 != nil {
				if err = ie.codebookVariantsList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode codebookVariantsList_r16", err)
				}
			}
			// encode pusch_RepetitionTypeA_r16 optional
			if ie.pusch_RepetitionTypeA_r16 != nil {
				if err = ie.pusch_RepetitionTypeA_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pusch_RepetitionTypeA_r16", err)
				}
			}
			// encode dci_DL_PriorityIndicator_r16 optional
			if ie.dci_DL_PriorityIndicator_r16 != nil {
				if err = ie.dci_DL_PriorityIndicator_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dci_DL_PriorityIndicator_r16", err)
				}
			}
			// encode dci_UL_PriorityIndicator_r16 optional
			if ie.dci_UL_PriorityIndicator_r16 != nil {
				if err = ie.dci_UL_PriorityIndicator_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dci_UL_PriorityIndicator_r16", err)
				}
			}
			// encode maxNumberPathlossRS_Update_r16 optional
			if ie.maxNumberPathlossRS_Update_r16 != nil {
				if err = ie.maxNumberPathlossRS_Update_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxNumberPathlossRS_Update_r16", err)
				}
			}
			// encode type2_HARQ_ACK_Codebook_r16 optional
			if ie.type2_HARQ_ACK_Codebook_r16 != nil {
				if err = ie.type2_HARQ_ACK_Codebook_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode type2_HARQ_ACK_Codebook_r16", err)
				}
			}
			// encode maxTotalResourcesForAcrossFreqRanges_r16 optional
			if ie.maxTotalResourcesForAcrossFreqRanges_r16 != nil {
				if err = ie.maxTotalResourcesForAcrossFreqRanges_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxTotalResourcesForAcrossFreqRanges_r16", err)
				}
			}
			// encode harqACK_separateMultiDCI_MultiTRP_r16 optional
			if ie.harqACK_separateMultiDCI_MultiTRP_r16 != nil {
				if err = ie.harqACK_separateMultiDCI_MultiTRP_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode harqACK_separateMultiDCI_MultiTRP_r16", err)
				}
			}
			// encode harqACK_jointMultiDCI_MultiTRP_r16 optional
			if ie.harqACK_jointMultiDCI_MultiTRP_r16 != nil {
				if err = ie.harqACK_jointMultiDCI_MultiTRP_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode harqACK_jointMultiDCI_MultiTRP_r16", err)
				}
			}
			// encode bwp_SwitchingMultiCCs_r16 optional
			if ie.bwp_SwitchingMultiCCs_r16 != nil {
				if err = ie.bwp_SwitchingMultiCCs_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode bwp_SwitchingMultiCCs_r16", err)
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
			optionals_ext_5 := []bool{ie.targetSMTC_SCG_r16 != nil, ie.supportRepetitionZeroOffsetRV_r16 != nil, ie.cbg_TransInOrderPUSCH_UL_r16 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode targetSMTC_SCG_r16 optional
			if ie.targetSMTC_SCG_r16 != nil {
				if err = ie.targetSMTC_SCG_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode targetSMTC_SCG_r16", err)
				}
			}
			// encode supportRepetitionZeroOffsetRV_r16 optional
			if ie.supportRepetitionZeroOffsetRV_r16 != nil {
				if err = ie.supportRepetitionZeroOffsetRV_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportRepetitionZeroOffsetRV_r16", err)
				}
			}
			// encode cbg_TransInOrderPUSCH_UL_r16 optional
			if ie.cbg_TransInOrderPUSCH_UL_r16 != nil {
				if err = ie.cbg_TransInOrderPUSCH_UL_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cbg_TransInOrderPUSCH_UL_r16", err)
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
			optionals_ext_6 := []bool{ie.bwp_SwitchingMultiDormancyCCs_r16 != nil, ie.supportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16 != nil, ie.pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 != nil}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode bwp_SwitchingMultiDormancyCCs_r16 optional
			if ie.bwp_SwitchingMultiDormancyCCs_r16 != nil {
				if err = ie.bwp_SwitchingMultiDormancyCCs_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode bwp_SwitchingMultiDormancyCCs_r16", err)
				}
			}
			// encode supportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16 optional
			if ie.supportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16 != nil {
				if err = ie.supportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16", err)
				}
			}
			// encode pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 optional
			if ie.pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 != nil {
				if err = ie.pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16", err)
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
			optionals_ext_7 := []bool{ie.newBeamIdentifications2PortCSI_RS_r16 != nil, ie.pathlossEstimation2PortCSI_RS_r16 != nil}
			for _, bit := range optionals_ext_7 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode newBeamIdentifications2PortCSI_RS_r16 optional
			if ie.newBeamIdentifications2PortCSI_RS_r16 != nil {
				if err = ie.newBeamIdentifications2PortCSI_RS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode newBeamIdentifications2PortCSI_RS_r16", err)
				}
			}
			// encode pathlossEstimation2PortCSI_RS_r16 optional
			if ie.pathlossEstimation2PortCSI_RS_r16 != nil {
				if err = ie.pathlossEstimation2PortCSI_RS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pathlossEstimation2PortCSI_RS_r16", err)
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
			optionals_ext_8 := []bool{ie.mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16 != nil}
			for _, bit := range optionals_ext_8 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16 optional
			if ie.mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16 != nil {
				if err = ie.mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16", err)
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
			optionals_ext_9 := []bool{ie.guardSymbolReportReception_IAB_r17 != nil, ie.restricted_IAB_DU_BeamReception_r17 != nil, ie.recommended_IAB_MT_BeamTransmission_r17 != nil, ie.case6_TimingAlignmentReception_IAB_r17 != nil, ie.case7_TimingAlignmentReception_IAB_r17 != nil, ie.dl_tx_PowerAdjustment_IAB_r17 != nil, ie.desired_ul_tx_PowerAdjustment_r17 != nil, ie.fdm_SoftResourceAvailability_DynamicIndication_r17 != nil, ie.updated_T_DeltaRangeRecption_r17 != nil, ie.slotBasedDynamicPUCCH_Rep_r17 != nil, ie.sps_HARQ_ACK_Deferral_r17 != nil, ie.unifiedJointTCI_commonUpdate_r17 != nil, ie.mTRP_PDCCH_singleSpan_r17 != nil, ie.supportedActivatedPRS_ProcessingWindow_r17 != nil, ie.cg_TimeDomainAllocationExtension_r17 != nil}
			for _, bit := range optionals_ext_9 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode guardSymbolReportReception_IAB_r17 optional
			if ie.guardSymbolReportReception_IAB_r17 != nil {
				if err = ie.guardSymbolReportReception_IAB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode guardSymbolReportReception_IAB_r17", err)
				}
			}
			// encode restricted_IAB_DU_BeamReception_r17 optional
			if ie.restricted_IAB_DU_BeamReception_r17 != nil {
				if err = ie.restricted_IAB_DU_BeamReception_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode restricted_IAB_DU_BeamReception_r17", err)
				}
			}
			// encode recommended_IAB_MT_BeamTransmission_r17 optional
			if ie.recommended_IAB_MT_BeamTransmission_r17 != nil {
				if err = ie.recommended_IAB_MT_BeamTransmission_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode recommended_IAB_MT_BeamTransmission_r17", err)
				}
			}
			// encode case6_TimingAlignmentReception_IAB_r17 optional
			if ie.case6_TimingAlignmentReception_IAB_r17 != nil {
				if err = ie.case6_TimingAlignmentReception_IAB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode case6_TimingAlignmentReception_IAB_r17", err)
				}
			}
			// encode case7_TimingAlignmentReception_IAB_r17 optional
			if ie.case7_TimingAlignmentReception_IAB_r17 != nil {
				if err = ie.case7_TimingAlignmentReception_IAB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode case7_TimingAlignmentReception_IAB_r17", err)
				}
			}
			// encode dl_tx_PowerAdjustment_IAB_r17 optional
			if ie.dl_tx_PowerAdjustment_IAB_r17 != nil {
				if err = ie.dl_tx_PowerAdjustment_IAB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dl_tx_PowerAdjustment_IAB_r17", err)
				}
			}
			// encode desired_ul_tx_PowerAdjustment_r17 optional
			if ie.desired_ul_tx_PowerAdjustment_r17 != nil {
				if err = ie.desired_ul_tx_PowerAdjustment_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode desired_ul_tx_PowerAdjustment_r17", err)
				}
			}
			// encode fdm_SoftResourceAvailability_DynamicIndication_r17 optional
			if ie.fdm_SoftResourceAvailability_DynamicIndication_r17 != nil {
				if err = ie.fdm_SoftResourceAvailability_DynamicIndication_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode fdm_SoftResourceAvailability_DynamicIndication_r17", err)
				}
			}
			// encode updated_T_DeltaRangeRecption_r17 optional
			if ie.updated_T_DeltaRangeRecption_r17 != nil {
				if err = ie.updated_T_DeltaRangeRecption_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode updated_T_DeltaRangeRecption_r17", err)
				}
			}
			// encode slotBasedDynamicPUCCH_Rep_r17 optional
			if ie.slotBasedDynamicPUCCH_Rep_r17 != nil {
				if err = ie.slotBasedDynamicPUCCH_Rep_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode slotBasedDynamicPUCCH_Rep_r17", err)
				}
			}
			// encode sps_HARQ_ACK_Deferral_r17 optional
			if ie.sps_HARQ_ACK_Deferral_r17 != nil {
				if err = ie.sps_HARQ_ACK_Deferral_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sps_HARQ_ACK_Deferral_r17", err)
				}
			}
			// encode unifiedJointTCI_commonUpdate_r17 optional
			if ie.unifiedJointTCI_commonUpdate_r17 != nil {
				if err = extWriter.WriteInteger(*ie.unifiedJointTCI_commonUpdate_r17, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
					return utils.WrapError("Encode unifiedJointTCI_commonUpdate_r17", err)
				}
			}
			// encode mTRP_PDCCH_singleSpan_r17 optional
			if ie.mTRP_PDCCH_singleSpan_r17 != nil {
				if err = ie.mTRP_PDCCH_singleSpan_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mTRP_PDCCH_singleSpan_r17", err)
				}
			}
			// encode supportedActivatedPRS_ProcessingWindow_r17 optional
			if ie.supportedActivatedPRS_ProcessingWindow_r17 != nil {
				if err = ie.supportedActivatedPRS_ProcessingWindow_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedActivatedPRS_ProcessingWindow_r17", err)
				}
			}
			// encode cg_TimeDomainAllocationExtension_r17 optional
			if ie.cg_TimeDomainAllocationExtension_r17 != nil {
				if err = ie.cg_TimeDomainAllocationExtension_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cg_TimeDomainAllocationExtension_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 10
		if extBitmap[9] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 10
			optionals_ext_10 := []bool{ie.ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17 != nil, ie.directionalCollisionDC_IAB_r17 != nil}
			for _, bit := range optionals_ext_10 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17 optional
			if ie.ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17 != nil {
				if err = ie.ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17", err)
				}
			}
			// encode directionalCollisionDC_IAB_r17 optional
			if ie.directionalCollisionDC_IAB_r17 != nil {
				if err = ie.directionalCollisionDC_IAB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode directionalCollisionDC_IAB_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 11
		if extBitmap[10] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 11
			optionals_ext_11 := []bool{ie.priorityIndicatorInDCI_Multicast_r17 != nil, ie.priorityIndicatorInDCI_SPS_Multicast_r17 != nil, ie.twoHARQ_ACK_CodebookForUnicastAndMulticast_r17 != nil, ie.multiPUCCH_HARQ_ACK_ForMulticastUnicast_r17 != nil, ie.srs_AdditionalRepetition_r17 != nil, ie.pusch_Repetition_CG_SDT_r17 != nil}
			for _, bit := range optionals_ext_11 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode priorityIndicatorInDCI_Multicast_r17 optional
			if ie.priorityIndicatorInDCI_Multicast_r17 != nil {
				if err = ie.priorityIndicatorInDCI_Multicast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode priorityIndicatorInDCI_Multicast_r17", err)
				}
			}
			// encode priorityIndicatorInDCI_SPS_Multicast_r17 optional
			if ie.priorityIndicatorInDCI_SPS_Multicast_r17 != nil {
				if err = ie.priorityIndicatorInDCI_SPS_Multicast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode priorityIndicatorInDCI_SPS_Multicast_r17", err)
				}
			}
			// encode twoHARQ_ACK_CodebookForUnicastAndMulticast_r17 optional
			if ie.twoHARQ_ACK_CodebookForUnicastAndMulticast_r17 != nil {
				if err = ie.twoHARQ_ACK_CodebookForUnicastAndMulticast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode twoHARQ_ACK_CodebookForUnicastAndMulticast_r17", err)
				}
			}
			// encode multiPUCCH_HARQ_ACK_ForMulticastUnicast_r17 optional
			if ie.multiPUCCH_HARQ_ACK_ForMulticastUnicast_r17 != nil {
				if err = ie.multiPUCCH_HARQ_ACK_ForMulticastUnicast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode multiPUCCH_HARQ_ACK_ForMulticastUnicast_r17", err)
				}
			}
			// encode srs_AdditionalRepetition_r17 optional
			if ie.srs_AdditionalRepetition_r17 != nil {
				if err = ie.srs_AdditionalRepetition_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_AdditionalRepetition_r17", err)
				}
			}
			// encode pusch_Repetition_CG_SDT_r17 optional
			if ie.pusch_Repetition_CG_SDT_r17 != nil {
				if err = ie.pusch_Repetition_CG_SDT_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pusch_Repetition_CG_SDT_r17", err)
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

func (ie *Phy_ParametersCommon) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_RS_CFRA_ForHOPresent bool
	if csi_RS_CFRA_ForHOPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dynamicPRB_BundlingDLPresent bool
	if dynamicPRB_BundlingDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sp_CSI_ReportPUCCHPresent bool
	if sp_CSI_ReportPUCCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sp_CSI_ReportPUSCHPresent bool
	if sp_CSI_ReportPUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nzp_CSI_RS_IntefMgmtPresent bool
	if nzp_CSI_RS_IntefMgmtPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var type2_SP_CSI_Feedback_LongPUCCHPresent bool
	if type2_SP_CSI_Feedback_LongPUCCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var precoderGranularityCORESETPresent bool
	if precoderGranularityCORESETPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dynamicHARQ_ACK_CodebookPresent bool
	if dynamicHARQ_ACK_CodebookPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var semiStaticHARQ_ACK_CodebookPresent bool
	if semiStaticHARQ_ACK_CodebookPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var spatialBundlingHARQ_ACKPresent bool
	if spatialBundlingHARQ_ACKPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dynamicBetaOffsetInd_HARQ_ACK_CSIPresent bool
	if dynamicBetaOffsetInd_HARQ_ACK_CSIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_Repetition_F1_3_4Present bool
	if pucch_Repetition_F1_3_4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ra_Type0_PUSCHPresent bool
	if ra_Type0_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dynamicSwitchRA_Type0_1_PDSCHPresent bool
	if dynamicSwitchRA_Type0_1_PDSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dynamicSwitchRA_Type0_1_PUSCHPresent bool
	if dynamicSwitchRA_Type0_1_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_MappingTypeAPresent bool
	if pdsch_MappingTypeAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_MappingTypeBPresent bool
	if pdsch_MappingTypeBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var interleavingVRB_ToPRB_PDSCHPresent bool
	if interleavingVRB_ToPRB_PDSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var interSlotFreqHopping_PUSCHPresent bool
	if interSlotFreqHopping_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var type1_PUSCH_RepetitionMultiSlotsPresent bool
	if type1_PUSCH_RepetitionMultiSlotsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var type2_PUSCH_RepetitionMultiSlotsPresent bool
	if type2_PUSCH_RepetitionMultiSlotsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_RepetitionMultiSlotsPresent bool
	if pusch_RepetitionMultiSlotsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_RepetitionMultiSlotsPresent bool
	if pdsch_RepetitionMultiSlotsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var downlinkSPSPresent bool
	if downlinkSPSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var configuredUL_GrantType1Present bool
	if configuredUL_GrantType1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var configuredUL_GrantType2Present bool
	if configuredUL_GrantType2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pre_EmptIndication_DLPresent bool
	if pre_EmptIndication_DLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var cbg_TransIndication_DLPresent bool
	if cbg_TransIndication_DLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var cbg_TransIndication_ULPresent bool
	if cbg_TransIndication_ULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var cbg_FlushIndication_DLPresent bool
	if cbg_FlushIndication_DLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dynamicHARQ_ACK_CodeB_CBG_Retx_DLPresent bool
	if dynamicHARQ_ACK_CodeB_CBG_Retx_DLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rateMatchingResrcSetSemi_StaticPresent bool
	if rateMatchingResrcSetSemi_StaticPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rateMatchingResrcSetDynamicPresent bool
	if rateMatchingResrcSetDynamicPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var bwp_SwitchingDelayPresent bool
	if bwp_SwitchingDelayPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if csi_RS_CFRA_ForHOPresent {
		ie.csi_RS_CFRA_ForHO = new(Phy_ParametersCommon_csi_RS_CFRA_ForHO)
		if err = ie.csi_RS_CFRA_ForHO.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RS_CFRA_ForHO", err)
		}
	}
	if dynamicPRB_BundlingDLPresent {
		ie.dynamicPRB_BundlingDL = new(Phy_ParametersCommon_dynamicPRB_BundlingDL)
		if err = ie.dynamicPRB_BundlingDL.Decode(r); err != nil {
			return utils.WrapError("Decode dynamicPRB_BundlingDL", err)
		}
	}
	if sp_CSI_ReportPUCCHPresent {
		ie.sp_CSI_ReportPUCCH = new(Phy_ParametersCommon_sp_CSI_ReportPUCCH)
		if err = ie.sp_CSI_ReportPUCCH.Decode(r); err != nil {
			return utils.WrapError("Decode sp_CSI_ReportPUCCH", err)
		}
	}
	if sp_CSI_ReportPUSCHPresent {
		ie.sp_CSI_ReportPUSCH = new(Phy_ParametersCommon_sp_CSI_ReportPUSCH)
		if err = ie.sp_CSI_ReportPUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode sp_CSI_ReportPUSCH", err)
		}
	}
	if nzp_CSI_RS_IntefMgmtPresent {
		ie.nzp_CSI_RS_IntefMgmt = new(Phy_ParametersCommon_nzp_CSI_RS_IntefMgmt)
		if err = ie.nzp_CSI_RS_IntefMgmt.Decode(r); err != nil {
			return utils.WrapError("Decode nzp_CSI_RS_IntefMgmt", err)
		}
	}
	if type2_SP_CSI_Feedback_LongPUCCHPresent {
		ie.type2_SP_CSI_Feedback_LongPUCCH = new(Phy_ParametersCommon_type2_SP_CSI_Feedback_LongPUCCH)
		if err = ie.type2_SP_CSI_Feedback_LongPUCCH.Decode(r); err != nil {
			return utils.WrapError("Decode type2_SP_CSI_Feedback_LongPUCCH", err)
		}
	}
	if precoderGranularityCORESETPresent {
		ie.precoderGranularityCORESET = new(Phy_ParametersCommon_precoderGranularityCORESET)
		if err = ie.precoderGranularityCORESET.Decode(r); err != nil {
			return utils.WrapError("Decode precoderGranularityCORESET", err)
		}
	}
	if dynamicHARQ_ACK_CodebookPresent {
		ie.dynamicHARQ_ACK_Codebook = new(Phy_ParametersCommon_dynamicHARQ_ACK_Codebook)
		if err = ie.dynamicHARQ_ACK_Codebook.Decode(r); err != nil {
			return utils.WrapError("Decode dynamicHARQ_ACK_Codebook", err)
		}
	}
	if semiStaticHARQ_ACK_CodebookPresent {
		ie.semiStaticHARQ_ACK_Codebook = new(Phy_ParametersCommon_semiStaticHARQ_ACK_Codebook)
		if err = ie.semiStaticHARQ_ACK_Codebook.Decode(r); err != nil {
			return utils.WrapError("Decode semiStaticHARQ_ACK_Codebook", err)
		}
	}
	if spatialBundlingHARQ_ACKPresent {
		ie.spatialBundlingHARQ_ACK = new(Phy_ParametersCommon_spatialBundlingHARQ_ACK)
		if err = ie.spatialBundlingHARQ_ACK.Decode(r); err != nil {
			return utils.WrapError("Decode spatialBundlingHARQ_ACK", err)
		}
	}
	if dynamicBetaOffsetInd_HARQ_ACK_CSIPresent {
		ie.dynamicBetaOffsetInd_HARQ_ACK_CSI = new(Phy_ParametersCommon_dynamicBetaOffsetInd_HARQ_ACK_CSI)
		if err = ie.dynamicBetaOffsetInd_HARQ_ACK_CSI.Decode(r); err != nil {
			return utils.WrapError("Decode dynamicBetaOffsetInd_HARQ_ACK_CSI", err)
		}
	}
	if pucch_Repetition_F1_3_4Present {
		ie.pucch_Repetition_F1_3_4 = new(Phy_ParametersCommon_pucch_Repetition_F1_3_4)
		if err = ie.pucch_Repetition_F1_3_4.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_Repetition_F1_3_4", err)
		}
	}
	if ra_Type0_PUSCHPresent {
		ie.ra_Type0_PUSCH = new(Phy_ParametersCommon_ra_Type0_PUSCH)
		if err = ie.ra_Type0_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode ra_Type0_PUSCH", err)
		}
	}
	if dynamicSwitchRA_Type0_1_PDSCHPresent {
		ie.dynamicSwitchRA_Type0_1_PDSCH = new(Phy_ParametersCommon_dynamicSwitchRA_Type0_1_PDSCH)
		if err = ie.dynamicSwitchRA_Type0_1_PDSCH.Decode(r); err != nil {
			return utils.WrapError("Decode dynamicSwitchRA_Type0_1_PDSCH", err)
		}
	}
	if dynamicSwitchRA_Type0_1_PUSCHPresent {
		ie.dynamicSwitchRA_Type0_1_PUSCH = new(Phy_ParametersCommon_dynamicSwitchRA_Type0_1_PUSCH)
		if err = ie.dynamicSwitchRA_Type0_1_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode dynamicSwitchRA_Type0_1_PUSCH", err)
		}
	}
	if pdsch_MappingTypeAPresent {
		ie.pdsch_MappingTypeA = new(Phy_ParametersCommon_pdsch_MappingTypeA)
		if err = ie.pdsch_MappingTypeA.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_MappingTypeA", err)
		}
	}
	if pdsch_MappingTypeBPresent {
		ie.pdsch_MappingTypeB = new(Phy_ParametersCommon_pdsch_MappingTypeB)
		if err = ie.pdsch_MappingTypeB.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_MappingTypeB", err)
		}
	}
	if interleavingVRB_ToPRB_PDSCHPresent {
		ie.interleavingVRB_ToPRB_PDSCH = new(Phy_ParametersCommon_interleavingVRB_ToPRB_PDSCH)
		if err = ie.interleavingVRB_ToPRB_PDSCH.Decode(r); err != nil {
			return utils.WrapError("Decode interleavingVRB_ToPRB_PDSCH", err)
		}
	}
	if interSlotFreqHopping_PUSCHPresent {
		ie.interSlotFreqHopping_PUSCH = new(Phy_ParametersCommon_interSlotFreqHopping_PUSCH)
		if err = ie.interSlotFreqHopping_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode interSlotFreqHopping_PUSCH", err)
		}
	}
	if type1_PUSCH_RepetitionMultiSlotsPresent {
		ie.type1_PUSCH_RepetitionMultiSlots = new(Phy_ParametersCommon_type1_PUSCH_RepetitionMultiSlots)
		if err = ie.type1_PUSCH_RepetitionMultiSlots.Decode(r); err != nil {
			return utils.WrapError("Decode type1_PUSCH_RepetitionMultiSlots", err)
		}
	}
	if type2_PUSCH_RepetitionMultiSlotsPresent {
		ie.type2_PUSCH_RepetitionMultiSlots = new(Phy_ParametersCommon_type2_PUSCH_RepetitionMultiSlots)
		if err = ie.type2_PUSCH_RepetitionMultiSlots.Decode(r); err != nil {
			return utils.WrapError("Decode type2_PUSCH_RepetitionMultiSlots", err)
		}
	}
	if pusch_RepetitionMultiSlotsPresent {
		ie.pusch_RepetitionMultiSlots = new(Phy_ParametersCommon_pusch_RepetitionMultiSlots)
		if err = ie.pusch_RepetitionMultiSlots.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_RepetitionMultiSlots", err)
		}
	}
	if pdsch_RepetitionMultiSlotsPresent {
		ie.pdsch_RepetitionMultiSlots = new(Phy_ParametersCommon_pdsch_RepetitionMultiSlots)
		if err = ie.pdsch_RepetitionMultiSlots.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_RepetitionMultiSlots", err)
		}
	}
	if downlinkSPSPresent {
		ie.downlinkSPS = new(Phy_ParametersCommon_downlinkSPS)
		if err = ie.downlinkSPS.Decode(r); err != nil {
			return utils.WrapError("Decode downlinkSPS", err)
		}
	}
	if configuredUL_GrantType1Present {
		ie.configuredUL_GrantType1 = new(Phy_ParametersCommon_configuredUL_GrantType1)
		if err = ie.configuredUL_GrantType1.Decode(r); err != nil {
			return utils.WrapError("Decode configuredUL_GrantType1", err)
		}
	}
	if configuredUL_GrantType2Present {
		ie.configuredUL_GrantType2 = new(Phy_ParametersCommon_configuredUL_GrantType2)
		if err = ie.configuredUL_GrantType2.Decode(r); err != nil {
			return utils.WrapError("Decode configuredUL_GrantType2", err)
		}
	}
	if pre_EmptIndication_DLPresent {
		ie.pre_EmptIndication_DL = new(Phy_ParametersCommon_pre_EmptIndication_DL)
		if err = ie.pre_EmptIndication_DL.Decode(r); err != nil {
			return utils.WrapError("Decode pre_EmptIndication_DL", err)
		}
	}
	if cbg_TransIndication_DLPresent {
		ie.cbg_TransIndication_DL = new(Phy_ParametersCommon_cbg_TransIndication_DL)
		if err = ie.cbg_TransIndication_DL.Decode(r); err != nil {
			return utils.WrapError("Decode cbg_TransIndication_DL", err)
		}
	}
	if cbg_TransIndication_ULPresent {
		ie.cbg_TransIndication_UL = new(Phy_ParametersCommon_cbg_TransIndication_UL)
		if err = ie.cbg_TransIndication_UL.Decode(r); err != nil {
			return utils.WrapError("Decode cbg_TransIndication_UL", err)
		}
	}
	if cbg_FlushIndication_DLPresent {
		ie.cbg_FlushIndication_DL = new(Phy_ParametersCommon_cbg_FlushIndication_DL)
		if err = ie.cbg_FlushIndication_DL.Decode(r); err != nil {
			return utils.WrapError("Decode cbg_FlushIndication_DL", err)
		}
	}
	if dynamicHARQ_ACK_CodeB_CBG_Retx_DLPresent {
		ie.dynamicHARQ_ACK_CodeB_CBG_Retx_DL = new(Phy_ParametersCommon_dynamicHARQ_ACK_CodeB_CBG_Retx_DL)
		if err = ie.dynamicHARQ_ACK_CodeB_CBG_Retx_DL.Decode(r); err != nil {
			return utils.WrapError("Decode dynamicHARQ_ACK_CodeB_CBG_Retx_DL", err)
		}
	}
	if rateMatchingResrcSetSemi_StaticPresent {
		ie.rateMatchingResrcSetSemi_Static = new(Phy_ParametersCommon_rateMatchingResrcSetSemi_Static)
		if err = ie.rateMatchingResrcSetSemi_Static.Decode(r); err != nil {
			return utils.WrapError("Decode rateMatchingResrcSetSemi_Static", err)
		}
	}
	if rateMatchingResrcSetDynamicPresent {
		ie.rateMatchingResrcSetDynamic = new(Phy_ParametersCommon_rateMatchingResrcSetDynamic)
		if err = ie.rateMatchingResrcSetDynamic.Decode(r); err != nil {
			return utils.WrapError("Decode rateMatchingResrcSetDynamic", err)
		}
	}
	if bwp_SwitchingDelayPresent {
		ie.bwp_SwitchingDelay = new(Phy_ParametersCommon_bwp_SwitchingDelay)
		if err = ie.bwp_SwitchingDelay.Decode(r); err != nil {
			return utils.WrapError("Decode bwp_SwitchingDelay", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 11 bits for 11 extension groups
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

			dummyPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode dummy optional
			if dummyPresent {
				ie.dummy = new(Phy_ParametersCommon_dummy)
				if err = ie.dummy.Decode(extReader); err != nil {
					return utils.WrapError("Decode dummy", err)
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

			maxNumberSearchSpacesPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			rateMatchingCtrlResrcSetDynamicPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxLayersMIMO_IndicationPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode maxNumberSearchSpaces optional
			if maxNumberSearchSpacesPresent {
				ie.maxNumberSearchSpaces = new(Phy_ParametersCommon_maxNumberSearchSpaces)
				if err = ie.maxNumberSearchSpaces.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxNumberSearchSpaces", err)
				}
			}
			// decode rateMatchingCtrlResrcSetDynamic optional
			if rateMatchingCtrlResrcSetDynamicPresent {
				ie.rateMatchingCtrlResrcSetDynamic = new(Phy_ParametersCommon_rateMatchingCtrlResrcSetDynamic)
				if err = ie.rateMatchingCtrlResrcSetDynamic.Decode(extReader); err != nil {
					return utils.WrapError("Decode rateMatchingCtrlResrcSetDynamic", err)
				}
			}
			// decode maxLayersMIMO_Indication optional
			if maxLayersMIMO_IndicationPresent {
				ie.maxLayersMIMO_Indication = new(Phy_ParametersCommon_maxLayersMIMO_Indication)
				if err = ie.maxLayersMIMO_Indication.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxLayersMIMO_Indication", err)
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

			spCellPlacementPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode spCellPlacement optional
			if spCellPlacementPresent {
				ie.spCellPlacement = new(CarrierAggregationVariant)
				if err = ie.spCellPlacement.Decode(extReader); err != nil {
					return utils.WrapError("Decode spCellPlacement", err)
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

			twoStepRACH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dci_Format1_2And0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			monitoringDCI_SameSearchSpace_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			type2_CG_ReleaseDCI_0_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			type2_CG_ReleaseDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sps_ReleaseDCI_1_1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sps_ReleaseDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			csi_TriggerStateNon_ActiveBWP_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			separateSMTC_InterIAB_Support_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			separateRACH_IAB_Support_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_flexibleDL_SlotFormatSemiStatic_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_flexibleDL_SlotFormatDynamics_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dft_S_OFDM_WaveformUL_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dci_25_AI_RNTI_Support_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			t_DeltaReceptionSupport_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			guardSymbolReportReception_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			harqACK_CB_SpatialBundlingPUCCH_Group_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			crossSlotScheduling_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumberSRS_PosPathLossEstimateAllServingCells_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			extendedCG_Periodicities_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			extendedSPS_Periodicities_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			codebookVariantsList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pusch_RepetitionTypeA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dci_DL_PriorityIndicator_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dci_UL_PriorityIndicator_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumberPathlossRS_Update_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			type2_HARQ_ACK_Codebook_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxTotalResourcesForAcrossFreqRanges_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			harqACK_separateMultiDCI_MultiTRP_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			harqACK_jointMultiDCI_MultiTRP_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			bwp_SwitchingMultiCCs_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode twoStepRACH_r16 optional
			if twoStepRACH_r16Present {
				ie.twoStepRACH_r16 = new(Phy_ParametersCommon_twoStepRACH_r16)
				if err = ie.twoStepRACH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode twoStepRACH_r16", err)
				}
			}
			// decode dci_Format1_2And0_2_r16 optional
			if dci_Format1_2And0_2_r16Present {
				ie.dci_Format1_2And0_2_r16 = new(Phy_ParametersCommon_dci_Format1_2And0_2_r16)
				if err = ie.dci_Format1_2And0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode dci_Format1_2And0_2_r16", err)
				}
			}
			// decode monitoringDCI_SameSearchSpace_r16 optional
			if monitoringDCI_SameSearchSpace_r16Present {
				ie.monitoringDCI_SameSearchSpace_r16 = new(Phy_ParametersCommon_monitoringDCI_SameSearchSpace_r16)
				if err = ie.monitoringDCI_SameSearchSpace_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode monitoringDCI_SameSearchSpace_r16", err)
				}
			}
			// decode type2_CG_ReleaseDCI_0_1_r16 optional
			if type2_CG_ReleaseDCI_0_1_r16Present {
				ie.type2_CG_ReleaseDCI_0_1_r16 = new(Phy_ParametersCommon_type2_CG_ReleaseDCI_0_1_r16)
				if err = ie.type2_CG_ReleaseDCI_0_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode type2_CG_ReleaseDCI_0_1_r16", err)
				}
			}
			// decode type2_CG_ReleaseDCI_0_2_r16 optional
			if type2_CG_ReleaseDCI_0_2_r16Present {
				ie.type2_CG_ReleaseDCI_0_2_r16 = new(Phy_ParametersCommon_type2_CG_ReleaseDCI_0_2_r16)
				if err = ie.type2_CG_ReleaseDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode type2_CG_ReleaseDCI_0_2_r16", err)
				}
			}
			// decode sps_ReleaseDCI_1_1_r16 optional
			if sps_ReleaseDCI_1_1_r16Present {
				ie.sps_ReleaseDCI_1_1_r16 = new(Phy_ParametersCommon_sps_ReleaseDCI_1_1_r16)
				if err = ie.sps_ReleaseDCI_1_1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sps_ReleaseDCI_1_1_r16", err)
				}
			}
			// decode sps_ReleaseDCI_1_2_r16 optional
			if sps_ReleaseDCI_1_2_r16Present {
				ie.sps_ReleaseDCI_1_2_r16 = new(Phy_ParametersCommon_sps_ReleaseDCI_1_2_r16)
				if err = ie.sps_ReleaseDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sps_ReleaseDCI_1_2_r16", err)
				}
			}
			// decode csi_TriggerStateNon_ActiveBWP_r16 optional
			if csi_TriggerStateNon_ActiveBWP_r16Present {
				ie.csi_TriggerStateNon_ActiveBWP_r16 = new(Phy_ParametersCommon_csi_TriggerStateNon_ActiveBWP_r16)
				if err = ie.csi_TriggerStateNon_ActiveBWP_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode csi_TriggerStateNon_ActiveBWP_r16", err)
				}
			}
			// decode separateSMTC_InterIAB_Support_r16 optional
			if separateSMTC_InterIAB_Support_r16Present {
				ie.separateSMTC_InterIAB_Support_r16 = new(Phy_ParametersCommon_separateSMTC_InterIAB_Support_r16)
				if err = ie.separateSMTC_InterIAB_Support_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode separateSMTC_InterIAB_Support_r16", err)
				}
			}
			// decode separateRACH_IAB_Support_r16 optional
			if separateRACH_IAB_Support_r16Present {
				ie.separateRACH_IAB_Support_r16 = new(Phy_ParametersCommon_separateRACH_IAB_Support_r16)
				if err = ie.separateRACH_IAB_Support_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode separateRACH_IAB_Support_r16", err)
				}
			}
			// decode ul_flexibleDL_SlotFormatSemiStatic_IAB_r16 optional
			if ul_flexibleDL_SlotFormatSemiStatic_IAB_r16Present {
				ie.ul_flexibleDL_SlotFormatSemiStatic_IAB_r16 = new(Phy_ParametersCommon_ul_flexibleDL_SlotFormatSemiStatic_IAB_r16)
				if err = ie.ul_flexibleDL_SlotFormatSemiStatic_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_flexibleDL_SlotFormatSemiStatic_IAB_r16", err)
				}
			}
			// decode ul_flexibleDL_SlotFormatDynamics_IAB_r16 optional
			if ul_flexibleDL_SlotFormatDynamics_IAB_r16Present {
				ie.ul_flexibleDL_SlotFormatDynamics_IAB_r16 = new(Phy_ParametersCommon_ul_flexibleDL_SlotFormatDynamics_IAB_r16)
				if err = ie.ul_flexibleDL_SlotFormatDynamics_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_flexibleDL_SlotFormatDynamics_IAB_r16", err)
				}
			}
			// decode dft_S_OFDM_WaveformUL_IAB_r16 optional
			if dft_S_OFDM_WaveformUL_IAB_r16Present {
				ie.dft_S_OFDM_WaveformUL_IAB_r16 = new(Phy_ParametersCommon_dft_S_OFDM_WaveformUL_IAB_r16)
				if err = ie.dft_S_OFDM_WaveformUL_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode dft_S_OFDM_WaveformUL_IAB_r16", err)
				}
			}
			// decode dci_25_AI_RNTI_Support_IAB_r16 optional
			if dci_25_AI_RNTI_Support_IAB_r16Present {
				ie.dci_25_AI_RNTI_Support_IAB_r16 = new(Phy_ParametersCommon_dci_25_AI_RNTI_Support_IAB_r16)
				if err = ie.dci_25_AI_RNTI_Support_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode dci_25_AI_RNTI_Support_IAB_r16", err)
				}
			}
			// decode t_DeltaReceptionSupport_IAB_r16 optional
			if t_DeltaReceptionSupport_IAB_r16Present {
				ie.t_DeltaReceptionSupport_IAB_r16 = new(Phy_ParametersCommon_t_DeltaReceptionSupport_IAB_r16)
				if err = ie.t_DeltaReceptionSupport_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode t_DeltaReceptionSupport_IAB_r16", err)
				}
			}
			// decode guardSymbolReportReception_IAB_r16 optional
			if guardSymbolReportReception_IAB_r16Present {
				ie.guardSymbolReportReception_IAB_r16 = new(Phy_ParametersCommon_guardSymbolReportReception_IAB_r16)
				if err = ie.guardSymbolReportReception_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode guardSymbolReportReception_IAB_r16", err)
				}
			}
			// decode harqACK_CB_SpatialBundlingPUCCH_Group_r16 optional
			if harqACK_CB_SpatialBundlingPUCCH_Group_r16Present {
				ie.harqACK_CB_SpatialBundlingPUCCH_Group_r16 = new(Phy_ParametersCommon_harqACK_CB_SpatialBundlingPUCCH_Group_r16)
				if err = ie.harqACK_CB_SpatialBundlingPUCCH_Group_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode harqACK_CB_SpatialBundlingPUCCH_Group_r16", err)
				}
			}
			// decode crossSlotScheduling_r16 optional
			if crossSlotScheduling_r16Present {
				ie.crossSlotScheduling_r16 = new(Phy_ParametersCommon_crossSlotScheduling_r16)
				if err = ie.crossSlotScheduling_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode crossSlotScheduling_r16", err)
				}
			}
			// decode maxNumberSRS_PosPathLossEstimateAllServingCells_r16 optional
			if maxNumberSRS_PosPathLossEstimateAllServingCells_r16Present {
				ie.maxNumberSRS_PosPathLossEstimateAllServingCells_r16 = new(Phy_ParametersCommon_maxNumberSRS_PosPathLossEstimateAllServingCells_r16)
				if err = ie.maxNumberSRS_PosPathLossEstimateAllServingCells_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxNumberSRS_PosPathLossEstimateAllServingCells_r16", err)
				}
			}
			// decode extendedCG_Periodicities_r16 optional
			if extendedCG_Periodicities_r16Present {
				ie.extendedCG_Periodicities_r16 = new(Phy_ParametersCommon_extendedCG_Periodicities_r16)
				if err = ie.extendedCG_Periodicities_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode extendedCG_Periodicities_r16", err)
				}
			}
			// decode extendedSPS_Periodicities_r16 optional
			if extendedSPS_Periodicities_r16Present {
				ie.extendedSPS_Periodicities_r16 = new(Phy_ParametersCommon_extendedSPS_Periodicities_r16)
				if err = ie.extendedSPS_Periodicities_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode extendedSPS_Periodicities_r16", err)
				}
			}
			// decode codebookVariantsList_r16 optional
			if codebookVariantsList_r16Present {
				ie.codebookVariantsList_r16 = new(CodebookVariantsList_r16)
				if err = ie.codebookVariantsList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode codebookVariantsList_r16", err)
				}
			}
			// decode pusch_RepetitionTypeA_r16 optional
			if pusch_RepetitionTypeA_r16Present {
				ie.pusch_RepetitionTypeA_r16 = new(Phy_ParametersCommon_pusch_RepetitionTypeA_r16)
				if err = ie.pusch_RepetitionTypeA_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pusch_RepetitionTypeA_r16", err)
				}
			}
			// decode dci_DL_PriorityIndicator_r16 optional
			if dci_DL_PriorityIndicator_r16Present {
				ie.dci_DL_PriorityIndicator_r16 = new(Phy_ParametersCommon_dci_DL_PriorityIndicator_r16)
				if err = ie.dci_DL_PriorityIndicator_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode dci_DL_PriorityIndicator_r16", err)
				}
			}
			// decode dci_UL_PriorityIndicator_r16 optional
			if dci_UL_PriorityIndicator_r16Present {
				ie.dci_UL_PriorityIndicator_r16 = new(Phy_ParametersCommon_dci_UL_PriorityIndicator_r16)
				if err = ie.dci_UL_PriorityIndicator_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode dci_UL_PriorityIndicator_r16", err)
				}
			}
			// decode maxNumberPathlossRS_Update_r16 optional
			if maxNumberPathlossRS_Update_r16Present {
				ie.maxNumberPathlossRS_Update_r16 = new(Phy_ParametersCommon_maxNumberPathlossRS_Update_r16)
				if err = ie.maxNumberPathlossRS_Update_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxNumberPathlossRS_Update_r16", err)
				}
			}
			// decode type2_HARQ_ACK_Codebook_r16 optional
			if type2_HARQ_ACK_Codebook_r16Present {
				ie.type2_HARQ_ACK_Codebook_r16 = new(Phy_ParametersCommon_type2_HARQ_ACK_Codebook_r16)
				if err = ie.type2_HARQ_ACK_Codebook_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode type2_HARQ_ACK_Codebook_r16", err)
				}
			}
			// decode maxTotalResourcesForAcrossFreqRanges_r16 optional
			if maxTotalResourcesForAcrossFreqRanges_r16Present {
				ie.maxTotalResourcesForAcrossFreqRanges_r16 = new(Phy_ParametersCommon_maxTotalResourcesForAcrossFreqRanges_r16)
				if err = ie.maxTotalResourcesForAcrossFreqRanges_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxTotalResourcesForAcrossFreqRanges_r16", err)
				}
			}
			// decode harqACK_separateMultiDCI_MultiTRP_r16 optional
			if harqACK_separateMultiDCI_MultiTRP_r16Present {
				ie.harqACK_separateMultiDCI_MultiTRP_r16 = new(Phy_ParametersCommon_harqACK_separateMultiDCI_MultiTRP_r16)
				if err = ie.harqACK_separateMultiDCI_MultiTRP_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode harqACK_separateMultiDCI_MultiTRP_r16", err)
				}
			}
			// decode harqACK_jointMultiDCI_MultiTRP_r16 optional
			if harqACK_jointMultiDCI_MultiTRP_r16Present {
				ie.harqACK_jointMultiDCI_MultiTRP_r16 = new(Phy_ParametersCommon_harqACK_jointMultiDCI_MultiTRP_r16)
				if err = ie.harqACK_jointMultiDCI_MultiTRP_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode harqACK_jointMultiDCI_MultiTRP_r16", err)
				}
			}
			// decode bwp_SwitchingMultiCCs_r16 optional
			if bwp_SwitchingMultiCCs_r16Present {
				ie.bwp_SwitchingMultiCCs_r16 = new(Phy_ParametersCommon_bwp_SwitchingMultiCCs_r16)
				if err = ie.bwp_SwitchingMultiCCs_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode bwp_SwitchingMultiCCs_r16", err)
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

			targetSMTC_SCG_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportRepetitionZeroOffsetRV_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cbg_TransInOrderPUSCH_UL_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode targetSMTC_SCG_r16 optional
			if targetSMTC_SCG_r16Present {
				ie.targetSMTC_SCG_r16 = new(Phy_ParametersCommon_targetSMTC_SCG_r16)
				if err = ie.targetSMTC_SCG_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode targetSMTC_SCG_r16", err)
				}
			}
			// decode supportRepetitionZeroOffsetRV_r16 optional
			if supportRepetitionZeroOffsetRV_r16Present {
				ie.supportRepetitionZeroOffsetRV_r16 = new(Phy_ParametersCommon_supportRepetitionZeroOffsetRV_r16)
				if err = ie.supportRepetitionZeroOffsetRV_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportRepetitionZeroOffsetRV_r16", err)
				}
			}
			// decode cbg_TransInOrderPUSCH_UL_r16 optional
			if cbg_TransInOrderPUSCH_UL_r16Present {
				ie.cbg_TransInOrderPUSCH_UL_r16 = new(Phy_ParametersCommon_cbg_TransInOrderPUSCH_UL_r16)
				if err = ie.cbg_TransInOrderPUSCH_UL_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode cbg_TransInOrderPUSCH_UL_r16", err)
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

			bwp_SwitchingMultiDormancyCCs_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode bwp_SwitchingMultiDormancyCCs_r16 optional
			if bwp_SwitchingMultiDormancyCCs_r16Present {
				ie.bwp_SwitchingMultiDormancyCCs_r16 = new(Phy_ParametersCommon_bwp_SwitchingMultiDormancyCCs_r16)
				if err = ie.bwp_SwitchingMultiDormancyCCs_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode bwp_SwitchingMultiDormancyCCs_r16", err)
				}
			}
			// decode supportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16 optional
			if supportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16Present {
				ie.supportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16 = new(Phy_ParametersCommon_supportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16)
				if err = ie.supportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportRetx_Diff_CoresetPool_Multi_DCI_TRP_r16", err)
				}
			}
			// decode pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 optional
			if pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16Present {
				ie.pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16 = new(Phy_ParametersCommon_pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16)
				if err = ie.pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdcch_MonitoringAnyOccasionsWithSpanGapCrossCarrierSch_r16", err)
				}
			}
		}
		// decode extension group 7
		if len(extBitmap) > 6 && extBitmap[6] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			newBeamIdentifications2PortCSI_RS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pathlossEstimation2PortCSI_RS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode newBeamIdentifications2PortCSI_RS_r16 optional
			if newBeamIdentifications2PortCSI_RS_r16Present {
				ie.newBeamIdentifications2PortCSI_RS_r16 = new(Phy_ParametersCommon_newBeamIdentifications2PortCSI_RS_r16)
				if err = ie.newBeamIdentifications2PortCSI_RS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode newBeamIdentifications2PortCSI_RS_r16", err)
				}
			}
			// decode pathlossEstimation2PortCSI_RS_r16 optional
			if pathlossEstimation2PortCSI_RS_r16Present {
				ie.pathlossEstimation2PortCSI_RS_r16 = new(Phy_ParametersCommon_pathlossEstimation2PortCSI_RS_r16)
				if err = ie.pathlossEstimation2PortCSI_RS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pathlossEstimation2PortCSI_RS_r16", err)
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

			mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16 optional
			if mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16Present {
				ie.mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16 = new(Phy_ParametersCommon_mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16)
				if err = ie.mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode mux_HARQ_ACK_withoutPUCCH_onPUSCH_r16", err)
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

			guardSymbolReportReception_IAB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			restricted_IAB_DU_BeamReception_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			recommended_IAB_MT_BeamTransmission_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			case6_TimingAlignmentReception_IAB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			case7_TimingAlignmentReception_IAB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dl_tx_PowerAdjustment_IAB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			desired_ul_tx_PowerAdjustment_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			fdm_SoftResourceAvailability_DynamicIndication_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			updated_T_DeltaRangeRecption_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			slotBasedDynamicPUCCH_Rep_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sps_HARQ_ACK_Deferral_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			unifiedJointTCI_commonUpdate_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mTRP_PDCCH_singleSpan_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedActivatedPRS_ProcessingWindow_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cg_TimeDomainAllocationExtension_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode guardSymbolReportReception_IAB_r17 optional
			if guardSymbolReportReception_IAB_r17Present {
				ie.guardSymbolReportReception_IAB_r17 = new(Phy_ParametersCommon_guardSymbolReportReception_IAB_r17)
				if err = ie.guardSymbolReportReception_IAB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode guardSymbolReportReception_IAB_r17", err)
				}
			}
			// decode restricted_IAB_DU_BeamReception_r17 optional
			if restricted_IAB_DU_BeamReception_r17Present {
				ie.restricted_IAB_DU_BeamReception_r17 = new(Phy_ParametersCommon_restricted_IAB_DU_BeamReception_r17)
				if err = ie.restricted_IAB_DU_BeamReception_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode restricted_IAB_DU_BeamReception_r17", err)
				}
			}
			// decode recommended_IAB_MT_BeamTransmission_r17 optional
			if recommended_IAB_MT_BeamTransmission_r17Present {
				ie.recommended_IAB_MT_BeamTransmission_r17 = new(Phy_ParametersCommon_recommended_IAB_MT_BeamTransmission_r17)
				if err = ie.recommended_IAB_MT_BeamTransmission_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode recommended_IAB_MT_BeamTransmission_r17", err)
				}
			}
			// decode case6_TimingAlignmentReception_IAB_r17 optional
			if case6_TimingAlignmentReception_IAB_r17Present {
				ie.case6_TimingAlignmentReception_IAB_r17 = new(Phy_ParametersCommon_case6_TimingAlignmentReception_IAB_r17)
				if err = ie.case6_TimingAlignmentReception_IAB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode case6_TimingAlignmentReception_IAB_r17", err)
				}
			}
			// decode case7_TimingAlignmentReception_IAB_r17 optional
			if case7_TimingAlignmentReception_IAB_r17Present {
				ie.case7_TimingAlignmentReception_IAB_r17 = new(Phy_ParametersCommon_case7_TimingAlignmentReception_IAB_r17)
				if err = ie.case7_TimingAlignmentReception_IAB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode case7_TimingAlignmentReception_IAB_r17", err)
				}
			}
			// decode dl_tx_PowerAdjustment_IAB_r17 optional
			if dl_tx_PowerAdjustment_IAB_r17Present {
				ie.dl_tx_PowerAdjustment_IAB_r17 = new(Phy_ParametersCommon_dl_tx_PowerAdjustment_IAB_r17)
				if err = ie.dl_tx_PowerAdjustment_IAB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dl_tx_PowerAdjustment_IAB_r17", err)
				}
			}
			// decode desired_ul_tx_PowerAdjustment_r17 optional
			if desired_ul_tx_PowerAdjustment_r17Present {
				ie.desired_ul_tx_PowerAdjustment_r17 = new(Phy_ParametersCommon_desired_ul_tx_PowerAdjustment_r17)
				if err = ie.desired_ul_tx_PowerAdjustment_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode desired_ul_tx_PowerAdjustment_r17", err)
				}
			}
			// decode fdm_SoftResourceAvailability_DynamicIndication_r17 optional
			if fdm_SoftResourceAvailability_DynamicIndication_r17Present {
				ie.fdm_SoftResourceAvailability_DynamicIndication_r17 = new(Phy_ParametersCommon_fdm_SoftResourceAvailability_DynamicIndication_r17)
				if err = ie.fdm_SoftResourceAvailability_DynamicIndication_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode fdm_SoftResourceAvailability_DynamicIndication_r17", err)
				}
			}
			// decode updated_T_DeltaRangeRecption_r17 optional
			if updated_T_DeltaRangeRecption_r17Present {
				ie.updated_T_DeltaRangeRecption_r17 = new(Phy_ParametersCommon_updated_T_DeltaRangeRecption_r17)
				if err = ie.updated_T_DeltaRangeRecption_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode updated_T_DeltaRangeRecption_r17", err)
				}
			}
			// decode slotBasedDynamicPUCCH_Rep_r17 optional
			if slotBasedDynamicPUCCH_Rep_r17Present {
				ie.slotBasedDynamicPUCCH_Rep_r17 = new(Phy_ParametersCommon_slotBasedDynamicPUCCH_Rep_r17)
				if err = ie.slotBasedDynamicPUCCH_Rep_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode slotBasedDynamicPUCCH_Rep_r17", err)
				}
			}
			// decode sps_HARQ_ACK_Deferral_r17 optional
			if sps_HARQ_ACK_Deferral_r17Present {
				ie.sps_HARQ_ACK_Deferral_r17 = new(Phy_ParametersCommon_sps_HARQ_ACK_Deferral_r17)
				if err = ie.sps_HARQ_ACK_Deferral_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sps_HARQ_ACK_Deferral_r17", err)
				}
			}
			// decode unifiedJointTCI_commonUpdate_r17 optional
			if unifiedJointTCI_commonUpdate_r17Present {
				var tmp_int_unifiedJointTCI_commonUpdate_r17 int64
				if tmp_int_unifiedJointTCI_commonUpdate_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
					return utils.WrapError("Decode unifiedJointTCI_commonUpdate_r17", err)
				}
				ie.unifiedJointTCI_commonUpdate_r17 = &tmp_int_unifiedJointTCI_commonUpdate_r17
			}
			// decode mTRP_PDCCH_singleSpan_r17 optional
			if mTRP_PDCCH_singleSpan_r17Present {
				ie.mTRP_PDCCH_singleSpan_r17 = new(Phy_ParametersCommon_mTRP_PDCCH_singleSpan_r17)
				if err = ie.mTRP_PDCCH_singleSpan_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mTRP_PDCCH_singleSpan_r17", err)
				}
			}
			// decode supportedActivatedPRS_ProcessingWindow_r17 optional
			if supportedActivatedPRS_ProcessingWindow_r17Present {
				ie.supportedActivatedPRS_ProcessingWindow_r17 = new(Phy_ParametersCommon_supportedActivatedPRS_ProcessingWindow_r17)
				if err = ie.supportedActivatedPRS_ProcessingWindow_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedActivatedPRS_ProcessingWindow_r17", err)
				}
			}
			// decode cg_TimeDomainAllocationExtension_r17 optional
			if cg_TimeDomainAllocationExtension_r17Present {
				ie.cg_TimeDomainAllocationExtension_r17 = new(Phy_ParametersCommon_cg_TimeDomainAllocationExtension_r17)
				if err = ie.cg_TimeDomainAllocationExtension_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode cg_TimeDomainAllocationExtension_r17", err)
				}
			}
		}
		// decode extension group 10
		if len(extBitmap) > 9 && extBitmap[9] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			directionalCollisionDC_IAB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17 optional
			if ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17Present {
				ie.ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17 = new(Phy_ParametersCommon_ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17)
				if err = ie.ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ta_BasedPDC_TN_NonSharedSpectrumChAccess_r17", err)
				}
			}
			// decode directionalCollisionDC_IAB_r17 optional
			if directionalCollisionDC_IAB_r17Present {
				ie.directionalCollisionDC_IAB_r17 = new(Phy_ParametersCommon_directionalCollisionDC_IAB_r17)
				if err = ie.directionalCollisionDC_IAB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode directionalCollisionDC_IAB_r17", err)
				}
			}
		}
		// decode extension group 11
		if len(extBitmap) > 10 && extBitmap[10] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			priorityIndicatorInDCI_Multicast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			priorityIndicatorInDCI_SPS_Multicast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			twoHARQ_ACK_CodebookForUnicastAndMulticast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			multiPUCCH_HARQ_ACK_ForMulticastUnicast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_AdditionalRepetition_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pusch_Repetition_CG_SDT_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode priorityIndicatorInDCI_Multicast_r17 optional
			if priorityIndicatorInDCI_Multicast_r17Present {
				ie.priorityIndicatorInDCI_Multicast_r17 = new(Phy_ParametersCommon_priorityIndicatorInDCI_Multicast_r17)
				if err = ie.priorityIndicatorInDCI_Multicast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode priorityIndicatorInDCI_Multicast_r17", err)
				}
			}
			// decode priorityIndicatorInDCI_SPS_Multicast_r17 optional
			if priorityIndicatorInDCI_SPS_Multicast_r17Present {
				ie.priorityIndicatorInDCI_SPS_Multicast_r17 = new(Phy_ParametersCommon_priorityIndicatorInDCI_SPS_Multicast_r17)
				if err = ie.priorityIndicatorInDCI_SPS_Multicast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode priorityIndicatorInDCI_SPS_Multicast_r17", err)
				}
			}
			// decode twoHARQ_ACK_CodebookForUnicastAndMulticast_r17 optional
			if twoHARQ_ACK_CodebookForUnicastAndMulticast_r17Present {
				ie.twoHARQ_ACK_CodebookForUnicastAndMulticast_r17 = new(Phy_ParametersCommon_twoHARQ_ACK_CodebookForUnicastAndMulticast_r17)
				if err = ie.twoHARQ_ACK_CodebookForUnicastAndMulticast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode twoHARQ_ACK_CodebookForUnicastAndMulticast_r17", err)
				}
			}
			// decode multiPUCCH_HARQ_ACK_ForMulticastUnicast_r17 optional
			if multiPUCCH_HARQ_ACK_ForMulticastUnicast_r17Present {
				ie.multiPUCCH_HARQ_ACK_ForMulticastUnicast_r17 = new(Phy_ParametersCommon_multiPUCCH_HARQ_ACK_ForMulticastUnicast_r17)
				if err = ie.multiPUCCH_HARQ_ACK_ForMulticastUnicast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode multiPUCCH_HARQ_ACK_ForMulticastUnicast_r17", err)
				}
			}
			// decode srs_AdditionalRepetition_r17 optional
			if srs_AdditionalRepetition_r17Present {
				ie.srs_AdditionalRepetition_r17 = new(Phy_ParametersCommon_srs_AdditionalRepetition_r17)
				if err = ie.srs_AdditionalRepetition_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode srs_AdditionalRepetition_r17", err)
				}
			}
			// decode pusch_Repetition_CG_SDT_r17 optional
			if pusch_Repetition_CG_SDT_r17Present {
				ie.pusch_Repetition_CG_SDT_r17 = new(Phy_ParametersCommon_pusch_Repetition_CG_SDT_r17)
				if err = ie.pusch_Repetition_CG_SDT_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pusch_Repetition_CG_SDT_r17", err)
				}
			}
		}
	}
	return nil
}
