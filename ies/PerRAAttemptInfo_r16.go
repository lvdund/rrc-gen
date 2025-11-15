package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PerRAAttemptInfo_r16 struct {
	contentionDetected_r16   *bool                                          `optional`
	dlRSRPAboveThreshold_r16 *bool                                          `optional`
	fallbackToFourStepRA_r17 *PerRAAttemptInfo_r16_fallbackToFourStepRA_r17 `optional,ext-1`
}

func (ie *PerRAAttemptInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.fallbackToFourStepRA_r17 != nil
	preambleBits := []bool{hasExtensions, ie.contentionDetected_r16 != nil, ie.dlRSRPAboveThreshold_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.contentionDetected_r16 != nil {
		if err = w.WriteBoolean(*ie.contentionDetected_r16); err != nil {
			return utils.WrapError("Encode contentionDetected_r16", err)
		}
	}
	if ie.dlRSRPAboveThreshold_r16 != nil {
		if err = w.WriteBoolean(*ie.dlRSRPAboveThreshold_r16); err != nil {
			return utils.WrapError("Encode dlRSRPAboveThreshold_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.fallbackToFourStepRA_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap PerRAAttemptInfo_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.fallbackToFourStepRA_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode fallbackToFourStepRA_r17 optional
			if ie.fallbackToFourStepRA_r17 != nil {
				if err = ie.fallbackToFourStepRA_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode fallbackToFourStepRA_r17", err)
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

func (ie *PerRAAttemptInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var contentionDetected_r16Present bool
	if contentionDetected_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dlRSRPAboveThreshold_r16Present bool
	if dlRSRPAboveThreshold_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if contentionDetected_r16Present {
		var tmp_bool_contentionDetected_r16 bool
		if tmp_bool_contentionDetected_r16, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode contentionDetected_r16", err)
		}
		ie.contentionDetected_r16 = &tmp_bool_contentionDetected_r16
	}
	if dlRSRPAboveThreshold_r16Present {
		var tmp_bool_dlRSRPAboveThreshold_r16 bool
		if tmp_bool_dlRSRPAboveThreshold_r16, err = r.ReadBoolean(); err != nil {
			return utils.WrapError("Decode dlRSRPAboveThreshold_r16", err)
		}
		ie.dlRSRPAboveThreshold_r16 = &tmp_bool_dlRSRPAboveThreshold_r16
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

			fallbackToFourStepRA_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode fallbackToFourStepRA_r17 optional
			if fallbackToFourStepRA_r17Present {
				ie.fallbackToFourStepRA_r17 = new(PerRAAttemptInfo_r16_fallbackToFourStepRA_r17)
				if err = ie.fallbackToFourStepRA_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode fallbackToFourStepRA_r17", err)
				}
			}
		}
	}
	return nil
}
