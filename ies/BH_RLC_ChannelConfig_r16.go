package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BH_RLC_ChannelConfig_r16 struct {
	bh_LogicalChannelIdentity_r16 *BH_LogicalChannelIdentity_r16               `optional`
	bh_RLC_ChannelID_r16          BH_RLC_ChannelID_r16                         `madatory`
	reestablishRLC_r16            *BH_RLC_ChannelConfig_r16_reestablishRLC_r16 `optional`
	rlc_Config_r16                *RLC_Config                                  `optional`
	mac_LogicalChannelConfig_r16  *LogicalChannelConfig                        `optional`
}

func (ie *BH_RLC_ChannelConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.bh_LogicalChannelIdentity_r16 != nil, ie.reestablishRLC_r16 != nil, ie.rlc_Config_r16 != nil, ie.mac_LogicalChannelConfig_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.bh_LogicalChannelIdentity_r16 != nil {
		if err = ie.bh_LogicalChannelIdentity_r16.Encode(w); err != nil {
			return utils.WrapError("Encode bh_LogicalChannelIdentity_r16", err)
		}
	}
	if err = ie.bh_RLC_ChannelID_r16.Encode(w); err != nil {
		return utils.WrapError("Encode bh_RLC_ChannelID_r16", err)
	}
	if ie.reestablishRLC_r16 != nil {
		if err = ie.reestablishRLC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode reestablishRLC_r16", err)
		}
	}
	if ie.rlc_Config_r16 != nil {
		if err = ie.rlc_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rlc_Config_r16", err)
		}
	}
	if ie.mac_LogicalChannelConfig_r16 != nil {
		if err = ie.mac_LogicalChannelConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mac_LogicalChannelConfig_r16", err)
		}
	}
	return nil
}

func (ie *BH_RLC_ChannelConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var bh_LogicalChannelIdentity_r16Present bool
	if bh_LogicalChannelIdentity_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var reestablishRLC_r16Present bool
	if reestablishRLC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rlc_Config_r16Present bool
	if rlc_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mac_LogicalChannelConfig_r16Present bool
	if mac_LogicalChannelConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if bh_LogicalChannelIdentity_r16Present {
		ie.bh_LogicalChannelIdentity_r16 = new(BH_LogicalChannelIdentity_r16)
		if err = ie.bh_LogicalChannelIdentity_r16.Decode(r); err != nil {
			return utils.WrapError("Decode bh_LogicalChannelIdentity_r16", err)
		}
	}
	if err = ie.bh_RLC_ChannelID_r16.Decode(r); err != nil {
		return utils.WrapError("Decode bh_RLC_ChannelID_r16", err)
	}
	if reestablishRLC_r16Present {
		ie.reestablishRLC_r16 = new(BH_RLC_ChannelConfig_r16_reestablishRLC_r16)
		if err = ie.reestablishRLC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode reestablishRLC_r16", err)
		}
	}
	if rlc_Config_r16Present {
		ie.rlc_Config_r16 = new(RLC_Config)
		if err = ie.rlc_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rlc_Config_r16", err)
		}
	}
	if mac_LogicalChannelConfig_r16Present {
		ie.mac_LogicalChannelConfig_r16 = new(LogicalChannelConfig)
		if err = ie.mac_LogicalChannelConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mac_LogicalChannelConfig_r16", err)
		}
	}
	return nil
}
