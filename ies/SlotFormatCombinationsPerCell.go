package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SlotFormatCombinationsPerCell struct {
	servingCellId          ServCellIndex                                         `madatory`
	subcarrierSpacing      SubcarrierSpacing                                     `madatory`
	subcarrierSpacing2     *SubcarrierSpacing                                    `optional`
	slotFormatCombinations []SlotFormatCombination                               `lb:1,ub:maxNrofSlotFormatCombinationsPerSet,optional`
	positionInDCI          *int64                                                `lb:0,ub:maxSFI_DCI_PayloadSize_1,optional`
	enableConfiguredUL_r16 *SlotFormatCombinationsPerCell_enableConfiguredUL_r16 `optional,ext-1`
}

func (ie *SlotFormatCombinationsPerCell) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.enableConfiguredUL_r16 != nil
	preambleBits := []bool{hasExtensions, ie.subcarrierSpacing2 != nil, len(ie.slotFormatCombinations) > 0, ie.positionInDCI != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.servingCellId.Encode(w); err != nil {
		return utils.WrapError("Encode servingCellId", err)
	}
	if err = ie.subcarrierSpacing.Encode(w); err != nil {
		return utils.WrapError("Encode subcarrierSpacing", err)
	}
	if ie.subcarrierSpacing2 != nil {
		if err = ie.subcarrierSpacing2.Encode(w); err != nil {
			return utils.WrapError("Encode subcarrierSpacing2", err)
		}
	}
	if len(ie.slotFormatCombinations) > 0 {
		tmp_slotFormatCombinations := utils.NewSequence[*SlotFormatCombination]([]*SlotFormatCombination{}, uper.Constraint{Lb: 1, Ub: maxNrofSlotFormatCombinationsPerSet}, false)
		for _, i := range ie.slotFormatCombinations {
			tmp_slotFormatCombinations.Value = append(tmp_slotFormatCombinations.Value, &i)
		}
		if err = tmp_slotFormatCombinations.Encode(w); err != nil {
			return utils.WrapError("Encode slotFormatCombinations", err)
		}
	}
	if ie.positionInDCI != nil {
		if err = w.WriteInteger(*ie.positionInDCI, &uper.Constraint{Lb: 0, Ub: maxSFI_DCI_PayloadSize_1}, false); err != nil {
			return utils.WrapError("Encode positionInDCI", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.enableConfiguredUL_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SlotFormatCombinationsPerCell", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.enableConfiguredUL_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode enableConfiguredUL_r16 optional
			if ie.enableConfiguredUL_r16 != nil {
				if err = ie.enableConfiguredUL_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enableConfiguredUL_r16", err)
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

func (ie *SlotFormatCombinationsPerCell) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var subcarrierSpacing2Present bool
	if subcarrierSpacing2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var slotFormatCombinationsPresent bool
	if slotFormatCombinationsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var positionInDCIPresent bool
	if positionInDCIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.servingCellId.Decode(r); err != nil {
		return utils.WrapError("Decode servingCellId", err)
	}
	if err = ie.subcarrierSpacing.Decode(r); err != nil {
		return utils.WrapError("Decode subcarrierSpacing", err)
	}
	if subcarrierSpacing2Present {
		ie.subcarrierSpacing2 = new(SubcarrierSpacing)
		if err = ie.subcarrierSpacing2.Decode(r); err != nil {
			return utils.WrapError("Decode subcarrierSpacing2", err)
		}
	}
	if slotFormatCombinationsPresent {
		tmp_slotFormatCombinations := utils.NewSequence[*SlotFormatCombination]([]*SlotFormatCombination{}, uper.Constraint{Lb: 1, Ub: maxNrofSlotFormatCombinationsPerSet}, false)
		fn_slotFormatCombinations := func() *SlotFormatCombination {
			return new(SlotFormatCombination)
		}
		if err = tmp_slotFormatCombinations.Decode(r, fn_slotFormatCombinations); err != nil {
			return utils.WrapError("Decode slotFormatCombinations", err)
		}
		ie.slotFormatCombinations = []SlotFormatCombination{}
		for _, i := range tmp_slotFormatCombinations.Value {
			ie.slotFormatCombinations = append(ie.slotFormatCombinations, *i)
		}
	}
	if positionInDCIPresent {
		var tmp_int_positionInDCI int64
		if tmp_int_positionInDCI, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxSFI_DCI_PayloadSize_1}, false); err != nil {
			return utils.WrapError("Decode positionInDCI", err)
		}
		ie.positionInDCI = &tmp_int_positionInDCI
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

			enableConfiguredUL_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode enableConfiguredUL_r16 optional
			if enableConfiguredUL_r16Present {
				ie.enableConfiguredUL_r16 = new(SlotFormatCombinationsPerCell_enableConfiguredUL_r16)
				if err = ie.enableConfiguredUL_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode enableConfiguredUL_r16", err)
				}
			}
		}
	}
	return nil
}
