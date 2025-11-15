package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PSSCH_TxConfig_r16 struct {
	sl_TypeTxSync_r16             *SL_TypeTxSync_r16                         `optional`
	sl_ThresUE_Speed_r16          SL_PSSCH_TxConfig_r16_sl_ThresUE_Speed_r16 `madatory`
	sl_ParametersAboveThres_r16   SL_PSSCH_TxParameters_r16                  `madatory`
	sl_ParametersBelowThres_r16   SL_PSSCH_TxParameters_r16                  `madatory`
	sl_ParametersAboveThres_v1650 *SL_MinMaxMCS_List_r16                     `optional,ext-1`
	sl_ParametersBelowThres_v1650 *SL_MinMaxMCS_List_r16                     `optional,ext-1`
}

func (ie *SL_PSSCH_TxConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.sl_ParametersAboveThres_v1650 != nil || ie.sl_ParametersBelowThres_v1650 != nil
	preambleBits := []bool{hasExtensions, ie.sl_TypeTxSync_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_TypeTxSync_r16 != nil {
		if err = ie.sl_TypeTxSync_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TypeTxSync_r16", err)
		}
	}
	if err = ie.sl_ThresUE_Speed_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_ThresUE_Speed_r16", err)
	}
	if err = ie.sl_ParametersAboveThres_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_ParametersAboveThres_r16", err)
	}
	if err = ie.sl_ParametersBelowThres_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_ParametersBelowThres_r16", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.sl_ParametersAboveThres_v1650 != nil || ie.sl_ParametersBelowThres_v1650 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SL_PSSCH_TxConfig_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.sl_ParametersAboveThres_v1650 != nil, ie.sl_ParametersBelowThres_v1650 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode sl_ParametersAboveThres_v1650 optional
			if ie.sl_ParametersAboveThres_v1650 != nil {
				if err = ie.sl_ParametersAboveThres_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_ParametersAboveThres_v1650", err)
				}
			}
			// encode sl_ParametersBelowThres_v1650 optional
			if ie.sl_ParametersBelowThres_v1650 != nil {
				if err = ie.sl_ParametersBelowThres_v1650.Encode(extWriter); err != nil {
					return utils.WrapError("Encode sl_ParametersBelowThres_v1650", err)
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

func (ie *SL_PSSCH_TxConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_TypeTxSync_r16Present bool
	if sl_TypeTxSync_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_TypeTxSync_r16Present {
		ie.sl_TypeTxSync_r16 = new(SL_TypeTxSync_r16)
		if err = ie.sl_TypeTxSync_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TypeTxSync_r16", err)
		}
	}
	if err = ie.sl_ThresUE_Speed_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_ThresUE_Speed_r16", err)
	}
	if err = ie.sl_ParametersAboveThres_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_ParametersAboveThres_r16", err)
	}
	if err = ie.sl_ParametersBelowThres_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_ParametersBelowThres_r16", err)
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

			sl_ParametersAboveThres_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			sl_ParametersBelowThres_v1650Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode sl_ParametersAboveThres_v1650 optional
			if sl_ParametersAboveThres_v1650Present {
				ie.sl_ParametersAboveThres_v1650 = new(SL_MinMaxMCS_List_r16)
				if err = ie.sl_ParametersAboveThres_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_ParametersAboveThres_v1650", err)
				}
			}
			// decode sl_ParametersBelowThres_v1650 optional
			if sl_ParametersBelowThres_v1650Present {
				ie.sl_ParametersBelowThres_v1650 = new(SL_MinMaxMCS_List_r16)
				if err = ie.sl_ParametersBelowThres_v1650.Decode(extReader); err != nil {
					return utils.WrapError("Decode sl_ParametersBelowThres_v1650", err)
				}
			}
		}
	}
	return nil
}
