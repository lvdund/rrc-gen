package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BAP_Config_r16 struct {
	bap_Address_r16              *uper.BitString                             `lb:10,ub:10,optional`
	defaultUL_BAP_RoutingID_r16  *BAP_RoutingID_r16                          `optional`
	defaultUL_BH_RLC_Channel_r16 *BH_RLC_ChannelID_r16                       `optional`
	flowControlFeedbackType_r16  *BAP_Config_r16_flowControlFeedbackType_r16 `optional`
}

func (ie *BAP_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.bap_Address_r16 != nil, ie.defaultUL_BAP_RoutingID_r16 != nil, ie.defaultUL_BH_RLC_Channel_r16 != nil, ie.flowControlFeedbackType_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.bap_Address_r16 != nil {
		if err = w.WriteBitString(ie.bap_Address_r16.Bytes, uint(ie.bap_Address_r16.NumBits), &uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			return utils.WrapError("Encode bap_Address_r16", err)
		}
	}
	if ie.defaultUL_BAP_RoutingID_r16 != nil {
		if err = ie.defaultUL_BAP_RoutingID_r16.Encode(w); err != nil {
			return utils.WrapError("Encode defaultUL_BAP_RoutingID_r16", err)
		}
	}
	if ie.defaultUL_BH_RLC_Channel_r16 != nil {
		if err = ie.defaultUL_BH_RLC_Channel_r16.Encode(w); err != nil {
			return utils.WrapError("Encode defaultUL_BH_RLC_Channel_r16", err)
		}
	}
	if ie.flowControlFeedbackType_r16 != nil {
		if err = ie.flowControlFeedbackType_r16.Encode(w); err != nil {
			return utils.WrapError("Encode flowControlFeedbackType_r16", err)
		}
	}
	return nil
}

func (ie *BAP_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var bap_Address_r16Present bool
	if bap_Address_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var defaultUL_BAP_RoutingID_r16Present bool
	if defaultUL_BAP_RoutingID_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var defaultUL_BH_RLC_Channel_r16Present bool
	if defaultUL_BH_RLC_Channel_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var flowControlFeedbackType_r16Present bool
	if flowControlFeedbackType_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if bap_Address_r16Present {
		var tmp_bs_bap_Address_r16 []byte
		var n_bap_Address_r16 uint
		if tmp_bs_bap_Address_r16, n_bap_Address_r16, err = r.ReadBitString(&uper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			return utils.WrapError("Decode bap_Address_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_bap_Address_r16,
			NumBits: uint64(n_bap_Address_r16),
		}
		ie.bap_Address_r16 = &tmp_bitstring
	}
	if defaultUL_BAP_RoutingID_r16Present {
		ie.defaultUL_BAP_RoutingID_r16 = new(BAP_RoutingID_r16)
		if err = ie.defaultUL_BAP_RoutingID_r16.Decode(r); err != nil {
			return utils.WrapError("Decode defaultUL_BAP_RoutingID_r16", err)
		}
	}
	if defaultUL_BH_RLC_Channel_r16Present {
		ie.defaultUL_BH_RLC_Channel_r16 = new(BH_RLC_ChannelID_r16)
		if err = ie.defaultUL_BH_RLC_Channel_r16.Decode(r); err != nil {
			return utils.WrapError("Decode defaultUL_BH_RLC_Channel_r16", err)
		}
	}
	if flowControlFeedbackType_r16Present {
		ie.flowControlFeedbackType_r16 = new(BAP_Config_r16_flowControlFeedbackType_r16)
		if err = ie.flowControlFeedbackType_r16.Decode(r); err != nil {
			return utils.WrapError("Decode flowControlFeedbackType_r16", err)
		}
	}
	return nil
}
