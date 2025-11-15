package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AvailabilityCombinationsPerCell_r16 struct {
	availabilityCombinationsPerCellIndex_r16 AvailabilityCombinationsPerCellIndex_r16 `madatory`
	iab_DU_CellIdentity_r16                  CellIdentity                             `madatory`
	positionInDCI_AI_r16                     *int64                                   `lb:0,ub:maxAI_DCI_PayloadSize_1_r16,optional`
	availabilityCombinations_r16             []AvailabilityCombination_r16            `lb:1,ub:maxNrofAvailabilityCombinationsPerSet_r16,madatory`
	availabilityCombinationsRB_Groups_r17    []AvailabilityCombinationRB_Groups_r17   `lb:1,ub:maxNrofAvailabilityCombinationsPerSet_r16,optional,ext-1`
	positionInDCI_AI_RBGroups_v1720          *int64                                   `lb:0,ub:maxAI_DCI_PayloadSize_1_r16,optional,ext-2`
}

func (ie *AvailabilityCombinationsPerCell_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := len(ie.availabilityCombinationsRB_Groups_r17) > 0 || ie.positionInDCI_AI_RBGroups_v1720 != nil
	preambleBits := []bool{hasExtensions, ie.positionInDCI_AI_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.availabilityCombinationsPerCellIndex_r16.Encode(w); err != nil {
		return utils.WrapError("Encode availabilityCombinationsPerCellIndex_r16", err)
	}
	if err = ie.iab_DU_CellIdentity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode iab_DU_CellIdentity_r16", err)
	}
	if ie.positionInDCI_AI_r16 != nil {
		if err = w.WriteInteger(*ie.positionInDCI_AI_r16, &uper.Constraint{Lb: 0, Ub: maxAI_DCI_PayloadSize_1_r16}, false); err != nil {
			return utils.WrapError("Encode positionInDCI_AI_r16", err)
		}
	}
	tmp_availabilityCombinations_r16 := utils.NewSequence[*AvailabilityCombination_r16]([]*AvailabilityCombination_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofAvailabilityCombinationsPerSet_r16}, false)
	for _, i := range ie.availabilityCombinations_r16 {
		tmp_availabilityCombinations_r16.Value = append(tmp_availabilityCombinations_r16.Value, &i)
	}
	if err = tmp_availabilityCombinations_r16.Encode(w); err != nil {
		return utils.WrapError("Encode availabilityCombinations_r16", err)
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{len(ie.availabilityCombinationsRB_Groups_r17) > 0, ie.positionInDCI_AI_RBGroups_v1720 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap AvailabilityCombinationsPerCell_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.availabilityCombinationsRB_Groups_r17) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode availabilityCombinationsRB_Groups_r17 optional
			if len(ie.availabilityCombinationsRB_Groups_r17) > 0 {
				tmp_availabilityCombinationsRB_Groups_r17 := utils.NewSequence[*AvailabilityCombinationRB_Groups_r17]([]*AvailabilityCombinationRB_Groups_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAvailabilityCombinationsPerSet_r16}, false)
				for _, i := range ie.availabilityCombinationsRB_Groups_r17 {
					tmp_availabilityCombinationsRB_Groups_r17.Value = append(tmp_availabilityCombinationsRB_Groups_r17.Value, &i)
				}
				if err = tmp_availabilityCombinationsRB_Groups_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode availabilityCombinationsRB_Groups_r17", err)
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
			optionals_ext_2 := []bool{ie.positionInDCI_AI_RBGroups_v1720 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode positionInDCI_AI_RBGroups_v1720 optional
			if ie.positionInDCI_AI_RBGroups_v1720 != nil {
				if err = extWriter.WriteInteger(*ie.positionInDCI_AI_RBGroups_v1720, &uper.Constraint{Lb: 0, Ub: maxAI_DCI_PayloadSize_1_r16}, false); err != nil {
					return utils.WrapError("Encode positionInDCI_AI_RBGroups_v1720", err)
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

func (ie *AvailabilityCombinationsPerCell_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var positionInDCI_AI_r16Present bool
	if positionInDCI_AI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.availabilityCombinationsPerCellIndex_r16.Decode(r); err != nil {
		return utils.WrapError("Decode availabilityCombinationsPerCellIndex_r16", err)
	}
	if err = ie.iab_DU_CellIdentity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode iab_DU_CellIdentity_r16", err)
	}
	if positionInDCI_AI_r16Present {
		var tmp_int_positionInDCI_AI_r16 int64
		if tmp_int_positionInDCI_AI_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxAI_DCI_PayloadSize_1_r16}, false); err != nil {
			return utils.WrapError("Decode positionInDCI_AI_r16", err)
		}
		ie.positionInDCI_AI_r16 = &tmp_int_positionInDCI_AI_r16
	}
	tmp_availabilityCombinations_r16 := utils.NewSequence[*AvailabilityCombination_r16]([]*AvailabilityCombination_r16{}, uper.Constraint{Lb: 1, Ub: maxNrofAvailabilityCombinationsPerSet_r16}, false)
	fn_availabilityCombinations_r16 := func() *AvailabilityCombination_r16 {
		return new(AvailabilityCombination_r16)
	}
	if err = tmp_availabilityCombinations_r16.Decode(r, fn_availabilityCombinations_r16); err != nil {
		return utils.WrapError("Decode availabilityCombinations_r16", err)
	}
	ie.availabilityCombinations_r16 = []AvailabilityCombination_r16{}
	for _, i := range tmp_availabilityCombinations_r16.Value {
		ie.availabilityCombinations_r16 = append(ie.availabilityCombinations_r16, *i)
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

			availabilityCombinationsRB_Groups_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode availabilityCombinationsRB_Groups_r17 optional
			if availabilityCombinationsRB_Groups_r17Present {
				tmp_availabilityCombinationsRB_Groups_r17 := utils.NewSequence[*AvailabilityCombinationRB_Groups_r17]([]*AvailabilityCombinationRB_Groups_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofAvailabilityCombinationsPerSet_r16}, false)
				fn_availabilityCombinationsRB_Groups_r17 := func() *AvailabilityCombinationRB_Groups_r17 {
					return new(AvailabilityCombinationRB_Groups_r17)
				}
				if err = tmp_availabilityCombinationsRB_Groups_r17.Decode(extReader, fn_availabilityCombinationsRB_Groups_r17); err != nil {
					return utils.WrapError("Decode availabilityCombinationsRB_Groups_r17", err)
				}
				ie.availabilityCombinationsRB_Groups_r17 = []AvailabilityCombinationRB_Groups_r17{}
				for _, i := range tmp_availabilityCombinationsRB_Groups_r17.Value {
					ie.availabilityCombinationsRB_Groups_r17 = append(ie.availabilityCombinationsRB_Groups_r17, *i)
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

			positionInDCI_AI_RBGroups_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode positionInDCI_AI_RBGroups_v1720 optional
			if positionInDCI_AI_RBGroups_v1720Present {
				var tmp_int_positionInDCI_AI_RBGroups_v1720 int64
				if tmp_int_positionInDCI_AI_RBGroups_v1720, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxAI_DCI_PayloadSize_1_r16}, false); err != nil {
					return utils.WrapError("Decode positionInDCI_AI_RBGroups_v1720", err)
				}
				ie.positionInDCI_AI_RBGroups_v1720 = &tmp_int_positionInDCI_AI_RBGroups_v1720
			}
		}
	}
	return nil
}
