package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB4 struct {
	interFreqCarrierFreqList       InterFreqCarrierFreqList        `madatory`
	lateNonCriticalExtension       *[]byte                         `optional`
	interFreqCarrierFreqList_v1610 *InterFreqCarrierFreqList_v1610 `optional,ext-1`
	interFreqCarrierFreqList_v1700 *InterFreqCarrierFreqList_v1700 `optional,ext-2`
	interFreqCarrierFreqList_v1720 *InterFreqCarrierFreqList_v1720 `optional,ext-3`
	interFreqCarrierFreqList_v1730 *InterFreqCarrierFreqList_v1730 `optional,ext-4`
}

func (ie *SIB4) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.interFreqCarrierFreqList_v1610 != nil || ie.interFreqCarrierFreqList_v1700 != nil || ie.interFreqCarrierFreqList_v1720 != nil || ie.interFreqCarrierFreqList_v1730 != nil
	preambleBits := []bool{hasExtensions, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.interFreqCarrierFreqList.Encode(w); err != nil {
		return utils.WrapError("Encode interFreqCarrierFreqList", err)
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 4 bits for 4 extension groups
		extBitmap := []bool{ie.interFreqCarrierFreqList_v1610 != nil, ie.interFreqCarrierFreqList_v1700 != nil, ie.interFreqCarrierFreqList_v1720 != nil, ie.interFreqCarrierFreqList_v1730 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SIB4", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.interFreqCarrierFreqList_v1610 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode interFreqCarrierFreqList_v1610 optional
			if ie.interFreqCarrierFreqList_v1610 != nil {
				if err = ie.interFreqCarrierFreqList_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode interFreqCarrierFreqList_v1610", err)
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
			optionals_ext_2 := []bool{ie.interFreqCarrierFreqList_v1700 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode interFreqCarrierFreqList_v1700 optional
			if ie.interFreqCarrierFreqList_v1700 != nil {
				if err = ie.interFreqCarrierFreqList_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode interFreqCarrierFreqList_v1700", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.interFreqCarrierFreqList_v1720 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode interFreqCarrierFreqList_v1720 optional
			if ie.interFreqCarrierFreqList_v1720 != nil {
				if err = ie.interFreqCarrierFreqList_v1720.Encode(extWriter); err != nil {
					return utils.WrapError("Encode interFreqCarrierFreqList_v1720", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 4
		if extBitmap[3] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 4
			optionals_ext_4 := []bool{ie.interFreqCarrierFreqList_v1730 != nil}
			for _, bit := range optionals_ext_4 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode interFreqCarrierFreqList_v1730 optional
			if ie.interFreqCarrierFreqList_v1730 != nil {
				if err = ie.interFreqCarrierFreqList_v1730.Encode(extWriter); err != nil {
					return utils.WrapError("Encode interFreqCarrierFreqList_v1730", err)
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

func (ie *SIB4) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.interFreqCarrierFreqList.Decode(r); err != nil {
		return utils.WrapError("Decode interFreqCarrierFreqList", err)
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}

	if extensionBit {
		// Read extension bitmap: 4 bits for 4 extension groups
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

			interFreqCarrierFreqList_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode interFreqCarrierFreqList_v1610 optional
			if interFreqCarrierFreqList_v1610Present {
				ie.interFreqCarrierFreqList_v1610 = new(InterFreqCarrierFreqList_v1610)
				if err = ie.interFreqCarrierFreqList_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode interFreqCarrierFreqList_v1610", err)
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

			interFreqCarrierFreqList_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode interFreqCarrierFreqList_v1700 optional
			if interFreqCarrierFreqList_v1700Present {
				ie.interFreqCarrierFreqList_v1700 = new(InterFreqCarrierFreqList_v1700)
				if err = ie.interFreqCarrierFreqList_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode interFreqCarrierFreqList_v1700", err)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			interFreqCarrierFreqList_v1720Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode interFreqCarrierFreqList_v1720 optional
			if interFreqCarrierFreqList_v1720Present {
				ie.interFreqCarrierFreqList_v1720 = new(InterFreqCarrierFreqList_v1720)
				if err = ie.interFreqCarrierFreqList_v1720.Decode(extReader); err != nil {
					return utils.WrapError("Decode interFreqCarrierFreqList_v1720", err)
				}
			}
		}
		// decode extension group 4
		if len(extBitmap) > 3 && extBitmap[3] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			interFreqCarrierFreqList_v1730Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode interFreqCarrierFreqList_v1730 optional
			if interFreqCarrierFreqList_v1730Present {
				ie.interFreqCarrierFreqList_v1730 = new(InterFreqCarrierFreqList_v1730)
				if err = ie.interFreqCarrierFreqList_v1730.Decode(extReader); err != nil {
					return utils.WrapError("Decode interFreqCarrierFreqList_v1730", err)
				}
			}
		}
	}
	return nil
}
