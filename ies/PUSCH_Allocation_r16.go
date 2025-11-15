package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUSCH_Allocation_r16 struct {
	mappingType_r16            *PUSCH_Allocation_r16_mappingType_r16            `optional`
	startSymbolAndLength_r16   *int64                                           `lb:0,ub:127,optional`
	startSymbol_r16            *int64                                           `lb:0,ub:13,optional`
	length_r16                 *int64                                           `lb:1,ub:14,optional`
	numberOfRepetitions_r16    *PUSCH_Allocation_r16_numberOfRepetitions_r16    `optional`
	numberOfRepetitionsExt_r17 *PUSCH_Allocation_r16_numberOfRepetitionsExt_r17 `optional,ext-1`
	numberOfSlotsTBoMS_r17     *PUSCH_Allocation_r16_numberOfSlotsTBoMS_r17     `optional,ext-1`
	extendedK2_r17             *int64                                           `lb:0,ub:128,optional,ext-1`
}

func (ie *PUSCH_Allocation_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.numberOfRepetitionsExt_r17 != nil || ie.numberOfSlotsTBoMS_r17 != nil || ie.extendedK2_r17 != nil
	preambleBits := []bool{hasExtensions, ie.mappingType_r16 != nil, ie.startSymbolAndLength_r16 != nil, ie.startSymbol_r16 != nil, ie.length_r16 != nil, ie.numberOfRepetitions_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.mappingType_r16 != nil {
		if err = ie.mappingType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mappingType_r16", err)
		}
	}
	if ie.startSymbolAndLength_r16 != nil {
		if err = w.WriteInteger(*ie.startSymbolAndLength_r16, &uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Encode startSymbolAndLength_r16", err)
		}
	}
	if ie.startSymbol_r16 != nil {
		if err = w.WriteInteger(*ie.startSymbol_r16, &uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
			return utils.WrapError("Encode startSymbol_r16", err)
		}
	}
	if ie.length_r16 != nil {
		if err = w.WriteInteger(*ie.length_r16, &uper.Constraint{Lb: 1, Ub: 14}, false); err != nil {
			return utils.WrapError("Encode length_r16", err)
		}
	}
	if ie.numberOfRepetitions_r16 != nil {
		if err = ie.numberOfRepetitions_r16.Encode(w); err != nil {
			return utils.WrapError("Encode numberOfRepetitions_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.numberOfRepetitionsExt_r17 != nil || ie.numberOfSlotsTBoMS_r17 != nil || ie.extendedK2_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PUSCH_Allocation_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.numberOfRepetitionsExt_r17 != nil, ie.numberOfSlotsTBoMS_r17 != nil, ie.extendedK2_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode numberOfRepetitionsExt_r17 optional
			if ie.numberOfRepetitionsExt_r17 != nil {
				if err = ie.numberOfRepetitionsExt_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode numberOfRepetitionsExt_r17", err)
				}
			}
			// encode numberOfSlotsTBoMS_r17 optional
			if ie.numberOfSlotsTBoMS_r17 != nil {
				if err = ie.numberOfSlotsTBoMS_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode numberOfSlotsTBoMS_r17", err)
				}
			}
			// encode extendedK2_r17 optional
			if ie.extendedK2_r17 != nil {
				if err = extWriter.WriteInteger(*ie.extendedK2_r17, &uper.Constraint{Lb: 0, Ub: 128}, false); err != nil {
					return utils.WrapError("Encode extendedK2_r17", err)
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

func (ie *PUSCH_Allocation_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var mappingType_r16Present bool
	if mappingType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var startSymbolAndLength_r16Present bool
	if startSymbolAndLength_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var startSymbol_r16Present bool
	if startSymbol_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var length_r16Present bool
	if length_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var numberOfRepetitions_r16Present bool
	if numberOfRepetitions_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if mappingType_r16Present {
		ie.mappingType_r16 = new(PUSCH_Allocation_r16_mappingType_r16)
		if err = ie.mappingType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mappingType_r16", err)
		}
	}
	if startSymbolAndLength_r16Present {
		var tmp_int_startSymbolAndLength_r16 int64
		if tmp_int_startSymbolAndLength_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
			return utils.WrapError("Decode startSymbolAndLength_r16", err)
		}
		ie.startSymbolAndLength_r16 = &tmp_int_startSymbolAndLength_r16
	}
	if startSymbol_r16Present {
		var tmp_int_startSymbol_r16 int64
		if tmp_int_startSymbol_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
			return utils.WrapError("Decode startSymbol_r16", err)
		}
		ie.startSymbol_r16 = &tmp_int_startSymbol_r16
	}
	if length_r16Present {
		var tmp_int_length_r16 int64
		if tmp_int_length_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 14}, false); err != nil {
			return utils.WrapError("Decode length_r16", err)
		}
		ie.length_r16 = &tmp_int_length_r16
	}
	if numberOfRepetitions_r16Present {
		ie.numberOfRepetitions_r16 = new(PUSCH_Allocation_r16_numberOfRepetitions_r16)
		if err = ie.numberOfRepetitions_r16.Decode(r); err != nil {
			return utils.WrapError("Decode numberOfRepetitions_r16", err)
		}
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

			numberOfRepetitionsExt_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			numberOfSlotsTBoMS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			extendedK2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode numberOfRepetitionsExt_r17 optional
			if numberOfRepetitionsExt_r17Present {
				ie.numberOfRepetitionsExt_r17 = new(PUSCH_Allocation_r16_numberOfRepetitionsExt_r17)
				if err = ie.numberOfRepetitionsExt_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode numberOfRepetitionsExt_r17", err)
				}
			}
			// decode numberOfSlotsTBoMS_r17 optional
			if numberOfSlotsTBoMS_r17Present {
				ie.numberOfSlotsTBoMS_r17 = new(PUSCH_Allocation_r16_numberOfSlotsTBoMS_r17)
				if err = ie.numberOfSlotsTBoMS_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode numberOfSlotsTBoMS_r17", err)
				}
			}
			// decode extendedK2_r17 optional
			if extendedK2_r17Present {
				var tmp_int_extendedK2_r17 int64
				if tmp_int_extendedK2_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 128}, false); err != nil {
					return utils.WrapError("Decode extendedK2_r17", err)
				}
				ie.extendedK2_r17 = &tmp_int_extendedK2_r17
			}
		}
	}
	return nil
}
