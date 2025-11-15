package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RACH_ConfigGeneric struct {
	prach_ConfigurationIndex                 int64                                                        `lb:0,ub:255,madatory`
	msg1_FDM                                 RACH_ConfigGeneric_msg1_FDM                                  `madatory`
	msg1_FrequencyStart                      int64                                                        `lb:0,ub:maxNrofPhysicalResourceBlocks_1,madatory`
	zeroCorrelationZoneConfig                int64                                                        `lb:0,ub:15,madatory`
	preambleReceivedTargetPower              int64                                                        `lb:-202,ub:-60,madatory`
	preambleTransMax                         RACH_ConfigGeneric_preambleTransMax                          `madatory`
	powerRampingStep                         RACH_ConfigGeneric_powerRampingStep                          `madatory`
	ra_ResponseWindow                        RACH_ConfigGeneric_ra_ResponseWindow                         `madatory`
	prach_ConfigurationPeriodScaling_IAB_r16 *RACH_ConfigGeneric_prach_ConfigurationPeriodScaling_IAB_r16 `optional,ext-1`
	prach_ConfigurationFrameOffset_IAB_r16   *int64                                                       `lb:0,ub:63,optional,ext-1`
	prach_ConfigurationSOffset_IAB_r16       *int64                                                       `lb:0,ub:39,optional,ext-1`
	ra_ResponseWindow_v1610                  *RACH_ConfigGeneric_ra_ResponseWindow_v1610                  `optional,ext-1`
	prach_ConfigurationIndex_v1610           *int64                                                       `lb:256,ub:262,optional,ext-1`
	ra_ResponseWindow_v1700                  *RACH_ConfigGeneric_ra_ResponseWindow_v1700                  `optional,ext-2`
}

func (ie *RACH_ConfigGeneric) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.prach_ConfigurationPeriodScaling_IAB_r16 != nil || ie.prach_ConfigurationFrameOffset_IAB_r16 != nil || ie.prach_ConfigurationSOffset_IAB_r16 != nil || ie.ra_ResponseWindow_v1610 != nil || ie.prach_ConfigurationIndex_v1610 != nil || ie.ra_ResponseWindow_v1700 != nil
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.prach_ConfigurationIndex, &uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return utils.WrapError("WriteInteger prach_ConfigurationIndex", err)
	}
	if err = ie.msg1_FDM.Encode(w); err != nil {
		return utils.WrapError("Encode msg1_FDM", err)
	}
	if err = w.WriteInteger(ie.msg1_FrequencyStart, &uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
		return utils.WrapError("WriteInteger msg1_FrequencyStart", err)
	}
	if err = w.WriteInteger(ie.zeroCorrelationZoneConfig, &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger zeroCorrelationZoneConfig", err)
	}
	if err = w.WriteInteger(ie.preambleReceivedTargetPower, &uper.Constraint{Lb: -202, Ub: -60}, false); err != nil {
		return utils.WrapError("WriteInteger preambleReceivedTargetPower", err)
	}
	if err = ie.preambleTransMax.Encode(w); err != nil {
		return utils.WrapError("Encode preambleTransMax", err)
	}
	if err = ie.powerRampingStep.Encode(w); err != nil {
		return utils.WrapError("Encode powerRampingStep", err)
	}
	if err = ie.ra_ResponseWindow.Encode(w); err != nil {
		return utils.WrapError("Encode ra_ResponseWindow", err)
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.prach_ConfigurationPeriodScaling_IAB_r16 != nil || ie.prach_ConfigurationFrameOffset_IAB_r16 != nil || ie.prach_ConfigurationSOffset_IAB_r16 != nil || ie.ra_ResponseWindow_v1610 != nil || ie.prach_ConfigurationIndex_v1610 != nil, ie.ra_ResponseWindow_v1700 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RACH_ConfigGeneric", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.prach_ConfigurationPeriodScaling_IAB_r16 != nil, ie.prach_ConfigurationFrameOffset_IAB_r16 != nil, ie.prach_ConfigurationSOffset_IAB_r16 != nil, ie.ra_ResponseWindow_v1610 != nil, ie.prach_ConfigurationIndex_v1610 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode prach_ConfigurationPeriodScaling_IAB_r16 optional
			if ie.prach_ConfigurationPeriodScaling_IAB_r16 != nil {
				if err = ie.prach_ConfigurationPeriodScaling_IAB_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode prach_ConfigurationPeriodScaling_IAB_r16", err)
				}
			}
			// encode prach_ConfigurationFrameOffset_IAB_r16 optional
			if ie.prach_ConfigurationFrameOffset_IAB_r16 != nil {
				if err = extWriter.WriteInteger(*ie.prach_ConfigurationFrameOffset_IAB_r16, &uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
					return utils.WrapError("Encode prach_ConfigurationFrameOffset_IAB_r16", err)
				}
			}
			// encode prach_ConfigurationSOffset_IAB_r16 optional
			if ie.prach_ConfigurationSOffset_IAB_r16 != nil {
				if err = extWriter.WriteInteger(*ie.prach_ConfigurationSOffset_IAB_r16, &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
					return utils.WrapError("Encode prach_ConfigurationSOffset_IAB_r16", err)
				}
			}
			// encode ra_ResponseWindow_v1610 optional
			if ie.ra_ResponseWindow_v1610 != nil {
				if err = ie.ra_ResponseWindow_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ra_ResponseWindow_v1610", err)
				}
			}
			// encode prach_ConfigurationIndex_v1610 optional
			if ie.prach_ConfigurationIndex_v1610 != nil {
				if err = extWriter.WriteInteger(*ie.prach_ConfigurationIndex_v1610, &uper.Constraint{Lb: 256, Ub: 262}, false); err != nil {
					return utils.WrapError("Encode prach_ConfigurationIndex_v1610", err)
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
			optionals_ext_2 := []bool{ie.ra_ResponseWindow_v1700 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode ra_ResponseWindow_v1700 optional
			if ie.ra_ResponseWindow_v1700 != nil {
				if err = ie.ra_ResponseWindow_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode ra_ResponseWindow_v1700", err)
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

func (ie *RACH_ConfigGeneric) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_prach_ConfigurationIndex int64
	if tmp_int_prach_ConfigurationIndex, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return utils.WrapError("ReadInteger prach_ConfigurationIndex", err)
	}
	ie.prach_ConfigurationIndex = tmp_int_prach_ConfigurationIndex
	if err = ie.msg1_FDM.Decode(r); err != nil {
		return utils.WrapError("Decode msg1_FDM", err)
	}
	var tmp_int_msg1_FrequencyStart int64
	if tmp_int_msg1_FrequencyStart, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
		return utils.WrapError("ReadInteger msg1_FrequencyStart", err)
	}
	ie.msg1_FrequencyStart = tmp_int_msg1_FrequencyStart
	var tmp_int_zeroCorrelationZoneConfig int64
	if tmp_int_zeroCorrelationZoneConfig, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger zeroCorrelationZoneConfig", err)
	}
	ie.zeroCorrelationZoneConfig = tmp_int_zeroCorrelationZoneConfig
	var tmp_int_preambleReceivedTargetPower int64
	if tmp_int_preambleReceivedTargetPower, err = r.ReadInteger(&uper.Constraint{Lb: -202, Ub: -60}, false); err != nil {
		return utils.WrapError("ReadInteger preambleReceivedTargetPower", err)
	}
	ie.preambleReceivedTargetPower = tmp_int_preambleReceivedTargetPower
	if err = ie.preambleTransMax.Decode(r); err != nil {
		return utils.WrapError("Decode preambleTransMax", err)
	}
	if err = ie.powerRampingStep.Decode(r); err != nil {
		return utils.WrapError("Decode powerRampingStep", err)
	}
	if err = ie.ra_ResponseWindow.Decode(r); err != nil {
		return utils.WrapError("Decode ra_ResponseWindow", err)
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

			prach_ConfigurationPeriodScaling_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			prach_ConfigurationFrameOffset_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			prach_ConfigurationSOffset_IAB_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			ra_ResponseWindow_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			prach_ConfigurationIndex_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode prach_ConfigurationPeriodScaling_IAB_r16 optional
			if prach_ConfigurationPeriodScaling_IAB_r16Present {
				ie.prach_ConfigurationPeriodScaling_IAB_r16 = new(RACH_ConfigGeneric_prach_ConfigurationPeriodScaling_IAB_r16)
				if err = ie.prach_ConfigurationPeriodScaling_IAB_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode prach_ConfigurationPeriodScaling_IAB_r16", err)
				}
			}
			// decode prach_ConfigurationFrameOffset_IAB_r16 optional
			if prach_ConfigurationFrameOffset_IAB_r16Present {
				var tmp_int_prach_ConfigurationFrameOffset_IAB_r16 int64
				if tmp_int_prach_ConfigurationFrameOffset_IAB_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 63}, false); err != nil {
					return utils.WrapError("Decode prach_ConfigurationFrameOffset_IAB_r16", err)
				}
				ie.prach_ConfigurationFrameOffset_IAB_r16 = &tmp_int_prach_ConfigurationFrameOffset_IAB_r16
			}
			// decode prach_ConfigurationSOffset_IAB_r16 optional
			if prach_ConfigurationSOffset_IAB_r16Present {
				var tmp_int_prach_ConfigurationSOffset_IAB_r16 int64
				if tmp_int_prach_ConfigurationSOffset_IAB_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
					return utils.WrapError("Decode prach_ConfigurationSOffset_IAB_r16", err)
				}
				ie.prach_ConfigurationSOffset_IAB_r16 = &tmp_int_prach_ConfigurationSOffset_IAB_r16
			}
			// decode ra_ResponseWindow_v1610 optional
			if ra_ResponseWindow_v1610Present {
				ie.ra_ResponseWindow_v1610 = new(RACH_ConfigGeneric_ra_ResponseWindow_v1610)
				if err = ie.ra_ResponseWindow_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode ra_ResponseWindow_v1610", err)
				}
			}
			// decode prach_ConfigurationIndex_v1610 optional
			if prach_ConfigurationIndex_v1610Present {
				var tmp_int_prach_ConfigurationIndex_v1610 int64
				if tmp_int_prach_ConfigurationIndex_v1610, err = extReader.ReadInteger(&uper.Constraint{Lb: 256, Ub: 262}, false); err != nil {
					return utils.WrapError("Decode prach_ConfigurationIndex_v1610", err)
				}
				ie.prach_ConfigurationIndex_v1610 = &tmp_int_prach_ConfigurationIndex_v1610
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			ra_ResponseWindow_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode ra_ResponseWindow_v1700 optional
			if ra_ResponseWindow_v1700Present {
				ie.ra_ResponseWindow_v1700 = new(RACH_ConfigGeneric_ra_ResponseWindow_v1700)
				if err = ie.ra_ResponseWindow_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode ra_ResponseWindow_v1700", err)
				}
			}
		}
	}
	return nil
}
