package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BAP_Parameters_r16 struct {
	flowControlBH_RLC_ChannelBased_r16 *BAP_Parameters_r16_flowControlBH_RLC_ChannelBased_r16 `optional`
	flowControlRouting_ID_Based_r16    *BAP_Parameters_r16_flowControlRouting_ID_Based_r16    `optional`
}

func (ie *BAP_Parameters_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.flowControlBH_RLC_ChannelBased_r16 != nil, ie.flowControlRouting_ID_Based_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.flowControlBH_RLC_ChannelBased_r16 != nil {
		if err = ie.flowControlBH_RLC_ChannelBased_r16.Encode(w); err != nil {
			return utils.WrapError("Encode flowControlBH_RLC_ChannelBased_r16", err)
		}
	}
	if ie.flowControlRouting_ID_Based_r16 != nil {
		if err = ie.flowControlRouting_ID_Based_r16.Encode(w); err != nil {
			return utils.WrapError("Encode flowControlRouting_ID_Based_r16", err)
		}
	}
	return nil
}

func (ie *BAP_Parameters_r16) Decode(r *uper.UperReader) error {
	var err error
	var flowControlBH_RLC_ChannelBased_r16Present bool
	if flowControlBH_RLC_ChannelBased_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var flowControlRouting_ID_Based_r16Present bool
	if flowControlRouting_ID_Based_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if flowControlBH_RLC_ChannelBased_r16Present {
		ie.flowControlBH_RLC_ChannelBased_r16 = new(BAP_Parameters_r16_flowControlBH_RLC_ChannelBased_r16)
		if err = ie.flowControlBH_RLC_ChannelBased_r16.Decode(r); err != nil {
			return utils.WrapError("Decode flowControlBH_RLC_ChannelBased_r16", err)
		}
	}
	if flowControlRouting_ID_Based_r16Present {
		ie.flowControlRouting_ID_Based_r16 = new(BAP_Parameters_r16_flowControlRouting_ID_Based_r16)
		if err = ie.flowControlRouting_ID_Based_r16.Decode(r); err != nil {
			return utils.WrapError("Decode flowControlRouting_ID_Based_r16", err)
		}
	}
	return nil
}
