package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type InterRAT_Parameters struct {
	eutra        *EUTRA_Parameters        `optional`
	utra_FDD_r16 *UTRA_FDD_Parameters_r16 `optional,ext-1`
}

func (ie *InterRAT_Parameters) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.utra_FDD_r16 != nil
	preambleBits := []bool{hasExtensions, ie.eutra != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.eutra != nil {
		if err = ie.eutra.Encode(w); err != nil {
			return utils.WrapError("Encode eutra", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.utra_FDD_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap InterRAT_Parameters", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.utra_FDD_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode utra_FDD_r16 optional
			if ie.utra_FDD_r16 != nil {
				if err = ie.utra_FDD_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode utra_FDD_r16", err)
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

func (ie *InterRAT_Parameters) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var eutraPresent bool
	if eutraPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if eutraPresent {
		ie.eutra = new(EUTRA_Parameters)
		if err = ie.eutra.Decode(r); err != nil {
			return utils.WrapError("Decode eutra", err)
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

			utra_FDD_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode utra_FDD_r16 optional
			if utra_FDD_r16Present {
				ie.utra_FDD_r16 = new(UTRA_FDD_Parameters_r16)
				if err = ie.utra_FDD_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode utra_FDD_r16", err)
				}
			}
		}
	}
	return nil
}
