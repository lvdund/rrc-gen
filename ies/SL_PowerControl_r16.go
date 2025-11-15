package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PowerControl_r16 struct {
	sl_MaxTransPower_r16     int64                                         `lb:-30,ub:33,madatory`
	sl_Alpha_PSSCH_PSCCH_r16 *SL_PowerControl_r16_sl_Alpha_PSSCH_PSCCH_r16 `optional`
	dl_Alpha_PSSCH_PSCCH_r16 *SL_PowerControl_r16_dl_Alpha_PSSCH_PSCCH_r16 `optional`
	sl_P0_PSSCH_PSCCH_r16    *int64                                        `lb:-16,ub:15,optional`
	dl_P0_PSSCH_PSCCH_r16    *int64                                        `lb:-16,ub:15,optional`
	dl_Alpha_PSFCH_r16       *SL_PowerControl_r16_dl_Alpha_PSFCH_r16       `optional`
	dl_P0_PSFCH_r16          *int64                                        `lb:-16,ub:15,optional`
	dl_P0_PSSCH_PSCCH_r17    *int64                                        `lb:-202,ub:24,optional,ext-1`
	sl_P0_PSSCH_PSCCH_r17    *int64                                        `lb:-202,ub:24,optional,ext-1`
	dl_P0_PSFCH_r17          *int64                                        `lb:-202,ub:24,optional,ext-1`
}

func (ie *SL_PowerControl_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.dl_P0_PSSCH_PSCCH_r17 != nil || ie.sl_P0_PSSCH_PSCCH_r17 != nil || ie.dl_P0_PSFCH_r17 != nil
	preambleBits := []bool{hasExtensions, ie.sl_Alpha_PSSCH_PSCCH_r16 != nil, ie.dl_Alpha_PSSCH_PSCCH_r16 != nil, ie.sl_P0_PSSCH_PSCCH_r16 != nil, ie.dl_P0_PSSCH_PSCCH_r16 != nil, ie.dl_Alpha_PSFCH_r16 != nil, ie.dl_P0_PSFCH_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.sl_MaxTransPower_r16, &uper.Constraint{Lb: -30, Ub: 33}, false); err != nil {
		return utils.WrapError("WriteInteger sl_MaxTransPower_r16", err)
	}
	if ie.sl_Alpha_PSSCH_PSCCH_r16 != nil {
		if err = ie.sl_Alpha_PSSCH_PSCCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_Alpha_PSSCH_PSCCH_r16", err)
		}
	}
	if ie.dl_Alpha_PSSCH_PSCCH_r16 != nil {
		if err = ie.dl_Alpha_PSSCH_PSCCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dl_Alpha_PSSCH_PSCCH_r16", err)
		}
	}
	if ie.sl_P0_PSSCH_PSCCH_r16 != nil {
		if err = w.WriteInteger(*ie.sl_P0_PSSCH_PSCCH_r16, &uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode sl_P0_PSSCH_PSCCH_r16", err)
		}
	}
	if ie.dl_P0_PSSCH_PSCCH_r16 != nil {
		if err = w.WriteInteger(*ie.dl_P0_PSSCH_PSCCH_r16, &uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode dl_P0_PSSCH_PSCCH_r16", err)
		}
	}
	if ie.dl_Alpha_PSFCH_r16 != nil {
		if err = ie.dl_Alpha_PSFCH_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dl_Alpha_PSFCH_r16", err)
		}
	}
	if ie.dl_P0_PSFCH_r16 != nil {
		if err = w.WriteInteger(*ie.dl_P0_PSFCH_r16, &uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode dl_P0_PSFCH_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.dl_P0_PSSCH_PSCCH_r17 != nil || ie.sl_P0_PSSCH_PSCCH_r17 != nil || ie.dl_P0_PSFCH_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SL_PowerControl_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.dl_P0_PSSCH_PSCCH_r17 != nil, ie.sl_P0_PSSCH_PSCCH_r17 != nil, ie.dl_P0_PSFCH_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode dl_P0_PSSCH_PSCCH_r17 optional
			if ie.dl_P0_PSSCH_PSCCH_r17 != nil {
				if err = extWriter.WriteInteger(*ie.dl_P0_PSSCH_PSCCH_r17, &uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
					return utils.WrapError("Encode dl_P0_PSSCH_PSCCH_r17", err)
				}
			}
			// encode sl_P0_PSSCH_PSCCH_r17 optional
			if ie.sl_P0_PSSCH_PSCCH_r17 != nil {
				if err = extWriter.WriteInteger(*ie.sl_P0_PSSCH_PSCCH_r17, &uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
					return utils.WrapError("Encode sl_P0_PSSCH_PSCCH_r17", err)
				}
			}
			// encode dl_P0_PSFCH_r17 optional
			if ie.dl_P0_PSFCH_r17 != nil {
				if err = extWriter.WriteInteger(*ie.dl_P0_PSFCH_r17, &uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
					return utils.WrapError("Encode dl_P0_PSFCH_r17", err)
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

func (ie *SL_PowerControl_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_Alpha_PSSCH_PSCCH_r16Present bool
	if sl_Alpha_PSSCH_PSCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_Alpha_PSSCH_PSCCH_r16Present bool
	if dl_Alpha_PSSCH_PSCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_P0_PSSCH_PSCCH_r16Present bool
	if sl_P0_PSSCH_PSCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_P0_PSSCH_PSCCH_r16Present bool
	if dl_P0_PSSCH_PSCCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_Alpha_PSFCH_r16Present bool
	if dl_Alpha_PSFCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dl_P0_PSFCH_r16Present bool
	if dl_P0_PSFCH_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_sl_MaxTransPower_r16 int64
	if tmp_int_sl_MaxTransPower_r16, err = r.ReadInteger(&uper.Constraint{Lb: -30, Ub: 33}, false); err != nil {
		return utils.WrapError("ReadInteger sl_MaxTransPower_r16", err)
	}
	ie.sl_MaxTransPower_r16 = tmp_int_sl_MaxTransPower_r16
	if sl_Alpha_PSSCH_PSCCH_r16Present {
		ie.sl_Alpha_PSSCH_PSCCH_r16 = new(SL_PowerControl_r16_sl_Alpha_PSSCH_PSCCH_r16)
		if err = ie.sl_Alpha_PSSCH_PSCCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_Alpha_PSSCH_PSCCH_r16", err)
		}
	}
	if dl_Alpha_PSSCH_PSCCH_r16Present {
		ie.dl_Alpha_PSSCH_PSCCH_r16 = new(SL_PowerControl_r16_dl_Alpha_PSSCH_PSCCH_r16)
		if err = ie.dl_Alpha_PSSCH_PSCCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dl_Alpha_PSSCH_PSCCH_r16", err)
		}
	}
	if sl_P0_PSSCH_PSCCH_r16Present {
		var tmp_int_sl_P0_PSSCH_PSCCH_r16 int64
		if tmp_int_sl_P0_PSSCH_PSCCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode sl_P0_PSSCH_PSCCH_r16", err)
		}
		ie.sl_P0_PSSCH_PSCCH_r16 = &tmp_int_sl_P0_PSSCH_PSCCH_r16
	}
	if dl_P0_PSSCH_PSCCH_r16Present {
		var tmp_int_dl_P0_PSSCH_PSCCH_r16 int64
		if tmp_int_dl_P0_PSSCH_PSCCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode dl_P0_PSSCH_PSCCH_r16", err)
		}
		ie.dl_P0_PSSCH_PSCCH_r16 = &tmp_int_dl_P0_PSSCH_PSCCH_r16
	}
	if dl_Alpha_PSFCH_r16Present {
		ie.dl_Alpha_PSFCH_r16 = new(SL_PowerControl_r16_dl_Alpha_PSFCH_r16)
		if err = ie.dl_Alpha_PSFCH_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dl_Alpha_PSFCH_r16", err)
		}
	}
	if dl_P0_PSFCH_r16Present {
		var tmp_int_dl_P0_PSFCH_r16 int64
		if tmp_int_dl_P0_PSFCH_r16, err = r.ReadInteger(&uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode dl_P0_PSFCH_r16", err)
		}
		ie.dl_P0_PSFCH_r16 = &tmp_int_dl_P0_PSFCH_r16
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

			dl_P0_PSSCH_PSCCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sl_P0_PSSCH_PSCCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			dl_P0_PSFCH_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode dl_P0_PSSCH_PSCCH_r17 optional
			if dl_P0_PSSCH_PSCCH_r17Present {
				var tmp_int_dl_P0_PSSCH_PSCCH_r17 int64
				if tmp_int_dl_P0_PSSCH_PSCCH_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
					return utils.WrapError("Decode dl_P0_PSSCH_PSCCH_r17", err)
				}
				ie.dl_P0_PSSCH_PSCCH_r17 = &tmp_int_dl_P0_PSSCH_PSCCH_r17
			}
			// decode sl_P0_PSSCH_PSCCH_r17 optional
			if sl_P0_PSSCH_PSCCH_r17Present {
				var tmp_int_sl_P0_PSSCH_PSCCH_r17 int64
				if tmp_int_sl_P0_PSSCH_PSCCH_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
					return utils.WrapError("Decode sl_P0_PSSCH_PSCCH_r17", err)
				}
				ie.sl_P0_PSSCH_PSCCH_r17 = &tmp_int_sl_P0_PSSCH_PSCCH_r17
			}
			// decode dl_P0_PSFCH_r17 optional
			if dl_P0_PSFCH_r17Present {
				var tmp_int_dl_P0_PSFCH_r17 int64
				if tmp_int_dl_P0_PSFCH_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: -202, Ub: 24}, false); err != nil {
					return utils.WrapError("Decode dl_P0_PSFCH_r17", err)
				}
				ie.dl_P0_PSFCH_r17 = &tmp_int_dl_P0_PSFCH_r17
			}
		}
	}
	return nil
}
