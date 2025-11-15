package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PhysicalCellGroupConfig struct {
	harq_ACK_SpatialBundlingPUCCH                           *PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUCCH                           `optional`
	harq_ACK_SpatialBundlingPUSCH                           *PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUSCH                           `optional`
	p_NR_FR1                                                *P_Max                                                                           `optional`
	pdsch_HARQ_ACK_Codebook                                 PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook                                  `madatory`
	tpc_SRS_RNTI                                            *RNTI_Value                                                                      `optional`
	tpc_PUCCH_RNTI                                          *RNTI_Value                                                                      `optional`
	tpc_PUSCH_RNTI                                          *RNTI_Value                                                                      `optional`
	sp_CSI_RNTI                                             *RNTI_Value                                                                      `optional`
	cs_RNTI                                                 *RNTI_Value                                                                      `optional,setuprelease`
	mcs_C_RNTI                                              *RNTI_Value                                                                      `optional,ext-1`
	p_UE_FR1                                                *P_Max                                                                           `optional,ext-1`
	xScale                                                  *PhysicalCellGroupConfig_xScale                                                  `optional,ext-2`
	pdcch_BlindDetection                                    *PDCCH_BlindDetection                                                            `optional,ext-3,setuprelease`
	dcp_Config_r16                                          *DCP_Config_r16                                                                  `optional,ext-4,setuprelease`
	harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16   *PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16   `optional,ext-4`
	harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16   *PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16   `optional,ext-4`
	pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16         *PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16         `optional,ext-4`
	p_NR_FR2_r16                                            *P_Max                                                                           `optional,ext-4`
	p_UE_FR2_r16                                            *P_Max                                                                           `optional,ext-4`
	nrdc_PCmode_FR1_r16                                     *PhysicalCellGroupConfig_nrdc_PCmode_FR1_r16                                     `optional,ext-4`
	nrdc_PCmode_FR2_r16                                     *PhysicalCellGroupConfig_nrdc_PCmode_FR2_r16                                     `optional,ext-4`
	pdsch_HARQ_ACK_Codebook_r16                             *PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook_r16                             `optional,ext-4`
	nfi_TotalDAI_Included_r16                               *PhysicalCellGroupConfig_nfi_TotalDAI_Included_r16                               `optional,ext-4`
	ul_TotalDAI_Included_r16                                *PhysicalCellGroupConfig_ul_TotalDAI_Included_r16                                `optional,ext-4`
	pdsch_HARQ_ACK_OneShotFeedback_r16                      *PhysicalCellGroupConfig_pdsch_HARQ_ACK_OneShotFeedback_r16                      `optional,ext-4`
	pdsch_HARQ_ACK_OneShotFeedbackNDI_r16                   *PhysicalCellGroupConfig_pdsch_HARQ_ACK_OneShotFeedbackNDI_r16                   `optional,ext-4`
	pdsch_HARQ_ACK_OneShotFeedbackCBG_r16                   *PhysicalCellGroupConfig_pdsch_HARQ_ACK_OneShotFeedbackCBG_r16                   `optional,ext-4`
	downlinkAssignmentIndexDCI_0_2_r16                      *PhysicalCellGroupConfig_downlinkAssignmentIndexDCI_0_2_r16                      `optional,ext-4`
	downlinkAssignmentIndexDCI_1_2_r16                      *PhysicalCellGroupConfig_downlinkAssignmentIndexDCI_1_2_r16                      `optional,ext-4`
	pdsch_HARQ_ACK_CodebookList_r16                         *PDSCH_HARQ_ACK_CodebookList_r16                                                 `optional,ext-4,setuprelease`
	ackNackFeedbackMode_r16                                 *PhysicalCellGroupConfig_ackNackFeedbackMode_r16                                 `optional,ext-4`
	pdcch_BlindDetectionCA_CombIndicator_r16                *PDCCH_BlindDetectionCA_CombIndicator_r16                                        `optional,ext-4,setuprelease`
	pdcch_BlindDetection2_r16                               *PDCCH_BlindDetection2_r16                                                       `optional,ext-4,setuprelease`
	pdcch_BlindDetection3_r16                               *PDCCH_BlindDetection3_r16                                                       `optional,ext-4,setuprelease`
	bdFactorR_r16                                           *PhysicalCellGroupConfig_bdFactorR_r16                                           `optional,ext-4`
	pdsch_HARQ_ACK_EnhType3ToAddModList_r17                 []PDSCH_HARQ_ACK_EnhType3_r17                                                    `lb:1,ub:maxNrofEnhType3HARQ_ACK_r17,optional,ext-5`
	pdsch_HARQ_ACK_EnhType3ToReleaseList_r17                []PDSCH_HARQ_ACK_EnhType3Index_r17                                               `lb:1,ub:maxNrofEnhType3HARQ_ACK_r17,optional,ext-5`
	pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17        []PDSCH_HARQ_ACK_EnhType3_r17                                                    `lb:1,ub:maxNrofEnhType3HARQ_ACK_r17,optional,ext-5`
	pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17       []PDSCH_HARQ_ACK_EnhType3Index_r17                                               `lb:1,ub:maxNrofEnhType3HARQ_ACK_r17,optional,ext-5`
	pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 *PhysicalCellGroupConfig_pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 `optional,ext-5`
	pdsch_HARQ_ACK_EnhType3DCI_Field_r17                    *PhysicalCellGroupConfig_pdsch_HARQ_ACK_EnhType3DCI_Field_r17                    `optional,ext-5`
	pdsch_HARQ_ACK_Retx_r17                                 *PhysicalCellGroupConfig_pdsch_HARQ_ACK_Retx_r17                                 `optional,ext-5`
	pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17              *PhysicalCellGroupConfig_pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17              `optional,ext-5`
	pucch_sSCell_r17                                        *SCellIndex                                                                      `optional,ext-5`
	pucch_sSCellSecondaryPUCCHgroup_r17                     *SCellIndex                                                                      `optional,ext-5`
	pucch_sSCellDyn_r17                                     *PhysicalCellGroupConfig_pucch_sSCellDyn_r17                                     `optional,ext-5`
	pucch_sSCellDynSecondaryPUCCHgroup_r17                  *PhysicalCellGroupConfig_pucch_sSCellDynSecondaryPUCCHgroup_r17                  `optional,ext-5`
	pucch_sSCellPattern_r17                                 []int64                                                                          `lb:1,ub:maxNrofSlots,e_lb:0,e_ub:1,optional,ext-5`
	pucch_sSCellPatternSecondaryPUCCHgroup_r17              []int64                                                                          `lb:1,ub:maxNrofSlots,e_lb:0,e_ub:1,optional,ext-5`
	uci_MuxWithDiffPrio_r17                                 *PhysicalCellGroupConfig_uci_MuxWithDiffPrio_r17                                 `optional,ext-5`
	uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17              *PhysicalCellGroupConfig_uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17              `optional,ext-5`
	simultaneousPUCCH_PUSCH_r17                             *PhysicalCellGroupConfig_simultaneousPUCCH_PUSCH_r17                             `optional,ext-5`
	simultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17         *PhysicalCellGroupConfig_simultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17         `optional,ext-5`
	prioLowDG_HighCG_r17                                    *PhysicalCellGroupConfig_prioLowDG_HighCG_r17                                    `optional,ext-5`
	prioHighDG_LowCG_r17                                    *PhysicalCellGroupConfig_prioHighDG_LowCG_r17                                    `optional,ext-5`
	twoQCLTypeDforPDCCHRepetition_r17                       *PhysicalCellGroupConfig_twoQCLTypeDforPDCCHRepetition_r17                       `optional,ext-5`
	multicastConfig_r17                                     *MulticastConfig_r17                                                             `optional,ext-5,setuprelease`
	pdcch_BlindDetectionCA_CombIndicator_r17                *PDCCH_BlindDetectionCA_CombIndicator_r17                                        `optional,ext-5,setuprelease`
	simultaneousSR_PUSCH_diffPUCCH_Groups_r17               *PhysicalCellGroupConfig_simultaneousSR_PUSCH_diffPUCCH_Groups_r17               `optional,ext-6`
	intraBandNC_PRACH_simulTx_r17                           *PhysicalCellGroupConfig_intraBandNC_PRACH_simulTx_r17                           `optional,ext-7`
}

func (ie *PhysicalCellGroupConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.mcs_C_RNTI != nil || ie.p_UE_FR1 != nil || ie.xScale != nil || ie.pdcch_BlindDetection != nil || ie.dcp_Config_r16 != nil || ie.harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16 != nil || ie.harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16 != nil || ie.pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16 != nil || ie.p_NR_FR2_r16 != nil || ie.p_UE_FR2_r16 != nil || ie.nrdc_PCmode_FR1_r16 != nil || ie.nrdc_PCmode_FR2_r16 != nil || ie.pdsch_HARQ_ACK_Codebook_r16 != nil || ie.nfi_TotalDAI_Included_r16 != nil || ie.ul_TotalDAI_Included_r16 != nil || ie.pdsch_HARQ_ACK_OneShotFeedback_r16 != nil || ie.pdsch_HARQ_ACK_OneShotFeedbackNDI_r16 != nil || ie.pdsch_HARQ_ACK_OneShotFeedbackCBG_r16 != nil || ie.downlinkAssignmentIndexDCI_0_2_r16 != nil || ie.downlinkAssignmentIndexDCI_1_2_r16 != nil || ie.pdsch_HARQ_ACK_CodebookList_r16 != nil || ie.ackNackFeedbackMode_r16 != nil || ie.pdcch_BlindDetectionCA_CombIndicator_r16 != nil || ie.pdcch_BlindDetection2_r16 != nil || ie.pdcch_BlindDetection3_r16 != nil || ie.bdFactorR_r16 != nil || len(ie.pdsch_HARQ_ACK_EnhType3ToAddModList_r17) > 0 || len(ie.pdsch_HARQ_ACK_EnhType3ToReleaseList_r17) > 0 || len(ie.pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17) > 0 || len(ie.pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17) > 0 || ie.pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 != nil || ie.pdsch_HARQ_ACK_EnhType3DCI_Field_r17 != nil || ie.pdsch_HARQ_ACK_Retx_r17 != nil || ie.pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17 != nil || ie.pucch_sSCell_r17 != nil || ie.pucch_sSCellSecondaryPUCCHgroup_r17 != nil || ie.pucch_sSCellDyn_r17 != nil || ie.pucch_sSCellDynSecondaryPUCCHgroup_r17 != nil || len(ie.pucch_sSCellPattern_r17) > 0 || len(ie.pucch_sSCellPatternSecondaryPUCCHgroup_r17) > 0 || ie.uci_MuxWithDiffPrio_r17 != nil || ie.uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17 != nil || ie.simultaneousPUCCH_PUSCH_r17 != nil || ie.simultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17 != nil || ie.prioLowDG_HighCG_r17 != nil || ie.prioHighDG_LowCG_r17 != nil || ie.twoQCLTypeDforPDCCHRepetition_r17 != nil || ie.multicastConfig_r17 != nil || ie.pdcch_BlindDetectionCA_CombIndicator_r17 != nil || ie.simultaneousSR_PUSCH_diffPUCCH_Groups_r17 != nil || ie.intraBandNC_PRACH_simulTx_r17 != nil
	preambleBits := []bool{hasExtensions, ie.harq_ACK_SpatialBundlingPUCCH != nil, ie.harq_ACK_SpatialBundlingPUSCH != nil, ie.p_NR_FR1 != nil, ie.tpc_SRS_RNTI != nil, ie.tpc_PUCCH_RNTI != nil, ie.tpc_PUSCH_RNTI != nil, ie.sp_CSI_RNTI != nil, ie.cs_RNTI != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.harq_ACK_SpatialBundlingPUCCH != nil {
		if err = ie.harq_ACK_SpatialBundlingPUCCH.Encode(w); err != nil {
			return utils.WrapError("Encode harq_ACK_SpatialBundlingPUCCH", err)
		}
	}
	if ie.harq_ACK_SpatialBundlingPUSCH != nil {
		if err = ie.harq_ACK_SpatialBundlingPUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode harq_ACK_SpatialBundlingPUSCH", err)
		}
	}
	if ie.p_NR_FR1 != nil {
		if err = ie.p_NR_FR1.Encode(w); err != nil {
			return utils.WrapError("Encode p_NR_FR1", err)
		}
	}
	if err = ie.pdsch_HARQ_ACK_Codebook.Encode(w); err != nil {
		return utils.WrapError("Encode pdsch_HARQ_ACK_Codebook", err)
	}
	if ie.tpc_SRS_RNTI != nil {
		if err = ie.tpc_SRS_RNTI.Encode(w); err != nil {
			return utils.WrapError("Encode tpc_SRS_RNTI", err)
		}
	}
	if ie.tpc_PUCCH_RNTI != nil {
		if err = ie.tpc_PUCCH_RNTI.Encode(w); err != nil {
			return utils.WrapError("Encode tpc_PUCCH_RNTI", err)
		}
	}
	if ie.tpc_PUSCH_RNTI != nil {
		if err = ie.tpc_PUSCH_RNTI.Encode(w); err != nil {
			return utils.WrapError("Encode tpc_PUSCH_RNTI", err)
		}
	}
	if ie.sp_CSI_RNTI != nil {
		if err = ie.sp_CSI_RNTI.Encode(w); err != nil {
			return utils.WrapError("Encode sp_CSI_RNTI", err)
		}
	}
	if ie.cs_RNTI != nil {
		tmp_cs_RNTI := utils.SetupRelease[*RNTI_Value]{
			Setup: ie.cs_RNTI,
		}
		if err = tmp_cs_RNTI.Encode(w); err != nil {
			return utils.WrapError("Encode cs_RNTI", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 7 bits for 7 extension groups
		extBitmap := []bool{ie.mcs_C_RNTI != nil || ie.p_UE_FR1 != nil, ie.xScale != nil, ie.pdcch_BlindDetection != nil, ie.dcp_Config_r16 != nil || ie.harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16 != nil || ie.harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16 != nil || ie.pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16 != nil || ie.p_NR_FR2_r16 != nil || ie.p_UE_FR2_r16 != nil || ie.nrdc_PCmode_FR1_r16 != nil || ie.nrdc_PCmode_FR2_r16 != nil || ie.pdsch_HARQ_ACK_Codebook_r16 != nil || ie.nfi_TotalDAI_Included_r16 != nil || ie.ul_TotalDAI_Included_r16 != nil || ie.pdsch_HARQ_ACK_OneShotFeedback_r16 != nil || ie.pdsch_HARQ_ACK_OneShotFeedbackNDI_r16 != nil || ie.pdsch_HARQ_ACK_OneShotFeedbackCBG_r16 != nil || ie.downlinkAssignmentIndexDCI_0_2_r16 != nil || ie.downlinkAssignmentIndexDCI_1_2_r16 != nil || ie.pdsch_HARQ_ACK_CodebookList_r16 != nil || ie.ackNackFeedbackMode_r16 != nil || ie.pdcch_BlindDetectionCA_CombIndicator_r16 != nil || ie.pdcch_BlindDetection2_r16 != nil || ie.pdcch_BlindDetection3_r16 != nil || ie.bdFactorR_r16 != nil, len(ie.pdsch_HARQ_ACK_EnhType3ToAddModList_r17) > 0 || len(ie.pdsch_HARQ_ACK_EnhType3ToReleaseList_r17) > 0 || len(ie.pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17) > 0 || len(ie.pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17) > 0 || ie.pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 != nil || ie.pdsch_HARQ_ACK_EnhType3DCI_Field_r17 != nil || ie.pdsch_HARQ_ACK_Retx_r17 != nil || ie.pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17 != nil || ie.pucch_sSCell_r17 != nil || ie.pucch_sSCellSecondaryPUCCHgroup_r17 != nil || ie.pucch_sSCellDyn_r17 != nil || ie.pucch_sSCellDynSecondaryPUCCHgroup_r17 != nil || len(ie.pucch_sSCellPattern_r17) > 0 || len(ie.pucch_sSCellPatternSecondaryPUCCHgroup_r17) > 0 || ie.uci_MuxWithDiffPrio_r17 != nil || ie.uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17 != nil || ie.simultaneousPUCCH_PUSCH_r17 != nil || ie.simultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17 != nil || ie.prioLowDG_HighCG_r17 != nil || ie.prioHighDG_LowCG_r17 != nil || ie.twoQCLTypeDforPDCCHRepetition_r17 != nil || ie.multicastConfig_r17 != nil || ie.pdcch_BlindDetectionCA_CombIndicator_r17 != nil, ie.simultaneousSR_PUSCH_diffPUCCH_Groups_r17 != nil, ie.intraBandNC_PRACH_simulTx_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PhysicalCellGroupConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.mcs_C_RNTI != nil, ie.p_UE_FR1 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode mcs_C_RNTI optional
			if ie.mcs_C_RNTI != nil {
				if err = ie.mcs_C_RNTI.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mcs_C_RNTI", err)
				}
			}
			// encode p_UE_FR1 optional
			if ie.p_UE_FR1 != nil {
				if err = ie.p_UE_FR1.Encode(extWriter); err != nil {
					return utils.WrapError("Encode p_UE_FR1", err)
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
			optionals_ext_2 := []bool{ie.xScale != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode xScale optional
			if ie.xScale != nil {
				if err = ie.xScale.Encode(extWriter); err != nil {
					return utils.WrapError("Encode xScale", err)
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
			optionals_ext_3 := []bool{ie.pdcch_BlindDetection != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode pdcch_BlindDetection optional
			if ie.pdcch_BlindDetection != nil {
				tmp_pdcch_BlindDetection := utils.SetupRelease[*PDCCH_BlindDetection]{
					Setup: ie.pdcch_BlindDetection,
				}
				if err = tmp_pdcch_BlindDetection.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdcch_BlindDetection", err)
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
			optionals_ext_4 := []bool{ie.dcp_Config_r16 != nil, ie.harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16 != nil, ie.harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16 != nil, ie.pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16 != nil, ie.p_NR_FR2_r16 != nil, ie.p_UE_FR2_r16 != nil, ie.nrdc_PCmode_FR1_r16 != nil, ie.nrdc_PCmode_FR2_r16 != nil, ie.pdsch_HARQ_ACK_Codebook_r16 != nil, ie.nfi_TotalDAI_Included_r16 != nil, ie.ul_TotalDAI_Included_r16 != nil, ie.pdsch_HARQ_ACK_OneShotFeedback_r16 != nil, ie.pdsch_HARQ_ACK_OneShotFeedbackNDI_r16 != nil, ie.pdsch_HARQ_ACK_OneShotFeedbackCBG_r16 != nil, ie.downlinkAssignmentIndexDCI_0_2_r16 != nil, ie.downlinkAssignmentIndexDCI_1_2_r16 != nil, ie.pdsch_HARQ_ACK_CodebookList_r16 != nil, ie.ackNackFeedbackMode_r16 != nil, ie.pdcch_BlindDetectionCA_CombIndicator_r16 != nil, ie.pdcch_BlindDetection2_r16 != nil, ie.pdcch_BlindDetection3_r16 != nil, ie.bdFactorR_r16 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode dcp_Config_r16 optional
			if ie.dcp_Config_r16 != nil {
				tmp_dcp_Config_r16 := utils.SetupRelease[*DCP_Config_r16]{
					Setup: ie.dcp_Config_r16,
				}
				if err = tmp_dcp_Config_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dcp_Config_r16", err)
				}
			}
			// encode harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16 optional
			if ie.harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16 != nil {
				if err = ie.harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16", err)
				}
			}
			// encode harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16 optional
			if ie.harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16 != nil {
				if err = ie.harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16", err)
				}
			}
			// encode pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16 optional
			if ie.pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16 != nil {
				if err = ie.pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16", err)
				}
			}
			// encode p_NR_FR2_r16 optional
			if ie.p_NR_FR2_r16 != nil {
				if err = ie.p_NR_FR2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode p_NR_FR2_r16", err)
				}
			}
			// encode p_UE_FR2_r16 optional
			if ie.p_UE_FR2_r16 != nil {
				if err = ie.p_UE_FR2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode p_UE_FR2_r16", err)
				}
			}
			// encode nrdc_PCmode_FR1_r16 optional
			if ie.nrdc_PCmode_FR1_r16 != nil {
				if err = ie.nrdc_PCmode_FR1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nrdc_PCmode_FR1_r16", err)
				}
			}
			// encode nrdc_PCmode_FR2_r16 optional
			if ie.nrdc_PCmode_FR2_r16 != nil {
				if err = ie.nrdc_PCmode_FR2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nrdc_PCmode_FR2_r16", err)
				}
			}
			// encode pdsch_HARQ_ACK_Codebook_r16 optional
			if ie.pdsch_HARQ_ACK_Codebook_r16 != nil {
				if err = ie.pdsch_HARQ_ACK_Codebook_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_HARQ_ACK_Codebook_r16", err)
				}
			}
			// encode nfi_TotalDAI_Included_r16 optional
			if ie.nfi_TotalDAI_Included_r16 != nil {
				if err = ie.nfi_TotalDAI_Included_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nfi_TotalDAI_Included_r16", err)
				}
			}
			// encode ul_TotalDAI_Included_r16 optional
			if ie.ul_TotalDAI_Included_r16 != nil {
				if err = ie.ul_TotalDAI_Included_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_TotalDAI_Included_r16", err)
				}
			}
			// encode pdsch_HARQ_ACK_OneShotFeedback_r16 optional
			if ie.pdsch_HARQ_ACK_OneShotFeedback_r16 != nil {
				if err = ie.pdsch_HARQ_ACK_OneShotFeedback_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_HARQ_ACK_OneShotFeedback_r16", err)
				}
			}
			// encode pdsch_HARQ_ACK_OneShotFeedbackNDI_r16 optional
			if ie.pdsch_HARQ_ACK_OneShotFeedbackNDI_r16 != nil {
				if err = ie.pdsch_HARQ_ACK_OneShotFeedbackNDI_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_HARQ_ACK_OneShotFeedbackNDI_r16", err)
				}
			}
			// encode pdsch_HARQ_ACK_OneShotFeedbackCBG_r16 optional
			if ie.pdsch_HARQ_ACK_OneShotFeedbackCBG_r16 != nil {
				if err = ie.pdsch_HARQ_ACK_OneShotFeedbackCBG_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_HARQ_ACK_OneShotFeedbackCBG_r16", err)
				}
			}
			// encode downlinkAssignmentIndexDCI_0_2_r16 optional
			if ie.downlinkAssignmentIndexDCI_0_2_r16 != nil {
				if err = ie.downlinkAssignmentIndexDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode downlinkAssignmentIndexDCI_0_2_r16", err)
				}
			}
			// encode downlinkAssignmentIndexDCI_1_2_r16 optional
			if ie.downlinkAssignmentIndexDCI_1_2_r16 != nil {
				if err = ie.downlinkAssignmentIndexDCI_1_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode downlinkAssignmentIndexDCI_1_2_r16", err)
				}
			}
			// encode pdsch_HARQ_ACK_CodebookList_r16 optional
			if ie.pdsch_HARQ_ACK_CodebookList_r16 != nil {
				tmp_pdsch_HARQ_ACK_CodebookList_r16 := utils.SetupRelease[*PDSCH_HARQ_ACK_CodebookList_r16]{
					Setup: ie.pdsch_HARQ_ACK_CodebookList_r16,
				}
				if err = tmp_pdsch_HARQ_ACK_CodebookList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_HARQ_ACK_CodebookList_r16", err)
				}
			}
			// encode ackNackFeedbackMode_r16 optional
			if ie.ackNackFeedbackMode_r16 != nil {
				if err = ie.ackNackFeedbackMode_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ackNackFeedbackMode_r16", err)
				}
			}
			// encode pdcch_BlindDetectionCA_CombIndicator_r16 optional
			if ie.pdcch_BlindDetectionCA_CombIndicator_r16 != nil {
				tmp_pdcch_BlindDetectionCA_CombIndicator_r16 := utils.SetupRelease[*PDCCH_BlindDetectionCA_CombIndicator_r16]{
					Setup: ie.pdcch_BlindDetectionCA_CombIndicator_r16,
				}
				if err = tmp_pdcch_BlindDetectionCA_CombIndicator_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdcch_BlindDetectionCA_CombIndicator_r16", err)
				}
			}
			// encode pdcch_BlindDetection2_r16 optional
			if ie.pdcch_BlindDetection2_r16 != nil {
				tmp_pdcch_BlindDetection2_r16 := utils.SetupRelease[*PDCCH_BlindDetection2_r16]{
					Setup: ie.pdcch_BlindDetection2_r16,
				}
				if err = tmp_pdcch_BlindDetection2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdcch_BlindDetection2_r16", err)
				}
			}
			// encode pdcch_BlindDetection3_r16 optional
			if ie.pdcch_BlindDetection3_r16 != nil {
				tmp_pdcch_BlindDetection3_r16 := utils.SetupRelease[*PDCCH_BlindDetection3_r16]{
					Setup: ie.pdcch_BlindDetection3_r16,
				}
				if err = tmp_pdcch_BlindDetection3_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdcch_BlindDetection3_r16", err)
				}
			}
			// encode bdFactorR_r16 optional
			if ie.bdFactorR_r16 != nil {
				if err = ie.bdFactorR_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode bdFactorR_r16", err)
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
			optionals_ext_5 := []bool{len(ie.pdsch_HARQ_ACK_EnhType3ToAddModList_r17) > 0, len(ie.pdsch_HARQ_ACK_EnhType3ToReleaseList_r17) > 0, len(ie.pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17) > 0, len(ie.pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17) > 0, ie.pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 != nil, ie.pdsch_HARQ_ACK_EnhType3DCI_Field_r17 != nil, ie.pdsch_HARQ_ACK_Retx_r17 != nil, ie.pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17 != nil, ie.pucch_sSCell_r17 != nil, ie.pucch_sSCellSecondaryPUCCHgroup_r17 != nil, ie.pucch_sSCellDyn_r17 != nil, ie.pucch_sSCellDynSecondaryPUCCHgroup_r17 != nil, len(ie.pucch_sSCellPattern_r17) > 0, len(ie.pucch_sSCellPatternSecondaryPUCCHgroup_r17) > 0, ie.uci_MuxWithDiffPrio_r17 != nil, ie.uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17 != nil, ie.simultaneousPUCCH_PUSCH_r17 != nil, ie.simultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17 != nil, ie.prioLowDG_HighCG_r17 != nil, ie.prioHighDG_LowCG_r17 != nil, ie.twoQCLTypeDforPDCCHRepetition_r17 != nil, ie.multicastConfig_r17 != nil, ie.pdcch_BlindDetectionCA_CombIndicator_r17 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode pdsch_HARQ_ACK_EnhType3ToAddModList_r17 optional
			if len(ie.pdsch_HARQ_ACK_EnhType3ToAddModList_r17) > 0 {
				tmp_pdsch_HARQ_ACK_EnhType3ToAddModList_r17 := utils.NewSequence[*PDSCH_HARQ_ACK_EnhType3_r17]([]*PDSCH_HARQ_ACK_EnhType3_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofEnhType3HARQ_ACK_r17}, false)
				for _, i := range ie.pdsch_HARQ_ACK_EnhType3ToAddModList_r17 {
					tmp_pdsch_HARQ_ACK_EnhType3ToAddModList_r17.Value = append(tmp_pdsch_HARQ_ACK_EnhType3ToAddModList_r17.Value, &i)
				}
				if err = tmp_pdsch_HARQ_ACK_EnhType3ToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_HARQ_ACK_EnhType3ToAddModList_r17", err)
				}
			}
			// encode pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 optional
			if len(ie.pdsch_HARQ_ACK_EnhType3ToReleaseList_r17) > 0 {
				tmp_pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 := utils.NewSequence[*PDSCH_HARQ_ACK_EnhType3Index_r17]([]*PDSCH_HARQ_ACK_EnhType3Index_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofEnhType3HARQ_ACK_r17}, false)
				for _, i := range ie.pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 {
					tmp_pdsch_HARQ_ACK_EnhType3ToReleaseList_r17.Value = append(tmp_pdsch_HARQ_ACK_EnhType3ToReleaseList_r17.Value, &i)
				}
				if err = tmp_pdsch_HARQ_ACK_EnhType3ToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_HARQ_ACK_EnhType3ToReleaseList_r17", err)
				}
			}
			// encode pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 optional
			if len(ie.pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17) > 0 {
				tmp_pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 := utils.NewSequence[*PDSCH_HARQ_ACK_EnhType3_r17]([]*PDSCH_HARQ_ACK_EnhType3_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofEnhType3HARQ_ACK_r17}, false)
				for _, i := range ie.pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 {
					tmp_pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17.Value = append(tmp_pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17.Value, &i)
				}
				if err = tmp_pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17", err)
				}
			}
			// encode pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 optional
			if len(ie.pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17) > 0 {
				tmp_pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 := utils.NewSequence[*PDSCH_HARQ_ACK_EnhType3Index_r17]([]*PDSCH_HARQ_ACK_EnhType3Index_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofEnhType3HARQ_ACK_r17}, false)
				for _, i := range ie.pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 {
					tmp_pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17.Value = append(tmp_pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17.Value, &i)
				}
				if err = tmp_pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17", err)
				}
			}
			// encode pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 optional
			if ie.pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 != nil {
				if err = ie.pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17", err)
				}
			}
			// encode pdsch_HARQ_ACK_EnhType3DCI_Field_r17 optional
			if ie.pdsch_HARQ_ACK_EnhType3DCI_Field_r17 != nil {
				if err = ie.pdsch_HARQ_ACK_EnhType3DCI_Field_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_HARQ_ACK_EnhType3DCI_Field_r17", err)
				}
			}
			// encode pdsch_HARQ_ACK_Retx_r17 optional
			if ie.pdsch_HARQ_ACK_Retx_r17 != nil {
				if err = ie.pdsch_HARQ_ACK_Retx_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_HARQ_ACK_Retx_r17", err)
				}
			}
			// encode pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17 optional
			if ie.pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17 != nil {
				if err = ie.pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17", err)
				}
			}
			// encode pucch_sSCell_r17 optional
			if ie.pucch_sSCell_r17 != nil {
				if err = ie.pucch_sSCell_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pucch_sSCell_r17", err)
				}
			}
			// encode pucch_sSCellSecondaryPUCCHgroup_r17 optional
			if ie.pucch_sSCellSecondaryPUCCHgroup_r17 != nil {
				if err = ie.pucch_sSCellSecondaryPUCCHgroup_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pucch_sSCellSecondaryPUCCHgroup_r17", err)
				}
			}
			// encode pucch_sSCellDyn_r17 optional
			if ie.pucch_sSCellDyn_r17 != nil {
				if err = ie.pucch_sSCellDyn_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pucch_sSCellDyn_r17", err)
				}
			}
			// encode pucch_sSCellDynSecondaryPUCCHgroup_r17 optional
			if ie.pucch_sSCellDynSecondaryPUCCHgroup_r17 != nil {
				if err = ie.pucch_sSCellDynSecondaryPUCCHgroup_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pucch_sSCellDynSecondaryPUCCHgroup_r17", err)
				}
			}
			// encode pucch_sSCellPattern_r17 optional
			if len(ie.pucch_sSCellPattern_r17) > 0 {
				tmp_pucch_sSCellPattern_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
				for _, i := range ie.pucch_sSCellPattern_r17 {
					tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 1}, false)
					tmp_pucch_sSCellPattern_r17.Value = append(tmp_pucch_sSCellPattern_r17.Value, &tmp_ie)
				}
				if err = tmp_pucch_sSCellPattern_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pucch_sSCellPattern_r17", err)
				}
			}
			// encode pucch_sSCellPatternSecondaryPUCCHgroup_r17 optional
			if len(ie.pucch_sSCellPatternSecondaryPUCCHgroup_r17) > 0 {
				tmp_pucch_sSCellPatternSecondaryPUCCHgroup_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
				for _, i := range ie.pucch_sSCellPatternSecondaryPUCCHgroup_r17 {
					tmp_ie := utils.NewINTEGER(int64(i), uper.Constraint{Lb: 0, Ub: 1}, false)
					tmp_pucch_sSCellPatternSecondaryPUCCHgroup_r17.Value = append(tmp_pucch_sSCellPatternSecondaryPUCCHgroup_r17.Value, &tmp_ie)
				}
				if err = tmp_pucch_sSCellPatternSecondaryPUCCHgroup_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pucch_sSCellPatternSecondaryPUCCHgroup_r17", err)
				}
			}
			// encode uci_MuxWithDiffPrio_r17 optional
			if ie.uci_MuxWithDiffPrio_r17 != nil {
				if err = ie.uci_MuxWithDiffPrio_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uci_MuxWithDiffPrio_r17", err)
				}
			}
			// encode uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17 optional
			if ie.uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17 != nil {
				if err = ie.uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17", err)
				}
			}
			// encode simultaneousPUCCH_PUSCH_r17 optional
			if ie.simultaneousPUCCH_PUSCH_r17 != nil {
				if err = ie.simultaneousPUCCH_PUSCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simultaneousPUCCH_PUSCH_r17", err)
				}
			}
			// encode simultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17 optional
			if ie.simultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17 != nil {
				if err = ie.simultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17", err)
				}
			}
			// encode prioLowDG_HighCG_r17 optional
			if ie.prioLowDG_HighCG_r17 != nil {
				if err = ie.prioLowDG_HighCG_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode prioLowDG_HighCG_r17", err)
				}
			}
			// encode prioHighDG_LowCG_r17 optional
			if ie.prioHighDG_LowCG_r17 != nil {
				if err = ie.prioHighDG_LowCG_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode prioHighDG_LowCG_r17", err)
				}
			}
			// encode twoQCLTypeDforPDCCHRepetition_r17 optional
			if ie.twoQCLTypeDforPDCCHRepetition_r17 != nil {
				if err = ie.twoQCLTypeDforPDCCHRepetition_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode twoQCLTypeDforPDCCHRepetition_r17", err)
				}
			}
			// encode multicastConfig_r17 optional
			if ie.multicastConfig_r17 != nil {
				tmp_multicastConfig_r17 := utils.SetupRelease[*MulticastConfig_r17]{
					Setup: ie.multicastConfig_r17,
				}
				if err = tmp_multicastConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode multicastConfig_r17", err)
				}
			}
			// encode pdcch_BlindDetectionCA_CombIndicator_r17 optional
			if ie.pdcch_BlindDetectionCA_CombIndicator_r17 != nil {
				tmp_pdcch_BlindDetectionCA_CombIndicator_r17 := utils.SetupRelease[*PDCCH_BlindDetectionCA_CombIndicator_r17]{
					Setup: ie.pdcch_BlindDetectionCA_CombIndicator_r17,
				}
				if err = tmp_pdcch_BlindDetectionCA_CombIndicator_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdcch_BlindDetectionCA_CombIndicator_r17", err)
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
			optionals_ext_6 := []bool{ie.simultaneousSR_PUSCH_diffPUCCH_Groups_r17 != nil}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode simultaneousSR_PUSCH_diffPUCCH_Groups_r17 optional
			if ie.simultaneousSR_PUSCH_diffPUCCH_Groups_r17 != nil {
				if err = ie.simultaneousSR_PUSCH_diffPUCCH_Groups_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simultaneousSR_PUSCH_diffPUCCH_Groups_r17", err)
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
			optionals_ext_7 := []bool{ie.intraBandNC_PRACH_simulTx_r17 != nil}
			for _, bit := range optionals_ext_7 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode intraBandNC_PRACH_simulTx_r17 optional
			if ie.intraBandNC_PRACH_simulTx_r17 != nil {
				if err = ie.intraBandNC_PRACH_simulTx_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode intraBandNC_PRACH_simulTx_r17", err)
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

func (ie *PhysicalCellGroupConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var harq_ACK_SpatialBundlingPUCCHPresent bool
	if harq_ACK_SpatialBundlingPUCCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var harq_ACK_SpatialBundlingPUSCHPresent bool
	if harq_ACK_SpatialBundlingPUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var p_NR_FR1Present bool
	if p_NR_FR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tpc_SRS_RNTIPresent bool
	if tpc_SRS_RNTIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tpc_PUCCH_RNTIPresent bool
	if tpc_PUCCH_RNTIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tpc_PUSCH_RNTIPresent bool
	if tpc_PUSCH_RNTIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sp_CSI_RNTIPresent bool
	if sp_CSI_RNTIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var cs_RNTIPresent bool
	if cs_RNTIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if harq_ACK_SpatialBundlingPUCCHPresent {
		ie.harq_ACK_SpatialBundlingPUCCH = new(PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUCCH)
		if err = ie.harq_ACK_SpatialBundlingPUCCH.Decode(r); err != nil {
			return utils.WrapError("Decode harq_ACK_SpatialBundlingPUCCH", err)
		}
	}
	if harq_ACK_SpatialBundlingPUSCHPresent {
		ie.harq_ACK_SpatialBundlingPUSCH = new(PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUSCH)
		if err = ie.harq_ACK_SpatialBundlingPUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode harq_ACK_SpatialBundlingPUSCH", err)
		}
	}
	if p_NR_FR1Present {
		ie.p_NR_FR1 = new(P_Max)
		if err = ie.p_NR_FR1.Decode(r); err != nil {
			return utils.WrapError("Decode p_NR_FR1", err)
		}
	}
	if err = ie.pdsch_HARQ_ACK_Codebook.Decode(r); err != nil {
		return utils.WrapError("Decode pdsch_HARQ_ACK_Codebook", err)
	}
	if tpc_SRS_RNTIPresent {
		ie.tpc_SRS_RNTI = new(RNTI_Value)
		if err = ie.tpc_SRS_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode tpc_SRS_RNTI", err)
		}
	}
	if tpc_PUCCH_RNTIPresent {
		ie.tpc_PUCCH_RNTI = new(RNTI_Value)
		if err = ie.tpc_PUCCH_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode tpc_PUCCH_RNTI", err)
		}
	}
	if tpc_PUSCH_RNTIPresent {
		ie.tpc_PUSCH_RNTI = new(RNTI_Value)
		if err = ie.tpc_PUSCH_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode tpc_PUSCH_RNTI", err)
		}
	}
	if sp_CSI_RNTIPresent {
		ie.sp_CSI_RNTI = new(RNTI_Value)
		if err = ie.sp_CSI_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode sp_CSI_RNTI", err)
		}
	}
	if cs_RNTIPresent {
		tmp_cs_RNTI := utils.SetupRelease[*RNTI_Value]{}
		if err = tmp_cs_RNTI.Decode(r); err != nil {
			return utils.WrapError("Decode cs_RNTI", err)
		}
		ie.cs_RNTI = tmp_cs_RNTI.Setup
	}

	if extensionBit {
		// Read extension bitmap: 7 bits for 7 extension groups
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

			mcs_C_RNTIPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			p_UE_FR1Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode mcs_C_RNTI optional
			if mcs_C_RNTIPresent {
				ie.mcs_C_RNTI = new(RNTI_Value)
				if err = ie.mcs_C_RNTI.Decode(extReader); err != nil {
					return utils.WrapError("Decode mcs_C_RNTI", err)
				}
			}
			// decode p_UE_FR1 optional
			if p_UE_FR1Present {
				ie.p_UE_FR1 = new(P_Max)
				if err = ie.p_UE_FR1.Decode(extReader); err != nil {
					return utils.WrapError("Decode p_UE_FR1", err)
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

			xScalePresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode xScale optional
			if xScalePresent {
				ie.xScale = new(PhysicalCellGroupConfig_xScale)
				if err = ie.xScale.Decode(extReader); err != nil {
					return utils.WrapError("Decode xScale", err)
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

			pdcch_BlindDetectionPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode pdcch_BlindDetection optional
			if pdcch_BlindDetectionPresent {
				tmp_pdcch_BlindDetection := utils.SetupRelease[*PDCCH_BlindDetection]{}
				if err = tmp_pdcch_BlindDetection.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdcch_BlindDetection", err)
				}
				ie.pdcch_BlindDetection = tmp_pdcch_BlindDetection.Setup
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			dcp_Config_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			p_NR_FR2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			p_UE_FR2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nrdc_PCmode_FR1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nrdc_PCmode_FR2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_HARQ_ACK_Codebook_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nfi_TotalDAI_Included_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_TotalDAI_Included_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_HARQ_ACK_OneShotFeedback_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_HARQ_ACK_OneShotFeedbackNDI_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_HARQ_ACK_OneShotFeedbackCBG_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			downlinkAssignmentIndexDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			downlinkAssignmentIndexDCI_1_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_HARQ_ACK_CodebookList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ackNackFeedbackMode_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdcch_BlindDetectionCA_CombIndicator_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdcch_BlindDetection2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdcch_BlindDetection3_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			bdFactorR_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode dcp_Config_r16 optional
			if dcp_Config_r16Present {
				tmp_dcp_Config_r16 := utils.SetupRelease[*DCP_Config_r16]{}
				if err = tmp_dcp_Config_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode dcp_Config_r16", err)
				}
				ie.dcp_Config_r16 = tmp_dcp_Config_r16.Setup
			}
			// decode harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16 optional
			if harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16Present {
				ie.harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16 = new(PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16)
				if err = ie.harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode harq_ACK_SpatialBundlingPUCCH_secondaryPUCCHgroup_r16", err)
				}
			}
			// decode harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16 optional
			if harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16Present {
				ie.harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16 = new(PhysicalCellGroupConfig_harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16)
				if err = ie.harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode harq_ACK_SpatialBundlingPUSCH_secondaryPUCCHgroup_r16", err)
				}
			}
			// decode pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16 optional
			if pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16Present {
				ie.pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16 = new(PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16)
				if err = ie.pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_HARQ_ACK_Codebook_secondaryPUCCHgroup_r16", err)
				}
			}
			// decode p_NR_FR2_r16 optional
			if p_NR_FR2_r16Present {
				ie.p_NR_FR2_r16 = new(P_Max)
				if err = ie.p_NR_FR2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode p_NR_FR2_r16", err)
				}
			}
			// decode p_UE_FR2_r16 optional
			if p_UE_FR2_r16Present {
				ie.p_UE_FR2_r16 = new(P_Max)
				if err = ie.p_UE_FR2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode p_UE_FR2_r16", err)
				}
			}
			// decode nrdc_PCmode_FR1_r16 optional
			if nrdc_PCmode_FR1_r16Present {
				ie.nrdc_PCmode_FR1_r16 = new(PhysicalCellGroupConfig_nrdc_PCmode_FR1_r16)
				if err = ie.nrdc_PCmode_FR1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode nrdc_PCmode_FR1_r16", err)
				}
			}
			// decode nrdc_PCmode_FR2_r16 optional
			if nrdc_PCmode_FR2_r16Present {
				ie.nrdc_PCmode_FR2_r16 = new(PhysicalCellGroupConfig_nrdc_PCmode_FR2_r16)
				if err = ie.nrdc_PCmode_FR2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode nrdc_PCmode_FR2_r16", err)
				}
			}
			// decode pdsch_HARQ_ACK_Codebook_r16 optional
			if pdsch_HARQ_ACK_Codebook_r16Present {
				ie.pdsch_HARQ_ACK_Codebook_r16 = new(PhysicalCellGroupConfig_pdsch_HARQ_ACK_Codebook_r16)
				if err = ie.pdsch_HARQ_ACK_Codebook_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_HARQ_ACK_Codebook_r16", err)
				}
			}
			// decode nfi_TotalDAI_Included_r16 optional
			if nfi_TotalDAI_Included_r16Present {
				ie.nfi_TotalDAI_Included_r16 = new(PhysicalCellGroupConfig_nfi_TotalDAI_Included_r16)
				if err = ie.nfi_TotalDAI_Included_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode nfi_TotalDAI_Included_r16", err)
				}
			}
			// decode ul_TotalDAI_Included_r16 optional
			if ul_TotalDAI_Included_r16Present {
				ie.ul_TotalDAI_Included_r16 = new(PhysicalCellGroupConfig_ul_TotalDAI_Included_r16)
				if err = ie.ul_TotalDAI_Included_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_TotalDAI_Included_r16", err)
				}
			}
			// decode pdsch_HARQ_ACK_OneShotFeedback_r16 optional
			if pdsch_HARQ_ACK_OneShotFeedback_r16Present {
				ie.pdsch_HARQ_ACK_OneShotFeedback_r16 = new(PhysicalCellGroupConfig_pdsch_HARQ_ACK_OneShotFeedback_r16)
				if err = ie.pdsch_HARQ_ACK_OneShotFeedback_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_HARQ_ACK_OneShotFeedback_r16", err)
				}
			}
			// decode pdsch_HARQ_ACK_OneShotFeedbackNDI_r16 optional
			if pdsch_HARQ_ACK_OneShotFeedbackNDI_r16Present {
				ie.pdsch_HARQ_ACK_OneShotFeedbackNDI_r16 = new(PhysicalCellGroupConfig_pdsch_HARQ_ACK_OneShotFeedbackNDI_r16)
				if err = ie.pdsch_HARQ_ACK_OneShotFeedbackNDI_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_HARQ_ACK_OneShotFeedbackNDI_r16", err)
				}
			}
			// decode pdsch_HARQ_ACK_OneShotFeedbackCBG_r16 optional
			if pdsch_HARQ_ACK_OneShotFeedbackCBG_r16Present {
				ie.pdsch_HARQ_ACK_OneShotFeedbackCBG_r16 = new(PhysicalCellGroupConfig_pdsch_HARQ_ACK_OneShotFeedbackCBG_r16)
				if err = ie.pdsch_HARQ_ACK_OneShotFeedbackCBG_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_HARQ_ACK_OneShotFeedbackCBG_r16", err)
				}
			}
			// decode downlinkAssignmentIndexDCI_0_2_r16 optional
			if downlinkAssignmentIndexDCI_0_2_r16Present {
				ie.downlinkAssignmentIndexDCI_0_2_r16 = new(PhysicalCellGroupConfig_downlinkAssignmentIndexDCI_0_2_r16)
				if err = ie.downlinkAssignmentIndexDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode downlinkAssignmentIndexDCI_0_2_r16", err)
				}
			}
			// decode downlinkAssignmentIndexDCI_1_2_r16 optional
			if downlinkAssignmentIndexDCI_1_2_r16Present {
				ie.downlinkAssignmentIndexDCI_1_2_r16 = new(PhysicalCellGroupConfig_downlinkAssignmentIndexDCI_1_2_r16)
				if err = ie.downlinkAssignmentIndexDCI_1_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode downlinkAssignmentIndexDCI_1_2_r16", err)
				}
			}
			// decode pdsch_HARQ_ACK_CodebookList_r16 optional
			if pdsch_HARQ_ACK_CodebookList_r16Present {
				tmp_pdsch_HARQ_ACK_CodebookList_r16 := utils.SetupRelease[*PDSCH_HARQ_ACK_CodebookList_r16]{}
				if err = tmp_pdsch_HARQ_ACK_CodebookList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_HARQ_ACK_CodebookList_r16", err)
				}
				ie.pdsch_HARQ_ACK_CodebookList_r16 = tmp_pdsch_HARQ_ACK_CodebookList_r16.Setup
			}
			// decode ackNackFeedbackMode_r16 optional
			if ackNackFeedbackMode_r16Present {
				ie.ackNackFeedbackMode_r16 = new(PhysicalCellGroupConfig_ackNackFeedbackMode_r16)
				if err = ie.ackNackFeedbackMode_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ackNackFeedbackMode_r16", err)
				}
			}
			// decode pdcch_BlindDetectionCA_CombIndicator_r16 optional
			if pdcch_BlindDetectionCA_CombIndicator_r16Present {
				tmp_pdcch_BlindDetectionCA_CombIndicator_r16 := utils.SetupRelease[*PDCCH_BlindDetectionCA_CombIndicator_r16]{}
				if err = tmp_pdcch_BlindDetectionCA_CombIndicator_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdcch_BlindDetectionCA_CombIndicator_r16", err)
				}
				ie.pdcch_BlindDetectionCA_CombIndicator_r16 = tmp_pdcch_BlindDetectionCA_CombIndicator_r16.Setup
			}
			// decode pdcch_BlindDetection2_r16 optional
			if pdcch_BlindDetection2_r16Present {
				tmp_pdcch_BlindDetection2_r16 := utils.SetupRelease[*PDCCH_BlindDetection2_r16]{}
				if err = tmp_pdcch_BlindDetection2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdcch_BlindDetection2_r16", err)
				}
				ie.pdcch_BlindDetection2_r16 = tmp_pdcch_BlindDetection2_r16.Setup
			}
			// decode pdcch_BlindDetection3_r16 optional
			if pdcch_BlindDetection3_r16Present {
				tmp_pdcch_BlindDetection3_r16 := utils.SetupRelease[*PDCCH_BlindDetection3_r16]{}
				if err = tmp_pdcch_BlindDetection3_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdcch_BlindDetection3_r16", err)
				}
				ie.pdcch_BlindDetection3_r16 = tmp_pdcch_BlindDetection3_r16.Setup
			}
			// decode bdFactorR_r16 optional
			if bdFactorR_r16Present {
				ie.bdFactorR_r16 = new(PhysicalCellGroupConfig_bdFactorR_r16)
				if err = ie.bdFactorR_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode bdFactorR_r16", err)
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

			pdsch_HARQ_ACK_EnhType3ToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_HARQ_ACK_EnhType3ToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_HARQ_ACK_EnhType3DCI_Field_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_HARQ_ACK_Retx_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pucch_sSCell_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pucch_sSCellSecondaryPUCCHgroup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pucch_sSCellDyn_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pucch_sSCellDynSecondaryPUCCHgroup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pucch_sSCellPattern_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pucch_sSCellPatternSecondaryPUCCHgroup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uci_MuxWithDiffPrio_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			simultaneousPUCCH_PUSCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			simultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			prioLowDG_HighCG_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			prioHighDG_LowCG_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			twoQCLTypeDforPDCCHRepetition_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			multicastConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdcch_BlindDetectionCA_CombIndicator_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode pdsch_HARQ_ACK_EnhType3ToAddModList_r17 optional
			if pdsch_HARQ_ACK_EnhType3ToAddModList_r17Present {
				tmp_pdsch_HARQ_ACK_EnhType3ToAddModList_r17 := utils.NewSequence[*PDSCH_HARQ_ACK_EnhType3_r17]([]*PDSCH_HARQ_ACK_EnhType3_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofEnhType3HARQ_ACK_r17}, false)
				fn_pdsch_HARQ_ACK_EnhType3ToAddModList_r17 := func() *PDSCH_HARQ_ACK_EnhType3_r17 {
					return new(PDSCH_HARQ_ACK_EnhType3_r17)
				}
				if err = tmp_pdsch_HARQ_ACK_EnhType3ToAddModList_r17.Decode(extReader, fn_pdsch_HARQ_ACK_EnhType3ToAddModList_r17); err != nil {
					return utils.WrapError("Decode pdsch_HARQ_ACK_EnhType3ToAddModList_r17", err)
				}
				ie.pdsch_HARQ_ACK_EnhType3ToAddModList_r17 = []PDSCH_HARQ_ACK_EnhType3_r17{}
				for _, i := range tmp_pdsch_HARQ_ACK_EnhType3ToAddModList_r17.Value {
					ie.pdsch_HARQ_ACK_EnhType3ToAddModList_r17 = append(ie.pdsch_HARQ_ACK_EnhType3ToAddModList_r17, *i)
				}
			}
			// decode pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 optional
			if pdsch_HARQ_ACK_EnhType3ToReleaseList_r17Present {
				tmp_pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 := utils.NewSequence[*PDSCH_HARQ_ACK_EnhType3Index_r17]([]*PDSCH_HARQ_ACK_EnhType3Index_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofEnhType3HARQ_ACK_r17}, false)
				fn_pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 := func() *PDSCH_HARQ_ACK_EnhType3Index_r17 {
					return new(PDSCH_HARQ_ACK_EnhType3Index_r17)
				}
				if err = tmp_pdsch_HARQ_ACK_EnhType3ToReleaseList_r17.Decode(extReader, fn_pdsch_HARQ_ACK_EnhType3ToReleaseList_r17); err != nil {
					return utils.WrapError("Decode pdsch_HARQ_ACK_EnhType3ToReleaseList_r17", err)
				}
				ie.pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 = []PDSCH_HARQ_ACK_EnhType3Index_r17{}
				for _, i := range tmp_pdsch_HARQ_ACK_EnhType3ToReleaseList_r17.Value {
					ie.pdsch_HARQ_ACK_EnhType3ToReleaseList_r17 = append(ie.pdsch_HARQ_ACK_EnhType3ToReleaseList_r17, *i)
				}
			}
			// decode pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 optional
			if pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17Present {
				tmp_pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 := utils.NewSequence[*PDSCH_HARQ_ACK_EnhType3_r17]([]*PDSCH_HARQ_ACK_EnhType3_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofEnhType3HARQ_ACK_r17}, false)
				fn_pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 := func() *PDSCH_HARQ_ACK_EnhType3_r17 {
					return new(PDSCH_HARQ_ACK_EnhType3_r17)
				}
				if err = tmp_pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17.Decode(extReader, fn_pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17); err != nil {
					return utils.WrapError("Decode pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17", err)
				}
				ie.pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 = []PDSCH_HARQ_ACK_EnhType3_r17{}
				for _, i := range tmp_pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17.Value {
					ie.pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17 = append(ie.pdsch_HARQ_ACK_EnhType3SecondaryToAddModList_r17, *i)
				}
			}
			// decode pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 optional
			if pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17Present {
				tmp_pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 := utils.NewSequence[*PDSCH_HARQ_ACK_EnhType3Index_r17]([]*PDSCH_HARQ_ACK_EnhType3Index_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofEnhType3HARQ_ACK_r17}, false)
				fn_pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 := func() *PDSCH_HARQ_ACK_EnhType3Index_r17 {
					return new(PDSCH_HARQ_ACK_EnhType3Index_r17)
				}
				if err = tmp_pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17.Decode(extReader, fn_pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17); err != nil {
					return utils.WrapError("Decode pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17", err)
				}
				ie.pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 = []PDSCH_HARQ_ACK_EnhType3Index_r17{}
				for _, i := range tmp_pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17.Value {
					ie.pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17 = append(ie.pdsch_HARQ_ACK_EnhType3SecondaryToReleaseList_r17, *i)
				}
			}
			// decode pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 optional
			if pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17Present {
				ie.pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17 = new(PhysicalCellGroupConfig_pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17)
				if err = ie.pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_HARQ_ACK_EnhType3DCI_FieldSecondaryPUCCHgroup_r17", err)
				}
			}
			// decode pdsch_HARQ_ACK_EnhType3DCI_Field_r17 optional
			if pdsch_HARQ_ACK_EnhType3DCI_Field_r17Present {
				ie.pdsch_HARQ_ACK_EnhType3DCI_Field_r17 = new(PhysicalCellGroupConfig_pdsch_HARQ_ACK_EnhType3DCI_Field_r17)
				if err = ie.pdsch_HARQ_ACK_EnhType3DCI_Field_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_HARQ_ACK_EnhType3DCI_Field_r17", err)
				}
			}
			// decode pdsch_HARQ_ACK_Retx_r17 optional
			if pdsch_HARQ_ACK_Retx_r17Present {
				ie.pdsch_HARQ_ACK_Retx_r17 = new(PhysicalCellGroupConfig_pdsch_HARQ_ACK_Retx_r17)
				if err = ie.pdsch_HARQ_ACK_Retx_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_HARQ_ACK_Retx_r17", err)
				}
			}
			// decode pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17 optional
			if pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17Present {
				ie.pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17 = new(PhysicalCellGroupConfig_pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17)
				if err = ie.pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_HARQ_ACK_RetxSecondaryPUCCHgroup_r17", err)
				}
			}
			// decode pucch_sSCell_r17 optional
			if pucch_sSCell_r17Present {
				ie.pucch_sSCell_r17 = new(SCellIndex)
				if err = ie.pucch_sSCell_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pucch_sSCell_r17", err)
				}
			}
			// decode pucch_sSCellSecondaryPUCCHgroup_r17 optional
			if pucch_sSCellSecondaryPUCCHgroup_r17Present {
				ie.pucch_sSCellSecondaryPUCCHgroup_r17 = new(SCellIndex)
				if err = ie.pucch_sSCellSecondaryPUCCHgroup_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pucch_sSCellSecondaryPUCCHgroup_r17", err)
				}
			}
			// decode pucch_sSCellDyn_r17 optional
			if pucch_sSCellDyn_r17Present {
				ie.pucch_sSCellDyn_r17 = new(PhysicalCellGroupConfig_pucch_sSCellDyn_r17)
				if err = ie.pucch_sSCellDyn_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pucch_sSCellDyn_r17", err)
				}
			}
			// decode pucch_sSCellDynSecondaryPUCCHgroup_r17 optional
			if pucch_sSCellDynSecondaryPUCCHgroup_r17Present {
				ie.pucch_sSCellDynSecondaryPUCCHgroup_r17 = new(PhysicalCellGroupConfig_pucch_sSCellDynSecondaryPUCCHgroup_r17)
				if err = ie.pucch_sSCellDynSecondaryPUCCHgroup_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pucch_sSCellDynSecondaryPUCCHgroup_r17", err)
				}
			}
			// decode pucch_sSCellPattern_r17 optional
			if pucch_sSCellPattern_r17Present {
				tmp_pucch_sSCellPattern_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
				fn_pucch_sSCellPattern_r17 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 1}, false)
					return &ie
				}
				if err = tmp_pucch_sSCellPattern_r17.Decode(extReader, fn_pucch_sSCellPattern_r17); err != nil {
					return utils.WrapError("Decode pucch_sSCellPattern_r17", err)
				}
				ie.pucch_sSCellPattern_r17 = []int64{}
				for _, i := range tmp_pucch_sSCellPattern_r17.Value {
					ie.pucch_sSCellPattern_r17 = append(ie.pucch_sSCellPattern_r17, int64(i.Value))
				}
			}
			// decode pucch_sSCellPatternSecondaryPUCCHgroup_r17 optional
			if pucch_sSCellPatternSecondaryPUCCHgroup_r17Present {
				tmp_pucch_sSCellPatternSecondaryPUCCHgroup_r17 := utils.NewSequence[*utils.INTEGER]([]*utils.INTEGER{}, uper.Constraint{Lb: 1, Ub: maxNrofSlots}, false)
				fn_pucch_sSCellPatternSecondaryPUCCHgroup_r17 := func() *utils.INTEGER {
					ie := utils.NewINTEGER(0, uper.Constraint{Lb: 0, Ub: 1}, false)
					return &ie
				}
				if err = tmp_pucch_sSCellPatternSecondaryPUCCHgroup_r17.Decode(extReader, fn_pucch_sSCellPatternSecondaryPUCCHgroup_r17); err != nil {
					return utils.WrapError("Decode pucch_sSCellPatternSecondaryPUCCHgroup_r17", err)
				}
				ie.pucch_sSCellPatternSecondaryPUCCHgroup_r17 = []int64{}
				for _, i := range tmp_pucch_sSCellPatternSecondaryPUCCHgroup_r17.Value {
					ie.pucch_sSCellPatternSecondaryPUCCHgroup_r17 = append(ie.pucch_sSCellPatternSecondaryPUCCHgroup_r17, int64(i.Value))
				}
			}
			// decode uci_MuxWithDiffPrio_r17 optional
			if uci_MuxWithDiffPrio_r17Present {
				ie.uci_MuxWithDiffPrio_r17 = new(PhysicalCellGroupConfig_uci_MuxWithDiffPrio_r17)
				if err = ie.uci_MuxWithDiffPrio_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode uci_MuxWithDiffPrio_r17", err)
				}
			}
			// decode uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17 optional
			if uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17Present {
				ie.uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17 = new(PhysicalCellGroupConfig_uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17)
				if err = ie.uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode uci_MuxWithDiffPrioSecondaryPUCCHgroup_r17", err)
				}
			}
			// decode simultaneousPUCCH_PUSCH_r17 optional
			if simultaneousPUCCH_PUSCH_r17Present {
				ie.simultaneousPUCCH_PUSCH_r17 = new(PhysicalCellGroupConfig_simultaneousPUCCH_PUSCH_r17)
				if err = ie.simultaneousPUCCH_PUSCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode simultaneousPUCCH_PUSCH_r17", err)
				}
			}
			// decode simultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17 optional
			if simultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17Present {
				ie.simultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17 = new(PhysicalCellGroupConfig_simultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17)
				if err = ie.simultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode simultaneousPUCCH_PUSCH_SecondaryPUCCHgroup_r17", err)
				}
			}
			// decode prioLowDG_HighCG_r17 optional
			if prioLowDG_HighCG_r17Present {
				ie.prioLowDG_HighCG_r17 = new(PhysicalCellGroupConfig_prioLowDG_HighCG_r17)
				if err = ie.prioLowDG_HighCG_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode prioLowDG_HighCG_r17", err)
				}
			}
			// decode prioHighDG_LowCG_r17 optional
			if prioHighDG_LowCG_r17Present {
				ie.prioHighDG_LowCG_r17 = new(PhysicalCellGroupConfig_prioHighDG_LowCG_r17)
				if err = ie.prioHighDG_LowCG_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode prioHighDG_LowCG_r17", err)
				}
			}
			// decode twoQCLTypeDforPDCCHRepetition_r17 optional
			if twoQCLTypeDforPDCCHRepetition_r17Present {
				ie.twoQCLTypeDforPDCCHRepetition_r17 = new(PhysicalCellGroupConfig_twoQCLTypeDforPDCCHRepetition_r17)
				if err = ie.twoQCLTypeDforPDCCHRepetition_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode twoQCLTypeDforPDCCHRepetition_r17", err)
				}
			}
			// decode multicastConfig_r17 optional
			if multicastConfig_r17Present {
				tmp_multicastConfig_r17 := utils.SetupRelease[*MulticastConfig_r17]{}
				if err = tmp_multicastConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode multicastConfig_r17", err)
				}
				ie.multicastConfig_r17 = tmp_multicastConfig_r17.Setup
			}
			// decode pdcch_BlindDetectionCA_CombIndicator_r17 optional
			if pdcch_BlindDetectionCA_CombIndicator_r17Present {
				tmp_pdcch_BlindDetectionCA_CombIndicator_r17 := utils.SetupRelease[*PDCCH_BlindDetectionCA_CombIndicator_r17]{}
				if err = tmp_pdcch_BlindDetectionCA_CombIndicator_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdcch_BlindDetectionCA_CombIndicator_r17", err)
				}
				ie.pdcch_BlindDetectionCA_CombIndicator_r17 = tmp_pdcch_BlindDetectionCA_CombIndicator_r17.Setup
			}
		}
		// decode extension group 6
		if len(extBitmap) > 5 && extBitmap[5] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			simultaneousSR_PUSCH_diffPUCCH_Groups_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode simultaneousSR_PUSCH_diffPUCCH_Groups_r17 optional
			if simultaneousSR_PUSCH_diffPUCCH_Groups_r17Present {
				ie.simultaneousSR_PUSCH_diffPUCCH_Groups_r17 = new(PhysicalCellGroupConfig_simultaneousSR_PUSCH_diffPUCCH_Groups_r17)
				if err = ie.simultaneousSR_PUSCH_diffPUCCH_Groups_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode simultaneousSR_PUSCH_diffPUCCH_Groups_r17", err)
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

			intraBandNC_PRACH_simulTx_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode intraBandNC_PRACH_simulTx_r17 optional
			if intraBandNC_PRACH_simulTx_r17Present {
				ie.intraBandNC_PRACH_simulTx_r17 = new(PhysicalCellGroupConfig_intraBandNC_PRACH_simulTx_r17)
				if err = ie.intraBandNC_PRACH_simulTx_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode intraBandNC_PRACH_simulTx_r17", err)
				}
			}
		}
	}
	return nil
}
