package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_RS_ResourceConfigMobility struct {
	subcarrierSpacing        SubcarrierSpacing     `madatory`
	csi_RS_CellList_Mobility []CSI_RS_CellMobility `lb:1,ub:maxNrofCSI_RS_CellsRRM,madatory`
	refServCellIndex         *ServCellIndex        `optional,ext-1`
}

func (ie *CSI_RS_ResourceConfigMobility) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.refServCellIndex != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.subcarrierSpacing.Encode(w); err != nil {
		return utils.WrapError("Encode subcarrierSpacing", err)
	}
	tmp_csi_RS_CellList_Mobility := utils.NewSequence[*CSI_RS_CellMobility]([]*CSI_RS_CellMobility{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_CellsRRM}, false)
	for _, i := range ie.csi_RS_CellList_Mobility {
		tmp_csi_RS_CellList_Mobility.Value = append(tmp_csi_RS_CellList_Mobility.Value, &i)
	}
	if err = tmp_csi_RS_CellList_Mobility.Encode(w); err != nil {
		return utils.WrapError("Encode csi_RS_CellList_Mobility", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.refServCellIndex != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CSI_RS_ResourceConfigMobility", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.refServCellIndex != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode refServCellIndex optional
			if ie.refServCellIndex != nil {
				if err = ie.refServCellIndex.Encode(extWriter); err != nil {
					return utils.WrapError("Encode refServCellIndex", err)
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

func (ie *CSI_RS_ResourceConfigMobility) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.subcarrierSpacing.Decode(r); err != nil {
		return utils.WrapError("Decode subcarrierSpacing", err)
	}
	tmp_csi_RS_CellList_Mobility := utils.NewSequence[*CSI_RS_CellMobility]([]*CSI_RS_CellMobility{}, uper.Constraint{Lb: 1, Ub: maxNrofCSI_RS_CellsRRM}, false)
	fn_csi_RS_CellList_Mobility := func() *CSI_RS_CellMobility {
		return new(CSI_RS_CellMobility)
	}
	if err = tmp_csi_RS_CellList_Mobility.Decode(r, fn_csi_RS_CellList_Mobility); err != nil {
		return utils.WrapError("Decode csi_RS_CellList_Mobility", err)
	}
	ie.csi_RS_CellList_Mobility = []CSI_RS_CellMobility{}
	for _, i := range tmp_csi_RS_CellList_Mobility.Value {
		ie.csi_RS_CellList_Mobility = append(ie.csi_RS_CellList_Mobility, *i)
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
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

			refServCellIndexPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode refServCellIndex optional
			if refServCellIndexPresent {
				ie.refServCellIndex = new(ServCellIndex)
				if err = ie.refServCellIndex.Decode(extReader); err != nil {
					return utils.WrapError("Decode refServCellIndex", err)
				}
			}
		}
	}
	return nil
}
