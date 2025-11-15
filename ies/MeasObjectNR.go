package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasObjectNR struct {
	ssbFrequency                    *ARFCN_ValueNR                    `optional`
	ssbSubcarrierSpacing            *SubcarrierSpacing                `optional`
	smtc1                           *SSB_MTC                          `optional`
	smtc2                           *SSB_MTC2                         `optional`
	refFreqCSI_RS                   *ARFCN_ValueNR                    `optional`
	referenceSignalConfig           ReferenceSignalConfig             `madatory`
	absThreshSS_BlocksConsolidation *ThresholdNR                      `optional`
	absThreshCSI_RS_Consolidation   *ThresholdNR                      `optional`
	nrofSS_BlocksToAverage          *int64                            `lb:2,ub:maxNrofSS_BlocksToAverage,optional`
	nrofCSI_RS_ResourcesToAverage   *int64                            `lb:2,ub:maxNrofCSI_RS_ResourcesToAverage,optional`
	quantityConfigIndex             int64                             `lb:1,ub:maxNrofQuantityConfig,madatory`
	offsetMO                        Q_OffsetRangeList                 `madatory`
	cellsToRemoveList               *PCI_List                         `optional`
	cellsToAddModList               *CellsToAddModList                `optional`
	excludedCellsToRemoveList       *PCI_RangeIndexList               `optional`
	excludedCellsToAddModList       []PCI_RangeElement                `lb:1,ub:maxNrofPCI_Ranges,optional`
	allowedCellsToRemoveList        *PCI_RangeIndexList               `optional`
	allowedCellsToAddModList        []PCI_RangeElement                `lb:1,ub:maxNrofPCI_Ranges,optional`
	freqBandIndicatorNR             *FreqBandIndicatorNR              `optional,ext-1`
	measCycleSCell                  *MeasObjectNR_measCycleSCell      `optional,ext-1`
	smtc3list_r16                   *SSB_MTC3List_r16                 `optional,ext-2`
	rmtc_Config_r16                 *RMTC_Config_r16                  `optional,ext-2,setuprelease`
	t312_r16                        *T312_r16                         `optional,ext-2,setuprelease`
	associatedMeasGapSSB_r17        *MeasGapId_r17                    `optional,ext-3`
	associatedMeasGapCSIRS_r17      *MeasGapId_r17                    `optional,ext-3`
	smtc4list_r17                   *SSB_MTC4List_r17                 `optional,ext-3`
	measCyclePSCell_r17             *MeasObjectNR_measCyclePSCell_r17 `optional,ext-3`
	cellsToAddModListExt_v1710      *CellsToAddModListExt_v1710       `optional,ext-3`
	associatedMeasGapSSB2_v1720     *MeasGapId_r17                    `optional,ext-4`
	associatedMeasGapCSIRS2_v1720   *MeasGapId_r17                    `optional,ext-4`
}

func (ie *MeasObjectNR) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.freqBandIndicatorNR != nil || ie.measCycleSCell != nil || ie.smtc3list_r16 != nil || ie.rmtc_Config_r16 != nil || ie.t312_r16 != nil || ie.associatedMeasGapSSB_r17 != nil || ie.associatedMeasGapCSIRS_r17 != nil || ie.smtc4list_r17 != nil || ie.measCyclePSCell_r17 != nil || ie.cellsToAddModListExt_v1710 != nil || ie.associatedMeasGapSSB2_v1720 != nil || ie.associatedMeasGapCSIRS2_v1720 != nil
	preambleBits := []bool{hasExtensions, ie.ssbFrequency != nil, ie.ssbSubcarrierSpacing != nil, ie.smtc1 != nil, ie.smtc2 != nil, ie.refFreqCSI_RS != nil, ie.absThreshSS_BlocksConsolidation != nil, ie.absThreshCSI_RS_Consolidation != nil, ie.nrofSS_BlocksToAverage != nil, ie.nrofCSI_RS_ResourcesToAverage != nil, ie.cellsToRemoveList != nil, ie.cellsToAddModList != nil, ie.excludedCellsToRemoveList != nil, len(ie.excludedCellsToAddModList) > 0, ie.allowedCellsToRemoveList != nil, len(ie.allowedCellsToAddModList) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ssbFrequency != nil {
		if err = ie.ssbFrequency.Encode(w); err != nil {
			return utils.WrapError("Encode ssbFrequency", err)
		}
	}
	if ie.ssbSubcarrierSpacing != nil {
		if err = ie.ssbSubcarrierSpacing.Encode(w); err != nil {
			return utils.WrapError("Encode ssbSubcarrierSpacing", err)
		}
	}
	if ie.smtc1 != nil {
		if err = ie.smtc1.Encode(w); err != nil {
			return utils.WrapError("Encode smtc1", err)
		}
	}
	if ie.smtc2 != nil {
		if err = ie.smtc2.Encode(w); err != nil {
			return utils.WrapError("Encode smtc2", err)
		}
	}
	if ie.refFreqCSI_RS != nil {
		if err = ie.refFreqCSI_RS.Encode(w); err != nil {
			return utils.WrapError("Encode refFreqCSI_RS", err)
		}
	}
	if err = ie.referenceSignalConfig.Encode(w); err != nil {
		return utils.WrapError("Encode referenceSignalConfig", err)
	}
	if ie.absThreshSS_BlocksConsolidation != nil {
		if err = ie.absThreshSS_BlocksConsolidation.Encode(w); err != nil {
			return utils.WrapError("Encode absThreshSS_BlocksConsolidation", err)
		}
	}
	if ie.absThreshCSI_RS_Consolidation != nil {
		if err = ie.absThreshCSI_RS_Consolidation.Encode(w); err != nil {
			return utils.WrapError("Encode absThreshCSI_RS_Consolidation", err)
		}
	}
	if ie.nrofSS_BlocksToAverage != nil {
		if err = w.WriteInteger(*ie.nrofSS_BlocksToAverage, &uper.Constraint{Lb: 2, Ub: maxNrofSS_BlocksToAverage}, false); err != nil {
			return utils.WrapError("Encode nrofSS_BlocksToAverage", err)
		}
	}
	if ie.nrofCSI_RS_ResourcesToAverage != nil {
		if err = w.WriteInteger(*ie.nrofCSI_RS_ResourcesToAverage, &uper.Constraint{Lb: 2, Ub: maxNrofCSI_RS_ResourcesToAverage}, false); err != nil {
			return utils.WrapError("Encode nrofCSI_RS_ResourcesToAverage", err)
		}
	}
	if err = w.WriteInteger(ie.quantityConfigIndex, &uper.Constraint{Lb: 1, Ub: maxNrofQuantityConfig}, false); err != nil {
		return utils.WrapError("WriteInteger quantityConfigIndex", err)
	}
	if err = ie.offsetMO.Encode(w); err != nil {
		return utils.WrapError("Encode offsetMO", err)
	}
	if ie.cellsToRemoveList != nil {
		if err = ie.cellsToRemoveList.Encode(w); err != nil {
			return utils.WrapError("Encode cellsToRemoveList", err)
		}
	}
	if ie.cellsToAddModList != nil {
		if err = ie.cellsToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode cellsToAddModList", err)
		}
	}
	if ie.excludedCellsToRemoveList != nil {
		if err = ie.excludedCellsToRemoveList.Encode(w); err != nil {
			return utils.WrapError("Encode excludedCellsToRemoveList", err)
		}
	}
	if len(ie.excludedCellsToAddModList) > 0 {
		tmp_excludedCellsToAddModList := utils.NewSequence[*PCI_RangeElement]([]*PCI_RangeElement{}, uper.Constraint{Lb: 1, Ub: maxNrofPCI_Ranges}, false)
		for _, i := range ie.excludedCellsToAddModList {
			tmp_excludedCellsToAddModList.Value = append(tmp_excludedCellsToAddModList.Value, &i)
		}
		if err = tmp_excludedCellsToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode excludedCellsToAddModList", err)
		}
	}
	if ie.allowedCellsToRemoveList != nil {
		if err = ie.allowedCellsToRemoveList.Encode(w); err != nil {
			return utils.WrapError("Encode allowedCellsToRemoveList", err)
		}
	}
	if len(ie.allowedCellsToAddModList) > 0 {
		tmp_allowedCellsToAddModList := utils.NewSequence[*PCI_RangeElement]([]*PCI_RangeElement{}, uper.Constraint{Lb: 1, Ub: maxNrofPCI_Ranges}, false)
		for _, i := range ie.allowedCellsToAddModList {
			tmp_allowedCellsToAddModList.Value = append(tmp_allowedCellsToAddModList.Value, &i)
		}
		if err = tmp_allowedCellsToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode allowedCellsToAddModList", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 4 bits for 4 extension groups
		extBitmap := []bool{ie.freqBandIndicatorNR != nil || ie.measCycleSCell != nil, ie.smtc3list_r16 != nil || ie.rmtc_Config_r16 != nil || ie.t312_r16 != nil, ie.associatedMeasGapSSB_r17 != nil || ie.associatedMeasGapCSIRS_r17 != nil || ie.smtc4list_r17 != nil || ie.measCyclePSCell_r17 != nil || ie.cellsToAddModListExt_v1710 != nil, ie.associatedMeasGapSSB2_v1720 != nil || ie.associatedMeasGapCSIRS2_v1720 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MeasObjectNR", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.freqBandIndicatorNR != nil, ie.measCycleSCell != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode freqBandIndicatorNR optional
			if ie.freqBandIndicatorNR != nil {
				if err = ie.freqBandIndicatorNR.Encode(extWriter); err != nil {
					return utils.WrapError("Encode freqBandIndicatorNR", err)
				}
			}
			// encode measCycleSCell optional
			if ie.measCycleSCell != nil {
				if err = ie.measCycleSCell.Encode(extWriter); err != nil {
					return utils.WrapError("Encode measCycleSCell", err)
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
			optionals_ext_2 := []bool{ie.smtc3list_r16 != nil, ie.rmtc_Config_r16 != nil, ie.t312_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode smtc3list_r16 optional
			if ie.smtc3list_r16 != nil {
				if err = ie.smtc3list_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode smtc3list_r16", err)
				}
			}
			// encode rmtc_Config_r16 optional
			if ie.rmtc_Config_r16 != nil {
				tmp_rmtc_Config_r16 := utils.SetupRelease[*RMTC_Config_r16]{
					Setup: ie.rmtc_Config_r16,
				}
				if err = tmp_rmtc_Config_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rmtc_Config_r16", err)
				}
			}
			// encode t312_r16 optional
			if ie.t312_r16 != nil {
				tmp_t312_r16 := utils.SetupRelease[*T312_r16]{
					Setup: ie.t312_r16,
				}
				if err = tmp_t312_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode t312_r16", err)
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
			optionals_ext_3 := []bool{ie.associatedMeasGapSSB_r17 != nil, ie.associatedMeasGapCSIRS_r17 != nil, ie.smtc4list_r17 != nil, ie.measCyclePSCell_r17 != nil, ie.cellsToAddModListExt_v1710 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode associatedMeasGapSSB_r17 optional
			if ie.associatedMeasGapSSB_r17 != nil {
				if err = ie.associatedMeasGapSSB_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode associatedMeasGapSSB_r17", err)
				}
			}
			// encode associatedMeasGapCSIRS_r17 optional
			if ie.associatedMeasGapCSIRS_r17 != nil {
				if err = ie.associatedMeasGapCSIRS_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode associatedMeasGapCSIRS_r17", err)
				}
			}
			// encode smtc4list_r17 optional
			if ie.smtc4list_r17 != nil {
				if err = ie.smtc4list_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode smtc4list_r17", err)
				}
			}
			// encode measCyclePSCell_r17 optional
			if ie.measCyclePSCell_r17 != nil {
				if err = ie.measCyclePSCell_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode measCyclePSCell_r17", err)
				}
			}
			// encode cellsToAddModListExt_v1710 optional
			if ie.cellsToAddModListExt_v1710 != nil {
				if err = ie.cellsToAddModListExt_v1710.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cellsToAddModListExt_v1710", err)
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
			optionals_ext_4 := []bool{ie.associatedMeasGapSSB2_v1720 != nil, ie.associatedMeasGapCSIRS2_v1720 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode associatedMeasGapSSB2_v1720 optional
			if ie.associatedMeasGapSSB2_v1720 != nil {
				if err = ie.associatedMeasGapSSB2_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode associatedMeasGapSSB2_v1720", err)
				}
			}
			// encode associatedMeasGapCSIRS2_v1720 optional
			if ie.associatedMeasGapCSIRS2_v1720 != nil {
				if err = ie.associatedMeasGapCSIRS2_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode associatedMeasGapCSIRS2_v1720", err)
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

func (ie *MeasObjectNR) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ssbFrequencyPresent bool
	if ssbFrequencyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ssbSubcarrierSpacingPresent bool
	if ssbSubcarrierSpacingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var smtc1Present bool
	if smtc1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var smtc2Present bool
	if smtc2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var refFreqCSI_RSPresent bool
	if refFreqCSI_RSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var absThreshSS_BlocksConsolidationPresent bool
	if absThreshSS_BlocksConsolidationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var absThreshCSI_RS_ConsolidationPresent bool
	if absThreshCSI_RS_ConsolidationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nrofSS_BlocksToAveragePresent bool
	if nrofSS_BlocksToAveragePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nrofCSI_RS_ResourcesToAveragePresent bool
	if nrofCSI_RS_ResourcesToAveragePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var cellsToRemoveListPresent bool
	if cellsToRemoveListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var cellsToAddModListPresent bool
	if cellsToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var excludedCellsToRemoveListPresent bool
	if excludedCellsToRemoveListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var excludedCellsToAddModListPresent bool
	if excludedCellsToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var allowedCellsToRemoveListPresent bool
	if allowedCellsToRemoveListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var allowedCellsToAddModListPresent bool
	if allowedCellsToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ssbFrequencyPresent {
		ie.ssbFrequency = new(ARFCN_ValueNR)
		if err = ie.ssbFrequency.Decode(r); err != nil {
			return utils.WrapError("Decode ssbFrequency", err)
		}
	}
	if ssbSubcarrierSpacingPresent {
		ie.ssbSubcarrierSpacing = new(SubcarrierSpacing)
		if err = ie.ssbSubcarrierSpacing.Decode(r); err != nil {
			return utils.WrapError("Decode ssbSubcarrierSpacing", err)
		}
	}
	if smtc1Present {
		ie.smtc1 = new(SSB_MTC)
		if err = ie.smtc1.Decode(r); err != nil {
			return utils.WrapError("Decode smtc1", err)
		}
	}
	if smtc2Present {
		ie.smtc2 = new(SSB_MTC2)
		if err = ie.smtc2.Decode(r); err != nil {
			return utils.WrapError("Decode smtc2", err)
		}
	}
	if refFreqCSI_RSPresent {
		ie.refFreqCSI_RS = new(ARFCN_ValueNR)
		if err = ie.refFreqCSI_RS.Decode(r); err != nil {
			return utils.WrapError("Decode refFreqCSI_RS", err)
		}
	}
	if err = ie.referenceSignalConfig.Decode(r); err != nil {
		return utils.WrapError("Decode referenceSignalConfig", err)
	}
	if absThreshSS_BlocksConsolidationPresent {
		ie.absThreshSS_BlocksConsolidation = new(ThresholdNR)
		if err = ie.absThreshSS_BlocksConsolidation.Decode(r); err != nil {
			return utils.WrapError("Decode absThreshSS_BlocksConsolidation", err)
		}
	}
	if absThreshCSI_RS_ConsolidationPresent {
		ie.absThreshCSI_RS_Consolidation = new(ThresholdNR)
		if err = ie.absThreshCSI_RS_Consolidation.Decode(r); err != nil {
			return utils.WrapError("Decode absThreshCSI_RS_Consolidation", err)
		}
	}
	if nrofSS_BlocksToAveragePresent {
		var tmp_int_nrofSS_BlocksToAverage int64
		if tmp_int_nrofSS_BlocksToAverage, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: maxNrofSS_BlocksToAverage}, false); err != nil {
			return utils.WrapError("Decode nrofSS_BlocksToAverage", err)
		}
		ie.nrofSS_BlocksToAverage = &tmp_int_nrofSS_BlocksToAverage
	}
	if nrofCSI_RS_ResourcesToAveragePresent {
		var tmp_int_nrofCSI_RS_ResourcesToAverage int64
		if tmp_int_nrofCSI_RS_ResourcesToAverage, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: maxNrofCSI_RS_ResourcesToAverage}, false); err != nil {
			return utils.WrapError("Decode nrofCSI_RS_ResourcesToAverage", err)
		}
		ie.nrofCSI_RS_ResourcesToAverage = &tmp_int_nrofCSI_RS_ResourcesToAverage
	}
	var tmp_int_quantityConfigIndex int64
	if tmp_int_quantityConfigIndex, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofQuantityConfig}, false); err != nil {
		return utils.WrapError("ReadInteger quantityConfigIndex", err)
	}
	ie.quantityConfigIndex = tmp_int_quantityConfigIndex
	if err = ie.offsetMO.Decode(r); err != nil {
		return utils.WrapError("Decode offsetMO", err)
	}
	if cellsToRemoveListPresent {
		ie.cellsToRemoveList = new(PCI_List)
		if err = ie.cellsToRemoveList.Decode(r); err != nil {
			return utils.WrapError("Decode cellsToRemoveList", err)
		}
	}
	if cellsToAddModListPresent {
		ie.cellsToAddModList = new(CellsToAddModList)
		if err = ie.cellsToAddModList.Decode(r); err != nil {
			return utils.WrapError("Decode cellsToAddModList", err)
		}
	}
	if excludedCellsToRemoveListPresent {
		ie.excludedCellsToRemoveList = new(PCI_RangeIndexList)
		if err = ie.excludedCellsToRemoveList.Decode(r); err != nil {
			return utils.WrapError("Decode excludedCellsToRemoveList", err)
		}
	}
	if excludedCellsToAddModListPresent {
		tmp_excludedCellsToAddModList := utils.NewSequence[*PCI_RangeElement]([]*PCI_RangeElement{}, uper.Constraint{Lb: 1, Ub: maxNrofPCI_Ranges}, false)
		fn_excludedCellsToAddModList := func() *PCI_RangeElement {
			return new(PCI_RangeElement)
		}
		if err = tmp_excludedCellsToAddModList.Decode(r, fn_excludedCellsToAddModList); err != nil {
			return utils.WrapError("Decode excludedCellsToAddModList", err)
		}
		ie.excludedCellsToAddModList = []PCI_RangeElement{}
		for _, i := range tmp_excludedCellsToAddModList.Value {
			ie.excludedCellsToAddModList = append(ie.excludedCellsToAddModList, *i)
		}
	}
	if allowedCellsToRemoveListPresent {
		ie.allowedCellsToRemoveList = new(PCI_RangeIndexList)
		if err = ie.allowedCellsToRemoveList.Decode(r); err != nil {
			return utils.WrapError("Decode allowedCellsToRemoveList", err)
		}
	}
	if allowedCellsToAddModListPresent {
		tmp_allowedCellsToAddModList := utils.NewSequence[*PCI_RangeElement]([]*PCI_RangeElement{}, uper.Constraint{Lb: 1, Ub: maxNrofPCI_Ranges}, false)
		fn_allowedCellsToAddModList := func() *PCI_RangeElement {
			return new(PCI_RangeElement)
		}
		if err = tmp_allowedCellsToAddModList.Decode(r, fn_allowedCellsToAddModList); err != nil {
			return utils.WrapError("Decode allowedCellsToAddModList", err)
		}
		ie.allowedCellsToAddModList = []PCI_RangeElement{}
		for _, i := range tmp_allowedCellsToAddModList.Value {
			ie.allowedCellsToAddModList = append(ie.allowedCellsToAddModList, *i)
		}
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

			freqBandIndicatorNRPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			measCycleSCellPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode freqBandIndicatorNR optional
			if freqBandIndicatorNRPresent {
				ie.freqBandIndicatorNR = new(FreqBandIndicatorNR)
				if err = ie.freqBandIndicatorNR.Decode(extReader); err != nil {
					return utils.WrapError("Decode freqBandIndicatorNR", err)
				}
			}
			// decode measCycleSCell optional
			if measCycleSCellPresent {
				ie.measCycleSCell = new(MeasObjectNR_measCycleSCell)
				if err = ie.measCycleSCell.Decode(extReader); err != nil {
					return utils.WrapError("Decode measCycleSCell", err)
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

			smtc3list_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			rmtc_Config_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			t312_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode smtc3list_r16 optional
			if smtc3list_r16Present {
				ie.smtc3list_r16 = new(SSB_MTC3List_r16)
				if err = ie.smtc3list_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode smtc3list_r16", err)
				}
			}
			// decode rmtc_Config_r16 optional
			if rmtc_Config_r16Present {
				tmp_rmtc_Config_r16 := utils.SetupRelease[*RMTC_Config_r16]{}
				if err = tmp_rmtc_Config_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode rmtc_Config_r16", err)
				}
				ie.rmtc_Config_r16 = tmp_rmtc_Config_r16.Setup
			}
			// decode t312_r16 optional
			if t312_r16Present {
				tmp_t312_r16 := utils.SetupRelease[*T312_r16]{}
				if err = tmp_t312_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode t312_r16", err)
				}
				ie.t312_r16 = tmp_t312_r16.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			associatedMeasGapSSB_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			associatedMeasGapCSIRS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			smtc4list_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			measCyclePSCell_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cellsToAddModListExt_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode associatedMeasGapSSB_r17 optional
			if associatedMeasGapSSB_r17Present {
				ie.associatedMeasGapSSB_r17 = new(MeasGapId_r17)
				if err = ie.associatedMeasGapSSB_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode associatedMeasGapSSB_r17", err)
				}
			}
			// decode associatedMeasGapCSIRS_r17 optional
			if associatedMeasGapCSIRS_r17Present {
				ie.associatedMeasGapCSIRS_r17 = new(MeasGapId_r17)
				if err = ie.associatedMeasGapCSIRS_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode associatedMeasGapCSIRS_r17", err)
				}
			}
			// decode smtc4list_r17 optional
			if smtc4list_r17Present {
				ie.smtc4list_r17 = new(SSB_MTC4List_r17)
				if err = ie.smtc4list_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode smtc4list_r17", err)
				}
			}
			// decode measCyclePSCell_r17 optional
			if measCyclePSCell_r17Present {
				ie.measCyclePSCell_r17 = new(MeasObjectNR_measCyclePSCell_r17)
				if err = ie.measCyclePSCell_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode measCyclePSCell_r17", err)
				}
			}
			// decode cellsToAddModListExt_v1710 optional
			if cellsToAddModListExt_v1710Present {
				ie.cellsToAddModListExt_v1710 = new(CellsToAddModListExt_v1710)
				if err = ie.cellsToAddModListExt_v1710.Decode(extReader); err != nil {
					return utils.WrapError("Decode cellsToAddModListExt_v1710", err)
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

			associatedMeasGapSSB2_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			associatedMeasGapCSIRS2_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode associatedMeasGapSSB2_v1720 optional
			if associatedMeasGapSSB2_v1720Present {
				ie.associatedMeasGapSSB2_v1720 = new(MeasGapId_r17)
				if err = ie.associatedMeasGapSSB2_v1720.Decode(extReader); err != nil {
					return utils.WrapError("Decode associatedMeasGapSSB2_v1720", err)
				}
			}
			// decode associatedMeasGapCSIRS2_v1720 optional
			if associatedMeasGapCSIRS2_v1720Present {
				ie.associatedMeasGapCSIRS2_v1720 = new(MeasGapId_r17)
				if err = ie.associatedMeasGapCSIRS2_v1720.Decode(extReader); err != nil {
					return utils.WrapError("Decode associatedMeasGapCSIRS2_v1720", err)
				}
			}
		}
	}
	return nil
}
