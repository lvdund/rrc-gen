package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB9 struct {
	timeInfo                 *SIB9_timeInfo         `lb:2,ub:2,optional`
	lateNonCriticalExtension *[]byte                `optional`
	referenceTimeInfo_r16    *ReferenceTimeInfo_r16 `optional,ext-1`
}

func (ie *SIB9) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.referenceTimeInfo_r16 != nil
	preambleBits := []bool{hasExtensions, ie.timeInfo != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.timeInfo != nil {
		if err = ie.timeInfo.Encode(w); err != nil {
			return utils.WrapError("Encode timeInfo", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.referenceTimeInfo_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SIB9", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.referenceTimeInfo_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode referenceTimeInfo_r16 optional
			if ie.referenceTimeInfo_r16 != nil {
				if err = ie.referenceTimeInfo_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode referenceTimeInfo_r16", err)
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

func (ie *SIB9) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var timeInfoPresent bool
	if timeInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if timeInfoPresent {
		ie.timeInfo = new(SIB9_timeInfo)
		if err = ie.timeInfo.Decode(r); err != nil {
			return utils.WrapError("Decode timeInfo", err)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
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

			referenceTimeInfo_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode referenceTimeInfo_r16 optional
			if referenceTimeInfo_r16Present {
				ie.referenceTimeInfo_r16 = new(ReferenceTimeInfo_r16)
				if err = ie.referenceTimeInfo_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode referenceTimeInfo_r16", err)
				}
			}
		}
	}
	return nil
}
