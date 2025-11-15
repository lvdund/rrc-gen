package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_RLC_ChannelConfig_r17 struct {
	sl_RLC_ChannelID_r17            SL_RLC_ChannelID_r17         `madatory`
	sl_RLC_Config_r17               *SL_RLC_Config_r16           `optional`
	sl_MAC_LogicalChannelConfig_r17 *SL_LogicalChannelConfig_r16 `optional`
	sl_PacketDelayBudget_r17        *int64                       `lb:0,ub:1023,optional`
}

func (ie *SL_RLC_ChannelConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_RLC_Config_r17 != nil, ie.sl_MAC_LogicalChannelConfig_r17 != nil, ie.sl_PacketDelayBudget_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.sl_RLC_ChannelID_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sl_RLC_ChannelID_r17", err)
	}
	if ie.sl_RLC_Config_r17 != nil {
		if err = ie.sl_RLC_Config_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_RLC_Config_r17", err)
		}
	}
	if ie.sl_MAC_LogicalChannelConfig_r17 != nil {
		if err = ie.sl_MAC_LogicalChannelConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_MAC_LogicalChannelConfig_r17", err)
		}
	}
	if ie.sl_PacketDelayBudget_r17 != nil {
		if err = w.WriteInteger(*ie.sl_PacketDelayBudget_r17, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode sl_PacketDelayBudget_r17", err)
		}
	}
	return nil
}

func (ie *SL_RLC_ChannelConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var sl_RLC_Config_r17Present bool
	if sl_RLC_Config_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_MAC_LogicalChannelConfig_r17Present bool
	if sl_MAC_LogicalChannelConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PacketDelayBudget_r17Present bool
	if sl_PacketDelayBudget_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.sl_RLC_ChannelID_r17.Decode(r); err != nil {
		return utils.WrapError("Decode sl_RLC_ChannelID_r17", err)
	}
	if sl_RLC_Config_r17Present {
		ie.sl_RLC_Config_r17 = new(SL_RLC_Config_r16)
		if err = ie.sl_RLC_Config_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_RLC_Config_r17", err)
		}
	}
	if sl_MAC_LogicalChannelConfig_r17Present {
		ie.sl_MAC_LogicalChannelConfig_r17 = new(SL_LogicalChannelConfig_r16)
		if err = ie.sl_MAC_LogicalChannelConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_MAC_LogicalChannelConfig_r17", err)
		}
	}
	if sl_PacketDelayBudget_r17Present {
		var tmp_int_sl_PacketDelayBudget_r17 int64
		if tmp_int_sl_PacketDelayBudget_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode sl_PacketDelayBudget_r17", err)
		}
		ie.sl_PacketDelayBudget_r17 = &tmp_int_sl_PacketDelayBudget_r17
	}
	return nil
}
