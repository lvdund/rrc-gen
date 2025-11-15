package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConfigRestrictInfoSCG struct {
	allowedBC_ListMRDC                     *BandCombinationInfoList                         `optional`
	powerCoordination_FR1                  *ConfigRestrictInfoSCG_powerCoordination_FR1     `optional`
	servCellIndexRangeSCG                  *ConfigRestrictInfoSCG_servCellIndexRangeSCG     `optional`
	maxMeasFreqsSCG                        *int64                                           `lb:1,ub:maxMeasFreqsMN,optional`
	dummy                                  *int64                                           `lb:1,ub:maxMeasIdentitiesMN,optional`
	selectedBandEntriesMNList              []SelectedBandEntriesMN                          `lb:1,ub:maxBandComb,optional,ext-1`
	pdcch_BlindDetectionSCG                *int64                                           `lb:1,ub:15,optional,ext-1`
	maxNumberROHC_ContextSessionsSN        *int64                                           `lb:0,ub:16384,optional,ext-1`
	maxIntraFreqMeasIdentitiesSCG          *int64                                           `lb:1,ub:maxMeasIdentitiesMN,optional,ext-2`
	maxInterFreqMeasIdentitiesSCG          *int64                                           `lb:1,ub:maxMeasIdentitiesMN,optional,ext-2`
	p_maxNR_FR1_MCG_r16                    *P_Max                                           `optional,ext-3`
	powerCoordination_FR2_r16              *ConfigRestrictInfoSCG_powerCoordination_FR2_r16 `optional,ext-3`
	nrdc_PC_mode_FR1_r16                   *ConfigRestrictInfoSCG_nrdc_PC_mode_FR1_r16      `optional,ext-3`
	nrdc_PC_mode_FR2_r16                   *ConfigRestrictInfoSCG_nrdc_PC_mode_FR2_r16      `optional,ext-3`
	maxMeasSRS_ResourceSCG_r16             *int64                                           `lb:0,ub:maxNrofCLI_SRS_Resources_r16,optional,ext-3`
	maxMeasCLI_ResourceSCG_r16             *int64                                           `lb:0,ub:maxNrofCLI_RSSI_Resources_r16,optional,ext-3`
	maxNumberEHC_ContextsSN_r16            *int64                                           `lb:0,ub:65536,optional,ext-3`
	allowedReducedConfigForOverheating_r16 *OverheatingAssistance                           `optional,ext-3`
	maxToffset_r16                         *T_Offset_r16                                    `optional,ext-3`
	allowedReducedConfigForOverheating_r17 *OverheatingAssistance_r17                       `optional,ext-4`
	maxNumberUDC_DRB_r17                   *int64                                           `lb:0,ub:2,optional,ext-4`
	maxNumberCPCCandidates_r17             *int64                                           `lb:0,ub:maxNrofCondCells_1_r17,optional,ext-4`
}

func (ie *ConfigRestrictInfoSCG) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := len(ie.selectedBandEntriesMNList) > 0 || ie.pdcch_BlindDetectionSCG != nil || ie.maxNumberROHC_ContextSessionsSN != nil || ie.maxIntraFreqMeasIdentitiesSCG != nil || ie.maxInterFreqMeasIdentitiesSCG != nil || ie.p_maxNR_FR1_MCG_r16 != nil || ie.powerCoordination_FR2_r16 != nil || ie.nrdc_PC_mode_FR1_r16 != nil || ie.nrdc_PC_mode_FR2_r16 != nil || ie.maxMeasSRS_ResourceSCG_r16 != nil || ie.maxMeasCLI_ResourceSCG_r16 != nil || ie.maxNumberEHC_ContextsSN_r16 != nil || ie.allowedReducedConfigForOverheating_r16 != nil || ie.maxToffset_r16 != nil || ie.allowedReducedConfigForOverheating_r17 != nil || ie.maxNumberUDC_DRB_r17 != nil || ie.maxNumberCPCCandidates_r17 != nil
	preambleBits := []bool{hasExtensions, ie.allowedBC_ListMRDC != nil, ie.powerCoordination_FR1 != nil, ie.servCellIndexRangeSCG != nil, ie.maxMeasFreqsSCG != nil, ie.dummy != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.allowedBC_ListMRDC != nil {
		if err = ie.allowedBC_ListMRDC.Encode(w); err != nil {
			return utils.WrapError("Encode allowedBC_ListMRDC", err)
		}
	}
	if ie.powerCoordination_FR1 != nil {
		if err = ie.powerCoordination_FR1.Encode(w); err != nil {
			return utils.WrapError("Encode powerCoordination_FR1", err)
		}
	}
	if ie.servCellIndexRangeSCG != nil {
		if err = ie.servCellIndexRangeSCG.Encode(w); err != nil {
			return utils.WrapError("Encode servCellIndexRangeSCG", err)
		}
	}
	if ie.maxMeasFreqsSCG != nil {
		if err = w.WriteInteger(*ie.maxMeasFreqsSCG, &uper.Constraint{Lb: 1, Ub: maxMeasFreqsMN}, false); err != nil {
			return utils.WrapError("Encode maxMeasFreqsSCG", err)
		}
	}
	if ie.dummy != nil {
		if err = w.WriteInteger(*ie.dummy, &uper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
			return utils.WrapError("Encode dummy", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 4 bits for 4 extension groups
		extBitmap := []bool{len(ie.selectedBandEntriesMNList) > 0 || ie.pdcch_BlindDetectionSCG != nil || ie.maxNumberROHC_ContextSessionsSN != nil, ie.maxIntraFreqMeasIdentitiesSCG != nil || ie.maxInterFreqMeasIdentitiesSCG != nil, ie.p_maxNR_FR1_MCG_r16 != nil || ie.powerCoordination_FR2_r16 != nil || ie.nrdc_PC_mode_FR1_r16 != nil || ie.nrdc_PC_mode_FR2_r16 != nil || ie.maxMeasSRS_ResourceSCG_r16 != nil || ie.maxMeasCLI_ResourceSCG_r16 != nil || ie.maxNumberEHC_ContextsSN_r16 != nil || ie.allowedReducedConfigForOverheating_r16 != nil || ie.maxToffset_r16 != nil, ie.allowedReducedConfigForOverheating_r17 != nil || ie.maxNumberUDC_DRB_r17 != nil || ie.maxNumberCPCCandidates_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap ConfigRestrictInfoSCG", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.selectedBandEntriesMNList) > 0, ie.pdcch_BlindDetectionSCG != nil, ie.maxNumberROHC_ContextSessionsSN != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode selectedBandEntriesMNList optional
			if len(ie.selectedBandEntriesMNList) > 0 {
				tmp_selectedBandEntriesMNList := utils.NewSequence[*SelectedBandEntriesMN]([]*SelectedBandEntriesMN{}, uper.Constraint{Lb: 1, Ub: maxBandComb}, false)
				for _, i := range ie.selectedBandEntriesMNList {
					tmp_selectedBandEntriesMNList.Value = append(tmp_selectedBandEntriesMNList.Value, &i)
				}
				if err = tmp_selectedBandEntriesMNList.Encode(extWriter); err != nil {
					return utils.WrapError("Encode selectedBandEntriesMNList", err)
				}
			}
			// encode pdcch_BlindDetectionSCG optional
			if ie.pdcch_BlindDetectionSCG != nil {
				if err = extWriter.WriteInteger(*ie.pdcch_BlindDetectionSCG, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
					return utils.WrapError("Encode pdcch_BlindDetectionSCG", err)
				}
			}
			// encode maxNumberROHC_ContextSessionsSN optional
			if ie.maxNumberROHC_ContextSessionsSN != nil {
				if err = extWriter.WriteInteger(*ie.maxNumberROHC_ContextSessionsSN, &uper.Constraint{Lb: 0, Ub: 16384}, false); err != nil {
					return utils.WrapError("Encode maxNumberROHC_ContextSessionsSN", err)
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
			optionals_ext_2 := []bool{ie.maxIntraFreqMeasIdentitiesSCG != nil, ie.maxInterFreqMeasIdentitiesSCG != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode maxIntraFreqMeasIdentitiesSCG optional
			if ie.maxIntraFreqMeasIdentitiesSCG != nil {
				if err = extWriter.WriteInteger(*ie.maxIntraFreqMeasIdentitiesSCG, &uper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
					return utils.WrapError("Encode maxIntraFreqMeasIdentitiesSCG", err)
				}
			}
			// encode maxInterFreqMeasIdentitiesSCG optional
			if ie.maxInterFreqMeasIdentitiesSCG != nil {
				if err = extWriter.WriteInteger(*ie.maxInterFreqMeasIdentitiesSCG, &uper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
					return utils.WrapError("Encode maxInterFreqMeasIdentitiesSCG", err)
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
			optionals_ext_3 := []bool{ie.p_maxNR_FR1_MCG_r16 != nil, ie.powerCoordination_FR2_r16 != nil, ie.nrdc_PC_mode_FR1_r16 != nil, ie.nrdc_PC_mode_FR2_r16 != nil, ie.maxMeasSRS_ResourceSCG_r16 != nil, ie.maxMeasCLI_ResourceSCG_r16 != nil, ie.maxNumberEHC_ContextsSN_r16 != nil, ie.allowedReducedConfigForOverheating_r16 != nil, ie.maxToffset_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode p_maxNR_FR1_MCG_r16 optional
			if ie.p_maxNR_FR1_MCG_r16 != nil {
				if err = ie.p_maxNR_FR1_MCG_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode p_maxNR_FR1_MCG_r16", err)
				}
			}
			// encode powerCoordination_FR2_r16 optional
			if ie.powerCoordination_FR2_r16 != nil {
				if err = ie.powerCoordination_FR2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode powerCoordination_FR2_r16", err)
				}
			}
			// encode nrdc_PC_mode_FR1_r16 optional
			if ie.nrdc_PC_mode_FR1_r16 != nil {
				if err = ie.nrdc_PC_mode_FR1_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nrdc_PC_mode_FR1_r16", err)
				}
			}
			// encode nrdc_PC_mode_FR2_r16 optional
			if ie.nrdc_PC_mode_FR2_r16 != nil {
				if err = ie.nrdc_PC_mode_FR2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nrdc_PC_mode_FR2_r16", err)
				}
			}
			// encode maxMeasSRS_ResourceSCG_r16 optional
			if ie.maxMeasSRS_ResourceSCG_r16 != nil {
				if err = extWriter.WriteInteger(*ie.maxMeasSRS_ResourceSCG_r16, &uper.Constraint{Lb: 0, Ub: maxNrofCLI_SRS_Resources_r16}, false); err != nil {
					return utils.WrapError("Encode maxMeasSRS_ResourceSCG_r16", err)
				}
			}
			// encode maxMeasCLI_ResourceSCG_r16 optional
			if ie.maxMeasCLI_ResourceSCG_r16 != nil {
				if err = extWriter.WriteInteger(*ie.maxMeasCLI_ResourceSCG_r16, &uper.Constraint{Lb: 0, Ub: maxNrofCLI_RSSI_Resources_r16}, false); err != nil {
					return utils.WrapError("Encode maxMeasCLI_ResourceSCG_r16", err)
				}
			}
			// encode maxNumberEHC_ContextsSN_r16 optional
			if ie.maxNumberEHC_ContextsSN_r16 != nil {
				if err = extWriter.WriteInteger(*ie.maxNumberEHC_ContextsSN_r16, &uper.Constraint{Lb: 0, Ub: 65536}, false); err != nil {
					return utils.WrapError("Encode maxNumberEHC_ContextsSN_r16", err)
				}
			}
			// encode allowedReducedConfigForOverheating_r16 optional
			if ie.allowedReducedConfigForOverheating_r16 != nil {
				if err = ie.allowedReducedConfigForOverheating_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode allowedReducedConfigForOverheating_r16", err)
				}
			}
			// encode maxToffset_r16 optional
			if ie.maxToffset_r16 != nil {
				if err = ie.maxToffset_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxToffset_r16", err)
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
			optionals_ext_4 := []bool{ie.allowedReducedConfigForOverheating_r17 != nil, ie.maxNumberUDC_DRB_r17 != nil, ie.maxNumberCPCCandidates_r17 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode allowedReducedConfigForOverheating_r17 optional
			if ie.allowedReducedConfigForOverheating_r17 != nil {
				if err = ie.allowedReducedConfigForOverheating_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode allowedReducedConfigForOverheating_r17", err)
				}
			}
			// encode maxNumberUDC_DRB_r17 optional
			if ie.maxNumberUDC_DRB_r17 != nil {
				if err = extWriter.WriteInteger(*ie.maxNumberUDC_DRB_r17, &uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
					return utils.WrapError("Encode maxNumberUDC_DRB_r17", err)
				}
			}
			// encode maxNumberCPCCandidates_r17 optional
			if ie.maxNumberCPCCandidates_r17 != nil {
				if err = extWriter.WriteInteger(*ie.maxNumberCPCCandidates_r17, &uper.Constraint{Lb: 0, Ub: maxNrofCondCells_1_r17}, false); err != nil {
					return utils.WrapError("Encode maxNumberCPCCandidates_r17", err)
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

func (ie *ConfigRestrictInfoSCG) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var allowedBC_ListMRDCPresent bool
	if allowedBC_ListMRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var powerCoordination_FR1Present bool
	if powerCoordination_FR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var servCellIndexRangeSCGPresent bool
	if servCellIndexRangeSCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var maxMeasFreqsSCGPresent bool
	if maxMeasFreqsSCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var dummyPresent bool
	if dummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if allowedBC_ListMRDCPresent {
		ie.allowedBC_ListMRDC = new(BandCombinationInfoList)
		if err = ie.allowedBC_ListMRDC.Decode(r); err != nil {
			return utils.WrapError("Decode allowedBC_ListMRDC", err)
		}
	}
	if powerCoordination_FR1Present {
		ie.powerCoordination_FR1 = new(ConfigRestrictInfoSCG_powerCoordination_FR1)
		if err = ie.powerCoordination_FR1.Decode(r); err != nil {
			return utils.WrapError("Decode powerCoordination_FR1", err)
		}
	}
	if servCellIndexRangeSCGPresent {
		ie.servCellIndexRangeSCG = new(ConfigRestrictInfoSCG_servCellIndexRangeSCG)
		if err = ie.servCellIndexRangeSCG.Decode(r); err != nil {
			return utils.WrapError("Decode servCellIndexRangeSCG", err)
		}
	}
	if maxMeasFreqsSCGPresent {
		var tmp_int_maxMeasFreqsSCG int64
		if tmp_int_maxMeasFreqsSCG, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxMeasFreqsMN}, false); err != nil {
			return utils.WrapError("Decode maxMeasFreqsSCG", err)
		}
		ie.maxMeasFreqsSCG = &tmp_int_maxMeasFreqsSCG
	}
	if dummyPresent {
		var tmp_int_dummy int64
		if tmp_int_dummy, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
			return utils.WrapError("Decode dummy", err)
		}
		ie.dummy = &tmp_int_dummy
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

			selectedBandEntriesMNListPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdcch_BlindDetectionSCGPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumberROHC_ContextSessionsSNPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode selectedBandEntriesMNList optional
			if selectedBandEntriesMNListPresent {
				tmp_selectedBandEntriesMNList := utils.NewSequence[*SelectedBandEntriesMN]([]*SelectedBandEntriesMN{}, uper.Constraint{Lb: 1, Ub: maxBandComb}, false)
				fn_selectedBandEntriesMNList := func() *SelectedBandEntriesMN {
					return new(SelectedBandEntriesMN)
				}
				if err = tmp_selectedBandEntriesMNList.Decode(extReader, fn_selectedBandEntriesMNList); err != nil {
					return utils.WrapError("Decode selectedBandEntriesMNList", err)
				}
				ie.selectedBandEntriesMNList = []SelectedBandEntriesMN{}
				for _, i := range tmp_selectedBandEntriesMNList.Value {
					ie.selectedBandEntriesMNList = append(ie.selectedBandEntriesMNList, *i)
				}
			}
			// decode pdcch_BlindDetectionSCG optional
			if pdcch_BlindDetectionSCGPresent {
				var tmp_int_pdcch_BlindDetectionSCG int64
				if tmp_int_pdcch_BlindDetectionSCG, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
					return utils.WrapError("Decode pdcch_BlindDetectionSCG", err)
				}
				ie.pdcch_BlindDetectionSCG = &tmp_int_pdcch_BlindDetectionSCG
			}
			// decode maxNumberROHC_ContextSessionsSN optional
			if maxNumberROHC_ContextSessionsSNPresent {
				var tmp_int_maxNumberROHC_ContextSessionsSN int64
				if tmp_int_maxNumberROHC_ContextSessionsSN, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 16384}, false); err != nil {
					return utils.WrapError("Decode maxNumberROHC_ContextSessionsSN", err)
				}
				ie.maxNumberROHC_ContextSessionsSN = &tmp_int_maxNumberROHC_ContextSessionsSN
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			maxIntraFreqMeasIdentitiesSCGPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxInterFreqMeasIdentitiesSCGPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode maxIntraFreqMeasIdentitiesSCG optional
			if maxIntraFreqMeasIdentitiesSCGPresent {
				var tmp_int_maxIntraFreqMeasIdentitiesSCG int64
				if tmp_int_maxIntraFreqMeasIdentitiesSCG, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
					return utils.WrapError("Decode maxIntraFreqMeasIdentitiesSCG", err)
				}
				ie.maxIntraFreqMeasIdentitiesSCG = &tmp_int_maxIntraFreqMeasIdentitiesSCG
			}
			// decode maxInterFreqMeasIdentitiesSCG optional
			if maxInterFreqMeasIdentitiesSCGPresent {
				var tmp_int_maxInterFreqMeasIdentitiesSCG int64
				if tmp_int_maxInterFreqMeasIdentitiesSCG, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
					return utils.WrapError("Decode maxInterFreqMeasIdentitiesSCG", err)
				}
				ie.maxInterFreqMeasIdentitiesSCG = &tmp_int_maxInterFreqMeasIdentitiesSCG
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			p_maxNR_FR1_MCG_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			powerCoordination_FR2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nrdc_PC_mode_FR1_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nrdc_PC_mode_FR2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxMeasSRS_ResourceSCG_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxMeasCLI_ResourceSCG_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumberEHC_ContextsSN_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			allowedReducedConfigForOverheating_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxToffset_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode p_maxNR_FR1_MCG_r16 optional
			if p_maxNR_FR1_MCG_r16Present {
				ie.p_maxNR_FR1_MCG_r16 = new(P_Max)
				if err = ie.p_maxNR_FR1_MCG_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode p_maxNR_FR1_MCG_r16", err)
				}
			}
			// decode powerCoordination_FR2_r16 optional
			if powerCoordination_FR2_r16Present {
				ie.powerCoordination_FR2_r16 = new(ConfigRestrictInfoSCG_powerCoordination_FR2_r16)
				if err = ie.powerCoordination_FR2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode powerCoordination_FR2_r16", err)
				}
			}
			// decode nrdc_PC_mode_FR1_r16 optional
			if nrdc_PC_mode_FR1_r16Present {
				ie.nrdc_PC_mode_FR1_r16 = new(ConfigRestrictInfoSCG_nrdc_PC_mode_FR1_r16)
				if err = ie.nrdc_PC_mode_FR1_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode nrdc_PC_mode_FR1_r16", err)
				}
			}
			// decode nrdc_PC_mode_FR2_r16 optional
			if nrdc_PC_mode_FR2_r16Present {
				ie.nrdc_PC_mode_FR2_r16 = new(ConfigRestrictInfoSCG_nrdc_PC_mode_FR2_r16)
				if err = ie.nrdc_PC_mode_FR2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode nrdc_PC_mode_FR2_r16", err)
				}
			}
			// decode maxMeasSRS_ResourceSCG_r16 optional
			if maxMeasSRS_ResourceSCG_r16Present {
				var tmp_int_maxMeasSRS_ResourceSCG_r16 int64
				if tmp_int_maxMeasSRS_ResourceSCG_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofCLI_SRS_Resources_r16}, false); err != nil {
					return utils.WrapError("Decode maxMeasSRS_ResourceSCG_r16", err)
				}
				ie.maxMeasSRS_ResourceSCG_r16 = &tmp_int_maxMeasSRS_ResourceSCG_r16
			}
			// decode maxMeasCLI_ResourceSCG_r16 optional
			if maxMeasCLI_ResourceSCG_r16Present {
				var tmp_int_maxMeasCLI_ResourceSCG_r16 int64
				if tmp_int_maxMeasCLI_ResourceSCG_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofCLI_RSSI_Resources_r16}, false); err != nil {
					return utils.WrapError("Decode maxMeasCLI_ResourceSCG_r16", err)
				}
				ie.maxMeasCLI_ResourceSCG_r16 = &tmp_int_maxMeasCLI_ResourceSCG_r16
			}
			// decode maxNumberEHC_ContextsSN_r16 optional
			if maxNumberEHC_ContextsSN_r16Present {
				var tmp_int_maxNumberEHC_ContextsSN_r16 int64
				if tmp_int_maxNumberEHC_ContextsSN_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65536}, false); err != nil {
					return utils.WrapError("Decode maxNumberEHC_ContextsSN_r16", err)
				}
				ie.maxNumberEHC_ContextsSN_r16 = &tmp_int_maxNumberEHC_ContextsSN_r16
			}
			// decode allowedReducedConfigForOverheating_r16 optional
			if allowedReducedConfigForOverheating_r16Present {
				ie.allowedReducedConfigForOverheating_r16 = new(OverheatingAssistance)
				if err = ie.allowedReducedConfigForOverheating_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode allowedReducedConfigForOverheating_r16", err)
				}
			}
			// decode maxToffset_r16 optional
			if maxToffset_r16Present {
				ie.maxToffset_r16 = new(T_Offset_r16)
				if err = ie.maxToffset_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxToffset_r16", err)
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

			allowedReducedConfigForOverheating_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumberUDC_DRB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumberCPCCandidates_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode allowedReducedConfigForOverheating_r17 optional
			if allowedReducedConfigForOverheating_r17Present {
				ie.allowedReducedConfigForOverheating_r17 = new(OverheatingAssistance_r17)
				if err = ie.allowedReducedConfigForOverheating_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode allowedReducedConfigForOverheating_r17", err)
				}
			}
			// decode maxNumberUDC_DRB_r17 optional
			if maxNumberUDC_DRB_r17Present {
				var tmp_int_maxNumberUDC_DRB_r17 int64
				if tmp_int_maxNumberUDC_DRB_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
					return utils.WrapError("Decode maxNumberUDC_DRB_r17", err)
				}
				ie.maxNumberUDC_DRB_r17 = &tmp_int_maxNumberUDC_DRB_r17
			}
			// decode maxNumberCPCCandidates_r17 optional
			if maxNumberCPCCandidates_r17Present {
				var tmp_int_maxNumberCPCCandidates_r17 int64
				if tmp_int_maxNumberCPCCandidates_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofCondCells_1_r17}, false); err != nil {
					return utils.WrapError("Decode maxNumberCPCCandidates_r17", err)
				}
				ie.maxNumberCPCCandidates_r17 = &tmp_int_maxNumberCPCCandidates_r17
			}
		}
	}
	return nil
}
