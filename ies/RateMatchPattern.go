package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RateMatchPattern struct {
	rateMatchPatternId     RateMatchPatternId            `madatory`
	patternType            *RateMatchPattern_patternType `lb:275,ub:275,optional`
	subcarrierSpacing      *SubcarrierSpacing            `optional,ext`
	dummy                  RateMatchPattern_dummy        `madatory,ext`
	controlResourceSet_r16 *ControlResourceSetId_r16     `optional,ext-1`
}

func (ie *RateMatchPattern) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.controlResourceSet_r16 != nil
	preambleBits := []bool{hasExtensions, ie.patternType != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.rateMatchPatternId.Encode(w); err != nil {
		return utils.WrapError("Encode rateMatchPatternId", err)
	}
	if ie.patternType != nil {
		if err = ie.patternType.Encode(w); err != nil {
			return utils.WrapError("Encode patternType", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.controlResourceSet_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RateMatchPattern", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.controlResourceSet_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode controlResourceSet_r16 optional
			if ie.controlResourceSet_r16 != nil {
				if err = ie.controlResourceSet_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode controlResourceSet_r16", err)
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

func (ie *RateMatchPattern) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var patternTypePresent bool
	if patternTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.rateMatchPatternId.Decode(r); err != nil {
		return utils.WrapError("Decode rateMatchPatternId", err)
	}
	if patternTypePresent {
		ie.patternType = new(RateMatchPattern_patternType)
		if err = ie.patternType.Decode(r); err != nil {
			return utils.WrapError("Decode patternType", err)
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

			controlResourceSet_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode controlResourceSet_r16 optional
			if controlResourceSet_r16Present {
				ie.controlResourceSet_r16 = new(ControlResourceSetId_r16)
				if err = ie.controlResourceSet_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode controlResourceSet_r16", err)
				}
			}
		}
	}
	return nil
}
