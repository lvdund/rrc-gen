package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRB_RLC_ConfigBroadcast_r17 struct {
	logicalChannelIdentity_r17 LogicalChannelIdentity                          `madatory`
	sn_FieldLength_r17         *MRB_RLC_ConfigBroadcast_r17_sn_FieldLength_r17 `optional`
	t_Reassembly_r17           *T_Reassembly                                   `optional`
}

func (ie *MRB_RLC_ConfigBroadcast_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sn_FieldLength_r17 != nil, ie.t_Reassembly_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.logicalChannelIdentity_r17.Encode(w); err != nil {
		return utils.WrapError("Encode logicalChannelIdentity_r17", err)
	}
	if ie.sn_FieldLength_r17 != nil {
		if err = ie.sn_FieldLength_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sn_FieldLength_r17", err)
		}
	}
	if ie.t_Reassembly_r17 != nil {
		if err = ie.t_Reassembly_r17.Encode(w); err != nil {
			return utils.WrapError("Encode t_Reassembly_r17", err)
		}
	}
	return nil
}

func (ie *MRB_RLC_ConfigBroadcast_r17) Decode(r *uper.UperReader) error {
	var err error
	var sn_FieldLength_r17Present bool
	if sn_FieldLength_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var t_Reassembly_r17Present bool
	if t_Reassembly_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.logicalChannelIdentity_r17.Decode(r); err != nil {
		return utils.WrapError("Decode logicalChannelIdentity_r17", err)
	}
	if sn_FieldLength_r17Present {
		ie.sn_FieldLength_r17 = new(MRB_RLC_ConfigBroadcast_r17_sn_FieldLength_r17)
		if err = ie.sn_FieldLength_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sn_FieldLength_r17", err)
		}
	}
	if t_Reassembly_r17Present {
		ie.t_Reassembly_r17 = new(T_Reassembly)
		if err = ie.t_Reassembly_r17.Decode(r); err != nil {
			return utils.WrapError("Decode t_Reassembly_r17", err)
		}
	}
	return nil
}
