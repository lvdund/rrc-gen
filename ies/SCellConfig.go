package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SCellConfig struct {
	sCellIndex                       SCellIndex                                `madatory`
	sCellConfigCommon                *ServingCellConfigCommon                  `optional`
	sCellConfigDedicated             *ServingCellConfig                        `optional`
	smtc                             *SSB_MTC                                  `optional,ext-1`
	sCellState_r16                   *SCellConfig_sCellState_r16               `optional,ext-2`
	secondaryDRX_GroupConfig_r16     *SCellConfig_secondaryDRX_GroupConfig_r16 `optional,ext-2`
	preConfGapStatus_r17             *uper.BitString                           `lb:maxNrofGapId_r17,ub:maxNrofGapId_r17,optional,ext-3`
	goodServingCellEvaluationBFD_r17 *GoodServingCellEvaluation_r17            `optional,ext-3`
	sCellSIB20_r17                   *SCellSIB20_r17                           `optional,ext-3,setuprelease`
}

func (ie *SCellConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.smtc != nil || ie.sCellState_r16 != nil || ie.secondaryDRX_GroupConfig_r16 != nil || ie.preConfGapStatus_r17 != nil || ie.goodServingCellEvaluationBFD_r17 != nil || ie.sCellSIB20_r17 != nil
	preambleBits := []bool{hasExtensions, ie.sCellConfigCommon != nil, ie.sCellConfigDedicated != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.sCellIndex.Encode(w); err != nil {
		return utils.WrapError("Encode sCellIndex", err)
	}
	if ie.sCellConfigCommon != nil {
		if err = ie.sCellConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode sCellConfigCommon", err)
		}
	}
	if ie.sCellConfigDedicated != nil {
		if err = ie.sCellConfigDedicated.Encode(w); err != nil {
			return utils.WrapError("Encode sCellConfigDedicated", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.smtc != nil, ie.sCellState_r16 != nil || ie.secondaryDRX_GroupConfig_r16 != nil, ie.preConfGapStatus_r17 != nil || ie.goodServingCellEvaluationBFD_r17 != nil || ie.sCellSIB20_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SCellConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.smtc != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode smtc optional
			if ie.smtc != nil {
				if err = ie.smtc.Encode(extWriter); err != nil {
					return utils.WrapError("Encode smtc", err)
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
			optionals_ext_2 := []bool{ie.sCellState_r16 != nil, ie.secondaryDRX_GroupConfig_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sCellState_r16 optional
			if ie.sCellState_r16 != nil {
				if err = ie.sCellState_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sCellState_r16", err)
				}
			}
			// encode secondaryDRX_GroupConfig_r16 optional
			if ie.secondaryDRX_GroupConfig_r16 != nil {
				if err = ie.secondaryDRX_GroupConfig_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode secondaryDRX_GroupConfig_r16", err)
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
			optionals_ext_3 := []bool{ie.preConfGapStatus_r17 != nil, ie.goodServingCellEvaluationBFD_r17 != nil, ie.sCellSIB20_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode preConfGapStatus_r17 optional
			if ie.preConfGapStatus_r17 != nil {
				if err = extWriter.WriteBitString(ie.preConfGapStatus_r17.Bytes, uint(ie.preConfGapStatus_r17.NumBits), &uper.Constraint{Lb: maxNrofGapId_r17, Ub: maxNrofGapId_r17}, false); err != nil {
					return utils.WrapError("Encode preConfGapStatus_r17", err)
				}
			}
			// encode goodServingCellEvaluationBFD_r17 optional
			if ie.goodServingCellEvaluationBFD_r17 != nil {
				if err = ie.goodServingCellEvaluationBFD_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode goodServingCellEvaluationBFD_r17", err)
				}
			}
			// encode sCellSIB20_r17 optional
			if ie.sCellSIB20_r17 != nil {
				tmp_sCellSIB20_r17 := utils.SetupRelease[*SCellSIB20_r17]{
					Setup: ie.sCellSIB20_r17,
				}
				if err = tmp_sCellSIB20_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sCellSIB20_r17", err)
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

func (ie *SCellConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var sCellConfigCommonPresent bool
	if sCellConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sCellConfigDedicatedPresent bool
	if sCellConfigDedicatedPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.sCellIndex.Decode(r); err != nil {
		return utils.WrapError("Decode sCellIndex", err)
	}
	if sCellConfigCommonPresent {
		ie.sCellConfigCommon = new(ServingCellConfigCommon)
		if err = ie.sCellConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode sCellConfigCommon", err)
		}
	}
	if sCellConfigDedicatedPresent {
		ie.sCellConfigDedicated = new(ServingCellConfig)
		if err = ie.sCellConfigDedicated.Decode(r); err != nil {
			return utils.WrapError("Decode sCellConfigDedicated", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
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

			smtcPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode smtc optional
			if smtcPresent {
				ie.smtc = new(SSB_MTC)
				if err = ie.smtc.Decode(extReader); err != nil {
					return utils.WrapError("Decode smtc", err)
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

			sCellState_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			secondaryDRX_GroupConfig_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sCellState_r16 optional
			if sCellState_r16Present {
				ie.sCellState_r16 = new(SCellConfig_sCellState_r16)
				if err = ie.sCellState_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode sCellState_r16", err)
				}
			}
			// decode secondaryDRX_GroupConfig_r16 optional
			if secondaryDRX_GroupConfig_r16Present {
				ie.secondaryDRX_GroupConfig_r16 = new(SCellConfig_secondaryDRX_GroupConfig_r16)
				if err = ie.secondaryDRX_GroupConfig_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode secondaryDRX_GroupConfig_r16", err)
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

			preConfGapStatus_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			goodServingCellEvaluationBFD_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sCellSIB20_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode preConfGapStatus_r17 optional
			if preConfGapStatus_r17Present {
				var tmp_bs_preConfGapStatus_r17 []byte
				var n_preConfGapStatus_r17 uint
				if tmp_bs_preConfGapStatus_r17, n_preConfGapStatus_r17, err = extReader.ReadBitString(&uper.Constraint{Lb: maxNrofGapId_r17, Ub: maxNrofGapId_r17}, false); err != nil {
					return utils.WrapError("Decode preConfGapStatus_r17", err)
				}
				tmp_bitstring := uper.BitString{
					Bytes:   tmp_bs_preConfGapStatus_r17,
					NumBits: uint64(n_preConfGapStatus_r17),
				}
				ie.preConfGapStatus_r17 = &tmp_bitstring
			}
			// decode goodServingCellEvaluationBFD_r17 optional
			if goodServingCellEvaluationBFD_r17Present {
				ie.goodServingCellEvaluationBFD_r17 = new(GoodServingCellEvaluation_r17)
				if err = ie.goodServingCellEvaluationBFD_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode goodServingCellEvaluationBFD_r17", err)
				}
			}
			// decode sCellSIB20_r17 optional
			if sCellSIB20_r17Present {
				tmp_sCellSIB20_r17 := utils.SetupRelease[*SCellSIB20_r17]{}
				if err = tmp_sCellSIB20_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode sCellSIB20_r17", err)
				}
				ie.sCellSIB20_r17 = tmp_sCellSIB20_r17.Setup
			}
		}
	}
	return nil
}
