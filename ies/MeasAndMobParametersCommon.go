package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersCommon struct {
	supportedGapPattern                          *uper.BitString                                                          `lb:22,ub:22,optional`
	ssb_RLM                                      *MeasAndMobParametersCommon_ssb_RLM                                      `optional`
	ssb_AndCSI_RS_RLM                            *MeasAndMobParametersCommon_ssb_AndCSI_RS_RLM                            `optional`
	eventB_MeasAndReport                         *MeasAndMobParametersCommon_eventB_MeasAndReport                         `optional,ext-1`
	handoverFDD_TDD                              *MeasAndMobParametersCommon_handoverFDD_TDD                              `optional,ext-1`
	eutra_CGI_Reporting                          *MeasAndMobParametersCommon_eutra_CGI_Reporting                          `optional,ext-1`
	nr_CGI_Reporting                             *MeasAndMobParametersCommon_nr_CGI_Reporting                             `optional,ext-1`
	independentGapConfig                         *MeasAndMobParametersCommon_independentGapConfig                         `optional,ext-2`
	periodicEUTRA_MeasAndReport                  *MeasAndMobParametersCommon_periodicEUTRA_MeasAndReport                  `optional,ext-2`
	handoverFR1_FR2                              *MeasAndMobParametersCommon_handoverFR1_FR2                              `optional,ext-2`
	maxNumberCSI_RS_RRM_RS_SINR                  *MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR                  `optional,ext-2`
	nr_CGI_Reporting_ENDC                        *MeasAndMobParametersCommon_nr_CGI_Reporting_ENDC                        `optional,ext-3`
	eutra_CGI_Reporting_NEDC                     *MeasAndMobParametersCommon_eutra_CGI_Reporting_NEDC                     `optional,ext-4`
	eutra_CGI_Reporting_NRDC                     *MeasAndMobParametersCommon_eutra_CGI_Reporting_NRDC                     `optional,ext-4`
	nr_CGI_Reporting_NEDC                        *MeasAndMobParametersCommon_nr_CGI_Reporting_NEDC                        `optional,ext-4`
	nr_CGI_Reporting_NRDC                        *MeasAndMobParametersCommon_nr_CGI_Reporting_NRDC                        `optional,ext-4`
	reportAddNeighMeasForPeriodic_r16            *MeasAndMobParametersCommon_reportAddNeighMeasForPeriodic_r16            `optional,ext-5`
	condHandoverParametersCommon_r16             *MeasAndMobParametersCommon_condHandoverParametersCommon_r16             `optional,ext-5`
	nr_NeedForGap_Reporting_r16                  *MeasAndMobParametersCommon_nr_NeedForGap_Reporting_r16                  `optional,ext-5`
	supportedGapPattern_NRonly_r16               *uper.BitString                                                          `lb:10,ub:10,optional,ext-5`
	supportedGapPattern_NRonly_NEDC_r16          *MeasAndMobParametersCommon_supportedGapPattern_NRonly_NEDC_r16          `optional,ext-5`
	maxNumberCLI_RSSI_r16                        *MeasAndMobParametersCommon_maxNumberCLI_RSSI_r16                        `optional,ext-5`
	maxNumberCLI_SRS_RSRP_r16                    *MeasAndMobParametersCommon_maxNumberCLI_SRS_RSRP_r16                    `optional,ext-5`
	maxNumberPerSlotCLI_SRS_RSRP_r16             *MeasAndMobParametersCommon_maxNumberPerSlotCLI_SRS_RSRP_r16             `optional,ext-5`
	mfbi_IAB_r16                                 *MeasAndMobParametersCommon_mfbi_IAB_r16                                 `optional,ext-5`
	dummy                                        *MeasAndMobParametersCommon_dummy                                        `optional,ext-5`
	nr_CGI_Reporting_NPN_r16                     *MeasAndMobParametersCommon_nr_CGI_Reporting_NPN_r16                     `optional,ext-5`
	idleInactiveEUTRA_MeasReport_r16             *MeasAndMobParametersCommon_idleInactiveEUTRA_MeasReport_r16             `optional,ext-5`
	idleInactive_ValidityArea_r16                *MeasAndMobParametersCommon_idleInactive_ValidityArea_r16                `optional,ext-5`
	eutra_AutonomousGaps_r16                     *MeasAndMobParametersCommon_eutra_AutonomousGaps_r16                     `optional,ext-5`
	eutra_AutonomousGaps_NEDC_r16                *MeasAndMobParametersCommon_eutra_AutonomousGaps_NEDC_r16                `optional,ext-5`
	eutra_AutonomousGaps_NRDC_r16                *MeasAndMobParametersCommon_eutra_AutonomousGaps_NRDC_r16                `optional,ext-5`
	pcellT312_r16                                *MeasAndMobParametersCommon_pcellT312_r16                                `optional,ext-5`
	supportedGapPattern_r16                      *uper.BitString                                                          `lb:2,ub:2,optional,ext-5`
	concurrentMeasGap_r17                        *MeasAndMobParametersCommon_concurrentMeasGap_r17                        `optional,ext-6`
	nr_NeedForGapNCSG_Reporting_r17              *MeasAndMobParametersCommon_nr_NeedForGapNCSG_Reporting_r17              `optional,ext-6`
	eutra_NeedForGapNCSG_Reporting_r17           *MeasAndMobParametersCommon_eutra_NeedForGapNCSG_Reporting_r17           `optional,ext-6`
	ncsg_MeasGapPerFR_r17                        *MeasAndMobParametersCommon_ncsg_MeasGapPerFR_r17                        `optional,ext-6`
	ncsg_MeasGapPatterns_r17                     *uper.BitString                                                          `lb:24,ub:24,optional,ext-6`
	ncsg_MeasGapNR_Patterns_r17                  *uper.BitString                                                          `lb:24,ub:24,optional,ext-6`
	preconfiguredUE_AutonomousMeasGap_r17        *MeasAndMobParametersCommon_preconfiguredUE_AutonomousMeasGap_r17        `optional,ext-6`
	preconfiguredNW_ControlledMeasGap_r17        *MeasAndMobParametersCommon_preconfiguredNW_ControlledMeasGap_r17        `optional,ext-6`
	handoverFR1_FR2_2_r17                        *MeasAndMobParametersCommon_handoverFR1_FR2_2_r17                        `optional,ext-6`
	handoverFR2_1_FR2_2_r17                      *MeasAndMobParametersCommon_handoverFR2_1_FR2_2_r17                      `optional,ext-6`
	independentGapConfigPRS_r17                  *MeasAndMobParametersCommon_independentGapConfigPRS_r17                  `optional,ext-6`
	rrm_RelaxationRRC_ConnectedRedCap_r17        *MeasAndMobParametersCommon_rrm_RelaxationRRC_ConnectedRedCap_r17        `optional,ext-6`
	parallelMeasurementGap_r17                   *MeasAndMobParametersCommon_parallelMeasurementGap_r17                   `optional,ext-6`
	condHandoverWithSCG_NRDC_r17                 *MeasAndMobParametersCommon_condHandoverWithSCG_NRDC_r17                 `optional,ext-6`
	gNB_ID_LengthReporting_r17                   *MeasAndMobParametersCommon_gNB_ID_LengthReporting_r17                   `optional,ext-6`
	gNB_ID_LengthReporting_ENDC_r17              *MeasAndMobParametersCommon_gNB_ID_LengthReporting_ENDC_r17              `optional,ext-6`
	gNB_ID_LengthReporting_NEDC_r17              *MeasAndMobParametersCommon_gNB_ID_LengthReporting_NEDC_r17              `optional,ext-6`
	gNB_ID_LengthReporting_NRDC_r17              *MeasAndMobParametersCommon_gNB_ID_LengthReporting_NRDC_r17              `optional,ext-6`
	gNB_ID_LengthReporting_NPN_r17               *MeasAndMobParametersCommon_gNB_ID_LengthReporting_NPN_r17               `optional,ext-6`
	parallelSMTC_r17                             *MeasAndMobParametersCommon_parallelSMTC_r17                             `optional,ext-7`
	concurrentMeasGapEUTRA_r17                   *MeasAndMobParametersCommon_concurrentMeasGapEUTRA_r17                   `optional,ext-7`
	serviceLinkPropDelayDiffReporting_r17        *MeasAndMobParametersCommon_serviceLinkPropDelayDiffReporting_r17        `optional,ext-7`
	ncsg_SymbolLevelScheduleRestrictionInter_r17 *MeasAndMobParametersCommon_ncsg_SymbolLevelScheduleRestrictionInter_r17 `optional,ext-7`
	eventD1_MeasReportTrigger_r17                *MeasAndMobParametersCommon_eventD1_MeasReportTrigger_r17                `optional,ext-8`
	independentGapConfig_maxCC_r17               *MeasAndMobParametersCommon_independentGapConfig_maxCC_r17               `lb:1,ub:32,optional,ext-8`
}

func (ie *MeasAndMobParametersCommon) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.eventB_MeasAndReport != nil || ie.handoverFDD_TDD != nil || ie.eutra_CGI_Reporting != nil || ie.nr_CGI_Reporting != nil || ie.independentGapConfig != nil || ie.periodicEUTRA_MeasAndReport != nil || ie.handoverFR1_FR2 != nil || ie.maxNumberCSI_RS_RRM_RS_SINR != nil || ie.nr_CGI_Reporting_ENDC != nil || ie.eutra_CGI_Reporting_NEDC != nil || ie.eutra_CGI_Reporting_NRDC != nil || ie.nr_CGI_Reporting_NEDC != nil || ie.nr_CGI_Reporting_NRDC != nil || ie.reportAddNeighMeasForPeriodic_r16 != nil || ie.condHandoverParametersCommon_r16 != nil || ie.nr_NeedForGap_Reporting_r16 != nil || ie.supportedGapPattern_NRonly_r16 != nil || ie.supportedGapPattern_NRonly_NEDC_r16 != nil || ie.maxNumberCLI_RSSI_r16 != nil || ie.maxNumberCLI_SRS_RSRP_r16 != nil || ie.maxNumberPerSlotCLI_SRS_RSRP_r16 != nil || ie.mfbi_IAB_r16 != nil || ie.dummy != nil || ie.nr_CGI_Reporting_NPN_r16 != nil || ie.idleInactiveEUTRA_MeasReport_r16 != nil || ie.idleInactive_ValidityArea_r16 != nil || ie.eutra_AutonomousGaps_r16 != nil || ie.eutra_AutonomousGaps_NEDC_r16 != nil || ie.eutra_AutonomousGaps_NRDC_r16 != nil || ie.pcellT312_r16 != nil || ie.supportedGapPattern_r16 != nil || ie.concurrentMeasGap_r17 != nil || ie.nr_NeedForGapNCSG_Reporting_r17 != nil || ie.eutra_NeedForGapNCSG_Reporting_r17 != nil || ie.ncsg_MeasGapPerFR_r17 != nil || ie.ncsg_MeasGapPatterns_r17 != nil || ie.ncsg_MeasGapNR_Patterns_r17 != nil || ie.preconfiguredUE_AutonomousMeasGap_r17 != nil || ie.preconfiguredNW_ControlledMeasGap_r17 != nil || ie.handoverFR1_FR2_2_r17 != nil || ie.handoverFR2_1_FR2_2_r17 != nil || ie.independentGapConfigPRS_r17 != nil || ie.rrm_RelaxationRRC_ConnectedRedCap_r17 != nil || ie.parallelMeasurementGap_r17 != nil || ie.condHandoverWithSCG_NRDC_r17 != nil || ie.gNB_ID_LengthReporting_r17 != nil || ie.gNB_ID_LengthReporting_ENDC_r17 != nil || ie.gNB_ID_LengthReporting_NEDC_r17 != nil || ie.gNB_ID_LengthReporting_NRDC_r17 != nil || ie.gNB_ID_LengthReporting_NPN_r17 != nil || ie.parallelSMTC_r17 != nil || ie.concurrentMeasGapEUTRA_r17 != nil || ie.serviceLinkPropDelayDiffReporting_r17 != nil || ie.ncsg_SymbolLevelScheduleRestrictionInter_r17 != nil || ie.eventD1_MeasReportTrigger_r17 != nil || ie.independentGapConfig_maxCC_r17 != nil
	preambleBits := []bool{hasExtensions, ie.supportedGapPattern != nil, ie.ssb_RLM != nil, ie.ssb_AndCSI_RS_RLM != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.supportedGapPattern != nil {
		if err = w.WriteBitString(ie.supportedGapPattern.Bytes, uint(ie.supportedGapPattern.NumBits), &uper.Constraint{Lb: 22, Ub: 22}, false); err != nil {
			return utils.WrapError("Encode supportedGapPattern", err)
		}
	}
	if ie.ssb_RLM != nil {
		if err = ie.ssb_RLM.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_RLM", err)
		}
	}
	if ie.ssb_AndCSI_RS_RLM != nil {
		if err = ie.ssb_AndCSI_RS_RLM.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_AndCSI_RS_RLM", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 8 bits for 8 extension groups
		extBitmap := []bool{ie.eventB_MeasAndReport != nil || ie.handoverFDD_TDD != nil || ie.eutra_CGI_Reporting != nil || ie.nr_CGI_Reporting != nil, ie.independentGapConfig != nil || ie.periodicEUTRA_MeasAndReport != nil || ie.handoverFR1_FR2 != nil || ie.maxNumberCSI_RS_RRM_RS_SINR != nil, ie.nr_CGI_Reporting_ENDC != nil, ie.eutra_CGI_Reporting_NEDC != nil || ie.eutra_CGI_Reporting_NRDC != nil || ie.nr_CGI_Reporting_NEDC != nil || ie.nr_CGI_Reporting_NRDC != nil, ie.reportAddNeighMeasForPeriodic_r16 != nil || ie.condHandoverParametersCommon_r16 != nil || ie.nr_NeedForGap_Reporting_r16 != nil || ie.supportedGapPattern_NRonly_r16 != nil || ie.supportedGapPattern_NRonly_NEDC_r16 != nil || ie.maxNumberCLI_RSSI_r16 != nil || ie.maxNumberCLI_SRS_RSRP_r16 != nil || ie.maxNumberPerSlotCLI_SRS_RSRP_r16 != nil || ie.mfbi_IAB_r16 != nil || ie.dummy != nil || ie.nr_CGI_Reporting_NPN_r16 != nil || ie.idleInactiveEUTRA_MeasReport_r16 != nil || ie.idleInactive_ValidityArea_r16 != nil || ie.eutra_AutonomousGaps_r16 != nil || ie.eutra_AutonomousGaps_NEDC_r16 != nil || ie.eutra_AutonomousGaps_NRDC_r16 != nil || ie.pcellT312_r16 != nil || ie.supportedGapPattern_r16 != nil, ie.concurrentMeasGap_r17 != nil || ie.nr_NeedForGapNCSG_Reporting_r17 != nil || ie.eutra_NeedForGapNCSG_Reporting_r17 != nil || ie.ncsg_MeasGapPerFR_r17 != nil || ie.ncsg_MeasGapPatterns_r17 != nil || ie.ncsg_MeasGapNR_Patterns_r17 != nil || ie.preconfiguredUE_AutonomousMeasGap_r17 != nil || ie.preconfiguredNW_ControlledMeasGap_r17 != nil || ie.handoverFR1_FR2_2_r17 != nil || ie.handoverFR2_1_FR2_2_r17 != nil || ie.independentGapConfigPRS_r17 != nil || ie.rrm_RelaxationRRC_ConnectedRedCap_r17 != nil || ie.parallelMeasurementGap_r17 != nil || ie.condHandoverWithSCG_NRDC_r17 != nil || ie.gNB_ID_LengthReporting_r17 != nil || ie.gNB_ID_LengthReporting_ENDC_r17 != nil || ie.gNB_ID_LengthReporting_NEDC_r17 != nil || ie.gNB_ID_LengthReporting_NRDC_r17 != nil || ie.gNB_ID_LengthReporting_NPN_r17 != nil, ie.parallelSMTC_r17 != nil || ie.concurrentMeasGapEUTRA_r17 != nil || ie.serviceLinkPropDelayDiffReporting_r17 != nil || ie.ncsg_SymbolLevelScheduleRestrictionInter_r17 != nil, ie.eventD1_MeasReportTrigger_r17 != nil || ie.independentGapConfig_maxCC_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasAndMobParametersCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.eventB_MeasAndReport != nil, ie.handoverFDD_TDD != nil, ie.eutra_CGI_Reporting != nil, ie.nr_CGI_Reporting != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode eventB_MeasAndReport optional
			if ie.eventB_MeasAndReport != nil {
				if err = ie.eventB_MeasAndReport.Encode(extWriter); err != nil {
					return utils.WrapError("Encode eventB_MeasAndReport", err)
				}
			}
			// encode handoverFDD_TDD optional
			if ie.handoverFDD_TDD != nil {
				if err = ie.handoverFDD_TDD.Encode(extWriter); err != nil {
					return utils.WrapError("Encode handoverFDD_TDD", err)
				}
			}
			// encode eutra_CGI_Reporting optional
			if ie.eutra_CGI_Reporting != nil {
				if err = ie.eutra_CGI_Reporting.Encode(extWriter); err != nil {
					return utils.WrapError("Encode eutra_CGI_Reporting", err)
				}
			}
			// encode nr_CGI_Reporting optional
			if ie.nr_CGI_Reporting != nil {
				if err = ie.nr_CGI_Reporting.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nr_CGI_Reporting", err)
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
			optionals_ext_2 := []bool{ie.independentGapConfig != nil, ie.periodicEUTRA_MeasAndReport != nil, ie.handoverFR1_FR2 != nil, ie.maxNumberCSI_RS_RRM_RS_SINR != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode independentGapConfig optional
			if ie.independentGapConfig != nil {
				if err = ie.independentGapConfig.Encode(extWriter); err != nil {
					return utils.WrapError("Encode independentGapConfig", err)
				}
			}
			// encode periodicEUTRA_MeasAndReport optional
			if ie.periodicEUTRA_MeasAndReport != nil {
				if err = ie.periodicEUTRA_MeasAndReport.Encode(extWriter); err != nil {
					return utils.WrapError("Encode periodicEUTRA_MeasAndReport", err)
				}
			}
			// encode handoverFR1_FR2 optional
			if ie.handoverFR1_FR2 != nil {
				if err = ie.handoverFR1_FR2.Encode(extWriter); err != nil {
					return utils.WrapError("Encode handoverFR1_FR2", err)
				}
			}
			// encode maxNumberCSI_RS_RRM_RS_SINR optional
			if ie.maxNumberCSI_RS_RRM_RS_SINR != nil {
				if err = ie.maxNumberCSI_RS_RRM_RS_SINR.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxNumberCSI_RS_RRM_RS_SINR", err)
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
			optionals_ext_3 := []bool{ie.nr_CGI_Reporting_ENDC != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode nr_CGI_Reporting_ENDC optional
			if ie.nr_CGI_Reporting_ENDC != nil {
				if err = ie.nr_CGI_Reporting_ENDC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nr_CGI_Reporting_ENDC", err)
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
			optionals_ext_4 := []bool{ie.eutra_CGI_Reporting_NEDC != nil, ie.eutra_CGI_Reporting_NRDC != nil, ie.nr_CGI_Reporting_NEDC != nil, ie.nr_CGI_Reporting_NRDC != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode eutra_CGI_Reporting_NEDC optional
			if ie.eutra_CGI_Reporting_NEDC != nil {
				if err = ie.eutra_CGI_Reporting_NEDC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode eutra_CGI_Reporting_NEDC", err)
				}
			}
			// encode eutra_CGI_Reporting_NRDC optional
			if ie.eutra_CGI_Reporting_NRDC != nil {
				if err = ie.eutra_CGI_Reporting_NRDC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode eutra_CGI_Reporting_NRDC", err)
				}
			}
			// encode nr_CGI_Reporting_NEDC optional
			if ie.nr_CGI_Reporting_NEDC != nil {
				if err = ie.nr_CGI_Reporting_NEDC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nr_CGI_Reporting_NEDC", err)
				}
			}
			// encode nr_CGI_Reporting_NRDC optional
			if ie.nr_CGI_Reporting_NRDC != nil {
				if err = ie.nr_CGI_Reporting_NRDC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nr_CGI_Reporting_NRDC", err)
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
			optionals_ext_5 := []bool{ie.reportAddNeighMeasForPeriodic_r16 != nil, ie.condHandoverParametersCommon_r16 != nil, ie.nr_NeedForGap_Reporting_r16 != nil, ie.supportedGapPattern_NRonly_r16 != nil, ie.supportedGapPattern_NRonly_NEDC_r16 != nil, ie.maxNumberCLI_RSSI_r16 != nil, ie.maxNumberCLI_SRS_RSRP_r16 != nil, ie.maxNumberPerSlotCLI_SRS_RSRP_r16 != nil, ie.mfbi_IAB_r16 != nil, ie.dummy != nil, ie.nr_CGI_Reporting_NPN_r16 != nil, ie.idleInactiveEUTRA_MeasReport_r16 != nil, ie.idleInactive_ValidityArea_r16 != nil, ie.eutra_AutonomousGaps_r16 != nil, ie.eutra_AutonomousGaps_NEDC_r16 != nil, ie.eutra_AutonomousGaps_NRDC_r16 != nil, ie.pcellT312_r16 != nil, ie.supportedGapPattern_r16 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode reportAddNeighMeasForPeriodic_r16 optional
			if ie.reportAddNeighMeasForPeriodic_r16 != nil {
				if err = ie.reportAddNeighMeasForPeriodic_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode reportAddNeighMeasForPeriodic_r16", err)
				}
			}
			// encode condHandoverParametersCommon_r16 optional
			if ie.condHandoverParametersCommon_r16 != nil {
				if err = ie.condHandoverParametersCommon_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode condHandoverParametersCommon_r16", err)
				}
			}
			// encode nr_NeedForGap_Reporting_r16 optional
			if ie.nr_NeedForGap_Reporting_r16 != nil {
				if err = ie.nr_NeedForGap_Reporting_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nr_NeedForGap_Reporting_r16", err)
				}
			}
			// encode supportedGapPattern_NRonly_r16 optional
			if ie.supportedGapPattern_NRonly_r16 != nil {
				if err = extWriter.WriteBitString(ie.supportedGapPattern_NRonly_r16.Bytes, uint(ie.supportedGapPattern_NRonly_r16.NumBits), &uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
					return utils.WrapError("Encode supportedGapPattern_NRonly_r16", err)
				}
			}
			// encode supportedGapPattern_NRonly_NEDC_r16 optional
			if ie.supportedGapPattern_NRonly_NEDC_r16 != nil {
				if err = ie.supportedGapPattern_NRonly_NEDC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode supportedGapPattern_NRonly_NEDC_r16", err)
				}
			}
			// encode maxNumberCLI_RSSI_r16 optional
			if ie.maxNumberCLI_RSSI_r16 != nil {
				if err = ie.maxNumberCLI_RSSI_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxNumberCLI_RSSI_r16", err)
				}
			}
			// encode maxNumberCLI_SRS_RSRP_r16 optional
			if ie.maxNumberCLI_SRS_RSRP_r16 != nil {
				if err = ie.maxNumberCLI_SRS_RSRP_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxNumberCLI_SRS_RSRP_r16", err)
				}
			}
			// encode maxNumberPerSlotCLI_SRS_RSRP_r16 optional
			if ie.maxNumberPerSlotCLI_SRS_RSRP_r16 != nil {
				if err = ie.maxNumberPerSlotCLI_SRS_RSRP_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxNumberPerSlotCLI_SRS_RSRP_r16", err)
				}
			}
			// encode mfbi_IAB_r16 optional
			if ie.mfbi_IAB_r16 != nil {
				if err = ie.mfbi_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode mfbi_IAB_r16", err)
				}
			}
			// encode dummy optional
			if ie.dummy != nil {
				if err = ie.dummy.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dummy", err)
				}
			}
			// encode nr_CGI_Reporting_NPN_r16 optional
			if ie.nr_CGI_Reporting_NPN_r16 != nil {
				if err = ie.nr_CGI_Reporting_NPN_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nr_CGI_Reporting_NPN_r16", err)
				}
			}
			// encode idleInactiveEUTRA_MeasReport_r16 optional
			if ie.idleInactiveEUTRA_MeasReport_r16 != nil {
				if err = ie.idleInactiveEUTRA_MeasReport_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode idleInactiveEUTRA_MeasReport_r16", err)
				}
			}
			// encode idleInactive_ValidityArea_r16 optional
			if ie.idleInactive_ValidityArea_r16 != nil {
				if err = ie.idleInactive_ValidityArea_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode idleInactive_ValidityArea_r16", err)
				}
			}
			// encode eutra_AutonomousGaps_r16 optional
			if ie.eutra_AutonomousGaps_r16 != nil {
				if err = ie.eutra_AutonomousGaps_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode eutra_AutonomousGaps_r16", err)
				}
			}
			// encode eutra_AutonomousGaps_NEDC_r16 optional
			if ie.eutra_AutonomousGaps_NEDC_r16 != nil {
				if err = ie.eutra_AutonomousGaps_NEDC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode eutra_AutonomousGaps_NEDC_r16", err)
				}
			}
			// encode eutra_AutonomousGaps_NRDC_r16 optional
			if ie.eutra_AutonomousGaps_NRDC_r16 != nil {
				if err = ie.eutra_AutonomousGaps_NRDC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode eutra_AutonomousGaps_NRDC_r16", err)
				}
			}
			// encode pcellT312_r16 optional
			if ie.pcellT312_r16 != nil {
				if err = ie.pcellT312_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pcellT312_r16", err)
				}
			}
			// encode supportedGapPattern_r16 optional
			if ie.supportedGapPattern_r16 != nil {
				if err = extWriter.WriteBitString(ie.supportedGapPattern_r16.Bytes, uint(ie.supportedGapPattern_r16.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
					return utils.WrapError("Encode supportedGapPattern_r16", err)
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
			optionals_ext_6 := []bool{ie.concurrentMeasGap_r17 != nil, ie.nr_NeedForGapNCSG_Reporting_r17 != nil, ie.eutra_NeedForGapNCSG_Reporting_r17 != nil, ie.ncsg_MeasGapPerFR_r17 != nil, ie.ncsg_MeasGapPatterns_r17 != nil, ie.ncsg_MeasGapNR_Patterns_r17 != nil, ie.preconfiguredUE_AutonomousMeasGap_r17 != nil, ie.preconfiguredNW_ControlledMeasGap_r17 != nil, ie.handoverFR1_FR2_2_r17 != nil, ie.handoverFR2_1_FR2_2_r17 != nil, ie.independentGapConfigPRS_r17 != nil, ie.rrm_RelaxationRRC_ConnectedRedCap_r17 != nil, ie.parallelMeasurementGap_r17 != nil, ie.condHandoverWithSCG_NRDC_r17 != nil, ie.gNB_ID_LengthReporting_r17 != nil, ie.gNB_ID_LengthReporting_ENDC_r17 != nil, ie.gNB_ID_LengthReporting_NEDC_r17 != nil, ie.gNB_ID_LengthReporting_NRDC_r17 != nil, ie.gNB_ID_LengthReporting_NPN_r17 != nil}
			for _, bit := range optionals_ext_6 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode concurrentMeasGap_r17 optional
			if ie.concurrentMeasGap_r17 != nil {
				if err = ie.concurrentMeasGap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode concurrentMeasGap_r17", err)
				}
			}
			// encode nr_NeedForGapNCSG_Reporting_r17 optional
			if ie.nr_NeedForGapNCSG_Reporting_r17 != nil {
				if err = ie.nr_NeedForGapNCSG_Reporting_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nr_NeedForGapNCSG_Reporting_r17", err)
				}
			}
			// encode eutra_NeedForGapNCSG_Reporting_r17 optional
			if ie.eutra_NeedForGapNCSG_Reporting_r17 != nil {
				if err = ie.eutra_NeedForGapNCSG_Reporting_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode eutra_NeedForGapNCSG_Reporting_r17", err)
				}
			}
			// encode ncsg_MeasGapPerFR_r17 optional
			if ie.ncsg_MeasGapPerFR_r17 != nil {
				if err = ie.ncsg_MeasGapPerFR_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ncsg_MeasGapPerFR_r17", err)
				}
			}
			// encode ncsg_MeasGapPatterns_r17 optional
			if ie.ncsg_MeasGapPatterns_r17 != nil {
				if err = extWriter.WriteBitString(ie.ncsg_MeasGapPatterns_r17.Bytes, uint(ie.ncsg_MeasGapPatterns_r17.NumBits), &uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
					return utils.WrapError("Encode ncsg_MeasGapPatterns_r17", err)
				}
			}
			// encode ncsg_MeasGapNR_Patterns_r17 optional
			if ie.ncsg_MeasGapNR_Patterns_r17 != nil {
				if err = extWriter.WriteBitString(ie.ncsg_MeasGapNR_Patterns_r17.Bytes, uint(ie.ncsg_MeasGapNR_Patterns_r17.NumBits), &uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
					return utils.WrapError("Encode ncsg_MeasGapNR_Patterns_r17", err)
				}
			}
			// encode preconfiguredUE_AutonomousMeasGap_r17 optional
			if ie.preconfiguredUE_AutonomousMeasGap_r17 != nil {
				if err = ie.preconfiguredUE_AutonomousMeasGap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode preconfiguredUE_AutonomousMeasGap_r17", err)
				}
			}
			// encode preconfiguredNW_ControlledMeasGap_r17 optional
			if ie.preconfiguredNW_ControlledMeasGap_r17 != nil {
				if err = ie.preconfiguredNW_ControlledMeasGap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode preconfiguredNW_ControlledMeasGap_r17", err)
				}
			}
			// encode handoverFR1_FR2_2_r17 optional
			if ie.handoverFR1_FR2_2_r17 != nil {
				if err = ie.handoverFR1_FR2_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode handoverFR1_FR2_2_r17", err)
				}
			}
			// encode handoverFR2_1_FR2_2_r17 optional
			if ie.handoverFR2_1_FR2_2_r17 != nil {
				if err = ie.handoverFR2_1_FR2_2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode handoverFR2_1_FR2_2_r17", err)
				}
			}
			// encode independentGapConfigPRS_r17 optional
			if ie.independentGapConfigPRS_r17 != nil {
				if err = ie.independentGapConfigPRS_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode independentGapConfigPRS_r17", err)
				}
			}
			// encode rrm_RelaxationRRC_ConnectedRedCap_r17 optional
			if ie.rrm_RelaxationRRC_ConnectedRedCap_r17 != nil {
				if err = ie.rrm_RelaxationRRC_ConnectedRedCap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rrm_RelaxationRRC_ConnectedRedCap_r17", err)
				}
			}
			// encode parallelMeasurementGap_r17 optional
			if ie.parallelMeasurementGap_r17 != nil {
				if err = ie.parallelMeasurementGap_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode parallelMeasurementGap_r17", err)
				}
			}
			// encode condHandoverWithSCG_NRDC_r17 optional
			if ie.condHandoverWithSCG_NRDC_r17 != nil {
				if err = ie.condHandoverWithSCG_NRDC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode condHandoverWithSCG_NRDC_r17", err)
				}
			}
			// encode gNB_ID_LengthReporting_r17 optional
			if ie.gNB_ID_LengthReporting_r17 != nil {
				if err = ie.gNB_ID_LengthReporting_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode gNB_ID_LengthReporting_r17", err)
				}
			}
			// encode gNB_ID_LengthReporting_ENDC_r17 optional
			if ie.gNB_ID_LengthReporting_ENDC_r17 != nil {
				if err = ie.gNB_ID_LengthReporting_ENDC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode gNB_ID_LengthReporting_ENDC_r17", err)
				}
			}
			// encode gNB_ID_LengthReporting_NEDC_r17 optional
			if ie.gNB_ID_LengthReporting_NEDC_r17 != nil {
				if err = ie.gNB_ID_LengthReporting_NEDC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode gNB_ID_LengthReporting_NEDC_r17", err)
				}
			}
			// encode gNB_ID_LengthReporting_NRDC_r17 optional
			if ie.gNB_ID_LengthReporting_NRDC_r17 != nil {
				if err = ie.gNB_ID_LengthReporting_NRDC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode gNB_ID_LengthReporting_NRDC_r17", err)
				}
			}
			// encode gNB_ID_LengthReporting_NPN_r17 optional
			if ie.gNB_ID_LengthReporting_NPN_r17 != nil {
				if err = ie.gNB_ID_LengthReporting_NPN_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode gNB_ID_LengthReporting_NPN_r17", err)
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
			optionals_ext_7 := []bool{ie.parallelSMTC_r17 != nil, ie.concurrentMeasGapEUTRA_r17 != nil, ie.serviceLinkPropDelayDiffReporting_r17 != nil, ie.ncsg_SymbolLevelScheduleRestrictionInter_r17 != nil}
			for _, bit := range optionals_ext_7 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode parallelSMTC_r17 optional
			if ie.parallelSMTC_r17 != nil {
				if err = ie.parallelSMTC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode parallelSMTC_r17", err)
				}
			}
			// encode concurrentMeasGapEUTRA_r17 optional
			if ie.concurrentMeasGapEUTRA_r17 != nil {
				if err = ie.concurrentMeasGapEUTRA_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode concurrentMeasGapEUTRA_r17", err)
				}
			}
			// encode serviceLinkPropDelayDiffReporting_r17 optional
			if ie.serviceLinkPropDelayDiffReporting_r17 != nil {
				if err = ie.serviceLinkPropDelayDiffReporting_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode serviceLinkPropDelayDiffReporting_r17", err)
				}
			}
			// encode ncsg_SymbolLevelScheduleRestrictionInter_r17 optional
			if ie.ncsg_SymbolLevelScheduleRestrictionInter_r17 != nil {
				if err = ie.ncsg_SymbolLevelScheduleRestrictionInter_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ncsg_SymbolLevelScheduleRestrictionInter_r17", err)
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
			optionals_ext_8 := []bool{ie.eventD1_MeasReportTrigger_r17 != nil, ie.independentGapConfig_maxCC_r17 != nil}
			for _, bit := range optionals_ext_8 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode eventD1_MeasReportTrigger_r17 optional
			if ie.eventD1_MeasReportTrigger_r17 != nil {
				if err = ie.eventD1_MeasReportTrigger_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode eventD1_MeasReportTrigger_r17", err)
				}
			}
			// encode independentGapConfig_maxCC_r17 optional
			if ie.independentGapConfig_maxCC_r17 != nil {
				if err = ie.independentGapConfig_maxCC_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode independentGapConfig_maxCC_r17", err)
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

func (ie *MeasAndMobParametersCommon) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var supportedGapPatternPresent bool
	if supportedGapPatternPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_RLMPresent bool
	if ssb_RLMPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_AndCSI_RS_RLMPresent bool
	if ssb_AndCSI_RS_RLMPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if supportedGapPatternPresent {
		var tmp_bs_supportedGapPattern []byte
		var n_supportedGapPattern uint
		if tmp_bs_supportedGapPattern, n_supportedGapPattern, err = r.ReadBitString(&uper.Constraint{Lb: 22, Ub: 22}, false); err != nil {
			return utils.WrapError("Decode supportedGapPattern", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_supportedGapPattern,
			NumBits: uint64(n_supportedGapPattern),
		}
		ie.supportedGapPattern = &tmp_bitstring
	}
	if ssb_RLMPresent {
		ie.ssb_RLM = new(MeasAndMobParametersCommon_ssb_RLM)
		if err = ie.ssb_RLM.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_RLM", err)
		}
	}
	if ssb_AndCSI_RS_RLMPresent {
		ie.ssb_AndCSI_RS_RLM = new(MeasAndMobParametersCommon_ssb_AndCSI_RS_RLM)
		if err = ie.ssb_AndCSI_RS_RLM.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_AndCSI_RS_RLM", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 8 bits for 8 extension groups
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

			eventB_MeasAndReportPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			handoverFDD_TDDPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			eutra_CGI_ReportingPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nr_CGI_ReportingPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode eventB_MeasAndReport optional
			if eventB_MeasAndReportPresent {
				ie.eventB_MeasAndReport = new(MeasAndMobParametersCommon_eventB_MeasAndReport)
				if err = ie.eventB_MeasAndReport.Decode(extReader); err != nil {
					return utils.WrapError("Decode eventB_MeasAndReport", err)
				}
			}
			// decode handoverFDD_TDD optional
			if handoverFDD_TDDPresent {
				ie.handoverFDD_TDD = new(MeasAndMobParametersCommon_handoverFDD_TDD)
				if err = ie.handoverFDD_TDD.Decode(extReader); err != nil {
					return utils.WrapError("Decode handoverFDD_TDD", err)
				}
			}
			// decode eutra_CGI_Reporting optional
			if eutra_CGI_ReportingPresent {
				ie.eutra_CGI_Reporting = new(MeasAndMobParametersCommon_eutra_CGI_Reporting)
				if err = ie.eutra_CGI_Reporting.Decode(extReader); err != nil {
					return utils.WrapError("Decode eutra_CGI_Reporting", err)
				}
			}
			// decode nr_CGI_Reporting optional
			if nr_CGI_ReportingPresent {
				ie.nr_CGI_Reporting = new(MeasAndMobParametersCommon_nr_CGI_Reporting)
				if err = ie.nr_CGI_Reporting.Decode(extReader); err != nil {
					return utils.WrapError("Decode nr_CGI_Reporting", err)
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

			independentGapConfigPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			periodicEUTRA_MeasAndReportPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			handoverFR1_FR2Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumberCSI_RS_RRM_RS_SINRPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode independentGapConfig optional
			if independentGapConfigPresent {
				ie.independentGapConfig = new(MeasAndMobParametersCommon_independentGapConfig)
				if err = ie.independentGapConfig.Decode(extReader); err != nil {
					return utils.WrapError("Decode independentGapConfig", err)
				}
			}
			// decode periodicEUTRA_MeasAndReport optional
			if periodicEUTRA_MeasAndReportPresent {
				ie.periodicEUTRA_MeasAndReport = new(MeasAndMobParametersCommon_periodicEUTRA_MeasAndReport)
				if err = ie.periodicEUTRA_MeasAndReport.Decode(extReader); err != nil {
					return utils.WrapError("Decode periodicEUTRA_MeasAndReport", err)
				}
			}
			// decode handoverFR1_FR2 optional
			if handoverFR1_FR2Present {
				ie.handoverFR1_FR2 = new(MeasAndMobParametersCommon_handoverFR1_FR2)
				if err = ie.handoverFR1_FR2.Decode(extReader); err != nil {
					return utils.WrapError("Decode handoverFR1_FR2", err)
				}
			}
			// decode maxNumberCSI_RS_RRM_RS_SINR optional
			if maxNumberCSI_RS_RRM_RS_SINRPresent {
				ie.maxNumberCSI_RS_RRM_RS_SINR = new(MeasAndMobParametersCommon_maxNumberCSI_RS_RRM_RS_SINR)
				if err = ie.maxNumberCSI_RS_RRM_RS_SINR.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxNumberCSI_RS_RRM_RS_SINR", err)
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

			nr_CGI_Reporting_ENDCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode nr_CGI_Reporting_ENDC optional
			if nr_CGI_Reporting_ENDCPresent {
				ie.nr_CGI_Reporting_ENDC = new(MeasAndMobParametersCommon_nr_CGI_Reporting_ENDC)
				if err = ie.nr_CGI_Reporting_ENDC.Decode(extReader); err != nil {
					return utils.WrapError("Decode nr_CGI_Reporting_ENDC", err)
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

			eutra_CGI_Reporting_NEDCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			eutra_CGI_Reporting_NRDCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nr_CGI_Reporting_NEDCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nr_CGI_Reporting_NRDCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode eutra_CGI_Reporting_NEDC optional
			if eutra_CGI_Reporting_NEDCPresent {
				ie.eutra_CGI_Reporting_NEDC = new(MeasAndMobParametersCommon_eutra_CGI_Reporting_NEDC)
				if err = ie.eutra_CGI_Reporting_NEDC.Decode(extReader); err != nil {
					return utils.WrapError("Decode eutra_CGI_Reporting_NEDC", err)
				}
			}
			// decode eutra_CGI_Reporting_NRDC optional
			if eutra_CGI_Reporting_NRDCPresent {
				ie.eutra_CGI_Reporting_NRDC = new(MeasAndMobParametersCommon_eutra_CGI_Reporting_NRDC)
				if err = ie.eutra_CGI_Reporting_NRDC.Decode(extReader); err != nil {
					return utils.WrapError("Decode eutra_CGI_Reporting_NRDC", err)
				}
			}
			// decode nr_CGI_Reporting_NEDC optional
			if nr_CGI_Reporting_NEDCPresent {
				ie.nr_CGI_Reporting_NEDC = new(MeasAndMobParametersCommon_nr_CGI_Reporting_NEDC)
				if err = ie.nr_CGI_Reporting_NEDC.Decode(extReader); err != nil {
					return utils.WrapError("Decode nr_CGI_Reporting_NEDC", err)
				}
			}
			// decode nr_CGI_Reporting_NRDC optional
			if nr_CGI_Reporting_NRDCPresent {
				ie.nr_CGI_Reporting_NRDC = new(MeasAndMobParametersCommon_nr_CGI_Reporting_NRDC)
				if err = ie.nr_CGI_Reporting_NRDC.Decode(extReader); err != nil {
					return utils.WrapError("Decode nr_CGI_Reporting_NRDC", err)
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

			reportAddNeighMeasForPeriodic_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			condHandoverParametersCommon_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nr_NeedForGap_Reporting_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedGapPattern_NRonly_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedGapPattern_NRonly_NEDC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumberCLI_RSSI_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumberCLI_SRS_RSRP_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumberPerSlotCLI_SRS_RSRP_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			mfbi_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dummyPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nr_CGI_Reporting_NPN_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			idleInactiveEUTRA_MeasReport_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			idleInactive_ValidityArea_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			eutra_AutonomousGaps_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			eutra_AutonomousGaps_NEDC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			eutra_AutonomousGaps_NRDC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pcellT312_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			supportedGapPattern_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode reportAddNeighMeasForPeriodic_r16 optional
			if reportAddNeighMeasForPeriodic_r16Present {
				ie.reportAddNeighMeasForPeriodic_r16 = new(MeasAndMobParametersCommon_reportAddNeighMeasForPeriodic_r16)
				if err = ie.reportAddNeighMeasForPeriodic_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode reportAddNeighMeasForPeriodic_r16", err)
				}
			}
			// decode condHandoverParametersCommon_r16 optional
			if condHandoverParametersCommon_r16Present {
				ie.condHandoverParametersCommon_r16 = new(MeasAndMobParametersCommon_condHandoverParametersCommon_r16)
				if err = ie.condHandoverParametersCommon_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode condHandoverParametersCommon_r16", err)
				}
			}
			// decode nr_NeedForGap_Reporting_r16 optional
			if nr_NeedForGap_Reporting_r16Present {
				ie.nr_NeedForGap_Reporting_r16 = new(MeasAndMobParametersCommon_nr_NeedForGap_Reporting_r16)
				if err = ie.nr_NeedForGap_Reporting_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode nr_NeedForGap_Reporting_r16", err)
				}
			}
			// decode supportedGapPattern_NRonly_r16 optional
			if supportedGapPattern_NRonly_r16Present {
				var tmp_bs_supportedGapPattern_NRonly_r16 []byte
				var n_supportedGapPattern_NRonly_r16 uint
				if tmp_bs_supportedGapPattern_NRonly_r16, n_supportedGapPattern_NRonly_r16, err = extReader.ReadBitString(&uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
					return utils.WrapError("Decode supportedGapPattern_NRonly_r16", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_supportedGapPattern_NRonly_r16,
					NumBits: uint64(n_supportedGapPattern_NRonly_r16),
				}
				ie.supportedGapPattern_NRonly_r16 = &tmp_bitstring
			}
			// decode supportedGapPattern_NRonly_NEDC_r16 optional
			if supportedGapPattern_NRonly_NEDC_r16Present {
				ie.supportedGapPattern_NRonly_NEDC_r16 = new(MeasAndMobParametersCommon_supportedGapPattern_NRonly_NEDC_r16)
				if err = ie.supportedGapPattern_NRonly_NEDC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode supportedGapPattern_NRonly_NEDC_r16", err)
				}
			}
			// decode maxNumberCLI_RSSI_r16 optional
			if maxNumberCLI_RSSI_r16Present {
				ie.maxNumberCLI_RSSI_r16 = new(MeasAndMobParametersCommon_maxNumberCLI_RSSI_r16)
				if err = ie.maxNumberCLI_RSSI_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxNumberCLI_RSSI_r16", err)
				}
			}
			// decode maxNumberCLI_SRS_RSRP_r16 optional
			if maxNumberCLI_SRS_RSRP_r16Present {
				ie.maxNumberCLI_SRS_RSRP_r16 = new(MeasAndMobParametersCommon_maxNumberCLI_SRS_RSRP_r16)
				if err = ie.maxNumberCLI_SRS_RSRP_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxNumberCLI_SRS_RSRP_r16", err)
				}
			}
			// decode maxNumberPerSlotCLI_SRS_RSRP_r16 optional
			if maxNumberPerSlotCLI_SRS_RSRP_r16Present {
				ie.maxNumberPerSlotCLI_SRS_RSRP_r16 = new(MeasAndMobParametersCommon_maxNumberPerSlotCLI_SRS_RSRP_r16)
				if err = ie.maxNumberPerSlotCLI_SRS_RSRP_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxNumberPerSlotCLI_SRS_RSRP_r16", err)
				}
			}
			// decode mfbi_IAB_r16 optional
			if mfbi_IAB_r16Present {
				ie.mfbi_IAB_r16 = new(MeasAndMobParametersCommon_mfbi_IAB_r16)
				if err = ie.mfbi_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode mfbi_IAB_r16", err)
				}
			}
			// decode dummy optional
			if dummyPresent {
				ie.dummy = new(MeasAndMobParametersCommon_dummy)
				if err = ie.dummy.Decode(extReader); err != nil {
					return utils.WrapError("Decode dummy", err)
				}
			}
			// decode nr_CGI_Reporting_NPN_r16 optional
			if nr_CGI_Reporting_NPN_r16Present {
				ie.nr_CGI_Reporting_NPN_r16 = new(MeasAndMobParametersCommon_nr_CGI_Reporting_NPN_r16)
				if err = ie.nr_CGI_Reporting_NPN_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode nr_CGI_Reporting_NPN_r16", err)
				}
			}
			// decode idleInactiveEUTRA_MeasReport_r16 optional
			if idleInactiveEUTRA_MeasReport_r16Present {
				ie.idleInactiveEUTRA_MeasReport_r16 = new(MeasAndMobParametersCommon_idleInactiveEUTRA_MeasReport_r16)
				if err = ie.idleInactiveEUTRA_MeasReport_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode idleInactiveEUTRA_MeasReport_r16", err)
				}
			}
			// decode idleInactive_ValidityArea_r16 optional
			if idleInactive_ValidityArea_r16Present {
				ie.idleInactive_ValidityArea_r16 = new(MeasAndMobParametersCommon_idleInactive_ValidityArea_r16)
				if err = ie.idleInactive_ValidityArea_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode idleInactive_ValidityArea_r16", err)
				}
			}
			// decode eutra_AutonomousGaps_r16 optional
			if eutra_AutonomousGaps_r16Present {
				ie.eutra_AutonomousGaps_r16 = new(MeasAndMobParametersCommon_eutra_AutonomousGaps_r16)
				if err = ie.eutra_AutonomousGaps_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode eutra_AutonomousGaps_r16", err)
				}
			}
			// decode eutra_AutonomousGaps_NEDC_r16 optional
			if eutra_AutonomousGaps_NEDC_r16Present {
				ie.eutra_AutonomousGaps_NEDC_r16 = new(MeasAndMobParametersCommon_eutra_AutonomousGaps_NEDC_r16)
				if err = ie.eutra_AutonomousGaps_NEDC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode eutra_AutonomousGaps_NEDC_r16", err)
				}
			}
			// decode eutra_AutonomousGaps_NRDC_r16 optional
			if eutra_AutonomousGaps_NRDC_r16Present {
				ie.eutra_AutonomousGaps_NRDC_r16 = new(MeasAndMobParametersCommon_eutra_AutonomousGaps_NRDC_r16)
				if err = ie.eutra_AutonomousGaps_NRDC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode eutra_AutonomousGaps_NRDC_r16", err)
				}
			}
			// decode pcellT312_r16 optional
			if pcellT312_r16Present {
				ie.pcellT312_r16 = new(MeasAndMobParametersCommon_pcellT312_r16)
				if err = ie.pcellT312_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pcellT312_r16", err)
				}
			}
			// decode supportedGapPattern_r16 optional
			if supportedGapPattern_r16Present {
				var tmp_bs_supportedGapPattern_r16 []byte
				var n_supportedGapPattern_r16 uint
				if tmp_bs_supportedGapPattern_r16, n_supportedGapPattern_r16, err = extReader.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
					return utils.WrapError("Decode supportedGapPattern_r16", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_supportedGapPattern_r16,
					NumBits: uint64(n_supportedGapPattern_r16),
				}
				ie.supportedGapPattern_r16 = &tmp_bitstring
			}
		}
		// decode extension group 6
		if len(extBitmap) > 5 && extBitmap[5] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			concurrentMeasGap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nr_NeedForGapNCSG_Reporting_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			eutra_NeedForGapNCSG_Reporting_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ncsg_MeasGapPerFR_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ncsg_MeasGapPatterns_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ncsg_MeasGapNR_Patterns_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			preconfiguredUE_AutonomousMeasGap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			preconfiguredNW_ControlledMeasGap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			handoverFR1_FR2_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			handoverFR2_1_FR2_2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			independentGapConfigPRS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			rrm_RelaxationRRC_ConnectedRedCap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			parallelMeasurementGap_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			condHandoverWithSCG_NRDC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			gNB_ID_LengthReporting_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			gNB_ID_LengthReporting_ENDC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			gNB_ID_LengthReporting_NEDC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			gNB_ID_LengthReporting_NRDC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			gNB_ID_LengthReporting_NPN_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode concurrentMeasGap_r17 optional
			if concurrentMeasGap_r17Present {
				ie.concurrentMeasGap_r17 = new(MeasAndMobParametersCommon_concurrentMeasGap_r17)
				if err = ie.concurrentMeasGap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode concurrentMeasGap_r17", err)
				}
			}
			// decode nr_NeedForGapNCSG_Reporting_r17 optional
			if nr_NeedForGapNCSG_Reporting_r17Present {
				ie.nr_NeedForGapNCSG_Reporting_r17 = new(MeasAndMobParametersCommon_nr_NeedForGapNCSG_Reporting_r17)
				if err = ie.nr_NeedForGapNCSG_Reporting_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode nr_NeedForGapNCSG_Reporting_r17", err)
				}
			}
			// decode eutra_NeedForGapNCSG_Reporting_r17 optional
			if eutra_NeedForGapNCSG_Reporting_r17Present {
				ie.eutra_NeedForGapNCSG_Reporting_r17 = new(MeasAndMobParametersCommon_eutra_NeedForGapNCSG_Reporting_r17)
				if err = ie.eutra_NeedForGapNCSG_Reporting_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode eutra_NeedForGapNCSG_Reporting_r17", err)
				}
			}
			// decode ncsg_MeasGapPerFR_r17 optional
			if ncsg_MeasGapPerFR_r17Present {
				ie.ncsg_MeasGapPerFR_r17 = new(MeasAndMobParametersCommon_ncsg_MeasGapPerFR_r17)
				if err = ie.ncsg_MeasGapPerFR_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ncsg_MeasGapPerFR_r17", err)
				}
			}
			// decode ncsg_MeasGapPatterns_r17 optional
			if ncsg_MeasGapPatterns_r17Present {
				var tmp_bs_ncsg_MeasGapPatterns_r17 []byte
				var n_ncsg_MeasGapPatterns_r17 uint
				if tmp_bs_ncsg_MeasGapPatterns_r17, n_ncsg_MeasGapPatterns_r17, err = extReader.ReadBitString(&uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
					return utils.WrapError("Decode ncsg_MeasGapPatterns_r17", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_ncsg_MeasGapPatterns_r17,
					NumBits: uint64(n_ncsg_MeasGapPatterns_r17),
				}
				ie.ncsg_MeasGapPatterns_r17 = &tmp_bitstring
			}
			// decode ncsg_MeasGapNR_Patterns_r17 optional
			if ncsg_MeasGapNR_Patterns_r17Present {
				var tmp_bs_ncsg_MeasGapNR_Patterns_r17 []byte
				var n_ncsg_MeasGapNR_Patterns_r17 uint
				if tmp_bs_ncsg_MeasGapNR_Patterns_r17, n_ncsg_MeasGapNR_Patterns_r17, err = extReader.ReadBitString(&uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
					return utils.WrapError("Decode ncsg_MeasGapNR_Patterns_r17", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_ncsg_MeasGapNR_Patterns_r17,
					NumBits: uint64(n_ncsg_MeasGapNR_Patterns_r17),
				}
				ie.ncsg_MeasGapNR_Patterns_r17 = &tmp_bitstring
			}
			// decode preconfiguredUE_AutonomousMeasGap_r17 optional
			if preconfiguredUE_AutonomousMeasGap_r17Present {
				ie.preconfiguredUE_AutonomousMeasGap_r17 = new(MeasAndMobParametersCommon_preconfiguredUE_AutonomousMeasGap_r17)
				if err = ie.preconfiguredUE_AutonomousMeasGap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode preconfiguredUE_AutonomousMeasGap_r17", err)
				}
			}
			// decode preconfiguredNW_ControlledMeasGap_r17 optional
			if preconfiguredNW_ControlledMeasGap_r17Present {
				ie.preconfiguredNW_ControlledMeasGap_r17 = new(MeasAndMobParametersCommon_preconfiguredNW_ControlledMeasGap_r17)
				if err = ie.preconfiguredNW_ControlledMeasGap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode preconfiguredNW_ControlledMeasGap_r17", err)
				}
			}
			// decode handoverFR1_FR2_2_r17 optional
			if handoverFR1_FR2_2_r17Present {
				ie.handoverFR1_FR2_2_r17 = new(MeasAndMobParametersCommon_handoverFR1_FR2_2_r17)
				if err = ie.handoverFR1_FR2_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode handoverFR1_FR2_2_r17", err)
				}
			}
			// decode handoverFR2_1_FR2_2_r17 optional
			if handoverFR2_1_FR2_2_r17Present {
				ie.handoverFR2_1_FR2_2_r17 = new(MeasAndMobParametersCommon_handoverFR2_1_FR2_2_r17)
				if err = ie.handoverFR2_1_FR2_2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode handoverFR2_1_FR2_2_r17", err)
				}
			}
			// decode independentGapConfigPRS_r17 optional
			if independentGapConfigPRS_r17Present {
				ie.independentGapConfigPRS_r17 = new(MeasAndMobParametersCommon_independentGapConfigPRS_r17)
				if err = ie.independentGapConfigPRS_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode independentGapConfigPRS_r17", err)
				}
			}
			// decode rrm_RelaxationRRC_ConnectedRedCap_r17 optional
			if rrm_RelaxationRRC_ConnectedRedCap_r17Present {
				ie.rrm_RelaxationRRC_ConnectedRedCap_r17 = new(MeasAndMobParametersCommon_rrm_RelaxationRRC_ConnectedRedCap_r17)
				if err = ie.rrm_RelaxationRRC_ConnectedRedCap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode rrm_RelaxationRRC_ConnectedRedCap_r17", err)
				}
			}
			// decode parallelMeasurementGap_r17 optional
			if parallelMeasurementGap_r17Present {
				ie.parallelMeasurementGap_r17 = new(MeasAndMobParametersCommon_parallelMeasurementGap_r17)
				if err = ie.parallelMeasurementGap_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode parallelMeasurementGap_r17", err)
				}
			}
			// decode condHandoverWithSCG_NRDC_r17 optional
			if condHandoverWithSCG_NRDC_r17Present {
				ie.condHandoverWithSCG_NRDC_r17 = new(MeasAndMobParametersCommon_condHandoverWithSCG_NRDC_r17)
				if err = ie.condHandoverWithSCG_NRDC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode condHandoverWithSCG_NRDC_r17", err)
				}
			}
			// decode gNB_ID_LengthReporting_r17 optional
			if gNB_ID_LengthReporting_r17Present {
				ie.gNB_ID_LengthReporting_r17 = new(MeasAndMobParametersCommon_gNB_ID_LengthReporting_r17)
				if err = ie.gNB_ID_LengthReporting_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode gNB_ID_LengthReporting_r17", err)
				}
			}
			// decode gNB_ID_LengthReporting_ENDC_r17 optional
			if gNB_ID_LengthReporting_ENDC_r17Present {
				ie.gNB_ID_LengthReporting_ENDC_r17 = new(MeasAndMobParametersCommon_gNB_ID_LengthReporting_ENDC_r17)
				if err = ie.gNB_ID_LengthReporting_ENDC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode gNB_ID_LengthReporting_ENDC_r17", err)
				}
			}
			// decode gNB_ID_LengthReporting_NEDC_r17 optional
			if gNB_ID_LengthReporting_NEDC_r17Present {
				ie.gNB_ID_LengthReporting_NEDC_r17 = new(MeasAndMobParametersCommon_gNB_ID_LengthReporting_NEDC_r17)
				if err = ie.gNB_ID_LengthReporting_NEDC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode gNB_ID_LengthReporting_NEDC_r17", err)
				}
			}
			// decode gNB_ID_LengthReporting_NRDC_r17 optional
			if gNB_ID_LengthReporting_NRDC_r17Present {
				ie.gNB_ID_LengthReporting_NRDC_r17 = new(MeasAndMobParametersCommon_gNB_ID_LengthReporting_NRDC_r17)
				if err = ie.gNB_ID_LengthReporting_NRDC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode gNB_ID_LengthReporting_NRDC_r17", err)
				}
			}
			// decode gNB_ID_LengthReporting_NPN_r17 optional
			if gNB_ID_LengthReporting_NPN_r17Present {
				ie.gNB_ID_LengthReporting_NPN_r17 = new(MeasAndMobParametersCommon_gNB_ID_LengthReporting_NPN_r17)
				if err = ie.gNB_ID_LengthReporting_NPN_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode gNB_ID_LengthReporting_NPN_r17", err)
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

			parallelSMTC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			concurrentMeasGapEUTRA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			serviceLinkPropDelayDiffReporting_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ncsg_SymbolLevelScheduleRestrictionInter_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode parallelSMTC_r17 optional
			if parallelSMTC_r17Present {
				ie.parallelSMTC_r17 = new(MeasAndMobParametersCommon_parallelSMTC_r17)
				if err = ie.parallelSMTC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode parallelSMTC_r17", err)
				}
			}
			// decode concurrentMeasGapEUTRA_r17 optional
			if concurrentMeasGapEUTRA_r17Present {
				ie.concurrentMeasGapEUTRA_r17 = new(MeasAndMobParametersCommon_concurrentMeasGapEUTRA_r17)
				if err = ie.concurrentMeasGapEUTRA_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode concurrentMeasGapEUTRA_r17", err)
				}
			}
			// decode serviceLinkPropDelayDiffReporting_r17 optional
			if serviceLinkPropDelayDiffReporting_r17Present {
				ie.serviceLinkPropDelayDiffReporting_r17 = new(MeasAndMobParametersCommon_serviceLinkPropDelayDiffReporting_r17)
				if err = ie.serviceLinkPropDelayDiffReporting_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode serviceLinkPropDelayDiffReporting_r17", err)
				}
			}
			// decode ncsg_SymbolLevelScheduleRestrictionInter_r17 optional
			if ncsg_SymbolLevelScheduleRestrictionInter_r17Present {
				ie.ncsg_SymbolLevelScheduleRestrictionInter_r17 = new(MeasAndMobParametersCommon_ncsg_SymbolLevelScheduleRestrictionInter_r17)
				if err = ie.ncsg_SymbolLevelScheduleRestrictionInter_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ncsg_SymbolLevelScheduleRestrictionInter_r17", err)
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

			eventD1_MeasReportTrigger_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			independentGapConfig_maxCC_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode eventD1_MeasReportTrigger_r17 optional
			if eventD1_MeasReportTrigger_r17Present {
				ie.eventD1_MeasReportTrigger_r17 = new(MeasAndMobParametersCommon_eventD1_MeasReportTrigger_r17)
				if err = ie.eventD1_MeasReportTrigger_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode eventD1_MeasReportTrigger_r17", err)
				}
			}
			// decode independentGapConfig_maxCC_r17 optional
			if independentGapConfig_maxCC_r17Present {
				ie.independentGapConfig_maxCC_r17 = new(MeasAndMobParametersCommon_independentGapConfig_maxCC_r17)
				if err = ie.independentGapConfig_maxCC_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode independentGapConfig_maxCC_r17", err)
				}
			}
		}
	}
	return nil
}
