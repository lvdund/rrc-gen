package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ConfigRestrictModReqSCG struct {
	requestedBC_MRDC                   *BandCombinationInfoSN `optional`
	requestedP_MaxFR1                  *P_Max                 `optional`
	requestedPDCCH_BlindDetectionSCG   *int64                 `lb:1,ub:15,optional,ext-1`
	requestedP_MaxEUTRA                *P_Max                 `optional,ext-1`
	requestedP_MaxFR2_r16              *P_Max                 `optional,ext-2`
	requestedMaxInterFreqMeasIdSCG_r16 *int64                 `lb:1,ub:maxMeasIdentitiesMN,optional,ext-2`
	requestedMaxIntraFreqMeasIdSCG_r16 *int64                 `lb:1,ub:maxMeasIdentitiesMN,optional,ext-2`
	requestedToffset_r16               *T_Offset_r16          `optional,ext-2`
}

func (ie *ConfigRestrictModReqSCG) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.requestedPDCCH_BlindDetectionSCG != nil || ie.requestedP_MaxEUTRA != nil || ie.requestedP_MaxFR2_r16 != nil || ie.requestedMaxInterFreqMeasIdSCG_r16 != nil || ie.requestedMaxIntraFreqMeasIdSCG_r16 != nil || ie.requestedToffset_r16 != nil
	preambleBits := []bool{hasExtensions, ie.requestedBC_MRDC != nil, ie.requestedP_MaxFR1 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.requestedBC_MRDC != nil {
		if err = ie.requestedBC_MRDC.Encode(w); err != nil {
			return utils.WrapError("Encode requestedBC_MRDC", err)
		}
	}
	if ie.requestedP_MaxFR1 != nil {
		if err = ie.requestedP_MaxFR1.Encode(w); err != nil {
			return utils.WrapError("Encode requestedP_MaxFR1", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.requestedPDCCH_BlindDetectionSCG != nil || ie.requestedP_MaxEUTRA != nil, ie.requestedP_MaxFR2_r16 != nil || ie.requestedMaxInterFreqMeasIdSCG_r16 != nil || ie.requestedMaxIntraFreqMeasIdSCG_r16 != nil || ie.requestedToffset_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap ConfigRestrictModReqSCG", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.requestedPDCCH_BlindDetectionSCG != nil, ie.requestedP_MaxEUTRA != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode requestedPDCCH_BlindDetectionSCG optional
			if ie.requestedPDCCH_BlindDetectionSCG != nil {
				if err = extWriter.WriteInteger(*ie.requestedPDCCH_BlindDetectionSCG, &uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
					return utils.WrapError("Encode requestedPDCCH_BlindDetectionSCG", err)
				}
			}
			// encode requestedP_MaxEUTRA optional
			if ie.requestedP_MaxEUTRA != nil {
				if err = ie.requestedP_MaxEUTRA.Encode(extWriter); err != nil {
					return utils.WrapError("Encode requestedP_MaxEUTRA", err)
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
			optionals_ext_2 := []bool{ie.requestedP_MaxFR2_r16 != nil, ie.requestedMaxInterFreqMeasIdSCG_r16 != nil, ie.requestedMaxIntraFreqMeasIdSCG_r16 != nil, ie.requestedToffset_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode requestedP_MaxFR2_r16 optional
			if ie.requestedP_MaxFR2_r16 != nil {
				if err = ie.requestedP_MaxFR2_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode requestedP_MaxFR2_r16", err)
				}
			}
			// encode requestedMaxInterFreqMeasIdSCG_r16 optional
			if ie.requestedMaxInterFreqMeasIdSCG_r16 != nil {
				if err = extWriter.WriteInteger(*ie.requestedMaxInterFreqMeasIdSCG_r16, &uper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
					return utils.WrapError("Encode requestedMaxInterFreqMeasIdSCG_r16", err)
				}
			}
			// encode requestedMaxIntraFreqMeasIdSCG_r16 optional
			if ie.requestedMaxIntraFreqMeasIdSCG_r16 != nil {
				if err = extWriter.WriteInteger(*ie.requestedMaxIntraFreqMeasIdSCG_r16, &uper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
					return utils.WrapError("Encode requestedMaxIntraFreqMeasIdSCG_r16", err)
				}
			}
			// encode requestedToffset_r16 optional
			if ie.requestedToffset_r16 != nil {
				if err = ie.requestedToffset_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode requestedToffset_r16", err)
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

func (ie *ConfigRestrictModReqSCG) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var requestedBC_MRDCPresent bool
	if requestedBC_MRDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var requestedP_MaxFR1Present bool
	if requestedP_MaxFR1Present, err = r.ReadBool(); err != nil {
		return err
	}
	if requestedBC_MRDCPresent {
		ie.requestedBC_MRDC = new(BandCombinationInfoSN)
		if err = ie.requestedBC_MRDC.Decode(r); err != nil {
			return utils.WrapError("Decode requestedBC_MRDC", err)
		}
	}
	if requestedP_MaxFR1Present {
		ie.requestedP_MaxFR1 = new(P_Max)
		if err = ie.requestedP_MaxFR1.Decode(r); err != nil {
			return utils.WrapError("Decode requestedP_MaxFR1", err)
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

			requestedPDCCH_BlindDetectionSCGPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			requestedP_MaxEUTRAPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode requestedPDCCH_BlindDetectionSCG optional
			if requestedPDCCH_BlindDetectionSCGPresent {
				var tmp_int_requestedPDCCH_BlindDetectionSCG int64
				if tmp_int_requestedPDCCH_BlindDetectionSCG, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
					return utils.WrapError("Decode requestedPDCCH_BlindDetectionSCG", err)
				}
				ie.requestedPDCCH_BlindDetectionSCG = &tmp_int_requestedPDCCH_BlindDetectionSCG
			}
			// decode requestedP_MaxEUTRA optional
			if requestedP_MaxEUTRAPresent {
				ie.requestedP_MaxEUTRA = new(P_Max)
				if err = ie.requestedP_MaxEUTRA.Decode(extReader); err != nil {
					return utils.WrapError("Decode requestedP_MaxEUTRA", err)
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

			requestedP_MaxFR2_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			requestedMaxInterFreqMeasIdSCG_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			requestedMaxIntraFreqMeasIdSCG_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			requestedToffset_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode requestedP_MaxFR2_r16 optional
			if requestedP_MaxFR2_r16Present {
				ie.requestedP_MaxFR2_r16 = new(P_Max)
				if err = ie.requestedP_MaxFR2_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode requestedP_MaxFR2_r16", err)
				}
			}
			// decode requestedMaxInterFreqMeasIdSCG_r16 optional
			if requestedMaxInterFreqMeasIdSCG_r16Present {
				var tmp_int_requestedMaxInterFreqMeasIdSCG_r16 int64
				if tmp_int_requestedMaxInterFreqMeasIdSCG_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
					return utils.WrapError("Decode requestedMaxInterFreqMeasIdSCG_r16", err)
				}
				ie.requestedMaxInterFreqMeasIdSCG_r16 = &tmp_int_requestedMaxInterFreqMeasIdSCG_r16
			}
			// decode requestedMaxIntraFreqMeasIdSCG_r16 optional
			if requestedMaxIntraFreqMeasIdSCG_r16Present {
				var tmp_int_requestedMaxIntraFreqMeasIdSCG_r16 int64
				if tmp_int_requestedMaxIntraFreqMeasIdSCG_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxMeasIdentitiesMN}, false); err != nil {
					return utils.WrapError("Decode requestedMaxIntraFreqMeasIdSCG_r16", err)
				}
				ie.requestedMaxIntraFreqMeasIdSCG_r16 = &tmp_int_requestedMaxIntraFreqMeasIdSCG_r16
			}
			// decode requestedToffset_r16 optional
			if requestedToffset_r16Present {
				ie.requestedToffset_r16 = new(T_Offset_r16)
				if err = ie.requestedToffset_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode requestedToffset_r16", err)
				}
			}
		}
	}
	return nil
}
