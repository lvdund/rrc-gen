package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MAC_ParametersCommon struct {
	lcp_Restriction                           *MAC_ParametersCommon_lcp_Restriction                           `optional`
	dummy                                     *MAC_ParametersCommon_dummy                                     `optional`
	lch_ToSCellRestriction                    *MAC_ParametersCommon_lch_ToSCellRestriction                    `optional`
	recommendedBitRate                        *MAC_ParametersCommon_recommendedBitRate                        `optional,ext-1`
	recommendedBitRateQuery                   *MAC_ParametersCommon_recommendedBitRateQuery                   `optional,ext-1`
	recommendedBitRateMultiplier_r16          *MAC_ParametersCommon_recommendedBitRateMultiplier_r16          `optional,ext-2`
	preEmptiveBSR_r16                         *MAC_ParametersCommon_preEmptiveBSR_r16                         `optional,ext-2`
	autonomousTransmission_r16                *MAC_ParametersCommon_autonomousTransmission_r16                `optional,ext-2`
	lch_PriorityBasedPrioritization_r16       *MAC_ParametersCommon_lch_PriorityBasedPrioritization_r16       `optional,ext-2`
	lch_ToConfiguredGrantMapping_r16          *MAC_ParametersCommon_lch_ToConfiguredGrantMapping_r16          `optional,ext-2`
	lch_ToGrantPriorityRestriction_r16        *MAC_ParametersCommon_lch_ToGrantPriorityRestriction_r16        `optional,ext-2`
	singlePHR_P_r16                           *MAC_ParametersCommon_singlePHR_P_r16                           `optional,ext-2`
	ul_LBT_FailureDetectionRecovery_r16       *MAC_ParametersCommon_ul_LBT_FailureDetectionRecovery_r16       `optional,ext-2`
	tdd_MPE_P_MPR_Reporting_r16               *MAC_ParametersCommon_tdd_MPE_P_MPR_Reporting_r16               `optional,ext-2`
	lcid_ExtensionIAB_r16                     *MAC_ParametersCommon_lcid_ExtensionIAB_r16                     `optional,ext-2`
	spCell_BFR_CBRA_r16                       *MAC_ParametersCommon_spCell_BFR_CBRA_r16                       `optional,ext-3`
	srs_ResourceId_Ext_r16                    *MAC_ParametersCommon_srs_ResourceId_Ext_r16                    `optional,ext-4`
	enhancedUuDRX_forSidelink_r17             *MAC_ParametersCommon_enhancedUuDRX_forSidelink_r17             `optional,ext-5`
	mg_ActivationRequestPRS_Meas_r17          *MAC_ParametersCommon_mg_ActivationRequestPRS_Meas_r17          `optional,ext-5`
	mg_ActivationCommPRS_Meas_r17             *MAC_ParametersCommon_mg_ActivationCommPRS_Meas_r17             `optional,ext-5`
	intraCG_Prioritization_r17                *MAC_ParametersCommon_intraCG_Prioritization_r17                `optional,ext-5`
	jointPrioritizationCG_Retx_Timer_r17      *MAC_ParametersCommon_jointPrioritizationCG_Retx_Timer_r17      `optional,ext-5`
	survivalTime_r17                          *MAC_ParametersCommon_survivalTime_r17                          `optional,ext-5`
	lcg_ExtensionIAB_r17                      *MAC_ParametersCommon_lcg_ExtensionIAB_r17                      `optional,ext-5`
	harq_FeedbackDisabled_r17                 *MAC_ParametersCommon_harq_FeedbackDisabled_r17                 `optional,ext-5`
	uplink_Harq_ModeB_r17                     *MAC_ParametersCommon_uplink_Harq_ModeB_r17                     `optional,ext-5`
	sr_TriggeredBy_TA_Report_r17              *MAC_ParametersCommon_sr_TriggeredBy_TA_Report_r17              `optional,ext-5`
	extendedDRX_CycleInactive_r17             *MAC_ParametersCommon_extendedDRX_CycleInactive_r17             `optional,ext-5`
	simultaneousSR_PUSCH_DiffPUCCH_groups_r17 *MAC_ParametersCommon_simultaneousSR_PUSCH_DiffPUCCH_groups_r17 `optional,ext-5`
	lastTransmissionUL_r17                    *MAC_ParametersCommon_lastTransmissionUL_r17                    `optional,ext-5`
}

func (ie *MAC_ParametersCommon) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.recommendedBitRate != nil || ie.recommendedBitRateQuery != nil || ie.recommendedBitRateMultiplier_r16 != nil || ie.preEmptiveBSR_r16 != nil || ie.autonomousTransmission_r16 != nil || ie.lch_PriorityBasedPrioritization_r16 != nil || ie.lch_ToConfiguredGrantMapping_r16 != nil || ie.lch_ToGrantPriorityRestriction_r16 != nil || ie.singlePHR_P_r16 != nil || ie.ul_LBT_FailureDetectionRecovery_r16 != nil || ie.tdd_MPE_P_MPR_Reporting_r16 != nil || ie.lcid_ExtensionIAB_r16 != nil || ie.spCell_BFR_CBRA_r16 != nil || ie.srs_ResourceId_Ext_r16 != nil || ie.enhancedUuDRX_forSidelink_r17 != nil || ie.mg_ActivationRequestPRS_Meas_r17 != nil || ie.mg_ActivationCommPRS_Meas_r17 != nil || ie.intraCG_Prioritization_r17 != nil || ie.jointPrioritizationCG_Retx_Timer_r17 != nil || ie.survivalTime_r17 != nil || ie.lcg_ExtensionIAB_r17 != nil || ie.harq_FeedbackDisabled_r17 != nil || ie.uplink_Harq_ModeB_r17 != nil || ie.sr_TriggeredBy_TA_Report_r17 != nil || ie.extendedDRX_CycleInactive_r17 != nil || ie.simultaneousSR_PUSCH_DiffPUCCH_groups_r17 != nil || ie.lastTransmissionUL_r17 != nil
	preambleBits := []bool{hasExtensions, ie.lcp_Restriction != nil, ie.dummy != nil, ie.lch_ToSCellRestriction != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.lcp_Restriction != nil {
		if err = ie.lcp_Restriction.Encode(w); err != nil {
			return utils.WrapError("Encode lcp_Restriction", err)
		}
	}
	if ie.dummy != nil {
		if err = ie.dummy.Encode(w); err != nil {
			return utils.WrapError("Encode dummy", err)
		}
	}
	if ie.lch_ToSCellRestriction != nil {
		if err = ie.lch_ToSCellRestriction.Encode(w); err != nil {
			return utils.WrapError("Encode lch_ToSCellRestriction", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 5 bits for 5 extension groups
		extBitmap := []bool{ie.recommendedBitRate != nil || ie.recommendedBitRateQuery != nil, ie.recommendedBitRateMultiplier_r16 != nil || ie.preEmptiveBSR_r16 != nil || ie.autonomousTransmission_r16 != nil || ie.lch_PriorityBasedPrioritization_r16 != nil || ie.lch_ToConfiguredGrantMapping_r16 != nil || ie.lch_ToGrantPriorityRestriction_r16 != nil || ie.singlePHR_P_r16 != nil || ie.ul_LBT_FailureDetectionRecovery_r16 != nil || ie.tdd_MPE_P_MPR_Reporting_r16 != nil || ie.lcid_ExtensionIAB_r16 != nil, ie.spCell_BFR_CBRA_r16 != nil, ie.srs_ResourceId_Ext_r16 != nil, ie.enhancedUuDRX_forSidelink_r17 != nil || ie.mg_ActivationRequestPRS_Meas_r17 != nil || ie.mg_ActivationCommPRS_Meas_r17 != nil || ie.intraCG_Prioritization_r17 != nil || ie.jointPrioritizationCG_Retx_Timer_r17 != nil || ie.survivalTime_r17 != nil || ie.lcg_ExtensionIAB_r17 != nil || ie.harq_FeedbackDisabled_r17 != nil || ie.uplink_Harq_ModeB_r17 != nil || ie.sr_TriggeredBy_TA_Report_r17 != nil || ie.extendedDRX_CycleInactive_r17 != nil || ie.simultaneousSR_PUSCH_DiffPUCCH_groups_r17 != nil || ie.lastTransmissionUL_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MAC_ParametersCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.recommendedBitRate != nil, ie.recommendedBitRateQuery != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode recommendedBitRate optional
			if ie.recommendedBitRate != nil {
				if err = ie.recommendedBitRate.Encode(extWriter); err != nil {
					return utils.WrapError("Encode recommendedBitRate", err)
				}
			}
			// encode recommendedBitRateQuery optional
			if ie.recommendedBitRateQuery != nil {
				if err = ie.recommendedBitRateQuery.Encode(extWriter); err != nil {
					return utils.WrapError("Encode recommendedBitRateQuery", err)
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
			optionals_ext_2 := []bool{ie.recommendedBitRateMultiplier_r16 != nil, ie.preEmptiveBSR_r16 != nil, ie.autonomousTransmission_r16 != nil, ie.lch_PriorityBasedPrioritization_r16 != nil, ie.lch_ToConfiguredGrantMapping_r16 != nil, ie.lch_ToGrantPriorityRestriction_r16 != nil, ie.singlePHR_P_r16 != nil, ie.ul_LBT_FailureDetectionRecovery_r16 != nil, ie.tdd_MPE_P_MPR_Reporting_r16 != nil, ie.lcid_ExtensionIAB_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode recommendedBitRateMultiplier_r16 optional
			if ie.recommendedBitRateMultiplier_r16 != nil {
				if err = ie.recommendedBitRateMultiplier_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode recommendedBitRateMultiplier_r16", err)
				}
			}
			// encode preEmptiveBSR_r16 optional
			if ie.preEmptiveBSR_r16 != nil {
				if err = ie.preEmptiveBSR_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode preEmptiveBSR_r16", err)
				}
			}
			// encode autonomousTransmission_r16 optional
			if ie.autonomousTransmission_r16 != nil {
				if err = ie.autonomousTransmission_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode autonomousTransmission_r16", err)
				}
			}
			// encode lch_PriorityBasedPrioritization_r16 optional
			if ie.lch_PriorityBasedPrioritization_r16 != nil {
				if err = ie.lch_PriorityBasedPrioritization_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode lch_PriorityBasedPrioritization_r16", err)
				}
			}
			// encode lch_ToConfiguredGrantMapping_r16 optional
			if ie.lch_ToConfiguredGrantMapping_r16 != nil {
				if err = ie.lch_ToConfiguredGrantMapping_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode lch_ToConfiguredGrantMapping_r16", err)
				}
			}
			// encode lch_ToGrantPriorityRestriction_r16 optional
			if ie.lch_ToGrantPriorityRestriction_r16 != nil {
				if err = ie.lch_ToGrantPriorityRestriction_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode lch_ToGrantPriorityRestriction_r16", err)
				}
			}
			// encode singlePHR_P_r16 optional
			if ie.singlePHR_P_r16 != nil {
				if err = ie.singlePHR_P_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode singlePHR_P_r16", err)
				}
			}
			// encode ul_LBT_FailureDetectionRecovery_r16 optional
			if ie.ul_LBT_FailureDetectionRecovery_r16 != nil {
				if err = ie.ul_LBT_FailureDetectionRecovery_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ul_LBT_FailureDetectionRecovery_r16", err)
				}
			}
			// encode tdd_MPE_P_MPR_Reporting_r16 optional
			if ie.tdd_MPE_P_MPR_Reporting_r16 != nil {
				if err = ie.tdd_MPE_P_MPR_Reporting_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode tdd_MPE_P_MPR_Reporting_r16", err)
				}
			}
			// encode lcid_ExtensionIAB_r16 optional
			if ie.lcid_ExtensionIAB_r16 != nil {
				if err = ie.lcid_ExtensionIAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode lcid_ExtensionIAB_r16", err)
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
			optionals_ext_3 := []bool{ie.spCell_BFR_CBRA_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode spCell_BFR_CBRA_r16 optional
			if ie.spCell_BFR_CBRA_r16 != nil {
				if err = ie.spCell_BFR_CBRA_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode spCell_BFR_CBRA_r16", err)
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
			optionals_ext_4 := []bool{ie.srs_ResourceId_Ext_r16 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode srs_ResourceId_Ext_r16 optional
			if ie.srs_ResourceId_Ext_r16 != nil {
				if err = ie.srs_ResourceId_Ext_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode srs_ResourceId_Ext_r16", err)
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
			optionals_ext_5 := []bool{ie.enhancedUuDRX_forSidelink_r17 != nil, ie.mg_ActivationRequestPRS_Meas_r17 != nil, ie.mg_ActivationCommPRS_Meas_r17 != nil, ie.intraCG_Prioritization_r17 != nil, ie.jointPrioritizationCG_Retx_Timer_r17 != nil, ie.survivalTime_r17 != nil, ie.lcg_ExtensionIAB_r17 != nil, ie.harq_FeedbackDisabled_r17 != nil, ie.uplink_Harq_ModeB_r17 != nil, ie.sr_TriggeredBy_TA_Report_r17 != nil, ie.extendedDRX_CycleInactive_r17 != nil, ie.simultaneousSR_PUSCH_DiffPUCCH_groups_r17 != nil, ie.lastTransmissionUL_r17 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode enhancedUuDRX_forSidelink_r17 optional
			if ie.enhancedUuDRX_forSidelink_r17 != nil {
				if err = ie.enhancedUuDRX_forSidelink_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enhancedUuDRX_forSidelink_r17", err)
				}
			}
			// encode mg_ActivationRequestPRS_Meas_r17 optional
			if ie.mg_ActivationRequestPRS_Meas_r17 != nil {
				if err = ie.mg_ActivationRequestPRS_Meas_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mg_ActivationRequestPRS_Meas_r17", err)
				}
			}
			// encode mg_ActivationCommPRS_Meas_r17 optional
			if ie.mg_ActivationCommPRS_Meas_r17 != nil {
				if err = ie.mg_ActivationCommPRS_Meas_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mg_ActivationCommPRS_Meas_r17", err)
				}
			}
			// encode intraCG_Prioritization_r17 optional
			if ie.intraCG_Prioritization_r17 != nil {
				if err = ie.intraCG_Prioritization_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode intraCG_Prioritization_r17", err)
				}
			}
			// encode jointPrioritizationCG_Retx_Timer_r17 optional
			if ie.jointPrioritizationCG_Retx_Timer_r17 != nil {
				if err = ie.jointPrioritizationCG_Retx_Timer_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode jointPrioritizationCG_Retx_Timer_r17", err)
				}
			}
			// encode survivalTime_r17 optional
			if ie.survivalTime_r17 != nil {
				if err = ie.survivalTime_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode survivalTime_r17", err)
				}
			}
			// encode lcg_ExtensionIAB_r17 optional
			if ie.lcg_ExtensionIAB_r17 != nil {
				if err = ie.lcg_ExtensionIAB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode lcg_ExtensionIAB_r17", err)
				}
			}
			// encode harq_FeedbackDisabled_r17 optional
			if ie.harq_FeedbackDisabled_r17 != nil {
				if err = ie.harq_FeedbackDisabled_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode harq_FeedbackDisabled_r17", err)
				}
			}
			// encode uplink_Harq_ModeB_r17 optional
			if ie.uplink_Harq_ModeB_r17 != nil {
				if err = ie.uplink_Harq_ModeB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uplink_Harq_ModeB_r17", err)
				}
			}
			// encode sr_TriggeredBy_TA_Report_r17 optional
			if ie.sr_TriggeredBy_TA_Report_r17 != nil {
				if err = ie.sr_TriggeredBy_TA_Report_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sr_TriggeredBy_TA_Report_r17", err)
				}
			}
			// encode extendedDRX_CycleInactive_r17 optional
			if ie.extendedDRX_CycleInactive_r17 != nil {
				if err = ie.extendedDRX_CycleInactive_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode extendedDRX_CycleInactive_r17", err)
				}
			}
			// encode simultaneousSR_PUSCH_DiffPUCCH_groups_r17 optional
			if ie.simultaneousSR_PUSCH_DiffPUCCH_groups_r17 != nil {
				if err = ie.simultaneousSR_PUSCH_DiffPUCCH_groups_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simultaneousSR_PUSCH_DiffPUCCH_groups_r17", err)
				}
			}
			// encode lastTransmissionUL_r17 optional
			if ie.lastTransmissionUL_r17 != nil {
				if err = ie.lastTransmissionUL_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode lastTransmissionUL_r17", err)
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

func (ie *MAC_ParametersCommon) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var lcp_RestrictionPresent bool
	if lcp_RestrictionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dummyPresent bool
	if dummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var lch_ToSCellRestrictionPresent bool
	if lch_ToSCellRestrictionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if lcp_RestrictionPresent {
		ie.lcp_Restriction = new(MAC_ParametersCommon_lcp_Restriction)
		if err = ie.lcp_Restriction.Decode(r); err != nil {
			return utils.WrapError("Decode lcp_Restriction", err)
		}
	}
	if dummyPresent {
		ie.dummy = new(MAC_ParametersCommon_dummy)
		if err = ie.dummy.Decode(r); err != nil {
			return utils.WrapError("Decode dummy", err)
		}
	}
	if lch_ToSCellRestrictionPresent {
		ie.lch_ToSCellRestriction = new(MAC_ParametersCommon_lch_ToSCellRestriction)
		if err = ie.lch_ToSCellRestriction.Decode(r); err != nil {
			return utils.WrapError("Decode lch_ToSCellRestriction", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 5 bits for 5 extension groups
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

			recommendedBitRatePresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			recommendedBitRateQueryPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode recommendedBitRate optional
			if recommendedBitRatePresent {
				ie.recommendedBitRate = new(MAC_ParametersCommon_recommendedBitRate)
				if err = ie.recommendedBitRate.Decode(extReader); err != nil {
					return utils.WrapError("Decode recommendedBitRate", err)
				}
			}
			// decode recommendedBitRateQuery optional
			if recommendedBitRateQueryPresent {
				ie.recommendedBitRateQuery = new(MAC_ParametersCommon_recommendedBitRateQuery)
				if err = ie.recommendedBitRateQuery.Decode(extReader); err != nil {
					return utils.WrapError("Decode recommendedBitRateQuery", err)
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

			recommendedBitRateMultiplier_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			preEmptiveBSR_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			autonomousTransmission_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			lch_PriorityBasedPrioritization_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			lch_ToConfiguredGrantMapping_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			lch_ToGrantPriorityRestriction_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			singlePHR_P_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ul_LBT_FailureDetectionRecovery_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			tdd_MPE_P_MPR_Reporting_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			lcid_ExtensionIAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode recommendedBitRateMultiplier_r16 optional
			if recommendedBitRateMultiplier_r16Present {
				ie.recommendedBitRateMultiplier_r16 = new(MAC_ParametersCommon_recommendedBitRateMultiplier_r16)
				if err = ie.recommendedBitRateMultiplier_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode recommendedBitRateMultiplier_r16", err)
				}
			}
			// decode preEmptiveBSR_r16 optional
			if preEmptiveBSR_r16Present {
				ie.preEmptiveBSR_r16 = new(MAC_ParametersCommon_preEmptiveBSR_r16)
				if err = ie.preEmptiveBSR_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode preEmptiveBSR_r16", err)
				}
			}
			// decode autonomousTransmission_r16 optional
			if autonomousTransmission_r16Present {
				ie.autonomousTransmission_r16 = new(MAC_ParametersCommon_autonomousTransmission_r16)
				if err = ie.autonomousTransmission_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode autonomousTransmission_r16", err)
				}
			}
			// decode lch_PriorityBasedPrioritization_r16 optional
			if lch_PriorityBasedPrioritization_r16Present {
				ie.lch_PriorityBasedPrioritization_r16 = new(MAC_ParametersCommon_lch_PriorityBasedPrioritization_r16)
				if err = ie.lch_PriorityBasedPrioritization_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode lch_PriorityBasedPrioritization_r16", err)
				}
			}
			// decode lch_ToConfiguredGrantMapping_r16 optional
			if lch_ToConfiguredGrantMapping_r16Present {
				ie.lch_ToConfiguredGrantMapping_r16 = new(MAC_ParametersCommon_lch_ToConfiguredGrantMapping_r16)
				if err = ie.lch_ToConfiguredGrantMapping_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode lch_ToConfiguredGrantMapping_r16", err)
				}
			}
			// decode lch_ToGrantPriorityRestriction_r16 optional
			if lch_ToGrantPriorityRestriction_r16Present {
				ie.lch_ToGrantPriorityRestriction_r16 = new(MAC_ParametersCommon_lch_ToGrantPriorityRestriction_r16)
				if err = ie.lch_ToGrantPriorityRestriction_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode lch_ToGrantPriorityRestriction_r16", err)
				}
			}
			// decode singlePHR_P_r16 optional
			if singlePHR_P_r16Present {
				ie.singlePHR_P_r16 = new(MAC_ParametersCommon_singlePHR_P_r16)
				if err = ie.singlePHR_P_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode singlePHR_P_r16", err)
				}
			}
			// decode ul_LBT_FailureDetectionRecovery_r16 optional
			if ul_LBT_FailureDetectionRecovery_r16Present {
				ie.ul_LBT_FailureDetectionRecovery_r16 = new(MAC_ParametersCommon_ul_LBT_FailureDetectionRecovery_r16)
				if err = ie.ul_LBT_FailureDetectionRecovery_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ul_LBT_FailureDetectionRecovery_r16", err)
				}
			}
			// decode tdd_MPE_P_MPR_Reporting_r16 optional
			if tdd_MPE_P_MPR_Reporting_r16Present {
				ie.tdd_MPE_P_MPR_Reporting_r16 = new(MAC_ParametersCommon_tdd_MPE_P_MPR_Reporting_r16)
				if err = ie.tdd_MPE_P_MPR_Reporting_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode tdd_MPE_P_MPR_Reporting_r16", err)
				}
			}
			// decode lcid_ExtensionIAB_r16 optional
			if lcid_ExtensionIAB_r16Present {
				ie.lcid_ExtensionIAB_r16 = new(MAC_ParametersCommon_lcid_ExtensionIAB_r16)
				if err = ie.lcid_ExtensionIAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode lcid_ExtensionIAB_r16", err)
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

			spCell_BFR_CBRA_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode spCell_BFR_CBRA_r16 optional
			if spCell_BFR_CBRA_r16Present {
				ie.spCell_BFR_CBRA_r16 = new(MAC_ParametersCommon_spCell_BFR_CBRA_r16)
				if err = ie.spCell_BFR_CBRA_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode spCell_BFR_CBRA_r16", err)
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

			srs_ResourceId_Ext_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode srs_ResourceId_Ext_r16 optional
			if srs_ResourceId_Ext_r16Present {
				ie.srs_ResourceId_Ext_r16 = new(MAC_ParametersCommon_srs_ResourceId_Ext_r16)
				if err = ie.srs_ResourceId_Ext_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode srs_ResourceId_Ext_r16", err)
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

			enhancedUuDRX_forSidelink_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mg_ActivationRequestPRS_Meas_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mg_ActivationCommPRS_Meas_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			intraCG_Prioritization_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			jointPrioritizationCG_Retx_Timer_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			survivalTime_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			lcg_ExtensionIAB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			harq_FeedbackDisabled_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uplink_Harq_ModeB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sr_TriggeredBy_TA_Report_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			extendedDRX_CycleInactive_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			simultaneousSR_PUSCH_DiffPUCCH_groups_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			lastTransmissionUL_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode enhancedUuDRX_forSidelink_r17 optional
			if enhancedUuDRX_forSidelink_r17Present {
				ie.enhancedUuDRX_forSidelink_r17 = new(MAC_ParametersCommon_enhancedUuDRX_forSidelink_r17)
				if err = ie.enhancedUuDRX_forSidelink_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode enhancedUuDRX_forSidelink_r17", err)
				}
			}
			// decode mg_ActivationRequestPRS_Meas_r17 optional
			if mg_ActivationRequestPRS_Meas_r17Present {
				ie.mg_ActivationRequestPRS_Meas_r17 = new(MAC_ParametersCommon_mg_ActivationRequestPRS_Meas_r17)
				if err = ie.mg_ActivationRequestPRS_Meas_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mg_ActivationRequestPRS_Meas_r17", err)
				}
			}
			// decode mg_ActivationCommPRS_Meas_r17 optional
			if mg_ActivationCommPRS_Meas_r17Present {
				ie.mg_ActivationCommPRS_Meas_r17 = new(MAC_ParametersCommon_mg_ActivationCommPRS_Meas_r17)
				if err = ie.mg_ActivationCommPRS_Meas_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode mg_ActivationCommPRS_Meas_r17", err)
				}
			}
			// decode intraCG_Prioritization_r17 optional
			if intraCG_Prioritization_r17Present {
				ie.intraCG_Prioritization_r17 = new(MAC_ParametersCommon_intraCG_Prioritization_r17)
				if err = ie.intraCG_Prioritization_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode intraCG_Prioritization_r17", err)
				}
			}
			// decode jointPrioritizationCG_Retx_Timer_r17 optional
			if jointPrioritizationCG_Retx_Timer_r17Present {
				ie.jointPrioritizationCG_Retx_Timer_r17 = new(MAC_ParametersCommon_jointPrioritizationCG_Retx_Timer_r17)
				if err = ie.jointPrioritizationCG_Retx_Timer_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode jointPrioritizationCG_Retx_Timer_r17", err)
				}
			}
			// decode survivalTime_r17 optional
			if survivalTime_r17Present {
				ie.survivalTime_r17 = new(MAC_ParametersCommon_survivalTime_r17)
				if err = ie.survivalTime_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode survivalTime_r17", err)
				}
			}
			// decode lcg_ExtensionIAB_r17 optional
			if lcg_ExtensionIAB_r17Present {
				ie.lcg_ExtensionIAB_r17 = new(MAC_ParametersCommon_lcg_ExtensionIAB_r17)
				if err = ie.lcg_ExtensionIAB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode lcg_ExtensionIAB_r17", err)
				}
			}
			// decode harq_FeedbackDisabled_r17 optional
			if harq_FeedbackDisabled_r17Present {
				ie.harq_FeedbackDisabled_r17 = new(MAC_ParametersCommon_harq_FeedbackDisabled_r17)
				if err = ie.harq_FeedbackDisabled_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode harq_FeedbackDisabled_r17", err)
				}
			}
			// decode uplink_Harq_ModeB_r17 optional
			if uplink_Harq_ModeB_r17Present {
				ie.uplink_Harq_ModeB_r17 = new(MAC_ParametersCommon_uplink_Harq_ModeB_r17)
				if err = ie.uplink_Harq_ModeB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode uplink_Harq_ModeB_r17", err)
				}
			}
			// decode sr_TriggeredBy_TA_Report_r17 optional
			if sr_TriggeredBy_TA_Report_r17Present {
				ie.sr_TriggeredBy_TA_Report_r17 = new(MAC_ParametersCommon_sr_TriggeredBy_TA_Report_r17)
				if err = ie.sr_TriggeredBy_TA_Report_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sr_TriggeredBy_TA_Report_r17", err)
				}
			}
			// decode extendedDRX_CycleInactive_r17 optional
			if extendedDRX_CycleInactive_r17Present {
				ie.extendedDRX_CycleInactive_r17 = new(MAC_ParametersCommon_extendedDRX_CycleInactive_r17)
				if err = ie.extendedDRX_CycleInactive_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode extendedDRX_CycleInactive_r17", err)
				}
			}
			// decode simultaneousSR_PUSCH_DiffPUCCH_groups_r17 optional
			if simultaneousSR_PUSCH_DiffPUCCH_groups_r17Present {
				ie.simultaneousSR_PUSCH_DiffPUCCH_groups_r17 = new(MAC_ParametersCommon_simultaneousSR_PUSCH_DiffPUCCH_groups_r17)
				if err = ie.simultaneousSR_PUSCH_DiffPUCCH_groups_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode simultaneousSR_PUSCH_DiffPUCCH_groups_r17", err)
				}
			}
			// decode lastTransmissionUL_r17 optional
			if lastTransmissionUL_r17Present {
				ie.lastTransmissionUL_r17 = new(MAC_ParametersCommon_lastTransmissionUL_r17)
				if err = ie.lastTransmissionUL_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode lastTransmissionUL_r17", err)
				}
			}
		}
	}
	return nil
}
