package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type IMS_ParametersCommon struct {
	voiceOverEUTRA_5GC             *IMS_ParametersCommon_voiceOverEUTRA_5GC             `optional`
	voiceOverSCG_BearerEUTRA_5GC   *IMS_ParametersCommon_voiceOverSCG_BearerEUTRA_5GC   `optional,ext-1`
	voiceFallbackIndicationEPS_r16 *IMS_ParametersCommon_voiceFallbackIndicationEPS_r16 `optional,ext-2`
}

func (ie *IMS_ParametersCommon) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.voiceOverSCG_BearerEUTRA_5GC != nil || ie.voiceFallbackIndicationEPS_r16 != nil
	preambleBits := []bool{hasExtensions, ie.voiceOverEUTRA_5GC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.voiceOverEUTRA_5GC != nil {
		if err = ie.voiceOverEUTRA_5GC.Encode(w); err != nil {
			return utils.WrapError("Encode voiceOverEUTRA_5GC", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.voiceOverSCG_BearerEUTRA_5GC != nil, ie.voiceFallbackIndicationEPS_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap IMS_ParametersCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.voiceOverSCG_BearerEUTRA_5GC != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode voiceOverSCG_BearerEUTRA_5GC optional
			if ie.voiceOverSCG_BearerEUTRA_5GC != nil {
				if err = ie.voiceOverSCG_BearerEUTRA_5GC.Encode(extWriter); err != nil {
					return utils.WrapError("Encode voiceOverSCG_BearerEUTRA_5GC", err)
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
			optionals_ext_2 := []bool{ie.voiceFallbackIndicationEPS_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode voiceFallbackIndicationEPS_r16 optional
			if ie.voiceFallbackIndicationEPS_r16 != nil {
				if err = ie.voiceFallbackIndicationEPS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode voiceFallbackIndicationEPS_r16", err)
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

func (ie *IMS_ParametersCommon) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var voiceOverEUTRA_5GCPresent bool
	if voiceOverEUTRA_5GCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if voiceOverEUTRA_5GCPresent {
		ie.voiceOverEUTRA_5GC = new(IMS_ParametersCommon_voiceOverEUTRA_5GC)
		if err = ie.voiceOverEUTRA_5GC.Decode(r); err != nil {
			return utils.WrapError("Decode voiceOverEUTRA_5GC", err)
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

			voiceOverSCG_BearerEUTRA_5GCPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode voiceOverSCG_BearerEUTRA_5GC optional
			if voiceOverSCG_BearerEUTRA_5GCPresent {
				ie.voiceOverSCG_BearerEUTRA_5GC = new(IMS_ParametersCommon_voiceOverSCG_BearerEUTRA_5GC)
				if err = ie.voiceOverSCG_BearerEUTRA_5GC.Decode(extReader); err != nil {
					return utils.WrapError("Decode voiceOverSCG_BearerEUTRA_5GC", err)
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

			voiceFallbackIndicationEPS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode voiceFallbackIndicationEPS_r16 optional
			if voiceFallbackIndicationEPS_r16Present {
				ie.voiceFallbackIndicationEPS_r16 = new(IMS_ParametersCommon_voiceFallbackIndicationEPS_r16)
				if err = ie.voiceFallbackIndicationEPS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode voiceFallbackIndicationEPS_r16", err)
				}
			}
		}
	}
	return nil
}
