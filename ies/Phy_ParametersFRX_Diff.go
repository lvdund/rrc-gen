package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersFRX_Diff struct {
	dynamicSFI                                        *Phy_ParametersFRX_Diff_dynamicSFI                                        `optional`
	dummy1                                            *uper.BitString                                                           `lb:2,ub:2,optional`
	twoFL_DMRS                                        *uper.BitString                                                           `lb:2,ub:2,optional`
	dummy2                                            *uper.BitString                                                           `lb:2,ub:2,optional`
	dummy3                                            *uper.BitString                                                           `lb:2,ub:2,optional`
	supportedDMRS_TypeDL                              *Phy_ParametersFRX_Diff_supportedDMRS_TypeDL                              `optional`
	supportedDMRS_TypeUL                              *Phy_ParametersFRX_Diff_supportedDMRS_TypeUL                              `optional`
	semiOpenLoopCSI                                   *Phy_ParametersFRX_Diff_semiOpenLoopCSI                                   `optional`
	csi_ReportWithoutPMI                              *Phy_ParametersFRX_Diff_csi_ReportWithoutPMI                              `optional`
	csi_ReportWithoutCQI                              *Phy_ParametersFRX_Diff_csi_ReportWithoutCQI                              `optional`
	onePortsPTRS                                      *uper.BitString                                                           `lb:2,ub:2,optional`
	twoPUCCH_F0_2_ConsecSymbols                       *Phy_ParametersFRX_Diff_twoPUCCH_F0_2_ConsecSymbols                       `optional`
	pucch_F2_WithFH                                   *Phy_ParametersFRX_Diff_pucch_F2_WithFH                                   `optional`
	pucch_F3_WithFH                                   *Phy_ParametersFRX_Diff_pucch_F3_WithFH                                   `optional`
	pucch_F4_WithFH                                   *Phy_ParametersFRX_Diff_pucch_F4_WithFH                                   `optional`
	pucch_F0_2WithoutFH                               *Phy_ParametersFRX_Diff_pucch_F0_2WithoutFH                               `optional`
	pucch_F1_3_4WithoutFH                             *Phy_ParametersFRX_Diff_pucch_F1_3_4WithoutFH                             `optional`
	mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot            *Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot            `optional`
	uci_CodeBlockSegmentation                         *Phy_ParametersFRX_Diff_uci_CodeBlockSegmentation                         `optional`
	onePUCCH_LongAndShortFormat                       *Phy_ParametersFRX_Diff_onePUCCH_LongAndShortFormat                       `optional`
	twoPUCCH_AnyOthersInSlot                          *Phy_ParametersFRX_Diff_twoPUCCH_AnyOthersInSlot                          `optional`
	intraSlotFreqHopping_PUSCH                        *Phy_ParametersFRX_Diff_intraSlotFreqHopping_PUSCH                        `optional`
	pusch_LBRM                                        *Phy_ParametersFRX_Diff_pusch_LBRM                                        `optional`
	pdcch_BlindDetectionCA                            *int64                                                                    `lb:4,ub:16,optional`
	tpc_PUSCH_RNTI                                    *Phy_ParametersFRX_Diff_tpc_PUSCH_RNTI                                    `optional`
	tpc_PUCCH_RNTI                                    *Phy_ParametersFRX_Diff_tpc_PUCCH_RNTI                                    `optional`
	tpc_SRS_RNTI                                      *Phy_ParametersFRX_Diff_tpc_SRS_RNTI                                      `optional`
	absoluteTPC_Command                               *Phy_ParametersFRX_Diff_absoluteTPC_Command                               `optional`
	twoDifferentTPC_Loop_PUSCH                        *Phy_ParametersFRX_Diff_twoDifferentTPC_Loop_PUSCH                        `optional`
	twoDifferentTPC_Loop_PUCCH                        *Phy_ParametersFRX_Diff_twoDifferentTPC_Loop_PUCCH                        `optional`
	pusch_HalfPi_BPSK                                 *Phy_ParametersFRX_Diff_pusch_HalfPi_BPSK                                 `optional`
	pucch_F3_4_HalfPi_BPSK                            *Phy_ParametersFRX_Diff_pucch_F3_4_HalfPi_BPSK                            `optional`
	almostContiguousCP_OFDM_UL                        *Phy_ParametersFRX_Diff_almostContiguousCP_OFDM_UL                        `optional`
	sp_CSI_RS                                         *Phy_ParametersFRX_Diff_sp_CSI_RS                                         `optional`
	sp_CSI_IM                                         *Phy_ParametersFRX_Diff_sp_CSI_IM                                         `optional`
	tdd_MultiDL_UL_SwitchPerSlot                      *Phy_ParametersFRX_Diff_tdd_MultiDL_UL_SwitchPerSlot                      `optional`
	multipleCORESET                                   *Phy_ParametersFRX_Diff_multipleCORESET                                   `optional`
	csi_RS_IM_ReceptionForFeedback                    *CSI_RS_IM_ReceptionForFeedback                                           `optional,ext-1`
	csi_RS_ProcFrameworkForSRS                        *CSI_RS_ProcFrameworkForSRS                                               `optional,ext-1`
	csi_ReportFramework                               *CSI_ReportFramework                                                      `optional,ext-1`
	mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot             *Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot             `optional,ext-1`
	mux_SR_HARQ_ACK_PUCCH                             *Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_PUCCH                             `optional,ext-1`
	mux_MultipleGroupCtrlCH_Overlap                   *Phy_ParametersFRX_Diff_mux_MultipleGroupCtrlCH_Overlap                   `optional,ext-1`
	dl_SchedulingOffset_PDSCH_TypeA                   *Phy_ParametersFRX_Diff_dl_SchedulingOffset_PDSCH_TypeA                   `optional,ext-1`
	dl_SchedulingOffset_PDSCH_TypeB                   *Phy_ParametersFRX_Diff_dl_SchedulingOffset_PDSCH_TypeB                   `optional,ext-1`
	ul_SchedulingOffset                               *Phy_ParametersFRX_Diff_ul_SchedulingOffset                               `optional,ext-1`
	dl_64QAM_MCS_TableAlt                             *Phy_ParametersFRX_Diff_dl_64QAM_MCS_TableAlt                             `optional,ext-1`
	ul_64QAM_MCS_TableAlt                             *Phy_ParametersFRX_Diff_ul_64QAM_MCS_TableAlt                             `optional,ext-1`
	cqi_TableAlt                                      *Phy_ParametersFRX_Diff_cqi_TableAlt                                      `optional,ext-1`
	oneFL_DMRS_TwoAdditionalDMRS_UL                   *Phy_ParametersFRX_Diff_oneFL_DMRS_TwoAdditionalDMRS_UL                   `optional,ext-1`
	twoFL_DMRS_TwoAdditionalDMRS_UL                   *Phy_ParametersFRX_Diff_twoFL_DMRS_TwoAdditionalDMRS_UL                   `optional,ext-1`
	oneFL_DMRS_ThreeAdditionalDMRS_UL                 *Phy_ParametersFRX_Diff_oneFL_DMRS_ThreeAdditionalDMRS_UL                 `optional,ext-1`
	pdcch_BlindDetectionNRDC                          *Phy_ParametersFRX_Diff_pdcch_BlindDetectionNRDC                          `lb:1,ub:15,optional,ext-2`
	mux_HARQ_ACK_PUSCH_DiffSymbol                     *Phy_ParametersFRX_Diff_mux_HARQ_ACK_PUSCH_DiffSymbol                     `optional,ext-2`
	type1_HARQ_ACK_Codebook_r16                       *Phy_ParametersFRX_Diff_type1_HARQ_ACK_Codebook_r16                       `optional,ext-3`
	enhancedPowerControl_r16                          *Phy_ParametersFRX_Diff_enhancedPowerControl_r16                          `optional,ext-3`
	simultaneousTCI_ActMultipleCC_r16                 *Phy_ParametersFRX_Diff_simultaneousTCI_ActMultipleCC_r16                 `optional,ext-3`
	simultaneousSpatialRelationMultipleCC_r16         *Phy_ParametersFRX_Diff_simultaneousSpatialRelationMultipleCC_r16         `optional,ext-3`
	cli_RSSI_FDM_DL_r16                               *Phy_ParametersFRX_Diff_cli_RSSI_FDM_DL_r16                               `optional,ext-3`
	cli_SRS_RSRP_FDM_DL_r16                           *Phy_ParametersFRX_Diff_cli_SRS_RSRP_FDM_DL_r16                           `optional,ext-3`
	maxLayersMIMO_Adaptation_r16                      *Phy_ParametersFRX_Diff_maxLayersMIMO_Adaptation_r16                      `optional,ext-3`
	aggregationFactorSPS_DL_r16                       *Phy_ParametersFRX_Diff_aggregationFactorSPS_DL_r16                       `optional,ext-3`
	maxTotalResourcesForOneFreqRange_r16              *Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16              `optional,ext-3`
	csi_ReportFrameworkExt_r16                        *CSI_ReportFrameworkExt_r16                                               `optional,ext-3`
	twoTCI_Act_servingCellInCC_List_r16               *Phy_ParametersFRX_Diff_twoTCI_Act_servingCellInCC_List_r16               `optional,ext-4`
	cri_RI_CQI_WithoutNon_PMI_PortInd_r16             *Phy_ParametersFRX_Diff_cri_RI_CQI_WithoutNon_PMI_PortInd_r16             `optional,ext-5`
	cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 *Phy_ParametersFRX_Diff_cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 `optional,ext-6`
}

func (ie *Phy_ParametersFRX_Diff) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.csi_RS_IM_ReceptionForFeedback != nil || ie.csi_RS_ProcFrameworkForSRS != nil || ie.csi_ReportFramework != nil || ie.mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot != nil || ie.mux_SR_HARQ_ACK_PUCCH != nil || ie.mux_MultipleGroupCtrlCH_Overlap != nil || ie.dl_SchedulingOffset_PDSCH_TypeA != nil || ie.dl_SchedulingOffset_PDSCH_TypeB != nil || ie.ul_SchedulingOffset != nil || ie.dl_64QAM_MCS_TableAlt != nil || ie.ul_64QAM_MCS_TableAlt != nil || ie.cqi_TableAlt != nil || ie.oneFL_DMRS_TwoAdditionalDMRS_UL != nil || ie.twoFL_DMRS_TwoAdditionalDMRS_UL != nil || ie.oneFL_DMRS_ThreeAdditionalDMRS_UL != nil || ie.pdcch_BlindDetectionNRDC != nil || ie.mux_HARQ_ACK_PUSCH_DiffSymbol != nil || ie.type1_HARQ_ACK_Codebook_r16 != nil || ie.enhancedPowerControl_r16 != nil || ie.simultaneousTCI_ActMultipleCC_r16 != nil || ie.simultaneousSpatialRelationMultipleCC_r16 != nil || ie.cli_RSSI_FDM_DL_r16 != nil || ie.cli_SRS_RSRP_FDM_DL_r16 != nil || ie.maxLayersMIMO_Adaptation_r16 != nil || ie.aggregationFactorSPS_DL_r16 != nil || ie.maxTotalResourcesForOneFreqRange_r16 != nil || ie.csi_ReportFrameworkExt_r16 != nil || ie.twoTCI_Act_servingCellInCC_List_r16 != nil || ie.cri_RI_CQI_WithoutNon_PMI_PortInd_r16 != nil || ie.cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 != nil
	preambleBits := []bool{hasExtensions, ie.dynamicSFI != nil, ie.dummy1 != nil, ie.twoFL_DMRS != nil, ie.dummy2 != nil, ie.dummy3 != nil, ie.supportedDMRS_TypeDL != nil, ie.supportedDMRS_TypeUL != nil, ie.semiOpenLoopCSI != nil, ie.csi_ReportWithoutPMI != nil, ie.csi_ReportWithoutCQI != nil, ie.onePortsPTRS != nil, ie.twoPUCCH_F0_2_ConsecSymbols != nil, ie.pucch_F2_WithFH != nil, ie.pucch_F3_WithFH != nil, ie.pucch_F4_WithFH != nil, ie.pucch_F0_2WithoutFH != nil, ie.pucch_F1_3_4WithoutFH != nil, ie.mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot != nil, ie.uci_CodeBlockSegmentation != nil, ie.onePUCCH_LongAndShortFormat != nil, ie.twoPUCCH_AnyOthersInSlot != nil, ie.intraSlotFreqHopping_PUSCH != nil, ie.pusch_LBRM != nil, ie.pdcch_BlindDetectionCA != nil, ie.tpc_PUSCH_RNTI != nil, ie.tpc_PUCCH_RNTI != nil, ie.tpc_SRS_RNTI != nil, ie.absoluteTPC_Command != nil, ie.twoDifferentTPC_Loop_PUSCH != nil, ie.twoDifferentTPC_Loop_PUCCH != nil, ie.pusch_HalfPi_BPSK != nil, ie.pucch_F3_4_HalfPi_BPSK != nil, ie.almostContiguousCP_OFDM_UL != nil, ie.sp_CSI_RS != nil, ie.sp_CSI_IM != nil, ie.tdd_MultiDL_UL_SwitchPerSlot != nil, ie.multipleCORESET != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dynamicSFI != nil {
		if err = ie.dynamicSFI.Encode(w); err != nil {
			return utils.WrapError("Encode dynamicSFI", err)
		}
	}
	if ie.dummy1 != nil {
		if err = w.WriteBitString(ie.dummy1.Bytes, uint(ie.dummy1.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode dummy1", err)
		}
	}
	if ie.twoFL_DMRS != nil {
		if err = w.WriteBitString(ie.twoFL_DMRS.Bytes, uint(ie.twoFL_DMRS.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode twoFL_DMRS", err)
		}
	}
	if ie.dummy2 != nil {
		if err = w.WriteBitString(ie.dummy2.Bytes, uint(ie.dummy2.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode dummy2", err)
		}
	}
	if ie.dummy3 != nil {
		if err = w.WriteBitString(ie.dummy3.Bytes, uint(ie.dummy3.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode dummy3", err)
		}
	}
	if ie.supportedDMRS_TypeDL != nil {
		if err = ie.supportedDMRS_TypeDL.Encode(w); err != nil {
			return utils.WrapError("Encode supportedDMRS_TypeDL", err)
		}
	}
	if ie.supportedDMRS_TypeUL != nil {
		if err = ie.supportedDMRS_TypeUL.Encode(w); err != nil {
			return utils.WrapError("Encode supportedDMRS_TypeUL", err)
		}
	}
	if ie.semiOpenLoopCSI != nil {
		if err = ie.semiOpenLoopCSI.Encode(w); err != nil {
			return utils.WrapError("Encode semiOpenLoopCSI", err)
		}
	}
	if ie.csi_ReportWithoutPMI != nil {
		if err = ie.csi_ReportWithoutPMI.Encode(w); err != nil {
			return utils.WrapError("Encode csi_ReportWithoutPMI", err)
		}
	}
	if ie.csi_ReportWithoutCQI != nil {
		if err = ie.csi_ReportWithoutCQI.Encode(w); err != nil {
			return utils.WrapError("Encode csi_ReportWithoutCQI", err)
		}
	}
	if ie.onePortsPTRS != nil {
		if err = w.WriteBitString(ie.onePortsPTRS.Bytes, uint(ie.onePortsPTRS.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode onePortsPTRS", err)
		}
	}
	if ie.twoPUCCH_F0_2_ConsecSymbols != nil {
		if err = ie.twoPUCCH_F0_2_ConsecSymbols.Encode(w); err != nil {
			return utils.WrapError("Encode twoPUCCH_F0_2_ConsecSymbols", err)
		}
	}
	if ie.pucch_F2_WithFH != nil {
		if err = ie.pucch_F2_WithFH.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_F2_WithFH", err)
		}
	}
	if ie.pucch_F3_WithFH != nil {
		if err = ie.pucch_F3_WithFH.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_F3_WithFH", err)
		}
	}
	if ie.pucch_F4_WithFH != nil {
		if err = ie.pucch_F4_WithFH.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_F4_WithFH", err)
		}
	}
	if ie.pucch_F0_2WithoutFH != nil {
		if err = ie.pucch_F0_2WithoutFH.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_F0_2WithoutFH", err)
		}
	}
	if ie.pucch_F1_3_4WithoutFH != nil {
		if err = ie.pucch_F1_3_4WithoutFH.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_F1_3_4WithoutFH", err)
		}
	}
	if ie.mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot != nil {
		if err = ie.mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot.Encode(w); err != nil {
			return utils.WrapError("Encode mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot", err)
		}
	}
	if ie.uci_CodeBlockSegmentation != nil {
		if err = ie.uci_CodeBlockSegmentation.Encode(w); err != nil {
			return utils.WrapError("Encode uci_CodeBlockSegmentation", err)
		}
	}
	if ie.onePUCCH_LongAndShortFormat != nil {
		if err = ie.onePUCCH_LongAndShortFormat.Encode(w); err != nil {
			return utils.WrapError("Encode onePUCCH_LongAndShortFormat", err)
		}
	}
	if ie.twoPUCCH_AnyOthersInSlot != nil {
		if err = ie.twoPUCCH_AnyOthersInSlot.Encode(w); err != nil {
			return utils.WrapError("Encode twoPUCCH_AnyOthersInSlot", err)
		}
	}
	if ie.intraSlotFreqHopping_PUSCH != nil {
		if err = ie.intraSlotFreqHopping_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode intraSlotFreqHopping_PUSCH", err)
		}
	}
	if ie.pusch_LBRM != nil {
		if err = ie.pusch_LBRM.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_LBRM", err)
		}
	}
	if ie.pdcch_BlindDetectionCA != nil {
		if err = w.WriteInteger(*ie.pdcch_BlindDetectionCA, &uper.Constraint{Lb: 4, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode pdcch_BlindDetectionCA", err)
		}
	}
	if ie.tpc_PUSCH_RNTI != nil {
		if err = ie.tpc_PUSCH_RNTI.Encode(w); err != nil {
			return utils.WrapError("Encode tpc_PUSCH_RNTI", err)
		}
	}
	if ie.tpc_PUCCH_RNTI != nil {
		if err = ie.tpc_PUCCH_RNTI.Encode(w); err != nil {
			return utils.WrapError("Encode tpc_PUCCH_RNTI", err)
		}
	}
	if ie.tpc_SRS_RNTI != nil {
		if err = ie.tpc_SRS_RNTI.Encode(w); err != nil {
			return utils.WrapError("Encode tpc_SRS_RNTI", err)
		}
	}
	if ie.absoluteTPC_Command != nil {
		if err = ie.absoluteTPC_Command.Encode(w); err != nil {
			return utils.WrapError("Encode absoluteTPC_Command", err)
		}
	}
	if ie.twoDifferentTPC_Loop_PUSCH != nil {
		if err = ie.twoDifferentTPC_Loop_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode twoDifferentTPC_Loop_PUSCH", err)
		}
	}
	if ie.twoDifferentTPC_Loop_PUCCH != nil {
		if err = ie.twoDifferentTPC_Loop_PUCCH.Encode(w); err != nil {
			return utils.WrapError("Encode twoDifferentTPC_Loop_PUCCH", err)
		}
	}
	if ie.pusch_HalfPi_BPSK != nil {
		if err = ie.pusch_HalfPi_BPSK.Encode(w); err != nil {
			return utils.WrapError("Encode pusch_HalfPi_BPSK", err)
		}
	}
	if ie.pucch_F3_4_HalfPi_BPSK != nil {
		if err = ie.pucch_F3_4_HalfPi_BPSK.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_F3_4_HalfPi_BPSK", err)
		}
	}
	if ie.almostContiguousCP_OFDM_UL != nil {
		if err = ie.almostContiguousCP_OFDM_UL.Encode(w); err != nil {
			return utils.WrapError("Encode almostContiguousCP_OFDM_UL", err)
		}
	}
	if ie.sp_CSI_RS != nil {
		if err = ie.sp_CSI_RS.Encode(w); err != nil {
			return utils.WrapError("Encode sp_CSI_RS", err)
		}
	}
	if ie.sp_CSI_IM != nil {
		if err = ie.sp_CSI_IM.Encode(w); err != nil {
			return utils.WrapError("Encode sp_CSI_IM", err)
		}
	}
	if ie.tdd_MultiDL_UL_SwitchPerSlot != nil {
		if err = ie.tdd_MultiDL_UL_SwitchPerSlot.Encode(w); err != nil {
			return utils.WrapError("Encode tdd_MultiDL_UL_SwitchPerSlot", err)
		}
	}
	if ie.multipleCORESET != nil {
		if err = ie.multipleCORESET.Encode(w); err != nil {
			return utils.WrapError("Encode multipleCORESET", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 6 bits for 6 extension groups
		extBitmap := []bool{ie.csi_RS_IM_ReceptionForFeedback != nil || ie.csi_RS_ProcFrameworkForSRS != nil || ie.csi_ReportFramework != nil || ie.mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot != nil || ie.mux_SR_HARQ_ACK_PUCCH != nil || ie.mux_MultipleGroupCtrlCH_Overlap != nil || ie.dl_SchedulingOffset_PDSCH_TypeA != nil || ie.dl_SchedulingOffset_PDSCH_TypeB != nil || ie.ul_SchedulingOffset != nil || ie.dl_64QAM_MCS_TableAlt != nil || ie.ul_64QAM_MCS_TableAlt != nil || ie.cqi_TableAlt != nil || ie.oneFL_DMRS_TwoAdditionalDMRS_UL != nil || ie.twoFL_DMRS_TwoAdditionalDMRS_UL != nil || ie.oneFL_DMRS_ThreeAdditionalDMRS_UL != nil, ie.pdcch_BlindDetectionNRDC != nil || ie.mux_HARQ_ACK_PUSCH_DiffSymbol != nil, ie.type1_HARQ_ACK_Codebook_r16 != nil || ie.enhancedPowerControl_r16 != nil || ie.simultaneousTCI_ActMultipleCC_r16 != nil || ie.simultaneousSpatialRelationMultipleCC_r16 != nil || ie.cli_RSSI_FDM_DL_r16 != nil || ie.cli_SRS_RSRP_FDM_DL_r16 != nil || ie.maxLayersMIMO_Adaptation_r16 != nil || ie.aggregationFactorSPS_DL_r16 != nil || ie.maxTotalResourcesForOneFreqRange_r16 != nil || ie.csi_ReportFrameworkExt_r16 != nil, ie.twoTCI_Act_servingCellInCC_List_r16 != nil, ie.cri_RI_CQI_WithoutNon_PMI_PortInd_r16 != nil, ie.cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap Phy_ParametersFRX_Diff", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.csi_RS_IM_ReceptionForFeedback != nil, ie.csi_RS_ProcFrameworkForSRS != nil, ie.csi_ReportFramework != nil, ie.mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot != nil, ie.mux_SR_HARQ_ACK_PUCCH != nil, ie.mux_MultipleGroupCtrlCH_Overlap != nil, ie.dl_SchedulingOffset_PDSCH_TypeA != nil, ie.dl_SchedulingOffset_PDSCH_TypeB != nil, ie.ul_SchedulingOffset != nil, ie.dl_64QAM_MCS_TableAlt != nil, ie.ul_64QAM_MCS_TableAlt != nil, ie.cqi_TableAlt != nil, ie.oneFL_DMRS_TwoAdditionalDMRS_UL != nil, ie.twoFL_DMRS_TwoAdditionalDMRS_UL != nil, ie.oneFL_DMRS_ThreeAdditionalDMRS_UL != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
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
			// encode mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot optional
			if ie.mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot != nil {
				if err = ie.mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot", err)
				}
			}
			// encode mux_SR_HARQ_ACK_PUCCH optional
			if ie.mux_SR_HARQ_ACK_PUCCH != nil {
				if err = ie.mux_SR_HARQ_ACK_PUCCH.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mux_SR_HARQ_ACK_PUCCH", err)
				}
			}
			// encode mux_MultipleGroupCtrlCH_Overlap optional
			if ie.mux_MultipleGroupCtrlCH_Overlap != nil {
				if err = ie.mux_MultipleGroupCtrlCH_Overlap.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mux_MultipleGroupCtrlCH_Overlap", err)
				}
			}
			// encode dl_SchedulingOffset_PDSCH_TypeA optional
			if ie.dl_SchedulingOffset_PDSCH_TypeA != nil {
				if err = ie.dl_SchedulingOffset_PDSCH_TypeA.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dl_SchedulingOffset_PDSCH_TypeA", err)
				}
			}
			// encode dl_SchedulingOffset_PDSCH_TypeB optional
			if ie.dl_SchedulingOffset_PDSCH_TypeB != nil {
				if err = ie.dl_SchedulingOffset_PDSCH_TypeB.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dl_SchedulingOffset_PDSCH_TypeB", err)
				}
			}
			// encode ul_SchedulingOffset optional
			if ie.ul_SchedulingOffset != nil {
				if err = ie.ul_SchedulingOffset.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_SchedulingOffset", err)
				}
			}
			// encode dl_64QAM_MCS_TableAlt optional
			if ie.dl_64QAM_MCS_TableAlt != nil {
				if err = ie.dl_64QAM_MCS_TableAlt.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dl_64QAM_MCS_TableAlt", err)
				}
			}
			// encode ul_64QAM_MCS_TableAlt optional
			if ie.ul_64QAM_MCS_TableAlt != nil {
				if err = ie.ul_64QAM_MCS_TableAlt.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_64QAM_MCS_TableAlt", err)
				}
			}
			// encode cqi_TableAlt optional
			if ie.cqi_TableAlt != nil {
				if err = ie.cqi_TableAlt.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cqi_TableAlt", err)
				}
			}
			// encode oneFL_DMRS_TwoAdditionalDMRS_UL optional
			if ie.oneFL_DMRS_TwoAdditionalDMRS_UL != nil {
				if err = ie.oneFL_DMRS_TwoAdditionalDMRS_UL.Encode(extWriter); err != nil {
					return utils.WrapError("Encode oneFL_DMRS_TwoAdditionalDMRS_UL", err)
				}
			}
			// encode twoFL_DMRS_TwoAdditionalDMRS_UL optional
			if ie.twoFL_DMRS_TwoAdditionalDMRS_UL != nil {
				if err = ie.twoFL_DMRS_TwoAdditionalDMRS_UL.Encode(extWriter); err != nil {
					return utils.WrapError("Encode twoFL_DMRS_TwoAdditionalDMRS_UL", err)
				}
			}
			// encode oneFL_DMRS_ThreeAdditionalDMRS_UL optional
			if ie.oneFL_DMRS_ThreeAdditionalDMRS_UL != nil {
				if err = ie.oneFL_DMRS_ThreeAdditionalDMRS_UL.Encode(extWriter); err != nil {
					return utils.WrapError("Encode oneFL_DMRS_ThreeAdditionalDMRS_UL", err)
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
			optionals_ext_2 := []bool{ie.pdcch_BlindDetectionNRDC != nil, ie.mux_HARQ_ACK_PUSCH_DiffSymbol != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode pdcch_BlindDetectionNRDC optional
			if ie.pdcch_BlindDetectionNRDC != nil {
				if err = ie.pdcch_BlindDetectionNRDC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdcch_BlindDetectionNRDC", err)
				}
			}
			// encode mux_HARQ_ACK_PUSCH_DiffSymbol optional
			if ie.mux_HARQ_ACK_PUSCH_DiffSymbol != nil {
				if err = ie.mux_HARQ_ACK_PUSCH_DiffSymbol.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mux_HARQ_ACK_PUSCH_DiffSymbol", err)
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
			optionals_ext_3 := []bool{ie.type1_HARQ_ACK_Codebook_r16 != nil, ie.enhancedPowerControl_r16 != nil, ie.simultaneousTCI_ActMultipleCC_r16 != nil, ie.simultaneousSpatialRelationMultipleCC_r16 != nil, ie.cli_RSSI_FDM_DL_r16 != nil, ie.cli_SRS_RSRP_FDM_DL_r16 != nil, ie.maxLayersMIMO_Adaptation_r16 != nil, ie.aggregationFactorSPS_DL_r16 != nil, ie.maxTotalResourcesForOneFreqRange_r16 != nil, ie.csi_ReportFrameworkExt_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode type1_HARQ_ACK_Codebook_r16 optional
			if ie.type1_HARQ_ACK_Codebook_r16 != nil {
				if err = ie.type1_HARQ_ACK_Codebook_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode type1_HARQ_ACK_Codebook_r16", err)
				}
			}
			// encode enhancedPowerControl_r16 optional
			if ie.enhancedPowerControl_r16 != nil {
				if err = ie.enhancedPowerControl_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enhancedPowerControl_r16", err)
				}
			}
			// encode simultaneousTCI_ActMultipleCC_r16 optional
			if ie.simultaneousTCI_ActMultipleCC_r16 != nil {
				if err = ie.simultaneousTCI_ActMultipleCC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simultaneousTCI_ActMultipleCC_r16", err)
				}
			}
			// encode simultaneousSpatialRelationMultipleCC_r16 optional
			if ie.simultaneousSpatialRelationMultipleCC_r16 != nil {
				if err = ie.simultaneousSpatialRelationMultipleCC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simultaneousSpatialRelationMultipleCC_r16", err)
				}
			}
			// encode cli_RSSI_FDM_DL_r16 optional
			if ie.cli_RSSI_FDM_DL_r16 != nil {
				if err = ie.cli_RSSI_FDM_DL_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cli_RSSI_FDM_DL_r16", err)
				}
			}
			// encode cli_SRS_RSRP_FDM_DL_r16 optional
			if ie.cli_SRS_RSRP_FDM_DL_r16 != nil {
				if err = ie.cli_SRS_RSRP_FDM_DL_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cli_SRS_RSRP_FDM_DL_r16", err)
				}
			}
			// encode maxLayersMIMO_Adaptation_r16 optional
			if ie.maxLayersMIMO_Adaptation_r16 != nil {
				if err = ie.maxLayersMIMO_Adaptation_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxLayersMIMO_Adaptation_r16", err)
				}
			}
			// encode aggregationFactorSPS_DL_r16 optional
			if ie.aggregationFactorSPS_DL_r16 != nil {
				if err = ie.aggregationFactorSPS_DL_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode aggregationFactorSPS_DL_r16", err)
				}
			}
			// encode maxTotalResourcesForOneFreqRange_r16 optional
			if ie.maxTotalResourcesForOneFreqRange_r16 != nil {
				if err = ie.maxTotalResourcesForOneFreqRange_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxTotalResourcesForOneFreqRange_r16", err)
				}
			}
			// encode csi_ReportFrameworkExt_r16 optional
			if ie.csi_ReportFrameworkExt_r16 != nil {
				if err = ie.csi_ReportFrameworkExt_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode csi_ReportFrameworkExt_r16", err)
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
			optionals_ext_4 := []bool{ie.twoTCI_Act_servingCellInCC_List_r16 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode twoTCI_Act_servingCellInCC_List_r16 optional
			if ie.twoTCI_Act_servingCellInCC_List_r16 != nil {
				if err = ie.twoTCI_Act_servingCellInCC_List_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode twoTCI_Act_servingCellInCC_List_r16", err)
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
			optionals_ext_5 := []bool{ie.cri_RI_CQI_WithoutNon_PMI_PortInd_r16 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode cri_RI_CQI_WithoutNon_PMI_PortInd_r16 optional
			if ie.cri_RI_CQI_WithoutNon_PMI_PortInd_r16 != nil {
				if err = ie.cri_RI_CQI_WithoutNon_PMI_PortInd_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cri_RI_CQI_WithoutNon_PMI_PortInd_r16", err)
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
			optionals_ext_6 := []bool{ie.cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 != nil}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 optional
			if ie.cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 != nil {
				if err = ie.cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17", err)
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

func (ie *Phy_ParametersFRX_Diff) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var dynamicSFIPresent bool
	if dynamicSFIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy1Present bool
	if dummy1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var twoFL_DMRSPresent bool
	if twoFL_DMRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy2Present bool
	if dummy2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dummy3Present bool
	if dummy3Present, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedDMRS_TypeDLPresent bool
	if supportedDMRS_TypeDLPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedDMRS_TypeULPresent bool
	if supportedDMRS_TypeULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var semiOpenLoopCSIPresent bool
	if semiOpenLoopCSIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_ReportWithoutPMIPresent bool
	if csi_ReportWithoutPMIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_ReportWithoutCQIPresent bool
	if csi_ReportWithoutCQIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var onePortsPTRSPresent bool
	if onePortsPTRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPUCCH_F0_2_ConsecSymbolsPresent bool
	if twoPUCCH_F0_2_ConsecSymbolsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_F2_WithFHPresent bool
	if pucch_F2_WithFHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_F3_WithFHPresent bool
	if pucch_F3_WithFHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_F4_WithFHPresent bool
	if pucch_F4_WithFHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_F0_2WithoutFHPresent bool
	if pucch_F0_2WithoutFHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_F1_3_4WithoutFHPresent bool
	if pucch_F1_3_4WithoutFHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlotPresent bool
	if mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlotPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var uci_CodeBlockSegmentationPresent bool
	if uci_CodeBlockSegmentationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var onePUCCH_LongAndShortFormatPresent bool
	if onePUCCH_LongAndShortFormatPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var twoPUCCH_AnyOthersInSlotPresent bool
	if twoPUCCH_AnyOthersInSlotPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var intraSlotFreqHopping_PUSCHPresent bool
	if intraSlotFreqHopping_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_LBRMPresent bool
	if pusch_LBRMPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcch_BlindDetectionCAPresent bool
	if pdcch_BlindDetectionCAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tpc_PUSCH_RNTIPresent bool
	if tpc_PUSCH_RNTIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tpc_PUCCH_RNTIPresent bool
	if tpc_PUCCH_RNTIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tpc_SRS_RNTIPresent bool
	if tpc_SRS_RNTIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var absoluteTPC_CommandPresent bool
	if absoluteTPC_CommandPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var twoDifferentTPC_Loop_PUSCHPresent bool
	if twoDifferentTPC_Loop_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var twoDifferentTPC_Loop_PUCCHPresent bool
	if twoDifferentTPC_Loop_PUCCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pusch_HalfPi_BPSKPresent bool
	if pusch_HalfPi_BPSKPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_F3_4_HalfPi_BPSKPresent bool
	if pucch_F3_4_HalfPi_BPSKPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var almostContiguousCP_OFDM_ULPresent bool
	if almostContiguousCP_OFDM_ULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sp_CSI_RSPresent bool
	if sp_CSI_RSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sp_CSI_IMPresent bool
	if sp_CSI_IMPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tdd_MultiDL_UL_SwitchPerSlotPresent bool
	if tdd_MultiDL_UL_SwitchPerSlotPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var multipleCORESETPresent bool
	if multipleCORESETPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if dynamicSFIPresent {
		ie.dynamicSFI = new(Phy_ParametersFRX_Diff_dynamicSFI)
		if err = ie.dynamicSFI.Decode(r); err != nil {
			return utils.WrapError("Decode dynamicSFI", err)
		}
	}
	if dummy1Present {
		var tmp_bs_dummy1 []byte
		var n_dummy1 uint
		if tmp_bs_dummy1, n_dummy1, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode dummy1", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_dummy1,
			NumBits: uint64(n_dummy1),
		}
		ie.dummy1 = &tmp_bitstring
	}
	if twoFL_DMRSPresent {
		var tmp_bs_twoFL_DMRS []byte
		var n_twoFL_DMRS uint
		if tmp_bs_twoFL_DMRS, n_twoFL_DMRS, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode twoFL_DMRS", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_twoFL_DMRS,
			NumBits: uint64(n_twoFL_DMRS),
		}
		ie.twoFL_DMRS = &tmp_bitstring
	}
	if dummy2Present {
		var tmp_bs_dummy2 []byte
		var n_dummy2 uint
		if tmp_bs_dummy2, n_dummy2, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode dummy2", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_dummy2,
			NumBits: uint64(n_dummy2),
		}
		ie.dummy2 = &tmp_bitstring
	}
	if dummy3Present {
		var tmp_bs_dummy3 []byte
		var n_dummy3 uint
		if tmp_bs_dummy3, n_dummy3, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode dummy3", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_dummy3,
			NumBits: uint64(n_dummy3),
		}
		ie.dummy3 = &tmp_bitstring
	}
	if supportedDMRS_TypeDLPresent {
		ie.supportedDMRS_TypeDL = new(Phy_ParametersFRX_Diff_supportedDMRS_TypeDL)
		if err = ie.supportedDMRS_TypeDL.Decode(r); err != nil {
			return utils.WrapError("Decode supportedDMRS_TypeDL", err)
		}
	}
	if supportedDMRS_TypeULPresent {
		ie.supportedDMRS_TypeUL = new(Phy_ParametersFRX_Diff_supportedDMRS_TypeUL)
		if err = ie.supportedDMRS_TypeUL.Decode(r); err != nil {
			return utils.WrapError("Decode supportedDMRS_TypeUL", err)
		}
	}
	if semiOpenLoopCSIPresent {
		ie.semiOpenLoopCSI = new(Phy_ParametersFRX_Diff_semiOpenLoopCSI)
		if err = ie.semiOpenLoopCSI.Decode(r); err != nil {
			return utils.WrapError("Decode semiOpenLoopCSI", err)
		}
	}
	if csi_ReportWithoutPMIPresent {
		ie.csi_ReportWithoutPMI = new(Phy_ParametersFRX_Diff_csi_ReportWithoutPMI)
		if err = ie.csi_ReportWithoutPMI.Decode(r); err != nil {
			return utils.WrapError("Decode csi_ReportWithoutPMI", err)
		}
	}
	if csi_ReportWithoutCQIPresent {
		ie.csi_ReportWithoutCQI = new(Phy_ParametersFRX_Diff_csi_ReportWithoutCQI)
		if err = ie.csi_ReportWithoutCQI.Decode(r); err != nil {
			return utils.WrapError("Decode csi_ReportWithoutCQI", err)
		}
	}
	if onePortsPTRSPresent {
		var tmp_bs_onePortsPTRS []byte
		var n_onePortsPTRS uint
		if tmp_bs_onePortsPTRS, n_onePortsPTRS, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode onePortsPTRS", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_onePortsPTRS,
			NumBits: uint64(n_onePortsPTRS),
		}
		ie.onePortsPTRS = &tmp_bitstring
	}
	if twoPUCCH_F0_2_ConsecSymbolsPresent {
		ie.twoPUCCH_F0_2_ConsecSymbols = new(Phy_ParametersFRX_Diff_twoPUCCH_F0_2_ConsecSymbols)
		if err = ie.twoPUCCH_F0_2_ConsecSymbols.Decode(r); err != nil {
			return utils.WrapError("Decode twoPUCCH_F0_2_ConsecSymbols", err)
		}
	}
	if pucch_F2_WithFHPresent {
		ie.pucch_F2_WithFH = new(Phy_ParametersFRX_Diff_pucch_F2_WithFH)
		if err = ie.pucch_F2_WithFH.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_F2_WithFH", err)
		}
	}
	if pucch_F3_WithFHPresent {
		ie.pucch_F3_WithFH = new(Phy_ParametersFRX_Diff_pucch_F3_WithFH)
		if err = ie.pucch_F3_WithFH.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_F3_WithFH", err)
		}
	}
	if pucch_F4_WithFHPresent {
		ie.pucch_F4_WithFH = new(Phy_ParametersFRX_Diff_pucch_F4_WithFH)
		if err = ie.pucch_F4_WithFH.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_F4_WithFH", err)
		}
	}
	if pucch_F0_2WithoutFHPresent {
		ie.pucch_F0_2WithoutFH = new(Phy_ParametersFRX_Diff_pucch_F0_2WithoutFH)
		if err = ie.pucch_F0_2WithoutFH.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_F0_2WithoutFH", err)
		}
	}
	if pucch_F1_3_4WithoutFHPresent {
		ie.pucch_F1_3_4WithoutFH = new(Phy_ParametersFRX_Diff_pucch_F1_3_4WithoutFH)
		if err = ie.pucch_F1_3_4WithoutFH.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_F1_3_4WithoutFH", err)
		}
	}
	if mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlotPresent {
		ie.mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot = new(Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot)
		if err = ie.mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot.Decode(r); err != nil {
			return utils.WrapError("Decode mux_SR_HARQ_ACK_CSI_PUCCH_MultiPerSlot", err)
		}
	}
	if uci_CodeBlockSegmentationPresent {
		ie.uci_CodeBlockSegmentation = new(Phy_ParametersFRX_Diff_uci_CodeBlockSegmentation)
		if err = ie.uci_CodeBlockSegmentation.Decode(r); err != nil {
			return utils.WrapError("Decode uci_CodeBlockSegmentation", err)
		}
	}
	if onePUCCH_LongAndShortFormatPresent {
		ie.onePUCCH_LongAndShortFormat = new(Phy_ParametersFRX_Diff_onePUCCH_LongAndShortFormat)
		if err = ie.onePUCCH_LongAndShortFormat.Decode(r); err != nil {
			return utils.WrapError("Decode onePUCCH_LongAndShortFormat", err)
		}
	}
	if twoPUCCH_AnyOthersInSlotPresent {
		ie.twoPUCCH_AnyOthersInSlot = new(Phy_ParametersFRX_Diff_twoPUCCH_AnyOthersInSlot)
		if err = ie.twoPUCCH_AnyOthersInSlot.Decode(r); err != nil {
			return utils.WrapError("Decode twoPUCCH_AnyOthersInSlot", err)
		}
	}
	if intraSlotFreqHopping_PUSCHPresent {
		ie.intraSlotFreqHopping_PUSCH = new(Phy_ParametersFRX_Diff_intraSlotFreqHopping_PUSCH)
		if err = ie.intraSlotFreqHopping_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode intraSlotFreqHopping_PUSCH", err)
		}
	}
	if pusch_LBRMPresent {
		ie.pusch_LBRM = new(Phy_ParametersFRX_Diff_pusch_LBRM)
		if err = ie.pusch_LBRM.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_LBRM", err)
		}
	}
	if pdcch_BlindDetectionCAPresent {
		var tmp_int_pdcch_BlindDetectionCA int64
		if tmp_int_pdcch_BlindDetectionCA, err = r.ReadInteger(&uper.Constraint{Lb: 4, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode pdcch_BlindDetectionCA", err)
		}
		ie.pdcch_BlindDetectionCA = &tmp_int_pdcch_BlindDetectionCA
	}
	if tpc_PUSCH_RNTIPresent {
		ie.tpc_PUSCH_RNTI = new(Phy_ParametersFRX_Diff_tpc_PUSCH_RNTI)
		if err = ie.tpc_PUSCH_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode tpc_PUSCH_RNTI", err)
		}
	}
	if tpc_PUCCH_RNTIPresent {
		ie.tpc_PUCCH_RNTI = new(Phy_ParametersFRX_Diff_tpc_PUCCH_RNTI)
		if err = ie.tpc_PUCCH_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode tpc_PUCCH_RNTI", err)
		}
	}
	if tpc_SRS_RNTIPresent {
		ie.tpc_SRS_RNTI = new(Phy_ParametersFRX_Diff_tpc_SRS_RNTI)
		if err = ie.tpc_SRS_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode tpc_SRS_RNTI", err)
		}
	}
	if absoluteTPC_CommandPresent {
		ie.absoluteTPC_Command = new(Phy_ParametersFRX_Diff_absoluteTPC_Command)
		if err = ie.absoluteTPC_Command.Decode(r); err != nil {
			return utils.WrapError("Decode absoluteTPC_Command", err)
		}
	}
	if twoDifferentTPC_Loop_PUSCHPresent {
		ie.twoDifferentTPC_Loop_PUSCH = new(Phy_ParametersFRX_Diff_twoDifferentTPC_Loop_PUSCH)
		if err = ie.twoDifferentTPC_Loop_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode twoDifferentTPC_Loop_PUSCH", err)
		}
	}
	if twoDifferentTPC_Loop_PUCCHPresent {
		ie.twoDifferentTPC_Loop_PUCCH = new(Phy_ParametersFRX_Diff_twoDifferentTPC_Loop_PUCCH)
		if err = ie.twoDifferentTPC_Loop_PUCCH.Decode(r); err != nil {
			return utils.WrapError("Decode twoDifferentTPC_Loop_PUCCH", err)
		}
	}
	if pusch_HalfPi_BPSKPresent {
		ie.pusch_HalfPi_BPSK = new(Phy_ParametersFRX_Diff_pusch_HalfPi_BPSK)
		if err = ie.pusch_HalfPi_BPSK.Decode(r); err != nil {
			return utils.WrapError("Decode pusch_HalfPi_BPSK", err)
		}
	}
	if pucch_F3_4_HalfPi_BPSKPresent {
		ie.pucch_F3_4_HalfPi_BPSK = new(Phy_ParametersFRX_Diff_pucch_F3_4_HalfPi_BPSK)
		if err = ie.pucch_F3_4_HalfPi_BPSK.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_F3_4_HalfPi_BPSK", err)
		}
	}
	if almostContiguousCP_OFDM_ULPresent {
		ie.almostContiguousCP_OFDM_UL = new(Phy_ParametersFRX_Diff_almostContiguousCP_OFDM_UL)
		if err = ie.almostContiguousCP_OFDM_UL.Decode(r); err != nil {
			return utils.WrapError("Decode almostContiguousCP_OFDM_UL", err)
		}
	}
	if sp_CSI_RSPresent {
		ie.sp_CSI_RS = new(Phy_ParametersFRX_Diff_sp_CSI_RS)
		if err = ie.sp_CSI_RS.Decode(r); err != nil {
			return utils.WrapError("Decode sp_CSI_RS", err)
		}
	}
	if sp_CSI_IMPresent {
		ie.sp_CSI_IM = new(Phy_ParametersFRX_Diff_sp_CSI_IM)
		if err = ie.sp_CSI_IM.Decode(r); err != nil {
			return utils.WrapError("Decode sp_CSI_IM", err)
		}
	}
	if tdd_MultiDL_UL_SwitchPerSlotPresent {
		ie.tdd_MultiDL_UL_SwitchPerSlot = new(Phy_ParametersFRX_Diff_tdd_MultiDL_UL_SwitchPerSlot)
		if err = ie.tdd_MultiDL_UL_SwitchPerSlot.Decode(r); err != nil {
			return utils.WrapError("Decode tdd_MultiDL_UL_SwitchPerSlot", err)
		}
	}
	if multipleCORESETPresent {
		ie.multipleCORESET = new(Phy_ParametersFRX_Diff_multipleCORESET)
		if err = ie.multipleCORESET.Decode(r); err != nil {
			return utils.WrapError("Decode multipleCORESET", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 6 bits for 6 extension groups
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
			mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlotPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mux_SR_HARQ_ACK_PUCCHPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mux_MultipleGroupCtrlCH_OverlapPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dl_SchedulingOffset_PDSCH_TypeAPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dl_SchedulingOffset_PDSCH_TypeBPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_SchedulingOffsetPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dl_64QAM_MCS_TableAltPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_64QAM_MCS_TableAltPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cqi_TableAltPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			oneFL_DMRS_TwoAdditionalDMRS_ULPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			twoFL_DMRS_TwoAdditionalDMRS_ULPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			oneFL_DMRS_ThreeAdditionalDMRS_ULPresent, err := extReader.ReadBool()
			if err != nil {
				return err
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
			// decode mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot optional
			if mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlotPresent {
				ie.mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot = new(Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot)
				if err = ie.mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot.Decode(extReader); err != nil {
					return utils.WrapError("Decode mux_SR_HARQ_ACK_CSI_PUCCH_OncePerSlot", err)
				}
			}
			// decode mux_SR_HARQ_ACK_PUCCH optional
			if mux_SR_HARQ_ACK_PUCCHPresent {
				ie.mux_SR_HARQ_ACK_PUCCH = new(Phy_ParametersFRX_Diff_mux_SR_HARQ_ACK_PUCCH)
				if err = ie.mux_SR_HARQ_ACK_PUCCH.Decode(extReader); err != nil {
					return utils.WrapError("Decode mux_SR_HARQ_ACK_PUCCH", err)
				}
			}
			// decode mux_MultipleGroupCtrlCH_Overlap optional
			if mux_MultipleGroupCtrlCH_OverlapPresent {
				ie.mux_MultipleGroupCtrlCH_Overlap = new(Phy_ParametersFRX_Diff_mux_MultipleGroupCtrlCH_Overlap)
				if err = ie.mux_MultipleGroupCtrlCH_Overlap.Decode(extReader); err != nil {
					return utils.WrapError("Decode mux_MultipleGroupCtrlCH_Overlap", err)
				}
			}
			// decode dl_SchedulingOffset_PDSCH_TypeA optional
			if dl_SchedulingOffset_PDSCH_TypeAPresent {
				ie.dl_SchedulingOffset_PDSCH_TypeA = new(Phy_ParametersFRX_Diff_dl_SchedulingOffset_PDSCH_TypeA)
				if err = ie.dl_SchedulingOffset_PDSCH_TypeA.Decode(extReader); err != nil {
					return utils.WrapError("Decode dl_SchedulingOffset_PDSCH_TypeA", err)
				}
			}
			// decode dl_SchedulingOffset_PDSCH_TypeB optional
			if dl_SchedulingOffset_PDSCH_TypeBPresent {
				ie.dl_SchedulingOffset_PDSCH_TypeB = new(Phy_ParametersFRX_Diff_dl_SchedulingOffset_PDSCH_TypeB)
				if err = ie.dl_SchedulingOffset_PDSCH_TypeB.Decode(extReader); err != nil {
					return utils.WrapError("Decode dl_SchedulingOffset_PDSCH_TypeB", err)
				}
			}
			// decode ul_SchedulingOffset optional
			if ul_SchedulingOffsetPresent {
				ie.ul_SchedulingOffset = new(Phy_ParametersFRX_Diff_ul_SchedulingOffset)
				if err = ie.ul_SchedulingOffset.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_SchedulingOffset", err)
				}
			}
			// decode dl_64QAM_MCS_TableAlt optional
			if dl_64QAM_MCS_TableAltPresent {
				ie.dl_64QAM_MCS_TableAlt = new(Phy_ParametersFRX_Diff_dl_64QAM_MCS_TableAlt)
				if err = ie.dl_64QAM_MCS_TableAlt.Decode(extReader); err != nil {
					return utils.WrapError("Decode dl_64QAM_MCS_TableAlt", err)
				}
			}
			// decode ul_64QAM_MCS_TableAlt optional
			if ul_64QAM_MCS_TableAltPresent {
				ie.ul_64QAM_MCS_TableAlt = new(Phy_ParametersFRX_Diff_ul_64QAM_MCS_TableAlt)
				if err = ie.ul_64QAM_MCS_TableAlt.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_64QAM_MCS_TableAlt", err)
				}
			}
			// decode cqi_TableAlt optional
			if cqi_TableAltPresent {
				ie.cqi_TableAlt = new(Phy_ParametersFRX_Diff_cqi_TableAlt)
				if err = ie.cqi_TableAlt.Decode(extReader); err != nil {
					return utils.WrapError("Decode cqi_TableAlt", err)
				}
			}
			// decode oneFL_DMRS_TwoAdditionalDMRS_UL optional
			if oneFL_DMRS_TwoAdditionalDMRS_ULPresent {
				ie.oneFL_DMRS_TwoAdditionalDMRS_UL = new(Phy_ParametersFRX_Diff_oneFL_DMRS_TwoAdditionalDMRS_UL)
				if err = ie.oneFL_DMRS_TwoAdditionalDMRS_UL.Decode(extReader); err != nil {
					return utils.WrapError("Decode oneFL_DMRS_TwoAdditionalDMRS_UL", err)
				}
			}
			// decode twoFL_DMRS_TwoAdditionalDMRS_UL optional
			if twoFL_DMRS_TwoAdditionalDMRS_ULPresent {
				ie.twoFL_DMRS_TwoAdditionalDMRS_UL = new(Phy_ParametersFRX_Diff_twoFL_DMRS_TwoAdditionalDMRS_UL)
				if err = ie.twoFL_DMRS_TwoAdditionalDMRS_UL.Decode(extReader); err != nil {
					return utils.WrapError("Decode twoFL_DMRS_TwoAdditionalDMRS_UL", err)
				}
			}
			// decode oneFL_DMRS_ThreeAdditionalDMRS_UL optional
			if oneFL_DMRS_ThreeAdditionalDMRS_ULPresent {
				ie.oneFL_DMRS_ThreeAdditionalDMRS_UL = new(Phy_ParametersFRX_Diff_oneFL_DMRS_ThreeAdditionalDMRS_UL)
				if err = ie.oneFL_DMRS_ThreeAdditionalDMRS_UL.Decode(extReader); err != nil {
					return utils.WrapError("Decode oneFL_DMRS_ThreeAdditionalDMRS_UL", err)
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

			pdcch_BlindDetectionNRDCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mux_HARQ_ACK_PUSCH_DiffSymbolPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode pdcch_BlindDetectionNRDC optional
			if pdcch_BlindDetectionNRDCPresent {
				ie.pdcch_BlindDetectionNRDC = new(Phy_ParametersFRX_Diff_pdcch_BlindDetectionNRDC)
				if err = ie.pdcch_BlindDetectionNRDC.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdcch_BlindDetectionNRDC", err)
				}
			}
			// decode mux_HARQ_ACK_PUSCH_DiffSymbol optional
			if mux_HARQ_ACK_PUSCH_DiffSymbolPresent {
				ie.mux_HARQ_ACK_PUSCH_DiffSymbol = new(Phy_ParametersFRX_Diff_mux_HARQ_ACK_PUSCH_DiffSymbol)
				if err = ie.mux_HARQ_ACK_PUSCH_DiffSymbol.Decode(extReader); err != nil {
					return utils.WrapError("Decode mux_HARQ_ACK_PUSCH_DiffSymbol", err)
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

			type1_HARQ_ACK_Codebook_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			enhancedPowerControl_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			simultaneousTCI_ActMultipleCC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			simultaneousSpatialRelationMultipleCC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cli_RSSI_FDM_DL_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cli_SRS_RSRP_FDM_DL_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxLayersMIMO_Adaptation_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			aggregationFactorSPS_DL_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxTotalResourcesForOneFreqRange_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			csi_ReportFrameworkExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode type1_HARQ_ACK_Codebook_r16 optional
			if type1_HARQ_ACK_Codebook_r16Present {
				ie.type1_HARQ_ACK_Codebook_r16 = new(Phy_ParametersFRX_Diff_type1_HARQ_ACK_Codebook_r16)
				if err = ie.type1_HARQ_ACK_Codebook_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode type1_HARQ_ACK_Codebook_r16", err)
				}
			}
			// decode enhancedPowerControl_r16 optional
			if enhancedPowerControl_r16Present {
				ie.enhancedPowerControl_r16 = new(Phy_ParametersFRX_Diff_enhancedPowerControl_r16)
				if err = ie.enhancedPowerControl_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode enhancedPowerControl_r16", err)
				}
			}
			// decode simultaneousTCI_ActMultipleCC_r16 optional
			if simultaneousTCI_ActMultipleCC_r16Present {
				ie.simultaneousTCI_ActMultipleCC_r16 = new(Phy_ParametersFRX_Diff_simultaneousTCI_ActMultipleCC_r16)
				if err = ie.simultaneousTCI_ActMultipleCC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode simultaneousTCI_ActMultipleCC_r16", err)
				}
			}
			// decode simultaneousSpatialRelationMultipleCC_r16 optional
			if simultaneousSpatialRelationMultipleCC_r16Present {
				ie.simultaneousSpatialRelationMultipleCC_r16 = new(Phy_ParametersFRX_Diff_simultaneousSpatialRelationMultipleCC_r16)
				if err = ie.simultaneousSpatialRelationMultipleCC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode simultaneousSpatialRelationMultipleCC_r16", err)
				}
			}
			// decode cli_RSSI_FDM_DL_r16 optional
			if cli_RSSI_FDM_DL_r16Present {
				ie.cli_RSSI_FDM_DL_r16 = new(Phy_ParametersFRX_Diff_cli_RSSI_FDM_DL_r16)
				if err = ie.cli_RSSI_FDM_DL_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode cli_RSSI_FDM_DL_r16", err)
				}
			}
			// decode cli_SRS_RSRP_FDM_DL_r16 optional
			if cli_SRS_RSRP_FDM_DL_r16Present {
				ie.cli_SRS_RSRP_FDM_DL_r16 = new(Phy_ParametersFRX_Diff_cli_SRS_RSRP_FDM_DL_r16)
				if err = ie.cli_SRS_RSRP_FDM_DL_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode cli_SRS_RSRP_FDM_DL_r16", err)
				}
			}
			// decode maxLayersMIMO_Adaptation_r16 optional
			if maxLayersMIMO_Adaptation_r16Present {
				ie.maxLayersMIMO_Adaptation_r16 = new(Phy_ParametersFRX_Diff_maxLayersMIMO_Adaptation_r16)
				if err = ie.maxLayersMIMO_Adaptation_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxLayersMIMO_Adaptation_r16", err)
				}
			}
			// decode aggregationFactorSPS_DL_r16 optional
			if aggregationFactorSPS_DL_r16Present {
				ie.aggregationFactorSPS_DL_r16 = new(Phy_ParametersFRX_Diff_aggregationFactorSPS_DL_r16)
				if err = ie.aggregationFactorSPS_DL_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode aggregationFactorSPS_DL_r16", err)
				}
			}
			// decode maxTotalResourcesForOneFreqRange_r16 optional
			if maxTotalResourcesForOneFreqRange_r16Present {
				ie.maxTotalResourcesForOneFreqRange_r16 = new(Phy_ParametersFRX_Diff_maxTotalResourcesForOneFreqRange_r16)
				if err = ie.maxTotalResourcesForOneFreqRange_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxTotalResourcesForOneFreqRange_r16", err)
				}
			}
			// decode csi_ReportFrameworkExt_r16 optional
			if csi_ReportFrameworkExt_r16Present {
				ie.csi_ReportFrameworkExt_r16 = new(CSI_ReportFrameworkExt_r16)
				if err = ie.csi_ReportFrameworkExt_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode csi_ReportFrameworkExt_r16", err)
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

			twoTCI_Act_servingCellInCC_List_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode twoTCI_Act_servingCellInCC_List_r16 optional
			if twoTCI_Act_servingCellInCC_List_r16Present {
				ie.twoTCI_Act_servingCellInCC_List_r16 = new(Phy_ParametersFRX_Diff_twoTCI_Act_servingCellInCC_List_r16)
				if err = ie.twoTCI_Act_servingCellInCC_List_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode twoTCI_Act_servingCellInCC_List_r16", err)
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

			cri_RI_CQI_WithoutNon_PMI_PortInd_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode cri_RI_CQI_WithoutNon_PMI_PortInd_r16 optional
			if cri_RI_CQI_WithoutNon_PMI_PortInd_r16Present {
				ie.cri_RI_CQI_WithoutNon_PMI_PortInd_r16 = new(Phy_ParametersFRX_Diff_cri_RI_CQI_WithoutNon_PMI_PortInd_r16)
				if err = ie.cri_RI_CQI_WithoutNon_PMI_PortInd_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode cri_RI_CQI_WithoutNon_PMI_PortInd_r16", err)
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

			cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 optional
			if cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17Present {
				ie.cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17 = new(Phy_ParametersFRX_Diff_cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17)
				if err = ie.cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode cqi_4_BitsSubbandTN_NonSharedSpectrumChAccess_r17", err)
				}
			}
		}
	}
	return nil
}
