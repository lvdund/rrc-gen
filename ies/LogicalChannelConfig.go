package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LogicalChannelConfig struct {
	ul_SpecificParameters     *LogicalChannelConfig_ul_SpecificParameters `lb:1,ub:maxNrofServingCells_1,optional`
	channelAccessPriority_r16 *int64                                      `lb:1,ub:4,optional,ext-3`
	bitRateMultiplier_r16     *LogicalChannelConfig_bitRateMultiplier_r16 `optional,ext-3`
}

func (ie *LogicalChannelConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.channelAccessPriority_r16 != nil || ie.bitRateMultiplier_r16 != nil
	preambleBits := []bool{hasExtensions, ie.ul_SpecificParameters != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ul_SpecificParameters != nil {
		if err = ie.ul_SpecificParameters.Encode(w); err != nil {
			return utils.WrapError("Encode ul_SpecificParameters", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.channelAccessPriority_r16 != nil || ie.bitRateMultiplier_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap LogicalChannelConfig", err)
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{ie.channelAccessPriority_r16 != nil, ie.bitRateMultiplier_r16 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode channelAccessPriority_r16 optional
			if ie.channelAccessPriority_r16 != nil {
				if err = extWriter.WriteInteger(*ie.channelAccessPriority_r16, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
					return utils.WrapError("Encode channelAccessPriority_r16", err)
				}
			}
			// encode bitRateMultiplier_r16 optional
			if ie.bitRateMultiplier_r16 != nil {
				if err = ie.bitRateMultiplier_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode bitRateMultiplier_r16", err)
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

func (ie *LogicalChannelConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_SpecificParametersPresent bool
	if ul_SpecificParametersPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ul_SpecificParametersPresent {
		ie.ul_SpecificParameters = new(LogicalChannelConfig_ul_SpecificParameters)
		if err = ie.ul_SpecificParameters.Decode(r); err != nil {
			return utils.WrapError("Decode ul_SpecificParameters", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			channelAccessPriority_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			bitRateMultiplier_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode channelAccessPriority_r16 optional
			if channelAccessPriority_r16Present {
				var tmp_int_channelAccessPriority_r16 int64
				if tmp_int_channelAccessPriority_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
					return utils.WrapError("Decode channelAccessPriority_r16", err)
				}
				ie.channelAccessPriority_r16 = &tmp_int_channelAccessPriority_r16
			}
			// decode bitRateMultiplier_r16 optional
			if bitRateMultiplier_r16Present {
				ie.bitRateMultiplier_r16 = new(LogicalChannelConfig_bitRateMultiplier_r16)
				if err = ie.bitRateMultiplier_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode bitRateMultiplier_r16", err)
				}
			}
		}
	}
	return nil
}
