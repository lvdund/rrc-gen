package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Uu_RelayRLC_ChannelConfig_r17 struct {
	uu_LogicalChannelIdentity_r17 *LogicalChannelIdentity                           `optional`
	uu_RelayRLC_ChannelID_r17     Uu_RelayRLC_ChannelID_r17                         `madatory`
	reestablishRLC_r17            *Uu_RelayRLC_ChannelConfig_r17_reestablishRLC_r17 `optional`
	rlc_Config_r17                *RLC_Config                                       `optional`
	mac_LogicalChannelConfig_r17  *LogicalChannelConfig                             `optional`
}

func (ie *Uu_RelayRLC_ChannelConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.uu_LogicalChannelIdentity_r17 != nil, ie.reestablishRLC_r17 != nil, ie.rlc_Config_r17 != nil, ie.mac_LogicalChannelConfig_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.uu_LogicalChannelIdentity_r17 != nil {
		if err = ie.uu_LogicalChannelIdentity_r17.Encode(w); err != nil {
			return utils.WrapError("Encode uu_LogicalChannelIdentity_r17", err)
		}
	}
	if err = ie.uu_RelayRLC_ChannelID_r17.Encode(w); err != nil {
		return utils.WrapError("Encode uu_RelayRLC_ChannelID_r17", err)
	}
	if ie.reestablishRLC_r17 != nil {
		if err = ie.reestablishRLC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode reestablishRLC_r17", err)
		}
	}
	if ie.rlc_Config_r17 != nil {
		if err = ie.rlc_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode rlc_Config_r17", err)
		}
	}
	if ie.mac_LogicalChannelConfig_r17 != nil {
		if err = ie.mac_LogicalChannelConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode mac_LogicalChannelConfig_r17", err)
		}
	}
	return nil
}

func (ie *Uu_RelayRLC_ChannelConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var uu_LogicalChannelIdentity_r17Present bool
	if uu_LogicalChannelIdentity_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var reestablishRLC_r17Present bool
	if reestablishRLC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rlc_Config_r17Present bool
	if rlc_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mac_LogicalChannelConfig_r17Present bool
	if mac_LogicalChannelConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if uu_LogicalChannelIdentity_r17Present {
		ie.uu_LogicalChannelIdentity_r17 = new(LogicalChannelIdentity)
		if err = ie.uu_LogicalChannelIdentity_r17.Decode(r); err != nil {
			return utils.WrapError("Decode uu_LogicalChannelIdentity_r17", err)
		}
	}
	if err = ie.uu_RelayRLC_ChannelID_r17.Decode(r); err != nil {
		return utils.WrapError("Decode uu_RelayRLC_ChannelID_r17", err)
	}
	if reestablishRLC_r17Present {
		ie.reestablishRLC_r17 = new(Uu_RelayRLC_ChannelConfig_r17_reestablishRLC_r17)
		if err = ie.reestablishRLC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode reestablishRLC_r17", err)
		}
	}
	if rlc_Config_r17Present {
		ie.rlc_Config_r17 = new(RLC_Config)
		if err = ie.rlc_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode rlc_Config_r17", err)
		}
	}
	if mac_LogicalChannelConfig_r17Present {
		ie.mac_LogicalChannelConfig_r17 = new(LogicalChannelConfig)
		if err = ie.mac_LogicalChannelConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode mac_LogicalChannelConfig_r17", err)
		}
	}
	return nil
}
