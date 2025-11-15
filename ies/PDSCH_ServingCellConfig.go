package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_ServingCellConfig struct {
	codeBlockGroupTransmission               *PDSCH_CodeBlockGroupTransmission                         `optional,setuprelease`
	xOverhead                                *PDSCH_ServingCellConfig_xOverhead                        `optional`
	nrofHARQ_ProcessesForPDSCH               *PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH       `optional`
	pucch_Cell                               *ServCellIndex                                            `optional`
	maxMIMO_Layers                           *int64                                                    `lb:1,ub:8,optional,ext-1`
	processingType2Enabled                   *bool                                                     `optional,ext-1`
	pdsch_CodeBlockGroupTransmissionList_r16 *PDSCH_CodeBlockGroupTransmissionList_r16                 `optional,ext-2,setuprelease`
	downlinkHARQ_FeedbackDisabled_r17        *DownlinkHARQ_FeedbackDisabled_r17                        `optional,ext-3,setuprelease`
	nrofHARQ_ProcessesForPDSCH_v1700         *PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_v1700 `optional,ext-3`
}

func (ie *PDSCH_ServingCellConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.maxMIMO_Layers != nil || ie.processingType2Enabled != nil || ie.pdsch_CodeBlockGroupTransmissionList_r16 != nil || ie.downlinkHARQ_FeedbackDisabled_r17 != nil || ie.nrofHARQ_ProcessesForPDSCH_v1700 != nil
	preambleBits := []bool{hasExtensions, ie.codeBlockGroupTransmission != nil, ie.xOverhead != nil, ie.nrofHARQ_ProcessesForPDSCH != nil, ie.pucch_Cell != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.codeBlockGroupTransmission != nil {
		tmp_codeBlockGroupTransmission := utils.SetupRelease[*PDSCH_CodeBlockGroupTransmission]{
			Setup: ie.codeBlockGroupTransmission,
		}
		if err = tmp_codeBlockGroupTransmission.Encode(w); err != nil {
			return utils.WrapError("Encode codeBlockGroupTransmission", err)
		}
	}
	if ie.xOverhead != nil {
		if err = ie.xOverhead.Encode(w); err != nil {
			return utils.WrapError("Encode xOverhead", err)
		}
	}
	if ie.nrofHARQ_ProcessesForPDSCH != nil {
		if err = ie.nrofHARQ_ProcessesForPDSCH.Encode(w); err != nil {
			return utils.WrapError("Encode nrofHARQ_ProcessesForPDSCH", err)
		}
	}
	if ie.pucch_Cell != nil {
		if err = ie.pucch_Cell.Encode(w); err != nil {
			return utils.WrapError("Encode pucch_Cell", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.maxMIMO_Layers != nil || ie.processingType2Enabled != nil, ie.pdsch_CodeBlockGroupTransmissionList_r16 != nil, ie.downlinkHARQ_FeedbackDisabled_r17 != nil || ie.nrofHARQ_ProcessesForPDSCH_v1700 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PDSCH_ServingCellConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.maxMIMO_Layers != nil, ie.processingType2Enabled != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode maxMIMO_Layers optional
			if ie.maxMIMO_Layers != nil {
				if err = extWriter.WriteInteger(*ie.maxMIMO_Layers, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
					return utils.WrapError("Encode maxMIMO_Layers", err)
				}
			}
			// encode processingType2Enabled optional
			if ie.processingType2Enabled != nil {
				if err = extWriter.WriteBoolean(*ie.processingType2Enabled); err != nil {
					return utils.WrapError("Encode processingType2Enabled", err)
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
			optionals_ext_2 := []bool{ie.pdsch_CodeBlockGroupTransmissionList_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode pdsch_CodeBlockGroupTransmissionList_r16 optional
			if ie.pdsch_CodeBlockGroupTransmissionList_r16 != nil {
				tmp_pdsch_CodeBlockGroupTransmissionList_r16 := utils.SetupRelease[*PDSCH_CodeBlockGroupTransmissionList_r16]{
					Setup: ie.pdsch_CodeBlockGroupTransmissionList_r16,
				}
				if err = tmp_pdsch_CodeBlockGroupTransmissionList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_CodeBlockGroupTransmissionList_r16", err)
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
			optionals_ext_3 := []bool{ie.downlinkHARQ_FeedbackDisabled_r17 != nil, ie.nrofHARQ_ProcessesForPDSCH_v1700 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode downlinkHARQ_FeedbackDisabled_r17 optional
			if ie.downlinkHARQ_FeedbackDisabled_r17 != nil {
				tmp_downlinkHARQ_FeedbackDisabled_r17 := utils.SetupRelease[*DownlinkHARQ_FeedbackDisabled_r17]{
					Setup: ie.downlinkHARQ_FeedbackDisabled_r17,
				}
				if err = tmp_downlinkHARQ_FeedbackDisabled_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode downlinkHARQ_FeedbackDisabled_r17", err)
				}
			}
			// encode nrofHARQ_ProcessesForPDSCH_v1700 optional
			if ie.nrofHARQ_ProcessesForPDSCH_v1700 != nil {
				if err = ie.nrofHARQ_ProcessesForPDSCH_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nrofHARQ_ProcessesForPDSCH_v1700", err)
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

func (ie *PDSCH_ServingCellConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var codeBlockGroupTransmissionPresent bool
	if codeBlockGroupTransmissionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var xOverheadPresent bool
	if xOverheadPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nrofHARQ_ProcessesForPDSCHPresent bool
	if nrofHARQ_ProcessesForPDSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pucch_CellPresent bool
	if pucch_CellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if codeBlockGroupTransmissionPresent {
		tmp_codeBlockGroupTransmission := utils.SetupRelease[*PDSCH_CodeBlockGroupTransmission]{}
		if err = tmp_codeBlockGroupTransmission.Decode(r); err != nil {
			return utils.WrapError("Decode codeBlockGroupTransmission", err)
		}
		ie.codeBlockGroupTransmission = tmp_codeBlockGroupTransmission.Setup
	}
	if xOverheadPresent {
		ie.xOverhead = new(PDSCH_ServingCellConfig_xOverhead)
		if err = ie.xOverhead.Decode(r); err != nil {
			return utils.WrapError("Decode xOverhead", err)
		}
	}
	if nrofHARQ_ProcessesForPDSCHPresent {
		ie.nrofHARQ_ProcessesForPDSCH = new(PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH)
		if err = ie.nrofHARQ_ProcessesForPDSCH.Decode(r); err != nil {
			return utils.WrapError("Decode nrofHARQ_ProcessesForPDSCH", err)
		}
	}
	if pucch_CellPresent {
		ie.pucch_Cell = new(ServCellIndex)
		if err = ie.pucch_Cell.Decode(r); err != nil {
			return utils.WrapError("Decode pucch_Cell", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
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

			maxMIMO_LayersPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			processingType2EnabledPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode maxMIMO_Layers optional
			if maxMIMO_LayersPresent {
				var tmp_int_maxMIMO_Layers int64
				if tmp_int_maxMIMO_Layers, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
					return utils.WrapError("Decode maxMIMO_Layers", err)
				}
				ie.maxMIMO_Layers = &tmp_int_maxMIMO_Layers
			}
			// decode processingType2Enabled optional
			if processingType2EnabledPresent {
				var tmp_bool_processingType2Enabled bool
				if tmp_bool_processingType2Enabled, err = extReader.ReadBoolean(); err != nil {
					return utils.WrapError("Decode processingType2Enabled", err)
				}
				ie.processingType2Enabled = &tmp_bool_processingType2Enabled
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			pdsch_CodeBlockGroupTransmissionList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode pdsch_CodeBlockGroupTransmissionList_r16 optional
			if pdsch_CodeBlockGroupTransmissionList_r16Present {
				tmp_pdsch_CodeBlockGroupTransmissionList_r16 := utils.SetupRelease[*PDSCH_CodeBlockGroupTransmissionList_r16]{}
				if err = tmp_pdsch_CodeBlockGroupTransmissionList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_CodeBlockGroupTransmissionList_r16", err)
				}
				ie.pdsch_CodeBlockGroupTransmissionList_r16 = tmp_pdsch_CodeBlockGroupTransmissionList_r16.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			downlinkHARQ_FeedbackDisabled_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			nrofHARQ_ProcessesForPDSCH_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode downlinkHARQ_FeedbackDisabled_r17 optional
			if downlinkHARQ_FeedbackDisabled_r17Present {
				tmp_downlinkHARQ_FeedbackDisabled_r17 := utils.SetupRelease[*DownlinkHARQ_FeedbackDisabled_r17]{}
				if err = tmp_downlinkHARQ_FeedbackDisabled_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode downlinkHARQ_FeedbackDisabled_r17", err)
				}
				ie.downlinkHARQ_FeedbackDisabled_r17 = tmp_downlinkHARQ_FeedbackDisabled_r17.Setup
			}
			// decode nrofHARQ_ProcessesForPDSCH_v1700 optional
			if nrofHARQ_ProcessesForPDSCH_v1700Present {
				ie.nrofHARQ_ProcessesForPDSCH_v1700 = new(PDSCH_ServingCellConfig_nrofHARQ_ProcessesForPDSCH_v1700)
				if err = ie.nrofHARQ_ProcessesForPDSCH_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode nrofHARQ_ProcessesForPDSCH_v1700", err)
				}
			}
		}
	}
	return nil
}
