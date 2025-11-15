package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersFRX_Diff struct {
	ss_SINR_Meas                                   *MeasAndMobParametersFRX_Diff_ss_SINR_Meas                                   `optional`
	csi_RSRP_AndRSRQ_MeasWithSSB                   *MeasAndMobParametersFRX_Diff_csi_RSRP_AndRSRQ_MeasWithSSB                   `optional`
	csi_RSRP_AndRSRQ_MeasWithoutSSB                *MeasAndMobParametersFRX_Diff_csi_RSRP_AndRSRQ_MeasWithoutSSB                `optional`
	csi_SINR_Meas                                  *MeasAndMobParametersFRX_Diff_csi_SINR_Meas                                  `optional`
	csi_RS_RLM                                     *MeasAndMobParametersFRX_Diff_csi_RS_RLM                                     `optional`
	handoverInterF                                 *MeasAndMobParametersFRX_Diff_handoverInterF                                 `optional,ext-1`
	handoverLTE_EPC                                *MeasAndMobParametersFRX_Diff_handoverLTE_EPC                                `optional,ext-1`
	handoverLTE_5GC                                *MeasAndMobParametersFRX_Diff_handoverLTE_5GC                                `optional,ext-1`
	maxNumberResource_CSI_RS_RLM                   *MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM                   `optional,ext-2`
	simultaneousRxDataSSB_DiffNumerology           *MeasAndMobParametersFRX_Diff_simultaneousRxDataSSB_DiffNumerology           `optional,ext-3`
	nr_AutonomousGaps_r16                          *MeasAndMobParametersFRX_Diff_nr_AutonomousGaps_r16                          `optional,ext-4`
	nr_AutonomousGaps_ENDC_r16                     *MeasAndMobParametersFRX_Diff_nr_AutonomousGaps_ENDC_r16                     `optional,ext-4`
	nr_AutonomousGaps_NEDC_r16                     *MeasAndMobParametersFRX_Diff_nr_AutonomousGaps_NEDC_r16                     `optional,ext-4`
	nr_AutonomousGaps_NRDC_r16                     *MeasAndMobParametersFRX_Diff_nr_AutonomousGaps_NRDC_r16                     `optional,ext-4`
	dummy                                          *MeasAndMobParametersFRX_Diff_dummy                                          `optional,ext-4`
	cli_RSSI_Meas_r16                              *MeasAndMobParametersFRX_Diff_cli_RSSI_Meas_r16                              `optional,ext-4`
	cli_SRS_RSRP_Meas_r16                          *MeasAndMobParametersFRX_Diff_cli_SRS_RSRP_Meas_r16                          `optional,ext-4`
	interFrequencyMeas_NoGap_r16                   *MeasAndMobParametersFRX_Diff_interFrequencyMeas_NoGap_r16                   `optional,ext-4`
	simultaneousRxDataSSB_DiffNumerology_Inter_r16 *MeasAndMobParametersFRX_Diff_simultaneousRxDataSSB_DiffNumerology_Inter_r16 `optional,ext-4`
	idleInactiveNR_MeasReport_r16                  *MeasAndMobParametersFRX_Diff_idleInactiveNR_MeasReport_r16                  `optional,ext-4`
	idleInactiveNR_MeasBeamReport_r16              *MeasAndMobParametersFRX_Diff_idleInactiveNR_MeasBeamReport_r16              `optional,ext-4`
	increasedNumberofCSIRSPerMO_r16                *MeasAndMobParametersFRX_Diff_increasedNumberofCSIRSPerMO_r16                `optional,ext-5`
}

func (ie *MeasAndMobParametersFRX_Diff) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.handoverInterF != nil || ie.handoverLTE_EPC != nil || ie.handoverLTE_5GC != nil || ie.maxNumberResource_CSI_RS_RLM != nil || ie.simultaneousRxDataSSB_DiffNumerology != nil || ie.nr_AutonomousGaps_r16 != nil || ie.nr_AutonomousGaps_ENDC_r16 != nil || ie.nr_AutonomousGaps_NEDC_r16 != nil || ie.nr_AutonomousGaps_NRDC_r16 != nil || ie.dummy != nil || ie.cli_RSSI_Meas_r16 != nil || ie.cli_SRS_RSRP_Meas_r16 != nil || ie.interFrequencyMeas_NoGap_r16 != nil || ie.simultaneousRxDataSSB_DiffNumerology_Inter_r16 != nil || ie.idleInactiveNR_MeasReport_r16 != nil || ie.idleInactiveNR_MeasBeamReport_r16 != nil || ie.increasedNumberofCSIRSPerMO_r16 != nil
	preambleBits := []bool{hasExtensions, ie.ss_SINR_Meas != nil, ie.csi_RSRP_AndRSRQ_MeasWithSSB != nil, ie.csi_RSRP_AndRSRQ_MeasWithoutSSB != nil, ie.csi_SINR_Meas != nil, ie.csi_RS_RLM != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ss_SINR_Meas != nil {
		if err = ie.ss_SINR_Meas.Encode(w); err != nil {
			return utils.WrapError("Encode ss_SINR_Meas", err)
		}
	}
	if ie.csi_RSRP_AndRSRQ_MeasWithSSB != nil {
		if err = ie.csi_RSRP_AndRSRQ_MeasWithSSB.Encode(w); err != nil {
			return utils.WrapError("Encode csi_RSRP_AndRSRQ_MeasWithSSB", err)
		}
	}
	if ie.csi_RSRP_AndRSRQ_MeasWithoutSSB != nil {
		if err = ie.csi_RSRP_AndRSRQ_MeasWithoutSSB.Encode(w); err != nil {
			return utils.WrapError("Encode csi_RSRP_AndRSRQ_MeasWithoutSSB", err)
		}
	}
	if ie.csi_SINR_Meas != nil {
		if err = ie.csi_SINR_Meas.Encode(w); err != nil {
			return utils.WrapError("Encode csi_SINR_Meas", err)
		}
	}
	if ie.csi_RS_RLM != nil {
		if err = ie.csi_RS_RLM.Encode(w); err != nil {
			return utils.WrapError("Encode csi_RS_RLM", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 5 bits for 5 extension groups
		extBitmap := []bool{ie.handoverInterF != nil || ie.handoverLTE_EPC != nil || ie.handoverLTE_5GC != nil, ie.maxNumberResource_CSI_RS_RLM != nil, ie.simultaneousRxDataSSB_DiffNumerology != nil, ie.nr_AutonomousGaps_r16 != nil || ie.nr_AutonomousGaps_ENDC_r16 != nil || ie.nr_AutonomousGaps_NEDC_r16 != nil || ie.nr_AutonomousGaps_NRDC_r16 != nil || ie.dummy != nil || ie.cli_RSSI_Meas_r16 != nil || ie.cli_SRS_RSRP_Meas_r16 != nil || ie.interFrequencyMeas_NoGap_r16 != nil || ie.simultaneousRxDataSSB_DiffNumerology_Inter_r16 != nil || ie.idleInactiveNR_MeasReport_r16 != nil || ie.idleInactiveNR_MeasBeamReport_r16 != nil, ie.increasedNumberofCSIRSPerMO_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasAndMobParametersFRX_Diff", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.handoverInterF != nil, ie.handoverLTE_EPC != nil, ie.handoverLTE_5GC != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode handoverInterF optional
			if ie.handoverInterF != nil {
				if err = ie.handoverInterF.Encode(extWriter); err != nil {
					return utils.WrapError("Encode handoverInterF", err)
				}
			}
			// encode handoverLTE_EPC optional
			if ie.handoverLTE_EPC != nil {
				if err = ie.handoverLTE_EPC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode handoverLTE_EPC", err)
				}
			}
			// encode handoverLTE_5GC optional
			if ie.handoverLTE_5GC != nil {
				if err = ie.handoverLTE_5GC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode handoverLTE_5GC", err)
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
			optionals_ext_2 := []bool{ie.maxNumberResource_CSI_RS_RLM != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode maxNumberResource_CSI_RS_RLM optional
			if ie.maxNumberResource_CSI_RS_RLM != nil {
				if err = ie.maxNumberResource_CSI_RS_RLM.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxNumberResource_CSI_RS_RLM", err)
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
			optionals_ext_3 := []bool{ie.simultaneousRxDataSSB_DiffNumerology != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode simultaneousRxDataSSB_DiffNumerology optional
			if ie.simultaneousRxDataSSB_DiffNumerology != nil {
				if err = ie.simultaneousRxDataSSB_DiffNumerology.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simultaneousRxDataSSB_DiffNumerology", err)
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
			optionals_ext_4 := []bool{ie.nr_AutonomousGaps_r16 != nil, ie.nr_AutonomousGaps_ENDC_r16 != nil, ie.nr_AutonomousGaps_NEDC_r16 != nil, ie.nr_AutonomousGaps_NRDC_r16 != nil, ie.dummy != nil, ie.cli_RSSI_Meas_r16 != nil, ie.cli_SRS_RSRP_Meas_r16 != nil, ie.interFrequencyMeas_NoGap_r16 != nil, ie.simultaneousRxDataSSB_DiffNumerology_Inter_r16 != nil, ie.idleInactiveNR_MeasReport_r16 != nil, ie.idleInactiveNR_MeasBeamReport_r16 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode nr_AutonomousGaps_r16 optional
			if ie.nr_AutonomousGaps_r16 != nil {
				if err = ie.nr_AutonomousGaps_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nr_AutonomousGaps_r16", err)
				}
			}
			// encode nr_AutonomousGaps_ENDC_r16 optional
			if ie.nr_AutonomousGaps_ENDC_r16 != nil {
				if err = ie.nr_AutonomousGaps_ENDC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nr_AutonomousGaps_ENDC_r16", err)
				}
			}
			// encode nr_AutonomousGaps_NEDC_r16 optional
			if ie.nr_AutonomousGaps_NEDC_r16 != nil {
				if err = ie.nr_AutonomousGaps_NEDC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nr_AutonomousGaps_NEDC_r16", err)
				}
			}
			// encode nr_AutonomousGaps_NRDC_r16 optional
			if ie.nr_AutonomousGaps_NRDC_r16 != nil {
				if err = ie.nr_AutonomousGaps_NRDC_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nr_AutonomousGaps_NRDC_r16", err)
				}
			}
			// encode dummy optional
			if ie.dummy != nil {
				if err = ie.dummy.Encode(extWriter); err != nil {
					return utils.WrapError("Encode dummy", err)
				}
			}
			// encode cli_RSSI_Meas_r16 optional
			if ie.cli_RSSI_Meas_r16 != nil {
				if err = ie.cli_RSSI_Meas_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cli_RSSI_Meas_r16", err)
				}
			}
			// encode cli_SRS_RSRP_Meas_r16 optional
			if ie.cli_SRS_RSRP_Meas_r16 != nil {
				if err = ie.cli_SRS_RSRP_Meas_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cli_SRS_RSRP_Meas_r16", err)
				}
			}
			// encode interFrequencyMeas_NoGap_r16 optional
			if ie.interFrequencyMeas_NoGap_r16 != nil {
				if err = ie.interFrequencyMeas_NoGap_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode interFrequencyMeas_NoGap_r16", err)
				}
			}
			// encode simultaneousRxDataSSB_DiffNumerology_Inter_r16 optional
			if ie.simultaneousRxDataSSB_DiffNumerology_Inter_r16 != nil {
				if err = ie.simultaneousRxDataSSB_DiffNumerology_Inter_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode simultaneousRxDataSSB_DiffNumerology_Inter_r16", err)
				}
			}
			// encode idleInactiveNR_MeasReport_r16 optional
			if ie.idleInactiveNR_MeasReport_r16 != nil {
				if err = ie.idleInactiveNR_MeasReport_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode idleInactiveNR_MeasReport_r16", err)
				}
			}
			// encode idleInactiveNR_MeasBeamReport_r16 optional
			if ie.idleInactiveNR_MeasBeamReport_r16 != nil {
				if err = ie.idleInactiveNR_MeasBeamReport_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode idleInactiveNR_MeasBeamReport_r16", err)
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
			optionals_ext_5 := []bool{ie.increasedNumberofCSIRSPerMO_r16 != nil}
			for _, bit := range optionals_ext_5 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode increasedNumberofCSIRSPerMO_r16 optional
			if ie.increasedNumberofCSIRSPerMO_r16 != nil {
				if err = ie.increasedNumberofCSIRSPerMO_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode increasedNumberofCSIRSPerMO_r16", err)
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

func (ie *MeasAndMobParametersFRX_Diff) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ss_SINR_MeasPresent bool
	if ss_SINR_MeasPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_RSRP_AndRSRQ_MeasWithSSBPresent bool
	if csi_RSRP_AndRSRQ_MeasWithSSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_RSRP_AndRSRQ_MeasWithoutSSBPresent bool
	if csi_RSRP_AndRSRQ_MeasWithoutSSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_SINR_MeasPresent bool
	if csi_SINR_MeasPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_RS_RLMPresent bool
	if csi_RS_RLMPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ss_SINR_MeasPresent {
		ie.ss_SINR_Meas = new(MeasAndMobParametersFRX_Diff_ss_SINR_Meas)
		if err = ie.ss_SINR_Meas.Decode(r); err != nil {
			return utils.WrapError("Decode ss_SINR_Meas", err)
		}
	}
	if csi_RSRP_AndRSRQ_MeasWithSSBPresent {
		ie.csi_RSRP_AndRSRQ_MeasWithSSB = new(MeasAndMobParametersFRX_Diff_csi_RSRP_AndRSRQ_MeasWithSSB)
		if err = ie.csi_RSRP_AndRSRQ_MeasWithSSB.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RSRP_AndRSRQ_MeasWithSSB", err)
		}
	}
	if csi_RSRP_AndRSRQ_MeasWithoutSSBPresent {
		ie.csi_RSRP_AndRSRQ_MeasWithoutSSB = new(MeasAndMobParametersFRX_Diff_csi_RSRP_AndRSRQ_MeasWithoutSSB)
		if err = ie.csi_RSRP_AndRSRQ_MeasWithoutSSB.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RSRP_AndRSRQ_MeasWithoutSSB", err)
		}
	}
	if csi_SINR_MeasPresent {
		ie.csi_SINR_Meas = new(MeasAndMobParametersFRX_Diff_csi_SINR_Meas)
		if err = ie.csi_SINR_Meas.Decode(r); err != nil {
			return utils.WrapError("Decode csi_SINR_Meas", err)
		}
	}
	if csi_RS_RLMPresent {
		ie.csi_RS_RLM = new(MeasAndMobParametersFRX_Diff_csi_RS_RLM)
		if err = ie.csi_RS_RLM.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RS_RLM", err)
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

			handoverInterFPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			handoverLTE_EPCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			handoverLTE_5GCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode handoverInterF optional
			if handoverInterFPresent {
				ie.handoverInterF = new(MeasAndMobParametersFRX_Diff_handoverInterF)
				if err = ie.handoverInterF.Decode(extReader); err != nil {
					return utils.WrapError("Decode handoverInterF", err)
				}
			}
			// decode handoverLTE_EPC optional
			if handoverLTE_EPCPresent {
				ie.handoverLTE_EPC = new(MeasAndMobParametersFRX_Diff_handoverLTE_EPC)
				if err = ie.handoverLTE_EPC.Decode(extReader); err != nil {
					return utils.WrapError("Decode handoverLTE_EPC", err)
				}
			}
			// decode handoverLTE_5GC optional
			if handoverLTE_5GCPresent {
				ie.handoverLTE_5GC = new(MeasAndMobParametersFRX_Diff_handoverLTE_5GC)
				if err = ie.handoverLTE_5GC.Decode(extReader); err != nil {
					return utils.WrapError("Decode handoverLTE_5GC", err)
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

			maxNumberResource_CSI_RS_RLMPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode maxNumberResource_CSI_RS_RLM optional
			if maxNumberResource_CSI_RS_RLMPresent {
				ie.maxNumberResource_CSI_RS_RLM = new(MeasAndMobParametersFRX_Diff_maxNumberResource_CSI_RS_RLM)
				if err = ie.maxNumberResource_CSI_RS_RLM.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxNumberResource_CSI_RS_RLM", err)
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

			simultaneousRxDataSSB_DiffNumerologyPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode simultaneousRxDataSSB_DiffNumerology optional
			if simultaneousRxDataSSB_DiffNumerologyPresent {
				ie.simultaneousRxDataSSB_DiffNumerology = new(MeasAndMobParametersFRX_Diff_simultaneousRxDataSSB_DiffNumerology)
				if err = ie.simultaneousRxDataSSB_DiffNumerology.Decode(extReader); err != nil {
					return utils.WrapError("Decode simultaneousRxDataSSB_DiffNumerology", err)
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

			nr_AutonomousGaps_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nr_AutonomousGaps_ENDC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nr_AutonomousGaps_NEDC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nr_AutonomousGaps_NRDC_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dummyPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cli_RSSI_Meas_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cli_SRS_RSRP_Meas_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			interFrequencyMeas_NoGap_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			simultaneousRxDataSSB_DiffNumerology_Inter_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			idleInactiveNR_MeasReport_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			idleInactiveNR_MeasBeamReport_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode nr_AutonomousGaps_r16 optional
			if nr_AutonomousGaps_r16Present {
				ie.nr_AutonomousGaps_r16 = new(MeasAndMobParametersFRX_Diff_nr_AutonomousGaps_r16)
				if err = ie.nr_AutonomousGaps_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode nr_AutonomousGaps_r16", err)
				}
			}
			// decode nr_AutonomousGaps_ENDC_r16 optional
			if nr_AutonomousGaps_ENDC_r16Present {
				ie.nr_AutonomousGaps_ENDC_r16 = new(MeasAndMobParametersFRX_Diff_nr_AutonomousGaps_ENDC_r16)
				if err = ie.nr_AutonomousGaps_ENDC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode nr_AutonomousGaps_ENDC_r16", err)
				}
			}
			// decode nr_AutonomousGaps_NEDC_r16 optional
			if nr_AutonomousGaps_NEDC_r16Present {
				ie.nr_AutonomousGaps_NEDC_r16 = new(MeasAndMobParametersFRX_Diff_nr_AutonomousGaps_NEDC_r16)
				if err = ie.nr_AutonomousGaps_NEDC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode nr_AutonomousGaps_NEDC_r16", err)
				}
			}
			// decode nr_AutonomousGaps_NRDC_r16 optional
			if nr_AutonomousGaps_NRDC_r16Present {
				ie.nr_AutonomousGaps_NRDC_r16 = new(MeasAndMobParametersFRX_Diff_nr_AutonomousGaps_NRDC_r16)
				if err = ie.nr_AutonomousGaps_NRDC_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode nr_AutonomousGaps_NRDC_r16", err)
				}
			}
			// decode dummy optional
			if dummyPresent {
				ie.dummy = new(MeasAndMobParametersFRX_Diff_dummy)
				if err = ie.dummy.Decode(extReader); err != nil {
					return utils.WrapError("Decode dummy", err)
				}
			}
			// decode cli_RSSI_Meas_r16 optional
			if cli_RSSI_Meas_r16Present {
				ie.cli_RSSI_Meas_r16 = new(MeasAndMobParametersFRX_Diff_cli_RSSI_Meas_r16)
				if err = ie.cli_RSSI_Meas_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode cli_RSSI_Meas_r16", err)
				}
			}
			// decode cli_SRS_RSRP_Meas_r16 optional
			if cli_SRS_RSRP_Meas_r16Present {
				ie.cli_SRS_RSRP_Meas_r16 = new(MeasAndMobParametersFRX_Diff_cli_SRS_RSRP_Meas_r16)
				if err = ie.cli_SRS_RSRP_Meas_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode cli_SRS_RSRP_Meas_r16", err)
				}
			}
			// decode interFrequencyMeas_NoGap_r16 optional
			if interFrequencyMeas_NoGap_r16Present {
				ie.interFrequencyMeas_NoGap_r16 = new(MeasAndMobParametersFRX_Diff_interFrequencyMeas_NoGap_r16)
				if err = ie.interFrequencyMeas_NoGap_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode interFrequencyMeas_NoGap_r16", err)
				}
			}
			// decode simultaneousRxDataSSB_DiffNumerology_Inter_r16 optional
			if simultaneousRxDataSSB_DiffNumerology_Inter_r16Present {
				ie.simultaneousRxDataSSB_DiffNumerology_Inter_r16 = new(MeasAndMobParametersFRX_Diff_simultaneousRxDataSSB_DiffNumerology_Inter_r16)
				if err = ie.simultaneousRxDataSSB_DiffNumerology_Inter_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode simultaneousRxDataSSB_DiffNumerology_Inter_r16", err)
				}
			}
			// decode idleInactiveNR_MeasReport_r16 optional
			if idleInactiveNR_MeasReport_r16Present {
				ie.idleInactiveNR_MeasReport_r16 = new(MeasAndMobParametersFRX_Diff_idleInactiveNR_MeasReport_r16)
				if err = ie.idleInactiveNR_MeasReport_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode idleInactiveNR_MeasReport_r16", err)
				}
			}
			// decode idleInactiveNR_MeasBeamReport_r16 optional
			if idleInactiveNR_MeasBeamReport_r16Present {
				ie.idleInactiveNR_MeasBeamReport_r16 = new(MeasAndMobParametersFRX_Diff_idleInactiveNR_MeasBeamReport_r16)
				if err = ie.idleInactiveNR_MeasBeamReport_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode idleInactiveNR_MeasBeamReport_r16", err)
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

			increasedNumberofCSIRSPerMO_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode increasedNumberofCSIRSPerMO_r16 optional
			if increasedNumberofCSIRSPerMO_r16Present {
				ie.increasedNumberofCSIRSPerMO_r16 = new(MeasAndMobParametersFRX_Diff_increasedNumberofCSIRSPerMO_r16)
				if err = ie.increasedNumberofCSIRSPerMO_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode increasedNumberofCSIRSPerMO_r16", err)
				}
			}
		}
	}
	return nil
}
