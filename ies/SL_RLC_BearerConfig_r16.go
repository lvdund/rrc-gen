package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RLC_BearerConfig_r16 struct {
	sl_RLC_BearerConfigIndex_r16    SL_RLC_BearerConfigIndex_r16 `madatory`
	sl_ServedRadioBearer_r16        *SLRB_Uu_ConfigIndex_r16     `optional`
	sl_RLC_Config_r16               *SL_RLC_Config_r16           `optional`
	sl_MAC_LogicalChannelConfig_r16 *SL_LogicalChannelConfig_r16 `optional`
}

func (ie *SL_RLC_BearerConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_ServedRadioBearer_r16 != nil, ie.sl_RLC_Config_r16 != nil, ie.sl_MAC_LogicalChannelConfig_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.sl_RLC_BearerConfigIndex_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_RLC_BearerConfigIndex_r16", err)
	}
	if ie.sl_ServedRadioBearer_r16 != nil {
		if err = ie.sl_ServedRadioBearer_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_ServedRadioBearer_r16", err)
		}
	}
	if ie.sl_RLC_Config_r16 != nil {
		if err = ie.sl_RLC_Config_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RLC_Config_r16", err)
		}
	}
	if ie.sl_MAC_LogicalChannelConfig_r16 != nil {
		if err = ie.sl_MAC_LogicalChannelConfig_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MAC_LogicalChannelConfig_r16", err)
		}
	}
	return nil
}

func (ie *SL_RLC_BearerConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_ServedRadioBearer_r16Present bool
	if sl_ServedRadioBearer_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_RLC_Config_r16Present bool
	if sl_RLC_Config_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MAC_LogicalChannelConfig_r16Present bool
	if sl_MAC_LogicalChannelConfig_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.sl_RLC_BearerConfigIndex_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_RLC_BearerConfigIndex_r16", err)
	}
	if sl_ServedRadioBearer_r16Present {
		ie.sl_ServedRadioBearer_r16 = new(SLRB_Uu_ConfigIndex_r16)
		if err = ie.sl_ServedRadioBearer_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_ServedRadioBearer_r16", err)
		}
	}
	if sl_RLC_Config_r16Present {
		ie.sl_RLC_Config_r16 = new(SL_RLC_Config_r16)
		if err = ie.sl_RLC_Config_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_RLC_Config_r16", err)
		}
	}
	if sl_MAC_LogicalChannelConfig_r16Present {
		ie.sl_MAC_LogicalChannelConfig_r16 = new(SL_LogicalChannelConfig_r16)
		if err = ie.sl_MAC_LogicalChannelConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MAC_LogicalChannelConfig_r16", err)
		}
	}
	return nil
}
