package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_MeasurementsAvailable_r16 struct {
	logMeasAvailable_r16          *UE_MeasurementsAvailable_r16_logMeasAvailable_r16         `optional`
	logMeasAvailableBT_r16        *UE_MeasurementsAvailable_r16_logMeasAvailableBT_r16       `optional`
	logMeasAvailableWLAN_r16      *UE_MeasurementsAvailable_r16_logMeasAvailableWLAN_r16     `optional`
	connEstFailInfoAvailable_r16  *UE_MeasurementsAvailable_r16_connEstFailInfoAvailable_r16 `optional`
	rlf_InfoAvailable_r16         *UE_MeasurementsAvailable_r16_rlf_InfoAvailable_r16        `optional`
	successHO_InfoAvailable_r17   *UE_MeasurementsAvailable_r16_successHO_InfoAvailable_r17  `optional,ext-1`
	sigLogMeasConfigAvailable_r17 *bool                                                      `optional,ext-1`
}

func (ie *UE_MeasurementsAvailable_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.successHO_InfoAvailable_r17 != nil || ie.sigLogMeasConfigAvailable_r17 != nil
	preambleBits := []bool{hasExtensions, ie.logMeasAvailable_r16 != nil, ie.logMeasAvailableBT_r16 != nil, ie.logMeasAvailableWLAN_r16 != nil, ie.connEstFailInfoAvailable_r16 != nil, ie.rlf_InfoAvailable_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.logMeasAvailable_r16 != nil {
		if err = ie.logMeasAvailable_r16.Encode(w); err != nil {
			return utils.WrapError("Encode logMeasAvailable_r16", err)
		}
	}
	if ie.logMeasAvailableBT_r16 != nil {
		if err = ie.logMeasAvailableBT_r16.Encode(w); err != nil {
			return utils.WrapError("Encode logMeasAvailableBT_r16", err)
		}
	}
	if ie.logMeasAvailableWLAN_r16 != nil {
		if err = ie.logMeasAvailableWLAN_r16.Encode(w); err != nil {
			return utils.WrapError("Encode logMeasAvailableWLAN_r16", err)
		}
	}
	if ie.connEstFailInfoAvailable_r16 != nil {
		if err = ie.connEstFailInfoAvailable_r16.Encode(w); err != nil {
			return utils.WrapError("Encode connEstFailInfoAvailable_r16", err)
		}
	}
	if ie.rlf_InfoAvailable_r16 != nil {
		if err = ie.rlf_InfoAvailable_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rlf_InfoAvailable_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.successHO_InfoAvailable_r17 != nil || ie.sigLogMeasConfigAvailable_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap UE_MeasurementsAvailable_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.successHO_InfoAvailable_r17 != nil, ie.sigLogMeasConfigAvailable_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode successHO_InfoAvailable_r17 optional
			if ie.successHO_InfoAvailable_r17 != nil {
				if err = ie.successHO_InfoAvailable_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode successHO_InfoAvailable_r17", err)
				}
			}
			// encode sigLogMeasConfigAvailable_r17 optional
			if ie.sigLogMeasConfigAvailable_r17 != nil {
				if err = extWriter.WriteBoolean(*ie.sigLogMeasConfigAvailable_r17); err != nil {
					return utils.WrapError("Encode sigLogMeasConfigAvailable_r17", err)
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

func (ie *UE_MeasurementsAvailable_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var logMeasAvailable_r16Present bool
	if logMeasAvailable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var logMeasAvailableBT_r16Present bool
	if logMeasAvailableBT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var logMeasAvailableWLAN_r16Present bool
	if logMeasAvailableWLAN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var connEstFailInfoAvailable_r16Present bool
	if connEstFailInfoAvailable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rlf_InfoAvailable_r16Present bool
	if rlf_InfoAvailable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if logMeasAvailable_r16Present {
		ie.logMeasAvailable_r16 = new(UE_MeasurementsAvailable_r16_logMeasAvailable_r16)
		if err = ie.logMeasAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode logMeasAvailable_r16", err)
		}
	}
	if logMeasAvailableBT_r16Present {
		ie.logMeasAvailableBT_r16 = new(UE_MeasurementsAvailable_r16_logMeasAvailableBT_r16)
		if err = ie.logMeasAvailableBT_r16.Decode(r); err != nil {
			return utils.WrapError("Decode logMeasAvailableBT_r16", err)
		}
	}
	if logMeasAvailableWLAN_r16Present {
		ie.logMeasAvailableWLAN_r16 = new(UE_MeasurementsAvailable_r16_logMeasAvailableWLAN_r16)
		if err = ie.logMeasAvailableWLAN_r16.Decode(r); err != nil {
			return utils.WrapError("Decode logMeasAvailableWLAN_r16", err)
		}
	}
	if connEstFailInfoAvailable_r16Present {
		ie.connEstFailInfoAvailable_r16 = new(UE_MeasurementsAvailable_r16_connEstFailInfoAvailable_r16)
		if err = ie.connEstFailInfoAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode connEstFailInfoAvailable_r16", err)
		}
	}
	if rlf_InfoAvailable_r16Present {
		ie.rlf_InfoAvailable_r16 = new(UE_MeasurementsAvailable_r16_rlf_InfoAvailable_r16)
		if err = ie.rlf_InfoAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rlf_InfoAvailable_r16", err)
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

			successHO_InfoAvailable_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sigLogMeasConfigAvailable_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode successHO_InfoAvailable_r17 optional
			if successHO_InfoAvailable_r17Present {
				ie.successHO_InfoAvailable_r17 = new(UE_MeasurementsAvailable_r16_successHO_InfoAvailable_r17)
				if err = ie.successHO_InfoAvailable_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode successHO_InfoAvailable_r17", err)
				}
			}
			// decode sigLogMeasConfigAvailable_r17 optional
			if sigLogMeasConfigAvailable_r17Present {
				var tmp_bool_sigLogMeasConfigAvailable_r17 bool
				if tmp_bool_sigLogMeasConfigAvailable_r17, err = extReader.ReadBoolean(); err != nil {
					return utils.WrapError("Decode sigLogMeasConfigAvailable_r17", err)
				}
				ie.sigLogMeasConfigAvailable_r17 = &tmp_bool_sigLogMeasConfigAvailable_r17
			}
		}
	}
	return nil
}
