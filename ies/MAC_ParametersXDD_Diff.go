package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MAC_ParametersXDD_Diff struct {
	skipUplinkTxDynamic                *MAC_ParametersXDD_Diff_skipUplinkTxDynamic                `optional`
	logicalChannelSR_DelayTimer        *MAC_ParametersXDD_Diff_logicalChannelSR_DelayTimer        `optional`
	longDRX_Cycle                      *MAC_ParametersXDD_Diff_longDRX_Cycle                      `optional`
	shortDRX_Cycle                     *MAC_ParametersXDD_Diff_shortDRX_Cycle                     `optional`
	multipleSR_Configurations          *MAC_ParametersXDD_Diff_multipleSR_Configurations          `optional`
	multipleConfiguredGrants           *MAC_ParametersXDD_Diff_multipleConfiguredGrants           `optional`
	secondaryDRX_Group_r16             *MAC_ParametersXDD_Diff_secondaryDRX_Group_r16             `optional,ext-1`
	enhancedSkipUplinkTxDynamic_r16    *MAC_ParametersXDD_Diff_enhancedSkipUplinkTxDynamic_r16    `optional,ext-2`
	enhancedSkipUplinkTxConfigured_r16 *MAC_ParametersXDD_Diff_enhancedSkipUplinkTxConfigured_r16 `optional,ext-2`
}

func (ie *MAC_ParametersXDD_Diff) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.secondaryDRX_Group_r16 != nil || ie.enhancedSkipUplinkTxDynamic_r16 != nil || ie.enhancedSkipUplinkTxConfigured_r16 != nil
	preambleBits := []bool{hasExtensions, ie.skipUplinkTxDynamic != nil, ie.logicalChannelSR_DelayTimer != nil, ie.longDRX_Cycle != nil, ie.shortDRX_Cycle != nil, ie.multipleSR_Configurations != nil, ie.multipleConfiguredGrants != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.skipUplinkTxDynamic != nil {
		if err = ie.skipUplinkTxDynamic.Encode(w); err != nil {
			return utils.WrapError("Encode skipUplinkTxDynamic", err)
		}
	}
	if ie.logicalChannelSR_DelayTimer != nil {
		if err = ie.logicalChannelSR_DelayTimer.Encode(w); err != nil {
			return utils.WrapError("Encode logicalChannelSR_DelayTimer", err)
		}
	}
	if ie.longDRX_Cycle != nil {
		if err = ie.longDRX_Cycle.Encode(w); err != nil {
			return utils.WrapError("Encode longDRX_Cycle", err)
		}
	}
	if ie.shortDRX_Cycle != nil {
		if err = ie.shortDRX_Cycle.Encode(w); err != nil {
			return utils.WrapError("Encode shortDRX_Cycle", err)
		}
	}
	if ie.multipleSR_Configurations != nil {
		if err = ie.multipleSR_Configurations.Encode(w); err != nil {
			return utils.WrapError("Encode multipleSR_Configurations", err)
		}
	}
	if ie.multipleConfiguredGrants != nil {
		if err = ie.multipleConfiguredGrants.Encode(w); err != nil {
			return utils.WrapError("Encode multipleConfiguredGrants", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.secondaryDRX_Group_r16 != nil, ie.enhancedSkipUplinkTxDynamic_r16 != nil || ie.enhancedSkipUplinkTxConfigured_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap MAC_ParametersXDD_Diff", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.secondaryDRX_Group_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode secondaryDRX_Group_r16 optional
			if ie.secondaryDRX_Group_r16 != nil {
				if err = ie.secondaryDRX_Group_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode secondaryDRX_Group_r16", err)
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
			optionals_ext_2 := []bool{ie.enhancedSkipUplinkTxDynamic_r16 != nil, ie.enhancedSkipUplinkTxConfigured_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode enhancedSkipUplinkTxDynamic_r16 optional
			if ie.enhancedSkipUplinkTxDynamic_r16 != nil {
				if err = ie.enhancedSkipUplinkTxDynamic_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enhancedSkipUplinkTxDynamic_r16", err)
				}
			}
			// encode enhancedSkipUplinkTxConfigured_r16 optional
			if ie.enhancedSkipUplinkTxConfigured_r16 != nil {
				if err = ie.enhancedSkipUplinkTxConfigured_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode enhancedSkipUplinkTxConfigured_r16", err)
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

func (ie *MAC_ParametersXDD_Diff) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var skipUplinkTxDynamicPresent bool
	if skipUplinkTxDynamicPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var logicalChannelSR_DelayTimerPresent bool
	if logicalChannelSR_DelayTimerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var longDRX_CyclePresent bool
	if longDRX_CyclePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var shortDRX_CyclePresent bool
	if shortDRX_CyclePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var multipleSR_ConfigurationsPresent bool
	if multipleSR_ConfigurationsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var multipleConfiguredGrantsPresent bool
	if multipleConfiguredGrantsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if skipUplinkTxDynamicPresent {
		ie.skipUplinkTxDynamic = new(MAC_ParametersXDD_Diff_skipUplinkTxDynamic)
		if err = ie.skipUplinkTxDynamic.Decode(r); err != nil {
			return utils.WrapError("Decode skipUplinkTxDynamic", err)
		}
	}
	if logicalChannelSR_DelayTimerPresent {
		ie.logicalChannelSR_DelayTimer = new(MAC_ParametersXDD_Diff_logicalChannelSR_DelayTimer)
		if err = ie.logicalChannelSR_DelayTimer.Decode(r); err != nil {
			return utils.WrapError("Decode logicalChannelSR_DelayTimer", err)
		}
	}
	if longDRX_CyclePresent {
		ie.longDRX_Cycle = new(MAC_ParametersXDD_Diff_longDRX_Cycle)
		if err = ie.longDRX_Cycle.Decode(r); err != nil {
			return utils.WrapError("Decode longDRX_Cycle", err)
		}
	}
	if shortDRX_CyclePresent {
		ie.shortDRX_Cycle = new(MAC_ParametersXDD_Diff_shortDRX_Cycle)
		if err = ie.shortDRX_Cycle.Decode(r); err != nil {
			return utils.WrapError("Decode shortDRX_Cycle", err)
		}
	}
	if multipleSR_ConfigurationsPresent {
		ie.multipleSR_Configurations = new(MAC_ParametersXDD_Diff_multipleSR_Configurations)
		if err = ie.multipleSR_Configurations.Decode(r); err != nil {
			return utils.WrapError("Decode multipleSR_Configurations", err)
		}
	}
	if multipleConfiguredGrantsPresent {
		ie.multipleConfiguredGrants = new(MAC_ParametersXDD_Diff_multipleConfiguredGrants)
		if err = ie.multipleConfiguredGrants.Decode(r); err != nil {
			return utils.WrapError("Decode multipleConfiguredGrants", err)
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

			secondaryDRX_Group_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode secondaryDRX_Group_r16 optional
			if secondaryDRX_Group_r16Present {
				ie.secondaryDRX_Group_r16 = new(MAC_ParametersXDD_Diff_secondaryDRX_Group_r16)
				if err = ie.secondaryDRX_Group_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode secondaryDRX_Group_r16", err)
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

			enhancedSkipUplinkTxDynamic_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			enhancedSkipUplinkTxConfigured_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode enhancedSkipUplinkTxDynamic_r16 optional
			if enhancedSkipUplinkTxDynamic_r16Present {
				ie.enhancedSkipUplinkTxDynamic_r16 = new(MAC_ParametersXDD_Diff_enhancedSkipUplinkTxDynamic_r16)
				if err = ie.enhancedSkipUplinkTxDynamic_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode enhancedSkipUplinkTxDynamic_r16", err)
				}
			}
			// decode enhancedSkipUplinkTxConfigured_r16 optional
			if enhancedSkipUplinkTxConfigured_r16Present {
				ie.enhancedSkipUplinkTxConfigured_r16 = new(MAC_ParametersXDD_Diff_enhancedSkipUplinkTxConfigured_r16)
				if err = ie.enhancedSkipUplinkTxConfigured_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode enhancedSkipUplinkTxConfigured_r16", err)
				}
			}
		}
	}
	return nil
}
