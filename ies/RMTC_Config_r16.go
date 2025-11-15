package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RMTC_Config_r16 struct {
	rmtc_Periodicity_r16      RMTC_Config_r16_rmtc_Periodicity_r16       `madatory`
	rmtc_SubframeOffset_r16   *int64                                     `lb:0,ub:639,optional`
	measDurationSymbols_r16   RMTC_Config_r16_measDurationSymbols_r16    `madatory`
	rmtc_Frequency_r16        ARFCN_ValueNR                              `madatory`
	ref_SCS_CP_r16            RMTC_Config_r16_ref_SCS_CP_r16             `madatory`
	rmtc_Bandwidth_r17        *RMTC_Config_r16_rmtc_Bandwidth_r17        `optional,ext-1`
	measDurationSymbols_v1700 *RMTC_Config_r16_measDurationSymbols_v1700 `optional,ext-1`
	ref_SCS_CP_v1700          *RMTC_Config_r16_ref_SCS_CP_v1700          `optional,ext-1`
	tci_StateInfo_r17         *RMTC_Config_r16_tci_StateInfo_r17         `optional,ext-1`
	ref_BWPId_r17             *BWP_Id                                    `optional,ext-2`
}

func (ie *RMTC_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.rmtc_Bandwidth_r17 != nil || ie.measDurationSymbols_v1700 != nil || ie.ref_SCS_CP_v1700 != nil || ie.tci_StateInfo_r17 != nil || ie.ref_BWPId_r17 != nil
	preambleBits := []bool{hasExtensions, ie.rmtc_SubframeOffset_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.rmtc_Periodicity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode rmtc_Periodicity_r16", err)
	}
	if ie.rmtc_SubframeOffset_r16 != nil {
		if err = w.WriteInteger(*ie.rmtc_SubframeOffset_r16, &uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Encode rmtc_SubframeOffset_r16", err)
		}
	}
	if err = ie.measDurationSymbols_r16.Encode(w); err != nil {
		return utils.WrapError("Encode measDurationSymbols_r16", err)
	}
	if err = ie.rmtc_Frequency_r16.Encode(w); err != nil {
		return utils.WrapError("Encode rmtc_Frequency_r16", err)
	}
	if err = ie.ref_SCS_CP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ref_SCS_CP_r16", err)
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.rmtc_Bandwidth_r17 != nil || ie.measDurationSymbols_v1700 != nil || ie.ref_SCS_CP_v1700 != nil || ie.tci_StateInfo_r17 != nil, ie.ref_BWPId_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RMTC_Config_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.rmtc_Bandwidth_r17 != nil, ie.measDurationSymbols_v1700 != nil, ie.ref_SCS_CP_v1700 != nil, ie.tci_StateInfo_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode rmtc_Bandwidth_r17 optional
			if ie.rmtc_Bandwidth_r17 != nil {
				if err = ie.rmtc_Bandwidth_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rmtc_Bandwidth_r17", err)
				}
			}
			// encode measDurationSymbols_v1700 optional
			if ie.measDurationSymbols_v1700 != nil {
				if err = ie.measDurationSymbols_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode measDurationSymbols_v1700", err)
				}
			}
			// encode ref_SCS_CP_v1700 optional
			if ie.ref_SCS_CP_v1700 != nil {
				if err = ie.ref_SCS_CP_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ref_SCS_CP_v1700", err)
				}
			}
			// encode tci_StateInfo_r17 optional
			if ie.tci_StateInfo_r17 != nil {
				if err = ie.tci_StateInfo_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode tci_StateInfo_r17", err)
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
			optionals_ext_2 := []bool{ie.ref_BWPId_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ref_BWPId_r17 optional
			if ie.ref_BWPId_r17 != nil {
				if err = ie.ref_BWPId_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ref_BWPId_r17", err)
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

func (ie *RMTC_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var rmtc_SubframeOffset_r16Present bool
	if rmtc_SubframeOffset_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.rmtc_Periodicity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode rmtc_Periodicity_r16", err)
	}
	if rmtc_SubframeOffset_r16Present {
		var tmp_int_rmtc_SubframeOffset_r16 int64
		if tmp_int_rmtc_SubframeOffset_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 639}, false); err != nil {
			return utils.WrapError("Decode rmtc_SubframeOffset_r16", err)
		}
		ie.rmtc_SubframeOffset_r16 = &tmp_int_rmtc_SubframeOffset_r16
	}
	if err = ie.measDurationSymbols_r16.Decode(r); err != nil {
		return utils.WrapError("Decode measDurationSymbols_r16", err)
	}
	if err = ie.rmtc_Frequency_r16.Decode(r); err != nil {
		return utils.WrapError("Decode rmtc_Frequency_r16", err)
	}
	if err = ie.ref_SCS_CP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ref_SCS_CP_r16", err)
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

			rmtc_Bandwidth_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			measDurationSymbols_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ref_SCS_CP_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			tci_StateInfo_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode rmtc_Bandwidth_r17 optional
			if rmtc_Bandwidth_r17Present {
				ie.rmtc_Bandwidth_r17 = new(RMTC_Config_r16_rmtc_Bandwidth_r17)
				if err = ie.rmtc_Bandwidth_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode rmtc_Bandwidth_r17", err)
				}
			}
			// decode measDurationSymbols_v1700 optional
			if measDurationSymbols_v1700Present {
				ie.measDurationSymbols_v1700 = new(RMTC_Config_r16_measDurationSymbols_v1700)
				if err = ie.measDurationSymbols_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode measDurationSymbols_v1700", err)
				}
			}
			// decode ref_SCS_CP_v1700 optional
			if ref_SCS_CP_v1700Present {
				ie.ref_SCS_CP_v1700 = new(RMTC_Config_r16_ref_SCS_CP_v1700)
				if err = ie.ref_SCS_CP_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode ref_SCS_CP_v1700", err)
				}
			}
			// decode tci_StateInfo_r17 optional
			if tci_StateInfo_r17Present {
				ie.tci_StateInfo_r17 = new(RMTC_Config_r16_tci_StateInfo_r17)
				if err = ie.tci_StateInfo_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode tci_StateInfo_r17", err)
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

			ref_BWPId_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ref_BWPId_r17 optional
			if ref_BWPId_r17Present {
				ie.ref_BWPId_r17 = new(BWP_Id)
				if err = ie.ref_BWPId_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode ref_BWPId_r17", err)
				}
			}
		}
	}
	return nil
}
