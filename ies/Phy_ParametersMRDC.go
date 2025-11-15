package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersMRDC struct {
	naics_Capability_List             []NAICS_Capability_Entry                              `lb:1,ub:maxNrofNAICS_Entries,optional`
	spCellPlacement                   *CarrierAggregationVariant                            `optional,ext-1`
	tdd_PCellUL_TX_AllUL_Subframe_r16 *Phy_ParametersMRDC_tdd_PCellUL_TX_AllUL_Subframe_r16 `optional,ext-2`
	fdd_PCellUL_TX_AllUL_Subframe_r16 *Phy_ParametersMRDC_fdd_PCellUL_TX_AllUL_Subframe_r16 `optional,ext-2`
}

func (ie *Phy_ParametersMRDC) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.spCellPlacement != nil || ie.tdd_PCellUL_TX_AllUL_Subframe_r16 != nil || ie.fdd_PCellUL_TX_AllUL_Subframe_r16 != nil
	preambleBits := []bool{hasExtensions, len(ie.naics_Capability_List) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.naics_Capability_List) > 0 {
		tmp_naics_Capability_List := utils.NewSequence[*NAICS_Capability_Entry]([]*NAICS_Capability_Entry{}, uper.Constraint{Lb: 1, Ub: maxNrofNAICS_Entries}, false)
		for _, i := range ie.naics_Capability_List {
			tmp_naics_Capability_List.Value = append(tmp_naics_Capability_List.Value, &i)
		}
		if err = tmp_naics_Capability_List.Encode(w); err != nil {
			return utils.WrapError("Encode naics_Capability_List", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.spCellPlacement != nil, ie.tdd_PCellUL_TX_AllUL_Subframe_r16 != nil || ie.fdd_PCellUL_TX_AllUL_Subframe_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap Phy_ParametersMRDC", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.spCellPlacement != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode spCellPlacement optional
			if ie.spCellPlacement != nil {
				if err = ie.spCellPlacement.Encode(extWriter); err != nil {
					return utils.WrapError("Encode spCellPlacement", err)
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
			optionals_ext_2 := []bool{ie.tdd_PCellUL_TX_AllUL_Subframe_r16 != nil, ie.fdd_PCellUL_TX_AllUL_Subframe_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode tdd_PCellUL_TX_AllUL_Subframe_r16 optional
			if ie.tdd_PCellUL_TX_AllUL_Subframe_r16 != nil {
				if err = ie.tdd_PCellUL_TX_AllUL_Subframe_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode tdd_PCellUL_TX_AllUL_Subframe_r16", err)
				}
			}
			// encode fdd_PCellUL_TX_AllUL_Subframe_r16 optional
			if ie.fdd_PCellUL_TX_AllUL_Subframe_r16 != nil {
				if err = ie.fdd_PCellUL_TX_AllUL_Subframe_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode fdd_PCellUL_TX_AllUL_Subframe_r16", err)
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

func (ie *Phy_ParametersMRDC) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var naics_Capability_ListPresent bool
	if naics_Capability_ListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if naics_Capability_ListPresent {
		tmp_naics_Capability_List := utils.NewSequence[*NAICS_Capability_Entry]([]*NAICS_Capability_Entry{}, uper.Constraint{Lb: 1, Ub: maxNrofNAICS_Entries}, false)
		fn_naics_Capability_List := func() *NAICS_Capability_Entry {
			return new(NAICS_Capability_Entry)
		}
		if err = tmp_naics_Capability_List.Decode(r, fn_naics_Capability_List); err != nil {
			return utils.WrapError("Decode naics_Capability_List", err)
		}
		ie.naics_Capability_List = []NAICS_Capability_Entry{}
		for _, i := range tmp_naics_Capability_List.Value {
			ie.naics_Capability_List = append(ie.naics_Capability_List, *i)
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

			spCellPlacementPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode spCellPlacement optional
			if spCellPlacementPresent {
				ie.spCellPlacement = new(CarrierAggregationVariant)
				if err = ie.spCellPlacement.Decode(extReader); err != nil {
					return utils.WrapError("Decode spCellPlacement", err)
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

			tdd_PCellUL_TX_AllUL_Subframe_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			fdd_PCellUL_TX_AllUL_Subframe_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode tdd_PCellUL_TX_AllUL_Subframe_r16 optional
			if tdd_PCellUL_TX_AllUL_Subframe_r16Present {
				ie.tdd_PCellUL_TX_AllUL_Subframe_r16 = new(Phy_ParametersMRDC_tdd_PCellUL_TX_AllUL_Subframe_r16)
				if err = ie.tdd_PCellUL_TX_AllUL_Subframe_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode tdd_PCellUL_TX_AllUL_Subframe_r16", err)
				}
			}
			// decode fdd_PCellUL_TX_AllUL_Subframe_r16 optional
			if fdd_PCellUL_TX_AllUL_Subframe_r16Present {
				ie.fdd_PCellUL_TX_AllUL_Subframe_r16 = new(Phy_ParametersMRDC_fdd_PCellUL_TX_AllUL_Subframe_r16)
				if err = ie.fdd_PCellUL_TX_AllUL_Subframe_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode fdd_PCellUL_TX_AllUL_Subframe_r16", err)
				}
			}
		}
	}
	return nil
}
