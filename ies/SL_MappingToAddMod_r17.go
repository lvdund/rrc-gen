package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_MappingToAddMod_r17 struct {
	sl_RemoteUE_RB_Identity_r17 SL_RemoteUE_RB_Identity_r17 `madatory`
	sl_EgressRLC_ChannelUu_r17  *Uu_RelayRLC_ChannelID_r17  `optional`
	sl_EgressRLC_ChannelPC5_r17 *SL_RLC_ChannelID_r17       `optional`
}

func (ie *SL_MappingToAddMod_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_EgressRLC_ChannelUu_r17 != nil, ie.sl_EgressRLC_ChannelPC5_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.sl_RemoteUE_RB_Identity_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sl_RemoteUE_RB_Identity_r17", err)
	}
	if ie.sl_EgressRLC_ChannelUu_r17 != nil {
		if err = ie.sl_EgressRLC_ChannelUu_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_EgressRLC_ChannelUu_r17", err)
		}
	}
	if ie.sl_EgressRLC_ChannelPC5_r17 != nil {
		if err = ie.sl_EgressRLC_ChannelPC5_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sl_EgressRLC_ChannelPC5_r17", err)
		}
	}
	return nil
}

func (ie *SL_MappingToAddMod_r17) Decode(r *uper.UperReader) error {
	var err error
	var sl_EgressRLC_ChannelUu_r17Present bool
	if sl_EgressRLC_ChannelUu_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_EgressRLC_ChannelPC5_r17Present bool
	if sl_EgressRLC_ChannelPC5_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.sl_RemoteUE_RB_Identity_r17.Decode(r); err != nil {
		return utils.WrapError("Decode sl_RemoteUE_RB_Identity_r17", err)
	}
	if sl_EgressRLC_ChannelUu_r17Present {
		ie.sl_EgressRLC_ChannelUu_r17 = new(Uu_RelayRLC_ChannelID_r17)
		if err = ie.sl_EgressRLC_ChannelUu_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_EgressRLC_ChannelUu_r17", err)
		}
	}
	if sl_EgressRLC_ChannelPC5_r17Present {
		ie.sl_EgressRLC_ChannelPC5_r17 = new(SL_RLC_ChannelID_r17)
		if err = ie.sl_EgressRLC_ChannelPC5_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sl_EgressRLC_ChannelPC5_r17", err)
		}
	}
	return nil
}
