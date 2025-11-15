package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SSB_ConfigMobility struct {
	ssb_ToMeasure                         *SSB_ToMeasure                         `optional,setuprelease`
	deriveSSB_IndexFromCell               bool                                   `madatory`
	ss_RSSI_Measurement                   *SS_RSSI_Measurement                   `optional`
	ssb_PositionQCL_Common_r16            *SSB_PositionQCL_Relation_r16          `optional,ext-1`
	ssb_PositionQCL_CellsToAddModList_r16 *SSB_PositionQCL_CellsToAddModList_r16 `optional,ext-1`
	ssb_PositionQCL_CellsToRemoveList_r16 *PCI_List                              `optional,ext-1`
	deriveSSB_IndexFromCellInter_r17      *ServCellIndex                         `optional,ext-2`
	ssb_PositionQCL_Common_r17            *SSB_PositionQCL_Relation_r17          `optional,ext-2`
	ssb_PositionQCL_Cells_r17             *SSB_PositionQCL_CellList_r17          `optional,ext-2,setuprelease`
	cca_CellsToAddModList_r17             *PCI_List                              `optional,ext-3`
	cca_CellsToRemoveList_r17             *PCI_List                              `optional,ext-3`
}

func (ie *SSB_ConfigMobility) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.ssb_PositionQCL_Common_r16 != nil || ie.ssb_PositionQCL_CellsToAddModList_r16 != nil || ie.ssb_PositionQCL_CellsToRemoveList_r16 != nil || ie.deriveSSB_IndexFromCellInter_r17 != nil || ie.ssb_PositionQCL_Common_r17 != nil || ie.ssb_PositionQCL_Cells_r17 != nil || ie.cca_CellsToAddModList_r17 != nil || ie.cca_CellsToRemoveList_r17 != nil
	preambleBits := []bool{hasExtensions, ie.ssb_ToMeasure != nil, ie.ss_RSSI_Measurement != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ssb_ToMeasure != nil {
		tmp_ssb_ToMeasure := utils.SetupRelease[*SSB_ToMeasure]{
			Setup: ie.ssb_ToMeasure,
		}
		if err = tmp_ssb_ToMeasure.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_ToMeasure", err)
		}
	}
	if err = w.WriteBoolean(ie.deriveSSB_IndexFromCell); err != nil {
		return utils.WrapError("WriteBoolean deriveSSB_IndexFromCell", err)
	}
	if ie.ss_RSSI_Measurement != nil {
		if err = ie.ss_RSSI_Measurement.Encode(w); err != nil {
			return utils.WrapError("Encode ss_RSSI_Measurement", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.ssb_PositionQCL_Common_r16 != nil || ie.ssb_PositionQCL_CellsToAddModList_r16 != nil || ie.ssb_PositionQCL_CellsToRemoveList_r16 != nil, ie.deriveSSB_IndexFromCellInter_r17 != nil || ie.ssb_PositionQCL_Common_r17 != nil || ie.ssb_PositionQCL_Cells_r17 != nil, ie.cca_CellsToAddModList_r17 != nil || ie.cca_CellsToRemoveList_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SSB_ConfigMobility", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.ssb_PositionQCL_Common_r16 != nil, ie.ssb_PositionQCL_CellsToAddModList_r16 != nil, ie.ssb_PositionQCL_CellsToRemoveList_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ssb_PositionQCL_Common_r16 optional
			if ie.ssb_PositionQCL_Common_r16 != nil {
				if err = ie.ssb_PositionQCL_Common_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ssb_PositionQCL_Common_r16", err)
				}
			}
			// encode ssb_PositionQCL_CellsToAddModList_r16 optional
			if ie.ssb_PositionQCL_CellsToAddModList_r16 != nil {
				if err = ie.ssb_PositionQCL_CellsToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ssb_PositionQCL_CellsToAddModList_r16", err)
				}
			}
			// encode ssb_PositionQCL_CellsToRemoveList_r16 optional
			if ie.ssb_PositionQCL_CellsToRemoveList_r16 != nil {
				if err = ie.ssb_PositionQCL_CellsToRemoveList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ssb_PositionQCL_CellsToRemoveList_r16", err)
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
			optionals_ext_2 := []bool{ie.deriveSSB_IndexFromCellInter_r17 != nil, ie.ssb_PositionQCL_Common_r17 != nil, ie.ssb_PositionQCL_Cells_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode deriveSSB_IndexFromCellInter_r17 optional
			if ie.deriveSSB_IndexFromCellInter_r17 != nil {
				if err = ie.deriveSSB_IndexFromCellInter_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode deriveSSB_IndexFromCellInter_r17", err)
				}
			}
			// encode ssb_PositionQCL_Common_r17 optional
			if ie.ssb_PositionQCL_Common_r17 != nil {
				if err = ie.ssb_PositionQCL_Common_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ssb_PositionQCL_Common_r17", err)
				}
			}
			// encode ssb_PositionQCL_Cells_r17 optional
			if ie.ssb_PositionQCL_Cells_r17 != nil {
				tmp_ssb_PositionQCL_Cells_r17 := utils.SetupRelease[*SSB_PositionQCL_CellList_r17]{
					Setup: ie.ssb_PositionQCL_Cells_r17,
				}
				if err = tmp_ssb_PositionQCL_Cells_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ssb_PositionQCL_Cells_r17", err)
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
			optionals_ext_3 := []bool{ie.cca_CellsToAddModList_r17 != nil, ie.cca_CellsToRemoveList_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode cca_CellsToAddModList_r17 optional
			if ie.cca_CellsToAddModList_r17 != nil {
				if err = ie.cca_CellsToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cca_CellsToAddModList_r17", err)
				}
			}
			// encode cca_CellsToRemoveList_r17 optional
			if ie.cca_CellsToRemoveList_r17 != nil {
				if err = ie.cca_CellsToRemoveList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cca_CellsToRemoveList_r17", err)
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

func (ie *SSB_ConfigMobility) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ssb_ToMeasurePresent bool
	if ssb_ToMeasurePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ss_RSSI_MeasurementPresent bool
	if ss_RSSI_MeasurementPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ssb_ToMeasurePresent {
		tmp_ssb_ToMeasure := utils.SetupRelease[*SSB_ToMeasure]{}
		if err = tmp_ssb_ToMeasure.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_ToMeasure", err)
		}
		ie.ssb_ToMeasure = tmp_ssb_ToMeasure.Setup
	}
	var tmp_bool_deriveSSB_IndexFromCell bool
	if tmp_bool_deriveSSB_IndexFromCell, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean deriveSSB_IndexFromCell", err)
	}
	ie.deriveSSB_IndexFromCell = tmp_bool_deriveSSB_IndexFromCell
	if ss_RSSI_MeasurementPresent {
		ie.ss_RSSI_Measurement = new(SS_RSSI_Measurement)
		if err = ie.ss_RSSI_Measurement.Decode(r); err != nil {
			return utils.WrapError("Decode ss_RSSI_Measurement", err)
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

			ssb_PositionQCL_Common_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ssb_PositionQCL_CellsToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ssb_PositionQCL_CellsToRemoveList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ssb_PositionQCL_Common_r16 optional
			if ssb_PositionQCL_Common_r16Present {
				ie.ssb_PositionQCL_Common_r16 = new(SSB_PositionQCL_Relation_r16)
				if err = ie.ssb_PositionQCL_Common_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ssb_PositionQCL_Common_r16", err)
				}
			}
			// decode ssb_PositionQCL_CellsToAddModList_r16 optional
			if ssb_PositionQCL_CellsToAddModList_r16Present {
				ie.ssb_PositionQCL_CellsToAddModList_r16 = new(SSB_PositionQCL_CellsToAddModList_r16)
				if err = ie.ssb_PositionQCL_CellsToAddModList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ssb_PositionQCL_CellsToAddModList_r16", err)
				}
			}
			// decode ssb_PositionQCL_CellsToRemoveList_r16 optional
			if ssb_PositionQCL_CellsToRemoveList_r16Present {
				ie.ssb_PositionQCL_CellsToRemoveList_r16 = new(PCI_List)
				if err = ie.ssb_PositionQCL_CellsToRemoveList_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode ssb_PositionQCL_CellsToRemoveList_r16", err)
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

			deriveSSB_IndexFromCellInter_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ssb_PositionQCL_Common_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ssb_PositionQCL_Cells_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode deriveSSB_IndexFromCellInter_r17 optional
			if deriveSSB_IndexFromCellInter_r17Present {
				ie.deriveSSB_IndexFromCellInter_r17 = new(ServCellIndex)
				if err = ie.deriveSSB_IndexFromCellInter_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode deriveSSB_IndexFromCellInter_r17", err)
				}
			}
			// decode ssb_PositionQCL_Common_r17 optional
			if ssb_PositionQCL_Common_r17Present {
				ie.ssb_PositionQCL_Common_r17 = new(SSB_PositionQCL_Relation_r17)
				if err = ie.ssb_PositionQCL_Common_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ssb_PositionQCL_Common_r17", err)
				}
			}
			// decode ssb_PositionQCL_Cells_r17 optional
			if ssb_PositionQCL_Cells_r17Present {
				tmp_ssb_PositionQCL_Cells_r17 := utils.SetupRelease[*SSB_PositionQCL_CellList_r17]{}
				if err = tmp_ssb_PositionQCL_Cells_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ssb_PositionQCL_Cells_r17", err)
				}
				ie.ssb_PositionQCL_Cells_r17 = tmp_ssb_PositionQCL_Cells_r17.Setup
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			cca_CellsToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cca_CellsToRemoveList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode cca_CellsToAddModList_r17 optional
			if cca_CellsToAddModList_r17Present {
				ie.cca_CellsToAddModList_r17 = new(PCI_List)
				if err = ie.cca_CellsToAddModList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode cca_CellsToAddModList_r17", err)
				}
			}
			// decode cca_CellsToRemoveList_r17 optional
			if cca_CellsToRemoveList_r17Present {
				ie.cca_CellsToRemoveList_r17 = new(PCI_List)
				if err = ie.cca_CellsToRemoveList_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode cca_CellsToRemoveList_r17", err)
				}
			}
		}
	}
	return nil
}
