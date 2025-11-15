package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RLC_BearerConfig struct {
	logicalChannelIdentity        LogicalChannelIdentity              `madatory`
	servedRadioBearer             *RLC_BearerConfig_servedRadioBearer `optional`
	reestablishRLC                *RLC_BearerConfig_reestablishRLC    `optional`
	rlc_Config                    *RLC_Config                         `optional`
	mac_LogicalChannelConfig      *LogicalChannelConfig               `optional`
	rlc_Config_v1610              *RLC_Config_v1610                   `optional,ext-1`
	rlc_Config_v1700              *RLC_Config_v1700                   `optional,ext-2`
	logicalChannelIdentityExt_r17 *LogicalChannelIdentityExt_r17      `optional,ext-2`
	multicastRLC_BearerConfig_r17 *MulticastRLC_BearerConfig_r17      `optional,ext-2`
	servedRadioBearerSRB4_r17     *SRB_Identity_v1700                 `optional,ext-2`
}

func (ie *RLC_BearerConfig) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.rlc_Config_v1610 != nil || ie.rlc_Config_v1700 != nil || ie.logicalChannelIdentityExt_r17 != nil || ie.multicastRLC_BearerConfig_r17 != nil || ie.servedRadioBearerSRB4_r17 != nil
	preambleBits := []bool{hasExtensions, ie.servedRadioBearer != nil, ie.reestablishRLC != nil, ie.rlc_Config != nil, ie.mac_LogicalChannelConfig != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.logicalChannelIdentity.Encode(w); err != nil {
		return utils.WrapError("Encode logicalChannelIdentity", err)
	}
	if ie.servedRadioBearer != nil {
		if err = ie.servedRadioBearer.Encode(w); err != nil {
			return utils.WrapError("Encode servedRadioBearer", err)
		}
	}
	if ie.reestablishRLC != nil {
		if err = ie.reestablishRLC.Encode(w); err != nil {
			return utils.WrapError("Encode reestablishRLC", err)
		}
	}
	if ie.rlc_Config != nil {
		if err = ie.rlc_Config.Encode(w); err != nil {
			return utils.WrapError("Encode rlc_Config", err)
		}
	}
	if ie.mac_LogicalChannelConfig != nil {
		if err = ie.mac_LogicalChannelConfig.Encode(w); err != nil {
			return utils.WrapError("Encode mac_LogicalChannelConfig", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.rlc_Config_v1610 != nil, ie.rlc_Config_v1700 != nil || ie.logicalChannelIdentityExt_r17 != nil || ie.multicastRLC_BearerConfig_r17 != nil || ie.servedRadioBearerSRB4_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap RLC_BearerConfig", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.rlc_Config_v1610 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode rlc_Config_v1610 optional
			if ie.rlc_Config_v1610 != nil {
				if err = ie.rlc_Config_v1610.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rlc_Config_v1610", err)
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
			optionals_ext_2 := []bool{ie.rlc_Config_v1700 != nil, ie.logicalChannelIdentityExt_r17 != nil, ie.multicastRLC_BearerConfig_r17 != nil, ie.servedRadioBearerSRB4_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode rlc_Config_v1700 optional
			if ie.rlc_Config_v1700 != nil {
				if err = ie.rlc_Config_v1700.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rlc_Config_v1700", err)
				}
			}
			// encode logicalChannelIdentityExt_r17 optional
			if ie.logicalChannelIdentityExt_r17 != nil {
				if err = ie.logicalChannelIdentityExt_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode logicalChannelIdentityExt_r17", err)
				}
			}
			// encode multicastRLC_BearerConfig_r17 optional
			if ie.multicastRLC_BearerConfig_r17 != nil {
				if err = ie.multicastRLC_BearerConfig_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode multicastRLC_BearerConfig_r17", err)
				}
			}
			// encode servedRadioBearerSRB4_r17 optional
			if ie.servedRadioBearerSRB4_r17 != nil {
				if err = ie.servedRadioBearerSRB4_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode servedRadioBearerSRB4_r17", err)
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

func (ie *RLC_BearerConfig) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var servedRadioBearerPresent bool
	if servedRadioBearerPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var reestablishRLCPresent bool
	if reestablishRLCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rlc_ConfigPresent bool
	if rlc_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var mac_LogicalChannelConfigPresent bool
	if mac_LogicalChannelConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.logicalChannelIdentity.Decode(r); err != nil {
		return utils.WrapError("Decode logicalChannelIdentity", err)
	}
	if servedRadioBearerPresent {
		ie.servedRadioBearer = new(RLC_BearerConfig_servedRadioBearer)
		if err = ie.servedRadioBearer.Decode(r); err != nil {
			return utils.WrapError("Decode servedRadioBearer", err)
		}
	}
	if reestablishRLCPresent {
		ie.reestablishRLC = new(RLC_BearerConfig_reestablishRLC)
		if err = ie.reestablishRLC.Decode(r); err != nil {
			return utils.WrapError("Decode reestablishRLC", err)
		}
	}
	if rlc_ConfigPresent {
		ie.rlc_Config = new(RLC_Config)
		if err = ie.rlc_Config.Decode(r); err != nil {
			return utils.WrapError("Decode rlc_Config", err)
		}
	}
	if mac_LogicalChannelConfigPresent {
		ie.mac_LogicalChannelConfig = new(LogicalChannelConfig)
		if err = ie.mac_LogicalChannelConfig.Decode(r); err != nil {
			return utils.WrapError("Decode mac_LogicalChannelConfig", err)
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

			rlc_Config_v1610Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode rlc_Config_v1610 optional
			if rlc_Config_v1610Present {
				ie.rlc_Config_v1610 = new(RLC_Config_v1610)
				if err = ie.rlc_Config_v1610.Decode(extReader); err != nil {
					return utils.WrapError("Decode rlc_Config_v1610", err)
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

			rlc_Config_v1700Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			logicalChannelIdentityExt_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			multicastRLC_BearerConfig_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			servedRadioBearerSRB4_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode rlc_Config_v1700 optional
			if rlc_Config_v1700Present {
				ie.rlc_Config_v1700 = new(RLC_Config_v1700)
				if err = ie.rlc_Config_v1700.Decode(extReader); err != nil {
					return utils.WrapError("Decode rlc_Config_v1700", err)
				}
			}
			// decode logicalChannelIdentityExt_r17 optional
			if logicalChannelIdentityExt_r17Present {
				ie.logicalChannelIdentityExt_r17 = new(LogicalChannelIdentityExt_r17)
				if err = ie.logicalChannelIdentityExt_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode logicalChannelIdentityExt_r17", err)
				}
			}
			// decode multicastRLC_BearerConfig_r17 optional
			if multicastRLC_BearerConfig_r17Present {
				ie.multicastRLC_BearerConfig_r17 = new(MulticastRLC_BearerConfig_r17)
				if err = ie.multicastRLC_BearerConfig_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode multicastRLC_BearerConfig_r17", err)
				}
			}
			// decode servedRadioBearerSRB4_r17 optional
			if servedRadioBearerSRB4_r17Present {
				ie.servedRadioBearerSRB4_r17 = new(SRB_Identity_v1700)
				if err = ie.servedRadioBearerSRB4_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode servedRadioBearerSRB4_r17", err)
				}
			}
		}
	}
	return nil
}
