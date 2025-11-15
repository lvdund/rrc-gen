package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_ServingCellConfig struct {
	codeBlockGroupTransmission     *PUSCH_CodeBlockGroupTransmission                       `optional,setuprelease`
	rateMatching                   *PUSCH_ServingCellConfig_rateMatching                   `optional`
	xOverhead                      *PUSCH_ServingCellConfig_xOverhead                      `optional`
	maxMIMO_Layers                 *int64                                                  `lb:1,ub:4,optional,ext-1`
	processingType2Enabled         *bool                                                   `optional,ext-1`
	maxMIMO_LayersDCI_0_2_r16      *MaxMIMO_LayersDCI_0_2_r16                              `optional,ext-2,setuprelease`
	nrofHARQ_ProcessesForPUSCH_r17 *PUSCH_ServingCellConfig_nrofHARQ_ProcessesForPUSCH_r17 `optional,ext-3`
	uplinkHARQ_mode_r17            *UplinkHARQ_mode_r17                                    `optional,ext-3,setuprelease`
}

func (ie *PUSCH_ServingCellConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.maxMIMO_Layers != nil || ie.processingType2Enabled != nil || ie.maxMIMO_LayersDCI_0_2_r16 != nil || ie.nrofHARQ_ProcessesForPUSCH_r17 != nil || ie.uplinkHARQ_mode_r17 != nil
	preambleBits := []bool{hasExtensions, ie.codeBlockGroupTransmission != nil, ie.rateMatching != nil, ie.xOverhead != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.codeBlockGroupTransmission != nil {
		tmp_codeBlockGroupTransmission := utils.SetupRelease[*PUSCH_CodeBlockGroupTransmission]{
			Setup: ie.codeBlockGroupTransmission,
		}
		if err = tmp_codeBlockGroupTransmission.Encode(w); err != nil {
			return utils.WrapError("Encode codeBlockGroupTransmission", err)
		}
	}
	if ie.rateMatching != nil {
		if err = ie.rateMatching.Encode(w); err != nil {
			return utils.WrapError("Encode rateMatching", err)
		}
	}
	if ie.xOverhead != nil {
		if err = ie.xOverhead.Encode(w); err != nil {
			return utils.WrapError("Encode xOverhead", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.maxMIMO_Layers != nil || ie.processingType2Enabled != nil, ie.maxMIMO_LayersDCI_0_2_r16 != nil, ie.nrofHARQ_ProcessesForPUSCH_r17 != nil || ie.uplinkHARQ_mode_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PUSCH_ServingCellConfig", err)
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
				if err = extWriter.WriteInteger(*ie.maxMIMO_Layers, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
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
			optionals_ext_2 := []bool{ie.maxMIMO_LayersDCI_0_2_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode maxMIMO_LayersDCI_0_2_r16 optional
			if ie.maxMIMO_LayersDCI_0_2_r16 != nil {
				tmp_maxMIMO_LayersDCI_0_2_r16 := utils.SetupRelease[*MaxMIMO_LayersDCI_0_2_r16]{
					Setup: ie.maxMIMO_LayersDCI_0_2_r16,
				}
				if err = tmp_maxMIMO_LayersDCI_0_2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxMIMO_LayersDCI_0_2_r16", err)
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
			optionals_ext_3 := []bool{ie.nrofHARQ_ProcessesForPUSCH_r17 != nil, ie.uplinkHARQ_mode_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode nrofHARQ_ProcessesForPUSCH_r17 optional
			if ie.nrofHARQ_ProcessesForPUSCH_r17 != nil {
				if err = ie.nrofHARQ_ProcessesForPUSCH_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode nrofHARQ_ProcessesForPUSCH_r17", err)
				}
			}
			// encode uplinkHARQ_mode_r17 optional
			if ie.uplinkHARQ_mode_r17 != nil {
				tmp_uplinkHARQ_mode_r17 := utils.SetupRelease[*UplinkHARQ_mode_r17]{
					Setup: ie.uplinkHARQ_mode_r17,
				}
				if err = tmp_uplinkHARQ_mode_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode uplinkHARQ_mode_r17", err)
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

func (ie *PUSCH_ServingCellConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var codeBlockGroupTransmissionPresent bool
	if codeBlockGroupTransmissionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rateMatchingPresent bool
	if rateMatchingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var xOverheadPresent bool
	if xOverheadPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if codeBlockGroupTransmissionPresent {
		tmp_codeBlockGroupTransmission := utils.SetupRelease[*PUSCH_CodeBlockGroupTransmission]{}
		if err = tmp_codeBlockGroupTransmission.Decode(r); err != nil {
			return utils.WrapError("Decode codeBlockGroupTransmission", err)
		}
		ie.codeBlockGroupTransmission = tmp_codeBlockGroupTransmission.Setup
	}
	if rateMatchingPresent {
		ie.rateMatching = new(PUSCH_ServingCellConfig_rateMatching)
		if err = ie.rateMatching.Decode(r); err != nil {
			return utils.WrapError("Decode rateMatching", err)
		}
	}
	if xOverheadPresent {
		ie.xOverhead = new(PUSCH_ServingCellConfig_xOverhead)
		if err = ie.xOverhead.Decode(r); err != nil {
			return utils.WrapError("Decode xOverhead", err)
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
				if tmp_int_maxMIMO_Layers, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
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

			maxMIMO_LayersDCI_0_2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode maxMIMO_LayersDCI_0_2_r16 optional
			if maxMIMO_LayersDCI_0_2_r16Present {
				tmp_maxMIMO_LayersDCI_0_2_r16 := utils.SetupRelease[*MaxMIMO_LayersDCI_0_2_r16]{}
				if err = tmp_maxMIMO_LayersDCI_0_2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxMIMO_LayersDCI_0_2_r16", err)
				}
				ie.maxMIMO_LayersDCI_0_2_r16 = tmp_maxMIMO_LayersDCI_0_2_r16.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			nrofHARQ_ProcessesForPUSCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			uplinkHARQ_mode_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode nrofHARQ_ProcessesForPUSCH_r17 optional
			if nrofHARQ_ProcessesForPUSCH_r17Present {
				ie.nrofHARQ_ProcessesForPUSCH_r17 = new(PUSCH_ServingCellConfig_nrofHARQ_ProcessesForPUSCH_r17)
				if err = ie.nrofHARQ_ProcessesForPUSCH_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode nrofHARQ_ProcessesForPUSCH_r17", err)
				}
			}
			// decode uplinkHARQ_mode_r17 optional
			if uplinkHARQ_mode_r17Present {
				tmp_uplinkHARQ_mode_r17 := utils.SetupRelease[*UplinkHARQ_mode_r17]{}
				if err = tmp_uplinkHARQ_mode_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode uplinkHARQ_mode_r17", err)
				}
				ie.uplinkHARQ_mode_r17 = tmp_uplinkHARQ_mode_r17.Setup
			}
		}
	}
	return nil
}
