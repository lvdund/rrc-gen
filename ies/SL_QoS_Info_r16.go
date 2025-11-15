package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_QoS_Info_r16 struct {
	sl_QoS_FlowIdentity_r16 SL_QoS_FlowIdentity_r16 `madatory`
	sl_QoS_Profile_r16      *SL_QoS_Profile_r16     `optional`
}

func (ie *SL_QoS_Info_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_QoS_Profile_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.sl_QoS_FlowIdentity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_QoS_FlowIdentity_r16", err)
	}
	if ie.sl_QoS_Profile_r16 != nil {
		if err = ie.sl_QoS_Profile_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_QoS_Profile_r16", err)
		}
	}
	return nil
}

func (ie *SL_QoS_Info_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_QoS_Profile_r16Present bool
	if sl_QoS_Profile_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.sl_QoS_FlowIdentity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_QoS_FlowIdentity_r16", err)
	}
	if sl_QoS_Profile_r16Present {
		ie.sl_QoS_Profile_r16 = new(SL_QoS_Profile_r16)
		if err = ie.sl_QoS_Profile_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_QoS_Profile_r16", err)
		}
	}
	return nil
}
