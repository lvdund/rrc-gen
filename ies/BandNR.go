package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandNR struct {
	bandNR                                                         FreqBandIndicatorNR                                                    `madatory`
	modifiedMPR_Behaviour                                          *uper.BitString                                                        `lb:8,ub:8,optional`
	mimo_ParametersPerBand                                         *MIMO_ParametersPerBand                                                `optional`
	extendedCP                                                     *BandNR_extendedCP                                                     `optional`
	multipleTCI                                                    *BandNR_multipleTCI                                                    `optional`
	bwp_WithoutRestriction                                         *BandNR_bwp_WithoutRestriction                                         `optional`
	bwp_SameNumerology                                             *BandNR_bwp_SameNumerology                                             `optional`
	bwp_DiffNumerology                                             *BandNR_bwp_DiffNumerology                                             `optional`
	crossCarrierScheduling_SameSCS                                 *BandNR_crossCarrierScheduling_SameSCS                                 `optional`
	pdsch_256QAM_FR2                                               *BandNR_pdsch_256QAM_FR2                                               `optional`
	pusch_256QAM                                                   *BandNR_pusch_256QAM                                                   `optional`
	ue_PowerClass                                                  *BandNR_ue_PowerClass                                                  `optional`
	rateMatchingLTE_CRS                                            *BandNR_rateMatchingLTE_CRS                                            `optional`
	channelBWs_DL                                                  *BandNR_channelBWs_DL                                                  `lb:10,ub:10,optional`
	channelBWs_UL                                                  *BandNR_channelBWs_UL                                                  `lb:10,ub:10,optional`
	maxUplinkDutyCycle_PC2_FR1                                     *BandNR_maxUplinkDutyCycle_PC2_FR1                                     `optional,ext-1`
	pucch_SpatialRelInfoMAC_CE                                     *BandNR_pucch_SpatialRelInfoMAC_CE                                     `optional,ext-2`
	powerBoosting_pi2BPSK                                          *BandNR_powerBoosting_pi2BPSK                                          `optional,ext-2`
	maxUplinkDutyCycle_FR2                                         *BandNR_maxUplinkDutyCycle_FR2                                         `optional,ext-3`
	channelBWs_DL_v1590                                            *BandNR_channelBWs_DL_v1590                                            `lb:16,ub:16,optional,ext-4`
	channelBWs_UL_v1590                                            *BandNR_channelBWs_UL_v1590                                            `lb:16,ub:16,optional,ext-4`
	asymmetricBandwidthCombinationSet                              *uper.BitString                                                        `lb:1,ub:32,optional,ext-5`
	sharedSpectrumChAccessParamsPerBand_r16                        *SharedSpectrumChAccessParamsPerBand_r16                               `optional,ext-6`
	cancelOverlappingPUSCH_r16                                     *BandNR_cancelOverlappingPUSCH_r16                                     `optional,ext-6`
	multipleRateMatchingEUTRA_CRS_r16                              *BandNR_multipleRateMatchingEUTRA_CRS_r16                              `lb:2,ub:6,optional,ext-6`
	overlapRateMatchingEUTRA_CRS_r16                               *BandNR_overlapRateMatchingEUTRA_CRS_r16                               `optional,ext-6`
	pdsch_MappingTypeB_Alt_r16                                     *BandNR_pdsch_MappingTypeB_Alt_r16                                     `optional,ext-6`
	oneSlotPeriodicTRS_r16                                         *BandNR_oneSlotPeriodicTRS_r16                                         `optional,ext-6`
	olpc_SRS_Pos_r16                                               *OLPC_SRS_Pos_r16                                                      `optional,ext-6`
	spatialRelationsSRS_Pos_r16                                    *SpatialRelationsSRS_Pos_r16                                           `optional,ext-6`
	simulSRS_MIMO_TransWithinBand_r16                              *BandNR_simulSRS_MIMO_TransWithinBand_r16                              `optional,ext-6`
	channelBW_DL_IAB_r16                                           *BandNR_channelBW_DL_IAB_r16                                           `optional,ext-6`
	channelBW_UL_IAB_r16                                           *BandNR_channelBW_UL_IAB_r16                                           `optional,ext-6`
	rasterShift7dot5_IAB_r16                                       *BandNR_rasterShift7dot5_IAB_r16                                       `optional,ext-6`
	ue_PowerClass_v1610                                            *BandNR_ue_PowerClass_v1610                                            `optional,ext-6`
	condHandover_r16                                               *BandNR_condHandover_r16                                               `optional,ext-6`
	condHandoverFailure_r16                                        *BandNR_condHandoverFailure_r16                                        `optional,ext-6`
	condHandoverTwoTriggerEvents_r16                               *BandNR_condHandoverTwoTriggerEvents_r16                               `optional,ext-6`
	condPSCellChange_r16                                           *BandNR_condPSCellChange_r16                                           `optional,ext-6`
	condPSCellChangeTwoTriggerEvents_r16                           *BandNR_condPSCellChangeTwoTriggerEvents_r16                           `optional,ext-6`
	mpr_PowerBoost_FR2_r16                                         *BandNR_mpr_PowerBoost_FR2_r16                                         `optional,ext-6`
	activeConfiguredGrant_r16                                      *BandNR_activeConfiguredGrant_r16                                      `lb:2,ub:32,optional,ext-6`
	jointReleaseConfiguredGrantType2_r16                           *BandNR_jointReleaseConfiguredGrantType2_r16                           `optional,ext-6`
	sps_r16                                                        *BandNR_sps_r16                                                        `lb:1,ub:8,optional,ext-6`
	jointReleaseSPS_r16                                            *BandNR_jointReleaseSPS_r16                                            `optional,ext-6`
	simulSRS_TransWithinBand_r16                                   *BandNR_simulSRS_TransWithinBand_r16                                   `optional,ext-6`
	trs_AdditionalBandwidth_r16                                    *BandNR_trs_AdditionalBandwidth_r16                                    `optional,ext-6`
	handoverIntraF_IAB_r16                                         *BandNR_handoverIntraF_IAB_r16                                         `optional,ext-6`
	simulTX_SRS_AntSwitchingIntraBandUL_CA_r16                     *SimulSRS_ForAntennaSwitching_r16                                      `optional,ext-7`
	sharedSpectrumChAccessParamsPerBand_v1630                      *SharedSpectrumChAccessParamsPerBand_v1630                             `optional,ext-7`
	handoverUTRA_FDD_r16                                           *BandNR_handoverUTRA_FDD_r16                                           `optional,ext-8`
	enhancedUL_TransientPeriod_r16                                 *BandNR_enhancedUL_TransientPeriod_r16                                 `optional,ext-8`
	sharedSpectrumChAccessParamsPerBand_v1640                      *SharedSpectrumChAccessParamsPerBand_v1640                             `optional,ext-8`
	type1_PUSCH_RepetitionMultiSlots_v1650                         *BandNR_type1_PUSCH_RepetitionMultiSlots_v1650                         `optional,ext-9`
	type2_PUSCH_RepetitionMultiSlots_v1650                         *BandNR_type2_PUSCH_RepetitionMultiSlots_v1650                         `optional,ext-9`
	pusch_RepetitionMultiSlots_v1650                               *BandNR_pusch_RepetitionMultiSlots_v1650                               `optional,ext-9`
	configuredUL_GrantType1_v1650                                  *BandNR_configuredUL_GrantType1_v1650                                  `optional,ext-9`
	configuredUL_GrantType2_v1650                                  *BandNR_configuredUL_GrantType2_v1650                                  `optional,ext-9`
	sharedSpectrumChAccessParamsPerBand_v1650                      *SharedSpectrumChAccessParamsPerBand_v1650                             `optional,ext-9`
	enhancedSkipUplinkTxConfigured_v1660                           *BandNR_enhancedSkipUplinkTxConfigured_v1660                           `optional,ext-10`
	enhancedSkipUplinkTxDynamic_v1660                              *BandNR_enhancedSkipUplinkTxDynamic_v1660                              `optional,ext-10`
	maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16                         *BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16                         `optional,ext-11`
	txDiversity_r16                                                *BandNR_txDiversity_r16                                                `optional,ext-11`
	pdsch_1024QAM_FR1_r17                                          *BandNR_pdsch_1024QAM_FR1_r17                                          `optional,ext-12`
	ue_PowerClass_v1700                                            *BandNR_ue_PowerClass_v1700                                            `optional,ext-12`
	fr2_2_AccessParamsPerBand_r17                                  *FR2_2_AccessParamsPerBand_r17                                         `optional,ext-12`
	rlm_Relaxation_r17                                             *BandNR_rlm_Relaxation_r17                                             `optional,ext-12`
	bfd_Relaxation_r17                                             *BandNR_bfd_Relaxation_r17                                             `optional,ext-12`
	cg_SDT_r17                                                     *BandNR_cg_SDT_r17                                                     `optional,ext-12`
	locationBasedCondHandover_r17                                  *BandNR_locationBasedCondHandover_r17                                  `optional,ext-12`
	timeBasedCondHandover_r17                                      *BandNR_timeBasedCondHandover_r17                                      `optional,ext-12`
	eventA4BasedCondHandover_r17                                   *BandNR_eventA4BasedCondHandover_r17                                   `optional,ext-12`
	mn_InitiatedCondPSCellChangeNRDC_r17                           *BandNR_mn_InitiatedCondPSCellChangeNRDC_r17                           `optional,ext-12`
	sn_InitiatedCondPSCellChangeNRDC_r17                           *BandNR_sn_InitiatedCondPSCellChangeNRDC_r17                           `optional,ext-12`
	pdcch_SkippingWithoutSSSG_r17                                  *BandNR_pdcch_SkippingWithoutSSSG_r17                                  `optional,ext-12`
	sssg_Switching_1BitInd_r17                                     *BandNR_sssg_Switching_1BitInd_r17                                     `optional,ext-12`
	sssg_Switching_2BitInd_r17                                     *BandNR_sssg_Switching_2BitInd_r17                                     `optional,ext-12`
	pdcch_SkippingWithSSSG_r17                                     *BandNR_pdcch_SkippingWithSSSG_r17                                     `optional,ext-12`
	searchSpaceSetGrp_switchCap2_r17                               *BandNR_searchSpaceSetGrp_switchCap2_r17                               `optional,ext-12`
	uplinkPreCompensation_r17                                      *BandNR_uplinkPreCompensation_r17                                      `optional,ext-12`
	uplink_TA_Reporting_r17                                        *BandNR_uplink_TA_Reporting_r17                                        `optional,ext-12`
	max_HARQ_ProcessNumber_r17                                     *BandNR_max_HARQ_ProcessNumber_r17                                     `optional,ext-12`
	type2_HARQ_Codebook_r17                                        *BandNR_type2_HARQ_Codebook_r17                                        `optional,ext-12`
	type1_HARQ_Codebook_r17                                        *BandNR_type1_HARQ_Codebook_r17                                        `optional,ext-12`
	type3_HARQ_Codebook_r17                                        *BandNR_type3_HARQ_Codebook_r17                                        `optional,ext-12`
	ue_specific_K_Offset_r17                                       *BandNR_ue_specific_K_Offset_r17                                       `optional,ext-12`
	multiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17                      *BandNR_multiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17                      `optional,ext-12`
	multiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17                      *BandNR_multiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17                      `optional,ext-12`
	parallelPRS_MeasRRC_Inactive_r17                               *BandNR_parallelPRS_MeasRRC_Inactive_r17                               `optional,ext-12`
	nr_UE_TxTEG_ID_MaxSupport_r17                                  *BandNR_nr_UE_TxTEG_ID_MaxSupport_r17                                  `optional,ext-12`
	prs_ProcessingRRC_Inactive_r17                                 *BandNR_prs_ProcessingRRC_Inactive_r17                                 `optional,ext-12`
	prs_ProcessingWindowType1A_r17                                 *BandNR_prs_ProcessingWindowType1A_r17                                 `optional,ext-12`
	prs_ProcessingWindowType1B_r17                                 *BandNR_prs_ProcessingWindowType1B_r17                                 `optional,ext-12`
	prs_ProcessingWindowType2_r17                                  *BandNR_prs_ProcessingWindowType2_r17                                  `optional,ext-12`
	srs_AllPosResourcesRRC_Inactive_r17                            *SRS_AllPosResourcesRRC_Inactive_r17                                   `optional,ext-12`
	olpc_SRS_PosRRC_Inactive_r17                                   *OLPC_SRS_Pos_r16                                                      `optional,ext-12`
	spatialRelationsSRS_PosRRC_Inactive_r17                        *SpatialRelationsSRS_Pos_r16                                           `optional,ext-12`
	maxNumberPUSCH_TypeA_Repetition_r17                            *BandNR_maxNumberPUSCH_TypeA_Repetition_r17                            `optional,ext-12`
	puschTypeA_RepetitionsAvailSlot_r17                            *BandNR_puschTypeA_RepetitionsAvailSlot_r17                            `optional,ext-12`
	tb_ProcessingMultiSlotPUSCH_r17                                *BandNR_tb_ProcessingMultiSlotPUSCH_r17                                `optional,ext-12`
	tb_ProcessingRepMultiSlotPUSCH_r17                             *BandNR_tb_ProcessingRepMultiSlotPUSCH_r17                             `optional,ext-12`
	maxDurationDMRS_Bundling_r17                                   *BandNR_maxDurationDMRS_Bundling_r17                                   `optional,ext-12`
	pusch_RepetitionMsg3_r17                                       *BandNR_pusch_RepetitionMsg3_r17                                       `optional,ext-12`
	sharedSpectrumChAccessParamsPerBand_v1710                      *SharedSpectrumChAccessParamsPerBand_v1710                             `optional,ext-12`
	parallelMeasurementWithoutRestriction_r17                      *BandNR_parallelMeasurementWithoutRestriction_r17                      `optional,ext-12`
	maxNumber_NGSO_SatellitesWithinOneSMTC_r17                     *BandNR_maxNumber_NGSO_SatellitesWithinOneSMTC_r17                     `optional,ext-12`
	k1_RangeExtension_r17                                          *BandNR_k1_RangeExtension_r17                                          `optional,ext-12`
	aperiodicCSI_RS_FastScellActivation_r17                        *BandNR_aperiodicCSI_RS_FastScellActivation_r17                        `optional,ext-12`
	aperiodicCSI_RS_AdditionalBandwidth_r17                        *BandNR_aperiodicCSI_RS_AdditionalBandwidth_r17                        `optional,ext-12`
	bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17                         *BandNR_bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17                         `optional,ext-12`
	halfDuplexFDD_TypeA_RedCap_r17                                 *BandNR_halfDuplexFDD_TypeA_RedCap_r17                                 `optional,ext-12`
	posSRS_RRC_Inactive_OutsideInitialUL_BWP_r17                   *PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17                          `optional,ext-12`
	channelBWs_DL_SCS_480kHz_FR2_2_r17                             *uper.BitString                                                        `lb:8,ub:8,optional,ext-12`
	channelBWs_UL_SCS_480kHz_FR2_2_r17                             *uper.BitString                                                        `lb:8,ub:8,optional,ext-12`
	channelBWs_DL_SCS_960kHz_FR2_2_r17                             *uper.BitString                                                        `lb:8,ub:8,optional,ext-12`
	channelBWs_UL_SCS_960kHz_FR2_2_r17                             *uper.BitString                                                        `lb:8,ub:8,optional,ext-12`
	ul_GapFR2_r17                                                  *BandNR_ul_GapFR2_r17                                                  `optional,ext-12`
	oneShotHARQ_feedbackTriggeredByDCI_1_2_r17                     *BandNR_oneShotHARQ_feedbackTriggeredByDCI_1_2_r17                     `optional,ext-12`
	oneShotHARQ_feedbackPhy_Priority_r17                           *BandNR_oneShotHARQ_feedbackPhy_Priority_r17                           `optional,ext-12`
	enhancedType3_HARQ_CodebookFeedback_r17                        *BandNR_enhancedType3_HARQ_CodebookFeedback_r17                        `optional,ext-12`
	triggeredHARQ_CodebookRetx_r17                                 *BandNR_triggeredHARQ_CodebookRetx_r17                                 `optional,ext-12`
	ue_OneShotUL_TimingAdj_r17                                     *BandNR_ue_OneShotUL_TimingAdj_r17                                     `optional,ext-13`
	pucch_Repetition_F0_2_r17                                      *BandNR_pucch_Repetition_F0_2_r17                                      `optional,ext-13`
	cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17                *BandNR_cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17                `optional,ext-13`
	mux_HARQ_ACK_DiffPriorities_r17                                *BandNR_mux_HARQ_ACK_DiffPriorities_r17                                `optional,ext-13`
	ta_BasedPDC_NTN_SharedSpectrumChAccess_r17                     *BandNR_ta_BasedPDC_NTN_SharedSpectrumChAccess_r17                     `optional,ext-13`
	ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17               *BandNR_ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17               `optional,ext-13`
	maxNumberG_RNTI_r17                                            *int64                                                                 `lb:2,ub:8,optional,ext-13`
	dynamicMulticastDCI_Format4_2_r17                              *BandNR_dynamicMulticastDCI_Format4_2_r17                              `optional,ext-13`
	maxModulationOrderForMulticast_r17                             *BandNR_maxModulationOrderForMulticast_r17                             `optional,ext-13`
	dynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 *BandNR_dynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 `optional,ext-13`
	dynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17   *BandNR_dynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17   `optional,ext-13`
	nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17               *BandNR_nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17               `optional,ext-13`
	ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17           *BandNR_ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17           `optional,ext-13`
	maxNumberG_CS_RNTI_r17                                         *int64                                                                 `lb:2,ub:8,optional,ext-13`
	re_LevelRateMatchingForMulticast_r17                           *BandNR_re_LevelRateMatchingForMulticast_r17                           `optional,ext-13`
	pdsch_1024QAM_2MIMO_FR1_r17                                    *BandNR_pdsch_1024QAM_2MIMO_FR1_r17                                    `optional,ext-13`
	prs_MeasurementWithoutMG_r17                                   *BandNR_prs_MeasurementWithoutMG_r17                                   `optional,ext-13`
	maxNumber_LEO_SatellitesPerCarrier_r17                         *int64                                                                 `lb:3,ub:4,optional,ext-13`
	prs_ProcessingCapabilityOutsideMGinPPW_r17                     []PRS_ProcessingCapabilityOutsideMGinPPWperType_r17                    `lb:1,ub:3,optional,ext-13`
	srs_SemiPersistent_PosResourcesRRC_Inactive_r17                *BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17                `optional,ext-13`
	channelBWs_DL_SCS_120kHz_FR2_2_r17                             *uper.BitString                                                        `lb:8,ub:8,optional,ext-13`
	channelBWs_UL_SCS_120kHz_FR2_2_r17                             *uper.BitString                                                        `lb:8,ub:8,optional,ext-13`
	dmrs_BundlingPUSCH_RepTypeA_r17                                *BandNR_dmrs_BundlingPUSCH_RepTypeA_r17                                `optional,ext-14`
	dmrs_BundlingPUSCH_RepTypeB_r17                                *BandNR_dmrs_BundlingPUSCH_RepTypeB_r17                                `optional,ext-14`
	dmrs_BundlingPUSCH_multiSlot_r17                               *BandNR_dmrs_BundlingPUSCH_multiSlot_r17                               `optional,ext-14`
	dmrs_BundlingPUCCH_Rep_r17                                     *BandNR_dmrs_BundlingPUCCH_Rep_r17                                     `optional,ext-14`
	interSlotFreqHopInterSlotBundlingPUSCH_r17                     *BandNR_interSlotFreqHopInterSlotBundlingPUSCH_r17                     `optional,ext-14`
	interSlotFreqHopPUCCH_r17                                      *BandNR_interSlotFreqHopPUCCH_r17                                      `optional,ext-14`
	dmrs_BundlingRestart_r17                                       *BandNR_dmrs_BundlingRestart_r17                                       `optional,ext-14`
	dmrs_BundlingNonBackToBackTX_r17                               *BandNR_dmrs_BundlingNonBackToBackTX_r17                               `optional,ext-14`
}

func (ie *BandNR) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.maxUplinkDutyCycle_PC2_FR1 != nil || ie.pucch_SpatialRelInfoMAC_CE != nil || ie.powerBoosting_pi2BPSK != nil || ie.maxUplinkDutyCycle_FR2 != nil || ie.channelBWs_DL_v1590 != nil || ie.channelBWs_UL_v1590 != nil || ie.asymmetricBandwidthCombinationSet != nil || ie.sharedSpectrumChAccessParamsPerBand_r16 != nil || ie.cancelOverlappingPUSCH_r16 != nil || ie.multipleRateMatchingEUTRA_CRS_r16 != nil || ie.overlapRateMatchingEUTRA_CRS_r16 != nil || ie.pdsch_MappingTypeB_Alt_r16 != nil || ie.oneSlotPeriodicTRS_r16 != nil || ie.olpc_SRS_Pos_r16 != nil || ie.spatialRelationsSRS_Pos_r16 != nil || ie.simulSRS_MIMO_TransWithinBand_r16 != nil || ie.channelBW_DL_IAB_r16 != nil || ie.channelBW_UL_IAB_r16 != nil || ie.rasterShift7dot5_IAB_r16 != nil || ie.ue_PowerClass_v1610 != nil || ie.condHandover_r16 != nil || ie.condHandoverFailure_r16 != nil || ie.condHandoverTwoTriggerEvents_r16 != nil || ie.condPSCellChange_r16 != nil || ie.condPSCellChangeTwoTriggerEvents_r16 != nil || ie.mpr_PowerBoost_FR2_r16 != nil || ie.activeConfiguredGrant_r16 != nil || ie.jointReleaseConfiguredGrantType2_r16 != nil || ie.sps_r16 != nil || ie.jointReleaseSPS_r16 != nil || ie.simulSRS_TransWithinBand_r16 != nil || ie.trs_AdditionalBandwidth_r16 != nil || ie.handoverIntraF_IAB_r16 != nil || ie.simulTX_SRS_AntSwitchingIntraBandUL_CA_r16 != nil || ie.sharedSpectrumChAccessParamsPerBand_v1630 != nil || ie.handoverUTRA_FDD_r16 != nil || ie.enhancedUL_TransientPeriod_r16 != nil || ie.sharedSpectrumChAccessParamsPerBand_v1640 != nil || ie.type1_PUSCH_RepetitionMultiSlots_v1650 != nil || ie.type2_PUSCH_RepetitionMultiSlots_v1650 != nil || ie.pusch_RepetitionMultiSlots_v1650 != nil || ie.configuredUL_GrantType1_v1650 != nil || ie.configuredUL_GrantType2_v1650 != nil || ie.sharedSpectrumChAccessParamsPerBand_v1650 != nil || ie.enhancedSkipUplinkTxConfigured_v1660 != nil || ie.enhancedSkipUplinkTxDynamic_v1660 != nil || ie.maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16 != nil || ie.txDiversity_r16 != nil || ie.pdsch_1024QAM_FR1_r17 != nil || ie.ue_PowerClass_v1700 != nil || ie.fr2_2_AccessParamsPerBand_r17 != nil || ie.rlm_Relaxation_r17 != nil || ie.bfd_Relaxation_r17 != nil || ie.cg_SDT_r17 != nil || ie.locationBasedCondHandover_r17 != nil || ie.timeBasedCondHandover_r17 != nil || ie.eventA4BasedCondHandover_r17 != nil || ie.mn_InitiatedCondPSCellChangeNRDC_r17 != nil || ie.sn_InitiatedCondPSCellChangeNRDC_r17 != nil || ie.pdcch_SkippingWithoutSSSG_r17 != nil || ie.sssg_Switching_1BitInd_r17 != nil || ie.sssg_Switching_2BitInd_r17 != nil || ie.pdcch_SkippingWithSSSG_r17 != nil || ie.searchSpaceSetGrp_switchCap2_r17 != nil || ie.uplinkPreCompensation_r17 != nil || ie.uplink_TA_Reporting_r17 != nil || ie.max_HARQ_ProcessNumber_r17 != nil || ie.type2_HARQ_Codebook_r17 != nil || ie.type1_HARQ_Codebook_r17 != nil || ie.type3_HARQ_Codebook_r17 != nil || ie.ue_specific_K_Offset_r17 != nil || ie.multiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil || ie.multiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil || ie.parallelPRS_MeasRRC_Inactive_r17 != nil || ie.nr_UE_TxTEG_ID_MaxSupport_r17 != nil || ie.prs_ProcessingRRC_Inactive_r17 != nil || ie.prs_ProcessingWindowType1A_r17 != nil || ie.prs_ProcessingWindowType1B_r17 != nil || ie.prs_ProcessingWindowType2_r17 != nil || ie.srs_AllPosResourcesRRC_Inactive_r17 != nil || ie.olpc_SRS_PosRRC_Inactive_r17 != nil || ie.spatialRelationsSRS_PosRRC_Inactive_r17 != nil || ie.maxNumberPUSCH_TypeA_Repetition_r17 != nil || ie.puschTypeA_RepetitionsAvailSlot_r17 != nil || ie.tb_ProcessingMultiSlotPUSCH_r17 != nil || ie.tb_ProcessingRepMultiSlotPUSCH_r17 != nil || ie.maxDurationDMRS_Bundling_r17 != nil || ie.pusch_RepetitionMsg3_r17 != nil || ie.sharedSpectrumChAccessParamsPerBand_v1710 != nil || ie.parallelMeasurementWithoutRestriction_r17 != nil || ie.maxNumber_NGSO_SatellitesWithinOneSMTC_r17 != nil || ie.k1_RangeExtension_r17 != nil || ie.aperiodicCSI_RS_FastScellActivation_r17 != nil || ie.aperiodicCSI_RS_AdditionalBandwidth_r17 != nil || ie.bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17 != nil || ie.halfDuplexFDD_TypeA_RedCap_r17 != nil || ie.posSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 != nil || ie.channelBWs_DL_SCS_480kHz_FR2_2_r17 != nil || ie.channelBWs_UL_SCS_480kHz_FR2_2_r17 != nil || ie.channelBWs_DL_SCS_960kHz_FR2_2_r17 != nil || ie.channelBWs_UL_SCS_960kHz_FR2_2_r17 != nil || ie.ul_GapFR2_r17 != nil || ie.oneShotHARQ_feedbackTriggeredByDCI_1_2_r17 != nil || ie.oneShotHARQ_feedbackPhy_Priority_r17 != nil || ie.enhancedType3_HARQ_CodebookFeedback_r17 != nil || ie.triggeredHARQ_CodebookRetx_r17 != nil || ie.ue_OneShotUL_TimingAdj_r17 != nil || ie.pucch_Repetition_F0_2_r17 != nil || ie.cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17 != nil || ie.mux_HARQ_ACK_DiffPriorities_r17 != nil || ie.ta_BasedPDC_NTN_SharedSpectrumChAccess_r17 != nil || ie.ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17 != nil || ie.maxNumberG_RNTI_r17 != nil || ie.dynamicMulticastDCI_Format4_2_r17 != nil || ie.maxModulationOrderForMulticast_r17 != nil || ie.dynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 != nil || ie.dynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17 != nil || ie.nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17 != nil || ie.ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17 != nil || ie.maxNumberG_CS_RNTI_r17 != nil || ie.re_LevelRateMatchingForMulticast_r17 != nil || ie.pdsch_1024QAM_2MIMO_FR1_r17 != nil || ie.prs_MeasurementWithoutMG_r17 != nil || ie.maxNumber_LEO_SatellitesPerCarrier_r17 != nil || len(ie.prs_ProcessingCapabilityOutsideMGinPPW_r17) > 0 || ie.srs_SemiPersistent_PosResourcesRRC_Inactive_r17 != nil || ie.channelBWs_DL_SCS_120kHz_FR2_2_r17 != nil || ie.channelBWs_UL_SCS_120kHz_FR2_2_r17 != nil || ie.dmrs_BundlingPUSCH_RepTypeA_r17 != nil || ie.dmrs_BundlingPUSCH_RepTypeB_r17 != nil || ie.dmrs_BundlingPUSCH_multiSlot_r17 != nil || ie.dmrs_BundlingPUCCH_Rep_r17 != nil || ie.interSlotFreqHopInterSlotBundlingPUSCH_r17 != nil || ie.interSlotFreqHopPUCCH_r17 != nil || ie.dmrs_BundlingRestart_r17 != nil || ie.dmrs_BundlingNonBackToBackTX_r17 != nil
	preambleBits := []bool{hasExtensions, ie.modifiedMPR_Behaviour != nil, ie.mimo_ParametersPerBand != nil, ie.extendedCP != nil, ie.multipleTCI != nil, ie.bwp_WithoutRestriction != nil, ie.bwp_SameNumerology != nil, ie.bwp_DiffNumerology != nil, ie.crossCarrierScheduling_SameSCS != nil, ie.pdsch_256QAM_FR2 != nil, ie.pusch_256QAM != nil, ie.ue_PowerClass != nil, ie.rateMatchingLTE_CRS != nil, ie.channelBWs_DL != nil, ie.channelBWs_UL != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.bandNR.Encode(w); err != nil {
		return utils.WrapError("Encode bandNR", err)
	}
	if ie.modifiedMPR_Behaviour != nil {
		if err = w.WriteBitString(ie.modifiedMPR_Behaviour.Bytes, uint(ie.modifiedMPR_Behaviour.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode modifiedMPR_Behaviour", err)
		}
	}
	if ie.mimo_ParametersPerBand != nil {
		if err = ie.mimo_ParametersPerBand.Encode(w); err != nil {
			return utils.WrapError("Encode mimo_ParametersPerBand", err)
		}
	}
	if ie.extendedCP != nil {
		if err = ie.extendedCP.Encode(w); err != nil {
			return utils.WrapError("Encode extendedCP", err)
		}
	}
	if ie.multipleTCI != nil {
		if err = ie.multipleTCI.Encode(w); err != nil {
			return utils.WrapError("Encode multipleTCI", err)
		}
	}
	if ie.bwp_WithoutRestriction != nil {
		if err = ie.bwp_WithoutRestriction.Encode(w); err != nil {
			return utils.WrapError("Encode bwp_WithoutRestriction", err)
		}
	}
	if ie.bwp_SameNumerology != nil {
		if err = ie.bwp_SameNumerology.Encode(w); err != nil {
			return utils.WrapError("Encode bwp_SameNumerology", err)
		}
	}
	if ie.bwp_DiffNumerology != nil {
		if err = ie.bwp_DiffNumerology.Encode(w); err != nil {
			return utils.WrapError("Encode bwp_DiffNumerology", err)
		}
	}
	if ie.crossCarrierScheduling_SameSCS != nil {
		if err = ie.crossCarrierScheduling_SameSCS.Encode(w); err != nil {
			return utils.WrapError("Encode crossCarrierScheduling_SameSCS", err)
		}
	}
	if ie.pdsch_256QAM_FR2 != nil {
		if err = ie.pdsch_256QAM_FR2.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_256QAM_FR2", err)
		}
	}
	if ie.pusch_256QAM != nil {
		if err = ie.pusch_256QAM.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_256QAM", err)
		}
	}
	if ie.ue_PowerClass != nil {
		if err = ie.ue_PowerClass.Encode(w); err != nil {
			return utils.WrapError("Encode ue_PowerClass", err)
		}
	}
	if ie.rateMatchingLTE_CRS != nil {
		if err = ie.rateMatchingLTE_CRS.Encode(w); err != nil {
			return utils.WrapError("Encode rateMatchingLTE_CRS", err)
		}
	}
	if ie.channelBWs_DL != nil {
		if err = ie.channelBWs_DL.Encode(w); err != nil {
			return utils.WrapError("Encode channelBWs_DL", err)
		}
	}
	if ie.channelBWs_UL != nil {
		if err = ie.channelBWs_UL.Encode(w); err != nil {
			return utils.WrapError("Encode channelBWs_UL", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 14 bits for 14 extension groups
		extBitmap := []bool{ie.maxUplinkDutyCycle_PC2_FR1 != nil, ie.pucch_SpatialRelInfoMAC_CE != nil || ie.powerBoosting_pi2BPSK != nil, ie.maxUplinkDutyCycle_FR2 != nil, ie.channelBWs_DL_v1590 != nil || ie.channelBWs_UL_v1590 != nil, ie.asymmetricBandwidthCombinationSet != nil, ie.sharedSpectrumChAccessParamsPerBand_r16 != nil || ie.cancelOverlappingPUSCH_r16 != nil || ie.multipleRateMatchingEUTRA_CRS_r16 != nil || ie.overlapRateMatchingEUTRA_CRS_r16 != nil || ie.pdsch_MappingTypeB_Alt_r16 != nil || ie.oneSlotPeriodicTRS_r16 != nil || ie.olpc_SRS_Pos_r16 != nil || ie.spatialRelationsSRS_Pos_r16 != nil || ie.simulSRS_MIMO_TransWithinBand_r16 != nil || ie.channelBW_DL_IAB_r16 != nil || ie.channelBW_UL_IAB_r16 != nil || ie.rasterShift7dot5_IAB_r16 != nil || ie.ue_PowerClass_v1610 != nil || ie.condHandover_r16 != nil || ie.condHandoverFailure_r16 != nil || ie.condHandoverTwoTriggerEvents_r16 != nil || ie.condPSCellChange_r16 != nil || ie.condPSCellChangeTwoTriggerEvents_r16 != nil || ie.mpr_PowerBoost_FR2_r16 != nil || ie.activeConfiguredGrant_r16 != nil || ie.jointReleaseConfiguredGrantType2_r16 != nil || ie.sps_r16 != nil || ie.jointReleaseSPS_r16 != nil || ie.simulSRS_TransWithinBand_r16 != nil || ie.trs_AdditionalBandwidth_r16 != nil || ie.handoverIntraF_IAB_r16 != nil, ie.simulTX_SRS_AntSwitchingIntraBandUL_CA_r16 != nil || ie.sharedSpectrumChAccessParamsPerBand_v1630 != nil, ie.handoverUTRA_FDD_r16 != nil || ie.enhancedUL_TransientPeriod_r16 != nil || ie.sharedSpectrumChAccessParamsPerBand_v1640 != nil, ie.type1_PUSCH_RepetitionMultiSlots_v1650 != nil || ie.type2_PUSCH_RepetitionMultiSlots_v1650 != nil || ie.pusch_RepetitionMultiSlots_v1650 != nil || ie.configuredUL_GrantType1_v1650 != nil || ie.configuredUL_GrantType2_v1650 != nil || ie.sharedSpectrumChAccessParamsPerBand_v1650 != nil, ie.enhancedSkipUplinkTxConfigured_v1660 != nil || ie.enhancedSkipUplinkTxDynamic_v1660 != nil, ie.maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16 != nil || ie.txDiversity_r16 != nil, ie.pdsch_1024QAM_FR1_r17 != nil || ie.ue_PowerClass_v1700 != nil || ie.fr2_2_AccessParamsPerBand_r17 != nil || ie.rlm_Relaxation_r17 != nil || ie.bfd_Relaxation_r17 != nil || ie.cg_SDT_r17 != nil || ie.locationBasedCondHandover_r17 != nil || ie.timeBasedCondHandover_r17 != nil || ie.eventA4BasedCondHandover_r17 != nil || ie.mn_InitiatedCondPSCellChangeNRDC_r17 != nil || ie.sn_InitiatedCondPSCellChangeNRDC_r17 != nil || ie.pdcch_SkippingWithoutSSSG_r17 != nil || ie.sssg_Switching_1BitInd_r17 != nil || ie.sssg_Switching_2BitInd_r17 != nil || ie.pdcch_SkippingWithSSSG_r17 != nil || ie.searchSpaceSetGrp_switchCap2_r17 != nil || ie.uplinkPreCompensation_r17 != nil || ie.uplink_TA_Reporting_r17 != nil || ie.max_HARQ_ProcessNumber_r17 != nil || ie.type2_HARQ_Codebook_r17 != nil || ie.type1_HARQ_Codebook_r17 != nil || ie.type3_HARQ_Codebook_r17 != nil || ie.ue_specific_K_Offset_r17 != nil || ie.multiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil || ie.multiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil || ie.parallelPRS_MeasRRC_Inactive_r17 != nil || ie.nr_UE_TxTEG_ID_MaxSupport_r17 != nil || ie.prs_ProcessingRRC_Inactive_r17 != nil || ie.prs_ProcessingWindowType1A_r17 != nil || ie.prs_ProcessingWindowType1B_r17 != nil || ie.prs_ProcessingWindowType2_r17 != nil || ie.srs_AllPosResourcesRRC_Inactive_r17 != nil || ie.olpc_SRS_PosRRC_Inactive_r17 != nil || ie.spatialRelationsSRS_PosRRC_Inactive_r17 != nil || ie.maxNumberPUSCH_TypeA_Repetition_r17 != nil || ie.puschTypeA_RepetitionsAvailSlot_r17 != nil || ie.tb_ProcessingMultiSlotPUSCH_r17 != nil || ie.tb_ProcessingRepMultiSlotPUSCH_r17 != nil || ie.maxDurationDMRS_Bundling_r17 != nil || ie.pusch_RepetitionMsg3_r17 != nil || ie.sharedSpectrumChAccessParamsPerBand_v1710 != nil || ie.parallelMeasurementWithoutRestriction_r17 != nil || ie.maxNumber_NGSO_SatellitesWithinOneSMTC_r17 != nil || ie.k1_RangeExtension_r17 != nil || ie.aperiodicCSI_RS_FastScellActivation_r17 != nil || ie.aperiodicCSI_RS_AdditionalBandwidth_r17 != nil || ie.bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17 != nil || ie.halfDuplexFDD_TypeA_RedCap_r17 != nil || ie.posSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 != nil || ie.channelBWs_DL_SCS_480kHz_FR2_2_r17 != nil || ie.channelBWs_UL_SCS_480kHz_FR2_2_r17 != nil || ie.channelBWs_DL_SCS_960kHz_FR2_2_r17 != nil || ie.channelBWs_UL_SCS_960kHz_FR2_2_r17 != nil || ie.ul_GapFR2_r17 != nil || ie.oneShotHARQ_feedbackTriggeredByDCI_1_2_r17 != nil || ie.oneShotHARQ_feedbackPhy_Priority_r17 != nil || ie.enhancedType3_HARQ_CodebookFeedback_r17 != nil || ie.triggeredHARQ_CodebookRetx_r17 != nil, ie.ue_OneShotUL_TimingAdj_r17 != nil || ie.pucch_Repetition_F0_2_r17 != nil || ie.cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17 != nil || ie.mux_HARQ_ACK_DiffPriorities_r17 != nil || ie.ta_BasedPDC_NTN_SharedSpectrumChAccess_r17 != nil || ie.ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17 != nil || ie.maxNumberG_RNTI_r17 != nil || ie.dynamicMulticastDCI_Format4_2_r17 != nil || ie.maxModulationOrderForMulticast_r17 != nil || ie.dynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 != nil || ie.dynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17 != nil || ie.nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17 != nil || ie.ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17 != nil || ie.maxNumberG_CS_RNTI_r17 != nil || ie.re_LevelRateMatchingForMulticast_r17 != nil || ie.pdsch_1024QAM_2MIMO_FR1_r17 != nil || ie.prs_MeasurementWithoutMG_r17 != nil || ie.maxNumber_LEO_SatellitesPerCarrier_r17 != nil || len(ie.prs_ProcessingCapabilityOutsideMGinPPW_r17) > 0 || ie.srs_SemiPersistent_PosResourcesRRC_Inactive_r17 != nil || ie.channelBWs_DL_SCS_120kHz_FR2_2_r17 != nil || ie.channelBWs_UL_SCS_120kHz_FR2_2_r17 != nil, ie.dmrs_BundlingPUSCH_RepTypeA_r17 != nil || ie.dmrs_BundlingPUSCH_RepTypeB_r17 != nil || ie.dmrs_BundlingPUSCH_multiSlot_r17 != nil || ie.dmrs_BundlingPUCCH_Rep_r17 != nil || ie.interSlotFreqHopInterSlotBundlingPUSCH_r17 != nil || ie.interSlotFreqHopPUCCH_r17 != nil || ie.dmrs_BundlingRestart_r17 != nil || ie.dmrs_BundlingNonBackToBackTX_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap BandNR", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.maxUplinkDutyCycle_PC2_FR1 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode maxUplinkDutyCycle_PC2_FR1 optional
			if ie.maxUplinkDutyCycle_PC2_FR1 != nil {
				if err = ie.maxUplinkDutyCycle_PC2_FR1.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxUplinkDutyCycle_PC2_FR1", err)
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
			optionals_ext_2 := []bool{ie.pucch_SpatialRelInfoMAC_CE != nil, ie.powerBoosting_pi2BPSK != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode pucch_SpatialRelInfoMAC_CE optional
			if ie.pucch_SpatialRelInfoMAC_CE != nil {
				if err = ie.pucch_SpatialRelInfoMAC_CE.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pucch_SpatialRelInfoMAC_CE", err)
				}
			}
			// encode powerBoosting_pi2BPSK optional
			if ie.powerBoosting_pi2BPSK != nil {
				if err = ie.powerBoosting_pi2BPSK.Encode(extWriter); err != nil {
					return utils.WrapError("Encode powerBoosting_pi2BPSK", err)
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
			optionals_ext_3 := []bool{ie.maxUplinkDutyCycle_FR2 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode maxUplinkDutyCycle_FR2 optional
			if ie.maxUplinkDutyCycle_FR2 != nil {
				if err = ie.maxUplinkDutyCycle_FR2.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxUplinkDutyCycle_FR2", err)
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
			optionals_ext_4 := []bool{ie.channelBWs_DL_v1590 != nil, ie.channelBWs_UL_v1590 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode channelBWs_DL_v1590 optional
			if ie.channelBWs_DL_v1590 != nil {
				if err = ie.channelBWs_DL_v1590.Encode(extWriter); err != nil {
					return utils.WrapError("Encode channelBWs_DL_v1590", err)
				}
			}
			// encode channelBWs_UL_v1590 optional
			if ie.channelBWs_UL_v1590 != nil {
				if err = ie.channelBWs_UL_v1590.Encode(extWriter); err != nil {
					return utils.WrapError("Encode channelBWs_UL_v1590", err)
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
			optionals_ext_5 := []bool{ie.asymmetricBandwidthCombinationSet != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode asymmetricBandwidthCombinationSet optional
			if ie.asymmetricBandwidthCombinationSet != nil {
				if err = extWriter.WriteBitString(ie.asymmetricBandwidthCombinationSet.Bytes, uint(ie.asymmetricBandwidthCombinationSet.NumBits), &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
					return utils.WrapError("Encode asymmetricBandwidthCombinationSet", err)
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
			optionals_ext_6 := []bool{ie.sharedSpectrumChAccessParamsPerBand_r16 != nil, ie.cancelOverlappingPUSCH_r16 != nil, ie.multipleRateMatchingEUTRA_CRS_r16 != nil, ie.overlapRateMatchingEUTRA_CRS_r16 != nil, ie.pdsch_MappingTypeB_Alt_r16 != nil, ie.oneSlotPeriodicTRS_r16 != nil, ie.olpc_SRS_Pos_r16 != nil, ie.spatialRelationsSRS_Pos_r16 != nil, ie.simulSRS_MIMO_TransWithinBand_r16 != nil, ie.channelBW_DL_IAB_r16 != nil, ie.channelBW_UL_IAB_r16 != nil, ie.rasterShift7dot5_IAB_r16 != nil, ie.ue_PowerClass_v1610 != nil, ie.condHandover_r16 != nil, ie.condHandoverFailure_r16 != nil, ie.condHandoverTwoTriggerEvents_r16 != nil, ie.condPSCellChange_r16 != nil, ie.condPSCellChangeTwoTriggerEvents_r16 != nil, ie.mpr_PowerBoost_FR2_r16 != nil, ie.activeConfiguredGrant_r16 != nil, ie.jointReleaseConfiguredGrantType2_r16 != nil, ie.sps_r16 != nil, ie.jointReleaseSPS_r16 != nil, ie.simulSRS_TransWithinBand_r16 != nil, ie.trs_AdditionalBandwidth_r16 != nil, ie.handoverIntraF_IAB_r16 != nil}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sharedSpectrumChAccessParamsPerBand_r16 optional
			if ie.sharedSpectrumChAccessParamsPerBand_r16 != nil {
				if err = ie.sharedSpectrumChAccessParamsPerBand_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sharedSpectrumChAccessParamsPerBand_r16", err)
				}
			}
			// encode cancelOverlappingPUSCH_r16 optional
			if ie.cancelOverlappingPUSCH_r16 != nil {
				if err = ie.cancelOverlappingPUSCH_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cancelOverlappingPUSCH_r16", err)
				}
			}
			// encode multipleRateMatchingEUTRA_CRS_r16 optional
			if ie.multipleRateMatchingEUTRA_CRS_r16 != nil {
				if err = ie.multipleRateMatchingEUTRA_CRS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode multipleRateMatchingEUTRA_CRS_r16", err)
				}
			}
			// encode overlapRateMatchingEUTRA_CRS_r16 optional
			if ie.overlapRateMatchingEUTRA_CRS_r16 != nil {
				if err = ie.overlapRateMatchingEUTRA_CRS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode overlapRateMatchingEUTRA_CRS_r16", err)
				}
			}
			// encode pdsch_MappingTypeB_Alt_r16 optional
			if ie.pdsch_MappingTypeB_Alt_r16 != nil {
				if err = ie.pdsch_MappingTypeB_Alt_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_MappingTypeB_Alt_r16", err)
				}
			}
			// encode oneSlotPeriodicTRS_r16 optional
			if ie.oneSlotPeriodicTRS_r16 != nil {
				if err = ie.oneSlotPeriodicTRS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode oneSlotPeriodicTRS_r16", err)
				}
			}
			// encode olpc_SRS_Pos_r16 optional
			if ie.olpc_SRS_Pos_r16 != nil {
				if err = ie.olpc_SRS_Pos_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode olpc_SRS_Pos_r16", err)
				}
			}
			// encode spatialRelationsSRS_Pos_r16 optional
			if ie.spatialRelationsSRS_Pos_r16 != nil {
				if err = ie.spatialRelationsSRS_Pos_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode spatialRelationsSRS_Pos_r16", err)
				}
			}
			// encode simulSRS_MIMO_TransWithinBand_r16 optional
			if ie.simulSRS_MIMO_TransWithinBand_r16 != nil {
				if err = ie.simulSRS_MIMO_TransWithinBand_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simulSRS_MIMO_TransWithinBand_r16", err)
				}
			}
			// encode channelBW_DL_IAB_r16 optional
			if ie.channelBW_DL_IAB_r16 != nil {
				if err = ie.channelBW_DL_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode channelBW_DL_IAB_r16", err)
				}
			}
			// encode channelBW_UL_IAB_r16 optional
			if ie.channelBW_UL_IAB_r16 != nil {
				if err = ie.channelBW_UL_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode channelBW_UL_IAB_r16", err)
				}
			}
			// encode rasterShift7dot5_IAB_r16 optional
			if ie.rasterShift7dot5_IAB_r16 != nil {
				if err = ie.rasterShift7dot5_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rasterShift7dot5_IAB_r16", err)
				}
			}
			// encode ue_PowerClass_v1610 optional
			if ie.ue_PowerClass_v1610 != nil {
				if err = ie.ue_PowerClass_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ue_PowerClass_v1610", err)
				}
			}
			// encode condHandover_r16 optional
			if ie.condHandover_r16 != nil {
				if err = ie.condHandover_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode condHandover_r16", err)
				}
			}
			// encode condHandoverFailure_r16 optional
			if ie.condHandoverFailure_r16 != nil {
				if err = ie.condHandoverFailure_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode condHandoverFailure_r16", err)
				}
			}
			// encode condHandoverTwoTriggerEvents_r16 optional
			if ie.condHandoverTwoTriggerEvents_r16 != nil {
				if err = ie.condHandoverTwoTriggerEvents_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode condHandoverTwoTriggerEvents_r16", err)
				}
			}
			// encode condPSCellChange_r16 optional
			if ie.condPSCellChange_r16 != nil {
				if err = ie.condPSCellChange_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode condPSCellChange_r16", err)
				}
			}
			// encode condPSCellChangeTwoTriggerEvents_r16 optional
			if ie.condPSCellChangeTwoTriggerEvents_r16 != nil {
				if err = ie.condPSCellChangeTwoTriggerEvents_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode condPSCellChangeTwoTriggerEvents_r16", err)
				}
			}
			// encode mpr_PowerBoost_FR2_r16 optional
			if ie.mpr_PowerBoost_FR2_r16 != nil {
				if err = ie.mpr_PowerBoost_FR2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mpr_PowerBoost_FR2_r16", err)
				}
			}
			// encode activeConfiguredGrant_r16 optional
			if ie.activeConfiguredGrant_r16 != nil {
				if err = ie.activeConfiguredGrant_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode activeConfiguredGrant_r16", err)
				}
			}
			// encode jointReleaseConfiguredGrantType2_r16 optional
			if ie.jointReleaseConfiguredGrantType2_r16 != nil {
				if err = ie.jointReleaseConfiguredGrantType2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode jointReleaseConfiguredGrantType2_r16", err)
				}
			}
			// encode sps_r16 optional
			if ie.sps_r16 != nil {
				if err = ie.sps_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sps_r16", err)
				}
			}
			// encode jointReleaseSPS_r16 optional
			if ie.jointReleaseSPS_r16 != nil {
				if err = ie.jointReleaseSPS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode jointReleaseSPS_r16", err)
				}
			}
			// encode simulSRS_TransWithinBand_r16 optional
			if ie.simulSRS_TransWithinBand_r16 != nil {
				if err = ie.simulSRS_TransWithinBand_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simulSRS_TransWithinBand_r16", err)
				}
			}
			// encode trs_AdditionalBandwidth_r16 optional
			if ie.trs_AdditionalBandwidth_r16 != nil {
				if err = ie.trs_AdditionalBandwidth_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode trs_AdditionalBandwidth_r16", err)
				}
			}
			// encode handoverIntraF_IAB_r16 optional
			if ie.handoverIntraF_IAB_r16 != nil {
				if err = ie.handoverIntraF_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode handoverIntraF_IAB_r16", err)
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
			optionals_ext_7 := []bool{ie.simulTX_SRS_AntSwitchingIntraBandUL_CA_r16 != nil, ie.sharedSpectrumChAccessParamsPerBand_v1630 != nil}
			for _, bit := range optionals_ext_7 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode simulTX_SRS_AntSwitchingIntraBandUL_CA_r16 optional
			if ie.simulTX_SRS_AntSwitchingIntraBandUL_CA_r16 != nil {
				if err = ie.simulTX_SRS_AntSwitchingIntraBandUL_CA_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simulTX_SRS_AntSwitchingIntraBandUL_CA_r16", err)
				}
			}
			// encode sharedSpectrumChAccessParamsPerBand_v1630 optional
			if ie.sharedSpectrumChAccessParamsPerBand_v1630 != nil {
				if err = ie.sharedSpectrumChAccessParamsPerBand_v1630.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sharedSpectrumChAccessParamsPerBand_v1630", err)
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
			optionals_ext_8 := []bool{ie.handoverUTRA_FDD_r16 != nil, ie.enhancedUL_TransientPeriod_r16 != nil, ie.sharedSpectrumChAccessParamsPerBand_v1640 != nil}
			for _, bit := range optionals_ext_8 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode handoverUTRA_FDD_r16 optional
			if ie.handoverUTRA_FDD_r16 != nil {
				if err = ie.handoverUTRA_FDD_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode handoverUTRA_FDD_r16", err)
				}
			}
			// encode enhancedUL_TransientPeriod_r16 optional
			if ie.enhancedUL_TransientPeriod_r16 != nil {
				if err = ie.enhancedUL_TransientPeriod_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enhancedUL_TransientPeriod_r16", err)
				}
			}
			// encode sharedSpectrumChAccessParamsPerBand_v1640 optional
			if ie.sharedSpectrumChAccessParamsPerBand_v1640 != nil {
				if err = ie.sharedSpectrumChAccessParamsPerBand_v1640.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sharedSpectrumChAccessParamsPerBand_v1640", err)
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
			optionals_ext_9 := []bool{ie.type1_PUSCH_RepetitionMultiSlots_v1650 != nil, ie.type2_PUSCH_RepetitionMultiSlots_v1650 != nil, ie.pusch_RepetitionMultiSlots_v1650 != nil, ie.configuredUL_GrantType1_v1650 != nil, ie.configuredUL_GrantType2_v1650 != nil, ie.sharedSpectrumChAccessParamsPerBand_v1650 != nil}
			for _, bit := range optionals_ext_9 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode type1_PUSCH_RepetitionMultiSlots_v1650 optional
			if ie.type1_PUSCH_RepetitionMultiSlots_v1650 != nil {
				if err = ie.type1_PUSCH_RepetitionMultiSlots_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode type1_PUSCH_RepetitionMultiSlots_v1650", err)
				}
			}
			// encode type2_PUSCH_RepetitionMultiSlots_v1650 optional
			if ie.type2_PUSCH_RepetitionMultiSlots_v1650 != nil {
				if err = ie.type2_PUSCH_RepetitionMultiSlots_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode type2_PUSCH_RepetitionMultiSlots_v1650", err)
				}
			}
			// encode pusch_RepetitionMultiSlots_v1650 optional
			if ie.pusch_RepetitionMultiSlots_v1650 != nil {
				if err = ie.pusch_RepetitionMultiSlots_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pusch_RepetitionMultiSlots_v1650", err)
				}
			}
			// encode configuredUL_GrantType1_v1650 optional
			if ie.configuredUL_GrantType1_v1650 != nil {
				if err = ie.configuredUL_GrantType1_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode configuredUL_GrantType1_v1650", err)
				}
			}
			// encode configuredUL_GrantType2_v1650 optional
			if ie.configuredUL_GrantType2_v1650 != nil {
				if err = ie.configuredUL_GrantType2_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode configuredUL_GrantType2_v1650", err)
				}
			}
			// encode sharedSpectrumChAccessParamsPerBand_v1650 optional
			if ie.sharedSpectrumChAccessParamsPerBand_v1650 != nil {
				if err = ie.sharedSpectrumChAccessParamsPerBand_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sharedSpectrumChAccessParamsPerBand_v1650", err)
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
			optionals_ext_10 := []bool{ie.enhancedSkipUplinkTxConfigured_v1660 != nil, ie.enhancedSkipUplinkTxDynamic_v1660 != nil}
			for _, bit := range optionals_ext_10 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode enhancedSkipUplinkTxConfigured_v1660 optional
			if ie.enhancedSkipUplinkTxConfigured_v1660 != nil {
				if err = ie.enhancedSkipUplinkTxConfigured_v1660.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enhancedSkipUplinkTxConfigured_v1660", err)
				}
			}
			// encode enhancedSkipUplinkTxDynamic_v1660 optional
			if ie.enhancedSkipUplinkTxDynamic_v1660 != nil {
				if err = ie.enhancedSkipUplinkTxDynamic_v1660.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enhancedSkipUplinkTxDynamic_v1660", err)
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
			optionals_ext_11 := []bool{ie.maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16 != nil, ie.txDiversity_r16 != nil}
			for _, bit := range optionals_ext_11 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16 optional
			if ie.maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16 != nil {
				if err = ie.maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16", err)
				}
			}
			// encode txDiversity_r16 optional
			if ie.txDiversity_r16 != nil {
				if err = ie.txDiversity_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode txDiversity_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 12
		if extBitmap[11] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 12
			optionals_ext_12 := []bool{ie.pdsch_1024QAM_FR1_r17 != nil, ie.ue_PowerClass_v1700 != nil, ie.fr2_2_AccessParamsPerBand_r17 != nil, ie.rlm_Relaxation_r17 != nil, ie.bfd_Relaxation_r17 != nil, ie.cg_SDT_r17 != nil, ie.locationBasedCondHandover_r17 != nil, ie.timeBasedCondHandover_r17 != nil, ie.eventA4BasedCondHandover_r17 != nil, ie.mn_InitiatedCondPSCellChangeNRDC_r17 != nil, ie.sn_InitiatedCondPSCellChangeNRDC_r17 != nil, ie.pdcch_SkippingWithoutSSSG_r17 != nil, ie.sssg_Switching_1BitInd_r17 != nil, ie.sssg_Switching_2BitInd_r17 != nil, ie.pdcch_SkippingWithSSSG_r17 != nil, ie.searchSpaceSetGrp_switchCap2_r17 != nil, ie.uplinkPreCompensation_r17 != nil, ie.uplink_TA_Reporting_r17 != nil, ie.max_HARQ_ProcessNumber_r17 != nil, ie.type2_HARQ_Codebook_r17 != nil, ie.type1_HARQ_Codebook_r17 != nil, ie.type3_HARQ_Codebook_r17 != nil, ie.ue_specific_K_Offset_r17 != nil, ie.multiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil, ie.multiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil, ie.parallelPRS_MeasRRC_Inactive_r17 != nil, ie.nr_UE_TxTEG_ID_MaxSupport_r17 != nil, ie.prs_ProcessingRRC_Inactive_r17 != nil, ie.prs_ProcessingWindowType1A_r17 != nil, ie.prs_ProcessingWindowType1B_r17 != nil, ie.prs_ProcessingWindowType2_r17 != nil, ie.srs_AllPosResourcesRRC_Inactive_r17 != nil, ie.olpc_SRS_PosRRC_Inactive_r17 != nil, ie.spatialRelationsSRS_PosRRC_Inactive_r17 != nil, ie.maxNumberPUSCH_TypeA_Repetition_r17 != nil, ie.puschTypeA_RepetitionsAvailSlot_r17 != nil, ie.tb_ProcessingMultiSlotPUSCH_r17 != nil, ie.tb_ProcessingRepMultiSlotPUSCH_r17 != nil, ie.maxDurationDMRS_Bundling_r17 != nil, ie.pusch_RepetitionMsg3_r17 != nil, ie.sharedSpectrumChAccessParamsPerBand_v1710 != nil, ie.parallelMeasurementWithoutRestriction_r17 != nil, ie.maxNumber_NGSO_SatellitesWithinOneSMTC_r17 != nil, ie.k1_RangeExtension_r17 != nil, ie.aperiodicCSI_RS_FastScellActivation_r17 != nil, ie.aperiodicCSI_RS_AdditionalBandwidth_r17 != nil, ie.bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17 != nil, ie.halfDuplexFDD_TypeA_RedCap_r17 != nil, ie.posSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 != nil, ie.channelBWs_DL_SCS_480kHz_FR2_2_r17 != nil, ie.channelBWs_UL_SCS_480kHz_FR2_2_r17 != nil, ie.channelBWs_DL_SCS_960kHz_FR2_2_r17 != nil, ie.channelBWs_UL_SCS_960kHz_FR2_2_r17 != nil, ie.ul_GapFR2_r17 != nil, ie.oneShotHARQ_feedbackTriggeredByDCI_1_2_r17 != nil, ie.oneShotHARQ_feedbackPhy_Priority_r17 != nil, ie.enhancedType3_HARQ_CodebookFeedback_r17 != nil, ie.triggeredHARQ_CodebookRetx_r17 != nil}
			for _, bit := range optionals_ext_12 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode pdsch_1024QAM_FR1_r17 optional
			if ie.pdsch_1024QAM_FR1_r17 != nil {
				if err = ie.pdsch_1024QAM_FR1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_1024QAM_FR1_r17", err)
				}
			}
			// encode ue_PowerClass_v1700 optional
			if ie.ue_PowerClass_v1700 != nil {
				if err = ie.ue_PowerClass_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ue_PowerClass_v1700", err)
				}
			}
			// encode fr2_2_AccessParamsPerBand_r17 optional
			if ie.fr2_2_AccessParamsPerBand_r17 != nil {
				if err = ie.fr2_2_AccessParamsPerBand_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode fr2_2_AccessParamsPerBand_r17", err)
				}
			}
			// encode rlm_Relaxation_r17 optional
			if ie.rlm_Relaxation_r17 != nil {
				if err = ie.rlm_Relaxation_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rlm_Relaxation_r17", err)
				}
			}
			// encode bfd_Relaxation_r17 optional
			if ie.bfd_Relaxation_r17 != nil {
				if err = ie.bfd_Relaxation_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode bfd_Relaxation_r17", err)
				}
			}
			// encode cg_SDT_r17 optional
			if ie.cg_SDT_r17 != nil {
				if err = ie.cg_SDT_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cg_SDT_r17", err)
				}
			}
			// encode locationBasedCondHandover_r17 optional
			if ie.locationBasedCondHandover_r17 != nil {
				if err = ie.locationBasedCondHandover_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode locationBasedCondHandover_r17", err)
				}
			}
			// encode timeBasedCondHandover_r17 optional
			if ie.timeBasedCondHandover_r17 != nil {
				if err = ie.timeBasedCondHandover_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode timeBasedCondHandover_r17", err)
				}
			}
			// encode eventA4BasedCondHandover_r17 optional
			if ie.eventA4BasedCondHandover_r17 != nil {
				if err = ie.eventA4BasedCondHandover_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode eventA4BasedCondHandover_r17", err)
				}
			}
			// encode mn_InitiatedCondPSCellChangeNRDC_r17 optional
			if ie.mn_InitiatedCondPSCellChangeNRDC_r17 != nil {
				if err = ie.mn_InitiatedCondPSCellChangeNRDC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mn_InitiatedCondPSCellChangeNRDC_r17", err)
				}
			}
			// encode sn_InitiatedCondPSCellChangeNRDC_r17 optional
			if ie.sn_InitiatedCondPSCellChangeNRDC_r17 != nil {
				if err = ie.sn_InitiatedCondPSCellChangeNRDC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sn_InitiatedCondPSCellChangeNRDC_r17", err)
				}
			}
			// encode pdcch_SkippingWithoutSSSG_r17 optional
			if ie.pdcch_SkippingWithoutSSSG_r17 != nil {
				if err = ie.pdcch_SkippingWithoutSSSG_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdcch_SkippingWithoutSSSG_r17", err)
				}
			}
			// encode sssg_Switching_1BitInd_r17 optional
			if ie.sssg_Switching_1BitInd_r17 != nil {
				if err = ie.sssg_Switching_1BitInd_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sssg_Switching_1BitInd_r17", err)
				}
			}
			// encode sssg_Switching_2BitInd_r17 optional
			if ie.sssg_Switching_2BitInd_r17 != nil {
				if err = ie.sssg_Switching_2BitInd_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sssg_Switching_2BitInd_r17", err)
				}
			}
			// encode pdcch_SkippingWithSSSG_r17 optional
			if ie.pdcch_SkippingWithSSSG_r17 != nil {
				if err = ie.pdcch_SkippingWithSSSG_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdcch_SkippingWithSSSG_r17", err)
				}
			}
			// encode searchSpaceSetGrp_switchCap2_r17 optional
			if ie.searchSpaceSetGrp_switchCap2_r17 != nil {
				if err = ie.searchSpaceSetGrp_switchCap2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode searchSpaceSetGrp_switchCap2_r17", err)
				}
			}
			// encode uplinkPreCompensation_r17 optional
			if ie.uplinkPreCompensation_r17 != nil {
				if err = ie.uplinkPreCompensation_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uplinkPreCompensation_r17", err)
				}
			}
			// encode uplink_TA_Reporting_r17 optional
			if ie.uplink_TA_Reporting_r17 != nil {
				if err = ie.uplink_TA_Reporting_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uplink_TA_Reporting_r17", err)
				}
			}
			// encode max_HARQ_ProcessNumber_r17 optional
			if ie.max_HARQ_ProcessNumber_r17 != nil {
				if err = ie.max_HARQ_ProcessNumber_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode max_HARQ_ProcessNumber_r17", err)
				}
			}
			// encode type2_HARQ_Codebook_r17 optional
			if ie.type2_HARQ_Codebook_r17 != nil {
				if err = ie.type2_HARQ_Codebook_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode type2_HARQ_Codebook_r17", err)
				}
			}
			// encode type1_HARQ_Codebook_r17 optional
			if ie.type1_HARQ_Codebook_r17 != nil {
				if err = ie.type1_HARQ_Codebook_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode type1_HARQ_Codebook_r17", err)
				}
			}
			// encode type3_HARQ_Codebook_r17 optional
			if ie.type3_HARQ_Codebook_r17 != nil {
				if err = ie.type3_HARQ_Codebook_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode type3_HARQ_Codebook_r17", err)
				}
			}
			// encode ue_specific_K_Offset_r17 optional
			if ie.ue_specific_K_Offset_r17 != nil {
				if err = ie.ue_specific_K_Offset_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ue_specific_K_Offset_r17", err)
				}
			}
			// encode multiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17 optional
			if ie.multiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil {
				if err = ie.multiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode multiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17", err)
				}
			}
			// encode multiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17 optional
			if ie.multiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17 != nil {
				if err = ie.multiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode multiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17", err)
				}
			}
			// encode parallelPRS_MeasRRC_Inactive_r17 optional
			if ie.parallelPRS_MeasRRC_Inactive_r17 != nil {
				if err = ie.parallelPRS_MeasRRC_Inactive_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode parallelPRS_MeasRRC_Inactive_r17", err)
				}
			}
			// encode nr_UE_TxTEG_ID_MaxSupport_r17 optional
			if ie.nr_UE_TxTEG_ID_MaxSupport_r17 != nil {
				if err = ie.nr_UE_TxTEG_ID_MaxSupport_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nr_UE_TxTEG_ID_MaxSupport_r17", err)
				}
			}
			// encode prs_ProcessingRRC_Inactive_r17 optional
			if ie.prs_ProcessingRRC_Inactive_r17 != nil {
				if err = ie.prs_ProcessingRRC_Inactive_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode prs_ProcessingRRC_Inactive_r17", err)
				}
			}
			// encode prs_ProcessingWindowType1A_r17 optional
			if ie.prs_ProcessingWindowType1A_r17 != nil {
				if err = ie.prs_ProcessingWindowType1A_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode prs_ProcessingWindowType1A_r17", err)
				}
			}
			// encode prs_ProcessingWindowType1B_r17 optional
			if ie.prs_ProcessingWindowType1B_r17 != nil {
				if err = ie.prs_ProcessingWindowType1B_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode prs_ProcessingWindowType1B_r17", err)
				}
			}
			// encode prs_ProcessingWindowType2_r17 optional
			if ie.prs_ProcessingWindowType2_r17 != nil {
				if err = ie.prs_ProcessingWindowType2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode prs_ProcessingWindowType2_r17", err)
				}
			}
			// encode srs_AllPosResourcesRRC_Inactive_r17 optional
			if ie.srs_AllPosResourcesRRC_Inactive_r17 != nil {
				if err = ie.srs_AllPosResourcesRRC_Inactive_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_AllPosResourcesRRC_Inactive_r17", err)
				}
			}
			// encode olpc_SRS_PosRRC_Inactive_r17 optional
			if ie.olpc_SRS_PosRRC_Inactive_r17 != nil {
				if err = ie.olpc_SRS_PosRRC_Inactive_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode olpc_SRS_PosRRC_Inactive_r17", err)
				}
			}
			// encode spatialRelationsSRS_PosRRC_Inactive_r17 optional
			if ie.spatialRelationsSRS_PosRRC_Inactive_r17 != nil {
				if err = ie.spatialRelationsSRS_PosRRC_Inactive_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode spatialRelationsSRS_PosRRC_Inactive_r17", err)
				}
			}
			// encode maxNumberPUSCH_TypeA_Repetition_r17 optional
			if ie.maxNumberPUSCH_TypeA_Repetition_r17 != nil {
				if err = ie.maxNumberPUSCH_TypeA_Repetition_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxNumberPUSCH_TypeA_Repetition_r17", err)
				}
			}
			// encode puschTypeA_RepetitionsAvailSlot_r17 optional
			if ie.puschTypeA_RepetitionsAvailSlot_r17 != nil {
				if err = ie.puschTypeA_RepetitionsAvailSlot_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode puschTypeA_RepetitionsAvailSlot_r17", err)
				}
			}
			// encode tb_ProcessingMultiSlotPUSCH_r17 optional
			if ie.tb_ProcessingMultiSlotPUSCH_r17 != nil {
				if err = ie.tb_ProcessingMultiSlotPUSCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode tb_ProcessingMultiSlotPUSCH_r17", err)
				}
			}
			// encode tb_ProcessingRepMultiSlotPUSCH_r17 optional
			if ie.tb_ProcessingRepMultiSlotPUSCH_r17 != nil {
				if err = ie.tb_ProcessingRepMultiSlotPUSCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode tb_ProcessingRepMultiSlotPUSCH_r17", err)
				}
			}
			// encode maxDurationDMRS_Bundling_r17 optional
			if ie.maxDurationDMRS_Bundling_r17 != nil {
				if err = ie.maxDurationDMRS_Bundling_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxDurationDMRS_Bundling_r17", err)
				}
			}
			// encode pusch_RepetitionMsg3_r17 optional
			if ie.pusch_RepetitionMsg3_r17 != nil {
				if err = ie.pusch_RepetitionMsg3_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pusch_RepetitionMsg3_r17", err)
				}
			}
			// encode sharedSpectrumChAccessParamsPerBand_v1710 optional
			if ie.sharedSpectrumChAccessParamsPerBand_v1710 != nil {
				if err = ie.sharedSpectrumChAccessParamsPerBand_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sharedSpectrumChAccessParamsPerBand_v1710", err)
				}
			}
			// encode parallelMeasurementWithoutRestriction_r17 optional
			if ie.parallelMeasurementWithoutRestriction_r17 != nil {
				if err = ie.parallelMeasurementWithoutRestriction_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode parallelMeasurementWithoutRestriction_r17", err)
				}
			}
			// encode maxNumber_NGSO_SatellitesWithinOneSMTC_r17 optional
			if ie.maxNumber_NGSO_SatellitesWithinOneSMTC_r17 != nil {
				if err = ie.maxNumber_NGSO_SatellitesWithinOneSMTC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxNumber_NGSO_SatellitesWithinOneSMTC_r17", err)
				}
			}
			// encode k1_RangeExtension_r17 optional
			if ie.k1_RangeExtension_r17 != nil {
				if err = ie.k1_RangeExtension_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode k1_RangeExtension_r17", err)
				}
			}
			// encode aperiodicCSI_RS_FastScellActivation_r17 optional
			if ie.aperiodicCSI_RS_FastScellActivation_r17 != nil {
				if err = ie.aperiodicCSI_RS_FastScellActivation_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode aperiodicCSI_RS_FastScellActivation_r17", err)
				}
			}
			// encode aperiodicCSI_RS_AdditionalBandwidth_r17 optional
			if ie.aperiodicCSI_RS_AdditionalBandwidth_r17 != nil {
				if err = ie.aperiodicCSI_RS_AdditionalBandwidth_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode aperiodicCSI_RS_AdditionalBandwidth_r17", err)
				}
			}
			// encode bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17 optional
			if ie.bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17 != nil {
				if err = ie.bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17", err)
				}
			}
			// encode halfDuplexFDD_TypeA_RedCap_r17 optional
			if ie.halfDuplexFDD_TypeA_RedCap_r17 != nil {
				if err = ie.halfDuplexFDD_TypeA_RedCap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode halfDuplexFDD_TypeA_RedCap_r17", err)
				}
			}
			// encode posSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 optional
			if ie.posSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 != nil {
				if err = ie.posSRS_RRC_Inactive_OutsideInitialUL_BWP_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode posSRS_RRC_Inactive_OutsideInitialUL_BWP_r17", err)
				}
			}
			// encode channelBWs_DL_SCS_480kHz_FR2_2_r17 optional
			if ie.channelBWs_DL_SCS_480kHz_FR2_2_r17 != nil {
				if err = extWriter.WriteBitString(ie.channelBWs_DL_SCS_480kHz_FR2_2_r17.Bytes, uint(ie.channelBWs_DL_SCS_480kHz_FR2_2_r17.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Encode channelBWs_DL_SCS_480kHz_FR2_2_r17", err)
				}
			}
			// encode channelBWs_UL_SCS_480kHz_FR2_2_r17 optional
			if ie.channelBWs_UL_SCS_480kHz_FR2_2_r17 != nil {
				if err = extWriter.WriteBitString(ie.channelBWs_UL_SCS_480kHz_FR2_2_r17.Bytes, uint(ie.channelBWs_UL_SCS_480kHz_FR2_2_r17.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Encode channelBWs_UL_SCS_480kHz_FR2_2_r17", err)
				}
			}
			// encode channelBWs_DL_SCS_960kHz_FR2_2_r17 optional
			if ie.channelBWs_DL_SCS_960kHz_FR2_2_r17 != nil {
				if err = extWriter.WriteBitString(ie.channelBWs_DL_SCS_960kHz_FR2_2_r17.Bytes, uint(ie.channelBWs_DL_SCS_960kHz_FR2_2_r17.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Encode channelBWs_DL_SCS_960kHz_FR2_2_r17", err)
				}
			}
			// encode channelBWs_UL_SCS_960kHz_FR2_2_r17 optional
			if ie.channelBWs_UL_SCS_960kHz_FR2_2_r17 != nil {
				if err = extWriter.WriteBitString(ie.channelBWs_UL_SCS_960kHz_FR2_2_r17.Bytes, uint(ie.channelBWs_UL_SCS_960kHz_FR2_2_r17.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Encode channelBWs_UL_SCS_960kHz_FR2_2_r17", err)
				}
			}
			// encode ul_GapFR2_r17 optional
			if ie.ul_GapFR2_r17 != nil {
				if err = ie.ul_GapFR2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_GapFR2_r17", err)
				}
			}
			// encode oneShotHARQ_feedbackTriggeredByDCI_1_2_r17 optional
			if ie.oneShotHARQ_feedbackTriggeredByDCI_1_2_r17 != nil {
				if err = ie.oneShotHARQ_feedbackTriggeredByDCI_1_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode oneShotHARQ_feedbackTriggeredByDCI_1_2_r17", err)
				}
			}
			// encode oneShotHARQ_feedbackPhy_Priority_r17 optional
			if ie.oneShotHARQ_feedbackPhy_Priority_r17 != nil {
				if err = ie.oneShotHARQ_feedbackPhy_Priority_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode oneShotHARQ_feedbackPhy_Priority_r17", err)
				}
			}
			// encode enhancedType3_HARQ_CodebookFeedback_r17 optional
			if ie.enhancedType3_HARQ_CodebookFeedback_r17 != nil {
				if err = ie.enhancedType3_HARQ_CodebookFeedback_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enhancedType3_HARQ_CodebookFeedback_r17", err)
				}
			}
			// encode triggeredHARQ_CodebookRetx_r17 optional
			if ie.triggeredHARQ_CodebookRetx_r17 != nil {
				if err = ie.triggeredHARQ_CodebookRetx_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode triggeredHARQ_CodebookRetx_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 13
		if extBitmap[12] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 13
			optionals_ext_13 := []bool{ie.ue_OneShotUL_TimingAdj_r17 != nil, ie.pucch_Repetition_F0_2_r17 != nil, ie.cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17 != nil, ie.mux_HARQ_ACK_DiffPriorities_r17 != nil, ie.ta_BasedPDC_NTN_SharedSpectrumChAccess_r17 != nil, ie.ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17 != nil, ie.maxNumberG_RNTI_r17 != nil, ie.dynamicMulticastDCI_Format4_2_r17 != nil, ie.maxModulationOrderForMulticast_r17 != nil, ie.dynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 != nil, ie.dynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17 != nil, ie.nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17 != nil, ie.ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17 != nil, ie.maxNumberG_CS_RNTI_r17 != nil, ie.re_LevelRateMatchingForMulticast_r17 != nil, ie.pdsch_1024QAM_2MIMO_FR1_r17 != nil, ie.prs_MeasurementWithoutMG_r17 != nil, ie.maxNumber_LEO_SatellitesPerCarrier_r17 != nil, len(ie.prs_ProcessingCapabilityOutsideMGinPPW_r17) > 0, ie.srs_SemiPersistent_PosResourcesRRC_Inactive_r17 != nil, ie.channelBWs_DL_SCS_120kHz_FR2_2_r17 != nil, ie.channelBWs_UL_SCS_120kHz_FR2_2_r17 != nil}
			for _, bit := range optionals_ext_13 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ue_OneShotUL_TimingAdj_r17 optional
			if ie.ue_OneShotUL_TimingAdj_r17 != nil {
				if err = ie.ue_OneShotUL_TimingAdj_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ue_OneShotUL_TimingAdj_r17", err)
				}
			}
			// encode pucch_Repetition_F0_2_r17 optional
			if ie.pucch_Repetition_F0_2_r17 != nil {
				if err = ie.pucch_Repetition_F0_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pucch_Repetition_F0_2_r17", err)
				}
			}
			// encode cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17 optional
			if ie.cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17 != nil {
				if err = ie.cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17", err)
				}
			}
			// encode mux_HARQ_ACK_DiffPriorities_r17 optional
			if ie.mux_HARQ_ACK_DiffPriorities_r17 != nil {
				if err = ie.mux_HARQ_ACK_DiffPriorities_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mux_HARQ_ACK_DiffPriorities_r17", err)
				}
			}
			// encode ta_BasedPDC_NTN_SharedSpectrumChAccess_r17 optional
			if ie.ta_BasedPDC_NTN_SharedSpectrumChAccess_r17 != nil {
				if err = ie.ta_BasedPDC_NTN_SharedSpectrumChAccess_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ta_BasedPDC_NTN_SharedSpectrumChAccess_r17", err)
				}
			}
			// encode ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17 optional
			if ie.ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17 != nil {
				if err = ie.ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17", err)
				}
			}
			// encode maxNumberG_RNTI_r17 optional
			if ie.maxNumberG_RNTI_r17 != nil {
				if err = extWriter.WriteInteger(*ie.maxNumberG_RNTI_r17, &uper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
					return utils.WrapError("Encode maxNumberG_RNTI_r17", err)
				}
			}
			// encode dynamicMulticastDCI_Format4_2_r17 optional
			if ie.dynamicMulticastDCI_Format4_2_r17 != nil {
				if err = ie.dynamicMulticastDCI_Format4_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dynamicMulticastDCI_Format4_2_r17", err)
				}
			}
			// encode maxModulationOrderForMulticast_r17 optional
			if ie.maxModulationOrderForMulticast_r17 != nil {
				if err = ie.maxModulationOrderForMulticast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxModulationOrderForMulticast_r17", err)
				}
			}
			// encode dynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 optional
			if ie.dynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 != nil {
				if err = ie.dynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17", err)
				}
			}
			// encode dynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17 optional
			if ie.dynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17 != nil {
				if err = ie.dynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17", err)
				}
			}
			// encode nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17 optional
			if ie.nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17 != nil {
				if err = ie.nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17", err)
				}
			}
			// encode ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17 optional
			if ie.ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17 != nil {
				if err = ie.ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17", err)
				}
			}
			// encode maxNumberG_CS_RNTI_r17 optional
			if ie.maxNumberG_CS_RNTI_r17 != nil {
				if err = extWriter.WriteInteger(*ie.maxNumberG_CS_RNTI_r17, &uper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
					return utils.WrapError("Encode maxNumberG_CS_RNTI_r17", err)
				}
			}
			// encode re_LevelRateMatchingForMulticast_r17 optional
			if ie.re_LevelRateMatchingForMulticast_r17 != nil {
				if err = ie.re_LevelRateMatchingForMulticast_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode re_LevelRateMatchingForMulticast_r17", err)
				}
			}
			// encode pdsch_1024QAM_2MIMO_FR1_r17 optional
			if ie.pdsch_1024QAM_2MIMO_FR1_r17 != nil {
				if err = ie.pdsch_1024QAM_2MIMO_FR1_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_1024QAM_2MIMO_FR1_r17", err)
				}
			}
			// encode prs_MeasurementWithoutMG_r17 optional
			if ie.prs_MeasurementWithoutMG_r17 != nil {
				if err = ie.prs_MeasurementWithoutMG_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode prs_MeasurementWithoutMG_r17", err)
				}
			}
			// encode maxNumber_LEO_SatellitesPerCarrier_r17 optional
			if ie.maxNumber_LEO_SatellitesPerCarrier_r17 != nil {
				if err = extWriter.WriteInteger(*ie.maxNumber_LEO_SatellitesPerCarrier_r17, &uper.Constraint{Lb: 3, Ub: 4}, false); err != nil {
					return utils.WrapError("Encode maxNumber_LEO_SatellitesPerCarrier_r17", err)
				}
			}
			// encode prs_ProcessingCapabilityOutsideMGinPPW_r17 optional
			if len(ie.prs_ProcessingCapabilityOutsideMGinPPW_r17) > 0 {
				tmp_prs_ProcessingCapabilityOutsideMGinPPW_r17 := utils.NewSequence[*PRS_ProcessingCapabilityOutsideMGinPPWperType_r17]([]*PRS_ProcessingCapabilityOutsideMGinPPWperType_r17{}, uper.Constraint{Lb: 1, Ub: 3}, false)
				for _, i := range ie.prs_ProcessingCapabilityOutsideMGinPPW_r17 {
					tmp_prs_ProcessingCapabilityOutsideMGinPPW_r17.Value = append(tmp_prs_ProcessingCapabilityOutsideMGinPPW_r17.Value, &i)
				}
				if err = tmp_prs_ProcessingCapabilityOutsideMGinPPW_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode prs_ProcessingCapabilityOutsideMGinPPW_r17", err)
				}
			}
			// encode srs_SemiPersistent_PosResourcesRRC_Inactive_r17 optional
			if ie.srs_SemiPersistent_PosResourcesRRC_Inactive_r17 != nil {
				if err = ie.srs_SemiPersistent_PosResourcesRRC_Inactive_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_SemiPersistent_PosResourcesRRC_Inactive_r17", err)
				}
			}
			// encode channelBWs_DL_SCS_120kHz_FR2_2_r17 optional
			if ie.channelBWs_DL_SCS_120kHz_FR2_2_r17 != nil {
				if err = extWriter.WriteBitString(ie.channelBWs_DL_SCS_120kHz_FR2_2_r17.Bytes, uint(ie.channelBWs_DL_SCS_120kHz_FR2_2_r17.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Encode channelBWs_DL_SCS_120kHz_FR2_2_r17", err)
				}
			}
			// encode channelBWs_UL_SCS_120kHz_FR2_2_r17 optional
			if ie.channelBWs_UL_SCS_120kHz_FR2_2_r17 != nil {
				if err = extWriter.WriteBitString(ie.channelBWs_UL_SCS_120kHz_FR2_2_r17.Bytes, uint(ie.channelBWs_UL_SCS_120kHz_FR2_2_r17.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Encode channelBWs_UL_SCS_120kHz_FR2_2_r17", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 14
		if extBitmap[13] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 14
			optionals_ext_14 := []bool{ie.dmrs_BundlingPUSCH_RepTypeA_r17 != nil, ie.dmrs_BundlingPUSCH_RepTypeB_r17 != nil, ie.dmrs_BundlingPUSCH_multiSlot_r17 != nil, ie.dmrs_BundlingPUCCH_Rep_r17 != nil, ie.interSlotFreqHopInterSlotBundlingPUSCH_r17 != nil, ie.interSlotFreqHopPUCCH_r17 != nil, ie.dmrs_BundlingRestart_r17 != nil, ie.dmrs_BundlingNonBackToBackTX_r17 != nil}
			for _, bit := range optionals_ext_14 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode dmrs_BundlingPUSCH_RepTypeA_r17 optional
			if ie.dmrs_BundlingPUSCH_RepTypeA_r17 != nil {
				if err = ie.dmrs_BundlingPUSCH_RepTypeA_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dmrs_BundlingPUSCH_RepTypeA_r17", err)
				}
			}
			// encode dmrs_BundlingPUSCH_RepTypeB_r17 optional
			if ie.dmrs_BundlingPUSCH_RepTypeB_r17 != nil {
				if err = ie.dmrs_BundlingPUSCH_RepTypeB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dmrs_BundlingPUSCH_RepTypeB_r17", err)
				}
			}
			// encode dmrs_BundlingPUSCH_multiSlot_r17 optional
			if ie.dmrs_BundlingPUSCH_multiSlot_r17 != nil {
				if err = ie.dmrs_BundlingPUSCH_multiSlot_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dmrs_BundlingPUSCH_multiSlot_r17", err)
				}
			}
			// encode dmrs_BundlingPUCCH_Rep_r17 optional
			if ie.dmrs_BundlingPUCCH_Rep_r17 != nil {
				if err = ie.dmrs_BundlingPUCCH_Rep_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dmrs_BundlingPUCCH_Rep_r17", err)
				}
			}
			// encode interSlotFreqHopInterSlotBundlingPUSCH_r17 optional
			if ie.interSlotFreqHopInterSlotBundlingPUSCH_r17 != nil {
				if err = ie.interSlotFreqHopInterSlotBundlingPUSCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode interSlotFreqHopInterSlotBundlingPUSCH_r17", err)
				}
			}
			// encode interSlotFreqHopPUCCH_r17 optional
			if ie.interSlotFreqHopPUCCH_r17 != nil {
				if err = ie.interSlotFreqHopPUCCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode interSlotFreqHopPUCCH_r17", err)
				}
			}
			// encode dmrs_BundlingRestart_r17 optional
			if ie.dmrs_BundlingRestart_r17 != nil {
				if err = ie.dmrs_BundlingRestart_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dmrs_BundlingRestart_r17", err)
				}
			}
			// encode dmrs_BundlingNonBackToBackTX_r17 optional
			if ie.dmrs_BundlingNonBackToBackTX_r17 != nil {
				if err = ie.dmrs_BundlingNonBackToBackTX_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dmrs_BundlingNonBackToBackTX_r17", err)
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

func (ie *BandNR) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var modifiedMPR_BehaviourPresent bool
	if modifiedMPR_BehaviourPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var mimo_ParametersPerBandPresent bool
	if mimo_ParametersPerBandPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var extendedCPPresent bool
	if extendedCPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var multipleTCIPresent bool
	if multipleTCIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var bwp_WithoutRestrictionPresent bool
	if bwp_WithoutRestrictionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var bwp_SameNumerologyPresent bool
	if bwp_SameNumerologyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var bwp_DiffNumerologyPresent bool
	if bwp_DiffNumerologyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var crossCarrierScheduling_SameSCSPresent bool
	if crossCarrierScheduling_SameSCSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_256QAM_FR2Present bool
	if pdsch_256QAM_FR2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_256QAMPresent bool
	if pusch_256QAMPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ue_PowerClassPresent bool
	if ue_PowerClassPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rateMatchingLTE_CRSPresent bool
	if rateMatchingLTE_CRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var channelBWs_DLPresent bool
	if channelBWs_DLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var channelBWs_ULPresent bool
	if channelBWs_ULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.bandNR.Decode(r); err != nil {
		return utils.WrapError("Decode bandNR", err)
	}
	if modifiedMPR_BehaviourPresent {
		var tmp_bs_modifiedMPR_Behaviour []byte
		var n_modifiedMPR_Behaviour uint
		if tmp_bs_modifiedMPR_Behaviour, n_modifiedMPR_Behaviour, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode modifiedMPR_Behaviour", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_modifiedMPR_Behaviour,
			NumBits: uint64(n_modifiedMPR_Behaviour),
		}
		ie.modifiedMPR_Behaviour = &tmp_bitstring
	}
	if mimo_ParametersPerBandPresent {
		ie.mimo_ParametersPerBand = new(MIMO_ParametersPerBand)
		if err = ie.mimo_ParametersPerBand.Decode(r); err != nil {
			return utils.WrapError("Decode mimo_ParametersPerBand", err)
		}
	}
	if extendedCPPresent {
		ie.extendedCP = new(BandNR_extendedCP)
		if err = ie.extendedCP.Decode(r); err != nil {
			return utils.WrapError("Decode extendedCP", err)
		}
	}
	if multipleTCIPresent {
		ie.multipleTCI = new(BandNR_multipleTCI)
		if err = ie.multipleTCI.Decode(r); err != nil {
			return utils.WrapError("Decode multipleTCI", err)
		}
	}
	if bwp_WithoutRestrictionPresent {
		ie.bwp_WithoutRestriction = new(BandNR_bwp_WithoutRestriction)
		if err = ie.bwp_WithoutRestriction.Decode(r); err != nil {
			return utils.WrapError("Decode bwp_WithoutRestriction", err)
		}
	}
	if bwp_SameNumerologyPresent {
		ie.bwp_SameNumerology = new(BandNR_bwp_SameNumerology)
		if err = ie.bwp_SameNumerology.Decode(r); err != nil {
			return utils.WrapError("Decode bwp_SameNumerology", err)
		}
	}
	if bwp_DiffNumerologyPresent {
		ie.bwp_DiffNumerology = new(BandNR_bwp_DiffNumerology)
		if err = ie.bwp_DiffNumerology.Decode(r); err != nil {
			return utils.WrapError("Decode bwp_DiffNumerology", err)
		}
	}
	if crossCarrierScheduling_SameSCSPresent {
		ie.crossCarrierScheduling_SameSCS = new(BandNR_crossCarrierScheduling_SameSCS)
		if err = ie.crossCarrierScheduling_SameSCS.Decode(r); err != nil {
			return utils.WrapError("Decode crossCarrierScheduling_SameSCS", err)
		}
	}
	if pdsch_256QAM_FR2Present {
		ie.pdsch_256QAM_FR2 = new(BandNR_pdsch_256QAM_FR2)
		if err = ie.pdsch_256QAM_FR2.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_256QAM_FR2", err)
		}
	}
	if pusch_256QAMPresent {
		ie.pusch_256QAM = new(BandNR_pusch_256QAM)
		if err = ie.pusch_256QAM.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_256QAM", err)
		}
	}
	if ue_PowerClassPresent {
		ie.ue_PowerClass = new(BandNR_ue_PowerClass)
		if err = ie.ue_PowerClass.Decode(r); err != nil {
			return utils.WrapError("Decode ue_PowerClass", err)
		}
	}
	if rateMatchingLTE_CRSPresent {
		ie.rateMatchingLTE_CRS = new(BandNR_rateMatchingLTE_CRS)
		if err = ie.rateMatchingLTE_CRS.Decode(r); err != nil {
			return utils.WrapError("Decode rateMatchingLTE_CRS", err)
		}
	}
	if channelBWs_DLPresent {
		ie.channelBWs_DL = new(BandNR_channelBWs_DL)
		if err = ie.channelBWs_DL.Decode(r); err != nil {
			return utils.WrapError("Decode channelBWs_DL", err)
		}
	}
	if channelBWs_ULPresent {
		ie.channelBWs_UL = new(BandNR_channelBWs_UL)
		if err = ie.channelBWs_UL.Decode(r); err != nil {
			return utils.WrapError("Decode channelBWs_UL", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 14 bits for 14 extension groups
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

			maxUplinkDutyCycle_PC2_FR1Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode maxUplinkDutyCycle_PC2_FR1 optional
			if maxUplinkDutyCycle_PC2_FR1Present {
				ie.maxUplinkDutyCycle_PC2_FR1 = new(BandNR_maxUplinkDutyCycle_PC2_FR1)
				if err = ie.maxUplinkDutyCycle_PC2_FR1.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxUplinkDutyCycle_PC2_FR1", err)
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

			pucch_SpatialRelInfoMAC_CEPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			powerBoosting_pi2BPSKPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode pucch_SpatialRelInfoMAC_CE optional
			if pucch_SpatialRelInfoMAC_CEPresent {
				ie.pucch_SpatialRelInfoMAC_CE = new(BandNR_pucch_SpatialRelInfoMAC_CE)
				if err = ie.pucch_SpatialRelInfoMAC_CE.Decode(extReader); err != nil {
					return utils.WrapError("Decode pucch_SpatialRelInfoMAC_CE", err)
				}
			}
			// decode powerBoosting_pi2BPSK optional
			if powerBoosting_pi2BPSKPresent {
				ie.powerBoosting_pi2BPSK = new(BandNR_powerBoosting_pi2BPSK)
				if err = ie.powerBoosting_pi2BPSK.Decode(extReader); err != nil {
					return utils.WrapError("Decode powerBoosting_pi2BPSK", err)
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

			maxUplinkDutyCycle_FR2Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode maxUplinkDutyCycle_FR2 optional
			if maxUplinkDutyCycle_FR2Present {
				ie.maxUplinkDutyCycle_FR2 = new(BandNR_maxUplinkDutyCycle_FR2)
				if err = ie.maxUplinkDutyCycle_FR2.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxUplinkDutyCycle_FR2", err)
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

			channelBWs_DL_v1590Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			channelBWs_UL_v1590Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode channelBWs_DL_v1590 optional
			if channelBWs_DL_v1590Present {
				ie.channelBWs_DL_v1590 = new(BandNR_channelBWs_DL_v1590)
				if err = ie.channelBWs_DL_v1590.Decode(extReader); err != nil {
					return utils.WrapError("Decode channelBWs_DL_v1590", err)
				}
			}
			// decode channelBWs_UL_v1590 optional
			if channelBWs_UL_v1590Present {
				ie.channelBWs_UL_v1590 = new(BandNR_channelBWs_UL_v1590)
				if err = ie.channelBWs_UL_v1590.Decode(extReader); err != nil {
					return utils.WrapError("Decode channelBWs_UL_v1590", err)
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

			asymmetricBandwidthCombinationSetPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode asymmetricBandwidthCombinationSet optional
			if asymmetricBandwidthCombinationSetPresent {
				var tmp_bs_asymmetricBandwidthCombinationSet []byte
				var n_asymmetricBandwidthCombinationSet uint
				if tmp_bs_asymmetricBandwidthCombinationSet, n_asymmetricBandwidthCombinationSet, err = extReader.ReadBitString(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
					return utils.WrapError("Decode asymmetricBandwidthCombinationSet", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_asymmetricBandwidthCombinationSet,
					NumBits: uint64(n_asymmetricBandwidthCombinationSet),
				}
				ie.asymmetricBandwidthCombinationSet = &tmp_bitstring
			}
		}
		// decode extension group 6
		if len(extBitmap) > 5 && extBitmap[5] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			sharedSpectrumChAccessParamsPerBand_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cancelOverlappingPUSCH_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			multipleRateMatchingEUTRA_CRS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			overlapRateMatchingEUTRA_CRS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_MappingTypeB_Alt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			oneSlotPeriodicTRS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			olpc_SRS_Pos_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			spatialRelationsSRS_Pos_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			simulSRS_MIMO_TransWithinBand_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			channelBW_DL_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			channelBW_UL_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			rasterShift7dot5_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ue_PowerClass_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			condHandover_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			condHandoverFailure_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			condHandoverTwoTriggerEvents_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			condPSCellChange_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			condPSCellChangeTwoTriggerEvents_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mpr_PowerBoost_FR2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			activeConfiguredGrant_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			jointReleaseConfiguredGrantType2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sps_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			jointReleaseSPS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			simulSRS_TransWithinBand_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			trs_AdditionalBandwidth_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			handoverIntraF_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sharedSpectrumChAccessParamsPerBand_r16 optional
			if sharedSpectrumChAccessParamsPerBand_r16Present {
				ie.sharedSpectrumChAccessParamsPerBand_r16 = new(SharedSpectrumChAccessParamsPerBand_r16)
				if err = ie.sharedSpectrumChAccessParamsPerBand_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sharedSpectrumChAccessParamsPerBand_r16", err)
				}
			}
			// decode cancelOverlappingPUSCH_r16 optional
			if cancelOverlappingPUSCH_r16Present {
				ie.cancelOverlappingPUSCH_r16 = new(BandNR_cancelOverlappingPUSCH_r16)
				if err = ie.cancelOverlappingPUSCH_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode cancelOverlappingPUSCH_r16", err)
				}
			}
			// decode multipleRateMatchingEUTRA_CRS_r16 optional
			if multipleRateMatchingEUTRA_CRS_r16Present {
				ie.multipleRateMatchingEUTRA_CRS_r16 = new(BandNR_multipleRateMatchingEUTRA_CRS_r16)
				if err = ie.multipleRateMatchingEUTRA_CRS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode multipleRateMatchingEUTRA_CRS_r16", err)
				}
			}
			// decode overlapRateMatchingEUTRA_CRS_r16 optional
			if overlapRateMatchingEUTRA_CRS_r16Present {
				ie.overlapRateMatchingEUTRA_CRS_r16 = new(BandNR_overlapRateMatchingEUTRA_CRS_r16)
				if err = ie.overlapRateMatchingEUTRA_CRS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode overlapRateMatchingEUTRA_CRS_r16", err)
				}
			}
			// decode pdsch_MappingTypeB_Alt_r16 optional
			if pdsch_MappingTypeB_Alt_r16Present {
				ie.pdsch_MappingTypeB_Alt_r16 = new(BandNR_pdsch_MappingTypeB_Alt_r16)
				if err = ie.pdsch_MappingTypeB_Alt_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_MappingTypeB_Alt_r16", err)
				}
			}
			// decode oneSlotPeriodicTRS_r16 optional
			if oneSlotPeriodicTRS_r16Present {
				ie.oneSlotPeriodicTRS_r16 = new(BandNR_oneSlotPeriodicTRS_r16)
				if err = ie.oneSlotPeriodicTRS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode oneSlotPeriodicTRS_r16", err)
				}
			}
			// decode olpc_SRS_Pos_r16 optional
			if olpc_SRS_Pos_r16Present {
				ie.olpc_SRS_Pos_r16 = new(OLPC_SRS_Pos_r16)
				if err = ie.olpc_SRS_Pos_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode olpc_SRS_Pos_r16", err)
				}
			}
			// decode spatialRelationsSRS_Pos_r16 optional
			if spatialRelationsSRS_Pos_r16Present {
				ie.spatialRelationsSRS_Pos_r16 = new(SpatialRelationsSRS_Pos_r16)
				if err = ie.spatialRelationsSRS_Pos_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode spatialRelationsSRS_Pos_r16", err)
				}
			}
			// decode simulSRS_MIMO_TransWithinBand_r16 optional
			if simulSRS_MIMO_TransWithinBand_r16Present {
				ie.simulSRS_MIMO_TransWithinBand_r16 = new(BandNR_simulSRS_MIMO_TransWithinBand_r16)
				if err = ie.simulSRS_MIMO_TransWithinBand_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode simulSRS_MIMO_TransWithinBand_r16", err)
				}
			}
			// decode channelBW_DL_IAB_r16 optional
			if channelBW_DL_IAB_r16Present {
				ie.channelBW_DL_IAB_r16 = new(BandNR_channelBW_DL_IAB_r16)
				if err = ie.channelBW_DL_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode channelBW_DL_IAB_r16", err)
				}
			}
			// decode channelBW_UL_IAB_r16 optional
			if channelBW_UL_IAB_r16Present {
				ie.channelBW_UL_IAB_r16 = new(BandNR_channelBW_UL_IAB_r16)
				if err = ie.channelBW_UL_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode channelBW_UL_IAB_r16", err)
				}
			}
			// decode rasterShift7dot5_IAB_r16 optional
			if rasterShift7dot5_IAB_r16Present {
				ie.rasterShift7dot5_IAB_r16 = new(BandNR_rasterShift7dot5_IAB_r16)
				if err = ie.rasterShift7dot5_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode rasterShift7dot5_IAB_r16", err)
				}
			}
			// decode ue_PowerClass_v1610 optional
			if ue_PowerClass_v1610Present {
				ie.ue_PowerClass_v1610 = new(BandNR_ue_PowerClass_v1610)
				if err = ie.ue_PowerClass_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode ue_PowerClass_v1610", err)
				}
			}
			// decode condHandover_r16 optional
			if condHandover_r16Present {
				ie.condHandover_r16 = new(BandNR_condHandover_r16)
				if err = ie.condHandover_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode condHandover_r16", err)
				}
			}
			// decode condHandoverFailure_r16 optional
			if condHandoverFailure_r16Present {
				ie.condHandoverFailure_r16 = new(BandNR_condHandoverFailure_r16)
				if err = ie.condHandoverFailure_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode condHandoverFailure_r16", err)
				}
			}
			// decode condHandoverTwoTriggerEvents_r16 optional
			if condHandoverTwoTriggerEvents_r16Present {
				ie.condHandoverTwoTriggerEvents_r16 = new(BandNR_condHandoverTwoTriggerEvents_r16)
				if err = ie.condHandoverTwoTriggerEvents_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode condHandoverTwoTriggerEvents_r16", err)
				}
			}
			// decode condPSCellChange_r16 optional
			if condPSCellChange_r16Present {
				ie.condPSCellChange_r16 = new(BandNR_condPSCellChange_r16)
				if err = ie.condPSCellChange_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode condPSCellChange_r16", err)
				}
			}
			// decode condPSCellChangeTwoTriggerEvents_r16 optional
			if condPSCellChangeTwoTriggerEvents_r16Present {
				ie.condPSCellChangeTwoTriggerEvents_r16 = new(BandNR_condPSCellChangeTwoTriggerEvents_r16)
				if err = ie.condPSCellChangeTwoTriggerEvents_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode condPSCellChangeTwoTriggerEvents_r16", err)
				}
			}
			// decode mpr_PowerBoost_FR2_r16 optional
			if mpr_PowerBoost_FR2_r16Present {
				ie.mpr_PowerBoost_FR2_r16 = new(BandNR_mpr_PowerBoost_FR2_r16)
				if err = ie.mpr_PowerBoost_FR2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode mpr_PowerBoost_FR2_r16", err)
				}
			}
			// decode activeConfiguredGrant_r16 optional
			if activeConfiguredGrant_r16Present {
				ie.activeConfiguredGrant_r16 = new(BandNR_activeConfiguredGrant_r16)
				if err = ie.activeConfiguredGrant_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode activeConfiguredGrant_r16", err)
				}
			}
			// decode jointReleaseConfiguredGrantType2_r16 optional
			if jointReleaseConfiguredGrantType2_r16Present {
				ie.jointReleaseConfiguredGrantType2_r16 = new(BandNR_jointReleaseConfiguredGrantType2_r16)
				if err = ie.jointReleaseConfiguredGrantType2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode jointReleaseConfiguredGrantType2_r16", err)
				}
			}
			// decode sps_r16 optional
			if sps_r16Present {
				ie.sps_r16 = new(BandNR_sps_r16)
				if err = ie.sps_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sps_r16", err)
				}
			}
			// decode jointReleaseSPS_r16 optional
			if jointReleaseSPS_r16Present {
				ie.jointReleaseSPS_r16 = new(BandNR_jointReleaseSPS_r16)
				if err = ie.jointReleaseSPS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode jointReleaseSPS_r16", err)
				}
			}
			// decode simulSRS_TransWithinBand_r16 optional
			if simulSRS_TransWithinBand_r16Present {
				ie.simulSRS_TransWithinBand_r16 = new(BandNR_simulSRS_TransWithinBand_r16)
				if err = ie.simulSRS_TransWithinBand_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode simulSRS_TransWithinBand_r16", err)
				}
			}
			// decode trs_AdditionalBandwidth_r16 optional
			if trs_AdditionalBandwidth_r16Present {
				ie.trs_AdditionalBandwidth_r16 = new(BandNR_trs_AdditionalBandwidth_r16)
				if err = ie.trs_AdditionalBandwidth_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode trs_AdditionalBandwidth_r16", err)
				}
			}
			// decode handoverIntraF_IAB_r16 optional
			if handoverIntraF_IAB_r16Present {
				ie.handoverIntraF_IAB_r16 = new(BandNR_handoverIntraF_IAB_r16)
				if err = ie.handoverIntraF_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode handoverIntraF_IAB_r16", err)
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

			simulTX_SRS_AntSwitchingIntraBandUL_CA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sharedSpectrumChAccessParamsPerBand_v1630Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode simulTX_SRS_AntSwitchingIntraBandUL_CA_r16 optional
			if simulTX_SRS_AntSwitchingIntraBandUL_CA_r16Present {
				ie.simulTX_SRS_AntSwitchingIntraBandUL_CA_r16 = new(SimulSRS_ForAntennaSwitching_r16)
				if err = ie.simulTX_SRS_AntSwitchingIntraBandUL_CA_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode simulTX_SRS_AntSwitchingIntraBandUL_CA_r16", err)
				}
			}
			// decode sharedSpectrumChAccessParamsPerBand_v1630 optional
			if sharedSpectrumChAccessParamsPerBand_v1630Present {
				ie.sharedSpectrumChAccessParamsPerBand_v1630 = new(SharedSpectrumChAccessParamsPerBand_v1630)
				if err = ie.sharedSpectrumChAccessParamsPerBand_v1630.Decode(extReader); err != nil {
					return utils.WrapError("Decode sharedSpectrumChAccessParamsPerBand_v1630", err)
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

			handoverUTRA_FDD_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			enhancedUL_TransientPeriod_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sharedSpectrumChAccessParamsPerBand_v1640Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode handoverUTRA_FDD_r16 optional
			if handoverUTRA_FDD_r16Present {
				ie.handoverUTRA_FDD_r16 = new(BandNR_handoverUTRA_FDD_r16)
				if err = ie.handoverUTRA_FDD_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode handoverUTRA_FDD_r16", err)
				}
			}
			// decode enhancedUL_TransientPeriod_r16 optional
			if enhancedUL_TransientPeriod_r16Present {
				ie.enhancedUL_TransientPeriod_r16 = new(BandNR_enhancedUL_TransientPeriod_r16)
				if err = ie.enhancedUL_TransientPeriod_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode enhancedUL_TransientPeriod_r16", err)
				}
			}
			// decode sharedSpectrumChAccessParamsPerBand_v1640 optional
			if sharedSpectrumChAccessParamsPerBand_v1640Present {
				ie.sharedSpectrumChAccessParamsPerBand_v1640 = new(SharedSpectrumChAccessParamsPerBand_v1640)
				if err = ie.sharedSpectrumChAccessParamsPerBand_v1640.Decode(extReader); err != nil {
					return utils.WrapError("Decode sharedSpectrumChAccessParamsPerBand_v1640", err)
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

			type1_PUSCH_RepetitionMultiSlots_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			type2_PUSCH_RepetitionMultiSlots_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pusch_RepetitionMultiSlots_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			configuredUL_GrantType1_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			configuredUL_GrantType2_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sharedSpectrumChAccessParamsPerBand_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode type1_PUSCH_RepetitionMultiSlots_v1650 optional
			if type1_PUSCH_RepetitionMultiSlots_v1650Present {
				ie.type1_PUSCH_RepetitionMultiSlots_v1650 = new(BandNR_type1_PUSCH_RepetitionMultiSlots_v1650)
				if err = ie.type1_PUSCH_RepetitionMultiSlots_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode type1_PUSCH_RepetitionMultiSlots_v1650", err)
				}
			}
			// decode type2_PUSCH_RepetitionMultiSlots_v1650 optional
			if type2_PUSCH_RepetitionMultiSlots_v1650Present {
				ie.type2_PUSCH_RepetitionMultiSlots_v1650 = new(BandNR_type2_PUSCH_RepetitionMultiSlots_v1650)
				if err = ie.type2_PUSCH_RepetitionMultiSlots_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode type2_PUSCH_RepetitionMultiSlots_v1650", err)
				}
			}
			// decode pusch_RepetitionMultiSlots_v1650 optional
			if pusch_RepetitionMultiSlots_v1650Present {
				ie.pusch_RepetitionMultiSlots_v1650 = new(BandNR_pusch_RepetitionMultiSlots_v1650)
				if err = ie.pusch_RepetitionMultiSlots_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode pusch_RepetitionMultiSlots_v1650", err)
				}
			}
			// decode configuredUL_GrantType1_v1650 optional
			if configuredUL_GrantType1_v1650Present {
				ie.configuredUL_GrantType1_v1650 = new(BandNR_configuredUL_GrantType1_v1650)
				if err = ie.configuredUL_GrantType1_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode configuredUL_GrantType1_v1650", err)
				}
			}
			// decode configuredUL_GrantType2_v1650 optional
			if configuredUL_GrantType2_v1650Present {
				ie.configuredUL_GrantType2_v1650 = new(BandNR_configuredUL_GrantType2_v1650)
				if err = ie.configuredUL_GrantType2_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode configuredUL_GrantType2_v1650", err)
				}
			}
			// decode sharedSpectrumChAccessParamsPerBand_v1650 optional
			if sharedSpectrumChAccessParamsPerBand_v1650Present {
				ie.sharedSpectrumChAccessParamsPerBand_v1650 = new(SharedSpectrumChAccessParamsPerBand_v1650)
				if err = ie.sharedSpectrumChAccessParamsPerBand_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode sharedSpectrumChAccessParamsPerBand_v1650", err)
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

			enhancedSkipUplinkTxConfigured_v1660Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			enhancedSkipUplinkTxDynamic_v1660Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode enhancedSkipUplinkTxConfigured_v1660 optional
			if enhancedSkipUplinkTxConfigured_v1660Present {
				ie.enhancedSkipUplinkTxConfigured_v1660 = new(BandNR_enhancedSkipUplinkTxConfigured_v1660)
				if err = ie.enhancedSkipUplinkTxConfigured_v1660.Decode(extReader); err != nil {
					return utils.WrapError("Decode enhancedSkipUplinkTxConfigured_v1660", err)
				}
			}
			// decode enhancedSkipUplinkTxDynamic_v1660 optional
			if enhancedSkipUplinkTxDynamic_v1660Present {
				ie.enhancedSkipUplinkTxDynamic_v1660 = new(BandNR_enhancedSkipUplinkTxDynamic_v1660)
				if err = ie.enhancedSkipUplinkTxDynamic_v1660.Decode(extReader); err != nil {
					return utils.WrapError("Decode enhancedSkipUplinkTxDynamic_v1660", err)
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

			maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			txDiversity_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16 optional
			if maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16Present {
				ie.maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16 = new(BandNR_maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16)
				if err = ie.maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxUplinkDutyCycle_PC1dot5_MPE_FR1_r16", err)
				}
			}
			// decode txDiversity_r16 optional
			if txDiversity_r16Present {
				ie.txDiversity_r16 = new(BandNR_txDiversity_r16)
				if err = ie.txDiversity_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode txDiversity_r16", err)
				}
			}
		}
		// decode extension group 12
		if len(extBitmap) > 11 && extBitmap[11] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			pdsch_1024QAM_FR1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ue_PowerClass_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			fr2_2_AccessParamsPerBand_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			rlm_Relaxation_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			bfd_Relaxation_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cg_SDT_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			locationBasedCondHandover_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			timeBasedCondHandover_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			eventA4BasedCondHandover_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mn_InitiatedCondPSCellChangeNRDC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sn_InitiatedCondPSCellChangeNRDC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdcch_SkippingWithoutSSSG_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sssg_Switching_1BitInd_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sssg_Switching_2BitInd_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdcch_SkippingWithSSSG_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			searchSpaceSetGrp_switchCap2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uplinkPreCompensation_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uplink_TA_Reporting_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			max_HARQ_ProcessNumber_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			type2_HARQ_Codebook_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			type1_HARQ_Codebook_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			type3_HARQ_Codebook_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ue_specific_K_Offset_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			multiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			multiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			parallelPRS_MeasRRC_Inactive_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nr_UE_TxTEG_ID_MaxSupport_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			prs_ProcessingRRC_Inactive_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			prs_ProcessingWindowType1A_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			prs_ProcessingWindowType1B_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			prs_ProcessingWindowType2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_AllPosResourcesRRC_Inactive_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			olpc_SRS_PosRRC_Inactive_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			spatialRelationsSRS_PosRRC_Inactive_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumberPUSCH_TypeA_Repetition_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			puschTypeA_RepetitionsAvailSlot_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			tb_ProcessingMultiSlotPUSCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			tb_ProcessingRepMultiSlotPUSCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxDurationDMRS_Bundling_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pusch_RepetitionMsg3_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sharedSpectrumChAccessParamsPerBand_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			parallelMeasurementWithoutRestriction_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumber_NGSO_SatellitesWithinOneSMTC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			k1_RangeExtension_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			aperiodicCSI_RS_FastScellActivation_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			aperiodicCSI_RS_AdditionalBandwidth_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			halfDuplexFDD_TypeA_RedCap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			posSRS_RRC_Inactive_OutsideInitialUL_BWP_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			channelBWs_DL_SCS_480kHz_FR2_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			channelBWs_UL_SCS_480kHz_FR2_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			channelBWs_DL_SCS_960kHz_FR2_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			channelBWs_UL_SCS_960kHz_FR2_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_GapFR2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			oneShotHARQ_feedbackTriggeredByDCI_1_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			oneShotHARQ_feedbackPhy_Priority_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			enhancedType3_HARQ_CodebookFeedback_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			triggeredHARQ_CodebookRetx_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode pdsch_1024QAM_FR1_r17 optional
			if pdsch_1024QAM_FR1_r17Present {
				ie.pdsch_1024QAM_FR1_r17 = new(BandNR_pdsch_1024QAM_FR1_r17)
				if err = ie.pdsch_1024QAM_FR1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_1024QAM_FR1_r17", err)
				}
			}
			// decode ue_PowerClass_v1700 optional
			if ue_PowerClass_v1700Present {
				ie.ue_PowerClass_v1700 = new(BandNR_ue_PowerClass_v1700)
				if err = ie.ue_PowerClass_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode ue_PowerClass_v1700", err)
				}
			}
			// decode fr2_2_AccessParamsPerBand_r17 optional
			if fr2_2_AccessParamsPerBand_r17Present {
				ie.fr2_2_AccessParamsPerBand_r17 = new(FR2_2_AccessParamsPerBand_r17)
				if err = ie.fr2_2_AccessParamsPerBand_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode fr2_2_AccessParamsPerBand_r17", err)
				}
			}
			// decode rlm_Relaxation_r17 optional
			if rlm_Relaxation_r17Present {
				ie.rlm_Relaxation_r17 = new(BandNR_rlm_Relaxation_r17)
				if err = ie.rlm_Relaxation_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode rlm_Relaxation_r17", err)
				}
			}
			// decode bfd_Relaxation_r17 optional
			if bfd_Relaxation_r17Present {
				ie.bfd_Relaxation_r17 = new(BandNR_bfd_Relaxation_r17)
				if err = ie.bfd_Relaxation_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode bfd_Relaxation_r17", err)
				}
			}
			// decode cg_SDT_r17 optional
			if cg_SDT_r17Present {
				ie.cg_SDT_r17 = new(BandNR_cg_SDT_r17)
				if err = ie.cg_SDT_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode cg_SDT_r17", err)
				}
			}
			// decode locationBasedCondHandover_r17 optional
			if locationBasedCondHandover_r17Present {
				ie.locationBasedCondHandover_r17 = new(BandNR_locationBasedCondHandover_r17)
				if err = ie.locationBasedCondHandover_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode locationBasedCondHandover_r17", err)
				}
			}
			// decode timeBasedCondHandover_r17 optional
			if timeBasedCondHandover_r17Present {
				ie.timeBasedCondHandover_r17 = new(BandNR_timeBasedCondHandover_r17)
				if err = ie.timeBasedCondHandover_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode timeBasedCondHandover_r17", err)
				}
			}
			// decode eventA4BasedCondHandover_r17 optional
			if eventA4BasedCondHandover_r17Present {
				ie.eventA4BasedCondHandover_r17 = new(BandNR_eventA4BasedCondHandover_r17)
				if err = ie.eventA4BasedCondHandover_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode eventA4BasedCondHandover_r17", err)
				}
			}
			// decode mn_InitiatedCondPSCellChangeNRDC_r17 optional
			if mn_InitiatedCondPSCellChangeNRDC_r17Present {
				ie.mn_InitiatedCondPSCellChangeNRDC_r17 = new(BandNR_mn_InitiatedCondPSCellChangeNRDC_r17)
				if err = ie.mn_InitiatedCondPSCellChangeNRDC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mn_InitiatedCondPSCellChangeNRDC_r17", err)
				}
			}
			// decode sn_InitiatedCondPSCellChangeNRDC_r17 optional
			if sn_InitiatedCondPSCellChangeNRDC_r17Present {
				ie.sn_InitiatedCondPSCellChangeNRDC_r17 = new(BandNR_sn_InitiatedCondPSCellChangeNRDC_r17)
				if err = ie.sn_InitiatedCondPSCellChangeNRDC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sn_InitiatedCondPSCellChangeNRDC_r17", err)
				}
			}
			// decode pdcch_SkippingWithoutSSSG_r17 optional
			if pdcch_SkippingWithoutSSSG_r17Present {
				ie.pdcch_SkippingWithoutSSSG_r17 = new(BandNR_pdcch_SkippingWithoutSSSG_r17)
				if err = ie.pdcch_SkippingWithoutSSSG_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdcch_SkippingWithoutSSSG_r17", err)
				}
			}
			// decode sssg_Switching_1BitInd_r17 optional
			if sssg_Switching_1BitInd_r17Present {
				ie.sssg_Switching_1BitInd_r17 = new(BandNR_sssg_Switching_1BitInd_r17)
				if err = ie.sssg_Switching_1BitInd_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sssg_Switching_1BitInd_r17", err)
				}
			}
			// decode sssg_Switching_2BitInd_r17 optional
			if sssg_Switching_2BitInd_r17Present {
				ie.sssg_Switching_2BitInd_r17 = new(BandNR_sssg_Switching_2BitInd_r17)
				if err = ie.sssg_Switching_2BitInd_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sssg_Switching_2BitInd_r17", err)
				}
			}
			// decode pdcch_SkippingWithSSSG_r17 optional
			if pdcch_SkippingWithSSSG_r17Present {
				ie.pdcch_SkippingWithSSSG_r17 = new(BandNR_pdcch_SkippingWithSSSG_r17)
				if err = ie.pdcch_SkippingWithSSSG_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdcch_SkippingWithSSSG_r17", err)
				}
			}
			// decode searchSpaceSetGrp_switchCap2_r17 optional
			if searchSpaceSetGrp_switchCap2_r17Present {
				ie.searchSpaceSetGrp_switchCap2_r17 = new(BandNR_searchSpaceSetGrp_switchCap2_r17)
				if err = ie.searchSpaceSetGrp_switchCap2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode searchSpaceSetGrp_switchCap2_r17", err)
				}
			}
			// decode uplinkPreCompensation_r17 optional
			if uplinkPreCompensation_r17Present {
				ie.uplinkPreCompensation_r17 = new(BandNR_uplinkPreCompensation_r17)
				if err = ie.uplinkPreCompensation_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode uplinkPreCompensation_r17", err)
				}
			}
			// decode uplink_TA_Reporting_r17 optional
			if uplink_TA_Reporting_r17Present {
				ie.uplink_TA_Reporting_r17 = new(BandNR_uplink_TA_Reporting_r17)
				if err = ie.uplink_TA_Reporting_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode uplink_TA_Reporting_r17", err)
				}
			}
			// decode max_HARQ_ProcessNumber_r17 optional
			if max_HARQ_ProcessNumber_r17Present {
				ie.max_HARQ_ProcessNumber_r17 = new(BandNR_max_HARQ_ProcessNumber_r17)
				if err = ie.max_HARQ_ProcessNumber_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode max_HARQ_ProcessNumber_r17", err)
				}
			}
			// decode type2_HARQ_Codebook_r17 optional
			if type2_HARQ_Codebook_r17Present {
				ie.type2_HARQ_Codebook_r17 = new(BandNR_type2_HARQ_Codebook_r17)
				if err = ie.type2_HARQ_Codebook_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode type2_HARQ_Codebook_r17", err)
				}
			}
			// decode type1_HARQ_Codebook_r17 optional
			if type1_HARQ_Codebook_r17Present {
				ie.type1_HARQ_Codebook_r17 = new(BandNR_type1_HARQ_Codebook_r17)
				if err = ie.type1_HARQ_Codebook_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode type1_HARQ_Codebook_r17", err)
				}
			}
			// decode type3_HARQ_Codebook_r17 optional
			if type3_HARQ_Codebook_r17Present {
				ie.type3_HARQ_Codebook_r17 = new(BandNR_type3_HARQ_Codebook_r17)
				if err = ie.type3_HARQ_Codebook_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode type3_HARQ_Codebook_r17", err)
				}
			}
			// decode ue_specific_K_Offset_r17 optional
			if ue_specific_K_Offset_r17Present {
				ie.ue_specific_K_Offset_r17 = new(BandNR_ue_specific_K_Offset_r17)
				if err = ie.ue_specific_K_Offset_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ue_specific_K_Offset_r17", err)
				}
			}
			// decode multiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17 optional
			if multiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17Present {
				ie.multiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17 = new(BandNR_multiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17)
				if err = ie.multiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode multiPDSCH_SingleDCI_FR2_1_SCS_120kHz_r17", err)
				}
			}
			// decode multiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17 optional
			if multiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17Present {
				ie.multiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17 = new(BandNR_multiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17)
				if err = ie.multiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode multiPUSCH_SingleDCI_FR2_1_SCS_120kHz_r17", err)
				}
			}
			// decode parallelPRS_MeasRRC_Inactive_r17 optional
			if parallelPRS_MeasRRC_Inactive_r17Present {
				ie.parallelPRS_MeasRRC_Inactive_r17 = new(BandNR_parallelPRS_MeasRRC_Inactive_r17)
				if err = ie.parallelPRS_MeasRRC_Inactive_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode parallelPRS_MeasRRC_Inactive_r17", err)
				}
			}
			// decode nr_UE_TxTEG_ID_MaxSupport_r17 optional
			if nr_UE_TxTEG_ID_MaxSupport_r17Present {
				ie.nr_UE_TxTEG_ID_MaxSupport_r17 = new(BandNR_nr_UE_TxTEG_ID_MaxSupport_r17)
				if err = ie.nr_UE_TxTEG_ID_MaxSupport_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode nr_UE_TxTEG_ID_MaxSupport_r17", err)
				}
			}
			// decode prs_ProcessingRRC_Inactive_r17 optional
			if prs_ProcessingRRC_Inactive_r17Present {
				ie.prs_ProcessingRRC_Inactive_r17 = new(BandNR_prs_ProcessingRRC_Inactive_r17)
				if err = ie.prs_ProcessingRRC_Inactive_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode prs_ProcessingRRC_Inactive_r17", err)
				}
			}
			// decode prs_ProcessingWindowType1A_r17 optional
			if prs_ProcessingWindowType1A_r17Present {
				ie.prs_ProcessingWindowType1A_r17 = new(BandNR_prs_ProcessingWindowType1A_r17)
				if err = ie.prs_ProcessingWindowType1A_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode prs_ProcessingWindowType1A_r17", err)
				}
			}
			// decode prs_ProcessingWindowType1B_r17 optional
			if prs_ProcessingWindowType1B_r17Present {
				ie.prs_ProcessingWindowType1B_r17 = new(BandNR_prs_ProcessingWindowType1B_r17)
				if err = ie.prs_ProcessingWindowType1B_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode prs_ProcessingWindowType1B_r17", err)
				}
			}
			// decode prs_ProcessingWindowType2_r17 optional
			if prs_ProcessingWindowType2_r17Present {
				ie.prs_ProcessingWindowType2_r17 = new(BandNR_prs_ProcessingWindowType2_r17)
				if err = ie.prs_ProcessingWindowType2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode prs_ProcessingWindowType2_r17", err)
				}
			}
			// decode srs_AllPosResourcesRRC_Inactive_r17 optional
			if srs_AllPosResourcesRRC_Inactive_r17Present {
				ie.srs_AllPosResourcesRRC_Inactive_r17 = new(SRS_AllPosResourcesRRC_Inactive_r17)
				if err = ie.srs_AllPosResourcesRRC_Inactive_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode srs_AllPosResourcesRRC_Inactive_r17", err)
				}
			}
			// decode olpc_SRS_PosRRC_Inactive_r17 optional
			if olpc_SRS_PosRRC_Inactive_r17Present {
				ie.olpc_SRS_PosRRC_Inactive_r17 = new(OLPC_SRS_Pos_r16)
				if err = ie.olpc_SRS_PosRRC_Inactive_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode olpc_SRS_PosRRC_Inactive_r17", err)
				}
			}
			// decode spatialRelationsSRS_PosRRC_Inactive_r17 optional
			if spatialRelationsSRS_PosRRC_Inactive_r17Present {
				ie.spatialRelationsSRS_PosRRC_Inactive_r17 = new(SpatialRelationsSRS_Pos_r16)
				if err = ie.spatialRelationsSRS_PosRRC_Inactive_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode spatialRelationsSRS_PosRRC_Inactive_r17", err)
				}
			}
			// decode maxNumberPUSCH_TypeA_Repetition_r17 optional
			if maxNumberPUSCH_TypeA_Repetition_r17Present {
				ie.maxNumberPUSCH_TypeA_Repetition_r17 = new(BandNR_maxNumberPUSCH_TypeA_Repetition_r17)
				if err = ie.maxNumberPUSCH_TypeA_Repetition_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxNumberPUSCH_TypeA_Repetition_r17", err)
				}
			}
			// decode puschTypeA_RepetitionsAvailSlot_r17 optional
			if puschTypeA_RepetitionsAvailSlot_r17Present {
				ie.puschTypeA_RepetitionsAvailSlot_r17 = new(BandNR_puschTypeA_RepetitionsAvailSlot_r17)
				if err = ie.puschTypeA_RepetitionsAvailSlot_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode puschTypeA_RepetitionsAvailSlot_r17", err)
				}
			}
			// decode tb_ProcessingMultiSlotPUSCH_r17 optional
			if tb_ProcessingMultiSlotPUSCH_r17Present {
				ie.tb_ProcessingMultiSlotPUSCH_r17 = new(BandNR_tb_ProcessingMultiSlotPUSCH_r17)
				if err = ie.tb_ProcessingMultiSlotPUSCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode tb_ProcessingMultiSlotPUSCH_r17", err)
				}
			}
			// decode tb_ProcessingRepMultiSlotPUSCH_r17 optional
			if tb_ProcessingRepMultiSlotPUSCH_r17Present {
				ie.tb_ProcessingRepMultiSlotPUSCH_r17 = new(BandNR_tb_ProcessingRepMultiSlotPUSCH_r17)
				if err = ie.tb_ProcessingRepMultiSlotPUSCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode tb_ProcessingRepMultiSlotPUSCH_r17", err)
				}
			}
			// decode maxDurationDMRS_Bundling_r17 optional
			if maxDurationDMRS_Bundling_r17Present {
				ie.maxDurationDMRS_Bundling_r17 = new(BandNR_maxDurationDMRS_Bundling_r17)
				if err = ie.maxDurationDMRS_Bundling_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxDurationDMRS_Bundling_r17", err)
				}
			}
			// decode pusch_RepetitionMsg3_r17 optional
			if pusch_RepetitionMsg3_r17Present {
				ie.pusch_RepetitionMsg3_r17 = new(BandNR_pusch_RepetitionMsg3_r17)
				if err = ie.pusch_RepetitionMsg3_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pusch_RepetitionMsg3_r17", err)
				}
			}
			// decode sharedSpectrumChAccessParamsPerBand_v1710 optional
			if sharedSpectrumChAccessParamsPerBand_v1710Present {
				ie.sharedSpectrumChAccessParamsPerBand_v1710 = new(SharedSpectrumChAccessParamsPerBand_v1710)
				if err = ie.sharedSpectrumChAccessParamsPerBand_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode sharedSpectrumChAccessParamsPerBand_v1710", err)
				}
			}
			// decode parallelMeasurementWithoutRestriction_r17 optional
			if parallelMeasurementWithoutRestriction_r17Present {
				ie.parallelMeasurementWithoutRestriction_r17 = new(BandNR_parallelMeasurementWithoutRestriction_r17)
				if err = ie.parallelMeasurementWithoutRestriction_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode parallelMeasurementWithoutRestriction_r17", err)
				}
			}
			// decode maxNumber_NGSO_SatellitesWithinOneSMTC_r17 optional
			if maxNumber_NGSO_SatellitesWithinOneSMTC_r17Present {
				ie.maxNumber_NGSO_SatellitesWithinOneSMTC_r17 = new(BandNR_maxNumber_NGSO_SatellitesWithinOneSMTC_r17)
				if err = ie.maxNumber_NGSO_SatellitesWithinOneSMTC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxNumber_NGSO_SatellitesWithinOneSMTC_r17", err)
				}
			}
			// decode k1_RangeExtension_r17 optional
			if k1_RangeExtension_r17Present {
				ie.k1_RangeExtension_r17 = new(BandNR_k1_RangeExtension_r17)
				if err = ie.k1_RangeExtension_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode k1_RangeExtension_r17", err)
				}
			}
			// decode aperiodicCSI_RS_FastScellActivation_r17 optional
			if aperiodicCSI_RS_FastScellActivation_r17Present {
				ie.aperiodicCSI_RS_FastScellActivation_r17 = new(BandNR_aperiodicCSI_RS_FastScellActivation_r17)
				if err = ie.aperiodicCSI_RS_FastScellActivation_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode aperiodicCSI_RS_FastScellActivation_r17", err)
				}
			}
			// decode aperiodicCSI_RS_AdditionalBandwidth_r17 optional
			if aperiodicCSI_RS_AdditionalBandwidth_r17Present {
				ie.aperiodicCSI_RS_AdditionalBandwidth_r17 = new(BandNR_aperiodicCSI_RS_AdditionalBandwidth_r17)
				if err = ie.aperiodicCSI_RS_AdditionalBandwidth_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode aperiodicCSI_RS_AdditionalBandwidth_r17", err)
				}
			}
			// decode bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17 optional
			if bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17Present {
				ie.bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17 = new(BandNR_bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17)
				if err = ie.bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode bwp_WithoutCD_SSB_OrNCD_SSB_RedCap_r17", err)
				}
			}
			// decode halfDuplexFDD_TypeA_RedCap_r17 optional
			if halfDuplexFDD_TypeA_RedCap_r17Present {
				ie.halfDuplexFDD_TypeA_RedCap_r17 = new(BandNR_halfDuplexFDD_TypeA_RedCap_r17)
				if err = ie.halfDuplexFDD_TypeA_RedCap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode halfDuplexFDD_TypeA_RedCap_r17", err)
				}
			}
			// decode posSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 optional
			if posSRS_RRC_Inactive_OutsideInitialUL_BWP_r17Present {
				ie.posSRS_RRC_Inactive_OutsideInitialUL_BWP_r17 = new(PosSRS_RRC_Inactive_OutsideInitialUL_BWP_r17)
				if err = ie.posSRS_RRC_Inactive_OutsideInitialUL_BWP_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode posSRS_RRC_Inactive_OutsideInitialUL_BWP_r17", err)
				}
			}
			// decode channelBWs_DL_SCS_480kHz_FR2_2_r17 optional
			if channelBWs_DL_SCS_480kHz_FR2_2_r17Present {
				var tmp_bs_channelBWs_DL_SCS_480kHz_FR2_2_r17 []byte
				var n_channelBWs_DL_SCS_480kHz_FR2_2_r17 uint
				if tmp_bs_channelBWs_DL_SCS_480kHz_FR2_2_r17, n_channelBWs_DL_SCS_480kHz_FR2_2_r17, err = extReader.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Decode channelBWs_DL_SCS_480kHz_FR2_2_r17", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_channelBWs_DL_SCS_480kHz_FR2_2_r17,
					NumBits: uint64(n_channelBWs_DL_SCS_480kHz_FR2_2_r17),
				}
				ie.channelBWs_DL_SCS_480kHz_FR2_2_r17 = &tmp_bitstring
			}
			// decode channelBWs_UL_SCS_480kHz_FR2_2_r17 optional
			if channelBWs_UL_SCS_480kHz_FR2_2_r17Present {
				var tmp_bs_channelBWs_UL_SCS_480kHz_FR2_2_r17 []byte
				var n_channelBWs_UL_SCS_480kHz_FR2_2_r17 uint
				if tmp_bs_channelBWs_UL_SCS_480kHz_FR2_2_r17, n_channelBWs_UL_SCS_480kHz_FR2_2_r17, err = extReader.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Decode channelBWs_UL_SCS_480kHz_FR2_2_r17", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_channelBWs_UL_SCS_480kHz_FR2_2_r17,
					NumBits: uint64(n_channelBWs_UL_SCS_480kHz_FR2_2_r17),
				}
				ie.channelBWs_UL_SCS_480kHz_FR2_2_r17 = &tmp_bitstring
			}
			// decode channelBWs_DL_SCS_960kHz_FR2_2_r17 optional
			if channelBWs_DL_SCS_960kHz_FR2_2_r17Present {
				var tmp_bs_channelBWs_DL_SCS_960kHz_FR2_2_r17 []byte
				var n_channelBWs_DL_SCS_960kHz_FR2_2_r17 uint
				if tmp_bs_channelBWs_DL_SCS_960kHz_FR2_2_r17, n_channelBWs_DL_SCS_960kHz_FR2_2_r17, err = extReader.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Decode channelBWs_DL_SCS_960kHz_FR2_2_r17", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_channelBWs_DL_SCS_960kHz_FR2_2_r17,
					NumBits: uint64(n_channelBWs_DL_SCS_960kHz_FR2_2_r17),
				}
				ie.channelBWs_DL_SCS_960kHz_FR2_2_r17 = &tmp_bitstring
			}
			// decode channelBWs_UL_SCS_960kHz_FR2_2_r17 optional
			if channelBWs_UL_SCS_960kHz_FR2_2_r17Present {
				var tmp_bs_channelBWs_UL_SCS_960kHz_FR2_2_r17 []byte
				var n_channelBWs_UL_SCS_960kHz_FR2_2_r17 uint
				if tmp_bs_channelBWs_UL_SCS_960kHz_FR2_2_r17, n_channelBWs_UL_SCS_960kHz_FR2_2_r17, err = extReader.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Decode channelBWs_UL_SCS_960kHz_FR2_2_r17", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_channelBWs_UL_SCS_960kHz_FR2_2_r17,
					NumBits: uint64(n_channelBWs_UL_SCS_960kHz_FR2_2_r17),
				}
				ie.channelBWs_UL_SCS_960kHz_FR2_2_r17 = &tmp_bitstring
			}
			// decode ul_GapFR2_r17 optional
			if ul_GapFR2_r17Present {
				ie.ul_GapFR2_r17 = new(BandNR_ul_GapFR2_r17)
				if err = ie.ul_GapFR2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_GapFR2_r17", err)
				}
			}
			// decode oneShotHARQ_feedbackTriggeredByDCI_1_2_r17 optional
			if oneShotHARQ_feedbackTriggeredByDCI_1_2_r17Present {
				ie.oneShotHARQ_feedbackTriggeredByDCI_1_2_r17 = new(BandNR_oneShotHARQ_feedbackTriggeredByDCI_1_2_r17)
				if err = ie.oneShotHARQ_feedbackTriggeredByDCI_1_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode oneShotHARQ_feedbackTriggeredByDCI_1_2_r17", err)
				}
			}
			// decode oneShotHARQ_feedbackPhy_Priority_r17 optional
			if oneShotHARQ_feedbackPhy_Priority_r17Present {
				ie.oneShotHARQ_feedbackPhy_Priority_r17 = new(BandNR_oneShotHARQ_feedbackPhy_Priority_r17)
				if err = ie.oneShotHARQ_feedbackPhy_Priority_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode oneShotHARQ_feedbackPhy_Priority_r17", err)
				}
			}
			// decode enhancedType3_HARQ_CodebookFeedback_r17 optional
			if enhancedType3_HARQ_CodebookFeedback_r17Present {
				ie.enhancedType3_HARQ_CodebookFeedback_r17 = new(BandNR_enhancedType3_HARQ_CodebookFeedback_r17)
				if err = ie.enhancedType3_HARQ_CodebookFeedback_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode enhancedType3_HARQ_CodebookFeedback_r17", err)
				}
			}
			// decode triggeredHARQ_CodebookRetx_r17 optional
			if triggeredHARQ_CodebookRetx_r17Present {
				ie.triggeredHARQ_CodebookRetx_r17 = new(BandNR_triggeredHARQ_CodebookRetx_r17)
				if err = ie.triggeredHARQ_CodebookRetx_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode triggeredHARQ_CodebookRetx_r17", err)
				}
			}
		}
		// decode extension group 13
		if len(extBitmap) > 12 && extBitmap[12] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ue_OneShotUL_TimingAdj_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pucch_Repetition_F0_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mux_HARQ_ACK_DiffPriorities_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ta_BasedPDC_NTN_SharedSpectrumChAccess_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumberG_RNTI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dynamicMulticastDCI_Format4_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxModulationOrderForMulticast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumberG_CS_RNTI_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			re_LevelRateMatchingForMulticast_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_1024QAM_2MIMO_FR1_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			prs_MeasurementWithoutMG_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumber_LEO_SatellitesPerCarrier_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			prs_ProcessingCapabilityOutsideMGinPPW_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			srs_SemiPersistent_PosResourcesRRC_Inactive_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			channelBWs_DL_SCS_120kHz_FR2_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			channelBWs_UL_SCS_120kHz_FR2_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ue_OneShotUL_TimingAdj_r17 optional
			if ue_OneShotUL_TimingAdj_r17Present {
				ie.ue_OneShotUL_TimingAdj_r17 = new(BandNR_ue_OneShotUL_TimingAdj_r17)
				if err = ie.ue_OneShotUL_TimingAdj_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ue_OneShotUL_TimingAdj_r17", err)
				}
			}
			// decode pucch_Repetition_F0_2_r17 optional
			if pucch_Repetition_F0_2_r17Present {
				ie.pucch_Repetition_F0_2_r17 = new(BandNR_pucch_Repetition_F0_2_r17)
				if err = ie.pucch_Repetition_F0_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pucch_Repetition_F0_2_r17", err)
				}
			}
			// decode cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17 optional
			if cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17Present {
				ie.cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17 = new(BandNR_cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17)
				if err = ie.cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode cqi_4_BitsSubbandNTN_SharedSpectrumChAccess_r17", err)
				}
			}
			// decode mux_HARQ_ACK_DiffPriorities_r17 optional
			if mux_HARQ_ACK_DiffPriorities_r17Present {
				ie.mux_HARQ_ACK_DiffPriorities_r17 = new(BandNR_mux_HARQ_ACK_DiffPriorities_r17)
				if err = ie.mux_HARQ_ACK_DiffPriorities_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mux_HARQ_ACK_DiffPriorities_r17", err)
				}
			}
			// decode ta_BasedPDC_NTN_SharedSpectrumChAccess_r17 optional
			if ta_BasedPDC_NTN_SharedSpectrumChAccess_r17Present {
				ie.ta_BasedPDC_NTN_SharedSpectrumChAccess_r17 = new(BandNR_ta_BasedPDC_NTN_SharedSpectrumChAccess_r17)
				if err = ie.ta_BasedPDC_NTN_SharedSpectrumChAccess_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ta_BasedPDC_NTN_SharedSpectrumChAccess_r17", err)
				}
			}
			// decode ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17 optional
			if ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17Present {
				ie.ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17 = new(BandNR_ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17)
				if err = ie.ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ack_NACK_FeedbackForMulticastWithDCI_Enabler_r17", err)
				}
			}
			// decode maxNumberG_RNTI_r17 optional
			if maxNumberG_RNTI_r17Present {
				var tmp_int_maxNumberG_RNTI_r17 int64
				if tmp_int_maxNumberG_RNTI_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
					return utils.WrapError("Decode maxNumberG_RNTI_r17", err)
				}
				ie.maxNumberG_RNTI_r17 = &tmp_int_maxNumberG_RNTI_r17
			}
			// decode dynamicMulticastDCI_Format4_2_r17 optional
			if dynamicMulticastDCI_Format4_2_r17Present {
				ie.dynamicMulticastDCI_Format4_2_r17 = new(BandNR_dynamicMulticastDCI_Format4_2_r17)
				if err = ie.dynamicMulticastDCI_Format4_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dynamicMulticastDCI_Format4_2_r17", err)
				}
			}
			// decode maxModulationOrderForMulticast_r17 optional
			if maxModulationOrderForMulticast_r17Present {
				ie.maxModulationOrderForMulticast_r17 = new(BandNR_maxModulationOrderForMulticast_r17)
				if err = ie.maxModulationOrderForMulticast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxModulationOrderForMulticast_r17", err)
				}
			}
			// decode dynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 optional
			if dynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17Present {
				ie.dynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17 = new(BandNR_dynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17)
				if err = ie.dynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dynamicSlotRepetitionMulticastTN_NonSharedSpectrumChAccess_r17", err)
				}
			}
			// decode dynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17 optional
			if dynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17Present {
				ie.dynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17 = new(BandNR_dynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17)
				if err = ie.dynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dynamicSlotRepetitionMulticastNTN_SharedSpectrumChAccess_r17", err)
				}
			}
			// decode nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17 optional
			if nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17Present {
				ie.nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17 = new(BandNR_nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17)
				if err = ie.nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode nack_OnlyFeedbackForMulticastWithDCI_Enabler_r17", err)
				}
			}
			// decode ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17 optional
			if ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17Present {
				ie.ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17 = new(BandNR_ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17)
				if err = ie.ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ack_NACK_FeedbackForSPS_MulticastWithDCI_Enabler_r17", err)
				}
			}
			// decode maxNumberG_CS_RNTI_r17 optional
			if maxNumberG_CS_RNTI_r17Present {
				var tmp_int_maxNumberG_CS_RNTI_r17 int64
				if tmp_int_maxNumberG_CS_RNTI_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 2, Ub: 8}, false); err != nil {
					return utils.WrapError("Decode maxNumberG_CS_RNTI_r17", err)
				}
				ie.maxNumberG_CS_RNTI_r17 = &tmp_int_maxNumberG_CS_RNTI_r17
			}
			// decode re_LevelRateMatchingForMulticast_r17 optional
			if re_LevelRateMatchingForMulticast_r17Present {
				ie.re_LevelRateMatchingForMulticast_r17 = new(BandNR_re_LevelRateMatchingForMulticast_r17)
				if err = ie.re_LevelRateMatchingForMulticast_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode re_LevelRateMatchingForMulticast_r17", err)
				}
			}
			// decode pdsch_1024QAM_2MIMO_FR1_r17 optional
			if pdsch_1024QAM_2MIMO_FR1_r17Present {
				ie.pdsch_1024QAM_2MIMO_FR1_r17 = new(BandNR_pdsch_1024QAM_2MIMO_FR1_r17)
				if err = ie.pdsch_1024QAM_2MIMO_FR1_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_1024QAM_2MIMO_FR1_r17", err)
				}
			}
			// decode prs_MeasurementWithoutMG_r17 optional
			if prs_MeasurementWithoutMG_r17Present {
				ie.prs_MeasurementWithoutMG_r17 = new(BandNR_prs_MeasurementWithoutMG_r17)
				if err = ie.prs_MeasurementWithoutMG_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode prs_MeasurementWithoutMG_r17", err)
				}
			}
			// decode maxNumber_LEO_SatellitesPerCarrier_r17 optional
			if maxNumber_LEO_SatellitesPerCarrier_r17Present {
				var tmp_int_maxNumber_LEO_SatellitesPerCarrier_r17 int64
				if tmp_int_maxNumber_LEO_SatellitesPerCarrier_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 3, Ub: 4}, false); err != nil {
					return utils.WrapError("Decode maxNumber_LEO_SatellitesPerCarrier_r17", err)
				}
				ie.maxNumber_LEO_SatellitesPerCarrier_r17 = &tmp_int_maxNumber_LEO_SatellitesPerCarrier_r17
			}
			// decode prs_ProcessingCapabilityOutsideMGinPPW_r17 optional
			if prs_ProcessingCapabilityOutsideMGinPPW_r17Present {
				tmp_prs_ProcessingCapabilityOutsideMGinPPW_r17 := utils.NewSequence[*PRS_ProcessingCapabilityOutsideMGinPPWperType_r17]([]*PRS_ProcessingCapabilityOutsideMGinPPWperType_r17{}, uper.Constraint{Lb: 1, Ub: 3}, false)
				fn_prs_ProcessingCapabilityOutsideMGinPPW_r17 := func() *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17 {
					return new(PRS_ProcessingCapabilityOutsideMGinPPWperType_r17)
				}
				if err = tmp_prs_ProcessingCapabilityOutsideMGinPPW_r17.Decode(extReader, fn_prs_ProcessingCapabilityOutsideMGinPPW_r17); err != nil {
					return utils.WrapError("Decode prs_ProcessingCapabilityOutsideMGinPPW_r17", err)
				}
				ie.prs_ProcessingCapabilityOutsideMGinPPW_r17 = []PRS_ProcessingCapabilityOutsideMGinPPWperType_r17{}
				for _, i := range tmp_prs_ProcessingCapabilityOutsideMGinPPW_r17.Value {
					ie.prs_ProcessingCapabilityOutsideMGinPPW_r17 = append(ie.prs_ProcessingCapabilityOutsideMGinPPW_r17, *i)
				}
			}
			// decode srs_SemiPersistent_PosResourcesRRC_Inactive_r17 optional
			if srs_SemiPersistent_PosResourcesRRC_Inactive_r17Present {
				ie.srs_SemiPersistent_PosResourcesRRC_Inactive_r17 = new(BandNR_srs_SemiPersistent_PosResourcesRRC_Inactive_r17)
				if err = ie.srs_SemiPersistent_PosResourcesRRC_Inactive_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode srs_SemiPersistent_PosResourcesRRC_Inactive_r17", err)
				}
			}
			// decode channelBWs_DL_SCS_120kHz_FR2_2_r17 optional
			if channelBWs_DL_SCS_120kHz_FR2_2_r17Present {
				var tmp_bs_channelBWs_DL_SCS_120kHz_FR2_2_r17 []byte
				var n_channelBWs_DL_SCS_120kHz_FR2_2_r17 uint
				if tmp_bs_channelBWs_DL_SCS_120kHz_FR2_2_r17, n_channelBWs_DL_SCS_120kHz_FR2_2_r17, err = extReader.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Decode channelBWs_DL_SCS_120kHz_FR2_2_r17", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_channelBWs_DL_SCS_120kHz_FR2_2_r17,
					NumBits: uint64(n_channelBWs_DL_SCS_120kHz_FR2_2_r17),
				}
				ie.channelBWs_DL_SCS_120kHz_FR2_2_r17 = &tmp_bitstring
			}
			// decode channelBWs_UL_SCS_120kHz_FR2_2_r17 optional
			if channelBWs_UL_SCS_120kHz_FR2_2_r17Present {
				var tmp_bs_channelBWs_UL_SCS_120kHz_FR2_2_r17 []byte
				var n_channelBWs_UL_SCS_120kHz_FR2_2_r17 uint
				if tmp_bs_channelBWs_UL_SCS_120kHz_FR2_2_r17, n_channelBWs_UL_SCS_120kHz_FR2_2_r17, err = extReader.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
					return utils.WrapError("Decode channelBWs_UL_SCS_120kHz_FR2_2_r17", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_channelBWs_UL_SCS_120kHz_FR2_2_r17,
					NumBits: uint64(n_channelBWs_UL_SCS_120kHz_FR2_2_r17),
				}
				ie.channelBWs_UL_SCS_120kHz_FR2_2_r17 = &tmp_bitstring
			}
		}
		// decode extension group 14
		if len(extBitmap) > 13 && extBitmap[13] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			dmrs_BundlingPUSCH_RepTypeA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dmrs_BundlingPUSCH_RepTypeB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dmrs_BundlingPUSCH_multiSlot_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dmrs_BundlingPUCCH_Rep_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			interSlotFreqHopInterSlotBundlingPUSCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			interSlotFreqHopPUCCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dmrs_BundlingRestart_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dmrs_BundlingNonBackToBackTX_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode dmrs_BundlingPUSCH_RepTypeA_r17 optional
			if dmrs_BundlingPUSCH_RepTypeA_r17Present {
				ie.dmrs_BundlingPUSCH_RepTypeA_r17 = new(BandNR_dmrs_BundlingPUSCH_RepTypeA_r17)
				if err = ie.dmrs_BundlingPUSCH_RepTypeA_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dmrs_BundlingPUSCH_RepTypeA_r17", err)
				}
			}
			// decode dmrs_BundlingPUSCH_RepTypeB_r17 optional
			if dmrs_BundlingPUSCH_RepTypeB_r17Present {
				ie.dmrs_BundlingPUSCH_RepTypeB_r17 = new(BandNR_dmrs_BundlingPUSCH_RepTypeB_r17)
				if err = ie.dmrs_BundlingPUSCH_RepTypeB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dmrs_BundlingPUSCH_RepTypeB_r17", err)
				}
			}
			// decode dmrs_BundlingPUSCH_multiSlot_r17 optional
			if dmrs_BundlingPUSCH_multiSlot_r17Present {
				ie.dmrs_BundlingPUSCH_multiSlot_r17 = new(BandNR_dmrs_BundlingPUSCH_multiSlot_r17)
				if err = ie.dmrs_BundlingPUSCH_multiSlot_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dmrs_BundlingPUSCH_multiSlot_r17", err)
				}
			}
			// decode dmrs_BundlingPUCCH_Rep_r17 optional
			if dmrs_BundlingPUCCH_Rep_r17Present {
				ie.dmrs_BundlingPUCCH_Rep_r17 = new(BandNR_dmrs_BundlingPUCCH_Rep_r17)
				if err = ie.dmrs_BundlingPUCCH_Rep_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dmrs_BundlingPUCCH_Rep_r17", err)
				}
			}
			// decode interSlotFreqHopInterSlotBundlingPUSCH_r17 optional
			if interSlotFreqHopInterSlotBundlingPUSCH_r17Present {
				ie.interSlotFreqHopInterSlotBundlingPUSCH_r17 = new(BandNR_interSlotFreqHopInterSlotBundlingPUSCH_r17)
				if err = ie.interSlotFreqHopInterSlotBundlingPUSCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode interSlotFreqHopInterSlotBundlingPUSCH_r17", err)
				}
			}
			// decode interSlotFreqHopPUCCH_r17 optional
			if interSlotFreqHopPUCCH_r17Present {
				ie.interSlotFreqHopPUCCH_r17 = new(BandNR_interSlotFreqHopPUCCH_r17)
				if err = ie.interSlotFreqHopPUCCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode interSlotFreqHopPUCCH_r17", err)
				}
			}
			// decode dmrs_BundlingRestart_r17 optional
			if dmrs_BundlingRestart_r17Present {
				ie.dmrs_BundlingRestart_r17 = new(BandNR_dmrs_BundlingRestart_r17)
				if err = ie.dmrs_BundlingRestart_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dmrs_BundlingRestart_r17", err)
				}
			}
			// decode dmrs_BundlingNonBackToBackTX_r17 optional
			if dmrs_BundlingNonBackToBackTX_r17Present {
				ie.dmrs_BundlingNonBackToBackTX_r17 = new(BandNR_dmrs_BundlingNonBackToBackTX_r17)
				if err = ie.dmrs_BundlingNonBackToBackTX_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode dmrs_BundlingNonBackToBackTX_r17", err)
				}
			}
		}
	}
	return nil
}
