package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_TimeDomainResourceAllocation_r16 struct {
	k0_r16                   *int64                                                         `lb:0,ub:32,optional`
	mappingType_r16          PDSCH_TimeDomainResourceAllocation_r16_mappingType_r16         `madatory`
	startSymbolAndLength_r16 int64                                                          `lb:0,ub:127,madatory`
	repetitionNumber_r16     *PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_r16   `optional`
	k0_v1710                 *int64                                                         `lb:33,ub:128,optional,ext-1`
	repetitionNumber_v1730   *PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_v1730 `optional,ext-2`
}

func (ie *PDSCH_TimeDomainResourceAllocation_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.k0_v1710 != nil || ie.repetitionNumber_v1730 != nil
	preambleBits := []bool{hasExtensions, ie.k0_r16 != nil, ie.repetitionNumber_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.k0_r16 != nil {
		if err = w.WriteInteger(*ie.k0_r16, &uper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode k0_r16", err)
		}
	}
	if err = ie.mappingType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode mappingType_r16", err)
	}
	if err = w.WriteInteger(ie.startSymbolAndLength_r16, &uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
		return utils.WrapError("WriteInteger startSymbolAndLength_r16", err)
	}
	if ie.repetitionNumber_r16 != nil {
		if err = ie.repetitionNumber_r16.Encode(w); err != nil {
			return utils.WrapError("Encode repetitionNumber_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.k0_v1710 != nil, ie.repetitionNumber_v1730 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PDSCH_TimeDomainResourceAllocation_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.k0_v1710 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode k0_v1710 optional
			if ie.k0_v1710 != nil {
				if err = extWriter.WriteInteger(*ie.k0_v1710, &uper.Constraint{Lb: 33, Ub: 128}, false); err != nil {
					return utils.WrapError("Encode k0_v1710", err)
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
			optionals_ext_2 := []bool{ie.repetitionNumber_v1730 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode repetitionNumber_v1730 optional
			if ie.repetitionNumber_v1730 != nil {
				if err = ie.repetitionNumber_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode repetitionNumber_v1730", err)
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

func (ie *PDSCH_TimeDomainResourceAllocation_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var k0_r16Present bool
	if k0_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var repetitionNumber_r16Present bool
	if repetitionNumber_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if k0_r16Present {
		var tmp_int_k0_r16 int64
		if tmp_int_k0_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode k0_r16", err)
		}
		ie.k0_r16 = &tmp_int_k0_r16
	}
	if err = ie.mappingType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode mappingType_r16", err)
	}
	var tmp_int_startSymbolAndLength_r16 int64
	if tmp_int_startSymbolAndLength_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
		return utils.WrapError("ReadInteger startSymbolAndLength_r16", err)
	}
	ie.startSymbolAndLength_r16 = tmp_int_startSymbolAndLength_r16
	if repetitionNumber_r16Present {
		ie.repetitionNumber_r16 = new(PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_r16)
		if err = ie.repetitionNumber_r16.Decode(r); err != nil {
			return utils.WrapError("Decode repetitionNumber_r16", err)
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

			k0_v1710Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode k0_v1710 optional
			if k0_v1710Present {
				var tmp_int_k0_v1710 int64
				if tmp_int_k0_v1710, err = extReader.ReadInteger(&uper.Constraint{Lb: 33, Ub: 128}, false); err != nil {
					return utils.WrapError("Decode k0_v1710", err)
				}
				ie.k0_v1710 = &tmp_int_k0_v1710
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			repetitionNumber_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode repetitionNumber_v1730 optional
			if repetitionNumber_v1730Present {
				ie.repetitionNumber_v1730 = new(PDSCH_TimeDomainResourceAllocation_r16_repetitionNumber_v1730)
				if err = ie.repetitionNumber_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode repetitionNumber_v1730", err)
				}
			}
		}
	}
	return nil
}
