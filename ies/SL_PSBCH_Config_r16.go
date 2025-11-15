package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PSBCH_Config_r16 struct {
	dl_P0_PSBCH_r16    *int64                                  `lb:-16,ub:15,optional`
	dl_Alpha_PSBCH_r16 *SL_PSBCH_Config_r16_dl_Alpha_PSBCH_r16 `optional`
	dl_P0_PSBCH_r17    *int64                                  `lb:-202,ub:24,optional,ext-1`
}

func (ie *SL_PSBCH_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.dl_P0_PSBCH_r17 != nil
	preambleBits := []bool{hasExtensions, ie.dl_P0_PSBCH_r16 != nil, ie.dl_Alpha_PSBCH_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dl_P0_PSBCH_r16 != nil {
		if err = w.WriteInteger(*ie.dl_P0_PSBCH_r16, &uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode dl_P0_PSBCH_r16", err)
		}
	}
	if ie.dl_Alpha_PSBCH_r16 != nil {
		if err = ie.dl_Alpha_PSBCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dl_Alpha_PSBCH_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.dl_P0_PSBCH_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SL_PSBCH_Config_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.dl_P0_PSBCH_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode dl_P0_PSBCH_r17 optional
			if ie.dl_P0_PSBCH_r17 != nil {
				if err = extWriter.WriteInteger(*ie.dl_P0_PSBCH_r17, &uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
					return utils.WrapError("Encode dl_P0_PSBCH_r17", err)
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

func (ie *SL_PSBCH_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_P0_PSBCH_r16Present bool
	if dl_P0_PSBCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_Alpha_PSBCH_r16Present bool
	if dl_Alpha_PSBCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if dl_P0_PSBCH_r16Present {
		var tmp_int_dl_P0_PSBCH_r16 int64
		if tmp_int_dl_P0_PSBCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode dl_P0_PSBCH_r16", err)
		}
		ie.dl_P0_PSBCH_r16 = &tmp_int_dl_P0_PSBCH_r16
	}
	if dl_Alpha_PSBCH_r16Present {
		ie.dl_Alpha_PSBCH_r16 = new(SL_PSBCH_Config_r16_dl_Alpha_PSBCH_r16)
		if err = ie.dl_Alpha_PSBCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dl_Alpha_PSBCH_r16", err)
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

			dl_P0_PSBCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode dl_P0_PSBCH_r17 optional
			if dl_P0_PSBCH_r17Present {
				var tmp_int_dl_P0_PSBCH_r17 int64
				if tmp_int_dl_P0_PSBCH_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
					return utils.WrapError("Decode dl_P0_PSBCH_r17", err)
				}
				ie.dl_P0_PSBCH_r17 = &tmp_int_dl_P0_PSBCH_r17
			}
		}
	}
	return nil
}
