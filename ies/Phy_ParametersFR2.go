package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersFR2 struct {
	dummy                                               *Phy_ParametersFR2_dummy                                               `optional`
	pdsch_RE_MappingFR2_PerSymbol                       *Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSymbol                       `optional`
	pCell_FR2                                           *Phy_ParametersFR2_pCell_FR2                                           `optional,ext-1`
	pdsch_RE_MappingFR2_PerSlot                         *Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot                         `optional,ext-1`
	defaultSpatialRelationPathlossRS_r16                *Phy_ParametersFR2_defaultSpatialRelationPathlossRS_r16                `optional,ext-2`
	spatialRelationUpdateAP_SRS_r16                     *Phy_ParametersFR2_spatialRelationUpdateAP_SRS_r16                     `optional,ext-2`
	maxNumberSRS_PosSpatialRelationsAllServingCells_r16 *Phy_ParametersFR2_maxNumberSRS_PosSpatialRelationsAllServingCells_r16 `optional,ext-2`
}

func (ie *Phy_ParametersFR2) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.pCell_FR2 != nil || ie.pdsch_RE_MappingFR2_PerSlot != nil || ie.defaultSpatialRelationPathlossRS_r16 != nil || ie.spatialRelationUpdateAP_SRS_r16 != nil || ie.maxNumberSRS_PosSpatialRelationsAllServingCells_r16 != nil
	preambleBits := []bool{hasExtensions, ie.dummy != nil, ie.pdsch_RE_MappingFR2_PerSymbol != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dummy != nil {
		if err = ie.dummy.Encode(w); err != nil {
			return utils.WrapError("Encode dummy", err)
		}
	}
	if ie.pdsch_RE_MappingFR2_PerSymbol != nil {
		if err = ie.pdsch_RE_MappingFR2_PerSymbol.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_RE_MappingFR2_PerSymbol", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.pCell_FR2 != nil || ie.pdsch_RE_MappingFR2_PerSlot != nil, ie.defaultSpatialRelationPathlossRS_r16 != nil || ie.spatialRelationUpdateAP_SRS_r16 != nil || ie.maxNumberSRS_PosSpatialRelationsAllServingCells_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap Phy_ParametersFR2", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.pCell_FR2 != nil, ie.pdsch_RE_MappingFR2_PerSlot != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode pCell_FR2 optional
			if ie.pCell_FR2 != nil {
				if err = ie.pCell_FR2.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pCell_FR2", err)
				}
			}
			// encode pdsch_RE_MappingFR2_PerSlot optional
			if ie.pdsch_RE_MappingFR2_PerSlot != nil {
				if err = ie.pdsch_RE_MappingFR2_PerSlot.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pdsch_RE_MappingFR2_PerSlot", err)
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
			optionals_ext_2 := []bool{ie.defaultSpatialRelationPathlossRS_r16 != nil, ie.spatialRelationUpdateAP_SRS_r16 != nil, ie.maxNumberSRS_PosSpatialRelationsAllServingCells_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode defaultSpatialRelationPathlossRS_r16 optional
			if ie.defaultSpatialRelationPathlossRS_r16 != nil {
				if err = ie.defaultSpatialRelationPathlossRS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode defaultSpatialRelationPathlossRS_r16", err)
				}
			}
			// encode spatialRelationUpdateAP_SRS_r16 optional
			if ie.spatialRelationUpdateAP_SRS_r16 != nil {
				if err = ie.spatialRelationUpdateAP_SRS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode spatialRelationUpdateAP_SRS_r16", err)
				}
			}
			// encode maxNumberSRS_PosSpatialRelationsAllServingCells_r16 optional
			if ie.maxNumberSRS_PosSpatialRelationsAllServingCells_r16 != nil {
				if err = ie.maxNumberSRS_PosSpatialRelationsAllServingCells_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode maxNumberSRS_PosSpatialRelationsAllServingCells_r16", err)
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

func (ie *Phy_ParametersFR2) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var dummyPresent bool
	if dummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_RE_MappingFR2_PerSymbolPresent bool
	if pdsch_RE_MappingFR2_PerSymbolPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if dummyPresent {
		ie.dummy = new(Phy_ParametersFR2_dummy)
		if err = ie.dummy.Decode(r); err != nil {
			return utils.WrapError("Decode dummy", err)
		}
	}
	if pdsch_RE_MappingFR2_PerSymbolPresent {
		ie.pdsch_RE_MappingFR2_PerSymbol = new(Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSymbol)
		if err = ie.pdsch_RE_MappingFR2_PerSymbol.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_RE_MappingFR2_PerSymbol", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
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

			pCell_FR2Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pdsch_RE_MappingFR2_PerSlotPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode pCell_FR2 optional
			if pCell_FR2Present {
				ie.pCell_FR2 = new(Phy_ParametersFR2_pCell_FR2)
				if err = ie.pCell_FR2.Decode(extReader); err != nil {
					return utils.WrapError("Decode pCell_FR2", err)
				}
			}
			// decode pdsch_RE_MappingFR2_PerSlot optional
			if pdsch_RE_MappingFR2_PerSlotPresent {
				ie.pdsch_RE_MappingFR2_PerSlot = new(Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot)
				if err = ie.pdsch_RE_MappingFR2_PerSlot.Decode(extReader); err != nil {
					return utils.WrapError("Decode pdsch_RE_MappingFR2_PerSlot", err)
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

			defaultSpatialRelationPathlossRS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			spatialRelationUpdateAP_SRS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			maxNumberSRS_PosSpatialRelationsAllServingCells_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode defaultSpatialRelationPathlossRS_r16 optional
			if defaultSpatialRelationPathlossRS_r16Present {
				ie.defaultSpatialRelationPathlossRS_r16 = new(Phy_ParametersFR2_defaultSpatialRelationPathlossRS_r16)
				if err = ie.defaultSpatialRelationPathlossRS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode defaultSpatialRelationPathlossRS_r16", err)
				}
			}
			// decode spatialRelationUpdateAP_SRS_r16 optional
			if spatialRelationUpdateAP_SRS_r16Present {
				ie.spatialRelationUpdateAP_SRS_r16 = new(Phy_ParametersFR2_spatialRelationUpdateAP_SRS_r16)
				if err = ie.spatialRelationUpdateAP_SRS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode spatialRelationUpdateAP_SRS_r16", err)
				}
			}
			// decode maxNumberSRS_PosSpatialRelationsAllServingCells_r16 optional
			if maxNumberSRS_PosSpatialRelationsAllServingCells_r16Present {
				ie.maxNumberSRS_PosSpatialRelationsAllServingCells_r16 = new(Phy_ParametersFR2_maxNumberSRS_PosSpatialRelationsAllServingCells_r16)
				if err = ie.maxNumberSRS_PosSpatialRelationsAllServingCells_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode maxNumberSRS_PosSpatialRelationsAllServingCells_r16", err)
				}
			}
		}
	}
	return nil
}
