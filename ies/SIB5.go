package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB5 struct {
	carrierFreqListEUTRA          *CarrierFreqListEUTRA               `optional`
	t_ReselectionEUTRA            T_Reselection                       `madatory`
	t_ReselectionEUTRA_SF         *SpeedStateScaleFactors             `optional`
	lateNonCriticalExtension      *[]byte                             `optional`
	carrierFreqListEUTRA_v1610    *CarrierFreqListEUTRA_v1610         `optional,ext-1`
	carrierFreqListEUTRA_v1700    *CarrierFreqListEUTRA_v1700         `optional,ext-2`
	idleModeMeasVoiceFallback_r17 *SIB5_idleModeMeasVoiceFallback_r17 `optional,ext-2`
}

func (ie *SIB5) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.carrierFreqListEUTRA_v1610 != nil || ie.carrierFreqListEUTRA_v1700 != nil || ie.idleModeMeasVoiceFallback_r17 != nil
	preambleBits := []bool{hasExtensions, ie.carrierFreqListEUTRA != nil, ie.t_ReselectionEUTRA_SF != nil, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.carrierFreqListEUTRA != nil {
		if err = ie.carrierFreqListEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode carrierFreqListEUTRA", err)
		}
	}
	if err = ie.t_ReselectionEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode t_ReselectionEUTRA", err)
	}
	if ie.t_ReselectionEUTRA_SF != nil {
		if err = ie.t_ReselectionEUTRA_SF.Encode(w); err != nil {
			return utils.WrapError("Encode t_ReselectionEUTRA_SF", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.carrierFreqListEUTRA_v1610 != nil, ie.carrierFreqListEUTRA_v1700 != nil || ie.idleModeMeasVoiceFallback_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SIB5", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.carrierFreqListEUTRA_v1610 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode carrierFreqListEUTRA_v1610 optional
			if ie.carrierFreqListEUTRA_v1610 != nil {
				if err = ie.carrierFreqListEUTRA_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode carrierFreqListEUTRA_v1610", err)
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
			optionals_ext_2 := []bool{ie.carrierFreqListEUTRA_v1700 != nil, ie.idleModeMeasVoiceFallback_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode carrierFreqListEUTRA_v1700 optional
			if ie.carrierFreqListEUTRA_v1700 != nil {
				if err = ie.carrierFreqListEUTRA_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode carrierFreqListEUTRA_v1700", err)
				}
			}
			// encode idleModeMeasVoiceFallback_r17 optional
			if ie.idleModeMeasVoiceFallback_r17 != nil {
				if err = ie.idleModeMeasVoiceFallback_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode idleModeMeasVoiceFallback_r17", err)
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

func (ie *SIB5) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var carrierFreqListEUTRAPresent bool
	if carrierFreqListEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var t_ReselectionEUTRA_SFPresent bool
	if t_ReselectionEUTRA_SFPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if carrierFreqListEUTRAPresent {
		ie.carrierFreqListEUTRA = new(CarrierFreqListEUTRA)
		if err = ie.carrierFreqListEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode carrierFreqListEUTRA", err)
		}
	}
	if err = ie.t_ReselectionEUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode t_ReselectionEUTRA", err)
	}
	if t_ReselectionEUTRA_SFPresent {
		ie.t_ReselectionEUTRA_SF = new(SpeedStateScaleFactors)
		if err = ie.t_ReselectionEUTRA_SF.Decode(r); err != nil {
			return utils.WrapError("Decode t_ReselectionEUTRA_SF", err)
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

			carrierFreqListEUTRA_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode carrierFreqListEUTRA_v1610 optional
			if carrierFreqListEUTRA_v1610Present {
				ie.carrierFreqListEUTRA_v1610 = new(CarrierFreqListEUTRA_v1610)
				if err = ie.carrierFreqListEUTRA_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode carrierFreqListEUTRA_v1610", err)
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

			carrierFreqListEUTRA_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			idleModeMeasVoiceFallback_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode carrierFreqListEUTRA_v1700 optional
			if carrierFreqListEUTRA_v1700Present {
				ie.carrierFreqListEUTRA_v1700 = new(CarrierFreqListEUTRA_v1700)
				if err = ie.carrierFreqListEUTRA_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode carrierFreqListEUTRA_v1700", err)
				}
			}
			// decode idleModeMeasVoiceFallback_r17 optional
			if idleModeMeasVoiceFallback_r17Present {
				ie.idleModeMeasVoiceFallback_r17 = new(SIB5_idleModeMeasVoiceFallback_r17)
				if err = ie.idleModeMeasVoiceFallback_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode idleModeMeasVoiceFallback_r17", err)
				}
			}
		}
	}
	return nil
}
