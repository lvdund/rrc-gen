package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RACH_ConfigGenericTwoStepRA_r16 struct {
	msgA_PRACH_ConfigurationIndex_r16    *int64                                                             `lb:0,ub:262,optional`
	msgA_RO_FDM_r16                      *RACH_ConfigGenericTwoStepRA_r16_msgA_RO_FDM_r16                   `optional`
	msgA_RO_FrequencyStart_r16           *int64                                                             `lb:0,ub:maxNrofPhysicalResourceBlocks_1,optional`
	msgA_ZeroCorrelationZoneConfig_r16   *int64                                                             `lb:0,ub:15,optional`
	msgA_PreamblePowerRampingStep_r16    *RACH_ConfigGenericTwoStepRA_r16_msgA_PreamblePowerRampingStep_r16 `optional`
	msgA_PreambleReceivedTargetPower_r16 *int64                                                             `lb:-202,ub:-60,optional`
	msgB_ResponseWindow_r16              *RACH_ConfigGenericTwoStepRA_r16_msgB_ResponseWindow_r16           `optional`
	preambleTransMax_r16                 *RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16              `optional`
	msgB_ResponseWindow_v1700            *RACH_ConfigGenericTwoStepRA_r16_msgB_ResponseWindow_v1700         `optional,ext-1`
}

func (ie *RACH_ConfigGenericTwoStepRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.msgB_ResponseWindow_v1700 != nil
	preambleBits := []bool{hasExtensions, ie.msgA_PRACH_ConfigurationIndex_r16 != nil, ie.msgA_RO_FDM_r16 != nil, ie.msgA_RO_FrequencyStart_r16 != nil, ie.msgA_ZeroCorrelationZoneConfig_r16 != nil, ie.msgA_PreamblePowerRampingStep_r16 != nil, ie.msgA_PreambleReceivedTargetPower_r16 != nil, ie.msgB_ResponseWindow_r16 != nil, ie.preambleTransMax_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.msgA_PRACH_ConfigurationIndex_r16 != nil {
		if err = w.WriteInteger(*ie.msgA_PRACH_ConfigurationIndex_r16, &uper.Constraint{Lb: 0, Ub: 262}, false); err != nil {
			return utils.WrapError("Encode msgA_PRACH_ConfigurationIndex_r16", err)
		}
	}
	if ie.msgA_RO_FDM_r16 != nil {
		if err = ie.msgA_RO_FDM_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_RO_FDM_r16", err)
		}
	}
	if ie.msgA_RO_FrequencyStart_r16 != nil {
		if err = w.WriteInteger(*ie.msgA_RO_FrequencyStart_r16, &uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Encode msgA_RO_FrequencyStart_r16", err)
		}
	}
	if ie.msgA_ZeroCorrelationZoneConfig_r16 != nil {
		if err = w.WriteInteger(*ie.msgA_ZeroCorrelationZoneConfig_r16, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode msgA_ZeroCorrelationZoneConfig_r16", err)
		}
	}
	if ie.msgA_PreamblePowerRampingStep_r16 != nil {
		if err = ie.msgA_PreamblePowerRampingStep_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgA_PreamblePowerRampingStep_r16", err)
		}
	}
	if ie.msgA_PreambleReceivedTargetPower_r16 != nil {
		if err = w.WriteInteger(*ie.msgA_PreambleReceivedTargetPower_r16, &uper.Constraint{Lb: -202, Ub: -60}, false); err != nil {
			return utils.WrapError("Encode msgA_PreambleReceivedTargetPower_r16", err)
		}
	}
	if ie.msgB_ResponseWindow_r16 != nil {
		if err = ie.msgB_ResponseWindow_r16.Encode(w); err != nil {
			return utils.WrapError("Encode msgB_ResponseWindow_r16", err)
		}
	}
	if ie.preambleTransMax_r16 != nil {
		if err = ie.preambleTransMax_r16.Encode(w); err != nil {
			return utils.WrapError("Encode preambleTransMax_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.msgB_ResponseWindow_v1700 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RACH_ConfigGenericTwoStepRA_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.msgB_ResponseWindow_v1700 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode msgB_ResponseWindow_v1700 optional
			if ie.msgB_ResponseWindow_v1700 != nil {
				if err = ie.msgB_ResponseWindow_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode msgB_ResponseWindow_v1700", err)
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

func (ie *RACH_ConfigGenericTwoStepRA_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_PRACH_ConfigurationIndex_r16Present bool
	if msgA_PRACH_ConfigurationIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_RO_FDM_r16Present bool
	if msgA_RO_FDM_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_RO_FrequencyStart_r16Present bool
	if msgA_RO_FrequencyStart_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_ZeroCorrelationZoneConfig_r16Present bool
	if msgA_ZeroCorrelationZoneConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_PreamblePowerRampingStep_r16Present bool
	if msgA_PreamblePowerRampingStep_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgA_PreambleReceivedTargetPower_r16Present bool
	if msgA_PreambleReceivedTargetPower_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var msgB_ResponseWindow_r16Present bool
	if msgB_ResponseWindow_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var preambleTransMax_r16Present bool
	if preambleTransMax_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if msgA_PRACH_ConfigurationIndex_r16Present {
		var tmp_int_msgA_PRACH_ConfigurationIndex_r16 int64
		if tmp_int_msgA_PRACH_ConfigurationIndex_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 262}, false); err != nil {
			return utils.WrapError("Decode msgA_PRACH_ConfigurationIndex_r16", err)
		}
		ie.msgA_PRACH_ConfigurationIndex_r16 = &tmp_int_msgA_PRACH_ConfigurationIndex_r16
	}
	if msgA_RO_FDM_r16Present {
		ie.msgA_RO_FDM_r16 = new(RACH_ConfigGenericTwoStepRA_r16_msgA_RO_FDM_r16)
		if err = ie.msgA_RO_FDM_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_RO_FDM_r16", err)
		}
	}
	if msgA_RO_FrequencyStart_r16Present {
		var tmp_int_msgA_RO_FrequencyStart_r16 int64
		if tmp_int_msgA_RO_FrequencyStart_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
			return utils.WrapError("Decode msgA_RO_FrequencyStart_r16", err)
		}
		ie.msgA_RO_FrequencyStart_r16 = &tmp_int_msgA_RO_FrequencyStart_r16
	}
	if msgA_ZeroCorrelationZoneConfig_r16Present {
		var tmp_int_msgA_ZeroCorrelationZoneConfig_r16 int64
		if tmp_int_msgA_ZeroCorrelationZoneConfig_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode msgA_ZeroCorrelationZoneConfig_r16", err)
		}
		ie.msgA_ZeroCorrelationZoneConfig_r16 = &tmp_int_msgA_ZeroCorrelationZoneConfig_r16
	}
	if msgA_PreamblePowerRampingStep_r16Present {
		ie.msgA_PreamblePowerRampingStep_r16 = new(RACH_ConfigGenericTwoStepRA_r16_msgA_PreamblePowerRampingStep_r16)
		if err = ie.msgA_PreamblePowerRampingStep_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgA_PreamblePowerRampingStep_r16", err)
		}
	}
	if msgA_PreambleReceivedTargetPower_r16Present {
		var tmp_int_msgA_PreambleReceivedTargetPower_r16 int64
		if tmp_int_msgA_PreambleReceivedTargetPower_r16, err = r.ReadInteger(&uper.Constraint{Lb: -202, Ub: -60}, false); err != nil {
			return utils.WrapError("Decode msgA_PreambleReceivedTargetPower_r16", err)
		}
		ie.msgA_PreambleReceivedTargetPower_r16 = &tmp_int_msgA_PreambleReceivedTargetPower_r16
	}
	if msgB_ResponseWindow_r16Present {
		ie.msgB_ResponseWindow_r16 = new(RACH_ConfigGenericTwoStepRA_r16_msgB_ResponseWindow_r16)
		if err = ie.msgB_ResponseWindow_r16.Decode(r); err != nil {
			return utils.WrapError("Decode msgB_ResponseWindow_r16", err)
		}
	}
	if preambleTransMax_r16Present {
		ie.preambleTransMax_r16 = new(RACH_ConfigGenericTwoStepRA_r16_preambleTransMax_r16)
		if err = ie.preambleTransMax_r16.Decode(r); err != nil {
			return utils.WrapError("Decode preambleTransMax_r16", err)
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

			msgB_ResponseWindow_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode msgB_ResponseWindow_v1700 optional
			if msgB_ResponseWindow_v1700Present {
				ie.msgB_ResponseWindow_v1700 = new(RACH_ConfigGenericTwoStepRA_r16_msgB_ResponseWindow_v1700)
				if err = ie.msgB_ResponseWindow_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode msgB_ResponseWindow_v1700", err)
				}
			}
		}
	}
	return nil
}
