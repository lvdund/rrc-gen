package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PDCP_Config_r16 struct {
	sl_DiscardTimer_r16   *SL_PDCP_Config_r16_sl_DiscardTimer_r16   `optional`
	sl_PDCP_SN_Size_r16   *SL_PDCP_Config_r16_sl_PDCP_SN_Size_r16   `optional`
	sl_OutOfOrderDelivery *SL_PDCP_Config_r16_sl_OutOfOrderDelivery `optional`
}

func (ie *SL_PDCP_Config_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_DiscardTimer_r16 != nil, ie.sl_PDCP_SN_Size_r16 != nil, ie.sl_OutOfOrderDelivery != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_DiscardTimer_r16 != nil {
		if err = ie.sl_DiscardTimer_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_DiscardTimer_r16", err)
		}
	}
	if ie.sl_PDCP_SN_Size_r16 != nil {
		if err = ie.sl_PDCP_SN_Size_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_PDCP_SN_Size_r16", err)
		}
	}
	if ie.sl_OutOfOrderDelivery != nil {
		if err = ie.sl_OutOfOrderDelivery.Encode(w); err != nil {
			return utils.WrapError("Encode sl_OutOfOrderDelivery", err)
		}
	}
	return nil
}

func (ie *SL_PDCP_Config_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_DiscardTimer_r16Present bool
	if sl_DiscardTimer_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_PDCP_SN_Size_r16Present bool
	if sl_PDCP_SN_Size_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sl_OutOfOrderDeliveryPresent bool
	if sl_OutOfOrderDeliveryPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_DiscardTimer_r16Present {
		ie.sl_DiscardTimer_r16 = new(SL_PDCP_Config_r16_sl_DiscardTimer_r16)
		if err = ie.sl_DiscardTimer_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_DiscardTimer_r16", err)
		}
	}
	if sl_PDCP_SN_Size_r16Present {
		ie.sl_PDCP_SN_Size_r16 = new(SL_PDCP_Config_r16_sl_PDCP_SN_Size_r16)
		if err = ie.sl_PDCP_SN_Size_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_PDCP_SN_Size_r16", err)
		}
	}
	if sl_OutOfOrderDeliveryPresent {
		ie.sl_OutOfOrderDelivery = new(SL_PDCP_Config_r16_sl_OutOfOrderDelivery)
		if err = ie.sl_OutOfOrderDelivery.Decode(r); err != nil {
			return utils.WrapError("Decode sl_OutOfOrderDelivery", err)
		}
	}
	return nil
}
